# Dropbox

{% code fullWidth="true" %}
```
名称：
   singularity datasource add dropbox - Dropbox

用法：
   singularity datasource add dropbox [命令选项] <数据集名称> <源路径>

描述：
   --dropbox-auth-url
      认证服务器URL。
      
      留空以使用提供程序的默认值。

   --dropbox-batch-commit-timeout
      等待批处理完成提交的最长时间

   --dropbox-batch-mode
      文件上传批处理同步|异步|关闭。
      
      这会设置rclone使用的批处理模式。
      
      有三个可能的值：
      
      - 关闭 - 不使用批处理
      - 同步 - 批量上传并检查完成（默认值）
      - 异步 - 批量上传且不检查完成
      
      当rclone退出时，它会关闭任何未完成的批次，这可能会导致延迟。
      

   --dropbox-batch-size
      上传批处理中的最大文件数量。
      
      这会设置要上传的文件批次的大小，必须小于1000。
      
      默认情况下，此值为0，表示根据batch_mode的设置计算批次大小。
      
      - batch_mode: async - 默认的batch_size为100
      - batch_mode: sync - 默认的batch_size与--transfers相同
      - batch_mode: off - 不使用
      
      当上传大量小文件时，设置此值非常有用，因为它可以加快上传速度。
      您可以使用--transfers 32来最大化吞吐量。
      

   --dropbox-batch-timeout
      在上传之前，允许空闲上传批次的最长时间。
      
      如果上传批次空闲时间超过此时间，将上传批次。
      
      默认值为0，这意味着rclone将根据所使用的batch_mode选择一个合适的默认值。
      
      - batch_mode: async - 默认的batch_timeout为500毫秒
      - batch_mode: sync - 默认的batch_timeout为10秒
      - batch_mode: off - 不使用
      

   --dropbox-chunk-size
      上传分块的大小（< 150Mi）。
      
      大于此大小的任何文件将按此大小进行分块上传。
      
      注意，分块在内存中进行缓冲（一次一个），以便rclone可以处理重试。
      将此值设置得更大会略微增加速度（在测试中，对于128MiB而言，最多增加10%），
      但代价是更多的内存使用。如果内存紧张，可以将其设置得更小。

   --dropbox-client-id
      OAuth客户端ID。
      
      通常留空。

   --dropbox-client-secret
      OAuth客户端秘钥。
      
      通常留空。

   --dropbox-encoding
      后端的编码。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

   --dropbox-impersonate
      使用商业账户时模拟此用户。
      
      请注意，如果要使用账户模拟，您应该在运行“rclone config”时确保设置此标志，
      因为它将导致rclone请求“members.read”范围，而通常情况下它不会。
      这是为了将成员的电子邮件地址查找为dropbox在API中使用的内部ID。
      
      使用“members.read”范围需要Dropbox团队管理员在OAuth流程中批准。
      
      您将需要使用自己的应用（设置自己的client_id和client_secret）来使用此选项，
      因为当前rclone的默认权限集不包括“members.read”。
      一旦v1.55或更高版本在所有地方都在使用，就可以添加此权限。

   --dropbox-shared-files
      让rclone处理单个共享文件。
      
      在此模式下，rclone的功能极为有限-只支持列出操作（ls，lsl等）和读取操作（例如下载）。
      在此模式下，其他所有操作都将被禁用。

   --dropbox-shared-folders
      让rclone处理共享文件夹。
            
      当此标志与无路径一起使用时，仅支持List操作，并且将列出所有可用的共享文件夹。
      如果指定了路径，则第一部分将被解释为共享文件夹的名称。
      然后，rclone尝试将此共享文件夹挂载到根命名空间。
      成功后，rclone将正常进行处理。共享文件夹现在几乎与普通文件夹相同，
      支持所有普通操作。
      
      请注意，我们不会在之后卸载共享文件夹，
      因此在首次使用特定共享文件夹后，可以省略--dropbox-shared-folders。

   --dropbox-token
      作为JSON字符串的OAuth访问令牌。

   --dropbox-token-url
      令牌服务器URL。
      
      留空以使用提供程序的默认值。


选项：
   --help, -h  显示帮助信息

   数据准备选项

   --delete-after-export    [危险] 导出数据集到CAR文件后删除数据集中的文件。 （默认值：false）
   --rescan-interval value  当距离上次成功扫描过去了此间隔时，自动重新扫描源目录（默认值：禁用）
   --scanning-state value   设置初始扫描状态（默认值：ready）

   dropbox选项

   --dropbox-auth-url value              认证服务器URL。[$DROPBOX_AUTH_URL]
   --dropbox-batch-commit-timeout value  等待批处理完成提交的最长时间（默认值：“10m0s”）[$DROPBOX_BATCH_COMMIT_TIMEOUT]
   --dropbox-batch-mode value            文件上传批处理同步|异步|关闭（默认值：“sync”）[$DROPBOX_BATCH_MODE]
   --dropbox-batch-size value            上传批处理中的最大文件数量（默认值：“0”）[$DROPBOX_BATCH_SIZE]
   --dropbox-batch-timeout value         在上传之前，允许空闲上传批次的最长时间（默认值：“0s”）[$DROPBOX_BATCH_TIMEOUT]
   --dropbox-chunk-size value            上传分块的大小（< 150Mi）（默认值：“48Mi”）[$DROPBOX_CHUNK_SIZE]
   --dropbox-client-id value             OAuth客户端ID。[$DROPBOX_CLIENT_ID]
   --dropbox-client-secret value         OAuth客户端秘钥。[$DROPBOX_CLIENT_SECRET]
   --dropbox-encoding value              后端的编码（默认值：“Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot”）[$DROPBOX_ENCODING]
   --dropbox-impersonate value           使用商业账户时模拟此用户。[$DROPBOX_IMPERSONATE]
   --dropbox-shared-files value          让rclone处理单个共享文件（默认值：“false”）[$DROPBOX_SHARED_FILES]
   --dropbox-shared-folders value        让rclone处理共享文件夹（默认值：“false”）[$DROPBOX_SHARED_FOLDERS]
   --dropbox-token value                 作为JSON字符串的OAuth访问令牌。[$DROPBOX_TOKEN]
   --dropbox-token-url value             令牌服务器URL。[$DROPBOX_TOKEN_URL]

```
{% endcode %}