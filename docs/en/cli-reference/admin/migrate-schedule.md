# Migrate schedule from old singularity mongodb

{% code fullWidth="true" %}
```
NAME:
   singularity singularity singularity singularity admin migrate-schedule - Migrate schedule from old singularity mongodb

USAGE:
   singularity singularity singularity singularity admin migrate-schedule [command options] [arguments...]

DESCRIPTION:
   Migrate schedules from singularity V1 to V2. Note that
     1. You must complete dataset migration first
     2. All new schedules will be created with status 'paused'
     3. The deal states will not be migrated over as it will be populated with deal tracker automatically
     4. --output-csv is no longer supported. We will provide a new tool in the future
     5. # of replicas is no longer supported as part of the schedule. We will make this a configurable policy in the future
     6. --force is no longer supported. We may add similar support to ignore all policy restrictions in the future
     7. --offline is no longer supported. It will be always offline deal for legacy market and online deal for boost market if URL template is configured

OPTIONS:
   --mongo-connection-string value  MongoDB connection string (default: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --help, -h                       show help
```
{% endcode %}
