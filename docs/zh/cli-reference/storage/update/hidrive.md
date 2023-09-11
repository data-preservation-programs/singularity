# HiDrive

{% code fullWidth="true" %}
```
名称：
   singularity storage update hidrive - HiDrive

用法：
   singularity storage update hidrive [命令选项] <名称|ID>

说明：
   --client-id
      OAuth 客户端 ID。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌(JSON格式)。

   --auth-url
      授权服务器 URL。
      
      留空以使用提供程序默认值。

   --token-url
      令牌服务器 URL。
      
      留空以使用提供程序默认值。

   --scope-access
      请求 HiDrive 访问权限时 rclone 应使用的访问权限。

      示例：
         | rw | 对资源进行读写访问。
         | ro | 对资源进行只读访问。

   --scope-role
      请求 HiDrive 访问权限时 rclone 应使用的用户级别。

      示例：
         | user  | 用户级别访问管理权限。
         |       | 在大多数情况下，这已足够。
         | admin | 具有完整访问管理权限。
         | owner | 具有完全访问管理权限。

   --root-prefix
      所有路径的根/父文件夹。
      
      填写以使用指定文件夹作为所有路径给定给远程的父文件夹。
      这样 rclone 可以将任何文件夹作为其起始点。

      示例：
         | /       | rclone 可访问的最顶层目录。
         |         | 如果 rclone 使用普通 HiDrive 用户账户，这将等同于"root"。
         | root    | HiDrive 用户账户的最顶层目录。
         | <unset> | 表示路径没有根前缀。
         |         | 使用此选项时，您始终需要指定有效父文件夹的路径，例如 "remote:/path/to/dir" 或 "remote:root/path/to/dir"。

   --endpoint
      该服务的终端点。
      
      这是将进行 API 调用的 URL。

   --disable-fetching-member-count
      除非绝对必要，否则不获取目录中的对象数。
      
      如果不获取子目录中的对象数，请求可能更快。

   --chunk-size
      分块上传的块大小。
      
      大于配置的截断值或未知大小的文件将以此大小的块上传。
      
      此值的上限为 2147483647 字节(约 2.000Gi)。
      这是单个上传操作支持的最大字节数。
      将此值设置为上限之上或负值将导致上传失败。
      
      将此值设置为较大值可能会增加上传速度，但会增加内存使用量。
      将此值设置为较小值可能会减少内存使用量。

   --upload-cutoff
      分块上传的截断/阈值。
      
      文件大小超过此值将以配置的块大小的块上传。
      
      此值的上限为 2147483647 字节(约 2.000Gi)。
      这是单个上传操作支持的最大字节数。
      将此值设置为上限之上将导致上传失败。

   --upload-concurrency
      分块上传的并发数。
      
      这是同一文件同时运行的传输数量上限。
      将此值设置为小于 1 的值将导致上传死锁。
      
      如果您在高速链接上上传少量大文件，并且这些上传未充分利用带宽，
      那么提高此值可能有助于加快传输速度。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

选项：
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --scope-access value   请求 HiDrive 访问权限时 rclone 应使用的访问权限。(默认值: "rw")[$SCOPE_ACCESS]

   高级选项

   --auth-url value                 授权服务器 URL。[$AUTH_URL]
   --chunk-size value               分块上传的块大小。(默认值: "48Mi")[$CHUNK_SIZE]
   --disable-fetching-member-count  除非绝对必要，否则不获取目录中的对象数。(默认值: false)[$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 后端的编码方式。(默认值: "Slash,Dot")[$ENCODING]
   --endpoint value                 该服务的终端点。(默认值: "https://api.hidrive.strato.com/2.1")[$ENDPOINT]
   --root-prefix value              所有路径的根/父文件夹。(默认值: "/")[$ROOT_PREFIX]
   --scope-role value               请求 HiDrive 访问权限时 rclone 应使用的用户级别。(默认值: "user")[$SCOPE_ROLE]
   --token value                    OAuth 访问令牌(JSON格式)。[$TOKEN]
   --token-url value                令牌服务器 URL。[$TOKEN_URL]
   --upload-concurrency value       分块上传的并发数。(默认值: 4)[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分块上传的截断/阈值。(默认值: "96Mi")[$UPLOAD_CUTOFF]

```
{% endcode %}