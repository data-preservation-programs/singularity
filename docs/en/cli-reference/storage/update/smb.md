# SMB / CIFS

{% code fullWidth="true" %}
```
NAME:
   singularity storage update smb - SMB / CIFS

USAGE:
   singularity storage update smb [command options] <name|id>

DESCRIPTION:
   --host
      SMB server hostname to connect to.
      
      E.g. "example.com".

   --user
      SMB username.

   --port
      SMB port number.

   --pass
      SMB password.

   --domain
      Domain name for NTLM authentication.

   --spn
      Service principal name.
      
      Rclone presents this name to the server. Some servers use this as further
      authentication, and it often needs to be set for clusters. For example:
      
          cifs/remotehost:1020
      
      Leave blank if not sure.
      

   --idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --hide-special-share
      Hide special shares (e.g. print$) which users aren't supposed to access.

   --case-insensitive
      Whether the server is configured to be case-insensitive.
      
      Always true on Windows shares.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --domain value  Domain name for NTLM authentication. (default: "WORKGROUP") [$DOMAIN]
   --help, -h      show help
   --host value    SMB server hostname to connect to. [$HOST]
   --pass value    SMB password. [$PASS]
   --port value    SMB port number. (default: 445) [$PORT]
   --spn value     Service principal name. [$SPN]
   --user value    SMB username. (default: "$USER") [$USER]

   Advanced

   --case-insensitive    Whether the server is configured to be case-insensitive. (default: true) [$CASE_INSENSITIVE]
   --encoding value      The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --hide-special-share  Hide special shares (e.g. print$) which users aren't supposed to access. (default: true) [$HIDE_SPECIAL_SHARE]
   --idle-timeout value  Max time before closing idle connections. (default: "1m0s") [$IDLE_TIMEOUT]

```
{% endcode %}
