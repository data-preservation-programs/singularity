# Migrate dataset from old singularity mongodb

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity admin migrate-dataset - Migrate dataset from old singularity mongodb

USAGE:
   singularity singularity singularity singularity admin migrate-dataset [command options] [arguments...]

DESCRIPTION:
   Migrate datasets from singularity V1 to V2. Those steps include
     1. Create source storage and output storage and attach them to a dataprep in V2.
     2. Create all folder structures and files in the new dataset.
   Caveats:
     1. The created preparation won't be compatible with the new dataset worker.
        So do not attempt to resume a data preparation or push new files onto migrated dataset.
        You can make deals or browse the dataset without issues.
     2. The folder CID won't be generated or migrated due to the complexity

OPTIONS:
   --mongo-connection-string value  MongoDB connection string (default: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --skip-files                     Skip migrating details about files and folders. This will make the migration much faster. Useful if you only want to make deals. (default: false)
   --help, -h                       show help
```
{% endcode %}
