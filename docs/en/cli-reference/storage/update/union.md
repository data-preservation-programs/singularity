# Union merges the contents of several upstream fs

{% code fullWidth="true" %}
```
NAME:
   singularity storage update union - Union merges the contents of several upstream fs

USAGE:
   singularity storage update union [command options] <name|id>

DESCRIPTION:
   --upstreams
      List of space separated upstreams.
      
      Can be 'upstreama:test/dir upstreamb:', '"upstreama:test/space:ro dir" upstreamb:', etc.

   --action-policy
      Policy to choose upstream on ACTION category.

   --create-policy
      Policy to choose upstream on CREATE category.

   --search-policy
      Policy to choose upstream on SEARCH category.

   --cache-time
      Cache time of usage and free space (in seconds).
      
      This option is only useful when a path preserving policy is used.

   --min-free-space
      Minimum viable free space for lfs/eplfs policies.
      
      If a remote has less than this much free space then it won't be
      considered for use in lfs or eplfs policies.

   --description
      Description of the remote.


OPTIONS:
   --action-policy value  Policy to choose upstream on ACTION category. (default: "epall") [$ACTION_POLICY]
   --cache-time value     Cache time of usage and free space (in seconds). (default: 120) [$CACHE_TIME]
   --create-policy value  Policy to choose upstream on CREATE category. (default: "epmfs") [$CREATE_POLICY]
   --help, -h             show help
   --search-policy value  Policy to choose upstream on SEARCH category. (default: "ff") [$SEARCH_POLICY]
   --upstreams value      List of space separated upstreams. [$UPSTREAMS]

   Advanced

   --description value     Description of the remote. [$DESCRIPTION]
   --min-free-space value  Minimum viable free space for lfs/eplfs policies. (default: "1Gi") [$MIN_FREE_SPACE]

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
