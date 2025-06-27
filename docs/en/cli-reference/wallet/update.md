# Update wallet details

{% code fullWidth="true" %}
```
NAME:
   singularity wallet update - Update wallet details

USAGE:
   singularity wallet update [command options] <address>

DESCRIPTION:
   Update non-essential details of an existing wallet.

   This command allows you to update the following wallet properties:
   - Name (optional wallet label)
   - Contact information (email for SP)
   - Location (region, country for SP)

   Essential properties like the wallet address, private key, and balance cannot be modified.

   EXAMPLES:
       # Update the actor name
       singularity wallet update f1abc123... --name "My Main Wallet"

       # Update multiple fields at once
       singularity wallet update f1xyz789... --name "Storage Provider" --location "US-East"

OPTIONS:
   --name value      Set the readable label for the wallet
   --contact value   Set the contact information (email) for the wallet
   --location value  Set the location (region, country) for the wallet
   --help, -h        show help
```
{% endcode %}
