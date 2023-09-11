# IONOS Cloud

{% code fullWidth="true" %}
```
NAME:
    singularity 存储更新 s3 ionos - IONOS Cloud

用法: singularity 存储更新 s3 ionos [命令选项] <名称或ID>

描述:
    --env-auth
    从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
    
    仅当access_key_id和secret_access_key为空白时生效。

    示例:
        | false | 在下一步中输入AWS凭证。
        | true  | 从环境中获取AWS凭证（环境变量或IAM）。

    --access-key-id
    AWS访问密钥ID。
    
    留空以进行匿名访问或使用运行时凭证。

    --secret-access-key
    AWS秘密访问密钥（密码）。
    
    留空以进行匿名访问或使用运行时凭证。

    --region
    您的Bucket将被创建并存储数据的区域。
    

    示例:
        | de           | 德国法兰克福
        | eu-central-2 | 德国柏林
        | eu-south-2   | 西班牙洛格罗尼奥

    --endpoint
    IONOS S3对象存储的端点。
    
    指定来自同一区域的端点。

    示例:
        | s3-eu-central-1.ionoscloud.com | 德国法兰克福
        | s3-eu-central-2.ionoscloud.com | 德国柏林
        | s3-eu-south-2.ionoscloud.com   | 西班牙洛格罗尼奥

    --acl
    创建存储桶和存储或复制对象时使用的预定义ACL。
    
    此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建存储桶。
    
    有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
    
    请注意，当在服务器端复制对象时，将应用此ACL，因为S3不会复制源对象的ACL，而是写入新的 ACL。
    
    如果ACL是一个空字符串，那么将不会添加X-Amz-Acl: 头，并且将使用默认 (private)。

    --bucket-acl
    创建存储桶时使用的预定义ACL。
    
    有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
    
    请注意，仅在创建存储桶时应用此ACL。 如果未设置它，则使用 "acl" 作为替代。
    
    如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加X-Amz-Acl: 头，将使用默认 (private)。

    示例:
        | private            | 所有者获得FULL_CONTROL权限。
        |                    | 没有其他人拥有访问权限（默认）。
        | public-read        | 所有者获得FULL_CONTROL权限。
        |                    | AllUsers组获得读取权限。
        | public-read-write  | 所有者获得FULL_CONTROL权限。
        |                    | AllUsers组获得读取和写入权限。
        |                    | 通常不建议在存储桶上授予此权限。
        | authenticated-read | 所有者获得FULL_CONTROL权限。
        |                    | AuthenticatedUsers组获得读取权限。

    --upload-cutoff
    切换到分块上传的截止点。
    
    大于此大小的任何文件将以 chunk_size 的大小分块上传。
    最小值为0，最大值为5 GiB。

    --chunk-size
    用于上传的分块大小。
    
    当上传大于upload_cutoff的文件或大小未知的文件（例如来自 "rclone rcat" 或使用 "rclone mount" 或谷歌照片或谷歌文档上传的文件）时，将使用此分块大小进行分块上传。
    
    请注意，每个传输将在内存中缓冲 "--s3-upload-concurrency" 个 chunks。
    
    如果您正在通过高速链接传输大文件，并且拥有足够的内存，增加此值将加快传输速度。
    
    Rclone会自动增加分块大小，以保持在10,000个分块的限制之下上传大文件的策略。
    
    未知大小的文件将使用配置的分块大小上传。由于默认分块大小为5 MiB，并且最多可以有10,000个分块，因此默认情况下，可以流式上传的文件的最大大小为48 GiB。 如果您想流式上传更大的文件，则需要增加 chunk_size。
    
    增加分块大小会降低使用 "-P" 标志显示的进度统计的准确性。 当Rclone将块标记为已发送时，它是通过AWS SDK缓冲的，只有在块传输完成后，它才会实际上已经上传。 更大的分块大小意味着更大的AWS SDK缓冲区，进度报告更可能与实际情况有所不同。

    --max-upload-parts
    多部分上传中的最大部分数。
    
    此选项定义了在执行多部分上传时使用的最大多部分块数。
    
    如果服务不支持AWS S3规范的10,000个分块，则可以使用此选项。
    
    Rclone会自动增加分块大小，以保持在此分块数限制之下上传已知大小的大文件。
    
    --copy-cutoff
    切换到分块复制的截止点。
    
    需要服务器端复制的大于此大小的任何文件将以此大小的分块复制。

    --disable-checksum
    不在对象元数据中存储MD5校验和。
    
    通常，rclone会计算上传文件的MD5校验和，以便将其添加到对象的元数据中。 这对于数据完整性检查非常有用，但对于大文件，开始上传可能会花费很长时间。

    --shared-credentials-file
    共享凭证文件的路径。
    
    如果 env_auth = true，则rclone可以使用共享凭证文件。
    
    如果此变量为空，则rclone将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。 如果环境变量的值为空，则默认为当前用户的主目录。

        Linux/OSX: "$HOME/.aws/credentials"
        Windows:   "%USERPROFILE%\.aws\credentials"

    --profile
    在共享凭证文件中使用的配置文件。
    
    如果 env_auth = true，则rclone可以使用共享凭证文件。 此变量控制在该文件中使用哪个配置文件。
    
    如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default" （如果该环境变量也未设置）。

    --session-token
    AWS会话令牌。

    --upload-concurrency
    进行多部分上传的并发数。
    
    这是同时上传的相同文件的块数。
    
    如果您通过高速链接上传较少的大文件，并且这些上传没有完全使用您的带宽，那么增加此数值可能有助于加快传输速度。

    --force-path-style
    如果设置为true，使用路径样式访问；如果设置为false，使用虚拟托管样式访问。
    
    如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式。 有关更多信息，请参阅 [AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

    --v2-auth
    如果为true，则使用v2认证。

    如果为false（默认值），则rclone将使用v4认证。 如果设置了它，rclone将使用v2认证。

    仅在v4签名无效时使用，例如早于 Jewel/v10 CEPH 的版本。

    --list-chunk
    列举的块大小（每个 ListObject S3 请求的响应列表）。
    
    此选项也称为"MaxKeys"、"max-items"或"page-size"，在AWS S3规范中。
    大多数服务将列表响应截断为1000个对象，即使请求的数量超过了这个数字。
    在AWS S3中，这是一个全局最大值，无法更改，参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
    在Ceph中，可以使用"rgw list buckets max chunk"选项提高此值。

    --list-version
    使用的ListObjects版本：1、2或0（自动）。
    
    当S3首次推出时，它仅提供了用于枚举存储桶中的对象的ListObjects调用。
    
    但是，在2016年5月，引入了ListObjectsV2调用。 这个调用的性能更高，应该尽可能使用。
    
    如果设置为默认值0，则rclone将根据设置的提供商猜测使用哪个列表对象方法进行调用。 如果猜测错误，则可以在此手动设置。

    --list-url-encode
    是否对列表进行url编码：true/false/unset
    
    一些提供商支持URL编码列表，在文件名中使用控制字符时，这是更可靠的方法。 如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的编码，但您可以在此处覆盖rclone的选择。

    --no-check-bucket
    如果设置，不会尝试检查Bucket是否存在或创建它。
    
    当尝试最小化rclone在您已知Bucket已经存在时执行的事务数量时，这很有用。
    
    如果您使用的用户没有创建Bucket的权限，则可能需要使用此选项。 在v1.52.0之前，由于错误，此选项将被静默通过。

    --no-head
    如果设置，则不会执行HEAD请求来检查完整性。
    
    这在尝试最小化rclone执行的事务数量时非常有用。
    
    设置它意味着如果在使用PUT上传对象后收到200 OK消息，则rclone将假定它已正确上传。

    特别是它将假设：
    
    - 元数据，包括修改时间、存储类和内容类型与上传的相同
    - 大小与上传的相同
    
    它从单个部分PUT的响应中读取以下项目：
    
    - MD5SUM
    - 已上传日期
    
    对于多部分上传，不会读取这些项目。
    
    如果上传具有未知长度的源对象，则rclone **将**执行 HEAD 请求。
    
    设置此标志会增加无法检测到的上传失败的可能性，特别是大小不正确的情况，默认情况下不建议正常操作。 实际上，即使有此标志，上传失败的机会非常小。

    --no-head-object
    如果设置，则在获取对象时不执行HEAD请求。

    --encoding
    后端的编码方式。
    
    有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

    --memory-pool-flush-time
    内部内存缓冲池刷新的时间间隔。
    
    需要额外缓冲区（例如分片上传）的上传将使用内存池进行分配。
    此选项控制多久将从池中删除未使用的缓冲区。

    --memory-pool-use-mmap
    是否在内部内存池中使用mmap缓冲区。

    --disable-http2
    禁用S3后端的http2使用。
    
    s3（特别是minio）后端与HTTP/2目前存在未解决的问题。 S3后端默认启用HTTP/2，但可以在此禁用。 解决问题后，将删除此标志。
    
    请参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

    --download-url
    下载的自定义端点。
    通常设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的较低出口流量费用。

    --use-multipart-etag
    是否在分块上传中使用ETag进行验证。
    
    此值应为true、false或留空以使用提供商的默认值。

    --use-presigned-request
    是否使用预签名请求或PutObject进行单部分上传。
    
    如果此值为false，则rclone将使用AWS SDK的PutObject上传对象。
    
    rclone版本 < 1.59使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。 除非在特殊情况下或进行测试，否则不应使用。

    --versions
    在目录列表中包含旧版本。

    --version-at
    按指定时间显示文件版本。
    
    参数应为日期，"2006-01-02"，日期时间 "2006-01-02 15:04:05" 或之前的时间段，例如 "100d" 或 "1h"。
    
    请注意，使用此功能时，不允许进行文件写入操作，因此无法上传文件或删除文件。
    
    有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

    --decompress
    如果设置，将解压缩gzip编码的对象。
    
    可以使用 "Content-Encoding: gzip" 将对象上传到S3。 通常，rclone会将这些文件作为压缩对象下载。
    
    如果设置了此标志，则rclone会在接收到 "Content-Encoding: gzip" 的文件时解压缩文件。 这意味着rclone无法检查文件的大小和哈希值，但文件内容将被解压缩。

    --might-gzip
    如果后端可能压缩对象，则设置此值。
    
    通常情况下，提供商在下载对象时不会更改对象。 如果一个对象没有使用 `Content-Encoding: gzip` 进行上传，那么在下载时也不会设置。
    
    但是，一些提供商可能会压缩对象，即使它们没有使用 `Content-Encoding: gzip` 进行上传（例如Cloudflare）。
    
    如果设置了此标志，并且rclone下载具有设置 `Content-Encoding: gzip` 和分块传输编码的对象，那么rclone将会实时解压缩该对象。
    
    如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

    --no-system-metadata
    禁止设置和读取系统元数据。

选项:
    --access-key-id value      AWS访问密钥ID。 [$ACCESS_KEY_ID]
    --acl value                创建存储桶和存储或复制对象时使用的预定义ACL。 [$ACL]
    --endpoint value           IONOS S3对象存储的端点。 [$ENDPOINT]
    --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
    --help, -h                 显示帮助
    --region value             您的Bucket将被创建并存储数据的区域。 [$REGION]
    --secret-access-key value  AWS秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

    高级

    --bucket-acl value               创建存储桶时使用的预定义ACL。 [$BUCKET_ACL]
    --chunk-size value               用于上传的分块大小。 (default: "5Mi") [$CHUNK_SIZE]
    --copy-cutoff value              切换到分块复制的截止点。 (default: "4.656Gi") [$COPY_CUTOFF]
    --decompress                     如果设置，将解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
    --disable-checksum               不在对象元数据中存储MD5校验和。 (default: false) [$DISABLE_CHECKSUM]
    --disable-http2                  禁用S3后端的http2使用。 (default: false) [$DISABLE_HTTP2]
    --download-url value             下载的自定义端点。 [$DOWNLOAD_URL]
    --encoding value                 后端的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
    --force-path-style               如果设置为true，使用路径样式访问；如果设置为false，使用虚拟托管样式访问。 (default: true) [$FORCE_PATH_STYLE]
    --list-chunk value               列举的块大小（每个ListObject S3请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
    --list-url-encode value          是否对列表进行url编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
    --list-version value             使用的ListObjects版本：1、2或0（自动）。 (default: 0) [$LIST_VERSION]
    --max-upload-parts value         多部分上传中的最大部分数。 (default: 10000) [$MAX_UPLOAD_PARTS]
    --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
    --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
    --might-gzip value               如果后端可能压缩对象，则设置此值，(default: "unset") [$MIGHT_GZIP]
    --no-check-bucket                如果设置，不会尝试检查Bucket是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
    --no-head                        如果设置，则不会执行HEAD请求来检查完整性。 (default: false) [$NO_HEAD]
    --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (default: false) [$NO_HEAD_OBJECT]
    --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
    --profile value                  在共享凭证文件中使用的配置文件。 [$PROFILE]
    --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
    --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
    --upload-concurrency value       进行多部分上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
    --upload-cutoff value            切换到分块上传的截止点。 (default: "200Mi") [$UPLOAD_CUTOFF]
    --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
    --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
    --v2-auth                        如果为true，则使用v2认证 (default: false) [$V2_AUTH]
    --version-at value               按指定时间显示文件版本。 (default: "off") [$VERSION_AT]
    --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]

```
{% endcode %}