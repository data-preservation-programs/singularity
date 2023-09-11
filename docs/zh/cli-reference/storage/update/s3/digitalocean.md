# DigitalOcean Spaces

{% code fullWidth="true" %}
```
名称:
   singularity storage update s3 digitalocean - DigitalOcean空间

用法:
   singularity storage update s3 digitalocean [命令选项] <名称|ID>

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅在access_key_id和secret_access_key为空时适用。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问凭证ID。

      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问凭证（密码）。

      留空以进行匿名访问或运行时凭证。

   --region
      要连接的区域。

      如果您使用的是S3克隆，并且没有区域，请留空。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空的区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如，在Jewel/v10 CEPH之前。

   --endpoint
      S3 API的终端节点。

      使用S3克隆时必填。

      示例:
         | syd1.digitaloceanspaces.com | DigitalOcean Sydney 1 Spaces
         | sfo3.digitaloceanspaces.com | DigitalOcean San Francisco 3 Spaces
         | fra1.digitaloceanspaces.com | DigitalOcean Frankfurt 1 Spaces
         | nyc3.digitaloceanspaces.com | DigitalOcean New York 3 Spaces
         | ams3.digitaloceanspaces.com | DigitalOcean Amsterdam 3 Spaces
         | sgp1.digitaloceanspaces.com | DigitalOcean Singapore 1 Spaces

   --location-constraint
      位置约束 - 必须设置为与区域匹配。

      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶、存储或复制对象时使用的现有权限控制策略(ACL)。

      此ACL用于创建对象，如果未设置bucket_acl，则也用于创建存储桶。

      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，当对S3进行服务器端复制时，此ACL会被应用，因为S3不会复制源的ACL，而是写入一个新的ACL。

      如果acl为空字符串，则不添加X-Amz-Acl:头，将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的现有权限控制策略(ACL)。

      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"代替。

      如果"acl"和"bucket_acl"都是空字符串，则不添加X-Amz-Acl:头，并且将使用默认值（private）。

      示例:
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 其他人没有访问权限（默认值）。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | 所有用户组具有READ权限。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | 所有用户组具有READ和WRITE权限。
         |                    | 通常不推荐在存储桶上进行设置。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers用户组具有READ权限。

   --upload-cutoff
      切换到分块上传的文件截止点。

      任何大于此大小的文件将以chunk_size的分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。

      对于大于upload_cutoff的文件或大小未知的文件（例如，来自"rclone rcat"、通过"rclone mount"或Google照片或Google文档上传的文件），将使用此分块大小进行分块上传。

      请注意，每个传输会在内存中缓冲"--s3-upload-concurrency"大小的块。
      
      如果您正在通过高速链接传输大文件并且具有足够的内存，增加此值将加快传输速度。

      Rclone将根据需要自动增加分块大小，以保持在10,000个块的限制之下。

      未知大小的文件将以配置的chunk_size上传。由于默认的chunk_size为5 MiB，并且最多可以有10,000个块，因此，默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加chunk_size。

      增加chunk_size会减少使用"-P"标志显示的进度统计的准确性。当AWS SDK将缓冲的块视为已发送时，Rclone将其视为已发送，而实际上它可能仍在上传。更大的chunk_size意味着更大的AWS SDK缓冲区和与真实情况相去甚远的进度报告。

   --max-upload-parts
      分块上传中的最大块数。

      此选项定义分块上传时使用的最大多块数量。

      如果某个服务不支持AWS S3规范的10,000个多块，则这对您很有用。

      Rclone将根据需要自动增加分块大小，以保持在此多块数量的限制之下。

   --copy-cutoff
      切换到分块复制的文件截止点。

      任何需要进行服务器端复制的大于此大小的文件将以此大小的分块复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。

      通常情况下，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但会导致大文件上传的长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth = true，则rclone可以使用共享凭证文件。

      如果这个变量为空，rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果env的值为空，它将默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭证文件。此变量用于控制在该文件中使用哪个配置文件。

      如果为空，则默认为"AWS_PROFILE"环境变量或"default"（如果该环境变量也未设置）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传同一文件的块数。

      如果您正在通过高速链接上传少量大文件，并且这些上传未完全利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，使用路径样式访问; 如果为false，使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问; 如果为false，则rclone将使用虚拟路径样式。有关详细信息，请参阅[the AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求将其设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      默认情况下，rclone将使用v4身份验证。如果设置了此选项，则rclone将使用v2身份验证。

      仅在v4签名无法工作时使用，例如，在Jewel/v10 CEPH之前。

   --list-chunk
      列出分页的大小（每个ListObject S3请求的响应列表的大小）。

      此选项也被称为"MaxKeys"、"max-items"或"page-size"，来自AWS S3规范。
      大多数服务将响应列表截断为1000个对象，即使请求的更多。

      在AWS S3中，这是全局最大值，无法更改，详见[Amazon S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。

      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。

      当S3最初发布时，它仅提供ListObjects调用以枚举存储桶中的对象。

      然而，在2016年5月，引入了ListObjectsV2调用。这是一个更高性能的调用，如果有可能，应该使用它。

      如果设置为默认值0，rclone将根据设置的提供商猜测要调用的ListObjects方法。如果它猜错了，则可以在此手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      一些提供商支持URL编码清单，如果可用，则在文件名中使用控制字符时，此方法更可靠。如果将其设置为unset（默认值），则rclone将根据提供商设置选择应用何种编码，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查桶是否存在或创建它。

      如果您知道桶已经存在，这对于尽量减少rclone执行的事务数很有用。

      如果使用的用户没有桶创建权限，则也可能需要此选项。v1.52.0之前的版本由于错误而会默默地执行此操作。

   --no-head
      如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。

      这对于尽量减少rclone执行的事务数很有用。

      设置此选项意味着如果rclone在PUT之后接收到200 OK消息，那么它将假设已正确上传。

      特别是，它将假设：

      - 元数据，包括修改时间、存储类和内容类型与上传时相同。
      - 大小与上传时相同。

      它从单个部分PUT的响应中读取以下条目：

      - MD5SUM
      - 上传日期

      对于多部分上传，不读取这些条目。

      如果上传了大小未知的源对象，则rclone**将**执行HEAD请求。

      设置此标志增加了未检测到的上传失败的机会，特别是错误的大小，因此不推荐在正常操作中使用。实际上，即使在设置此标志的情况下，未检测到的上传失败的概率非常小。

   --no-head-object
      如果设置，则获取对象之前不执行HEAD请求。

   --encoding
      后端的编码。

      有关更多信息，请参见概述中的[encoding部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的时间间隔。

      需要额外缓冲区（例如，多部分）的上传将使用内存池进行分配。
      此选项控制多久内未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前，s3（特别是minio）后端存在一个未解决的问题，与HTTP/2有关。默认情况下，s3后端启用HTTP/2，但可以在此禁用。当问题解决后，此标志将被删除。

      请参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      自定义下载的终端节点。
      通常将其设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的数据的更便宜的出口。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证。

      这应该设置为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单一部分上传。

      如果设置为false，rclone将使用AWS SDK中的PutObject来上传对象。

      rclone的版本1.59以前使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除了特殊情况或测试外，不应该使用此选项。

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示在指定时间点的文件版本。

      参数应为日期（"2006-01-02"）、日期时间（"2006-01-02 15:04:05"）或从当前时间点开始推算的持续时间，例如"100d"或"1h"。

      请注意，使用此选项时，不允许进行文件写操作，因此无法上传文件或删除文件。

      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。

      可以使用"Content-Encoding: gzip"将对象上传到S3。通常，rclone会将这些文件作为压缩对象进行下载。

      如果设置了此标志，则rclone在接收到带有"Content-Encoding: gzip"的文件时会解压缩它们。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会对对象进行gzip压缩，请设置此标志。

      通常情况下，提供商在下载对象时不会更改对象。如果对象没有使用"Content-Encoding: gzip"进行上传，则下载对象时不会设置该标志。

      然而，一些提供商可能会对对象进行gzip压缩，即使它们没有使用"Content-Encoding: gzip"进行上传（例如云帆）。

      如果设置了此标志，并且rclone下载了带有设置了"Content-Encoding: gzip"和分块传输编码的对象，rclone将在接收到对象时即时对其进行解压缩。

      如果设置为unset（默认值），则rclone将根据提供商设置选择应用何种压缩，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id value        AWS访问凭证ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的现有ACL。[$ACL]
   --endpoint value             S3 API的终端节点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。默认值：false [$ENV_AUTH]
   --help, -h                   显示帮助信息
   --location-constraint value  位置约束 - 必须设置为与区域匹配。[$LOCATION_CONSTRAINT]
   --region value               要连接的区域。[$REGION]
   --secret-access-key value    AWS秘密访问凭证（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的现有ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的文件截止点。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             自定义下载的终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，使用路径样式访问; 如果为false，使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出分页的大小（每个ListObject S3请求的响应列表的大小）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1,2或0表示自动。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大块数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池刷新的时间间隔。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会对对象进行gzip压缩，请设置此标志。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查桶是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则获取对象之前不执行HEAD请求。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            转换为分块上传的文件截止点。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单一部分上传 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值：false) [$V2_AUTH]
   --version-at value               显示在指定时间点的文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值：false) [$VERSIONS]

```
{% endcode %}