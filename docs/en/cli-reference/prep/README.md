# Create and manage dataset preparations

{% code fullWidth="true" %}
```
NAME:
   singularity prep - Create and manage dataset preparations

USAGE:
   singularity prep command [command options] [arguments...]

COMMANDS:
   create         Create a new preparation
   list           List all preparations
   status         Get the preparation job status of a preparation
   rename         Rename a preparation
   attach-source  Attach a source storage to a preparation
   attach-output  Attach a output storage to a preparation
   detach-output  Detach a output storage to a preparation
   start-scan     Start scanning of the source storage
   pause-scan     Pause a scanning job
   start-pack     Start / Restart all pack jobs or a specific one
   pause-pack     Pause all pack jobs or a specific one
   start-daggen   Start a DAG generation that creates a snapshot of all folder structures
   pause-daggen   Pause a DAG generation job
   list-pieces    List all generated pieces for a preparation
   add-piece      Manually add piece info to a preparation. This is useful for pieces prepared by external tools.
   explore        Explore prepared source by path
   attach-wallet  Attach a wallet to a preparation
   list-wallets   List attached wallets with a preparation
   detach-wallet  Detach a wallet to a preparation
   remove         Remove a preparation
   help, h        Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
