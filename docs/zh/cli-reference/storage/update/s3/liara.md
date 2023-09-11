# Liara 对象存储

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 liara - Liara 对象存储

USAGE:
   singularity storage update s3 liara [命令选项] <名称|ID>

DESCRIPTION:
   --env-auth
      从运行时获取 AWS 凭证（从环境变量或 EC2/ECS 元数据获取，如果没有环境变量）。
      
      仅在 access_key_id 和 secret_access_key 为空时生效。

      示例:
         | false | 在下一步中输入 AWS 凭证。
         | true  | 从环境中获取 AWS 凭证（环境变量或 IAM）。

   --access-key-id
      AWS 访问密钥 ID。
      
      保留为空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS 秘密访问密钥（密码）。
      
      保留为空以进行匿名访问或运行时凭证。

   --endpoint
      Liara 对象存储 API 的终端节点。

      示例:
         | storage.iran.liara.space | 默认终端节点
         |                          | 伊朗

   --acl
      创建存储桶、存储或复制对象时使用的预定义 ACL。
      
      此 ACL 用于创建对象，并且如果未设置 bucket_acl，则用于创建存储桶。
      
      有关详情，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，在使用 S3 进行服务器端复制对象时，S3 不会复制源对象的 ACL，而是写入新的 ACL。
      
      如果 acl 是空字符串，则不会添加 X-Amz-Acl: 标头，并将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预定义 ACL。
      
      有关详情，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，在仅创建存储桶时应用此 ACL。如果未设置它，则使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 都是空字符串，则不会添加 X-Amz-Acl:
      标头，并将使用默认值（private）。

      示例:
         | private            | 拥有者拥有完全控制权。
         |                    | 没有其他用户具有访问权限（默认）。
         | public-read        | 拥有者拥有完全控制权。
         |                    | AllUsers 用户组具有读取权限。
         | public-read-write  | 拥有者拥有完全控制权。
         |                    | AllUsers 用户组具有读写权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者拥有完全控制权。
         |                    | AuthenticatedUsers 用户组具有读取权限。

   --storage-class
      在 Liara 中存储新对象时要使用的存储类别。

      示例:
         | STANDARD | 标准存储类别

   --upload-cutoff
      切换到分块上传的截断点。
      
      任何大于此大小的文件都将按块大小进行上传。最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于 upload_cutoff 的文件或具有未知大小的文件（例如使用 "rclone rcat" 或使用 "rclone mount" 或 Google 照片或 Google 文档上传）时，将使用该块大小进行分块上传。
      
      注意，每次传输都会在内存中缓冲 "--s3-upload-concurrency" 个这样大小的块。
      
      如果你正在通过高速链接传输大文件，并且有足够的内存，则增加此值将加快传输速度。
      
      Rclone 将根据需要增加块大小，以确保在上传已知大小的大文件时保持在 10,000 个块的限制以下。
      
      未知大小的文件将使用配置的 chunk_size 进行上传。默认块大小为 5 MiB，最多可以有 10,000 个块，这意味着默认情况下你可以流式上传的文件的最大大小为 48 GiB。如果要流式上传较大的文件，则需要增加块大小。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。当 AWS SDK 缓冲一个块时，rclone 会将该块视为已发送，而实际上可能仍在上传。较大的块大小意味着更大的 AWS SDK 缓冲区，并且进度报告离实际情况更远。
      

   --max-upload-parts
      分块上传的最大块数。
      
      该选项定义分块上传时要使用的最大块数。
      
      如果某个服务不支持 AWS S3 的规范（10,000 个块），这将非常有用。
      
      Rclone 将根据需要增加块大小，以确保在上传已知大小的大文件时保持在此块数限制以下。
      

   --copy-cutoff
      切换到分块复制的截断点。
      
      任何大于此大小需要进行服务器端复制的文件都将按该大小进行复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常情况下，rclone 会在上传之前计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但会导致大型文件启动上传时出现长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。
      
      如果此变量为空，则 rclone 将查找
      "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭据文件。此
      变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果也未设置环境变量）。
      

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果你正在通过高速链接上传少量大文件，并且这些上传未充分利用你的带宽，那么增加这个值可能会帮助加快传输速度。

   --force-path-style
      如果为 true，请使用路径样式访问；如果为 false，请使用虚拟主机样式访问。
      
      如果为 true（默认），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟路径样式。有关详情，请参阅 [AWS S3
      文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      。
      
      某些提供者（例如 AWS、Aliyun OSS、Netease COS 或 Tencent COS）要求此值设置为
      false - rclone 会根据提供者设置自动完成此操作。

   --v2-auth
      如果为 true，请使用 v2 认证。
      
      如果为 false（默认），则 rclone 将使用 v4 认证。如果设置了此项，则 rclone 将使用 v2 认证。
      
      仅在 v4 签名无效时使用，例如 Jewel/v10 CEPH 之前的版本。

   --list-chunk
      列表块的大小（每个 ListObject S3 请求的响应列表）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务即使请求的值大于 1000，也会截断响应列表为 1000 个对象。
      在 AWS S3 中，这是一个全局最大值，无法更改，请参阅 [AWS
      S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项增加此值。
      

   --list-version
      要使用的 ListObjects 版本：1、2 或 0 用于自动。
      
      当 S3 最初推出时，它只提供了 ListObjects 调用来列举存储桶中的对象。
      
      但是，在2016年5月，ListObjectsV2 调用被引入。这是性能更高的方法，应尽可能使用。
      
      如果设置为默认值 0，则 rclone 将根据设置的提供者猜测要调用的列表对象方法。如果猜错了，可以在这里手动设置。
      

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      一些提供者支持 URL 编码列表，如果可用，则在文件名中使用控制字符时，这是更可靠的方法。如果将其设置为未设置（默认值），则 rclone 将根据提供者设置选择适用的方法，但可以在此处覆盖 rclone 的选择。
      

   --no-check-bucket
      如果设置了此项，则不会尝试检查存储桶是否存在或创建它。
      
      如果你知道存储桶已经存在，则此项可以有助于尽量减少 rclone 执行的事务数。
      
      如果你使用的用户没有创建存储桶的权限，则可能也需要使用此项。在 v1.52.0 之前，此操作将静默通过由于一个 bug。
      

   --no-head
      如果设置了此项，则不会检查上传的对象的头部以检查完整性。
      
      如果尽量减少 rclone 执行的事务数，则此项可以有助于达到目的。
      
      设置此项意味着如果 rclone 在使用 PUT 上传对象后收到 200 OK 的响应，则它将假设对象已正确上传。
      
      特别是它将假设：
      
      - 元数据（包括修改时间、存储类别和内容类型）与上传的对象相同
      - 大小与上传的对象相同
      
      它从单个部分 PUT 的响应中读取以下内容：
      
      - MD5 校验和
      - 上传日期
      
      对于分块上传，不会读取这些内容。
      
      如果上传一个未知长度的源对象，则 rclone **将**执行 HEAD 请求。
      
      设置此标志会增加无法检测到的上传故障的机会，特别是大小不正确的文件，因此不建议在正常操作中使用。实际上，即使设置此标志，出现无法检测到的上传故障的机会也非常小。
      

   --no-head-object
      如果设置了此项，则在获取对象时不执行 HEAD 请求。

   --encoding
      后端的编码方式。
      
      有关详情，请参阅[概览中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区的上传（例如分块上传）将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。
      
      目前，s3 后端（特别是 minio）与 HTTP/2 存在未解决的问题。默认情况下，S3 后端启用了 HTTP/2，但可以在此禁用。解决问题后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      通常将此设置为 CloudFront CDN URL，因为 AWS S3 通过 CloudFront 网络下载的数据提供了更便宜的外发流量。

   --use-multipart-etag
      是否在分块上传中使用 ETag 进行验证
      
      此项应为 true、false 或保持未设置以使用提供者的默认值。
      

   --use-presigned-request
      是否使用预签名请求还是 PutObject 进行单个部分上传
      
      如果此项为 false，则 rclone 将使用 AWS SDK 的 PutObject 来上传对象。
      
      低于 rclone 1.59 版本使用预签名请求来上传单个部分对象，将此标志设置为 true 将重新启用该功能。除非在特殊情况下或进行测试，否则不应该使用此项。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定时间显示文件版本。
      
      参数应该是日期 "2006-01-02"、日期时间 "2006-01-02
      15:04:05" 或距离现在那么长时间的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见 [时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置了此项，则将解压缩 gzip 编码的对象。
      
      可以将对象上传到 S3，并设置 "Content-Encoding: gzip"。通常，rclone 会将这些文件下载为压缩对象。
      
      如果设置了此标志，则 rclone 会在接收时解压缩这些文件，其中 "Content-Encoding: gzip"。这意味着 rclone 无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能对对象进行 gzip 压缩，请设置此项。
      
      通常情况下，提供者在下载对象时不会更改对象。如果对象未使用 `Content-Encoding: gzip` 进行上传，则在下载时不会设置该标头。
      
      但是，某些提供者可能会 gzip 压缩对象，即使它们没有使用 `Content-Encoding: gzip` 进行上传（例如 Cloudflare）。
      
      接到错误消息如下表示存在此问题：
      
         ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此项并且 rclone 下载了带有设置了 Content-Encoding: gzip 且分块传输编码的对象，则 rclone 将在接收时动态解压缩该对象。
      
      如果设置为 unset（默认值），则 rclone 将根据提供者设置选择适用的方法，但可以在此处覆盖 rclone 的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value      AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶、存储或复制对象时使用的预定义 ACL。[$ACL]
   --endpoint value           Liara 对象存储 API 的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取 AWS 凭证（从环境变量或 EC2/ECS 元数据获取，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  AWS 秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      在 Liara 中存储新对象时要使用的存储类别。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义 ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截断点。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置了此项，则将解压缩 gzip 编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 http2 使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，请使用路径样式访问；如果为 false，请使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个 ListObject S3 请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0 用于自动。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传的最大块数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行 gzip 压缩，请设置此项。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置了此项，则不会尝试检查存储桶是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置了此项，则不会检查上传的对象的头部以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置了此项，则在获取对象时不执行 HEAD 请求。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截断点。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用 ETag 进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求还是 PutObject 进行单个部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，请使用 v2 认证。 (default: false) [$V2_AUTH]
   --version-at value               按指定时间显示文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]


```
{% endcode %}