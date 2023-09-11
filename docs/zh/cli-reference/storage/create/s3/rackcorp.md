# RackCorp 对象存储

{% code fullWidth="true" %}
```
名称：
   singularity storage create s3 rackcorp - RackCorp对象存储

用法：
   singularity storage create s3 rackcorp [command options] [arguments...]

说明：
   --env-auth
      从运行时获取AWS凭证（从环境变量或EC2/ECS元数据获取）。
      
      仅在access_key_id和secret_access_key为空时生效。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --region
      区域-您的存储桶将被创建和存储数据的位置。
      

      示例：
         | global    | 全球CDN（所有地点）区域
         | au        | 澳大利亚（所有省份）
         | au-nsw    | 新南威尔士（澳大利亚）区域
         | au-qld    | 昆士兰（澳大利亚）区域
         | au-vic    | 维多利亚（澳大利亚）区域
         | au-wa     | 珀斯（澳大利亚）区域
         | ph        | 马尼拉（菲律宾）区域
         | th        | 曼谷（泰国）区域
         | hk        | 香港区域
         | mn        | 乌兰巴托（蒙古）区域
         | kg        | 比什凯克（吉尔吉斯斯坦）区域
         | id        | 雅加达（印度尼西亚）区域
         | jp        | 东京（日本）区域
         | sg        | 新加坡区域
         | de        | 法兰克福（德国）区域
         | us        | 美国（AnyCast）区域
         | us-east-1 | 纽约（美国）区域
         | us-west-1 | Freemont (USA) Region
         | nz        | 奥克兰（新西兰）区域

   --endpoint
      RackCorp对象存储的端点。

      示例：
         | s3.rackcorp.com           | 全球（AnyCast）端点
         | au.s3.rackcorp.com        | 澳大利亚（Anycast）端点
         | au-nsw.s3.rackcorp.com    | 悉尼（澳大利亚）端点
         | au-qld.s3.rackcorp.com    | 布里斯班（澳大利亚）端点
         | au-vic.s3.rackcorp.com    | 墨尔本（澳大利亚）端点
         | au-wa.s3.rackcorp.com     | 珀斯（澳大利亚）端点
         | ph.s3.rackcorp.com        | 马尼拉（菲律宾）端点
         | th.s3.rackcorp.com        | 曼谷（泰国）端点
         | hk.s3.rackcorp.com        | 香港端点
         | mn.s3.rackcorp.com        | 乌兰巴托（蒙古）端点
         | kg.s3.rackcorp.com        | 比什凯克（吉尔吉斯斯坦）端点
         | id.s3.rackcorp.com        | 雅加达（印度尼西亚）端点
         | jp.s3.rackcorp.com        | 东京（日本）端点
         | sg.s3.rackcorp.com        | 新加坡端点
         | de.s3.rackcorp.com        | 法兰克福（德国）端点
         | us.s3.rackcorp.com        | 美国（AnyCast）端点
         | us-east-1.s3.rackcorp.com | 纽约（美国）端点
         | us-west-1.s3.rackcorp.com | Freemont (USA) Endpoint
         | nz.s3.rackcorp.com        | 奥克兰（新西兰）端点

   --location-constraint
      位置约束-您的存储桶将被创建和存储数据的位置。
      

      示例：
         | global    | 全球CDN区域
         | au        | 澳大利亚（所有位置）
         | au-nsw    | 新南威尔士（澳大利亚）区域
         | au-qld    | 昆士兰（澳大利亚）区域
         | au-vic    | 维多利亚（澳大利亚）区域
         | au-wa     | 珀斯（澳大利亚）区域
         | ph        | 马尼拉（菲律宾）区域
         | th        | 曼谷（泰国）区域
         | hk        | 香港区域
         | mn        | 乌兰巴托（蒙古）区域
         | kg        | 比什凯克（吉尔吉斯斯坦）区域
         | id        | 雅加达（印度尼西亚）区域
         | jp        | 东京（日本）区域
         | sg        | 新加坡区域
         | de        | 法兰克福（德国）区域
         | us        | 美国（AnyCast）区域
         | us-east-1 | 纽约（美国）区域
         | us-west-1 | Freemont (USA) Region
         | nz        | 奥克兰（新西兰）区域

   --acl
      在创建桶、存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，如果未设置bucket_acl，则还用于创建桶。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，当服务器端复制对象时，S3不会复制源时的ACL，而是写入一个新的ACL。
      
      如果ACL是一个空字符串，则不添加X-Amz-Acl:标题，并使用默认值（私有）。

   --bucket-acl
      创建存储桶时使用的默认ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，仅在创建存储桶时应用此ACL。如果未设置它，则使用"acl"而不是"bucket_acl"。
      
      如果"acl"和"bucket_acl"都是空字符串，则不添加X-Amz-Acl:
      标题，并使用默认值（私有）。

      示例：
         | private            | 所有者获得FULL_CONTROL权限。
         |                    | 没有其他人拥有访问权限（默认）。
         | public-read        | 所有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 不建议在存储桶上进行此操作。
         | authenticated-read | 所有者获得FULL_CONTROL权限。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止值。
      
      大于此文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或具有未知大小（例如来自"rclone rcat"或使用"rclone mount"或Google照片或Google文档上传的文件）时，它们将使用此块大小作为分块上传。
      
      注意，每次传输在内存中缓冲此大小的"--s3-upload-concurrency"个块。
      
      如果您正在通过高速链路传输大文件并且内存足够，则增加此值将加快传输速度。
      
      在上传已知大小的大文件时，rclone将自动增加块大小以保持在10000个块的限制以下。
      
      未知大小的文件将使用配置的chunk_size进行上传。由于默认块大小为5 MiB，最多可以有10000个块，这意味着默认情况下您可以流式传输的文件的最大大小为48 GiB。如果要流式传输大文件，则需要增加chunk_size。
      
      增加块大小会降低使用"-P"标志显示的进度统计的准确性。当块被AWS SDK缓冲时，rclone将其视为已发送的块，事实上它可能仍在上传。较大的块大小意味着较大的AWS SDK缓冲区和更不准确的进度报告。

   --max-upload-parts
      分块上传中的最大块数。
      
      当执行分块上传时，此选项定义要使用的分块数的最大值。
      
      如果服务不支持AWS S3规范的10000个块，此选项可能很有用。
      
      当上传已知大小的大文件时，rclone将自动增加块大小以保持在此块数限制以下。

   --copy-cutoff
      切换到分块复制的截止值。
      
      需要服务器端复制的大于此大小的文件将被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于开始上传大文件可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"如果该环境变量也未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      同一文件的这些块并发上传。
      
      如果您通过高速链路上传少量大文件，并且这些上传没有充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果设置为true，则使用路径样式访问，如果设置为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为
      false- rclone将根据提供程序制定的设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，则rclone将使用v2身份验证。
      
      仅在v4签名无效时使用，例如pre Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求超过了此数量，也会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，不能更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      当S3最初发布时，它只提供了用于枚举存储桶中的对象的ListObjects调用。
      
      但是，在2016年5月引入了ListObjectsV2调用。这是更高性能的调用，如果可能的话，应该使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供程序猜测要调用哪个列表对象方法。如果猜测错误，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持对列表进行URL编码，如果可用，则在使用控制字符时这更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但是您可以在此处重写rclone的选择。

   --no-check-bucket
      如果设置，不要尝试检查存储桶是否存在或创建它。
      
      如果您已经知道存储桶已经存在，这可能对于尽量减少rclone执行的事务数量很有用。
      
      如果使用的用户不具有创建存储桶权限，则也可能需要此操作。在v1.52.0之前，由于错误，此操作将无声传递。

   --no-head
      如果设置，不要对已上传的对象进行HEAD检查以检查完整性。
      
      这对于尽量减少rclone执行的事务数量可能很有用。
      
      设置后，这意味着如果rclone在使用PUT上传对象后接收到200 OK消息，它将假定它已正确上传。
      
      特别地，它将假定：
      
      - 上传之前的元数据，包括修改时间、存储类和内容类型都与上传一样
      - 大小与上传相同
      
      它从单个部分PUT的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于分块上传，不会读取这些项。
      
      如果上传的源对象大小未知，则rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传故障的机会，特别是大小不正确的故障，因此不建议在正常操作中使用。实际上，即使使用此标志，产生未检测到的上传失败的几率也非常小。

   --no-head-object
      如果设置，不要在获取对象时在GET之前执行HEAD。

   --encoding
      后端的编码。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区（例如分块）的上传将使用内存池进行分配。
      此选项控制在何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在一个未解决的与HTTP/2相关的问题。S3后端默认启用HTTP/2，但可以在此处禁用。解决这个问题后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      通常，它被设置为CloudFront CDN URL，因为AWS S3为通过CloudFront网络下载的数据提供更便宜的出口流量。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应该是true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果该标志为false，则rclone将使用AWS SDK中的PutObject上传对象。
      
      rclone的版本 < 1.59使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或进行测试，否则不应该需要此选项。

   --versions
      目录列表中是否包含旧版本。

   --version-at
      显示指定时间点的文件版本。
      
      参数应为日期，“2006-01-02”、日期时间“2006-01-02
      15:04:05”或距离那个时间点的持续时间，例如“100d”或“1h”。
      
      请注意，在使用此选项时，不允许进行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"在S3上上传对象。通常，rclone会下载这些文件作为压缩对象。
      
      如果设置了此标志，则rclone将在接收到"Content-Encoding: gzip"的文件时进行解压缩。这意味着rclone不能检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能gzip对象，请设置此标志。
      
      通常，提供者不会在下载时更改对象。如果未使用“Content-Encoding: gzip”上传对象，则不会在下载时设置。
      
      但是，某些提供者可能会gzip对象，即使它们不是使用“Content-Encoding: gzip”上传的（例如Cloudflare）。
      
      这种情况的症状是接收到类似错误的结果
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了带有设置了Content-Encoding: gzip和块传输编码的对象，则rclone将在传输过程中实时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但您可以在此处重写rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  在创建桶、存储或复制对象时使用的预定义ACL。[$ACL]
   --debug                      联机调试模式。
   --endpoint value             RackCorp对象存储的端点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（从环境变量或EC2/ECS元数据获取）。 (default: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束-您的存储桶将被创建和存储数据的位置。[$LOCATION_CONSTRAINT]
   --log-level value            日志级别 (default: "INFO")
   --path value                 存储路径
   --region value               区域-您的存储桶将被创建和存储数据的位置。[$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的默认ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果true，则使用路径样式访问；如果false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大块数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，请设置此标志。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，不要尝试检查存储桶是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，不要对已上传的对象进行HEAD检查以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，不要在获取对象时在GET之前执行HEAD。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果true，则使用v2身份验证。 (default: false) [$V2_AUTH]
   --version-at value               显示指定时间点的文件版本。 (default: "off") [$VERSION_AT]
   --versions                       目录列表中是否包含旧版本。 (default: false) [$VERSIONS]

   通用

   --name value  存储的名称（默认为自动生成）
   --path value  存储的路径

```
{% endcode %}