# Minio 对象存储

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 minio - Minio 对象存储

用法：
   singularity storage create s3 minio [command options] [arguments...]

描述：
   --env-auth
      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据）。
      
      仅当 access_key_id 和 secret_access_key 都为空时应用。

      示例：
         | false | 在下一步输入 AWS 凭证。
         | true  | 从环境（环境变量或 IAM）中获取 AWS 凭证。

   --access-key-id
      AWS 访问密钥 ID。
      
      对于匿名访问或运行时凭证，请留空。

   --secret-access-key
      AWS 秘密访问密钥（密码）。
      
      对于匿名访问或运行时凭证，请留空。

   --region
      要连接的区域。
      
      如果您使用的是 S3 克隆，并且没有区域，请留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用 v4 签名和空区域。
         | other-v2-signature | 仅在 v4 签名无法使用时使用此选项。
         |                    | 例如，Jewel/v10 之前的 CEPH。

   --endpoint
      S3 API 的接入点。
      
      使用 S3 克隆时需要提供此选项。

   --location-constraint
      区域限制 - 必须与区域相匹配。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      在创建存储桶、存储对象或复制对象时使用的权限控制。
      
      此权限控制用于创建对象，并且如果没有设置 bucket_acl，则也用于创建存储桶。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，此权限控制将在服务器端复制对象时应用，
      因为 S3 不会复制源的权限控制，而是写入一个新的权限控制。
      
      如果权限控制为空字符串，则不添加 X-Amz-Acl: 标头，并使用默认权限（private）。

   --bucket-acl
      在创建存储桶时使用的权限控制。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此权限控制。如果未设置，则使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 都为空字符串，则不添加 X-Amz-Acl: 标头，并使用默认权限（private）。

      示例：
         | private            | 所有者获得 FULL_CONTROL 权限。
         |                    | 没有其他用户有访问权限（默认）。
         | public-read        | 所有者获得 FULL_CONTROL 权限。
         |                    | 所有用户组获得 READ 权限。
         | public-read-write  | 所有者获得 FULL_CONTROL 权限。
         |                    | 所有用户组获得 READ 和 WRITE 权限。
         |                    | 不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者获得 FULL_CONTROL 权限。
         |                    | 所有认证用户组获得 READ 权限。

   --server-side-encryption
      在将对象存储到 S3 时使用的服务器端加密算法。

      示例：
         | <unset> | 禁用
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用 SSE-C，则在将对象存储到 S3 时使用的服务器端加密算法。

      示例：
         | <unset> | 禁用
         | AES256  | AES256

   --sse-kms-key-id
      如果使用 KMS ID，则必须提供密钥的 ARN。

      示例：
         | <unset>                 | 禁用
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      如果使用 SSE-C，则可以提供用于加密/解密数据的秘密加密密钥。
      
      或者，可以提供 --sse-customer-key-base64。

      示例：
         | <unset> | 禁用

   --sse-customer-key-base64
      如果使用 SSE-C，则必须以 Base64 格式提供用于加密/解密数据的秘密加密密钥。
      
      或者，可以提供 --sse-customer-key。

      示例：
         | <unset> | 禁用

   --sse-customer-key-md5
      如果使用 SSE-C，则可以提供秘密加密密钥的 MD5 校验和（可选）。
      
      如果留空，则会自动从提供的 sse_customer_key 计算此值。
      

      示例：
         | <unset> | 禁用

   --upload-cutoff
      切换到分块上传的最小文件大小。
      
      大于此大小的任何文件将以块的形式上传，每块大小为 chunk_size。
      允许的最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于 upload_cutoff 的文件或大小未知的文件（例如使用 "rclone rcat" 或使用 "rclone mount" 或上传的谷歌照片或谷歌文档）时，将使用此块大小进行分块上传。
      
      请注意，每次传输会在内存中缓冲这个大小的 "--s3-upload-concurrency"个块。
      
      如果您正在通过高速链接传输大文件并且具有足够的内存，则增加此大小将加快传输速度。
      
      当上传已知大小的大文件时，rclone 会自动增加分块大小，以保持在 10,000 个分块限制之下。
      
      未知大小的文件使用配置的 chunk_size 进行上传。
      由于默认的分块大小为 5 MiB，并且最多可以有 10,000 个分块，这意味着默认情况下您可以流式上传的文件的最大大小为 48 GiB。如果要流式上传更大的文件，则需要增加 chunk_size。
      
      增加块大小会减少使用 "-P" 标志时显示的进度统计数据的准确性。
      当 AWS SDK 缓冲了一个块时，rclone 将把其视为已发送，但实际上可能仍在上传中。
      更大的块大小意味着更大的 AWS SDK 缓冲区和与真相更偏离的进度报告。
      

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义进行多部分上传时要使用的最大多部分块数。
      
      如果服务不支持 AWS S3 中的 10,000 个多部分块的规范，则此选项可能有用。
      
      当上传已知大小的大文件时，rclone 会自动增加分块大小，以保持在此部分数限制之下。
      

   --copy-cutoff
      切换到多部分复制的最小文件大小。
      
      需要服务器端复制的大于此大小的文件将以此大小的块进行复制。
      
      允许的最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象的元数据一起存储。
      
      通常，rclone 会在上传之前计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这可以用于进行数据完整性检查，但是对于大文件来说，可能会导致长时间的上传延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。
      
      如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，则将使用当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果环境变量也未设置）。
      

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      同一文件的这么多块同时上传。
      
      如果您在高速链接上上传数量较小的大文件，并且这些上传未完全利用您的带宽，则增加此数字可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认），则 rclone 将使用路径样式访问，如果为 false，则 rclone 将使用虚拟主机样式访问。有关更多信息，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供者（例如 AWS、阿里云 OSS、网易云 COS 或腾讯云 COS）要求将其设置为 false - rclone 会根据提供者的设置自动执行此操作。

   --v2-auth
      如果为 true，则使用 v2 认证。
      
      如果为 false（默认），则 rclone 将使用 v4 认证。如果设置，则 rclone 将使用 v2 认证。
      
      仅当 v4 签名无法使用时才使用此选项，例如在 Jewel/v10 之前的 CEPH 中使用。

   --list-chunk
      列表块的大小（每个 ListObject S3 请求的响应列表大小）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务对响应列表进行截断，即使请求的大小超过了这个限制。
      在 AWS S3 中，这是一个全局最大值，无法更改，请参阅 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以通过 "rgw list buckets max chunk" 选项进行增加。
      

   --list-version
      要使用的 ListObjects 版本：1、2 或 0（自动）。
      
      当 S3 最初发布时，它只提供了 ListObjects 调用来枚举存储桶中的对象。
      
      但是，在 2016 年 5 月，引入了 ListObjectsV2 调用。这是更高性能的，应尽可能使用。
      
      如果设置为默认值 0，则 rclone 将根据设置的提供者猜测要调用哪个列表对象方法。如果它的猜测错误，则可以在此处手动设置它。
      

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      某些提供者支持对列表进行 URL 编码，若可用，则在文件名中使用控制字符更可靠。如果设置为 unset（默认值），则 rclone 会根据提供者的设置选择要应用的方式，但您可以在此处覆盖 rclone 的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果您知道存储桶已存在，尽量减少 rclone 进行的事务数，这可能有用。
      
      如果您使用的用户没有创建存储桶的权限，也可能需要使用此选项。在 v1.52.0 之前，由于错误，此选项会静默通过。
      

   --no-head
      如果设置，则不进行 HEAD 操作以检查已上传对象的完整性。
      
      如果尽量减少 rclone 执行的事务数量，这可能有用。
      
      如果 rclone 在 PUT 操作后收到 200 OK 消息，则会假定对象已正确上传。
      
      具体来说，它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传一致
      - 大小与上传一致
      
      它从单个部分 PUT 的响应中读取以下项目：
      
      - MD5SUM
      - 上传的日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传一个未知长度的源对象，则 rclone **将**执行 HEAD 请求。
      
      设置此标志会增加未检测到的上传故障的可能性，特别是大小不正确，因此不建议在正常操作中使用。实际上，即使使用此标志，未检测到的上传故障的几率也非常小。
      

   --no-head-object
      如果设置，则在 GET 对象时不执行 HEAD 操作。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新频率。
      
      需要额外缓冲区的上传（例如分块上传）将使用内存池进行分配。
      此选项控制多久删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端使用 http2。
      
      目前，s3 后端（特别是 Minio 后端）存在未解决的问题与 HTTP/2 相关。HTTP/2 在 S3 后端中默认启用，但可以在此禁用。问题解决后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631，
      

   --download-url
      下载的自定义端点。
      这通常设置为 CloudFront CDN URL，因为 AWS S3 提供通过 CloudFront 网络下载数据的更便宜的出口带宽。

   --use-multipart-etag
      是否在多部分上传中使用 ETag 进行验证
      
      这应该设置为 true、false 或留空以使用提供者的默认值。
      

   --use-presigned-request
      是否在单块上传时使用预签名请求或 PutObject。
      
      如果为 false，则 rclone 会使用 AWS SDK 的 PutObject 上传对象。
      
      rclone < 1.59 的版本使用预签名请求上传单部分对象，将此标志设置为 true 将重新启用该功能。除非特殊情况或测试，否则不应该使用此功能。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应为日期 "2006-01-02"、日期时间 "2006-01-02
      15:04:05" 或表示以前时间的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此选项时，不允许进行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅 [time 选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，将解压缩 gzip 编码的对象。
      
      可以将对象以 "Content-Encoding: gzip" 上传到 S3。通常情况下，rclone 会将这些文件下载为压缩对象。
      
      如果设置了此标志，则 rclone 会在接收到具有 "Content-Encoding: gzip" 的文件时进行解压缩。这意味着 rclone 无法检查大小和哈希，但是文件内容将被解压缩。
      

   --might-gzip
      如果存储后端可能对对象进行 gzip 压缩，请设置此标志。
      
      通常情况下，提供者在下载时不会更改对象。如果对象没有使用 `Content-Encoding: gzip` 上传，那么在下载时也不会设置该标头。
      
      但是，一些提供者甚至在对象没有使用 `Content-Encoding: gzip` 上传时也可能对其进行 gzip 压缩（例如 Cloudflare）。
      
      这种情况下的一种症状是接收到如下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了该标志，并且 rclone 下载了具有设置 `Content-Encoding: gzip` 和分块传输编码的对象，则 rclone 将在接收时即时解压缩对象。
      
      如果设置为 unset（默认值），则 rclone 将根据提供者的设置选择要应用的方式，但是您可以在此处覆盖 rclone 的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value           AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                     在创建存储桶和存储或复制对象时使用的权限控制。[$ACL]
   --endpoint value                S3 API 的接入点。[$ENDPOINT]
   --env-auth                      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据）。 (默认值：false) [$ENV_AUTH]
   --help, -h                      显示帮助
   --location-constraint value     区域限制 - 必须与区域相匹配。[$LOCATION_CONSTRAINT]
   --region value                  要连接的区域。[$REGION]
   --secret-access-key value       AWS 秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  在将对象存储到 S3 时使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          如果使用 KMS ID，则必须提供密钥的 ARN。[$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               在创建存储桶时使用的权限控制。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分复制的最小文件大小。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩 gzip 编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象的元数据一起存储。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端使用 http2。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个 ListObject S3 请求的响应列表大小）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0（自动）。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新频率。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果存储后端可能对对象进行 gzip 压缩，请设置此标志。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不进行 HEAD 操作以检查已上传对象的完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则在 GET 对象时不执行 HEAD 操作。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用 SSE-C，则在将对象存储到 S3 时使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         如果使用 SSE-C，则可以提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用 SSE-C，则必须以 Base64 格式提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用 SSE-C，则可以提供秘密加密密钥的 MD5 校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       多部分上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的最小文件大小。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用 ETag 进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否在单块上传时使用预签名请求或 PutObject。 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 认证。 (默认值：false) [$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值：false) [$VERSIONS]

   General

   --name value  存储的名称（自动生成）
   --path value  存储的路径

```
{% endcode %}