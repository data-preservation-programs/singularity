# 华为对象存储服务

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 huaweiobs - 华为对象存储服务

用法：
   singularity storage update s3 huaweiobs [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时有效。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）中获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      留空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      留空表示匿名访问或运行时凭证。

   --region
      要连接的地区-创建桶并存储或拷贝对象的位置。需要与您的终端节点相同。
      

      示例：
         | af-south-1     | AF-Johannesburg
         | ap-southeast-2 | AP-Bangkok
         | ap-southeast-3 | AP-Singapore
         | cn-east-3      | CN East-Shanghai1
         | cn-east-2      | CN East-Shanghai2
         | cn-north-1     | CN North-Beijing1
         | cn-north-4     | CN North-Beijing4
         | cn-south-1     | CN South-Guangzhou
         | ap-southeast-1 | CN-Hong Kong
         | sa-argentina-1 | LA-Buenos Aires1
         | sa-peru-1      | LA-Lima1
         | na-mexico-1    | LA-Mexico City1
         | sa-chile-1     | LA-Santiago2
         | sa-brazil-1    | LA-Sao Paulo1
         | ru-northwest-2 | RU-Moscow2

   --endpoint
      OBS API的终端节点。

      示例：
         | obs.af-south-1.myhuaweicloud.com     | AF-Johannesburg
         | obs.ap-southeast-2.myhuaweicloud.com | AP-Bangkok
         | obs.ap-southeast-3.myhuaweicloud.com | AP-Singapore
         | obs.cn-east-3.myhuaweicloud.com      | CN East-Shanghai1
         | obs.cn-east-2.myhuaweicloud.com      | CN East-Shanghai2
         | obs.cn-north-1.myhuaweicloud.com     | CN North-Beijing1
         | obs.cn-north-4.myhuaweicloud.com     | CN North-Beijing4
         | obs.cn-south-1.myhuaweicloud.com     | CN South-Guangzhou
         | obs.ap-southeast-1.myhuaweicloud.com | CN-Hong Kong
         | obs.sa-argentina-1.myhuaweicloud.com | LA-Buenos Aires1
         | obs.sa-peru-1.myhuaweicloud.com      | LA-Lima1
         | obs.na-mexico-1.myhuaweicloud.com    | LA-Mexico City1
         | obs.sa-chile-1.myhuaweicloud.com     | LA-Santiago2
         | obs.sa-brazil-1.myhuaweicloud.com    | LA-Sao Paulo1
         | obs.ru-northwest-2.myhuaweicloud.com | RU-Moscow2

   --acl
      创建桶和存储或拷贝对象时使用的预设访问控制列表（ACL）。
      
      此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建桶。
      
      有关更多信息，请参见https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，会应用此ACL，因为S3不会复制源中的ACL，而是写入新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl：标头，并使用默认值（私有）。

   --bucket-acl
      创建桶时使用的预设访问控制列表（ACL）。
      
      有关更多信息，请参见https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      如果未设置此项，则使用"acl"。
      
      如果"acl"和"bucket_acl"均为empty字符串，则不会添加X-Amz-Acl：标头，并使用默认值（私有）。

      示例：
         | private            | 拥有者获得FULL_CONTROL。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 拥有者获得FULL_CONTROL。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 拥有者获得FULL_CONTROL。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不建议在桶上授予此权限。
         | authenticated-read | 拥有者获得FULL_CONTROL。
         |                    | AuthenticatedUsers组获得读取权限。

   --upload-cutoff
      切换到分块上传的截止点。
      
      大于此大小的文件将按照chunk_size的大小分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小不明确的文件时（例如，来自"rclone rcat"、使用"rclone mount"上传的文件或Google照片或Google文档），将使用此分块大小进行上传。
      
      请注意，每个上传的"--s3-upload-concurrency"块大小将缓存在内存中。
      
      如果您正在通过高速链接传输大文件，并且具有足够的内存，则增加此值可以加快传输速度。
      
      当上传已知大小的大文件时，rclone将自动增加分块大小，以保持在10,000个块的限制以下。
      
      如果上传大小未知，则使用配置的分块大小上传。
      由于默认的分块大小为5 MiB，并且最多可以有10,000个块，因此默认情况下，您可以流式传输的文件的最大大小为48 GiB。如果要流式上传更大的文件，则需要增加chunk_size的值。
      
      增加分块大小会降低"-P"标志显示的进度统计数据的准确性。当块由AWS SDK缓冲时，rclone将其视为已发送的块，而实际上可能仍在上载。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况偏离更大。
      

   --max-upload-parts
      多部分上传中的最大块数。
      
      此选项定义在执行多部分上传时使用的最大多部分块数。
      
      如果某个服务不支持AWS S3的10,000个块规范，则可以使用。
      
      rclone将自动增加分块的大小以保持在此块数的限制以下，以在上传已知大小的大文件时使用。
      

   --copy-cutoff
      切换到分块复制的截止点。
      
      需要以此大小的块复制大于此大小的文件。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常情况下，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说，可能会导致长时间的上传延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享的凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量的值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      在共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享的凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果未设置该环境变量）。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是并发上传的相同文件块的数量。
      
      如果您通过高速链接上传较少数量的大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参见[rclone使用桶](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      一些提供商（例如，AWS、阿里云OSS、网易COS或腾讯COS）要求将此设置为false-rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，则rclone将使用v2身份验证。
      
      仅当v4签名无法使用（例如，Jewel/v10 CEPH之前）时才使用此选项。

   --list-chunk
      列举块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为"MaxKeys"、"max-items"或"page-size"，源自AWS S3规范。
      大多数服务即使请求超过1000个对象，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，详见[AWS S3](https://docs.aws.amazon.com/zh_cn/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动选择。
      
      当S3最初启动时，它仅提供了ListObjects调用以列举存储桶中的对象。
      
      但是在2016年5月，引入了ListObjectsV2调用。这是更高性能的方法，如果可能，请使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供商来猜测调用哪种列表对象方法。如果它猜错了，那么可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果提供了这个选项，使用控制字符的文件名时，它更可靠。如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但你可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已经存在，那么这可能对试图将rclone执行的事务数最小化非常有用。
      
      如果没有存储桶创建权限的用户使用，也可能需要此选项。在v1.52.0之前，由于一个bug，这个选项会产生静默效果。

   --no-head
      如果设置，则不通过HEAD请求检查已上传对象的完整性。
      
      如果想要最小化rclone执行的事务数，那么这可能很有用。
      
      设置它意味着如果rclone在使用PUT上传对象后收到200 OK消息，则它将假设它已经正确上传。
      
      特别地，它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传一致
      - 大小与上传一致
      
      对于单部分PUT的响应，rclone读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不读取这些内容。
      
      如果上传源对象的长度未知，则rclone**会**执行HEAD请求。
      
      设置此标志会增加无法检测的上传故障的几率，特别是大小不正确的情况，因此不推荐在正常操作中使用。实际上，即使使用此标志，检测到不正确的上传故障的几率也非常小。
      

   --no-head-object
      如果设置，则在获取对象时不进行HEAD请求。

   --encoding
      后端的编码。
      
      请参阅[概览中的编码部分](/overview/#encoding)了解更多信息。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。
      
      需要额外缓冲区的上传（例如分块）将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中移除。

   --memory-pool-use-mmap
      是否在内部内存缓冲区池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP/2使用。
      
      目前s3（特别是minio）后端和HTTP/2之间存在未解决的问题。对于s3后端，默认启用HTTP/2，但可以在此处禁用。解决此问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据提供更便宜的外传。

   --use-multipart-etag
      是否在multipart上传中使用ETag进行验证
      
      应设置为true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      rclone的版本<1.59使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或者用于测试，否则不应该需要打开此选项。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      指定指定时间点的文件版本。
      
      参数应为日期，"2006-01-02"，日期时间"2006-01-02 15:04:05"或表示相隔多久的持续时间，例如"100d"或"1h"。
      
      请注意，在使用此选项时，不允许执行任何文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[time选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则会解压缩gzip编码的对象。
      
      可以使用"Content-Encoding：gzip"在S3上上传对象。通常情况下，rclone会将这些文件作为压缩的对象下载。
      
      如果设置了此标志，则rclone将在接收到具有"Content-Encoding：gzip"的文件时进行解压缩。这意味着rclone无法检查大小和哈希，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能会gzip压缩对象，请设置此标志。
      
      通常，提供程序在下载时不会更改对象。如果对象未使用"Content-Encoding：gzip"上传，那么在下载时它不会设置。
      
      但是，某些提供商即使在上传时没有使用"Content-Encoding：gzip"（例如Cloudflare），也可能gzip压缩对象。
      
      这种情况的症状可能是接收到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载带有Content-Encoding：gzip和块传输编码的对象，则rclone将动态解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建桶和存储或拷贝对象时使用的预设访问控制列表（ACL）。[$ACL]
   --endpoint value           OBS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --region value             要连接的地区-创建桶并存储或拷贝对象的位置。需要与您的终端节点相同。[$REGION]
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建桶时使用的预设访问控制列表（ACL）。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止点。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则会解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列举块的大小（每个ListObject S3请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0表示自动选择。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大块数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲区池中使用mmap缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会gzip压缩对象，请设置此标志。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不通过HEAD请求检查已上传对象的完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不进行HEAD请求。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  在共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在multipart上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (default: false) [$V2_AUTH]
   --version-at value               指定指定时间点的文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (default: false) [$VERSIONS]

```
{% endcode %}