# Hadoop distributed file system

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add hdfs - Hadoop distributed file system

USAGE:
   singularity datasource add hdfs [command options] <dataset_name> <source_path>

DESCRIPTION:
   --hdfs-service-principal-name
      Kerberos service principal name for the namenode.
      
      Enables KERBEROS authentication. Specifies the Service Principal Name
      (SERVICE/FQDN) for the namenode. E.g. \"hdfs/namenode.hadoop.docker\"
      for namenode running as service 'hdfs' with FQDN 'namenode.hadoop.docker'.

   --hdfs-data-transfer-protection
      Kerberos data transfer protection: authentication|integrity|privacy.
      
      Specifies whether or not authentication, data signature integrity
      checks, and wire encryption is required when communicating the the
      datanodes. Possible values are 'authentication', 'integrity' and
      'privacy'. Used only with KERBEROS enabled.

      Examples:
         | privacy | Ensure authentication, integrity and encryption enabled.

   --hdfs-encoding
      The encoding for the backend.
      
      See the [encoding section in the overview](/overview/#encoding) for more info.

   --hdfs-namenode
      Hadoop name node and port.
      
      E.g. "namenode:8020" to connect to host namenode at port 8020.

   --hdfs-username
      Hadoop user name.

      Examples:
         | root | Connect to hdfs as root.


OPTIONS:
   --help, -h  show help

   Data Preparation Options

   --delete-after-export    [Dangerous] Delete the files of the dataset after exporting it to CAR files.  (default: false)
   --rescan-interval value  Automatically rescan the source directory when this interval has passed from last successful scan (default: disabled)

   Options for hdfs

   --hdfs-data-transfer-protection value  Kerberos data transfer protection: authentication|integrity|privacy. [$HDFS_DATA_TRANSFER_PROTECTION]
   --hdfs-encoding value                  The encoding for the backend. (default: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]
   --hdfs-namenode value                  Hadoop name node and port. [$HDFS_NAMENODE]
   --hdfs-service-principal-name value    Kerberos service principal name for the namenode. [$HDFS_SERVICE_PRINCIPAL_NAME]
   --hdfs-username value                  Hadoop user name. [$HDFS_USERNAME]

```
{% endcode %}
