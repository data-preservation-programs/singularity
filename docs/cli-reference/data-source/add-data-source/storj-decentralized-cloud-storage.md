# Storj Decentralized Cloud Storage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add storj - Storj Decentralized Cloud Storage

USAGE:
   singularity datasource add storj [command options] <dataset_name> <source_path>

DESCRIPTION:
   --storj-access-grant
      [Provider] - existing
         Access grant.

   --storj-satellite-address
      [Provider] - new
         Satellite address.

         Custom satellite address should match the format: `<nodeid>@<address>:<port>`.

         Examples:
            | us1.storj.io | US1
            | eu1.storj.io | EU1
            | ap1.storj.io | AP1

   --storj-api-key
      [Provider] - new
         API key.

   --storj-passphrase
      [Provider] - new
         Encryption passphrase.

         To access existing objects enter passphrase used for uploading.

   --storj-provider
      Choose an authentication method.

      Examples:
         | existing | Use an existing access grant.
         | new      | Create a new access grant from satellite address, API key, and passphrase.


OPTIONS:
   --help, -h                       show help
   --storj-access-grant value       Access grant. [$STORJ_ACCESS_GRANT]
   --storj-api-key value            API key. [$STORJ_API_KEY]
   --storj-passphrase value         Encryption passphrase. [$STORJ_PASSPHRASE]
   --storj-provider value           Choose an authentication method. (default: "existing") [$STORJ_PROVIDER]
   --storj-satellite-address value  Satellite address. (default: "us1.storj.io") [$STORJ_SATELLITE_ADDRESS]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
