# Get wallet balance information

{% code fullWidth="true" %}
```
NAME:
   singularity wallet balance - Get wallet balance information

USAGE:
   singularity wallet balance [command options] <wallet_address>

DESCRIPTION:
   Get FIL balance and FIL+ datacap balance for a specific wallet address.
   This command queries the Lotus network to retrieve current balance information.

   Examples:
     singularity wallet balance f12syf7zd3lfsv43aj2kb454ymaqw7debhumjnbqa
     singularity wallet balance --json f1abc123...def456

   The command returns:
   - FIL balance in human-readable format (e.g., "1.000000 FIL")
   - Raw balance in attoFIL for precise calculations
   - FIL+ datacap balance in GiB format (e.g., "1024.50 GiB") 
   - Raw datacap in bytes

   If there are issues retrieving either balance, partial results will be shown with error details.

OPTIONS:
   --help, -h  show help
```
{% endcode %}
