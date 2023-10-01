# CARファイルのダウンロードを提供するため、リモートメタデータAPIに接続するHTTPサーバー

{% code fullWidth="true" %}
```
名前:
   singularity run download-server - CARファイルのダウンロードを提供するため、リモートメタデータAPIに接続するHTTPサーバー

使用方法:
   singularity run download-server [コマンドオプション] [引数...]

説明:
   使い方の例:
     singularity run download-server --metadata-api "http://remote-metadata-api:7777" --bind "127.0.0.1:8888"

オプション:
   --help, -h  ヘルプの表示

   1Fichier

   --fichier-api-key value          1FichierのAPIキー。[https://1fichier.com/console/params.pl]で取得できます。 [$FICHIER_API_KEY]
   --fichier-file-password value    共有ファイルのダウンロード時にパスワードが必要な場合、このパラメータを追加します。[$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  パスワードで保護された共有フォルダ内のファイルの一覧を取得する場合、このパラメータを追加します。[$FICHIER_FOLDER_PASSWORD]

   Akamai NetStorage

   --netstorage-secret value  NetStorageアカウントのシークレット/G2Oキーを設定します。[$NETSTORAGE_SECRET]

   Amazon Drive

   --acd-client-secret value  OAuthクライアントシークレット。[$ACD_CLIENT_SECRET]
   --acd-token value          JSON形式のOAuthアクセストークン。[$ACD_TOKEN]
   --acd-token-url value      トークンサーバーのURL。[$ACD_TOKEN_URL]

   AWS、Alibaba、Ceph、China Mobile、Cloudflare、ArvanCloud、DigitalOcean、Dreamhost、Huawei OBS、IBM COS、IDrive e2、IONOS Cloud、Liara、Lyve Cloud、Minio、Netease、RackCorp、Scaleway、SeaweedFS、StackPath、Storj、Tencent COS、Qiniu、Wasabiを含むAmazon S3互換ストレージプロバイダ

   --s3-access-key-id value            AWSのアクセスキーID。[$S3_ACCESS_KEY_ID]
   --s3-secret-access-key value        AWSのシークレットアクセスキー（パスワード）。[$S3_SECRET_ACCESS_KEY]
   --s3-session-token value            AWSのセッショントークン。[$S3_SESSION_TOKEN]
   --s3-sse-customer-key value         SSE-Cを使用する場合、暗号化/復号化に使用する秘密暗号キーを指定します。[$S3_SSE_CUSTOMER_KEY]
   --s3-sse-customer-key-base64 value  SSE-Cを使用する場合、暗号化/復号化に使用する秘密暗号キーをBase64形式で指定します。[$S3_SSE_CUSTOMER_KEY_BASE64]
   --s3-sse-customer-key-md5 value     SSE-Cを使用する場合、秘密暗号キーのMD5チェックサムを指定できます（オプション）。[$S3_SSE_CUSTOMER_KEY_MD5]
   --s3-sse-kms-key-id value           KMS IDを使用する場合、キーのARNを指定する必要があります。[$S3_SSE_KMS_KEY_ID]

   Backblaze B2

   --b2-key value  アプリケーションキー。[$B2_KEY]

   Box

   --box-access-token value   Box Appプライマリのアクセストークン。[$BOX_ACCESS_TOKEN]
   --box-client-secret value  OAuthクライアントシークレット。[$BOX_CLIENT_SECRET]
   --box-token value          JSON形式のOAuthアクセストークン。[$BOX_TOKEN]
   --box-token-url value      トークンサーバーのURL。[$BOX_TOKEN_URL]

   クライアント設定

   --client-ca-cert value                           サーバーを検証するために使用するCA証明書へのパス。削除するには空の文字列を使用します。
   --client-cert value                              相互TLS認証のためのクライアントSSL証明書（PEM形式）へのパス。削除するには空の文字列を使用します。
   --client-connect-timeout value                   HTTPクライアントの接続タイムアウト（デフォルト: 1分）
   --client-expect-continue-timeout value           HTTPのexpect / 100-continueを使用する場合のタイムアウト（デフォルト: 1秒）
   --client-header value [ --client-header value ]  すべてのトランザクションに対してHTTPヘッダを設定します（例: キー=値）。これにより、既存のヘッダ値が置き換えられます。ヘッダを削除するには、--http-header "key="を使用します。すべてのヘッダを削除するには、--http-header ""を使用します。
   --client-insecure-skip-verify                    サーバーのSSL証明書を検証しない（安全ではない）（デフォルト: false）
   --client-key value                               相互TLS認証のためのクライアントSSL秘密鍵（PEM形式）へのパス。削除するには空の文字列を使用します。
   --client-no-gzip                                 Accept-Encoding: gzipを設定しない（デフォルト: false）
   --client-scan-concurrency value                  データソースをスキャンする際の並行リストリクエストの最大数（デフォルト: 1）
   --client-timeout value                           IOアイドルタイムアウト（デフォルト: 5分）
   --client-use-server-mod-time                     可能な場合はサーバーの変更日時を使用する（デフォルト: false）
   --client-user-agent value                        指定した文字列にユーザーエージェントを設定します。削除するには空の文字列を使用します（デフォルト: rclone/v1.62.2-DEV）

   Dropbox

   --dropbox-client-secret value  OAuthクライアントシークレット。[$DROPBOX_CLIENT_SECRET]
   --dropbox-token value          JSON形式のOAuthアクセストークン。[$DROPBOX_TOKEN]
   --dropbox-token-url value      トークンサーバーのURL。[$DROPBOX_TOKEN_URL]

   Enterprise File Fabric

   --filefabric-permanent-token value  永続認証トークン。[$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-token value            セッショントークン。[$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     トークンの有効期限。[$FILEFABRIC_TOKEN_EXPIRY]

   FTP

   --ftp-ask-password  必要に応じてFTPパスワードの入力を許可します。 (デフォルト: false) [$FTP_ASK_PASSWORD]
   --ftp-pass value    FTPパスワード。[$FTP_PASS]

   一般設定

   --bind value          HTTPサーバーをバインドするアドレス（デフォルト: "127.0.0.1:8888"）
   --metadata-api value  メタデータAPIのURL（デフォルト: "http://127.0.0.1:7777"）

   Google Cloud Storage（これはGoogle Driveではありません）

   --gcs-client-secret value  OAuthクライアントシークレット。[$GCS_CLIENT_SECRET]
   --gcs-token value          JSON形式のOAuthアクセストークン。[$GCS_TOKEN]
   --gcs-token-url value      トークンサーバーのURL。[$GCS_TOKEN_URL]

   Google Drive

   --drive-client-secret value  OAuthクライアントシークレット。[$DRIVE_CLIENT_SECRET]
   --drive-resource-key value   リンク共有ファイルにアクセスするためのリソースキー。[$DRIVE_RESOURCE_KEY]
   --drive-token value          JSON形式のOAuthアクセストークン。[$DRIVE_TOKEN]
   --drive-token-url value      トークンサーバーのURL。[$DRIVE_TOKEN_URL]

   Google Photos

   --gphotos-client-secret value  OAuthクライアントシークレット。[$GPHOTOS_CLIENT_SECRET]
   --gphotos-token value          JSON形式のOAuthアクセストークン。[$GPHOTOS_TOKEN]
   --gphotos-token-url value      トークンサーバーのURL。[$GPHOTOS_TOKEN_URL]

   HiDrive

   --hidrive-client-secret value  OAuthクライアントシークレット。[$HIDRIVE_CLIENT_SECRET]
   --hidrive-token value          JSON形式のOAuthアクセストークン。[$HIDRIVE_TOKEN]
   --hidrive-token-url value      トークンサーバーのURL。[$HIDRIVE_TOKEN_URL]

   Internet Archive

   --internetarchive-access-key-id value      IAS3アクセスキー。[$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-secret-access-key value  IAS3シークレットキー（パスワード）。[$INTERNETARCHIVE_SECRET_ACCESS_KEY]

   Koofr、Digi Storage および他のKoofr互換ストレージプロバイダ

   --koofr-password value  rcloneのパスワード（[https://storage.rcs-rds.ro/app/admin/preferences/password]で生成できます）。[$KOOFR_PASSWORD]

   Mail.ru Cloud

   --mailru-pass value  パスワード。[$MAILRU_PASS]

   Mega

   --mega-pass value  パスワード。[$MEGA_PASS]

   Microsoft Azure Blob Storage

   --azureblob-client-certificate-password value  オプションの証明書ファイルのパスワード。[$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-secret value                サービスプリンシパルのクライアントシークレットのいずれか。[$AZUREBLOB_CLIENT_SECRET]
   --azureblob-key value                          ストレージアカウントの共有キー。[$AZUREBLOB_KEY]
   --azureblob-password value                     ユーザーのパスワード。[$AZUREBLOB_PASSWORD]

   Microsoft OneDrive

   --onedrive-client-secret value  OAuthクライアントシークレット。[$ONEDRIVE_CLIENT_SECRET]
   --onedrive-link-password value  リンクコマンドで作成されたリンクのパスワードを設定します。[$ONEDRIVE_LINK_PASSWORD]
   --onedrive-token value          JSON形式のOAuthアクセストークン。[$ONEDRIVE_TOKEN]
   --onedrive-token-url value      トークンサーバーのURL。[$ONEDRIVE_TOKEN_URL]

   OpenDrive

   --opendrive-password value  パスワード。[$OPENDRIVE_PASSWORD]

   OpenStack Swift（Rackspace Cloud Files、Memset Memstore、OVH）

   --swift-application-credential-secret value  アプリケーション資格情報のシークレット（OS_APPLICATION_CREDENTIAL_SECRET）。[$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth-token value                     別の認証に必要な認証トークン（オプション）（OS_AUTH_TOKEN）。[$SWIFT_AUTH_TOKEN]
   --swift-key value                            APIキーまたはパスワード（OS_PASSWORD）。[$SWIFT_KEY]

   Oracle Cloud Infrastructure Object Storage

   --oos-sse-customer-key value         SSE-Cを使用する場合、オプションのヘッダーで256ビットの暗号化キーを指定します。[$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    SSE-Cを使用する場合、ベース64でエンコードされたAES-256暗号化キーの文字列を含むファイル。[$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのベース64エンコードされたSHA256ハッシュを指定できます。[$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           マスターキーを使用する場合、キーのARNを指定する必要があります（OOS_SSE_KMS_KEY_ID）。

   Pcloud

   --pcloud-client-secret value  OAuthクライアントシークレット。[$PCLOUD_CLIENT_SECRET]
   --pcloud-password value       Pcloudのパスワード。[$PCLOUD_PASSWORD]
   --pcloud-token value          JSON形式のOAuthアクセストークン。[$PCLOUD_TOKEN]
   --pcloud-token-url value      トークンサーバーのURL。[$PCLOUD_TOKEN_URL]

   QingCloud Object Storage

   --qingstor-access-key-id value      QingStorのアクセスキーID。[$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-secret-access-key value  QingStorのシークレットアクセスキー（パスワード）。[$QINGSTOR_SECRET_ACCESS_KEY]

   リトライ戦略

   --client-low-level-retries value  低レベルのクライアントエラーの最大リトライ回数（デフォルト: 10）
   --client-retry-backoff value      IO読み取りエラーのリトライ時の一定のディレイバックオフ（デフォルト: 1秒）
   --client-retry-backoff-exp value  IO読み取りエラーのリトライ時の指数関数的なディレイバックオフ（デフォルト: 1.0）
   --client-retry-delay value        IO読み取りエラーのリトライ前の初期ディレイ（デフォルト: 1秒）
   --client-retry-max value          IO読み取りエラーの最大リトライ回数（デフォルト: 10）
   --client-skip-inaccessible        開かれることのないファイルをスキップする（デフォルト: false）

   SMB / CIFS

   --smb-pass value  SMBのパスワード。[$SMB_PASS]

   SSH/SFTP

   --sftp-ask-password         必要に応じてSFTPパスワードの入力を許可します。 (デフォルト: false) [$SFTP_ASK_PASSWORD]
   --sftp-key-exchange value   優先度順に並べたスペース区切りの鍵交換アルゴリズムのリスト。[$SFTP_KEY_EXCHANGE]
   --sftp-key-file value       PEM形式の秘密鍵ファイルへのパス。[$SFTP_KEY_FILE]
   --sftp-key-file-pass value  PEM形式の秘密鍵ファイルを復号するためのパスフレーズ。[$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value        生のPEM形式の秘密鍵。[$SFTP_KEY_PEM]
   --sftp-key-use-agent        ssh-agentの使用を強制します。 (デフォルト: false) [$SFTP_KEY_USE_AGENT]
   --sftp-pass value           SSHパスワード。ssh-agentを使用する場合は空白のままにします。 [$SFTP_PASS]
   --sftp-pubkey-file value    パブリックキーファイルへのオプションのパス。[$SFTP_PUBKEY_FILE]

   Sia Decentralized Cloud

   --sia-api-password value  SiaデーモンAPIパスワード。[$SIA_API_PASSWORD]

   Storj Decentralized Cloud Storage

   --storj-api-key value     APIキー。[$STORJ_API_KEY]
   --storj-passphrase value  暗号化のパスフレーズ。[$STORJ_PASSPHRASE]

   Sugarsync

   --sugarsync-access-key-id value       Sugarsync Access Key ID. [$SUGARSYNC_ACCESS_KEY_ID]
   --sugarsync-private-access-key value  Sugarsync Private Access Key. [$SUGARSYNC_PRIVATE_ACCESS_KEY]
   --sugarsync-refresh-token value       Sugarsync refresh token. [$SUGARSYNC_REFRESH_TOKEN]

   Uptobox

   --uptobox-access-token value  アクセストークン。[$UPTOBOX_ACCESS_TOKEN]

   WebDAV

   --webdav-bearer-token value          ユーザー/パスワードの代わりにベアラトークン（Macaroonなど）を使用します。[$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  ベアラトークンを取得するために実行するコマンド。[$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-pass value                  パスワード。[$WEBDAV_PASS]

   Yandex Disk

   --yandex-client-secret value  OAuthクライアントシークレット。[$YANDEX_CLIENT_SECRET]
   --yandex-token value          JSON形式のOAuthアクセストークン。[$YANDEX_TOKEN]
   --yandex-token-url value      トークンサーバーのURL。[$YANDEX_TOKEN_URL]

   Zoho

   --zoho-client-secret value  OAuthクライアントシークレット。[$ZOHO_CLIENT_SECRET]
   --zoho-token value          JSON形式のOAuthアクセストークン。[$ZOHO_TOKEN]
   --zoho-token-url value      トークンサーバーのURL。[$ZOHO_TOKEN_URL]

   premiumize.me

   --premiumizeme-api-key value  APIキー。[$PREMIUMIZEME_API_KEY]

   seafile

   --seafile-auth-token value   認証トークン。[$SEAFILE_AUTH_TOKEN]
   --seafile-library-key value  暗号化されたライブラリのパスワード（オプション）。[$SEAFILE_LIBRARY_KEY]
   --seafile-pass value         パスワード。[$SEAFILE_PASS]

```
{% endcode %}