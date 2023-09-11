# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity storage update dropbox - Dropbox

USAGE:
   singularity storage update dropbox [命令选项] <名称|ID>

DESCRIPTION:
   --client-id
      OAuth 客户端 ID。
      
      通常留空。

   --client-secret
      OAuth 客户端密钥。
      
      通常留空。

   --token
      OAuth 访问令牌的 JSON 对象。

   --auth-url
      认证服务器 URL。
      
      如果留空，则使用提供者的默认设置。

   --token-url
      令牌服务器 URL。
      
      如果留空，则使用提供者的默认设置。

   --chunk-size
      上传分片大小（< 150Mi）。
      
      所有大于此大小的文件将被分片上传。
      
      请注意，分片会在内存中进行缓冲（每次只有一个），以便 rclone 可以进行重试。
      增大此值会稍微提高速度（在测试中对于 128 MiB 至多提高 10%），但会增加内存使用量。
      如果内存不足，可以减小此值。

   --impersonate
      在使用企业账户时冒名顶替此用户。
      
      请注意，如果要使用冒名顶替功能，应确保在运行 "rclone config" 时设置此标志，
      因为这将导致 rclone 请求 "members.read" 作用域，而这是通常不会请求的。
      这是为了将成员的电子邮件地址查找为 Dropbox 在 API 中使用的内部 ID。
      
      使用 "members.read" 作用域将需要 Dropbox 团队管理员在 OAuth 流程中进行批准。
      
      要使用此选项，您必须使用自己的应用程序（设置自己的 client_id 和 client_secret），
      因为 rclone 的默认权限集中不包括 "members.read"。
      一旦 everywhere 都在使用 v1.55 或更高版本后，可以添加此权限。

   --shared-files
      指示 rclone 处理单个共享文件。
      
      在此模式下，rclone 的功能非常有限 - 仅支持列表（ls，lsl 等）操作和读取操作（例如下载）。
      在此模式下，所有其他操作将被禁用。

   --shared-folders
      指示 rclone 处理共享文件夹。
      
      当此标志与无路径一起使用时，仅支持列表操作，并列出所有可用的共享文件夹。
      如果指定了路径，则第一部分将被解释为共享文件夹的名称。
      然后，rclone 将尝试将此共享文件夹挂载到根命名空间。
      如果成功，rclone 将继续正常进行。此时共享文件夹基本上就是普通文件夹，支持所有正常操作。
      
      请注意，我们不会在之后卸载共享文件夹，因此之后可以省略 --shared-folders 标志，
      除非使用特定共享文件夹的第一个用法。

   --batch-mode
      文件批处理上传的模式 sync|async|off。
      
      这设置了 rclone 使用的批处理模式。
      
      有 3 种可能的值：
      
      - off - 不使用批处理
      - sync - 批处理上传并检查完成（默认值）
      - async - 批处理上传并不检查完成
      
      当 rclone 退出时，它将关闭任何未完成的批处理，这可能会造成一些延迟。

   --batch-size
      上传批处理中的文件最大数量。
      
      这设置了要上传的文件的批处理大小。
      它必须小于 1000。
      
      默认值为 0，这意味着 rclone 会根据 batch_mode 的设置计算批处理大小。
      
      - batch_mode: async - 默认 batch_size 为 100
      - batch_mode: sync - 默认 batch_size 与 --transfers 相同
      - batch_mode: off - 不使用
      
      当您上传大量的小文件时，设置此值是个好主意，因为它会加快上传速度。
      您可以使用 --transfers 32 来最大化吞吐量。

   --batch-timeout
      允许一个空闲的上传批处理最长时间后再上传。
      
      如果一个上传批处理空闲时间超过此值，它将被上传。
      
      默认值为 0，这意味着 rclone 将根据所使用的 batch_mode 选择合适的默认值。
      
      - batch_mode: async - 默认 batch_timeout 为 500ms
      - batch_mode: sync - 默认 batch_timeout 为 10s
      - batch_mode: off - 不使用

   --batch-commit-timeout
      等待批处理完成提交的最长时间

   --encoding
      后端的编码方式。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

OPTIONS:
   --client-id value      OAuth 客户端 ID。[$CLIENT_ID]
   --client-secret value  OAuth 客户端密钥。[$CLIENT_SECRET]
   --help, -h             显示帮助信息

   高级选项

   --auth-url value              认证服务器 URL。[$AUTH_URL]
   --batch-commit-timeout value  等待批处理完成提交的最长时间（默认值："10m0s"）[$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            文件批处理上传的模式 sync|async|off（默认值："sync"）[$BATCH_MODE]
   --batch-size value            上传批处理中的文件最大数量（默认值：0）[$BATCH_SIZE]
   --batch-timeout value         允许一个空闲的上传批处理最长时间后再上传（默认值："0s"）[$BATCH_TIMEOUT]
   --chunk-size value            上传分片大小（< 150Mi）（默认值："48Mi"）[$CHUNK_SIZE]
   --encoding value              后端的编码方式（默认值："Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"）[$ENCODING]
   --impersonate value           在使用企业账户时冒名顶替此用户。[$IMPERSONATE]
   --shared-files                指示 rclone 处理单个共享文件（默认值：false）[$SHARED_FILES]
   --shared-folders              指示 rclone 处理共享文件夹（默认值：false）[$SHARED_FOLDERS]
   --token value                 OAuth 访问令牌的 JSON 对象。[$TOKEN]
   --token-url value             令牌服务器 URL。[$TOKEN_URL]

```
{% endcode %}