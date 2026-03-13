# Start a deal pusher that monitors deal schedules and pushes deals to storage providers

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-pusher - Start a deal pusher that monitors deal schedules and pushes deals to storage providers

USAGE:
   singularity run deal-pusher [command options]

OPTIONS:
   --deal-attempts value, -d value           Number of times to attempt a deal before giving up (default: 3)
   --max-replication-factor value, -M value  Max number of replicas for each individual PieceCID across all clients and providers (default: Unlimited)
   --pdp-batch-size value                    Number of roots to include in each PDP add-roots transaction (default: 128)
   --pdp-max-pieces-per-proofset value       Maximum pieces per proof set before handing off to the storage provider (default: 1024)
   --pdp-confirmation-depth value            Number of block confirmations required for PDP transactions (default: 5)
   --pdp-poll-interval value                 Polling interval for PDP transaction confirmation checks (default: 30s)
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
