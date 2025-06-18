# Create a new wallet

{% code fullWidth="true" %}
```
NAME:
   singularity wallet create - Create a new wallet

USAGE:
   singularity wallet create [command options] [type]

DESCRIPTION:
   Create a new Filecoin wallet or storage provider contact entry.

   The command automatically detects the wallet type based on provided arguments:
   - For UserWallet: Creates a wallet with offline keypair generation
   - For SPWallet: Creates a contact entry for a storage provider

   SUPPORTED KEY TYPES (for UserWallet):
     secp256k1    ECDSA using the secp256k1 curve (default, most common)
     bls          BLS signature scheme (Boneh-Lynn-Shacham)

   EXAMPLES:
     # Create a secp256k1 UserWallet (default)
     singularity wallet create

     # Create a secp256k1 UserWallet explicitly
     singularity wallet create secp256k1

     # Create a BLS UserWallet
     singularity wallet create bls

     # Create an SPWallet contact entry
     singularity wallet create --address f3abc123... --actor-id f01234 --name "Example SP"

   The newly created wallet address and other details will be displayed upon successful creation.

OPTIONS:
   --address value   Storage provider wallet address (creates SPWallet contact)
   --actor-id value  Storage provider actor ID (e.g., f01234)
   --name value      Optional display name
   --contact value   Optional contact information
   --location value  Optional provider location
   --help, -h        show help
```
{% endcode %}
