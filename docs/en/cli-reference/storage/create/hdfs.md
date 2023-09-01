# Hadoop distributed file system

{% code fullWidth="true" %}
```
NAME:
   singularity storage create hdfs - Hadoop distributed file system

USAGE:
   singularity storage create hdfs [command options] <name> <path>

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

```
{% endcode %}
