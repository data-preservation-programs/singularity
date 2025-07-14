# Wallet management

Singularity provides comprehensive wallet management capabilities for Filecoin operations. You can create new wallets, import existing ones, and manage wallet metadata for better organization.

## Wallet Metadata Support

Wallets support optional metadata fields for enhanced organization:
- **Name**: Human-readable display name
- **Contact**: Contact information (email, etc.)  
- **Location**: Geographic location or region

These fields help identify and organize wallets, especially useful for storage providers managing multiple wallets across different regions.

{% code fullWidth="true" %}
```
NAME:
   singularity wallet - Wallet management

USAGE:
   singularity wallet command [command options]

COMMANDS:
   create   Create a new wallet
   import   Import a wallet from exported private key
   init     Initialize a wallet
   list     List all imported wallets
   remove   Remove a wallet
   update   Update wallet details
   help, h  Shows a list of commands or help for one command

OPTIONS:
   --help, -h  show help
```
{% endcode %}
