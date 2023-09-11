# Dreamhost DreamObjects

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 dreamhost - Dreamhost DreamObjects

用法：
   singularity storage update s3 dreamhost [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。

      仅当access_key_id和secret_access_key为空时适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。

      为空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。

      为空以进行匿名访问或运行时凭证。

   --region
      连接的区域。

      如果您使用S3克隆并且没有区域，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空白区域。
         | other-v2-signature | 仅当v4签名无效时使用。例如.之前的 Jewel/v10 CEPH。

   --endpoint
      S3 API的终端节点。

      使用S3克隆时必填。

      示例：
         | objects-us-east-1.dream.io | Dream Objects终端节点

   --location-constraint
      位置约束-必须设置以匹配区域。

      如果不确定，请留空。仅用于创建存储桶。

   --acl
      创建存储桶和存储或复制对象时使用的对象对ACL。

      此ACL用于创建对象以及如果未设置bucket_acl，则用于创建存储桶。

      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，S3在服务器端复制对象时应用此ACL，而不是从源复制ACL写入新的ACL。

      如果acl是空字符串，则不会添加X-Amz-Acl：header，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的对象对ACL。

      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。

      如果“acl”和“bucket_acl”是空字符串，则不会添加X-Amz-Acl：header，并且将使用默认值（private）。

      示例：
         | private            | 访问所有者为FULL_CONTROL。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 访问所有者为FULL_CONTROL。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 访问所有者为FULL_CONTROL。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 访问所有者为FULL_CONTROL。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止值。

      大于此大小的文件将按chunk_size切块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时使用的分块大小。

      当上传大于upload_cutoff的文件或大小不明确的文件（例如使用“rclone rcat”或使用“rclone mount”或google photos或google docs上传的文件）时，将使用此分块大小进行分块上传。

      请注意，每个传输将缓冲此大小的“--s3-upload-concurrency”个块。

      如果您要通过高速链路传输大文件并且拥有足够的内存，那么增加此值将加快传输速度。

      Rclone会自动增加分块大小以便用于已知大小的大文件以保持在10000块限制以下。

      未知大小的文件使用配置的chunk_size进行上传。由于默认的分块大小是5 MiB，并且最多可以有10000个块，这意味着默认情况下可以流式传输的文件的最大大小为48 GiB。如果要流式传输更大的文件，则需要增加chunk_size。

      增加分块大小会降低使用“-P”标志显示的进度统计的准确性。当rclone将块缓冲到AWS SDK时，rclone将分块视为已发送，而实际上可能仍在上传。较大的块尺寸意味着更大的AWS SDK缓冲区，并且与真实情况可能更偏离的进度报告。

   --max-upload-parts
      多部分上传的最大部分数。

      该选项定义使用多部分上传时要使用的最大多部分块数。

      当上传已知大小的大文件以保持在这些块数限制以下时，rclone会自动增加分块大小。

   --copy-cutoff
      切换到分块复制的截止值。

      需要服务器端复制的大于此大小的文件将以此大小的块复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和存储在对象元数据中。

      通常，在上传之前，rclone会计算输入的MD5校验和，并将其添加到对象的元数据中。这对于数据完整性检查很有用，但对于大文件来说，会导致启动上传过程的时间很长。

   --shared-credentials-file
      共享凭据文件的路径。

      如果env_auth = true，则rclone可以使用共享凭证文件。

      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空，则默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。

      如果为空，则默认为环境变量“AWS_PROFILE”或“default”，如果该环境变量也未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。

      这是同时上传的相同文件的块数。

      如果您在高速链路上上传较少数量的大文件，并且这些上传未完全利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问，如果为false，则rclone将使用虚拟路径样式。有关详细信息，请参见[the AWS S3 docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果未设置（默认值），则rclone将使用v4身份验证。如果设置，则rclone将使用v2身份验证。

      仅当v4签名无效时使用，例如.在 Jewel/v10 CEPH之前。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。

      这个选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。

      大多数服务都会截断响应列表以返回1000个对象，即使请求了更多。

      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。

      在Ceph中，可以使用“rgw list buckets max chunk”选项来增加这一值。

   --list-version
      要使用的ListObjects的版本：1、2或0自动。

      当S3最初推出时，它只提供了一个“ListObjects”调用来枚举存储桶中的对象。

      但是，从2016年5月开始，引入了ListObjectsV2调用。这是更高性能的调用，如果可能的话应该使用该调用。

      如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪种列出对象的方法。如果猜测错误，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      一些提供者支持URL编码列表，在使用控制字符的文件名时，这更可靠。如果设置为unset（默认值），则rclone将根据提供者设置来决定要应用什么，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，不要检查桶是否存在或创建它。

      如果您知道存储桶已经存在，这可能很有用，以尽量减少rclone执行的事务数。

      如果使用的用户没有创建桶的权限，也可能需要此选项。在v1.52.0之前，由于一个错误，这将默默地通过了。

   --no-head
      如果设置，则不对上传的对象进行HEAD请求以检查完整性。

      这在试图最小化rclone的事务数量时会很有用。

      如果在PUT上传对象后接收到200 OK消息，则rclone会假设它已经成功上传。

      特别是它将假设：

      - 元数据，包括修改时间、存储类和内容类型与上传的一样
      - 大小与上传的一样

      它从单个部分PUT的响应中读取以下项目：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项目。

      如果上传未知长度的源对象，则rclone**将**执行HEAD请求。

      设置此标志会增加未检测到的上传失败的可能性，特别是大小不正确的失败，因此不推荐在正常操作中使用。实际上，即使设置此标志，检测不到上传失败的几率也非常小。

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求前执行GET请求。

   --encoding
      后端的编码。

      请参见概述中的[编码部分](/overview/#encoding)获取更多信息。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。

      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。

      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前s3（特别是minio）后端存在一个未解决的问题，与HTTP/2相关。对于S3后端，默认启用了HTTP/2，但可以在此禁用。在问题解决之后，此标志将被删除。

      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终结点。

      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据具有更低的出口费用。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证

      这应为true、false或留空以使用提供者的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject上传单部分。

      如果为false，则rclone将使用AWS SDK的PutObject上传对象。

      rclone的版本 < 1.59使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。除了特殊情况或测试之外，这不是必要的。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。

      参数应为一个日期，“2006-01-02”，日期时间 “2006-01-02 15:04:05”或与之前时间相隔的持续时间，例如“100d”或“1h”。

      请注意，在使用此功能时，不允许进行文件写入操作，因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。

      可以将对象上传到S3并设置“Content-Encoding：gzip”。通常情况下，rclone将以压缩对象的形式下载这些文件。

      如果设置了此标志，则在接收到具有“Content-Encoding：gzip”的文件时，rclone将解压缩这些文件。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能压缩对象，请设置此标志。

      通常，当下载对象时，提供者不会更改对象。如果对象未使用'Content-Encoding：gzip'上传，则在下载时也不会设置。

      但是，某些提供商可能会压缩对象，即使它们未使用'Content-Encoding：gzip'上传（例如Cloudflare）。

      这种情况的症状可能是收到错误，例如

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置了此标志，并且rclone使用Content-Encoding: gzip和分块传输编码下载对象，则rclone将动态地解压缩对象。

      如果设置为unset（默认值），则rclone将根据提供者设置来决定要应用什么，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      抑制系统元数据的设置和读取


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的对象对ACL。[$ACL]
   --endpoint value             S3 API的终端节点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。 （默认值：false）[$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束-必须设置以匹配区域。[$LOCATION_CONSTRAINT]
   --region value               连接的区域。[$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               创建存储桶时使用的对象对ACL。[$BUCKET_ACL]
   --chunk-size value               上传时使用的分块大小。 （默认值：“5Mi”）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 （默认值：“4.656Gi”）[$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和存储在对象元数据中。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终结点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 （默认值：“Slash,InvalidUtf8,Dot”）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值：“unset”）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0自动。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传的最大部分数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 （默认值：“1m0s”）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请设置此标志。 （默认值：“unset”）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，不要检查桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对上传的对象进行HEAD请求以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求前执行GET请求。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             抑制系统元数据的设置和读取 （默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 （默认值：“200Mi”）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证。（默认值：“unset”） [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject上传单部分。 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 （默认值：“off”）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 （默认值：false）[$VERSIONS]

```
{% endcode %}