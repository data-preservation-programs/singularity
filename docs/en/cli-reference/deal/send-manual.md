# Send a manual deal proposal to boost or legacy market

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual - Send a manual deal proposal to boost or legacy market

USAGE:
   singularity deal send-manual [command options] CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE

DESCRIPTION:
   Send a manual deal proposal to boost or legacy market
     Example: singularity deal send-manual f01234 f05678 bagaxxxx 32GiB
   Notes:
     * The client address must have been imported to the wallet using 'singularity wallet import'
     * The deal proposal will not be saved in the database however will eventually be tracked if the deal tracker is running
     * There is a quick address verification using GLIF API which can be made faster by setting LOTUS_API and LOTUS_TOKEN to your own lotus node

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

```
{% endcode %}
