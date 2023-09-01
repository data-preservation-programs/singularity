# Create a new preparation

{% code fullWidth="true" %}
```
NAME:
   singularity prep create - Create a new preparation

USAGE:
   singularity prep create [command options] [arguments...]

CATEGORY:
   Preparation Management

OPTIONS:
   --delete-after-export              Whether to delete the source files after export to CAR files (default: false)
   --help, -h                         show help
   --max-size value                   The maximum size of a single CAR file (default: "31.5GiB")
   --output value [ --output value ]  The name of the output storage to be used for the preparation
   --piece-size value                 The target piece size of the CAR files used for piece commitment calculation (default: Determined by --max-size)
   --source value [ --source value ]  The name of the source storage to be used for the preparation

   Quick creation with local output paths

   --local-output value [ --local-output value ]  The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path

```
{% endcode %}
