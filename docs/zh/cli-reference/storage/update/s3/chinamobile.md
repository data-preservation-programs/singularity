# 中国移动 Ecloud 弹性对象存储（EOS）

{% code fullWidth="true" %}
```
命令名称：
   singularity storage update s3 chinamobile - 中国移动 Ecloud 弹性对象存储（EOS）

用法：
   singularity storage update s3 chinamobile [命令选项] <名称|ID>

说明：
   --env-auth
      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。

      仅当 access_key_id 和 secret_access_key 为空时适用。

      示例：
         | false | 在下一步中输入 AWS 凭证。
         | true  | 从环境中获取 AWS 凭证（环境变量或 IAM）。

   --access-key-id
      AWS 访问密钥 ID。

      为空表示匿名访问或运行时凭证。

   --secret-access-key
      AWS 秘密访问密钥（密码）。

      为空表示匿名访问或运行时凭证。

   --endpoint
      中国移动 Ecloud 弹性对象存储（EOS）API 的终端节点。

      示例：
         | eos-wuxi-1.cmecloud.cn      | 默认终端节点-如果您不确定，可以选择此项。
         |                             | 中国东部（苏州）
         | eos-jinan-1.cmecloud.cn     | 中国东部（济南）
         | eos-ningbo-1.cmecloud.cn    | 中国东部（杭州）
         | eos-shanghai-1.cmecloud.cn  | 中国东部（上海-1）
         | eos-zhengzhou-1.cmecloud.cn | 中国中部（郑州）
         | eos-hunan-1.cmecloud.cn     | 中国中部（长沙-1）
         | eos-zhuzhou-1.cmecloud.cn   | 中国中部（长沙-2）
         | eos-guangzhou-1.cmecloud.cn | 中国南部（广州-2）
         | eos-dongguan-1.cmecloud.cn  | 中国南部（广州-3）
         | eos-beijing-1.cmecloud.cn   | 中国北部（北京-1）
         | eos-beijing-2.cmecloud.cn   | 中国北部（北京-2）
         | eos-beijing-4.cmecloud.cn   | 中国北部（北京-3）
         | eos-huhehaote-1.cmecloud.cn | 中国北部（呼和浩特）
         | eos-chengdu-1.cmecloud.cn   | 中国西南（成都）
         | eos-chongqing-1.cmecloud.cn | 中国西南（重庆）
         | eos-guiyang-1.cmecloud.cn   | 中国西南（贵阳）
         | eos-xian-1.cmecloud.cn      | 中国西南（西安）
         | eos-yunnan.cmecloud.cn      | 中国云南（昆明）
         | eos-yunnan-2.cmecloud.cn    | 中国云南（昆明-2）
         | eos-tianjin-1.cmecloud.cn   | 中国天津（天津）
         | eos-jilin-1.cmecloud.cn     | 中国吉林（长春）
         | eos-hubei-1.cmecloud.cn     | 中国湖北（襄阳）
         | eos-jiangxi-1.cmecloud.cn   | 中国江西（南昌）
         | eos-gansu-1.cmecloud.cn     | 中国甘肃（兰州）
         | eos-shanxi-1.cmecloud.cn    | 中国山西（太原）
         | eos-liaoning-1.cmecloud.cn  | 中国辽宁（沈阳）
         | eos-hebei-1.cmecloud.cn     | 中国河北（石家庄）
         | eos-fujian-1.cmecloud.cn    | 中国福建（厦门）
         | eos-guangxi-1.cmecloud.cn   | 中国广西（南宁）
         | eos-anhui-1.cmecloud.cn     | 中国安徽（淮南）

   --location-constraint
      区域约束-必须与终端节点匹配。

      仅在创建存储桶时使用。

      示例：
         | wuxi1      | 中国东部（苏州）
         | jinan1     | 中国东部（济南）
         | ningbo1    | 中国东部（杭州）
         | shanghai1  | 中国东部（上海-1）
         | zhengzhou1 | 中国中部（郑州）
         | hunan1     | 中国中部（长沙-1）
         | zhuzhou1   | 中国中部（长沙-2）
         | guangzhou1 | 中国南部（广州-2）
         | dongguan1  | 中国南部（广州-3）
         | beijing1   | 中国北部（北京-1）
         | beijing2   | 中国北部（北京-2）
         | beijing4   | 中国北部（北京-3）
         | huhehaote1 | 中国北部（呼和浩特）
         | chengdu1   | 中国西南（成都）
         | chongqing1 | 中国西南（重庆）
         | guiyang1   | 中国西南（贵阳）
         | xian1      | 中国西南（西安）
         | yunnan     | 中国云南（昆明）
         | yunnan2    | 中国云南（昆明-2）
         | tianjin1   | 中国天津（天津）
         | jilin1     | 中国吉林（长春）
         | hubei1     | 中国湖北（襄阳）
         | jiangxi1   | 中国江西（南昌）
         | gansu1     | 中国甘肃（兰州）
         | shanxi1    | 中国山西（太原）
         | liaoning1  | 中国辽宁（沈阳）
         | hebei1     | 中国河北（石家庄）
         | fujian1    | 中国福建（厦门）
         | guangxi1   | 中国广西（南宁）
         | anhui1     | 中国安徽（淮南）

   --acl
      创建存储桶和存储或复制对象时使用的预设 ACL。

      此 ACL 用于创建对象，如果未设置 bucket_acl，也可用于创建存储桶。

      更多信息请参阅 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意，此 ACL 适用于在 S3 中进行服务器端复制对象，因为 S3 不会复制源中的 ACL，而是写入全新的 ACL。

      如果 acl 是空字符串，则不会添加 X-Amz-Acl: 标头，并且将使用默认值（private）。

   --bucket-acl
      创建存储桶时使用的预设 ACL。

      更多信息请参阅 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意，此 ACL 仅在创建存储桶时使用。如果未设置它，则使用 "acl"。

      如果 "acl" 和 "bucket_acl" 是空字符串，则不会添加 X-Amz-Acl:
      标头，并且将使用默认值（private）。

      示例：
         | private            | 所有者获得 FULL_CONTROL 权限。
         |                    | 没有其他人的访问权限（默认值）。
         | public-read        | 所有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 用户组获得读取权限。
         | public-read-write  | 所有者获得 FULL_CONTROL 权限。
         |                    | AllUsers 用户组获得读取和写入权限。
         |                    | 一般不建议在存储桶上授权此操作。
         | authenticated-read | 所有者获得 FULL_CONTROL 权限。
         |                    | AuthenticatedUsers 用户组获得读取权限。

   --server-side-encryption
      存储此对象时所使用的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用 SSE-C，则存储此对象时所使用的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-customer-key
      若要使用 SSE-C，则可以提供用于加密/解密数据的密钥。

      或者，您可以提供 --sse-customer-key-base64。

      示例：
         | <unset> | None

   --sse-customer-key-base64
      如果使用 SSE-C，则必须以 base64 格式提供用于加密/解密数据的密钥。

      或者，您可以提供 --sse-customer-key。

      示例：
         | <unset> | None

   --sse-customer-key-md5
      如果使用 SSE-C，则可以提供密钥的 MD5 校验和（可选）。

      如果留空，它将从提供的 sse_customer_key 自动计算。

      示例：
         | <unset> | None

   --storage-class
      存储新对象时所使用的存储类。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类
         | GLACIER     | 归档存储模式
         | STANDARD_IA | 低频访问存储模式

   --upload-cutoff
      切换到分块上传的上传截止点。

      任何大于此大小的文件将以 chunk_size 的大小进行分块上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的分块大小。

      上传大于 upload_cutoff 的文件或大小未知的文件（例如来自 "rclone rcat" 或使用 "rclone mount" 或 google
      photos 或 google docs 上传的文件）将使用此分块大小进行分块上传。

      注意，"--s3-upload-concurrency" 每个传输将在内存中缓冲此大小的块。

      如果在高速连接上传输大文件且内存足够，则增加此值可以加快传输速度。

      当上传已知大小的大文件时，rclone 将自动增加分块大小，以保持在 10,000 个分块限制以下。

      未知大小的文件将使用配置的
      分块大小上传。由于默认分块大小为 5 MiB，最多可有
      10,000 个分块，这意味着默认情况下您可以流式传输的文件的最大大小为 48 GiB。如果要流式传输
      较大的文件，则需要增加 chunk_size。

      增加 chunk_size 会降低使用 "-P" 标志查看的进度
      统计的准确性。当状态标识为“已发送”时，rclone 将块视为未发送，实际上传可能仍在进行。
      大的块大小意味着更大的 AWS SDK 缓冲区，进度报告与实际情况越来越不符。

   --max-upload-parts
      分块上传中的最大部分数。

      此选项定义进行分块上传时要使用的最大多部分块数。

      如果某个服务不支持 AWS S3 规范要求的 10,000 个分块，这将非常有用。

      当上传已知大小的大文件时，rclone 将自动增加分块大小，以保持在此块数限制以下。

   --copy-cutoff
      切换到分块复制的复制截止点。

      必须以此大小进行分块复制的文件都比此值大。

      最小值为 0，最大值为 5 GiB。

   --disable-checksum
      不在对象元数据中存储 MD5 校验和。

      通常情况下，在上传之前，rclone 会计算输入的 MD5 校验和，以将其添加到对象的元数据中。这对于数据完整性检查很有用，但是对于大文件来说，开始上传可能需要很长时间。

   --shared-credentials-file
      共享凭据文件的路径。

      如果 env_auth = true，则 rclone 可以使用共享凭据文件。

      如果此变量为空，则 rclone 将查找
      “AWS_SHARED_CREDENTIALS_FILE” 环境变量。如果环境变量的值为空，则默认为当前用户的主目录。

          Linux/OSX：“$HOME/.aws/credentials”
          Windows：“%USERPROFILE%\.aws\credentials”

   --profile
      共享凭据文件中要使用的配置文件。

      如果 env_auth = true，则 rclone 可以使用共享凭据文件。此
      变量控制在文件中使用哪个配置文件。

      如果为空，则默认为环境变量 "AWS_PROFILE" 或 "default"（如果环境变量未设置）。

   --session-token
      AWS 会话 Token。

   --upload-concurrency
      分块上传的并发数。

      这是同时上传的相同文件块的数量。

      如果您使用高速连接上传少量大文件，并且这些上传没有充分利用带宽，那么增加此值可能有助于提高传输速度。

   --force-path-style
      如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。

      如果为真（默认），则 rclone 将使用路径样式访问；如果为 false，则 rclone 将使用虚拟主机样式访问。有关更多信息，请参阅 [AWS S3 文档](https://docs.aws.amazon.com/zh_cn/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。

      某些提供商（例如 AWS、阿里云 OSS、网易云 COS 或腾讯云 COS）要求设置为
      false - rclone 将根据提供商的设置自动执行此操作。

   --v2-auth
      如果为真，则使用 v2 认证。

      如果为 false（默认），则 rclone 将使用 v4 认证。如果设置此值，则 rclone 将使用 v2 认证。

      仅在 v4 签名不起作用时使用，例如在 Jewel/v10 CEPH 之前。

   --list-chunk
      列举数据块的大小（每个 ListObject S3 请求的响应列表）。

      此选项也称为 AWS S3 规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务即使请求更多的对象，也会将响应列表截断为 1000 个对象。
      在 AWS S3 中，这是一个全局最大值，无法更改，请参阅 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在 Ceph 中，可以使用 "rgw list buckets max chunk" 选项增加此值。

   --list-version
      要使用的 ListObjects 版本：1、2 或 0 表示自动选择。

      当 S3 最初推出时，仅提供了 ListObjects 调用以枚举存储桶中的对象。

      但是，在2016年5月引入了 ListObjectsV2 调用。它具有更高的性能，如果可能，请使用此调用。

      如果设置为默认值 0，则 rclone 将根据设置的提供商猜测要调用哪个列表对象方法。如果猜测错误，则可以在此处手动设置它。

   --list-url-encode
      是否对列表进行 URL 编码：true/false/unset

      某些提供商支持 URL 编码列表，如果可用，则在使用控制字符的文件名时，这是更可靠的选择。如果将其设置为 unset（默认值），则 rclone 将根据提供商设置来选择应用的方式，但是您可以在此处覆盖 rclone 的选择。

   --no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建它。

      如果您知道存储桶已经存在，这样做可以帮助您尽可能减少 rclone 执行的事务数量。

      如果使用的用户没有存储桶创建权限，则可能需要使用此选项。v1.52.0 之前的版本由于错误导致了静默通过。

   --no-head
      如果设置，不在上传对象之前对其进行 HEAD 操作以检查完整性。

      这可用于帮助您尽可能减少 rclone 执行的事务数量。

      设置此标志意味着如果 rclone 在 PUT 之后收到 200 OK 消息，则它将认为对象已被正确上传。

      特别是，它将假设：

      - 元数据，包括修改时间、存储类和内容类型与上传的一样
      - 大小与上传的一样

      对于单个部分 PUT，它会从响应中读取以下项：

      - MD5SUM
      - 上传日期

      对于多部分上传，不会读取这些项。

      如果上传了未知长度的源对象，则 rclone 将**会**执行 HEAD 请求。

      设置此标志增加了无法检测到的上传失败的几率，特别是大小不正确，因此不建议在常规操作中使用。实际上，在未设置此标志的情况下，检测到上传失败的机会非常小。

   --no-head-object
      如果设置，则在获取对象之前不执行 HEAD 操作。

   --encoding
      后端的编码。

      有关更多信息，请参阅[概述中的编码部分](https://github.com/rclone/rclone/blob/master/docs/overview.md#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。

      需要额外缓冲区的上传（例如多部分）将使用内存池进行分配。
      此选项控制何时从池中删除未使用的缓冲区。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --disable-http2
      禁用 S3 后端的 http2 使用。

      目前，s3（具体来说是 minio）后端与 HTTP/2 存在一个未解决的问题。默认情况下，s3 后端启用了 HTTP/2，但可以在此禁用。待问题解决后，将删除此标志。

      参见：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      下载的自定义终端节点。
      通常将其设置为 CloudFront CDN URL，因为 AWS S3 提供通过 CloudFront 网络下载数据的更低价格。

   --use-multipart-etag
      是否在分块上传中使用 ETag 进行验证

      这应该是 true、false 或 unset，以使用提供商的默认值。

   --use-presigned-request
      是否在单个部分上传时使用预签名请求或 PutObject。

      如果为 false，则 rclone 将使用 AWS SDK 中的 PutObject 来上传对象。

      版本 1.59 之前的 rclone 使用预签名请求来上传单个部分对象，将此标志设置为 true 将重新启用该功能。除非特殊情况或测试需要，否则不应该使用此选项。

   --versions
      在目录列表中包括旧版本。

   --version-at
      根据指定时间显示文件版本。

      参数应为日期、"2006-01-02"、日期时间 "2006-01-02
      15:04:05" 或很久以前的持续时间，例如 "100d" 或 "1h"。

      请注意，在使用此选项时，不允许进行文件写入操作，因此无法上传文件或删除文件。

      有关有效格式，请参阅[时间选项文档](https://github.com/rclone/rclone/blob/master/docs/content/docs.md#time-option)。

   --decompress
      如果设置，则将对经过 gzip 编码的对象进行解压缩。

      可以使用 "Content-Encoding: gzip" 设置在 S3 中上传对象。通常情况下，rclone 将以压缩的对象形式下载这些文件。

      如果设置了此标志，则 rclone 会在收到它们时对具有 "Content-Encoding: gzip" 的文件进行解压缩。这意味着 rclone 无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行 gzip 压缩，请设置此项。

      通常，在下载时，提供者不会更改对象。如果一个对象在上传时没有使用 `Content-Encoding: gzip` 进行上传，则在下载时也不会设置它。

      但是，某些提供商可能会对对象进行 gzip 压缩，即使它们没有使用 `Content-Encoding: gzip` 进行上传（例如 Cloudflare）。

      如果设置了此标志，并且 rclone 下载了一个带有设置了 `Content-Encoding: gzip` 和分块传输编码的对象，则 rclone 将实时解压缩该对象。

      如果设置为 unset（默认值），则 rclone 将根据提供商的设置来选择应用的方式，但您可以在此处覆盖 rclone 的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value           AWS 访问密钥 ID。 [$ACCESS_KEY_ID]
   --acl value                     创建存储桶和存储或复制对象时使用的预设 ACL。 [$ACL]
   --endpoint value                中国移动 Ecloud 弹性对象存储（EOS）API 的终端节点。 [$ENDPOINT]
   --env-auth                      从运行时获取 AWS 凭证（环境变量或 EC2/ECS 元数据，如果没有环境变量）。 (默认值：false) [$ENV_AUTH]
   --help, -h                      显示帮助信息
   --location-constraint value     区域约束-必须与终端节点匹配。 [$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS 秘密访问密钥（密码）。 [$SECRET_ACCESS_KEY]
   --server-side-encryption value  将此对象存储到 S3 时使用的服务器端加密算法。 [$SERVER_SIDE_ENCRYPTION]
   --storage-class value           存储新对象时要使用的存储类。 [$STORAGE_CLASS]

   高级选项

   --bucket-acl value               创建存储桶时使用的预设 ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的分块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的复制截止点。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将对经过 gzip 编码的对象进行解压缩。 (默认值：false) [$DECOMPRESS]
   --disable-checksum               不在对象元数据中存储 MD5 校验和。 (默认值：false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用 S3 后端的 HTTP/2 使用。 (默认值：false) [$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为 true，则使用路径样式访问；如果为 false，则使用虚拟主机样式访问。 (默认值：true) [$FORCE_PATH_STYLE]
   --list-chunk value               列举数据块的大小（每个 ListObject S3 请求的响应列表）。 (默认值：1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行 URL 编码：true/false/unset (默认值："unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的 ListObjects 版本：1、2 或 0 表示自动选择。 (默认值：0) [$LIST_VERSION]
   --max-upload-parts value         分块上传中的最大部分数。 (默认值：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (默认值："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。 (默认值：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行 gzip 压缩，请设置此项。 (默认值："unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不在上传对象之前对其进行 HEAD 操作以检查完整性。 (默认值：false) [$NO_HEAD]
   --no-head-object                 如果设置，则在获取对象之前不执行 HEAD 操作。 (默认值：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值：false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS 会话 Token。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用 SSE-C，则存储此对象时所使用的服务器端加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         若要使用 SSE-C，则可以提供用于加密/解密数据的密钥。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用 SSE-C，则必须以 base64 格式提供用于加密/解密数据的密钥。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用 SSE-C，则可以提供密钥的 MD5 校验和（可选）。 [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       分块上传的并发数。 (默认值：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的上传截止点。 (默认值："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在分块上传中使用 ETag 进行验证 (默认值："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否在单个部分上传时使用预签名请求或 PutObject。 (默认值：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为 true，则使用 v2 认证。 (默认值：false) [$V2_AUTH]
   --version-at value               根据指定时间显示文件版本。 (默认值："off") [$VERSION_AT]
   --versions                       在目录列表中包括旧版本。 (默认值：false) [$VERSIONS]

```
{% endcode %}