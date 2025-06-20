# Create a new deal template

{% code fullWidth="true" %}
```
NAME:
   singularity deal-template create - Create a new deal template

USAGE:
   singularity deal-template create [command options]

CATEGORY:
   Deal Template Management

OPTIONS:
   --name value                     Name of the deal template
   --description value              Description of the deal template
   --deal-price-per-gb value        Price in FIL per GiB for storage deals (default: 0)
   --deal-price-per-gb-epoch value  Price in FIL per GiB per epoch for storage deals (default: 0)
   --deal-price-per-deal value      Price in FIL per deal for storage deals (default: 0)
   --deal-duration value            Duration for storage deals (e.g., 535 days) (default: 0s)
   --deal-start-delay value         Start delay for storage deals (e.g., 72h) (default: 0s)
   --deal-verified                  Whether deals should be verified (default: false)
   --deal-keep-unsealed             Whether to keep unsealed copy of deals (default: false)
   --deal-announce-to-ipni          Whether to announce deals to IPNI (default: false)
   --deal-provider value            Storage Provider ID for deals (e.g., f01000)
   --deal-url-template value        URL template for deals
   --deal-http-headers value        HTTP headers for deals in JSON format
   --help, -h                       show help
```
{% endcode %}
