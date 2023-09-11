# Wasabi对象存储

{% code fullWidth="true" %}
```
命名:
   singularity storage create s3 wasabi - Wasabi对象存储

用法:
   singularity storage create s3 wasabi [command options] [arguments...]

描述:
   --env-auth
      从运行环境中获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时适用。

      示例:
         | false | 下一步输入AWS凭据。
         | true  | 从环境（环境变量或IAM）获取AWS凭据。

   --access-key-id
      AWS访问密钥ID。
      
      如果要匿名访问或使用运行时凭据，请将其留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      如果要匿名访问或使用运行时凭据，请将其留空。

   --region
      要连接的区域。
      
      如果你使用的是S3克隆，并且没有设置区域，请将其留空。

      示例：
         | <未设置>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名无效时使用此选项。
         |                    | 例如，Jewel/v10 CEPH之前。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时必需。

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
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP Northeast 1（东京）端点
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP Northeast 2（大阪）端点
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP Southeast 1（新加坡）
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP Southeast 2（悉尼）

   --location-constraint
      位置约束-必须设置为与区域匹配。
      
      如果不确定，请将其留空。仅在创建存储桶时使用。

   --acl
      创建存储桶、存储或复制对象时使用的预设ACL。
      
      此ACL用于创建对象，如果未设置bucket_acl，则用于创建存储桶。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，在服务器端复制对象时，S3不会复制源上的ACL，而是写入一个新的ACL。
      
      如果acl是一个空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预设ACL。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。
      
      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

      示例：
         | private            | 属主拥有FULL_CONTROL。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 属主拥有FULL_CONTROL。
         |                    | AllUsers组有读取权限。
         | public-read-write  | 属主拥有FULL_CONTROL。
         |                    | AllUsers组有读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 属主拥有FULL_CONTROL。
         |                    | AuthenticatedUsers组有读取权限。

   --upload-cutoff
      切换到分块上传的截止点。
      
      大于此大小的文件将使用chunk_size分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或谷歌照片或谷歌文档上传的文件）时，
      它们将使用此块大小进行分块上传。
      
      请注意，每个传输的内存中有"--s3-upload-concurrency"大小的块。
      
      如果您正在高速链接上传输大型文件，并且有足够的内存，则增加此值可以加快传输速度。
      
      当上传已知大小的大文件时，rclone会自动增大块大小，以保持在10000块的限制以下。
      
      不知道大小的文件使用配置的chunk_size进行上传。由于默认的块大小为5 MiB，并且最多可以有10000个块，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。
      如果您希望流式上传更大的文件，则需要增加chunk_size。
      
      增加块的大小会降低使用"-P"标志显示的进度统计的准确性。在发送的当前块被AWS SDK缓冲时，rclone将该块视为已发送，而实际上它可能仍在上传。
      较大的块大小意味着更大的AWS SDK缓冲区和进度报告与实际情况的偏离。
      

   --max-upload-parts
      multipart上传中的最大部分数。
      
      该选项定义在执行multipart上传时使用的最大多部分块数。
      
      如果某个服务不支持AWS S3规范中的10,000个块，这将非常有用。
      
      当上传已知大小的大文件时，rclone会自动增大块大小，以保持低于此块数的限制。
      

   --copy-cutoff
      切换到分块复制的截止点。
      
      需要服务器端复制的大于此大小的文件将按此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和与对象元数据一起存储。
      
      通常，rclone将在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，
      但对于大文件来说，可能会导致长时间的上传前的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量的值为空，则将默认为当前用户的主目录。
      
          Linux/OSX：“$HOME/.aws/credentials”
          Windows：“%USERPROFILE%\.aws\credentials”
      

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，它将默认为环境变量“AWS_PROFILE”或“default”（如果未设置该环境变量）。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      multipart上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您正在高速链接上上传少量大文件，并且这些上传未能充分利用带宽，则增加此值可以帮助加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见AWS S3文档。
      
      某些提供程序（例如AWS，阿里云OSS，网易COS或腾讯COS）要求将其设置为false-rclone将根据提供程序设置进行自动设置。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，rclone将使用v2身份验证。
      
      仅在v4签名无效时使用，例如，jewel/v10 CEPH之前。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      该选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。
      大多数服务即使请求更多的内容，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，不能修改，请参见AWS S3。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加它。
      

   --list-version
      要使用的ListObjects的版本：1、2或0代表自动。
      
      当S3最初推出时，它仅提供了ListObjects调用来枚举存储桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。这样做的性能要高得多，应该在可能的情况下使用。
      
      如果设置为默认值0，则rclone将根据设置的提供方式猜测调用哪个列表对象方法。如果它猜错了，那么可以在此处手动设置它。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      有些提供商支持URL编码列表，如果可用，则在使用控制字符的文件名时，这是比较可靠的。如果将其设置为unset（默认值），
      则rclone将根据供应商设置选择要应用的选项，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果要在最小化rclone事务数的同时尽量减少rclone的操作数、或者您知道存储桶已存在，这可能很有用。
      
      如果您使用的用户没有存储桶创建权限，则可能需要这样做。在v1.52.0之前，由于一个错误，这将被静默通过。
      

   --no-head
      如果设置，则不对已上传的对象进行HEAD请求以检查完整性。
      
      这在尝试最小化rclone的事务数时很有用。
      
      设置后，意味着如果rclone在PUT上传对象后收到200 OK消息，它将认为该对象已上传正确。
      
      特别是它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传一样
      - 大小与上传的大小相同
      
      它从单一部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象的大小未知，则rclone将执行HEAD请求。
      
      设置此标志将增加未检测到上传失败的机会，特别是大小不正确，因此不建议正常操作。实际上，即使使用此标志，检测到的上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求以获取GET之前的状态。

   --encoding
      后端的编码。
      
      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的频率。
      
      需要使用额外缓冲区（例如，大块上传）的上传将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在一个未解决的问题，与HTTP/2相关。
      默认情况下，s3后端启用HTTP/2，但可以在此禁用。解决此问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义端点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据提供了更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      应该为true、false或留空以使用提供程序的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果此标志设置为false，则rclone将使用AWS SDK中的PutObject来上传对象。
      
      rclone的版本< 1.59使用预签名请求来上传单部分对象，将此标志设置为true将重新启用该功能。
      除非特殊情况或测试，通常不需要这样做。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示在指定时间点的文件版本。
      
      参数应该是日期“2006-01-02”、日期时间“2006-01-02
      15:04:05”或long ago的持续时间，例如“100d”或“1h”。
      
      请注意，使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”在S3上上传对象。通常，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，rclone将在接收到具有“Content-Encoding: gzip”的文件时解压缩它们。这意味着rclone无法检查大小和哈希，
      但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能压缩对象，则设置此项。
      
      通常情况下，提供商在下载对象时不会更改对象。如果对象没有使用“Content-Encoding: gzip”进行上传，那么下载时将不会设置。
      
      但是，一些提供商可能会压缩对象，即使它们没有使用“Content-Encoding: gzip”进行上传（例如Cloudflare）。
      
      这将会收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了具有设置了Content-Encoding: gzip和分块传输编码的对象，那么rclone将动态地执行对象解压缩。
      
      如果设置为unset（默认值），则rclone将根据提供程序设置选择要应用的选项，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      阻止设置和读取系统元数据


选项:
   --access-key-id value        AWS访问密钥ID。 [$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预设ACL。 [$ACL]
   --endpoint value             S3 API的端点。 [$ENDPOINT]
   --env-auth                   从运行环境中获取AWS凭据（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束-必须设置为与区域匹配。 [$LOCATION_CONSTRAINT]
   --region value               要连接的区域。 [$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               创建存储桶时使用的预设ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset。 (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1,2或0代表auto。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         multipart上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此项，如果后端可能压缩对象。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对已上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求以获取GET之前的状态。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             阻止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       multipart上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示在指定时间点的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name value  存储的名称（默认：自动生成）
   --path value  存储的路径

```
{% endcode %}