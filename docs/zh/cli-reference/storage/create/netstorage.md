# Akamai NetStorage

{% code fullWidth="true" %}
```
名称：
   singularity storage create netstorage - Akamai NetStorage

使用：
   singularity storage create netstorage [命令选项] [参数...]

描述：
   --protocol
      选择使用 HTTP 或者 HTTPS 协议。
      
      大多数用户应该选择 HTTPS，这是默认选项。
      HTTP 主要用于调试目的。

      示例：
         | http  | 使用 HTTP 协议
         | https | 使用 HTTPS 协议

   --host
      NetStorage 主机的域名和路径。
      
      格式应为 `<域名>/<内部文件夹>`

   --account
      设置 NetStorage 帐户名

   --secret
      设置 NetStorage 帐户密钥/G2O 密钥进行身份验证。
      
      请选择“是”以设置自己的密码，然后输入您的密钥。

选项：
   --account value  设置 NetStorage 帐户名 [$ACCOUNT]
   --help, -h       显示帮助信息
   --host value     NetStorage 主机的域名和路径 [$HOST]
   --secret value   设置 NetStorage 帐户密钥/G2O 密钥进行身份验证 [$SECRET]

   高级选项

   --protocol value  选择使用 HTTP 或者 HTTPS 协议。 (默认值: "https") [$PROTOCOL]

   通用选项

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}