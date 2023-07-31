# 从元数据API下载CAR文件

{% code fullWidth="true" %}
```
名称：
   singularity download - 从元数据API下载CAR文件

用法：
   singularity download [命令选项] PIECE_CID

类别：
   实用工具

选项：
   通用选项

   --api value                    元数据API的URL（默认值："http://127.0.0.1:7777"）
   --concurrency value, -j value  并发下载数量（默认值：10）
   --out-dir value, -o value      写入CAR文件的目录（默认值："。"）

   acd选项

   --acd-auth-url value            认证服务器的URL。 [$ACD_AUTH_URL]
   --acd-client-id value           OAuth客户端ID。 [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth客户端密钥。 [$ACD_CLIENT_SECRET]
   --acd-encoding value            后端编码方式。（默认值："Slash,InvalidUtf8,Dot"） [$ACD_ENCODING]
   --acd-templink-threshold value  通过临时链接下载文件的大小阈值。（默认值："9Gi"） [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth访问令牌的JSON密钥。 [$ACD_TOKEN]
   --acd-token-url value           令牌服务器URL。 [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完成上传后等待每个GB的附加时间以查看是否出现（默认值："3m0s"） [$ACD_UPLOAD_WAIT_PER_GB]

   azureblob选项

   --azureblob-access-tier value                    Blob的访问层级：hot、cool或archive。 [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure存储帐户名称。 [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            在覆盖之前删除存档层级的块。（默认值："false"） [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     上传块大小。（默认值："4Mi"） [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    证书文件的密码（可选）。 [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        包括私钥在内的PEM或PKCS12证书文件的路径。 [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      在使用的客户端的ID。 [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  要使用的服务主体的一个客户端密钥。 [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  在使用证书auth时发送证书链。(默认值："false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               不要将MD5校验和与对象元数据一起存储。（默认值："false"） [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       后端编码方式。（默认值："Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"） [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       服务的终结点。 [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       从运行环境中读取凭证（环境变量、CLI或MSI）。 （默认值："false"） [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            存储帐户的共享密钥。 [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     Blob列表的大小。 （默认值："5000"） [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         内部内存缓冲池刷新的频率。 （默认值："1m0s"） [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           是否在内部内存缓冲池中使用mmap缓冲区。（默认值："false"） [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  要使用的用户分配的MSI的对象ID（若有）。 [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  要使用的用户分配的MSI的Azure资源ID（若有）。 [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  要使用的用户分配的MSI的对象ID（若有）。 [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             如果设置了，不要尝试检查容器是否存在或创建它。 （默认值："false"） [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 如果设置了，在获取对象时不要做HEAD请求。（默认值："false"） [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       用户的密码 [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  容器的公共访问级别：blob或container。 [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        仅用于容器级别访问的SAS URL。 [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         包含用于服务主体的证书的文件的路径。 [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         服务主体的租户ID。也称为其目录ID。 [$AZUREBLOB