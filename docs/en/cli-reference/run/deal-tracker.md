# Start a deal tracker that tracks the deal for all relevant wallets

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-tracker - Start a deal tracker that tracks the deal for all relevant wallets

USAGE:
   singularity run deal-tracker [command options]

OPTIONS:
   --no-automigrate                   skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons (default: false)
   --market-deal-url value, -m value  The URL for ZST compressed state market deals json. Set to empty to use Lotus API. (default: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         How often to check for new deals (default: 1h0m0s)
   --once                             Run once and exit (default: false)
   --eth-rpc value                    Ethereum RPC endpoint for FEVM (required for DDO allocation tracking and paid f05 transaction receipt tracking) [$ETH_RPC_URL]
   --ddo-contract value               DDO Diamond proxy contract address [$DDO_CONTRACT_ADDRESS]
   --help, -h                         show help
```
{% endcode %}
