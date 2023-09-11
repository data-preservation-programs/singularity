# 阿里云对象存储系统（OSS）Aliyun

{% code fullWidth="true" %}
```
命名:
   singularity storage update s3 alibaba - 阿里云对象存储系统（OSS）Aliyun

使用方法:
   singularity storage update s3 alibaba [命令选项] <名称|ID>

说明:
   --env-auth
      从运行环境（环境变量或EC2/ECS元数据）获取AWS凭证。

      仅在access_key_id和secret_access_key为空时适用。

      示例:
         | false | 下一步输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS Access Key ID。

      留空以匿名访问或使用运行时凭证。

   --secret-access-key
      AWS Secret Access Key（密码）。

      留空以匿名访问或使用运行时凭证。

   --endpoint
      OSS API的终端节点。

      示例:
         | oss-accelerate.aliyuncs.com          | 全球加速
         | oss-accelerate-overseas.aliyuncs.com | 全球加速（中国境外）
         | oss-cn-hangzhou.aliyuncs.com         | 华东1（杭州）
         | oss-cn-shanghai.aliyuncs.com         | 华东2（上海）
         | oss-cn-qingdao.aliyuncs.com          | 华北1（青岛）
         | oss-cn-beijing.aliyuncs.com          | 华北2（北京）
         | oss-cn-zhangjiakou.aliyuncs.com      | 华北3（张家口）
         | oss-cn-huhehaote.aliyuncs.com        | 华北5（呼和浩特）
         | oss-cn-wulanchabu.aliyuncs.com       | 华北6（乌兰察布）
         | oss-cn-shenzhen.aliyuncs.com         | 华南1（深圳）
         | oss-cn-heyuan.aliyuncs.com           | 华南2（河源）
         | oss-cn-guangzhou.aliyuncs.com        | 华南3（广州）
         | oss-cn-chengdu.aliyuncs.com          | 西南1（成都）
         | oss-cn-hongkong.aliyuncs.com         | 中国香港（香港）
         | oss-us-west-1.aliyuncs.com           | 美国西部1（硅谷）
         | oss-us-east-1.aliyuncs.com           | 美国东部1（弗吉尼亚）
         | oss-ap-southeast-1.aliyuncs.com      | 东南亚1（新加坡）
         | oss-ap-southeast-2.aliyuncs.com      | 亚太东南2（悉尼）
         | oss-ap-southeast-3.aliyuncs.com      | 东南亚3（吉隆坡）
         | oss-ap-southeast-5.aliyuncs.com      | 亚太东南5（雅加达）
         | oss-ap-northeast-1.aliyuncs.com      | 亚太东北1（日本）
         | oss-ap-south-1.aliyuncs.com          | 亚太南部1（孟买）
         | oss-eu-central-1.aliyuncs.com        | 欧洲中部1（法兰克福）
         | oss-eu-west-1.aliyuncs.com           | 西欧（伦敦）
         | oss-me-east-1.aliyuncs.com           | 中东1（迪拜）

   --acl
      存储桶创建和存储或复制对象时使用的预定义ACL。

      此ACL用于创建对象，并在未设置bucket_acl时，用于创建存储桶。

      更多信息请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，此ACL用于在服务器端复制对象时应用，
      因为S3不会复制源文件的ACL，而是写入一个新的ACL。

      如果acl是空字符串，则不会添加X-Amz-Acl:头，
      并将使用默认值（私有）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。

      更多信息请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。

      如果"acl"和"bucket_acl"是空字符串，则不会添加X-Amz-Acl:
      头，并将使用默认值（私有）。

      示例:
         | private            | 拥有者完全控制权限。
         |                    | 无其他用户有访问权限（默认值）。
         | public-read        | 拥有者完全控制权限。
         |                    | 全局用户组具有读取权限。
         | public-read-write  | 拥有者完全控制权限。
         |                    | 全局用户组具有读取和写入权限。
         |                    | 这通常不推荐在存储桶上授予。
         | authenticated-read | 拥有者完全控制权限。
         |                    | 授权用户组具有读取权限。

   --storage-class
      存储新对象时要使用的存储类型。

      示例:
         | <未设置>    | 默认值
         | STANDARD   | 标准存储类型
         | GLACIER    | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换为分块上传的文件大小阈值。

      任何大于此阈值的文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。
   --chunk-size
      用于上传的分块大小。

      当上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或google
      相片或google文档上传的文件）时，将使用此块大小进行分块上传。

      请注意，"--s3-upload-concurrency"每个传输将缓冲此大小的块在内存中。

      如果您正在高速链接上传输大文件，并且您的内存足够，
      那么增加此值将加快传输速度。

      当上传已知大小的大文件时，rclone将自动增加块大小，
      以保持不超过10,000个块的限制。

      未知大小的文件将使用配置的块大小上传。
      由于默认块大小为5 MiB，最多可以有10,000个块，
      这意味着默认情况下可以流式上传的文件的最大大小为48 GiB。
      如果您希望流式上传更大的文件，则需要增加chunk_size。

      增加块大小会降低使用"-P"标志时显示的进度统计的准确性。
      当通过AWS SDK缓冲块时，rclone将块视为已发送，
      而实际上它可能仍在上传中。
      块大小越大，AWS SDK缓冲区和进度报告的准确性就越小。

   --max-upload-parts
      分块上传中的最大部分数。

      该选项定义在执行分块上传时要使用多少个分块。
      当服务不支持AWS S3规范的10,000个块时，这可能很有用。

      当上传已知大小的大文件时，rclone将自动增加块大小，
      以保持不超过此块数限制。

   --copy-cutoff
      切换为分块复制的文件大小阈值。

      任何需要服务器端复制的大于此阈值的文件将以此大小的块复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。

      通常，在上传之前，rclone会计算输入的MD5校验和，
      以便可以将其添加到对象的元数据中。这对于数据完整性检查非常好，
      但对于大文件开始上传可能会导致较长的延迟。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth = true，则rclone可以使用共享凭证文件。

      如果此变量为空，则rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE"环境变量。
      如果环境值为空，则它将默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭证文件。此
      变量控制在该文件中使用的配置文件。

      如果为空，则默认为环境变量"AWS_PROFILE"或"default"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传的相同文件的块数。

      如果您在高速链接上上传少量大文件，并且这些上传未充分利用您的带宽，
      那么增加此值可能会有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟主机样式访问。有关详细信息，
      请参阅[the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如AWS，Aliyun OSS，Netease COS或Tencent COS）要求将其设置为
      false - rclone会根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了此标志，则rclone将使用v2身份验证。

      仅在v4签名无效时使用，例如早于Jewel/v10 CEPH。

   --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。

      该选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。
      大多数服务将响应列表截断为1000个对象，即使请求超过1000个。

      在AWS S3中，这是一个全局上限，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。

      在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0为自动。

      当S3最初发布时，只提供了ListObjects调用以枚举存储桶中的对象。

      然而，在2016年5月，引入了ListObjectsV2调用。这样做
      可以提供更高性能，并且如果可能应该使用。

      如果设置为默认值0，rclone将根据提供商的设置猜测应调用哪个列表对象方法。
      如果猜测错误，则可能在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      一些提供商支持URL编码列表，若可用，则在文件名中使用控制字符时，这是更可靠的。
      如果设置为unset（默认值），则rclone将根据提供商的设置来选择如何应用，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。

      如果知道存储桶已经存在，则可以在尽量减少rclone所做的事务数量时，
      这可能很有用。

      如果使用的用户没有创建存储桶的权限，则也可能需要它。
      在v1.52.0之前，由于错误，此操作将静默通过。

   --no-head
      如果设置，则不会对已上传对象进行HEAD请求以检查完整性。

      如果试图将rclone所做的事务数量最小化，这可能很有用。

      设置后，意味着在PUT上传对象后，如果rclone收到一个200 OK的消息，
      那么它将假设对象已经正确上传。

      特别是它将假设：

      - 元数据，包括修改时间、存储类和内容类型与上传的一样。
      - 大小与上传的一样。

      它从单个部分PUT的响应中读取以下项：

      - MD5SUM
      - 上传日期

      对于分片上传，不会读取这些项。

      如果上传未知长度的源对象，则rclone**将**执行HEAD请求。

      设置此标记会增加检测不到的上传错误的机会，
      特别是错误大小，因此不推荐在正常操作中使用它。
      实际上，即使设置了此标志，出现检测不到的上传错误的机会也很小。

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码。

      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池将刷新的频率。

      需要额外缓冲区的上传（例如，多部分）将使用内存池进行分配。
      此选项控制未使用的缓冲区将从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前，s3（特别是minio）后端存在未解决的问题
      和HTTP/2。S3后端默认启用HTTP/2，但可以在此处禁用。
      当问题解决时，此标志将被删除。

      参见：
      https://github.com/rclone/rclone/issues/4673,
      https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终节点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据
      提供更便宜的出口。

   --use-multipart-etag
      是否在分块上传中使用ETag进行校验。

      这应该为true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传。

      如果这是false，则rclone将使用AWS SDK中的PutObject上传对象。

      rclone < 1.59的版本在上传单个部分对象时使用预签名请求，
      将此标志设置为true将重新启用该功能。
      除了特殊情况或测试，通常不需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。

      参数应为日期，"2006-01-02"，日期时间"2006-01-02 15:04:05"，
      或之前的时间长度，例如"100d"或"1h"。
      
      请注意，在使用此选项时，不允许进行文件写入操作，
      因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。

      可以使用"Content-Encoding: gzip"将对象上传到S3。
      通常，rclone会将这些文件下载为压缩对象。

      如果设置了此标志，rclone将在接收到带有
      "Content-Encoding: gzip"的对象时解压缩它们。
      这意味着rclone无法检查大小和哈希，
      但文件内容将被解压缩。

   --might-gzip
      如果后端可能gzip对象，请设置此项。

      通常，提供程序在下载时不会更改对象。
      如果一个对象在上传时未使用“Content-Encoding: gzip”，
      那么在下载时也不会设置它。
      
      然而，某些提供程序可能会在未使用“Content-Encoding: gzip”
      的情况下对对象进行gzip压缩（例如Cloudflare）。

      这种情况的症状将看到以下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载一个带有
      设置了Content-Encoding: gzip且使用了分块传输编码的对象，
      那么rclone将在传输过程中解压缩对象。

      如果设置为unset（默认值），则rclone将根据提供商的设置来选择如何应用，
      但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id value      AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                存储桶创建和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value           OSS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行环境（环境变量或EC2/ECS元数据）获取AWS凭证。（默认值为false）[$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      存储新对象时要使用的存储类型。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。（默认值为"5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的文件大小阈值。（默认值为"4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。（默认值为false）[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。（默认值为false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。（默认值为false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。（默认值为"Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。（默认值为true）[$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。（默认值为1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset（默认值为"unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0为自动。（默认值为0）[$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大部分数。（默认值为10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池将刷新的频率。（默认值为"1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。（默认值为false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，请设置此项。（默认值为"unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。（默认值为false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对已上传对象进行HEAD请求以检查完整性。（默认值为false）[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。（默认值为false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据（默认值为false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。（默认值为4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的文件大小阈值。（默认值为"200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行校验。（默认值为"unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传。（默认值为false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。（默认值为false）[$V2_AUTH]
   --version-at value               显示指定时间的文件版本。（默认值为"off"）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。（默认值为false）[$VERSIONS]

```
{% endcode %}