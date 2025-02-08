# Nym Gateway Metrics Prometheus Exporter

This program is a Prometheus exporter written in Go that retrieves JSON data from the Harbourmaster API (https://harbourmaster.nymtech.net/v2/gateways/) for a Nym gateway and exposes it as Prometheus metrics. The JSON data is recursively flattened, and each field is converted into a corresponding Prometheus metric.

**Important Notes:**
- **Gateway Only:** Currently, this exporter supports **gateway** data only.
- **Update Interval:** Data is refreshed at **10-minute** intervals.
- **Harbourmaster Dependency:** This exporter is dependent on the Harbourmaster API. Any changes to the API may require corresponding updates to the code.

## Features

- **JSON Retrieval and Flattening**  
  For a given `identity_key`, the exporter fetches JSON data from the Harbourmaster API and recursively flattens its structure.

- **Prometheus Metrics Conversion**  
  - Numerical fields and booleans are exported as gauge metrics (with booleans converted to 1 for `true` and 0 for `false`).  
  - String fields are exported as info metrics (with a `_info` suffix) using a label named `value` to hold the original string.  
  - Metrics corresponding to the `last_probe_log` field are intentionally ignored.

- **Caching Mechanism**  
  The exporter caches the retrieved data for each `identity_key` for 10 minutes to minimize unnecessary API requests.

- **Endpoint**  
  Metrics are available at the `/metrics` endpoint. It expects an `identity_key` query parameter in the URL.  
  Example: `/metrics?identity_key=<your_identity_key>`

- **Port Configuration**  
  Use the `--port` command-line flag to specify the port on which the server listens (default is `9998`).


## Installation

1. Clone the repository or place the source files (e.g., `main.go`) in your project directory.
2. Execute the following command to download the necessary modules:

   ```sh
   go mod tidy
   ```

## Usage

1. Start the exporter:

   ```sh
   ./nym_exporter --port=9998
   ```

   The `--port` flag is optional; if not specified, the default port `9998` will be used.

2. Access the metrics endpoint by opening your browser or via cURL:

   ```
   http://localhost:9998/metrics?identity_key=<your_identity_key>
   ```

   Example:

   ```
   http://localhost:9998/metrics?identity_key=28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND
   ```

## Metrics Output

```
# HELP nym_blacklisted Metric for blacklisted
# TYPE nym_blacklisted gauge
nym_blacklisted 0
# HELP nym_bonded Metric for bonded
# TYPE nym_bonded gauge
nym_bonded 1
# HELP nym_config_score Metric for config_score
# TYPE nym_config_score gauge
nym_config_score 0
# HELP nym_description_details_info Info metric for description_details
# TYPE nym_description_details_info gauge
nym_description_details_info{value="This node is operated by HCloud, a Japanese limited liability company, in support of NYM's vision. 1Gbpsdedicated network."} 1
# HELP nym_description_moniker_info Info metric for description_moniker
# TYPE nym_description_moniker_info gauge
nym_description_moniker_info{value="HCloud Ltd | Gateway | JP"} 1
# HELP nym_description_security_contact_info Info metric for description_security_contact
# TYPE nym_description_security_contact_info gauge
nym_description_security_contact_info{value="contact@hcloud.ltd"} 1
# HELP nym_description_website_info Info metric for description_website
# TYPE nym_description_website_info gauge
nym_description_website_info{value="https://hcloud.ltd"} 1
# HELP nym_explorer_pretty_bond_block_height Metric for explorer_pretty_bond_block_height
# TYPE nym_explorer_pretty_bond_block_height gauge
nym_explorer_pretty_bond_block_height 1.6664752e+07
# HELP nym_explorer_pretty_bond_gateway_clients_port Metric for explorer_pretty_bond_gateway_clients_port
# TYPE nym_explorer_pretty_bond_gateway_clients_port gauge
nym_explorer_pretty_bond_gateway_clients_port 9000
# HELP nym_explorer_pretty_bond_gateway_host_info Info metric for explorer_pretty_bond_gateway_host
# TYPE nym_explorer_pretty_bond_gateway_host_info gauge
nym_explorer_pretty_bond_gateway_host_info{value="217.178.53.49"} 1
# HELP nym_explorer_pretty_bond_gateway_identity_key_info Info metric for explorer_pretty_bond_gateway_identity_key
# TYPE nym_explorer_pretty_bond_gateway_identity_key_info gauge
nym_explorer_pretty_bond_gateway_identity_key_info{value="28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_explorer_pretty_bond_gateway_location_info Info metric for explorer_pretty_bond_gateway_location
# TYPE nym_explorer_pretty_bond_gateway_location_info gauge
nym_explorer_pretty_bond_gateway_location_info{value="Japan"} 1
# HELP nym_explorer_pretty_bond_gateway_mix_port Metric for explorer_pretty_bond_gateway_mix_port
# TYPE nym_explorer_pretty_bond_gateway_mix_port gauge
nym_explorer_pretty_bond_gateway_mix_port 1789
# HELP nym_explorer_pretty_bond_gateway_sphinx_key_info Info metric for explorer_pretty_bond_gateway_sphinx_key
# TYPE nym_explorer_pretty_bond_gateway_sphinx_key_info gauge
nym_explorer_pretty_bond_gateway_sphinx_key_info{value="7YauXgjKA9qoCmU5somVxTRy7XyQtCSc8yh29Ueitv3J"} 1
# HELP nym_explorer_pretty_bond_gateway_version_info Info metric for explorer_pretty_bond_gateway_version
# TYPE nym_explorer_pretty_bond_gateway_version_info gauge
nym_explorer_pretty_bond_gateway_version_info{value="1.4.0"} 1
# HELP nym_explorer_pretty_bond_location_country_name_info Info metric for explorer_pretty_bond_location_country_name
# TYPE nym_explorer_pretty_bond_location_country_name_info gauge
nym_explorer_pretty_bond_location_country_name_info{value="Japan"} 1
# HELP nym_explorer_pretty_bond_location_latitude Metric for explorer_pretty_bond_location_latitude
# TYPE nym_explorer_pretty_bond_location_latitude gauge
nym_explorer_pretty_bond_location_latitude 35.6837
# HELP nym_explorer_pretty_bond_location_longitude Metric for explorer_pretty_bond_location_longitude
# TYPE nym_explorer_pretty_bond_location_longitude gauge
nym_explorer_pretty_bond_location_longitude 139.6805
# HELP nym_explorer_pretty_bond_location_three_letter_iso_country_code_info Info metric for explorer_pretty_bond_location_three_letter_iso_country_code
# TYPE nym_explorer_pretty_bond_location_three_letter_iso_country_code_info gauge
nym_explorer_pretty_bond_location_three_letter_iso_country_code_info{value="JPN"} 1
# HELP nym_explorer_pretty_bond_location_two_letter_iso_country_code_info Info metric for explorer_pretty_bond_location_two_letter_iso_country_code
# TYPE nym_explorer_pretty_bond_location_two_letter_iso_country_code_info gauge
nym_explorer_pretty_bond_location_two_letter_iso_country_code_info{value="JP"} 1
# HELP nym_explorer_pretty_bond_owner_info Info metric for explorer_pretty_bond_owner
# TYPE nym_explorer_pretty_bond_owner_info gauge
nym_explorer_pretty_bond_owner_info{value="n13q6pz3savlje2yvtcwc5eeej8pnfhkvrf5mh5t"} 1
# HELP nym_explorer_pretty_bond_pledge_amount_amount_info Info metric for explorer_pretty_bond_pledge_amount_amount
# TYPE nym_explorer_pretty_bond_pledge_amount_amount_info gauge
nym_explorer_pretty_bond_pledge_amount_amount_info{value="100000000"} 1
# HELP nym_explorer_pretty_bond_pledge_amount_denom_info Info metric for explorer_pretty_bond_pledge_amount_denom
# TYPE nym_explorer_pretty_bond_pledge_amount_denom_info gauge
nym_explorer_pretty_bond_pledge_amount_denom_info{value="unym"} 1
# HELP nym_gateway_identity_key_info Info metric for gateway_identity_key
# TYPE nym_gateway_identity_key_info gauge
nym_gateway_identity_key_info{value="28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_last_probe_result_gateway_info Info metric for last_probe_result_gateway
# TYPE nym_last_probe_result_gateway_info gauge
nym_last_probe_result_gateway_info{value="28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_last_probe_result_outcome_as_entry_can_connect Metric for last_probe_result_outcome_as_entry_can_connect
# TYPE nym_last_probe_result_outcome_as_entry_can_connect gauge
nym_last_probe_result_outcome_as_entry_can_connect 1
# HELP nym_last_probe_result_outcome_as_entry_can_route Metric for last_probe_result_outcome_as_entry_can_route
# TYPE nym_last_probe_result_outcome_as_entry_can_route gauge
nym_last_probe_result_outcome_as_entry_can_route 1
# HELP nym_last_probe_result_outcome_as_exit_can_connect Metric for last_probe_result_outcome_as_exit_can_connect
# TYPE nym_last_probe_result_outcome_as_exit_can_connect gauge
nym_last_probe_result_outcome_as_exit_can_connect 1
# HELP nym_last_probe_result_outcome_as_exit_can_route_ip_external_v4 Metric for last_probe_result_outcome_as_exit_can_route_ip_external_v4
# TYPE nym_last_probe_result_outcome_as_exit_can_route_ip_external_v4 gauge
nym_last_probe_result_outcome_as_exit_can_route_ip_external_v4 1
# HELP nym_last_probe_result_outcome_as_exit_can_route_ip_external_v6 Metric for last_probe_result_outcome_as_exit_can_route_ip_external_v6
# TYPE nym_last_probe_result_outcome_as_exit_can_route_ip_external_v6 gauge
nym_last_probe_result_outcome_as_exit_can_route_ip_external_v6 1
# HELP nym_last_probe_result_outcome_as_exit_can_route_ip_v4 Metric for last_probe_result_outcome_as_exit_can_route_ip_v4
# TYPE nym_last_probe_result_outcome_as_exit_can_route_ip_v4 gauge
nym_last_probe_result_outcome_as_exit_can_route_ip_v4 1
# HELP nym_last_probe_result_outcome_as_exit_can_route_ip_v6 Metric for last_probe_result_outcome_as_exit_can_route_ip_v6
# TYPE nym_last_probe_result_outcome_as_exit_can_route_ip_v6 gauge
nym_last_probe_result_outcome_as_exit_can_route_ip_v6 1
# HELP nym_last_probe_result_outcome_wg_can_handshake_v4 Metric for last_probe_result_outcome_wg_can_handshake_v4
# TYPE nym_last_probe_result_outcome_wg_can_handshake_v4 gauge
nym_last_probe_result_outcome_wg_can_handshake_v4 1
# HELP nym_last_probe_result_outcome_wg_can_handshake_v6 Metric for last_probe_result_outcome_wg_can_handshake_v6
# TYPE nym_last_probe_result_outcome_wg_can_handshake_v6 gauge
nym_last_probe_result_outcome_wg_can_handshake_v6 1
# HELP nym_last_probe_result_outcome_wg_can_register Metric for last_probe_result_outcome_wg_can_register
# TYPE nym_last_probe_result_outcome_wg_can_register gauge
nym_last_probe_result_outcome_wg_can_register 1
# HELP nym_last_probe_result_outcome_wg_can_resolve_dns_v4 Metric for last_probe_result_outcome_wg_can_resolve_dns_v4
# TYPE nym_last_probe_result_outcome_wg_can_resolve_dns_v4 gauge
nym_last_probe_result_outcome_wg_can_resolve_dns_v4 1
# HELP nym_last_probe_result_outcome_wg_can_resolve_dns_v6 Metric for last_probe_result_outcome_wg_can_resolve_dns_v6
# TYPE nym_last_probe_result_outcome_wg_can_resolve_dns_v6 gauge
nym_last_probe_result_outcome_wg_can_resolve_dns_v6 1
# HELP nym_last_probe_result_outcome_wg_ping_hosts_performance_v4 Metric for last_probe_result_outcome_wg_ping_hosts_performance_v4
# TYPE nym_last_probe_result_outcome_wg_ping_hosts_performance_v4 gauge
nym_last_probe_result_outcome_wg_ping_hosts_performance_v4 1
# HELP nym_last_probe_result_outcome_wg_ping_hosts_performance_v6 Metric for last_probe_result_outcome_wg_ping_hosts_performance_v6
# TYPE nym_last_probe_result_outcome_wg_ping_hosts_performance_v6 gauge
nym_last_probe_result_outcome_wg_ping_hosts_performance_v6 1
# HELP nym_last_probe_result_outcome_wg_ping_ips_performance_v4 Metric for last_probe_result_outcome_wg_ping_ips_performance_v4
# TYPE nym_last_probe_result_outcome_wg_ping_ips_performance_v4 gauge
nym_last_probe_result_outcome_wg_ping_ips_performance_v4 0.6666667
# HELP nym_last_probe_result_outcome_wg_ping_ips_performance_v6 Metric for last_probe_result_outcome_wg_ping_ips_performance_v6
# TYPE nym_last_probe_result_outcome_wg_ping_ips_performance_v6 gauge
nym_last_probe_result_outcome_wg_ping_ips_performance_v6 1
# HELP nym_last_testrun_utc_info Info metric for last_testrun_utc
# TYPE nym_last_testrun_utc_info gauge
nym_last_testrun_utc_info{value="2025-02-08T09:46:26+00:00"} 1
# HELP nym_last_updated_utc_info Info metric for last_updated_utc
# TYPE nym_last_updated_utc_info gauge
nym_last_updated_utc_info{value="2025-02-08T09:46:26+00:00"} 1
# HELP nym_performance Metric for performance
# TYPE nym_performance gauge
nym_performance 91
# HELP nym_routing_score Metric for routing_score
# TYPE nym_routing_score gauge
nym_routing_score 0
# HELP nym_self_described_authenticator_address_info Info metric for self_described_authenticator_address
# TYPE nym_self_described_authenticator_address_info gauge
nym_self_described_authenticator_address_info{value="9Nqp2m4kWNx7NfXQPmL1AyG2E4AVh83XXLJemA3tkvSS.6HFL8xR8JuFea7ersz2DdiQC1StbWkBiDk9LbAneeZRM@28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_self_described_auxiliary_details_accepted_operator_terms_and_conditions Metric for self_described_auxiliary_details_accepted_operator_terms_and_conditions
# TYPE nym_self_described_auxiliary_details_accepted_operator_terms_and_conditions gauge
nym_self_described_auxiliary_details_accepted_operator_terms_and_conditions 1
# HELP nym_self_described_auxiliary_details_location_info Info metric for self_described_auxiliary_details_location
# TYPE nym_self_described_auxiliary_details_location_info gauge
nym_self_described_auxiliary_details_location_info{value="JP"} 1
# HELP nym_self_described_build_information_binary_name_info Info metric for self_described_build_information_binary_name
# TYPE nym_self_described_build_information_binary_name_info gauge
nym_self_described_build_information_binary_name_info{value="nym-node"} 1
# HELP nym_self_described_build_information_build_timestamp_info Info metric for self_described_build_information_build_timestamp
# TYPE nym_self_described_build_information_build_timestamp_info gauge
nym_self_described_build_information_build_timestamp_info{value="2025-02-04T09:35:42.399220545Z"} 1
# HELP nym_self_described_build_information_build_version_info Info metric for self_described_build_information_build_version
# TYPE nym_self_described_build_information_build_version_info gauge
nym_self_described_build_information_build_version_info{value="1.4.0"} 1
# HELP nym_self_described_build_information_cargo_profile_info Info metric for self_described_build_information_cargo_profile
# TYPE nym_self_described_build_information_cargo_profile_info gauge
nym_self_described_build_information_cargo_profile_info{value="release"} 1
# HELP nym_self_described_build_information_cargo_triple_info Info metric for self_described_build_information_cargo_triple
# TYPE nym_self_described_build_information_cargo_triple_info gauge
nym_self_described_build_information_cargo_triple_info{value="x86_64-unknown-linux-gnu"} 1
# HELP nym_self_described_build_information_commit_branch_info Info metric for self_described_build_information_commit_branch
# TYPE nym_self_described_build_information_commit_branch_info gauge
nym_self_described_build_information_commit_branch_info{value="HEAD"} 1
# HELP nym_self_described_build_information_commit_sha_info Info metric for self_described_build_information_commit_sha
# TYPE nym_self_described_build_information_commit_sha_info gauge
nym_self_described_build_information_commit_sha_info{value="4c2bf3642e8eec0d55c7753e14429d73ac2d0e99"} 1
# HELP nym_self_described_build_information_commit_timestamp_info Info metric for self_described_build_information_commit_timestamp
# TYPE nym_self_described_build_information_commit_timestamp_info gauge
nym_self_described_build_information_commit_timestamp_info{value="2025-02-04T10:29:48.000000000+01:00"} 1
# HELP nym_self_described_build_information_rustc_channel_info Info metric for self_described_build_information_rustc_channel
# TYPE nym_self_described_build_information_rustc_channel_info gauge
nym_self_described_build_information_rustc_channel_info{value="stable"} 1
# HELP nym_self_described_build_information_rustc_version_info Info metric for self_described_build_information_rustc_version
# TYPE nym_self_described_build_information_rustc_version_info gauge
nym_self_described_build_information_rustc_version_info{value="1.84.1"} 1
# HELP nym_self_described_declared_role_entry Metric for self_described_declared_role_entry
# TYPE nym_self_described_declared_role_entry gauge
nym_self_described_declared_role_entry 1
# HELP nym_self_described_declared_role_exit_ipr Metric for self_described_declared_role_exit_ipr
# TYPE nym_self_described_declared_role_exit_ipr gauge
nym_self_described_declared_role_exit_ipr 1
# HELP nym_self_described_declared_role_exit_nr Metric for self_described_declared_role_exit_nr
# TYPE nym_self_described_declared_role_exit_nr gauge
nym_self_described_declared_role_exit_nr 1
# HELP nym_self_described_declared_role_mixnode Metric for self_described_declared_role_mixnode
# TYPE nym_self_described_declared_role_mixnode gauge
nym_self_described_declared_role_mixnode 0
# HELP nym_self_described_host_information_hostname_info Info metric for self_described_host_information_hostname
# TYPE nym_self_described_host_information_hostname_info gauge
nym_self_described_host_information_hostname_info{value="nym-exit.hcloud.ltd"} 1
# HELP nym_self_described_host_information_ip_address_info Info metric for self_described_host_information_ip_address
# TYPE nym_self_described_host_information_ip_address_info gauge
nym_self_described_host_information_ip_address_info{value="217.178.53.49"} 1
# HELP nym_self_described_host_information_keys_ed25519_info Info metric for self_described_host_information_keys_ed25519
# TYPE nym_self_described_host_information_keys_ed25519_info gauge
nym_self_described_host_information_keys_ed25519_info{value="28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_self_described_host_information_keys_x25519_info Info metric for self_described_host_information_keys_x25519
# TYPE nym_self_described_host_information_keys_x25519_info gauge
nym_self_described_host_information_keys_x25519_info{value="7YauXgjKA9qoCmU5somVxTRy7XyQtCSc8yh29Ueitv3J"} 1
# HELP nym_self_described_ip_packet_router_address_info Info metric for self_described_ip_packet_router_address
# TYPE nym_self_described_ip_packet_router_address_info gauge
nym_self_described_ip_packet_router_address_info{value="EUP1cHzYqpNPmQVRJQ9G9HAvmxb3woQamH3zpL4DBY6r.5G8NipY4qnrNjXusVFMbujqspK6QwtFPpaT6PhKv5P8m@28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_self_described_last_polled_info Info metric for self_described_last_polled
# TYPE nym_self_described_last_polled_info gauge
nym_self_described_last_polled_info{value="2025-02-08 08:45:04.450067303 +00:00:00"} 1
# HELP nym_self_described_mixnet_websockets_ws_port Metric for self_described_mixnet_websockets_ws_port
# TYPE nym_self_described_mixnet_websockets_ws_port gauge
nym_self_described_mixnet_websockets_ws_port 9000
# HELP nym_self_described_mixnet_websockets_wss_port Metric for self_described_mixnet_websockets_wss_port
# TYPE nym_self_described_mixnet_websockets_wss_port gauge
nym_self_described_mixnet_websockets_wss_port 9001
# HELP nym_self_described_network_requester_address_info Info metric for self_described_network_requester_address
# TYPE nym_self_described_network_requester_address_info gauge
nym_self_described_network_requester_address_info{value="9zJ5cbqMjVbL6CXs1nUyRNYcMHbYdyY7q1HrjG1SViDb.9UrAkEpx2ZKTz7S8FLCeNy9XZdqGrcHL5DdSMZuemYhZ@28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND"} 1
# HELP nym_self_described_network_requester_uses_exit_policy Metric for self_described_network_requester_uses_exit_policy
# TYPE nym_self_described_network_requester_uses_exit_policy gauge
nym_self_described_network_requester_uses_exit_policy 1
# HELP nym_self_described_wireguard_port Metric for self_described_wireguard_port
# TYPE nym_self_described_wireguard_port gauge
nym_self_described_wireguard_port 51822
# HELP nym_self_described_wireguard_public_key_info Info metric for self_described_wireguard_public_key
# TYPE nym_self_described_wireguard_public_key_info gauge
nym_self_described_wireguard_public_key_info{value="E1gb9o4EhYLtEttzBDrxg38heFmTJLA5fMwJSmBY4pau"} 1
```

## Notes

- This exporter currently handles **gateway** data only.
- Data is refreshed at intervals of **10 minutes**.
- The exporter is **dependent** on the Harbourmaster API, so any changes there might require updates to this tool.

## Support

Stake my Gateway or Mixnode if you like.

- https://explorer.nym.spectredao.net/nodes/71HcWeM4o1TQaL56vctCCweRXCG8P6c3xiPCj1ugM8zn

- https://explorer.nym.spectredao.net/nodes/28tXg9mEW4mifgU1TdetVVAN5PvmhtLpHzFRMfJBT6ND

## TODO

- [ ] Add Mixnode data
