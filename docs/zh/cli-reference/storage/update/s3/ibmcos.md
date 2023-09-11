# IBM COS S3

{% code fullWidth="true" %}
```
名称：
   singularity storage update s3 ibmcos - IBM COS S3

用法：
   singularity storage update s3 ibmcos [命令选项] <名称|ID>

描述：
   --env-auth
      从运行时获取AWS凭证（如果access_key_id和secret_access_key为空，则从环境变量或EC2/ECS元数据获取）。
      
      仅在access_key_id和secret_access_key为空时生效。

      示例：
         | false | 在下一步中输入AWS凭证。
         | true  | 从运行时环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS访问密钥ID。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS密钥访问密钥（密码）。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --region
      要连接的区域。
      
      如果使用S3克隆且没有区域，请留空。

      示例：
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空区域。
         | other-v2-signature | 仅在v4签名不起作用时使用此选项。
         |                    | 例如，旧版本的Jewel/v10 CEPH。

   --endpoint
      IBM COS S3 API的终端节点。
      
      如果使用IBM COS本地部署，请指定。

      示例：
         | s3.us.cloud-object-storage.appdomain.cloud               | 美国跨区域终端节点
         | s3.dal.us.cloud-object-storage.appdomain.cloud           | 美国跨区域达拉斯终端节点
         | s3.wdc.us.cloud-object-storage.appdomain.cloud           | 美国跨区域华盛顿特区终端节点
         | s3.sjc.us.cloud-object-storage.appdomain.cloud           | 美国跨区域圣何塞终端节点
         | s3.private.us.cloud-object-storage.appdomain.cloud       | 美国跨区域私有终端节点
         | s3.private.dal.us.cloud-object-storage.appdomain.cloud   | 美国跨区域达拉斯私有终端节点
         | s3.private.wdc.us.cloud-object-storage.appdomain.cloud   | 美国跨区域华盛顿特区私有终端节点
         | s3.private.sjc.us.cloud-object-storage.appdomain.cloud   | 美国跨区域圣何塞私有终端节点
         | s3.us-east.cloud-object-storage.appdomain.cloud          | 美国东部区域终端节点
         | s3.private.us-east.cloud-object-storage.appdomain.cloud  | 美国东部区域私有终端节点
         | s3.us-south.cloud-object-storage.appdomain.cloud         | 美国南部区域终端节点
         | s3.private.us-south.cloud-object-storage.appdomain.cloud | 美国南部区域私有终端节点
         | s3.eu.cloud-object-storage.appdomain.cloud               | 欧洲跨区域终端节点
         | s3.fra.eu.cloud-object-storage.appdomain.cloud           | 欧洲跨区域法兰克福终端节点
         | s3.mil.eu.cloud-object-storage.appdomain.cloud           | 欧洲跨区域米兰终端节点
         | s3.ams.eu.cloud-object-storage.appdomain.cloud           | 欧洲跨区域阿姆斯特丹终端节点
         | s3.private.eu.cloud-object-storage.appdomain.cloud       | 欧洲跨区域私有终端节点
         | s3.private.fra.eu.cloud-object-storage.appdomain.cloud   | 欧洲跨区域法兰克福私有终端节点
         | s3.private.mil.eu.cloud-object-storage.appdomain.cloud   | 欧洲跨区域米兰私有终端节点
         | s3.private.ams.eu.cloud-object-storage.appdomain.cloud   | 欧洲跨区域阿姆斯特丹私有终端节点
         | s3.eu-gb.cloud-object-storage.appdomain.cloud            | 英国终端节点
         | s3.private.eu-gb.cloud-object-storage.appdomain.cloud    | 英国私有终端节点
         | s3.eu-de.cloud-object-storage.appdomain.cloud            | 欧洲DE区域终端节点
         | s3.private.eu-de.cloud-object-storage.appdomain.cloud    | 欧洲DE区域私有终端节点
         | s3.ap.cloud-object-storage.appdomain.cloud               | 亚太跨区域终端节点
         | s3.tok.ap.cloud-object-storage.appdomain.cloud           | 亚太跨区域东京终端节点
         | s3.hkg.ap.cloud-object-storage.appdomain.cloud           | 亚太跨区域香港终端节点
         | s3.seo.ap.cloud-object-storage.appdomain.cloud           | 亚太跨区域首尔终端节点
         | s3.private.ap.cloud-object-storage.appdomain.cloud       | 亚太跨区域私有终端节点
         | s3.private.tok.ap.cloud-object-storage.appdomain.cloud   | 亚太跨区域东京私有终端节点
         | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud   | 亚太跨区域香港私有终端节点
         | s3.private.seo.ap.cloud-object-storage.appdomain.cloud   | 亚太跨区域首尔私有终端节点
         | s3.jp-tok.cloud-object-storage.appdomain.cloud           | 亚太日本区域终端节点
         | s3.private.jp-tok.cloud-object-storage.appdomain.cloud   | 亚太日本区域私有终端节点
         | s3.au-syd.cloud-object-storage.appdomain.cloud           | 亚太澳大利亚区域终端节点
         | s3.private.au-syd.cloud-object-storage.appdomain.cloud   | 亚太澳大利亚区域私有终端节点
         | s3.ams03.cloud-object-storage.appdomain.cloud            | 阿姆斯特丹单区域终端节点
         | s3.private.ams03.cloud-object-storage.appdomain.cloud    | 阿姆斯特丹单区域私有终端节点
         | s3.che01.cloud-object-storage.appdomain.cloud            | 金奈单区域终端节点
         | s3.private.che01.cloud-object-storage.appdomain.cloud    | 金奈单区域私有终端节点
         | s3.mel01.cloud-object-storage.appdomain.cloud            | 墨尔本单区域终端节点
         | s3.private.mel01.cloud-object-storage.appdomain.cloud    | 墨尔本单区域私有终端节点
         | s3.osl01.cloud-object-storage.appdomain.cloud            | 奥斯陆单区域终端节点
         | s3.private.osl01.cloud-object-storage.appdomain.cloud    | 奥斯陆单区域私有终端节点
         | s3.tor01.cloud-object-storage.appdomain.cloud            | 多伦多单区域终端节点
         | s3.private.tor01.cloud-object-storage.appdomain.cloud    | 多伦多单区域私有终端节点
         | s3.seo01.cloud-object-storage.appdomain.cloud            | 首尔单区域终端节点
         | s3.private.seo01.cloud-object-storage.appdomain.cloud    | 首尔单区域私有终端节点
         | s3.mon01.cloud-object-storage.appdomain.cloud            | 蒙特利尔单区域终端节点
         | s3.private.mon01.cloud-object-storage.appdomain.cloud    | 蒙特利尔单区域私有终端节点
         | s3.mex01.cloud-object-storage.appdomain.cloud            | 墨西哥单区域终端节点
         | s3.private.mex01.cloud-object-storage.appdomain.cloud    | 墨西哥单区域私有终端节点
         | s3.sjc04.cloud-object-storage.appdomain.cloud            | 圣何塞单区域终端节点
         | s3.private.sjc04.cloud-object-storage.appdomain.cloud    | 圣何塞单区域私有终端节点
         | s3.mil01.cloud-object-storage.appdomain.cloud            | 米兰单区域终端节点
         | s3.private.mil01.cloud-object-storage.appdomain.cloud    | 米兰单区域私有终端节点
         | s3.hkg02.cloud-object-storage.appdomain.cloud            | 香港单区域终端节点
         | s3.private.hkg02.cloud-object-storage.appdomain.cloud    | 香港单区域私有终端节点
         | s3.par01.cloud-object-storage.appdomain.cloud            | 巴黎单区域终端节点
         | s3.private.par01.cloud-object-storage.appdomain.cloud    | 巴黎单区域私有终端节点
         | s3.sng01.cloud-object-storage.appdomain.cloud            | 新加坡单区域终端节点
         | s3.private.sng01.cloud-object-storage.appdomain.cloud    | 新加坡单区域私有终端节点

   --location-constraint
      区域约束 - 必须与使用IBM Cloud Public时的终端节点匹配。
      
      对于本地部署的COS，请不要从此列表中选择。

      示例：
         | us-standard       | 美国跨区域标准桶
         | us-vault          | 美国跨区域保险库桶
         | us-cold           | 美国跨区域冷存桶
         | us-flex           | 美国跨区域弹性桶
         | us-east-standard  | 美国东部标准桶
         | us-east-vault     | 美国东部保险库桶
         | us-east-cold      | 美国东部冷存桶
         | us-east-flex      | 美国东部弹性桶
         | us-south-standard | 美国南部标准桶
         | us-south-vault    | 美国南部保险库桶
         | us-south-cold     | 美国南部冷存桶
         | us-south-flex     | 美国南部弹性桶
         | eu-standard       | 欧洲跨区域标准桶
         | eu-vault          | 欧洲跨区域保险库桶
         | eu-cold           | 欧洲跨区域冷存桶
         | eu-flex           | 欧洲跨区域弹性桶
         | eu-gb-standard    | 英国标准桶
         | eu-gb-vault       | 英国保险库桶
         | eu-gb-cold        | 英国冷存桶
         | eu-gb-flex        | 英国弹性桶
         | ap-standard       | 亚太标准桶
         | ap-vault          | 亚太保险库桶
         | ap-cold           | 亚太冷存桶
         | ap-flex           | 亚太弹性桶
         | mel01-standard    | 墨尔本标准桶
         | mel01-vault       | 墨尔本保险库桶
         | mel01-cold        | 墨尔本冷存桶
         | mel01-flex        | 墨尔本弹性桶
         | tor01-standard    | 多伦多标准桶
         | tor01-vault       | 多伦多保险库桶
         | tor01-cold        | 多伦多冷存桶
         | tor01-flex        | 多伦多弹性桶

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。
      
      此ACL适用于创建对象，并且如果未设置bucket_acl，则还适用于创建存储桶。
      
      更多信息请参见[Amazon S3开发人员指南](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      注意，在进行服务端复制对象时，S3不会复制源对象的ACL，而是写入一个新的ACL。
      
      如果acl是空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

      示例：
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 没有其他人具有访问权限（默认）。
         |                    | 此acl适用于IBM Cloud（基础架构）、IBM Cloud（存储）和本地部署COS。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取权限。
         |                    | 此acl适用于IBM Cloud（基础架构）、IBM Cloud（存储）、本地部署IBM COS。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有读取和写入权限。
         |                    | 此acl适用于IBM Cloud（基础架构）和本地部署的IBM COS。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有读取权限。
         |                    | 不能在桶上使用。
         |                    | 此acl适用于IBM Cloud（基础架构）和本地部署的IBM COS。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      更多信息请参见[Amazon S3开发人员指南](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)。
      
      注意，仅在创建存储桶时应用此ACL。如果未设置它，则将使用“acl”。
      
      如果“acl”和“bucket_acl”都是空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

   --upload-cutoff
      切换到分块上传的截止点。
      
      大于此大小的文件将按照chunk_size进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或带有未知大小的文件（例如使用“rclone rcat”或使用“rclone mount”或Google
      照片或Google文档上传的文件）时，将使用该块大小进行多部分上传。
      
      请注意，“--s3-upload-concurrency”个这种大小的块会在每次传输中在内存中进行缓冲。
      
      如果您正在通过高速连接传输大型文件并且拥有足够的内存，则增加此大小将加快传输速度。
      
      Rclone将自动增加块大小，以便在上传已知大小的大文件时保持在10000个块的限制之下。
      
      未知大小的文件使用配置的chunk_size进行上传。由于默认块大小为5 MiB，并且最多可以有
      10000个块，这意味着默认情况下，您可以使用48 GiB的文件大小进行流式上传。如果要流式上传
      更大的文件，则需要增加chunk_size。
      
      增大块大小会减小使用“-P”标志显示的进度统计信息的准确性。当rclone对一个块进行缓冲时，
      它会认为块已被发送，而实际上可能还在上传中。较大的块大小意味着较大的AWS SDK缓冲区和进度
      报告与真实情况偏离更大。

   --max-upload-parts
      多部分上传中的最大部分数量。
      
      此选项定义了在执行多部分上传时要使用的最大分块数。
      
      如果某个服务不支持AWS S3的10000个分块的规范，则此选项可能很有用。
      
      当上传已知大小的大文件时，rclone将自动增加块大小以确保不超过此块数的限制。

   --copy-cutoff
      切换到分块复制的截止点。
      
      需要进行服务器端复制且大于此大小的文件将按照该大小的块复制。

   --disable-checksum
      不要在对象元数据中存储MD5校验和。
      
      通常情况下，在上传文件之前，rclone会计算输入的MD5校验和，以便将其添加到对象的元数据中。这非常适用于数据完整性
      检查，但对于大型文件启动上传可能会导致长时间延迟。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果env_auth为true，则rclone可以使用共享凭证文件。
      
      如果此变量为空，则rclone将查找“AWS_SHARED_CREDENTIALS_FILE”环境变量。
      如果环境值为空，则会默认使用当前用户的主目录。
      
          Linux/OSX：“$HOME/.aws/credentials”
          Windows：“%USERPROFILE%\.aws\credentials”

   --profile
      共享凭证文件中要使用的配置文件。
      
      如果env_auth为true，则rclone可以使用共享凭证文件。该变量控制在该文件中使用的配置文件。
      
      如果为空，则默认使用环境变量“AWS_PROFILE”或默认值“default”。

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传相同文件的分块数。
      
      如果使用高速链接上传少量大文件且这些上传未充分利用带宽，则增加此值可能有助于加快传输速度。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问；如果为false，则rclone将使用虚拟路径样式访问。
      有关更多信息，请参见[AWS S3文档]https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro）。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）要求将此设置为
      false - rclone将根据提供商设置自动执行此操作。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置了此值，则rclone将使用v2身份验证。
      
      仅在v4签名不起作用时才使用此项，例如旧版本的Jewel/v10 CEPH。

   --list-chunk
      列举块的大小（一个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的“MaxKeys”、“max-items”或“page-size”。
      大多数服务即使请求超过1000个对象，也会截断响应列表为1000个对象。
      在AWS S3中，这是一个全局最大值，无法更改，参见
      [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      在Ceph中，可以通过“rgw list buckets max chunk”选项增加此值。

   --list-version
      要使用的列表对象的版本：1、2或0为自动。
      
      当S3最初推出时，它仅提供了用于枚举存储桶中对象的ListObjects调用。
      
      然而，在2016年5月引入了ListObjectsV2调用。这是更高性能的方式，如果可能，应该使用它。
      
      如果设置为默认值0，则rclone将根据设置的提供者猜测要调用哪个列表对象方法。如果它猜错了，
      则可能在此处手动设置。

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持对列表进行URL编码，如果可用，则在文件名中使用控制字符时，这样处理更可靠。
      如果设置为unset（默认值），则rclone将根据提供者设置选择要应用的内容，但可以在此处覆盖rclone的选择。

   --no-check-bucket
      设置后，不尝试检查存储桶的存在或创建。
      
      这在试图最小化rclone的事务数时非常有用，如果您知道存储桶已存在，则可以使用它。
      
      如果使用的用户没有存储桶创建权限，则可能需要此项。在v1.52.0之前，此项将以静默方式传递，因为存在错误。

   --no-head
      设置后，不会对已上传的对象进行HEAD操作以检查完整性。
      
      这在试图最小化rclone的事务数时非常有用。
      
      设置它意味着如果rclone在使用PUT上传对象后收到200 OK消息，它将假定已正确上传该对象。
      
      特别是它将假定：
      
      - 元数据，包括修改时间、存储类别和内容类型与上传的一样
      - 大小与上传的一样
      
      对于单部分PUT的响应，它从响应中读取以下项目：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，它不会读取这些项目。
      
      如果上传源对象的长度未知，则rclone将**do HEAD请求**。
      
      设置此标志会增加未检测到的上传失败的几率，特别是错误的大小，因此不建议在正常操作中使用它。实际上，
      即使启用此标志，检测不到的上传失败的几率也非常小。
      

   --no-head-object
      如果设置，则在执行GET操作获取对象之前不会执行HEAD操作。

   --encoding
      后端使用的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --memory-pool-flush-time
      刷新内部内存缓冲池的频率。
      
      需要额外缓冲区（例如由多部分上传）的上传将使用内存池进行分配。
      此选项控制多长时间未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（具体来说是minio）后端存在一个未解决的问题，即与HTTP/2的问题。 S3后端默认启用HTTP/2，但可以在此处禁用。
      解决此问题后，将删除此标志。
      
      参见：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      下载的自定义终端节点。
      此通常设置为CloudFront CDN URL，因为通过CloudFront网络下载的数据经过AWS S3提供更便宜的外出流量。

   --use-multipart-etag
      是否在对多部分上传进行验证时使用ETag
      
      值可以是true、false或保持未设置以使用提供者的默认值。

   --use-presigned-request
      是否使用预先签名的请求还是PutObject进行单个部分上传
      
      如果为false，则rclone将使用AWS SDK的PutObject上传对象。
      
      rclone的版本<1.59使用预签名的请求上传单个部分对象，将此标志设置为true将重新启用该功能。除了特殊情况或测试之外，不应该需要这样做。

   --versions
      在目录列表中包括旧版本。

   --version-at
      按指定的时间显示文件版本。
      
      参数应该是一个日期，“2006-01-02”，日期时间“2006-01-02 15:04:05”或距离那个时间的持续时间，例如“100d”或“1h”。
      
      请注意，使用此功能时，不允许进行文件写操作，因此无法上传文件或删除它们。
      
      有关有效格式，请参见[时间选项文档](/docs/#time-option)。
      

   --decompress
      如果设置，这将解压缩gzip编码的对象。
      
      可以使用“Content-Encoding: gzip”设置向S3上传对象。通常，rclone以压缩的对象形式下载这些文件。
      
      如果设置此标志，rclone将在接收到以“Content-Encoding: gzip”形式的文件时进行解压缩。这意味着rclone不能检查大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能gzip对象，请设置此标志。
      
      通常，提供商不会在下载时更改对象。如果未使用“Content-Encoding: gzip”上传对象，则在下载时也不会设置它。
      
      但是，某些提供商甚至可能在未使用“Content-Encoding: gzip”上传对象的情况下对其进行gzip压缩（例如Cloudflare）。
      
      这样做的症状可能是接收到以下类似的错误：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置此标志，并且rclone下载带有Content-Encoding: gzip设置和分块传输编码的对象，则rclone将随时解压缩该对象。
      
      如果将此设置为unset（默认值），则rclone将根据提供者设置选择要应用的内容，但可以在此处覆盖rclone的选择。

   --no-system-metadata
      禁止设置和读取系统元数据


选项：
   --access-key-id value        AWS访问密钥ID。[$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预定义ACL。[$ACL]
   --endpoint value             IBM COS S3 API的终端节点。[$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（如果access_key_id和secret_access_key为空，则从环境变量或EC2/ECS元数据获取）。（默认值：false）[$ENV_AUTH]
   --help, -h                   显示帮助信息
   --location-constraint value  区域约束 - 必须与使用IBM Cloud Public时的终端节点匹配。[$LOCATION_CONSTRAINT]
   --region value               要连接的区域。[$REGION]
   --secret-access-key value    AWS密钥访问密钥（密码）。[$SECRET_ACCESS_KEY]

   高级选项

   --bucket-acl value               创建存储桶时使用的预定义ACL。[$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。（默认值：“5Mi”）[$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截止点。（默认值：“4.656Gi”）[$COPY_CUTOFF]
   --decompress                     如果设置此项，则将解压缩gzip编码的对象。（默认值：false）[$DECOMPRESS]
   --disable-checksum               不要在对象元数据中存储MD5校验和。（默认值：false）[$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。（默认值：false）[$DISABLE_HTTP2]
   --download-url value             下载的自定义终端节点。[$DOWNLOAD_URL]
   --encoding value                 后端使用的编码。（默认值：“Slash,InvalidUtf8,Dot”）[$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟托管样式访问。（默认值：true）[$FORCE_PATH_STYLE]
   --list-chunk value               列举块的大小（一个ListObject S3请求的响应列表）。（默认值：1000）[$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset（默认值：“unset”）[$LIST_URL_ENCODE]
   --list-version value             要使用的列表对象的版本：1、2或0为自动。（默认值：0）[$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数量。（默认值：10000）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   刷新内部内存缓冲池的频率。（默认值：“1m0s”）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区。（默认值：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能gzip对象，请设置此标志。（默认值：“unset”）[$MIGHT_GZIP]
   --no-check-bucket                设置后，不尝试检查存储桶的存在或创建。 （默认值：false）[$NO_CHECK_BUCKET]
   --no-head                        设置后，不会对已上传的对象进行HEAD操作以检查完整性。 （默认值：false）[$NO_HEAD]
   --no-head-object                 如果设置，则在执行GET操作获取对象之前不会执行HEAD操作。 （默认值：false）[$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据（默认值：false）[$NO_SYSTEM_METADATA]
   --profile value                  共享凭证文件中要使用的配置文件。[$PROFILE]
   --session-token value            AWS会话令牌。[$SESSION_TOKEN]
   --shared-credentials-file value  共享凭证文件的路径。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。（默认值：4）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点。（默认值：“200Mi”）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在对多部分上传进行验证时使用ETag（默认值：“unset”）[$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预先签名的请求还是PutObject进行单个部分上传（默认值：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。（默认值：false）[$V2_AUTH]
   --version-at value               按指定的时间显示文件版本。（默认值：“off”）[$VERSION_AT]
   --versions                       在目录列表中包括旧版本。（默认值：false）[$VERSIONS]

```
{% endcode %}