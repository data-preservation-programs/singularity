# 华为对象存储服务

{% code fullWidth="true" %}
```
名称:
   singularity storage create s3 huaweiobs - 华为对象存储服务

用法:
   singularity storage create s3 huaweiobs [命令选项] [参数...]

描述:
   --env-auth
      从运行时（环境变量或EC2/ECS元数据，如果没有环境变量）获取AWS凭证。
      
      仅在access_key_id和secret_access_key为空时适用。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      留空以进行匿名访问或运行时凭证。

   --secret-access-key
      AWS Secret访问密钥（密码）。
      
      留空以进行匿名访问或运行时凭证。

   --region
      要连接的区域。- 存储桶创建和数据存储的位置。必须与您的终端点相同。
      

      示例:
         | af-south-1     | AF-约翰内斯堡
         | ap-southeast-2 | AP-曼谷
         | ap-southeast-3 | AP-新加坡
         | cn-east-3      | 中国华东-上海1
         | cn-east-2      | 中国华东-上海2
         | cn-north-1     | 中国华北-北京1
         | cn-north-4     | 中国华北-北京4
         | cn-south-1     | 中国华南-广州
         | ap-southeast-1 | 中国香港
         | sa-argentina-1 | 拉美-布宜诺斯艾利斯1
         | sa-peru-1      | 拉美-利马1
         | na-mexico-1    | 拉美-墨西哥城1
         | sa-chile-1     | 拉美-圣地亚哥2
         | sa-brazil-1    | 拉美-圣保罗1
         | ru-northwest-2 | 俄罗斯-莫斯科2

   --endpoint
      OBS API的终端点。

      示例:
         | obs.af-south-1.myhuaweicloud.com     | AF-约翰内斯堡
         | obs.ap-southeast-2.myhuaweicloud.com | AP-曼谷
         | obs.ap-southeast-3.myhuaweicloud.com | AP-新加坡
         | obs.cn-east-3.myhuaweicloud.com      | 中国华东-上海1
         | obs.cn-east-2.myhuaweicloud.com      | 中国华东-上海2
         | obs.cn-north-1.myhuaweicloud.com     | 中国华北-北京1
         | obs.cn-north-4.myhuaweicloud.com     | 中国华北-北京4
         | obs.cn-south-1.myhuaweicloud.com     | 中国华南-广州
         | obs.ap-southeast-1.myhuaweicloud.com | 中国香港
         | obs.sa-argentina-1.myhuaweicloud.com | 拉美-布宜诺斯艾利斯1
         | obs.sa-peru-1.myhuaweicloud.com      | 拉美-利马1
         | obs.na-mexico-1.myhuaweicloud.com    | 拉美-墨西哥城1
         | obs.sa-chile-1.myhuaweicloud.com     | 拉美-圣地亚哥2
         | obs.sa-brazil-1.myhuaweicloud.com    | 拉美-圣保罗1
         | obs.ru-northwest-2.myhuaweicloud.com | 俄罗斯-莫斯科2

   --acl
      创建存储桶、存储或复制对象时使用的Canned ACL。
      
      此ACL用于创建对象，如果未设置bucket_acl，则还用于创建存储桶。
      
      若要了解更多信息，请访问 https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，此ACL仅在服务器端复制对象时应用，因为S3不会复制源的ACL，而是写入新的ACL。
      
      如果acl为空字符串，则不添加X-Amz-Acl:头，并使用默认（私有）。
      

   --bucket-acl
      创建存储桶时使用的Canned ACL。
      
      请访问 https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/acl-overview.html#canned-acl 了解更多信息
      
      请注意，仅在创建存储桶时应用此ACL。如果未设置它，则使用"acl"。
      
      如果acl和bucket_acl为空字符串，则不添加X-Amz-Acl:头，并使用默认（私有）。
      

      示例:
         | private            | 拥有者获得FULL_CONTROL。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 拥有者获得FULL_CONTROL。
         |                    | AllUsers组获取读访问权限。
         | public-read-write  | 拥有者获得FULL_CONTROL。
         |                    | AllUsers组获取读和写访问权限。
         |                    | 不建议在存储桶上授予此权限。
         | authenticated-read | 拥有者获得FULL_CONTROL。
         |                    | AuthenticatedUsers组获取读访问权限。

   --upload-cutoff
      切换到分块上传的截止值。
      
      大于此值的任何文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff或大小未知的文件（例如，从"rclone rcat"或通过"rclone mount"或Google图片/文档上传）时，
      它们将使用这个块大小作为分块上传。
      
      请注意，每个传输会在内存中缓冲"--s3-upload-concurrency"个这样大小的块。
      
      如果您正在通过高速链路传输大文件，并且具有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件以低于10000个块的速度时，Rclone将自动增加分块大小。
      
      未知大小的文件使用配置的块大小上传。由于默认块大小为5 MiB，最多可以有10,000个块，因此默认情况下可以流式上传的文件的最大大小为48 GiB。
      如果要流式上传更大文件，则需要增加块大小。
      
      增加块大小会降低使用“-P”标志显示的进度统计信息的精确度。Rclone在AWS SDK的群集中，当它被Rclone缓冲时，将块视为已发送，而实际上它可能仍在上传。
      更大的块大小意味着更大的AWS SDK缓冲区和与真实情况更偏离的进度报告。
      

   --max-upload-parts
      分块上传中的最大部件数。
      
      此选项定义了执行分块上传时要使用的最大分块数。
      
      如果服务不支持AWS S3规范的10,000个分块，则此选项可能会很有用。
      
      当上传已知大小的大文件以低于此块数限制的速度时，Rclone将自动增加分块大小。
      

   --copy-cutoff
      切换到分块复制的截止值。
      
      需要进行服务器端复制的大于此值的文件将以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。
      
      通常，rclone会在上传前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件开始上传可能会导致长时间的延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，那么rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境变量为空，它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      在共享凭证文件中使用的配置文件。
      
      如果env_auth = true，那么rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"默认"。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的同一文件的块数。
      
      如果您通过高速链路上传少量大文件，并且这些上传未充分利用带宽，
      那么增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问; 如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问;
      如果为false，则rclone将使用虚拟路径样式。请参阅[the AWS S3文档](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)以获取更多信息。
      
      某些提供商（例如，AWS，阿里云OSS，网易COS或腾讯COS）要求将其设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。
      如果设置，则rclone将使用v2身份验证。
      
      仅当v4签名无法工作时（例如，Jewel/v10之前的CEPH）才使用此选项。

   --list-chunk
      列表块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范的"MaxKeys"、"max-items"或"page-size"。
      大多数服务即使请求的列表项超过1000个，也会将响应列表截断为1000个对象。
      在AWS S3中，这是一个全局的最大值，并且无法更改，请参阅[AWS S3文档](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项增加它。
      

   --list-version
      要使用的ListObjects的版本：1、2或0表示自动。
      
      当S3最初发布时，它只提供了ListObjects调用以枚举存储桶中的对象。
      
      但是，从2016年5月开始，引入了ListObjectsV2调用。这样做性能更高，应尽可能使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个列表对象方法。
      如果猜错，则可以在此处手动设置它。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果可用，则在文件名中使用控制字符时，URL编码更可靠。
      如果设置为未设置（默认值），则rclone将根据提供商设置选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建它。
      
      如果知道存储桶已经存在时，这可能非常有用，可以尽量减少rclone执行的事务数量。
      
      如果使用的用户没有创建存储桶的权限，则可能需要该选项。 在v1.52.0之前，由于一个错误，此选项将被静默地传递。

   --no-head
      如果设置，则不对上载的对象进行HEAD以检查完整性。
      
      这可能在尽量减少rclone执行的事务数量时很有用。
      
      设置它意味着如果rclone在使用PUT上传对象后收到200 OK消息，则它将假设成功上传。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类和内容类型与上传一致
      - 大小与上传一致
      
      它从单个部分PUT的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部件上传，不会读取这些项。
      
      如果上传一个未知长度的源对象，则rclone **会**做一个HEAD请求。
      
      设置此标志将增加未检测到的上传失败的机会，
      特别是不正确的大小，因此不推荐正常运行时使用它。实际上，即使使用此标志，未检测到的上传失败的机会也很小。
      

   --no-head-object
      如果设置，则在GET对象时不执行HEAD。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外缓冲区（例如分块）的上传将使用内存池进行分配。
      此选项控制多久将未使用的缓冲区从池中删除。

   --memory-pool-use-mmap
      是否在内部内存缓冲池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在未解决的问题与HTTP/2有关。
      使用默认情况下，s3后端启用HTTP/2，但可以在这里禁用。
      解决问题后，将删除此标志。
      
      请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端点。
      通常将其设置为CloudFront CDN URL，因为AWS S3提供了通过CloudFront网络下载的更便宜的出口流量。

   --use-multipart-etag
      分块上传中是否使用ETag进行验证
      
      这应该设置为true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求还是PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject来上传对象。
      
      Rclone的版本< 1.59使用预签名请求来上传单部分对象，将此标志设置为true将重新启用此功能。
      除非在特殊情况下或用于测试，否则不应该需要这样做。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间的文件版本。
      
      参数应为日期（例如"2006-01-02"）、日期时间（例如"2006-01-02 15:04:05"）或早到那么久的持续时间，例如"100d"或"1h"。
      
      请注意，使用此选项时不允许文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则将对gzip编码的对象进行解压缩。
      
      可以使用"Content-Encoding: gzip"将对象上传到S3。通常，rclone会将这些文件作为压缩对象下载。
      
      如果设置了此标志，则rclone将在收到的时候解压这些文件，并使用"Content-Encoding: gzip"。这意味着rclone无法检查大小和哈希，但文件内容将被解压缩。
      

   --might-gzip
      如果可能，设置此标志可以gzip对象。
      
      通常，提供商在下载对象时不会更改对象。如果一个对象没有使用 `Content-Encoding: gzip` 上传，那么在下载时它就不会设置。
      
      但是，即使不是使用 `Content-Encoding: gzip` 进行上传（例如Cloudflare），某些提供商也可能对对象进行gzip操作。
      
      设置此标志后，rclone会在下载具有设置了 `Content-Encoding: gzip` 和分块传输编码的对象时即时解压缩对象。
      
      如果设置为unset（默认值），则rclone将根据提供商的设置选择要应用的内容，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


选项:
   --access-key-id value      AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                创建存储桶和存储或复制对象时使用的Canned ACL。[$ACL]
   --endpoint value           OBS API的终端点。[$ENDPOINT]
   --env-auth                 从运行时（环境变量或EC2/ECS元数据，如果没有环境变量）获取AWS凭证。 (默认值: false) [$ENV_AUTH]
   --help, -h                 显示帮助信息
   --region value             要连接的区域。- 存储桶创建和数据存储的位置。必须与您的终端点相同。[$REGION]
   --secret-access-key value  AWS密钥访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的Canned ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止值。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将对gzip编码的对象进行解压缩。 (default: false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP/2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             自定义下载终端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问; 如果为false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表块的大小（每个ListObject S3请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects的版本：1、2或0表示自动。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大部件数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内存缓冲池刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存缓冲池中使用mmap缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               设置此标志以gzip对象（如果可能）。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不对上载的对象进行HEAD以检查完整性。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在GET对象时不执行HEAD。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  在共享凭证文件中使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       分块上传的并发数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       分块上传中是否使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求还是PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (default: false) [$V2_AUTH]
   --version-at value               显示指定时间的文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]

   一般选项

   --name value  存储名称（默认值: 自动生成）
   --path value  存储路径

```
{% endcode %}