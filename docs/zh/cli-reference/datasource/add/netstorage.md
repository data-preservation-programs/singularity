# Akamai NetStorage

{% code fullWidth="true" %}
```
名称：
   singularity datasource add netstorage - Akamai NetStorage

使用方法：
   singularity datasource add netstorage [命令选项] <数据集名称> <源路径>

说明：
   --netstorage-account
      设置NetStorage帐户名称

   --netstorage-host
      连接到NetStorage主机的域名+路径。
      
      格式应为“<域名>/<内部文件夹>”

   --netstorage-protocol
      选择HTTP或HTTPS协议。
      
      大多数用户应选择HTTPS，这是默认值。
      HTTP主要用于调试目的。

      示例：
         | http  | HTTP协议
         | https | HTTPS协议

   --netstorage-secret
      设置用于身份验证的NetStorage帐户密码/密钥。
      
      请选"是"选项设置您自己的密码，然后输入您的密钥。


选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出数据集为CAR文件后删除数据集的文件。  (默认值: false)
   --rescan-interval value  上次成功扫描后自动重新扫描源目录的时间间隔 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 准备就绪)

   NetStorage选项

   --netstorage-account value   设置NetStorage帐户名称 [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      连接到NetStorage主机的域名+路径。 [$NETSTORAGE_HOST]
   --netstorage-protocol value  选择HTTP或HTTPS协议。 (默认值: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    设置用于身份验证的NetStorage帐户密码/密钥。 [$NETSTORAGE_SECRET]

```
{% endcode %}