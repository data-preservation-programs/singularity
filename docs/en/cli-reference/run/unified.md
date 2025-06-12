# Run unified auto-preparation service (workflow orchestration + worker management)

{% code fullWidth="true" %}
```
NAME:
   singularity run unified - Run unified auto-preparation service (workflow orchestration + worker management)

USAGE:
   singularity run unified [command options]

DESCRIPTION:
   The unified service combines workflow orchestration and worker lifecycle management.
     
   It automatically:
   - Manages dataset worker lifecycle (start/stop workers based on job availability)
   - Orchestrates job progression (scan → pack → daggen → deals)
   - Scales workers up/down based on job queue
   - Handles automatic deal creation when preparations complete

   This is the recommended way to run fully automated data preparation.

OPTIONS:
   --min-workers value               Minimum number of workers to keep running (default: 1)
   --max-workers value               Maximum number of workers to run (default: 5)
   --scale-up-threshold value        Number of ready jobs to trigger worker scale-up (default: 5)
   --scale-down-threshold value      Number of ready jobs below which to scale down workers (default: 2)
   --check-interval value            How often to check for scaling and workflow progression (default: 30s)
   --worker-idle-timeout value       How long a worker can be idle before shutdown (0 = never) (default: 5m0s)
   --disable-auto-scaling            Disable automatic worker scaling (default: false)
   --disable-workflow-orchestration  Disable automatic job progression (default: false)
   --disable-auto-deals              Disable automatic deal creation (default: false)
   --disable-scan-to-pack            Disable automatic scan → pack transitions (default: false)
   --disable-pack-to-daggen          Disable automatic pack → daggen transitions (default: false)
   --disable-daggen-to-deals         Disable automatic daggen → deals transitions (default: false)
   --help, -h                        show help
```
{% endcode %}
