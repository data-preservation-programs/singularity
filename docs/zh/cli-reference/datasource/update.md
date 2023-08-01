# 更新数据源的配置选项

## 用法

```
singularity datasource update [命令选项] <source_id>
```

## 选项

### 数据准备选项

- `--delete-after-export`：导出数据集为 CAR 文件后删除数据集的文件。（默认为 false）
- `--rescan-interval`：在最后一次成功扫描之后经过指定时间后自动重新扫描源目录。单位为秒。（默认为禁用）
- `--scanning-state`：设置初始扫描状态。（默认为 ready）

### acd 选项

- `--acd-auth-url`：Auth 服务器的 URL。[$ACD_AUTH_URL]
- `--acd-client-id`：OAuth 客户端 ID。[$ACD_CLIENT_ID]
- `--acd-client-secret`：OAuth 客户端密钥。[$ACD_CLIENT_SECRET]
- `--acd-encoding`：后端的编码方式。（默认为 "Slash,InvalidUtf8,Dot"）[$ACD_ENCODING]
- `--acd-templink-threshold`：文件大小超过此大小时将通过临时链接进行下载。单位为字节。（默认为 "9Gi"）[$ACD_TEMPLINK_THRESHOLD]
- `--acd-token`：OAuth 访问令牌，格式为 JSON。[$ACD_TOKEN]
- `--acd-token-url`：Token 服务器的 URL。[$ACD_TOKEN_URL]
- `--acd-upload-wait-per-gb`：在完整上传失败后每 GB 额外等待的时间。单位为秒。（默认为 "3m0s"）[$ACD_UPLOAD_WAIT_PER_GB]

### azureblob 选项

- `--azureblob-access-tier`：blob 的访问层级：hot、cool 或 archive。[$AZUREBLOB_ACCESS_TIER]
- `--azureblob-account`：Azure 存储帐户名称。[$AZUREBLOB_ACCOUNT]
- `--azureblob-archive-tier-delete`：覆盖归档层级的 blob 之前删除归档层级的 blob。（默认为 false）[$AZUREBLOB_ARCHIVE_TIER_DELETE]
- `--azureblob-chunk-size`：上传块的大小。（默认为 "4Mi"）[$AZUREBLOB_CHUNK_SIZE]
- `--azureblob-client-certificate-password`：证书文件的密码（可选）。[$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
- `--azureblob-client-certificate-path`：包括私钥的 PEM 或 PKCS12 证书文件的路径。[$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
- `--azureblob-client-id`：正在使用的客户端的 ID。[$AZUREBLOB_CLIENT_ID]
- `--azureblob-client-secret`：其中一个服务主体的客户端密钥。[$AZUREBLOB_CLIENT_SECRET]
- `--azureblob-client-send-certificate-chain`：当使用证书身份验证时发送证书链。（默认为 false）[$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
- `--azureblob-disable-checksum`：不要使用对象元数据存储 MD5 校验和。（默认为 false）[$AZUREBLOB_DISABLE_CHECKSUM]
- `--azureblob-encoding`：后端的编码方式。（默认为 "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"）[$AZUREBLOB_ENCODING]
- `--azureblob-endpoint`：服务的终结点。[$AZUREBLOB_ENDPOINT]
- `--azureblob-env-auth`：从运行时（环境变量、CLI 或 MSI）中读取凭据。（默认为 false）[$AZUREBLOB_ENV_AUTH]
- `--azureblob-key`：存储帐户共享密钥。[$AZUREBLOB_KEY]
- `--azureblob-list-chunk`：Blob 列表的大小。（默认为 "5000"）[$AZUREBLOB_LIST_CHUNK]
- `--azureblob-memory-pool-flush-time`：内部内存缓冲区池刷新的时间间隔。（默认为 "1m0s"）[$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
- `--azureblob-memory-pool-use-mmap`：是否在内部内存池中使用 mmap 缓冲区。（默认为 false）[$AZUREBLOB_MEMORY_POOL_USE_MMAP]
- `--azureblob-msi-client-id`：要使用的用户分配的 MSI 的对象 ID（如果有）。[$AZUREBLOB_MSI_CLIENT_ID]
- `--azureblob-msi-mi-res-id`：要使用的用户分配的 MSI 的 Azure 资源 ID（如果有）。[$AZUREBLOB_MSI_MI_RES_ID]
- `--azureblob-msi-object-id`：要使用的用户分配的 MSI 的对象 ID（如果有）。[$AZUREBLOB_MSI_OBJECT_ID]
- `--azureblob-no-check-container`：如果设置，不尝试检查容器是否存在或创建它。（默认为 false）[$AZUREBLOB_NO_CHECK_CONTAINER]
- `--azureblob-no-head-object`：如果设置，在获取对象时不执行 HEAD 操作。（默认为 false）[$AZUREBLOB_NO_HEAD_OBJECT]
- `--azureblob-password`：用户的密码[$AZUREBLOB_PASSWORD]
- `--azureblob-public-access`：容器的公共访问级别：blob 或 container。[$AZUREBLOB_PUBLIC_ACCESS]
- `--azureblob-sas-url`：仅用于容器级别访问的 SAS URL。[$AZUREBLOB_SAS_URL]
- `--azureblob-service-principal-file`：包含与服务主体一起使用的凭据的文件路径。[$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
- `--azureblob-tenant`：服务主体的租户 ID，也称为其目