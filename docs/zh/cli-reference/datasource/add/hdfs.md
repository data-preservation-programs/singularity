# Hadoop分布式文件系统

{% code fullWidth="true" %}
```
命令名称:
   singularity datasource add hdfs - Hadoop分布式文件系统

使用方法:
   singularity datasource add hdfs [命令选项] <数据集名称> <源路径>

描述:
   --hdfs-data-transfer-protection
      Kerberos数据传输保护：认证|完整性|隐私。
      
      指定在与数据节点通信时是否需要身份验证、数据签名完整性检查和传输加密。可选择的值为'authentication'，'integrity'和'privacy'。仅在启用KERBEROS时使用。

      示例:
         | privacy | 确保启用身份验证、完整性和加密。

   --hdfs-encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --hdfs-namenode
      Hadoop名称节点和端口。
      
      例如：连接到主机“namenode”端口“8020”，可以写作“namenode:8020”。

   --hdfs-service-principal-name
      名称节点的Kerberos服务主体名称。
      
      启用KERBEROS身份验证。指定名称节点的服务主体名称（SERVICE/FQDN）。例如，对于运行为“hdfs”服务，具有FQDN“namenode.hadoop.docker”的名称节点，可以写作“hdfs/namenode.hadoop.docker”。

   --hdfs-username
      Hadoop用户名。

      示例:
         | root | 以root身份连接到hdfs。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集为CAR文件后，删除数据集文件。 (默认值: false)
   --rescan-interval value  自动重新扫描源目录，当达到上次成功扫描时间间隔时 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 准备就绪)

   HDFS选项

   --hdfs-data-transfer-protection value  Kerberos数据传输保护：认证|完整性|隐私。[$HDFS_DATA_TRANSFER_PROTECTION]
   --hdfs-encoding value                  后端的编码方式。 (默认值: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]
   --hdfs-namenode value                  Hadoop名称节点和端口。[$HDFS_NAMENODE]
   --hdfs-service-principal-name value    名称节点的Kerberos服务主体名称。[$HDFS_SERVICE_PRINCIPAL_NAME]
   --hdfs-username value                  Hadoop用户名。 [$HDFS_USERNAME]

```
{% endcode %}