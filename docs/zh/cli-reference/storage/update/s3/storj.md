# Storj (S3兼容网关)

{% code fullWidth="true" %}
```
名称：
   singularity存储更新S3 storj - Storj（S3兼容网关）

用法：
   singularity存储更新S3 storj [命令选项] <名称|ID>

说明：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时才适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --endpoint
      Storj Gateway的端点。

      示例：
         | gateway.storjshare.io | 全球托管网关

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅在创建存储桶时应用此ACL。 如果未设置，则使用“acl”。
      
      如果“acl”和“bucket_acl”是空字符串，则不会添加X-Amz-Acl: header，并将使用默认值（private）。

      示例：
         | private            | 所有者获得FULL_CONTROL权限。
         |                    | 没有其他人有访问权（默认）。
         | public-read        | 所有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者获得FULL_CONTROL权限。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者获得FULL_CONTROL权限。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止大小。
      
      大于此大小的任何文件将按块大小进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如，来自“rclone rcat”或通过“rclone mount”或Google照片或Google文档上传的文件），它们将使用此块大小进行分块上传。
      
      请注意，"--s3-upload-concurrency"每个传输中的块大小在内存中缓冲。
      
      如果您正在高速链接上传输大型文件并且具有足够的内存，增加此值将加快传输速度。
      
      Rclone将根据已知大小的大文件而自动增加块大小，以保持在10000个块的限制以下。
      
      大小未知的文件使用配置的
      块大小上传。 由于默认的块大小为5 MiB，并且最多可以有
      10,000个块，这意味着默认情况下可以流式传输的文件的最大大小为
      48 GiB。 如果您希望流式传输
      更大的文件，那么您需要增加块大小。
      
      增加块大小会降低通过“-P”标志显示的进度
      统计的准确性。 Rclone将块视为已发送时
      它被AWS SDK缓冲，而实际上它可能仍在上传。
      较大的块大小意味着更大的AWS SDK缓冲区和进度
      报告与实际情况更偏离。
      

   --max-upload-parts
      多部分上传中的最多部分数。
      
      此选项定义在执行多部分上传时要使用的最多块数。
      
      当上传已知大小的大文件时，Rclone将根据此块数下限自动增加块大小。
      

   --copy-cutoff
      切换到分块复制的截止大小。
      
      需要进行服务器端复制的大于此大小的任何文件将按此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和存储到对象元数据中。
      
      通常，Rclone会在上传之前计算输入的MD5校验和，
      以便将其添加到对象元数据中。 这对于数据完整性检查很有用，
      但对于大文件要开始上传可能会导致长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则Rclone可以使用共享凭证文件。
      
      如果此变量为空，则Rclone将查找
      “AWS_SHARED_CREDENTIALS_FILE” env变量。 如果env变量为空，
      则它将默认为当前用户的主目录。
      
          Linux/OSX: “$HOME/.aws/credentials”
          Windows: “%USERPROFILE%\.aws\credentials”
      

   --profile
      在共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则Rclone可以使用共享凭证文件。 此
      变量用于控制在文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”或
      如果也没有设置该环境变量，则为“default”。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      同一文件的多个块将同时上传。
      
      如果您正在通过高速链接上传少量大文件
      并且这些上传没有充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。
      
      如果为true（默认值），则Rclone将使用路径样式访问；
      如果为false，则Rclone将使用虚拟路径样式。 请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      了解更多信息。
      
      某些提供商（例如AWS，Aliyun OSS，Netease COS或Tencent COS）要求此设置为
      false  - Rclone会根据提供商自动执行此操作
      设置。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则Rclone将使用v4身份验证。
      如果设置了它，Rclone将使用v2身份验证。
      
      仅当v4签名无效时才使用此选项，例如在Jewel/v10 CEPH之前。

   --list-chunk
      列出单个ListObject S3请求的列表块大小（响应列表）。
      
      此选项也称为“AWS S3规范”中的“MaxKeys”，“max-items”或“page-size”。
      大多数服务即使请求了超过1000个对象，它们也会截断响应列表。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      当S3最初推出时，它只提供了ListObjects调用
      用于列举存储桶中的对象。
      
      然而，在2016年5月，ListObjectsV2调用被引入。 这是
      性能更高，应尽可能使用。
   如果设置为默认值0，Rclone将根据提供者猜测调用
      使用哪种列举对象的方法。 如果猜错，则可能会手动设置此选项。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，则在使用控制字符的文件名时，这更可靠。
      如果设置为unset（默认设置），则Rclone将选择根据提供者设置来决定如何应用，但是您可以在此处覆盖Rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，通过这种方式可以尝试将Rclone执行的事务数量减至最小。
      
      如果使用的用户没有创建存储桶的权限，则也可能需要此选项。 在v1.52.0之前，由于错误，此操作将静默通过。
      

   --no-head
      如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。
      
      如果尝试将Rclone执行的事务数量减到最小，那么这可能很有用。
      
      设置后，如果Rclone在使用PUT上传对象后收到200 OK消息，则它将假设它已正确上传。
      
      具体来说，它将假设：
      
      - 元数据（包括修改时间，存储类别和内容类型）与已上传的文件相同
      - 大小与上传的大小相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象长度未知，则Rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的几率，
      特别是大小不正确，因此不建议正常操作中使用。 在实践中，即使使用此标志，检测不到的上传失败的机会也很小。
      

   --no-head-object
      如果设定，则在获取对象之前不执行HEAD。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区的上传（例如多部分上传）将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个未解决的http2问题。
      S3后端默认情况下启用了HTTP/2，但可以在此禁用。 当问题解决时，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      通常，将其设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的数据的更便宜的出站流量。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应为true、false或未设置以使用提供者的默认设置。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则Rclone将使用AWS SDK的PutObject上传对象。
      
      Rclone的版本 < 1.59使用预签名请求上传单个
      部分对象，将此标志设置为true将重新启用该功能。
      除非特殊情况或测试，否则不应此必要。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      按指定时间显示文件版本。
      
      参数应为日期，“2006-01-02”，datetime“2006-01-02
      15:04:05”或其之前的持续时间，例如“100d”或“1h”。
      
      请注意，使用此选项时不允许写入文件操作，
      因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设定，将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”将对象上传到S3中。
      通常，Rclone将下载这些文件作为压缩对象。
      
      如果设置了此标志，则Rclone将在接收到这些文件时使用“Content-Encoding: gzip”对其进行解压缩。 这意味着Rclone无法检查文件大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能会gzip对象，则设置此标志。
      
      通常，提供商在下载对象时不会更改对象。 如果
      对象没有用`Content-Encoding: gzip`上传，则在下载时它不会被设置。
      
      但是，即使没有用`Content-Encoding: gzip`上传，
      某些提供商可能会对对象进行gzip压缩（例如Cloudflare）。
      
      这种情况的症状将是收到类似的错误
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且Rclone下载了带有
      设置为Content-Encoding: gzip和chunked传输编码的对象，则Rclone将动态对对象进行解压缩。
      
      如果设置为unset（默认设置），则Rclone将选择
      根据提供者设置来决定如何应用，但是您可以在此处覆盖
      Rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value      AWS访问密钥ID。 [$ACCESS_KEY_ID]
   --endpoint value           Storj Gateway的端点。 [$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  AWS Secret Access Key（密码）。 [$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               创建存储桶时使用的预定义ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止大小。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设定，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和存储到对象元数据中。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出单个ListObject S3请求的列表块大小。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最多部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此标志，如果后端可能会gzip对象。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设定，则在获取对象之前不执行HEAD。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  在共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止大小。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               按指定时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}