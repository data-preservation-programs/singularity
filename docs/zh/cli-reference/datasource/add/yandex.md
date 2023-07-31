# Yandex Disk

{% code fullWidth="true" %}
```
名称:
   singularity datasource add yandex - Yandex Disk

用法:
   singularity datasource add yandex [命令选项] <数据集名称> <源路径>

描述:
   --yandex-auth-url
      验证服务器 URL。
      
      留空以使用提供程序默认值。

   --yandex-client-id
      OAuth 客户端 ID。
      
      通常空白。

   --yandex-client-secret
      OAuth 客户端密钥。
      
      通常空白。

   --yandex-encoding
      后端的编码。
      
      有关更多详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --yandex-hard-delete
      永久删除文件，而不是将它们放入回收站。

   --yandex-token
      OAuth 访问令牌（JSON 颗粒）。

   --yandex-token-url
      令牌服务器 URL。
      
      留空以使用提供程序默认值。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 导出数据集为 CAR 文件后删除数据集的文件。 (默认值: false)
   --rescan-interval value  上次成功扫描后，自动重新扫描源目录的时间间隔 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: 就绪)

   Yandex 选项

   --yandex-auth-url value       验证服务器 URL。[$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth 客户端 ID。[$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth 客户端密钥。[$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       后端的编码。 (默认值: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    永久删除文件，而不是将它们放入回收站。 (默认值: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth 访问令牌（JSON 颗粒）。[$YANDEX_TOKEN]
   --yandex-token-url value      令牌服务器 URL。[$YANDEX_TOKEN_URL]

```
{% endcode %}