# 七牛对象存储（Kodo）

{% code fullWidth="true" %}
```
命名：
   singularity storage create s3 qiniu - 七牛对象存储（Kodo）

使用方法：
   singularity storage create s3 qiniu [命令选项] [参数...]

说明：
   --env-auth
      从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据）。

      仅在access_key_id和secret_access_key为空时适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。

      如果要进行匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。

      如果要进行匿名访问或使用运行时凭证，请留空。

   --region
      要连接的区域。

      示例：
         | cn-east-1      | 默认终端节点 - 如果您不确定，请选择此选项。
         |                | 华东区域1。
         |                | 需要位置限制cn-east-1。
         | cn-east-2      | 华东区域2。
         |                | 需要位置限制cn-east-2。
         | cn-north-1     | 华北区域1。
         |                | 需要位置限制cn-north-1。
         | cn-south-1     | 华南区域1。
         |                | 需要位置限制cn-south-1。
         | us-north-1     | 北美区域。
         |                | 需要位置限制us-north-1。
         | ap-southeast-1 | 东南亚区域1。
         |                | 需要位置限制ap-southeast-1。
         | ap-northeast-1 | 东北亚区域1。
         |                | 需要位置限制ap-northeast-1。

   --endpoint
      七牛对象存储的终端节点。

      示例：
         | s3-cn-east-1.qiniucs.com      | 华东终端节点1
         | s3-cn-east-2.qiniucs.com      | 华东终端节点2
         | s3-cn-north-1.qiniucs.com     | 华北终端节点1
         | s3-cn-south-1.qiniucs.com     | 华南终端节点1
         | s3-us-north-1.qiniucs.com     | 北美终端节点1
         | s3-ap-southeast-1.qiniucs.com | 东南亚终端节点1
         | s3-ap-northeast-1.qiniucs.com | 东北亚终端节点1

   --location-constraint
      地理位置限制 - 必须设置与区域匹配。

      仅在创建存储桶时使用。

      示例：
         | cn-east-1      | 华东区域1
         | cn-east-2      | 华东区域2
         | cn-north-1     | 华北区域1
         | cn-south-1     | 华南区域1
         | us-north-1     | 北美区域1
         | ap-southeast-1 | 东南亚区域1
         | ap-northeast-1 | 东北亚区域1

   --acl
      创建存储桶和存储或复制对象时使用的预设访问控制列表（ACL）。

      此ACL用于创建对象，并且如果未设置bucket_acl，则也用于创建存储桶。

      更多信息请参考[这里](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      注意，当服务器端复制对象时，rclone会对ACL进行应用，因为S3不会复制来自源的ACL，而是会写入新的ACL。

      如果acl是一个空字符串，则不会添加X-Amz-Acl：标题，并使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预设访问控制列表（ACL）。

      更多信息请参考[这里](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      注意，仅在创建存储桶时应用此ACL。如果未设置，则使用“acl”。

      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl：标题，并使用默认值（private）。

      示例：
         | private            | 拥有者获得FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 拥有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得READ权限。
         | public-read-write  | 拥有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得READ和WRITE权限。
         |                    | 不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者获得FULL_CONTROL权限。
         |                    | AuthenticatedUsers组获得READ权限。

   --storage-class
      存储新对象时使用的存储类。

      示例：
         | STANDARD     | 标准存储类
         | LINE         | 低频访问存储模式
         | GLACIER      | 归档存储模式
         | DEEP_ARCHIVE | 深度归档存储模式

   --upload-cutoff
      切换到分块上传的临界点。

      任何大于此值的文件将以chunk_size的大小进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。

      当上传大于upload_cutoff的文件或未知大小的文件（例如使用“rclone rcat”上传的文件或使用“rclone mount”或Google
      相册或Google文档上传的文件）时，将使用此分块大小进行分块上传。

      请注意，每个传输的缓冲区中将缓冲大小为“--s3-upload-concurrency”个分块大小。

      如果您在高速链接上传输大文件并且具有足够的内存，则增加此值将加快传输速度。

      Rclone将在上传已知大小的大文件时自动增加分块大小，以保持在10000个分块的限制之下。

      大小未知的文件将使用配置的分块大小进行上传。由于默认的分块大小为5 MiB，并且最多可以有10000个分块，因此默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果要流式上传更大的文件，则需要增加chunk_size。

      增加分块大小会降低使用“-P”标记显示的进度统计的准确性。当Rclone将缓冲到AWS SDK的分块发送时，Rclone会将分块视为已发送，而实际上可能仍在上传。较大的分块大小意味着更大的AWS SDK缓冲区和与真实情况更偏离的进度报告。

   --max-upload-parts
      多部分上传中的最大部分数。

      该选项定义在执行多部分上传时使用的最大多部分分块数。

      如果服务不支持AWS S3的10000个分块规范，则此选项可能非常有用。

      Rclone将在上传已知大小的大文件时自动增加分块大小，以保持在此部分数限制之下。

   --copy-cutoff
      切换到分块复制的临界点。

      需要进行分块复制的大于此值的文件将以此大小进行复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和存储为对象元数据。

      通常情况下，rclone会在上传文件之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件启动上传过程可能会导致长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth为true，则rclone可以使用共享凭证文件。

      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量的值为空，则默认为当前用户的主目录。

            Linux/OSX: "$HOME/.aws/credentials"
            Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth为true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。

      如果为空，则默认为环境变量“AWS_PROFILE”或“default”（如果该环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发度。

      这是同时上传相同文件的块数。

      如果您在高速链接上上传大量大文件，并且这些上传未充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟主机样式访问。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求将其设置为false - rclone将根据提供商设置自动完成此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，rclone将使用v2身份验证。

      仅当v4签名不起作用时才使用此选项，例如旧版Jewel/v10 CEPH。

   --list-chunk
      列出的列表块的大小（响应每个ListObject S3请求的列表）。

      此选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。大多数服务即使请求超过了1000个对象，也会截断响应列表。在AWS
      S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。在Ceph中，可以通过“rgw
      list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0为自动。

      S3最初推出时，它只提供了“ListObjects”调用来枚举存储桶中的对象。

      但是，在2016年5月，引入了“ListObjectsV2”调用。这是性能更高，如果可能，应该使用它。

      如果设置为默认值0，则rclone将根据提供商设置猜测要调用哪个ListObjects方法。如果猜测错误，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      一些提供商支持URL编码列表，如果可用，则使用控制字符在文件名中更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的编码，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不会尝试检查存储桶是否存在或创建存储桶。

      当尝试尽量减少rclone的事务数量时，如果知道存储桶已存在，则这可能很有用。

      如果使用的用户没有存储桶创建权限，则可能也需要此选项。在v1.52.0之前，由于一个错误，这将导致默默地传递。

   --no-head
      如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。

      当尝试尽量减少rclone的事务数量时，这可能很有用。

      设置后，如果rclone在PUT之后收到200 OK消息，则会假设对象已正确上传。

      具体而言，它将假设：

      - 元数据，包括修改时间，存储类和内容类型与上传的元数据相同
      - 大小与上传的大小相同

      对于单个部分PUT的响应，它会读取以下项：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项。

      如果上传长度未知的源对象，则rclone将进行HEAD请求。

      设置此标志会增加未检测到的上传失败的机会，特别是错误大小的机会，因此不建议在正常操作中使用此标志。在实践中，即使使用此标志，未检测到的上传失败的机会也非常小。

   --no-head-object
      如果设置，则在执行GET操作获取对象之前不会执行HEAD请求。

   --encoding
      后端的编码方式。

      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。

      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2的使用。

      目前，s3（特别是minio）后端存在一个问题，无法解决http2的问题。为了解决这个问题，s3后端默认启用了HTTP/2，但可以在此禁用。解决了这个问题后，将删除该标志。

      请参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      自定义下载的终端节点。
      这通常设置为CloudFront CDN的URL，因为AWS S3通过CloudFront网络下载的数据获得更便宜的流出费用。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证。

      此值应设置为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传。

      如果此值为false，则rclone将使用AWS SDK的PutObject上传对象。

      rclone
      < 1.59的版本使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或进行测试，否则不应该需要此标志。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定时间显示文件版本。

       参数应为日期“2006-01-02”、日期时间“2006-01-02 15:04:05”或距那时太远的持续时间，例如“100d”或“1h”。

      请注意，使用此选项后，不允许进行任何文件写操作，因此您无法上传文件或删除文件。

      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置此项，将解压缩gzip编码的对象。

      可以将对象上传到S3时设置“Content-Encoding：gzip”。通常情况下，rclone将这些文件作为压缩对象下载。

      如果设置了此标志，则rclone将在接收到这些带有“Content-Encoding：gzip”的文件时进行解压缩。这意味着rclone不能检查大小和哈希，但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此值。

      通常情况下，提供者在下载对象时不会更改对象。如果一个对象在上传时没有使用`Content-Encoding：gzip`设置，那么在下载时也不会设置。

      但是，某些提供者即使在未上传时使用`Content-Encoding：gzip`（例如Cloudflare）也可能对对象进行gzip压缩。

      这种情况的症状可能是收到以下错误：

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置此标志，并且rclone下载具有设置了Content-Encoding：gzip和分块传输编码的对象，则rclone将在接收到对象时动态解压缩对象。

      如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的选项，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id 值        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl 值                  创建存储桶和存储或复制对象时使用的预设访问控制列表（ACL）。[$ACL]
   --endpoint 值             七牛对象存储的终端节点。[$ENDPOINT]
   --env-auth                从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据）。 (默认值: false) [$ENV_AUTH]
   --location-constraint 值  地理位置限制 - 必须设置与区域匹配。[$LOCATION_CONSTRAINT]
   --region 值               要连接的区域。[$REGION]
   --secret-access-key 值    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class 值        存储新对象时使用的存储类。[$STORAGE_CLASS]

   高级设置

   --bucket-acl 值               创建存储桶时使用的预设访问控制列表（ACL）。[$BUCKET_ACL]
   --chunk-size 值               用于上传的分块大小。 (缺省值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff 值              切换到分块复制的临界点。 (缺省值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                  如果设置此项，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum            不要将MD5校验和存储为对象元数据。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2               禁用S3后端的http2的使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url 值             自定义下载的终端节点。[$DOWNLOAD_URL]
   --encoding 值                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style            如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk 值               列出的列表块的大小（响应每个ListObject S3请求的列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode 值          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version 值             要使用的ListObjects版本：1、2或0为自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts 值         多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time 值   内部内存缓冲池将刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap        是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip 值               如果后端可能对对象进行gzip压缩，请设置此值。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket             如果设置，则不会尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                     如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object              如果设置，则在执行GET操作获取对象之前不会执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata          禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile 值                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token 值            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file 值  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency 值       多部分上传的并发度。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 值            切换到分块上传的临界点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag 值       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request       是否使用预签名请求或PutObject进行单个部分上传。 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                     如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at 值               按指定时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                    在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name 值  存储的名称（默认: 自动生成的）
   --path 值  存储的路径

```
{% endcode %}