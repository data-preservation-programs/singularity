# 从运行时（env）自动获取凭证，第一个提供认证的获胜

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos env_auth - 从运行时（env）自动获取凭证，第一个提供认证的获胜

用法:
   singularity storage update oos env_auth [命令选项] <名称|ID>

说明:
   --namespace
      对象存储命名空间

   --compartment
      对象存储区段 OCID

   --region
      对象存储区域

   --endpoint
      对象存储 API 的终端点。
      
      留空以使用区域的默认终端点。

   --storage-tier
      存储新对象时使用的存储级别。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例:
         | Standard         | 标准存储级别，此为默认级别
         | InfrequentAccess | 延迟访问存储级别
         | Archive          | 归档存储级别

   --upload-cutoff
      切换到分块上传的切换点。
      
      大于此大小的文件将以块大小进行分块上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      对于大于 upload_cutoff 的文件，或者大小未知的文件（例如来自“rclone rcat”、使用“rclone mount”或 Google
      照片或 Google 文档上传的文件），将使用这个块大小进行分块上传。
      
      请注意，“upload_concurrency” 个此大小的块会为每个传输在内存中进行缓冲。
      
      如果您通过高速连接传输大文件，并且有足够的内存，则增大此值将加快传输速度。
      
      Rclone 将在上传已知大小的大文件时自动增加块大小，以保持在 10,000 个块的限制之下。
      
      未知大小的文件将以配置的块大小上传。由于默认块大小为 5 MiB，最多可以有 10,000 个块，这意味着默认情况下可以流式上传的文件的最大大小为 48 GiB。如果要流式上传更大的文件，则需要增加块大小。
      
      增大块大小会降低使用“-P”标志显示的进度统计的准确性。
   

   --upload-concurrency
      多部分上传的并发数。
      
      这是并发上传的同一文件的块数。
      
      如果您通过高速链接上传少量的大文件，并且这些上传无法完全利用您的带宽，则增加此值可能有助于加快传输速度。

   --copy-cutoff
      切换到多部分拷贝的切换点。
      
      大于此大小的需要服务器端复制的文件将按该大小分块复制。
      
      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制操作的超时时间。
      
      复制是一个异步操作，指定超时时间以等待复制操作成功。
      

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常 rclone 会在上传之前计算输入的 MD5 校验和，并将其添加到对象的元数据中。对于数据完整性检查来说这是很好的，但对于大文件来说会导致长时间的上传延迟。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error
      如果为 true，则在失败时避免调用中止上传，将所有成功上传的部分留在 S3 上以供手动恢复。
      
      对于恢复跨不同会话的上传，应将其设置为 true。
      
      警告：存储不完整的多部分上传的部分会计入对象存储的空间使用，并且如果不清理，将增加额外费用。
      

   --no-check-bucket
      如果设置了此选项，则不会尝试检查存储桶是否存在或创建。
      
      这在将 rclone 的事务数量最小化时很有用，如果已经知道存储桶已经存在。
      
      如果您使用的用户没有创建存储桶的权限，也可能需要此选项。
      

   --sse-customer-key-file
      要使用 SSE-C，需要包含与对象关联的 AES-256 加密密钥的 base64 编码字符串的文件。
      请注意，只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      示例:
         | <unset> | None

   --sse-customer-key
      要使用 SSE-C，指定 base64 编码的 256 位加密密钥进行加密或解密的可选标头。
      请注意，只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。有关更多信息，请参见使用自己的加密密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      示例:
         | <unset> | None

   --sse-customer-key-sha256
      如果使用 SSE-C，请指定用于加密密钥的 base64 编码 SHA256 哈希的可选标头。
      此值用于检查加密密钥的完整性。有关详细信息，请参见使用自己的加密密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None

   --sse-kms-key-id
      如果在保险库中使用自己的主密钥，则此标头指定调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥所使用的主加密密钥的 OCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）。
      请注意，只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      示例:
         | <unset> | None

   --sse-customer-algorithm
      如果使用 SSE-C，则此可选标头指定 AES256 作为加密算法。
      对象存储支持“AES256”作为加密算法。有关详细信息，请参见
      使用自己的加密密钥进行服务器端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None
         | AES256  | AES256


选项:
   --compartment value  对象存储区段 OCID [$COMPARTMENT]
   --endpoint value     对象存储 API 的终端点。[$ENDPOINT]
   --help, -h           显示帮助信息
   --namespace value    对象存储命名空间 [$NAMESPACE]
   --region value       对象存储区域 [$REGION]

   高级选项

   --chunk-size value               用于上传的块大小。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分拷贝的切换点。 (default: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             复制操作的超时时间。 (default: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储。 (default: false) [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           如果为 true，则在失败时避免调用中止上传，将所有成功上传的部分留在 S3 上以供手动恢复。 (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置了此选项，则不会尝试检查存储桶是否存在或创建。 (default: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用 SSE-C，则此可选标头指定 AES256 作为加密算法。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用 SSE-C，指定 base64 编码的 256 位加密密钥进行加密或解密的 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用 SSE-C，需要包含与对象关联的 AES-256 加密密钥的 base64 编码字符串的文件。 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用 SSE-C，请指定用于加密密钥的 base64 编码 SHA256 哈希的可选标头。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在保险库中使用自己的主密钥，则此标头指定调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥所使用的主加密密钥的 OCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）。 [$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时使用的存储级别。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       多部分上传的并发数。 (default: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的切换点。 (default: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}