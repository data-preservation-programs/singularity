# Oracle Cloud Infrastructure对象存储

{% code fullWidth="true" %}
```
名称:
   singularity datasource add oos - Oracle Cloud Infrastructure对象存储

用法:
   singularity datasource add oos [命令选项] <数据集名称> <源路径>

描述:
   --oos-sse-customer-key-file
      要使用SSE-C，需要一个包含与对象关联的256位AES-256加密密钥的base64编码字符串的文件。
      请注意，只需要sse_customer_key_file | sse_customer_key | sse_kms_key_id中的一个。

      示例：
         |<unset> | 无

   --oos-sse-kms-key-id
      如果在库中使用您自己的主密钥，请使用此标头指定用于调用密钥管理服务以生成数据加密密钥或加密或解密数据加密密钥的主加密密钥的OCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）。
      请注意，只需要sse_customer_key_file | sse_customer_key | sse_kms_key_id中的一个。

      示例：
         |<unset> | 无

   --oos-config-file
      [Provider] - user_principal_auth
         OCI配置文件路径

         示例：
            | ~/.oci/config | oci配置文件位置

   --oos-upload-cutoff
      用于切换到分块上传的截止值。

      大于此大小的所有文件将以chunk_size的块上传。
      最小值为0，最大值为5 GiB。

   --oos-leave-parts-on-error
      如果为真，则在出错时避免调用中止上传，将所有成功上传的部分留在S3上以供手动恢复。

      对于在不同会话之间恢复上传应设置为true。

      警告：未完成的分块上传的部分计入对象存储的空间使用量，如果不清理，则会增加其他费用。
      

   --oos-chunk-size
      用于上传的块大小。

      当上传大于upload_cutoff的文件或文件大小未知（例如，从“rclone rcat”上传或使用“rclone mount”或google photos或google docs上传）时，它们将使用此块大小作为分块上传。

      请注意，“upload_concurrency”个此大小的块按传输分类在内存中缓冲。

      如果您正在通过高速链接传输大文件并且具有足够的内存，则增加此值将加快传输速度。

      上传已知大小的大文件时，Rclone会自动增加块大小以保持低于1万个块的限制。

      以未知大小的文件以配置的块大小上传。由于默认块大小为5 MiB，最多可以有10000个块，因此默认情况下可以流式上传的文件的最大大小为48 GiB。如果要流式上传更大的文件，则需要增加块大小。

      增加块大小会降低使用“-P”标志显示的进度统计数据的准确性。
     

   --oos-copy-timeout
      复制超时。

      复制为异步操作，请指定超时以等待复制成功

   --oos-sse-customer-key
      要使用SSE-C，指定用于加密或解密数据的base64编码的256位加密密钥的可选标头。请注意，仅需要sse_customer_key_file | sse_customer_key | sse_kms_key_id之一。有关详细信息，请参见使用自己的服务器端加密密钥

      示例：
         |<unset> | 无

   --oos-provider
      选择身份验证提供程序

      示例：
         | env_auth                | 自动从运行时（env）中提取凭据，第一个提供授权的凭据获胜
         | user_principal_auth     | 使用OCI用户和API密钥进行身份验证。
                                   | 您需要在配置文件中放置租户OCID、用户OCID、区域、路径、API密钥的指纹。
                                   | https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
         | instance_principal_auth | 使用实例主体来授权实例进行API调用。
                                   | 每个实例都拥有自己的身份标识，并使用从实例元数据中读取的证书进行身份验证。
                                   | https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
         | resource_principal_auth | 使用资源主体进行API调用
         | no_auth                 | 不需要凭据，这通常用于读取公共存储桶

   --oos-compartment
      [Provider] - user_principal_auth
         对象存储所属的部门OCID

   --oos-endpoint
      对象存储API的端点。
      
      保留为空以使用区域的默认端点。

   --oos-region
      对象存储区域

   --oos-config-profile
      [Provider] - user_principal_auth
         oci配置文件中的程序文件名称

         示例：
            | Default | 使用默认配置

   --oos-sse-customer-key-sha256
      如果使用SSE-C，则为可选标头，该标头指定加密密钥的base64编码SHA256哈希。此值用于检查加密密钥的完整性。有关详细信息，请参见使用自己的密钥进行服务器端加密。

      示例：
         |<unset> | 无

   --oos-copy-cutoff
      用于切换到多部分复制的截止