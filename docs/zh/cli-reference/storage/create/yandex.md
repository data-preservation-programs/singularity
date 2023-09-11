# Yandex Disk

{% code fullWidth="true" %}
```
名称：
   singularity storage create yandex - Yandex Disk

用法：
   singularity storage create yandex [命令选项] [参数]

描述：
   --client-id
      OAuth客户端ID。
      
      通常保持空白。

   --client-secret
      OAuth客户端秘钥。
      
      通常保持空白。

   --token
      OAuth访问令牌，以JSON格式。
   
   --auth-url
      认证服务器URL。
      
      保持空白以使用提供者的默认值。

   --token-url
      令牌服务器URL。
      
      保持空白以使用提供者的默认值。

   --hard-delete
      永久删除文件，而不将其放入回收站。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项：
   --client-id value      OAuth客户端ID。[$CLIENT_ID]
   --client-secret value  OAuth客户端秘钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息

   高级选项

   --auth-url value   认证服务器URL。[$AUTH_URL]
   --encoding value   后端的编码方式。（默认值："Slash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --hard-delete      永久删除文件，而不将其放入回收站。（默认值：false）[$HARD_DELETE]
   --token value      OAuth访问令牌，以JSON格式。[$TOKEN]
   --token-url value  令牌服务器URL。[$TOKEN_URL]

   通用选项

   --name value  存储名称（默认值：自动生成）
   --path value  存储路径

```
{% endcode %}