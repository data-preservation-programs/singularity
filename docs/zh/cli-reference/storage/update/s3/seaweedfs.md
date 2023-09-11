# SeaweedFS S3

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 seaweedfs - SeaweedFS S3

用法：
   singularity storage update s3 seaweedfs [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭据。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭据。

   --region
      连接的区域。
      
      如果使用S3克隆并且没有区域，则留空。

      示例：
         | <未设置>                | 不确定时使用此选项。
         |                        | 将使用v4签名和空区域。
         | other-v2-signature     | 仅当v4签名不起作用时使用此选项。
         |                        | 例如Jewel / v10之前的CEPH。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要。

      示例：
         | localhost:8333 | SeaweedFS S3本地主机

   --location-constraint
      位置限制 - 必须设置为匹配的区域。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的Canned ACL。
      
      此ACL用于创建对象，如果没有设置bucket_acl，则用于创建存储桶。
      
      获取更多信息请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，此ACL适用于在S3上进行服务器端复制对象时，
      S3不复制源中的ACL，而是写入新的ACL。
      
      如果acl是空字符串，则不添加X-Amz-Acl：标题，并且
      将使用默认值（private）。
      

   --bucket-acl
      创建存储桶时使用的Canned ACL。
      
      获取更多信息请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。如果没有设置，
      则使用"acl"。
      
      如果“acl”和“bucket_acl”都是空字符串，则不添加X-Amz-Acl：
      标题，并且将使用默认值（private）。
      

      示例：
         | private            | 所有者拥有FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止值。
      
      大于此大小的任何文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件
      （例如来自"rclone rcat"或使用"rclone mount"或google
      相册或google文档上传的文件）时，将使用此块大小进行分块上传。
      
      请注意，每个传输在内存中缓冲chunk_size的"--s3-upload-concurrency"个块。
      
      如果您正在通过高速链接传输大文件并且具有足够的内存，
      那么增加这个值将加快传输速度。
      
      Rclone会在上传已知大小的大文件时自动增加块大小，
      以保持在10000个分块限制以下。
      
      未知大小的文件以配置的chunk_size进行上传。
      由于默认的块大小为5 MiB，并且最多有10000个块，
      这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大的文件，则需要增加chunk_size。
      
      增加块大小会降低用"-P"标志显示的进度
      统计的准确性。当AWS SDK缓冲块时，
      Rclone会将块视为已发送，而实际上可能仍在上传。
      更大的块大小意味着更大的AWS SDK缓冲区和进度
      报告偏离真实情况更远。

   --max-upload-parts
      单个分块上传中的最大块数。
      
      此选项定义在执行多部分上传时使用的最大分块数。
      
      如果某个服务不支持AWS S3的10000个分块规范，
      这可能是有用的。
      
      Rclone会在上传已知大小的大文件时自动增加块大小，
      以保持在此分块数限制以下。

   --copy-cutoff
      切换到分块复制的截止值。
      
      大于此大小且需要进行服务端复制的任何文件
      都会被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性校验很好，
      但对于大文件的开始上传会造成长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果该变量为空，rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE"环境变量。如果env值为空
      则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此
      变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或
      如果未设置该环境变量，则默认为"default"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您通过高速链接上传少量大文件，
      并且这些上传没有充分利用您的带宽，
      那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式访问。
      有关更多信息，请参见AWS S3文档。
      
      某些提供商（例如AWS，Aliyun OSS，Netease COS或Tencent COS）可能需要将此设置为
      false- rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了它，则rclone将使用v2身份验证。
      
      仅在v4签名不起作用时使用，在例如Jewel / v10之前的CEPH。

   --list-chunk
      列举块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"，"max-items"或"page-size"。
      大多数服务即使请求超过了1000个对象，
      也会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，不能更改，请参见https：//docs.aws.amazon.com/cli/latest/reference/s3/ls.html。
      在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。
      

   --list-version
      要使用的ListObjects版本：1、2或0以进行自动选择。
      
      当S3最初发布时，它仅提供了ListObjects调用以
      枚举存储桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。
      这是更高性能的，应尽可能使用它。
      
      如果设置为默认值0，则rclone会根据提供程序
      设置猜测，调用哪个列表对象方法。如果猜错，则可以
      在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset。
      
      某些提供商支持URL编码列表，并且如果可用，使用控制字符使用此编码
      在文件名中更可靠。如果将其设置为unset（默认值），
      则rclone将根据提供商设置选择要应用的内容，但可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，不要尝试检查bucket是否存在或创建它。
      
      如果您知道bucket已经存在，则这可能是减少rclone执行的事务数量的好方法。
      
      如果使用的用户没有bucket创建权限，则可能还需要此选项。在v1.52.0之前，由于错误，
      这将被静默通过。
      

   --no-head
      如果设置，将不会对已上传的对象进行HEAD以检查完整性。
      
      如果尝试最小化rclone执行的事务数量时，这可能很有用。
      
      设置后，这意味着如果rclone在使用PUT上传对象后收到200 OK消息，
      则假设它已正确上传。
      
      特别是它会假设：
      
      - 元数据，包括修改时间，存储类和内容类型与上传时一样
      - 大小与上传时一样
      
      它从单个部分PUT的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传未知长度的源对象，则rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，
      尤其是错误的大小，因此不建议在正常操作中使用。
      实际上，即使使用该标志，未检测到的上传失败的机会也非常小。
      

   --no-head-object
      如果设置，获取对象之前不执行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见概述中的编码部分。

   --memory-pool-flush-time
      内部内存缓冲池将定期刷新的时间间隔。
      
      需要额外的缓冲区（例如multipart）的上传
      将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个未解决的问题
      和HTTP/2之间的关系。S3后端默认启用HTTP/2，但可以
      在此禁用。当问题解决后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端。
      通常将其设置为CloudFront CDN URL，因为AWS S3提供
      通过CloudFront网络下载的数据的更便宜的出口。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应为true，false或留空以使用提供者的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK中的PutObject上传
      对象。
      
      rclone的版本<1.59使用预签名请求来上传单个
      部分对象，将此标志设置为true将重新启用该功能。
      除了特殊情况或测试之外，不应需要此功能。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示在指定时间点的文件版本。
      
      参数应为日期，"2006-01-02"，日期时间"2006-01-02
      15:04:05"或从很久以前开始的持续时间，例如"100d"或"1h"。
      
      请注意，使用此选项时，不允许进行文件写操作，
      因此无法上传文件或删除文件。
      
      有关有效格式的格式，请参见时间选项文档。
      

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置上传对象到S3。
      通常，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在收到带有
      "Content-Encoding: gzip"的文件时对其进行解压缩。
      这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip编码，请设置此标志。
      
      通常，提供商不会在下载时更改对象。如果
      没有使用`Content-Encoding: gzip`上传对象，则不会设置在下载时。
      
      但是，即使没有使用`Content-Encoding: gzip`上传对象（例如Cloudflare），
      一些提供商可能会对对象进行gzip压缩。
      
      这样做的一个症状是收到以下错误
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载具有
      设置为Content-Encoding: gzip和分块传输编码的对象，
      则rclone将动态解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商设置选择
      要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的Canned ACL。[$ACL]
   --endpoint value             S3 API的端点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2 / ECS元数据，如果没有环境变量）。 (默认值：false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置限制 - 必须设置为匹配的区域。[$LOCATION_CONSTRAINT]
   --region value               连接的区域。[$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               创建存储桶时使用的Canned ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列举块的大小（每个ListObject S3请求的响应列表）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0以进行自动选择。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         单个分块上传中的最大块数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将定期刷新的时间间隔。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip编码，请设置此标志。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，不要尝试检查bucket是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，将不会对已上传的对象进行HEAD以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，获取对象之前不执行HEAD请求。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值：false) [$V2_AUTH]
   --version-at value               显示在指定时间点的文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值：false) [$VERSIONS]

```
{% endcode %}