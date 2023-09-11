# StackPath对象存储

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 stackpath - StackPath对象存储

USAGE:
   singularity storage update s3 stackpath [command options] <name|id>

DESCRIPTION:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据）。
      
      如果access_key_id和secret_access_key为空，则应用此项。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --region
      连接的区域。
      
      如果使用S3克隆并且没有区域，请留空。

      示例：
         | <unset>            | 如果不确定，请使用此项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅当v4签名无效时使用此项。
         |                    | 例如，Jewel/v10 CEPH之前版本。

   --endpoint
      StackPath对象存储的终端节点。

      示例：
         | s3.us-east-2.stackpathstorage.com    | 东部区域终端节点
         | s3.us-west-1.stackpathstorage.com    | 西部区域终端节点
         | s3.eu-central-1.stackpathstorage.com | 欧洲区域终端节点

   --acl
      创建存储桶、存储或复制对象时使用的预设ACL。
      
      此ACL用于创建对象，并且如果未设置bucket_acl，则用于创建存储桶。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，在服务器端复制对象时，此ACL将应用，因为S3不会复制源的ACL，而是编写新的ACL。
      
      如果acl为空字符串，则不会添加X-Amz-Acl:标头，并且将使用默认的（私有）。

   --bucket-acl
      创建存储桶时使用的预设ACL。
      
      有关更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。
      
      请注意，此ACL仅在创建存储桶时应用。如果未设置，则使用“acl”。
      
      如果“acl”和“bucket_acl”为空字符串，则不会添加X-Amz-Acl:
      标头，并且将使用默认的（私有）。

      示例：
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 没有其他人具有访问权限（默认）。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers群组具有读取权限。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers群组具有读取和写入权限。
         |                    | 不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers群组具有读取权限。

   --upload-cutoff
      切换到分块上传的截止值。
      
      大于此值的文件将以chunk_size的块进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或具有未知大小的文件（例如“rclone rcat”生成的文件，或使用“rclone mount”或Google相册或Google文档上传的文件）时，将使用此分块大小进行分块上传。
      
      请注意，每个传输会将"--s3-upload-concurrency"个具有此大小的块缓冲到内存中。
      
      如果正在通过高速链接传输大文件，并且内存足够，增加此值将加快传输速度。
      
      Rclone将根据需要自动增加分块大小，以保持在10000块限制以下的大文件进行上传。
      
      未知大小的文件将使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，最多可以有10000块，这意味着默认情况下您可以流式传输的文件的最大大小为48 GiB。如果要流式传输更大的文件，则需要增加chunk_size。
      
      增加chunk_size会影响使用“-P”标志显示的进度统计的准确性。当AWS SDK将所发送的块缓冲到内存中时，Rclone将其视为已发送，而实际上可能仍在上传。更大的块大小意味着更大的AWS SDK缓冲区和与实际情况偏离更大的进度报告。

   --max-upload-parts
      多部分上传中的最大部分数量。
      
      此选项定义在执行多部分上传时使用的最大多部分块数。
      
      如果某个服务不支持AWS S3的10000个块的规范，则这将非常有用。
      
      当上传已知大小的大文件时，Rclone将自动增加块大小，以保持低于此块数限制。
      

   --copy-cutoff
      切换到分块复制的截止值。
      
      任何大于该大小且需要进行服务器端复制的文件都将被分块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。
      
      通常，rclone会在上传之前计算输入数据的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间的上传延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量的值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用的配置文件。
      
      如果为空，它将默认为环境变量“AWS_PROFILE”或“default”（如果该环境变量也未设置）。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是并发上传的相同文件的块数。
      
      如果您正在通过高速链接上传较少数量的大文件，并且这些上传没有充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认）则rclone将使用路径样式访问，如果为false则rclone将使用虚拟路径样式。详细信息请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将其设置为false-基于提供商的设置，rclone将自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果这是false（默认值），则rclone将使用v4身份验证。如果设置了此项，则rclone将使用v2身份验证。
      
      仅在v4签名无效时使用此项，例如，早于Jewel/v10 CEPH版本。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求的数量多于1000个，也会将响应列表截断为1000个对象。在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。
      
      当S3最初发布时，它只提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月，引入了ListObjectsV2调用。这个调用的性能更高，如果有可能，应该使用它。
      
      如果设置为默认值0，则rclone将根据提供商设置猜测要调用哪种列举对象方法。如果猜测错误，那么可以在此处手动设置。
      

   --list-url-encode
      是否将列表URL进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果可用，使用控制字符时更可靠。如果设置为unset（默认值），则rclone将选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查桶是否存在或创建它。
      
      如果知道桶已经存在，这可以有助于最小化rclone执行的事务数。
      
      如果使用的用户没有创建存储桶的权限，则可能需要此选项。在v1.52.0之前，由于存在错误，这将默默地通过。
      

   --no-head
      如果设置，则不为已上传的对象执行HEAD请求以检查完整性。
      
      这可以有助于最小化rclone执行的事务数。
      
      设置后，如果rclone在使用PUT上传对象后收到200 OK消息，则会认为它已经正确上传。
      
      特别是，它将假定：
      
      - metadata，包括modtime，存储类别和内容类型与上传的文件相同
      - 大小与上传的文件相同
      
      它从单个部分PUT的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传长度未知的源对象，则rclone**将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传失败的机会，尤其是大小不正确。因此，不建议在正常操作中使用它。实际上，在使用此标志的情况下，未检测到上传失败的机会非常小。
      

   --no-head-object
      如果设置，则在获取对象时不执行HEAD请求以进行GET。

   --encoding
      后端使用的编码。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区（例如多部分）的上传将使用内存池进行分配。
      此选项控制多久会从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前s3（特别是minio）后端存在一个未解决的问题，与HTTP/2有关。默认情况下，S3后端启用HTTP/2，但可以在此禁用。该问题解决后，此标志将被删除。
      
      请参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      自定义下载的终端节点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据价格更便宜。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该为true、false或保留默认设置以使用提供程序的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject来上传对象。
      
      rclone < 1.59版本使用预签名请求上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或用于测试，否则不应该用于正常操作。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      按照指定时间显示文件版本。
      
      参数应为日期（"2006-01-02"），日期时间（"2006-01-02 15:04:05"）或到那时为止的持续时间，例如"100d"或"1h"。
      
      请注意，使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”设置将对象上传到S3。通常，rclone会以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则rclone将以它们接收到的形式解压缩这些带有“Content-Encoding: gzip”的文件。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此项。
      
      通常，提供者在下载对象时不会更改对象。如果即使在上传时没有使用“Content-Encoding: gzip”也会对对象进行gzip压缩（例如Cloudflare），则对象可能已由提供者进行gzip压缩。
      
      这种情况的一个症状是收到类似的错误
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载了设置了Content-Encoding: gzip并带有分块传输编码的对象，则rclone将实时解压缩对象。
      
      如果设置为unset（默认）则rclone将选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      抑制设置和读取系统元数据


OPTIONS:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的预设ACL。[$ACL]
   --endpoint value           StackPath对象存储的终端节点。[$ENDPOINT]
   --env-auth                 从运行时获取AWS凭证（环境变量或EC2/ECS元数据）。（默认值：false）[$ENV_AUTH]
   --help, -h                 显示帮助
   --region value             连接的区域。[$REGION]
   --secret-access-key value  AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预设ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 （默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 （默认值："4.656Gi"）[$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 （默认值：false）[$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 （默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 （默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端使用的编码。 （默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 （默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否将列表URL进行URL编码：true/false/unset （默认值："unset"）[$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0表示自动。 （默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数量。 （默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 （默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 （默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此项，以防后端可能对对象进行gzip压缩。 （默认值："unset"）[$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查桶是否存在或创建它。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不为已上传的对象执行HEAD请求以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象时不执行HEAD请求以进行GET。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             抑制设置和读取系统元数据 （默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。 （默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 （默认值："200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 （默认值："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 （默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false）[$V2_AUTH]
   --version-at value               按照指定时间显示文件版本。 （默认值："off"）[$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 （默认值：false）[$VERSIONS]

```
{% endcode %}