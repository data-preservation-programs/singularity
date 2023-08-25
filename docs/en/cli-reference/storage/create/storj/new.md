# Create a new access grant from satellite address, API key, and passphrase.

{% code fullWidth="true" %}
```
NAME:
   singularity storage create storj new - Create a new access grant from satellite address, API key, and passphrase.

USAGE:
   singularity storage create storj new [command options] <name> <path>

DESCRIPTION:
   --satellite_address
      Satellite address.
      
      Custom satellite address should match the format: `<nodeid>@<address>:<port>`.

      Examples:
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api_key
      API key.

   --passphrase
      Encryption passphrase.
      
      To access existing objects enter passphrase used for uploading.


OPTIONS:
   --satellite_address value  Satellite address. (default: "us1.storj.io") [$SATELLITE_ADDRESS]
   --api_key value            API key. [$API_KEY]
   --passphrase value         Encryption passphrase. [$PASSPHRASE]
   --help, -h                 show help
```
{% endcode %}
