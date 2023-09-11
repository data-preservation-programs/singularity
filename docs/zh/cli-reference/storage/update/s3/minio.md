# Minio对象存储

{% code fullWidth="true" %}
```
NAME:
    singularity storage update s3 minio - Minio对象存储

USAGE:
   singularity storage update s3 minio [command options] <name|id>

DESCRIPTION:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      只有当access_key_id和secret_access_key为空时才适用。
      
      示例：
         | false | 在下一步输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或使用运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或使用运行时凭证。

   --region
      要连接的区域。
      
      如果使用S3克隆并且没有区域，请留空。
      
      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空白区域。
         | other-v2-signature | 仅当v4签名无效时使用此选项。
         |                    | 例如Pre Jewel/v10 CEPH。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要提供。

   --location-constraint
      位置约束-必须设置为与区域匹配。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，并且如果未设置bucket_acl，则还用于创建存储桶。
      
      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，此ACL在服务器端复制对象时应用，因为S3不会复制源的ACL，而是自己写入新的ACL。
      
      如果ACL是空字符串，则不会添加X-Amz-Acl：标头，并且将使用默认值（私有）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。如果未设置它，则使用“acl”。
      
      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl：标头，并且将使用默认值（私有）。

      示例：
         | private            | 所有者拥有FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读权限。
         | public-read-write  | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读和写权限。
         |                    | 不建议在存储桶上设置此权限。
         | authenticated-read | 所有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读权限。

   --server-side-encryption
      在S3中存储此对象时使用的服务器端加密算法。

      示例：
         | <unset> | 无
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用SSE-C，存储此对象时使用的服务器端加密算法。

      示例：
         | <unset> | 无
         | AES256  | AES256

   --sse-kms-key-id
      如果使用KMS ID，则必须提供密钥的ARN。

      示例：
         | <unset>                 | 无
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      要使用SSE-C，可以提供用于加密/解密数据的秘密加密密钥。
      
      或者，您可以提供 --sse-customer-key-base64。

      示例：
         | <unset> | 无

   --sse-customer-key-base64
      如果使用SSE-C，则必须使用基于base64的格式提供秘密加密密钥，用于加密/解密数据。
      
      或者，您可以提供 --sse-customer-key。

      示例：
         | <unset> | 无

   --sse-customer-key-md5
      如果使用SSE-C，可以提供秘密加密密钥的MD5校验和（可选）。
      
      如果将其留空，则会自动从提供的 sse_customer_key 计算。

      示例：
         | <unset> | 无

   --upload-cutoff
      切换到分块上传的大小。
      
      超过此大小的文件将按chunk_size分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如来自“rclone rcat”或使用“rclone mount”或Google图片或Google文档上传的文件）时，
      将使用此块大小进行分块上传。
      
      请注意，“--s3-upload-concurrency”每个传输在内存中缓冲区了该大小的块。
      
      如果您通过高速链接传输大文件并且有足够的内存，增加此值将加快传输速度。
      
      当上传已知大小的大文件时，rclone将自动增加块大小以保持在10000块的限制以下。
      
      未知大小的文件使用配置的块大小进行上传。由于默认的块大小为5 MiB，并且最多有10000个块，因此默认情况下，您可以流式传输的文件的最大大小为48 GiB。。
      如果要传输更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用“-P”标志显示的进度统计的准确性。当缓冲的AWS SDK块已发送时，rclone会将块视为已发送，而实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和进度报告与真实情况偏离更大。
      

   --max-upload-parts
      多部分上传中的最大上传部件数。
      
      此选项定义在执行多部分上传时要使用的最大分块数。
      
      如果服务不支持AWS S3规范的10000个分块，则此选项可能很有用。
      
      当上传已知大小的大文件时，rclone将自动增加块大小以保持在此分块限制的下方。
      

   --copy-cutoff
      切换到多部分复制的大小。
      
      需要以服务器端复制的文件大于此大小的文件将被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      在对象元数据中不存储MD5校验和。
      
      通常，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间的延迟才能开始上传。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”或“default”如果该环境变量也未设置。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件块数。
      
      如果您在高速链接上传少量大文件，并且这些上传未充分利用您的带宽，则增加此数值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问，如果为false，则使用虚拟主题样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见[the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求设置为false-根据提供商的设置，rclone会自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，rclone将使用v2身份验证。
      
      仅在v4签名无效时使用此选项，例如Pre Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表的大小）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求数量超过1000个，也会将响应列表截断为1000个对象。
      在AWS S3中，这是最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过“rgw list buckets max chunk”选项进行增加。
      

   --list-version
      要使用的ListObjects版本：1、2或0（自动）。
      
      当S3最初发布时，它仅提供了ListObjects调用来枚举存储桶中的对象。
      
      但是，2016年5月引入了ListObjectsV2调用。这是更高性能的，并且应尽可能使用。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个list objects方法。如果它猜错了，那么可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果这可用，则在文件名中使用控制字符时，这更可靠。如果设置为unset（默认值），
      则rclone将根据提供商设置的内容选择要应用的选项，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建该存储桶。
      
      如果您知道存储桶已经存在，这对于尽量减少rclone执行的事务数量可能很有用。
      
      如果正在使用的用户没有创建存储桶的权限，则也可能需要这样做。v1.52.0之前，由于错误，此编写不会发出错误消息。

   --no-head
      如果设置，则不对上传对象进行HEAD请求以检查完整性。
      
      如果要尽量减少rclone执行的事务数量，则这可能很有用。
      
      设置后，这意味着如果在PUT上传对象之后，rclone接收到200 OK消息，则假设它已正确上传。
      
      特别是它将假设：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传时相同
      - 大小与上传时相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传源对象的长度未知，则rclone**将执行HEAD请求**。
      
      设置此标志会增加未检测到的上传失败的几率，特别是不正确的大小，因此不建议在正常操作中使用它。事实上，即使使用此标志，
      未检测到的上传失败的几率也非常小。

   --no-head-object
      如果设置，则在GET对象时不执行HEAD请求。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的间隔时间。
      
      需要对上传使用额外缓冲区（例如多部分）将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用内存映射缓冲区。

   --disable-http2
      为S3后端禁用使用http2。
      
      当前s3（具体地说是minio）后端存在一个未解决的问题，涉及HTTP/2。默认情况下，s3后端启用HTTP/2，但可以在此禁用。解决了该问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端。
      这通常设置为CloudFront CDN URL，因为AWS S3 提供了通过 CloudFront 网络下载数据的更便宜的出口。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该是true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用签署的请求还是PutObject进行单部分上传
      
      如果为false，rclone将使用AWS SDK中的PutObject上传对象。
      
      rclone < 1.59版本使用预签名请求来上传单部分对象，将此标志设置为true将重新启用该功能。除非特殊情况或测试，
      否则不应使用此选项。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定的时间显示文件版本。
      
      参数应为日期，“2006-01-02”，日期时间“2006-01-02 15:04:05”或表示很久以前的持续时间，例如“100d”或“1h”。
      
      请注意，在使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，这将解压缩gzip编码对象。
      
      可以使用“Content-Encoding: gzip”将对象上传到S3。通常，rclone将以压缩的对象形式下载这些文件。
      
      如果设置了此标志，那么rclone将在接收到“Content-Encoding: gzip”后解压缩这些文件。这意味着rclone无法检查大小和哈希，
      但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此标志。
      
      通常情况下，提供商在下载对象时不会更改对象。如果对象没有使用“Content-Encoding: gzip”上传，则在下载时它将不会设置。
      
      但是，一些提供商即使未使用“Content-Encoding: gzip”上传文件也可能对其进行gzip压缩（例如Cloudflare）。
      
      这将导致收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志并且rclone下载具有设置了Content-Encoding: gzip和分块传输编码的对象，
      则rclone将会即时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商设置的内容选择要应用的选项，但您可以在此覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value           AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                     创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value                S3 API的端点。[$ENDPOINT]
   --env-auth                      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                      显示帮助信息
   --location-constraint value     位置约束-必须设置为与区域匹配。[$LOCATION_CONSTRAINT]
   --region value                  要连接的区域。[$REGION]
   --secret-access-key value       AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  在S3中存储此对象时使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          如果使用KMS ID，则必须提供密钥的ARN。[$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分复制的大小。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               在对象元数据中不存储MD5校验和。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主题样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表的大小）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0（自动）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大上传部件数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新未使用的缓冲区的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用内存映射缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此标志如果后端可能对对象进行gzip压缩。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对上传对象进行HEAD请求以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用SSE-C，则存储此对象时使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用SSE-C，可以提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用SSE-C，则必须使用基于base64的格式提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用SSE-C，可以提供秘密加密密钥的MD5校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       多部分上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的大小。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用签署的请求还是PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (default: false) [$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]

```
{% endcode %}