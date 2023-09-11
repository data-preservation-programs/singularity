# Cloudflare R2 存储

{% code fullWidth="true" %}
```
NAME:
    singularity storage create s3 cloudflare - Cloudflare R2 存储

USAGE:
    singularity storage create s3 cloudflare [命令选项] [参数...]

DESCRIPTION:
    --env-auth
        从运行环境获取 AWS 凭证 (环境变量或 EC2/ECS 元数据，如果没有环境变量)。

        仅当 access_key_id 和 secret_access_key 为空时生效。

        示例:
        | false | 在下一步中输入 AWS 凭证。
        | true  | 从环境中获取 AWS 凭证 (环境变量或 IAM)。

    --access-key-id
        AWS 访问密钥 ID。

        如果要进行匿名访问或使用运行时凭证，请留空。

    --secret-access-key
        AWS 秘密访问密钥（密码）。

        如果要进行匿名访问或使用运行时凭证，请留空。

    --region
        要连接的区域。

        示例:
        | auto | R2 存储桶会自动分布在 Cloudflare 的数据中心以实现低延迟。

    --endpoint
        S3 API 的终端节点。

        使用 S3 克隆时必填。

    --bucket-acl
        创建存储桶时使用的预设 ACL。

        更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

        注意，该 ACL 仅在创建存储桶时应用。如果未设置，则使用 "acl"。

        如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加 X-Amz-Acl: 头，并将使用默认值（私有）。

        示例:
        | private            | 拥有者具有 FULL_CONTROL 权限。
        |                    | 无其他人有访问权限（默认值）。
        | public-read        | 拥有者具有 FULL_CONTROL 权限。
        |                    | AllUsers 组具有 READ 权限。
        | public-read-write  | 拥有者具有 FULL_CONTROL 权限。
        |                    | AllUsers 组具有 READ 和 WRITE 权限。
        |                    | 通常不建议在存储桶上授予该权限。
        | authenticated-read | 拥有者具有 FULL_CONTROL 权限。
        |                    | AuthenticatedUsers 组具有 READ 权限。

    --upload-cutoff
        切换为分块上传的文件截止大小。

        大于该大小的文件将分块上传，每块大小为 chunk_size。
        最小值为 0，最大值为 5 GiB。

    --chunk-size
        上传时使用的块大小。

        对于大于 upload_cutoff 的文件或大小未知的文件（例如来自 "rclone rcat"、"rclone mount"、Google 照片或 Google 文档），将使用此块大小进行分块上传。

        注意，每个传输的内存中会缓冲 "--s3-upload-concurrency" 个此大小的块。

        如果您正在通过高速链接传输大型文件，并且有足够的内存，则增加此值将加速传输。

        Rclone 会自动增加块大小，以保持在 10,000 块的限制之下，以上传已知大小的大型文件。

        未知大小的文件使用配置的块大小进行上传。由于默认的块大小为 5 MiB，最多可以有 10,000 个块，这意味着默认情况下，您可以流式传输的文件的最大大小为 48 GiB。如果您希望流式传输更大的文件，您需要增加块大小。

        增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。当 Rclone 缓冲 AWS SDK 时，它将发送给块，而实际上它可能仍在上传。较大的块大小意味着较大的 AWS SDK 缓冲区，并导致进度报告更远离实际情况。

    --max-upload-parts
        多块上传中使用的最大块数。

        此选项定义执行多块上传时要使用的最大分块数。

        如果服务不支持 AWS S3 规范的 10,000 块，则此选项可能很有用。

        Rclone 会自动增加块大小，以保持在此分块数限制之下，以上传已知大小的大型文件。

    --copy-cutoff
        切换为分块复制的文件截止大小。

        大于该大小的需要进行服务器端复制的文件将按照此大小分块复制。

        最小值为 0，最大值为 5 GiB。

    --disable-checksum
        不将 MD5 校验和与对象元数据一起存储。

        通常，rclone 会在上传之前计算输入文件的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查很有帮助，但会导致大文件开始上传时出现长时间的延迟。

    --shared-credentials-file
        共享凭证文件的路径。

        如果 env_auth = true，则 rclone 可以使用共享凭证文件。

        如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果 env 值为空，则默认为当前用户的主目录。

            Linux/OSX: "$HOME/.aws/credentials"
            Windows:   "%USERPROFILE%\.aws\credentials"

    --profile
        共享凭证文件中要使用的配置文件。

        如果 env_auth = true，则 rclone 可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。

        如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果该环境变量也未设置）。

    --session-token
        AWS 会话令牌。

    --upload-concurrency
        多块上传的并发度。

        这是同时上传的相同文件的块数。

        如果您通过高速链接上传较少数量的大型文件，而这些上传未完全利用您的带宽，则增加此值可能有助于加快传输速度。

    --force-path-style
        如果设置为 true，则使用路径样式访问；如果设置为 false，则使用虚拟主机样式访问。

        如果设置为 true（默认值），则 rclone 将使用路径样式访问；如果设置为 false，则 rclone 将使用虚拟路径样式。有关详细信息，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

        某些提供商（例如 AWS、Aliyun OSS、Netease COS 或 Tencent COS）需要将此设置为 false，rclone 将根据提供商设置自动完成此操作。

    --v2-auth
        如果设置为 true，则使用 v2 认证。

        如果设置为 false（默认值），则 rclone 将使用 v4 认证。如果设置了该值，则 rclone 将使用 v2 认证。

        仅在 v4 签名无效时使用，例如早期版本的 Jewel/v10 CEPH。

    --list-chunk
        目录清单的大小（每个 ListObject S3 请求的响应列表）。

        这个选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
        大多数服务即使请求的数量超过 1000，仍会截断响应列表为 1000 个对象。
        在 AWS S3 中，这是一个全局最大值，无法更改，请参阅 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
        在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项进行增加。

    --list-version
        要使用的 ListObjects 版本: 1、2 或 0（自动）。

        当 S3 最初推出时，它只提供了 ListObjects 调用以枚举存储桶中的对象。

        但是，在 2016 年 5 月，ListObjectsV2 调用被引入。这是更高性能的方法，如果可能的话应该使用。

        如果设置为默认值 0，则 rclone 将根据设置的提供商来猜测要调用哪个列举对象方法。如果它的猜测结果错误，则可以在此处手动设置。

    --list-url-encode
        是否对列表进行 URL 编码：true/false/unset

        如果提供者支持 URL 编码列表，并且可用时，使用控制字符的文件名更可靠。如果设置为 unset（默认值），那么 rclone 将根据提供者设置来进行选择，但是您可以在此处覆盖 rclone 的选择。

    --no-check-bucket
        如果设置了此标志，则不尝试检查存储桶是否存在或创建它。

        如果您知道存储桶已经存在，当尝试将 rclone 的交易数最小化时，此选项可能很有用。

        如果使用的用户没有存储桶创建权限，则可能需要此选项。在 v1.52.0 之前，由于一个错误，此操作将悄无声息地传递。

    --no-head
        如果设置了此标志，则不会使用 HEAD 操作来检查已上传对象的完整性。

        如果最小化 rclone 的交易数量很重要，此选项可能很有用。

        设置此标志后，如果 rclone 在 PUT 操作后收到 200 OK 消息，则会假定对象已正确上传。

        特别地，它假设：
        - 元数据（包括修改时间、存储类别和内容类型）与上传的内容相同。
        - 大小与上传的内容相同。

        它从单个部分 PUT 的响应中读取以下项目：
        - MD5SUM
        - 上传日期

        对于多部分上传，不会读取这些项目。

        如果上传源对象的长度未知，则 rclone **将**执行 HEAD 请求。

        设置此标志会增加检测不到的上传失败的机会，尤其是大小不正确的机会，因此不建议在正常操作中使用。实际上，即使设置了此标志，检测不到的上传失败的几率也非常小。

    --no-head-object
        如果设置了此标志，则获取对象之前不会执行 HEAD 操作。

    --encoding
        后端的编码方式。

        查看 [概览中的编码部分](/overview/#encoding) 以获取更多信息。

    --memory-pool-flush-time
        内部内存缓冲池刷新的频率。

        需要额外缓冲区的上传（例如分块上传）将使用内存池进行内存分配。
        此选项控制在内存池中何时删除未使用的缓冲区。

    --memory-pool-use-mmap
        是否在内部内存池中使用 mmap 缓冲区。

    --disable-http2
        禁用 S3 后端的 http2 使用。

        当前 s3（特别是 minio）后端存在问题，并且没有解决。s3 后端默认启用 HTTP/2，但可以在此处禁用。解决问题后，将删除此标志。

        参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

    --download-url
        下载的自定义终端节点。
        这通常设置为 CloudFront CDN 的 URL，因为 AWS S3 通过 CloudFront 网络下载数据的数据出站费用较低。

    --use-multipart-etag
        是否对分块上传使用 ETag 进行验证

        这应该是 true、false 或留空以使用提供商的默认值。

    --use-presigned-request
        是否使用预签名请求或 PutObject 来进行单部分上传

        如果此为 false，则 rclone 将使用 AWS SDK 的 PutObject 来上传对象。

        Rclone 的版本 < 1.59 使用预签名请求来上传单个部分对象，将此标志设置为 true 将重新启用该功能。除非特殊情况或测试，否则不应该使用此功能。

    --versions
        在目录列表中包含旧版本。

    --version-at
        显示指定时间点的文件版本。

        参数应为日期，例如 "2006-01-02"，日期时间 "2006-01-02 15:04:05"，或距离很久之前的持续时间，例如 "100d" 或 "1h"。

        请注意，使用此选项时不允许执行文件写操作，因此无法上传文件或删除文件。

        有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

    --decompress
        如果设置了此标志，则会解压缩经过 gzip 编码的对象。

        可以将对象以 "Content-Encoding: gzip" 上传到 S3。通常情况下，rclone 会将这些文件作为压缩对象下载。

        如果设置了此标志，则 rclone 将在接收到这些文件时将其解压缩为 "Content-Encoding: gzip"。这意味着 rclone 无法检查大小和哈希，但文件内容将会被解压缩。

    --might-gzip
        如果后端可能会对对象进行 gzip 压缩，则设置此标志。

        通常，提供者在下载时不会更改对象。如果一个对象没有使用 `Content-Encoding: gzip` 上传，则在下载时也不会设置。

        但是，某些提供商甚至可能在没有使用 `Content-Encoding: gzip` 的情况下对对象进行压缩（例如 Cloudflare）。

        所以，如果设置了此标志并且 rclone 使用支持分块传输编码的 `Content-Encoding: gzip` 和 `chunked` 的对象，则 rclone 将会边下载边解压缩这些对象。

        如果设置为 unset（默认值），则 rclone 将根据提供商设置来进行选择，但您可以在此处覆盖 rclone 的选择。

    --no-system-metadata
        禁止设置和读取系统元数据


OPTIONS:
    --access-key-id value      AWS 访问密钥 ID。[$ACCESS_KEY_ID]
    --endpoint value           S3 API 的终端节点。[$ENDPOINT]
    --env-auth                 从运行时环境获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。(默认值: false)[$ENV_AUTH]
    --help, -h                 显示帮助信息
    --region value             要连接的区域。[$REGION]
    --secret-access-key value  AWS 秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

    高级

    --bucket-acl value               创建存储桶时使用的预设 ACL。[$BUCKET_ACL]
    --chunk-size value               上传时使用的块大小。 (默认值: "5Mi")[$CHUNK_SIZE]
    --copy-cutoff value              切换为分块复制的文件截止大小。 (默认值: "4.656Gi")[$COPY_CUTOFF]
    --decompress                     如果设置了此标志，则会解压缩经过 gzip 编码的对象。 (默认值: false)[$DECOMPRESS]
    --disable-checksum               不将 MD5 校验和与对象元数据一起存储。 (默认值: false)[$DISABLE_CHECKSUM]
    --disable-http2                  禁用 S3 后端的 http2 使用。 (默认值: false)[$DISABLE_HTTP2]
    --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
    --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot")[$ENCODING]
    --force-path-style               如果设置为 true，则使用路径样式访问；如果设置为 false，则使用虚拟主机样式访问。 (默认值: true)[$FORCE_PATH_STYLE]
    --list-chunk value               目录清单的大小（每个 ListObject S3 请求的响应列表）。 (默认值: 1000)[$LIST_CHUNK]
    --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (默认值: "unset")[$LIST_URL_ENCODE]
    --list-version value             要使用的 ListObjects 版本: 1、2 或 0（自动）。 (默认值: 0)[$LIST_VERSION]
    --max-upload-parts value         多块上传中使用的最大块数。 (默认值: 10000)[$MAX_UPLOAD_PARTS]
    --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s")[$MEMORY_POOL_FLUSH_TIME]
    --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值: false)[$MEMORY_POOL_USE_MMAP]
    --might-gzip value               设置此标志，如果后端可能会对对象进行 gzip 压缩。 (默认值: "unset")[$MIGHT_GZIP]
    --no-check-bucket                如果设置了此标志，则不尝试检查存储桶是否存在或创建它。 (默认值: false)[$NO_CHECK_BUCKET]
    --no-head                        如果设置了此标志，则不会使用 HEAD 操作来检查已上传对象的完整性。 (默认值: false)[$NO_HEAD]
    --no-head-object                 如果设置了此标志，则获取对象之前不会执行 HEAD 操作。 (默认值: false)[$NO_HEAD_OBJECT]
    --no-system-metadata             禁止设置和读取系统元数据 (默认值: false)[$NO_SYSTEM_METADATA]
    --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
    --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
    --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
    --upload-concurrency value       多块上传的并发度。 (默认值: 4)[$UPLOAD_CONCURRENCY]
    --upload-cutoff value            切换为分块上传的文件截止大小。 (默认值: "200Mi")[$UPLOAD_CUTOFF]
    --use-multipart-etag value       是否对分块上传使用 ETag 进行验证 (默认值: "unset")[$USE_MULTIPART_ETAG]
    --use-presigned-request          是否使用预签名请求或 PutObject 来进行单部分上传 (默认值: false)[$USE_PRESIGNED_REQUEST]
    --v2-auth                        如果设置为 true，则使用 v2 认证。 (默认值: false)[$V2_AUTH]
    --version-at value               显示指定时间点的文件版本。 (默认值: "off")[$VERSION_AT]
    --versions                       在目录列表中包含旧版本。 (默认值: false)[$VERSIONS]

    常规

    --name value  存储的名称（默认值: 自动生成的）
    --path value  存储的路径
```
{% endcode %}