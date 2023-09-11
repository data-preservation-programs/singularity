# DigitalOcean Spaces

{% code fullWidth="true" %}
```
命令：
   singularity storage create s3 digitalocean - DigitalOcean Spaces

用法：
   singularity storage create s3 digitalocean [命令选项] [参数]

描述：
   --env-auth
      从运行环境获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时适用。

      例子：
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境中获取AWS凭据（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或使用运行时凭据。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或使用运行时凭据。

   --region
      连接的区域。
      
      如果使用的是S3克隆并且没有特定的区域，请留空。

      例子：
         | <未设置>                  | 如果不确定，请使用此选项。
         |                          | 将使用v4签名和空区域。
         | other-v2-signature       | 仅当v4签名不起作用时使用。
         |                          | 例如，pre Jewel/v10 CEPH。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时必填。

      例子：
         | syd1.digitaloceanspaces.com | DigitalOcean Spaces Sydney 1
         | sfo3.digitaloceanspaces.com | DigitalOcean Spaces San Francisco 3
         | fra1.digitaloceanspaces.com | DigitalOcean Spaces Frankfurt 1
         | nyc3.digitaloceanspaces.com | DigitalOcean Spaces New York 3
         | ams3.digitaloceanspaces.com | DigitalOcean Spaces Amsterdam 3
         | sgp1.digitaloceanspaces.com | DigitalOcean Spaces Singapore 1

   --location-constraint
      区域约束，必须与区域匹配。
      
      如果不确定，请留空。仅在创建桶时使用。

   --acl
      创建桶、存储或复制对象时使用的预设ACL。
      
      此ACL用于创建对象，并且如果桶的acl未设置，也用户创建桶。
      
      获取更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，S3不会复制源的ACL，而是写入一个新的ACL。
      
      如果ACL是一个空字符串，则不会添加X-Amz-Acl头，并且将使用默认（私有）。

   --bucket-acl
      创建桶时使用的预设ACL。
      
      获取更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅当创建桶时才会应用此ACL。如果未设置该属性，则使用“acl”。
      
      如果“acl”和“bucket_acl”均为空字符串，则不会添加X-Amz-Acl头，
      且将使用默认（私有）。

      例子：
         | private            | 所有者获得FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者获得FULL_CONTROL权限。
         |                    | AllUSers组拥有读取权限。
         | public-read-write  | 所有者获得FULL_CONTROL权限。
         |                    | AllUSers组拥有读取和写入权限。
         |                    | 通常不推荐在桶上进行此设置。
         | authenticated-read | 所有者获得FULL_CONTROL权限。
         |                    | AuthenticatedUsers组拥有读取权限。

   --upload-cutoff
      切换到分块上传的大小。
      
      任何大于此大小的文件都将以chunk_size的大小进行分块上传。
      最小为0，最大为5 GB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如，使用“rclone rcat”或使用“rclone mount”或Google照片或Google文档上传的文件），
      将使用此块大小进行分块上传。
      
      请注意，每个传输都会在内存中缓冲大小为“--s3-upload-concurrency”的块。
      
      如果您正在通过高速连接传输大型文件且内存足够，增加此值将加快传输速度。
      
      Rclone将在传输大型已知大小的文件时自动增加块大小，以确保不超过10000块的限制。
      
      未知大小的文件将使用配置的chunk_size进行上传。由于默认的块大小为5 MiB，并且最多可以有10000个块，
      这意味着默认情况下您可以流式上传的文件的最大大小为48 GB。如果您希望流式上传更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用“-P”标记时显示的进度统计结果的准确性。
      Rclone将块视为已发送，当AWS SDK将其缓冲时，实际上可能仍在上传中。
      更大的块大小意味着更大的AWS SDK缓冲区和进度报告与实际情况更不一致。

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义进行多部分上传时要使用的最大多部分块数。
      
      如果一个服务不支持AWS S3定义的10000个块的规范，
      则这样做可能很有用。
      
      Rclone将在传输已知大小的大型文件时自动增加块大小，以保持低于此块数的限制。

   --copy-cutoff
      切换到分块复制的大小。
      
      需要服务器端复制的任何大于此大小的文件都将以此大小进行分块复制。
      
      最小为0，最大为5 GB。

   --disable-checksum
      在对象元数据中不存储MD5校验和。
      
      通常，在上传之前，rclone会计算输入内容的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，
      但可能会导致大文件在开始上传时出现长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将寻找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量为空，
      则它将默认为当前用户的主目录。
      
          Linux / OSX:“$ HOME / .aws / credentials”
          Windows：“％USERPROFILE％\。aws \ credentials”

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。该变量控制在该文件中使用哪个配置文件。
      
      如果为空，则将默认为环境变量“AWS_PROFILE”或“default”如果该环境变量也未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传相同文件的块数。
      
      如果您通过高速连接上传少量大型文件，并且这些上传没有充分利用带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟路径样式。
      有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了该值，则rclone将使用v2身份验证。
      
      仅当v4签名不可用时使用，例如，pre Jewel/v10 CEPH。

   --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范的“MaxKeys”，“max-items”或“page-size”。
      大多数服务都限制响应列表为1000个对象，即使请求更多。
      在AWS S3中，这是一个全局最大值，不能更改，请参阅[AWS S3文档](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0为自动。
      
      当S3最初发布时，它只提供了用于枚举存储桶中对象的ListObjects调用。
      
      但是，在2016年5月引入了ListObjectsV2调用。这是性能更高的调用，如果可能的话应使用。
      
      如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪个列出对象方法。如果它猜错了，
      则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，则使用控制字符在文件名中更可靠。如果设置为unset（默认值），
      则rclone将根据提供商设置选择要应用的URL编码。
      
   --no-check-bucket
      如果设置，则不要尝试检查存储桶是否存在或创建它。
      
      如果您知道桶已经存在，则可以使用此功能来尽量减少rclone执行的事务数。
      
      如果使用的用户没有创建桶的权限，则可能需要这样做。在v1.52.0之前，由于缺陷，这会导致静默传递。
      
   --no-head
      如果设置，则不会在GET获取对象之前进行HEAD检查以检查完整性。
      
      这在尝试最小化rclone执行的事务数时很有用。
      
      设置它意味着如果rclone在使用PUT上传对象后收到200 OK消息，它将假定对象已正确上传。
      
      特别是它将假定：
      
      - 元数据（包括modtime、存储类和内容类型）与上传时的相同
      - 大小与上传时的相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传位置未知的源对象，那么rclone将**会**做一个HEAD请求。
      
      设置此标志会增加未检测到的上传失败的几率，特别是大小不正确的几率，
      因此不推荐在正常操作中使用。实际上，即使启用此标志，未检测到的上传失败的机会也非常小。
      
   --no-head-object
      如果设置，则在获取对象时不进行HEAD操作。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      要求额外缓冲区（例如分块上传）的上传将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个未解决的问题和HTTP/2的问题。
      HTTP/2默认启用s3后端，但可以在此禁用。当问题解决时，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631。

   --download-url
      下载的自定义终端。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据
      提供更低成本的出站流量。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这里可以设置为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      Rclone版本<1.59使用预签名请求上传单个部分对象，
      将此标志设置为true将重新启用此功能。除非在特殊情况下或用于测试，否则不应该使用。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应为日期“2006-01-02”，日期时间“2006-01-02 15:04:05”或相对之前的持续时间，例如“100d”或“1h”。
      
      请注意，使用此选项时不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      请参阅[时间选项文档](/docs/#time-option)以获取有效格式。

   --decompress
      如果设置，这将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”设置将对象上传到S3。通常rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在收到带有“Content-Encoding: gzip”的文件时进行解压缩。这意味着rclone
      无法检查大小和哈希值，但是文件内容将会被解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此项。
      
      通常提供者在下载对象时不会更改对象。如果一个对象在上传时没有通过“Content-Encoding: gzip”进行上传，
      则在下载时也不会设置它。
      
      但是，某些提供者甚至可能对对象进行gzip压缩，即使它们没有通过“Content-Encoding: gzip”进行上传（例如Cloudflare）。
      
      这会导致接收到错误，例如：
      
      ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置此标志，并且rclone下载了具有设置了Content-Encoding: gzip和分块传输编码的对象，
      则rclone将在传输过程中动态解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，
      但是您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  创建桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value             S3 API的端点。[$ENDPOINT]
   --env-auth                   从运行环境获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  区域约束，必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region value               连接的区域。[$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的大小。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，这将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               在对象元数据中不存储MD5校验和。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0为自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip压缩，请设置此项。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不要尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会在GET获取对象之前进行HEAD检查以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不进行HEAD操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的大小。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name value  存储名称（默认值: 自动生成）
   --path value  存储路径

```
{% endcode %}