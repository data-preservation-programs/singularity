# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity storage update gphotos - Google相册

USAGE:
   singularity storage update gphotos [命令选项] <名称|ID>

DESCRIPTION:
   --client-id
      OAuth客户端ID。

      通常留空。

   --client-secret
      OAuth客户端密钥。

      通常留空。

   --token
      作为JSON blob的OAuth访问令牌。

   --auth-url
      认证服务器URL。

      留空以使用提供程序默认值。

   --token-url
      令牌服务器URL。

      留空以使用提供程序默认值。

   --read-only
      设置Google相册后端为只读模式。

      如果选择只读模式，则rclone仅会请求对您的照片的只读访问权限，
      否则，rclone将请求对您的照片的完全访问权限。

   --read-size
      设置读取媒体项大小。

      通常rclone不会读取媒体项的大小，因为这需要另一个事务。
      这对于同步而言不是必需的。
      但是，rclone mount在读取文件前需要提前知道文件大小，
      所以当使用rclone mount时，建议设置此标志以读取媒体。

   --start-year
      使下载的照片限制在给定年份之后上传的照片。

   --include-archived
      同时查看和下载已归档的媒体。

      默认情况下，rclone不会请求已归档的媒体。
      因此，在同步时，已归档的媒体在目录列表或传输中不可见。

      请注意，无论其归档状态如何，相册中的媒体始终可见并同步。

      使用此标志时，已归档的媒体将始终在目录列表中可见并传输。

      不使用此标志时，已归档的媒体将不会在目录列表中可见，并且不会传输。

   --encoding
      后端的编码。

      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。


OPTIONS:
   --client-id value      OAuth客户端ID。[$CLIENT_ID]
   --client-secret value  OAuth客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --read-only            设置Google相册后端为只读模式。(默认值：false)[$READ_ONLY]

   高级选项

   --auth-url value    认证服务器URL。[$AUTH_URL]
   --encoding value    后端的编码。(默认值："Slash,CrLf,InvalidUtf8,Dot")[$ENCODING]
   --include-archived  同时查看和下载已归档的媒体。(默认值：false)[$INCLUDE_ARCHIVED]
   --read-size         设置读取媒体项的大小。(默认值：false)[$READ_SIZE]
   --start-year value  使下载的照片限制在给定年份之后上传的照片。(默认值：2000)[$START_YEAR]
   --token value       作为JSON blob的OAuth访问令牌。[$TOKEN]
   --token-url value   令牌服务器URL。[$TOKEN_URL]

```
{% endcode %}