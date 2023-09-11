# 七牛对象存储（Kodo）

{% code fullWidth="true" %}
```
命令：
   singularity storage update s3 qiniu - 七牛对象存储（Kodo）

用法：
   singularity storage update s3 qiniu [命令选项] <name|id>

说明：
   --env-auth
      从运行环境中获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅当access_key_id和secret_access_key为空时才适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。

      如果要进行匿名访问或者使用运行时凭证，请将其留空。

   --secret-access-key
      AWS密钥（密码）。

      如果要进行匿名访问或者使用运行时凭证，请将其留空。

   --region
      要连接的区域。

      示例：
         | cn-east-1      | 默认端点，如果不确定可以选择此选项。
         |                | 中国东部区域1。
         |                | 需要区域约束为cn-east-1。
         | cn-east-2      | 中国东部区域2。
         |                | 需要区域约束为cn-east-2。
         | cn-north-1     | 中国北部区域1。
         |                | 需要区域约束为cn-north-1。
         | cn-south-1     | 中国南部区域1。
         |                | 需要区域约束为cn-south-1。
         | us-north-1     | 北美区域。
         |                | 需要区域约束为us-north-1。
         | ap-southeast-1 | 东南亚区域1。
         |                | 需要区域约束为ap-southeast-1。
         | ap-northeast-1 | 东北亚区域1。
         |                | 需要区域约束为ap-northeast-1。

   --endpoint
      七牛对象存储的端点。

      示例：
         | s3-cn-east-1.qiniucs.com      | 中国东部区域1的端点
         | s3-cn-east-2.qiniucs.com      | 中国东部区域2的端点
         | s3-cn-north-1.qiniucs.com     | 中国北部区域1的端点
         | s3-cn-south-1.qiniucs.com     | 中国南部区域1的端点
         | s3-us-north-1.qiniucs.com     | 北美区域的端点
         | s3-ap-southeast-1.qiniucs.com | 东南亚区域1的端点
         | s3-ap-northeast-1.qiniucs.com | 东北亚区域1的端点

   --location-constraint
      区域约束，必须设置与区域匹配。

      仅在创建存储桶时使用。

      示例：
         | cn-east-1      | 中国东部区域1
         | cn-east-2      | 中国东部区域2
         | cn-north-1     | 中国北部区域1
         | cn-south-1     | 中国南部区域1
         | us-north-1     | 北美区域1
         | ap-southeast-1 | 东南亚区域1
         | ap-northeast-1 | 东北亚区域1

   --acl
      创建存储桶、存储对象或复制对象时使用的预定义ACL。

      此ACL用于创建对象，如果未设置bucket_acl，则用于创建存储桶。

      获取更多信息请访问[Amazon S3 ACL文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      需要注意的是，在服务器端复制对象时会应用此ACL，
      因为S3不会复制源对象的ACL，而是写入一个新的ACL。

      如果acl为空字符串，则不添加X-Amz-Acl:头，并使用默认（private）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。

      获取更多信息请访问[Amazon S3 ACL文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      需要注意的是，仅在创建存储桶时应用此ACL。如果未设置它，
      将使用"acl"代替。

      如果"acl"和"bucket_acl"都为空字符串，则不添加X-Amz-Acl:
      头，并使用默认（private）。

      示例：
         | private            | 拥有者具备FULL_CONTROL权限。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 拥有者具备FULL_CONTROL权限。
         |                    | 所有用户组具备READ权限。
         | public-read-write  | 拥有者具备FULL_CONTROL权限。
         |                    | 所有用户组具备READ和WRITE权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者具备FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具备READ权限。

   --storage-class
      存储新对象时使用的存储类型。

      示例：
         | STANDARD     | 标准存储类型
         | LINE         | 低访问频率存储模式
         | GLACIER      | 归档存储模式
         | DEEP_ARCHIVE | 深度归档存储模式

   --upload-cutoff
      切换到分片上传的文件截止值。

      超过此大小的文件将按分片大小进行分片上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时使用的分片大小。

      上传超过upload_cutoff的大文件或没有大小信息的文件
      （例如来自"rclone rcat"或使用"rclone mount"、Google照片或Google文档上传的文件）
      将使用此分片大小进行分片上传。

      请注意，每个传输会在内存中缓冲"--s3-upload-concurrency"个此大小的分片。

      如果您正在高速链接上传大文件，且内存充足，则增加此值将加快传输速度。

      当上传已知大小的大文件以保持低于10,000个分片的限制时，
      rclone会自动增加分片大小。

      未知大小的文件将使用配置的chunk_size进行上传。
      由于默认的分片大小为5 MiB，最多可以有10,000个分片，
      因此默认情况下，您可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大的文件，则需要增加chunk_size。

      增加分片大小会降低使用“-P”标志显示的进度统计数据的准确性。
      Rclone在分片被AWS SDK缓冲时将其视为已发送，
      而实际上可能仍在上传。更大的分片大小意味着更大的AWS SDK缓冲区和进度报告与实际情况更偏离。

   --max-upload-parts
      多部分上传中的最大部分数。

      此选项定义进行多部分上传时要使用的最大多部分分片数。

      如果某个服务不支持10,000个分片的AWS S3规范，
      这会很有用。

      当上传已知大小的大文件以保持低于指定的部分数限制时，
      rclone会自动增加分片大小。

   --copy-cutoff
      切换到分块复制的文件截止值。

      超过此大小需要服务器端复制的文件将按此大小的分块进行复制。
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和与对象元数据一起存储。

      通常rclone会在上传之前计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性检查很有用，
      但对于大文件启动上传可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth设置为true，则rclone可以使用共享凭证文件。

      如果该变量为空，rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。
      如果环境变量的值为空，则默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth设置为true，则rclone可以使用共享凭证文件。
      此变量控制在该文件中使用哪个配置文件。

      如果为空，则默认为环境变量"AWS_PROFILE"。
      如果环境变量未设置，则默认为"default"。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。

      这是同时上传同一文件的分片数。

      如果您正在高速链接上传大量大文件，
      而这些上传未充分利用您的带宽，
      则增加此值可能有助于提高传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟路径样式。
      有关更多信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（如AWS、阿里云OSS、网易云COS或腾讯云COS）要求将此设置为false，
      rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2鉴权。

      如果此值为false（默认值），rclone将使用v4鉴权。
      如果设置了此值，则rclone将使用v2鉴权。

      仅在无法使用v4签名的情况下使用，例如旧版本的Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。

      该选项也被称为"MaxKeys"、"max-items"或"AWS S3规范的"page-size"。

      大多数服务限制了每个请求的列表响应列表为1000个对象，
      即使请求的数量多于此值。在AWS S3中，这是一个全局最大值，
      不能更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。

      在Ceph中，可以通过"rgw list buckets max chunk"选项来增加该值。

   --list-version
      要使用的ListObjects版本：1、2或0（自动）。

      当S3最初发布时，它仅提供了ListObjects调用以枚举存储桶中的对象。

      但是，在2016年5月，ListObjectsV2调用被引入。这个调用的性能要高得多，如果可能的话应该使用它。

      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个list对象方法。
      如果它猜错了，那么可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      某些提供商支持URL编码的列表，如果可用，则在使用控制字符的文件名时这更加可靠。
      如果设置为unset（默认值），则rclone会根据提供商设置选择应用什么，
      但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，不会尝试检查存储桶是否存在或创建它。

      如果您知道存储桶已经存在，那么这对于将rclone的事务数量最小化会很有用。

      如果使用的用户没有创建存储桶的权限，则可能需要执行此操作。
      在v1.52.0之前，由于一个bug，此操作将会静默通过。

   --no-head
      如果设置，不会对上传的对象进行HEAD请求来检查完整性。

      这对于将rclone的事务数量最小化会很有用。

      设置此标志意味着如果rclone在PUT上传对象后收到200 OK消息，
      则假设对象已上传正确。

      特别地，它会假设：

      - 元数据，包括修改时间、存储类和内容类型与上传时一致
      - 大小与上传时一致

      对于单部分PUT请求，它从响应中读取以下条目：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些条目。

      如果上传未知长度的源对象，则rclone将执行HEAD请求。

      设置此标志会增加未检测到的上传失败的几率，
      特别是错误的大小，因此不推荐在正常操作中使用。
      实际上，即使使用此标志，发生未检测到的上传失败的几率也非常小。

   --no-head-object
      如果设置，则在获取对象之前不会执行HEAD请求。

   --encoding
      后端的编码方式。

      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。

      使用需要额外缓冲区的上传（例如分块上传）将使用内存缓冲池进行分配。
      此选项控制多久从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前，s3（特别是minio）后端存在一个无法解决的问题和HTTP/2。
      s3后端默认启用HTTP/2，但可以在此禁用。
      当问题解决后，此标志将被移除。

      请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      自定义下载的端点。
      这通常设置为CloudFront CDN URL，
      因为通过CloudFront网络下载的数据，
      AWS S3提供更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证

      这应为true、false或设置为未设置以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传

      如果为false，则rclone将使用AWS SDK的PutObject上传对象。

      rclone版本<1.59会使用预签名请求上传单个分片对象，
      设置此标志为true将重新启用该功能。
      除非在特殊情况下或进行测试，否则不应该使用。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示文件版本为指定时间的版本。

      参数应为日期，"2006-01-02"，
      日期时间"2006-01-02 15:04:05"或距离现在那么久的时间，例如"100d"或"1h"。

      请注意，使用此选项将不允许进行文件写操作，
      因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则会解压缩gzip编码的对象。

      可以向S3上传具有"Content-Encoding: gzip"的对象。
      通常情况下，rclone会将这些文件以压缩对象的形式下载。

      如果设置了此标志，则rclone将在接收到具有"Content-Encoding: gzip"的文件时解压缩它们。
      这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会gzip对象，请设置此标志。

      通常，提供商不会在下载时更改对象。
      如果对象上没有设置"Content-Encoding: gzip"，则在下载时它也不会设置。

      但是，有些服务提供商甚至在未使用"Content-Encoding: gzip"上上传对象时也会gzip对象（例如Cloudflare）。

      出现这种情况的症状可能是接收到以下错误：

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置了该标志，并且rclone通过Content-Encoding: gzip和分块传输编码下载了一个对象，
      则rclone将实时解压缩该对象。

      如果将此设置为未设置（默认值），则rclone将根据提供商的设置选择应用什么，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶、存储对象或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value             七牛对象存储的端点。[$ENDPOINT]
   --env-auth                   从运行环境中获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 （默认值：false）[$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  区域约束，必须设置与区域匹配。[$LOCATION_CONSTRAINT]
   --region value               要连接的区域。[$REGION]
   --secret-access-key value    AWS密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value        存储新对象时使用的存储类型。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               上传时使用的分片大小。（默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的文件截止值。（默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置，则会解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不将MD5校验和与对象元数据一起存储。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             自定义下载的端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 （默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0（自动）。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 （默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此标志如果后端可能会gzip对象。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，不会尝试检查存储桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，不会对上传的对象进行HEAD请求来检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象之前不会执行HEAD请求。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 （默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分片上传的文件截止值。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 （默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2鉴权。 （默认值：false）[$V2_AUTH]
   --version-at value               显示文件版本为指定时间的版本。 （默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 （默认值：false）[$VERSIONS]

```
{% endcode %}