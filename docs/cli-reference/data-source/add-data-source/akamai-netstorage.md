# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add netstorage - Akamai NetStorage

USAGE:
   singularity datasource add netstorage [command options] <dataset_name> <source_path>

DESCRIPTION:
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

   --netstorage-secret
      Set the NetStorage account secret/G2O key for authentication.

      Please choose the 'y' option to set your own password then enter your secret.


OPTIONS:
   --help, -h                  show help
   --netstorage-account value  Set the NetStorage account name [$NETSTORAGE_ACCOUNT]
   --netstorage-host value     Domain+path of NetStorage host to connect to. [$NETSTORAGE_HOST]
   --netstorage-secret value   Set the NetStorage account secret/G2O key for authentication. [$NETSTORAGE_SECRET]

   Advanced Options

   --netstorage-protocol value  Select between HTTP or HTTPS protocol. (default: "https") [$NETSTORAGE_PROTOCOL]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
