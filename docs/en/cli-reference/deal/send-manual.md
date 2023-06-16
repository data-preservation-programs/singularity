# Send a manual deal proposal to boost or legacy market

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual - Send a manual deal proposal to boost or legacy market

USAGE:
   singularity deal send-manual [command options] CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE

OPTIONS:
   --help, -h  show help

   Boost Only

   --file-size value                            File size in bytes for boost to fetch the CAR file (default: 0)
   --http-header value [ --http-header value ]  http headers to be passed with the request (i.e. key=value)
   --ipni                                       Whether to announce the deal to IPNI (default: true)
   --url-template value                         URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car

   Deal Proposal

   --duration value, -d value     Duration in days for deal length (default: 535)
   --keep-unsealed                Whether to keep unsealed copy (default: true)
   --price value, -p value        Price per 32GiB Deal over whole duration in Fil (default: 0)
   --root-cid value               Root CID that is required as part of the deal proposal, if empty, will be set to empty CID (default: Empty CID)
   --start-delay value, -s value  Deal start delay in days (default: 3)
   --verified                     Whether to propose deals as verified (default: true)

   Lotus

   --lotus-api value    Lotus RPC API endpoint, only used to get miner info (default: "https://api.node.glif.io/rpc/v1")
   --lotus-token value  Lotus RPC API token, only used to get miner info

```
{% endcode %}
