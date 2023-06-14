# Dataset management

{% code fullWidth="true" %}
```
NAME:
   singularity dataset - Dataset management

USAGE:
   singularity dataset command [command options] [arguments...]

COMMANDS:
   create         Create a new dataset
   list           List all datasets
   update         Update an existing dataset
   remove         Remove a specific dataset. This will not remove the CAR files.
   add-wallet     Associate a wallet with the dataset. The wallet needs to be imported first using the `singularity wallet import` command.
   list-wallet    List all associated wallets with the dataset
   remove-wallet  Remove an associated wallet from the dataset
   add-piece      Manually register a piece (CAR file) with the dataset for deal making purpose
   list-pieces    List all pieces for the dataset that are available for deal making
   help, h        Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
