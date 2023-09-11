# Tencent Cloud Object Storage（COS）

{% code fullWidth="true" %}
```
名称：
   singularity storage create s3 tencentcos - 腾讯云对象存储（COS）

用法：
   singularity storage create s3 tencentcos [命令选项] [参数...]

说明:
   --env-auth
      从运行时获取AWS凭证（环境变量或环境中无环境变量时使用EC2/ECS元数据）。
      
      仅在access_key_id和secret_access_key为空时有效。

      示例:
         | false | 下一步输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS Access Key ID。
      
      不填则表示匿名访问或使用运行时凭证。

   --secret-access-key
      AWS Secret Access Key (密码)。
      
      不填则表示匿名访问或使用运行时凭证。

   --endpoint
      腾讯COS API的终端。

      示例:
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
         | cos.accelerate.myqcloud.com       | 使用腾讯COS加速终端

   --acl
      创建存储桶、存储或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，如果未设置bucket_acl，则也用于创建桶。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，此ACL将会应用，因为S3不会复制源对象的ACL，而是写入一个新的ACL。
      
      如果acl是一个空字符串，则不会添加X-Amz-Acl：标头，将使用默认（私有）。

      示例:
         | default | 拥有者获得 FULL_CONTROL 权限。
         |         | 其他人没有访问权限（默认）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      如果未设置bucket_acl，则会应用此ACL到创建桶上。

      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl：
      标头，将使用默认（私有）。

      示例:
         | private            | 拥有者获得 FULL_CONTROL 权限。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 拥有者获得 FULL_CONTROL 权限。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 拥有者获得 FULL_CONTROL 权限。
         |                    | AllUsers组获得读取和写入权限。
         |                    | 一般不推荐将其授权给桶。
         | authenticated-read | 拥有者获得 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers组获得读取权限。

   --storage-class
      存储新对象时使用的存储类。

      示例:
         | <unset>     | 默认
         | STANDARD    | 标准存储类
         | ARCHIVE     | 存档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换到分片上传的文件大小临界值。
      
      大于此大小的文件将按照chunk_size分片上传。
      此值的最小值为0，最大值为5 GiB。

   --chunk-size
      上传使用的分片大小。
      
      当上传大于upload_cutoff的文件或者大小未知的文件（例如通过"rclone rcat"上传或通过"rclone mount"或谷歌照片或谷歌文档上传）时，将会按照此分片大小进行分片上传。
      
      请注意，每次传输都会在内存中缓冲此大小的"--s3-upload-concurrency"个分片。
      
      如果您正在通过高速链接传输大文件并且拥有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件以保持在10000个分片限制以下时，rclone将自动增加分片大小。
      
      未知大小的文件使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，最多可以有10000个分片，所以默认情况下您可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加chunk_size。
      
      增加分片大小会减少使用"-P"标志显示的进度统计的准确性。当rclone将分片缓冲到AWS SDK时，rclone将分片视为已发送，而实际上可能仍在上传。更大的分片大小意味着更大的AWS SDK缓冲区和进度报告与实际情况更加偏差。
      

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义在执行多部分上传时要使用的最大分片数。
      
      当上传已知大小的大文件以保持在这个片数限制以下时，rclone将自动增加分片大小。

   --copy-cutoff
      切换到分片复制的文件大小临界值。
      
      大于此大小的需要服务端复制的文件将按照此大小进行分片复制。
      
      此值的最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常rclone会在上传之前计算输入数据的MD5校验和，以便将其添加到对象的元数据中。这在数据完整性检查中非常有用，但是对于大文件开始上传时会导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空
      它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此
      变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或
      未设置此环境变量的情况下为"default"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发程度。
      
      这是同时上传的相同文件的分片数。
      
      如果您正在通过高速链接上传大量的大文件，并且这些上传未充分利用您的带宽，那么增加此值可能有助于提高传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式访问。请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      以了解更多信息。
      
      某些提供商（如AWS，阿里云OSS，网易COS或腾讯COS）要求此值设置为
      false - rclone将根据提供程序的设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），rclone将使用v4身份验证。
      如果设置，则rclone将使用v2身份验证。
      
      只有在v4签名无效时使用此功能，例如早于Jewel/v10 CEPH的版本。

   --list-chunk
      列出清单的大小（每个ListObject S3请求的列表响应）。
      
      此选项也称为“MaxKeys”，“max-items”或“AWS S3规范中的“page-size”。
      大多数服务即使请求的数量超过1000个，也会截断列表响应为1000个对象。
      在AWS S3中，这是一个全局最大值，不能更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项进行增加。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。
      
      当S3最初发布时，它仅提供了“ListObjects”调用以枚举存储桶中的对象。
      
      然而，在2016年5月，引入了“ListObjectsV2”调用。这是
      更高性能，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据提供程序的设置猜测要调用的
      列出对象方法。如果它猜错了，则可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      有些提供商支持URL对列表进行编码，如果可用，这在使用控制字符的文件名时更可靠。如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的方法，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建它。
      
      当试图最小化rclone执行的事务数量时，这可能很有用，如果您知道存储桶已经存在。
      
      如果使用的用户没有创建桶的权限，则也可能会需要这样做。在版本v1.52.0以前，由于一个错误，此操作不会发出任何警告。
      

   --no-head
      如果设置，则不对上传的对象进行HEAD操作以检查完整性。
      
      当试图最小化rclone执行的事务数量时，这可能很有用。
      
      设置它意味着如果rclone在使用PUT上传对象后收到200 OK消息，则会假设已经成功上传了对象。
      
      特别地，它将假定：
      
      - 元数据，包括modtime，存储类和内容类型与上传的相同
      - 大小与上传的相同
      
      对于单部分PUT，它从响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，它不会读取这些项目。
      
      如果上传一个未知大小的源对象，那么rclone **会**执行HEAD请求。
      
      设置此标志将增加未检测到的上传失败的机会，特别是不正确的大小，因此不推荐在正常操作中使用。实际上，即使设置了此标志，在上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在GET对象之前不执行HEAD操作。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池将被刷新的时间间隔。
      
      需要附加缓冲区的上传（例如分片）将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中移除。

   --memory-pool-use-mmap
      是否在内部内存缓冲池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个尚未解决的问题，该问题与HTTP/2有关。默认情况下，s3后端启用HTTP/2，但可以在此禁用。在解决这个问题之前，将删除此标志。
      
      请参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端点。
      此项通常设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载数据的更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该设置为true、false或不设置以使用提供商的默认值。
      

   --use-presigned-request
      是否使用签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK中的PutObject上传
      对象。
      
      rclone < 1.59的版本使用预签名请求上传单个
      部分对象，将此标志设置为true将重新启用该
      功能。除非特殊情况或测试，否则不应该使用这个标志。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定时间显示文件版本。
      
      参数应该是一个日期，"2006-01-02"，日期时间 "2006-01-02
      15:04:05" 或表示时间长久以前的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此功能时不允许进行文件写操作，
      因此不能上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置将对象上传到S3。通常情况下，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在接收到具有"Content-Encoding: gzip"设置的文件时进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能压缩对象，请设置此项。
      
      通常，提供者在下载时不会修改对象。如果下载一个对象时未上传`Content-Encoding: gzip`，则不会在下载时设置它。
      
      但是，即使未使用`Content-Encoding: gzip`上传，某些提供者也可能对对象进行压缩（例如Cloudflare）。
      
      这种情况的症状可能是收到类似的错误
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志并且rclone下载具有设置了`Content-Encoding: gzip`和chunked传输编码的对象，那么rclone将会即时解压缩对象。
      
      如果此项设置为unset（默认值），则rclone将根据提供商的设置选择要应用的方法，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value           腾讯COS API的终端。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或环境中无环境变量时使用EC2/ECS元数据）。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  AWS Secret Access Key (密码)。[$SECRET_ACCESS_KEY]
   --storage-class value      存储新对象时使用的存储类。[$STORAGE_CLASS]

   高级设置

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               上传使用的分片大小。（默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换到复制分片的临界大小。（默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 （默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列出清单的大小（每个ListObject S3请求的列表响应）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0表示自动。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池将被刷新的时间间隔。 （默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请设置此项。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，不尝试检查存储桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对上传的对象进行HEAD操作以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在GET对象之前不执行HEAD操作。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 （默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发程度。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分片上传的文件大小临界值。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 （默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用签名请求或PutObject进行单部分上传 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value               按指定时间显示文件版本。 （默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 （默认值：false）[$VERSIONS]

   通用

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```