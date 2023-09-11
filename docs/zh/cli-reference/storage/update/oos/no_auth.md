# 不需要凭证，这通常用于读取公共存储桶

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos no_auth - 不需要凭证，这通常用于读取公共存储桶

USAGE:
   singularity storage update oos no_auth [command options] <name|id>

DESCRIPTION:
   --namespace
      对象存储的命名空间

   --region
      对象存储的区域

   --endpoint
      对象存储 API 的终端节点。
      
      留空以使用区域的默认终端节点。

   --storage-tier
      存储新对象时要使用的存储类型。详见 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例:
         | Standard         | 标准存储类型，默认类型
         | InfrequentAccess | 低频存储类型
         | Archive          | 存档存储类型

   --upload-cutoff
      切换到分块上传的文件大小截断值。
      
      大于该值的文件将按照 chunk_size 进行分块上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      上传时要使用的块大小。
      
      上传大于 upload_cutoff 或大小未知的文件（例如使用 "rclone rcat" 上传或使用 "rclone mount" 或谷歌照片或谷歌文档上传）将使用此块大小进行多部分上传。
      
      请注意，每个传输将在内存中缓冲 upload_concurrency 个此大小的块。
      
      如果您正在通过高速链接传输大文件，并且内存足够，增加此值将加快传输速度。
      
      Rclone 将自动增加块大小，以确保在上传已知大小的大文件时不超过 10,000 个块的限制。
      
      对于未知大小的文件，将使用配置的 chunk_size 进行上传。因为默认的块大小为 5 MiB，并且最多可以有 10,000 个块，所以默认情况下您可以流式上传的文件的最大大小为 48 GiB。如果您希望流式上传更大的文件，则需要增加 chunk_size。
      
      增大块大小会降低 "-P" 标志显示的进度统计的准确性。
      

   --upload-concurrency
      多部分上传的并发数。
      
      同一文件的连续块的并发上传数。
      
      如果您正在通过高速链接上传少量大文件，并且这些上传未充分利用带宽，那么增加此值可能有助于加快传输速度。

   --copy-cutoff
      切换到多部分复制的文件大小截断值。
      
      大于该值的需要服务器端复制的文件将按照这个大小进行分块复制。
      
      最小值为 0，最大值为 5 GiB。

   --copy-timeout
      复制操作的超时时间。
      
      复制是一种异步操作，指定超时时间以等待复制成功。
      

   --disable-checksum
      不要将 MD5 校验和与对象元数据一同存储。
      
      通常在上传之前，rclone 会计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但可能导致大文件在开始上传时长时间延迟。

   --encoding
      后端的编码方式。
      
      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error
      如果为 true，则在失败时避免调用中止上传，使所有成功上传的部分保留在 S3 上，以便手动恢复。
      
      对于在不同会话之间恢复上传，应将其设置为 true。
      
      警告: 存储不完整的分块上传的部分会占用对象存储的空间用量，并且如果未清理，将增加额外的成本。
      

   --no-check-bucket
      如果设置了该标志，则不尝试检查存储桶是否存在或创建。
      
      当尝试最小化 rclone 执行的事务数量时，可以使用此选项，前提是您知道存储桶已存在。
      
      如果使用的用户没有创建存储桶的权限，也可能需要使用此选项。
      

   --sse-customer-key-file
      使用 SSE-C 时，包含与对象关联的 AES-256 加密密钥的 BASE64 编码字符串的文件。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id 只需要其中一个。

      示例:
         | <unset> | None

   --sse-customer-key
      使用 SSE-C 时，可选的标头，指定用于加密或解密数据的 BASE64 编码的 256 位加密密钥。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id 只需要其中一个。有关更多信息，请参见使用自己的密钥进行服务器端加密 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None

   --sse-customer-key-sha256
      如果使用 SSE-C，则会指定一个可选的标头，其中包含加密密钥的 BASE64 编码 SHA256 哈希。
      此值用于检查加密密钥的完整性。有关详细信息，请参见使用自己的密钥进行服务器端加密 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None

   --sse-kms-key-id
      如果在保管库中使用自己的主密钥，则此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密键的主加密密钥的 OCID。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id 只需要其中一个。

      示例:
         | <unset> | None

   --sse-customer-algorithm
      如果使用 SSE-C，则可选的标头，指定 "AES256" 作为加密算法。
      对象存储支持 "AES256" 作为加密算法。有关详细信息，请参见使用自己的密钥进行服务器端加密 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --endpoint value   对象存储 API 的终端节点。 [$ENDPOINT]
   --help, -h         显示帮助信息
   --namespace value  对象存储的命名空间 [$NAMESPACE]
   --region value     对象存储的区域 [$REGION]

   高级选项

   --chunk-size value                 上传时要使用的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value                切换到多部分复制的文件大小截断值。 (default: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value               复制操作的超时时间。 (default: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum                 不要将 MD5 校验和与对象元数据一同存储。 (default: false) [$DISABLE_CHECKSUM]
   --encoding value                   后端的编码方式。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error             如果为 true，则在失败时避免调用中止上传，使所有成功上传的部分保留在 S3 上，以便手动恢复。 (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                  如果设置了该标志，则不尝试检查存储桶是否存在或创建。 (default: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value     如果使用 SSE-C，则可选的标头，指定 "AES256" 作为加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value           使用 SSE-C 时，可选的标头，指定用于加密或解密数据的 BASE64 编码的 256 位加密密钥。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value      使用 SSE-C 时，包含与对象关联的 AES-256 加密密钥的 BASE64 编码字符串的文件。 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value    如果使用 SSE-C，则会指定一个可选的标头，其中包含加密密钥的 BASE64 编码 SHA256 哈希。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value             如果在保管库中使用自己的主密钥，则此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密键的主加密密钥的 OCID。 [$SSE_KMS_KEY_ID]
   --storage-tier value               存储新对象时要使用的存储类型。详见 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$STORAGE_TIER]
   --upload-concurrency value         多部分上传的并发数。 (default: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value              切换到分块上传的文件大小截断值。 (default: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}