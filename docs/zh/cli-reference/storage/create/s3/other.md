# 任何其他兼容S3的提供商

{% code fullWidth="true" %}
```
命令名称：
  singularity storage create s3 other - 任何其他S3兼容提供商

使用方法：
  singularity storage create s3 other [command options] [arguments...]

说明：
  --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅当access_key_id和secret_access_key为空时才适用。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

  --access-key-id
      AWS Access Key ID。

      留空以进行匿名访问或运行时凭证。

  --secret-access-key
      AWS Secret Access Key (密码)。
      
      留空以进行匿名访问或运行时凭证。

  --region
      要连接的区域。
      
      如果您使用的是S3克隆，并且您没有区域，则留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名无效时使用此选项。
         |                    | 例如，pre Jewel/v10 CEPH。

  --endpoint
      S3 API的端点。
      
      使用S3克隆时必填。

  --location-constraint
      位置约束-必须设置以匹配区域。
      
      如果不确定，请留空。仅在创建bucket时使用。

  --acl
      创建bucket、存储或复制对象时使用的预定义访问控制策略（Canned ACL）。
      
      此ACL用于创建对象，并且如果bucket_acl未设置，则还用于创建桶。
      
      获取更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，当服务器端复制对象时，S3不会复制源中的ACL，而是写入一份新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl：头，并且将使用默认值（private）。

  --bucket-acl
      创建bucket时使用的预定义访问控制策略（Canned ACL）。
      
      获取更多信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，仅当创建bucket时才应用此ACL。如果未设置，则使用"acl"。
      
      如果"acl"和"bucket_acl"都是空字符串，则不会添加X-Amz-Acl：头，并且将使用默认值（private）。

      示例：
         | private            | 所有者具有FULL_CONTROL。
         |                    | 没有其他人有访问权限（默认值）。
         | public-read        | 所有者具有FULL_CONTROL。
         |                    | AllUsers组获取读取访问权限。
         | public-read-write  | 所有者具有FULL_CONTROL。
         |                    | AllUsers组获取读写访问权限。
         |                    | 不建议在bucket上授予此权限。
         | authenticated-read | 所有者具有FULL_CONTROL。
         |                    | AuthenticatedUsers组获取读取访问权限。

  --upload-cutoff
      切换到分块上传的文件大小截断点。
      
      大于此大小的文件将使用chunk_size分块上传。
      最小值为0，最大值为5 GiB。

  --chunk-size
      用于上传的分块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如“rclone rcat”中的文件或通过“rclone mount”或Google Photos或Google Docs上传的文件），将使用此分块大小进行多部分上传。
      
      请注意，每个传输会缓冲内存中"--s3-upload-concurrency"块的大小。
      
      如果您正在通过高速链路传输大文件，并且具有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件以保持在10000个块限制以下时，Rclone会自动增加分块大小。
      
      未知大小的文件会使用配置的chunk_size进行上传。由于默认的分块大小为5 MiB，并且最多可以有10000个块，这意味着默认情况下您可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，则需要增加chunk_size。
      
      增加分块大小会降低"-P"标志显示的进度统计的准确性。当AWS SDK将块缓冲到内存时，Rclone将分块视为已发送，但实际上可能仍在上传。较大的块大小意味着更大的AWS SDK缓冲区和进度报告与实际情况的更大偏差。

  --max-upload-parts
      多部分上传的最大部分数。
      
      此选项定义执行多部分上传时要使用的最大多部分块数。
      
      如果某个服务不支持AWS S3规范中的10000个块，则可以设置此选项。
      
      当上传已知大小的大文件以保持在此块数限制以下时，Rclone会自动增加分块大小。
      

  --copy-cutoff
      切换到分块复制的文件大小截断点。
      
      需要服务器端复制的大于此大小的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

  --disable-checksum
      不在对象元数据中存储MD5校验和。
      
      通常，在上传之前，rclone会计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查很有用，但对于大文件来说可能会导致长时间的上传延迟。

  --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth=true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量值为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

  --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果该环境变量也未设置）。
      

  --session-token
      AWS会话token。

  --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您在高速链接上上传大量大型文件，并且这些上传没有充分利用带宽，则增加此值可能有助于加快传输速度。

  --force-path-style
      如果设置为true，则使用路径样式访问；如果设置为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，如果为false，rclone将使用虚拟路径样式。有关详细信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为
      false- rclone将根据提供商设置自动完成此操作。

  --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了该值，则rclone将使用v2身份验证。
      
      仅在v4签名无效时使用此选项，例如，pre Jewel/v10 CEPH。

  --list-chunk
      列出块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”，“ max-items”或“page-size”。
      大多数服务即使请求超过1000个对象，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，请参阅[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用“rgw list buckets max chunk”选项增加此值。
      

  --list-version
      要使用的ListObjects版本：1、2或0用于自动选择。
      
      当S3最初推出时，它只提供了用于枚举存储桶中的对象的ListObjects调用。
      
      但是，在2016年5月引入了ListObjectsV2调用。这是效率更高的方法，应尽可能使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用的列表对象方法。如果猜测错误，则可以在此处手动设置。
      

  --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，在使用控制字符时更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的编码，但您可以在此处覆盖rclone的选择。
      

  --no-check-bucket
      如果设置，则不要尝试检查bucket是否存在或创建。
      
      如果您知道bucket已经存在时，这可能非常有用以最小化rclone的事务数。
      
      如果您使用的用户没有创建bucket的权限，则可能需要进行设置。v1.52.0之前，由于存在bug，此选项会悄悄通过。
      

  --no-head
      如果设置，则不要在获取对象时执行HEAD来检查完整性。
      
      如果要最小化rclone执行的事务数，则可能非常有用。
      
      设置后，这意味着如果rclone在PUT上传对象后接收到200 OK消息，则会假设它已正确上传。
      
      特别是它会假设：
      
      - 元数据，包括修改时间，存储类别和内容类型与上传的值相同
      - 大小与上传的值相同
      
      它从单个部分PUT的响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，这些项目不会被读取。
      
      如果上传一个长度未知的源对象，则rclone **会**执行一个HEAD请求。
      
      设置此标志会增加检测不到的上传失败的几率，特别是错误的大小，因此不建议在常规操作中使用。实际上，即使使用此标志，检测不到的上传失败的几率也非常小。
      

  --no-head-object
      如果设置，则在获取对象时不进行HEAD请求。

  --encoding
      后端的编码方式。
      
      有关详细信息，请参见[概述中的编码部分](/overview/#encoding)。

  --memory-pool-flush-time
      内部内存缓冲区池刷新的时间间隔。
      
      需要额外缓冲区（例如分部分）的上传将使用内存池进行分配。
      此选项控制多久将从池中删除未使用的缓冲区。

  --memory-pool-use-mmap
      是否使用内部内存池中的mmap缓冲区。

  --disable-http2
      禁用S3后端的http2的使用。
      
      s3（特别是minio）后端和HTTP/2目前存在一个未解决的问题。默认情况下，S3后端启用了HTTP/2，但可以在此禁用它。当解决此问题时，将删除此标志。
      
      参考链接：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

  --download-url
      下载的自定义端点。
      这通常设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据提供更便宜的出口流量。

  --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该设置为true、false或留空以使用提供商的默认设置。
      

  --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，rclone将使用来自AWS SDK的PutObject上传对象。
      
      rclone < 1.59版本使用预签名请求来上传单个部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况或测试中，否则不应使用此选项。
      

  --versions
      在目录列表中包括旧版本。

  --version-at
      显示文件的指定时间的版本。
      
      参数应为日期"2006-01-02"、datetime "2006-01-02
      15:04:05"或该很久以前的持续时间，例如"100d"或"1h"。
      
      请注意，在使用此选项时，不允许进行文件写操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参见[time option docs](/docs/#time-option)。

  --decompress
      如果设置，则会对gzip编码的对象进行解压缩。
      
      可以通过设置“Content-Encoding: gzip”来将对象上传到S3。通常，rclone将以压缩对象的形式下载这些文件。
      
      如果设置了此标志，则rclone将在接收到以“Content-Encoding: gzip”编码的文件时进行解压缩。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。
      

  --might-gzip
      如果后端可能会gzip对象，则设置此标志。
      
      通常情况下，提供商在下载对象时不会更改它们。如果对象在上传时没有使用`Content-Encoding: gzip`进行上传，则不会下载时设置它。
      
      但是，一些提供商可能会对对象进行gzip压缩，即使它们在上传时没有使用`Content-Encoding: gzip`（例如Cloudflare）。
      
      这种情况的一个症状是收到如下的错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置了此标志，并且rclone下载带有已设置`Content-Encoding: gzip`和分块传输编码的对象，则rclone将即时解压缩该对象。
      
      如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的值，但您可以在此处覆盖rclone的选择。
      

  --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
  --access-key-id value        AWS Access Key ID。 [$ACCESS_KEY_ID]
  --acl value                  创建buckets和存储或复制对象时使用的预定义访问控制策略（Canned ACL）。 [$ACL]
  --endpoint value             S3 API的端点。 [$ENDPOINT]
  --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 （默认值：false） [$ENV_AUTH]
  --help, -h                   显示帮助
  --location-constraint value  位置约束-必须设置以匹配区域。 [$LOCATION_CONSTRAINT]
  --region value               要连接的区域。 [$REGION]
  --secret-access-key value    AWS Secret Access Key (密码)。 [$SECRET_ACCESS_KEY]

  高级

  --bucket-acl value               创建buckets时使用的预定义访问控制策略（Canned ACL）。 [$BUCKET_ACL]
  --chunk-size value               用于上传的分块大小。 （默认值："5Mi"） [$CHUNK_SIZE]
  --copy-cutoff value              切换到分块复制的文件大小截断点。 （默认值："4.656Gi"） [$COPY_CUTOFF]
  --decompress                     如果设置，则会对gzip编码的对象进行解压缩。 （默认值：false） [$DECOMPRESS]
  --disable-checksum               不在对象元数据中存储MD5校验和。 （默认值：false） [$DISABLE_CHECKSUM]
  --disable-http2                  禁用S3后端的http2的使用。 （默认值：false） [$DISABLE_HTTP2]
  --download-url value             下载的自定义端点。 [$DOWNLOAD_URL]
  --encoding value                 后端的编码方式。 （默认值："Slash,InvalidUtf8,Dot"） [$ENCODING]
  --force-path-style               如果设置为true，则使用路径样式访问；如果设置为false，则使用虚拟主机样式访问。 （默认值：true） [$FORCE_PATH_STYLE]
  --list-chunk value               列出块的大小（每个ListObject S3请求的响应列表）。 （默认值：1000） [$LIST_CHUNK]
  --list-url-encode value          是否对列表进行URL编码：true/false/unset （默认值："unset"） [$LIST_URL_ENCODE]
  --list-version value             要使用的ListObjects版本：1、2或0用于自动选择。 （默认值：0） [$LIST_VERSION]
  --max-upload-parts value         多部分上传的最大部分数。 （默认值：10000） [$MAX_UPLOAD_PARTS]
  --memory-pool-flush-time value   内部内存缓冲区池刷新的时间间隔。 （默认值："1m0s"） [$MEMORY_POOL_FLUSH_TIME]
  --memory-pool-use-mmap           是否使用内部内存池中的mmap缓冲区。 （默认值：false） [$MEMORY_POOL_USE_MMAP]
  --might-gzip value               设置此项如果后端可能会gzip对象。 （默认值："unset"） [$MIGHT_GZIP]
  --no-check-bucket                如果设置，则不要尝试检查bucket是否存在或创建。 （默认值：false） [$NO_CHECK_BUCKET]
  --no-head                        如果设置，则不要在获取对象时执行HEAD来检查完整性。 （默认值：false） [$NO_HEAD]
  --no-head-object                 如果设置，则在获取对象时不进行HEAD请求。 （默认值：false） [$NO_HEAD_OBJECT]
  --no-system-metadata             禁止设置和读取系统元数据 （默认值：false） [$NO_SYSTEM_METADATA]
  --profile value                  共享凭证文件中要使用的配置文件。 [$PROFILE]
  --session-token value            AWS会话token。 [$SESSION_TOKEN]
  --shared-credentials-file value  共享凭证文件的路径。 [$SHARED_CREDENTIALS_FILE]
  --upload-concurrency value       多部分上传的并发数。 （默认值：4） [$UPLOAD_CONCURRENCY]
  --upload-cutoff value            切换到分块上传的文件大小截断点。 （默认值："200Mi"） [$UPLOAD_CUTOFF]
  --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 （默认值："unset"） [$USE_MULTIPART_ETAG]
  --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 （默认值：false） [$USE_PRESIGNED_REQUEST]
  --v2-auth                        如果为true，则使用v2身份验证。 （默认值：false） [$V2_AUTH]
  --version-at value               显示文件的指定时间的版本。 （默认值："off"） [$VERSION_AT]
  --versions                       在目录列表中包括旧版本。 （默认值：false） [$VERSIONS]

  一般

  --name value  存储的名称（默认值：自动生成）
  --path value  存储的路径

```
{% endcode %}