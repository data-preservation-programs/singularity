# 更新数据源的配置选项

{% code fullWidth="true" %}
```
NAME:
   singularity datasource update - 更新数据源的配置选项

USAGE:
   singularity datasource update [command options] <source_id>

OPTIONS:
   --help, -h  显示帮助

   数据准备选项

   --delete-after-export    [危险操作] 导出到CAR文件之后删除数据集中的文件。 (默认值：false)
   --rescan-interval value  上一次成功扫描后，自动重新扫描源目录所需的间隔时间。 (默认值：禁用)

   acd的选项

   --acd-auth-url value            授权服务器URL。[$ACD_AUTH_URL]
   --acd-client-id value           OAuth客户端ID。[$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth客户端机密。[$ACD_CLIENT_SECRET]
   --acd-encoding value            后端编码方式。（默认值：“斜杆、无效的UTF8、点”）[$ACD_ENCODING]
   --acd-templink-threshold value  以其tempLink下载文件的文件（大小≥此大小）。 (默认值：“9Gi”) [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth访问令牌，以JSON格式存储。[$ACD_TOKEN]
   --acd-token-url value           令牌服务器URL。[$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  在完全上传失败后，每GiB增加的等待时间。 (默认值：“3m0s”) [$ACD_UPLOAD_WAIT_PER_GB]

   azureblob的选项

   --azureblob-access-tier value                    Blob的访问层级：hot、cool或archive。[$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure存储账户名称。[$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            在覆盖之前删除存档层级块。 (默认值：false)
--local-no-check-updated value   Do not check for updated files on the remote. (default: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-case-insensitive value  Make the path traversal case-insensitive. (default: "false") [$LOCAL_CASE_INSENSITIVE] 
   --local-no-modtime value         Do not read/write the modification time (can speed things up). (default: "false") [$LOCAL_NO_MODTIME] 
   --local-copy-links value         Follow symlinks and copy the pointed to item. (default: "false") [$LOCAL_COPY_LINKS]
   --local-no-check-dest value      Don't check the destination hasn't been modified since the last transfer. (default: "false") [$LOCAL_NO_CHECK_DEST] 
   --local-delete-excluded value    Delete files on destination excluded from sync. (default: "false") [$LOCAL_DELETE_EXCLUDED]
   --local-ignore-existing value    Ignore files that already exist on destination. (default: "false") [$LOCAL_IGNORE_EXISTING] 
   --local-skip-links value         Skip over symlinks. (default: "false") [$LOCAL_SKIP_LINKS] 
   --local-encoding value           The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$LOCAL_ENCODING]
--local-case-insensitive value           强制文件系统报告结果为不区分大小写。(默认值："false") [$LOCAL_CASE_INSENSITIVE]
--local-case-sensitive value             强制文件系统报告结果为区分大小写。(默认值："false") [$LOCAL_CASE_SENSITIVE]
--local-copy-links value, -L value       跟随符号链接并复制指向的内容。(默认值："false") [$LOCAL_COPY_LINKS]
--local-encoding value                   后端编码。(默认值："Slash,Dot") [$LOCAL_ENCODING]
--local-links value, -l value            将符号链接转换为/从常规文件，其具有'.rclonelink'扩展名。(默认值："false") [$LOCAL_LINKS]
--local-no-check-updated value           不检查传输过程中文件是否更改。(默认值："false") [$LOCAL_NO_CHECK_UPDATED]
--local-no-preallocate value             禁用传输文件的磁盘空间的预分配。(默认值："false") [$LOCAL_NO_PREALLOCATE]
--local-no-set-modtime value             禁用设置修改时间的功能。(默认
   --s3-location-constraint value      地域约束 - 必须设置为与区域匹配。 [$S3_LOCATION_CONSTRAINT]
   --s3-max-upload-parts value         分段上传中的最大部分数。 (default: "10000") [$S3_MAX_UPLOAD_PARTS]
   --s3-memory-pool-flush-time value   内部内存缓冲池将刷新的频率。 (default: "1m0s") [$S3_MEMORY_POOL_FLUSH_TIME]
   --s3-memory-pool-use-mmap value     是否在内部内存池中使用mmap缓冲区。 (default: "false") [$S3_MEMORY_POOL_USE_MMAP]
   --s3-might-gzip value               如果后端可能对对象进行gzip，则设置此选项。 (default: "unset") [$S3_MIGHT_GZIP]
   --s3-no-check-bucket value          如果设置，则不尝试检查桶是否存在或创建。 (default: "false") [$S3_NO_CHECK_BUCKET]
   --
--swift-tenant-id value                      租户 ID - 对于 v1 认证是可选的，否则需要此选项或租户 ID (OS_TENANT_ID)。 [$SWIFT_TENANT_ID]
   --swift-user value                           登录用户名 (OS_USERNAME)。 [$SWIFT_USER]
   --swift-user-id value                        登录的用户 ID - 可选的 - 大多数 swift 系统使用用户并留下此字段为空 (v3 认证) (OS_USER_ID)。 [$SWIFT_USER_ID]

   uptobox 的选项

   --uptobox-access-token value  您的访问令牌。 [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      后端编码 (default: "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot")。 [$UPTOBOX_ENCODING]

   WebDAV 的选项

   --webdav-bearer-token value          Bearer 令牌（例如 Macaroon）代替用户名/密码。 [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  用于获取 bearer 令牌的命令。 [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              后端编码。 [$WEBDAV_ENCODING]
   --webdav-headers value               为所有事务设置 HTTP 标头。 [$WEBDAV_HEADERS]
   --webdav-pass value                  密码。 [$WEBDAV_PASS]
   --webdav-url value                   要连接到的 http 主机的 URL。 [$WEBDAV_URL]
   --webdav-user value                  用户名。 [$WEBDAV_USER]
   --webdav-vendor value                您使用的 WebDAV 站点/服务/软件的名称。 [$WEBDAV_VENDOR]

   云雀 的选项

   --yandex-auth-url value       认证服务器的 URL。 [$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth 客户端 ID。 [$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth 客户端密钥。 [$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       后端编码 (default: "Slash,Del,Ctl,InvalidUtf8,Dot")。 [$YANDEX_ENCODING]
   --yandex-hard-delete value    永久删除文件而不是将其放入回收站 (default: "false")。 [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth 访问令牌的 JSON 代码块。 [$YANDEX_TOKEN]
   --yandex-token-url value      令牌服务器 URL。 [$YANDEX_TOKEN_URL]

   Zoho 的选项

   --zoho-auth-url value       认证服务器 URL。 [$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth 客户端 ID。 [$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth 客户端密钥。 [$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       后端编码 (default: "Del,Ctl,InvalidUtf8")。 [$ZOHO_ENCODING]
   --zoho-region value         要连接的 Zoho 区域。 [$ZOHO_REGION]
   --zoho-token value          OAuth 访问令牌的 JSON 代码块。 [$ZOHO_TOKEN]
   --zoho-token-url value      令牌服务器 URL。 [$ZOHO_TOKEN_URL]

```
{% endcode %}