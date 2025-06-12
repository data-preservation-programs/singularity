# Start the auto-deal daemon to automatically create deal schedules when preparations complete

{% code fullWidth="true" %}
```
NAME:
   singularity run autodeal - Start the auto-deal daemon to automatically create deal schedules when preparations complete

USAGE:
   singularity run autodeal [command options]

OPTIONS:
   --check-interval value  How often to check for ready preparations (default: 30s)
   --enable-batch-mode     Enable batch processing mode to scan for ready preparations (default: true)
   --exit-on-complete      Exit when there are no more preparations to process (default: false)
   --exit-on-error         Exit when any error occurs (default: false)
   --max-retries value     Maximum number of retries before backing off (0 = unlimited) (default: 3)
   --retry-interval value  Base interval for retry backoff (default: 5m0s)
   --enable-job-hooks      Enable automatic triggering on job completion (requires worker integration) (default: true)
   --help, -h              show help
```
{% endcode %}
