# SeaweedFS S3

{% code fullWidth="true" %}
```
名称：
   singularity storage create s3 seaweedfs - SeaweedFS S3

用法：
   singularity storage create s3 seaweedfs [命令选项] [参数...]

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有设置环境变量）。
      
      仅当`access_key_id`和`secret_access_key`为空时生效。

      示例：
         | false | 下一步输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      如果需要匿名访问或运行时凭证，请留空。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      如果需要匿名访问或运行时凭证，请留空。

   --region
      要连接的区域。
      
      如果使用的是S3克隆并且没有特定的区域，请留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如，旧版本的Jewel/v10 CEPH.

   --endpoint
      S3 API的终端节点。
      
      使用S3克隆时需要此选项。

      示例：
         | localhost:8333 | SeaweedFS S3本地终端节点

   --location-constraint
      地理位置约束 - 必须与区域匹配。
      
      如果不确定，请留空。仅用于创建存储桶。

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。
      
      此预定义ACL用于创建对象，如果未设置`bucket_acl`，也用于创建存储桶。
      
      有关详细信息，请访问[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      请注意，当使用S3服务器端复制对象时，会应用此ACL；
      S3不会复制源对象的ACL，而是写入新的ACL。
      
      如果ACL为空字符串，则不会添加X-Amz-Acl:标头，
      并且将使用默认值（私有）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关详细信息，请访问[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      请注意，仅在创建存储桶时应用此ACL。
      如果未设置，则使用`acl`选项。
      
      如果`acl`和`bucket_acl`都是空字符串，则不会添加X-Amz-Acl:标头，
      并且将使用默认值（私有）。

      示例：
         | private            | 所有者获得全部控制权。
         |                    | 没有其他用户具有访问权限（默认值）。
         | public-read        | 所有者获得全部控制权。
         |                    | AllUsers组具有读取访问权限。
         | public-read-write  | 所有者获得全部控制权。
         |                    | AllUsers组具有读写访问权限。
         |                    | 通常不建议在存储桶上设置此权限。
         | authenticated-read | 所有者获得全部控制权。
         |                    | AuthenticatedUsers组具有读取访问权限。

   --upload-cutoff
      切换到分块上传的文件大小截断值。
      
      大于该值的文件将会分块上传，块大小为`chunk_size`。
      最小值为0，最大值为5GB。

   --chunk-size
      用于上传的分块大小。
      
      上传大于`upload_cutoff`的文件或未知大小的文件（例如使用`rclone rcat`上传或使用`rclone mount`、谷歌照片或谷歌文档上传），
      将会使用此分块大小进行分块上传。
      
      请注意，每个传输的内存中会缓冲`--s3-upload-concurrency`个大小为此分块大小的文件块。
      
      如果您通过高速链接传输大文件并且有足够的内存，则增大此值将加快传输速度。
      
      当上传已知大小的大文件以保持小于10000个块的限制时，Rclone将自动增大分块大小。
      
      未知大小的文件将使用配置的分块大小上传。
      由于默认的分块大小为5 MiB，且最多可有10000块，因此默认情况下，您可以流式传输的文件的最大大小为48 GB。
      如果要流式传输更大的文件，则需要增加分块大小。
      
      增大分块大小会降低使用“-P”标志显示的进度统计的精确性。
      当Rclone将文件块缓冲到AWS SDK时，它将把块视为已发送，而实际上可能仍在上传。
      更大的块大小意味着更大的AWS SDK缓冲区，进度报告偏离实际情况越大。

   --max-upload-parts
      分块上传中的最大部分数。
      
      此选项定义执行分块上传时要使用的最大分块数。
      
      如果服务不支持AWS S3规范的10000个块，则此选项可能会很有用。
      
      当上传已知大小的大文件以保持小于此分块数限制时，Rclone将自动增大分块大小。

   --copy-cutoff
      切换到分块复制的文件大小截断值。
      
      需要进行服务器端复制的大于此值的文件将分块复制。
      
      最小值为0，最大值为5GB。

   --disable-checksum
      不要在对象元数据中存储MD5校验和。
      
      正常情况下，rclone会在上传之前计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性检查很有用，
      但对于大文件的开始上传可能会造成长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果`env_auth`为true，则rclone可以使用共享的凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果`env_auth`为true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”，如果该环境变量也未设置，则默认为“default”。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      同一文件的多个块并发上传。
      
      如果您使用高速链接上载较少数量的大文件，且这些上传未充分利用带宽，
      增大此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。
      
      如果为true（默认），rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。
      有关详细信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认），则rclone将使用v4身份验证。
      如果设置了该选项，则rclone将使用v2身份验证。
      
      仅在v4签名无效时使用，例如，旧版本的Jewel/v10 CEPH。

   --list-chunk
      列出清单时的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为“MaxKeys”、“max-items”或“page-size”，来自AWS S3规范。
      大多数服务在请求超过1000个对象时将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，详见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0自动。
      
      当最初发布S3时，它仅提供了ListObjects调用以枚举桶中的对象。
      
      但在2016年5月，引入了ListObjectsV2调用。这是更高性能的，应尽可能使用。
      
      如果设置为默认值0，则rclone将按照提供程序设置的方式猜测要调用的list objects方法。如果猜错了，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行url编码：true/false/unset
      
      某些提供商支持对列表进行URL编码，在使用文件名中的控制字符时这更可靠。
      如果设置为“unset”（默认值），则rclone将根据提供商设置选择应用的方法，但是，您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不会尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已存在，但希望尽量减少rclone的事务数量，则可能很有用。
      
      如果您使用的用户没有存储桶创建权限，则也可能需要，v1.52.0之前的版本由于bug而会默默透过，而不提供任何提示。

   --no-head
      如果设置，则不会HEAD已上传的对象以检查完整性。
      
      如果您尝试尽量减少rclone的事务数量，这可能很有用。
      
      设置后，意味着如果rclone在使用PUT上传对象后收到200 OK消息，则会假设它已正确上传。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传的内容相同。
      - 大小与上传的内容相同。
      
      它从单个块的PUT响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于分块上传，不会读取这些项目。
      
      如果上传未知长度的源对象，那么rclone将**会**执行HEAD请求。
      
      设置此标志会增加无法检测到的上传失败的机会，
      特别是错误的大小，因此不建议在正常操作中使用。实际上，即使设置了此标志，检测到的上传失败的机会也很小。

   --no-head-object
      如果设置，则在获取对象时不会先进行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新时间间隔。
      
      需要额外缓冲区的上传（例如分块上传）将使用内存池进行分配。
      此选项控制多久将清除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在一个未解决的问题与HTTP/2相关。
      s3后端默认启用HTTP/2，但可以在此禁用。
      在问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      通常将其设置为CloudFront CDN URL，因为通过CloudFront网络下载的数据可以获得更低的AWS S3出口费用。

   --use-multipart-etag
      分块上传中是否使用ETag进行验证。
      
      可以将其设置为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个分块上传。
      
      如果设置为false，则rclone将使用AWS SDK的PutObject来上传对象。
      
      1.59版本以下的rclone使用预签名请求来上传单个分块对象，
      将此标志设置为true将重新启用此功能。除非在特殊情况下或测试中，
      否则不需要设置此标志。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间点文件的版本。
      
      该参数应该是一个日期，例如“2006-01-02”，日期时间“2006-01-02 15:04:05”或距现在的时间差，例如“100d”或“1h”。
      
      请注意，使用此参数时，不允许进行文件写入操作，
      因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以将对象以“Content-Encoding: gzip”的形式上传到S3。通常情况下，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone会在收到“Content-Encoding: gzip”的文件时进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将解压缩。

   --might-gzip
      如果后端可能压缩了对象，则设置此选项。
      
      通常情况下，提供商在下载对象时不会更改对象。如果一个对象在上传时未设置“Content-Encoding: gzip”，那么在下载时也不会被设置。
      
      但是，一些提供商可能会压缩对象，即使它们没有使用“Content-Encoding: gzip”进行上传（例如Cloudflare）。
      
      如果设置了此标志，并且rclone使用了设置了“Content-Encoding: gzip”和分块传输编码的对象，
      那么rclone将会实时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但您可以在此覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id value       AWS Access Key ID（访问密钥ID）。[$ACCESS_KEY_ID]
   --acl value                 创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value            S3 API的终端节点。[$ENDPOINT]
   --env-auth                  从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有设置环境变量）。 （默认值：false）[$ENV_AUTH]
   --help, -h                  显示帮助
   --location-constraint value 地理位置约束 - 必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region value              要连接的区域。[$REGION]
   --secret-access-key value   AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value        创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value        用于上传的分块大小。 （默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value       切换到分块复制的文件大小截断值。 （默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress              如果设置，则将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum        不要在对象元数据中存储MD5校验和。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2           禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value      下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value          后端的编码方式。 （默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style        如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value        列出清单时的大小（每个ListObject S3请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value   是否对列表进行url编码：true/false/unset （默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value      要使用的ListObjects版本：1、2或0自动。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value  分块上传中的最大部分数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value 内部内存缓冲区池刷新时间间隔。 （默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap    是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value        如果后端可能压缩了对象，则设置此选项。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket         如果设置，则不会尝试检查存储桶是否存在或创建存储桶。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                 如果设置，则不会HEAD已上传的对象以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object          如果设置，则在获取对象时不会先进行HEAD请求。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata      禁止设置和读取系统元数据（默认值：false）[$NO_SYSTEM_METADATA]
   --profile value           共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value     AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value 共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value 分块上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value     切换到分块上传的文件大小截断值。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value 分块上传中是否使用ETag进行验证 （默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request   是否使用预签名请求或PutObject进行单个分块上传。 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                 如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value        显示指定时间点文件的版本。 （默认值："off"）[$VERSION_AT]
   --versions                在目录列表中包含旧版本。 （默认值：false）[$VERSIONS]

   通用

   --name value  存储的名称（默认：自动生成）
   --path value  存储的路径

```
{% endcode %}