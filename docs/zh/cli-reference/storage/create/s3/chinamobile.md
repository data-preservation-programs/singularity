# 中国移动Ecloud Elastic Object Storage（EOS）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 chinamobile - 中国移动Ecloud弹性对象存储（EOS）

用法:
   singularity storage create s3 chinamobile [command options] [arguments...]

描述:
   --env-auth
      从运行环境中获取AWS凭据（环境变量或无环境变量时，从EC2/ECS元数据获取）。

      仅当access_key_id和secret_access_key为空时适用。

      示例：
         | false | 在下一步中输入AWS凭据。
         | true  | 从环境（环境变量或IAM）获取AWS凭据。

   --access-key-id
      AWS Access Key ID。

      为空表示匿名访问或运行时凭据.

   --secret-access-key
      AWS Secret Access Key（密码）。

      为空表示匿名访问或运行时凭据。

   --endpoint
      中国移动Ecloud弹性对象存储（EOS）API的终端节点。

      示例：
         | eos-wuxi-1.cmecloud.cn      | 默认终端节点 - 如果您不确定，可以选择此选项。
         |                             | 华东区（苏州）
         | eos-jinan-1.cmecloud.cn     | 华东区（济南）
         | eos-ningbo-1.cmecloud.cn    | 华东区（杭州）
         | eos-shanghai-1.cmecloud.cn  | 华东区（上海-1）
         | eos-zhengzhou-1.cmecloud.cn | 华中区（郑州）
         | eos-hunan-1.cmecloud.cn     | 华中区（长沙-1）
         | eos-zhuzhou-1.cmecloud.cn   | 华中区（长沙-2）
         | eos-guangzhou-1.cmecloud.cn | 华南区（广州-2）
         | eos-dongguan-1.cmecloud.cn  | 华南区（广州-3）
         | eos-beijing-1.cmecloud.cn   | 华北区（北京-1）
         | eos-beijing-2.cmecloud.cn   | 华北区（北京-2）
         | eos-beijing-4.cmecloud.cn   | 华北区（北京-3）
         | eos-huhehaote-1.cmecloud.cn | 华北区（呼和浩特）
         | eos-chengdu-1.cmecloud.cn   | 西南区（成都）
         | eos-chongqing-1.cmecloud.cn | 西南区（重庆）
         | eos-guiyang-1.cmecloud.cn   | 西南区（贵阳）
         | eos-xian-1.cmecloud.cn      | 西北区（西安）
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
      位置约束，必须与终端节点匹配。

      仅在创建存储桶时使用。

      示例：
         | wuxi1      | 华东区（苏州）
         | jinan1     | 华东区（济南）
         | ningbo1    | 华东区（杭州）
         | shanghai1  | 华东区（上海-1）
         | zhengzhou1 | 华中区（郑州）
         | hunan1     | 华中区（长沙-1）
         | zhuzhou1   | 华中区（长沙-2）
         | guangzhou1 | 华南区（广州-2）
         | dongguan1  | 华南区（广州-3）
         | beijing1   | 华北区（北京-1）
         | beijing2   | 华北区（北京-2）
         | beijing4   | 华北区（北京-3）
         | huhehaote1 | 华北区（呼和浩特）
         | chengdu1   | 西南区（成都）
         | chongqing1 | 西南区（重庆）
         | guiyang1   | 西南区（贵阳）
         | xian1      | 西北区（西安）
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
      创建存储桶、存储或复制对象时使用的预定义ACL。

      此ACL用于创建对象，如果未设置bucket_acl，则用于创建存储桶。

      有关详细信息，请参阅https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，在S3复制对象时，S3不会复制源对象的ACL，而是写入一个新的ACL。

      如果acl是空字符串，则不会添加X-Amz-Acl:标头，并使用默认（private）。

   --bucket-acl
      创建存储桶时使用的预定义ACL。

      有关详细信息，请参阅https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl。

      请注意，仅在创建存储桶时使用此ACL。如果未设置，则使用"acl"。

      如果"acl"和"bucket_acl"都是空字符串，则不会添加X-Amz-Acl:标头，并使用默认（private）。

      示例：
         | private            | 所有者获得FULL_CONTROL。
         |                    | 其他人没有访问权限（默认）。
         | public-read        | 所有者获得FULL_CONTROL。
         |                    | AllUsers组获得读取权限。
         | public-read-write  | 所有者获得FULL_CONTROL。
         |                    | AllUsers组获得读写权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者获得FULL_CONTROL。
         |                    | AuthenticatedUsers组获得读取权限。

   --server-side-encryption
      存储对象在S3中使用的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      如果使用SSE-C，在将此对象存储在S3时所使用的服务器端加密算法。

      示例：
         | <unset> | None
         | AES256  | AES256

   --sse-customer-key
      如果使用SSE-C，可以提供用于加密/解密数据的加密密钥。

      或者可以提供--sse-customer-key-base64。

      示例：
         | <unset> | None

   --sse-customer-key-base64
      如果使用SSE-C，必须以base64格式提供加密/解密数据的加密密钥。

      或者可以提供--sse-customer-key。

      示例：
         | <unset> | None

   --sse-customer-key-md5
      如果使用SSE-C，可以提供加密密钥的MD5校验和（可选）。

      如果不填写，则会根据提供的sse_customer_key自动计算。

      示例：
         | <unset> | None

   --storage-class
      在ChinaMobile中存储新对象时要使用的存储类。

      示例：
         | <unset>     | 默认
         | STANDARD    | 标准存储类
         | GLACIER     | 存档存储类
         | STANDARD_IA | 低频访问存储类

   --upload-cutoff
      转换为分块上传的截止点。

      任何大于此大小的文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时使用的分块大小。

      当上传大于upload_cutoff的文件时，或者上传大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或Google相册或Google文档上传的文件）时，将使用此分块大小进行分块上传。

      请注意，每次传输都会在内存中缓冲chunk_size大小的分块，并行上传。

      如果您正在通过高速连接传输大文件，并且内存足够，增加此值将加快传输速度。

      Rclone会自动增加分块大小以确保在10,000个分块的限制之下。
      
      大小未知的文件以配置的chunk_size上传。由于默认的chunk_size为5 MiB，最多可以有10,000个分块，因此默认情况下您可以上传的文件的最大大小为48 GiB。如果要上传更大的文件，则需要增加chunk_size。
      
      增加分块大小会降低使用"-P"标志显示的进度统计的准确性。Rclone将在使用"-P"标志时将chunk视为已发送，但实际上可能仍在上传。较大的chunk_size意味着较大的AWS SDK缓冲区和进度报告与实际情况偏离的可能性更大。

   --max-upload-parts
      多部分上传中的最大部分数。

      此选项定义在执行多部分上传时要使用的最大多部分块数。

      如果某个服务不支持10,000个多部分块的AWS S3规范，则此选项可能非常有用。

      Rclone会自动增加分块大小以确保在这个多部分块限制之下。

   --copy-cutoff
      转换为多部分复制的截止点。

      任何大于此大小的需要在服务器端复制的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不将MD5校验和存储到对象元数据中。

      通常情况下，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但可能会导致大文件上传开始时出现长时间的延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。
      
      如果该变量为空，则rclone将查找"AWS_SHARED_CREDENTIALS_FILE"环境变量。如果环境值为空，则它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共享凭据文件中要使用的配置文件。
      
      如果env_auth = true，则rclone可以使用共享凭据文件。此变量控制在该文件中使用哪个配置文件。
      
      如果为空，则默认为环境变量"AWS_PROFILE"或"default"（如果未设置该环境变量）。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      此选项是同时上传相同文件的块数。
      
      如果您通过高速链接上传几个大文件，并且这些上传未充分利用带宽，那么增加该值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，如果为false，则rclone将使用虚拟路径样式访问。有关更多信息，请参阅[AWS S3文档](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证；

      如果为false（默认值），则rclone将使用v4身份验证。如果设置了它，rclone将使用v2身份验证。

      仅在v4签名无法正常工作时使用，例如在Jewel/v10 CEPH之前。

   --list-chunk
      列表的块大小（每个ListObject S3请求的响应列表大小）。
      
      此选项也被称为AWS S3规范中的 "MaxKeys"、"max-items" 或 "page-size"。
      大多数服务即使请求超过1000个响应列表，也会将其截断为1000个对象。
      在AWS S3中，这是一个全局最大值，不可更改，参见[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以使用"rgw list buckets max chunk"选项进行增加。

   --list-version
      使用的ListObjects版本：1、2或0（自动判断）。
      
      当S3最初发布时，它仅提供用于枚举存储桶中对象的ListObjects调用。
      
      然而，在2016年5月，引入了ListObjectsV2调用。这是一个更高性能的调用，应尽量使用它。
      
      如果设置为默认值0，rclone将根据设置的提供商猜测要调用哪个列表对象方法。如果猜测错误，则可以在此手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      一些提供商支持URL编码列表，在可用时，在文件名中使用控制字符时，这更可靠。如果设置为unset（默认值），则rclone将根据提供商设置选择要应用的设置，但是您可以在此处覆盖rclone的选择。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      这在尽量减少rclone进行的事务数量时很有用，如果您知道存储桶已经存在。

      如果使用的用户没有创建存储桶的权限，可能也需要启用此选项。在v1.52.0之前，这将由于错误而沉默通过。

   --no-head
      如果设置，则不检查已上传的对象的头部以检查完整性。
      
      当尽量减少rclone执行的事务数量时，这可能很有用。

      设置它意味着如果rclone在使用PUT上传对象后接收到200 OK消息，它将假设成功上传。

      特别是它将假设：

      - 元数据，包括modtime、存储类别和内容类型和上传一致。
      - 大小与上传的一致。
      
      它从单个部分PUT的响应中读取以下项：

      - MD5校验和
      - 上传日期
      
      对于多部分上传，不会读取这些项。

      如果上传的源对象长度未知，则rclone **将**执行HEAD请求。

      设置此标志将增加检测到的上传失败的可能性，特别是大小不正确，因此不建议在正常情况下使用。实际上，即使设置此标志，检测上传失败的机会也很小。

   --no-head-object
      如果设置，则在GET对象时不执行HEAD。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅概述中的[encoding section](/overview/#encoding)。

   --memory-pool-flush-time
      内部内存缓冲池刷新的时间间隔。
      
      需要额外缓冲区（例如分块）的上传将使用内存池进行分配。
      此选项控制将未使用的缓冲区从池中移除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的HTTP2使用。
      
      当前s3（特别是minio）后端与HTTP/2存在一个未解决的问题。默认情况下启用s3后端的HTTP/2，但可以在此禁用该功能。解决此问题后，将删除此标志。

      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631。

   --download-url
      自定义下载的端点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载数据的数据传出量更便宜。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证

      此设置应为true、false或保留默认值（unset），以使用提供商的默认值。

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      版本为1.59的rclone使用预先签名的请求上传单个部分的对象，设置此标志为true将重新启用该功能。除非在特殊情况下或进行测试，否则不应该这样做。

   --versions
      在目录列表中包含旧版本。

   --version-at
      根据指定的时间显示文件版本。
      
      参数应为日期，"2006-01-02"，日期时间 "2006-01-02 15:04:05" 或距现在的持续时间，例如 "100d" 或 "1h"。
      
      请注意，使用此选项时不允许进行文件写操作，因此不能上传文件或删除文件。
      
      有关有效格式，请参阅[时间选项文档](/docs/#time-option)。

   --decompress
      如果设置，将解压缩gzip编码的对象。
      
      可以在上传对象到S3时设置 "Content-Encoding: gzip"。通常，rclone会以压缩的对象文件进行下载。
      
      如果设置了此标志，则rclone会在接收到设置为 "Content-Encoding: gzip" 的文件时解压缩这些文件。这意味着rclone无法检查大小和哈希值，但文件内容将被解压缩。

   --might-gzip
      如果后端可能对对象进行gzip压缩，则设置此标志。

      通常情况下，提供商在下载时不会修改对象。如果对象在上传时没有设置 `Content-Encoding: gzip`，则在下载时也不会设置。

      但是，有些提供商甚至可以对未使用 `Content-Encoding: gzip` 进行上传的对象进行gzip压缩（例如Cloudflare）。

      如果设置了此标志，并且rclone下载了带有设置了 `Content-Encoding: gzip` 并且使用了chunked传输编码的对象，则rclone会动态解压缩对象。
      
      如果此设置为unset（默认值），则rclone将根据提供商设置选择要应用的设置，但是您可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value           AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                     创建存储桶、存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value                中国移动Ecloud弹性对象存储（EOS）API的终端节点。[$ENDPOINT]
   --env-auth                      从运行环境中获取AWS凭据（环境变量或无环境变量时，从EC2/ECS元数据获取）。 (默认值: false) [$ENV_AUTH]
   --help, -h                      显示帮助
   --location-constraint value     位置约束，必须与终端节点匹配。[$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS Secret Access Key（密码）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  存储对象在S3中使用的服务器端加密算法。[$SERVER_SIDE_ENCRYPTION]
   --storage-class value           在ChinaMobile中存储新对象时要使用的存储类。[$STORAGE_CLASS]

   高级

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               上传时使用的分块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              转换为多部分复制的截止点。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，将解压缩gzip编码的对象。 (默认值: false) [$DECOMPRESS]
   --disable-checksum               不将MD5校验和存储到对象元数据中。 (默认值: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的HTTP2使用。 (默认值: false) [$DISABLE_HTTP2]
   --download-url value             自定义下载的端点。[$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。 (默认值: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表的块大小（每个ListObject S3请求的响应列表大小）。 (默认值: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (默认值: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用的ListObjects版本：1、2或0（自动判断）。 (默认值: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (默认值: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的时间间隔。 (默认值: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。 (默认值: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能对对象进行gzip压缩，则设置此标志。 (默认值: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不检查已上传的对象的头部以检查完整性。 (默认值: false) [$NO_HEAD]
   --no-head-object                 如果设置，则在GET对象时不执行HEAD。 (默认值: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (默认值: false) [$NO_SYSTEM_METADATA]
   --profile value                  共享凭据文件中要使用的配置文件。 [$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   如果使用SSE-C，在将此对象存储在S3时所使用的服务器端加密算法。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         如果使用SSE-C，可以提供用于加密/解密数据的加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  如果使用SSE-C，必须以base64格式提供加密/解密数据的加密密钥。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     如果使用SSE-C，可以提供加密密钥的MD5校验和（可选）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       多部分上传的并发数。 (默认值: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            转换为分块上传的截止点。 (默认值: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (默认值: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (默认值: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (默认值: false) [$V2_AUTH]
   --version-at value               根据指定的时间显示文件版本。 (默认值: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (默认值: false) [$VERSIONS]

   通用

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}