# Download a CAR file from the metadata API

{% code fullWidth="true" %}
```
NAME:
   singularity download - Download a CAR file from the metadata API

USAGE:
   singularity download [command options] <piece_cid>

CATEGORY:
   Utility

OPTIONS:
   1Fichier

   --fichier-api-key value          Your API Key, get it from https://1fichier.com/console/params.pl. [$FICHIER_API_KEY]
   --fichier-file-password value    If you want to download a shared file that is password protected, add this parameter. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  If you want to list the files in a shared folder that is password protected, add this parameter. [$FICHIER_FOLDER_PASSWORD]

   Akamai NetStorage

   --netstorage-secret value  Set the NetStorage account secret/G2O key for authentication. [$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-client-secret value  OAuth Client Secret. [$ACD_CLIENT_SECRET]
   --acd-token value          OAuth Access Token as a JSON blob. [$ACD_TOKEN]
   --acd-token-url value      Token server url. [$ACD_TOKEN_URL]

   Amazon S3 Compliant Storage Providers including AWS, Alibaba, Ceph, China Mobile, Cloudflare, ArvanCloud, DigitalOcean, Dreamhost, Huawei OBS, IBM COS, IDrive e2, IONOS Cloud, Liara, Lyve Cloud, Minio, Netease, RackCorp, Scaleway, SeaweedFS, StackPath, Storj, Tencent COS, Qiniu and Wasabi

   --s3-access-key-id value            AWS Access Key ID. [$S3_ACCESS_KEY_ID]
   --s3-secret-access-key value        AWS Secret Access Key (password). [$S3_SECRET_ACCESS_KEY]
   --s3-session-token value            An AWS session token. [$S3_SESSION_TOKEN]
   --s3-sse-customer-key value         To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     If using SSE-C you may provide the secret encryption key MD5 checksum (optional). [$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sse-kms-key-id value           If using KMS ID you must provide the ARN of Key. [$S3_SSE_KMS_KEY_ID]

   Backblaze B2

   --b2-key value  Application Key. [$B2_KEY]

   Box

   --box-access-token value   Box App Primary Access Token [$BOX_ACCESS_TOKEN]
   --box-client-secret value  OAuth Client Secret. [$BOX_CLIENT_SECRET]
   --box-token value          OAuth Access Token as a JSON blob. [$BOX_TOKEN]
   --box-token-url value      Token server url. [$BOX_TOKEN_URL]

   Client Config

   --client-ca-cert value                           Path to CA certificate used to verify servers. To remove, use empty string.
   --client-cert value                              Path to Client SSL certificate (PEM) for mutual TLS auth. To remove, use empty string.
   --client-connect-timeout value                   HTTP Client Connect timeout (default: 1m0s)
   --client-expect-continue-timeout value           Timeout when using expect / 100-continue in HTTP (default: 1s)
   --client-header value [ --client-header value ]  Set HTTP header for all transactions (i.e. key=value). This will replace the existing header values. To remove a header, use --http-header "key="". To remove all headers, use --http-header ""
   --client-insecure-skip-verify                    Do not verify the server SSL certificate (insecure) (default: false)
   --client-key value                               Path to Client SSL private key (PEM) for mutual TLS auth. To remove, use empty string.
   --client-no-gzip                                 Don't set Accept-Encoding: gzip (default: false)
   --client-scan-concurrency value                  Max number of concurrent listing requests when scanning data source (default: 1)
   --client-timeout value                           IO idle timeout (default: 5m0s)
   --client-use-server-mod-time                     Use server modified time if possible (default: false)
   --client-user-agent value                        Set the user-agent to a specified string. To remove, use empty string. (default: rclone/v1.62.2-DEV)

   Dropbox

   --dropbox-client-secret value  OAuth Client Secret. [$DROPBOX_CLIENT_SECRET]
   --dropbox-token value          OAuth Access Token as a JSON blob. [$DROPBOX_TOKEN]
   --dropbox-token-url value      Token server url. [$DROPBOX_TOKEN_URL]

   Enterprise File Fabric

   --filefabric-permanent-token value  Permanent Authentication Token. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-token value            Session Token. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     Token expiry time. [$FILEFABRIC_TOKEN_EXPIRY]

   FTP

   --ftp-ask-password  Allow asking for FTP password when needed. (default: false) [$FTP_ASK_PASSWORD]
   --ftp-pass value    FTP password. [$FTP_PASS]

   General Config

   --api value          URL of the metadata API (default: "http://127.0.0.1:7777")
   --concurrency value  Number of concurrent downloads (default: 10)
   --out-dir value      Directory to write CAR files to (default: ".")
   --quiet              Suppress all output (default: false)

   Google Cloud Storage (this is not Google Drive)

   --gcs-client-secret value  OAuth Client Secret. [$GCS_CLIENT_SECRET]
   --gcs-token value          OAuth Access Token as a JSON blob. [$GCS_TOKEN]
   --gcs-token-url value      Token server url. [$GCS_TOKEN_URL]

   Google Drive

   --drive-client-secret value  OAuth Client Secret. [$DRIVE_CLIENT_SECRET]
   --drive-resource-key value   Resource key for accessing a link-shared file. [$DRIVE_RESOURCE_KEY]
   --drive-token value          OAuth Access Token as a JSON blob. [$DRIVE_TOKEN]
   --drive-token-url value      Token server url. [$DRIVE_TOKEN_URL]

   Google Photos

   --gphotos-client-secret value  OAuth Client Secret. [$GPHOTOS_CLIENT_SECRET]
   --gphotos-token value          OAuth Access Token as a JSON blob. [$GPHOTOS_TOKEN]
   --gphotos-token-url value      Token server url. [$GPHOTOS_TOKEN_URL]

   HiDrive

   --hidrive-client-secret value  OAuth Client Secret. [$HIDRIVE_CLIENT_SECRET]
   --hidrive-token value          OAuth Access Token as a JSON blob. [$HIDRIVE_TOKEN]
   --hidrive-token-url value      Token server url. [$HIDRIVE_TOKEN_URL]

   Internet Archive

   --internetarchive-access-key-id value      IAS3 Access Key. [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-secret-access-key value  IAS3 Secret Key (password). [$INTERNETARCHIVE_SECRET_ACCESS_KEY]

   Koofr, Digi Storage and other Koofr-compatible storage providers

   --koofr-password value  Your password for rclone (generate one at https://storage.rcs-rds.ro/app/admin/preferences/password). [$KOOFR_PASSWORD]

   Mail.ru Cloud

   --mailru-pass value  Password. [$MAILRU_PASS]

   Mega

   --mega-pass value  Password. [$MEGA_PASS]

   Microsoft Azure Blob Storage

   --azureblob-client-certificate-password value  Password for the certificate file (optional). [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-secret value                One of the service principal's client secrets [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-key value                          Storage Account Shared Key. [$AZUREBLOB_KEY]
   --azureblob-password value                     The user's password [$AZUREBLOB_PASSWORD]

   Microsoft OneDrive

   --onedrive-client-secret value  OAuth Client Secret. [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-link-password value  Set the password for links created by the link command. [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-token value          OAuth Access Token as a JSON blob. [$ONEDRIVE_TOKEN]
   --onedrive-token-url value      Token server url. [$ONEDRIVE_TOKEN_URL]

   OpenDrive

   --opendrive-password value  Password. [$OPENDRIVE_PASSWORD]

   OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

   --swift-application-credential-secret value  Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth-token value                     Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). [$SWIFT_AUTH_TOKEN]
   --swift-key value                            API key or password (OS_PASSWORD). [$SWIFT_KEY]

   Oracle Cloud Infrastructure Object Storage

   --oos-sse-customer-key value         To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           if using using your own master key in vault, this header specifies the [$OOS_SSE_KMS_KEY_ID]

   Pcloud

   --pcloud-client-secret value  OAuth Client Secret. [$PCLOUD_CLIENT_SECRET]
   --pcloud-password value       Your pcloud password. [$PCLOUD_PASSWORD]
   --pcloud-token value          OAuth Access Token as a JSON blob. [$PCLOUD_TOKEN]
   --pcloud-token-url value      Token server url. [$PCLOUD_TOKEN_URL]

   QingCloud Object Storage

   --qingstor-access-key-id value      QingStor Access Key ID. [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-secret-access-key value  QingStor Secret Access Key (password). [$QINGSTOR_SECRET_ACCESS_KEY]

   Retry Strategy

   --client-low-level-retries value  Maximum number of retries for low-level client errors (default: 10)
   --client-retry-backoff value      The constant delay backoff for retrying IO read errors (default: 1s)
   --client-retry-backoff-exp value  The exponential delay backoff for retrying IO read errors (default: 1.0)
   --client-retry-delay value        The initial delay before retrying IO read errors (default: 1s)
   --client-retry-max value          Max number of retries for IO read errors (default: 10)
   --client-skip-inaccessible        Skip inaccessible files when opening (default: false)

   SMB / CIFS

   --smb-pass value  SMB password. [$SMB_PASS]

   SSH/SFTP

   --sftp-ask-password         Allow asking for SFTP password when needed. (default: false) [$SFTP_ASK_PASSWORD]
   --sftp-key-exchange value   Space separated list of key exchange algorithms, ordered by preference. [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value       Path to PEM-encoded private key file. [$SFTP_KEY_FILE]
   --sftp-key-file-pass value  The passphrase to decrypt the PEM-encoded private key file. [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value        Raw PEM-encoded private key. [$SFTP_KEY_PEM]
   --sftp-key-use-agent        When set forces the usage of the ssh-agent. (default: false) [$SFTP_KEY_USE_AGENT]
   --sftp-pass value           SSH password, leave blank to use ssh-agent. [$SFTP_PASS]
   --sftp-pubkey-file value    Optional path to public key file. [$SFTP_PUBKEY_FILE]

   Sia Decentralized Cloud

   --sia-api-password value  Sia Daemon API Password. [$SIA_API_PASSWORD]

   Storj Decentralized Cloud Storage

   --storj-api-key value     API key. [$STORJ_API_KEY]
   --storj-passphrase value  Encryption passphrase. [$STORJ_PASSPHRASE]

   Sugarsync

   --sugarsync-access-key-id value       Sugarsync Access Key ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-private-access-key value  Sugarsync Private Access Key. [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value       Sugarsync refresh token. [$SUGARSYNC_REFRESH_TOKEN]

   Uptobox

   --uptobox-access-token value  Your access token. [$UPTOBOX_ACCESS_TOKEN]

   WebDAV

   --webdav-bearer-token value          Bearer token instead of user/pass (e.g. a Macaroon). [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  Command to run to get a bearer token. [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-pass value                  Password. [$WEBDAV_PASS]

   Yandex Disk

   --yandex-client-secret value  OAuth Client Secret. [$YANDEX_CLIENT_SECRET]
   --yandex-token value          OAuth Access Token as a JSON blob. [$YANDEX_TOKEN]
   --yandex-token-url value      Token server url. [$YANDEX_TOKEN_URL]

   Zoho

   --zoho-client-secret value  OAuth Client Secret. [$ZOHO_CLIENT_SECRET]
   --zoho-token value          OAuth Access Token as a JSON blob. [$ZOHO_TOKEN]
   --zoho-token-url value      Token server url. [$ZOHO_TOKEN_URL]

   premiumize.me

   --premiumizeme-api-key value  API Key. [$PREMIUMIZEME_API_KEY]

   seafile

   --seafile-auth-token value   Authentication token. [$SEAFILE_AUTH_TOKEN]
   --seafile-library-key value  Library password (for encrypted libraries only). [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value         Password. [$SEAFILE_PASS]

```
{% endcode %}
