# Dreamhost DreamObjects

{% code fullWidth="true" %}
```
名称:
   singularity storage create s3 dreamhost - Dreamhost DreamObjects

用法:
   singularity storage create s3 dreamhost [命令选项] [参数...]

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。
      
      仅适用于access_key_id和secret_access_key为空的情况。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空表示匿名访问或运行时凭证。

   --region
      要连接的区域。
      
      如果您使用的是S3的克隆版本并且没有区域，则保持空白。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如， 早期的 Jewel/v10 CEPH。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要此选项。

      示例:
         | objects-us-east-1.dream.io | Dream Objects 端点

   --location-constraint
      位置约束 - 必须与区域相匹配。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      存储桶创建、存储或复制对象时使用的预定义权限。
      
      此权限用于创建对象，并且如果未设置bucket_acl，则也用于创建存储桶。
      
      有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，在服务器端复制对象时，S3不会复制源的ACL，而是会写入一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预定义权限。
      
      有关详细信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。如果未设置bucket_acl，则使用“acl”。
      
      如果“acl”和“bucket_acl”是空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认值（private）。

      示例:
         | private            | 所有者拥有FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      转换为分块上传的截止值。
      
      大于此值的文件将以chunk_size的块进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传的文件大于upload_cutoff或大小未知的文件（例如，通过"rclone rcat"上传或使用"rclone mount"或google
      照片或谷歌文档上传的文件）时，将使用这个块大小进行分块上传。
      
      注意，每个传输下的"--s3-upload-concurrency"个大小为chunk_size的块会在内存中进行缓冲。
      
      如果您通过高速链接传输大文件并且具有足够的内存，则可以增加此值以加快传输速度。
      
      当上传已知大小的大文件以保持在10,000个块限制之下时，rclone会自动增加块大小。
      
      未知大小的文件使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，
      并且最多可以有10,000个块，因此默认情况下流式上传的文件的最大大小为48 GiB。 
      如果您希望流式上传更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用"-P"标志显示的进度统计的准确性。 
      当rclone使用AWS SDK缓冲了块时，rclone将块视为已发送，而实际上它可能仍在上传。
      更大的块大小意味着更大的AWS SDK缓冲区和与实际情况偏离更大的进度报告。

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义进行多部分上传时要使用的最大多部分块数。
      
      如果某个服务不支持AWS S3规范的10,000个块，则可以使用此选项。
      
      当上传已知大小的大文件以保持在此块数限制之下时，rclone会自动增加块大小。

   --copy-cutoff
      转换为分块复制的截止值。
      
      需要进行服务器端复制的大于此值的文件将以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，在上传之前，rclone会计算输入的MD5校验和，以便将其添加到对象的元数据中。 
      这对于数据完整性检查很有用，但可能会导致大文件的长时间延迟才能开始上传。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。
      如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果未设置该环境变量）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您使用高速连接上传少量大文件，并且这些上传未完全利用您的带宽，
      则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式访问。请参阅[AWS S3
      文档](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      了解更多信息。
      
      一些服务提供商（例如AWS，Aliyun OSS，Netease COS或Tencent COS）要求将其设置为
      false - rclone将根据提供商的设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2认证。
      
      如果为false（默认值），则rclone将使用v4认证。如果设置了它，则rclone将使用v2认证。
      
      仅当v4签名不起作用时使用此选项，例如早期的Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（用于每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。
      大多数服务在超过1000个对象时都会截断响应列表，即使请求超过该数量。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过"rgw list buckets max chunk"选项来增加此值。

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。
      
      当S3最初发布时，它仅提供了用于枚举存储桶中的对象的ListObjects调用。
      
      但是，从2016年5月开始，引入了ListObjectsV2调用。它的性能要高得多，如果可能的话应该使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供商来猜测应该调用哪个列表对象方法。
      如果它猜错了，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行url编码：true/false/unset
      
      某些提供商支持URL编码列表，如果可用，则使用控制字符在文件名中的URL编码要更可靠。
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的编码方式，
      但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，这可能对于尽量减少rclone操作的数量很有用。
      
      如果所使用的用户没有创建存储桶的权限，则也可能需要这样做。在v1.52.0之前，由于Bug的缘故，此操作不会报错。

   --no-head
      如果设置，则不对上传的对象进行HEAD检查以检查完整性。
      
      这可能有助于尽量减少rclone进行的事务数量。
      
      如果在PUT上传对象后rclone接收到200 OK消息，那么rclone将假设对象已正确上传。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传时相同
      - 大小与上传时相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传未知长度的源对象，则rclone **将**执行HEAD请求。
      
      设置此标志将增加未检测到的上传故障的机会，特别是错误的大小，因此不建议正常操作使用。
      实际上，即使设置了此标志，检测到未检测到的上传故障的机会也非常小。

   --no-head-object
      如果设置，将不在获取对象之前执行HEAD。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。
      
      需要额外缓冲区（例如，多部分）的上传会使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      s3（具体来说是minio）后端目前存在HTTP/2的一个未解决问题。默认情况下，s3后端启用HTTP/2，
      但是可以在这里禁用。当问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      这通常设置为CloudFront CDN URL，因为通过CloudFront网络下载的数据可以享受AWS S3提供的更便宜的出站流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证。
      
      这应该设置为true、false或保持unset以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject上传单个部分对象。
      
      如果为false，则rclone将使用AWS SDK的PutObject来上传对象。
      
      rclone的版本 <1.59使用预签名请求来上传单个部分对象，
      将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应使用此选项。

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示文件版本为指定时间。
      
      参数应为日期、"2006-01-02"、时间日期"2006-01-02 15:04:05"或距离现在多久，
      如"100d"或"1h"。
      
      请注意，在使用此选项时，不允许执行任何文件写操作，
      因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"将对象上传到S3。
      通常情况下，rclone会将这些文件作为压缩的对象下载。
      
      如果设置了此标志，则rclone会在接收对象时对这些具有
      "Content-Encoding: gzip"的文件进行解压缩。这意味着rclone
      无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象使用gzip进行压缩，则设置此标志。
      
      通常，提供商在下载对象时不会更改对象。
      如果一个对象未使用“Content-Encoding: gzip”进行上传，
      则在下载时也不会设置它。
      
      但是，某些提供商可能对对象使用gzip进行压缩，即使它们未使用
      “Content-Encoding: gzip”上传（例如Cloudflare）。
      
      如果设置了此标志并且rclone下载了具有设置为
      “Content-Encoding: gzip”的对象和分块传输编码的对象，
      则rclone将即时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的值，
      但是您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id 值        AWS访问秘钥ID。[$ACCESS_KEY_ID]
   --acl 值                  创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint 值             S3 API的端点。[$ENDPOINT]
   --env-auth               从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --location-constraint 值  位置约束-必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region 值               要连接的区域。[$REGION]
   --secret-access-key 值    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl 值               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size 值               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff 值              转换为分块复制的截止值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                   如果设置此标志，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum             不要将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url 值             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding 值                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style             如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk 值               列表块的大小（用于每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode 值          是否对列表进行url编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version 值             要使用的ListObjects的版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts 值         多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time 值   内部内存缓冲池刷新的时间间隔。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap         是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip 值               如果后端可能对对象使用gzip进行压缩，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket              如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                      如果设置，则不对上传的对象进行HEAD检查以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object               如果设置，将不在获取对象之前执行HEAD。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata           禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile 值                  共享凭证文件中使用的配置文件。 [$PROFILE]
   --session-token 值            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file 值  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency 值       多部分上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 值            转换为分块上传的截止值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag 值       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request        是否使用预签名请求或PutObject上传单个部分对象。 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                      如果为true，则使用v2认证。 (默认值: false) [$V2_AUTH]
   --version-at 值               显示文件版本为指定时间。 (默认值: "off") [$VERSION_AT]
   --versions                     在目录列表中包括旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name 值  存储的名称（默认值: 自动生成的）
   --path 值  存储的路径

```
{% endcode %}