# Box

{% code fullWidth="true" %}
```
NAME:
   singularity storage create box - Box

USAGE:
   singularity storage create box [command options] [arguments...]

DESCRIPTION:
   --client-id
      OAuth Client Id.
      
      Leave blank normally.

   --client-secret
      OAuth Client Secret.
      
      Leave blank normally.

   --token
      OAuth Access Token as a JSON blob.

   --auth-url
      Auth server URL.
      
      Leave blank to use the provider defaults.

   --token-url
      Token server url.
      
      Leave blank to use the provider defaults.

   --root-folder-id
      Fill in for rclone to use a non root folder as its starting point.

   --box-config-file
      Box App config.json location
      
      Leave blank normally.
      
      Leading `~` will be expanded in the file name as will environment variables such as `${RCLONE_CONFIG_DIR}`.

   --access-token
      Box App Primary Access Token
      
      Leave blank normally.

   --box-sub-type
      

      Examples:
         | user       | Rclone should act on behalf of a user.
         | enterprise | Rclone should act on behalf of a service account.

   --upload-cutoff
      Cutoff for switching to multipart upload (>= 50 MiB).

   --commit-retries
      Max number of times to try committing a multipart file.

   --list-chunk
      Size of listing chunk 1-1000.

   --owned-by
      Only show items owned by the login (email address) passed in.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --access-token value     Box App Primary Access Token [$ACCESS_TOKEN]
   --box-config-file value  Box App config.json location [$BOX_CONFIG_FILE]
   --box-sub-type value     (default: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuth Client Id. [$CLIENT_ID]
   --client-secret value    OAuth Client Secret. [$CLIENT_SECRET]
   --help, -h               show help

   Advanced

   --auth-url value        Auth server URL. [$AUTH_URL]
   --commit-retries value  Max number of times to try committing a multipart file. (default: 100) [$COMMIT_RETRIES]
   --encoding value        The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --list-chunk value      Size of listing chunk 1-1000. (default: 1000) [$LIST_CHUNK]
   --owned-by value        Only show items owned by the login (email address) passed in. [$OWNED_BY]
   --root-folder-id value  Fill in for rclone to use a non root folder as its starting point. (default: "0") [$ROOT_FOLDER_ID]
   --token value           OAuth Access Token as a JSON blob. [$TOKEN]
   --token-url value       Token server url. [$TOKEN_URL]
   --upload-cutoff value   Cutoff for switching to multipart upload (>= 50 MiB). (default: "50Mi") [$UPLOAD_CUTOFF]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value)
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string (default: rclone/v1.62.2-DEV)

   General

   --name value  Name of the storage (default: Auto generated)
   --path value  Path of the storage

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

```
{% endcode %}
