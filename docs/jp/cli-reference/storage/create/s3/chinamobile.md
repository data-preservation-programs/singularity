# 中国移动Ecloud弹性对象存储(EOS)

```
名称：
   singularity storage create s3 chinamobile - 中国移动Ecloud弹性对象存储(EOS)

用法：
   singularity storage create s3 chinamobile [command options] [arguments...]

描述：
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时有效。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境（环境变量或IAM）获取AWS凭证。

   --access-key-id
      AWS访问密钥ID。
      
      如需匿名访问或运行时凭证，请留空。

   --secret-access-key
      AWS秘密访问密钥（密码）。
      
      如需匿名访问或运行时凭证，请留空。

   --endpoint
      中国移动Ecloud弹性对象存储(EOS) API的终端节点。

      示例：
         | eos-wuxi-1.cmecloud.cn      | 默认终端节点，如果不确定时，可以选择这个。
         |                             | 中国东部（苏州）
         | eos-jinan-1.cmecloud.cn     | 中国东部（济南）
         | eos-ningbo-1.cmecloud.cn    | 中国东部（杭州）
         | eos-shanghai-1.cmecloud.cn  | 中国东部（上海-1）
         | eos-zhengzhou-1.cmecloud.cn | 中国中部（郑州）
         | eos-hunan-1.cmecloud.cn     | 中国中部（长沙-1）
         | eos-zhuzhou-1.cmecloud.cn   | 中国中部（长沙-2）
         | eos-guangzhou-1.cmecloud.cn | 中国南方（广州-2）
         | eos-dongguan-1.cmecloud.cn  | 中国南方（广州-3）
         | eos-beijing-1.cmecloud.cn   | 中国北方（北京-1）
         | eos-beijing-2.cmecloud.cn   | 中国北方（北京-2）
         | eos-beijing-4.cmecloud.cn   | 中国北方（北京-3）
         | eos-huhehaote-1.cmecloud.cn | 中国北方（呼和浩特）
         | eos-chengdu-1.cmecloud.cn   | 西南地区（成都）
         | eos-chongqing-1.cmecloud.cn | 西南地区（重庆）
         | eos-guiyang-1.cmecloud.cn   | 西南地区（贵阳）
         | eos-xian-1.cmecloud.cn      | 西北地区（西安）
         | eos-yunnan.cmecloud.cn      | 云南地区（昆明）
         | eos-yunnan-2.cmecloud.cn    | 云南地区（昆明-2）
         | eos-tianjin-1.cmecloud.cn   | 天津地区（天津）
         | eos-jilin-1.cmecloud.cn     | 吉林地区（长春）
         | eos-hubei-1.cmecloud.cn     | 湖北地区（襄阳）
         | eos-jiangxi-1.cmecloud.cn   | 江西地区（南昌）
         | eos-gansu-1.cmecloud.cn     | 甘肃地区（兰州）
         | eos-shanxi-1.cmecloud.cn    | 山西地区（太原）
         | eos-liaoning-1.cmecloud.cn  | 辽宁地区（沈阳）
         | eos-hebei-1.cmecloud.cn     | 河北地区（石家庄）
         | eos-fujian-1.cmecloud.cn    | 福建地区（厦门）
         | eos-guangxi-1.cmecloud.cn   | 广西地区（南宁）
         | eos-anhui-1.cmecloud.cn     | 安徽地区（淮南）

   --location-constraint
      存储位置限制 - 必须与终端节点匹配。
      
      仅用于创建存储桶。

      示例：
         | wuxi1      | 中国东部（苏州）
         | jinan1     | 中国东部（济南）
         | ningbo1    | 中国东部（杭州）
         | shanghai1  | 中国东部（上海-1）
         | zhengzhou1 | 中国中部（郑州）
         | hunan1     | 中国中部（长沙-1）
         | zhuzhou1   | 中国中部（长沙-2）
         | guangzhou1 | 中国南方（广州-2）
         | dongguan1  | 中国南方（广州-3）
         | beijing1   | 中国北方（北京-1）
         | beijing2   | 中国北方（北京-2）
         | beijing4   | 中国北方（北京-3）
         | huhehaote1 | 中国北方（呼和浩特）
         | chengdu1   | 西南地区（成都）
         | chongqing1 | 西南地区（重庆）
         | guiyang1   | 西南地区（贵阳）
         | xian1      | 西北地区（西安）
         | yunnan     | 云南地区（昆明）
         | yunnan2    | 云南地区（昆明-2）
         | tianjin1   | 天津地区（天津）
         | jilin1     | 吉林地区（长春）
         | hubei1     | 湖北地区（襄阳）
         | jiangxi1   | 江西地区（南昌）
         | gansu1     | 甘肃地区（兰州）
         | shanxi1    | 山西地区（太原）
         | liaoning1  | 辽宁地区（沈阳）
         | hebei1     | 河北地区（石家庄）
         | fujian1    | 福建地区（厦门）
         | guangxi1   | 广西地区（南宁）
         | anhui1     | 安徽地区（淮南）

   --acl
      在创建存储桶和存储对象或复制对象时使用的预定义ACL。
      
      此ACL用于创建对象，如果未设置 bucket_acl ，则用于创建存储桶。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      请注意，当服务器端复制对象时，S3不会从源复制ACL，而是写入新的ACL。
      
      如果acl为空字符串，则不会添加任何 X-Amz-Acl: 标头，默认将使用默认（private）。

   --bucket-acl
      在创建存储桶时使用的预定义ACL。
      
      有关更多信息，请访问 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      如果bucket_acl未设置，则仅在创建存储桶时应用.
      
      如果 acl 和 bucket_acl 都为空字符串，则不会添加任何 X-Amz-Acl:
      标头，默认将使用默认（private）。

      示例：
         | private            | 所有者获取FULL_CONTROL。
         |                    | 无其他用户具有访问权限（默认）。
         | public-read        | 所有者获取FULL_CONTROL。
         |                    | AllUsers组获取READ访问权限。
         | public-read-write  | 所有者获取FULL_CONTROL。
         |                    | AllUsers组获取读写访问权限。
         |                    | 一般不推荐在存储桶上授予此权限。
         | authenticated-read | 所有者获取FULL_CONTROL。
         |                    | AuthenticatedUsers组获取READ访问权限。

   --server-side-encryption
      存储此对象时使用的服务器端加密算法。

      示例：
         | <unset> | 无
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用SSE-C，则存储此对象时使用的服务器端加密算法。

      示例：
         | <unset> | 无
         | AES256  | AES256

   --sse-customer-key
      如需使用SSE-C，您可以提供用于加密/解密数据的秘密加密密钥。
      
      或者，您可以提供--sse-customer-key-base64。

      示例：
         | <unset> | 无

   --sse-customer-key-base64
      如果使用SSE-C，您必须以base64格式提供用于加密/解密数据的秘密加密密钥。
      
      或者，您可以提供--sse-customer-key。

      示例：
         | <unset> | 无

   --sse-customer-key-md5
      如果使用SSE-C，您可以提供秘密加密密钥的MD5校验和（可选）。
      
      如果留空，则会自动从提供的sse_customer_key计算得出。

      示例：
         | <unset> | 无

   --storage-class
      存储新对象时要使用的存储类别。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类别
         | GLACIER     | 归档存储模式
         | STANDARD_IA | 低访问频率存储模式

   --upload-cutoff
      切换到分块上传的上传大小截止值。
      
      任何大于此值的文件将按照chunk_size切分为多个块进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件（例如"rclone rcat"或使用"rclone mount"或谷歌照片或谷歌文档上传）时，将使用此块大小进行分块上传。
      
      请注意，每次传输都会在内存中缓冲--s3-upload-concurrency个此大小的块。
      
      如果您正在通过高速链接传输大文件，并且具有足够的内存，增加此值将加快传输速度。
      
      当上传已知大小的大型文件时，Rclone将自动增加块大小，以保持低于10,000个块的限制。
      
      未知大小的文件使用配置的chunk_size进行上传。由于默认的块大小为5 MiB，并且最多可以有10,000个块，因此默认情况下，您可以流式上传的文件的最大大小是48 GiB。如果希望流式上传更大的文件，则需要增加chunk_size。
      
      增加块大小将降低使用"-P"标志显示的进度统计的准确性。当Rclone将块缓冲到AWS SDK时，Rclone在发送块时将块视为已发送，而事实上可能仍在上传。较大的块大小意味着较大的AWS SDK缓冲区和进度报告与实际情况更偏离。
      

   --max-upload-parts
      分块上传中的最大块数。
      
      此选项定义在执行分块上传时使用的最大多个块数。
      
      当服务不支持AWS S3规范的10,000个块时，此选项可能很有用。
      
      Rclone将自动增加块大小，以保持低于此块数限制（对于已知大小的大文件）。
      

   --copy-cutoff
      切换到分块复制的复制大小截止值。
      
      需要复制的大于此值的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不在对象元数据中存储MD5校验和。
      
      通常，在上传对象之前，rclone将计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件来说，可能会导致开始上传的长时间延迟。

   --shared-credentials-file
      共享凭证文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。如果环境变量为空，则默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭证文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量“AWS_PROFILE”或“default”（如果未设置该环境变量）。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      分块上传的并发数。
      
      这是同时上传的相同文件块的数量。
      
      如果您正在通过高速链接上传少量大文件，并且这些上传未充分使用您的带宽，那么增加此参数可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟主机样式访问。有关更多信息，请参见[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将其设置为false - rclone将根据提供商的设置自动完成此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了此参数，则rclone将使用v2身份验证。
      
      仅在v4签名无法正常工作时使用，例如Jewel/v10 CEPH之前。

   --list-chunk
      展示块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"，"max-items"或"page-size"。
      大多数服务会将响应列表截断为1000个对象，即使请求的更多也是如此。
      在AWS S3中，这是一个全局最大值，不能更改，参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项增加此值。
      

   --list-version
      要使用的ListObjects版本：1、2或自动为0。
      
      S3最初只提供ListObjects调用以枚举存储桶中的对象。
      
      但是，在2016年5月引入了ListObjectsV2调用。这是更高性能的调用，如果可能的话，应使用此调用。
      
      如果设置为默认值0，则rclone将根据设置的提供商猜测要调用哪个ListObjects方法。如果它猜测错误，则可能需要手动在此处设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果可用，则在使用控制字符的文件名时，这更可靠。如果将其设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但是您可以在此处覆盖rclone的选择。
      

   --no-check-bucket
      如果设置，则不会尝试检查存储桶是否存在或创建存储桶。
      
      如果您知道存储桶已存在，您可以尝试最小化rclone执行的事务数，此选项可能很有用。
      
      如果所使用的用户没有创建存储桶的权限，则可能需要使用此选项。在v1.52.0之前，此选项不会引发错误，而是静默通过了。
      

   --no-head
      如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。
      
      如果要最小化rclone的执行事务数，此选项可能很有用。
      
      设置此标志意味着，如果rclone在使用PUT上传对象后收到了200 OK消息，它将假设它已正确上传。
      
      特别是，它将假设：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传时相同
      - 大小与上传时相同
      
      对于单个部分的PUT，它会读取以下内容：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些内容。
      
      如果要上传长度未知的源对象，则rclone **将**执行HEAD请求。
      
      设置此标志会增加未检测到的上传错误的几率，特别是尺寸不正确的几率，因此不建议进行正常操作。实际上，即使设置了此标志，未检测到的上传错误的几率也很小。
      

   --no-head-object
      如果设置，则在获取对象之前不进行HEAD请求。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的频率。
      
      需要更多缓冲区（例如：分块）的上传将使用内存池进行分配。
      此选项控制多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      当前s3（特别是minio）后端存在与HTTP/2相关的悬而未决的问题。该问题默认情况下启用了s3后端的HTTP/2，但是可以在此禁用。此问题解决后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      通常，将其设置为CloudFront CDN URL，因为AWS S3提供通过CloudFront网络下载的数据更便宜。

   --use-multipart-etag
      是否在分块上传中使用ETag进行验证
      
      这应该是true、false或留空以使用提供商的默认值。
      

   --use-presigned-request
      是否使用预签名请求还是PutObject进行单个部分上传
      
      如果设置为false，则rclone将使用AWS SDK中的PutObject来上传对象。
      
      rclone版本<1.59将使用预签名请求来上传单个部分对象，并将此标志设置为true将重新启用该功能。除非特殊情况或测试需要，否则不应该使用此功能。
      

   --versions
      在目录列表中包括旧版本。

   --version-at
      显示文件版本，以指定时间为准。
      
      参数应该是日期“2006-01-02”，日期时间“2006-01-02 15:04:05”或表示很久以前的持续时间
      例如 "100d" 或 "1h"。
      
      请注意，当使用此标志时，不允许进行文件写入操作，因此无法上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，则会解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”将对象上传到S3。通常，rclone将以压缩的对象方式下载这些文件。

      如果设置此标志，则rclone将在接收到以“Content-Encoding: gzip”编码的文件时对其进行解压缩。这意味着rclone无法检查大小和哈希值，但是文件内容将解压缩。
      

   --might-gzip
      如果后端可能对对象进行gzip压缩，请设置此选项。
      
      通常情况下，提供商在下载对象时不会更改对象。如果即使没有使用“Content-Encoding: gzip”上传对象，提供商也会对对象进行gzip压缩（例如Cloudflare）。
      
      这可能会导致收到以下错误：
      
          错误 corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置此标志并且rclone下载了具有设置了Content-Encoding: gzip和分块传输编码的对象，则rclone将即时解压缩对象。
      
      如果将其设置为unset（默认值），则rclone将根据提供商设置选择要应用的内容，但是您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      抑制设置和读取系统元数据


选项：
   --access-key-id value           AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                     在创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value                中国移动Ecloud弹性对象存储列 (EOS) API的终端节点。[$ENDPOINT]
   --env-auth                      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (默认值: false) [$ENV_AUTH]
   --help, -h                      显示帮助
   --location-constraint value     存储位置限制 - 必须与终端节点匹配。[$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS秘密访问密钥（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  存储此对象时使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --storage-class value           存储新对象时要使用的存储类别。[$STORAGE_CLASS]

   进阶

   --bucket-acl value               在创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的复制大小截止值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则会解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储MD5校验和。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问，如果为false，则使用虚拟主机样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               展示块的大小（每个ListObject S3请求的响应列表）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1,2或自动为0。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大块数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲区池刷新的频率。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip压缩，请设置此选项。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不会尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不会对已上传的对象进行HEAD请求以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象之前不进行HEAD请求。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             抑制设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用SSE-C，则存储此对象时使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         如需使用SSE-C，您可以提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用SSE-C，则必须以base64格式提供用于加密/解密数据的秘密加密密钥。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用SSE-C，则可以提供秘密加密密钥的MD5校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       分块上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的上传大小截止值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求还是PutObject进行单个部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               显示文件版本，正如在指定时间那样。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值: false) [$VERSIONS]

   常规

   --name value  存储的名称（默认为自动生成的）
   --path value  存储的路径

```