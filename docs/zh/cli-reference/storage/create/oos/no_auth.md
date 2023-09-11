# 不需要凭据，通常适用于读取公共存储桶

{% code fullWidth="true" %}
```
命令名称:
   创建存储框架OOS no_auth-不需要凭据，通常适用于读取公共存储桶

用法:
   singularity storage create oos no_auth [命令选项] [参数...]

说明:
   --namespace
      对象存储的命名空间

   --region
      对象存储的区域

   --endpoint
      对象存储 API 的终端点。
      
      留空以使用该区域的默认终端点。

   --storage-tier
      存储新对象时要使用的存储等级。参见https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例:
         | Standard         | 标准存储等级，这是默认等级
         | InfrequentAccess | 低频存储等级
         | Archive          | 存档存储等级

   --upload-cutoff
      切换为分块上传的大小阈值。
      
      大于此大小的任何文件将以块大小的形式上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时要使用的块大小。
      
      当上载大于 upload_cutoff 的文件或大小不明确的文件（例如通过“rclone rcat”或通过“rclone mount”或谷歌
      照片或谷歌文档上传的文件）时，将使用这个块大小进行分块上传。
      
      请注意，每个传输的内存中都会缓存此大小的"upload_concurrency"块。
      
      如果您正在通过高速链接传输大文件并且有足够的内存，则增加此值将加快传输速度。
      
      当Rclone将较大的已知大小的文件上传时，将自动增加块大小，以保持在10,000块的限制范围内。
      
      未知大小的文件使用配置的 chunk_size 进行上传。由于默认的 chunk_size 为 5 MiB，并且最多可以有10,000个块，所以默认情况下可以流式上传的文件的最大大小为48 GiB。 
      如果您希望流式上传更大的文件，则需要增加Chunk大小。
      
      增加块大小会降低带有“-P”标志显示的进度统计的准确性。
      

   --upload-concurrency
      分块上传的并发数。
      
      同时上传相同文件的多个块。
      
      如果您正在通过高速链接上传少量较大文件，并且这些上传未完全利用您的带宽，那么增加此值可以加快传输速度。

   --copy-cutoff
      切换为分块复制的大小阈值。
      
      大于此大小的需要在服务器端复制的文件将以此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制操作的超时时间。
      
      复制是异步操作，请指定超时时间以等待复制成功。
      

   --disable-checksum
      不要将 MD5 校验和存储在对象元数据中。
      
      通常，rclone在上传之前会计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件来说，可能会导致长时间的延迟才能开始上传。

   --encoding
      后端的编码。
      
      有关更多信息，请参见概述中的 [编码章节](/overview/#encoding)。

   --leave-parts-on-error
      如果为 true，则在失败时避免调用"abort upload"，将所有已成功上载的部分保留在 S3 中供手动恢复。
      
      对于在不同会话之间恢复上传，应将其设置为 true。
      
      警告: 将不完整的多部分上传的部分保存在对象存储中会计入空间使用，并在没有清理的情况下增加额外成本。
      

   --no-check-bucket
      如果设置了该项，则不尝试检查存储桶是否存在或创建存储桶。
      
      当尝试尽量减少 rclone 的事务数时，这可能非常有用，如果您知道存储桶已经存在。
      
      如果使用的用户没有存储桶创建权限，可能也需要设置该项。
      

   --sse-customer-key-file
      要使用 SSE-C，文件中包含与对象关联的 AES-256 加密密钥的 base64 编码字符串。请注意，需要 sse_customer_key_file | sse_customer_key | sse_kms_key_id 中的一个。

      示例:
         | <unset> | None

   --sse-customer-key
      要使用 SSE-C，可选头部，指定要用于加密或解密数据的 base64 编码的 256 位加密密钥。请注意，需要 sse_customer_key_file | sse_customer_key | sse_kms_key_id 中的一个。有关更多信息，请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      示例:
         | <unset> | None

   --sse-customer-key-sha256
      如果使用 SSE-C，可选头部，指定加密密钥的 base64 编码的 SHA256 哈希值。该值用于检查加密密钥的完整性。请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      示例:
         | <unset> | None

   --sse-kms-key-id
      如果在保管库中使用自己的主密钥，则此头部指定使用的主加密密钥的 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)，
      用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥。请注意，需要 sse_customer_key_file | sse_customer_key | sse_kms_key_id 中的一个。

      示例:
         | <unset> | None

   --sse-customer-algorithm
      如果使用 SSE-C，可选头部，指定加密算法为 "AES256"。
      对象存储支持 "AES256" 作为加密算法。有关更多信息，请参见使用自己的密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None
         | AES256  | AES256


选项:
   --endpoint value   对象存储 API 的终端点。[$ENDPOINT]
   --help, -h         显示帮助信息
   --namespace value  对象存储的命名空间。[$NAMESPACE]
   --region value     对象存储的区域。[$REGION]

   高级选项

   --chunk-size value               用于上传的块大小。 (默认值: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的大小阈值。 (默认值: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             复制操作的超时时间。 (默认值: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               不要将 MD5 校验和存储在对象元数据中。 (默认值: false) [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码。 (默认值: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           如果为 true，则在失败时避免调用 "abort upload"，将所有已成功上载的部分保留在 S3 中供手动恢复。 (默认值: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置了该项，则不尝试检查存储桶是否存在或创建存储桶。 (默认值: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用 SSE-C，可选头部，指定加密算法为 "AES256"。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用 SSE-C，可选头部，指定用于加密或解密数据的 base64 编码的 256 位加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用 SSE-C，文件中包含与对象关联的 AES-256 加密密钥的 base64 编码字符串。[$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用 SSE-C，则是可选头部，指定加密密钥的 base64 编码的SHA256哈希值。[$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在保管库中使用自己的主密钥，则此头部指定使用的主加密密钥的 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)。[$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时要使用的存储等级。参见 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (默认值: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       分块上传的并发数。 (默认值: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的大小阈值。 (默认值: "200Mi") [$UPLOAD_CUTOFF]

   通用选项

   --name value  存储的名称 (默认值: 自动生成的名称)
   --path value  存储的路径

```
{% endcode %}