# HiDrive

{% code fullWidth="true" %}
```
名称:
   singularity datasource add hidrive - HiDrive

使用方法:
   singularity datasource add hidrive [命令选项] <数据集名称> <源路径>

描述:
   --hidrive-auth-url
      认证服务器URL。
      
      留空以使用提供商的默认值。

   --hidrive-chunk-size
      分块上传的块大小。
      
      大于设定阈值（或未知大小的文件）的文件将以此大小的块上传。
      
      这个值的上限为2147483647字节（大约2 GB）。
      这是单个上传操作所能支持的最大字节数。
      将此值设置得超过上限或为负值将导致上传失败。
      
      将此值设得更大可能会提高上传速度，但会消耗更多的内存。
      将此值设得更小可以节省内存。

   --hidrive-client-id
      OAuth客户端ID。
      
      通常留空。

   --hidrive-client-secret
      OAuth客户端秘钥。
      
      通常留空。

   --hidrive-disable-fetching-member-count
      除非绝对必要，否则不获取目录中的对象数量。
      
      如果不获取子目录中的对象数量，则请求可能更快。

   --hidrive-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --hidrive-endpoint
      服务的端点。
      
      这是API调用将要发送至的URL。

   --hidrive-root-prefix
      所有路径的根/父文件夹。
      
      填写以将指定的文件夹用作所有路径在远程服务中的起始位置。
      这样，rclone可以以任何文件夹作为起点。

      示例:
         | /       | rclone可以访问的最高级目录。
                   | 如果rclone使用的是普通的HiDrive用户帐户，则与"root"等效。
         | root    | HiDrive用户帐户的最高级目录
         | <unset> | 表示没有根前缀用于您的路径。
                   | 使用该选项时，您将始终需要使用有效的父目录指定到该远程的路径，例如 "remote:/path/to/dir" 或 "remote:root/path/to/dir"。

   --hidrive-scope-access
      rclone向HiDrive请求访问时应使用的访问权限。

      示例:
         | rw | 读写资源的权限。
         | ro | 只读资源的权限。

   --hidrive-scope-role
      rclone向HiDrive请求访问时应使用的用户级别。

      示例:
         | user  | 管理权限的用户级访问。
                 | 在大多数情况下将足够。
         | admin | 扩展的管理权限访问。
         | owner | 完整的管理权限访问。

   --hidrive-token
      OAuth访问令牌，以JSON形式。

   --hidrive-token-url
      令牌服务器URL。
      
      留空以使用提供商的默认值。

   --hidrive-upload-concurrency
      分块上传的并发数。
      
      对于同时运行的相同文件的转移数，这是一个上限。
      将此值设为小于1的值将导致上传死锁。
      
      如果您在高速链接上上传少量大文件
      并且这些上传没有充分利用带宽，则增加此值可能有助于加快传输速度。

   --hidrive-upload-cutoff
      分块上传的阈值。
      
      大于此大小的文件将以设定的块大小进行分块上传。
      
      这个值的上限为2147483647字节（大约2 GB）。
      这是单个上传操作所能支持的最大字节数。
      将此值设得超过上限将导致上传失败。

选项:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 在将数据集导出为CAR文件后删除文件。  (默认: false)
   --rescan-interval value  当距离上次成功扫描的时间间隔达到此设定值时，自动重新扫描源目录（默认: 禁用）
   --scanning-state value   设置初始扫描状态（默认: 准备就绪）

   HiDrive选项

   --hidrive-auth-url value                       认证服务器URL。[$HIDRIVE_AUTH_URL]
   --hidrive-chunk-size value                     分块上传的块大小。（默认: "48Mi"）[$HIDRIVE_CHUNK_SIZE]
   --hidrive-client-id value                      OAuth客户端ID。[$HIDRIVE_CLIENT_ID]
   --hidrive-client-secret value                  OAuth客户端秘钥。[$HIDRIVE_CLIENT_SECRET]
   --hidrive-disable-fetching-member-count value  除非绝对必要，否则不获取目录中的对象数量。（默认: "false"）[$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
   --hidrive-encoding value                       后端的编码方式。（默认: "Slash,Dot"）[$HIDRIVE_ENCODING]
   --hidrive-endpoint value                       服务的端点。（默认: "https://api.hidrive.strato.com/2.1"）[$HIDRIVE_ENDPOINT]
   --hidrive-root-prefix value                    所有路径的根/父文件夹。（默认: "/"）[$HIDRIVE_ROOT_PREFIX]
   --hidrive-scope-access value                   rclone向HiDrive请求访问时应使用的访问权限。（默认: "rw"）[$HIDRIVE_SCOPE_ACCESS]
   --hidrive-scope-role value                     rclone向HiDrive请求访问时应使用的用户级别。（默认: "user"）[$HIDRIVE_SCOPE_ROLE]
   --hidrive-token value                          OAuth访问令牌，以JSON形式。[$HIDRIVE_TOKEN]
   --hidrive-token-url value                      令牌服务器URL。[$HIDRIVE_TOKEN_URL]
   --hidrive-upload-concurrency value             分块上传的并发数。（默认: "4"）[$HIDRIVE_UPLOAD_CONCURRENCY]
   --hidrive-upload-cutoff value                  分块上传的阈值。（默认: "96Mi"）[$HIDRIVE_UPLOAD_CUTOFF]

```
{% endcode %}