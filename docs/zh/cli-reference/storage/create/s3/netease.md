# 网易对象存储（NOS）

{% code fullWidth="true" %}
```
命令名称:
   singularity storage create s3 netease - 网易对象存储（NOS）

用法:
   singularity storage create s3 netease [command options] [arguments...]

描述:
   --env-auth
      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。
      
      仅在 access_key_id 和 secret_access_key 为空时适用。

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
      
      如果您使用的是 S3 克隆，并且没有区域，留空。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用 v4 签名和一个空区域。
         | other-v2-signature | 仅在 v4 签名不适用时使用此选项。
         |                    | 例如，在 Jewel/v10 之前的 CEPH 中。

   --endpoint
      S3 API 的终节点。
      
      在使用 S3 克隆时需要提供。

   --location-constraint
      位置约束 - 必须与区域匹配。
      
      如果不确定，请留空。仅在创建桶时使用。

   --acl
      创建桶和存储或复制对象时使用的预定义 ACL（访问控制列表）。
      
      此 ACL 用于创建对象，并且如果未设置 bucket_acl，则用于创建桶。
      
      有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，S3 不会复制源的 ACL，而是写入新的 ACL。
      
      如果 acl 的值为空字符串，则不添加 X-Amz-Acl: 标头，并且将使用默认（私有）。

   --bucket-acl
      创建桶时使用的预定义 ACL。
      
      有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建桶时应用此 ACL。如果未设置，则使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 的值均为空字符串，则不添加 X-Amz-Acl: 标头，并且将使用默认（私有）。

      示例:
         | private            | 拥有者具有FULL_CONTROL权限。
         |                    | 没有其他人具有访问权限（默认）。
         | public-read        | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers 组具有读取权限。
         | public-read-write  | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers 组具有读取和写入权限。
         |                    | 通常不建议在桶上授予此权限。
         | authenticated-read | 拥有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers 组具有读取权限。

   --upload-cutoff
      切换到分块上传的截止值。
      
      大于此大小的文件将以块大小进行上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于 upload_cutoff 的文件或大小未知的文件（例如，来自 "rclone rcat" 或使用 "rclone mount" 或 Google 相册或 Google 文档上传的文件）时，
      将使用此分块大小进行分块上传。
      
      请注意，每个传输缓冲区的大小为 "--s3-upload-concurrency" 个该大小的块。
      
      如果您在高速链接上传输大文件，并且有足够的内存，那么增加此值将加快传输速度。
      
      当上传已知大小的大文件时，Rclone 将自动增加块大小，以保持在 10,000 个块的限制之下。
      
      未知大小的文件使用配置的块大小进行上传。由于默认块大小为 5 MiB，并且最多有 10,000 个块，这意味着默认情况下您可以流式传输上传的文件的最大大小为 48 GiB。
      如果要流式上传更大的文件，则需要增加块大小。
      
      增加块大小会降低进度统计的准确性（使用 "-P" 标志显示）。当 AWS SDK 缓冲块时，Rclone 在发送块时将其视为已发送，但事实上可能仍在上传。
      更大的块大小意味着更大的 AWS SDK 缓冲区和进度报告与真实情况偏离更大。

   --max-upload-parts
      分块上传的最大部分数。
      
      此选项定义进行分块上传时要使用的最大分块数。
      
      如果服务不支持 AWS S3 规范的 10,000 个块，则此选项可能很有用。
      
      当上传已知大小的大文件时，Rclone 将自动增加块大小以保持在此分块数的限制之下。

   --copy-cutoff
      切换到分块复制的截止值。
      
      需要服务器端复制的大于此大小的文件将分块复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常，rclone 将在上传之前计算输入的 MD5 校验和，以便可以在对象的元数据中添加它。这对于数据完整性检查非常有用，
      但对于大文件来说，可能会导致上传开始的停顿。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。
      
      如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果该环境变量也未设置）。
      

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      同一文件的分块数的并发上传。
      
      如果您在高速链接上上传少量大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认值），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟路径样式访问。有关更多信息，请参见 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如 AWS、Aliyun OSS、网易 COS 或腾讯 COS）要求将其设置为 false - rclone 会根据提供商设置自动执行此操作。

   --v2-auth
      如果为 true，则使用 v2 认证。
      
      如果为 false（默认值），则 rclone 将使用 v4 认证。如果设置了这个标志，rclone 将使用 v2 认证。
      
      仅在 v4 签名不适用时使用，例如在 Jewel/v10 之前的 CEPH。

   --list-chunk
      列出块的大小（每个 ListObject S3 请求的响应列表）。
      
      此选项也称为 AWS S3 规范的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务会截断响应列表以包含 1000 个对象，即使请求的数量更多。
      在 AWS S3 中，这是一个全局最大值，无法更改，参见 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项进行增加。
      

   --list-version
      要使用的 ListObjects 的版本：1、2 或 0（自动）。
      
      当 S3 最初发布时，仅提供了 ListObjects 调用，用于枚举存储桶中的对象。
      
      然而，在 2016 年 5 月，引入了 ListObjectsV2 调用。这是更高性能的方法，如果可能，应该使用它。
      
      如果设置为默认值 0，rclone 将根据设置的提供商猜测要调用的列表对象方法。 如果猜测错误，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      有些提供商支持对列表进行 URL 编码，并且在使用控制字符时可靠性更高。如果设置为未设置（默认值），则 rclone 将根据提供商设置选择要应用的内容，但您可以在此处覆盖 rclone 的选择。

   --no-check-bucket
      如果设置，不尝试检查桶是否存在或创建。
      
      如果您知道桶已经存在，则此功能可帮助尽量减少 rclone 执行的事务数。
      
      如果使用的用户没有桶创建权限，也可能需要使用此功能。在 v1.52.0 之前，由于一个错误，这个选项会无声地通过。
      

   --no-head
      如果设置，则不会 HEAD 已上传的对象以检查完整性。
      
      如果尽量减少 rclone 执行的事务数，此功能可帮助您。
      
      设置它意味着如果 rclone 收到 PUT 后的 200 OK 消息，则会默认假定它已正确上传。
      
      特别是，它会假定：
      
      - 元数据（包括修改时间、存储类和内容类型）与上传的对象一样
      - 大小与上传的大小相同
      
      它从单个部分 PUT 的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传未知长度的源对象，则 rclone **会**执行 HEAD 请求。
      
      设置此标志会增加无法检测到的上传失败的机率，特别是大小不正确，因此不建议在正常操作中使用。实际上，即使使用此标志，检测不到的上传失败的机会非常小。

   --no-head-object
      如果设置，则获取对象之前不会执行 HEAD 操作。

   --encoding
      后端的编码方式。
      
      有关详情，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区（例如，multipart）的上传将使用内存池进行分配。
      此选项控制多久会从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 用法。
      
      s3（特别是 minio）后端与 HTTP/2 存在无法解决的问题。s3 后端默认启用 HTTP/2，但可以在此禁用。问题解决后，将删除此标志。
      
      详见: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      此通常设置为 CloudFront CDN URL，因为 AWS S3 提供通过 CloudFront 网络下载的数据的费用较低。

   --use-multipart-etag
      是否在分块上传中使用 ETag 进行验证。
      
      这个值应该为 true、false 或未设置以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或 PutObject 进行单部分上传。
      
      如果为 false，则 rclone 将使用 AWS SDK 的 PutObject 来上传对象。
      
      版本小于 1.59 的 rclone 使用预签名请求来上传单部分对象，将此标志设置为 true 将重新启用该功能。除非在特殊情况下或用于测试，否则不应该使用此功能。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示文件版本为指定时间的版本。
      
      该参数应为日期，例如 "2006-01-02"，时间 "2006-01-02 15:04:05" 或早之前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，在使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将对gzip编码的对象进行解压缩。
      
      可以使用 "Content-Encoding: gzip" 将对象上传到 S3。通常，rclone 将以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则 rclone 在接收到带有 "Content-Encoding: gzip" 的文件时会对其进行解压缩。这意味着 rclone 无法检查大小和哈希，
      但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行 gzip 压缩，则设置此标志。
      
      通常，提供程序在下载对象时不会更改对象。如果没有将 "Content-Encoding: gzip" 加到上传的对象中，下载时也不会设置在下载中。
      
      但是，有些提供程序可能会对对象进行压缩，即使它们不是使用 "Content-Encoding: gzip" 上传的（例如 Cloudflare）。
      
      这种情况的症状将是收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置此标志，并且 rclone 下载了具有设置了 "Content-Encoding: gzip" 和分块传输编码的对象，则 rclone 将即时解压缩对象。
      
      如果将其设置为未设置（默认值），则 rclone 将根据提供商设置选择要应用的内容。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id value        AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                  创建桶和存储或复制对象时使用的预定义 ACL。[$ACL]
   --endpoint value             S3 API 的终节点。[$ENDPOINT]
   --env-auth                   从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束 - 必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region value               连接的区域。[$REGION]
   --secret-access-key value    AWS 密钥访问密钥（密码）。[$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               创建桶时使用的预定义 ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将对gzip编码的对象进行解压缩。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 http2 用法。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码。 (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 的版本：1、2 或 0（自动）。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行 gzip 压缩，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查桶是否存在或创建。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会 HEAD 已上传的对象以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则获取对象之前不会执行 HEAD 操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用 ETag 进行验证。 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 进行单部分上传。 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 认证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示文件版本为指定时间的版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   General

   --name value  存储的名称 （默认为自动生成）
   --path value  存储的路径

```
{% endcode %}