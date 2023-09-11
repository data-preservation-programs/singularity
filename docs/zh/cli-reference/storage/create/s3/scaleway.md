# Scaleway对象存储

{% code fullWidth="true" %}
```
名称：
   singularity storage create s3 scaleway - Scaleway对象存储

用法：
   singularity storage create s3 scaleway [命令选项] [参数...]

描述：
   --env-auth
      从运行时获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅适用于access_key_id和secret_access_key为空白的情况。

      示例：
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境（环境变量或IAM）获取AWS凭据。

   --access-key-id
      AWS Access Key ID（访问密钥ID）。
      
      对于匿名访问或运行时凭据，请留空。

   --secret-access-key
      AWS Secret Access Key (密码)。
      
      对于匿名访问或运行时凭据，请留空。

   --region
      要连接的区域。

      示例：
         | nl-ams | 荷兰阿姆斯特丹
         | fr-par | 法国巴黎
         | pl-waw | 波兰华沙

   --endpoint
      Scaleway对象存储的终结点。

      示例：
         | s3.nl-ams.scw.cloud | 阿姆斯特丹终结点
         | s3.fr-par.scw.cloud | 巴黎终结点
         | s3.pl-waw.scw.cloud | 华沙终结点

   --acl
      创建桶和存储或复制对象时使用的默认ACL。
      
      此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建桶。
      
      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当S3服务器复制对象时，会应用此ACL，因为S3不复制源的ACL，而是写入新的ACL。
      
      如果ACL是一个空字符串，则不会添加X-Amz-Acl：header，并且将使用默认ACL（private）。

   --bucket-acl
      创建桶时使用的默认ACL。
      
      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建桶时应用此ACL。如果未设置，则使用“acl”替代。
      
      如果“acl”和“bucket_acl”是空字符串，则不会添加X-Amz-Acl：header，并且将使用默认ACL（private）。

      示例：
         | private            | 拥有者获取FULL_CONTROL。
         |                    | 没有其他用户有访问权限（默认）。
         | public-read        | 拥有者获取FULL_CONTROL。
         |                    | AllUsers组获取读取权限。
         | public-read-write  | 拥有者获取FULL_CONTROL。
         |                    | AllUsers组获取读取和写入权限。
         |                    | 通常不建议在操作存储桶时授予此权限。
         | authenticated-read | 拥有者获取FULL_CONTROL。
         |                    | AuthenticatedUsers组获取读取权限。

   --storage-class
      存储新对象时要使用的存储类。

      示例：
         | <unset>  | 默认值。
         | STANDARD | 标准类用于任何上传。
         |          | 适用于按需内容，如流媒体或CDN。
         | GLACIER  | 存档存储。
         |          | 价格较低，但需要首先恢复才能访问。

   --upload-cutoff
      切换到分块上传的文件切割点。
      
      大于此大小的文件将使用chunk_size分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传文件大于upload_cutoff或具有未知大小的文件（例如由"rclone rcat"或使用"rclone mount"或Google照片或Google文档上传）时，将使用此分块大小进行分块上传。
      
      请注意，“--s3-upload-concurrency”每个传输在内存中缓冲此大小的块。
      
      如果您正在高速链接上传输大型文件并且具有足够的内存，将此值增加将加快传输速度。
      
      Rclone将自动增加分块大小，以便在上传已知大小的大型文件时保持在10,000个分块限制以下。
      
      未知大小的文件将使用配置的分块大小进行上传。因为默认的块大小是5 MiB，最多可以有10,000个分块，这意味着默认情况下您可以流式传输的文件的最大大小为48 GiB。如果您希望流式传输更大的文件，则需要增加chunk_size。
      
      增加分块大小会降低使用"-P"标志显示的进度统计的准确性。当rclone将块缓冲到AWS SDK时，它会将块视为已发送，而实际上它可能仍在上传。更大的块大小意味着更大的AWS SDK缓冲区和进度报告与实际情况的偏离。
      

   --max-upload-parts
      分块上传中最多存在的分块数。
      
      该选项定义了执行分块上传时要使用的最大分块数。
      
      如果服务不支持AWS S3规范的10,000个分块，则此选项可能很有用。
      
      Rclone将自动增加分块大小，以便在上传已知大小的大型文件时保持在此分块数限制之下。
      

   --copy-cutoff
      切换到多部分复制的复制切割点。
      
      需要服务器端拷贝的大于此大小的文件将会被拷贝为此大小的分块。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入的MD5校验和，以便在对象的元数据中添加它。这对于数据完整性检查非常有用，但对于大型文件开始上传可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果env值为空，则它将使用当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"默认"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的同一文件的分块的数量。
      
      如果你通过高速链接上传大量的大型文件，并且这些上传没有充分利用你的带宽，那么增加这个值可能会帮助加快传输速度。

   --force-path-style
      如果设置为true，则使用路径样式访问；如果设置为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式。有关更多信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些服务提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求将此设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果设置为true，则使用v2身份验证；如果设置为false（默认值），则使用v4身份验证；如果设置，则rclone将使用v2身份验证。
      
      仅在v4签名无法使用时才使用此选项，例如早于Jewel/v10 CEPH。

   --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。
      大多数服务即使请求超过1000个对象也会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      在S3最初推出时，只提供了ListObjects调用来枚举存储桶中的对象。
      
      然而，在2016年5月，引入了ListObjectsV2调用。这是性能更高的版本，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪个列表对象方法。如果它猜错了，那么可以在此手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，在使用控制字符的文件名时，这更可靠。如果设置为未设置（默认值），那么rclone将根据提供者设置选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不会尝试检查桶是否存在或创建它。
      
      如果了解桶已经存在，可以尝试将rclone执行的事务数量最小化时，这可能很有用。
      
      如果使用的用户没有创建桶的权限，则也可能需要此选项。在v1.52.0之前，由于错误，这将无声地通过。
      

   --no-head
      如果设置，则不会对上传的对象进行HEAD请求以检查完整性。
      
      如果最终上载成功后，rclone会返回一个200 OK的消息，那么它将假设它已经正确上传。
      
      特别需要注意的是:
      
      - 元数据，包括修改时间，存储类和内容类型与上传的相同
      - 大小与上传的相同
      
      对于单个部分PUT的响应，它会读取以下项：
      
      - MD5SUM
      - 上载日期
      
      对于多部分上传，将不会读取这些项。
      
      如果上传长度未知的源对象，则rclone **将**执行HEAD请求。
      
      设置此标志将增加未检测到的上传失败的几率，特别是大小不正确，因此不建议在正常操作中使用。实际上，即使在设置此标志的情况下，发生未检测到的上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。
      
      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制多长时间内未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端和HTTP/2存在未解决的问题。 S3后端默认启用HTTP/2，但可以在此处禁用。问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终结点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据具有更低的出口费用。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该设置为true、false或留空以使用提供者的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传
      
      如果设置为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      rclone < 1.59的版本使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况或测试中，否则不应该使用。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定的时间显示文件版本。
      
      参数应为日期“2006-01-02”、日期时间“2006-01-02 15:04:05”或距离现在的持续时间，例如“100d”或“1h”。
      
      请注意，当使用此选项时，不允许执行任何文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"将对象上传到S3。通常，rclone会将这些文件作为压缩对象下载。
      
      如果设置了该标志，那么rclone将按数据接收时的"Content-Encoding: gzip"对这些文件进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能gzip对象，请设置此标志。
      
      通常情况下，提供者在下载对象时不会更改对象。如果即使未使用“Content-Encoding: gzip”上传对象，某个提供商也可能对对象进行gzip压缩（例如Cloudflare）。
      
      如果设置了此标志，并且rclone下载了带有设置了Content-Encoding: gzip和块传输编码的对象，则rclone将实时对对象进行解压缩。
      
      如果设置为未设置（默认值），则rclone将根据提供者设置选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      不要设置和读取系统元数据


选项：
   --access-key-id value      AWS Access Key ID（访问密钥ID）。
   --acl value                创建桶和存储或复制对象时使用的默认ACL。 [$ACL]
   --endpoint value           Scaleway对象存储的终结点。 [$ENDPOINT]
   --env-auth                 从运行时获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --region value             要连接的区域。 [$REGION]
   --secret-access-key value  AWS Secret Access Key (密码)。 [$SECRET_ACCESS_KEY]
   --storage-class value      存储新对象时要使用的存储类。 [$STORAGE_CLASS]

   高级设置

   --bucket-acl value               创建桶时使用的默认ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的复制切割点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终结点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果设置为true，则使用路径样式访问；如果设置为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中最多存在的分块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不会尝试检查桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             不要设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的文件切割点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单个部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果设置为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]
   
   通用设置
   
   --name value  存储名称（默认值: 自动生成）
   --path value  存储路径
```
{% endcode %}