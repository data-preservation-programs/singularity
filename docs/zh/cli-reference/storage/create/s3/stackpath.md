# StackPath对象存储

{% code fullWidth="true" %}
```
命令名称:
   singularity storage create s3 stackpath - StackPath Object Storage

用法:
   singularity storage create s3 stackpath [命令选项] [参数...]

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时才适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或获取运行时凭据。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或获取运行时凭据。

   --region
      连接的地区。
      
      如果使用S3克隆，并且没有地区，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空地区。
         | other-v2-signature | 仅在v4签名无效时使用此选项。
         |                    | 例如，Jewel/v10之前的CEPH。

   --endpoint
      StackPath对象存储的端点。

      示例：
         | s3.us-east-2.stackpathstorage.com    | 美国东部端点
         | s3.us-west-1.stackpathstorage.com    | 美国西部端点
         | s3.eu-central-1.stackpathstorage.com | 欧洲端点

   --acl
      创建桶和存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，并且如果bucket_acl为空，则用于创建桶。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当在服务器端复制对象时，此ACL仅适用于S3，
      因为S3不会复制源中的ACL，而是写入新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl:标头，将使用默认值（私有）。

   --bucket-acl
      创建桶时使用的预定义ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建桶时应用此ACL。如果未设置，则使用“acl”。
      
      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl：
      标头，将使用默认值（私有）。

      示例：
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 通常不建议在桶上授予此权限。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      切换到分块上传的临界值。
      
      大于此大小的文件将按块大小进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如由“rclone rcat”或
      使用“rclone mount”或Google照片或Google文档上传的文件）时，将使用此块大小进行上传。
      
      请注意，“--s3-upload-concurrency”个此大小的块在每次传输中在内存中进行缓冲。
      
      如果您正在通过高速链接传输大文件，并且有足够的内存，则增加此值将加快传输速度。
      
      Rclone将根据文件大小自动增加块大小，以保持小于10,000个块的限制。
      
      未知大小的文件使用配置的
      chunk_size进行上传。由于默认的块大小为5 MiB，最多可以有
      10,000个块，这意味着默认情况下您可以流式传输上传的文件的最大大小为48 GiB。
      如果要流式传输更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用“-P”选项时显示的进度
      统计数据的准确性。当AWS SDK将块作为已发送时，rclone将块视为已发送，
      而事实上可能仍在上传。更大的块大小意味着更大的AWS SDK缓冲区和进度
      报告与实际值偏离更多。
      

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义在执行多部分上传时要使用的最大多部分块数。
      
      如果某个服务不支持AWS S3的规范中的10,000个块，则这很有用。
      
      Rclone将根据文件大小自动增加块大小，以保持这个块数的限制。
      

   --copy-cutoff
      切换到分块复制的临界值。
      
      需要进行服务器端复制的大于此大小的文件将按该大小进行复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，
      但对于大文件而言可能会导致长时间的延迟才能开始上传。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth为true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找
      “AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空
      它将默认为当前用户的主目录。
      
          Linux/OSX: “$HOME/.aws/credentials”
          Windows:   “%USERPROFILE%\.aws\credentials”
      

   --profile
      在共享凭据文件中使用的配置文件。
      
      如果env_auth为true，则rclone可以使用共享凭据文件。此
      变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”或
      如果该环境变量也未设置，则默认为“default”。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您正在通过高速链接上传少量大文件，并且这些上传没有充分利用带宽，
      那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见[AWS S3
      文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      
      某些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求将此设置为
      false - rclone将根据提供者的设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了它，则rclone将使用v2身份验证。
      
      仅在v4签名无效时使用。

   --list-chunk
      列举块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务将响应列表截断为最多1000个对象，即使请求的多于这个数量。
      在AWS S3中，这是全局最大值，并且无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过“rgw list buckets max chunk”选项来增加此值。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。
      
      当S3最初推出时，它只提供了ListObjects调用以
      枚举存储桶中的对象。
      
      但是在2016年5月引入了ListObjectsV2调用。这是
      性能非常高的，应该尽可能使用。
      
      如果设置为默认值0，则rclone将根据提供者设置猜测要调用的
      List objects方法。如果它的猜测错误，则可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，并且可以在文件
      名中使用控制字符时，这更加可靠。如果设置为unset
      （默认值），则rclone将根据提供者设置选择要应用的内容，
      但是您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，不要尝试检查桶是否存在或创建桶。
      
      如果您知道桶已经存在，可以使用此功能来尝试最小化rclone执行的事务数。
      
      如果使用的用户没有桶创建权限，也可能需要此功能。在v1.52.0之前，由于错误，
      这将会静默传递。
      

   --no-head
      如果设置，不要对已上传对象进行HEAD验证。
      
      这在尝试尽量减少rclone执行的事务数时非常有用。
      
      设置它意味着，如果rclone在使用PUT上传对象后接收到200 OK消息，
      它将假设它已正确上传。
      
      特别是，它将假设：
      
      - 元数据（包括修改时间、存储类和内容类型）与上传时一样
      - 大小与上传时一样
      
      对于单分块PUT，它从响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传长度未知的源对象，则rclone**将**执行
      HEAD请求。
      
      设置此标志会增加检测不到上传故障的机会，
      特别是错误的大小，因此不建议正常操作使用此标志。实际上，
      即使使用此标志，检测不到上传故障的机会也很小。
      

   --no-head-object
      如果设置，不要在GET获取对象之前执行HEAD。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区（例如分块）的上传将使用内存池进行分配。
      此选项控制多久将未使用的缓冲区从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个未解决的问题
      和HTTP/2。对于s3后端，默认情况下启用了HTTP/2，但是可以
      在此禁用HTTP/2。当问题解决时，此标志将被删除。
      
      请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义端点。
      这通常设置为CloudFront CDN URL，因为AWS S3提供了
      通过CloudFront网络下载数据的更便宜的出口。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      此值应设置为true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject来上传
      对象。
      
      版本小于1.59的rclone使用预签名请求来上传单个
      部分对象，将此标志设置为true将重新启用该
      功能。除非特殊情况或测试，否则不应该需要这样做。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应为日期，“2006-01-02”，日期时间“2006-01-02
      15:04:05”或距今的持续时间，例如“100d”或“1h”。
      
      请注意，当使用此选项时，不允许进行文件写操作，
      因此无法上传文件或删除文件。
      
      请参见[时间选项文档](/docs/#time-option)以获取有效格式。
      

   --decompress
      如果设置，这将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”将对象上传到S3。通常，rclone
      将以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则rclone将在接收到“Content-Encoding: gzip”的文件时进行解压缩。
      这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此值。
      
      通常，提供者在下载对象时不会更改对象。如果
      对象没有使用“Content-Encoding: gzip”上传，则下载时它不会
      集。然而，即使没有使用“Content-Encoding: gzip”上传，
      某些提供商也可能对对象进行gzip压缩（例如Cloudflare）。
      
      这种情况的症状将是收到类似的错误
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone使用
      Content-Encoding: gzip 设置的、以块式传输编码的对象时，
      rclone将实时解压缩该对象。
      
      如果将其设置为unset（默认值），则rclone将选择
      根据提供者的设置选择要应用的内容，但您可以在此处覆盖
      rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id 值            AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl 值                      创建桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint 值                 StackPath对象存储的端点。[$ENDPOINT]
   --env-auth                    从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。(默认值: false) [$ENV_AUTH]
   --help, -h                    显示帮助信息
   --region 值                   连接的地区。[$REGION]
   --secret-access-key 值        AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl 值                创建桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size 值                用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff 值               切换到分块复制的临界值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                   如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum             不要将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url 值              下载的自定义端点。[$DOWNLOAD_URL]
   --encoding 值                  后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style             如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk 值                列举块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode 值           是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version 值              要使用的ListObjects的版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts 值          多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time 值    内部内存缓冲池将刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap         是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip 值                如果后端可能对对象进行gzip压缩，请设置此值。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket              如果设置，不要尝试检查桶是否存在或创建桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                      如果设置，不要对已上传对象进行HEAD验证。 (默认值: false) [$NO_HEAD]
   --no-head-object               如果设置，不要在GET获取对象之前执行HEAD。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata           禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile 值                   在共享凭据文件中使用的配置文件。 [$PROFILE]
   --session-token 值             AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file 值   共享凭据文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency 值        多部分上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 值             切换到分块上传的临界值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag 值        是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request        是否使用预签名请求或PutObject进行单个部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                      如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at 值                显示指定时间的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                     在目录列表中包括旧版本。 (默认值: false) [$VERSIONS]

   一般

   --name 值  存储的名称（默认值：自动生成）
   --path 值  存储的路径

```
{% endcode %}