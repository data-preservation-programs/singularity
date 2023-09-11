# IDrive e2

## 使用方法：
```
singularity storage create s3 idrive [command options] [arguments...]
```

## 描述：
- `--env-auth`：从运行时获取AWS凭证（环境变量或EC2/ECS元数据）。
  - 仅当access_key_id和secret_access_key为空时有效。
  - 示例：
    - `false`：在下一步中输入AWS凭证。
    - `true`：从环境变量（环境变量或IAM）获取AWS凭证。

- `--access-key-id`：AWS访问密钥ID。
  - 对于匿名访问或运行时凭证，保持为空。

- `--secret-access-key`：AWS秘密访问密钥（密码）。
  - 对于匿名访问或运行时凭证，保持为空。

- `--acl`：创建存储桶、存储或复制对象时使用的预设ACL。
  - 此ACL用于创建对象，如果bucket_acl未设置，则也用于创建存储桶。
  - 有关更多信息，请参阅https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
  - 需要说明的是，当S3进行服务器端复制对象时，将应用该ACL，因为S3不会复制源对象的ACL，而是写入一个新的。
  - 如果acl是空字符串，则不添加X-Amz-Acl: header，并且将使用默认值（私有）。

- `--bucket-acl`：创建存储桶时使用的预设ACL。
  - 有关更多信息，请参阅https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
  - 需要说明的是，仅在创建存储桶时才应用此ACL。如果没有设置，则使用"acl"。
  - 如果"acl"和"bucket_acl"是空字符串，则不添加X-Amz-Acl: header，并且将使用默认值（私有）。

- `--upload-cutoff`：切换到分块上传的阈值。
  - 任何大于此阈值的文件将按块大小(chunk_size)进行分块上传。
  - 最小值为0，最大值为5 GiB。

- `--chunk-size`：用于上传的块大小。
  - 当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或Google照片或Google文档上传的文件）时，将使用此块大小进行分块上传。
  - 请注意，每个传输会在内存中缓冲"--s3-upload-concurrency"个这样大小的块。
  - 如果您使用高速链接传输大文件，并且具有足够的内存，则增加此值将加快传输速度。
  - Rclone将在上传已知大小的大文件时自动增加块大小，以保持在10000个块的限制以下。
  - 未知大小的文件使用配置的chunk_size进行上传。由于默认chunk_size为5 MiB，并且最多可以有10,000个块，这意味着默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果要流式上传更大的文件，您需要增加chunk_size。
  - 增加块大小会降低使用"-P"标志显示的进度统计的准确性。当rclone将块缓冲到AWS SDK中时，rclone会将该块视为已发送，而实际上可能仍在上传中。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与真相更偏离。

- `--max-upload-parts`：分块上传中的最大部分数。
  - 此选项定义执行分块上传时要使用的最大分块数。
  - 如果服务不支持AWS S3规范的10000个块，则可以使用此选项。
  - 当上传已知大小的大文件时，Rclone将自动增加块大小以保持在分块数限制以内。

- `--copy-cutoff`：切换到分块复制的阈值。
  - 需要复制的大于此阈值的文件将按此大小的块进行复制。
  - 最小值为0，最大值为5 GiB。

- `--disable-checksum`：不将MD5校验和与对象元数据一起存储。
  - 通常，rclone将在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间的延迟才能开始上传。

- `--shared-credentials-file`：共享凭证文件的路径。
  - 如果env_auth = true，则rclone可以使用共享凭证文件。
  - 如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量值为空，则默认为当前用户的主目录。
  - Linux/OSX："$HOME/.aws/credentials"
  - Windows："%USERPROFILE%\.aws\credentials"

- `--profile`：在共享凭证文件中要使用的配置文件。
  - 如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
  - 如果为空，则默认为"AWS_PROFILE"环境变量或"default"（如果该环境变量也未设置）。

- `--session-token`：AWS会话令牌。

- `--upload-concurrency`：用于分块上传的并发数。
  - 这是同时上传同一文件的块数。
  - 如果您在高速链接上上传少量大型文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

