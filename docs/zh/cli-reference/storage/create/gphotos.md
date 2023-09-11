# Google Photos

{% code fullWidth="true" %}
```
名称:
   singularity storage create gphotos - Google Photos

使用方法:
   singularity storage create gphotos [命令选项] [参数...]

说明:
   --client-id
      OAuth 客户端 ID。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌，以 JSON 格式提供。

   --auth-url
      认证服务器 URL。
      
      留空则使用提供程序的默认值。

   --token-url
      令牌服务器 URL。
      
      留空则使用提供程序的默认值。

   --read-only
      设置此标志以使 Google Photos 后端只读。
      
      如果选择只读，则 rclone 仅请求对照片的只读访问权限，否则 rclone 将请求完全访问权限。

   --read-size
      设置此标志以读取媒体项的大小。
      
      通常，rclone 不会读取媒体项的大小，因为这将产生额外的交易。这对于同步操作是不必要的。然而，rclone mount 需要在读取文件之前提前知道文件的大小，因此在使用 rclone mount 时建议设置此标志以读取媒体。

   --start-year
      年份限制要下载的照片，仅限于上传日期在给定年份之后的照片。

   --include-archived
      还可以查看和下载已封存的媒体。
      
      默认情况下，rclone 不会请求封存的媒体。因此，在同步操作中，目录列表和传输过程中不会显示封存的媒体。
      
      请注意，相册中的媒体始终可见并会进行同步，无论其封存状态如何。
      
      使用此标志，已封存的媒体始终在目录列表中可见，并进行传输。
      
      不使用此标志，已封存的媒体将不会在目录列表中可见，也不会传输。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。


选项:
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             展示帮助信息
   --read-only            设置此标志以使 Google Photos 后端只读。（默认为 false）[$READ_ONLY]

   高级选项

   --auth-url value    认证服务器 URL。[$AUTH_URL]
   --encoding value    后端的编码方式。（默认为 "Slash,CrLf,InvalidUtf8,Dot"）[$ENCODING]
   --include-archived  还可以查看和下载已封存的媒体。（默认为 false）[$INCLUDE_ARCHIVED]
   --read-size         设置此标志以读取媒体项的大小。（默认为 false）[$READ_SIZE]
   --start-year value  年份限制要下载的照片，仅限于上传日期在给定年份之后的照片。（默认为 2000）[$START_YEAR]
   --token value       OAuth 访问令牌，以 JSON 格式提供。[$TOKEN]
   --token-url value   令牌服务器 URL。[$TOKEN_URL]

   通用选项

   --name value  存储的名称（默认为自动生成）
   --path value  存储的路径

```
{% endcode %}