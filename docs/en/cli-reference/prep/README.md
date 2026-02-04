# Create and manage dataset preparations

{% code fullWidth="true" %}
```
NAME:
   singularity prep - Create and manage dataset preparations

USAGE:
   singularity prep command [command options]

COMMANDS:
   rename   Rename a preparation
   remove   Remove a preparation
   help, h  Shows a list of commands or help for one command
   Job Management:
     status        Get the preparation job status of a preparation
     start-scan    Start scanning of the source storage
     pause-scan    Pause a scanning job
     start-pack    Start / Restart all pack jobs or a specific one
     pause-pack    Pause all pack jobs or a specific one
     start-daggen  Start a DAG generation that creates a snapshot of all folder structures
     pause-daggen  Pause a DAG generation job
   Piece Management:
     list-pieces   List all generated pieces for a preparation
     add-piece     Manually add piece info to a preparation. This is useful for pieces prepared by external tools.
     delete-piece  Delete a piece from a preparation
   Preparation Management:
     create         Create a new preparation
     list           List all preparations
     attach-source  Attach a source storage to a preparation
     attach-output  Attach a output storage to a preparation
     detach-output  Detach a output storage to a preparation
     explore        Explore prepared source by path
   Wallet Management:
     attach-wallet  Attach a wallet to a preparation
     list-wallets   List attached wallets with a preparation
     detach-wallet  Detach a wallet to a preparation

OPTIONS:
   --help, -h  show help
```
{% endcode %}
