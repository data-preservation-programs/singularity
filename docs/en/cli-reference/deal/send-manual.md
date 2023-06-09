# Send a manual deal proposal to boost or legacy market

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual - Send a manual deal proposal to boost or legacy market

USAGE:
   singularity deal send-manual [command options] CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE

OPTIONS:
   --help, -h       show help
   --timeout value  Timeout for the deal proposal (default: 1m)

   Boost Only

   --file-size value                            File size in bytes for boost to fetch the CAR file (default: 0)
   --http-header value [ --http-header value ]  http headers to be passed with the request (i.e. key=value)
   --ipni                                       Whether to announce the deal to IPNI (default: true)
   --url-template value                         URL template with PIECE_CID placeholder for boost to fetch the CAR file, i.e. http://127.0.0.1/piece/{PIECE_CID}.car

   Deal Proposal

   --duration value, -d value     Duration in epoch or in duration format, i.e. 1500000, 2400h (default: 12840h[535 days])
   --keep-unsealed                Whether to keep unsealed copy (default: true)
   --price-per-deal value         Price in FIL per deal (default: 0)
   --price-per-gb value           Price in FIL  per GiB (default: 0)
   --price-per-gb-epoch value     Price in FIL per GiB per epoch (default: 0)
   --root-cid value               Root CID that is required as part of the deal proposal, if empty, will be set to empty CID (default: Empty CID)
   --start-delay value, -s value  Deal start delay in epoch or in duration format, i.e. 1000, 72h (default: 72h[3 days])
   --verified                     Whether to propose deals as verified (default: true)

   Lotus

   --lotus-api value    Lotus RPC API endpoint, only used to get miner info (default: "https://api.node.glif.io/rpc/v1")
   --lotus-token value  Lotus RPC API token, only used to get miner info

```
{% endcode %}
