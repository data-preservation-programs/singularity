# Update an existing dataset

{% code fullWidth="true" %}
```
NAME:
   singularity dataset update - Update an existing dataset

USAGE:
   singularity dataset update [command options] <dataset_name>

OPTIONS:
   --help, -h  show help

   Encryption

   --encryption-recipient value [ --encryption-recipient value ]  Public key of the encryption recipient

   Inline Preparation

   --output-dir value, -o value [ --output-dir value, -o value ]  Output directory for CAR files (default: not needed)

   Preparation Parameters

   --max-size value, -M value    Maximum size of the CAR files to be created (default: "30GiB")
   --piece-size value, -s value  Target piece size of the CAR files used for piece commitment calculation (default: inferred)

```
{% endcode %}
