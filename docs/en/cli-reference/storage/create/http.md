# HTTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage create http - HTTP

USAGE:
   singularity storage create http [command options]

DESCRIPTION:
   --url
      URL of HTTP host to connect to.
      
      E.g. "https://example.com", or "https://user:pass@example.com" to use a username and password.

   --headers
      Set HTTP headers for all transactions.
      
      Use this to set additional HTTP headers for all transactions.
      
      The input format is comma separated list of key,value pairs.  Standard
      [CSV encoding](https://godoc.org/encoding/csv) may be used.
      
      For example, to set a Cookie use 'Cookie,name=value', or '"Cookie","name=value"'.
      
      You can set multiple headers, e.g. '"Cookie","name=value","Authorization","xxx"'.

   --no-slash
      Set this if the site doesn't end directories with /.
      
      Use this if your target website does not use / on the end of
      directories.
      
      A / on the end of a path is how rclone normally tells the difference
      between files and directories.  If this flag is set, then rclone will
      treat all files with Content-Type: text/html as directories and read
      URLs from them rather than downloading them.
      
      Note that this may cause rclone to confuse genuine HTML files with
      directories.

   --no-head
      Don't use HEAD requests.
      
      HEAD requests are mainly used to find file sizes in dir listing.
      If your site is being very slow to load then you can try this option.
      Normally rclone does a HEAD request for each potential file in a
      directory listing to:
      
      - find its size
      - check it really exists
      - check to see if it is a directory
      
      If you set this option, rclone will not do the HEAD request. This will mean
      that directory listings are much quicker, but rclone won't have the times or
      sizes of any files, and some files that don't exist may be in the listing.


OPTIONS:
   --help, -h   show help
   --url value  URL of HTTP host to connect to. [$URL]

   Advanced

   --headers value  Set HTTP headers for all transactions. [$HEADERS]
   --no-head        Don't use HEAD requests. (default: false) [$NO_HEAD]
   --no-slash       Set this if the site doesn't end directories with /. (default: false) [$NO_SLASH]

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
