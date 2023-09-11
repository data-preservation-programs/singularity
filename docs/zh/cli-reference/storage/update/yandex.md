# Yandex Disk

{% code fullWidth="true" %}
```
名称：
   singularity storage update yandex - Yandex Disk

用法：
   singularity storage update yandex [命令选项] <名称|id>

描述：
   --client-id
      OAuth 客户端标识。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌，以 JSON 格式存储。

   --auth-url
      认证服务器 URL。
      
      留空以使用提供商默认值。

   --token-url
      令牌服务器 URL。
      
      留空以使用提供商默认值。

   --hard-delete
      永久删除文件，而不是将其移到回收站。

   --encoding
      后端使用的编码方式。
      
      更多信息请参考[概述中的编码部分](/overview/#encoding)。

选项：
   --client-id value      OAuth 客户端标识。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息

   高级选项：

   --auth-url value   认证服务器 URL。[$AUTH_URL]
   --encoding value   后端使用的编码方式。 (默认值: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete      永久删除文件，而不是将其移到回收站。 (默认值：false) [$HARD_DELETE]
   --token value      OAuth 访问令牌，以 JSON 格式存储。[$TOKEN]
   --token-url value  令牌服务器 URL。[$TOKEN_URL]

```
{% endcode %}