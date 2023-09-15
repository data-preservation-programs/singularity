# SMB / CIFS

{% code fullWidth="true" %}
```
NAME:
   singularity storage create smb - SMB / CIFS

USAGE:
   singularity storage create smb [command options] [arguments...]

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

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

```
{% endcode %}
