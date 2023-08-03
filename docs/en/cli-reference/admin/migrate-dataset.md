# Migrate dataset from old singularity mongodb

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-dataset - Migrate dataset from old singularity mongodb

USAGE:
   singularity admin migrate-dataset [command options] [arguments...]

DESCRIPTION:
   Migrate datasets from singularity V1 to V2. The steps include
     1. Create a new dataset in V2.
     2. Create a new datasource which is either an S3 path or local path.
     3. Create all folder structures and files in the new dataset.
   Caveats:
     1. The created dataset won't be compatible with the new dataset worker.
        So do not attempt to trigger a data preparation on migrated dataset.
     2. The folder CID won't be generated or migrated

OPTIONS:
   --mongo-connection-string value  MongoDB connection string (default: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --help, -h                       show help
```
{% endcode %}
