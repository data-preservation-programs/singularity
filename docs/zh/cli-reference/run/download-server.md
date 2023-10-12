# 连接到远程元数据API的HTTP服务器，提供CAR文件下载

{% code fullWidth="true" %}
```
NAME:
   singularity run download-server - 一个连接到远程元数据API的HTTP服务器，用于提供CAR文件下载

使用方法:
   singularity run download-server [命令选项] [参数...]

描述:
   示例用法:
     singularity run download-server --metadata-api "http://remote-metadata-api:7777" --bind "127.0.0.1:8888"

选项:
   --help, -h  显示帮助信息

   1Fichier

   --fichier-api-key value          你的API密钥，从https://1fichier.com/console/params.pl获取。 [$FICHIER_API_KEY]
   --fichier-file-password value    如果你想下载一个受密码保护的共享文件，请添加此参数。 [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  如果你想列出一个受密码保护的共享文件夹中的文件，请添加此参数。 [$FICHIER_FOLDER_PASSWORD]

   Akamai NetStorage

   --netstorage-secret value  设置NetStorage帐户的密钥/G2O密钥以进行身份验证。 [$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-client-secret value  OAuth客户端密钥。 [$ACD_CLIENT_SECRET]
   --acd-token value          JSON blob格式的OAuth访问令牌。 [$ACD_TOKEN]
   --acd-token-url value      令牌服务器URL。 [$ACD_TOKEN_URL]

   支持Amazon S3的云存储提供商（包括AWS，阿里巴巴，Ceph，中国移动，Cloudflare，ArvanCloud，DigitalOcean，Dreamhost，华为OBS，IBM COS，IDrive e2，IONOS Cloud，Liara，Lyve Cloud，Minio，Netease，RackCorp，Scaleway，SeaweedFS，StackPath，Storj，腾讯COS，七牛和Wasabi）

   --s3-access-key-id value            AWS访问密钥ID。 [$S3_ACCESS_KEY_ID]
   --s3-secret-access-key value        AWS密钥（密码）。 [$S3_SECRET_ACCESS_KEY]
   --s3-session-token value            AWS会话令牌。 [$S3_SESSION_TOKEN]
   --s3-sse-customer-key value         如果要使用SSE-C，可以提供用于加密/解密数据的秘密加密密钥。 [$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  如果使用SSE-C，则必须提供以Base64格式编码的秘密加密密钥，用于加密/解密数据。 [$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     如果使用SSE-C，可以提供秘密加密密钥的MD5校验和（可选）。 [$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sse-kms-key-id value           如果使用KMS ID，则必须提供密钥的ARN。 [$S3_SSE_KMS_KEY_ID]

   Backblaze B2

   --b2-key value  应用程序密钥。 [$B2_KEY]

   Box

   --box-access-token value   Box应用程序主访问令牌 [$BOX_ACCESS_TOKEN]
   --box-client-secret value  OAuth客户端密钥。 [$BOX_CLIENT_SECRET]
   --box-token value          JSON blob格式的OAuth访问令牌。 [$BOX_TOKEN]
   --box-token-url value      令牌服务器URL。 [$BOX_TOKEN_URL]

   客户端配置

   --client-ca-cert value                           用于验证服务器的CA证书的路径。要删除，请使用空字符串。
   --client-cert value                              用于相互TLS身份验证的客户端SSL证书（PEM）的路径。要删除，请使用空字符串。
   --client-connect-timeout value                   HTTP客户端连接超时（默认值：1m0s）
   --client-expect-continue-timeout value           使用expect / 100-continue进行HTTP超时（默认值：1s）
   --client-header value [ --client-header value ]  为所有事务设置HTTP标头（即key=value）。这将替换现有的header值。要删除标题，请使用--http-header“key=”。要删除所有标题，请使用--http-header“”。
   --client-insecure-skip-verify                    不验证服务器SSL证书（不安全）（默认值：false）
   --client-key value                               客户端SSL私钥（PEM）的路径，用于相互TLS身份验证。要删除，请使用空字符串。
   --client-no-gzip                                 不设置Accept-Encoding：gzip（默认值：false）
   --client-scan-concurrency value                  扫描数据源时的最大并发列表请求数（默认值：1）
   --client-timeout value                           IO空闲超时（默认值：5m0s）
   --client-use-server-mod-time                     如果可能，请使用服务器修改时间（默认值：false）
   --client-user-agent value                        将用户代理设置为指定的字符串。要删除，请使用空字符串。（默认值：rclone/v1.62.2-DEV）

   Dropbox

   --dropbox-client-secret value  OAuth客户端密钥。 [$DROPBOX_CLIENT_SECRET]
   --dropbox-token value          JSON blob格式的OAuth访问令牌。 [$DROPBOX_TOKEN]
   --dropbox-token-url value      令牌服务器URL。 [$DROPBOX_TOKEN_URL]

   企业文件存储

   --filefabric-permanent-token value  永久身份验证令牌。 [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-token value            会话令牌。 [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     令牌到期时间。 [$FILEFABRIC_TOKEN_EXPIRY]

   FTP

   --ftp-ask-password  需要时允许询问FTP密码。 （默认值：false） [$FTP_ASK_PASSWORD]
   --ftp-pass value    FTP密码。 [$FTP_PASS]

   一般配置

   --bind value          将HTTP服务器绑定到的地址（默认值："127.0.0.1:8888"）
   --metadata-api value  元数据API的URL（默认值："http://127.0.0.1:7777"）

   Google云存储（不是Google Drive）

   --gcs-client-secret value  OAuth客户端密钥。 [$GCS_CLIENT_SECRET]
   --gcs-token value          JSON blob格式的OAuth访问令牌。 [$GCS_TOKEN]
   --gcs-token-url value      令牌服务器URL。 [$GCS_TOKEN_URL]

   Google Drive

   --drive-client-secret value  OAuth客户端密钥。 [$DRIVE_CLIENT_SECRET]
   --drive-resource-key value   用于访问共享链接的资源密钥。 [$DRIVE_RESOURCE_KEY]
   --drive-token value          JSON blob格式的OAuth访问令牌。 [$DRIVE_TOKEN]
   --drive-token-url value      令牌服务器URL。 [$DRIVE_TOKEN_URL]

   Google相册

   --gphotos-client-secret value  OAuth客户端密钥。 [$GPHOTOS_CLIENT_SECRET]
   --gphotos-token value          JSON blob格式的OAuth访问令牌。 [$GPHOTOS_TOKEN]
   --gphotos-token-url value      令牌服务器URL。 [$GPHOTOS_TOKEN_URL]

   HiDrive

   --hidrive-client-secret value  OAuth客户端密钥。 [$HIDRIVE_CLIENT_SECRET]
   --hidrive-token value          JSON blob格式的OAuth访问令牌。 [$HIDRIVE_TOKEN]
   --hidrive-token-url value      令牌服务器URL。 [$HIDRIVE_TOKEN_URL]

   互联网档案馆

   --internetarchive-access-key-id value      IAS3访问密钥。 [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-secret-access-key value  IAS3密钥（密码）。 [$INTERNETARCHIVE_SECRET_ACCESS_KEY]

   Koofr，Digi Storage和其他与Koofr兼容的存储提供商

   --koofr-password value  rclone的密码（在https://storage.rcs-rds.ro/app/admin/preferences/password中生成）。 [$KOOFR_PASSWORD]

   Mail.ru Cloud

   --mailru-pass value  密码。 [$MAILRU_PASS]

   Mega

   --mega-pass value  密码。 [$MEGA_PASS]

   Microsoft Azure Blob Storage

   --azureblob-client-certificate-password value  证书文件的密码（可选）。 [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-secret value                服务主体的客户机密码之一 [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-key value                          存储帐户共享密钥。 [$AZUREBLOB_KEY]
   --azureblob-password value                     用户的密码 [$AZUREBLOB_PASSWORD]

   Microsoft OneDrive

   --onedrive-client-secret value  OAuth客户端密钥。 [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-link-password value  为链接命令创建的链接设置密码。 [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-token value          JSON blob格式的OAuth访问令牌。 [$ONEDRIVE_TOKEN]
   --onedrive-token-url value      令牌服务器URL。 [$ONEDRIVE_TOKEN_URL]

   OpenDrive

   --opendrive-password value  密码。 [$OPENDRIVE_PASSWORD]

   OpenStack Swift（Rackspace Cloud Files，Memset Memstore，OVH）

   --swift-application-credential-secret value  应用程序凭据密钥（OS_APPLICATION_CREDENTIAL_SECRET）。 [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth-token value                     替代验证的身份验证令牌-可选（OS_AUTH_TOKEN）。 [$SWIFT_AUTH_TOKEN]
   --swift-key value                            API密钥或密码（OS_PASSWORD）。 [$SWIFT_KEY]

   Oracle Cloud基础架构对象存储

   --oos-sse-customer-key value         使用SSE-C，可选的标头，指定用于 [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    使用SSE-C的文件，包含与 [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  如果使用SSE-C，则基于64的加密 [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           如果在vault中使用自己的主密钥，则此标头指定 [$OOS_SSE_KMS_KEY_ID]

   Pcloud

   --pcloud-client-secret value  OAuth客户端密钥。 [$PCLOUD_CLIENT_SECRET]
   --pcloud-password value       pcloud密码。 [$PCLOUD_PASSWORD]
   --pcloud-token value          JSON blob格式的OAuth访问令牌。 [$PCLOUD_TOKEN]
   --pcloud-token-url value      令牌服务器URL。 [$PCLOUD_TOKEN_URL]

   QingCloud对象存储

   --qingstor-access-key-id value      QingStor访问密钥ID。 [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-secret-access-key value  QingStor的密钥（密码）。 [$QINGSTOR_SECRET_ACCESS_KEY]

   重试策略

   --client-low-level-retries value  低级客户端错误的最大重试次数（默认值：10）
   --client-retry-backoff value      读取IO错误的恒定延迟退避（默认值：1s）
   --client-retry-backoff-exp value  重试IO读取错误的指数退避（默认值：1.0）
   --client-retry-delay value        重试IO读取错误之前的初始延迟（默认值：1s）
   --client-retry-max value          IO读取错误的最大重试次数（默认值：10）
   --client-skip-inaccessible        打开时跳过不可访问的文件（默认值：false）

   SMB / CIFS

   --smb-pass value  SMB密码。 [$SMB_PASS]

   SSH/SFTP

   --sftp-ask-password         需要时允许询问SFTP密码。 （默认值：false） [$SFTP_ASK_PASSWORD]
   --sftp-key-exchange value   以空格分隔的密钥交换算法列表，按首选顺序排序。 [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value       PEM编码的私钥文件的路径。 [$SFTP_KEY_FILE]
   --sftp-key-file-pass value  解密PEM编码的私钥文件的密码。 [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value        原始的PEM编码的私钥。 [$SFTP_KEY_PEM]
   --sftp-key-use-agent        当设置时，强制使用ssh-agent。 （默认值：false） [$SFTP_KEY_USE_AGENT]
   --sftp-pass value           SSH密码，留空以使用ssh-agent。 [$SFTP_PASS]
   --sftp-pubkey-file value    公共密钥文件的可选路径。 [$SFTP_PUBKEY_FILE]

   Sia分散式云

   --sia-api-password value  Sia Daemon API密码。 [$SIA_API_PASSWORD]

   Storj分散式云存储

   --storj-api-key value     API密钥。 [$STORJ_API_KEY]
   --storj-passphrase value  加密密码。 [$STORJ_PASSPHRASE]

   Sugarsync

   --sugarsync-access-key-id value       Sugarsync访问密钥ID。 [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-private-access-key value  Sugarsync私有访问密钥。 [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value       Sugarsync刷新令牌。 [$SUGARSYNC_REFRESH_TOKEN]

   Uptobox

   --uptobox-access-token value 你的访问令牌。 [$UPTOBOX_ACCESS_TOKEN]

   WebDAV

   --webdav-bearer-token value          用户/密码代替带有基本令牌的令牌（如Macaroon）。 [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  运行以获得基本令牌的命令。 [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-pass value                  密码。 [$WEBDAV_PASS]

   Yandex磁盘

   --yandex-client-secret value  OAuth客户端密钥。 [$YANDEX_CLIENT_SECRET]
   --yandex-token value  OAuth访问令牌作为JSON blob。 [$YANDEX_TOKEN]
   --yandex-token-url value  令牌服务器URL。 [$YANDEX_TOKEN_URL]

   Zoho

   --zoho-client-secret value  OAuth客户端密钥。 [$ZOHO_CLIENT_SECRET]
   --zoho-token value          OAuth访问令牌作为JSON blob。 [$ZOHO_TOKEN]
   --zoho-token-url value      令牌服务器URL。 [$ZOHO_TOKEN_URL]

   premiumize.me

   --premiumizeme-api-key value  API密钥。 [$PREMIUMIZEME_API_KEY]

   Seafile

   --seafile-auth-token value   认证令牌。 [$SEAFILE_AUTH_TOKEN]
   --seafile-library-key value  图书馆密码（仅适用于加密图书馆）。 [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value         密码。 [$SEAFILE_PASS]

```
{% endcode %}