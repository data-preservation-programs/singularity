# 使用OCI用户和API密钥进行身份验证。
您需要在配置文件中填入租户OCID，用户OCID，区域，路径和API密钥的指纹。
[配置SDK](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm)

{% code fullWidth="true" %}
```
命令:
   singularity storage create oos user_principal_auth - 使用OCI用户和API密钥进行身份验证。
                                                        您需要在配置文件中填入租户OCID，用户OCID，区域，路径和API密钥的指纹。
                                                        [配置SDK](https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm)

使用方式:
   singularity storage create oos user_principal_auth [command options] [arguments...]

简介:
   --namespace
      对象存储命名空间

   --compartment
      对象存储的区域OCID

   --region
      对象存储区域

   --endpoint
      对象存储API的终端点。
      
      留空可使用区域的默认终端点。

   --config-file
      OCI配置文件的路径

      示例:
         | ~/.oci/config | oci配置文件路径

   --config-profile
      oci配置文件中的配置文件名称

      示例:
         | Default | 使用默认配置文件

   --storage-tier
      存储新对象时要使用的存储类。 [了解更多](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)

      示例:
         | Standard         | 标准存储类，这是默认存储类
         | InfrequentAccess | 低频访问存储类
         | Archive          | 存档存储类

   --upload-cutoff
      切换到分块上传的分割线。
      
      大于此值的文件将以分块尺寸上传。
      最小值为0，最大值为5 GiB。

   --chunk-size
      用于上传的分块大小。
      
      上传大于upload_cutoff或大小未知的文件（例如使用“rclone rcat”或“rclone mount”或Google
      照片或Google文档上传的文件），将使用此分块大小进行分块上传。
      
      请注意，每个传输器将在内存中缓冲upload_concurrency个这样大小的分块。
      
      如果您通过高速链路传输大型文件并且有足够的内存，则可以增加此值以加快传输速度。
      
      当上传大小已知的大文件以保持在10000个分块一下时，rclone会自动增加分块大小。

      大小未知的文件使用配置的chunk_size上传。默认的chunk_size为5 MiB，最多可有
      10000个分块，这意味着默认情况下，您可以在流式上传的文件大小最大为48 GiB。
      如果要流式上传更大的文件，您需要增加chunk_size。
      
      增加分块大小会降低使用“-P”标志时显示的传输进度的准确性。

   --upload-concurrency
      多段上传的并发数。
      
      同一文件的这些分块并发上传。
      
      如果在高速链路上上传少量大文件，而这些上传未充分利用带宽，则可以考虑
      增加此值以加快传输速度。

   --copy-cutoff（复制分割线）
      切换到多段复制的分割线。
      
      大于此值的需要在服务器端复制的文件将以此大小的分块复制。
      
      最小为0，最大为5 GiB。

   --copy-timeout（复制超时）
      复制的超时时间。
      
      复制是一个异步操作，指定超时时间以等待复制成功

   --disable-checksum（禁用校验和）
      不将MD5校验和与对象元数据一起存储。
      
      通常，rclone会在上传前计算输入的MD5校验和，以便将其添加到对象的元数据中。
      这对于数据完整性检查很有用，但是可能会导致大型文件长时间启动上传。

   --encoding（编码）
      后端的编码。
      
      有关详情，请参阅[概述中的编码部分](/overview/#encoding)。

   --leave-parts-on-error（错误时保留分块）
      如果为true，在失败时避免调用中止上传，使所有成功上传的分块留在S3上供手动恢复使用。
      
      对于在不同会话之间恢复上传，应将其设置为true。
      
      警告: 存储不完整的多段上传的部分会导致对象存储空间使用量增加，并在未清理时产生
      额外费用。

   --no-check-bucket（无需检查存储桶）
      如果设置了此选项，则不会检查存储桶是否存在或创建它。
      
      这在试图最小化rclone的事务数时可能很有用，如果您已经知道存储桶已经存在。

      如果使用的用户没有创建存储桶的权限，则可能需要此选项。

   --sse-customer-key-file（SSE-C客户端密钥文件）
      要使用SSE-C，指定与对象关联的AES-256加密密钥的Base64编码字符串的文件。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id只需要其中一个。

      示例:
         | <unset> | 无

   --sse-customer-key（SSE-C客户端密钥）
      要使用SSE-C，可选的标头，指定用于加密或解密数据的Base64编码的256位加密密钥。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id只需要其中一个。
      有关更多信息，请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | 无

   --sse-customer-key-sha256（SSE-C客户端密钥SHA256）
      如果使用SSE-C，可选的标头，指定加密密钥的Base64编码的SHA256哈希。
      此值用于检查加密密钥的完整性。
      有关更多信息，请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | 无

   --sse-kms-key-id（SSE-KMS密钥ID）
      如果使用自己的主密钥在保险库中，此标头指定用于调用
      密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的OCID。
      请注意，sse_customer_key_file|sse_customer_key|sse_kms_key_id只需要其中一个。

      示例:
         | <unset> | 无

   --sse-customer-algorithm（SSE-C算法）
      如果使用SSE-C，可选的标头，指定加密算法为“AES256”。
      对象存储支持“AES256”作为加密算法。
      有关更多信息，请参见使用自己的密钥进行服务器端加密
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

      示例:
         | <unset> | 无
         | AES256  | AES256


选项:
   --compartment value     对象存储的区域OCID [$COMPARTMENT]
   --config-file value     OCI配置文件的路径（默认值: "~/.oci/config"） [$CONFIG_FILE]
   --config-profile value  oci配置文件中的配置文件名称（默认值: "Default"） [$CONFIG_PROFILE]
   --endpoint value        对象存储API的终端点 [$ENDPOINT]
   --help, -h              显示帮助
   --namespace value       对象存储命名空间 [$NAMESPACE]
   --region value          对象存储区域 [$REGION]

   高级

   --chunk-size value               用于上传的分块大小（默认值: "5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              切换到多段复制的分割线（默认值: "4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             复制的超时时间（默认值: "1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               不将MD5校验和与对象元数据一起存储（默认值: false） [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码（默认值: "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           如果为true，在失败时避免调用中止上传，使所有成功上传的分块留在S3上供手动恢复使用（默认值: false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                如果设置了此选项，则不会检查存储桶是否存在或创建它（默认值: false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   如果使用SSE-C，可选的标头，指定加密算法为“AES256” [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         要使用SSE-C，可选的标头，指定用于加密或解密数据的Base64编码的256位加密密钥 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    要使用SSE-C，将包含与对象关联的AES-256加密密钥的Base64编码字符串的文件 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  如果使用SSE-C，可选的标头，指定加密密钥的Base64编码的SHA256哈希 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           如果使用自己的主密钥在保险库中，此标头指定用于调用 [$SSE_KMS_KEY_ID]
   --storage-tier value             存储新对象时要使用的存储类 [了解更多]https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (默认值: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       多段上传的并发数（默认值: 10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的分割线（默认值: "200Mi"） [$UPLOAD_CUTOFF]

   常规

   --name value  存储的名称（默认值: 自动生成）
   --path value  存储的路径

```
{% endcode %}