# Google照片

{% code fullWidth="true" %}
```
命令:
   singularity datasource add gphotos - Google照片

用法:
   singularity datasource add gphotos [命令选项] <数据集名称> <源路径>

描述:
   --gphotos-auth-url
      认证服务器URL。
      
      留空以使用提供程序的默认值。

   --gphotos-client-id
      OAuth客户端ID。
      
      通常留空。

   --gphotos-client-secret
      OAuth客户端密钥。
      
      通常留空。

   --gphotos-encoding
      后端的编码。
      
      更多信息请参见[概述中的编码部分](/overview/#encoding)。

   --gphotos-include-archived
      同时查看和下载已归档的媒体。
      
      默认情况下，rclone不会请求已归档的媒体。因此，在同步时，
      在目录列表或传输中，不会显示已归档的媒体。
      
      请注意，相册中的媒体始终可见并同步，无论其归档状态如何。
      
      使用此标志，已归档的媒体始终在目录列表和传输中可见。
      
      如果没有此标志，已归档的媒体将不会在目录列表中可见，并且不会被传输。

   --gphotos-read-only
      设置为使Google照片后端只读。
      
      如果选择只读，则rclone将仅请求对您的照片的只读访问权限；
      否则，rclone将请求完全访问权限。

   --gphotos-read-size
      设置为读取媒体项目的大小。
      
      通常情况下，rclone不会读取媒体项目的大小，因为这需要另一个事务。
      对于同步来说，这不是必需的。然而，当使用rclone mount时，rclone mount需要预先了解文件的大小，所以如果要读取媒体，则建议在使用rclone mount时设置此标志。

   --gphotos-start-year
      年份限定将要下载的照片为那些在指定年份之后上传的照片。

   --gphotos-token
      OAuth访问令牌作为一个JSON块。

   --gphotos-token-url
      令牌服务器URL。
      
      留空以使用提供程序的默认值。


选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险]在将数据集导出为CAR文件后删除数据集的文件。(默认值：false)
   --rescan-interval value  当从上次成功扫描以来经过此间隔时，自动重新扫描源目录。(默认值：禁用)
   --scanning-state value   设置初始扫描状态。(默认值：ready)

   gphotos选项

   --gphotos-auth-url value          认证服务器URL。[$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth客户端ID。[$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth客户端密钥。[$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          后端的编码。(默认值："Slash,CrLf,InvalidUtf8,Dot")[$GPHOTOS_ENCODING]
   --gphotos-include-archived value  同时查看和下载已归档的媒体。(默认值："false")[$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         设置为使Google照片后端只读。(默认值："false")[$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         设置为读取媒体项目的大小。(默认值："false")[$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        年份限定将要下载的照片为那些在指定年份之后上传的照片。(默认值："2000")[$GPHOTOS_START_YEAR]
   --gphotos-token value             OAuth访问令牌作为一个JSON块。[$GPHOTOS_TOKEN]
   --gphotos-token-url value         令牌服务器URL。[$GPHOTOS_TOKEN_URL]

```
{% endcode %}
