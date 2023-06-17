# Yandex Disk

{% code fullWidth="true" %}
```
名称：
   singularity datasource add yandex - 云盘Yandex

用法：
   singularity datasource add yandex [命令选项] <数据集名称> <源路径>

描述：
   --yandex-client-id
      OAuth客户端ID。
      
      通常留空。

   --yandex-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --yandex-token
      OAuth访问令牌的JSON字符格式。

   --yandex-auth-url
      认证服务器URL。
      
      如果使用提供程序默认情况则留空。

   --yandex-token-url
      令牌服务器URL。
      
      如果使用提供程序默认情况则留空。

   --yandex-hard-delete
      永久删除文件而不将其放入回收站。

   --yandex-encoding
      后端编码方式。
      
      有关详细信息，请参阅[概览中的编码部分](/overview/#encoding)。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出数据集到CAR文件后删除数据源中的文件。  (默认: false)
   --rescan-interval value  前次扫描后经过此时间间隔后自动重新扫描数据源目录 (默认: 禁用)

   Yandex选项

   --yandex-auth-url value       认证服务器URL。 [$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth客户端ID。 [$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth客户端密钥。 [$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       后端编码方式。 (默认: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    永久删除文件而不将其放入回收站。 (默认: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth访问令牌的JSON字符格式。 [$YANDEX_TOKEN]
   --yandex-token-url value      令牌服务器URL。 [$YANDEX_TOKEN_URL]

```
{% endcode %}