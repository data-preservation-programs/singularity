# AWS S3 操作文档

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 aws - 亚马逊云服务 (AWS) S3

USAGE:
   singularity storage update s3 aws [command options] <name|id>

DESCRIPTION:
   --env-auth
      通过运行时获取 AWS 凭证（从环境变量或 EC2/ECS 元数据获取）。
      
      仅在 access_key_id 和 secret_access_key 为空时生效。

      示例：
         | false | 在下一步中输入 AWS 凭证。
         | true  | 从环境中（环境变量或 IAM）获取 AWS 凭证。

   --access-key-id
      AWS 访问密钥 ID。
      
      留空以进行匿名访问或使用运行时凭证。

   --secret-access-key
      AWS 秘密访问密钥（密码）。
      
      留空以进行匿名访问或使用运行时凭证。

   --region
      连接的区域。

      示例：
         | us-east-1      | 默认端点-如果不确定，请选择此选项。
         |                | 美国区，弗吉尼亚北部或太平洋西北。
         |                | 将位置约束留空。
         | us-east-2      | 美国东部（俄亥俄州）区域。
         |                | 需要位置约束 us-east-2。
         | us-west-1      | 美国西部（加利福尼亚北部）区域。
         |                | 需要位置约束 us-west-1。
         | us-west-2      | 美国西部（俄勒冈）区域。
         |                | 需要位置约束 us-west-2。
         | ca-central-1   | 加拿大（中部）区域。
         |                | 需要位置约束 ca-central-1。
         | eu-west-1      | 欧洲（爱尔兰）区域。
         |                | 需要位置约束 EU 或 eu-west-1。
         | eu-west-2      | 欧洲（伦敦）区域。
         |                | 需要位置约束 eu-west-2。
         | eu-west-3      | 欧洲（巴黎）区域。
         |                | 需要位置约束 eu-west-3。
         | eu-north-1     | 欧洲（斯德哥尔摩）区域。
         |                | 需要位置约束 eu-north-1。
         | eu-south-1     | 欧洲（米兰）区域。
         |                | 需要位置约束 eu-south-1。
         | eu-central-1   | 欧洲（法兰克福）区域。
         |                | 需要位置约束 eu-central-1。
         | ap-southeast-1 | 亚太地区（新加坡）区域。
         |                | 需要位置约束 ap-southeast-1。
         | ap-southeast-2 | 亚太地区（悉尼）区域。
         |                | 需要位置约束 ap-southeast-2。
         | ap-northeast-1 | 亚太地区（东京）区域。
         |                | 需要位置约束 ap-northeast-1。
         | ap-northeast-2 | 亚太地区（首尔）。
         |                | 需要位置约束 ap-northeast-2。
         | ap-northeast-3 | 亚太地区（大阪-本地）。
         |                | 需要位置约束 ap-northeast-3。
         | ap-south-1     | 亚太地区（孟买）。
         |                | 需要位置约束 ap-south-1。
         | ap-east-1      | 亚太地区（香港）区域。
         |                | 需要位置约束 ap-east-1。
         | sa-east-1      | 南美洲（圣保罗）区域。
         |                | 需要位置约束 sa-east-1。
         | me-south-1     | 中东（巴林）区域。
         |                | 需要位置约束 me-south-1。
         | af-south-1     | 非洲（开普敦）区域。
         |                | 需要位置约束 af-south-1。
         | cn-north-1     | 中国（北京）区域。
         |                | 需要位置约束 cn-north-1。
         | cn-northwest-1 | 中国（宁夏）区域。
         |                | 需要位置约束 cn-northwest-1。
         | us-gov-east-1  | AWS政府云（美国-东部）区域。
         |                | 需要位置约束 us-gov-east-1。
         | us-gov-west-1  | AWS政府云（美国）区域。
         |                | 需要位置约束 us-gov-west-1。

   --endpoint
      S3 API 的端点。
      
      如果使用 AWS，则留空以使用该区域的默认端点。

   --location-constraint
      位置限制-必须设置为匹配区域。
      
      仅在创建存储桶时使用。

      示例：
         | <unset>        | 美国区，弗吉尼亚北部或太平洋西北没有空
         | us-east-2      | 美国东部（俄亥俄州）区域
         | us-west-1      | 美国西部（加利福尼亚北部）区域
         | us-west-2      | 美国西部（俄勒冈）区域
         | ca-central-1   | 加拿大（中部）区域
         | eu-west-1      | 欧洲（爱尔兰）区域
         | eu-west-2      | 欧洲（伦敦）区域
         | eu-west-3      | 欧洲（巴黎）区域
         | eu-north-1     | 欧洲（斯德哥尔摩）区域
         | eu-south-1     | 欧洲（米兰）区域
         | EU             | 欧洲区
         | ap-southeast-1 | 亚太地区（新加坡）区域
         | ap-southeast-2 | 亚太地区（悉尼）区域
         | ap-northeast-1 | 亚太地区（东京）区域
         | ap-northeast-2 | 亚太地区（首尔）区域
         | ap-northeast-3 | 亚太地区（大阪-本地）区域
         | ap-south-1     | 亚太地区（孟买）区域
         | ap-east-1      | 亚太地区（香港）区域
         | sa-east-1      | 南美洲（圣保罗）区域
         | me-south-1     | 中东（巴林）区域
         | af-south-1     | 非洲（开普敦）区域
         | cn-north-1     | 中国（北京）区域
         | cn-northwest-1 | 中国（宁夏）区域
         | us-gov-east-1  | AWS政府云（美国-东部）区域
         | us-gov-west-1  | AWS政府云（美国）区域

   --acl
      在创建存储桶并存储或复制对象时使用的预定义 ACL。
      
      此 ACL 用于创建对象，并且如果未设置 bucket_acl，则也用于创建存储桶。
      
      有关更多信息，请参见 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，在与 S3 进行服务器端复制对象时将应用此 ACL 因为 S3 不会复制来自源的 ACL，而是写入一个新的 ACL。
      
      如果 acl 是空字符串，则不添加 X-Amz-Acl: 标头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预定义 ACL。
      
      有关更多信息，请参见 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此 ACL。如果未设置此选项，将改用“acl”选项。
      
      如果“acl”和“bucket_acl”是空字符串，则不添加 X-Amz-Acl: 标头，并且将使用默认值（private）。

      示例：
         | private            | 拥有者具有 FULL_CONTROL。
         |                    | 无其他人具有访问权限（默认）。
         | public-read        | 拥有者具有 FULL_CONTROL。
         |                    | AllUsers 组具有 READ 访问权限。
         | public-read-write  | 拥有者具有 FULL_CONTROL。
         |                    | AllUsers 组具有 READ 和 WRITE 访问权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者具有 FULL_CONTROL。
         |                    | AuthenticatedUsers 组具有 READ 访问权限。

   --requester-pays
      与 S3 存储桶交互时启用请求者付费选项。

   --server-side-encryption
      存储对象在 S3 中的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用 SSE-C，存储对象在 S3 中的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-kms-key-id
      如果使用 KMS ID，则必须提供密钥的 ARN。

      示例：
         | <unset>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      要使用 SSE-C，可以提供用于加密/解密数据的密钥。
      
      或者您可以提供 --sse-customer-key-base64。

      示例：
         | <unset> | None

   --sse-customer-key-base64
      如果使用 SSE-C，则必须提供以 base64 格式编码的密钥，以便加密/解密数据。
      
      或者您可以提供 --sse-customer-key。

      示例：
         | <unset> | None

   --sse-customer-key-md5
      如果使用 SSE-C，可以提供密钥的 MD5 校验和（可选）。
      
      如果您将其留空，则会自动从提供的 sse_customer_key 计算出来。

      示例：
         | <unset> | None

   --storage-class
      在 S3 中存储新对象时要使用的存储类。

      示例：
         | <unset>             | 默认
         | STANDARD            | 标准存储类
         | REDUCED_REDUNDANCY  | 减少冗余存储类
         | STANDARD_IA         | IA 存储类
         | ONEZONE_IA          | 单个区域的 IA 存储类
         | GLACIER             | Glacier 存储类
         | DEEP_ARCHIVE        | Glacier Deep Archive 存储类
         | INTELLIGENT_TIERING | 智能分层存储类
         | GLACIER_IR          | Glacier Instant Retrieval 存储类

   --upload-cutoff
      切换为分块上传的大小。
      
      大于此值的任何文件将使用 chunk_size 进行分块上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于上传截止点或大小未知的文件时（例如来自 "rclone rcat" 或使用 "rclone mount" 或 Google 照片或 Google 文档上传的文件），
      它们将使用此块大小作为分块上传。
      
      请注意，每次传输内存中会缓冲这个大小的 "--s3-upload-concurrency" 块。
      
      如果您正在通过高速链路传输大文件，并且拥有足够的内存，那么增加此值可以加快传输速度。
      
      Rclone 会在上传已知大小的大文件时，根据需要自动增加块大小，以保持低于 10,000 个块的限制。
      
      未知大小的文件将以配置的块大小上传。由于默认的块大小为 5 MiB，最多可以有 10,000 块，这意味着默认情况下，您可以流式上传的文件的最大大小为 48 GiB。
      如果要流式上传更大的文件，您需要增大 chunk_size。
      
      增加块大小会降低 "-P" 标志显示的进度统计的准确性。当使用默认块大小上传大文件时，rclone 会在传输时将 chunk 视为已发送，而实际上它可能仍在上传。
      更大的块大小意味着更大的 AWS SDK 缓冲区和进度报告与真实情况的更大偏离。

   --max-upload-parts
      分块上传的最大块数。
      
      此选项定义进行分块上传时要使用的最大分块数量。
      
      如果某个服务不支持 AWS S3 规范的 10,000 个分块，这可能非常有用。
      
      Rclone 会在上传已知大小的大文件时，根据需要自动增加块大小，以保持此块数限制。

   --copy-cutoff
      切换为分块复制的大小。
      
      需要服务器端复制的大于此大小的任何文件将按此大小的分块复制。
      
      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常，在上传之前，rclone 会计算输入的 MD5 校验和，以便将其添加到对象的元数据中。对于大文件，这可能会导致长时间的延迟才能开始上传。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。
      
      如果此变量为空，则 rclone 将查找 "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果 env_auth = true，则 rclone 可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果该环境变量也没有设置）。

   --session-token
      AWS 会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传同一文件的块数。
      
      如果您通过高速链路大量上传大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。
      
      如果为 true（默认值），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟主机样式访问。有关更多信息，请参见 [AWS S3 文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些供应商（例如 AWS、阿里云 OSS、网易 COS 或腾讯 COS）要求将此设置为 false - rclone 将根据供应商设置自动完成此操作。

   --v2-auth
      如果为 true，则使用 v2 身份验证。
      
      如果为 false（默认值），则 rclone 将使用 v4 身份验证。如果设置了此值，则 rclone 将使用 v2 身份验证。
      
      仅在 v4 签名无法使用时（例如旧版 Jewel/v10 CEPH）才使用此选项。

   --use-accelerate-endpoint
      如果为 true，则使用 AWS S3 加速端点。
      
      请参见：[AWS S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)

   --leave-parts-on-error
      如果为 true，请在发生故障时避免调用中止上传，将所有成功上传的部分留在 S3 上供手动恢复。
      
      对于在不同会话之间恢复上传时，应将其设置为 true。
      
      警告：不完整多部分上传的各个部分计入 S3 上的空间使用情况，并且如果不清理掉会增加其他成本。

   --list-chunk
      列表块的大小（每个 ListObject S3 请求的响应列表）。
      
      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务在请求超过 1000 个对象时仍会截断响应列表。
      在 AWS S3 中，这是一个全局最大值，无法更改，请参见 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项增加此值。

   --list-version
      要使用的 ListObjects 的版本：1、2 或 0（自动）。
      
      当 S3 最初发布时，它仅提供了 ListObjects 调用以枚举存储桶中的对象。
      
      但是，在 2016 年 5 月，引入了 ListObjectsV2 调用。这更高效，如果可能，应使用该调用。
      
      如果设置为默认值 0，则 rclone 将根据设置的提供者猜测要调用的列表对象方法。如果它猜测错误，则可能在此处手动设置。

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset
      
      某些供应商支持 URL 编码列表，如果此选项可用，则在使用控制字符时，这是更可靠的方式。如果设置为默认值（未设置），则 rclone 将根据供应商设置选择要应用的内容，但您可以在此处覆盖 rclone 的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果知道存储桶已经存在，则这可能很有用，以尽量减少 rclone 的事务数。如果您使用的用户没有创建存储桶的权限，则可能需要设置这个选项。
      在 v1.52.0 之前，由于错误，此选项将通过静默传递。

   --no-head
      如果设置，则不对上传的对象进行 HEAD 操作以检查完整性。
      
      这可能很有用，以尽量减少 rclone 的事务数。
      
      如果上传一个对象，并且 PUT 操作后接收到 200 OK 的消息，则 rclone 将假设它被正确上传。

      特别是它将假设：
      
      - 元数据（包括修改时间、存储类别和内容类型）与上传的文件一致
      - 大小与上传的文件一致
      
      它会从单部分 PUT 的响应中读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于分块上传 响应 不会读取这些项。
      
      如果上传的源对象大小未知，则 rclone **会**执行 HEAD 请求。
      
      设置此标志会增加检测不到上传故障的几率，特别是文件大小不正确的几率，因此不建议在正常操作中使用。实际上，即使设置此标志，未检测到上传失败的几率非常小。

   --no-head-object
      如果设置，则在 GET 获取对象之前不执行 HEAD 操作。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见概述中的 [编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间。
      
      需要其他缓冲区（例如分块）的上传将使用内存池进行分配。
      此选项控制在何时从池中删除未使用的缓冲器。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 HTTP/2 使用。
      
      目前，s3（特别是 minio）后端与 HTTP/2 存在一个未解决的问题。默认情况下，s3 后端启用了 HTTP/2，但可以在此禁用。当问题解决后，将删除此标志。
      
      请参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      通常将其设置为 CloudFront CDN URL，因为 AWS S3 提供了通过 CloudFront 网络下载数据的更便宜的出口流量。

   --use-multipart-etag
      是否在分块上传中使用 ETag 进行验证
      
      它应该是 true、false 或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或 PutObject 来上传单块对象
      
      如果此值为 false，则 rclone 将使用 AWS SDK 的 PutObject 来上传对象。
      
      低于 1.59 版本的 rclone 使用预签名请求来上传单块对象，将此标志设置为 true 将重新启用该功能。除非在特殊情况下或用于测试，否则不应该需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      在指定的时间显示文件版本。
      
      参数应为日期 "2006-01-02"，日期时间 "2006-01-02 15:04:05" 或之前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩 gzip 编码的对象。
      
      可以使用 "Content-Encoding: gzip" 设置将对象上传到 S3。通常 rclone 将以压缩的形式下载这些文件。
      
      如果设置了此标志，则 rclone 将在接收到 "Content-Encoding: gzip" 的文件时对这些文件进行解压缩。这意味着 rclone 无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会为对象使用 gzip，则设置此标志。
      
      通常，供应商在下载对象时不会更改对象。如果对象在上传时未使用 `Content-Encoding: gzip` 上传，则在下载时也不会设置它。
      
      但是，一些供应商可能会压缩对象，即使它们未使用 `Content-Encoding: gzip` 上传（例如 Cloudflare）。
      
      这种情况的症状可能是收到类似于
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志并且 rclone 下载了具有设置了 Content-Encoding: gzip 和分块传输编码的对象，则 rclone 将在流上解压缩对象。
      
      如果此设置为未设置（默认值），则 rclone 将根据供应商设置选择要应用的内容，但您可以在此覆盖 rclone 的选择。

   --no-system-metadata
      抑制系统元数据的设置和读取

   --sts-endpoint
      STS 的端点。
      
      如果使用 AWS，则留空以使用该区域的默认端点。

OPTIONS:
   --access-key-id value           AWS 访问密钥 ID。[$ACCESS_KEY_ID]
   --acl value                     在创建存储桶和存储或复制对象时使用的预定义 ACL。[$ACL]
   --endpoint value                S3 API 的端点。[$ENDPOINT]
   --env-auth                      通过运行时获取 AWS 凭证（从环境变量或 EC2/ECS 元数据获取）。（默认值：false）[$ENV_AUTH]
   --help, -h                      显示帮助
   --location-constraint value     位置限制-必须设置为匹配区域。[$LOCATION_CONSTRAINT]
   --region value                  连接的区域。[$REGION]
   --secret-access-key value       AWS 秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  存储对象在 S3 中的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          如果使用 KMS ID，则必须提供密钥的 ARN。[$SSE_KMS_KEY_ID]
   --storage-class value           在 S3 中存储新对象时要使用的存储类。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               在创建存储桶时使用的预定义 ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。（默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换为整块复制的大小。（默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩 gzip 编码的对象。（默认值：false）[$DECOMPRESS]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。（默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 HTTP/2 使用。（默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。（默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。（默认值：true）[$FORCE_PATH_STYLE]
   --leave-parts-on-error           如果为 true，请在发生故障时避免调用中止上传，将所有成功上传的部分留在 S3 上供手动恢复。（默认值：false）[$LEAVE_PARTS_ON_ERROR]
   --list-chunk value               列表块的大小（每个 ListObject S3 请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset。（默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 的版本：1、2 或 0（自动）。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         分块上传的最大块数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间。（默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。（默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会为对象使用 gzip，则设置此标志。（默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。（默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对上传的对象进行 HEAD 操作以检查完整性。（默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在 GET 获取对象之前不执行 HEAD 操作。（默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             抑制系统元数据的设置和读取（默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --requester-pays                 与 S3 存储桶交互时启用请求者付费选项。（默认值：false）[$REQUESTER_PAYS]
   --session-token value            AWS 会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用 SSE-C，则在 S3 中存储对象时使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用 SSE-C，可以提供用于加密/解密数据的密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用 SSE-C，则必须提供以 base64 格式编码的密钥，以便加密/解密数据。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用 SSE-C，则可以提供密钥的 MD5 校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --sts-endpoint value             STS 的端点。[$STS_ENDPOINT]
   --upload-concurrency value       分块上传的并发数。（默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的大小。（默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-accelerate-endpoint        如果为 true，则使用 AWS S3 加速端点。（默认值：false）[$USE_ACCELERATE_ENDPOINT]
   --use-multipart-etag value       是否在分块上传中使用 ETag 进行验证（默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或 PutObject 来上传单块对象（默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 身份验证。（默认值：false）[$V2_AUTH]
   --version-at value               在指定的时间显示文件版本。（默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。（默认值：false）[$VERSIONS]

```
{% endcode %}