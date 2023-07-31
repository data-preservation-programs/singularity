# Azure Blob存储

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add azureblob - Microsoft Azure Blob Storage

USAGE:
   singularity datasource add azureblob [命令选项] <dataset_name> <source_path>

DESCRIPTION:
   --azureblob-access-tier
      Blob的访问层级: 热、高频或存档。
      
      存档的Blob可以通过将访问层级设置为热或高频来还原。
      如果未指定"访问层级"，rclone不会应用任何层级。
      rclone在上传时对Blob执行"Set Tier"操作，如果对象未被修改，将"访问层级"指定为新的层级不会产生任何影响。
      如果Blob在远程处于"存档层级"，则禁止从远程执行数据传输操作。用户应先通过将Blob进行分层以将其转换到"热"或"高频"状态来进行还原。

   --azureblob-account
      Azure存储帐户名称。
      
      将其设置为正在使用的Azure存储帐户名称。
      
      如果留空并且设置了env_auth，则将尝试从环境变量`AZURE_STORAGE_ACCOUNT_NAME`中读取。

   --azureblob-archive-tier-delete
      在覆盖之前删除存档层级的Blob。
      
      无法更新存档层级的Blob。因此，如果尝试更新存档层级的Blob，则rclone将生成以下错误：
      
          can't update archive tier blob without --azureblob-archive-tier-delete
      
      在设置了此标志后，在rclone尝试覆盖存档层级的Blob之前，它将删除现有的Blob，然后上传其替换项。如果上传失败（不同于更新普通Blob），这可能导致数据丢失，并且可能会产生更多费用，因为提前删除存档层级的Blob可能是可计费的。

   --azureblob-chunk-size
      上传块大小。
      
      请注意，这将存储在内存中，并且可能会同时存储在内存中的块数量为“--transfers” * “--azureblob-upload-concurrency”。

   --azureblob-client-certificate-password
      证书文件的密码（可选）。
      
      如果使用
      - 带密码的服务主体
      
      并且证书有密码，则可选择设置此项。

   --azureblob-client-certificate-path
      包含私钥的PEM或PKCS12证书文件的路径。
      
      如果使用
      - 带证书的服务主体

   --azureblob-client-id
      正在使用的客户端的ID。
      
      如果使用
      - 带客户端密码的服务主体
      - 带证书的服务主体
      - 使用用户名和密码的用户

   --azureblob-client-secret
      服务主体的某个客户端密码
      
      如果使用
      - 带客户端密码的服务主体

   --azureblob-client-send-certificate-chain
      在使用证书身份验证时发送证书链。
      
      指定身份验证请求是否包含x5c标题以支持基于主题名称/颁发者的身份验证。当设置为真时，身份验证请求将包含x5c标题。
      
      如果使用
      - 带证书的服务主体

   --azureblob-disable-checksum
      不将MD5校验和与对象元数据一起存储。
      
      通常，rclone在上传之前会计算输入的MD5校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但对于大文件来说可能会导致长时间的上传延迟。

   --azureblob-encoding
      后端的编码方式。
      
      有关更多信息，请参见[概述中的编码部分](/overview/#encoding)。

   --azureblob-endpoint
      服务的终结点。
      
      通常留空。

   --azureblob-env-auth
      从运行时读取凭据（环境变量，CLI或MSI）。
      
      有关完整信息，请参见[认证文档](/azureblob#authentication)。

   --azureblob-key
      存储帐户共享密钥。
      
      如果留空并且设置了SAS URL或Emulator，则使用它们。

   --azureblob-list-chunk
      Blob列表的大小。
      
      这设置了每个列表块中请求的Blob数量。默认值为最大值5000。允许“列出Blob”请求每兆字节2分钟完成。如果平均每兆字节的操作时间超过2分钟，则会超时。
      可以使用此选项限制要返回的Blob项数，以避免超时。

   --azureblob-memory-pool-flush-time
      内部内存缓冲区池将被刷新的频率。
      
      需要额外缓冲区（如多部分）的上传将使用内存池进行分配。
      此选项控制未使用的缓冲区将从池中删除的频率。

   --azureblob-memory-pool-use-mmap
      是否在内部内存池中使用mmap缓冲区。

   --azureblob-msi-client-id
      要使用的用户分配的MSI的对象ID（如果有）。
      
      如果未指定msi_object_id或msi_mi_res_id，则留空。

   --azureblob-msi-mi-res-id
      要使用的用户分配的MSI的Azure资源ID（如果有）。
      
      如果未指定msi_client_id或msi_object_id，则留空。

   --azureblob-msi-object-id
      要使用的用户分配的MSI的对象ID（如果有）。
      
      如果未指定msi_client_id或msi_mi_res_id，则留空。

   --azureblob-no-check-container
      如果设置，不会尝试检查容器是否存在或创建它。
      
      当尝试最小化rclone执行的事务数量时，此选项可能非常有用，如果您知道容器已经存在。

   --azureblob-no-head-object
      如果设置，获取对象时不执行HEAD请求。

   --azureblob-password
      用户的密码
      
      如果使用
      - 使用用户名和密码的用户

   --azureblob-public-access
      容器的公共访问级别：Blob或容器。

      示例:
         | <未设置>   | 只有经过授权的请求才能访问容器及其Blob。
                      | 这是默认值。
         | blob      | 可以通过匿名请求读取此容器中的Blob数据。
         | container | 允许容器和Blob数据完全公开读取。

   --azureblob-sas-url
      仅供容器级别访问的SAS URL。
      
      如果使用账户或密钥，或者使用Emulator，则留空。

   --azureblob-service-principal-file
      包含用于服务主体身份验证的凭据的文件的路径。
      
      通常留空。仅当您想要使用服务主体而不是交互式登录时才需要。
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      有关更多详细信息，请参见["创建Azure服务主体"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli)和["为访问Blob数据分配Azure角色"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli)页面。
      
      直接将凭据放入rclone配置文件中的`client_id`、`tenant`和`client_secret`键中，而不是设置`service_principal_file`可能更方便。

   --azureblob-tenant
      服务主体所在的租户的ID，也称为其目录ID。
      
      如果使用
      - 带客户端密码的服务主体
      - 带证书的服务主体
      - 使用用户名和密码的用户

   --azureblob-upload-concurrency
      多部分上传的并发数。
      
      这是同时上传文件的相同块数。
      
      如果您通过高速连接上传大量大文件，并且这些上传未充分利用带宽，则增加此值可能有助于加快传输速度。
      
      测试中，上传速度几乎与上传并发数呈线性增长。例如，要填充一个千兆管道，可能需要将此值增加到64。注意，这将使用更多的内存。
      
      请注意，块存储在内存中，并且可能会同时存储在内存中的块的数量为“--transfers” * “--azureblob-upload-concurrency”。

   --azureblob-upload-cutoff
      切换到分块上传的截断点（<= 256 MiB）（已弃用）。

   --azureblob-use-emulator
      如果提供了“true”，则使用本地存储模拟器。
      
      如果使用真实的Azure存储终结点，则留空。

   --azureblob-use-msi
      使用托管服务标识进行身份验证（仅适用于Azure）。
      
      当为true时，使用[托管服务标识](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)来对Azure存储进行身份验证，而不是使用SAS令牌或帐户密钥。
      
      如果运行此程序的VM（SS）具有系统分配的标识，将使用它作为默认值。如果资源没有系统分配的标识，但刚好有一个用户分配的标识，则将使用该用户分配的标识作为默认值。如果资源具有多个用户分配的标识，则必须明确指定要使用的标识，只能使用msi_object_id，msi_client_id或msi_mi_res_id参数之一。

   --azureblob-username
      用户名（通常是电子邮件地址）
      
      如果使用
      - 使用用户名和密码的用户

OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险]在将数据集导出为CAR文件后删除数据集的文件。  (默认值: false)
   --rescan-interval value  当上次成功扫描后经过此间隔时，自动重新扫描源目录 (默认值: 禁用)
   --scanning-state value   设置初始扫描状态 (默认值: ready)

   用于azureblob的选项

   --azureblob-access-tier value                    Blob的访问层级：热、高频或存档。 [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure存储帐户名称。 [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            在覆盖之前删除存档层级的Blob。 (默认值: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     上传块大小。 (默认值: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    证书文件的密码（可选）。 [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        包含私钥的PEM或PKCS12证书文件的路径。 [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      正在使用的客户端的ID。 [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  服务主体的某个客户端密码 [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  在使用证书身份验证时发送证书链。 (默认值: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               不将MD5校验和与对象元数据一起存储。 (默认值: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       后端的编码方式。 (默认值: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       服务的终结点。 [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       从运行时读取凭据（环境变量、CLI或MSI）。 (默认值: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            存储帐户共享密钥。 [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     Blob列表的大小。 (默认值: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         内部内存缓冲区池将被刷新的频率。 (默认值: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           是否在内部内存池中使用mmap缓冲区。 (默认值: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  要使用的用户分配的MSI的对象ID（如果有）。 [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  要使用的用户分配的MSI的Azure资源ID（如果有）。 [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  要使用的用户分配的MSI的对象ID（如果有）。 [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             如果设置，不会尝试检查容器是否存在或创建它。 (默认值: "false") [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 如果设置，获取对象时不执行HEAD请求。 (默认值: "false") [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       用户的密码 [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  容器的公共访问级别：Blob或容器。 [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        仅供容器级别访问的SAS URL。 [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         包含用于服务主体身份验证的凭据的文件的路径。 [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         服务主体所在的租户的ID，也称为其目录ID。 [$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             多部分上传的并发数。 (默认值: "16") [$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  切换到分块上传的截断点（<= 256 MiB）（已弃用）。 [$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulator value                   如果提供了“true”，则使用本地存储模拟器。 (默认值: "false") [$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        使用托管服务标识进行身份验证（仅适用于Azure）。 (默认值: "false") [$AZUREBLOB_USE_MSI]
   --azureblob-username value                       用户名（通常是电子邮件地址） [$AZUREBLOB_USERNAME]

```
{% endcode %}