# List all imported wallets

{% code fullWidth="true" %}
```
NAME:
   singularity wallet list - List all imported wallets

USAGE:
   singularity wallet list [command options]

DESCRIPTION:
   Display all imported wallets with their details including metadata.
   
   The output includes all available wallet information:
   - Address: The wallet's Filecoin address
   - Actor ID: The actor ID if available
   - Name: Display name (if set)
   - Contact: Contact information (if set)
   - Location: Location identifier (if set)
   - Balance: Current wallet balance
   - Type: Wallet type (UserWallet or SPWallet)

   EXAMPLE OUTPUT:
     Address: f1abc123def456...
     Actor ID: f01234
     Name: Main Storage Wallet
     Contact: admin@example.com  
     Location: US-East
     Balance: 10.5 FIL
     Type: UserWallet

OPTIONS:
   --help, -h  show help
```
{% endcode %}
