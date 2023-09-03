# Create a new access grant from satellite address, API key, and passphrase.

{% code fullWidth="true" %}
```
NAME:
   singularity storage update storj new - Create a new access grant from satellite address, API key, and passphrase.

USAGE:
   singularity storage update storj new [command options] <name|id>

DESCRIPTION:
   --satellite-address
      Satellite address.
      
      Custom satellite address should match the format: `<nodeid>@<address>:<port>`.

      Examples:
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key
      API key.

   --passphrase
      Encryption passphrase.
      
      To access existing objects enter passphrase used for uploading.


OPTIONS:
   --satellite-address value  Satellite address. (default: "us1.storj.io") [$SATELLITE_ADDRESS]
   --api-key value            API key. [$API_KEY]
   --passphrase value         Encryption passphrase. [$PASSPHRASE]
   --help, -h                 show help
```
{% endcode %}
