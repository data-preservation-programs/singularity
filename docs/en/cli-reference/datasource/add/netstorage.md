# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add netstorage - Akamai NetStorage

USAGE:
   singularity datasource add netstorage [command options] <dataset_name> <source_path>

DESCRIPTION:
   --netstorage-secret
      Set the NetStorage account secret/G2O key for authentication.
      
      Please choose the 'y' option to set your own password then enter your secret.

   --netstorage-protocol
      Select between HTTP or HTTPS protocol.
      
      Most users should choose HTTPS, which is the default.
      HTTP is provided primarily for debugging purposes.

      Examples:
         | http  | HTTP protocol
         | https | HTTPS protocol

   --netstorage-host
      Domain+path of NetStorage host to connect to.
      
      Format should be `<domain>/<internal folders>`

   --netstorage-account
      Set the NetStorage account name


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for netstorage

   --netstorage-account value   Set the NetStorage account name [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      Domain+path of NetStorage host to connect to. [$NETSTORAGE_HOST]
   --netstorage-protocol value  Select between HTTP or HTTPS protocol. (default: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    Set the NetStorage account secret/G2O key for authentication. [$NETSTORAGE_SECRET]

```
{% endcode %}
