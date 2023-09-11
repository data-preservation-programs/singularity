# IONOS Cloud

{% code fullWidth="true" %}
```
名称:
   singularity storage create s3 ionos - IONOS Cloud

用法:
   singularity storage create s3 ionos [命令选项] [参数...]

描述:
   --env-auth
      从运行时获取AWS凭据（环境变量或EC2/ECS元数据）。
      
      仅当access_key_id和secret_access_key为空时有效。

      示例：
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境中获取AWS凭据（环境变量或IAM）。

   --access-key-id
      AWS的访问密钥ID。
      
      如果要进行匿名访问或者使用运行时凭据，请留空。

   --secret-access-key
      AWS的秘密访问密钥（密码）。
      
      如果要进行匿名访问或者使用运行时凭据，请留空。

   --region
      您的存储桶将被创建和数据存储的区域。
      

      示例：
         | de           | 德国法兰克福
         | eu-central-2 | 德国柏林
         | eu-south-2   | 西班牙洛戈罗尼奥

   --endpoint
      IONOS S3对象存储的终结点。
      
      指定与所在区域相同的终结点。

      示例：
         | s3-eu-central-1.ionoscloud.com | 德国法兰克福
         | s3-eu-central-2.ionoscloud.com | 德国柏林
         | s3-eu-south-2.ionoscloud.com   | 西班牙洛戈罗尼奥

   --acl
      创建存储桶、存储或复制对象时使用的预设ACL。
      
      在创建对象时使用该ACL，如果没有设置bucket_acl，则在创建存储桶时也会使用该ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当S3服务器端复制对象时，会应用此ACL，因为S3不会复制源对象的ACL，而是会新写入一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl头，并使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预设ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      如果没有设置，则只在创建存储桶时使用“acl”。

      如果“acl”和“bucket_acl”都为空字符串，则不会添加X-Amz-Acl头，并使用默认值（private）。

      示例：
         | private            | 拥有者具有完全控制权。
         |                    | 无其他用户具有权限（默认设置）。
         | public-read        | 拥有者具有完全控制权。
         |                    | 所有用户组具有读取权限。
         | public-read-write  | 拥有者具有完全控制权。
         |                    | 所有用户组具有读取和写入权限。
         |                    | 通常不建议在存储桶上进行此操作。
         | authenticated-read | 拥有者具有完全控制权。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      切换为分块上传的文件大小。
      
      大于此大小的文件将以chunk_size的大小进行分块上传。
      最小值为0，最大值为5GiB。

   --chunk-size
      用于上传的分块大小。
      
      对于大于upload_cutoff的文件或大小未知的文件（例如来自“rclone rcat”或使用“rclone mount”或Google照片或Google文档上传的文件），将使用此分块大小进行分块上传。
      
      请注意，每个传输将在内存中缓冲chunk_size个分块。
      
      如果您在高速链接上传输大文件并且具有足够的内存，那么增加此值将加快传输速度。
      
      Rclone将在上传已知大小的大型文件时自动增加分块大小，以保持在10000个分块限制之内。
      
      未知大小的文件将使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，并且最多可以有10000个分块，因此默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果要流式传输更大的文件，您需要增加chunk_size。
      
      增加分块大小会降低使用“-P”标志显示的进度统计的准确性。 Rclone在通过AWS SDK缓冲分块时将分块视为已发送，而实际上可能仍在上传。更大的分块大小意味着更大的AWS SDK缓冲区和进度报告与真实情况更大的偏离。

   --max-upload-parts
      单个分块上传的最大分块数。
      
      此选项定义在执行分块上传时要使用的最大分块数。
      
      如果某个服务不支持AWS S3 10000个分块的规范，则此选项可能非常有用。
      
      当上传已知大小的大文件时，Rclone将自动增加分块大小以保持在此分块数限制之下。

   --copy-cutoff
      切换为分块复制的文件大小。
      
      大于此大小需要进行服务器端复制的文件将以此大小的块进行复制。
      
      最小值为0，最大值为5GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传之前计算输入数据的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间延迟才能开始上传。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境值为空，则它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"默认"如果该环境变量未设置。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      并发用于分块上传的块数。
      
      如果您通过高速链路上传少量大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式。有关更多信息，请参阅[the AWS S3
      文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置，则rclone将使用v2身份验证。
      
      仅当v4签名无法工作时使用，例如 Jewel/v10 CEPH 之前的版本。

   --list-chunk
      列表检索的大小（用于每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求了更多的内容，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项进行增加。

   --list-version
      要使用的ListObjects版本：1、2或0自动选择。
      
      当S3首次发布时，只提供了ListObjects调用，用于枚举存储桶中的对象。
      
      然而，在2016年5月，引入了ListObjectsV2调用。这是更高性能的方法，如果可能的话应使用。
      
      如果设置为默认值0，则rclone将根据所设置的提供程序猜测要调用的列表对象方法。如果它猜测错误，则可以在此手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供程序支持URL编码列表，如果可用，则在文件名中使用控制字符时更可靠。如果设置为unset（默认值），则rclone将根据提供程序的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，不会尝试检查存储桶是否存在或创建它。
      
      如果您知道存储桶已经存在，那么这可能很有用，以尽量减少rclone进行的事务数。
      
      如果使用的用户没有创建存储桶的权限，则可能也需要这样做。v1.52.0之前的版本由于错误而会默默通过此操作。

   --no-head
      如果设置，不会对上传的对象进行HEAD操作以检查完整性。
      
      如果设置了该选项，那么如果rclone在PUT操作后接收到200 OK消息，它将假设对象已经正确上传。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传相同
      - 大小与上传相同
      
      它从单个分块PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传的日期
      
      对于多部分上传，将不会读取这些项目。
      
      如果上传长度未知的源对象，则rclone将执行HEAD请求。
      
      设置此标志会增加未检测到的上传故障的机会，特别是大小不正确的上传，因此在正常使用中不建议使用此标志。实际上，即使使用此标志，检测到的上传故障的机会也非常小。

   --no-head-object
      如果设置，不会在GET操作之前执行HEAD操作。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区（例如分块上传）的上载操作将使用内存池进行分配。
      此选项用于控制多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP2使用。
      
      当您使用s3后端时，目前存在一个未解决的问题，与HTTP/2有关。HTTP/2默认启用s3后端，但可以在此处禁用。此问题解决后将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      自定义的下载终结点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据的传出流量成本更低。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      此值应为true、false或不设置，以使用提供程序的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject来上传单个分块对象
      
      如果为false，则rclone将使用AWS SDK中的PutObject来上传对象。
      
      Rclone版本< 1.59使用预签名请求来上传单个分块对象，将此标志设置为true将重新启用此功能。除非特殊情况或测试，否则不应该需要这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间点的文件版本。
      
      参数应为日期：“2006-01-02”、日期时间“2006-01-02
      15:04:05”或那么久之前的持续时间，例如“100d”或“1h”。
      
      请注意，在使用此功能时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则对gzip编码的对象进行解压缩。
      
      可以使用“Content-Encoding: gzip”将对象上传到S3。通常，rclone会将这些文件下载为压缩对象。
      
      如果设置了此标志，则rclone将在接收到带有“Content-Encoding: gzip”的文件时对其进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能gzip对象，请设置此标志。
      
      通常，提供程序在下载时不会更改对象。如果一个对象没有使用`Content-Encoding: gzip`进行上传，那么在下载时就不会设置该标志。
      
      不过，某些提供程序（例如Cloudflare）可能会gzip对象，即使它们没有使用“Content-Encoding: gzip”进行上传。
      
      这可能导致收到如下错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了带有Content-Encoding: gzip和分块传输编码的对象，则rclone将实时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供程序的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      抑制设置和读取系统元数据。

选项:
   --access-key-id value      AWS的访问密钥ID。 [$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的预设ACL。 [$ACL]
   --endpoint value           IONOS S3对象存储的终结点。 [$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --region value             您的存储桶将被创建和数据存储的区域。 [$REGION]
   --secret-access-key value  AWS的秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预设ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的文件大小。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则对gzip编码的对象进行解压缩。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             自定义的下载终结点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表检索的大小（用于每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0自动选择。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         单个分块上传的最大分块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，不会尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，不会对上传的对象进行HEAD操作以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，不会在GET操作之前执行HEAD操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             抑制设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       并发用于分块上传的块数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的文件大小。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject来上传单个分块对象 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示指定时间点的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}