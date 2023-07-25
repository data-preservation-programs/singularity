# Download a CAR file from the metadata API

{% code fullWidth="true" %}
```
NAME:
   singularity download - Download a CAR file from the metadata API

USAGE:
   singularity download [command options] PIECE_CID

CATEGORY:
   Utility

OPTIONS:
   General Options

   --api value                    URL of the metadata API (default: "http://127.0.0.1:7777")
   --concurrency value, -j value  Number of concurrent downloads (default: 10)
   --out-dir value, -o value      Directory to write CAR files to (default: ".")

   Options for acd

   --acd-auth-url value            Auth server URL. [$ACD_AUTH_URL]
   --acd-client-id value           OAuth Client Id. [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth Client Secret. [$ACD_CLIENT_SECRET]
   --acd-encoding value            The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  Files >= this size will be downloaded via their tempLink. (default: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth Access Token as a JSON blob. [$ACD_TOKEN]
   --acd-token-url value           Token server url. [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  Additional time per GiB to wait after a failed complete upload to see if it appears. (default: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

   Options for azureblob

   --azureblob-access-tier value                    Access tier of blob: hot, cool or archive. [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure Storage Account Name. [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            Delete archive tier blobs before overwriting. (default: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     Upload chunk size. (default: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    Password for the certificate file (optional). [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        Path to a PEM or PKCS12 certificate file including the private key. [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      The ID of the client in use. [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  One of the service principal's client secrets [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  Send the certificate chain when using certificate auth. (default: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               Don't store MD5 checksum with object metadata. (default: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       Endpoint for the service. [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       Read credentials from runtime (environment variables, CLI or MSI). (default: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            Storage Account Shared Key. [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     Size of blob list. (default: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         How often internal memory buffer pools will be flushed. (default: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           Whether to use mmap buffers in internal memory pool. (default: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  Object ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  Azure resource ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  Object ID of the user-assigned MSI to use, if any. [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             If set, don't attempt to check the container exists or create it. (default: "false") [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 If set, do not do HEAD before GET when getting objects. (default: "false") [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       The user's password [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  Public access level of a container: blob or container. [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        SAS URL for container level access only. [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         Path to file containing credentials for use with a service principal. [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         ID of the service principal's tenant. Also called its directory ID. [$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             Concurrency for multipart uploads. (default: "16") [$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  Cutoff for switching to chunked upload (<= 256 MiB) (deprecated). [$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulator value                   Uses local storage emulator if provided as 'true'. (default: "false") [$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        Use a managed service identity to authenticate (only works in Azure). (default: "false") [$AZUREBLOB_USE_MSI]
   --azureblob-username value                       User name (usually an email address) [$AZUREBLOB_USERNAME]

   Options for b2

   --b2-account value                 Account ID or Application Key ID. [$B2_ACCOUNT]
   --b2-chunk-size value              Upload chunk size. (default: "96Mi") [$B2_CHUNK_SIZE]
   --b2-copy-cutoff value             Cutoff for switching to multipart copy. (default: "4Gi") [$B2_COPY_CUTOFF]
   --b2-disable-checksum value        Disable checksums for large (> upload cutoff) files. (default: "false") [$B2_DISABLE_CHECKSUM]
   --b2-download-auth-duration value  Time before the authorization token will expire in s or suffix ms|s|m|h|d. (default: "1w") [$B2_DOWNLOAD_AUTH_DURATION]
   --b2-download-url value            Custom endpoint for downloads. [$B2_DOWNLOAD_URL]
   --b2-encoding value                The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$B2_ENCODING]
   --b2-endpoint value                Endpoint for the service. [$B2_ENDPOINT]
   --b2-hard-delete value             Permanently delete files on remote removal, otherwise hide files. (default: "false") [$B2_HARD_DELETE]
   --b2-key value                     Application Key. [$B2_KEY]
   --b2-memory-pool-flush-time value  How often internal memory buffer pools will be flushed. (default: "1m0s") [$B2_MEMORY_POOL_FLUSH_TIME]
   --b2-memory-pool-use-mmap value    Whether to use mmap buffers in internal memory pool. (default: "false") [$B2_MEMORY_POOL_USE_MMAP]
   --b2-test-mode value               A flag string for X-Bz-Test-Mode header for debugging. [$B2_TEST_MODE]
   --b2-upload-cutoff value           Cutoff for switching to chunked upload. (default: "200Mi") [$B2_UPLOAD_CUTOFF]
   --b2-version-at value              Show file versions as they were at the specified time. (default: "off") [$B2_VERSION_AT]
   --b2-versions value                Include old versions in directory listings. (default: "false") [$B2_VERSIONS]

   Options for box

   --box-access-token value     Box App Primary Access Token [$BOX_ACCESS_TOKEN]
   --box-auth-url value         Auth server URL. [$BOX_AUTH_URL]
   --box-box-config-file value  Box App config.json location [$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value     (default: "user") [$BOX_BOX_SUB_TYPE]
   --box-client-id value        OAuth Client Id. [$BOX_CLIENT_ID]
   --box-client-secret value    OAuth Client Secret. [$BOX_CLIENT_SECRET]
   --box-commit-retries value   Max number of times to try committing a multipart file. (default: "100") [$BOX_COMMIT_RETRIES]
   --box-encoding value         The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$BOX_ENCODING]
   --box-list-chunk value       Size of listing chunk 1-1000. (default: "1000") [$BOX_LIST_CHUNK]
   --box-owned-by value         Only show items owned by the login (email address) passed in. [$BOX_OWNED_BY]
   --box-root-folder-id value   Fill in for rclone to use a non root folder as its starting point. (default: "0") [$BOX_ROOT_FOLDER_ID]
   --box-token value            OAuth Access Token as a JSON blob. [$BOX_TOKEN]
   --box-token-url value        Token server url. [$BOX_TOKEN_URL]
   --box-upload-cutoff value    Cutoff for switching to multipart upload (>= 50 MiB). (default: "50Mi") [$BOX_UPLOAD_CUTOFF]

   Options for drive

   --drive-acknowledge-abuse value            Set to allow files which return cannotDownloadAbusiveFile to be downloaded. (default: "false") [$DRIVE_ACKNOWLEDGE_ABUSE]
   --drive-allow-import-name-change value     Allow the filetype to change when uploading Google docs. (default: "false") [$DRIVE_ALLOW_IMPORT_NAME_CHANGE]
   --drive-auth-owner-only value              Only consider files owned by the authenticated user. (default: "false") [$DRIVE_AUTH_OWNER_ONLY]
   --drive-auth-url value                     Auth server URL. [$DRIVE_AUTH_URL]
   --drive-chunk-size value                   Upload chunk size. (default: "8Mi") [$DRIVE_CHUNK_SIZE]
   --drive-client-id value                    Google Application Client Id [$DRIVE_CLIENT_ID]
   --drive-client-secret value                OAuth Client Secret. [$DRIVE_CLIENT_SECRET]
   --drive-copy-shortcut-content value        Server side copy contents of shortcuts instead of the shortcut. (default: "false") [$DRIVE_COPY_SHORTCUT_CONTENT]
   --drive-disable-http2 value                Disable drive using http2. (default: "true") [$DRIVE_DISABLE_HTTP2]
   --drive-encoding value                     The encoding for the backend. (default: "InvalidUtf8") [$DRIVE_ENCODING]
   --drive-export-formats value               Comma separated list of preferred formats for downloading Google docs. (default: "docx,xlsx,pptx,svg") [$DRIVE_EXPORT_FORMATS]
   --drive-formats value                      Deprecated: See export_formats. [$DRIVE_FORMATS]
   --drive-impersonate value                  Impersonate this user when using a service account. [$DRIVE_IMPERSONATE]
   --drive-import-formats value               Comma separated list of preferred formats for uploading Google docs. [$DRIVE_IMPORT_FORMATS]
   --drive-keep-revision-forever value        Keep new head revision of each file forever. (default: "false") [$DRIVE_KEEP_REVISION_FOREVER]
   --drive-list-chunk value                   Size of listing chunk 100-1000, 0 to disable. (default: "1000") [$DRIVE_LIST_CHUNK]
   --drive-pacer-burst value                  Number of API calls to allow without sleeping. (default: "100") [$DRIVE_PACER_BURST]
   --drive-pacer-min-sleep value              Minimum time to sleep between API calls. (default: "100ms") [$DRIVE_PACER_MIN_SLEEP]
   --drive-resource-key value                 Resource key for accessing a link-shared file. [$DRIVE_RESOURCE_KEY]
   --drive-root-folder-id value               ID of the root folder. [$DRIVE_ROOT_FOLDER_ID]
   --drive-scope value                        Scope that rclone should use when requesting access from drive. [$DRIVE_SCOPE]
   --drive-server-side-across-configs value   Allow server-side operations (e.g. copy) to work across different drive configs. (default: "false") [$DRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --drive-service-account-credentials value  Service Account Credentials JSON blob. [$DRIVE_SERVICE_ACCOUNT_CREDENTIALS]
   --drive-service-account-file value         Service Account Credentials JSON file path. [$DRIVE_SERVICE_ACCOUNT_FILE]
   --drive-shared-with-me value               Only show files that are shared with me. (default: "false") [$DRIVE_SHARED_WITH_ME]
   --drive-size-as-quota value                Show sizes as storage quota usage, not actual size. (default: "false") [$DRIVE_SIZE_AS_QUOTA]
   --drive-skip-checksum-gphotos value        Skip MD5 checksum on Google photos and videos only. (default: "false") [$DRIVE_SKIP_CHECKSUM_GPHOTOS]
   --drive-skip-dangling-shortcuts value      If set skip dangling shortcut files. (default: "false") [$DRIVE_SKIP_DANGLING_SHORTCUTS]
   --drive-skip-gdocs value                   Skip google documents in all listings. (default: "false") [$DRIVE_SKIP_GDOCS]
   --drive-skip-shortcuts value               If set skip shortcut files. (default: "false") [$DRIVE_SKIP_SHORTCUTS]
   --drive-starred-only value                 Only show files that are starred. (default: "false") [$DRIVE_STARRED_ONLY]
   --drive-stop-on-download-limit value       Make download limit errors be fatal. (default: "false") [$DRIVE_STOP_ON_DOWNLOAD_LIMIT]
   --drive-stop-on-upload-limit value         Make upload limit errors be fatal. (default: "false") [$DRIVE_STOP_ON_UPLOAD_LIMIT]
   --drive-team-drive value                   ID of the Shared Drive (Team Drive). [$DRIVE_TEAM_DRIVE]
   --drive-token value                        OAuth Access Token as a JSON blob. [$DRIVE_TOKEN]
   --drive-token-url value                    Token server url. [$DRIVE_TOKEN_URL]
   --drive-trashed-only value                 Only show files that are in the trash. (default: "false") [$DRIVE_TRASHED_ONLY]
   --drive-upload-cutoff value                Cutoff for switching to chunked upload. (default: "8Mi") [$DRIVE_UPLOAD_CUTOFF]
   --drive-use-created-date value             Use file created date instead of modified date. (default: "false") [$DRIVE_USE_CREATED_DATE]
   --drive-use-shared-date value              Use date file was shared instead of modified date. (default: "false") [$DRIVE_USE_SHARED_DATE]
   --drive-use-trash value                    Send files to the trash instead of deleting permanently. (default: "true") [$DRIVE_USE_TRASH]
   --drive-v2-download-min-size value         If Object's are greater, use drive v2 API to download. (default: "off") [$DRIVE_V2_DOWNLOAD_MIN_SIZE]

   Options for dropbox

   --dropbox-auth-url value              Auth server URL. [$DROPBOX_AUTH_URL]
   --dropbox-batch-commit-timeout value  Max time to wait for a batch to finish committing (default: "10m0s") [$DROPBOX_BATCH_COMMIT_TIMEOUT]
   --dropbox-batch-mode value            Upload file batching sync|async|off. (default: "sync") [$DROPBOX_BATCH_MODE]
   --dropbox-batch-size value            Max number of files in upload batch. (default: "0") [$DROPBOX_BATCH_SIZE]
   --dropbox-batch-timeout value         Max time to allow an idle upload batch before uploading. (default: "0s") [$DROPBOX_BATCH_TIMEOUT]
   --dropbox-chunk-size value            Upload chunk size (< 150Mi). (default: "48Mi") [$DROPBOX_CHUNK_SIZE]
   --dropbox-client-id value             OAuth Client Id. [$DROPBOX_CLIENT_ID]
   --dropbox-client-secret value         OAuth Client Secret. [$DROPBOX_CLIENT_SECRET]
   --dropbox-encoding value              The encoding for the backend. (default: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$DROPBOX_ENCODING]
   --dropbox-impersonate value           Impersonate this user when using a business account. [$DROPBOX_IMPERSONATE]
   --dropbox-shared-files value          Instructs rclone to work on individual shared files. (default: "false") [$DROPBOX_SHARED_FILES]
   --dropbox-shared-folders value        Instructs rclone to work on shared folders. (default: "false") [$DROPBOX_SHARED_FOLDERS]
   --dropbox-token value                 OAuth Access Token as a JSON blob. [$DROPBOX_TOKEN]
   --dropbox-token-url value             Token server url. [$DROPBOX_TOKEN_URL]

   Options for fichier

   --fichier-api-key value          Your API Key, get it from https://1fichier.com/console/params.pl. [$FICHIER_API_KEY]
   --fichier-encoding value         The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    If you want to download a shared file that is password protected, add this parameter. [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  If you want to list the files in a shared folder that is password protected, add this parameter. [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    If you want to download a shared folder, add this parameter. [$FICHIER_SHARED_FOLDER]

   Options for filefabric

   --filefabric-encoding value         The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-permanent-token value  Permanent Authentication Token. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   ID of the root folder. [$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-token value            Session Token. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     Token expiry time. [$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-url value              URL of the Enterprise File Fabric to connect to. [$FILEFABRIC_URL]
   --filefabric-version value          Version read from the file fabric. [$FILEFABRIC_VERSION]

   Options for ftp

   --ftp-ask-password value          Allow asking for FTP password when needed. (default: "false") [$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         Maximum time to wait for a response to close. (default: "1m0s") [$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           Maximum number of FTP simultaneous connections, 0 for unlimited. (default: "0") [$FTP_CONCURRENCY]
   --ftp-disable-epsv value          Disable using EPSV even if server advertises support. (default: "false") [$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          Disable using MLSD even if server advertises support. (default: "false") [$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         Disable TLS 1.3 (workaround for FTP servers with buggy TLS) (default: "false") [$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          Disable using UTF-8 even if server advertises support. (default: "false") [$FTP_DISABLE_UTF8]
   --ftp-encoding value              The encoding for the backend. (default: "Slash,Del,Ctl,RightSpace,Dot") [$FTP_ENCODING]
   --ftp-explicit-tls value          Use Explicit FTPS (FTP over TLS). (default: "false") [$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     Use LIST -a to force listing of hidden files and folders. This will disable the use of MLSD. (default: "false") [$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  FTP host to connect to. [$FTP_HOST]
   --ftp-idle-timeout value          Max time before closing idle connections. (default: "1m0s") [$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  Do not verify the TLS certificate of the server. (default: "false") [$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTP password. [$FTP_PASS]
   --ftp-port value                  FTP port number. (default: "21") [$FTP_PORT]
   --ftp-shut-timeout value          Maximum time to wait for data connection closing status. (default: "1m0s") [$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   Use Implicit FTPS (FTP over TLS). (default: "false") [$FTP_TLS]
   --ftp-tls-cache-size value        Size of TLS session cache for all control and data connections. (default: "32") [$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTP username. (default: "$USER") [$FTP_USER]
   --ftp-writing-mdtm value          Use MDTM to set modification time (VsFtpd quirk) (default: "false") [$FTP_WRITING_MDTM]

   Options for gcs

   --gcs-anonymous value             Access public buckets and objects without credentials. (default: "false") [$GCS_ANONYMOUS]
   --gcs-auth-url value              Auth server URL. [$GCS_AUTH_URL]
   --gcs-bucket-acl value            Access Control List for new buckets. [$GCS_BUCKET_ACL]
   --gcs-bucket-policy-only value    Access checks should use bucket-level IAM policies. (default: "false") [$GCS_BUCKET_POLICY_ONLY]
   --gcs-client-id value             OAuth Client Id. [$GCS_CLIENT_ID]
   --gcs-client-secret value         OAuth Client Secret. [$GCS_CLIENT_SECRET]
   --gcs-decompress value            If set this will decompress gzip encoded objects. (default: "false") [$GCS_DECOMPRESS]
   --gcs-encoding value              The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$GCS_ENCODING]
   --gcs-endpoint value              Endpoint for the service. [$GCS_ENDPOINT]
   --gcs-env-auth value              Get GCP IAM credentials from runtime (environment variables or instance meta data if no env vars). (default: "false") [$GCS_ENV_AUTH]
   --gcs-location value              Location for the newly created buckets. [$GCS_LOCATION]
   --gcs-no-check-bucket value       If set, don't attempt to check the bucket exists or create it. (default: "false") [$GCS_NO_CHECK_BUCKET]
   --gcs-object-acl value            Access Control List for new objects. [$GCS_OBJECT_ACL]
   --gcs-project-number value        Project number. [$GCS_PROJECT_NUMBER]
   --gcs-service-account-file value  Service Account Credentials JSON file path. [$GCS_SERVICE_ACCOUNT_FILE]
   --gcs-storage-class value         The storage class to use when storing objects in Google Cloud Storage. [$GCS_STORAGE_CLASS]
   --gcs-token value                 OAuth Access Token as a JSON blob. [$GCS_TOKEN]
   --gcs-token-url value             Token server url. [$GCS_TOKEN_URL]

   Options for gphotos

   --gphotos-auth-url value          Auth server URL. [$GPHOTOS_AUTH_URL]
   --gphotos-client-id value         OAuth Client Id. [$GPHOTOS_CLIENT_ID]
   --gphotos-client-secret value     OAuth Client Secret. [$GPHOTOS_CLIENT_SECRET]
   --gphotos-encoding value          The encoding for the backend. (default: "Slash,CrLf,InvalidUtf8,Dot") [$GPHOTOS_ENCODING]
   --gphotos-include-archived value  Also view and download archived media. (default: "false") [$GPHOTOS_INCLUDE_ARCHIVED]
   --gphotos-read-only value         Set to make the Google Photos backend read only. (default: "false") [$GPHOTOS_READ_ONLY]
   --gphotos-read-size value         Set to read the size of media items. (default: "false") [$GPHOTOS_READ_SIZE]
   --gphotos-start-year value        Year limits the photos to be downloaded to those which are uploaded after the given year. (default: "2000") [$GPHOTOS_START_YEAR]
   --gphotos-token value             OAuth Access Token as a JSON blob. [$GPHOTOS_TOKEN]
   --gphotos-token-url value         Token server url. [$GPHOTOS_TOKEN_URL]

   Options for hdfs

   --hdfs-data-transfer-protection value  Kerberos data transfer protection: authentication|integrity|privacy. [$HDFS_DATA_TRANSFER_PROTECTION]
   --hdfs-encoding value                  The encoding for the backend. (default: "Slash,Colon,Del,Ctl,InvalidUtf8,Dot") [$HDFS_ENCODING]
   --hdfs-namenode value                  Hadoop name node and port. [$HDFS_NAMENODE]
   --hdfs-service-principal-name value    Kerberos service principal name for the namenode. [$HDFS_SERVICE_PRINCIPAL_NAME]
   --hdfs-username value                  Hadoop user name. [$HDFS_USERNAME]

   Options for hidrive

   --hidrive-auth-url value                       Auth server URL. [$HIDRIVE_AUTH_URL]
   --hidrive-chunk-size value                     Chunksize for chunked uploads. (default: "48Mi") [$HIDRIVE_CHUNK_SIZE]
   --hidrive-client-id value                      OAuth Client Id. [$HIDRIVE_CLIENT_ID]
   --hidrive-client-secret value                  OAuth Client Secret. [$HIDRIVE_CLIENT_SECRET]
   --hidrive-disable-fetching-member-count value  Do not fetch number of objects in directories unless it is absolutely necessary. (default: "false") [$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
   --hidrive-encoding value                       The encoding for the backend. (default: "Slash,Dot") [$HIDRIVE_ENCODING]
   --hidrive-endpoint value                       Endpoint for the service. (default: "https://api.hidrive.strato.com/2.1") [$HIDRIVE_ENDPOINT]
   --hidrive-root-prefix value                    The root/parent folder for all paths. (default: "/") [$HIDRIVE_ROOT_PREFIX]
   --hidrive-scope-access value                   Access permissions that rclone should use when requesting access from HiDrive. (default: "rw") [$HIDRIVE_SCOPE_ACCESS]
   --hidrive-scope-role value                     User-level that rclone should use when requesting access from HiDrive. (default: "user") [$HIDRIVE_SCOPE_ROLE]
   --hidrive-token value                          OAuth Access Token as a JSON blob. [$HIDRIVE_TOKEN]
   --hidrive-token-url value                      Token server url. [$HIDRIVE_TOKEN_URL]
   --hidrive-upload-concurrency value             Concurrency for chunked uploads. (default: "4") [$HIDRIVE_UPLOAD_CONCURRENCY]
   --hidrive-upload-cutoff value                  Cutoff/Threshold for chunked uploads. (default: "96Mi") [$HIDRIVE_UPLOAD_CUTOFF]

   Options for http

   --http-headers value   Set HTTP headers for all transactions. [$HTTP_HEADERS]
   --http-no-head value   Don't use HEAD requests. (default: "false") [$HTTP_NO_HEAD]
   --http-no-slash value  Set this if the site doesn't end directories with /. (default: "false") [$HTTP_NO_SLASH]
   --http-url value       URL of HTTP host to connect to. [$HTTP_URL]

   Options for internetarchive

   --internetarchive-access-key-id value      IAS3 Access Key. [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   Don't ask the server to test against MD5 checksum calculated by rclone. (default: "true") [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           The encoding for the backend. (default: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3 Endpoint. (default: "https://s3.us.archive.org") [$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     Host of InternetArchive Frontend. (default: "https://archive.org") [$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3 Secret Key (password). [$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       Timeout for waiting the server's processing tasks (specifically archive and book_op) to finish. (default: "0s") [$INTERNETARCHIVE_WAIT_ARCHIVE]

   Options for jottacloud

   --jottacloud-encoding value             The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          Delete files permanently rather than putting them into the trash. (default: "false") [$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     Files bigger than this will be cached on disk to calculate the MD5 if required. (default: "10Mi") [$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          Avoid server side versioning by deleting files and recreating files instead of overwriting them. (default: "false") [$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         Only show files that are in the trash. (default: "false") [$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  Files bigger than this can be resumed if the upload fail's. (default: "10Mi") [$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

   Options for koofr

   --koofr-encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$KOOFR_ENCODING]
   --koofr-endpoint value  The Koofr API endpoint to use. [$KOOFR_ENDPOINT]
   --koofr-mountid value   Mount ID of the mount to use. [$KOOFR_MOUNTID]
   --koofr-password value  Your password for rclone (generate one at https://app.koofr.net/app/admin/preferences/password). [$KOOFR_PASSWORD]
   --koofr-provider value  Choose your storage provider. [$KOOFR_PROVIDER]
   --koofr-setmtime value  Does the backend support setting modification time. (default: "true") [$KOOFR_SETMTIME]
   --koofr-user value      Your user name. [$KOOFR_USER]

   Options for local

   --local-case-insensitive value       Force the filesystem to report itself as case insensitive. (default: "false") [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value         Force the filesystem to report itself as case sensitive. (default: "false") [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value             Follow symlinks and copy the pointed to item. (default: "false") [$LOCAL_COPY_LINKS]
   --local-encoding value               The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$LOCAL_ENCODING]
   --local-links value                  Translate symlinks to/from regular files with a '.rclonelink' extension. (default: "false") [$LOCAL_LINKS]
   --local-no-check-updated value       Don't check to see if the files change during upload. (default: "false") [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value         Disable preallocation of disk space for transferred files. (default: "false") [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value         Disable setting modtime. (default: "false") [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value              Disable sparse files for multi-thread downloads. (default: "false") [$LOCAL_NO_SPARSE]
   --local-nounc value                  Disable UNC (long path names) conversion on Windows. (default: "false") [$LOCAL_NOUNC]
   --local-one-file-system value        Don't cross filesystem boundaries (unix/macOS only). (default: "false") [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value             Don't warn about skipped symlinks. (default: "false") [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value  Apply unicode NFC normalization to paths and filenames. (default: "false") [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value        Assume the Stat size of links is zero (and read them instead) (deprecated). (default: "false") [$LOCAL_ZERO_SIZE_LINKS]

   Options for mailru

   --mailru-check-hash value             What should copy do if file checksum is mismatched or invalid. (default: "true") [$MAILRU_CHECK_HASH]
   --mailru-encoding value               The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$MAILRU_ENCODING]
   --mailru-pass value                   Password. [$MAILRU_PASS]
   --mailru-speedup-enable value         Skip full upload if there is another file with same data hash. (default: "true") [$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  Comma separated list of file name patterns eligible for speedup (put by hash). (default: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       This option allows you to disable speedup (put by hash) for large files. (default: "3Gi") [$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     Files larger than the size given below will always be hashed on disk. (default: "32Mi") [$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   User name (usually email). [$MAILRU_USER]

   Options for mega

   --mega-debug value        Output more debug from Mega. (default: "false") [$MEGA_DEBUG]
   --mega-encoding value     The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$MEGA_ENCODING]
   --mega-hard-delete value  Delete files permanently rather than putting them into the trash. (default: "false") [$MEGA_HARD_DELETE]
   --mega-pass value         Password. [$MEGA_PASS]
   --mega-use-https value    Use HTTPS for transfers. (default: "false") [$MEGA_USE_HTTPS]
   --mega-user value         User name. [$MEGA_USER]

   Options for netstorage

   --netstorage-account value   Set the NetStorage account name [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      Domain+path of NetStorage host to connect to. [$NETSTORAGE_HOST]
   --netstorage-protocol value  Select between HTTP or HTTPS protocol. (default: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    Set the NetStorage account secret/G2O key for authentication. [$NETSTORAGE_SECRET]

   Options for onedrive

   --onedrive-access-scopes value               Set scopes to be requested by rclone. (default: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ONEDRIVE_ACCESS_SCOPES]
   --onedrive-auth-url value                    Auth server URL. [$ONEDRIVE_AUTH_URL]
   --onedrive-chunk-size value                  Chunk size to upload files with - must be multiple of 320k (327,680 bytes). (default: "10Mi") [$ONEDRIVE_CHUNK_SIZE]
   --onedrive-client-id value                   OAuth Client Id. [$ONEDRIVE_CLIENT_ID]
   --onedrive-client-secret value               OAuth Client Secret. [$ONEDRIVE_CLIENT_SECRET]
   --onedrive-drive-id value                    The ID of the drive to use. [$ONEDRIVE_DRIVE_ID]
   --onedrive-drive-type value                  The type of the drive (personal | business | documentLibrary). [$ONEDRIVE_DRIVE_TYPE]
   --onedrive-encoding value                    The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ONEDRIVE_ENCODING]
   --onedrive-expose-onenote-files value        Set to make OneNote files show up in directory listings. (default: "false") [$ONEDRIVE_EXPOSE_ONENOTE_FILES]
   --onedrive-hash-type value                   Specify the hash in use for the backend. (default: "auto") [$ONEDRIVE_HASH_TYPE]
   --onedrive-link-password value               Set the password for links created by the link command. [$ONEDRIVE_LINK_PASSWORD]
   --onedrive-link-scope value                  Set the scope of the links created by the link command. (default: "anonymous") [$ONEDRIVE_LINK_SCOPE]
   --onedrive-link-type value                   Set the type of the links created by the link command. (default: "view") [$ONEDRIVE_LINK_TYPE]
   --onedrive-list-chunk value                  Size of listing chunk. (default: "1000") [$ONEDRIVE_LIST_CHUNK]
   --onedrive-no-versions value                 Remove all versions on modifying operations. (default: "false") [$ONEDRIVE_NO_VERSIONS]
   --onedrive-region value                      Choose national cloud region for OneDrive. (default: "global") [$ONEDRIVE_REGION]
   --onedrive-root-folder-id value              ID of the root folder. [$ONEDRIVE_ROOT_FOLDER_ID]
   --onedrive-server-side-across-configs value  Allow server-side operations (e.g. copy) to work across different onedrive configs. (default: "false") [$ONEDRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --onedrive-token value                       OAuth Access Token as a JSON blob. [$ONEDRIVE_TOKEN]
   --onedrive-token-url value                   Token server url. [$ONEDRIVE_TOKEN_URL]

   Options for oos

   --oos-chunk-size value               Chunk size to use for uploading. (default: "5Mi") [$OOS_CHUNK_SIZE]
   --oos-compartment value              Object storage compartment OCID [$OOS_COMPARTMENT]
   --oos-config-file value              Path to OCI config file (default: "~/.oci/config") [$OOS_CONFIG_FILE]
   --oos-config-profile value           Profile name inside the oci config file (default: "Default") [$OOS_CONFIG_PROFILE]
   --oos-copy-cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$OOS_COPY_CUTOFF]
   --oos-copy-timeout value             Timeout for copy. (default: "1m0s") [$OOS_COPY_TIMEOUT]
   --oos-disable-checksum value         Don't store MD5 checksum with object metadata. (default: "false") [$OOS_DISABLE_CHECKSUM]
   --oos-encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$OOS_ENCODING]
   --oos-endpoint value                 Endpoint for Object storage API. [$OOS_ENDPOINT]
   --oos-leave-parts-on-error value     If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: "false") [$OOS_LEAVE_PARTS_ON_ERROR]
   --oos-namespace value                Object storage namespace [$OOS_NAMESPACE]
   --oos-no-check-bucket value          If set, don't attempt to check the bucket exists or create it. (default: "false") [$OOS_NO_CHECK_BUCKET]
   --oos-provider value                 Choose your Auth Provider (default: "env_auth") [$OOS_PROVIDER]
   --oos-region value                   Object storage Region [$OOS_REGION]
   --oos-sse-customer-algorithm value   If using SSE-C, the optional header that specifies "AES256" as the encryption algorithm. [$OOS_SSE_CUSTOMER_ALGORITHM]
   --oos-sse-customer-key value         To use SSE-C, the optional header that specifies the base64-encoded 256-bit encryption key to use to [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    To use SSE-C, a file containing the base64-encoded string of the AES-256 encryption key associated [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  If using SSE-C, The optional header that specifies the base64-encoded SHA256 hash of the encryption [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           if using using your own master key in vault, this header specifies the [$OOS_SSE_KMS_KEY_ID]
   --oos-storage-tier value             The storage class to use when storing new objects in storage. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (default: "Standard") [$OOS_STORAGE_TIER]
   --oos-upload-concurrency value       Concurrency for multipart uploads. (default: "10") [$OOS_UPLOAD_CONCURRENCY]
   --oos-upload-cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$OOS_UPLOAD_CUTOFF]

   Options for opendrive

   --opendrive-chunk-size value  Files will be uploaded in chunks this size. (default: "10Mi") [$OPENDRIVE_CHUNK_SIZE]
   --opendrive-encoding value    The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$OPENDRIVE_ENCODING]
   --opendrive-password value    Password. [$OPENDRIVE_PASSWORD]
   --opendrive-username value    Username. [$OPENDRIVE_USERNAME]

   Options for pcloud

   --pcloud-auth-url value        Auth server URL. [$PCLOUD_AUTH_URL]
   --pcloud-client-id value       OAuth Client Id. [$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value   OAuth Client Secret. [$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value        The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PCLOUD_ENCODING]
   --pcloud-hostname value        Hostname to connect to. (default: "api.pcloud.com") [$PCLOUD_HOSTNAME]
   --pcloud-password value        Your pcloud password. [$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value  Fill in for rclone to use a non root folder as its starting point. (default: "d0") [$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value           OAuth Access Token as a JSON blob. [$PCLOUD_TOKEN]
   --pcloud-token-url value       Token server url. [$PCLOUD_TOKEN_URL]
   --pcloud-username value        Your pcloud username. [$PCLOUD_USERNAME]

   Options for premiumizeme

   --premiumizeme-encoding value  The encoding for the backend. (default: "Slash,DoubleQuote,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PREMIUMIZEME_ENCODING]

   Options for putio

   --putio-encoding value  The encoding for the backend. (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PUTIO_ENCODING]

   Options for qingstor

   --qingstor-access-key-id value       QingStor Access Key ID. [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-chunk-size value          Chunk size to use for uploading. (default: "4Mi") [$QINGSTOR_CHUNK_SIZE]
   --qingstor-connection-retries value  Number of connection retries. (default: "3") [$QINGSTOR_CONNECTION_RETRIES]
   --qingstor-encoding value            The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8") [$QINGSTOR_ENCODING]
   --qingstor-endpoint value            Enter an endpoint URL to connection QingStor API. [$QINGSTOR_ENDPOINT]
   --qingstor-env-auth value            Get QingStor credentials from runtime. (default: "false") [$QINGSTOR_ENV_AUTH]
   --qingstor-secret-access-key value   QingStor Secret Access Key (password). [$QINGSTOR_SECRET_ACCESS_KEY]
   --qingstor-upload-concurrency value  Concurrency for multipart uploads. (default: "1") [$QINGSTOR_UPLOAD_CONCURRENCY]
   --qingstor-upload-cutoff value       Cutoff for switching to chunked upload. (default: "200Mi") [$QINGSTOR_UPLOAD_CUTOFF]
   --qingstor-zone value                Zone to connect to. [$QINGSTOR_ZONE]

   Options for s3

   --s3-access-key-id value            AWS Access Key ID. [$S3_ACCESS_KEY_ID]
   --s3-acl value                      Canned ACL used when creating buckets and storing or copying objects. [$S3_ACL]
   --s3-bucket-acl value               Canned ACL used when creating buckets. [$S3_BUCKET_ACL]
   --s3-chunk-size value               Chunk size to use for uploading. (default: "5Mi") [$S3_CHUNK_SIZE]
   --s3-copy-cutoff value              Cutoff for switching to multipart copy. (default: "4.656Gi") [$S3_COPY_CUTOFF]
   --s3-decompress value               If set this will decompress gzip encoded objects. (default: "false") [$S3_DECOMPRESS]
   --s3-disable-checksum value         Don't store MD5 checksum with object metadata. (default: "false") [$S3_DISABLE_CHECKSUM]
   --s3-disable-http2 value            Disable usage of http2 for S3 backends. (default: "false") [$S3_DISABLE_HTTP2]
   --s3-download-url value             Custom endpoint for downloads. [$S3_DOWNLOAD_URL]
   --s3-encoding value                 The encoding for the backend. (default: "Slash,InvalidUtf8,Dot") [$S3_ENCODING]
   --s3-endpoint value                 Endpoint for S3 API. [$S3_ENDPOINT]
   --s3-env-auth value                 Get AWS credentials from runtime (environment variables or EC2/ECS meta data if no env vars). (default: "false") [$S3_ENV_AUTH]
   --s3-force-path-style value         If true use path style access if false use virtual hosted style. (default: "true") [$S3_FORCE_PATH_STYLE]
   --s3-leave-parts-on-error value     If true avoid calling abort upload on a failure, leaving all successfully uploaded parts on S3 for manual recovery. (default: "false") [$S3_LEAVE_PARTS_ON_ERROR]
   --s3-list-chunk value               Size of listing chunk (response list for each ListObject S3 request). (default: "1000") [$S3_LIST_CHUNK]
   --s3-list-url-encode value          Whether to url encode listings: true/false/unset (default: "unset") [$S3_LIST_URL_ENCODE]
   --s3-list-version value             Version of ListObjects to use: 1,2 or 0 for auto. (default: "0") [$S3_LIST_VERSION]
   --s3-location-constraint value      Location constraint - must be set to match the Region. [$S3_LOCATION_CONSTRAINT]
   --s3-max-upload-parts value         Maximum number of parts in a multipart upload. (default: "10000") [$S3_MAX_UPLOAD_PARTS]
   --s3-memory-pool-flush-time value   How often internal memory buffer pools will be flushed. (default: "1m0s") [$S3_MEMORY_POOL_FLUSH_TIME]
   --s3-memory-pool-use-mmap value     Whether to use mmap buffers in internal memory pool. (default: "false") [$S3_MEMORY_POOL_USE_MMAP]
   --s3-might-gzip value               Set this if the backend might gzip objects. (default: "unset") [$S3_MIGHT_GZIP]
   --s3-no-check-bucket value          If set, don't attempt to check the bucket exists or create it. (default: "false") [$S3_NO_CHECK_BUCKET]
   --s3-no-head value                  If set, don't HEAD uploaded objects to check integrity. (default: "false") [$S3_NO_HEAD]
   --s3-no-head-object value           If set, do not do HEAD before GET when getting objects. (default: "false") [$S3_NO_HEAD_OBJECT]
   --s3-no-system-metadata value       Suppress setting and reading of system metadata (default: "false") [$S3_NO_SYSTEM_METADATA]
   --s3-profile value                  Profile to use in the shared credentials file. [$S3_PROFILE]
   --s3-provider value                 Choose your S3 provider. [$S3_PROVIDER]
   --s3-region value                   Region to connect to. [$S3_REGION]
   --s3-requester-pays value           Enables requester pays option when interacting with S3 bucket. (default: "false") [$S3_REQUESTER_PAYS]
   --s3-secret-access-key value        AWS Secret Access Key (password). [$S3_SECRET_ACCESS_KEY]
   --s3-server-side-encryption value   The server-side encryption algorithm used when storing this object in S3. [$S3_SERVER_SIDE_ENCRYPTION]
   --s3-session-token value            An AWS session token. [$S3_SESSION_TOKEN]
   --s3-shared-credentials-file value  Path to the shared credentials file. [$S3_SHARED_CREDENTIALS_FILE]
   --s3-sse-customer-algorithm value   If using SSE-C, the server-side encryption algorithm used when storing this object in S3. [$S3_SSE_CUSTOMER_ALGORITHM]
   --s3-sse-customer-key value         To use SSE-C you may provide the secret encryption key used to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  If using SSE-C you must provide the secret encryption key encoded in base64 format to encrypt/decrypt your data. [$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     If using SSE-C you may provide the secret encryption key MD5 checksum (optional). [$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sse-kms-key-id value           If using KMS ID you must provide the ARN of Key. [$S3_SSE_KMS_KEY_ID]
   --s3-storage-class value            The storage class to use when storing new objects in S3. [$S3_STORAGE_CLASS]
   --s3-sts-endpoint value             Endpoint for STS. [$S3_STS_ENDPOINT]
   --s3-upload-concurrency value       Concurrency for multipart uploads. (default: "4") [$S3_UPLOAD_CONCURRENCY]
   --s3-upload-cutoff value            Cutoff for switching to chunked upload. (default: "200Mi") [$S3_UPLOAD_CUTOFF]
   --s3-use-accelerate-endpoint value  If true use the AWS S3 accelerated endpoint. (default: "false") [$S3_USE_ACCELERATE_ENDPOINT]
   --s3-use-multipart-etag value       Whether to use ETag in multipart uploads for verification (default: "unset") [$S3_USE_MULTIPART_ETAG]
   --s3-use-presigned-request value    Whether to use a presigned request or PutObject for single part uploads (default: "false") [$S3_USE_PRESIGNED_REQUEST]
   --s3-v2-auth value                  If true use v2 authentication. (default: "false") [$S3_V2_AUTH]
   --s3-version-at value               Show file versions as they were at the specified time. (default: "off") [$S3_VERSION_AT]
   --s3-versions value                 Include old versions in directory listings. (default: "false") [$S3_VERSIONS]

   Options for seafile

   --seafile-2fa value             Two-factor authentication ('true' if the account has 2FA enabled). (default: "false") [$SEAFILE_2FA]
   --seafile-create-library value  Should rclone create a library if it doesn't exist. (default: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        The encoding for the backend. (default: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$SEAFILE_ENCODING]
   --seafile-library value         Name of the library. [$SEAFILE_LIBRARY]
   --seafile-library-key value     Library password (for encrypted libraries only). [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value            Password. [$SEAFILE_PASS]
   --seafile-url value             URL of seafile host to connect to. [$SEAFILE_URL]
   --seafile-user value            User name (usually email address). [$SEAFILE_USER]

   Options for sftp

   --sftp-ask-password value               Allow asking for SFTP password when needed. (default: "false") [$SFTP_ASK_PASSWORD]
   --sftp-chunk-size value                 Upload and download chunk size. (default: "32Ki") [$SFTP_CHUNK_SIZE]
   --sftp-ciphers value                    Space separated list of ciphers to be used for session encryption, ordered by preference. [$SFTP_CIPHERS]
   --sftp-concurrency value                The maximum number of outstanding requests for one file (default: "64") [$SFTP_CONCURRENCY]
   --sftp-disable-concurrent-reads value   If set don't use concurrent reads. (default: "false") [$SFTP_DISABLE_CONCURRENT_READS]
   --sftp-disable-concurrent-writes value  If set don't use concurrent writes. (default: "false") [$SFTP_DISABLE_CONCURRENT_WRITES]
   --sftp-disable-hashcheck value          Disable the execution of SSH commands to determine if remote file hashing is available. (default: "false") [$SFTP_DISABLE_HASHCHECK]
   --sftp-host value                       SSH host to connect to. [$SFTP_HOST]
   --sftp-idle-timeout value               Max time before closing idle connections. (default: "1m0s") [$SFTP_IDLE_TIMEOUT]
   --sftp-key-exchange value               Space separated list of key exchange algorithms, ordered by preference. [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value                   Path to PEM-encoded private key file. [$SFTP_KEY_FILE]
   --sftp-key-file-pass value              The passphrase to decrypt the PEM-encoded private key file. [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value                    Raw PEM-encoded private key. [$SFTP_KEY_PEM]
   --sftp-key-use-agent value              When set forces the usage of the ssh-agent. (default: "false") [$SFTP_KEY_USE_AGENT]
   --sftp-known-hosts-file value           Optional path to known_hosts file. [$SFTP_KNOWN_HOSTS_FILE]
   --sftp-macs value                       Space separated list of MACs (message authentication code) algorithms, ordered by preference. [$SFTP_MACS]
   --sftp-md5sum-command value             The command used to read md5 hashes. [$SFTP_MD5SUM_COMMAND]
   --sftp-pass value                       SSH password, leave blank to use ssh-agent. [$SFTP_PASS]
   --sftp-path-override value              Override path used by SSH shell commands. [$SFTP_PATH_OVERRIDE]
   --sftp-port value                       SSH port number. (default: "22") [$SFTP_PORT]
   --sftp-pubkey-file value                Optional path to public key file. [$SFTP_PUBKEY_FILE]
   --sftp-server-command value             Specifies the path or command to run a sftp server on the remote host. [$SFTP_SERVER_COMMAND]
   --sftp-set-env value                    Environment variables to pass to sftp and commands [$SFTP_SET_ENV]
   --sftp-set-modtime value                Set the modified time on the remote if set. (default: "true") [$SFTP_SET_MODTIME]
   --sftp-sha1sum-command value            The command used to read sha1 hashes. [$SFTP_SHA1SUM_COMMAND]
   --sftp-shell-type value                 The type of SSH shell on remote server, if any. [$SFTP_SHELL_TYPE]
   --sftp-skip-links value                 Set to skip any symlinks and any other non regular files. (default: "false") [$SFTP_SKIP_LINKS]
   --sftp-subsystem value                  Specifies the SSH2 subsystem on the remote host. (default: "sftp") [$SFTP_SUBSYSTEM]
   --sftp-use-fstat value                  If set use fstat instead of stat. (default: "false") [$SFTP_USE_FSTAT]
   --sftp-use-insecure-cipher value        Enable the use of insecure ciphers and key exchange methods. (default: "false") [$SFTP_USE_INSECURE_CIPHER]
   --sftp-user value                       SSH username. (default: "$USER") [$SFTP_USER]

   Options for sharefile

   --sharefile-chunk-size value      Upload chunk size. (default: "64Mi") [$SHAREFILE_CHUNK_SIZE]
   --sharefile-encoding value        The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SHAREFILE_ENCODING]
   --sharefile-endpoint value        Endpoint for API calls. [$SHAREFILE_ENDPOINT]
   --sharefile-root-folder-id value  ID of the root folder. [$SHAREFILE_ROOT_FOLDER_ID]
   --sharefile-upload-cutoff value   Cutoff for switching to multipart upload. (default: "128Mi") [$SHAREFILE_UPLOAD_CUTOFF]

   Options for sia

   --sia-api-password value  Sia Daemon API Password. [$SIA_API_PASSWORD]
   --sia-api-url value       Sia daemon API URL, like http://sia.daemon.host:9980. (default: "http://127.0.0.1:9980") [$SIA_API_URL]
   --sia-encoding value      The encoding for the backend. (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$SIA_ENCODING]
   --sia-user-agent value    Siad User Agent (default: "Sia-Agent") [$SIA_USER_AGENT]

   Options for smb

   --smb-case-insensitive value    Whether the server is configured to be case-insensitive. (default: "true") [$SMB_CASE_INSENSITIVE]
   --smb-domain value              Domain name for NTLM authentication. (default: "WORKGROUP") [$SMB_DOMAIN]
   --smb-encoding value            The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,RightSpace,RightPeriod,InvalidUtf8,Dot") [$SMB_ENCODING]
   --smb-hide-special-share value  Hide special shares (e.g. print$) which users aren't supposed to access. (default: "true") [$SMB_HIDE_SPECIAL_SHARE]
   --smb-host value                SMB server hostname to connect to. [$SMB_HOST]
   --smb-idle-timeout value        Max time before closing idle connections. (default: "1m0s") [$SMB_IDLE_TIMEOUT]
   --smb-pass value                SMB password. [$SMB_PASS]
   --smb-port value                SMB port number. (default: "445") [$SMB_PORT]
   --smb-spn value                 Service principal name. [$SMB_SPN]
   --smb-user value                SMB username. (default: "$USER") [$SMB_USER]

   Options for storj

   --storj-access-grant value       Access grant. [$STORJ_ACCESS_GRANT]
   --storj-api-key value            API key. [$STORJ_API_KEY]
   --storj-passphrase value         Encryption passphrase. [$STORJ_PASSPHRASE]
   --storj-provider value           Choose an authentication method. (default: "existing") [$STORJ_PROVIDER]
   --storj-satellite-address value  Satellite address. (default: "us1.storj.io") [$STORJ_SATELLITE_ADDRESS]

   Options for sugarsync

   --sugarsync-access-key-id value         Sugarsync Access Key ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-app-id value                Sugarsync App ID. [$SUGARSYNC_APP_ID]
   --sugarsync-authorization value         Sugarsync authorization. [$SUGARSYNC_AUTHORIZATION]
   --sugarsync-authorization-expiry value  Sugarsync authorization expiry. [$SUGARSYNC_AUTHORIZATION_EXPIRY]
   --sugarsync-deleted-id value            Sugarsync deleted folder id. [$SUGARSYNC_DELETED_ID]
   --sugarsync-encoding value              The encoding for the backend. (default: "Slash,Ctl,InvalidUtf8,Dot") [$SUGARSYNC_ENCODING]
   --sugarsync-hard-delete value           Permanently delete files if true (default: "false") [$SUGARSYNC_HARD_DELETE]
   --sugarsync-private-access-key value    Sugarsync Private Access Key. [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value         Sugarsync refresh token. [$SUGARSYNC_REFRESH_TOKEN]
   --sugarsync-root-id value               Sugarsync root id. [$SUGARSYNC_ROOT_ID]
   --sugarsync-user value                  Sugarsync user. [$SUGARSYNC_USER]

   Options for swift

   --swift-application-credential-id value      Application Credential ID (OS_APPLICATION_CREDENTIAL_ID). [$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    Application Credential Name (OS_APPLICATION_CREDENTIAL_NAME). [$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  Application Credential Secret (OS_APPLICATION_CREDENTIAL_SECRET). [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           Authentication URL for server (OS_AUTH_URL). [$SWIFT_AUTH]
   --swift-auth-token value                     Auth Token from alternate authentication - optional (OS_AUTH_TOKEN). [$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   AuthVersion - optional - set to (1,2,3) if your auth URL has no version (ST_AUTH_VERSION). (default: "0") [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     Above this size files will be chunked into a _segments container. (default: "5Gi") [$SWIFT_CHUNK_SIZE]
   --swift-domain value                         User domain - optional (v3 auth) (OS_USER_DOMAIN_NAME) [$SWIFT_DOMAIN]
   --swift-encoding value                       The encoding for the backend. (default: "Slash,InvalidUtf8") [$SWIFT_ENCODING]
   --swift-endpoint-type value                  Endpoint type to choose from the service catalogue (OS_ENDPOINT_TYPE). (default: "public") [$SWIFT_ENDPOINT_TYPE]
   --swift-env-auth value                       Get swift credentials from environment variables in standard OpenStack form. (default: "false") [$SWIFT_ENV_AUTH]
   --swift-key value                            API key or password (OS_PASSWORD). [$SWIFT_KEY]
   --swift-leave-parts-on-error value           If true avoid calling abort upload on a failure. (default: "false") [$SWIFT_LEAVE_PARTS_ON_ERROR]
   --swift-no-chunk value                       Don't chunk files during streaming upload. (default: "false") [$SWIFT_NO_CHUNK]
   --swift-no-large-objects value               Disable support for static and dynamic large objects (default: "false") [$SWIFT_NO_LARGE_OBJECTS]
   --swift-region value                         Region name - optional (OS_REGION_NAME). [$SWIFT_REGION]
   --swift-storage-policy value                 The storage policy to use when creating a new container. [$SWIFT_STORAGE_POLICY]
   --swift-storage-url value                    Storage URL - optional (OS_STORAGE_URL). [$SWIFT_STORAGE_URL]
   --swift-tenant value                         Tenant name - optional for v1 auth, this or tenant_id required otherwise (OS_TENANT_NAME or OS_PROJECT_NAME). [$SWIFT_TENANT]
   --swift-tenant-domain value                  Tenant domain - optional (v3 auth) (OS_PROJECT_DOMAIN_NAME). [$SWIFT_TENANT_DOMAIN]
   --swift-tenant-id value                      Tenant ID - optional for v1 auth, this or tenant required otherwise (OS_TENANT_ID). [$SWIFT_TENANT_ID]
   --swift-user value                           User name to log in (OS_USERNAME). [$SWIFT_USER]
   --swift-user-id value                        User ID to log in - optional - most swift systems use user and leave this blank (v3 auth) (OS_USER_ID). [$SWIFT_USER_ID]

   Options for uptobox

   --uptobox-access-token value  Your access token. [$UPTOBOX_ACCESS_TOKEN]
   --uptobox-encoding value      The encoding for the backend. (default: "Slash,LtGt,DoubleQuote,BackQuote,Del,Ctl,LeftSpace,InvalidUtf8,Dot") [$UPTOBOX_ENCODING]

   Options for webdav

   --webdav-bearer-token value          Bearer token instead of user/pass (e.g. a Macaroon). [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  Command to run to get a bearer token. [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              The encoding for the backend. [$WEBDAV_ENCODING]
   --webdav-headers value               Set HTTP headers for all transactions. [$WEBDAV_HEADERS]
   --webdav-pass value                  Password. [$WEBDAV_PASS]
   --webdav-url value                   URL of http host to connect to. [$WEBDAV_URL]
   --webdav-user value                  User name. [$WEBDAV_USER]
   --webdav-vendor value                Name of the WebDAV site/service/software you are using. [$WEBDAV_VENDOR]

   Options for yandex

   --yandex-auth-url value       Auth server URL. [$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuth Client Id. [$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuth Client Secret. [$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       The encoding for the backend. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    Delete files permanently rather than putting them into the trash. (default: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuth Access Token as a JSON blob. [$YANDEX_TOKEN]
   --yandex-token-url value      Token server url. [$YANDEX_TOKEN_URL]

   Options for zoho

   --zoho-auth-url value       Auth server URL. [$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth Client Id. [$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth Client Secret. [$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       The encoding for the backend. (default: "Del,Ctl,InvalidUtf8") [$ZOHO_ENCODING]
   --zoho-region value         Zoho region to connect to. [$ZOHO_REGION]
   --zoho-token value          OAuth Access Token as a JSON blob. [$ZOHO_TOKEN]
   --zoho-token-url value      Token server url. [$ZOHO_TOKEN_URL]

```
{% endcode %}
