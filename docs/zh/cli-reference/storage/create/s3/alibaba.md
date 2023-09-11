# 阿里云对象存储服务 (OSS) 原名阿里云

{% code fullWidth="true" %}
```
命令名称:
   singularity storage create s3 alibaba - 阿里云对象存储服务 (OSS)

使用方法:
   singularity storage create s3 alibaba [命令选项] [参数]

描述:
   --env-auth
      从运行时（环境变量或EC2 / ECS元数据，如果没有环境变量）获取AWS凭证。

      仅在access_key_id和secret_access_key为空时适用。

      示例：
         | false | 输入下一步中的AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS Access Key ID。

      如果要匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS Secret Access Key（密码）。

      如果要匿名访问或使用运行时凭证，请留空。

   --endpoint
      OSS API的终端节点。

      示例：
         | oss-accelerate.aliyuncs.com          | 全球加速
         | oss-accelerate-overseas.aliyuncs.com | 全球加速（中国境外）
         | oss-cn-hangzhou.aliyuncs.com         | 华东1（杭州）
         | oss-cn-shanghai.aliyuncs.com         | 华东2（上海）
         | oss-cn-qingdao.aliyuncs.com          | 华北1（青岛）
         | oss-cn-beijing.aliyuncs.com          | 华北2（北京）
         | oss-cn-zhangjiakou.aliyuncs.com      | 华北3（张家口）
         | oss-cn-huhehaote.aliyuncs.com        | 华北5（呼和浩特）
         | oss-cn-wulanchabu.aliyuncs.com       | 华北6（乌兰察布）
         | oss-cn-shenzhen.aliyuncs.com         | 华南1（深圳）
         | oss-cn-heyuan.aliyuncs.com           | 华南2（河源）
         | oss-cn-guangzhou.aliyuncs.com        | 华南3（广州）
         | oss-cn-chengdu.aliyuncs.com          | 西南1（成都）
         | oss-cn-hongkong.aliyuncs.com         | 中国香港特别行政区（香港）
         | oss-us-west-1.aliyuncs.com           | 美国西部1（硅谷）
         | oss-us-east-1.aliyuncs.com           | 美国东部1（弗吉尼亚）
         | oss-ap-southeast-1.aliyuncs.com      | 东南亚东南1（新加坡）
         | oss-ap-southeast-2.aliyuncs.com      | 亚太地区东南2（悉尼）
         | oss-ap-southeast-3.aliyuncs.com      | 东南亚东南3（吉隆坡）
         | oss-ap-southeast-5.aliyuncs.com      | 亚太地区东南5（雅加达）
         | oss-ap-northeast-1.aliyuncs.com      | 亚太地区东北1（日本）
         | oss-ap-south-1.aliyuncs.com          | 亚太地区南部1（孟买）
         | oss-eu-central-1.aliyuncs.com        | 中欧1（法兰克福）
         | oss-eu-west-1.aliyuncs.com           | 西欧（伦敦）
         | oss-me-east-1.aliyuncs.com           | 中东1（迪拜）

   --acl
      创建存储桶和存储或复制对象时使用的 canned ACL。

      此 ACL 用于创建对象，并且如果未设置 bucket_acl，则用于创建存储桶。
      
      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，当进行服务器端复制对象时，此 ACL 会生效，因为S3不会复制源中的ACL，而是写入一个新的ACL。
      
      如果 acl 是空字符串，则不会添加 X-Amz-Acl: 头，并且将使用默认 (private)。

   --bucket-acl
      创建存储桶时使用的 canned ACL。

      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，仅在创建存储桶时应用此 ACL。如果未设置，则使用 "acl"。
      
      如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加 X-Amz-Acl: 头，并且将使用默认 (private)。

      示例：
         | private            | 所有者获得 FULL_CONTROL 权限。
         |                    | 任何其他人没有访问权限（默认）。
         | public-read        | 所有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 组获得 READ 访问权限。
         | public-read-write  | 所有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 组获得 READ 和 WRITE 访问权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者获得 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers 组获得 READ 访问权限。

   --storage-class
      向 OSS 存储新对象时要使用的存储类。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类
         | GLACIER     | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换为分块上传的截止点。

      大于此大小的任何文件将被分块以 chunk_size 传输。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。

      当上传大于 upload_cutoff 的文件或大小未知的文件（例如通过 "rclone rcat" 上传的文件或通过 "rclone mount" 或谷歌照片或谷歌文档上传的文件）时，将使用此分块大小以多部分上传的方式传输。
      
      请注意，每次传输内存中缓冲的 "--s3-upload-concurrency" 个分块都以此大小进行缓冲。
      
      如果您正在通过高速链接传输大文件，并且具有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件以保持在10,000个块限制以下时，rclone将自动增加块大小。
      
      具有未知大小的文件将使用配置的 chunk_size 进行上传。
      由于默认的 chunk_size 为5 MiB，最多可以有10,000个块，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果要流式传输更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。
      当 AWS SDK 缓冲块时，rclone将分块视为已发送，但实际上可能仍在上传。
      更大的块大小意味着更大的 AWS SDK 缓冲区，并且进度报告与实际情况可能更偏离。
      

   --max-upload-parts
      分块上传中的最大块数。

      此选项定义在执行分块上传时要使用的最大多部分块数。
      
      如果某个服务不支持10,000个块的 AWS S3 规范，则此选项可能很有用。
      
      当上传已知大小的大文件以保持在此块数上限以下时，rclone将自动增加块大小。
      

   --copy-cutoff
      切换为分块复制的截止点。

      需要分块复制的大于此大小的文件将被以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要在对象元数据中存储 MD5 校验和。

      通常，在上传之前，rclone会计算输入的MD5校验和，以便将其添加到对象的元数据中。这非常适用于数据完整性检查，但对于大文件开始上传可能导致较长的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果 env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将搜索 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。
      如果环境值为空，则它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。

      同一文件的并发块数。
      
      如果通过高速链接上传较少量的大文件，并且这些上传未能充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果设置为 true，则使用路径样式访问；如果设置为false，则使用虚拟托管样式访问。
      
      如果为 true（默认情况），则rclone将使用路径样式访问；
      如果为 false，则rclone将使用虚拟路径样式访问。请查看 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) 获取更多信息。
      
      某些提供商（例如 AWS、阿里云 OSS、网易 COS 或腾讯 COS）要求将其设置为 false，rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果设置为 true，则使用v2身份验证。

      如果为 false（默认情况），则rclone将使用v4身份验证。如果设置了该值，则rclone将使用v2身份验证。
      
      仅在v4签名无法工作时使用。例如 pre Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个 ListObject S3 请求的响应列表大小）。

      此选项也称为 "MaxKeys"、"max-items" 或 "page-size"，来源于 AWS S3 规范。
      大多数服务将响应列表截断为1000个对象，即使请求了多个对象。
      在 AWS S3 上，这是全局最大值，无法更改，参见 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以通过 "rgw list buckets max chunk" 选项进行增加。
      

   --list-version
      要使用的 ListObjects 版本：1、2 或 0（自动）。

      在 S3 最初发布时，它只提供了用于枚举存储桶中的对象的 ListObjects 调用。

      然而，于 2016 年 5 月，引入了 ListObjectsV2 调用。这个调用性能更好，应尽可能使用。
      
      如果设置为默认值 0，则rclone将根据设置的提供商猜测要调用哪个列出对象方法。如果它猜错了，那么可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行 URL 编码：true / false / unset
      
      一些提供商支持 URL 编码列表，如果可用，则在文件名中使用控制字符时这更可靠。
      如果设置为 unset（默认值），则rclone将根据提供商的设置来选择如何应用，但可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，并且想要尽量减少rclone的事务数，则这可能很有用。
      
      如果您使用的用户没有创建存储桶的权限，则可能需要这样做。在 v1.52.0 之前，这将无声地通过，因为存在一个错误。
      

   --no-head
      如果设置，则不进行 HEAD 请求以检查对象完整性。
      
      如果您想要尽量减少rclone的事务数，则这可能很有用。
      
      如果rclone上传对象后收到200 OK消息，并且没有 PUT 中的错误消息，那么rclone将假设文件已经正确上传。
      
      特别的，它将假设：
      
      - 元数据（包括 modtime、存储类别和内容类型）是上传时的元数据。
      - 大小与上传时的大小一致。
      
      它从单个部分 PUT 的响应中读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象的长度未知，则rclone **将**执行 HEAD 请求。
      
      设置此标志会增加未检测到的上传失败的机会，
      特别是大小不正确，因此不建议在正常操作中使用。实际上，即使有此标志，未检测到上传失败的几率非常小。
      

   --no-head-object
      如果设置，则在进行对象获取之前不执行 HEAD 请求。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见概述中的 [编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      刷新内部内存缓冲池的时间间隔。
      
      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。
      
      目前，s3（特别是 minio）后端存在一个无法解决的 http2 问题。
      默认情况下，s3 后端启用了 HTTP/2，但可以在此禁用。
      在问题解决后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      通常将其设置为 CloudFront CDN URL，因为通过 CloudFront 网络下载的数据可以享受 AWS S3 提供的更低廉的外发流量。

   --use-multipart-etag
      在多部分上传中是否使用 ETag 来进行验证。
      
      这应该为 true、false 或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求还是 PutObject 来上传单个部分对象。
      
      如果设置为 false，则rclone将使用 AWS SDK 的 PutObject 来上传对象。
      
      rclone 1.59 版本以下的版本使用预签名请求来上传单个部分对象，
      将这个标志设置为 true将重新启用该功能。
      除非情况特殊或用于测试，否则不应该有必要使用这个标志。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      指定时间点的文件版本。
      
      参数应为日期 "2006-01-02"、日期时间 "2006-01-02 15:04:05" 或距现在时间的持续时间，例如 "100d" 或 "1h"。
      
      请注意，在使用此功能时，不允许进行文件写入操作，因此不能上传文件或删除文件。
      
      有关有效格式，请参见 [时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩 gzip 编码的对象。
      
      可以使用 "Content-Encoding: gzip" 在上传对象到 S3 时设置对象。
      通常，rclone会将这些文件作为压缩对象进行下载。
      
      如果设置了此标志，rclone将在接收到带有 "Content-Encoding: gzip" 的文件时对其进行解压缩。这意味着 rclone 无法检查大小和哈希，但文件内容将会被解压缩。
      

   --might-gzip
      如果后端可能 gzip 对象，请设置此标志。
      
      通常，提供商在下载时不会更改对象。如果一个对象未使用 `Content-Encoding: gzip` 进行上传，则在下载时也不会设置该值。
      
      但是，一些提供商可能在没有使用 `Content-Encoding: gzip` 进行上传的情况下对对象进行 gzip 编码（例如 Cloudflare）。
      
      这种情况的症状是收到诸如
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      的错误消息。
      
      如果设置了此标志，并且 rclone 下载了带有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，
      那么 rclone 将在获取时对对象进行解压缩。
      
      如果未设置为 unset（默认值），则rclone将根据提供商的设置来选择如何应用，但可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项说明:
   --access-key-id value      AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的 canned ACL。[$ACL]
   --endpoint value           OSS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行时（环境变量或EC2 / ECS元数据，如果没有环境变量）获取AWS凭证。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --secret-access-key value  AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      向 OSS 存储新对象时要使用的存储类。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               创建存储桶时使用的 canned ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要在对象元数据中存储MD5校验和。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 http2 使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟托管样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个 ListObject S3 请求的响应列表大小）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0（自动）。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   刷新内部内存缓冲池的时间间隔。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此标志，以防后端可能对对象进行 gzip 编码。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不进行HEAD请求以检查对象完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在进行对象获取之前不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       在多部分上传中是否使用 ETag 来进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求还是 PutObject 来上传单个部分对象 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               指定时间点的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}