- `--force-path-style`：如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
  - 如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟主机样式访问。有关更多信息，请参阅[Amazon S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
  - 某些服务提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求此设置为false- rclone将根据提供商设置自动执行此操作。

- `--v2-auth`：如果为true，则使用v2身份验证。
  - 如果此值为false（默认值），则rclone将使用v4身份验证。如果设置了该值，则rclone将使用v2身份验证。
  - 仅当v4签名无效时才使用此选项，例如在Jewel/v10 CEPH之前。

- `--list-chunk`：列表块的大小（每个ListObject S3请求的响应列表）。
  - 此选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。
  - 大多数服务即使请求超过1000个对象，也会截断响应列表。
  - 在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
  - 在Ceph中，可以使用“rgw list buckets max chunk”选项来增加此值。 

- `--list-version`：要使用的ListObjects版本：1、2或0表示自动。
  - 当S3最初推出时，它只提供了ListObjects调用来枚举存储桶中的对象。
  - 但是在2016年5月推出了ListObjectsV2调用。这是更高性能的方式，应尽可能使用。
  - 如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪个list objects方法。如果猜错了，可以在此处手动设置。

- `--list-url-encode`：是否对列表进行URL编码：true/false/unset。
  - 一些服务支持URL编码列表，如果可用，则在文件名中使用控制字符时，这更可靠。如果设置为unset（默认值），则rclone将根据提供者设置的内容选择何时应用，但您可以在此处覆盖rclone的选择。

- `--no-check-bucket`：如果设置，则无需检查存储桶是否存在或创建。
  - 如果您知道存储桶已经存在，并且希望最小化rclone执行的事务数，则这可能很有用。
  - 如果您使用的用户没有创建存储桶的权限，则也可能需要此选项。在v1.52.0之前的版本由于错误而导致此选项会被静默通过。

- `--no-head`：如果设置，则不要在获取对象时进行HEAD检查以验证完整性。
  - 如果要尽量减少rclone的事务数量，则这可能很有用。
  - 设置后，如果rclone在PUT对象后收到200 OK消息，则会假设该对象已上传正确。
  - 特别是它将假设：
    - 元数据，包括修改时间，存储类别和内容类型与上传的内容相同。
    - 大小与上传的内容相同。
  - 它从单个部分PUT的响应中读取以下项目：
    - MD5SUM
    - 上传日期
  - 对于分块上传，不会读取这些项。
  - 如果上传未知长度的源对象，则rclone **将**执行HEAD请求。
  - 设置此标志会增加未检测到的上传失败的机会，尤其是尺寸不正确，因此不建议在正常操作中使用。实际上，即使使用此标志，未检测到的上传失败的机会也非常小。

- `--no-head-object`：如果设置，获取对象时不执行HEAD操作。

- `--encoding`：后端的编码。
  - 有关更多信息，请参阅概述中的[编码部分](/overview/#encoding)。

- `--memory-pool-flush-time`：内部内存缓冲池的刷新频率。
  - 需要使用额外缓冲区（例如multipart上传）的上传将使用内存池进行分配。
  - 该选项控制多久未使用的缓冲区将从池中删除。

- `--memory-pool-use-mmap`：是否在内部内存池中使用mmap缓冲区。

- `--disable-http2`：禁用S3后端的HTTP/2使用。
  - 目前，s3（特别是minio）后端存在一个未解决的问题，与HTTP/2有关。
  - s3后端默认启用HTTP/2，但可以在此禁用。
  - 解决此问题后，将删除此标志。
  - 请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631。

- `--download-url`：下载的自定义端点。
  - 这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据提供了更便宜的出站流量。

- `--use-multipart-etag`：是否在分块上传中使用ETag进行验证。
  - 这应该设置为true、false或保持未设置以使用提供商的默认值。

- `--use-presigned-request`：是否使用预签名请求或PutObject进行单个部分上传。
  - 如果此值为false，则rclone将使用AWS SDK中的PutObject上传对象。
  - rclone的版本< 1.59使用预签名请求上传单个部分对象，并将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应该这样做。

- `--versions`：在目录列表中包含旧版本。

- `--version-at`：显示指定时间的文件版本。
  - 参数应为一个日期，例如"2006-01-02"，日期时间"2006-01-02 15:04:05"或一段时间，例如"100d"或"1h"。
  - 请注意，在使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
  - 有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

- `--decompress`：如果设置，将解压缩gzip编码的对象。
  - 可以使用"Content-Encoding: gzip"在S3上上传对象。通常，rclone会将这些文件作为压缩对象下载。
  - 如果设置了此标志，则rclone将在接收到"Content-Encoding: gzip"的文件时对其进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

- `--might-gzip`：如果后端可能对对象进行gzip压缩，请设置此标志。
  - 通常，提供商不会在下载对象时更改对象。如果未使用`Content-Encoding: gzip`上传对象，则在下载时不会设置它。
  - 但是，某些提供商可能会对对象进行gzip压缩，即使它们在上传时未使用`Content-Encoding: gzip`（例如Cloudflare）。
  - 如果设置了此标志，并且rclone下载具有设置`Content-Encoding: gzip`和分块传输编码的对象，则rclone将动态地对对象进行解压缩。
  - 如果设置为unset（默认值），则rclone将根据提供者设置的内容选择何时应用，但您可以在此处覆盖rclone的选择。

- `--no-system-metadata`：禁止设置和读取系统元数据

## 选项：
- `--access-key-id value`：AWS访问密钥ID。[$ACCESS_KEY_ID]
- `--acl value`：创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
- `--env-auth`：从运行时获取AWS凭证（环境变量或EC2/ECS元数据）。默认值：false。[$ENV_AUTH]
- `--help, -h`：查看帮助
- `--secret-access-key value`：AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

- 高级选项：

- `--bucket-acl value`：创建存储桶时使用的预设ACL。[$BUCKET_ACL]
- `--chunk-size value`：用于上传的块大小。默认值："5Mi"。[$CHUNK_SIZE]
- `--copy-cutoff value`：切换到分块复制的阈值。默认值："4.656Gi"。[$COPY_CUTOFF]
- `--decompress`：如果设置这个标志，它将解压缩gzip编码的对象。默认值：false。[$DECOMPRESS]
- `--disable-checksum`：不将MD5校验和与对象元数据一起存储。默认值：false。[$DISABLE_CHECKSUM]
- `--disable-http2`：禁用S3后端的HTTP/2使用。默认值：false。[$DISABLE_HTTP2]
- `--download-url value`：下载的自定义端点。[$DOWNLOAD_URL]
- `--encoding value`：后端的编码。默认值："Slash,InvalidUtf8,Dot"。[$ENCODING]
- `--force-path-style`：如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。默认值：true。[$FORCE_PATH_STYLE]
- `--list-chunk value`：列表块的大小。默认值：1000。[$LIST_CHUNK]
- `--list-url-encode value`：是否对列表进行URL编码：true/false/unset。默认值："unset"。[$LIST_URL_ENCODE]
- `--list-version value`：要使用的ListObjects版本：1、2或0表示自动。默认值：0。[$LIST_VERSION]
- `--max-upload-parts value`：分块上传中的最大部分数。默认值：10000。[$MAX_UPLOAD_PARTS]
- `--memory-pool-flush-time value`：内部内存缓冲池的刷新频率。默认值："1m0s"。[$MEMORY_POOL_FLUSH_TIME]
- `--memory-pool-use-mmap`：是否在内部内存池中使用mmap缓冲区。默认值：false。[$MEMORY_POOL_USE_MMAP]
- `--might-gzip value`：设置此标志如果后端可能对对象进行gzip压缩。默认值："unset"。[$MIGHT_GZIP]
- `--no-check-bucket`：如果设置，则无需检查存储桶是否存在或创建。默认值：false。[$NO_CHECK_BUCKET]
- `--no-head`：如果设置，则不进行HEAD检查以验证完整性。默认值：false。[$NO_HEAD]
- `--no-head-object`：如果设置，则在获取对象时不进行HEAD检查。
- `--no-system-metadata`：禁止设置和读取系统元数据。默认值：false。[$NO_SYSTEM_METADATA]
- `--profile value`：在共享凭证文件中要使用的配置文件。[$PROFILE]
- `--session-token value`：AWS会话令牌。[$SESSION_TOKEN]
- `--shared-credentials-file value`：共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
- `--upload-concurrency value`：用于分块上传的并发数。默认值：4。[$UPLOAD_CONCURRENCY]
- `--upload-cutoff value`：切换到分块上传的阈值。默认值："200Mi"。[$UPLOAD_CUTOFF]
- `--use-multipart-etag value`：是否在分块上传中使用ETag进行验证。默认值："unset"。[$USE_MULTIPART_ETAG]
- `--use-presigned-request`：是否使用预签名请求或PutObject进行单个部分上传。默认值：false。[$USE_PRESIGNED_REQUEST]
- `--v2-auth`：如果为true，则使用v2身份验证。默认值：false。[$V2_AUTH]
- `--version-at value`：显示指定时间的文件版本。默认值："off"。[$VERSION_AT]
- `--versions`：在目录列表中包含旧版本。默认值：false。[$VERSIONS]

- 常规选项：

- `--name value`：存储的名称（默认值：自动生成）
- `--path value`：存储的路径