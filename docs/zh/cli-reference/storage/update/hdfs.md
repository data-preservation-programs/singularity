# Hadoop分布式文件系统

{% code fullWidth="true" %}
```
NAME:
   singularity storage update hdfs - Hadoop分布式文件系统

USAGE:
   singularity storage update hdfs [命令选项] <名称|ID>

DESCRIPTION:
   --namenode
      Hadoop名称节点和端口。
      
      如："namenode:8020" 连接到主机namenode的8020端口。

   --username
      Hadoop用户名。

      示例：
         | root | 以root身份连接到hdfs。

   --service-principal-name
      Kerberos服务主体名称用于名称节点。
      
      启用KERBEROS身份验证。指定名称节点的服务主体名称（SERVICE/FQDN）。
      例如，对于以'hdfs'服务和FQDN 'namenode.hadoop.docker'运行的名称节点，使用"hdfs/namenode.hadoop.docker"。

   --data-transfer-protection
      Kerberos数据传输保护：authentication|integrity|privacy。
      
      指定在与数据节点通信时是否需要身份验证、数据签名完整性检查和传输加密。
      可能的值为'authentication'、'integrity'和'privacy'。仅在启用KERBEROS时使用。

      示例：
         | privacy | 确保启用身份验证、完整性和加密。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。


OPTIONS:
   --help, -h        显示帮助信息
   --namenode value  Hadoop名称节点和端口。[$NAMENODE]
   --username value  Hadoop用户名。[$USERNAME]

   高级选项

   --data-transfer-protection value  Kerberos数据传输保护：authentication|integrity|privacy。[$DATA_TRANSFER_PROTECTION]
   --encoding value                  后端的编码。(默认值: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot")[$ENCODING]
   --service-principal-name value    Kerberos服务主体名称用于名称节点。[$SERVICE_PRINCIPAL_NAME]

```
{% endcode %}