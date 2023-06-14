# Create a new dataset

```
NAME:
   singularity dataset create - Create a new dataset

USAGE:
   singularity dataset create [command options] <dataset_name>

DESCRIPTION:
   <dataset_name> must be a unique identifier for a dataset
   The dataset is a top level object to distinguish different dataset.

OPTIONS:
   --help, -h  show help

   Encryption

   --encryption-recipient value [ --encryption-recipient value ]  [Alpha] Public key of the encryption recipient
   --encryption-script value                                      [WIP] EncryptionScript command to run for custom encryption

   Inline Preparation

   --output-dir value, -o value [ --output-dir value, -o value ]  Output directory for CAR files (default: not needed)

   Preparation Parameters

   --max-size value, -M value    Maximum size of the CAR files to be created (default: "31.5GiB")
   --piece-size value, -s value  Target piece size of the CAR files used for piece commitment calculation (default: inferred)

```
