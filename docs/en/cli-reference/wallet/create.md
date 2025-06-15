# Create a new wallet

{% code fullWidth="true" %}
```
NAME:
   singularity wallet create - Create a new wallet

USAGE:
   singularity wallet create [command options]

DESCRIPTION:
   Create a new Filecoin wallet using offline keypair generation.

   The wallet will be stored locally in the Singularity database and can be used for making deals and other operations. The private key is generated securely and stored encrypted.

   SUPPORTED KEY TYPES:
     secp256k1    ECDSA using the secp256k1 curve (default, most common)
     bls          BLS signature scheme (Boneh-Lynn-Shacham)

   EXAMPLES:
     # Create a secp256k1 wallet (default)
     singularity wallet create

     # Create a secp256k1 wallet explicitly
     singularity wallet create secp256k1

     # Create a BLS wallet
     singularity wallet create bls

   The newly created wallet address and other details will be displayed upon successful creation.

OPTIONS:
   --help, -h  show help
```
{% endcode %}
