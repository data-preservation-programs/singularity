# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
名称:
   singularity storage create s3 aws - 亚马逊云服务 (AWS) S3

用法:
   singularity storage create s3 aws [命令选项] [参数...]

描述:
   --env-auth
      从运行时获取AWS凭据（如果访问密钥ID和秘密访问密钥为空，则从环境变量或EC2/ECS元数据中获取）。
      
      仅当access_key_id和secret_access_key为空时才应用。

      示例:
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境中获取AWS凭据（环境变量或IAM）。

   --access-key-id
      AWS Access Key ID。
      
      匿名访问或运行时凭据，请留空。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      匿名访问或运行时凭据，请留空。

   --region
      要连接的区域。

      示例:
         | us-east-1      | 默认的端点 - 如果您不确定，请选择此项。
         |                | 美国区域，包括弗吉尼亚北部或北西太平洋。
         |                | 位置约束为空。
         | us-east-2      | 美国东部（俄亥俄州）区域。
         |                | 需要位置约束为us-east-2。
         | us-west-1      | 美国西部（加利福尼亚北部）区域。
         |                | 需要位置约束为us-west-1。
         | us-west-2      | 美国西部（俄勒冈州）区域。
         |                | 需要位置约束为us-west-2。
         | ca-central-1   | 加拿大（中部）区域。
         |                | 需要位置约束为ca-central-1。
         | eu-west-1      | 欧洲（爱尔兰）区域。
         |                | 需要位置约束为EU或eu-west-1。
         | eu-west-2      | 欧洲（伦敦）区域。
         |                | 需要位置约束为eu-west-2。
         | eu-west-3      | 欧洲（巴黎）区域。
         |                | 需要位置约束为eu-west-3。
         | eu-north-1     | 欧洲（斯德哥尔摩）区域。
         |                | 需要位置约束为eu-north-1。
         | eu-south-1     | 欧洲（米兰）区域。
         |                | 需要位置约束为eu-south-1。
         | eu-central-1   | 欧洲（法兰克福）区域。
         |                | 需要位置约束为eu-central-1。
         | ap-southeast-1 | 亚太地区（新加坡）区域。
         |                | 需要位置约束为ap-southeast-1。
         | ap-southeast-2 | 亚太地区（悉尼）区域。
         |                | 需要位置约束为ap-southeast-2。
         | ap-northeast-1 | 亚太地区（东京）区域。
         |                | 需要位置约束为ap-northeast-1。
         | ap-northeast-2 | 亚太地区（首尔）区域。
         |                | 需要位置约束为ap-northeast-2。
         | ap-northeast-3 | 亚太地区（大阪-本地）区域。
         |                | 需要位置约束为ap-northeast-3。
         | ap-south-1     | 亚太地区（孟买）区域。
         |                | 需要位置约束为ap-south-1。
         | ap-east-1      | 亚太地区（香港）区域。
         |                | 需要位置约束为ap-east-1。
         | sa-east-1      | 南美洲（圣保罗）区域。
         |                | 需要位置约束为sa-east-1。
         | me-south-1     | 中东（巴林）区域。
         |                | 需要位置约束为me-south-1。
         | af-south-1     | 非洲（开普敦）区域。
         |                | 需要位置约束为af-south-1。
         | cn-north-1     | 中国（北京）区域。
         |                | 需要位置约束为cn-north-1。
         | cn-northwest-1 | 中国（宁夏）区域。
         |                | 需要位置约束为cn-northwest-1。
         | us-gov-east-1  | AWS政府云（美国东部）区域。
         |                | 需要位置约束为us-gov-east-1。
         | us-gov-west-1  | AWS政府云（美国）区域。
         |                | 需要位置约束为us-gov-west-1。

   --endpoint
      S3 API的端点。
      
      如果使用AWS，则留空以使用该区域的默认端点。

   --location-constraint
      位置约束 - 必须设置为与区域匹配的值。
      
      仅在创建存储桶时使用。

      示例:
         | <unset>        | 美国区域，包括弗吉尼亚北部或北西太平洋。
         | us-east-2      | 美国东部（俄亥俄州）区域。
         | us-west-1      | 美国西部（加利福尼亚北部）区域。
         | us-west-2      | 美国西部（俄勒冈州）区域。
         | ca-central-1   | 加拿大（中部）区域。
         | eu-west-1      | 欧洲（爱尔兰）区域。
         | eu-west-2      | 欧洲（伦敦）区域。
         | eu-west-3      | 欧洲（巴黎）区域。
         | eu-north-1     | 欧洲（斯德哥尔摩）区域。
         | eu-south-1     | 欧洲（米兰）区域。
         | EU             | 欧洲区域。
         | ap-southeast-1 | 亚太地区（新加坡）区域。
         | ap-southeast-2 | 亚太地区（悉尼）区域。
         | ap-northeast-1 | 亚太地区（东京）区域。
         | ap-northeast-2 | 亚太地区（首尔）区域。
         | ap-northeast-3 | 亚太地区（大阪-本地）区域。
         | ap-south-1     | 亚太地区（孟买）区域。
         | ap-east-1      | 亚太地区（香港）区域。
         | sa-east-1      | 南美洲（圣保罗）区域。
         | me-south-1     | 中东（巴林）区域。
         | af-south-1     | 非洲（开普敦）区域。
         | cn-north-1     | 中国（北京）区域。
         | cn-northwest-1 | 中国（宁夏）区域。
         | us-gov-east-1  | AWS政府云（美国东部）区域。
         | us-gov-west-1  | AWS政府云（美国）区域。

   --acl
      创建存储桶和存储或复制对象时使用的预设ACL。
      
      对于创建对象以及未设置bucket_acl的情况下，也会使用该ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，S3不会从源复制ACL，而是生成一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预设ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅当创建存储桶时才使用该ACL。如果未设置，则使用"acl"。
      
      如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加X-Amz-Acl:
      标头，并且将使用默认值（private）。

      示例:
         | private            | 属主获得FULL_CONTROL权限。
         |                    | 没有其他用户可以访问（默认）。
         | public-read        | 属主获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 属主获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 一般不推荐将此权限授予存储桶。
         | authenticated-read | 属主获得FULL_CONTROL权限。
         |                    | AuthenticatedUsers组获得读取权限。

   --requester-pays
      与S3存储桶交互时启用请求方支付选项。

   --server-side-encryption
      存储此对象在S3中时使用的服务器端加密算法。

      示例:
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用SSE-C，则存储此对象在S3中时使用的服务器端加密算法。

      示例:
         | <unset> | None
         | AES256  | AES256

   --sse-kms-key-id
      如果使用KMS ID，则必须提供密钥的ARN。

      示例:
         | <unset>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      若要使用SSE-C，可以提供用于加密/解密数据的密钥。

      或者，您可以提供 --sse-customer-key-base64。

      示例:
         | <unset> | None

   --sse-customer-key-base64
      如果使用SSE-C，则必须提供以base64格式编码的密钥，以便加密/解密数据。

      或者，您可以提供 --sse-customer-key。

      示例:
         | <unset> | None

   --sse-customer-key-md5
      如果使用SSE-C，则可以提供密钥的MD5校验和（可选）。
      
      如果留空，则会从提供的sse_customer_key自动生成。

      示例:
         | <unset> | None

   --storage-class
      存储新对象时要使用的存储类别。

      示例:
         | <unset>             | 默认存储类别
         | STANDARD            | 标准存储类别
         | REDUCED_REDUNDANCY  | 低冗余存储类别
         | STANDARD_IA         | 标准低频访问存储类别
         | ONEZONE_IA          | 单区域低频访问存储类别
         | GLACIER             | 冰川存储类别
         | DEEP_ARCHIVE        | 冰川深度归档存储类别
         | INTELLIGENT_TIERING | 智能层级存储类别
         | GLACIER_IR          | 冰川即时检索存储类别

   --upload-cutoff
      切换到分块上传的截止限制。
      
      大于此大小的文件将按照chunk_size分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或Google
      Photos或Google Docs上传的文件）时，将使用此分块大小进行分块上传。
      
      请注意，每个传输在内存中缓冲"--s3-upload-concurrency"个该大小的chunk。
      
      如果您正在通过高速链接传输大型文件，并且有足够的内存，则增加此值将加快传输速度。
      
      当向已知大小的大型文件上传时，rclone将自动增加分块大小，以保持在10000个分块限制以下。
      
      未知大小的文件以配置的chunk_size上传。由于默认的chunk_size为5 MiB，并且最多可以有
      10,000个chunk，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果您需要
      流式上传更大的文件，则需要增加chunk_size。
      
      增大chunk大小会降低带有"-P"标志的进度统计的准确性。rclone在使用AWS SDK缓冲chunk时将chunk视为已发送，
      而实际上可能还在上传中。较大的chunk大小意味着较大的AWS SDK缓冲区，以及与真实情况偏离更多的进度报告。

   --max-upload-parts
      单个分块上传中的最大分块数。
      
      该选项定义了在进行分块上传时要使用的最大分块数。
      
      如果某个服务不支持AWS S3的规范中的10000个分块，则这对于我们很有用。
      
      当上传已知大小的大型文件时，rclone将自动增加分块大小，以保持该分块数以下。

   --copy-cutoff
      切换到分块复制的截止限制。
      
      需要复制的大于此大小的文件将被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和存储在对象元数据中。
      
      通常，rclone会在上传之前计算输入的MD5校验和，以便在对象的元数据中添加它。这对于数据完整性检查非常有用，
      但可能会导致大文件开始上传时出现长时间延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果 env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，rclone将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果 env_auth = true，则rclone可以使用共享凭据文件。此变量用于控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default" （如果该环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的相同文件块的数量。
      
      如果您正在通过高速链接上传少量大型文件，并且这些上传未充分利用带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟主机样式访问。有关更多信息，请参阅
      [AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      一些提供商（例如: AWS，阿里云OSS，网易COS或腾讯COS）要求将其设置为
      false - rclone将根据提供者设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了它，rclone将使用v2身份验证。
      
      仅当v4签名不起作用时（例如，针对Jewel/v10之前版本的CEPH）才使用此选项。

   --use-accelerate-endpoint
      如果为true，则使用AWS S3加速端点。
      
      请参阅[S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)。

   --leave-parts-on-error
      如果为true，则在失败时避免调用中止上传，为手动恢复保留所有成功上传的部分。
      
      对于在不同会话中恢复上传时，它应设置为true。
      
      警告: 存储不完整的分块上传占用S3上的空间，并且如果不清理将添加额外的费用。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。
      大多数服务即使请求超过此长度也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0进行自动选择。
      
      当S3最初发布时，它只提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月引入了ListObjectsV2调用。这个调用的性能要高得多，如果可能的话应该使用。
      
      如果设置为默认值0，则rclone将根据设置的提供者来猜测调用哪个列举对象方法。如果猜测错误，那么可以在此手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset。
      
      某些提供商支持对列表进行URL编码，在提供方设置允许的情况下，这是在文件名中使用控制字符时更可靠的方法。如果
      设置为未设置（默认值），则rclone将根据提供方设置来应用适当的方式，但可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，请不要尝试检查存储桶是否存在或创建存储桶。
      
      这在试图最小化rclone执行的事务数时非常有用，如果您知道存储桶已经存在。

   --no-head
      如果设置，请勿对上传的对象进行HEAD请求以检查完整性。
      
      这在试图使rclone的事务数最小化时非常有用。
      
      设置它意味着如果rclone在通过PUT上传对象后收到200 OK消息，那么它将假设它成功上传了。
      
      特别是，它将假设：
      
      - 元数据，包括最后修改时间，存储类别和内容类型与上传时相同。
      - 大小与上传时相同。
      
      它从单个part PUT 请求的响应中读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源的大小未知，则rclone **将**进行HEAD请求。
      
      设置此标志会增加未被检测到的上传失败的几率，特别是大小不正确，因此不推荐在正常操作中使用。实际上，即使有此标志，检测到的上传错误的几率也非常小。
      

   --no-head-object
      如果设置，请在GET获取对象时不进行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅概述中的[编码方法](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。
      
      需要额外缓冲区的上传（例如，分块上传）将使用内存池进行分配。
      此选项控制多久从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存缓冲池中使用mmap缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。
      
      目前 s3 (特别是 minio) 后端与 HTTP/2 存在未解决的问题。
      HTTP/2 在 S3 后端默认启用，但可以在这里禁用。
      在问题解决之前，这个标志将被保留。
      
      参见: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      自定义下载终端点。
      这通常设置为 CloudFront CDN URL，因为 AWS S3 提供通过 CloudFront 网络下载的更便宜的出口流量。

   --use-multipart-etag
      是否在分块上传中使用ETag以进行验证。
      
      此值应为true、false或未设置，以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传。
      
      如果为false，rclone将使用PutObject从AWS SDK上传对象。
      
      版本小于1.59的rclone使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除非特殊情况或测试需要，否则不应该需要设置此标志。

   --versions
      在目录列表中包括旧版本文件。

   --version-at
      显示文件版本，如指定的时间点。
      
      参数应为日期（"2006-01-02"）、日期时间（"2006-01-02 15:04:05"）或一个持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此选项时不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用 "Content-Encoding: gzip" 设置对象上传到S3。通常，
      rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在接收到这些文件时对带有 "Content-Encoding: gzip" 的文件进行解压缩。这意味着rclone无法检查大小和散列值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能压缩对象，请设置此标志。
      
      通常情况下，提供者在下载对象时不会更改对象。如果一个对象在上传时没有使用 `Content-Encoding: gzip` 上传，那么在下载时也不会设置它。
      
      但是，即使未使用 `Content-Encoding: gzip` 上传对象，一些提供商可能也会对对象进行 gzip 压缩（例如 Cloudflare）。
      
      这将导致收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了带有设置了 Content-Encoding: gzip 和分块传输编码的对象，则rclone将在传输过程中实时解压缩对象。
      
      如果将此设置为未设置（默认值），则rclone将根据提供商的设置选择应用何种转换，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

   --sts-endpoint
      STS的终端节点。
      
      如果留空，则使用默认的区域终端节点。

选项:
   --access-key-id value           AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                     创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value                S3 API的端点。[$ENDPOINT]
   --env-auth                      从运行时获取AWS凭据（如果访问密钥ID和秘密访问密钥为空，则从环境变量或EC2/ECS元数据中获取）。 (默认值：false) [$ENV_AUTH]
   --help, -h                      显示帮助
   --location-constraint value     位置约束 - 必须设置为与区域匹配的值。[$LOCATION_CONSTRAINT]
   --region value                  要连接的区域。[$REGION]
   --secret-access-key value       AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  存储此对象在S3中时使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          如果使用KMS ID，则必须提供密钥的ARN。[$SSE_KMS_KEY_ID]
   --storage-class value           存储新对象时要使用的存储类别。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值： "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止限制。 (默认值： "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值： false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和存储在对象元数据中。 (默认值： false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 http2 使用。 (默认值： false) [$DISABLE_HTTP2]
   --download-url value             自定义下载终端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值： "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值： true) [$FORCE_PATH_STYLE]
   --leave-parts-on-error           如果为true，则在失败时避免调用中止上传，为手动恢复保留所有成功上传的部分。 (默认值： false) [$LEAVE_PARTS_ON_ERROR]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (默认值： 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值： "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1, 2或0进行自动选择。 (默认值： 0) [$LIST_VERSION]
   --max-upload-parts value         单个分块上传中的最大分块数。 (默认值： 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 (默认值： "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲池中使用mmap缓冲区。 (默认值： false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请设置此标志。 (默认值： "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，请不要尝试检查存储桶是否存在或创建存储桶。 (默认值： false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，请勿对上传的对象进行HEAD请求以检查完整性。 (默认值： false) [$NO_HEAD]
   --no-head-object                 如果设置，请在GET获取对象时不进行HEAD请求。 (默认值： false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值： false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --requester-pays                 与S3存储桶交互时启用请求方支付选项。 (默认值： false) [$REQUESTER_PAYS]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用SSE-C，则存储此对象在S3中时使用的服务器端加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         若要使用SSE-C，可以提供用于加密/解密数据的密钥。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用SSE-C，则必须提供以base64格式编码的密钥，以便加密/解密数据。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用SSE-C，则可以提供密钥的MD5校验和（可选）。 [$SSE_CUSTOMER_KEY_MD5]
   --sts-endpoint value             STS的终端节点。 [$STS_ENDPOINT]
   --upload-concurrency value       分块上传的并发数。 (默认值： 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止限制。 (默认值： "200Mi") [$UPLOAD_CUTOFF]
   --use-accelerate-endpoint        如果为true，则使用AWS S3加速端点。 (默认值： false) [$USE_ACCELERATE_ENDPOINT]
   --use-multipart-etag value       是否在分块上传中使用ETag以进行验证 (默认值： "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单个部分上传 (默认值： false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值： false) [$V2_AUTH]
   --version-at value               显示文件版本，如指定的时间点。 (默认值： "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本文件。 (默认值： false) [$VERSIONS]

   通用

   --name value  存储的名称（默认：自动生成的）
   --path value  存储的路径

```
{% endcode %}