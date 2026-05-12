# Start a deal pusher that monitors deal schedules and pushes deals to storage providers

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-pusher - Start a deal pusher that monitors deal schedules and pushes deals to storage providers

USAGE:
   singularity run deal-pusher [command options]

OPTIONS:
   --no-automigrate                          skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons (default: false)
   --deal-attempts value, -d value           Number of times to attempt a deal before giving up (default: 3)
   --max-replication-factor value, -M value  Max number of replicas for each individual PieceCID across all clients and providers (default: Unlimited)
   --pdp-batch-size value                    Number of pieces to include in each /pdp/piece/pull request (default: 128)
   --pdp-max-pieces-per-proofset value       Maximum pieces per proof set before starting a new one (default: 1024)
   --pdp-pull-timeout value                  How long to wait for Curio to finish pulling a batch (per request) (default: 5m0s)
   --pdp-source-url-base value               HTTPS base URL where Curio fetches pieces from; sourceUrl is built as <base>/piece/<pieceCid> [$PDP_SOURCE_URL_BASE]
   --pdp-record-keeper value                 FWSS contract address (recordKeeper). Defaults to the network default from go-synapse. [$PDP_RECORD_KEEPER]
   --eth-rpc value                           Ethereum RPC endpoint for FEVM (required to execute PDP and DDO schedules on-chain) [$ETH_RPC_URL]
   --ddo-contract value                      DDO Diamond proxy contract address [$DDO_CONTRACT_ADDRESS]
   --ddo-payments-contract value             DDO Payments proxy contract address [$DDO_PAYMENTS_CONTRACT_ADDRESS]
   --ddo-payment-token value                 ERC20 payment token address (e.g. USDFC) [$DDO_PAYMENT_TOKEN]
   --ddo-batch-size value                    Number of pieces per DDO allocation transaction (default: 10)
   --ddo-confirmation-depth value            Number of block confirmations required for DDO transactions (default: 5)
   --ddo-poll-interval value                 Polling interval for DDO transaction confirmation checks (default: 30s)
   --ddo-term-min value                      Minimum term in epochs (~6 months default) (default: 518400)
   --ddo-term-max value                      Maximum term in epochs (~5 years default) (default: 5256000)
   --ddo-expiration-offset value             Expiration offset in epochs (default: 172800)
   --help, -h                                show help
```
{% endcode %}
