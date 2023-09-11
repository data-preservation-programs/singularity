# Dropbox

{% code fullWidth="true" %}
```
命令名称：
  singularity storage create dropbox - Dropbox

用法：
  singularity storage create dropbox [命令选项] [参数...]

描述：
  --client-id
     OAuth Client Id。
     
     通常留空。

  --client-secret
     OAuth Client Secret。
     
     通常留空。

  --token
     OAuth Access Token 作为 JSON 格式的字符串。

  --auth-url
     Auth server URL。
     
     留空以使用提供者的默认值。

  --token-url
     Token server URL。
     
     留空以使用提供者的默认值。

  --chunk-size
     上传块的大小（<150Mi）。
     
     大于此大小的所有文件将以此大小的块进行上传。
     
     请注意，块在内存中进行缓冲（一次一个），以便 rclone 可以处理重试。
     将此设置为较大的值会稍微增加速度（在测试中，128 MiB 的速度增加最多 10%），
     但会使用更多的内存。如果内存有限，可以将其设置得更小。

  --impersonate
     在使用商业帐户时模拟此用户。
     
     请注意，如果您想使用模拟用户，应确保在运行 "rclone config" 时设置此标志，
     因为这将导致 rclone 请求 "members.read" 作用域，该作用域默认情况下不会请求。
     这是为了将成员的电子邮件地址查找到 dropbox 在 API 中使用的内部 ID。
     
     使用 "members.read" 作用域需要 Dropbox 团队管理员在 OAuth 流程中进行批准。
     
     如果要使用此选项，则必须使用您自己的应用（设置自己的 client_id 和 client_secret）。
     因为 rclone 的默认权限集中不包括 "members.read"。一旦使用 v1.55 或更高版本，
     可以添加此权限。

  --shared-files
     告诉 rclone 处理单个共享文件。
     
     在此模式下，rclone 的功能非常有限 - 只能支持列表（ls、lsl 等）操作和读取操作（例如下载）。
     这种模式下将禁用所有其他操作。

  --shared-folders
     告诉 rclone 处理共享文件夹。
     
     当此标志与没有路径一起使用时，只支持 List 操作，并且将列出所有可用的共享文件夹。
     如果指定了路径，第一部分将被解释为共享文件夹的名称。
     然后，rclone 将尝试将此共享文件夹挂载到根命名空间。
     成功后，rclone 将按正常方式继续进行操作。
     
     请注意，我们不会在使用特定共享文件夹的第一次使用后取消挂载共享文件夹，
     因此可以省略 --dropbox-shared-folders。

  --batch-mode
     上传文件的批处理同步|异步|关闭。
     
     这设置了 rclone 使用的批处理模式。
     
     有 3 种可能的值
     
     - off - 无批处理
     - sync - 批量上传并检查完成（默认值）
     - async - 批量上传并不检查完成
     
     当退出时，rclone 将关闭任何未完成的批次，这可能会导致延迟。

  --batch-size
     上传批处理中的最大文件数。
     
     这设置了要上传的文件批处理的批量大小。必须小于1000。
     
     默认情况下，此值为 0，这意味着根据 batch_mode 的设置，rclone 会计算批量大小。
     
     - batch_mode: async - 默认的 batch_size 是 100
     - batch_mode: sync - 默认的 batch_size 与 --transfers 相同
     - batch_mode: off - 未使用
     
     当您上传大量的小文件时，设置此值是个好主意，因为它们将更快。
     您可以使用 --transfers 32 来提高吞吐量。

  --batch-timeout
     在上传前允许空闲上传批处理的最长时间。
     
     如果上传批次空闲时间超过此时间，则会上传。
     
     默认值为0，这意味着 rclone 将根据使用的 batch_mode 选择合适的默认值。
     
     - batch_mode: async - 默认的 batch_timeout 是 500ms
     - batch_mode: sync - 默认的 batch_timeout 是 10s
     - batch_mode: off - 未使用

  --batch-commit-timeout
     等待批次完成提交的最长时间

  --encoding
     后端的编码。
     
     有关更多信息，请参阅概述中的 [编码部分](/overview/#encoding)。


选项：
  --client-id value      OAuth Client Id。[$CLIENT_ID]
  --client-secret value  OAuth Client Secret。[$CLIENT_SECRET]
  --help, -h             显示帮助信息

  高级选项：

  --auth-url value              Auth server URL。[$AUTH_URL]
  --batch-commit-timeout value  Max time to wait for a batch to finish committing (default: "10m0s") [$BATCH_COMMIT_TIMEOUT]
  --batch-mode value            Upload file batching sync|async|off. (default: "sync") [$BATCH_MODE]
  --batch-size value            Max number of files in upload batch. (default: 0) [$BATCH_SIZE]
  --batch-timeout value         Max time to allow an idle upload batch before uploading. (default: "0s") [$BATCH_TIMEOUT]
  --chunk-size value            Upload chunk size (< 150Mi). (default: "48Mi") [$CHUNK_SIZE]
  --encoding value              The encoding for the backend. (default: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$ENCODING]
  --impersonate value           Impersonate this user when using a business account. [$IMPERSONATE]
  --shared-files                Instructs rclone to work on individual shared files. (default: false) [$SHARED_FILES]
  --shared-folders              Instructs rclone to work on shared folders. (default: false) [$SHARED_FOLDERS]
  --token value                 OAuth Access Token as a JSON blob. [$TOKEN]
  --token-url value             Token server url. [$TOKEN_URL]

  通用选项：

  --name value  存储的名称（默认值：自动生成）
  --path value  存储的路径

```
{% endcode %}