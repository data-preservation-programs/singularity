# Hadoop分布式文件系统

{% code fullWidth="true" %}
```
名称：
   singularity storage create hdfs - Hadoop分布式文件系统

用法：
   singularity storage create hdfs [命令选项] [参数...]

说明：
   --namenode
      Hadoop名称节点和端口。
      
      例如: "namenode:8020" 连接到主机名namenode的端口8020。

   --username
      Hadoop用户名。

      示例:
         | root | 作为root连接到hdfs。

   --service-principal-name
      Namenode的Kerberos服务主体名称。
      
      启用KERBEROS身份验证。指定namenode的服务主体名称(SERVICE/FQDN)。
      例如，对于作为服务 'hdfs' 运行且具有FQDN 'namenode.hadoop.docker' 的namenode，
      可以指定 "hdfs/namenode.hadoop.docker"。

   --data-transfer-protection
      Kerberos数据传输保护：authentication|integrity|privacy。
      
      指定在与数据节点通信时是否需要身份验证、数据签名完整性检查和数据传输加密。
      可能的值是 "authentication"、"integrity" 和 "privacy"。仅在启用KERBEROS时使用。

      示例:
         | privacy | 确保启用身份验证、完整性和加密。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

选项：
   --help, -h        显示帮助信息
   --namenode value  Hadoop名称节点和端口。[$NAMENODE]
   --username value  Hadoop用户名。[$USERNAME]

   高级选项

   --data-transfer-protection value  Kerberos数据传输保护：authentication|integrity|privacy。[$DATA_TRANSFER_PROTECTION]
   --encoding value                  后端的编码方式。（默认值: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --service-principal-name value    Namenode的Kerberos服务主体名称。[$SERVICE_PRINCIPAL_NAME]

   通用选项

   --name value  存储的名称（默认值:自动生成）
   --path value  存储的路径

```
{% endcode %}