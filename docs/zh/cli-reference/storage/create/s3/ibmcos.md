# IBM COS S3

{% code fullWidth="true" %}
```
NAME:
   创建s3 ibmcos - IBM COS S3

用法:
   singularity storage create s3 ibmcos [命令选项] [参数...]

描述:
   --env-auth
      从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。
      
      仅在access_key_id和secret_access_key为空时适用。

      示例:
         | false | 在下一步中输入AWS凭证。
         | true  | 从环境中获取AWS凭证（环境变量或IAM）。

   --access-key-id
      AWS Access Key ID。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --secret-access-key
      AWS Secret Access Key（密码）。
      
      如果要匿名访问或使用运行时凭证，请留空。

   --region
      连接的区域。
      
      如果您使用的是S3克隆，并且没有所在区域，请留空。

      示例:
         | <unset>            | 如果不确定，请使用此选项。
         |                    | 将使用v4签名和空的区域。
         | other-v2-signature | 只有在v4签名不起作用时才使用此选项。
         |                    | 例如，Jewel/v10 CEPH之前的版本。

   --endpoint
      IBM COS S3 API的终端点。
      
      如果使用IBM COS本地部署，请指定。

      示例:
         | s3.us.cloud-object-storage.appdomain.cloud                       | 跨区域美国终端节点
         | s3.dal.us.cloud-object-storage.appdomain.cloud                   | 跨区域达拉斯终端节点
         | s3.wdc.us.cloud-object-storage.appdomain.cloud                   | 跨区域华盛顿特区终端节点
         | s3.sjc.us.cloud-object-storage.appdomain.cloud                   | 跨区域圣何塞终端节点
         | s3.private.us.cloud-object-storage.appdomain.cloud               | 跨区域美国私有终端节点
         | s3.private.dal.us.cloud-object-storage.appdomain.cloud           | 跨区域达拉斯私有终端节点
         | s3.private.wdc.us.cloud-object-storage.appdomain.cloud           | 跨区域华盛顿特区私有终端节点
         | s3.private.sjc.us.cloud-object-storage.appdomain.cloud           | 跨区域圣何塞私有终端节点
         | s3.us-east.cloud-object-storage.appdomain.cloud                  | 美国东部区域终端节点
         | s3.private.us-east.cloud-object-storage.appdomain.cloud          | 美国东部区域私有终端节点
         | s3.us-south.cloud-object-storage.appdomain.cloud                 | 美国南部区域终端节点
         | s3.private.us-south.cloud-object-storage.appdomain.cloud         | 美国南部区域私有终端节点
         | s3.eu.cloud-object-storage.appdomain.cloud                       | 跨区域欧洲终端节点
         | s3.fra.eu.cloud-object-storage.appdomain.cloud                   | 跨区域法兰克福终端节点
         | s3.mil.eu.cloud-object-storage.appdomain.cloud                   | 跨区域米兰终端节点
         | s3.ams.eu.cloud-object-storage.appdomain.cloud                   | 跨区域阿姆斯特丹终端节点
         | s3.private.eu.cloud-object-storage.appdomain.cloud               | 跨区域欧洲私有终端节点
         | s3.private.fra.eu.cloud-object-storage.appdomain.cloud           | 跨区域法兰克福私有终端节点
         | s3.private.mil.eu.cloud-object-storage.appdomain.cloud           | 跨区域米兰私有终端节点
         | s3.private.ams.eu.cloud-object-storage.appdomain.cloud           | 跨区域阿姆斯特丹私有终端节点
         | s3.eu-gb.cloud-object-storage.appdomain.cloud                    | 英国终端节点
         | s3.private.eu-gb.cloud-object-storage.appdomain.cloud            | 英国私有终端节点
         | s3.eu-de.cloud-object-storage.appdomain.cloud                    | 欧洲德国区域终端节点
         | s3.private.eu-de.cloud-object-storage.appdomain.cloud            | 欧洲德国区域私有终端节点
         | s3.ap.cloud-object-storage.appdomain.cloud                       | 跨区域亚太终端节点
         | s3.tok.ap.cloud-object-storage.appdomain.cloud                   | 跨区域东京终端节点
         | s3.hkg.ap.cloud-object-storage.appdomain.cloud                   | 跨区域香港终端节点
         | s3.seo.ap.cloud-object-storage.appdomain.cloud                   | 跨区域首尔终端节点
         | s3.private.ap.cloud-object-storage.appdomain.cloud               | 跨区域亚太私有终端节点
         | s3.private.tok.ap.cloud-object-storage.appdomain.cloud           | 跨区域东京私有终端节点
         | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud           | 跨区域香港私有终端节点
         | s3.private.seo.ap.cloud-object-storage.appdomain.cloud           | 跨区域首尔私有终端节点
         | s3.jp-tok.cloud-object-storage.appdomain.cloud                   | 亚太区域日本终端节点
         | s3.private.jp-tok.cloud-object-storage.appdomain.cloud           | 亚太区域日本私有终端节点
         | s3.au-syd.cloud-object-storage.appdomain.cloud                   | 亚太区域澳大利亚终端节点
         | s3.private.au-syd.cloud-object-storage.appdomain.cloud           | 亚太区域澳大利亚私有终端节点
         | s3.ams03.cloud-object-storage.appdomain.cloud                    | 阿姆斯特丹单区终端节点
         | s3.private.ams03.cloud-object-storage.appdomain.cloud            | 阿姆斯特丹单区私有终端节点
         | s3.che01.cloud-object-storage.appdomain.cloud                    | 金奈单区终端节点
         | s3.private.che01.cloud-object-storage.appdomain.cloud            | 金奈单区私有终端节点
         | s3.mel01.cloud-object-storage.appdomain.cloud                    | 墨尔本单区终端节点
         | s3.private.mel01.cloud-object-storage.appdomain.cloud            | 墨尔本单区私有终端节点
         | s3.osl01.cloud-object-storage.appdomain.cloud                    | 奥斯陆单区终端节点
         | s3.private.osl01.cloud-object-storage.appdomain.cloud            | 奥斯陆单区私有终端节点
         | s3.tor01.cloud-object-storage.appdomain.cloud                    | 多伦多单区终端节点
         | s3.private.tor01.cloud-object-storage.appdomain.cloud            | 多伦多单区私有终端节点
         | s3.seo01.cloud-object-storage.appdomain.cloud                    | 首尔单区终端节点
         | s3.private.seo01.cloud-object-storage.appdomain.cloud            | 首尔单区私有终端节点
         | s3.mon01.cloud-object-storage.appdomain.cloud                    | 蒙特利尔单区终端节点
         | s3.private.mon01.cloud-object-storage.appdomain.cloud            | 蒙特利尔单区私有终端节点
         | s3.mex01.cloud-object-storage.appdomain.cloud                    | 墨西哥单区终端节点
         | s3.private.mex01.cloud-object-storage.appdomain.cloud            | 墨西哥单区私有终端节点
         | s3.sjc04.cloud-object-storage.appdomain.cloud                    | 圣何塞单区终端节点
         | s3.private.sjc04.cloud-object-storage.appdomain.cloud            | 圣何塞单区私有终端节点
         | s3.mil01.cloud-object-storage.appdomain.cloud                    | 米兰单区终端节点
         | s3.private.mil01.cloud-object-storage.appdomain.cloud            | 米兰单区私有终端节点
         | s3.hkg02.cloud-object-storage.appdomain.cloud                    | 香港单区终端节点
         | s3.private.hkg02.cloud-object-storage.appdomain.cloud            | 香港单区私有终端节点
         | s3.par01.cloud-object-storage.appdomain.cloud                    | 巴黎单区终端节点
         | s3.private.par01.cloud-object-storage.appdomain.cloud            | 巴黎单区私有终端节点
         | s3.sng01.cloud-object-storage.appdomain.cloud                    | 新加坡单区终端节点
         | s3.private.sng01.cloud-object-storage.appdomain.cloud            | 新加坡单区私有终端节点

   --location-constraint
      区域约束 - 在使用IBM Cloud Public时必须与终端点匹配。
      
      对于本地COS，请不要从此列表中选择，而是按回车键。

      示例:
         | us-standard       | 跨区域标准版
         | us-vault          | 跨区域Vault存储
         | us-cold           | 跨区域Cold存储
         | us-flex           | 跨区域Flex存储
         | us-east-standard  | 美国东部标准版
         | us-east-vault     | 美国东部Vault存储
         | us-east-cold      | 美国东部Cold存储
         | us-east-flex      | 美国东部Flex存储
         | us-south-standard | 美国南部标准版
         | us-south-vault    | 美国南部Vault存储
         | us-south-cold     | 美国南部Cold存储
         | us-south-flex     | 美国南部Flex存储
         | eu-standard       | 跨区域欧洲标准版
         | eu-vault          | 跨区域欧洲Vault存储
         | eu-cold           | 跨区域欧洲Cold存储
         | eu-flex           | 跨区域欧洲Flex存储
         | eu-gb-standard    | 英国标准版
         | eu-gb-vault       | 英国Vault存储
         | eu-gb-cold        | 英国Cold存储
         | eu-gb-flex        | 英国Flex存储
         | ap-standard       | 跨区域亚太标准版
         | ap-vault          | 跨区域亚太Vault存储
         | ap-cold           | 跨区域亚太Cold存储
         | ap-flex           | 跨区域亚太Flex存储
         | mel01-standard    | 墨尔本标准版
         | mel01-vault       | 墨尔本Vault存储
         | mel01-cold        | 墨尔本Cold存储
         | mel01-flex        | 墨尔本Flex存储
         | tor01-standard    | 多伦多标准版
         | tor01-vault       | 多伦多Vault存储
         | tor01-cold        | 多伦多Cold存储
         | tor01-flex        | 多伦多Flex存储

   --acl
      创建存储桶和存储或复制对象时使用的预定义ACL。
      
      对于创建对象，如果未设置bucket_acl，则使用此ACL。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      请注意，在通过S3服务器复制对象时将应用此ACL，因为S3不会复制源中的ACL，而是写入一份新的ACL。
      
      如果ACL为空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

      示例:
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 没有其他用户具有访问权限（默认值）。
         |                    | 此acl在IBM Cloud (Infra)、IBM Cloud (Storage)、On-Premise COS上都可用。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有READ权限。
         |                    | 此acl在IBM Cloud (Infra)、IBM Cloud (Storage)、On-Premise IBM COS上都可用。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有READ和WRITE权限。
         |                    | 此acl在IBM Cloud (Infra)、On-Premise IBM COS上都可用。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有READ权限。
         |                    | 不支持桶。
         |                    | 此acl在IBM Cloud (Infra)和On-Premise IBM COS上都可用。

   --bucket-acl
      创建存储桶时使用的预定义ACL。
      
      有关详细信息，请访问https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      如果未设置bucket_acl，则仅在创建存储桶时应用此ACL。
      
      如果"acl"和"bucket_acl"为空字符串，则不会添加X-Amz-Acl:头，并且将使用默认值（private）。

      示例:
         | private            | 所有者具有FULL_CONTROL权限。
         |                    | 没有其他用户具有访问权限（默认值）。
         | public-read        | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有READ权限。
         | public-read-write  | 所有者具有FULL_CONTROL权限。
         |                    | AllUsers组具有READ和WRITE权限。
         |                    | 通常不建议在存储桶上授予此权限。
         | authenticated-read | 所有者具有FULL_CONTROL权限。
         |                    | AuthenticatedUsers组具有READ权限。

   --upload-cutoff
      切换至分块上传的截止点。
      
      大于此大小的任何文件将分块传输，分块大小为chunk_size。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff的文件或大小不确定的文件（例如使用"rclone rcat"上传的文件或使用"rclone mount"或google
      photos或google docs上传的文件）时，将使用此块大小进行分块上传。
      
      请注意，每个传输的内存会缓冲 "--s3-upload-concurrency" 块大小。
      
      如果您正在高速链路上传输大文件并且具有足够的内存，则增加此大小将加快传输速度。
      
      Rclone将自动增加块大小，以便在上传已知大小的大文件时保持在10000块限制以下。
      
      未知大小的文件将使用配置的
      chunk_size上传。由于默认的块大小为5 MiB，最多可以有10,000个块，这意味着默认情况下您可以流式传输的文件的最大大小为48 GiB。如果要流式传输
      更大的文件，则需要增加chunk_size。
      
      增加块大小将降低使用"-P"标志显示的进度统计的精度。当
      Rclone将块视为发送时，它是由AWS SDK缓冲的，而实际上可能仍在上传。
      更大的块大小意味着更大的AWS SDK缓冲区和与实际情况更不符的进度
      报告。
      

   --max-upload-parts
      多部分上传中的最大部分数。
      
      此选项定义多部分上传时要使用的最大多部分块数。
      
      如果服务不支持AWS S3的10000块规范，这可能会很有用。
      
      Rclone将自动增加块大小，以便在上传已知大小的大文件时保持在此块数限制以下。
      

   --copy-cutoff
      切换至多部分复制的截止点。
      
      大于此大小的需要服务器端复制的文件将以此大小的块复制。
      
      最小值为0，最大值为5 GiB。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常，在上传之前，rclone会计算输入的MD5校验和，以便将其添加到对象的元数据中。这非常适合数据完整性检查，但可能导致大文件的上传开始时间较长。

   --shared-credentials-file
      共享凭据文件的路径。
      
      如果 env_auth=true，则rclone可以使用共享凭据文件。
      
      如果该变量为空，则rclone将查找
      "AWS_SHARED_CREDENTIALS_FILE" 环境变量。如果环境变量的值为空，它将默认为当前用户的主目录。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      在共享凭据文件中要使用的配置文件。
      
      如果 env_auth=true，则rclone可以使用共享凭据文件。此
      变量控制在文件中使用哪个配置文件。
      
      如果为空，它将默认为环境变量 "AWS_PROFILE" 或 "default" 如果该环境变量也未设置。
      

   --session-token
      AWS会话令牌。

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传的相同文件的块数。
      
      如果您正在高速链路上上传少量大文件，并且这些上传没有充分利用您的带宽，则增加此数值可能有助于加速传输。

   --force-path-style
      如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。
      
      如果为true（默认值），则rclone将使用路径样式访问，
      如果为false，则rclone将使用虚拟路径样式。请参阅[Amazon S3文档]（https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro）了解更多信息。
      
      某些提供商（例如AWS、Aliyun OSS、Netease COS或Tencent COS）需要将该设置为
      false - rclone将根据提供程序自动执行此操作
      设置。

   --v2-auth
      如果为true，则使用v2身份验证。
      
      如果为false（默认值），则rclone将使用v4身份验证。如果设置，则rclone将使用v2身份验证。
      
      仅在v4签名不起作用时使用，例如，Jewel/v10 CEPH之前的情况。

   --list-chunk
      列表分块的大小（每个ListObject S3请求的响应列表）。
      
      此选项也称为AWS S3规范中的"MaxKeys"、"max-items"或"page-size"。
      大多数服务会将响应列表截断为1000个对象，即使请求的数量更多。
      在AWS S3中，这是一个全局最大值，不能更改，请参阅[Amazon S3]（https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html）。
      在Ceph中，可以使用"rgw list buckets max chunk"选项进行增加。
      

   --list-version
      要使用的ListObjects版本：1、2或0表示自动。
      
      当最初发布S3时，只提供了ListObjects调用来枚举存储桶中的对象。
      
      但是，在2016年5月引入了ListObjectsV2调用。此调用性能更好，应尽可能使用。
      
      如果设置为默认值0，则rclone将根据提供程序设置猜测要调用哪个对象列表方法。如果猜测错误，则可以在此处手动设置。
      

   --list-url-encode
      是否对列表进行URL编码：true/false/unset
      
      某些提供商支持URL编码列表，如果可用，这是在文件
      名中使用控制字符时更可靠。如果设置为unset（默认值），则
      根据提供商设置，rclone将选择使用什么来应用，但可以在此处覆盖
      rclone的选择。
      

   --no-check-bucket
      如果设置，不要尝试检查存储桶是否存在或创建存储桶。
      
      如果知道存储桶已经存在，并且希望将rclone执行的事务数量最小化，这可能很有用。
      
      如果您使用的用户没有bucket
      创建权限，则可能需要这样做。v1.52.0之前的版本会因为一个错误而静默通过。
      

   --no-head
      如果设置，则不要HEAD已上传的对象以进行完整性检查。
      
      如果要将rclone在上传对象后收到200 OK消息，则假设它已正确上传，设置此标志。
      
      特别是，它将假定：
      
      - 元数据，包括修改时间、存储类和内容类型与上传的相同
      - 大小与上传的相同
      
      它从PUT请求的响应中读取以下项：
      
      - MD5SUM
      - 上传日期
      
      对于多部分上传，不会读取这些项。
      
      如果上传的源对象大小未知，则rclone **将**执行HEAD请求。
      
      设置此标志将增加未检测到的上传失败的概率，
      特别是错误的大小，因此不建议在正常操作中使用。实际上，即使在设置此标志的情况下，检测到上传失败的几率非常小。
      

   --no-head-object
      如果设置，获取对象之前不执行HEAD。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述中的编码方式]（/overview/#encoding）。

   --memory-pool-flush-time
      内部内存缓冲池刷新的频率。
      
      需要额外的缓冲区（例如多部分）上传将使用内存池进行分配。
      此选项控制多久将未使用的缓冲区从池中移除。

   --memory-pool-use-mmap
      是否在内部内存池中使用内存映射缓冲区。

   --disable-http2
      禁用S3后端的http2使用。
      
      目前，s3（特别是minio）后端存在一个未解决的问题，与HTTP/2有关。S3后端默认启用HTTP/2，但可以在此禁用。当问题解决后，将删除此标志。
      
      请参阅：https://github.com/rclone/rclone/issues/4673，https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      自定义下载的终点。
      通常将其设置为CloudFront CDN URL，因为AWS S3通过CloudFront网络下载的数据有更低的出口费用。

   --use-multipart-etag
      是否在多部分上传中使用ETag进行验证
      
      这应该是true、false或未设置以使用提供程序的默认值。
      

   --use-presigned-request
      是否使用预签名请求或PutObject进行单部分上传
      
      如果为false，则rclone将使用AWS SDK中的PutObject来上传对象。
      
      rclone的版本<1.59使用预签名请求来上传单部分对象，将此标志设置为true将重新启用该功能。除非在特殊情况下或者进行测试，否则不需要此标志。
      

   --versions
      在目录列表中包含旧版本。

   --version-at
      显示指定时间点的文件版本。
      
      参数应为日期格式（"2006-01-02"）或时间格式（"2006-01-02 15:04:05"），
      或者是这么久以前的持续时间，例如"100d"或"1h"。
      
      请注意，在使用此设置时，不允许执行文件写操作，
      因此无法上传或删除文件。
      
      有关有效格式，请参阅[时间选项文档]（/docs/#time-option）。
      

   --decompress
      如果设置，则将解压缩gzip编码的对象。
      
      可以使用"Content-Encoding: gzip"设置将对象上传到S3。通常情况下，rclone会以压缩的形式下载这些文件。
      
      如果设置此标志，则rclone将在接收到以"Content-Encoding: gzip"形式的文件时解压缩这些文件。这意味着rclone无法检查文件大小和哈希值，但文件内容将被解压缩。
      

   --might-gzip
      如果后端可能会压缩对象，请设置此项。
      
      通常情况下，提供者在下载对象时不会更改对象。如果
      对象未使用"Content-Encoding: gzip"上传，那么在下载时
      不会设置它。
      
      但某些提供者即使没有使用"Content-Encoding: gzip"上传对象（例如Cloudflare）也会对其进行gzip压缩。
      
      这样做的症状将收到以下类似错误的错误消息：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      如果设置此标志并且rclone以设置了
      "Content-Encoding: gzip" 和分块传输编码的对象下载了对象，那么rclone将会即时解压缩该对象。
      
      如果将其设置为unset（默认值），则rclone将选择根据提供者的设置应用何种，但您可以在此处覆盖rclone的选择。
      

   --no-system-metadata
      禁止设置和读取系统元数据


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  创建存储桶和存储或复制对象时使用的预定义ACL。 [$ACL]
   --endpoint value             IBM COS S3 API的终端点。 [$ENDPOINT]
   --env-auth                   从运行时获取AWS凭证（环境变量或EC2/ECS元数据，如果没有环境变量）。 (default: false) [$ENV_AUTH]
   --help, -h                   显示帮助
   --location-constraint value  区域约束 - 在使用IBM Cloud Public时必须与终端点匹配。 [$LOCATION_CONSTRAINT]
   --region value               连接的区域。 [$REGION]
   --secret-access-key value    AWS Secret Access Key（密码）。 [$SECRET_ACCESS_KEY]

   高级

   --bucket-acl value               创建存储桶时使用的预定义ACL。 [$BUCKET_ACL]
   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换至多部分复制的截止点。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     如果设置，则将解压缩gzip编码的对象。 (default: false) [$DECOMPRESS]
   --disable-checksum               不要与对象元数据一起存储MD5校验和。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  禁用S3后端的http2使用。 (default: false) [$DISABLE_HTTP2]
   --download-url value             自定义下载的终点。 [$DOWNLOAD_URL]
   --encoding value                 后端的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               如果为true，则使用路径样式访问；如果为false，则使用虚拟主机样式访问。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               列表分块的大小（每个ListObject S3请求的响应列表）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          是否对列表进行URL编码：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             要使用的ListObjects版本：1,2 or 0 for auto. (default: 0) [$LIST_VERSION]
   --max-upload-parts value         多部分上传中的最大部分数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部内存缓冲池刷新的频率。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用内存映射缓冲区。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               如果后端可能会压缩对象，请设置此项。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                如果设置，不要尝试检查存储桶是否存在或创建存储桶。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        如果设置，则不要HEAD已上传的对象以进行完整性检查。 (default: false) [$NO_HEAD]
   --no-head-object                 如果设置，获取对象之前不执行HEAD。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             禁止设置和读取系统元数据 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  在共享凭据文件中要使用的配置文件。 ($PROFILE]
   --session-token value            AWS会话令牌。 [$SESSION_TOKEN]
   --shared-credentials-file value  共享凭据文件的路径。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       多部分上传的并发数。  (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换至分块上传的截止点。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       是否在多部分上传中使用ETag进行验证 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          是否使用预签名请求或PutObject进行单部分上传 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        如果为true，则使用v2身份验证。 (default: false) [$V2_AUTH]
   --version-at value               显示指定时间点的文件版本。 (default: "off") [$VERSION_AT]
   --versions                       在目录列表中包含旧版本。 (default: false) [$VERSIONS]

   General

   --name value  存储的名称（默认为自动生成的）
   --path value  存储的路径

```
{% endcode %}