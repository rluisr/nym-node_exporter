package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	mu            sync.RWMutex
	latestMetrics = make(map[string]interface{})
	cache = make(map[string]CacheEntry)
)

type CacheEntry struct {
	data        map[string]interface{}
	lastUpdated time.Time
}

type JSONCollector struct{}

func (c *JSONCollector) Describe(ch chan<- *prometheus.Desc) {}

func (c *JSONCollector) Collect(ch chan<- prometheus.Metric) {
	mu.RLock()
	defer mu.RUnlock()
	for key, value := range latestMetrics {
		if key == "last_probe_log" {
			continue
		}
		// メトリクス名として "nym_" プレフィックスを付与し、Prometheus仕様に合わせサニタイズ
		metricName := sanitize("nym_" + key)
		// 数値または boolean の場合は Gauge として出力
		if num, ok := toFloat64(value); ok {
			desc := prometheus.NewDesc(metricName, "Metric for "+key, nil, nil)
			if m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, num); err == nil {
				ch <- m
			}
		} else if b, ok := value.(bool); ok {
			val := 0.0
			if b {
				val = 1.0
			}
			desc := prometheus.NewDesc(metricName, "Metric for "+key, nil, nil)
			if m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, val); err == nil {
				ch <- m
			}
		} else if s, ok := value.(string); ok {
			// 文字列の場合、info メトリクスとしてラベル "value" にその文字列を入れ固定値 1 として出力
			desc := prometheus.NewDesc(metricName+"_info", "Info metric for "+key, []string{"value"}, nil)
			if m, err := prometheus.NewConstMetric(desc, prometheus.GaugeValue, 1, s); err == nil {
				ch <- m
			}
		}
		// 他の型は無視
	}
}

func toFloat64(value interface{}) (float64, bool) {
	switch v := value.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	default:
		return 0, false
	}
}

func sanitize(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	return re.ReplaceAllString(s, "_")
}

func flatten(prefix string, data interface{}, out map[string]interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			newPrefix := key
			if prefix != "" {
				newPrefix = prefix + "_" + key
			}
			flatten(newPrefix, value, out)
		}
	case []interface{}:
		// 配列の場合、各要素を文字列化してカンマで連結
		var parts []string
		for _, item := range v {
			switch item.(type) {
			case map[string]interface{}, []interface{}:
				b, err := json.Marshal(item)
				if err == nil {
					parts = append(parts, string(b))
				}
			default:
				parts = append(parts, fmt.Sprintf("%v", item))
			}
		}
		out[prefix] = strings.Join(parts, ",")
	default:
		out[prefix] = v
	}
}

func fetchData(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	out := make(map[string]interface{})
	flatten("", data, out)
	return out, nil
}

// /metrics Handler
func metricsHandler(w http.ResponseWriter, r *http.Request) {
	identityKey := r.URL.Query().Get("identity_key")
	if identityKey == "" {
		http.Error(w, "Missing identity_key parameter", http.StatusBadRequest)
		return
	}
	url := "https://harbourmaster.nymtech.net/v2/gateways/" + identityKey

	// Check cache (fetch again if more than 10 minutes have passed)
	mu.Lock()
	entry, ok := cache[identityKey]
	needFetch := !ok || time.Since(entry.lastUpdated) > 10*time.Minute
	mu.Unlock()

	if needFetch {
		newData, err := fetchData(url)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to fetch data: %v", err), http.StatusInternalServerError)
			return
		}
		mu.Lock()
		cache[identityKey] = CacheEntry{
			data:        newData,
			lastUpdated: time.Now(),
		}
		latestMetrics = newData
		mu.Unlock()
	} else {
		mu.Lock()
		latestMetrics = entry.data
		mu.Unlock()
	}

	promhttp.Handler().ServeHTTP(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`
	<html>
		<body>
			<h1>Nym Node Exporter</h1>
			<p>
				<a href="https://github.com/rluisr/nym-node_exporter">GitHub</a>
			</p>
		</body>
	</html>
	`))
}

func main() {
	portPtr := flag.String("port", "9998", "Port to serve metrics on")
	flag.Parse()

	collector := &JSONCollector{}
	prometheus.MustRegister(collector)

	http.Handle("/", http.HandlerFunc(indexHandler))
	http.Handle("/metrics", http.HandlerFunc(metricsHandler))
	log.Println("Prometheus exporter is running on :" + *portPtr + "/metrics")
	log.Fatal(http.ListenAndServe(":"+*portPtr, nil))
}