# Migrate schedule from old singularity mongodb

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-schedule - Migrate schedule from old singularity mongodb

USAGE:
   singularity admin migrate-schedule [command options] [arguments...]

DESCRIPTION:
   Migrate schedules from singularity V1 to V2. Note that
     1. You must complete dataset migration first
     2. You will need to import all relevant private keys to the database
     3. All new schedules will be created with status 'paused'
     4. The deal states will not be migrated over as it will be populated with deal tracker automatically.

OPTIONS:
   --mongo-connection-string value  MongoDB connection string (default: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --help, -h                       show help
```
{% endcode %}
