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
   --eth-rpc value                           Ethereum RPC endpoint for FEVM (required to execute PDP schedules on-chain) [$ETH_RPC_URL]
   --pdp-contract-address value              Override PDPVerifier contract address (for devnet/testing) [$PDP_CONTRACT_ADDRESS]
   --help, -h                                show help
```
{% endcode %}
