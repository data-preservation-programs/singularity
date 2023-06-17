# Dropbox

{% code fullWidth="true" %}
```
名称:
    singularity datasource add dropbox - Dropbox

用法:
    singularity datasource add dropbox [命令行选项] <数据集名称> <源路径>

描述:
    --dropbox-token
        OAuth 访问令牌，Json 格式。
    --dropbox-auth-url
        Auth 服务器 URL。
        留空使用提供程序默认。
    --dropbox-encoding
        后端编码。
        更多信息请参阅[概述中的编码部分](/overview/#encoding)。
    --dropbox-chunk-size
        上传块大小 (< 150Mi)。
        任何比这更大的文件将分块上传。
        注意块会在内存中缓冲（一次只有一个），因此rclone可以处理重试。将此设置更大会稍微提高速度（在测试中最多为128 MiB的10%），代价是使用更多内存。如果内存紧张，可以将其设置得更小。
    --dropbox-batch-mode
        上载文件批量同步|异步|关闭。
        这将设置 rclone 使用的批量模式。
        有3个可选值
        - 关闭 - 不批处理
        - 同步 - 批量上传并检查完成（默认设置）
        - 异步 - 批量上传并不检查完成
        Rclone 会在退出时关闭所有未完成的批处理，这可能会导致延迟。
    --dropbox-token-url
        令牌服务器 URL。
        留空使用提供程序默认。
    --dropbox-shared-folders
        指示 rclone 处理共享文件夹。
        使用此标志而不使用任何路径仅支持 列出 操作，会列出所有可用的共享文件夹。如果指定路径，则第一部分将被解释为共享文件夹的名称。然后rclone将尝试将这个共享文件夹挂载到根名称空间。成功挂载后，共享文件夹基本上就是一个普通文件夹，支持所有常规操作。
        请注意，我们不会在之后卸载共享文件夹，因此在使用特定共享文件夹的第一次使用后，可以省略 --dropbox-shared-folders。
    --dropbox-batch-size
        上传批次中的文件的最大数量。
        这设置了要上传的文件的批量大小。它必须小于 1000。
        默认情况下，这为 0，这意味着根据批量模式的设置，rclone将计算批量大小。
        - 批量模式: 异步 - 默认批量大小为 100
        - 批量模式: 同步 - 默认批量大小与 --transfers 相同
        - 批量模式: 关闭 - 未使用
        Rclone 会在退出时关闭所有未完成的批处理，这可能会导致延迟。
        如果要上传许多小文件，则设置此项是个好主意，因为它将使它们更快。您可以使用 --transfers 32 来最大化吞吐量。
    --dropbox-batch-commit-timeout
        等待批次完成提交的最长时间。
    --dropbox-client-id
        OAuth 客户端 ID。
        通常为空。
    --dropbox-client-secret
        OAuth 客户端密钥。
        通常为空。
    --dropbox-impersonate
        在使用商务帐户时冒充此用户。
        请注意，如果要使用impersonate，应确保在运行“rclone配置”时设置了此标志，因为这将导致rclone请求“members.read”范围，它通常不会。这是将一个成员的电子邮件地址查找到 Dropbox 在 API 中使用的内部 ID 所需的。
        使用“members.read”范围将需要 Dropbox 团队管理员在OAuth流程中批准。
        您将不得不使用自己的应用程序（设置自己的client_id和client_secret）来使用此选项，因为当前rclone的默认权限集合不包括“members.read”。这可以在任何地方使用v1.55或更高版本后添加。
    --dropbox-shared-files
        指示 rclone 处理单个共享文件。
        在此模式下，rclone 的功能极为有限 - 仅支持列表（ls、lsl 等）操作和读操作（例如下载）。在此模式下，将禁用所有其他操作。
    --dropbox-batch-timeout
        允许空闲上传批量的最长时间。
        如果上传批处理空闲时间超过此时长，它将被上传。
        默认值为 0，这意味着 rclone 将根据所使用的批量模式选择合适的默认值。
        - 批量模式: 异步 - 默认批量超时为 500ms
        - 批量模式: 同步 - 默认批量超时为 10s
        - 批量模式: 关闭 - 未使用

命令行选项:
    --help, -h   显示帮助
    数据准备选项
    --delete-after-export [危险] 将数据集导出到 CAR 文件后删除其中的文件。 （默认值:false）
    --rescan-interval value 在自动扫描源目录时，扫描完成后会在此时间段后自动重新扫描。 （默认值:已禁用）
    DropBox 相关选项
    --dropbox-auth-url value Auth 服务器 URL。 [$DROPBOX_AUTH_URL]
    --dropbox-batch-commit-timeout value 等待批次完成提交的最长时间。（默认值:"10分钟") [$DROPBOX_BATCH_COMMIT_TIMEOUT]
    --dropbox-batch-mode value 上载文件批量同步|异步|关闭。（默认值:"同步") [$DROPBOX_BATCH_MODE]
    --dropbox-batch-size value 上传批次中的文件的最大数量。（默认值:"0") [$DROPBOX_BATCH_SIZE]
    --dropbox-batch-timeout value 允许空闲上传批量的最长时间。（默认值:"0s") [$DROPBOX_BATCH_TIMEOUT]
    --dropbox-chunk-size value 上传块大小 (< 150Mi)。（默认值:"48Mi") [$DROPBOX_CHUNK_SIZE]
    --dropbox-client-id value OAuth 客户端 ID。 [$DROPBOX_CLIENT_ID]
    --dropbox-client-secret value OAuth 客户端密钥。 [$DROPBOX_CLIENT_SECRET]
    --dropbox-encoding value 后端编码。（默认值:"Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$DROPBOX_ENCODING]
    --dropbox-impersonate value 在使用商务帐户时冒充此用户。 [$DROPBOX_IMPERSONATE]
    --dropbox-shared-files value 指示 rclone 处理单个共享文件。(默认值:"false") [$DROPBOX_SHARED_FILES]
    --dropbox-shared-folders value 指示 rclone 处理共享文件夹。(默认值:"false") [$DROPBOX_SHARED_FOLDERS]
    --dropbox-token value OAuth 访问令牌，Json 格式。 [$DROPBOX_TOKEN]
    --dropbox-token-url value 令牌服务器 URL。 [$DROPBOX_TOKEN_URL]
```
{% endcode %}