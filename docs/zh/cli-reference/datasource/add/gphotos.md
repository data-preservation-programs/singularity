# Google 照片

{% code fullWidth="true" %}
```
名称：
   singularity 数据源添加 gphotos - Google 照片

用法：
   singularity datasource add gphotos [命令选项] <数据集名称> <源路径>

说明：
   --gphotos-client-id
      OAuth 客户端 ID。
      
      通常留空。

   --gphotos-token
      OAuth 访问令牌，以 JSON 数据格式表示。

   --gphotos-auth-url
      鉴权服务器 URL。
      
      如果留空，则使用提供程序的默认值。

   --gphotos-token-url
      令牌服务器 URL。
      
      如果留空，则使用提供程序的默认值。

   --gphotos-read-only
      设置为使 Google 照片后端只读。
      
      如果选择只读，则 rclone 只请求照片的只读访问权限；否则，rclone 将请求完全访问权限。

   --gphotos-start-year
      年份限制下载的照片为上传自给定年份之后的照片。

   --gphotos-encoding
      后端的编码方式。
      
      了解更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --gphotos-client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --gphotos-read-size
      设置为读取媒体项的大小。
      
      通常情况下，rclone 不会读取媒体项的大小，因为这需要另一次交易。对于同步来说，这不是必要的。但是，使用 rclone 挂载需要提前知道文件的大小，因此建议在使用 rclone 挂载时设置此标志，如果要读取媒体，则需要设置它。

   --gphotos-include-archived
      同时查看和下载存档的媒体。
      
      默认情况下，rclone 不会请求存档媒体。因此，在同步时，存档媒体在目录列表或传输中不可见。
      
      请注意，无论其存档状态如何，相册中的媒体始终可见且已同步。
      
      通过此标志，存档媒体始终在目录列表和传输中可见。
      
      没有此标志，存档媒体将不会在目录列表中可见，也不会被传输。

选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 导出数据集到 CAR 文件后删除数据集的文件。  (默认值：false)
   --rescan-interval value  当从上次成功扫描的时间过去此间隔时自动重新扫描源目录（默认值：禁用）

   gphotos 选项

   --gphotos-auth-url value          鉴权服务器 URL。[$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth 客户端 ID。[$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth 客户端密钥。[$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          后端的编码方式。 (默认值："Slash,CrLf,InvalidUtf8,Dot") [$GPHOTOS_ENCODING]
   --gphotos-include-archived value  同时查看和下载存档的媒体。 (默认值："false") [$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         设置为使 Google 照片后端只读。 (默认值："false") [$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         设置为读取媒体项的大小。 (默认值："false") [$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        年份限制下载的照片为上传自给定年份之后的照片。 (默认值："2000") [$GPHOTOS_START_YEAR]
   --gphotos-token value             OAuth 访问令牌，以 JSON 数据格式表示。[$GPHOTOS_TOKEN]
   --gphotos-token-url value         令牌服务器 URL。[$GPHOTOS_TOKEN_URL]

```
{% endcode %}