# Sugarsync

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sugarsync - Sugarsync

USAGE:
   singularity storage update sugarsync [command options] <name|id>

DESCRIPTION:
   --app-id
      Sugarsync App ID.
      
      Leave blank to use rclone's.

   --access-key-id
      Sugarsync Access Key ID.
      
      Leave blank to use rclone's.

   --private-access-key
      Sugarsync Private Access Key.
      
      Leave blank to use rclone's.

   --hard-delete
      Permanently delete files if true
      otherwise put them in the deleted files.

   --refresh-token
      Sugarsync refresh token.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization
      Sugarsync authorization.
      
      Leave blank normally, will be auto configured by rclone.

   --authorization-expiry
      Sugarsync authorization expiry.
      
      Leave blank normally, will be auto configured by rclone.

   --user
      Sugarsync user.
      
      Leave blank normally, will be auto configured by rclone.

   --root-id
      Sugarsync root id.
      
      Leave blank normally, will be auto configured by rclone.

   --deleted-id
      Sugarsync deleted folder id.
      
      Leave blank normally, will be auto configured by rclone.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --description
      Description of the remote.


OPTIONS:
   --access-key-id value       Sugarsync Access Key ID. [$ACCESS_KEY_ID]
   --app-id value              Sugarsync App ID. [$APP_ID]
   --hard-delete               Permanently delete files if true (default: false) [$HARD_DELETE]
   --help, -h                  show help
   --private-access-key value  Sugarsync Private Access Key. [$PRIVATE_ACCESS_KEY]

   Advanced

   --authorization value         Sugarsync authorization. [$AUTHORIZATION]
   --authorization-expiry value  Sugarsync authorization expiry. [$AUTHORIZATION_EXPIRY]
   --deleted-id value            Sugarsync deleted folder id. [$DELETED_ID]
   --description value           Description of the remote. [$DESCRIPTION]
   --encoding value              The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --refresh-token value         Sugarsync refresh token. [$REFRESH_TOKEN]
   --root-id value               Sugarsync root id. [$ROOT_ID]
   --user value                  Sugarsync user. [$USER]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone default)

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
