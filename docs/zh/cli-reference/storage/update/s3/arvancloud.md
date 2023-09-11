# Arvan Cloud Object Storage（AOS）

{% code fullWidth="true" %}
```
名称:
   singularity storage update s3 arvancloud - Arvan Cloud Object Storage（AOS)（AOS）

用法:
   singularity storage update s3 arvancloud [命令选项] <name|id>

说明:
   --env-auth
      从运行时获取 AWS 凭据（环境变量或 EC2/ECS 元数据（如果没有环境变量））。
      
      仅适用于 access_key_id 和 secret_access_key 为空的情况。

      示例:
         | false | 在下一步中输入 AWS 凭据。
         | true  | 从环境中获取 AWS 凭据（环境变量或 IAM）。

   --access-key-id
      AWS 访问密钥 ID。
      
      留空以进行匿名访问或运行时凭据。

   --secret-access-key
      AWS 秘密访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭据。

   --endpoint
      Arvan Cloud 对象存储（AOS）API 的终端节点。

      示例:
         | s3.ir-thr-at1.arvanstorage.com | 默认终端节点——如果不确定，可以选择此项。
         |                                | 伊朗特罗斯（Asiatech）
         | s3.ir-tbz-sh1.arvanstorage.com | 伊朗大不里士（Shahriar）

   --location-constraint
      位置约束-必须与终端节点匹配。
      
      仅在创建存储桶时使用。

      示例:
         | ir-thr-at1 | 伊朗特罗斯（Asiatech）
         | ir-tbz-sh1 | 伊朗大不里士（Shahriar）

   --acl
      创建存储桶、存储或复制对象时使用的预设 ACL。
      
      该 ACL 用于创建对象，并且如果未设置 bucket_acl，则用于创建存储桶。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，在服务器端复制对象时，S3 不会复制源的 ACL，而是写入一份新的 ACL。
      
      如果 ACL 是空字符串，则不会添加 X-Amz-Acl: 标头，并且将使用默认的（私有）。

   --bucket-acl
      创建存储桶时使用的预设 ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，仅在创建存储桶时应用该 ACL。如果未设置该选项，则使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加 X-Amz-Acl: 标头，并将使用默认的（私有）。

      示例:
         | private            | 所有者拥有完全控制权限。
         |                    | 没有其他人具有访问权限（默认）。
         | public-read        | 所有者拥有完全控制权限。
         |                    | AllUsers 组具有读取权限。
         | public-read-write  | 所有者拥有完全控制权限。
         |                    | AllUsers 组具有读取和写入权限。
         |                    | 通常不建议在存储桶上设置此权限。
         | authenticated-read | 所有者拥有完全控制权限。
         |                    | AuthenticatedUsers 组具有读取权限。

   --storage-class
      在 ArvanCloud 中存储新对象时要使用的存储类。

      示例:
         | STANDARD | 标准存储类

   --upload-cutoff
      切换到分段上传的大小截止值。
      
      任何大于此大小的文件将被分成 chunk_size 的块进行上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于 upload_cutoff 的文件或大小未知的文件（例如 "rclone rcat" 或使用 "rclone mount" 或 Google 照片或 Google 文档上传的文件）时，它们将使用此块大小进行分段上传。
      
      请注意，对于每次传输，"--s3-upload-concurrency" 个此大小的块将在内存中缓冲。
      
      如果您正在通过高速链接传输大文件，并且具有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件以避免超过 10,000 个分段的限制时，rclone 将自动增加块大小。
      
      不知道大小的文件会使用配置的 chunk_size 进行上传。由于默认的块大小为 5 MiB，并且最多可以有 10,000 个分段，所以默认情况下可以流式上传的文件的最大大小为 48 GiB。如果要流式上传更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。当 AWS SDK 将缓冲的块发送到服务器时，rclone 会将块视为已发送，但实际上可能仍在上传。较大的块大小意味着较大的 AWS SDK 缓冲区和与真实情况更偏离的进度报告。

   --max-upload-parts
      多部分上传的最大部分数。
      
      此选项定义在进行多部分上传时要使用的最大分段块数。
      
      如果某个服务不支持 AWS S3 的 10,000 个分段规范，这就会有所帮助。
      
      当上传已知大小的大文件以避免超过此数量的分段限制时，rclone 将自动增加块大小。

   --copy-cutoff
      切换到分段复制的大小截止值。
      
      任何大于此大小并需要服务器端复制的文件将被分成此大小的块进行复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常，rclone 会在上传之前计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件来说，可能会导致开始上传的时间过长。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。
      
      如果此变量为空，则 rclone 将搜索 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。此变量控制在该文件中使用的配置文件。
      
      如果留空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果未设置该环境变量）。

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的分段块数。
      
      如果您正在通过高速链接上传大量大文件，并且这些上传没有充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认值），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟路径样式访问。有关更多信息，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如 AWS、Aliyun OSS、Netease COS 或 Tencent COS）要求将此设置为 false-rclone 将根据提供程序设置自动执行此操作。

   --v2-auth
      如果为 true，则使用 v2 验证。
      
      如果为 false（默认值），则 rclone 将使用 v4 验证。如果设置了该值，则 rclone 将使用 v2 验证。
      
      仅当 v4 签名无法工作（例如 Jewel/v10 CEPH 之前）时才使用此选项。

   --list-chunk
      列出的块大小（每个 ListObject S3 请求的响应列表）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。大多数服务会将响应列表截断为 1000 个对象，即使请求了更多。
      在 AWS S3 中，这是一个全局最大值，无法更改，请参阅 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项来增加此值。

   --list-version
      要使用的 ListObjects 版本：1、2 或 0（自动）。
      
      当 S3 最初发布时，它只提供了 ListObjects 调用来列举存储桶中的对象。
      
      然后，在 2016 年 5 月引入了 ListObjectsV2 调用。它的性能要高得多，如果可能的话，应该使用该调用。
      
      如果设置为默认值 0，则 rclone 会根据设置的提供程序猜测要调用的 ListObjects 方法。如果猜错，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset。
      
      某些提供商支持 URL 编码列表。如果可用，则使用控制字符时，可靠性更高。如果设置为 unset（默认值），则 rclone 将根据提供商设置选择要应用的方法，但您可以在此处覆盖 rclone 的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，则此选项可以用于尽量减少 rclone 的事务数量。
      
      如果使用的用户没有存储桶创建权限，则可能需要使用此选项。在 v1.52.0 之前，由于一个错误，这将被静默传递。

   --no-head
      如果设置，则不会对已上传的对象进行 HEAD 请求以检查完整性。
      
      如果尽量减少 rclone 的事务数量，则可以使用此选项。
      
      设置此选项意味着，如果 rclone 在执行 PUT 操作后收到 200 OK 消息，则会认为对象已成功上传。
      
      特别是它将假设：
      
      - 元数据（包括修改时间、存储类别和内容类型）与上传的内容相同
      - 大小与上传的内容相同
      
      它从单个部分 PUT 的响应中读取以下项目：
      
      - MD5SUM
      - 已上传日期
      
      对于分段上传，不会读取这些项目。
      
      如果上传的源对象的大小未知，则 rclone **将**执行 HEAD 请求。
      
      设置此标志会增加未检测到的上传失败的机会，尤其是大小不正确的情况，因此不建议在正常操作中使用该标志。实际上，即使在启用此标志的情况下，未检测到的上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行 HEAD 请求。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅概述中的 [编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区的上传（例如多部分）将使用内存池进行分配。
      此选项控制将未使用的缓冲区从内存池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。
      
      目前，s3（特别是 minio）后端存在一个未解决的问题，与 HTTP/2 相关。默认情况下，s3 后端启用了 HTTP/2，但可以在此处禁用。解决这个问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631。

   --download-url
      下载的自定义终端节点。
      这通常设置为 CloudFront CDN URL，因为 AWS S3 提供了通过 CloudFront 网络下载的更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用 ETag 进行验证
      
      这应该设置为 true、false 或留在 unset 状态以使用提供程序的默认设置。

   --use-presigned-request
      是否使用预签名请求或 PutObject 进行单个部分上传
      
      如果此标志为 false，则 rclone 将使用 AWS SDK 的 PutObject 上载对象。
      
      rclone 的版本 < 1.59 使用预签名请求来上传单个部分的对象，将此标志设置为 true 将重新启用该功能。除了特殊情况或用于测试外，不应该有必要使用此标志。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示文件版本与指定时间的版本。
      
      参数应为日期，"2006-01-02"，日期时间 "2006-01-02 15:04:05" 或距离那时那么久之前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，当使用此选项时，不允许进行任何文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见 [时间选项文档](/docs/#time-option)。

   --decompress
      如果设置了此值，将解压缩 gzip 编码的对象。
      
      可以将文件上传到 S3，并设置 "Content-Encoding: gzip"。通常情况下，rclone 会将这些文件以压缩对象的形式下载。
      
      如果设置了此标志，则 rclone 会在接收到带有 "Content-Encoding: gzip" 的文件时进行解压缩。这意味着 rclone 无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会压缩对象，请设置此值。
      
      通常，提供商在下载时不会更改对象。如果对象未使用 `Content-Encoding: gzip` 进行上传，则在下载时不会设置该值。
      
      但是，有些提供商即使未使用 `Content-Encoding: gzip` 进行上传也会对对象进行压缩（例如 Cloudflare）。
      
      这种情况的症状可能是收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志并且 rclone 下载具有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，则 rclone 将动态解压缩对象。
      
      如果设置为 unset（默认值），则 rclone 将根据提供商设置选择要应用的方法，但您可以在此处覆盖 rclone 的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id value        AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预设 ACL。[$ACL]
   --endpoint value             Arvan Cloud 对象存储（AOS）API 的终端节点。[$ENDPOINT]
   --env-auth                   从运行时获取 AWS 凭据（环境变量或 EC2/ECS 元数据（如果没有环境变量））。 （默认值：false）[$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束-必须与终端节点匹配。[$LOCATION_CONSTRAINT]
   --secret-access-key value    AWS 秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value        在 ArvanCloud 中存储新对象时要使用的存储类。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               创建存储桶时使用的预设 ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分段复制的大小截止值。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置了此值，将解压缩 gzip 编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 http2 使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出的块大小（每个 ListObject S3 请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0（自动）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传的最大部分数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会压缩对象，请设置此值。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对已上传的对象进行 HEAD 请求以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行 HEAD 请求。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分段上传的大小截止值。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用 ETag 进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 进行单个部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 验证。 (default: false) [$V2_AUTH]
   --version-at value               显示文件版本与指定时间的版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]

```
{% endcode %}