# Track PDP deals via Shovel event indexing (requires PostgreSQL)

{% code fullWidth="true" %}
```
NAME:
   singularity run pdp-tracker - Track PDP deals via Shovel event indexing (requires PostgreSQL)

USAGE:
   singularity run pdp-tracker [command options]

OPTIONS:
   --eth-rpc value            Ethereum RPC endpoint for FEVM (default: "https://api.node.glif.io/rpc/v1") [$ETH_RPC_URL]
   --pdp-poll-interval value  How often to check for new events in Shovel tables (default: 30s)
   --full-sync                Re-index events from contract deployment by resetting the Shovel cursor. Derived PDP state (proof sets, deals) is preserved and updated via upserts. Requires an archival RPC node. Involves one RPC call per historical proof set. (default: false)
   --help, -h                 show help
```
{% endcode %}
