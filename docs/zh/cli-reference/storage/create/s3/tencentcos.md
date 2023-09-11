# 腾讯云对象存储 (COS)

{% code fullWidth="true" %}
```
命令名称:
   singularity storage create s3 tencentcos - 腾讯云对象存储 (COS)

使用方法:
   singularity storage create s3 tencentcos [command options] [arguments...]

说明:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅当 access_key_id 和 secret_access_key为空时才适用。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。

      如果是匿名访问或运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。

      如果是匿名访问或运行时凭证，请留空。

   --endpoint
      腾讯COS API的端点。

      示例:
         | cos.ap-beijing.myqcloud.com       | 北京区域
         | cos.ap-nanjing.myqcloud.com       | 南京区域
         | cos.ap-shanghai.myqcloud.com      | 上海区域
         | cos.ap-guangzhou.myqcloud.com     | 广州区域
         | cos.ap-nanjing.myqcloud.com       | 南京区域
         | cos.ap-chengdu.myqcloud.com       | 成都区域
         | cos.ap-chongqing.myqcloud.com     | 重庆区域
         | cos.ap-hongkong.myqcloud.com      | 香港（中国）区域
         | cos.ap-singapore.myqcloud.com     | 新加坡区域
         | cos.ap-mumbai.myqcloud.com        | 孟买区域
         | cos.ap-seoul.myqcloud.com         | 首尔区域
         | cos.ap-bangkok.myqcloud.com       | 曼谷区域
         | cos.ap-tokyo.myqcloud.com         | 东京区域
         | cos.na-siliconvalley.myqcloud.com | 硅谷区域
         | cos.na-ashburn.myqcloud.com       | 弗吉尼亚区域
         | cos.na-toronto.myqcloud.com       | 多伦多区域
         | cos.eu-frankfurt.myqcloud.com     | 法兰克福区域
         | cos.eu-moscow.myqcloud.com        | 莫斯科区域
         | cos.accelerate.myqcloud.com       | 使用腾讯COS加速端点

   --acl
      创建存储桶和存储或复制对象时使用的预设ACL。

      此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建存储桶。

      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意，当服务端复制对象时，此ACL将被应用，因为 S3 不复制源的ACL，而是写入一个新的 ACL。

      如果ACL为空字符串，则不添加 X-Amz-Acl: 头，将使用默认桶（私有）。

      示例:
         | default | 拥有者获得 Full_CONTROL 权限。
         |         | 没有其他人有访问权限（默认）。

   --bucket-acl
      创建存储桶时使用的预设ACL。

      获取更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意，此ACL仅在创建存储桶时应用。如果未设置，则将使用 "acl"。

      如果 "acl" 和 "bucket_acl" 为空字符串，则不添加 X-Amz-Acl: 头，将使用默认桶（私有）。

      示例:
         | private            | 拥有者获得 FULL_CONTROL 权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 拥有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 用户组获得 READ 权限。
         | public-read-write  | 拥有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 用户组获得 READ 和 WRITE 权限。
         |                    | 不推荐在存储桶上授予此权限。
         | authenticated-read | 拥有者获得 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers 用户组获得 READ 权限。

   --storage-class
      在腾讯COS中存储新对象时要使用的存储类。

      示例:
         | <unset>     | 默认
         | STANDARD    | 标准存储类
         | ARCHIVE     | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换到分块上传的文件截止点。

      文件大小大于此值将以 chunk_size 的分块方式上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的分块大小。

      当上传大于upload_cutoff的文件或大小未知的文件（例如从 "rclone rcat" 或 "rclone mount" 上传的文件，或从 Google photos 或 Google docs 上传的文件）时，使用此分块大小进行分块上传。

      请注意，每个传输中的 "--s3-upload-concurrency" 个具有该大小的分块在内存中缓冲。

      如果您正在高速链接上传输大文件并且具有足够的内存，则增加此值将加快传输速度。

      当上传一个已知大小的大文件时，rclone将自动增加分块大小，以保持在 10,000 个分块的限制之下。

      未知大小的文件将使用配置的 chunk_size 进行上传。由于默认的块大小为 5 MiB，最多有 10,000 个块，这意味着默认情况下您可以流式上传的文件的最大大小为 48 GiB。如果要流式上传更大的文件，则需要增加 chunk_size。

      增加块大小会降低使用 "-P" 标志显示的进度统计的准确性。rclone 在 AWS SDK 缓冲块时，将发送块，而实际上可能仍在上传。更大的块大小意味着更大的 AWS SDK 缓冲区和与实际情况更为不符的进度报告。

   --max-upload-parts
      分块上传的最大部分数。

      此选项定义当执行分块上传时使用的多部件块的最大数目。

      如果服务不支持 AWS S3 的 10,000 个块规范，则此选项可能很有用。

      当上传已知大小的大文件时，rclone将自动增加分块大小，以保持在此块数量限制之下。

   --copy-cutoff
      切换到分块复制的文件截止点。

      需要分块复制的文件大小大于此值。

      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      在对象元数据中不存储MD5校验和。

      通常情况下，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这在数据完整性检查方面非常好，但对于大文件来说可能导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果 env_auth=true，则rclone可以使用共享凭证文件。

      如果此变量为空，则rclone将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境值为空，则默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果 env_auth=true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。

      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"，如果环境变量也没有设置，则为默认值。

   --session-token
      AWS会话Token。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传同一文件的块数。

      如果您正在高速链接上上传数量较少的大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。

      如果为 true（默认值），则rclone将使用路径样式访问；如果为 false，则rclone将使用虚拟路径样式访问。有关更多详细信息，请参阅 [AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些服务提供商（例如：AWS、阿里云 OSS、网易 COS 或腾讯 COS）要求此值设置为 false - rclone将根据提供商的设置自动完成。

   --v2-auth
      如果为 true，则使用v2认证。

      如果为 false（默认值），则rclone将使用v4认证。如果设置了此值，则rclone将使用v2认证。

      只有在v4签名不起作用时才使用此选项，例如：Jewel/v10 CEPH之前的版本。

   --list-chunk
      列出对象响应时的列表块大小（响应每个ListObject S3请求）。

      此选项也称为 AWS S3规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务即使要求更多，也会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅 [AWS S3文档](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用 "rgw list buckets max chunk" 选项增加此值。

   --list-version
      要使用的 ListObjects 版本：1、2 或 0自动。

      当S3最初发布时，它只提供了用于枚举存储桶中对象的ListObjects调用。

      但是，在2016年5月引入了ListObjectsV2调用。它具有更高的性能，应尽可能使用。

      如果设置为默认值0，则根据提供者设置猜测rclone将调用哪个list对象方法。如果它猜错了，则可以在此处手动设置。

   --list-url-encode
      是否对列表项进行URL编码：true/false/unset

      某些提供商支持URL编码，如果可用，则在使用控制字符的文件名时这更可靠。如果设置为未设置（默认值），则rclone将根据提供者设置选择要应用的设置，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。

      如果要尽量减少rclone执行的事务数目或者知道桶已经存在，则这可能很有用。

      如果要使用的用户没有桶创建权限，也可能需要使用此标志。在v1.52.0之前，由于一个错误，此操作将不会产生任何输出。

   --no-head
      如果设置，则不会 HEAD 已上传对象以检查完整性。

      如果要尽量减少rclone执行的事务数，则这可能很有用。

      如果rclone在使用 PUT 上传对象后收到200 OK消息，则会认为它已经成功上传。

      特别是它会假设：

      - 元数据，包括修改时间、存储类和内容类型与上传的元数据一致。
      - 大小与上传的大小一致。

      对于单个部分的 PUT，它从响应中读取以下项目：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项目。

      如果上传具有未知长度的源对象，则rclone **将**执行 HEAD 请求。

      设置此标志会增加无法检测到的上传失败的几率，特别是大小不正确的几率，因此不建议在正常操作中使用。实际上，即使在设置此标志时，无法检测到的上传失败的几率非常小。

   --no-head-object
      如果设置，则获取对象之前不会执行 HEAD。

   --encoding
      后端使用的编码方式。

      有关更多信息，请参阅 [概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。

      需要额外的缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP/2使用。

      当前，s3（特别是minio）后端和HTTP/2存在一个未解决的问题。默认情况下，s3后端启用HTTP/2，但可以在此禁用。解决此问题后，此标志将被移除。

      请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      这通常设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的数据的更低出站流量费用。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证。
      
      这应该设置为 true、false 或不设置以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或 PutObject 进行单部分上传。
      
      如果为 false，则rclone将使用 AWS SDK 的 PutObject 上传对象。
      
      rclone的旧版本 < 1.59 使用预签名请求上传单个部分对象，将此标志设置为 true 将重新启用该功能。除了特殊情况或测试之外，这不应该是必需的。

   --versions
      在目录列表中包括旧版本。

   --version-at
      以指定时间查看文件版本。
      
      该参数应为日期 "2006-01-02"、日期时间 "2006-01-02 15:04:05" 或表示多久之前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此参数时不允许进行文件写操作，因此您不能上传文件或删除文件。
      
      有关有效格式，请参阅 [时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以使用 "Content-Encoding: gzip" 上传对象到S3。通常情况下，rclone将以压缩对象形式下载这些文件。
      
      如果设置了此标志，则rclone将在收到具有 "Content-Encoding: gzip" 的文件时解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行gzip，则设置此值。
      
      通常情况下，提供商在下载对象时不会更改对象。如果一个对象即使在上传时没有使用 `Content-Encoding: gzip` 也会被gzip压缩（例如 Cloudflare），就可能会遇到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了具有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，则rclone将在接收时实时解压缩对象。
      
      如果设置为未设置（默认值），则rclone将根据提供者设置选择要应用的设置，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

OPTIONS:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value           腾讯COS API的端点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      在腾讯COS中存储新对象时要使用的存储类。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的文件截止点。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               在对象元数据中不存储MD5校验和。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端使用的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出对象响应时的列表块大小。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表项进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0自动。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传的最大部分数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip，则设置此值。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会 HEAD 已上传对象以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，则获取对象之前不会执行 HEAD。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话Token。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的文件截止点。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用v2认证。 (default: false) [$V2_AUTH]
   --version-at value               以指定时间查看文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (default: false) [$VERSIONS]

   General

   --name value  存储的名称（默认：自动生成）
   --path value  存储的路径

```
{% endcode %}