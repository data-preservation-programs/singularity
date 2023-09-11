# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
名称:
   singularity storage update azureblob - 微软Azure Blob Storage

用法:
   singularity storage update azureblob [命令选项] <名称|ID>

描述:
   --account
      Azure存储账户名称。
      
      设置当前的Azure存储账户名称。
      
      留空以使用SAS URL或Emulator，否则需要进行设置。
      
      如果此项留空，并且env_auth已经设置，则会尝试从环境变量`AZURE_STORAGE_ACCOUNT_NAME`读取。
      

   --env-auth
      从运行时读取凭据（环境变量，CLI或MSI）。
      
      请参阅[身份验证文档](/azureblob#authentication)了解详细信息。

   --key
      存储账户的共享密钥。
      
      留空以使用SAS URL或Emulator。

   --sas-url
      仅容器级访问的SAS URL。
      
      如果使用账户/密钥或Emulator，请留空。

   --tenant
      使用的服务主体的租户ID，也称为其目录ID。
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 带有客户端密钥的服务主体
      - 带有证书的服务主体
      - 用户名和密码的用户
      

   --client-id
      正在使用的客户端ID。
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 带有客户端密钥的服务主体
      - 带有证书的服务主体
      - 用户名和密码的用户
      

   --client-secret
      服务主体的一个客户端密钥。
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 带有客户端密钥的服务主体
      

   --client-certificate-path
      包含私钥的PEM或PKCS12证书文件的路径。
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 带有证书的服务主体
      

   --client-certificate-password
      证书文件的密码（可选）。
      
      如果使用以下任一项进行身份验证，则可选地设置这个值：
      - 带有证书的服务主体
      
      并且证书是有密码的。
      

   --client-send-certificate-chain
      在使用证书进行身份验证时发送证书链。
      
      指定身份验证请求是否包括x5c标头，用于支持基于主题名称/颁发者的身份验证。设置为true时，身份验证请求会发送x5c标头。
      
      如果使用以下任一项进行身份验证，则可选地设置这个值：
      - 带有证书的服务主体
      

   --username
      用户名（通常是电子邮件地址）
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 使用用户名和密码的用户
      

   --password
      用户的密码
      
      如果使用以下任一项进行身份验证，则需要设置这个值：
      - 使用用户名和密码的用户
      

   --service-principal-file
      含有服务主体凭据的文件路径。
      
      通常留空。仅在要使用服务主体而不是交互式登录时才需要。
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      有关更多详细信息，请参阅["创建Azure服务主体"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli)和["为访问Blob数据分配Azure角色"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli)页面。
      
      直接将凭据放入rclone配置文件的`client_id`，`tenant`和`client_secret`键中，而不是设置`service_principal_file`可能更方便。
      

   --use-msi
      使用托管服务标识进行身份验证（仅适用于Azure）。
      
      当设置为true时，使用[托管服务标识](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)进行Azure存储身份验证，而不是使用SAS令牌或账户密钥。
      
      如果运行此程序的VM（SS）有系统分配的标识，则将默认使用它。如果资源没有系统分配标识，但恰好有一个用户分配的标识，则将默认使用用户分配的标识。如果资源有多个用户分配的标识，则必须明确指定要使用的标识，可以使用`msi_object_id`、`msi_client_id`或`msi_mi_res_id`参数之一。

   --msi-object-id
      要使用的用户分配MSI的对象ID（如果有）。
      
      如果指定了msi_client_id或msi_mi_res_id，则留空。

   --msi-client-id
      要使用的用户分配MSI的对象ID（如果有）。
      
      如果指定了msi_object_id或msi_mi_res_id，则留空。

   --msi-mi-res-id
      要使用的用户分配MSI的Azure资源ID（如果有）。
      
      如果指定了msi_client_id或msi_object_id，则留空。

   --use-emulator
      如果值为'true'，则使用本地存储 emulator。
      
      如果使用真实的Azure存储终结点，请留空。

   --endpoint
      服务的终结点。
      
      通常留空。

   --upload-cutoff
      切换到分块上传的截断值（<= 256 MiB）（已弃用）。

   --chunk-size
      上传分块大小。
      
      请注意，这些块存储在内存中，并且可能存在"--transfers" * "--azureblob-upload-concurrency"数量的块同时存储在内存中。

   --upload-concurrency
      并行上传分块的并发数。
      
      这是同时上传相同文件的块的数量。
      
      如果您在高速链路上上传少量大文件，并且这些上传未充分利用您的带宽，则增加此值可能有助于加快传输速度。
      
      在测试中，上传速度几乎与上传并发线性增加。例如，要填充一个千兆管道，可能需要将此值提高到64。请注意，这将占用更多内存。
      
      请注意，这些块存储在内存中，并且可能存在"--transfers" * "--azureblob-upload-concurrency"数量的块同时存储在内存中。

   --list-chunk
      blob列表的大小。
      
      这设定了每个列表块请求的blob数量。默认设置为最大值5000。允许“列表 blobs”请求使用2分钟来完成每兆字节。如果操作平均超过2分钟/MB，则会超时 ([source](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval))。这可以用来限制返回的blob项数，以避免超时。

   --access-tier
      blob的访问层级：hot、cool或archive。
      
      归档的blob可以通过将访问层级设置为hot或cool来恢复。如果打算使用默认访问层级，请留空，该访问层级设置在账户级别上。
      
      如果没有指定“访问层级”，rclone不会应用任何层级。rclone在上传时执行“Set Tier”操作，如果对象未修改，则指定新层级的"访问层级"不会产生任何影响。如果远程端的blob处于“archive tier”，则不允许从远程端执行数据传输操作。用户应该首先通过将blob分层为“Hot”或“Cool”来还原。

   --archive-tier-delete
      在覆盖之前删除存档层级的blob。
      
      由于无法更新存档层级的blob，所以如果不设置此标志，如果尝试更新存档层级的blob，则rclone将会产生以下错误：
      
          无法更新存档层级的blob而没有--azureblob-archive-tier-delete标志
      
      设置此标志后，在rclone尝试用新的文件替换存档层级的blob之前，会先删除现有的blob。这可能会导致数据丢失（与更新普通blob不同），并且可能会产生更多费用，因为提前删除存档层级的blob可能会收费。
      

   --disable-checksum
      不要将MD5校验和与对象元数据一起存储。
      
      通常情况下，rclone会在上传之前计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件来说可能会造成长时间的延迟。

   --memory-pool-flush-time
      内部内存缓冲区池刷新的频率。
      
      需要使用附加的缓冲区（例如，分块上传）时，会使用内存池进行 内存分配。
      此选项控制过多久未使用的缓冲区将从池中删除。

   --memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --encoding
      后端的编码方式。
      
      有关更多信息，请参阅[概述的编码部分](/overview/#encoding)。

   --public-access
      容器的公共访问级别：blob或container。

      例子：
         | <unset>   | 只能通过授权请求访问容器及其blob。
         |           | 这是默认值。
         | blob      | 容器内的Blob数据可以通过匿名请求读取。
         | container | 允许对容器和Blob数据进行完全的公开读取访问。

   --no-check-container
      如果设置了此选项，则不会尝试检查容器是否存在或创建容器。
      
      当尝试尽量减少rclone执行的事务数时，如果您知道容器已经存在，则这可能非常有用。
      

   --no-head-object
      如果设置了此选项，则在获取对象之前不进行HEAD。
      当获取对象时节省HEAD请求可以提高性能。

选项:
   --account value                      Azure存储账户名称 [$ACCOUNT]
   --client-certificate-password value  证书文件的密码（可选） [$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      包含私钥的PEM或PKCS12证书文件的路径 [$CLIENT_CERTIFICATE_PATH]
   --client-id value                    正在使用的客户端ID [$CLIENT_ID]
   --client-secret value                服务主体的一个客户端密钥 [$CLIENT_SECRET]
   --env-auth                           从运行时读取凭据（环境变量，CLI或MSI）（默认: false） [$ENV_AUTH]
   --help, -h                           显示帮助
   --key value                          存储账户的共享密钥 [$KEY]
   --sas-url value                      仅容器级访问的SAS URL [$SAS_URL]
   --tenant value                       使用的服务主体的租户ID [$TENANT]

   高级选项

   --access-tier value              blob的访问层级：hot、cool或archive [$ACCESS_TIER]
   --archive-tier-delete            在覆盖之前删除存档层级的blob（默认: false） [$ARCHIVE_TIER_DELETE]
   --chunk-size value               上传分块大小（默认: "4Mi"） [$CHUNK_SIZE]
   --client-send-certificate-chain  在使用证书进行身份验证时发送证书链（默认: false） [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               不要将MD5校验和与对象元数据一起存储（默认: false） [$DISABLE_CHECKSUM]
   --encoding value                 后端的编码方式（默认: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"） [$ENCODING]
   --endpoint value                 服务的终结点 [$ENDPOINT]
   --list-chunk value               blob列表的大小（默认: 5000） [$LIST_CHUNK]
   --memory-pool-flush-time value   内部内存缓冲区池刷新的频率（默认: "1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用mmap缓冲区（默认: false） [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            要使用的用户分配MSI的对象ID（如果有） [$MSI_CLIENT_ID]
   --msi-mi-res-id value            要使用的用户分配MSI的Azure资源ID（如果有） [$MSI_MI_RES_ID]
   --msi-object-id value            要使用的用户分配MSI的对象ID（如果有） [$MSI_OBJECT_ID]
   --no-check-container             如果设置了此选项，则不会尝试检查容器是否存在或创建容器（默认: false） [$NO_CHECK_CONTAINER]
   --no-head-object                 如果设置了此选项，则在获取对象时不进行HEAD操作（默认: false） [$NO_HEAD_OBJECT]
   --password value                 用户的密码 [$PASSWORD]
   --public-access value            容器的公共访问级别：blob或container [$PUBLIC_ACCESS]
   --service-principal-file value   含有服务主体凭据的文件路径 [$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       并行上传分块的并发数（默认: 16） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截断值（<= 256 MiB）（已弃用） [$UPLOAD_CUTOFF]
   --use-emulator                   如果提供'true'，则使用本地存储 emulator（默认: false） [$USE_EMULATOR]
   --use-msi                        使用托管服务标识进行身份验证（仅在Azure中有效）（默认: false） [$USE_MSI]
   --username value                 用户名（通常是电子邮件地址） [$USERNAME]

```
{% endcode %}