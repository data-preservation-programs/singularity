# Start a dataset preparation worker to process dataset scanning and preparation tasks

{% code fullWidth="true" %}
```
NAME:
   singularity run dataset-worker - Start a dataset preparation worker to process dataset scanning and preparation tasks

USAGE:
   singularity run dataset-worker [command options]

OPTIONS:
   --concurrency value   Number of concurrent workers to run (default: 1)
   --enable-scan         Enable scanning of datasets (default: true)
   --enable-pack         Enable packing of datasets that calculates CIDs and packs them into CAR files (default: true)
   --enable-dag          Enable dag generation of datasets that maintains the directory structure of datasets (default: true)
   --exit-on-complete    Exit the worker when there is no more work to do (default: false)
   --exit-on-error       Exit the worker when there is any error (default: false)
   --min-interval value  How often to check for new jobs (minimum) (default: 5s)
   --max-interval value  How often to check for new jobs (maximum) (default: 2m40s)
   --help, -h            show help
```
{% endcode %}
