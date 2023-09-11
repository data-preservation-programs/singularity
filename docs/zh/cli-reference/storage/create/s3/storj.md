# Storj (S3兼容网关)

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 storj - Storj (S3兼容网关)

USAGE:
   singularity storage create s3 storj [命令选项] [参数...]

DESCRIPTION:
   --env-auth
      从运行环境中获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时才生效。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      为空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      为空表示匿名访问或运行时凭证。

   --endpoint
      Storj Gateway的端点。

      示例:
         | gateway.storjshare.io | 全球托管网关

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意，此ACL仅在创建存储桶时应用。
      如果未设置bucket_acl，则将使用"acl"参数。
      如果"acl"和"bucket_acl"均为空字符串，则不会添加"X-Amz-Acl:"头，并使用默认值（private）。

      示例:
         | private            | 拥有者具有FULL_CONTROL权限。
         |                    | 无其他用户有访问权限（默认）。
         | public-read        | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取访问权限。
         | public-read-write  | 拥有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入访问权限。
         |                    | 不建议将此权限授予存储桶。
         | authenticated-read | 拥有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取访问权限。

   --upload-cutoff
      切换为分块上传的文件大小阈值。
      
      大于此大小的文件将采用块大小为chunk_size进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时使用的块大小。
      
      当上传大于upload_cutoff的文件，或者上传未知大小的文件（例如通过"rclone rcat"上传
      或使用"rclone mount"上传或从Google相册或谷歌文档上传），将使用此块大小进行分块上传。
      
      请注意，每次传输在内存中缓冲"--s3-upload-concurrency"块大小的数据。
      
      如果您正使用高速链接传输大文件，并且有足够的内存，增加此参数将加快传输速度。
      
      当上传已知大小的大文件以确保低于10000个块的限制时，Rclone会自动增加块大小。
      
      未知大小的文件将使用配置的chunk_size进行上传。由于默认的块大小为5 MiB，最多可以有
      10000个块，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果要进行更大
      大小的流式上传，您需要增加chunk_size。
      
      增大块大小会降低使用"-P"标志显示的进度统计的准确性。Rclone将在缓冲到AWS SDK的块发送
      时将其视为已发送，而实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和进度
      报告与实际进度相差更大。

   --max-upload-parts
      多部分上传的最大部分数。
      
      此选项定义进行多部分上传时使用的最大多部分块数。
      
      如果服务不支持AWS S3规范的10000块，这将会很有用。
      
      当上传已知大小的大文件以确保低于此块数限制时，Rclone会自动增加块大小。

   --copy-cutoff
      切换为分块复制的文件大小阈值。
      
      需要进行服务器端复制的大于此大小的文件将按此大小进行分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，rclone在上传之前会计算输入的MD5校验和，并将其添加到对象的元数据中。这对于
      数据的完整性检查非常有用，但对于大文件开始上传而导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。
      如果环境变量的值为空，则会默认路径为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      在共享凭据文件中使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此
      变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”或“default”。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传相同文件的块数。
      
      如果您在高速链接上上传少量大文件，并且这些上传未完全利用带宽，
      则增加此数值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认设置），则rclone将使用路径样式访问；
      如果为false，则rclone将使用虚拟主机样式访问。
      有关更多信息，请参阅[the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为
      false，rclone将根据提供商的设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认设置），则rclone将使用v4身份验证。
      如果设置，则rclone将使用v2身份验证。
      
      仅在v4签名无法工作时使用，例如Jewel/v10 CEPH。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为"AWS S3"规范中的"MaxKeys"、“max-items”或“page-size”。
      大多数服务即使请求的对象数多于1000个，也会将响应列表截断为1000个。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      S3最初仅提供了ListObjects调用来列举存储桶中的对象。
      
      然而，在2016年5月引入了ListObjectsV2调用。
      这是高性能的，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据提供商设置猜测应调用哪个列出对象的方法。
      如果猜测错误，则可能需要在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，如果可用，则在文件名中使用控制字符时，这是更可靠的方法。
      如果设置为unset（默认设置），则rclone将根据提供商设置选择应用的方法，
      但您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果您知道存储桶已经存在，这可能对于尽量减少rclone事务数量很有用。
      
      如果使用的用户没有存储桶创建权限，则也可能需要这样做。
      在v1.52.0之前版本中，由于错误，此操作将悄悄地传递。

   --no-head
      如果设置，则不会在获取对象时进行HEAD操作以检查完整性。
      
      如果尝试尽量减少rclone的事务数量，这可能很有用。
      
      设置后，如果rclone在使用PUT上传对象后接收到200 OK消息，它将假设对象被正确上传。
      
      特别地，它将假设：
      
      - 元数据（包括修改时间、存储类和内容类型）与上传一致
      - 大小与上传一致
      
      对于单个部分PUT，它从响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传具有未知长度的源对象，则rclone将执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，
      特别是错误的大小，因此不建议在正常操作中使用。
      实际上，即使在设置此标志后，未检测到的上传失败的机会也非常小。

   --no-head-object
      如果设置，则在获取对象时不会先进行HEAD操作。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅概述中的[encoding section](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP/2使用。
      
      目前，s3（特别是minio）后端的HTTP/2存在未解决的问题。默认情况下，s3后端启用HTTP/2，
      但可以在此处禁用。当问题解决后，此标志将被删除。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义端点。
      通常，此值设置为CloudFront CDN URL，因为AWS S3通过
      CloudFront网络下载的数据具有更低的出口费用。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      此选项应为true、false或不设置以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject上传单一部分对象
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      Rclone < 1.59版本使用预签名请求上传单个部分对象，
      将此标志设置为true将重新启用该功能。
      除非在特殊情况下或进行测试，否则不应该需要此标志。

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间点的文件版本。
      
      参数应为日期（"2006-01-02"）、日期时间（"2006-01-02 15:04:05"）或该时间之前的持续时间，例如"100d"或"1h"。
      
      请注意，当使用此选项时，不允许进行文件写操作，
      因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[time选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以在S3上使用"Content-Encoding: gzip"上传对象。
      通常，rclone将以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则rclone将在接收到"Content-Encoding: gzip"的对象时对其进行解压缩。
      这意味着rclone无法检查大小和散列，但文件内容将被解压缩。

   --might-gzip
      如果后端可能会压缩对象，则设置此标志。
      
      通常，提供商在下载对象时不会更改对象。
      如果一个对象在上传时没有使用“Content-Encoding: gzip”，那么在下载时也不会设置该标志。
      
      但是，一些提供商可能会对对象进行gzip压缩，即使它们未使用“Content-Encoding: gzip”进行上传（例如Cloudflare）。
      
      如果设置了此标志，并且rclone下载了带有设置了"Content-Encoding: gzip"和chunked传输编码的对象，
      则rclone将在传输过程中对对象进行解压缩。
      
      如果设置为非设置值（默认设置），则rclone将根据提供商设置选择将应用什么，但您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --endpoint value           Storj Gateway的端点。[$ENDPOINT]
   --env-auth                 从运行环境中获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。(默认值：false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               上传时使用的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的文件大小阈值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1、2或0表示自动。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会压缩对象，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会在获取对象时进行HEAD操作以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不会先进行HEAD操作。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  在共享凭证文件中使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的文件大小阈值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject上传单一部分对象 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示指定时间点的文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   常规选项

   --name value  存储的名称（默认为自动生成）
   --path value  存储的路径

```
{% endcode %}