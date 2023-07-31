# Oracle Cloud Infrastructure 对象存储

{% code fullWidth="true" %}
```
命令名称：
   singularity datasource add oos - Oracle Cloud Infrastructure 对象存储

使用方法：
   singularity datasource add oos [命令选项] <数据集名称> <源路径>

描述：
   --oos-chunk-size
      用于上传的块大小。
      
      当上传大于上传截断或大小未知的文件（例如从“rclone rcat”或通过“rclone mount”或 Google 照片或 Google 文档上传）时，将使用该块大小执行多部分上传。
      
      请注意，“upload_concurrency”个这样大小的块将在每次传输时在内存中进行缓冲。
      
      如果您在高速链路上传输大文件并且内存足够，则增加此值将加快传输速度。
      
      当上传已知大小的大文件时，rclone 会自动增加块大小，以保持在10,000个块的限制以下。
      
      未知大小的文件使用配置的块大小进行上传。由于默认的块大小为5 MiB，且最多可以有10,000个块，因此默认情况下，您可以流式上传的文件的最大大小为48 GiB。如果要流式上传更大的文件，则需要增加块大小。
      
      增加块大小会降低通过“-P”标志显示的进度统计信息的准确性。

   --oos-compartment
      [Provider] - user_principal_auth
         对象存储区划 OCID

   --oos-config-file
      [Provider] - user_principal_auth
         OCI 配置文件的路径

         示例：
            | ~/.oci/config | oci 配置文件位置

   --oos-config-profile
      [Provider] - user_principal_auth
         oci 配置文件中的配置文件名称

         示例：
            | Default | 使用默认配置文件

   --oos-copy-cutoff
      用于切换到多分块复制的截断点。
      
      需要服务器端复制的任何大于此大小的文件将按照该大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --oos-copy-timeout
      复制的超时时间。
      
      复制是一个异步操作，请指定在成功复制前等待的超时时间。

   --oos-disable-checksum
      不将 MD5 校验和与对象元数据一起存储。
      
      通常，rclone 在上传之前会计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件的起始上传可能会导致长时间延迟。

   --oos-encoding
      后端的编码。
      
      有关更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --oos-endpoint
      对象存储 API 的终节点。
      
      留空以使用该区域的默认终节点。

   --oos-leave-parts-on-error
      如果为 true，则在失败时避免调用中止上传，并将所有成功上传的部分留在 S3 上以进行手动恢复。
      
      对于在不同会话之间恢复上传，应设置为 true。
      
      警告：如果不清除部分不完整的多重部分上传，则它会计入对象存储上的空间使用情况，并增加额外的费用。

   --oos-namespace
      对象存储命名空间

   --oos-no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。
      
      这在尝试最小化 rclone 进行的事务数量时可能很有用，如果已知存储桶已存在，则可以使用此选项。
      
      如果正在使用的用户没有存储桶创建权限，则可能也需要此选项。

   --oos-provider
      选择您的身份验证提供者

      示例：
         | env_auth                | 自动从运行时（env）中选择凭据，首先提供身份验证的凭据将生效
         | user_principal_auth     | 使用 OCI 用户和 API 密钥进行身份验证。
                                   | 您需要在配置文件中放入租户 OCID、用户 OCID、区域、路径和 API 密钥的指纹。
                                   | https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
         | instance_principal_auth | 使用实例原则授权实例进行 API 调用。
                                   | 每个实例都有自己的标识，并使用从实例元数据中读取的证书进行身份验证。
                                   | https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
         | resource_principal_auth | 使用资源原则进行 API 调用
         | no_auth                 | 无需凭据，这通常用于读取公共存储桶

   --oos-region
      对象存储区域

   --oos-sse-customer-algorithm
      如果使用 SSE-C，则此可选标头指定“AES256”作为加密算法。
      对象存储支持“AES256”作为加密算法。有关更多信息，请参见
      使用您自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <未设置> | 无
         | AES256  | AES256

   --oos-sse-customer-key
      若要使用 SSE-C，请指定 base64 编码的256位加密密钥进行加密或解密数据的可选标头。请注意只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。有关更多信息，请参见使用您自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <未设置> | 无

   --oos-sse-customer-key-file
      若要使用 SSE-C，请指定包含与对象关联的 AES-256 加密密钥的base64 编码字符串的文件。请注意只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      示例：
         | <未设置> | 无

   --oos-sse-customer-key-sha256
      如果使用 SSE-C，则此可选标头指定加密密钥的 base64 编码的 SHA256 哈希。此值用于检查加密密钥的完整性。有关更多信息，请参见使用您自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <未设置> | 无

   --oos-sse-kms-key-id
      如果在保险库中使用自己的主密钥，则此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)。请注意只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      示例：
         | <未设置> | 无

   --oos-storage-tier
      存储新对象时要使用的存储类别。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例：
         | Standard         | 标准存储类别，这是默认类别
         | InfrequentAccess | 低频访问存储类别
         | Archive          | 归档存储类别

   --oos-upload-concurrency
      多部分上传的并发数。
      
      这是同时上传同一文件的块数。
      
      如果通过高速链路上传少量大文件，并且这些上传无法充分利用带宽，则增加此值可能有助于加快传输速度。

   --oos-upload-cutoff
      切换到分块上传的截断点。
      
      大于此大小的任何文件都将按照 chunk_size 的大小进行分块上传。
      最小值为0，最大值为5 GiB。


选项：
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险] 在将数据集导出到 CAR 文件后删除数据集的文件。 (默认值：false)
   --rescan-interval value  当距上次成功扫描此时间间隔已过时，自动重新扫描源目录 (默认值：禁用)
   --scanning-state value   设置初始扫描状态 (默认值：ready)

   对于 oos 的选项

   --oos-chunk-size value               用于上传的块大小。 (默认值："5Mi") [$OOS_CHUNK_SIZE]
   --oos-compartment value              对象存储区划 OCID [$OOS_COMPARTMENT]
   --oos-config-file value              OCI 配置文件的路径 (默认值："~/.oci/config") [$OOS_CONFIG_FILE]
   --oos-config-profile value           oci 配置文件中的配置文件名称 (默认值："Default") [$OOS_CONFIG_PROFILE]
   --oos-copy-cutoff value              用于切换到多分块复制的截断点。 (默认值："4.656Gi") [$OOS_COPY_CUTOFF]
   --oos-copy-timeout value             复制的超时时间。 (默认值："1m0s") [$OOS_COPY_TIMEOUT]
   --oos-disable-checksum value         不将 MD5 校验和与对象元数据一起存储。 (默认值："false") [$OOS_DISABLE_CHECKSUM]
   --oos-encoding value                 后端的编码。 (默认值："Slash,InvalidUtf8,Dot") [$OOS_ENCODING]
   --oos-endpoint value                 对象存储 API 的终节点。 [$OOS_ENDPOINT]
   --oos-leave-parts-on-error value     如果为 true，则在失败时避免调用中止上传，并将所有成功上传的部分留在 S3 上以进行手动恢复。 (默认值："false") [$OOS_LEAVE_PARTS_ON_ERROR]
   --oos-namespace value                对象存储命名空间 [$OOS_NAMESPACE]
   --oos-no-check-bucket value          如果设置，则不尝试检查存储桶是否存在或创建存储桶。 (默认值："false") [$OOS_NO_CHECK_BUCKET]
   --oos-provider value                 选择您的身份验证提供者 (默认值："env_auth") [$OOS_PROVIDER]
   --oos-region value                   对象存储区域 [$OOS_REGION]
   --oos-sse-customer-algorithm value   如果使用 SSE-C，则此可选标头指定 "AES256" 作为加密算法。 [$OOS_SSE_CUSTOMER_ALGORITHM]
   --oos-sse-customer-key value         若要使用 SSE-C，请指定 base64 编码的256位加密密钥进行对称键加密或解密数据的可选标头。请注意只需 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。有关更多信息，请参见使用您自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。
   --oos-sse-customer-key-file value    若要使用 SSE-C，请指定包含与对象关联的 AES-256 加密密钥的 base64 编码字符串的文件。请注意只需 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。缺少请使用<sunset>代替，无请使用<unset>代替。
   --oos-sse-customer-key-sha256 value  如果使用 SSE-C，则此可选标头指定加密密钥的 base64 编码的 SHA256 哈希。此值用于检查加密密钥的完整性。有关更多信息，请参见使用您自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。
   --oos-sse-kms-key-id value           如果在保险库中使用自己的主密钥，则此标头指定其 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)。请注意只需 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。缺少请使用<sunset>代替，无请使用<unset>代替。
   --oos-storage-tier value             存储新对象时要使用的存储类别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm###Options" (默认值："Standard") [$OOS_STORAGE_TIER]
   --oos-upload-concurrency value       多部分上传的并发数。 (默认值："10") [$OOS_UPLOAD_CONCURRENCY]
   --oos-upload-cutoff value            用于切换到分块上传的截断点。(默认值："200Mi") [$OOS_UPLOAD_CUTOFF]

```
{% endcode %}