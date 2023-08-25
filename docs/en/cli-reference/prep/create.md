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
   --encryption-recipient value [ --encryption-recipient value ]  The public key of the encryption recipient
   --help, -h                                                     show help
   --max-size value                                               The maximum size of a single CAR file (default: "31.5GiB")
   --output value [ --output value ]                              The name of the output storage to be used for the preparation
   --piece-size value                                             The target piece size of the CAR files used for piece commitment calculation (default: "32GiB")
   --source value [ --source value ]                              The name of the source storage to be used for the preparation

   Quick creation with local paths

   --local-output value [ --local-output value ]  The local output path to be used for the preparation. This is a convenient flag that will create a output storage with the provided path
   --local-source value [ --local-source value ]  The local source path to be used for the preparation. This is a convenient flag that will create a source storage with the provided path

```
{% endcode %}
