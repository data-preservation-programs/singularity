# Akamai NetStorage

{% code fullWidth="true" %}
```
名称：
   singularity datasource add netstorage - Akamai NetStorage

用法：
   singularity datasource add netstorage [command options] <dataset_name> <source_path>

介绍：
   --netstorage-secret
      设置 NetStorage 帐户的密钥/G2O 密钥进行身份验证。
      
      请选择“y”选项，然后输入您的密钥以设置自己的密码。

   --netstorage-protocol
      选择 HTTP 或 HTTPS 协议。
      
      大多数用户应选择默认的 HTTPS。
      HTTP 主要提供用于调试的功能。

      示例：
         | http  | HTTP 协议
         | https | HTTPS 协议

   --netstorage-host
      NetStorage 主机的域名 + 路径。
      
      格式应为 `<domain>/<internal folders>`

   --netstorage-account
      设置 NetStorage 帐户名


选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 将数据集导出到 CAR 文件后删除数据集的文件。（默认值：false）
   --rescan-interval value  当自上次成功扫描后经过此间隔时间时，自动重新扫描源目录（默认值：禁用）

   NetStorage 选项

   --netstorage-account value   设置 NetStorage 帐户名 [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      NetStorage 主机的域名 + 路径。[$NETSTORAGE_HOST]
   --netstorage-protocol value  选择 HTTP 或 HTTPS 协议。（默认值：“https”）[$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    设置 NetStorage 帐户的密钥/G2O 密钥进行身份验证。[$NETSTORAGE_SECRET]

```
{% endcode %}