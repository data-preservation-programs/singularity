# Tencent Cloud Object Storage (COS)

{% code fullWidth="true" %}
```
名称：
   singularity 存储更新 s3 tencentcos - 腾讯云对象存储（COS）

用法：
   singularity 存储更新 s3 tencentcos [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（如果access_key_id和secret_access_key为空，则从环境变量或EC2/ECS元数据中获取）
      
      只适用于access_key_id和secret_access_key为空的情况。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境变量（env vars或IAM）中获取AWS凭证。

   --access-key-id
      AWS Access Key ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --endpoint
      腾讯COS API的终端节点。

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
      在创建存储桶和存储或复制对象时使用的预设ACL。
      
      此ACL用于创建对象以及（如果未设置bucket_acl）创建存储桶。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，此ACL将被应用，因为S3不会复制源的ACL，而是写入新ACL。
      
      如果acl为空字符串，则不添加X-Amz-Acl:标头，并且将使用默认（私有）。

      示例：
         | default | 所有者获取Full_CONTROL。
         |         | 除此之外，没有其他人有访问权限（默认）。

   --bucket-acl
      在创建存储桶时使用的预设ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，此ACL仅在创建存储桶时应用。如果未设置，则使用"acl"。
      
      如果"acl"和"bucket_acl"为空字符串，则不添加X-Amz-Acl:标头，并且将使用默认（私有）。

      示例：
         | private            | 所有者获取FULL_CONTROL。
         |                    | 除此之外，没有其他人有访问权限（默认）。
         | public-read        | 所有者获取FULL_CONTROL。
         |                    | AllUsers组获取READ访问权限。
         | public-read-write  | 所有者获取FULL_CONTROL。
         |                    | AllUsers组获取READ和WRITE访问权限。
         |                    | 一般不建议在存储桶上授予此权限。
         | authenticated-read | 所有者获取FULL_CONTROL。
         |                    | AuthenticatedUsers组获取READ访问权限。

   --storage-class
      在Tencent COS中存储新对象时使用的存储类型。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类型
         | ARCHIVE     | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换到分块上传的截止点。
      
      任何比此大小大的文件都将分块上传，每块的大小为chunk_size。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      上传大于upload_cutoff的文件或大小未知的文件（例如来自"rclone rcat" 或 "rclone mount"、谷歌照片或谷歌文档的文件）时，将使用此分块大小进行分块上传。
      
      请注意，每个传输会将"--s3-upload-concurrency"个此大小的分块缓存在内存中。
      
      如果您正在通过高速链路传输大文件并且具有足够的内存，那么增加此值将加快传输速度。
      
      在上传已知大小的大文件时，rclone将自动增加分块大小，以保持低于10,000个分块的限制。
      
      未知大小的文件使用配置的chunk_size进行上传。由于默认的分块大小为5 MiB，最多可以有10,000个分块，因此通过默认设置可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加分块大小。
      
      增加分块大小会降低使用"-P"标志显示的进度统计的准确性。当块被AWS SDK缓冲时，rclone会将其视为已发送的块，实际上可能仍在上传。较大的分块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况偏离的更多。

   --max-upload-parts
      分块上传中的最大块数。
      
      此选项定义了执行分块上传时要使用的最大multipart块数。
      
      如果服务不支持AWS S3规范的10,000个块，则这可能非常有用。
      
      在上传已知大小的大文件时，rclone将自动增加分块大小，以保持低于此块数限制。

   --copy-cutoff
      切换到多段复制的截止点。
      
      任何大于此大小的需要在服务器端复制的文件将被分块复制。

      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。
      
      通常，rclone会在上传对象之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件开始上传可能会造成长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果env值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传相同文件的块数量。
      
      如果您正在通过高速链路上传少量大文件，并且这些上传未充分利用带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径风格访问；如果为false，则使用虚拟主机风格访问。
      
      如果设置为true（默认值），则rclone将使用路径风格访问；如果设置为false，则rclone将使用虚拟路径风格。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求将此设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果此值为false（默认值），则rclone将使用v4身份验证。如果设置此值，rclone将使用v2身份验证。
      
      仅在v4签名无法工作时使用此选项，例如旧版本的Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。
      大多数服务无论请求的数量是多少，都会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过“rgw list buckets max chunk”选项进行增加。

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      当S3最初发布时，它仅提供了用于枚举存储桶中的对象的ListObjects调用。
      
      但是在2016年5月，引入了ListObjectsV2调用。这是性能更高的版本，如果可能，应该使用。
      
      如果设置为默认值0，则rclone会根据设置的提供者猜测要调用哪个列举对象方法。如果它猜错了，那么可以在此手动设置。

   --list-url-encode
      是否对网址进行编码：true/false/unset
      
      某些提供商支持URL编码文件列表，如果可用，则在使用控制字符文件名时更可靠。如果将其设置为未设置（默认值），则rclone将根据提供商的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不会尝试检查存储桶是否存在或创建。
      
      如果您知道存储桶已经存在，那么这可能对尽量减少rclone进行的事务数很有用。
      
      如果使用的用户没有创建存储桶的权限，则可能也需要此选项。在v1.52.0之前，这将被无声地传递，因为存在错误。

   --no-head
      如果设置，则不会对上传的对象进行HEAD请求以检验完整性。
      
      这在试图尽量减少rclone进行的事务数时很有用。
      
      设置后，如果rclone在使用PUT上传对象后收到200 OK消息，则会假设它已经上传成功。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的内容相同
      - 大小与上传的内容相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传源对象的长度未知，则rclone会执行HEAD请求。
      
      设置此标志会增加未检测到的上传错误的几率，特别是大小不正确，因此不推荐在正常操作中使用。实际上，即使设置此标志，未检测到的上传错误的几率也非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区的上传（例如多部分上传）将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端使用http2。
      
      目前s3（特别是minio）后端存在一个未解决的问题，与HTTP/2有关。s3后端默认启用HTTP/2，但可以在此禁用。等到问题解决后，将删除此标志。
      
      请参阅：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据的出境流量更便宜。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应该设置为true、false或留空，以使用提供者的默认设置。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单块上传
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      版本为rclone < 1.59的情况下，会使用预签名请求上传单个部分的对象，将此标志设置为true将重新启用该功能。这只在特殊情况下或用于测试时才有必要。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应该是一个日期，"2006-01-02"，日期时间"2006-01-02
      15:04:05"，或离现在时间的时间段，例如"100d"或"1h"。
      
      请注意，使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用设置"Content-Encoding: gzip"上传对象到S3。通常，rclone会将这些文件作为已压缩对象下载。
      
      如果设置了此标志，则rclone将在接收到设置为“Content-Encoding: gzip”的文件时解压缩这些文件。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能gzip压缩对象，请设置此标志。

      通常，提供者在下载时不会更改对象。如果即使未使用`Content-Encoding: gzip`上传了对象，下载时它不会被设置。

      但是，某些提供者甚至在未使用`Content-Encoding: gzip`上传对象时也会对其进行gzip压缩（例如Cloudflare）。

      这种情况的典型表现是接收到以下错误消息：

          错误 corrupted on transfer: sizes differ NNN vs MMM

      如果设置了此标志，并且rclone下载了具有设置为`Content-Encoding: gzip`并且带有分块传输编码的对象，则rclone将在接收到对象时即时进行解压缩。

      如果设置为未设置（默认值），则rclone将根据提供商的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value      AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                在创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value           腾讯COS API的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（如果access_key_id和secret_access_key为空，则从环境变量或EC2/ECS元数据中获取）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --secret-access-key value  AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --storage-class value      在Tencent COS中存储新对象时使用的存储类型。[$STORAGE_CLASS]

   高级选项

   --bucket-acl value               在创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到多段复制的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端使用http2。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径风格访问；如果为false，则使用虚拟主机风格访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对网址进行编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip压缩对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不会尝试检查存储桶是否存在或创建。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对上传的对象进行HEAD请求以检验完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单块上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}