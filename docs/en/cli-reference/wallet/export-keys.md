# Migrate private keys from database (legacy Actor.PrivateKey) to the filesystem keystore

{% code fullWidth="true" %}
```
NAME:
   singularity wallet export-keys - Migrate private keys from database (legacy Actor.PrivateKey) to the filesystem keystore

USAGE:
   singularity wallet export-keys [command options]

DESCRIPTION:
   Reads private keys stored in the legacy actors table and saves them to
   the filesystem keystore (~/.singularity/keystore or SINGULARITY_KEYSTORE).
   Creates Wallet records for each exported key and links them to the
   corresponding Actor.

   This command is idempotent — actors whose address already has a Wallet
   record are skipped. Keys that fail to parse are reported but do not
   abort the migration.

   After exporting, prompts to drop the orphaned private_key column from
   the actors table. This is irreversible — verify keys are in the keystore
   before confirming. For scripted use, pass --drop-db-keys --i-am-really-sure
   to skip the prompt.

OPTIONS:
   --drop-db-keys      drop the private_key column from the actors table after export (default: false)
   --i-am-really-sure  confirm column drop (required with --drop-db-keys) (default: false)
   --help, -h          show help
```
{% endcode %}
