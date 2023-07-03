# Run different singularity components

{% code fullWidth="true" %}
```
NAME:
   singularity run - Run different singularity components

USAGE:
   singularity run command [command options] [arguments...]

COMMANDS:
   api               Run the singularity API
   dataset-worker    Start a dataset preparation worker to process dataset scanning and preparation tasks
   content-provider  Start a content provider that serves retrieval requests
   deal-tracker      Start a deal tracker that tracks the deal for all relevant wallets
   dealmaker         Start a deal making/tracking worker to process deal making
   spade-api         Start a Spade compatible API for storage provider deal proposal self service
   help, h           Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
