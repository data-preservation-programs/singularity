# Liara对象存储

{% code fullWidth="true" %}
```
名称：
    singularity storage create s3 liara - Liara对象存储

用法：
    singularity storage create s3 liara [命令选项] [参数...]

说明：
    --env-auth
        从运行环境获取AWS凭证（如果访问密钥和秘密访问密钥为空，则从环境变量或EC2/ECS元数据获取）。
        
        仅适用于访问密钥ID和秘密访问密钥为空的情况。

        示例：
            | false | 在下一步输入AWS凭证。
            | true  | 从环境获取AWS凭证（环境变量或IAM）。

    --access-key-id
        AWS访问密钥ID。
        
        留空以进行匿名访问或运行时凭证。

    --secret-access-key
        AWS秘密访问密钥（密码）。
        
        留空以进行匿名访问或运行时凭证。

    --endpoint
        Liara对象存储API的终端节点。

        示例：
            | storage.iran.liara.space | 默认终端节点
            |                          | 伊朗

    --acl
        创建桶和存储或复制对象时使用的预设ACL。
        
        此ACL用于创建对象，并且如果未设置bucket_acl，则还用于创建桶。
        
        有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
        
        请注意，S3在服务器端复制对象时应用此ACL，
        不会从源复制ACL，而是编写新的ACL。
        
        如果ACL是空字符串，则不会添加X-Amz-Acl：标题，并且将使用默认值（私有）。

    --bucket-acl
        创建存储桶时要使用的预设ACL。
        
        有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
        
        请注意，仅在创建存储桶时应用此ACL。如果未设置，则会使用“acl”。
        
        如果“acl”和“bucket_acl”是空字符串，则不会添加X-Amz-Acl：
        标题，并且将使用默认值（私有）。

        示例：
            | private            | 所有者拥有完全控制权。
            |                    | 没有其他人有访问权限（默认）。
            | public-read        | 所有者拥有完全控制权。
            |                    | AllUsers组具有读取访问权限。
            | public-read-write  | 所有者拥有完全控制权。
            |                    | AllUsers组具有读取和写入访问权限。
            |                    | 通常不推荐在存储桶上授予此权限。
            | authenticated-read | 所有者拥有完全控制权。
            |                    | AuthenticatedUsers组具有读取访问权限。

    --storage-class
        在Liara中存储新对象时要使用的存储类别。

        示例：
            | STANDARD | 标准存储类别

    --upload-cutoff
        切换为分块上传的截止点。
        
        大于此大小的文件将分块上传，每块大小为chunk_size。
        最小值为0，最大值为5 GiB。

    --chunk-size
        用于上传的块大小。
        
        当上传大于upload_cutoff的文件或大小未知的文件（例如从“rclone rcat”上传或使用“rclone mount”或谷歌
        照片或谷歌文档上传的文件）时，
        将使用此块大小进行分块上传。
        
        请注意，每次传输内存中缓冲的是“--s3-upload-concurrency”大小的块。
        
        如果您正在通过高速链路传输大文件，并且有足够的内存，
        那么增加这个值将加快传输速度。
        
        当上传已知大小的大文件时，rclone会自动增加块大小，
        以确保保持在10000个块的限制以下。
        
        未知大小的文件使用配置的
        块大小进行上传。由于默认块大小为5 MiB，
        并且最多可能有10000个块，
        这意味着默认情况下可以流式传输的文件的最大大小为48 GiB。
        如果要流式传输更大的文件，则需要增加块大小。
        
        增加块大小会降低用“-P”标志显示的进度
        统计的准确性。当rclone将块作为已发送时，
        它是由AWS SDK缓冲的，实际上它可能仍在上传。
        更大的块大小意味着更大的AWS SDK缓冲区和进度
        报告与实际情况更偏离。
        

    --max-upload-parts
        多部分上传中的最大部分数。
        
        此选项定义在执行多部分上传时要使用的最大多部分块数。
        
        如果某个服务不支持AWS S3规范中的10000个多部分块，
        这将非常有用。
        
        当上传已知大小的大文件时，rclone会自动增加块大小，
        以保持低于这个块数限制。
        

    --copy-cutoff
        切换为分块复制的截止点。
        
        需要进行服务端复制的大于此大小的文件将以此大小的块复制。
        
        最小值为0，最大值为5 GiB。

    --disable-checksum
        不要将MD5校验和存储到对象元数据中。
        
        通常rclone会在上传之前计算输入的MD5校验和，
        以便可以将其添加到对象的元数据中。这对于数据完整性
        检查非常有用，但对于大文件启动上传可能会导致长时间的延迟。

    --shared-credentials-file
        共享凭证文件的路径。
        
        如果env_auth = true，则rclone可以使用共享凭证文件。
        
        如果此变量为空，则rclone将查找
        “AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量值为空，
        则默认为当前用户的主目录。
        
            Linux/OSX: “$HOME/.aws/credentials”
            Windows:   “%USERPROFILE%\.aws\credentials”
        

    --profile
        共享凭证文件中要使用的配置文件。
        
        如果env_auth = true，则rclone可以使用共享凭证文件。此
        变量控制在该文件中使用哪个配置文件。
        
        如果为空，则默认为环境变量“AWS_PROFILE”或
        如果该环境变量也未设置，则为“默认值”。
        

    --session-token
        AWS会话令牌。

    --upload-concurrency
        多部分上传的并发数。
        
        这是并行上传相同文件的块数。
        
        如果您在高速链路上上传大量大文件，
        并且这些上传未充分利用您的带宽，
        那么增加这个值可能有助于加快传输速度。

    --force-path-style
        如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。
        
        如果为true（默认值），则rclone将使用路径样式访问，
        如果为false，则rclone将使用虚拟路径样式。有关更多信息，请
        参阅[ AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
        
        某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将其设置为
        false - rclone将根据提供商的设置自动执行此操作。

    --v2-auth
        如果为true，则使用v2身份验证。
        
        如果为false（默认值），则rclone将使用v4身份验证。
        如果设置了它，则rclone将使用v2身份验证。
        
        仅在v4签名不起作用时使用此选项，
        例如早期版本的Jewel/v10 CEPH。

    --list-chunk
        列表块的大小（每个ListObject S3请求的响应列表）。
        
        此选项也称为“MaxKeys”，“max-items”或“page-size”（AWS S3规范）。
        大多数服务将响应列表截断为1000个对象，即使请求超过该数量。
        在AWS S3中，这是一个全局最大值，并且无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
        在Ceph中，可以通过“rgw list buckets max chunk”选项来增加此数值。
        

    --list-version
        要使用的ListObjects的版本：1、2或0（自动）。
        
        在S3最初启动时，只提供了用于枚举存储桶中对象的ListObjects调用。
        
        但是在2016年5月，引入了ListObjectsV2调用。这个调用
        性能更高，如果可能的话应该使用。
        
        如果设置为默认值0，则rclone将根据提供商设置猜测要调用的列表对象方法。
        如果猜测错误，则可以在此处手动设置。
        

    --list-url-encode
        是否对列表进行URL编码：true/false/unset
        
        某些提供商支持对列表进行URL编码，
        当文件名中使用控制字符时，这更可靠。如果将其设置为
        unset（默认值），则rclone将根据提供商设置选择要应用的设置，但是您可以在此处覆盖
        rclone的选择。
        

    --no-check-bucket
        如果设置，则不要尝试检查桶是否存在或创建桶。
        
        如果您知道桶已经存在，这样可以最小化rclone操作的数量。
        
        如果使用的用户没有桶创建权限，这也可能是必需的。
        在 v1.52.0 之前，由于一个错误，这将会静默通过。
        

    --no-head
        如果设置，则不要对已上传的对象进行HEAD请求以检查完整性。
        
        如果设置了这个标志，那么如果rclone在PUT之后收到200 OK消息，
        则会假设它已正确上传。
        
        特别是，它会假设：
        
        - 元数据，包括修改时间、存储类别和内容类型与上传的相同
        - 大小与上传的相同
          
        它从单个部分PUT的响应中读取以下内容：
         
        - MD5SUM 哈希值
        - 上传日期
          
        对于多部分上传，不会读取这些项。
        
        如果上传未知长度的源对象，则 rclone **将**执行 HEAD 请求。
        
        设置此标志会增加检测不到的上传故障的几率，
        特别是错误的大小，因此不建议在常规操作中使用。
        实际上，即使带有此标志，检测不到的上传故障的几率也非常小。
        

    --no-head-object
        如果设置，则在获取对象时不执行 HEAD 请求。

    --encoding
        后端的编码方式。
        
        有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

    --memory-pool-flush-time
        内部内存缓冲区池的定期刷新时间。
        
        需要额外缓冲区的上传（例如，多部分上传）将使用内存池进行分配。
        此选项控制池中未使用的缓冲区将被移除的频率。

    --memory-pool-use-mmap
        是否在内部内存池中使用mmap缓冲区。

    --disable-http2
        禁用S3后端的http2使用。
        
        目前，s3（特别是minio）后端存在尚未解决的问题
        并且针对HTTP/2使能。S3 后端的 HTTP/2 默认启用，但可以在此处禁用。
        在问题解决之后，此标志将被删除。
        
        参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

    --download-url
        下载的自定义端点。
        此设置通常用于CloudFront CDN URL，因为AWS S3通过
        CloudFront网络下载的数据提供更便宜的出站流量。

    --use-multipart-etag
        是否在多部分上传中使用ETag进行校验
        
        这个值应该是true、false或未设置，以使用提供者的默认值。
        

    --use-presigned-request
        是否使用预签名请求或PutObject进行单部分上传。
        
        如果为false，则rclone将使用AWS SDK的PutObject上传对象。
        
        rclone的版本 < 1.59 使用预签名请求上传单个
        部分对象，并将此标志设置为true将重新启用该功能。
        这除非在特殊情况下或用于测试，否则不应该是必需的。
        

    --versions
        在目录列表中包括旧版本。

    --version-at
        显示文件版本，如指定的时间所示。
        
        参数应为日期，“2006-01-02”，日期时间 “2006-01-02
        15:04:05”或那么久之前的时间，例如 “100d” 或 “1h”。
        
        请注意，使用此选项时，不允许进行文件写入操作，
        因此无法上传文件或删除文件。
        
        有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

    --decompress
        如果设置，将解压缩gzip编码的对象。
        
        可以使用“Content-Encoding: gzip”将对象上传到S3。
        通常情况下，rclone会将这些文件下载为压缩对象。
        
        如果设置了此标志，则rclone将在接收到包含
        “Content-Encoding: gzip”的文件时对这些文件进行解压缩。
        这意味着rclone无法检查大小和哈希值，
        但文件内容将被解压缩。

    --might-gzip
        如果后端可能会gzip对象，则设置此标志。
        
        通常情况下，提供者在下载时不会更改对象。如果
        对象未使用“Content-Encoding: gzip”上传，
        在下载时也不会设置它。
        
        但是，有些提供者即使在未使用`Content-Encoding: gzip`上传
        对象的情况下也可能gzip对象（例如Cloudflare）。
        
        接收到以下错误时的症状可能是：
        
            ERROR corrupted on transfer: sizes differ NNN vs MMM
        
        如果设置了此标志，并且rclone下载了设置了
        Content-Encoding：gzip和分块传输编码的对象，
        那么rclone会即时解压缩对象。
        
        如果这设置为unset（默认值），则rclone将根据提供商的设置选择
        要应用的设置，但是您可以在此处覆盖rclone的选择。
        

    --no-system-metadata
        禁止设置和读取系统元数据


选项：
    --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
    --acl value                创建桶和存储或复制对象时使用的预设ACL。[$ACL]
    --endpoint value           Liara对象存储API的终端节点。[$ENDPOINT]
    --env-auth                 从运行环境获取AWS凭证（如果访问密钥和秘密访问密钥为空，则从环境变量或EC2/ECS元数据获取）。（默认值：false）[$ENV_AUTH]
    --help, -h                 显示帮助信息
    --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
    --storage-class value      在Liara中存储新对象时要使用的存储级别。[$STORAGE_CLASS]

    高级选项：

    --bucket-acl value             创建存储桶时要使用的预设ACL。[$BUCKET_ACL]
    --chunk-size value             用于上传的块大小。（默认值：“5Mi”）[$CHUNK_SIZE]
    --copy-cutoff value            切换为分块复制的截止点。（默认值：“4.656Gi”）[$COPY_CUTOFF]
    --decompress                   如果设置此标志，将解压缩gzip编码的对象。（默认值：false）[$DECOMPRESS]
    --disable-checksum             不要将MD5校验和存储到对象元数据中。（默认值：false）[$DISABLE_CHECKSUM]
    --disable-http2                禁用S3后端的http2使用。（默认值：false）[$DISABLE_HTTP2]
    --download-url value           下载的自定义端点。[$DOWNLOAD_URL]
    --encoding value               后端的编码方式。（默认值：“Slash,InvalidUtf8,Dot”）[$ENCODING]
    --force-path-style             如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。（默认值：true）[$FORCE_PATH_STYLE]
    --list-chunk value             列表块的大小（每个ListObject S3请求的响应列表）。（默认值：1000）[$LIST_CHUNK]
    --list-url-encode value        是否对列表进行URL编码：true/false/unset（默认值：“unset”）[$LIST_URL_ENCODE]
    --list-version value           要使用的ListObjects的版本：1、2或0（自动）。 （默认值：0）[$LIST_VERSION]
    --max-upload-parts value       多部分上传中的最大部分数。（默认值：10000）[$MAX_UPLOAD_PARTS]
    --memory-pool-flush-time value 内部内存缓冲区池的定期刷新时间。（默认值：“1m0s”）[$MEMORY_POOL_FLUSH_TIME]
    --memory-pool-use-mmap         是否在内部内存池中使用mmap缓冲区。（默认值：false）[$MEMORY_POOL_USE_MMAP]
    --might-gzip value             设置此标志以指示后端可能会gzip对象。（默认值：“unset”）[$MIGHT_GZIP]
    --no-check-bucket              如果设置，则不要尝试检查桶是否存在或创建桶。（默认值：false）[$NO_CHECK_BUCKET]
    --no-head                      如果设置，则不要对已上传的对象进行HEAD请求以检查完整性。（默认值：false）[$NO_HEAD]
    --no-head-object               如果设置，则在获取对象时不执行HEAD请求。（默认值：false）[$NO_HEAD_OBJECT]
    --no-system-metadata           禁止设置和读取系统元数据（默认值：false）[$NO_SYSTEM_METADATA]
    --profile value                共享凭证文件中要使用的配置文件。[$PROFILE]
    --session-token value          AWS会话令牌。[$SESSION_TOKEN]
    --shared-credentials-file value 共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
    --upload-concurrency value     多部分上传的并发数。（默认值：4）[$UPLOAD_CONCURRENCY]
    --upload-cutoff value          切换为分块上传的截止点。（默认值：“200Mi”）[$UPLOAD_CUTOFF]
    --use-multipart-etag value     是否在多部分上传中使用ETag进行校验（默认值：“unset”）[$USE_MULTIPART_ETAG]
    --use-presigned-request        是否使用预签名请求或PutObject进行单部分上传（默认值：false）[$USE_PRESIGNED_REQUEST]
    --v2-auth                      如果为true，则使用v2身份验证。（默认值：false）[$V2_AUTH]
    --version-at value             显示文件版本，如指定的时间所示。 （默认值：“off”）[$VERSION_AT]
    --versions                     在目录列表中包括旧版本。（默认值：false）[$VERSIONS]

    通用选项：

    --name value  存储的名称（默认值：自动生成）
    --path value  存储的路径

```
{% endcode %}