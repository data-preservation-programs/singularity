# Import a wallet from exported private key

{% code fullWidth="true" %}
```
NAME:
   singularity wallet import - Import a wallet from exported private key

USAGE:
   singularity wallet import [command options] [path, or stdin if omitted]

DESCRIPTION:
   Import an existing Filecoin wallet from an exported private key file.
   
   The command supports importing wallets with optional metadata fields for
   better organization and identification:
   - Name: A human-readable display name for the wallet
   - Contact: Contact information (email, etc.)
   - Location: Geographic location or region

   EXAMPLES:
     Basic import from file
       singularity wallet import /path/to/private-key.json

     Import with metadata
       singularity wallet import \
         --name "Storage Provider Main Wallet" \
         --contact "admin@example.com" \
         --location "US-East" \
         /path/to/private-key.json

     Import from stdin
       cat private-key.json | singularity wallet import --name "My Wallet"

   The imported wallet will be automatically added to your local wallet
   collection and can be used for deal creation and other operations.

OPTIONS:
   --name value      Optional display name for the wallet
   --contact value   Optional contact information (email, etc.)
   --location value  Optional location or region identifier
   --help, -h        show help
```
{% endcode %}
