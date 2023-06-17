# Hadoop 分布式文件系统

{% code fullWidth="true" %}
```
名称:
   singularity 数据源添加 hdfs - Hadoop 分布式文件系统

使用方式:
   singularity datasource add hdfs [命令选项] <数据集名称> <源路径>

描述:
   --hdfs-data-transfer-protection
      Kerberos 数据传输保护: 认证 | 完整性 | 加密。
      
      指定是否需要在与数据节点通信时进行身份验证、数据签名完整性检查以及传输加密。可能的值有'authentication'，'integrity'和'privacy'。仅在启用 KERBEROS 时使用。

      示例:
         | privacy | 确保启用身份验证、完整性和加密。

   --hdfs-encoding
      背景端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --hdfs-namenode
      Hadoop 名称节点和端口。
      
      例如："namenode:8020"，连接到端口8020处的namenode主机。

   --hdfs-username
      Hadoop 用户名。

      示例:
         | root | 以root身份连接到hdfs。

   --hdfs-service-principal-name
      namenode 的 Kerberos 服务主体名称。
      
      启用 KERBEROS 身份验证。指定namenode的 Service Principal Name (SERVICE/FQDN)。例如，对于运行为服务'hdfs'且具有FQDN'name node.hadoop.docker'的namenode，Kerberos 服务主体名称为"hdfs/namenode.hadoop.docker"。


选项:
  --help, -h 显示帮助

  数据准备选项

  --delete-after-export [警告] 将数据集导出为CAR文件后删除数据集文件。  (默认值:false)

  --rescan-interval value 当上一次成功扫描之后经历了此时间间隔时，自动重新扫描源目录。 (默认值:禁用)

  用于 hdfs 的选项

  --hdfs-data-transfer-protection value Kerberos 数据传输保护: authentication|integrity|privacy. [$HDFS_DATA_TRANSFER_PROTECTION]

  --hdfs-encoding value 背景端的编码方式。(默认值:"Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]

  --hdfs-namenode value Hadoop 名称节点和端口 [$HDFS_NAMENODE]

  --hdfs-service-principal-name value namenode 的 Kerberos 服务主体名称。[$HDFS_SERVICE_PRINCIPAL_NAME]

  --hdfs-username value Hadoop 用户名。[$HDFS_USERNAME]

```
{% endcode %}