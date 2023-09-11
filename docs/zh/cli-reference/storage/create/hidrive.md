# HiDrive

{% code fullWidth="true" %}
```
命令:
   singularity storage create hidrive - HiDrive

使用方式:
   singularity storage create hidrive [选项] [参数...]

说明:
   --client-id
      OAuth 客户端 ID。
 
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌(JSON 格式)。

   --auth-url
      认证服务器 URL。
      
      通常留空以使用提供商的默认值。

   --token-url
      令牌服务器 URL。
      
      通常留空以使用提供商的默认值。

   --scope-access
      rclone 请求 HiDrive 访问权限时应使用的访问权限。
      
      示例:
         | rw | 读写访问权限。
         | ro | 只读访问权限。

   --scope-role
      rclone 请求 HiDrive 访问权限时应使用的用户级别。
      
      示例:
         | user  | 用户级别访问管理权限。
         |       | 在大多数情况下这将足够。
         | admin | 完全访问管理权限。
         | owner | 完全访问管理权限。

   --root-prefix
      所有路径的根/父文件夹。
      
      填入指定的文件夹作为远程路径的父文件夹，rclone 可以以任何文件夹为其起点。
      
      示例:
         | /       | rclone 可访问的最高级目录。
         |         | 如果 rclone 使用一个正常的 HiDrive 用户帐户，则与 "root" 等同。
         | root    | HiDrive 用户帐户的最高级目录
         | <unset> | 这指定您的路径没有根前缀。
         |         | 使用此选项时，您总是需要使用一个有效的父文件夹指定远程路径，例如 "remote:/path/to/dir" 或 "remote:root/path/to/dir"。

   --endpoint
      服务的端点。
      
      这是将发起 API 调用的 URL。

   --disable-fetching-member-count
      除非绝对必要，否则不获取目录中的对象数量。
      
      如果不获取子目录中对象数量，请求可能会更快。

   --chunk-size
      分块上传的块大小。
      
      任何大于所配置截断值的文件（或未知大小的文件）将按此大小分块上传。
      
      上限为 2147483647 字节（约为 2.000Gi）。
      这是单个上传操作最大支持的字节数。
      将此设置为上限以上的较大值或负值会导致上传失败。
      
      将此值设置为较大值可以提高上传速度，但会使用更多内存。
      将此设置为较小值可以节省内存。

   --upload-cutoff
      分块上传的截断/阈值。
      
      任何大于此值的文件将按照配置的块大小逐块上传。
      
      上限为 2147483647 字节（约为 2.000Gi）。
      这是单个上传操作最大支持的字节数。
      将此设置为上限以上的较大值会导致上传失败。

   --upload-concurrency
      分块上传的并发性。
      
      这是同一文件并发运行的传输数的上限。
      将此设置为小于 1 的值将导致上传死锁。
      
      如果您通过高速链接上传少量的大文件
      并且这些上传未能充分利用您的带宽，那么增加
      此值可以帮助加速传输。

   --encoding
      后端的编码。
      
      有关详细信息，请参阅概述中的[编码部分](/overview/#encoding)。


选项:
   --client-id value      OAuth 客户端 ID。 [$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。 [$CLIENT_SECRET]
   --help, -h             显示帮助信息
   --scope-access value   rclone 请求 HiDrive 访问权限时应使用的访问权限。 (默认值: "rw") [$SCOPE_ACCESS]

   高级选项

   --auth-url value                 认证服务器 URL。 [$AUTH_URL]
   --chunk-size value               分块上传的块大小。 (默认值: "48Mi") [$CHUNK_SIZE]
   --disable-fetching-member-count  除非绝对必要，否则不获取目录中的对象数量。 (默认值: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 后端的编码。 (默认值: "Slash,Dot") [$ENCODING]
   --endpoint value                 服务的端点。 (默认值: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              所有路径的根/父文件夹。 (默认值: "/") [$ROOT_PREFIX]
   --scope-role value               rclone 请求 HiDrive 访问权限时应使用的用户级别。 (默认值: "user") [$SCOPE_ROLE]
   --token value                    OAuth 访问令牌(JSON 格式)。 [$TOKEN]
   --token-url value                令牌服务器 URL。 [$TOKEN_URL]
   --upload-concurrency value       分块上传的并发性。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分块上传的截断/阈值。 (默认值: "96Mi") [$UPLOAD_CUTOFF]

   通用选项

   --name value  存储的名称 (默认值: 自动生成)
   --path value  存储的路径

```
{% endcode %}