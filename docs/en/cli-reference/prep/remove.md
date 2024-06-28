# Remove a preparation

{% code fullWidth="true" %}
```
NAME:
   singularity prep remove - Remove a preparation

USAGE:
   singularity prep remove [command options]<name|id>

DESCRIPTION:
   This will remove all relevant information, including:
     * All related jobs
     * All related piece info
     * Mapping used for Inline Preparation
     * All File and Directory data and CIDs
     * All Schedules
   This will not remove
     * All deals ever made

OPTIONS:
   --cars      Also remove prepared CAR files (default: false)
   --help, -h  show help
```
{% endcode %}
