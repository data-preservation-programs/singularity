# Wallet management


Singularity provides comprehensive wallet management capabilities for Filecoin operations. You can create new wallets, import existing ones, and manage wallet metadata for better organization.

## Wallet Metadata Support

Wallets support optional metadata fields for enhanced organization:
- **Name**: Human-readable display name
- **Contact**: Contact information (email, etc.)  
- **Location**: Geographic location or region

These fields help identify and organize wallets, especially useful for storage providers managing multiple wallets across different regions.


Singularity provides comprehensive wallet management capabilities for Filecoin operations. You can create new wallets, import existing ones, and check balances.

## Balance Information

The wallet balance command provides real-time FIL and FIL+ datacap balance information:
- **FIL Balance**: Current FIL balance in human-readable format
- **Raw Balance**: Precise balance in attoFIL for calculations
- **FIL+ DataCap**: Verified client datacap allowance in GiB
- **Raw DataCap**: Precise datacap in bytes

{% code fullWidth="true" %}
```
NAME:
   singularity wallet - Wallet management

USAGE:
   singularity wallet command [command options]

COMMANDS:
   balance  Get wallet balance information
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

### Examples

**Check wallet balance:**
```bash
# Get balance for a specific wallet
singularity wallet balance f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz

# Get balance in JSON format
singularity --json wallet balance f1abcde7zd3lfsv43aj2kb454ymaqw7debhumjxyz
```

**Create and manage wallets:**
```bash
# Create a new secp256k1 wallet (default)
singularity wallet create

# Create a BLS wallet
singularity wallet create bls

# List all wallets
singularity wallet list

# Import a wallet from private key
singularity wallet import

# Update wallet metadata
singularity wallet update f1abc123... --name "My Wallet" --contact "user@example.com"
```
