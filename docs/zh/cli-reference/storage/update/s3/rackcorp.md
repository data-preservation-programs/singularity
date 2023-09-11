# RackCorp对象存储

{% code fullWidth="true" %}
```
命令名称：
   singularity storage update s3 rackcorp - RackCorp对象存储

用法：
   singularity storage update s3 rackcorp [命令选项] <名称|id>

描述：
   --env-auth
      从运行时获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅当访问密钥ID和机密访问密钥为空时适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。

      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。

      留空以进行匿名访问或运行时凭证。

   --region
      region - 您的存储桶将被创建的位置以及您的数据存储的位置。

      示例：
         | global    | 全球CDN（所有地点）区域
         | au        | 澳大利亚（所有州）
         | au-nsw    | 新南威尔士州（澳大利亚）区域
         | au-qld    | 昆士兰州（澳大利亚）区域
         | au-vic    | 维多利亚州（澳大利亚）区域
         | au-wa     | 珀斯（澳大利亚）区域
         | ph        | 马尼拉（菲律宾）区域
         | th        | 曼谷（泰国）区域
         | hk        | 香港区域
         | mn        | 乌兰巴托（蒙古）区域
         | kg        | 比什凯克（吉尔吉斯斯坦）区域
         | id        | 雅加达（印度尼西亚）区域
         | jp        | 东京（日本）区域
         | sg        | 新加坡（新加坡）区域
         | de        | 法兰克福（德国）区域
         | us        | 美国（任意播送）区域
         | us-east-1 | 纽约（美国）区域
         | us-west-1 | 维尔蒙特（美国）区域
         | nz        | 奥克兰（新西兰）区域

   --endpoint
      RackCorp对象存储的终端节点。

      示例：
         | s3.rackcorp.com           | 全球（AnyCast）终端节点
         | au.s3.rackcorp.com        | 澳大利亚（Anycast）终端节点
         | au-nsw.s3.rackcorp.com    | 悉尼（澳大利亚）终端节点
         | au-qld.s3.rackcorp.com    | 布里斯班（澳大利亚）终端节点
         | au-vic.s3.rackcorp.com    | 墨尔本（澳大利亚）终端节点
         | au-wa.s3.rackcorp.com     | 珀斯（澳大利亚）终端节点
         | ph.s3.rackcorp.com        | 马尼拉（菲律宾）终端节点
         | th.s3.rackcorp.com        | 曼谷（泰国）终端节点
         | hk.s3.rackcorp.com        | 香港终端节点
         | mn.s3.rackcorp.com        | 乌兰巴托（蒙古）终端节点
         | kg.s3.rackcorp.com        | 比什凯克（吉尔吉斯斯坦）终端节点
         | id.s3.rackcorp.com        | 雅加达（印度尼西亚）终端节点
         | jp.s3.rackcorp.com        | 东京（日本）终端节点
         | sg.s3.rackcorp.com        | 新加坡终端节点
         | de.s3.rackcorp.com        | 法兰克福（德国）终端节点
         | us.s3.rackcorp.com        | 美国（AnyCast）终端节点
         | us-east-1.s3.rackcorp.com | 纽约（美国）终端节点
         | us-west-1.s3.rackcorp.com | 维尔蒙特（美国）终端节点
         | nz.s3.rackcorp.com        | 奥克兰（新西兰）终端节点

   --location-constraint
      地理约束 - 存储桶所在的地理位置和数据存储的位置。

      示例：
         | global    | 全球CDN区域
         | au        | 澳大利亚（所有州）
         | au-nsw    | 新南威尔士州（澳大利亚）区域
         | au-qld    | 昆士兰州（澳大利亚）区域
         | au-vic    | 维多利亚州（澳大利亚）区域
         | au-wa     | 珀斯（澳大利亚）区域
         | ph        | 马尼拉（菲律宾）区域
         | th        | 曼谷（泰国）区域
         | hk        | 香港区域
         | mn        | 乌兰巴托（蒙古）区域
         | kg        | 比什凯克（吉尔吉斯斯坦）区域
         | id        | 雅加达（印度尼西亚）区域
         | jp        | 东京（日本）区域
         | sg        | 新加坡（新加坡）区域
         | de        | 法兰克福（德国）区域
         | us        | 美国（任意播送）区域
         | us-east-1 | 纽约（美国）区域
         | us-west-1 | 维尔蒙特（美国）区域
         | nz        | 奥克兰（新西兰）区域

   --acl
      在创建存储桶、存储或复制对象时使用的预设ACL。

      此ACL用于创建对象，并且如果未设置bucket_acl，则也用于创建存储桶。

      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，S3在服务器端复制对象时会应用此ACL，而不是从源复制ACL。

      如果acl为空字符串，则不会添加X-Amz-Acl：标头，并将使用默认值（private）。

   --bucket-acl
      在创建存储桶时使用的预设ACL。

      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用“acl”。

      如果“acl”和“bucket_acl”为空字符串，则不会添加X-Amz-Acl：
      标头，并将使用默认值（private）。

      示例：
         | private            | 所有者获得FULL_CONTROL。
         |                    | 无其他人员有访问权限（默认）。
         | public-read        | 所有者获得FULL_CONTROL。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者获得FULL_CONTROL。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者获得FULL_CONTROL。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止点。

      大于此大小的任何文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。

      对于大于upload_cutoff的文件或大小未知的文件（例如，“rclone rcat”或使用“rclone mount”或google
      photos或google docs上传的文件），将使用此块大小进行分块上传。

      注意，“--s3-upload-concurrency”每个传输在内存中缓冲此块大小。

      如果您在高速链路上传输大文件，并且具有足够的内存，则增加此值将加快传输速度。

      Rclone将自动增加块大小，以确保大文件的分块数不超过10,000。

      大小未知的文件使用配置的chunk_size进行上传。由于默认块大小为5 MiB，最多可以有10,000个块，所以默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加chunk_size。

      增加块大小会降低使用“-P”标志显示的进度统计数据的准确性。rclone在将块上传到AWS SDK的缓冲区后会将其视为已发送，而实际上它可能仍在上传中。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况更为偏差。

   --max-upload-parts
      分块上传中的最大部分数。

      此选项定义在执行分块上传时使用的最大多部分块数。

      如果某个服务不支持AWS S3 10,000个多部分块的规范，则此参数可能很有用。

      Rclone将自动增加块大小，以确保大文件的分块数不超过此数。

   --copy-cutoff
      切换到分块复制的截止点。

      需要服务器端复制且大于此大小的文件将使用此大小的块进行复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和与对象元数据一起存储。

      通常情况下，rclone将在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说，可能会导致长时间的上传延迟。

   --shared-credentials-file
      共享凭据的文件路径。

      如果env_auth = true，则rclone可以使用共享凭据文件。

      如果此变量为空，则rclone将搜索“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空，则默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭据文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用的配置文件。

      如果为空，则默认为环境变量“AWS_PROFILE”或“default”（如果该环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传的相同文件的块数。

      如果您在高速链路上上传少量大文件，并且这些上传不充分利用带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟主机样式访问。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求此设置为false- rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认），则rclone将使用v4身份验证。如果设置了该值，则rclone将使用v2身份验证。

      仅在v4签名不起作用（例如，早期的Jewel/v10 CEPH）时使用。

   --list-chunk
      列举块的大小（每个ListObject S3请求的响应列表）。

      该选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求的列表超过1000个对象，也会对响应列表进行截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。

      当S3最初发布时，它仅提供了用于枚举存储桶中的对象的ListObjects调用。

      但是，在2016年5月，引入了ListObjectsV2调用。这是更高性能的，应尽可能使用。

      如果设置为默认值0，则rclone将根据设置的提供商来猜测应调用哪个列出对象方法。如果它的猜测不正确，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      一些提供商支持对列表进行URL编码，在使用控制字符时，这更可靠。如果设置为unset（默认），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。

      如果您知道桶已经存在，这可能对于尽量减少rclone事务的数量非常有用。

      如果使用的用户没有创建桶的权限，则可能也需要使用此选项。在v1.52.0之前，这将由于错误而默默传递。

   --no-head
      如果设置，则不会HEAD已上传的对象以检查完整性。

      这在尽量减少rclone事务的数量时非常有用。

      如果在上传PUT对象后收到200 OK消息，则rclone会假设它已正确上传。具体来说，它将假设：
      - 元数据，包括修改时间、存储类和内容类型与上传一致。
      - 大小与上传一致。

      它从发送一个单一部分PUT的响应中读取以下内容：
      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项目。

      如果上传一个大小未知的源对象，则rclone**会**执行HEAD请求。

      设置此标志会增加无法检测到的上传错误的机会，特别是大小不正确，因此不推荐在正常操作中使用。实际上，即使设置了此标志，检测不到的上传错误的几率非常小。

   --no-head-object
      如果设置，获取对象之前不进行HEAD操作。

   --encoding
      后端的编码。

      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将被刷新的时间间隔。

      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前s3（特别是minio）后端与HTTP/2存在一个未解决的问题。默认情况下，启用了s3后端的HTTP/2，但可以在此禁用。当问题得到解决时，此标志将被删除。
      请参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631。

   --download-url
      自定义下载的端点。
      通常将其设置为CloudFront CDN URL，因为通过CloudFront网络下载的数据可以享受使用AWS S3提供的更便宜的出口流量。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证

      这应为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用带预签名的请求或PutObject进行单一部分上传。

      如果此为false，则rclone将使用AWS SDK的PutObject来上传对象。

      rclone的版本 < 1.59 使用预签名的请求来上传单一部分的对象，并允许重新启用此功能。除了特殊情况或测试之外，通常情况下不需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示文件版本为指定时间。

      参数应为日期、“2006-01-02”、日期时间“2006-01-02
      15:04:05”或距离那时远的持续时间，例如“100d”或“1h”。

      请注意，在使用此选项时，不允许执行文件写入操作，因此无法上传文件或删除文件。

      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。

      可以使用“Content-Encoding: gzip”将对象上传到S3。通常情况下，rclone将以压缩对象的方式下载这些文件。

      如果设置了此标志，则rclone将收到“Content-Encoding: gzip”的对象接收时对其进行解压缩。这意味着rclone无法检查大小和哈希值，但是文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行gzip压缩，则设置此标志。

      通常情况下，提供商在下载对象时不会更改对象。如果对象没有使用`Content-Encoding: gzip`进行上传，则在下载时也不会设置它。

      但是，有些提供商（例如Cloudflare）即使未使用`Content-Encoding: gzip`进行上传，也可能对对象进行gzip压缩。

      这种情况下的一种症状可能是收到类似下面的错误
      ```
      ERROR corrupted on transfer: sizes differ NNN vs MMM
      ```

      如果设置了此标志，并且rclone下载带有`Content-Encoding: gzip`和分块传输编码的对象，则rclone将会实时解压缩对象。

      如果设置为unset （默认值），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  在创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value             RackCorp对象存储的终端节点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  地理约束 - 存储桶所在的地理位置和数据存储的位置。[$LOCATION_CONSTRAINT]
   --region value               region - 您的存储桶将被创建的位置以及您的数据存储的位置。[$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               在创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             自定义下载的端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列举块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将被刷新的时间间隔。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip压缩，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会HEAD已上传的对象以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，获取对象之前不进行HEAD操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据的文件路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用带预签名的请求或PutObject进行单一部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示文件版本为指定时间。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}