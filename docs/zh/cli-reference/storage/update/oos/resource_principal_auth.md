# 使用资源主体进行API调用

{% code fullWidth="true" %}
```
名称:
   singularity storage update oos resource_principal_auth - 使用资源主体进行API调用

用法:
   singularity storage update oos resource_principal_auth [命令选项] <名称|ID>

描述:
   --namespace
      对象存储命名空间

   --compartment
      对象存储区域OCID

   --region
      对象存储区域

   --endpoint
      对象存储API的终端点。
      
      留空以使用区域的默认终端点。

   --storage-tier
      存储新对象时要使用的存储级别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例：
         | Standard         | 标准的存储级别，这是默认的级别
         | InfrequentAccess | 不经常访问的存储级别
         | Archive          | 存档的存储级别

   --upload-cutoff
      切换到分块上传的截止点。
      
      大于此大小的任何文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时要使用的块大小。
      
      当上传大于upload_cutoff的文件或大小未知的文件时（例如，使用"rclone rcat"命令或通过"rclone mount"或谷歌照片或谷歌文档上传的文件），将使用此块大小进行分块上传。
      
      请注意，每个传输中缓冲大小为"upload_concurrency"的chunk大小。
      
      如果您正在通过高速链接传输大文件，并且有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件时，Rclone将自动增加块大小，以保持在10000块的限制以下。
      
      未知大小的文件以配置的chunk_size进行上传。由于默认的chunk_size是5 MiB，并且最多有10000个chunk，这意味着默认情况下，您可以流式传输的文件的最大大小为48 GiB。如果您希望流式传输更大的文件，则需要增加chunk_size。
      
      增加块大小会降低使用"-P"标志显示的进度统计的准确性。
      

   --upload-concurrency
      并发进行分块上传的数量。
      
      如果您正在通过高速链接上传少量的大文件，并且这些上传未充分利用带宽，则增加此值可能有助于加快传输速度。

   --copy-cutoff
      切换到分块复制的截止点。
      
      大于此大小并且需要进行服务器端复制的任何文件将以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制操作的超时时间。
      
      复制是异步操作，指定超时时间以等待复制成功。

   --disable-checksum
      不要将MD5校验和与对象的元数据一起存储。
      
      通常，在上传之前，rclone会计算输入的MD5校验和，并将其添加到对象的元数据中。这对于数据完整性检查很有用，但可能会导致大文件开始上传的长时间延迟。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error
      如果设置为true，则在出现故障时避免调用中止上传，以便将所有成功上传的部分留在S3上供手动恢复。
      
      对于在不同会话之间恢复上传时，应将其设置为true。
      
      警告：存储未完成的分块上传的部分会导致对象存储的空间使用量计入，并且如果未清理，将增加额外成本。

   --no-check-bucket
      如果设置，不尝试检查存储桶是否存在或创建它。
      
      如果您知道存储桶已经存在，这可能有助于最小化rclone执行的事务数。
      
      如果使用的用户没有创建存储桶的权限，可能也需要此选项。

   --sse-customer-key-file
      要使用SSE-C，可以将包含与对象关联的AES-256加密密钥的base64编码字符串的文件指定为该选项。
      请注意，只需要在sse_customer_key_file|sse_customer_key|sse_kms_key_id中选择一个。

      示例：
         | <unset> | 无

   --sse-customer-key
      要使用SSE-C，可以将base64编码的256位加密密钥的可选头部指定为该选项。
      请注意，只需要在sse_customer_key_file|sse_customer_key|sse_kms_key_id中选择一个。
      有关更多信息，请参见使用您自己的密钥进行服务端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <unset> | 无

   --sse-customer-key-sha256
      如果使用SSE-C，可选头部指定了加密密钥的base64编码的SHA256哈希。
      此值用于检查加密密钥的完整性。有关使用您自己的密钥进行服务端加密的更多信息，请参见 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <unset> | 无

   --sse-kms-key-id
      如果在金库中使用您自己的主密钥，则此选项指定用于调用密钥管理服务生成数据加密密钥或加密/解密数据加密密钥的主加密密钥的OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)。
      请注意，只需要在sse_customer_key_file|sse_customer_key|sse_kms_key_id中选择一个。

      示例：
         | <unset> | 无

   --sse-customer-algorithm
      如果使用SSE-C，可选头部将"AES256"指定为加密算法。
      对象存储支持"AES256"作为加密算法。有关更多信息，请参见使用您自己的密钥进行服务端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例：
         | <unset> | 无
         | AES256  | AES256


选项:
   --compartment 值  对象存储区域OCID [$COMPARTMENT]
   --endpoint 值     对象存储API的终端点 [$ENDPOINT]
   --help, -h       显示帮助信息
   --namespace 值    对象存储命名空间 [$NAMESPACE]
   --region 值       对象存储区域 [$REGION]

   高级选项

   --chunk-size 值               用于上传的块大小。 (默认值："5Mi") [$CHUNK_SIZE]
   --copy-cutoff 值              切换到分块复制的截止点。 (默认值："4.656Gi") [$COPY_CUTOFF]
   --copy-timeout 值             复制操作的超时时间。 (默认值："1m0s") [$COPY_TIMEOUT]
   --disable-checksum            不要将MD5校验和与对象的元数据一起存储。 (默认值：false) [$DISABLE_CHECKSUM]
   --encoding 值                 后端的编码方式。 (默认值："Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error         如果为true，则在出现故障时避免调用中止上传，以便将所有成功上传的部分留在S3上供手动恢复。 (默认值：false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket              如果设置，不尝试检查存储桶是否存在或创建它。 (默认值：false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm 值   如果使用SSE-C，可选头部将"AES256"指定为加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key 值         要使用SSE-C，可选头部将base64编码的256位加密密钥指定为 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file 值    要使用SSE-C，可以将包含与对象关联的AES-256加密密钥的base64编码字符串的文件指定为 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 值  如果使用SSE-C，则可选头部指定加密密钥的base64编码的SHA256哈希 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id 值           如果在金库中使用您自己的主密钥，则此头部指定了用于调用密钥管理服务生成数据加密密钥或加密/解密数据加密密钥的主加密密钥的OCID [$SSE_KMS_KEY_ID]
   --storage-tier 值             存储新对象时要使用的存储级别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (默认值："Standard") [$STORAGE_TIER]
   --upload-concurrency 值       并发进行分块上传的数量。 (默认值：10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 值            切换到分块上传的截止点。 (默认值："200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}