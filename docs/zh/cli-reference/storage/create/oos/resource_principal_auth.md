# 使用资源主体进行 API 调用

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos resource_principal_auth - 使用资源主体进行 API 调用

USAGE:
   singularity storage create oos resource_principal_auth [command options] [arguments...]

DESCRIPTION:
   --namespace
      对象存储的命名空间

   --compartment
      对象存储的租户 OCID

   --region
      对象存储的区域

   --endpoint
      对象存储 API 的终结点。
      
      留空以使用区域的默认终结点。

   --storage-tier
      存储新对象时要使用的存储类别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例子：
         | Standard         | 标准存储类别，这是默认类别
         | InfrequentAccess | 低频访问存储类别
         | Archive          | 存档存储类别

   --upload-cutoff
      切换到分块上传的上限。
      
      大于该大小的文件将以 chunk_size 的块上传。
      最小值为 0，最大值为 5 GiB。

   --chunk-size
      用于上传的块大小。
      
      上传超过 upload_cutoff 或大小未知的文件（例如来自 "rclone rcat" 或使用 "rclone mount" 或 Google
      照片或 Google 文档上传的文件）将使用此块大小进行多部分上传。
      
      注意，内存中会为每次传输缓冲 "upload_concurrency" 个此大小的块。
      
      如果您正在通过高速链接传输大文件，并且有足够的内存，那么增加此值将加快传输速度。
      
      Rclone 将在上传已知大小的大文件时自动增加块大小，以保持低于 10000 个块的限制。
      
      大小未知的文件将以配置的 chunk_size 进行上传。由于默认的块大小为 5 MiB，并且最多可有 10000 个块，
      这意味着默认情况下您可以按流方式上传的文件的最大大小为 48 GiB。如果您希望流式上传更大的文件，
      则需要增加 chunk_size。
      
      增加块大小会降低使用 "-P" 标志显示的进度统计数据的准确性。
      

   --upload-concurrency
      多部分上传的并发数。
      
      这是同时上传相同文件的块数。
      
      如果您正在通过高速链接上传少量大文件，并且这些文件没有充分利用您的带宽，
      那么增加此值可能有助于加快传输速度。

   --copy-cutoff
      切换到多部分拷贝的上限。
      
      需要服务器端拷贝的大于该大小的文件将按照此大小进行拷贝。
      
      最小值为 0，最大值为 5 GiB。

   --copy-timeout
      拷贝超时时间。
      
      拷贝是一个异步操作，指定超时时间以等待拷贝成功。
      

   --disable-checksum
      不要将 MD5 校验和与对象元数据一起存储。
      
      通常，在上传之前 rclone 会计算输入的 MD5 校验和，
      这样就可以将其添加到对象的元数据中。这对于数据完整性检查非常有用，
      但对于大文件来说可能会导致很长时间的延迟才能开始上传。

   --encoding
      后端的编码方式。
      
      请参阅概览中的 [编码部分](/overview/#encoding) 了解更多信息。

   --leave-parts-on-error
      如果为 true，在失败时避免调用中止上传，将所有成功上传的部分留在 S3 上供手动恢复使用。
      
      对于在不同会话之间恢复上传，应设置为 true。
      
      警告：不完整的多部分上传的部分会计入对象存储的空间使用量，并且如果没有清理将会增加额外费用。
      

   --no-check-bucket
      如果设置了此标志，就不会尝试检查存储桶是否存在或创建它。
      
      如果您已经知道存储桶已经存在，则可以使用此选项以尽量减少 rclone 的事务数。
      
      如果使用的用户没有存储桶创建权限，也可能需要此选项。

   --sse-customer-key-file
      要使用 SSE-C，需要一个包含与对象关联的 AES-256 加密密钥的 base64 编码字符串的文件。
      只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      例子：
         | <unset> | None

   --sse-customer-key
      要使用 SSE-C，需要一个可选的标题，指定要用于加密或解密数据的基64 编码 256 位加密密钥。
      只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。
      更多信息，请参阅使用自己的密钥进行端到端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      例子：
         | <unset> | None

   --sse-customer-key-sha256
      如果使用 SSE-C，则指定基于 base64 编码的加密密钥的 SHA256 哈希值的可选标题。
      该值用于检查加密密钥的完整性。请参阅使用自己的密钥进行端到端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      例子：
         | <unset> | None

   --sse-kms-key-id
      如果在保险库中使用自己的主密钥，该标题指定了调用密钥管理服务来生成数据加密密钥或加密、解密数据加密密钥时使用的主加密密钥的 OCID。
      只需要 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。

      例子：
         | <unset> | None

   --sse-customer-algorithm
      如果使用 SSE-C，则指定 "AES256" 作为加密算法的可选标题。
      对象存储支持 "AES256" 作为加密算法。请参阅
      使用自己的密钥进行端到端加密 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 了解更多信息。

      例子：
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --compartment value  对象存储的租户 OCID [$COMPARTMENT]
   --endpoint value     对象存储 API 的终结点。[$ENDPOINT]
   --help, -h           显示帮助信息
   --namespace value    对象存储的命名空间 [$NAMESPACE]
   --region value       对象存储的区域 [$REGION]

   高级选项

   --chunk-size value               用于上传的块大小，默认为 "5Mi" [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分拷贝的上限，默认为 "4.656Gi" [$COPY_CUTOFF]
   --copy-timeout value             拷贝超时时间，默认为 "1m0s" [$COPY_TIMEOUT]
   --disable-checksum               不要将 MD5 校验和与对象元数据一起存储，默认为 false [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码方式，默认为 "Slash,InvalidUtf8,Dot" [$ENCODING]
   --leave-parts-on-error           如果为 true，在失败时避免调用中止上传，将所有成功上传的部分留在 S3 上供手动恢复使用。默认为 false [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置了此标志，就不会尝试检查存储桶是否存在或创建它。默认为 false [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用 SSE-C，则指定加密算法为 "AES256" 的可选标题。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用 SSE-C，需要一个基于 base64 编码的 256 位加密密钥的可选标题。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用 SSE-C，需要一个包含与对象关联的 AES-256 加密密钥的 base64 编码字符串的文件。 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用 SSE-C，则指定基于 base64 编码的加密密钥的 SHA256 哈希值的可选标题。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在保险库中使用自己的主密钥，该标题指定了调用密钥管理服务来生成数据加密密钥或加密、解密数据加密密钥时使用的主加密密钥的 OCID。 [$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时要使用的存储类别。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm，默认为 "Standard" [$STORAGE_TIER]
   --upload-concurrency value       多部分上传的并发数，默认为 10 [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的上限，默认为 "200Mi" [$UPLOAD_CUTOFF]

   通用选项

   --name value  存储的名称（默认：自动生成）
   --path value  存储的路径

```
{% endcode %}