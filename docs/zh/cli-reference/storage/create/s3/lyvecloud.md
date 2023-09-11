# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 lyvecloud - Seagate Lyve Cloud

USAGE:
   singularity storage create s3 lyvecloud [command options] [arguments...]

DESCRIPTION:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。

      仅在access_key_id和secret_access_key为空时适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS Access Key ID。

      为空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS Secret Access Key（密码）。

      为空表示匿名访问或运行时凭证。

   --region
      连接的区域。

      如果您使用的是S3 clone并且没有区域，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅当v4签名不起作用时使用此选项。
         |                    | 例如：早期版本的Jewel/v10 CEPH。

   --endpoint
      S3 API的终端点。

      使用S3克隆时需要。

      示例：
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US East 1（弗吉尼亚）
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US West 1（加利福尼亚）
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1（新加坡）

   --location-constraint
      地理位置限制-必须设置为匹配区域。

      如果不确定，请留空。仅在创建存储桶时使用。

   --acl
      创建存储桶、存储或复制对象时使用的预定义ACL。

      此ACL用于创建对象，并且如果未设置bucket_acl，则也用于创建存储桶。

      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，当S3服务器端复制对象时，此ACL被应用，因为S3不复制源中的ACL，而是写入一个新的ACL。

      如果acl是一个空字符串，则不会添加X-Amz-Acl:标题，并且将使用默认（私有）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。

      更多信息请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，仅在创建存储桶时应用此ACL。如果未设置，则使用"acl"。

      如果acl和bucket_acl都为空字符串，则不会添加X-Amz-Acl:标题，并且将使用默认（私有）。

      示例：
         | private            | 所有者拥有FULL_CONTROL权限。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         | public-read-write  | 所有者拥有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 一般情况下不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者拥有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。

   --upload-cutoff
      切换到分块上传的大小。

      大于此大小的文件将分块上传，每块的大小为chunk_size。
      最小值为0，最大值为5GiB。

   --chunk-size
      上传时要使用的块大小。

      当上传大于upload_cutoff或未知大小的文件（例如，使用"rclone rcat"上传的文件或使用"rclone mount"或Google照片或Google文档上传的文件）时，它们将以此块大小作为多部分上传进行上传。

      请注意，“--s3-upload-concurrency”每个传输的chunks块大小是在内存中缓冲的。

      如果您正在高速链接上传输大文件并且有足够的内存，那么增加这个值将加快传输速度。

      当上传已知大小的大文件以保持在10000个chunks限制之下时，Rclone将自动增加块大小。

      未知大小的文件按配置的块大小进行上传。由于默认块大小为5 MiB，并且最多可以有10000个chunks，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果要流式上传更大的文件，则需要增加chunk_size。

      增加块大小会降低使用"-P"标志显示的进度统计数据的准确性。Rclone当SDK发送一个chunk时就会把这个chunk标记为已发送，但实际上它可能仍在上传。更大的块大小意味着更大的SDK缓冲区和进度报告与真实情况更不相符。

   --max-upload-parts
      多部分上传中的最大部分数。

      此选项定义在执行多部分上传时要使用的最大多部分块数。

      如果服务不支持AWS S3规范的10000个chunks，则此选项可能非常有用。

      当正在上传一个已知大小的大文件时，Rclone将自动增加块大小以使其保持在此块数限制之下。

   --copy-cutoff
      切换到多部分复制的大小。

      大于此大小的需要进行服务器端复制的文件将按此大小分块复制。

      最小值为0，最大值为5GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。

      通常情况下，Rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查是很好的，但对于大文件来说可能导致启动上传的时间过长。

   --shared-credentials-file
      共享凭证文件的路径。

      如果env_auth = true，那么Rclone可以使用共享凭证文件。

      如果此变量为空，则Rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量的值为空，则将默认为当前用户的主目录。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭证文件中要使用的配置文件。

      如果env_auth = true，那么Rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。

      如果为空，则默认值为环境变量“AWS_PROFILE”或“default”，如果环境变量也未设置。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。

      同一文件的多个chunks将同时上传。

      如果您正在高速链接上上传少量大文件，并且这些上传未完全利用您的带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径格式访问，如果为false，则使用虚拟主机格式访问。

      如果为true（默认值），则使用路径格式访问；如果为false，则使用虚拟路径格式访问。请参阅[the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)获取更多信息。

      某些提供商（例如：AWS，阿里云OSS，网易COS或腾讯COS）要求将此设置为false，rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。

      如果为false（默认值），则Rclone将使用v4身份验证。如果设置为true，则Rclone将使用v2身份验证。

      仅当v4签名无法正常工作时使用此选项，例如：早期版本的Jewel/v10 CEPH。

   --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”，“max-items”或“page-size”。
      大多数服务在请求多于1000个对象时都会截断响应列表。
      在AWS S3中，这是一个全局最大值，不可更改，请参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项来增加此值。

   --list-version
      要使用的ListObjects版本：1，2或0自动。

      当S3最初推出时，仅提供了用于枚举存储桶中的对象的ListObjects请求。

      然而，从2016年5月开始，引入了ListObjectsV2请求。它具有更高的性能，应尽可能使用。

      如果设置为默认值0，则Rclone将根据设置的提供商猜测要调用哪个list_objects方法。如果它猜测错误，则可能在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset

      某些提供商支持URL编码列表，如果可用，通过使用控制字符访问文件名时，这样更可靠。如果将其设置为unset（默认值），则Rclone将选择如何应用，但您可以在此处覆盖Rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。

      如果您知道存储桶已经存在并尝试尽量减少Rclone的事务数量时，这可能很有用。

      如果使用的用户没有创建存储桶的权限，则也可能需要此选项。在v1.52.0之前，这将被静默通过，因为那时有一个错误。

   --no-head
      如果设置，则不对已上传的对象执行HEAD请求以检查完整性。

      如果要尽量减少Rclone的事务数量，这可能很有用。
      
      设置后，表示如果在PUT上传对象后收到200 OK消息，则Rclone将认为已正确上传对象。

      特别地，它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的一样；
      - 大小与上传的一样；
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项目。
      
      如果上传长度未知的源对象，则Rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，特别是大小不正确，因此不建议在正常操作中使用。实际上，在没有此标志的情况下，发生上传失败的可能性非常小。

   --no-head-object
      如果设置，则在获取对象时不会进行HEAD请求。

   --encoding
      后端的编码方式。

      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。

      需要额外缓冲区（例如分块）的上传将使用内存缓冲池进行分配。
      此选项控制将未使用的缓冲区从内存缓冲池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存缓冲池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP/2使用。

      s3（特别是minio）后端和HTTP/2目前存在未解决的问题。默认情况下，s3后端启用HTTP/2，但可以在这里禁用。解决该问题后将删除此标志。

      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终结点。

      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据时提供更便宜的出口流量。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证。

      这应为true，false或保留设置为提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传。

      如果设置为false，Rclone将使用AWS SDK的PutObject上传对象。

      Rclone < 1.59版本使用预签名请求上传单部分对象，将此标志设置为true将重新启用该功能。这在特殊情况或用于测试之外不应该是必需的。

   --versions
      在目录列表中包含旧版本。

   --version-at
      按指定的时间显示文件版本。

      参数应为日期，"2006-01-02"，日期时间"2006-01-02
      15:04:05"或YYYY-MM-DD格式的时间间隔，例如"100d"或"1h"。

      请注意，在使用此选项时，不允许执行文件写操作，因此无法上传文件或删除文件。

      请参阅[时间选项文档](/docs/#time-option)获取有效格式。

   --decompress
      如果设置，将解压缩gzip编码的对象。

      可以使用"Content-Encoding: gzip"将对象上传到S3。通常情况下，Rclone会下载这些文件作为压缩的对象。

      如果设置了此标志，则Rclone将在接收到这些文件时使用"Content-Encoding: gzip"对其进行解压缩。这意味着Rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能压缩对象，请进行此设置。

      通常情况下，提供者在下载时不会更改对象。如果未使用`Content-Encoding: gzip`上传对象，则在下载时不会设置该项。

      但是，某些提供商即使在未使用`Content-Encoding: gzip`上传对象的情况下，也可能对对象进行压缩（例如Cloudflare）。

      如果设置了此标志，并且Rclone下载了一个具有设置`Content-Encoding: gzip`和分块传输编码的对象，则Rclone将动态解压缩该对象。

      如果设置为unset（默认值），则Rclone将选择如何应用，但您可以在此处覆盖Rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据

OPTIONS:
   --access-key-id value        AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value             S3 API的终端点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值：false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  地理位置限制-必须设置为匹配区域。[$LOCATION_CONSTRAINT]
   --region value               连接的区域。[$REGION]
   --secret-access-key value    AWS Secret Access Key (密码)。[$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               上传时要使用的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分复制的大小。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终结点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径格式访问；如果为false，则使用虚拟路径格式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1，2或0自动。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲池中使用mmap缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能压缩对象，请进行此设置。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对已上传的对象执行HEAD请求以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不会进行HEAD请求。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的大小。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值：false) [$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值：false) [$VERSIONS]

   General

   --name value  存储的名称 (默认值：自动生成的)
   --path value  存储的路径

```