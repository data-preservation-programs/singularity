# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```markdown
名称:
   singularity storage create azureblob - Microsoft Azure Blob Storage

用法:
   singularity storage create azureblob [command options] [arguments...]

描述:
   --account
      Azure 存储帐户名称。
      
      设置正在使用的 Azure 存储帐户名称。
      
      如果要使用 SAS URL 或仿真器，请保留为空白，否则需要设置。
      
      如果此项为空，并且 env_auth 被设置，则将从环境变量 `AZURE_STORAGE_ACCOUNT_NAME` 中读取（如果可能）。

   --env-auth
      从运行时（环境变量、CLI 或MSI）中读取凭据。
      
      有关完整信息，请参阅[身份验证文档](/azureblob#authentication)。

   --key
      存储帐户的共享密钥。
      
      要使用 SAS URL 或仿真器，请保留为空白。

   --sas-url
      仅限容器访问的 SAS URL。
      
      使用帐户/密钥或仿真器时，请保留为空白。

   --tenant
      服务主体租户的 ID。也称为其目录 ID。
      
      如果使用以下内容，
      - 使用客户端密钥的服务主体
      - 使用证书的服务主体
      - 使用用户名和密码的用户
      
      则设置此项。

   --client-id
      正在使用的客户端的 ID。
      
      如果使用以下内容，
      - 使用客户端密钥的服务主体
      - 使用证书的服务主体
      - 使用用户名和密码的用户
      
      则设置此项。

   --client-secret
      是服务主体的客户端秘密之一。
      
      如果使用以下内容，
      - 使用客户端密钥的服务主体
      
      则设置此项。

   --client-certificate-path
      包含私钥的 PEM 或 PKCS12 证书文件的路径。
      
      如果使用以下内容，
      - 使用证书的服务主体
      
      则设置此项。

   --client-certificate-password
      证书文件的密码（可选项）。
      
      如果使用以下内容，
      - 使用证书的服务主体
      
      并且证书有密码，则可选地设置此项。

   --client-send-certificate-chain
      在使用证书进行身份验证时发送证书链。
      
      指定身份验证请求是否包含 x5c 标头以支持基于主题名称 / 颁发者的验证。设置为 true 时，身份验证请求将包含 x5c 标头。
      
      如果使用以下内容，
      - 使用证书的服务主体
      
      则可选地设置此项。

   --username
      用户名（通常是电子邮件地址）。
      
      如果使用以下内容，
      - 使用用户名和密码的用户
      
      则设置此项。

   --password
      用户的密码。
      
      如果使用以下内容，
      - 使用用户名和密码的用户
      
      则设置此项。

   --service-principal-file
      包含用于服务主体的凭据的文件的路径。
      
      通常保留为空白。仅在希望使用服务主体而不是交互式登录时才需要。
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      有关详细信息，请参阅["创建 Azure 服务主体"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) 和["分配对 Blob 数据访问的 Azure 角色"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli) 页面。
      
      直接将凭据放入 rclone 配置文件的 `client_id`、`tenant` 和 `client_secret` 键中，而不是设置 `service_principal_file`，可能更加方便。

   --use-msi
      使用托管服务标识进行身份验证（仅在 Azure 中有效）。
      
      当为 true 时，使用[托管服务标识](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)来对 Azure 存储进行身份验证，而不是使用 SAS 令牌或帐户密钥。
      
      如果运行该程序的 VM（SS）具有系统分配的标识，则默认情况下将使用该标识。如果资源没有系统分配的标识，但恰好有一个用户分配的标识，则将默认使用用户分配的标识。如果资源有多个用户分配的标识，则必须使用指定的 msi_object_id、msi_client_id 或 msi_mi_res_id 参数明确指定要使用的标识。

   --msi-object-id
      要使用的用户分配的 MSI 的对象 ID（如果有）。
      
      如果指定了 msi_client_id 或 msi_mi_res_id，请保留为空白。

   --msi-client-id
      要使用的用户分配的 MSI 的客户端 ID（如果有）。
      
      如果指定了 msi_object_id 或 msi_mi_res_id，请保留为空白。

   --msi-mi-res-id
      要使用的用户分配的 MSI 的 Azure 资源 ID（如果有）。
      
      如果指定了 msi_client_id 或 msi_object_id，请保留为空白。

   --use-emulator
      如果提供为 'true'，则使用本地存储仿真器。
      
      如果使用真实的 Azure 存储端点，请保留为空白。

   --endpoint
      服务的端点。
      
      通常保留为空白。

   --upload-cutoff
      切换到分块上传的截止点（<= 256 MiB）（已弃用）。

   --chunk-size
      上传块大小。
      
      请注意，此项存储在内存中，同一文件可能使用 "--transfers" * "--azureblob-upload-concurrency" 个块同时存储在内存中。

   --upload-concurrency
      多部分上传的并发数。
      
      同时上传文件的相同块数的并发数。
      
      如果您通过高速链接上传少量的大文件，并且这些上传未充分利用您的带宽，则增加此数值可能有助于加快传输速度。
      
      在测试中，上传速度几乎与上传并发线性增长。例如，要填满一个千兆位编码器，可能需要将此数值提高到 64。请注意，这将使用更多内存。
      
      请注意，块存储在内存中，并且可能存在 "--transfers" * "--azureblob-upload-concurrency" 个块同时存储在内存中。

   --list-chunk
      存储列表的大小。
      
      这设置了每个列表块中请求的 blob 数量。默认值为最大数 5000。每个 “列出 blob” 请求允许使用 2 分钟每兆字节完成。如果操作平均每兆字节超过 2 分钟，则会超时（[来源](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)）。这可以用来限制要返回的 blob 项数，以避免超时。

   --access-tier
      Blob 的访问层：热、冷或存档。
      
      存档 blob 可以通过将访问层设置为热或冷来恢复。如果打算使用默认的访问层，则保留为空白。
      
      如果没有指定 "访问层"，rclone 不应用任何层。在上传时，rclone 在 blob 上执行 "Set Tier" 操作，如果对象未被修改，则将 "访问层" 指定为新的层将没有效果。如果远程的 blob 处于 "存档层"，则不允许尝试从远程执行数据传输操作。用户应该首先通过将 blob 划分到 "热" 或 "冷" 来恢复。

   --archive-tier-delete
      在覆盖之前删除存档层 blob。
      
      无法更新存档层 blob。因此，如果尝试更新存档层 blob，则 rclone 将生成错误：
      
          without --azureblob-archive-tier-delete，无法更新存档层 blob
      
      如果设置了此标志，则在 rclone 尝试覆盖存档层 blob 之前，将删除现有 blob 并上传其替代品。如果上传失败（与更新普通 blob 不同）可能导致数据丢失，并且可能会产生更多费用，因为较早地删除存档层 blob 可能会产生费用。

   --disable-checksum
      不要使用对象元数据存储 MD5 校验和。
      
      通常，rclone 会在上传之前计算输入的 MD5 校验和，以便将其添加到对象的元数据中。这对于数据完整性检查很有用，但对于大文件来说可能会导致长时间的延迟才能开始上传。

   --memory-pool-flush-time
      内部内存缓冲池将刷新的频率。
      
      需要额外缓冲区进行上传（如分块上传）将使用内存池进行分配。
      此选项控制要将未使用的缓冲区从池中删除的频率。

   --memory-pool-use-mmap
      是否在内部内存池中使用 mmap 缓冲区。

   --encoding
      后端的编码方式。
      
      更多信息，请参阅[概述中的编码部分](/overview/#encoding)。

   --public-access
      容器的公共访问级别：blob 或容器。

      示例：
         | <unset>   | 容器及其 Blob 仅能通过授权请求访问。
         |           | 这是默认值。
         | blob      | 此容器中的 Blob 数据可以通过匿名请求读取。
         | container | 容器和 Blob 数据可以完全公开读取。

   --no-check-container
      如果设置，则不尝试检查容器是否存在或创建容器。
      
      当希望最小化 rclone 的交易次数时，这很有用，如果您知道容器已经存在。

   --no-head-object
      如果设置，获取对象时不执行 HEAD 再 GET。

选项:
   --account value                      Azure 存储帐户名称。[$ACCOUNT]
   --client-certificate-password value  证书文件的密码（可选项）。[$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      包含包括私钥在内的 PEM 或 PKCS12 证书文件的路径。[$CLIENT_CERTIFICATE_PATH]
   --client-id value                    正在使用的客户端的 ID。[$CLIENT_ID]
   --client-secret value                是服务主体的客户端秘密之一。[$CLIENT_SECRET]
   --env-auth                           从运行时（环境变量、CLI 或 MSI）中读取凭据。（默认值：false）[$ENV_AUTH]
   --help, -h                           查看帮助
   --key value                          存储帐户的共享密钥。[$KEY]
   --sas-url value                      仅限于容器访问的 SAS URL。[$SAS_URL]
   --tenant value                       服务主体租户的 ID。也称为其目录 ID。[$TENANT]

   Advanced

   --access-tier value              Blob 的访问层：热、冷或存档。[$ACCESS_TIER]
   --archive-tier-delete            在覆盖之前删除存档层 blob。（默认值：false）[$ARCHIVE_TIER_DELETE]
   --chunk-size value               上传块大小。（默认值："4Mi"）[$CHUNK_SIZE]
   --client-send-certificate-chain  在使用证书进行身份验证时发送证书链。（默认值：false）[$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               不要使用对象元数据存储 MD5 校验和。（默认值：false）[$DISABLE_CHECKSUM]
   --encoding value                 后端的编码方式。（默认值："Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"）[$ENCODING]
   --endpoint value                 服务的端点。[$ENDPOINT]
   --list-chunk value               存储列表的大小。（默认值：5000）[$LIST_CHUNK]
   --memory-pool-flush-time value   内部内存缓冲池将刷新的频率。（默认值："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           是否在内部内存池中使用 mmap 缓冲区。（默认值：false）[$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            要使用的用户分配的 MSI 的客户端 ID（如果有）。[$MSI_CLIENT_ID]
   --msi-mi-res-id value            要使用的用户分配的 MSI 的 Azure 资源 ID（如果有）。[$MSI_MI_RES_ID]
   --msi-object-id value            要使用的用户分配的 MSI 的对象 ID（如果有）。[$MSI_OBJECT_ID]
   --no-check-container             如果设置，则不尝试检查容器是否存在或创建容器。（默认值：false）[$NO_CHECK_CONTAINER]
   --no-head-object                 如果设置，获取对象时不执行 HEAD 再 GET。（默认值：false）[$NO_HEAD_OBJECT]
   --password value                 用户的密码。[$PASSWORD]
   --public-access value            容器的公共访问级别：blob 或容器。[$PUBLIC_ACCESS]
   --service-principal-file value   包含用于服务主体的凭据的文件的路径。[$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       多部分上传的并发数。（默认值：16）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切换到分块上传的截止点（<= 256 MiB）（已弃用）。[$UPLOAD_CUTOFF]
   --use-emulator                   如果提供为 'true'，则使用本地存储仿真器。（默认值：false）[$USE_EMULATOR]
   --use-msi                        使用托管服务标识进行身份验证（仅在 Azure 中有效）。 （默认值：false）[$USE_MSI]
   --username value                 用户名（通常是电子邮件地址）。[$USERNAME]

   General

   --name value  存储的名称（默认值：自动生成）
   --path value  存储的路径
```
{% endcode %}