# Create a new preparation

{% code fullWidth="true" %}
```
NAME:
   singularity prep create - Create a new preparation

USAGE:
   singularity prep create [command options]

CATEGORY:
   Preparation Management

OPTIONS:
   --delete-after-export              Whether to delete the source files after export to CAR files (default: false)
   --help, -h                         show help
   --max-size value                   The maximum size of a single CAR file (default: "31.5GiB")
   --min-piece-size value             The minimum size of a piece. Pieces smaller than this will be padded up to this size. It's recommended to leave this as the default (default: 1MiB)
   --name value                       The name for the preparation (default: Auto generated)
   --no-dag                           Whether to disable maintaining folder dag structure for the sources. If disabled, DagGen will not be possible and folders will not have an associated CID. (default: false)
   --no-inline                        Whether to disable inline storage for the preparation. Can save database space but requires at least one output storage. (default: false)
   --output value [ --output value ]  The id or name of the output storage to be used for the preparation
   --piece-size value                 The target piece size of the CAR files used for piece commitment calculation (default: Determined by --max-size)
   --source value [ --source value ]  The id or name of the source storage to be used for the preparation

   Auto Deal Creation

   --auto-create-deals              Enable automatic deal schedule creation after preparation completion (default: false)
   --deal-announce-to-ipni          Whether to announce deals to IPNI (default: false)
   --deal-duration value            Duration for storage deals (e.g., 535 days) (default: 0s)
   --deal-http-headers value        HTTP headers for deals in JSON format
   --deal-keep-unsealed             Whether to keep unsealed copy of deals (default: false)
   --deal-price-per-deal value      Price in FIL per deal for storage deals (default: 0)
   --deal-price-per-gb value        Price in FIL per GiB for storage deals (default: 0)
   --deal-price-per-gb-epoch value  Price in FIL per GiB per epoch for storage deals (default: 0)
   --deal-provider value            Storage Provider ID for deals (e.g., f01000)
   --deal-start-delay value         Start delay for storage deals (e.g., 72h) (default: 0s)
   --deal-url-template value        URL template for deals
   --deal-verified                  Whether deals should be verified (default: false)

   Quick creation with local output paths

   --local-output value [ --local-output value ]  The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path

   Quick creation with local source paths

   --local-source value [ --local-source value ]  The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path

   Validation

   --sp-validation      Enable storage provider validation before deal creation (default: false)
   --wallet-validation  Enable wallet balance validation before deal creation (default: false)

```
{% endcode %}
