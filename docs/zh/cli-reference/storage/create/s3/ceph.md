# Ceph 对象存储

{% code fullWidth="true" %}
```
命令:
   singularity storage create s3 ceph - Ceph 对象存储

使用方法:
   singularity storage create s3 ceph [command options] [arguments...]

描述:
   --env-auth
      使用运行时获取 AWS 凭证（环境变量或环境模拟）。
      
      仅当 access_key_id 和 secret_access_key 为空时才适用。

      示例:
         | false | 下一步输入 AWS 凭证。
         | true  | 从环境获取 AWS 凭证（环境变量或 IAM）。

   --access-key-id
      AWS 访问密钥 ID。
      
      需要匿名访问或运行时凭证时可以留空。

   --secret-access-key
      AWS 密钥访问密码。
      
      需要匿名访问或运行时凭证时可以留空。

   --region
      要连接的区域。
      
      使用 S3 克隆且没有区域时可以留空。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用 v4 签名和空白区域。
         | other-v2-signature | 仅在 v4 签名不可用时使用此选项。
         |                    | 例如，早期的 Jewel/v10 CEPH。

   --endpoint
      S3 API 的终端节点。
      
      使用 S3 克隆时必填。

   --location-constraint
      位置约束 - 必须与区域匹配。
      
      不确定时可以留空。仅在创建存储桶时使用。

   --acl
      在创建存储桶、存储或复制对象时使用的 canned ACL。
      
      此 ACL 用于创建对象，并且当未设置 bucket_acl 时，也用于创建存储桶。
      
      有关详细信息，请访问 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      请注意，当 S3 服务器复制对象时，它不会复制源中的 ACL，而是输入一个新的 ACL。
      
      如果 acl 是一个空字符串，则不会添加 X-Amz-Acl 头，将使用默认值（private）。

   --bucket-acl
      在创建存储桶时使用的 canned ACL。
      
      有关详细信息，请访问 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      请注意，仅在创建存储桶时应用此 ACL。如果未设置它，则还将使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 都是空字符串，则不会添加 X-Amz-Acl
      头，并且将使用默认值（private）。

      示例:
         | private            | 拥有者具有 FULL_CONTROL 权限。
         |                    | 其他用户没有访问权限（默认）。
         | public-read        | 拥有者具有 FULL_CONTROL 权限。
         |                    | AllUsers 组具有读权限。
         | public-read-write  | 拥有者具有 FULL_CONTROL 权限。
         |                    | AllUsers 组具有读和写权限。
         |                    | 不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者具有 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers 组具有读权限。

   --server-side-encryption
      存储对象时使用的服务器端加密算法。

      示例:
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用 SSE-C，则存储对象时使用的服务器端加密算法。

      示例:
         | <unset> | None
         | AES256  | AES256

   --sse-kms-key-id
      如果使用 KMS ID，则必须提供密钥的 ARN。

      示例:
         | <unset>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      如果使用 SSE-C，则可以提供用于加密/解密数据的加密密钥。
      
      也可以提供 --sse-customer-key-base64。

      示例:
         | <unset> | None

   --sse-customer-key-base64
      如果使用 SSE-C，则必须提供以 base64 格式编码的加密密钥以加密/解密数据。
      
      也可以提供 --sse-customer-key。

      示例:
         | <unset> | None

   --sse-customer-key-md5
      如果使用 SSE-C，则可以提供加密密钥的 MD5 校验和（可选）。
      
      如果为空，则会自动从提供的 sse_customer_key 计算。

      示例:
         | <unset> | None

   --upload-cutoff
      切换到分块上传的大小阈值。
      
      大于此大小的文件将以块大小为 chunk_size 进行上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      在上传大于 upload_cutoff 的文件或大小未知的文件（例如使用 "rclone rcat" 上传或使用 "rclone mount" 或 google photos 或 google docs 上传）时，将使用这个块大小进行分块上传。
      
      请注意，"--s3-upload-concurrency" 每个传输在内存中缓冲这个大小的块。
      
      如果您在高速链接上传输大文件并且有足够的内存，增加此值将加快传输速度。
      
      Rclone 将根据文件大小增加块大小，以保持不超过 10,000 个块的限制。
      
      未知大小的文件使用配置的 chunk_size 进行上传。由于默认的 chunk_size 是 5 MiB，并且最多有 10,000 个块，这意味着默认情况下您可以流式上传的文件的最大大小为 48 GiB。如果您希望流式上传更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计信息的精确性。当 Rclone 缓冲 AWS SDK 时，会将块视为已发送，而实际上它可能仍在上传。更大的块大小意味着更大的 AWS SDK 缓冲区，并且与实际情况可能有所偏离的进度报告。

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义在执行多部分上传时要使用的最大多部分块数。
      
      如果某个服务不支持 AWS S3 规范中的 10,000 个多部分块，则这可能很有用。
      
      在上传已知大小的大文件时，Rclone 将自动增加块大小以保持在块数限制内。

   --copy-cutoff
      切换到分块复制的大小阈值。
      
      大于此大小的需进行服务器端复制的文件将以此大小的块进行复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      在对象元数据中不存储 MD5 校验和。
      
      通常情况下，rclone 在上传文件之前会计算输入数据的 MD5 校验和，以便将其添加到对象元数据中。这在进行数据完整性检查时非常有用，但可能导致大型文件开始上传的时间较长。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。
      
      如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。此变量控制在该文件中要使用的配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default" 的值。

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的同一文件块数。
      
      如果您正在使用高速链接上传少量大文件，并且这些上传没有完全利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认值），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟路径样式。有关详细信息，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供者（例如 AWS、Aliyun OSS、Netease COS 或 Tencent COS）要求将此设置为 false - rclone 将根据提供者设置自动执行此操作。

   --v2-auth
      如果为 true，则使用 v2 认证。
      
      如果为 false（默认值），则 rclone 将使用 v4 认证。如果设置了它，则 rclone 将使用 v2 认证。
      
      仅在 v4 签名不可用时使用。

   --list-chunk
      列表块的大小（每个 ListObject S3 请求的响应列表大小）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务即使请求更多对象，也会截断响应列表为 1000 个对象。
      在 AWS S3 中，这是全局最大值且不可更改，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项进行增加。

   --list-version
      要使用的 ListObjects 的版本：1、2 或 0 以自动选择。
      
      S3 最初只提供了用于枚举存储桶中对象的 ListObjects 调用。
      
      但是，2016 年 5 月，引入了 ListObjectsV2 调用。这是一种性能更高的调用，如果可能的话应该使用它。
      
      如果设置为默认值 0，则 rclone 将根据设置的提供者猜测要调用的 list objects 方法。如果它猜测错误，则可以在这里手动设置。

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      某些提供者支持对列表进行 URL 编码，如果可用，则在文件名中使用控制字符时这是更可靠的。如果设置为 unset（默认值），则 rclone 将根据提供者设置来选择应用什么。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，则这可能很有用，以减少 rclone 的事务数。
      
      如果您使用的用户没有桶创建权限，则可能需要进行此操作。在 v1.52.0 之前，因为存在一个错误，这将悄无声息地通过。
      

   --no-head
      如果设置，则不 HEAD 已上传的对象以检查完整性。
      
      如果尝试最小化 rclone 的事务数量，则这可能很有用。
      
      设置此标志意味着如果 rclone 在使用 PUT 上传对象后收到 200 OK 消息，则将假设该对象已正确上传。
      
      具体而言，它假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的内容相同。
      - 大小与上传的内容相同。
      
      对于单部分 PUT，它会从响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传的源对象长度未知，则 rclone **将**执行 HEAD 请求。
      
      设置此标志会增加检测不到的上传失败的几率，特别是大小不正确的几率，因此不建议在正常操作中使用。实际上，即使使用此标志，检测不到的上传失败的几率也很小。
      

   --no-head-object
      如果设置，则在获取对象之前不执行 HEAD。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区的上传（例如分块）将使用内存池进行分配。
      此选项控制多久会从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。
      
      目前 s3（具体来说是 minio）后端存在一个无法解决的问题，与 HTTP/2 相关。HTTP/2 默认启用 s3 后端，但可以在此禁用。解决该问题后，此标志将被删除。
      
      参见: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      自定义下载的终端节点。
      通常将其设置为 CloudFront CDN URL，因为 AWS S3 通过 CloudFront 网络下载的数据提供更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用 ETag 进行验证
      
      它应该是 true、false 或留空以使用提供者的默认值。

   --use-presigned-request
      是否使用预签名请求或 PutObject 进行单部分上传
      
      如果此选项为 false，则 rclone 将使用 AWS SDK 中的 PutObject 上传对象。
      
      版本低于 1.59 的 rclone 使用预签名请求上传单部分对象，将此标志设置为 true 将重新启用该功能。除非异常情况或测试，否则不应该需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间点的文件版本。
      
      参数应为日期，“2006-01-02”，日期时间“2006-01-02 15:04:05”或表示那么久以前的持续时间，例如“100d”或“1h”。
      
      请注意，在使用此选项时不允许执行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩 gzip 编码的对象。
      
      可以使用 "Content-Encoding: gzip" 在 S3 中上传对象。通常情况下，rclone 会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则 rclone 会在收到带有 "Content-Encoding: gzip" 的对象时进行解压缩。这意味着 rclone 无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能会对对象进行 gzip 压缩，请设置此选项。
      
      通常情况下，提供者在下载对象时不会更改对象。如果一个对象在上传时没有使用 `Content-Encoding: gzip`，则在下载时也不会设置该选项。
      
      但是，一些提供者可能会对对象进行 gzip 压缩，即使它们没有使用 `Content-Encoding: gzip` 进行上传（例如 Cloudflare）。
      
      如果设置了此标志，并且 rclone 下载了具有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，则 rclone 将在流式传输时即时解压缩对象。
      
      如果将其设置为 unset（默认值），则 rclone 将根据提供者的设置选择要应用的选项，但您可以在此处覆盖 rclone 的选择。
      

   --no-system-metadata
      禁止对系统元数据的设置和读取


OPTIONS:
   --access-key-id value           AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                     在创建存储桶和存储或复制对象时使用的 canned ACL。[$ACL]
   --endpoint value                S3 API 的终端节点。[$ENDPOINT]
   --env-auth                      使用运行时获取 AWS 凭证（环境变量或环境模拟）。 (默认值：false) [$ENV_AUTH]
   --help, -h                      显示帮助信息
   --location-constraint value     位置约束 - 必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region value                  要连接的区域。[$REGION]
   --secret-access-key value       AWS 密钥访问密码。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  在存储对象时使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          如果使用 KMS ID，则必须提供密钥的 ARN。[$SSE_KMS_KEY_ID]

   高级选项

   --bucket-acl value               在创建存储桶时使用的 canned ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的大小阈值。 (默认值： "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩 gzip 编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               在对象元数据中不存储 MD5 校验和。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 HTTP/2 使用。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             自定义下载的终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值： "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个 ListObject S3 请求的响应列表大小）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset。 (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 的版本：1、2 或 0 以自动选择。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此选项如果后端可能会对对象进行 gzip 压缩。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不 HEAD 已上传的对象以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象之前不执行 HEAD。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止对系统元数据的设置和读取 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用 SSE-C，则存储对象时使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         如果使用 SSE-C，则可以提供用于加密/解密数据的加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用 SSE-C，则必须提供以 base64 格式编码的加密密钥以加密/解密数据。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用 SSE-C，则可以提供加密密钥的 MD5 校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       多部分上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的大小阈值。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用 ETag 进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 进行单部分上传 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 认证。 (默认值：false) [$V2_AUTH]
   --version-at value               显示指定时间点的文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值：false) [$VERSIONS]

   General

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}