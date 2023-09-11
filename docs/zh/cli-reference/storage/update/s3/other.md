# 任何其他兼容S3的提供商

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 other - 任何其他兼容S3的提供商

用法：
   singularity storage update s3 other [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时中获取AWS凭据（从环境变量或EC2 / ECS元数据获取，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时适用。例子：
      
      | false  | 在下一步中输入AWS凭据。
      | true   | 从环境中获取AWS凭据（env变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭据。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭据。

   --region
      要连接的区域。
      
      如果您使用的是S3克隆，并且没有区域，则留空。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要。

   --location-constraint
      位置约束 - 必须设置以匹配区域。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，默认情况下用于创建存储桶。 有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，此ACL仅在服务器端复制对象时应用，因为S3不会复制源对象的ACL，而是写入新的ACL。
      
      如果acl为空字符串，则不添加X-Amz-Acl：头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，此ACL仅在创建存储桶时应用。 如果未设置此参数，则使用acl代替。
      
      如果"acl"和"bucket_acl"均为空字符串，则不添加X-Amz-Acl：头，并且将使用默认值（private）。

   --upload-cutoff
      切换到分块上传的切换点。
      
      大于此大小的任何文件将使用块大小进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或Google照片或Google文档上传的文件）时，将使用此块大小进行分块上传。
      
      请注意，每个传输的内存中都会缓冲这个大小的"--s3-upload-concurrency"块。
      
      如果您在高速链路上传输大型文件，并且您具有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件时，rclone会自动增加块大小，以保持在10000个块限制以下。
      
      未知大小的文件会使用配置的块大小进行上传。 由于默认的块大小为5 MiB，最多有10000个块，这意味着默认情况下流式传输上传的文件的最大大小为48 GiB。如果要流式传输上传更大的文件，则需要增加块大小。
      
      增加块大小会降低使用“-P”标志显示的进度统计信息的准确性。当rclone将块缓冲到AWS SDK时，rclone会将块视为已发送，而实际上它可能仍在上传。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况不符。

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义在执行多部分上传时要使用的最大多部分块数。
      
      如果某个服务不支持AWS S3规范的10000个块，则此选项可能很有用。
      
      当上传已知大小的大文件时，rclone会自动增加块大小，以保持在这个块数量限制以下。

   --copy-cutoff
      切换到分块复制的切换点。
      
      需要服务器端复制的大于此大小的任何文件将以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。 这对于数据完整性检查非常有用，但可能导致大文件开始上传之前出现长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE" env变量。 如果env值为空，它将默认为当前用户的主目录。
      
          Linux / OSX："$ HOME / .aws / credentials"
          Windows："% USERPROFILE% \ .aws \ credentials"

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。 此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果该环境变量也没有设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      同一文件的这么多个块将同时上传。
      
      如果您通过高速链接上传少量大文件，并且这些上传未充分利用带宽，则增加此数量可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问; 如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问; 如果为false，则rclone将使用虚拟路径样式。 有关详细信息，请参见 [AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为false- rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。 如果设置了此值，则rclone将使用v2身份验证。
      
      仅在v4签名不起作用时使用，例如老版本的Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表大小）。
      
      此选项也称为AWS S3规范中的 "MaxKeys"、"max-items"或 "page-size"。
      大多数服务都将响应列表截断为1000个对象，即使请求超过1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参见 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过"rgw list buckets max chunk"选项增大此数值。

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      当S3最初推出时，它仅提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月引入了ListObjectsV2调用。 这个性能更高，应尽可能使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪个列出对象方法。 如果猜测错误，则可能会在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，则在文件名中使用控制字符时，这种方法更可靠。 如果将此设置为unset（默认值），则rclone将根据提供者的设置选择要应用的内容，但可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果您知道存储桶已经存在，那么这可能会有助于最小化rclone的事务数量。
      
      如果您使用的用户没有存储桶创建权限，则也可能需要执行此操作。在v1.52.0之前，由于一个错误，此操作将会默默通过。

   --no-head
      如果设置，则不会在获取对象时进行HEAD以检查完整性。
      
      这对于最小化rclone所做的事务数量非常有用。
      
      设置此标志意味着如果rclone在使用PUT上传对象后收到200 OK消息，那么它将假定它已正确上传。
      
      特别情况下，它将假定：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传的一样
      - 大小与上传的一样
      
      它从单个部分PUT的响应中读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不读取这些内容。
      
      如果上传未知长度的源对象，那么rclone**将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确，因此不建议在正常操作中使用。 实际上，即使使用此标志，发生未检测到的上传失败的几率也非常小。

   --no-head-object
      如果设置，则在获取对象时不会先执行HEAD操作。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的时间间隔。
      
      需要额外缓冲区（例如多部分）的上传将使用内存池来进行分配。
      此选项控制未使用的缓冲区将被从池中移除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在未解决的问题和HTTP/2之间的问题。 S3后端默认启用HTTP/2，但可以在此处禁用。 解决此问题后，将删除此标志。

      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      此通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据的出站费用较低。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该是true，false或保持unset以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果此值为false，则rclone将使用AWS SDK中的PutObject上传对象。
      
      版本1.59之前的rclone使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。 除非特殊情况或测试需要，否则不应该使用此选项。

   --versions
      在目录列表中包括旧版本。

   --version-at
      将文件版本显示为指定时间点的版本。
      
      参数应为日期 "2006-01-02"、日期时间 "2006-01-02
      15:04:05" 或该时间以前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，在使用该选项时，不允许进行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅 [时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，这将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"在S3中将对象上传为压缩对象。
      
      如果设置了此标志，则rclone将在接收到这些使用"Content-Encoding: gzip"的文件时进行解压缩。 这意味着rclone无法检查大小和哈希值，但文件内容将解压缩。

   --might-gzip
      如果后端可能会使用gzip压缩对象，请设置此标志。
      
      通常情况下，提供者在下载时不会更改对象。如果一个对象在上传时没有使用`Content-Encoding：gzip`上传，那么在下载时也不会设置它。
      
      但是，一些提供商甚至在它们没有上传`Content-Encoding：gzip`的对象时也可能对对象进行gzip压缩（例如Cloudflare）。
      
      这种情况的症状可能是接收到像下面这样的错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，而且rclone使用了使用`Content-Encoding：gzip`和分块传输编码的对象，则rclone将会在传输中解压缩该对象。
      
      如果将其设置为unset（默认值），则rclone将根据提供者的设置选择要应用的内容，但可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。 [$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预定义ACL。 [$ACL]
   --endpoint value             S3 API的端点。 [$ENDPOINT]
   --env-auth                   从运行时中获取AWS凭据（从环境变量或EC2 / ECS元数据获取，如果没有环境变量）。 （默认值：false）[$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束 - 必须设置以匹配区域。 [$LOCATION_CONSTRAINT]
   --region value               要连接的区域。 [$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 （默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的切换点。 （默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置这个将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 （默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问; 如果为false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表大小）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。（默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。（默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池刷新的时间间隔。（默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会使用gzip压缩对象，请设置此标志。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会在获取对象时执行HEAD以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不会先执行HEAD操作。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 （默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的切换点。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证。 （默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传。 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value               将文件版本显示为指定时间点的版本。 （默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 （默认值：false）[$VERSIONS]

```
{% endcode %}