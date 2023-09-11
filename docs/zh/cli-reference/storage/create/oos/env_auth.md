# 自动从运行时环境中获取凭证，第一个提供认证的获胜

{% code fullWidth="true" %}
```
NAME：
   singularity storage create oos env_auth-自动从运行时环境中获取凭证，第一个提供认证的获胜

USAGE：
   singularity storage create oos env_auth [command options] [arguments...]

DESCRIPTION：
   --namespace
      对象存储的命名空间

   --compartment
      对象存储的区段OCID

   --region
      对象存储的区域

   --endpoint
      对象存储API的终端点。
      
      留空以使用该区域的默认终端点。

   --storage-tier
      存储新对象时要使用的存储级别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例子：
         |标准标准存储级别，这是默认级别
         |冷访问低频访问存储级别
         |归档归档存储级别

   --upload-cutoff
      切换为分块上传的上限。
      
      大于此大小的文件将以块大小进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的块大小。
      
      当上传大于upload_cutoff 的文件或大小未知的文件（例如，rclone rcat 或使用 rclone mount 或谷歌照片或谷歌文档上传的文件）时，
      将使用此块大小进行多部分上传。
      
      请注意，每个传输缓冲区为 "upload_concurrency" 个此大小的块。
      
      如果您正在通过高速链接传输大文件，并且有足够的内存，则增加此值将加快传输速度。
      
      当上传已知大小的大文件时，rclone 会自动增加块大小，以保持在10,000个块限制以下。
      
      大小未知的文件是使用配置的
      chunk_size 进行上传的。由于默认的块大小为 5 MiB，并且最多可以有 10,000 个块，
      这意味着默认情况下，您可以流式传输的文件的最大大小为 48 GiB。如果您希望流式上传
      更大的文件，则需要增加 chunk_size。
      
      增加块大小会降低在 "-P" 标志下显示的进度
      统计数据的准确性。

   --upload-concurrency
      多部分上传的并发性。
      
      同一文件的多个块将同时上传。
      
      如果您正在通过高速连接上传小量的大文件，并且这些文件未能充分利用带宽，
      那么增加此值可能有助于加快传输速度。

   --copy-cutoff
      切换到分块复制的上限。
      
      需要服务器端复制的大于此大小的文件将按此大小的块进行复制。
      
      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制超时。
      
      复制是异步操作，请指定超时以等待复制成功

   --disable-checksum
      不要将MD5校验和存储到对象元数据中。
      
      通常rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间延迟开始上传。

   --encoding
      后端的编码。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#Encoding)。

   --leave-parts-on-error
      如果为true，则在失败时避免调用中止上传，在S3上保留所有成功上传的部分以进行手动恢复。
      
      设置为true以便在不同的会话之间恢复上传。

      警告：在不完整的多部分上传中存储部分将计入对象存储的空间使用量，并且如果不进行清理，还将增加附加费用。


   --no-check-bucket
      如果设置，则不尝试检查桶是否存在或创建它。
      
      如果知道桶已经存在，则可以减少rclone的操作次数时，这可能是有用的。
      
      如果您使用的用户没有桶创建权限，可能也是必需的。

   --sse-customer-key-file
      要使用SSE-C，需要包含与对象关联的AES-256加密密钥的base64编码字符串的文件。
      请注意，只需要一个 sse_customer_key_file|sse_customer_key|sse_kms_key_id中的一个。

      例子：
         |<未设置> |无

   --sse-customer-key
      要使用SSE-C，此为可选标头，用于指定用于加密或解密数据的base64编码的256位加密密钥。
      请注意，只需要一个 sse_customer_key_file|sse_customer_key|sse_kms_key_id 中的一个。
      有关详细信息，请参阅 使用您自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      例子：
         |<未设置> |无

   --sse-customer-key-sha256
      如果使用SSE-C，则为可选标头，指定加密密钥的base64编码SHA256散列。
      此值用于检查加密密钥的完整性，请参阅 使用您自己的密钥进行
      服务器端加密（https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm）。

      例子：
         |<未设置> |无

   --sse-kms-key-id
      如果在保险库中使用自己的主密钥，此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的OCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）。
      请注意，只需要一个 sse_customer_key_file|sse_customer_key|sse_kms_key_id。

      例子：
         |<未设置> |无

   --sse-customer-algorithm
      如果使用SSE-C，则为可选标头，指定加密算法为"AES256"。
      对象存储支持"AES256"作为加密算法。有关详细信息，请参见
      使用您自己的密钥进行服务器端加密（https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm）。

      例子：
         |<未设置> |无
         |AES256   |AES256


OPTIONS：
   --compartment value  对象存储的区段OCID [$COMPARTMENT]
   --endpoint value     对象存储API的终端点 [$ENDPOINT]
   --help, -h           显示帮助
   --namespace value    对象存储的命名空间 [$NAMESPACE]
   --region value       对象存储的区域 [$REGION]

   高级选项

   --chunk-size value               用于上传的块大小。（默认值："5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              切换为分块复制的上限。（默认值："4.656Gi"）[$COPY_CUTOFF]
   --copy-timeout value             复制超时。（默认值："1m0s"）[$COPY_TIMEOUT]
   --disable-checksum               不要将MD5校验和存储到对象元数据中。（默认值：false）[$DISABLE_CHECKSUM]
   --encoding value                 后端的编码。（默认值："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --leave-parts-on-error           如果为true，则在失败时避免调用中止上传，将所有成功上传的部分留在S3上以进行手动恢复。（默认值：false）[$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置，则不尝试检查桶是否存在或创建它。（默认值：false）[$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用SSE-C，则为可选标头，指定加密算法为"AES256"。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用SSE-C，此为可选标头，用于指定用于加密或解密数据的base64编码的256位加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用SSE-C，需要包含与对象关联的AES-256加密密钥的base64编码字符串的文件。[$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用 SSE-C，则为可选标头，指定加密密钥的 base64 编码 SHA256 散列。[$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在保险库中使用自己的主密钥，此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的 OCID。[$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时要使用的存储级别。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm （默认值："Standard"）[$STORAGE_TIER]
   --upload-concurrency value       多部分上传的并发性。（默认值：10）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换为分块上传的上限。（默认值："200Mi"）[$UPLOAD_CUTOFF]

   常规选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}