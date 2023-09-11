# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 lyvecloud - Seagate Lyve Cloud

用法：
   singularity storage update s3 lyvecloud [command options] <name|id>

说明：
   --env-auth
      从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据获取）。
      
      访问密钥ID和秘密访问密钥为空时才适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）中获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --region
      要连接的区域。
      
      如果使用S3克隆且没有区域，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    |例如，Jewel/v10 CEPH之前。

   --endpoint
      S3 API的端点。
      
      使用S3克隆时需要。

      示例：
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud美东1（弗吉尼亚）
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud美西1（加利福尼亚）
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1（新加坡）

   --location-constraint
      位置约束-必须设置为与区域相匹配。
      
      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶和存储或复制对象时使用的预设ACL。
      
      该ACL用于创建对象，如果bucket_acl未设置，则也用于创建存储桶。
      
      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，在服务器端复制对象时，S3不会复制源对象的ACL，而是写入新的ACL。
      
      如果ACL为空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认（私有）。

   --bucket-acl
      创建存储桶时使用的预设ACL。
      
      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。 如果未设置，则使用"acl"。
      
      如果"acl"和"bucket_acl"都是空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认（私有）。

      示例：
         | private            | 所有者具有FULL_CONTROL。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 所有者具有FULL_CONTROL。
         |                    | AllUsers组具有读取访问权限。
         | public-read-write  | 所有者具有FULL_CONTROL。
         |                    | AllUsers组具有读写访问权限。
         |                    | 不建议在存储桶上授予此权限。
         | authenticated-read | 所有者具有FULL_CONTROL。
         |                    | AuthenticatedUsers组具有读取访问权限。

   --upload-cutoff
      切换到分块上传的截止点。
      
      任何大于此大小的文件将以chunk_size大小的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如，从"rclone rcat"或使用"rclone mount"或google照片或google文档上传的文件），将使用该分块大小进行分块上传。
      
      请注意，每个传输在内存中缓冲大小为"--s3-upload-concurrency"的块。

      如果您通过高速链路传输大文件且具有足够的内存，则增加此值将加快传输速度。

      Rclone将根据需要提高分块大小，以确保文件大小低于10,000个分块的限制。

      未知大小的文件使用配置的chunk_size进行上传。由于默认的分块大小为5 MiB，并且最多可以有10,000个分块，这意味着默认情况下您可以通过流式上传的文件的最大大小为48 GiB。 如果要上传更大的文件，则需要增加chunk_size。

      增加分块大小会降低使用"-P"标志显示的进度统计的准确性。 Rclone在将块发送给AWS SDK缓冲时将块视为已发送，而实际上可能仍在上传。 更大的块大小意味着更大的AWS SDK缓冲区和与真实情况更偏离的进度报告。

   --max-upload-parts
      分块上传中的最大部件数。
      
      此选项定义进行分块上传时要使用的最大分块数。
      
      如果某个服务不支持AWS S3规范的10,000个分块，则可以使用此选项。
      
      Rclone将根据需要提高分块大小，以确保文件大小低于此分块数的限制。

   --copy-cutoff
      切换到分块复制的截止点。
      
      任何大于此大小的需要进行服务器端复制的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和与对象元数据一起存储。
      
      通常，rclone将在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这在进行数据完整性检查时非常有用，但是对于大文件来说可能会导致开始上传的长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"如果该环境变量也未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传相同文件的块数。

      如果您通过高速链路上传较少数量的大文件，并且这些上传没有充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置，则rclone将使用v2身份验证。
      
      仅在v4签名不起作用时使用此选项，例如在Jewel/v10 CEPH之前。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范的"MaxKeys"、“max-items”或“page-size”。大多数服务即使请求多个对象，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过"rgw list buckets max chunk"选项增加此值。

   --list-version
      使用的ListObjects版本：1、2或0用于自动设置。
      
      当S3首次发布时，只提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。 这具有更高的性能，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据所设置的提供程序猜测调用哪个列表对象方法。如果猜测错误，则可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供程序支持URL编码的列表，在使用控制字符的文件名时，比起其他方法更可靠。如果设置为unset（默认值），则rclone将根据所设置的提供程序选择要应用的内容，但是您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如设置，不会尝试检查存储桶是否存在或创建。
      
      如果您知道存储桶已经存在，那么这可能会有助于最少化rclone所执行的事务数量。
      
      如果正在使用的用户没有创建存储桶的权限，则可能也需要设置这个选项。在v1.52.0之前，由于错误，此选项将悄悄地传递。

   --no-head
      如果设置，不会HEAD已上传的对象以检查完整性。
      
      如果rclone在PUT之后收到200 OK消息，则会假设它已正确上传。
      
      特别是它将假设：
      
      - 元数据，包括修改时间，存储类别和内容类型与上传的一样
      - 大小与上传的一样
      
      它从单个部分PUT的响应中读取以下项：
      
      - The MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传一个未知长度的源对象，则rclone**将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的可能性，特别是大小不正确，因此不推荐在正常操作中使用。实际上，即使在设置此标志的情况下，未检测到的上传故障的机会非常小。

   --no-head-object
      如果设置，则在获取对象时不会在GET之前执行HEAD。

   --encoding
      后端的编码。
      
      有关更多信息，请参见概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池将刷新的频率。
      
      需要额外缓冲区的上传（例如分块）将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      当前s3（特别是minio）后端存在问题解决不了和HTTP/2。 S3后端默认启用HTTP/2，但可以在此禁用。问题解决后，此标志将被删除。
      
      参考：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据提供更便宜的出口流量。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应该是true，false或留空以使用提供程序的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK中的PutObject来上传对象。
      
      rclone的版本<1.59使用预签名请求来上传单个部分的对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或进行测试，否则不应该必要。

   --versions
      在目录列表中包括旧版本。

   --version-at
      按指定时间显示文件版本。
      
      参数应为日期“2006-01-02”，日期时间“2006-01-02 15:04:05”或距离那时的持续时间，例如“100d”或“1h”。
      
      请注意，使用此选项时不允许执行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如设置，则将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"将对象上传到S3。通常情况下，rclone将将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在接收到带有"Content-Encoding: gzip"的文件时对其进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能gzip对象，则设置此标志。
      
      通常，当下载对象时，提供程序不会更改对象。如果对象未使用“Content-Encoding: gzip”上传，则在下载时不会设置它。
      
      但是，即使对象未使用“Content-Encoding: gzip”上传（例如Cloudflare），某些提供程序可能会对对象进行gzip。
      
      这种情况的症状将是收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载带有Content-Encoding: gzip和分块传输编码的对象，则rclone将在传输过程中即时解压缩对象。
      
      如果设置为unset（默认值），则rclone将根据所设置的提供程序选择要应用的内容，但是您可以在此处覆盖rclone的选择。

   --no-system-metadata
      抑制设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。 [$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预设ACL。 [$ACL]
   --endpoint value             S3 API的端点。 [$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据获取）。 （默认值：false） [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  位置约束-必须设置为与区域相匹配。 [$LOCATION_CONSTRAINT]
   --region value               要连接的区域。 [$REGION]
   --secret-access-key value    AWS秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

   高级
  
   --bucket-acl value               创建存储桶时使用的预设ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 （默认值：“5Mi”） [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止点。 （默认值：“4.656Gi”） [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 （默认值：false） [$DECOMPRESS]
   --disable-checksum               不将MD5校验和与对象元数据一起存储。 （默认值：false） [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false） [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码。 （默认值：“Slash,InvalidUtf8,Dot”） [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true） [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000） [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值：“unset”） [$LIST_URL_ENCODE]
   --list-version value             使用的ListObjects版本：1、2或0用于自动设置。 （默认值：0） [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大部件数。 （默认值：10000） [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池将刷新的频率。 （默认值：“1m0s”） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，则设置此标志。 （默认值：“unset”） [$MIGHT_GZIP]
   --no-check-bucket                如设置，不会尝试检查存储桶是否存在或创建。 （默认值：false） [$NO_CHECK_BUCKET]
   --no-head                        如果设置，不会HEAD已上传的对象以检查完整性。 （默认值：false） [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不会在GET之前执行HEAD。 （默认值：false） [$NO_HEAD_OBJECT]
   --no-system-metadata             抑制设置和读取系统元数据（默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。 （默认值：“200Mi”） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 （默认值：“unset”） [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传（默认值：false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false） [$V2_AUTH]
   --version-at value               按指定时间显示文件版本。 （默认值：“off”） [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 （默认值：false） [$VERSIONS]

```
{% endcode %}