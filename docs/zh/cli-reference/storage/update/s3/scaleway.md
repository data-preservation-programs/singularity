# Scaleway 对象存储

{% code fullWidth="true" %}
```
命令名称:
   singularity storage update s3 scaleway - Scaleway 对象存储

使用方法:
   singularity storage update s3 scaleway [command options] <名称|ID>

描述:
   --env-auth
      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。
      
      仅当 access_key_id 和 secret_access_key 为空时适用。

      示例:
         | false | 在下一步中输入 AWS 凭证。
         | true  | 从环境中获取 AWS 凭证（环境变量或 IAM）。

   --access-key-id
      AWS 访问密钥 ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS 密钥访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --region
      连接的区域。

      示例:
         | nl-ams | 荷兰阿姆斯特丹
         | fr-par | 法国巴黎
         | pl-waw | 波兰华沙

   --endpoint
      Scaleway 对象存储的终端节点。

      示例:
         | s3.nl-ams.scw.cloud | 阿姆斯特丹终端节点
         | s3.fr-par.scw.cloud | 巴黎终端节点
         | s3.pl-waw.scw.cloud | 华沙终端节点

   --acl
      在创建存储桶、存储或复制对象时使用的绑定 ACL。
      
      此 ACL 用于创建对象，并且如果未设置 bucket_acl 时，也用于创建存储桶。
      
      有关更多信息，请访问 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      请注意，此 ACL 仅在 S3 执行服务器端复制对象时应用，因为 S3 不会从源复制 ACL，而是写入新的 ACL。
      
      如果 acl 是空字符串，则不会添加 X-Amz-Acl: 头，并且会使用默认值（private）。

   --bucket-acl
      在创建存储桶时使用的绑定 ACL。
      
      有关更多信息，请访问 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      请注意，此 ACL 仅在创建存储桶时应用。 如果未设置，则使用 "acl" 替代。
      
      如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加 X-Amz-Acl: 头，并且会使用默认值（private）。

      示例:
         | private            | 属主获取 FULL_CONTROL 权限。
         |                    | 没有其他用户访问权限（默认）。
         | public-read        | 属主获取 FULL_CONTROL 权限。
         |                    | AllUsers 组获得 READ 权限。
         | public-read-write  | 属主获取 FULL_CONTROL 权限。
         |                    | AllUsers 组获得 READ 和 WRITE 权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 属主获取 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers 组获得 READ 权限。

   --storage-class
      存储新对象时要使用的存储类。

      示例:
         | <unset>  | 默认值。
         | STANDARD | 标准类别，适用于按需内容（例如流媒体或 CDN）。
         |          | 这是默认的存储类别。
         | GLACIER  | 存档存储。
         |          | 价格更低，但必须首先恢复才能访问。

   --upload-cutoff
      切换到分块上传的文件截止大小。
      
      大于此大小的文件将以 chunk_size 的块上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于 upload_cutoff 的文件或大小未知的文件（例如通过 "rclone rcat" 上传或使用 "rclone mount" 或谷歌
      照片或谷歌文档上传的文件）时，将使用此块大小进行分块上传。
      
      请注意，每个传输缓冲区 "--s3-upload-concurrency" 个此大小的块。
      
      如果您通过高速链路传输大文件并且具有足够的内存，则增加此大小将加快传输速度。
      
      当上传已知大小的大文件时，rclone 将自动增加块大小，以使其保持在 10,000 个块的限制以内。
      
      大小未知的文件使用配置的块大小进行上传。由于默认的块大小为 5 MiB，最多有 10,000 个块，这意味着默认情况下可以按流式上传的文件的最大大小为 48 GiB。
      如果您希望按流式上传更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计信息的准确性。当 AWS SDK 缓冲区将 chunk 发送到服务器时，rclone 将其视为已发送，而实际上可能仍在上传。
      更大的块大小意味着更大的 AWS SDK 缓冲区和进度报告与实际情况更偏离。
      

   --max-upload-parts
      分块上传的最大分块数。
      
      此选项定义执行分块上传时要使用的最大分块数。
      
      如果某服务不支持 AWS S3 的 10,000 个分块规范，则此选项可能很有用。
      
      当上传已知大小的大文件时，rclone 将自动增加块大小，以使其保持在此分块数限制以内。
      

   --copy-cutoff
      切换到分块复制的文件截止大小。
      
      大于此大小的需要进行服务端复制的文件将按此大小进行复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常，rclone 在上传之前会计算输入文件的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件的开始上传可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。
      
      如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。该变量控制在该文件中使用的配置文件。
      
      如果留空，将默认使用环境变量 "AWS_PROFILE" 或 "default"，如果未设置该环境变量的情况下。
      

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您通过高速链路上传少量大文件并且这些上传未充分利用带宽，则增加此数字可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认值），则 rclone 将使用路径样式访问，如果为 false，则 rclone 将使用虚拟路径样式访问。有关详细信息，请参见 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如 AWS、Aliyun OSS、Netease COS 或 Tencent COS）要求将此设置为 false - rclone 将根据提供商设置自动执行此操作。

   --v2-auth
      如果为 true，则使用 v2 认证。
      
      如果为 false（默认值），则 rclone 将使用 v4 认证。如果设置了该值，则 rclone 将使用 v2 认证。
      
      仅在 v4 签名无效（例如在 Jewel/v10 CEPH 之前）时使用此选项。

   --list-chunk
      列表分块大小（每个 ListObject S3 请求的响应列表）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务将响应列表截断为 1000 个对象，即使请求了更多对象。
      在 AWS S3 中，这是一个全局最大值，无法更改，有关详情，请参见 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项进行增加。
      

   --list-version
      要使用的 ListObjects 版本：1、2 或 0（自动）。
      
      当 S3 最初发布时，它仅提供 ListObjects 调用以枚举存储桶中的对象。
      
      但在 2016 年 5 月，ListObjectsV2 调用被引入。这是性能更高的版本，如果可能应该使用它。
      
      如果设置为默认值 0，则 rclone 将根据提供商设置的并猜测要调用的 List Objects 方法。如果它的猜测错误，则可以在此处手动设置它。
      

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      某些提供商支持 URL 编码列表。当使用控制字符在文件名中时，这是更可靠的选择。如果设置为 unset（默认值），则 rclone 将根据提供商设置选择要应用的内容，但您可以在此处覆盖 rclone 的选择。
      

   --no-check-bucket
      如果设置，则不会尝试检查桶是否存在或创建桶。
      
      如果您知道存储桶已经存在，并且希望尽量减少 rclone 的交易数，这很有用。
      
      如果使用的用户没有创建存储桶的权限，则可能需要使用该选项。在 v1.52.0 之前，由于一个错误，这将会悄悄地传递。
      

   --no-head
      如果设置，则不会对上传的对象执行 HEAD 操作以检查完整性。
      
      如果尽量减少 rclone 的事务数量，这很有用。
      
      设置此标志意味着如果在使用 PUT 上传对象后收到 200 OK 消息，则 rclone 将假定它已正确上传。
      
      特别是它将假定：
      
      - 元数据（包括修改时间、存储类别和内容类型）与上传相同
      - 大小与上传的大小相同
      
      它从单个部分 PUT 的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传的源对象长度未知，则 rclone **将**执行 HEAD 请求。
      
      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确的机会，因此不建议在正常操作中使用此标志。实际上，即使在设置此标志的情况下，检测到未检测到的上传失败的机会也非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行 HEAD 操作。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概览中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新频率。
      
      需要额外缓冲区（如分块）的上传将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 HTTP/2 使用。
      
      目前的问题是 s3（特别是 minio）后端和 HTTP/2 之间存在未解决的问题。默认情况下，S3 后端启用 HTTP/2，但可以在此处禁用。问题解决后，将删除此标志。
      
      参见: [https://github.com/rclone/rclone/issues/4673](https://github.com/rclone/rclone/issues/4673)、[https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)
      

   --download-url
      下载的自定义终端节点。
      这通常设置为 CloudFront CDN URL，因为 AWS S3 通过 CloudFront 网络下载的数据提供更便宜的流出。

   --use-multipart-etag
      是否在分块上传中使用 ETag 进行验证
      
      此值应为 true、false 或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求或 PutObject 进行单个部分上传
      
      如果此值为 false，则 rclone 将使用 AWS SDK 中的 PutObject 上传对象。
      
      rclone 的版本号 < 1.59 会使用预签名请求上传单个部分对象，设置此标志为 true 将重新启用该功能。除非在特殊情况下或进行测试，否则不应该需要此功能。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示文件版本，如指定时间时的文件版本。
      
      参数应该是一个日期，"2006-01-02"，日期时间 "2006-01-02
      15:04:05" 或距那个时间以前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      登录 [time option docs](/docs/#time-option) 以查看有效格式。
      

   --decompress
      如果设置，将解压缩 gzip 编码的对象。
      
      可以使用 "Content-Encoding: gzip" 设置将对象上传到 S3。通常情况下，rclone 会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则 rclone 将在接收到带有 "Content-Encoding: gzip" 的文件时解压缩这些文件。这意味着 rclone
      无法检查大小和哈希值，但是文件内容将被解压缩。
      

   --might-gzip
      如果后端可能会压缩对象，请设置此标志。
      
      通常情况下，提供商不会在下载时更改对象。如果一个对象在上传时没有使用 `Content-Encoding: gzip` 进行上传，那么在下载时也不会设置它。
      
      但是，某些提供商可能会压缩对象，即使它们没有使用 `Content-Encoding: gzip` 进行上传（例如 Cloudflare）。
      
      这种情况的现象将会收到类似以下内容的错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志并且 rclone 下载带有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，则 rclone 将实时解压缩对象。
      
      如果设置为 unset（默认值），则 rclone 将根据提供商设置选择要应用的内容，但您可以在此处覆盖 rclone 的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id value      AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                在创建存储桶时使用的绑定 ACL。[$ACL]
   --endpoint value           Scaleway 对象存储的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --region value             连接的区域。[$REGION]
   --secret-access-key value  AWS 密钥访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      存储新对象时要使用的存储类别。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               在创建存储桶时使用的绑定 ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到 multipart copy 的文件截止大小。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩 gzip 编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 HTTP/2 使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表分块大小（每个 ListObject S3 请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1,2 or 0 for auto. (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传的最大分块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内存缓冲池的刷新频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会压缩对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不会尝试检查桶是否存在或创建桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对上传的对象执行 HEAD 操作以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行 HEAD 操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的文件截止大小。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用 ETag 进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 进行单个部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 认证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示文件版本，如指定时间时的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}