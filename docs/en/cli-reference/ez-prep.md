# Prepare a dataset from a local path

{% code fullWidth="true" %}
```
NAME:
   singularity ez-prep - Prepare a dataset from a local path

USAGE:
   singularity ez-prep [command options] <path>

CATEGORY:
   Utility

DESCRIPTION:
   This commands can be used to prepare a dataset from a local path with minimum configurable parameters.
   For more advanced usage, please use the subcommands under `storage` and `data-prep`.
   You can also use this command for benchmarking with in-memory database and inline preparation, i.e.
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

OPTIONS:
   --max-size value, -M value       Maximum size of the CAR files to be created (default: "31.5GiB")
   --output-dir value, -o value     Output directory for CAR files. To use inline preparation, use an empty string (default: "./cars")
   --concurrency value, -j value    Concurrency for packing (default: 1)
   --database-file value, -f value  The database file to store the metadata. To use in memory database, use an empty string. (default: ./ezprep-<name>.db)
   --help, -h                       show help
```
{% endcode %}
