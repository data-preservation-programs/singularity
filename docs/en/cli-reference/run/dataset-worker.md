# Start a dataset preparation worker to process dataset scanning and preparation tasks

{% code fullWidth="true" %}
```
NAME:
   singularity run dataset-worker - Start a dataset preparation worker to process dataset scanning and preparation tasks

USAGE:
   singularity run dataset-worker [command options] [arguments...]

OPTIONS:
   --concurrency value  Number of concurrent workers to run (default: 1) [$DATASET_WORKER_CONCURRENCY]
   --enable-scan        Enable scanning of datasets (default: true) [$DATASET_WORKER_ENABLE_SCAN]
   --enable-pack        Enable packing of datasets that calculates CIDs and packs them into CAR files (default: true) [$DATASET_WORKER_ENABLE_PACK]
   --enable-dag         Enable dag generation of datasets that maintains the directory structure of datasets (default: true) [$DATASET_WORKER_ENABLE_DAG]
   --exit-on-complete   Exit the worker when there is no more work to do (default: false) [$DATASET_WORKER_EXIT_ON_COMPLETE]
   --exit-on-error      Exit the worker when there is any error (default: false) [$DATASET_WORKER_EXIT_ON_ERROR]
   --help, -h           show help
```
{% endcode %}
