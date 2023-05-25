# SMB / CIFS

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add smb - SMB / CIFS

USAGE:
   singularity datasource add smb [command options] <dataset_name> <source_path>

DESCRIPTION:
   --smb-case-insensitive
      Whether the server is configured to be case-insensitive.

      Always true on Windows shares.

   --smb-encoding
      The encoding for the backend.

      See the [encoding section in the overview](/overview/#encoding) for more info.

   --smb-host
      SMB server hostname to connect to.

      E.g. "example.com".

   --smb-pass
      SMB password.

   --smb-domain
      Domain name for NTLM authentication.

   --smb-spn
      Service principal name.

      Rclone presents this name to the server. Some servers use this as further
      authentication, and it often needs to be set for clusters. For example:

          cifs/remotehost:1020

      Leave blank if not sure.


   --smb-idle-timeout
      Max time before closing idle connections.

      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.

      Set to 0 to keep connections indefinitely.


   --smb-hide-special-share
      Hide special shares (e.g. print$) which users aren't supposed to access.

   --smb-user
      SMB username.

   --smb-port
      SMB port number.


OPTIONS:
   --help, -h          show help
   --smb-domain value  Domain name for NTLM authentication. (default: "WORKGROUP") [$SMB_DOMAIN]
   --smb-host value    SMB server hostname to connect to. [$SMB_HOST]
   --smb-pass value    SMB password. [$SMB_PASS]
   --smb-port value    SMB port number. (default: "445") [$SMB_PORT]
   --smb-spn value     Service principal name. [$SMB_SPN]
   --smb-user value    SMB username. (default: "shane") [$SMB_USER]

   Advanced Options

   --smb-case-insensitive value    Whether the server is configured to be case-insensitive. (default: "true") [$SMB_CASE_INSENSITIVE]
   --smb-encoding value            The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
   --smb-hide-special-share value  Hide special shares (e.g. print$) which users aren't supposed to access. (default: "true") [$SMB_HIDE_SPECIAL_SHARE]
   --smb-idle-timeout value        Max time before closing idle connections. (default: "1m0s") [$SMB_IDLE_TIMEOUT]

   Data Preparation Options

   --delete-after-export  [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)

```
{% endcode %}
