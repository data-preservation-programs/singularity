# SMB / CIFS

```
NAME:
   singularity datasource add smb - SMB / CIFS

USAGE:
   singularity datasource add smb [command options] <dataset_name> <source_path>

DESCRIPTION:
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
      

   --smb-hide-special-share
      Hide special shares (e.g. print$) which users aren't supposed to access.

   --smb-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --smb-host
      SMB server hostname to connect to.
      
      E.g. "example.com".

   --smb-port
      SMB port number.

   --smb-idle-timeout
      Max time before closing idle connections.
      
      If no connections have been returned to the connection pool in the time
      given, rclone will empty the connection pool.
      
      Set to 0 to keep connections indefinitely.
      

   --smb-case-insensitive
      Whether the server is configured to be case-insensitive.
      
      Always true on Windows shares.

   --smb-user
      SMB username.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for smb

   --smb-case-insensitive value    Whether the server is configured to be case-insensitive. (default: "true") [$SMB_CASE_INSENSITIVE]
   --smb-domain value              Domain name for NTLM authentication. (default: "WORKGROUP") [$SMB_DOMAIN]
   --smb-encoding value            The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
   --smb-hide-special-share value  Hide special shares (e.g. print$) which users aren't supposed to access. (default: "true") [$SMB_HIDE_SPECIAL_SHARE]
   --smb-host value                SMB server hostname to connect to. [$SMB_HOST]
   --smb-idle-timeout value        Max time before closing idle connections. (default: "1m0s") [$SMB_IDLE_TIMEOUT]
   --smb-pass value                SMB password. [$SMB_PASS]
   --smb-port value                SMB port number. (default: "445") [$SMB_PORT]
   --smb-spn value                 Service principal name. [$SMB_SPN]
   --smb-user value                SMB username. (default: "shane") [$SMB_USER]

```
