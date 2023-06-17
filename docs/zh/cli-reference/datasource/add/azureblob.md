# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
名称：
    singularity 数据源：加入 AzureBlob - Microsoft Azure Blob 存储

使用方法：
    singularity datasource add azureblob [命令选项] <数据集名称> <源路径>

描述：
    --azureblob-env -auth 
        从运行时（环境变量，CLI 或 MSI）中读取凭据。 
        
        有关完整信息，请参见 [认证文档](/azureblob#authentication)。

    --azureblob-upload-cutoff
        切换到分块上传的截止点 (<= 256 MiB)（已弃用）。

   --azureblob-encoding
        后端的编码方式。
        
        有关更多信息，请参见[总览中的编码部分](/overview/#encoding)。

    --azureblob-archive-tier-delete 
        在覆盖之前删除归档层 Blob。
        
        归档层的 Blob 无法进行更新。因此，如果您尝试更新归档层的 Blob，则 rclone 将产生错误：
        
            不能在没有 --azureblob-archive-tier-delete 的情况下更新归档层 Blob。
        
        如果设置了此标志，则在 rclone 尝试覆盖归档层 Blob 之前，它将在上传其替换项之前删除现有 Blob。 这存在数据丢失的潜在风险（与更新普通 Blob 不同），而且可能会导致更多的成本，因为删除归档层 Blob 可能要收费。

    --azureblob-public-access
        容器的公开访问级别：blob 或 container。

        示例：
             | <未设置>   | 只有经过授权的请求才能访问容器及其 Blob。
                         | 这是默认值。
             | blob      | 此容器中的 Blob 数据可以通过匿名请求读取。
             | container | 允许容器和 Blob 数据完全公开阅读访问。

    --azureblob-client-certificate-password 
        证书文件的密码（可选）。

        如果使用以下项：
        - 带证书的服务主体

        并且该证书有密码，则可以选择设置此项。

    --azureblob-username 
        用户名（通常是电子邮件地址）

        如果使用以下项之一，则设置此项：
        - 带用户名和密码的用户

    --azureblob-chunk-size 
        上载块大小。

        请注意，此项存储在内存中，并且在内存中可能最多存储
        "--transfers" * "--azureblob-upload-concurrency" 个块。

    --azureblob-upload-concurrency 
        多部分上传的并发数。

        这是同时上传文件的相同部分块数。

        如果您在高速链接上上传少量的大文件，而且这些文件未能充分利用您的带宽，则增加此值可能有助于加速传输。

        在测试中，上传速度基本上随着上传并发的增加而线性增加。例如，为了填满千兆管道，可能需要将其提高到 64. 请注意，这将使用更多的内存。

        请注意，块存储在内存中，最多可能在内存中同时存储
        "--transfers" * "--azureblob-upload-concurrency" 个块。

    --azureblob-client-certificate-path
        PEM 或 PKCS12 证书文件（包括私钥）的路径。

        如果使用以下项之一，则设置此项：
        - 带证书的服务主体

    --azureblob-memory-pool-use-mmap 
        是否在内部内存池中使用 mmap 缓冲区。

    --azureblob-no-check-container 
        如果设置，则不尝试检查容器是否存在或创建它。

        当您尝试将 rclone 执行的事务数最小化时，此项可能会很有用，如果您知道容器已经存在，则可以使用此项。

    --azureblob-no-head-object 
        如果设置，则在获取对象时不执行 HEAD 操作。

    --azureblob-memory-pool-flush-time 
        内部存储器缓冲池刷新频率。

        需要使用上传的额外缓冲区（例如多部分）时，将使用内存池进行分配。此选项控制将不使用的缓冲区删除时的频率。

    --azureblob-tenant 
        服务主体的租户 ID。 也称为其目录 ID。

        如果使用以下项之一，则设置此项：
        - 带客户端密钥的服务主体
        - 带证书的服务主体
        - 带用户名和密码的用户

    --azureblob-client-secret 
        服务主体的一个客户机密钥。

        如果使用以下项之一，则设置此项：
        - 带客户端密钥的服务主体

    --azureblob-password 
        用户的密码。

        如果使用以下项，则设置此项：
        - 带用户名和密码的用户

    --azureblob-disable-checksum 
        不要在对象元数据中存储 MD5 校验和。

        通常，rclone 将在上传之前计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查非常有用，但是对于大文件的开始上传可能会导致长时间的延迟。

    --azureblob-list-chunk 
        Blob 列表的大小。

        这将设置每个列表块中请求的 Blob 数。默认值为最大值 5000。每分钟允许“列举 Blob”请求花费 2 分钟才能完成。如果平均操作时间超过每兆字节 2 分钟，则会超时
        ([source](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)).这可用于限制要返回的 Blob 项数，以避免超时。

    --azureblob-key 
        存储帐户共享密钥。

        留空以使用 SAS URL 或仿真器。

    --azureblob-sas-url 
        仅限容器级别访问的 SAS URL。

        如果使用帐户/密钥或仿真器，请留空。

    --azureblob-client-id 
        正在使用的客户端 ID。

        如果使用以下项之一，则设置此项：
        - 带客户端密钥的服务主体
        - 带证书的服务主体
        - 带用户名和密码的用户

    --azureblob-client-send-certificate-chain 
        在使用证书身份验证时发送证书链。

        指定身份验证请求是否将包括 x5c 头，以支持基于主体名称/颁发者的身份验证。当设置为 true 时，身份验证请求将包括 x5c 头。

        如果使用以下项之一，则可以选择设置此项：
        - 带证书的服务主体

    --azureblob-service-principal-file 
        包含与服务主体一起使用的凭据的文件的路径。

        通常情况下，不需要填写。仅当您想使用服务主体而不是交互式登录时才需要。

        $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
        
        以下页面提供了更多详细信息：["创建 Azure 服务主体"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-
```
azureblob命令选项：

   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作]将数据集导出为CAR文件后删除文件。  (默认值：false)
   --rescan-interval value  当上一次成功扫描的时间已经超过此间隔时，自动重新扫描源目录（默认值：禁用）

   azureblob选项

   --azureblob-access-tier value                    Blob的存储层级：hot, cool或archive。 [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure存储账号名称。 [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            在覆盖前删除归档层级blob。（默认值："false"）[$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     上传块大小。 (默认值："4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    证书文件的密码 （可选）。[$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        PEM或PKCS12证书文件的路径，包括私钥。[$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      使用的客户端ID。[$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  服务主体的一个客户端 secret[$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  当使用证书进行身份验证时，发送证书链。(默认值："false")[$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               不将MD5校验和与对象元数据一起存储。(默认值："false")[$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       后端编码方式。(默认值："Slash，BackSlash，Del，Ctl，RightPeriod，InvalidUtf8")[$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       服务的终结点。[$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       从运行时读取凭据（环境变量，CLI或MSI）。 (默认值："false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            存储账号的共享密钥。 [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     Blob列表的大小。 (默认值："5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         内部存储器缓存池刷新的时间间隔。 (默认值："1m0s")[$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           内部存储器池是否使用mmap缓存。(默认值："false")[$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  用户分配的MSI的Object ID。(如果有的话) [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  用户分配的MSI的Azure资源ID。(如果有的话) [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  用户分配的MSI的Object ID。（如果有的话）[$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             如果设置，则不尝试检查容器是否存在或创建它。（默认值："false"） [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 如果设置，则在获取对象时不执行HEAD操作。（默认值:"false" ） [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       用户的密码[$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  容器的公共访问级别：blob或container。[$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        SAS URL仅适用于容器级别的访问。[$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         包含与服务主体一起使用的凭据的文件的路径。 [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         服务主体租户的ID。也称为其目录ID。[$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             多部分上传并发性。 (默认值："16")[$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  切换到分块上传的截止值(<=256 MiB)(已弃用)。[$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulator value                   如果提供'true'，则使用本地存储模拟器。(默认值："false")[$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        使用托管服务标识进行身份验证（仅在Azure中有效）。(默认值："false") ($AZUREBLOB_USE_MSI)
   --azureblob-username value                       用户名称(通常是电子邮件地址)[$AZUREBLOB_USERNAME]
   
   --azureblob-access-tier
      Blob的存储层级：hot, cool或archive。
      
      将Blob存储级别设置为热、冷或归档。如果想使用默认存储级别，即在账户级别设置的级别，可以不设置此参数。
      
      如果没有设置“存储层级”，rclone不会应用任何级别。在上传文件时，rclone会执行“Set Tier”操作，如果对象没有被修改，则指定新的“存储层级”不会产生任何影响。如果远程仓库的Blob处于“归档”层级，则不允许尝试从远程执行数据传输操作。用户应先通过调整Blob的存储层级将其恢复为“热”或“冷”，然后再执行数据传输操作。

```
