# 腾讯云对象存储（COS）

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 tencentcos - 腾讯云对象存储（COS）

使用方法：
   singularity storage update s3 tencentcos [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在访问密钥ID和秘密访问密钥为空白时生效。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --endpoint
      腾讯云COS API的终端节点。

      示例：
         | cos.ap-beijing.myqcloud.com       | 北京区域
         | cos.ap-nanjing.myqcloud.com       | 南京区域
         | cos.ap-shanghai.myqcloud.com      | 上海区域
         | cos.ap-guangzhou.myqcloud.com     | 广州区域
         | cos.ap-nanjing.myqcloud.com       | 南京区域
         | cos.ap-chengdu.myqcloud.com       | 成都区域
         | cos.ap-chongqing.myqcloud.com     | 重庆区域
         | cos.ap-hongkong.myqcloud.com      | 香港（中国）区域
         | cos.ap-singapore.myqcloud.com     | 新加坡区域
         | cos.ap-mumbai.myqcloud.com        | 孟买区域
         | cos.ap-seoul.myqcloud.com         | 首尔区域
         | cos.ap-bangkok.myqcloud.com       | 曼谷区域
         | cos.ap-tokyo.myqcloud.com         | 东京区域
         | cos.na-siliconvalley.myqcloud.com | 硅谷区域
         | cos.na-ashburn.myqcloud.com       | 弗吉尼亚区域
         | cos.na-toronto.myqcloud.com       | 多伦多区域
         | cos.eu-frankfurt.myqcloud.com     | 法兰克福区域
         | cos.eu-moscow.myqcloud.com        | 莫斯科区域
         | cos.accelerate.myqcloud.com       | 使用腾讯COS加速终端节点

   --acl
      在创建桶和存储或复制对象时使用的预设ACL。
      
      该ACL用于创建对象，如果bucket_acl未设置，则还用于创建桶。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，当服务器端复制对象时，将应用该ACL，因为S3不会复制源的ACL，而是写入新的ACL。
      
      如果acl是一个空字符串，则不会添加X-Amz-Acl:头，将使用默认（私有）。

      示例：
         | default | 所有者获得完全控制权。
         |         | 没有其他用户可以访问（默认）。

   --bucket-acl
      在创建桶时使用的预设ACL。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      如果未设置bucket_acl，则仅在创建桶时应用此ACL。
      
      如果"acl"和"bucket_acl"是空字符串，则不会添加X-Amz-Acl:头，将使用默认（私有）。

      示例：
         | private            | 所有者获得完全控制权。
         |                    | 没有其他用户可以访问（默认）。
         | public-read        | 所有者获得完全控制权。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者获得完全控制权。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 通常不建议在桶上授予此权限。
         | authenticated-read | 所有者获得完全控制权。
         |                    | AuthenticatedUsers组获得读取权限。

   --storage-class
      把新对象存储在腾讯云COS中使用的储存类型。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类型
         | ARCHIVE     | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换到分块上传的截断点。
      
      任何大于此大小的文件将分块上传，每块大小为chunk_size。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如使用`rclone rcat`上传的文件或使用`rclone mount`上传的文件或谷歌照片或谷歌文档）时，将使用该分块大小进行分块上传。
      
      请注意，每个传输的内存中缓冲了“--s3-upload-concurrency”大小的分块。
      
      如果您正在高速链接上传输大型文件，并且具有足够的内存，那么增加此值将加快传输速度。
      
      当上传已知大小的大文件时，Rclone将自动增加分块大小，以使其保持在10000块的限制以下。
      
      大小未知的文件使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，并且最多可以有10,000个块，这意味着默认情况下您可以流式传输的文件的最大大小为48 GiB。如果要流式传输更大的文件，则需要增加chunk_size。
      
      增加分块大小会降低配置选项“-P”标志显示的进度统计的准确性。当使用AWS SDK缓冲缓冲时，Rclone将分块视为已发送，实际上可能仍在上传。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况越发错误。

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义在执行多部分上传时使用的最大多部分块数。
      
      如果服务不支持AWS S3的10,000个块的规范，则这可能很有用。
      
      当上传已知大小的大文件时，Rclone将自动增加分块大小，以保持在此块数限制以下。
      

   --copy-cutoff
      切换到分块复制的截断点。
      
      需要分块服务器端复制的大于此大小的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，在上传之前，Rclone会计算输入的MD5校验和，以便可以将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致较长的延迟以便启动上传。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则Rclone可以使用共享凭证文件。
      
      如果此变量为空，则Rclone会查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      在共享凭证文件中使用的配置文件。
      
      如果env_auth = true，则Rclone可以使用共享凭证文件。该变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"，如果该环境变量也未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您正在高速链接上上传少量大文件，并且这些上传没有充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则Rclone将使用路径样式访问；如果为false，则Rclone将使用虚拟路径样式访问。有关详细信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）需要将此设置为false - Rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则Rclone将使用v4身份验证。如果设置了它，则Rclone将使用v2身份验证。
      
      仅在v4签名不起作用的情况下使用，例如前Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为"MaxKeys"、"max-items"或"page-size"，来自AWS S3规范。
      大多数服务将列表响应截断为1000个对象，即使请求的数量更多。在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项进行增加。

   --list-version
      要使用的ListObjects的版本：1、2或0自动选择。
      
      当S3最初发布时，它只提供了ListObjects调用来枚举桶中的对象。
      
      然而，在2016年5月引入了ListObjectsV2调用。这个调用性能更好，应尽可能使用它。
      
      如果设置为默认值0，则Rclone将根据提供者设置猜测应调用哪个list objects方法。如果它的猜测错误，则可以在此处手动设置它。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，在可用时在使用控制字符在文件名中时，这更可靠。如果设置为unset（默认值），则Rclone将根据提供者设置选择要应用的选项，但您可以在此处覆盖Rclone的选择。
      

   --no-check-bucket
      如果设置，则不会尝试检查桶是否存在或创建它。
      
      如果您知道桶已经存在，尝试尽量减少Rclone执行的事务数量时，这可能很有用。
      
      如果您使用的用户没有桶创建权限，则可能也需要此选项。在v1.52.0之前，此操作将由于错误而默默传递。
      

   --no-head
      如果设置，则不会HEAD已上传的对象以检查完整性。
      
      如果您尝试尽量减少Rclone执行的事务数量，这可能很有用。
      
      设置后，意味着如果Rclone在PUT上传对象后收到200 OK消息，则假设它已正确上传。
      
      特别是，它将假设：
      
      - 上传时的元数据，包括修改时间、存储类型和内容类型与上传时相同
      - 大小与上传时相同
      
      它从单个部分PUT请求的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象的长度未知，则Rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确的机会，因此不建议进行正常操作。实际上，即使在此标志的情况下，检测不到的上传失败的机会也非常小。
      

   --no-head-object
      如果设置，则在GET对象时不执行HEAD。

   --encoding
      后端的编码。
      
      有关更多信息，请参阅概述中的[编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区（例如大文件的多部分）将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端与HTTP/2存在未解决的问题。HTTP/2默认启用于s3后端，但可以在此处禁用。当问题解决时，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631。
      

   --download-url
      下载的自定义终端节点。
      这通常设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载数据的更便宜的流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该是true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单个部分上传
      
      如果为false，则Rclone将使用来自AWS SDK的PutObject上传对象。
      
      Rclone的版本< 1.59使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应该需要这样做。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示文件版本与指定时间相同的文件版本。
      
      参数应为日期、"2006-01-02"、日期时间"2006-01-02 15:04:05"或距离那时多久之前的持续时间，例如"100d"或"1h"。
      
      请注意，使用此选项时不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置此参数，将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置将对象上传到S3。通常，Rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则在接收到带有"Content-Encoding: gzip"的文件时，Rclone将解压缩这些文件。这意味着Rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能使用gzip压缩对象，则设置此参数。
      
      通常，提供商在下载对象时不会更改对象。如果即使未使用`Content-Encoding: gzip`上传对象，提供商也会对对象进行gzip压缩（例如Cloudflare）。
      
      如果设置了此标志，并且Rclone使用设置了`Content-Encoding: gzip`和分块传输编码的对象，则Rclone将在传输过程中解压缩该对象。
      
      如果设置为unset（默认值），则Rclone将根据提供者设置选择要应用的选项，但您可以在此处覆盖Rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                在创建桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value           腾讯云COS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助信息
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      把新对象存储在腾讯云COS中使用的储存类型。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               在创建桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 （默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截断点。 （默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置此参数，将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 （默认值："Slash，InvalidUtf8，Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0自动选择。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 （默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能使用gzip压缩对象，则设置此参数。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不会尝试检查桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会HEAD已上传的对象以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在GET对象时不执行HEAD。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据（默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  在共享凭证文件中使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截断点。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证（默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单个部分上传（默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value               显示文件版本与指定时间相同的文件版本。 （默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 （默认值：false）[$VERSIONS]

```
{% endcode %}