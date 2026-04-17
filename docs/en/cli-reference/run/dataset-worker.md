# Start a dataset preparation worker to process dataset scanning and preparation tasks

{% code fullWidth="true" %}
```
NAME:
   singularity run dataset-worker - Start a dataset preparation worker to process dataset scanning and preparation tasks

USAGE:
   singularity run dataset-worker [command options]

OPTIONS:
   --no-automigrate      skip automatic database migration and correctness checks on startup; only use if you run 'admin init' on every upgrade or manually before starting daemons (default: false)
   --concurrency value   Number of concurrent workers to run (default: 1)
   --enable-scan         Enable scanning of datasets (default: true)
   --enable-pack         Enable packing of datasets that calculates CIDs and packs them into CAR files (default: true)
   --enable-dag          Enable dag generation of datasets that maintains the directory structure of datasets (default: true)
   --enable-reaper       Enable the orphan-record reaper. Exactly one dataset-worker process per deployment should enable this; running multiple reapers contends on the same rows and can livelock. (default: true)
   --exit-on-complete    Exit the worker when there is no more work to do (default: false)
   --exit-on-error       Exit the worker when there is any error (default: false)
   --min-interval value  How often to check for new jobs (minimum) (default: 5s)
   --max-interval value  How often to check for new jobs (maximum) (default: 2m40s)
   --help, -h            show help
```
{% endcode %}
