# 使用实例主体授权实例进行API调用
每个实例都有自己的身份，并使用从实例元数据中读取的证书进行身份验证。
https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos instance_principal_auth - 使用实例主体授权实例进行API调用
                                                            每个实例都有自己的身份，并使用从实例元数据中读取的证书进行身份验证。
                                                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

USAGE:
   singularity storage create oos instance_principal_auth [command options] [arguments...]

DESCRIPTION:
   --namespace
      对象存储命名空间

   --compartment
      对象存储部门OCID

   --region
      对象存储区域

   --endpoint
      对象存储API的终端。

      留空以使用该区域的默认终端。

   --storage-tier
      存储类别，在存储新对象时使用。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例子:
         | 标准            | 标准存储类别，这是默认类别
         | 低频访问        | 低频访问存储类别
         | 存档            | 存档存储类别

   --upload-cutoff
      切换到分块上传的截止值。

      大于此大小的任何文件将以块大小进行上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      上传时使用的块大小。

      当文件大于upload_cutoff或大小未知的文件（例如来自"rclone rcat"或使用"rclone mount"或Google照片或Google文档上传的文件）时，
      将使用此块大小进行分块上传。

      请注意，每次传输内存中都会缓冲"upload_concurrency"大小的块。

      如果您正在通过高速链接传输大型文件，并且有足够的内存，那么增加此值将加速传输。

      Rclone将自动增加块大小以在10,000块限制以下上传已知大小的大文件。

      未知大小的文件将使用配置的块大小进行上传。由于默认块大小为5 MiB，最多可有10,000块，这意味着默认情况下可以以流式上传的最大文件大小为48 GiB。 
      如果您希望流式上传更大的文件，则需要增加chunk_size。

      增加块大小会降低使用“-P”标志显示的进度统计的准确性。


   --upload-concurrency
      多部分上传的并发数。

      这是同时上传同一文件的块数量。

      如果您正在高速链接上上传数量较少的大文件，并且这些上传不完全利用带宽，则增加此值可能有助于加速传输。

   --copy-cutoff
      切换到多部分复制的截止值。

      大于此大小并且需要在服务器端复制的任何文件将按照此大小分块复制。

      最小值为0，最大值为5 GiB。

   --copy-timeout
      复制的超时时间。

      复制是一个异步操作，指定超时时间以等待复制成功。

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。

      通常，rclone在上传之前会计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常好，但可能会导致大文件开始上传的时间延长。

   --encoding
      后端的编码方式。

      有关详细信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error
      如果为true，则在失败时避免调用中止上传，将所有成功上传的部分留在S3上供手动恢复。

      对于不同会话之间的恢复上传，应将其设置为true。

      警告：在不完整的多部分上传的部分占用对象存储空间，并且如果不及时清理，将增加额外的费用。

   --no-check-bucket
      如果设置，则不尝试检查桶是否存在或创建它。

      如果您知道桶已存在，则可以减少rclone的事务数。这对于尝试最小化rclone的事务数很有用。

      如果您使用的用户没有桶创建权限，也可能需要这样做。

   --sse-customer-key-file
      若要使用SSE-C，请使用包含与对象关联的AES-256加密密钥的base64编码字符串的文件。
      请注意，sse_customer_key_file | sse_customer_key | sse_kms_key_id中只需要其中之一。

      例如:
         | <unset> | None

   --sse-customer-key
      若要使用SSE-C，则是可选头，指定要用于加密或解密数据的base64编码的256位加密密钥。
      请注意，sse_customer_key_file | sse_customer_key | sse_kms_key_id中只需要其中之一。
      有关更多信息，请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      例如:
         | <unset> | None

   --sse-customer-key-sha256
      如果使用SSE-C，则是可选头，指定加密密钥的base64编码SHA256哈希值。
      此值用于检查加密密钥的完整性。
      有关使用自己的密钥进行服务器端加密的详细信息，请参见
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      例如:
         | <unset> | None

   --sse-kms-key-id
      如果在保险库中使用自己的主密钥，则此头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的OCID。
      sse_customer_key_file | sse_customer_key | sse_kms_key_id中只需要其中之一。

      例如:
         | <unset> | None

   --sse-customer-algorithm
      如果使用SSE-C，则是可选头，指定加密算法为"AES256"。
      对象存储支持"AES256"作为加密算法。有关更多信息，请参阅
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      例如:
         | <unset> | None
         | AES256  | AES256

OPTIONS:
   --compartment value  对象存储部门OCID [$COMPARTMENT]
   --endpoint value     对象存储API的终端 [$ENDPOINT]
   --help, -h           显示帮助信息
   --namespace value    对象存储命名空间 [$NAMESPACE]
   --region value       对象存储区域 [$REGION]

   高级选项

   --chunk-size value               上传时使用的块大小（默认值："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              切换到多部分复制的截止值（默认值："4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             复制的超时时间（默认值："1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储（默认值：false） [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码方式（默认值："Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           如果为true，则在失败时避免调用中止上传，将所有成功上传的部分留在S3上供手动恢复（默认值：false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置，则不尝试检查桶是否存在或创建它（默认值：false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用SSE-C，则是可选头，指定加密算法为"AES256"。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         若要使用SSE-C，则是可选头，指定要用于加密或解密数据的base64编码的256位加密密钥。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    若要使用SSE-C，请使用包含与对象关联的AES-256加密密钥的base64编码字符串的文件。[$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用SSE-C，则是可选头，指定加密密钥的base64编码SHA256哈希值。[$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果在保险库中使用自己的主密钥，则此头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的OCID。[$SSE_KMS_KEY_ID]
   --storage-tier value             存储类别，在存储新对象时使用。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（默认值："Standard"） [$STORAGE_TIER]
   --upload-concurrency value       多部分上传的并发数（默认值：10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止值（默认值："200Mi"） [$UPLOAD_CUTOFF]

   通用选项

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径

```
{% endcode %}