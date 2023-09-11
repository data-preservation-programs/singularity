# Wasabi对象存储

{% code fullWidth="true" %}
```
命令格式:
   singularity storage update s3 wasabi - Wasabi Object Storage

用法:
   singularity storage update s3 wasabi [command options] <name|id>

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅适用于access_key_id和secret_access_key为空的情况。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。

      针对匿名访问或运行时凭证请留空。

   --secret-access-key
      AWS机密访问密钥。

      针对匿名访问或运行时凭证请留空。

   --region
      要连接的区域。

      如果您使用的是S3克隆，且没有区域，请留空。

      示例:
         | <unset>            | 使用此值如果不确定。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如，Jewel/v10 CEPH之前。

   --endpoint
      S3 API的终端点。

      使用S3克隆时必须提供。

      示例:
         | s3.wasabisys.com                | Wasabi US East 1（弗吉尼亚北部）
         | s3.us-east-2.wasabisys.com      | Wasabi US East 2（弗吉尼亚北部）
         | s3.us-central-1.wasabisys.com   | Wasabi US Central 1（德克萨斯州）
         | s3.us-west-1.wasabisys.com      | Wasabi US West 1（俄勒冈州）
         | s3.ca-central-1.wasabisys.com   | Wasabi CA Central 1（多伦多）
         | s3.eu-central-1.wasabisys.com   | Wasabi EU Central 1（阿姆斯特丹）
         | s3.eu-central-2.wasabisys.com   | Wasabi EU Central 2（法兰克福）
         | s3.eu-west-1.wasabisys.com      | Wasabi EU West 1（伦敦）
         | s3.eu-west-2.wasabisys.com      | Wasabi EU West 2（巴黎）
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP Northeast 1（东京）终端节点
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP Northeast 2（大阪）终端节点
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP Southeast 1（新加坡）
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP Southeast 2（悉尼）

   --location-constraint
      位置约束，必须与区域匹配。

      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的预设ACL。

      此预设ACL用于创建对象，如果bucket_acl未设置，则也用于创建存储桶。

      有关更多信息，请访问[Amazon S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      请注意，当服务器端复制对象时，此ACL会应用。因为S3不会复制源对象的ACL，
      而是写入一个新的ACL。

      如果acl是空字符串，则不会添加X-Amz-Acl:头，将使用默认的（private）。

   --bucket-acl
      创建存储桶时使用的预设ACL。

      有关更多信息，请访问[Amazon S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。

      如果未设置bucket_acl，则仅在创建存储桶时应用“acl”。

      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl:头，
      将使用默认的（private）。

      示例:
         | private            | 所有者拥有FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认设置）。
         | public-read        | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers群组有读取权限。
         | public-read-write  | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers群组有读取和写入权限。
         |                    | 不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers群组有读取权限。

   --upload-cutoff
      切换到分块上传的文件大小阈值。

      大于此大小的文件将按块（chunk_size）上传。
      最小值为0，最大值为5GiB。

   --chunk-size
      用于上传的块大小。

      当上传文件大于upload_cutoff或大小未知（例如"rclone rcat"，
      或使用"rclone mount"、谷歌照片或谷歌文档上传的文件）时，
      将使用此块大小进行分块上传。

      请注意，每个传输每次会在内存中缓冲"--s3-upload-concurrency"个这样大小的块。

      如果您正在通过高速连接传输大文件，并且具有足够的内存，
      增大此值将加快传输速度。

      当上传已知大小的大文件以保持低于10,000块的限制时，
      rclone将自动增加块大小。

      未知大小的文件将使用配置的chunk_size上传。由于默认块大小为5 MiB，
      最多有10,000个块，
      这意味着默认情况下，您可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大的文件，则需要增加chunk_size。

      增大块大小会降低使用"-P"标志时显示的进度统计的准确性。
      当AWS SDK缓冲块时，rclone将发送块，并将其视为已发送，
      实际上还可能正在上传。较大的块大小意味着更大的AWS SDK缓冲区，
      并且进度报告更违背事实。

   --max-upload-parts
      多部分上传的最大分块数。

      此选项定义进行多部分上传时要使用的最大分块数。

      如果服务商不支持AWS S3规范中的10,000个分块，
      这将非常有用。

      当上传已知大小的大文件以保持低于此分块数的限制时，
      rclone将自动增加块大小。

   --copy-cutoff
      切换到分块复制的文件大小阈值。

      需要进行服务器端复制的大于此大小的文件将按照此大小进行复制。

      最小值为0，最大值为5GiB。

   --disable-checksum
      不将MD5校验和存储在对象元数据中。

      通常，在上传之前，rclone会计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性检查很有用，
      但对于大文件的起始上传可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth = true，则rclone可以使用共享凭证文件。

      如果此变量为空，则rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量的值为空，
      则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      
   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。

      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传的文件的块数。

      如果您通过高速链接上传少量大文件，并且这些上传未完全利用您的带宽，
      那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；
      如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟路径样式访问。
      有关更多信息，请参阅[the AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      一些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将其设置为
      false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了此选项，则rclone将使用v2身份验证。

      仅在v4签名不起作用时使用此选项，例如，Jewel/v10 CEPH之前。

   --list-chunk
      列举块的大小（每次ListObject S3请求的响应列表）。

      此选项也称为"MaxKeys"、"max-items"或AWS S3规范中的"page-size"。

      大多数服务即使请求了更多对象，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局的限制，无法更改，详情请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。

      在Ceph中，可以使用"rgw list buckets max chunk"选项来增加此值。

   --list-version
      要使用的ListObjects版本：1,2或0用于自动判断。

      当S3首次推出时，仅提供了ListObjects调用以列举存储桶中的对象。

      然而，在2016年5月，引入了ListObjectsV2调用。这是性能更高的方法，
      应尽可能使用。

      如果设置为默认值0，则根据设置的提供商猜测调用哪个列表对象方法。
      如果猜测错误，则可以在此处手动设置。

   --list-url-encode
      是否URL编码列表：true/false/unset

      某些提供商支持URL编码列表，如果可用，则在使用控制字符的文件名时，这种方法更可靠。
      如果设置为unset（默认值），则rclone将根据提供商的设置选择应用的方法，
      但可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建。

      如果您知道存储桶已经存在，此选项可以有助于最小化rclone执行的事务数。

      如果您正在使用的用户没有创建存储桶的权限，也可能需要将其设置为true。
      在v1.52.0之前的版本中，由于错误，此设置将被静默忽略。

   --no-head
      如果设置，则不通过HEAD请求来检查已上传的对象的完整性。

      如果尝试最小化rclone执行的事务数，这可能会很有用。

      设置后，意味着在PUT上传对象后，如果rclone接收到200 OK消息，
      则会认为其已正确上传。

      特别是，它会认为：

      - 元数据（包括修改时间、存储类和内容类型）与上传的一样
      - 大小与上传的一样

      它会从以下响应中读取单个分块PUT的以下项：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项。

      如果上传的源对象长度未知，则rclone将**会**执行HEAD请求。

      设置此标志会增加未检测到的上传失败的几率，
      特别是错误的大小，因此不推荐在正常操作中使用。实际上，
      即使使用此标志，未检测到的上传失败的几率非常小。

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码。

      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的频率。

      需要额外缓冲区的上传（例如多部分）将使用内存池进行分配。
      此选项控制多久将未使用的缓冲区从内存池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP2使用。

      目前，s3（特别是minio）后端存在一个未解决的问题和HTTP/2的问题。
      s3后端默认启用HTTP/2，但可以在此禁用。
      解决问题后，此标志将被删除。

      请参阅: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端点。
      通常将其设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的数据的较低费用。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证。

      这应该是true、false或将其设置为默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传。

      如果为false，则rclone将使用AWS SDK中的PutObject上传对象。

      rclone的版本 < 1.59使用预签名请求上传单个部分对象，
      并将此标志设置为true将重新启用该功能。
      除非特殊情况或用于测试，否则不应该需要它。

   --versions
      在目录列表中包含旧版本。

   --version-at
      以指定的时间显示文件版本。

      参数应该是一个日期，"2006-01-02"，datetime "2006-01-02
      15:04:05"，或以那么久以前的持续时间，例如"100d"或"1h"。

      请注意，在使用此选项时，不允许执行文件写入操作，
      因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则将解压缩gzip编码的对象。

      可以使用"Content-Encoding: gzip"设置上传对象到S3。
      通常，rclone会将这些文件以压缩的对象形式下载。

      如果设置了此标志，则rclone将在接收到这些以"Content-Encoding: gzip"接收到的文件时进行解压缩。
      这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会压缩对象，请设置此标志。

      通常，提供商在下载对象时不会更改对象。
      如果一个对象没有使用“Content-Encoding: gzip”进行上传，那么在下载时它也不会被设置。

      但是，即使没有使用“Content-Encoding: gzip”上传对象（例如，Cloudflare），
      某些提供商也可能对对象进行gzip压缩。

      这种情况的症状可能是收到如下错误：

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置了此标志，并且rclone使用"Content-Encoding: gzip"和分块传输编码下载对象，
      则rclone将在获取时即时解压对象。

      如果设置为unset（默认值），则rclone将根据提供商的设置选择应用的值，
      但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id的值       AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl的值                创建存储桶以及存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint的值           S3 API的终端点。[$ENDPOINT]
   --env-auth               从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h               显示帮助
   --location-constraint的值 位置约束，必须与区域匹配。[$LOCATION_CONSTRAINT]
   --region的值             要连接的区域。[$REGION]
   --secret-access-key的值  AWS机密访问密钥。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl的值               创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size的值               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff的值              切换到分块复制的文件大小阈值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                   如果设置，则将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum             不将MD5校验和存储在对象元数据中。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                禁用S3后端的HTTP2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url的值             下载的自定义终端点。[$DOWNLOAD_URL]
   --encoding的值                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style             如果为true，则使用路径样式访问；
                                 如果为false，则使用虚拟主机样式访问。(默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk的值               列举块的大小。（每次ListObject S3请求的响应列表） (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode的值          是否URL编码列表：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version的值             要使用的ListObjects版本：1,2或0用于自动判断。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts的值         多部分上传的最大分块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time的值   内部内存缓冲区池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap         是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip的值               如果后端可能会压缩对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket              如果设置，则不尝试检查存储桶是否存在或创建。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                      如果设置，则不通过HEAD请求来检查已上传的对象的完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object               如果设置，则在获取对象时不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata           禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile的值                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token的值            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file的值  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency的值       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff的值            切换到分块上传的文件大小阈值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag的值       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request        是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                      如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at的值               以指定的时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                     在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}