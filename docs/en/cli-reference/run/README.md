# run different singularity components

{% code fullWidth="true" %}
```
NAME:
   singularity run - run different singularity components

USAGE:
   singularity run command [command options]

COMMANDS:
   api               Run the singularity API
   dataset-worker    Start a dataset preparation worker to process dataset scanning and preparation tasks
   content-provider  Start a content provider that serves retrieval requests
   deal-tracker      Start a deal tracker that tracks the deal for all relevant wallets
   deal-pusher       Start a deal pusher that monitors deal schedules and pushes deals to storage providers
   download-server   An HTTP server connecting to remote metadata API to offer CAR file downloads
   unified           Start the unified auto-preparation service with worker management and workflow orchestration
   help, h           Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
