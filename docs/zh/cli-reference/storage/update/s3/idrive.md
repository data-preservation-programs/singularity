# IDrive e2

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 idrive - IDrive e2

USAGE:
   singularity storage update s3 idrive [command options] <name|id>

DESCRIPTION:
   --env-auth
      从运行时获取AWS凭证 (环境变量或EC2/ECS元数据, 如果没有环境变量)。
      
      当access_key_id和secret_access_key为空时才有效。

      示例:
         | false | 下一步输入AWS凭证。
         | true  | 从环境获取AWS凭证 (优先环境变量或IAM)。

   --access-key-id
      AWS访问密钥ID。
      
      留空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥 (密码)。
      
      留空表示匿名访问或运行时凭证。

   --acl
      创建存储桶、存储或复制对象时指定的预设ACL。
      
      此ACL用于创建对象，如果bucket_acl未设置，也用于创建存储桶。
      
      请参考以下链接了解更多详情：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，当使用S3进行服务器端拷贝时，此ACL会被应用，
      因为S3不会复制源对象的ACL，而是新写一个。
      
      如果ACL为空字符串，则不会添加X-Amz-Acl:头，并且会使用默认的(private)。
      

   --bucket-acl
      创建存储桶时指定的预设ACL。
      
      请参考以下链接了解更多详情：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，此ACL仅在创建存储桶时应用。
      如果未设置，则使用"acl"。
      
      如果"acl"和"bucket_acl"都为空字符串，则不会添加X-Amz-Acl:头，并且会使用默认的(private)。
      

      示例:
         | private            | 拥有者获取完全控制权限。
         |                    | 没有其他用户对此对象具有访问权限 (默认)。
         | public-read        | 拥有者获取完全控制权限。
         |                    | AllUsers用户组获得读取权限。
         | public-read-write  | 拥有者获取完全控制权限。
         |                    | AllUsers用户组获得读取和写入权限。
         |                    | 一般不推荐在存储桶上执行此操作。
         | authenticated-read | 拥有者获取完全控制权限。
         |                    | AuthenticatedUsers用户组获得读取权限。

   --upload-cutoff
      切换到分块上传的文件大小阈值。
      
      大于此大小的文件将使用chunk_size分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或具有未知大小的文件
      (例如通过"rclone rcat"或"rclone mount"或Google Photos或Google Docs上传的文件)时，
      将使用此大小进行分块上传。
      
      请注意，每个传输在内存中缓冲的是"--s3-upload-concurrency"个这样大小的块。
      
      如果您正在通过高速链接传输大文件，并且具有足够的内存，
      则增加此值可以加快传输速度。
      
      当上传已知大小的大文件以保持在10000个块以下时，rclone会自动增加分块大小。
      
      未知大小的文件以配置的chunk_size进行上传。
      由于默认的chunk_size为5 MiB，并且最多可以有10000个chunk，
      这意味着默认情况下的流式上传文件的最大大小为48 GiB。
      如果要流式上传更大的文件，需要增加chunk_size。
      
      增加的分块大小会减少使用"-P"参数显示的进度统计的准确性。
      当使用AWS SDK的缓冲区缓冲时，rclone认为已发送chunk，但实际上可能仍在上传。
      更大的分块大小意味着更大的AWS SDK缓冲区，并且进度报告与实际情况越来越不符。
      

   --max-upload-parts
      分块上传的最大块数。
      
      此选项定义了进行分块上传时使用的最大分块数。
      
      如果服务不支持AWS S3规范中的10000个块，此选项可能很有用。
      
      当上传已知大小的大文件以保持在此分块数以下时，
      rclone会自动增加分块大小。
      

   --copy-cutoff
      切换到分块复制的文件大小阈值。
      
      需要进行服务器端复制的大于此大小的文件将被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone将在上传之前计算输入的MD5校验和，
      以便将其添加到对象的元数据中。这对于数据完整性检查很有用，
      但对于大文件来说可能会导致长时间的上传延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果该变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。
      如果环境变量为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\\.aws\\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果环境变量也未设置）。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      它表示同时上传的相同文件的分块数。
      
      如果您正在通过高速链接上传少量大文件，并且这些上传没有充分利用带宽，
      那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，使用路径形式访问；如果为false，使用虚拟主机形式访问。
      
      如果为true（默认值），rclone将使用路径形式访问；
      如果为false，则rclone将使用虚拟主机形式访问。
      有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      一些提供商 (例如AWS、Aliyun OSS、Netease COS或Tencent COS) 要求将其设置为false -
      rclone会根据提供商设置自动完成此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置了它，则rclone将使用v2身份验证。
      
      仅在v4签名无法工作时使用，例如旧的Jewel/v10 CEPH。

   --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。
      大多数服务在请求超过1000个对象时都会截断响应列表到1000个。
      在AWS S3中，这是一个全局限制，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以增加"rgw list buckets max chunk"选项。

   --list-version
      要使用的ListObjects版本：1、2或0为自动。
      
      当S3最初推出时，它只提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月引入了ListObjectsV2调用。这是高性能的，
      如果可能的话应该使用。
      
      如果设置为默认值0，则rclone将根据设置的提供程序猜测调用哪个对象列表方法。
      如果猜测不正确，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行url编码：true/false/unset
      
      某些提供商支持对列表进行URL编码，如果可用，则在文件名中使用控制字符时，这是更可靠的方法。
      如果设置为unset（默认值），则rclone将根据提供商设置选择应用什么，但是您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置了，则不尝试检查桶是否存在或创建它。
      
      如果您知道存储桶已存在，这可以减少rclone执行的事务数量。
      
      如果使用的用户没有创建存储桶的权限，则可能需要这样做。 v1.52.0之前的版本由于错误，这将会无声地通过。
      

   --no-head
      如果设置了，则不会在获取对象之前执行HEAD请求以检查完整性。
      
      如果要最小化rclone执行的事务数量，这可能会很有用。
      
      设置后，意味着如果rclone在通过PUT上传对象后收到200 OK消息，
      则会假定对象已正确上传。
      
      特别是，它会假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的相同
      - 大小与上传的相同
      
      对于单部分PUT请求，它会从响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，这些项目不会被读取。
      
      如果上传一个未知长度的源对象，则rclone **会**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，
      特别是错误大小的机会，因此不推荐在正常操作中使用它。
      实际上，即使使用此标志，检测不到的上传失败的机会也非常小。
      

   --no-head-object
      如果设置了，则在获取对象之前不执行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池的刷新频率。
      
      需要额外缓冲区（例如分块上传）的上传将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存缓冲池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一些未解决的问题和HTTP/2之间的兼容性问题。
      s3后端默认启用HTTP/2，但可以在此处禁用。
      问题解决后，此标志将被移除。
      
      请参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      用于下载的自定义端点。
      这通常设置为CloudFront CDN的URL，因为AWS S3提供了通过CloudFront网络下载数据的更低出站流量费用。

   --use-multipart-etag
      分块上传时是否使用ETag进行校验
      
      此选项应为true、false或不设置，以使用提供商的默认设置。
      

   --use-presigned-request
      单部分上传时是否使用预签名请求或PutObject
      
      如果设置为false，rclone将使用AWS SDK的PutObject来上传对象。
      
      rclone的版本< 1.59使用预签名请求来上传单部分对象，
      设置此标志为true将重新启用该功能。除非在特殊情况或测试中，
      否则不应该需要此选项。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定的时间显示文件版本。
      
      参数应为日期，"2006-01-02"，日期时间"2006-01-02 15:04:05"或距离现在的时间段，例如"100d"或"1h"。
      
      请注意，当使用此选项时，不允许进行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则会解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置将对象上传到S3。通常，rclone会将这些文件下载为压缩对象。
      
      如果设置了此标志，则rclone在接收到这些具有"Content-Encoding: gzip"的文件时，会对其进行解压缩。
      这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能gzip对象，则设置此标志。
      
      通常，在下载时，提供商不会更改对象。如果对象没有使用`Content-Encoding: gzip`上传，则下载时不会设置该项。
      
      但是，某些提供商可能会在没有使用`Content-Encoding: gzip`上传的情况下对对象进行gzip压缩（例如Cloudflare）。
      
      出现以下错误时的一个症状是：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载具有设置Content-Encoding: gzip和块传输编码的对象，
      那么rclone会在传输期间即时解压缩对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择应用什么，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的预设ACL。 [$ACL]
   --env-auth                 从运行时获取AWS凭证 (环境变量或EC2/ECS元数据, 如果没有环境变量)。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --secret-access-key value  AWS Secret Access Key (password). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               创建存储桶时使用的预设ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的文件大小阈值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则会解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             用于下载的自定义端点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，使用路径形式访问；如果为false，使用虚拟主机形式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行url编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0为自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传的最大块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池的刷新频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置了，则不尝试检查桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置了，则不会在获取对象之前执行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置了，则在获取对象之前不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的文件大小阈值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       分块上传时是否使用ETag进行校验 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          单部分上传时是否使用预签名请求或PutObject (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}