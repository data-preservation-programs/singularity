# 使用实例主体授权实例进行API调用
每个实例都有自己的身份，并使用从实例元数据中读取的证书进行身份验证。
https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos instance_principal_auth - 使用实例主体授权实例进行API调用
                                                            每个实例都有自己的身份，并使用从实例元数据中读取的证书进行身份验证。
                                                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

USAGE:
   singularity storage update oos instance_principal_auth [command options] <name|id>

DESCRIPTION:
   --namespace
      对象存储命名空间

   --compartment
      对象存储部门OCID

   --region
      对象存储区域

   --endpoint
      对象存储API的终端点。

      留空以使用该区域的默认终端点。

   --storage-tier
      存储新对象时要使用的存储类。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      示例:
         | Standard         | 标准存储级别，这是默认级别
         | InfrequentAccess | 低频访问存储级别
         | Archive          | 存档存储级别

   --upload-cutoff
      切换到分块上传的截断点。

      大于该大小的任何文件将以chunk_size为单位上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。

      当上传大于upload_cutoff的文件或大小未知的文件（例如从"rclone rcat"或使用"rclone mount"或Google相片或Google文档上传的文件）时，将使用此分块大小进行分块上传。

      请注意，每个传输在内存中缓冲"upload_concurrency"个chunk_size大小的分块。

      如果您正在高速链路上传大型文件且有足够的内存，则增加此大小将加快传输速度。

      Rclone将在上传已知大小的大型文件时自动增加分块大小，以保持在10,000个分块限制以下。

      未知大小的文件使用配置的chunk_size进行上传。由于默认的chunk_size为5 MiB，最多可以有10,000个块，这意味着，默认情况下可以流式上传的文件的最大大小为48 GiB。如果您希望流式上传更大的文件，那么需要增加chunk_size。

      增加分块大小会降低使用"-P"标志显示的进度统计的准确性。
      

   --upload-concurrency
      分块上传的并发度。

      即同时上传相同文件的块数。

      如果您正在通过高速链路上传大量的大文件，并且这些上传未能充分利用您的带宽，那么增加这个值可能有助于加快传输速度。

   --copy-cutoff
      切换到分块复制的截断点。

      需要进行服务器端复制的大于该大小的任何文件将以此大小的块复制。

      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制超时时间。

      复制是一项异步操作，请指定超时时间以等待复制成功。

   --disable-checksum
      不在对象元数据中存储MD5校验和。

      通常，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但可能导致对于大文件的上传开始产生很长的延迟。

   --encoding
      后端的编码。

      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error
      如果设置为true，则在失败时避免调用中止上传操作，使S3上的所有已成功上传的部分均可手动恢复。

      对于恢复跨不同会话的上传，应该设置为true。

      警告：在对象存储上存储不完整的分块上传的部分会计入空间使用量，并将在清理时增加额外成本。

   --no-check-bucket
      如果设置，则不尝试检查存储桶是否存在或创建存储桶。

      如果您知道存储桶已存在，这对于试图最小化rclone执行的事务数量非常有用。

      如果您使用的用户没有存储桶创建权限，也可能需要这个选项。

   --sse-customer-key-file
      要使用SSE-C，包含与对象关联的AES-256加密密钥的base64编码字符串的文件。请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id中只需要一个选项。

      示例:
         | <unset> | None

   --sse-customer-key
      要使用SSE-C，可选的报头，指定用于加密或解密数据的base64编码256位加密密钥。

      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id中只需要一个选项。有关更多信息，请参见使用您自己的密钥进行服务器端加密部分的文档
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None

   --sse-customer-key-sha256
      如果使用SSE-C，可选的报头，指定加密密钥的base64编码SHA256哈希值。

      此值用于检查加密密钥的完整性。有关更多信息，请参见使用您自己的密钥进行服务器端加密部分的文档
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None

   --sse-kms-key-id
      如果在vault中使用您自己的主密钥，则此报头指定用于调用密钥管理服务以生成数据加密密钥或对数据加密或解密的主加密密钥的OCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）。
      sse_customer_key_file|sse_customer_key|sse_kms_key_id中只需要一个选项。

      示例:
         | <unset> | None

   --sse-customer-algorithm
      如果使用SSE-C，可选的报头，将加密算法指定为"AES256"。

      对象存储支持"AES256"作为加密算法。有关更多信息，请参见使用您自己的密钥进行服务器端加密部分的文档
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --compartment value  对象存储部门OCID [$COMPARTMENT]
   --endpoint value     对象存储API的终端点 [$ENDPOINT]
   --help, -h           显示帮助信息
   --namespace value    对象存储命名空间 [$NAMESPACE]
   --region value       对象存储区域 [$REGION]

   Advanced

   --chunk-size value               用于上传的分块大小。（默认值为"5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              切换到分块复制的截断点。（默认值为"4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             复制超时时间。（默认值为"1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               不在对象元数据中存储MD5校验和。（默认值为false） [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码。（默认值为"Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           如果设置为true，则在失败时避免调用中止上传操作，使S3上的所有已成功上传的部分均可手动恢复。（默认值为false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置，则不尝试检查存储桶是否存在或创建存储桶。（默认值为false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用SSE-C，可选的报头，将加密算法指定为"AES256"。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用SSE-C，可选的报头，指定用于加密或解密数据的base64编码256位加密密钥。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用SSE-C，包含与对象关联的AES-256加密密钥的base64编码字符串的文件。 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用SSE-C，可选的报头，指定加密密钥的base64编码SHA256哈希值。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在vault中使用您自己的主密钥，则此报头指定用于调用密钥管理服务以生成数据加密密钥或对数据加密或解密的主加密密钥的OCID。 [$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时要使用的存储类。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm （默认值为"Standard"） [$STORAGE_TIER]
   --upload-concurrency value       分块上传的并发度。（默认值为10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截断点。（默认值为"200Mi"） [$UPLOAD_CUTOFF]

```
{% endcode %}