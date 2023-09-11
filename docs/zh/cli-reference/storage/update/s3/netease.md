# 网易云对象存储（NOS）

{% code fullWidth="true" %}
```
命令名称：
   singularity storage update s3 netease - 网易云对象存储（NOS）

用法：
   singularity storage update s3 netease [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅当access_key_id和secret_access_key为空时才适用。

      示例：
         | false | 下一步输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。

      如果要使用匿名访问或运行时凭证，留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。

      如果要使用匿名访问或运行时凭证，留空。

   --region
      连接的区域。

      如果使用的是S3克隆，并且您没有区域，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如，Jewel/v10之前的CEPH。

   --endpoint
      S3 API的终端节点。

      使用S3克隆时需要此选项。

   --location-constraint
      位置约束 - 必须与区域匹配。

      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。

      此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建存储桶。

      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，当对象由S3进行服务器端复制时，此ACL将应用，因为S3不复制源的ACL，而是写入一个新的ACL。

      如果acl是空字符串，则不添加X-Amz-Acl:头部，将使用默认（私有）ACL。

   --bucket-acl
      创建存储桶时使用的预定义ACL。

      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。

      如果"acl"和"bucket_acl"都是空字符串，则不添加X-Amz-Acl:头部，将使用默认（私有）ACL。

      示例：
         | private            | 拥有者具有FULL_CONTROL权限。
         |                    | 没有其他人可以访问（默认）。
         | public-read        | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 一般不推荐在存储桶上授予此权限。
         | authenticated-read | 拥有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      切换到分块上传的大小阈值。

      大于此大小的文件将以chunk_size的大小进行分块上传。
      最小值为0，最大值为5GiB。

   --chunk-size
      用于上传操作的块大小。

      当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"、使用"rclone mount"或google photos或google docs上传的文件）时，将使用此块大小进行分块上传。

      请注意，每个传输在内存中缓冲"--s3-upload-concurrency"个该大小的块。

      如果您正在通过高速链接传输大文件并且具有足够的内存，则增加此值将加快传输速度。

      Rclone会根据需要增加块大小，以避免超过10,000个块的限制。

      尺寸已知的文件使用配置的块大小进行上传。由于默认块大小为5MiB，并且最多可以有10,000个块，因此默认情况下，您可以流式传输的文件的最大大小为48GiB。如果您希望流式传输更大的文件，则需要增加块大小。

      增加块大小会降低使用"-P"标志显示的进度统计的准确性。Rclone在将块缓冲到AWS SDK时，将已发送的块视为已发送，实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和与真实情况越来越偏离的进度报告。

   --max-upload-parts
      多部分上传中的最大分块数。

      此选项定义进行多部分上传时要使用的最大分块数。

      如果某个服务不支持AWS S3规范的10,000个块，则可以使用此选项。

      Rclone会根据需要增加块大小，以避免超过此分块数的限制。

   --copy-cutoff
      切换到分块复制的大小阈值。

      需要进行服务器端复制的大于此大小的文件将以此大小的块进行复制。

      最小值为0，最大值为5GiB。

   --disable-checksum
      不将MD5校验和存储到对象元数据中。

      通常，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件的上传可能导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth=true，rclone可以使用共享凭证文件。

      如果该变量为空，rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空，则默认为当前用户的主目录。

           Linux/OSX: "$HOME/.aws/credentials"
           Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth=true，rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。

      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。

      这是同时上传的相同文件的块数。

      如果您正在通过高速链接上传较少数量的大文件，并且这些上传没有充分利用您的带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认设置），则rclone将使用路径样式访问，如果为false，则rclone将使用虚拟路径样式。有关详细信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如AWS、Aliyun OSS、网易云、腾讯云）要求将其设置为false - rclone将基于提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认设置），则rclone将使用v4身份验证。如果设置了此选项，则rclone将使用v2身份验证。

      仅在v4签名不起作用，例如Jewel/v10之前的CEPH时使用此选项。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。

      此选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。

      大多数服务将响应列表截断为1000个对象，即使请求的数量更多。在AWS S3中，这是一个全局限制，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。

      当S3最初发布时，它只提供了ListObjects调用来枚举存储桶中的对象。

      但是，在2016年5月，引入了ListObjectsV2调用。这样做可以提高性能，并且应尽可能使用。

      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用的列表对象方法。如果猜错，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset。

      某些提供商支持URL编码列表，如果可用，则在使用文件名中包含控制字符时，这种方式更可靠。如果设置为unset（默认设置），则rclone将根据提供商设置选择应用什么。

   --no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建。

      如果要尽量减少rclone执行的事务数，可以使用此选项，前提是您知道存储桶已经存在。

      如果您使用的用户没有创建存储桶的权限，也可能需要使用此选项。在v1.52.0之前，该选项会由于错误而默默通过。

   --no-head
      如果设置，不会HEAD已上传的对象以检查完整性。

      如果尽量减少rclone执行的事务数，则可使用此选项。

      设置后，如果rclone在使用PUT上传对象后接收到200 OK消息，则会认为该对象已正确上传。

      特别是，它假设：
      
      - 元数据，包括modtime、存储类和内容类型与上传时相同。
      - 大小与上传时相同。
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传对象长度未知，则rclone将**会**执行HEAD请求。
      
      设置此标志会增加无法检测到的上传失败的机会，特别是大小不正确的机会，因此不建议在正常操作中使用它。事实上，即使使用此标志，检测到的上传失败的机会也非常小。

   --no-head-object
      如果设置，则在获取对象之前不进行HEAD请求。

   --encoding
      后端的编码方式。

      有关详细信息，请参阅概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将被刷新的时间间隔。

      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中移除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端使用http2的功能。

      当前s3（特别是minio）后端与HTTP/2存在一个未解决的问题。对于s3后端，默认启用了HTTP/2，但可以在此禁用。解决该问题后，此标志将被删除。

      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      通常将其设置为CloudFront CDN URL，因为AWS S3为通过CloudFront网络下载的数据提供更便宜的出口。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证。

      此选项应设置为true、false或留空以使用提供商的默认设置。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传。

      如果设置为false，则rclone将使用AWS SDK中的PutObject来上传对象。

      rclone的版本<1.59使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应该将其设置为true。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。

      参数应为日期、“2006-01-02”格式的日期时间或该时间以前的持续时间，例如“100d”或“1h”。

      请注意，在使用此选项时，不允许执行文件写入操作，因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将会解压缩gzip编码的对象。

      可以使用"Content-Encoding: gzip"将对象上传到S3。通常，rclone会将这些文件作为压缩对象下载。

      如果设置了此标志，则rclone将在收到这些带有"Content-Encoding: gzip"的文件时解压缩它们。这意味着rclone无法检查大小和哈希值，但文件内容将解压缩。

   --might-gzip
      如果后端可能会压缩对象，则设置此标志。

      通常情况下，提供程序在下载对象时不会更改对象。如果一个对象没有使用`Content-Encoding: gzip`上传，那么在下载时它也不会设置。

      但是，某些提供商可能会压缩对象，即使它们没有使用`Content-Encoding: gzip`上传（例如Cloudflare）。

      这种情况的症状会出现错误，如下所示：

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置此标志，并且rclone使用设置了Content-Encoding: gzip和分块传输编码的对象，则rclone将在传输过程中实时解压缩该对象。

      如果设置为unset（默认设置），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项：
   --access-key-id value        AWS访问密钥ID。 [$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预定义ACL。 [$ACL]
   --endpoint value             S3 API的终端节点。 [$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值：false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束 - 必须与区域匹配。 [$LOCATION_CONSTRAINT]
   --region value               连接的区域。 [$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传操作的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的大小阈值。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将会解压缩gzip编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和存储到对象元数据中。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端使用http2的功能。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1,2或0表示自动。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大分块数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将被刷新的时间间隔。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会压缩对象，则设置此标志。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，不尝试检查存储桶是否存在或创建。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，不会HEAD已上传的对象以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象之前不进行HEAD请求。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的大小阈值。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单个部分上传 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值：false) [$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值：false) [$VERSIONS]

```
{% endcode %}