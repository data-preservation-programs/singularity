# Akamai NetStorage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create netstorage - Akamai NetStorage

USAGE:
   singularity storage create netstorage [command options] [arguments...]

DESCRIPTION:
   --protocol
      Select between HTTP or HTTPS protocol.
      
      Most users should choose HTTPS, which is the default.
      HTTP is provided primarily for debugging purposes.

      Examples:
         | http  | HTTP protocol
         | https | HTTPS protocol

   --host
      Domain+path of NetStorage host to connect to.
      
      Format should be `<domain>/<internal folders>`

   --account
      Set the NetStorage account name

   --secret
      Set the NetStorage account secret/G2O key for authentication.
      
      Please choose the 'y' option to set your own password then enter your secret.


OPTIONS:
   --account value  Set the NetStorage account name [$ACCOUNT]
   --help, -h       show help
   --host value     Domain+path of NetStorage host to connect to. [$HOST]
   --secret value   Set the NetStorage account secret/G2O key for authentication. [$SECRET]

   Advanced

   --protocol value  Select between HTTP or HTTPS protocol. (default: "https") [$PROTOCOL]

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

```
{% endcode %}
