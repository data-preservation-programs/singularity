# Enterprise File Fabric

{% code fullWidth="true" %}
```
NAME:
   singularity storage update filefabric - Enterprise File Fabric

USAGE:
   singularity storage update filefabric [command options] <name|id>

DESCRIPTION:
   --url
      URL of the Enterprise File Fabric to connect to.

      Examples:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | Connect to your Enterprise File Fabric

   --root-folder-id
      ID of the root folder.
      
      Leave blank normally.
      
      Fill in to make rclone start with directory of a given ID.
      

   --permanent-token
      Permanent Authentication Token.
      
      A Permanent Authentication Token can be created in the Enterprise File
      Fabric, on the users Dashboard under Security, there is an entry
      you'll see called "My Authentication Tokens". Click the Manage button
      to create one.
      
      These tokens are normally valid for several years.
      
      For more info see: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      Session Token.
      
      This is a session token which rclone caches in the config file. It is
      usually valid for 1 hour.
      
      Don't set this value - rclone will set it automatically.
      

   --token-expiry
      Token expiry time.
      
      Don't set this value - rclone will set it automatically.
      

   --version
      Version read from the file fabric.
      
      Don't set this value - rclone will set it automatically.
      

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h               show help
   --permanent-token value  Permanent Authentication Token. [$PERMANENT_TOKEN]
   --root-folder-id value   ID of the root folder. [$ROOT_FOLDER_ID]
   --url value              URL of the Enterprise File Fabric to connect to. [$URL]

   Advanced

   --encoding value      The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         Session Token. [$TOKEN]
   --token-expiry value  Token expiry time. [$TOKEN_EXPIRY]
   --version value       Version read from the file fabric. [$VERSION]

   HTTP Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

```
{% endcode %}
