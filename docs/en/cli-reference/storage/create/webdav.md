# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularity storage create webdav - WebDAV

USAGE:
   singularity storage create webdav [command options]

DESCRIPTION:
   --url
      URL of http host to connect to.
      
      E.g. https://example.com.

   --vendor
      Name of the WebDAV site/service/software you are using.

      Examples:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, authenticated by Microsoft account
         | sharepoint-ntlm | Sharepoint with NTLM authentication, usually self-hosted or on-premises
         | other           | Other site/service or software

   --user
      User name.
      
      In case NTLM authentication is used, the username should be in the format 'Domain\User'.

   --pass
      Password.

   --bearer-token
      Bearer token instead of user/pass (e.g. a Macaroon).

   --bearer-token-command
      Command to run to get a bearer token.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.
      
      Default encoding is Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8 for sharepoint-ntlm or identity otherwise.

   --headers
      Set HTTP headers for all transactions.
      
      Use this to set additional HTTP headers for all transactions
      
      The input format is comma separated list of key,value pairs.  Standard
      [CSV encoding](https://godoc.org/encoding/csv) may be used.
      
      For example, to set a Cookie use 'Cookie,name=value', or '"Cookie","name=value"'.
      
      You can set multiple headers, e.g. '"Cookie","name=value","Authorization","xxx"'.
      


OPTIONS:
   --bearer-token value  Bearer token instead of user/pass (e.g. a Macaroon). [$BEARER_TOKEN]
   --help, -h            show help
   --pass value          Password. [$PASS]
   --url value           URL of http host to connect to. [$URL]
   --user value          User name. [$USER]
   --vendor value        Name of the WebDAV site/service/software you are using. [$VENDOR]

   Advanced

   --bearer-token-command value  Command to run to get a bearer token. [$BEARER_TOKEN_COMMAND]
   --encoding value              The encoding for the backend. [$ENCODING]
   --headers value               Set HTTP headers for all transactions. [$HEADERS]

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
