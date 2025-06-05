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

   Quick creation with local output paths

   --local-output value [ --local-output value ]  The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path

   Quick creation with local source paths

   --local-source value [ --local-source value ]  The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path

```
{% endcode %}
