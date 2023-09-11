# Cloudflare R2存储

{% code fullWidth="true" %}
```
名称:
   singularity storage update s3 cloudflare - Cloudflare R2存储

用法:
   singularity storage update s3 cloudflare [命令选项] <名称|ID>

描述:
   --env-auth
      从运行环境获取AWS凭据（如果访问密钥ID和访问密钥为空，则从环境变量或EC2 / ECS元数据获取凭据）。

      这仅适用于访问密钥ID和密钥为空的情况下。

      示例:
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境变量（env vars或IAM）获取AWS凭据。

   --access-key-id
      AWS访问密钥ID。

      留空以进行匿名访问或运行时凭据。

   --secret-access-key
      AWS Secret Access Key（密码）。

      留空以进行匿名访问或运行时凭据。

   --region
      要连接的区域。

      示例:
         | auto | R2存储桶会自动分布在Cloudflare的数据中心以获得低延迟。

   --endpoint
      S3 API的终端。

      使用S3克隆时需要。

   --bucket-acl
      在创建存储桶时使用的预设ACL。

      有关详细信息，请参见https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，当创建存储桶时只应用此ACL。如果未设置，则使用“acl”。
      
      如果“acl”和“bucket_acl”为空字符串，则不会添加X-Amz-Acl: header，并将使用默认值（private）。

      示例:
         | private            | 拥有者具有完全控制权限。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 拥有者具有完全控制权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 拥有者具有完全控制权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者具有完全控制权限。
         |                    | AuthenticatedUsers组有读取权限。

   --upload-cutoff
      切换到分块上传的上传截止点。

      大于此大小的任何文件将以chunk_size的大小进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。

      当上传大于upload_cutoff的文件或大小未知的文件（例如来自“rclone rcat”或使用“rclone mount”或Google照片或Google文档上传的文件）时，
      将使用此块大小进行分块上传。

      请注意，“--s3-upload-concurrency”每个传输在内存中缓冲的块大小。
      
      如果您正在高速链路上传输大文件并且您有足够的内存，增加这个值将加快传输速度。
      
      rclone将根据文件大小自动增加块大小，以保持在10,000个块的限制之下。
      
      大小未知的文件以配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，并且最多可以有10,000个块，因此默认情况下您可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用“-P”标志显示的进度统计的准确性。当AWS SDK缓冲块大小时，rclone会将块视为已发送，而实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和进度报告与实际情况更有差异。

   --max-upload-parts
      允许的分块上传的最大块数。

      此选项定义在执行分块上传时要使用的最大块数。

      如果某个服务不支持AWS S3规范的10,000个块，则此选项可能很有用。

      rclone将根据文件大小自动增加块大小，以保持在此块数限制之下。

   --copy-cutoff
      切换到分块复制的复制截止点。

      大于此大小且需要在服务器端复制的任何文件都将按此大小进行分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。

      通常，rclone将在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件来说可能会导致较长的延迟以开始上传。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth = true，则rclone可以使用共享凭据文件。

      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。
      如果环境变量的值为空，则默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，则rclone可以使用共享凭证文件。
      此变量控制在该文件中使用哪个配置文件。

      如果为空，则默认为环境变量“AWS_PROFILE”或“default”（如果也未设置该环境变量）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传相同文件的块的数量。

      如果您正在高速链路上上传少量大文件，并且这些上传未充分利用您的带宽，那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。

      如果为true（默认值），则rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟主机样式访问。有关详细信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      一些提供商（例如AWS、阿里云OSS、网易COS或腾讯COS）要求此设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2认证。

      如果为false（默认值），则rclone将使用v4认证。
      如果它被设置，则rclone将使用v2认证。

      仅当v4签名无法正常工作时使用，例如旧版/v10 CEPH。

   --list-chunk
      列出的块的大小（每个ListObject S3请求的响应列表）。

      此选项也称为AWS S3规范的“MaxKeys”，“max-items”或“page-size”。
      大多数服务即使请求超过这些大小也会截断响应列表以包含1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的ListObjects版本：1、2或0表示自动选择。

      当S3最初发布时，它只提供了ListObjects调用，用于枚举存储桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。此调用性能更高，应尽可能使用。

      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个ListObjects方法。如果它猜错，则可以在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset。

      某些提供商支持URL编码列表，在使用控制字符的文件名时，这更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。

      如果您知道存储桶已经存在，并且希望最小化rclone的事务数量时，这可能很有用。

      如果您使用的用户没有创建存储桶的权限，则可能需要将其设置为true。在v1.52.0之前，由于该问题，此设置将静默传递。

   --no-head
      如果设置，则不对已上传的对象进行HEAD请求以检查完整性。

      如果您尝试最小化rclone执行的事务数量，则这可能很有用。

      设置后，意味着如果rclone在PUT上传对象后收到200 OK消息，则它将假定对象已正确上传。

      特别是，它将假定：

      - 元数据，包括修改时间，存储类和内容类型与上传的相同
      - 大小与上传的相同

      它从单个部分PUT的响应中读取以下项目：
      - MD5SUM
      - 上传日期

      对于分块上传，不会读取这些项目。

      如果上传一个长度未知的源对象，则rclone **将**执行HEAD请求。

      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确，因此不建议在正常操作中使用此标志。实际上，即使使用此标志，未检测到的上传失败的几率也非常小。

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求。

   --encoding
      后端的编码。

      有关详细信息，请参见[概览中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。

      需要额外缓冲区的上传（例如多部分上传）将使用内存池进行分配。
      此选项控制多久没有使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。

      目前，s3（特别是minio）后端与HTTP/2存在一个无解的问题。默认情况下，s3后端启用HTTP/2，但可以在此处禁用。当问题解决时，将删除此标志。

      参见: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      下载自定义终端。

      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据提供更便宜的流出。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证。

      这应该是true、false或留空以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传。

      如果此设置为false，则rclone将使用AWS SDK的PutObject来上传对象。

      rclone的版本<1.59使用预签名请求来上传单部分对象，并将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应该需要使用此标志。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定的时间显示文件版本。

      参数应为日期，“2006-01-02”，日期时间“2006-01-02 15:04:05”或距离那之前的持续时间，例如“100d”或“1h”。

      请注意，在使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。

      有关有效格式，请参见[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，则会解压缩gzip编码的对象。

      可以使用“Content-Encoding: gzip”将对象上传到S3。
      
      通常，rclone将将这些文件作为压缩对象进行下载。

      如果设置此标志，则rclone在接收到这些文件时将通过“Content-Encoding: gzip”解压缩它们。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能压缩对象，请设置此标志。

      通常，提供程序在下载对象时不会更改对象。如果一个对象在上传时没有使用“Content-Encoding: gzip”，那么在下载时也不会设置它。

      但是，某些提供程序甚至在对象未使用“Content-Encoding: gzip”上传时（例如Cloudflare）也会压缩对象。

      这样做的一个证据是收到如下错误：

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      如果设置此标志并且rclone下载具有设置了Content-Encoding: gzip和分块传输编码的对象，则rclone将即时解压缩对象。

      如果将此设置设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

选项:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --endpoint value           S3 API的终端。[$ENDPOINT]
   --env-auth                 从运行环境获取AWS凭据（如果访问密钥ID和访问密钥为空，则从环境变量或EC2 / ECS元数据获取凭据）。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助
   --region value             要连接的区域。[$REGION]
   --secret-access-key value  AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               在创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的复制截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则会解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载自定义终端。[$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出的块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动选择。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         允许的分块上传的最大块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内存缓冲池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对已上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的上传截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2认证。 (默认值: false) [$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

```
{% endcode %}