# Ceph对象存储

{% code fullWidth="true" %}
```
名称:
   singularity storage update s3 ceph - Ceph对象存储

用法:
   singularity storage update s3 ceph [命令选项] <名称|ID>

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时有效。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      如果要进行匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      如果要进行匿名访问或使用运行时凭证，请留空。

   --region
      要连接的区域。
      
      如果您使用的是S3克隆，并且没有区域，请留空。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名无法正常工作时使用。
         |                    | 例如，早期的Jewel/v10 CEPH。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要。

   --location-constraint
      位置约束 - 必须设置与区域匹配。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶、存储或复制对象时使用的默认ACL。
      
      此ACL用于创建对象，并且如果未设置bucket_acl也用于创建存储桶。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，在服务器端复制对象时，S3不会复制源对象的ACL，而是写入一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl:头，将使用默认（私有）。

   --bucket-acl
      创建存储桶时使用的默认ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。如果未设置，则改为使用“acl”。
      
      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl:
      头，默认情况下使用默认（私有）。

      示例:
         | private            | 所有者有完全控制权。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者有完全控制权。
         |                    | AllUsers组具有读取访问权限。
         | public-read-write  | 所有者有完全控制权。
         |                    | AllUsers组具有读取和写入访问权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者有完全控制权。
         |                    | AuthenticatedUsers组具有读取访问权限。

   --server-side-encryption
      存储对象在S3中使用的服务器端加密算法。

      示例:
         | <unset> | 无
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用SSE-C，在存储对象时使用的服务器端加密算法。

      示例:
         | <unset> | 无
         | AES256  | AES256

   --sse-kms-key-id
      如果使用KMS ID，则必须提供Key的ARN。

      示例:
         | <unset>                 | 无
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      若要使用SSE-C，您可以提供用于加密/解密数据的秘密加密密钥。
      
      或者，您可以提供--sse-customer-key-base64。

      示例:
         | <unset> | 无

   --sse-customer-key-base64
      如果使用SSE-C，则必须提供以base64格式编码的秘密加密密钥，用于加密/解密数据。
      
      或者，您可以提供--sse-customer-key。

      示例:
         | <unset> | 无

   --sse-customer-key-md5
      如果使用SSE-C，可以提供秘密加密密钥的MD5校验和（可选）。
      
      如果将其留空，则此校验和将从提供的sse_customer_key自动计算。

      示例:
         | <unset> | 无

   --upload-cutoff
      切换到分块上传的截止点。
      
      大于此大小的任何文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如使用“rclone rcat”或使用“rclone mount”或Google
      相册或Google文档上传的文件），将使用此块大小进行分块上传。
      
      请注意，每次传输都会缓冲“--s3-upload-concurrency”个此大小的块。
      
      如果您正在使用高速链接传输大文件，并且内存足够，增加此值将加快传输速度。
      
      Rclone在上传已知大小的大文件时，会自动增加块大小，以保持在10,000个块的限制之下。
      
      未知大小的文件使用配置的
      chunk_size 进行上传。由于默认的块大小为5 MiB，并且最多可以有10,000个块，
      这意味着默认情况下可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低使用“-P”标志显示的进度统计的准确性。当
      由于设置在AWS SDK中缓冲了块而未发送时，Rclone将块视为已发送，
      实际上可以仍在上传。较大的块大小意味着较大的AWS SDK缓冲区和进度
      报告与实际情况更有偏差。
      

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义进行多部分上传时要使用的最大多部分块数。
      
      如果有服务不支持10,000个块的AWS S3规范，则可能很有用。
      
      当上传已知大小的大文件时，Rclone会自动增加块大小，以保持在该块数限制之下。
      

   --copy-cutoff
      切换到分块拷贝的截止点。
      
      需要进行服务器端拷贝的大于此大小的任何文件将被分块拷贝。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和与对象元数据一起存储。
      
      通常在上传之前，rclone会计算输入文件的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说会导致长时间的延迟，以便开始上传。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境值为空
      它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\\.aws\\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则将默认为环境变量"AWS_PROFILE"或
      如果该环境变量也未设置，则默认为"default"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      同一文件的这些块的数量将同时上传。
      
      如果您通过高速链接上传少量大文件，并且这些上传未充分利用带宽，则增加此值可能会帮助加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见 [AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      。
      
      一些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为
      false - rclone将根据提供商设置自动完成此操作。

   --v2-auth
      如果为true，则使用v2认证。
      
      如果为false（默认值），则rclone将使用v4认证。如果设置了它，则rclone将使用v2认证。
      
      仅在v4签名无法正常工作时使用，例如，早期的Jewel/v10 CEPH。

   --list-chunk
      要列出的大小（每个ListObject S3请求的响应列表大小）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求超过1000个对象，也会将响应列表截断为1000个。
      在AWS S3中，这是一个全局最大值，不能更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加这个值。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动选择。
      
      当S3最初推出时，它只提供了ListObjects调用来枚举存储桶中的对象。
      
      然而，在2016年5月，引入了ListObjectsV2调用。这是
      更高的性能，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个列表对象方法。如果猜测错误，则可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，使用控制字符的文件名会更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果要尽量减少rclone的交易次数，或者知道存储桶已经存在，则这可能很有用。
      
      如果使用的用户没有创建存储桶的权限，则可能也需要使用此标志。在v1.52.0之前，由于错误，这个操作会忽略。
      

   --no-head
      如果设置，则不对上传的对象进行HEAD请求以检查完整性。
      
      如果尽量减少rclone的交易次数，这可能很有用。
      
      设置它意味着如果rclone在PUT上传对象后收到200 OK消息，则会假设它已正确上传。
      
      特别是，它会假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的文件相同
      - 大小与上传的文件相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象的长度未知，则rclone将执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，
      特别是大小不正确，因此不推荐在正常操作中使用。实际上，即使使用此标志，未检测到的上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在获取对象之前不执行HEAD请求。

   --encoding
      后端的编码。
      
      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的频率。
      
      需要附加缓冲区（例如分部分地）的上传将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端使用http2。
      
      目前，s3（特别是minio）后端存在一个未解决的HTTP/2问题。
      S3后端默认启用HTTP/2，但可以在此处禁用。问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据提供更便宜的出口。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该设置为true、false或留空以使用提供商的默认设置。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传
      
      如果为false，则rclone将使用来自AWS SDK的PutObject来上传对象。
      
      rclone < 1.59 的版本使用预签名请求来上传单个部分的对象，将此标志设置为true将重新启用该 functionality。
      这在例外情况下或进行测试时才是必需的，普通情况下不应该使用。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应该是一个日期，"2006-01-02"，时间日期 "2006-01-02
      15:04:05"，或者表示多久以前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此功能时，不允许进行文件写入操作，
      因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置将对象上传到S3。通常，rclone将以压缩的对象形式下载这些文件。
      
      如果设置了此标志，则rclone将在接收到带有"Content-Encoding: gzip"的文件时进行解压缩。这意味着rclone无法检查文件的大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此标志。
      
      通常情况下，提供程序在下载对象时不会修改对象。如果
      对象没有使用`Content-Encoding: gzip`上传，那么在下载时它也不会被设置。
      
      但是，一些提供程序可能会对对象进行gzip压缩，即使它们没有使用`Content-Encoding: gzip`上传（例如Cloudflare）。
      
      这种情况的一种症状是接收到以下错误信息：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone使用
      Content-Encoding: gzip 设置和分块传输编码下载对象，则rclone将动态解压缩该对象。
      
      如果将此设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但是您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id 值           AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl 值                     创建存储桶和存储或复制对象时使用的默认ACL。[$ACL]
   --endpoint 值                S3 API的端点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint 值     位置约束 - 必须设置与区域匹配。[$LOCATION_CONSTRAINT]
   --region 值                  要连接的区域。[$REGION]
   --secret-access-key 值       AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption 值  存储对象在S3中使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id 值          如果使用KMS ID，则必须提供Key的ARN。[$SSE_KMS_KEY_ID]

   高级选项

   --bucket-acl 值               创建存储桶时使用的默认ACL。[$BUCKET_ACL]
   --chunk-size 值               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff 值              切换到分块拷贝的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                  如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum            不将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2               禁用S3后端使用http2。 (默认值: false) [$DISABLE_HTTP2]
   --download-url 值             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding 值                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style            如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk 值               要列出的大小（每个ListObject S3请求的响应列表大小）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode 值          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version 值             要使用的ListObjects的版本：1、2或0表示自动选择。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts 值         多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time 值   内部内存缓冲区池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap        是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip 值               如果后端可能对对象进行gzip压缩，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket             如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                     如果设置，则不对上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object              如果设置，则在获取对象之前不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata          禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile 值                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token 值            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file 值  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm 值   如果使用SSE-C，在存储对象时使用的服务器端加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value      若要使用SSE-C，您可以提供用于加密/解密数据的秘密加密密钥。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用SSE-C，则必须提供以base64格式编码的秘密加密密钥，用于加密/解密数据。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 值     如果使用SSE-C，可以提供秘密加密密钥的MD5校验和（可选）。 [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency 值       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 值            切换到分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag 值       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request       是否使用预签名请求或PutObject进行单个部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                     如果为true，则使用v2认证。 (默认值: false) [$V2_AUTH]
   --version-at 值               显示指定时间的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                    在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

``` 
{% endcode %}