# Akamai NetStorage

{% code fullWidth="true" %}
```
名称:
   singularity storage update netstorage - Akamai NetStorage

使用方法:
   singularity storage update netstorage [命令选项] <名称|ID>

说明:
   --protocol
      选择使用HTTP还是HTTPS协议。
      
      大多数用户应选择HTTPS，默认为此选项。
      HTTP主要用于调试目的。

      示例:
         | http  | 使用HTTP协议
         | https | 使用HTTPS协议

   --host
      连接到NetStorage主机的域名和路径。
      
      格式应为“<域名>/<内部文件夹>”

   --account
      设置NetStorage帐户名称

   --secret
      设置NetStorage帐户的密钥/ G2O密钥以进行身份验证。
      
      请选择“y”选项以设置您自己的密码，然后输入您的密钥。


选项:
   --account value  设置NetStorage帐户名称 [$ACCOUNT]
   --help, -h       显示帮助
   --host value     连接到NetStorage主机的域名和路径。 [$HOST]
   --secret value   设置NetStorage帐户的密钥/ G2O密钥以进行身份验证。 [$SECRET]

   高级选项

   --protocol value  选择使用HTTP还是HTTPS协议。 (默认: "https") [$PROTOCOL]

```
{% endcode %}