# 华为对象存储服务

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 huaweiobs - 华为对象存储服务

用法：
   singularity storage update s3 huaweiobs [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据获取）。
      
      仅在访问密钥ID和秘密访问密钥为空时适用。

      示例：
         | false | 在下一步输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      如果需要匿名访问或运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      如果需要匿名访问或运行时凭证，请留空。

   --region
      要连接的区域 - 存储桶将被创建并存储数据的位置。需要与您的终端节点相同。
      

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
      创建存储桶并存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，并且如果未设置bucket_acl，则也用于创建存储桶。
      
      有关更多信息，请访问[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      请注意，在服务器端复制对象时将应用此ACL，因为S3不会复制源的ACL，而是写入一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl: header，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。
      
      如果"acl"和"bucket_acl"是空字符串，则不会添加X-Amz-Acl: header，并且将使用默认值（private）。

      示例：
         | private            | 拥有者拥有FULL_CONTROL权限。
         |                    | 没有其他人有访问权限（默认）。
         | public-read        | 拥有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 拥有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      切换到分块上传的阈值。
      
      大于此大小的文件将以分块大小上传。最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如从"rclone rcat"或使用"rclone mount"或google照片或google文档上传的文件）时，将使用此分块大小进行分块上传。
      
      请注意，每次传输内存中缓冲的是--s3-upload-concurrency个此大小的块。
      
      如果您正在使用高速链接传输大文件且具有足够的内存，则增加此值将加快传输速度。
      
      rclone会自动增加分块大小，以确保在低于10000的块数限制下上传已知大小的大文件。
      
      未知大小的文件将使用配置的分块大小进行上传。由于默认分块大小为5 MiB，并且最多可以有10000个块，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加分块大小。
      
      增加分块大小将降低使用"-P"标志显示的进度统计数据的准确性。当rclone将块缓冲到AWS SDK时，rclone会认为块已发送，而实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和与实际情况偏离更大的进度报告。

   --max-upload-parts
      分块上传的最大部件数。
      
      此选项定义了执行分块上传时使用的最大分块数。
      
      如果某个服务不支持AWS S3规范的10000个块限制，则可以使用此选项。
      
      当上传已知大小的大文件时，rclone会自动增加分块大小，以便保持低于此块数限制。
      

   --copy-cutoff
      切换到分块复制的阈值。
      
      需要服务器端复制的大于此大小的文件将按此大小分块复制。
      
      最小为0，最大为5 GiB。

   --disable-checksum
      不要将MD5校验和存储在对象元数据中。
      
      通常，rclone将在上传之前计算输入的MD5校验和，并将其添加到对象的元数据中。这对于数据完整性检查很有用，但对于大文件启动上传可能会导致长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"默认"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的相同文件块的数量。
      
      如果在高速链接上传少量大文件且这些上传未能充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式。有关更多信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将其设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了此值，则rclone将使用v2身份验证。
      
      只有在v4签名无法工作时（例如v10 CEPH之前的Jewel版本）才使用此选项。

   --list-chunk
      列出块的大小（用于每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求大于该数量也会截断响应列表到1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

   --list-version
      要使用的ListObjects的版本：1、2或0（自动）。
      
      当S3最初发布时，它只提供了ListObjects调用以枚举桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。这个调用的性能要高得多，如果可能的话应该使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个列出对象的方法。如果猜测错误，则可以在此处手动设置。

      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，则在使用控制字符的文件名时，这种编码方式更可靠。
      如果设置为未设置（默认值），则rclone将根据提供商设置选择应用什么，但可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，不要尝试检查存储桶是否存在或创建存储桶。
      
      如果要尽量减少rclone执行的事务数量，或者已知存储桶已存在，则这将非常有用。

   --no-head
      如果设置，不要对上传的对象进行HEAD检查以检查完整性。
      
      这在尽量减少rclone执行的事务数量时非常有用。
      
      设置此标志意味着如果rclone在使用PUT上传对象后收到200 OK消息，则假设该对象上传成功。
      
      特别是它将假定：
      
      - 元数据，包括修改时间、存储类和内容类型与上传相同
      - 大小与上传相同
      
      它从单部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于分块上传，不会读取这些项目。
      
      如果上传的源对象大小未知，则rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确的情况，因此不建议在正常操作中使用。实际上，即使设置了此标志，未检测到的上传失败的机会也非常小。

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      进行额外缓冲（例如分块）的上传将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端和HTTP/2存在未解决的问题。默认情况下，s3后端启用了HTTP/2，但可以在此禁用。此问题解决后，将删除此标志。
      
      请参阅：[https://github.com/rclone/rclone/issues/4673](https://github.com/rclone/rclone/issues/4673)，[https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)

   --download-url
      下载的自定义终端节点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据的出口更便宜。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应该为true、false或未设置，以使用提供者的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK中的PutObject上传对象。
      
      rclone的版本<1.59使用预签名请求上传单部分对象，将此标志设置为true将重新启用该功能。除非属于特殊情况或进行测试，否则不应该需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示文件在指定时间的版本。
      
      参数应为日期（"2006-01-02"）、日期时间（"2006-01-02
      15:04:05"）或以前多长时间的持续时间，例如"100d"或"1h"。
      
      请注意，当使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置此标志，则将解压缩gzip编码的对象。
      
      可以使用设置了“Content-Encoding: gzip”的对象上传到S3。通常情况下，rclone将以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则rclone将在接收到设置了“Content-Encoding: gzip”的文件时进行解压缩。这意味着rclone无法检查大小和哈希，但文件内容将被解压缩。

   --might-gzip
      如果后端可能压缩对象，请设置此标志。
      
      通常情况下，提供商不会更改下载的对象。如果一个对象未以“Content-Encoding: gzip”上传，那么在下载时它也不会设置。
      
      然而，某些提供商即使没有使用“Content-Encoding: gzip”上传了对象（例如Cloudflare），也可能对对象进行gzip压缩。
      
      这样做将会导致接收到如下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了设置了Content-Encoding: gzip和chunked传输编码的对象，则rclone将动态地解压缩对象。
      
      如果设置为未设置（默认值），则rclone将根据提供商设置选择应用什么，但可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项：
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶并存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value           OBS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（如果没有环境变量，则从环境变量或EC2/ECS元数据获取）。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助信息
   --region value             要连接的区域 - 存储桶将被创建并存储数据的位置。需要与您的终端节点相同。[$REGION]
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   进阶

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi")[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的阈值。 (默认值: "4.656Gi")[$COPY_CUTOFF]
   --decompress                     如果设置此标志，则将解压缩gzip编码的对象。 (默认值: false)[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和存储在对象元数据中。 (默认值: false)[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false)[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot")[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true)[$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（用于每个ListObject S3请求的响应列表）。 (默认值: 1000)[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset")[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1,2或0（自动） (默认值: 0)[$LIST_VERSION]
   --max-upload-parts value         分块上传的最大部件数。 (默认值: 10000)[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s")[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false)[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请设置此标志。 (默认值: "unset")[$MIGHT_GZIP]
   --no-check-bucket                如果设置，不要尝试检查存储桶是否存在或创建存储桶。 (默认值: false)[$NO_CHECK_BUCKET]
   --no-head                        如果设置，不要对上传的对象进行HEAD检查以检查完整性。 (默认值: false)[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (默认值: false)[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false)[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4)[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的阈值。 (默认值: "200Mi")[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset")[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false)[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false)[$V2_AUTH]
   --version-at value               显示文件在指定时间的版本。 (默认值: "off")[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false)[$VERSIONS]

```
{% endcode %}