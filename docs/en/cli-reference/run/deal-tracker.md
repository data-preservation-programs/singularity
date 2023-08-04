# Start a deal tracker that tracks the deal for all relevant wallets

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-tracker - Start a deal tracker that tracks the deal for all relevant wallets

USAGE:
   singularity run deal-tracker [command options] [arguments...]

OPTIONS:
   --market-deal-url value, -m value  The URL for ZST compressed state market deals json. Set to empty to use Lotus API. (default: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         How often to check for new deals (default: 1h0m0s)
   --once, -o                         Run once and exit (default: false)
   --help, -h                         show help
```
{% endcode %}
