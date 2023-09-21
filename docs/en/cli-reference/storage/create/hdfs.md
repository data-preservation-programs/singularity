# Hadoop distributed file system

{% code fullWidth="true" %}
```
NAME:
   singularity storage create hdfs - Hadoop distributed file system

USAGE:
   singularity storage create hdfs [command options] [arguments...]

DESCRIPTION:
   --namenode
      Hadoop name node and port.
      
      E.g. "namenode:8020" to connect to host namenode at port 8020.

   --username
      Hadoop user name.

      Examples:
         | root | Connect to hdfs as root.

   --service-principal-name
      Kerberos service principal name for the namenode.
      
      Enables KERBEROS authentication. Specifies the Service Principal Name
      (SERVICE/FQDN) for the namenode. E.g. \"hdfs/namenode.hadoop.docker\"
      for namenode running as service 'hdfs' with FQDN 'namenode.hadoop.docker'.

   --data-transfer-protection
      Kerberos data transfer protection: authentication|integrity|privacy.
      
      Specifies whether or not authentication, data signature integrity
      checks, and wire encryption is required when communicating the the
      datanodes. Possible values are 'authentication', 'integrity' and
      'privacy'. Used only with KERBEROS enabled.

      Examples:
         | privacy | Ensure authentication, integrity and encryption enabled.

   --encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.


OPTIONS:
   --help, -h        show help
   --namenode value  Hadoop name node and port. [$NAMENODE]
   --username value  Hadoop user name. [$USERNAME]

   Advanced

   --data-transfer-protection value  Kerberos data transfer protection: authentication|integrity|privacy. [$DATA_TRANSFER_PROTECTION]
   --encoding value                  The encoding for the backend. (default: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --service-principal-name value    Kerberos service principal name for the namenode. [$SERVICE_PRINCIPAL_NAME]

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
