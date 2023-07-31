# ソースの設定オプションを更新します

{% code fullWidth="true" %}
```
NAME:
   singularity datasource update - ソースの設定オプションを更新します

USAGE:
   singularity datasource update [command options] <source_id>

OPTIONS:
   --help, -h  ヘルプを表示します。

   Data Preparation Options

   --delete-after-export    CAR ファイルにエクスポート後、データセットのファイルを削除します（危険）。  (default: false)
   --rescan-interval value  前回のスキャンからこのインターバルが経過した場合、ソースディレクトリを自動的に再スキャンします（デフォルト：無効）
   --scanning-state value   スキャンの初期状態を設定します (デフォルト: ready)

   acd のオプション

   --acd-auth-url value            認証サーバーの URL。 [$ACD_AUTH_URL]
   --acd-client-id value           OAuth クライアント ID。 [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth クライアントシークレット。 [$ACD_CLIENT_SECRET]
   --acd-encoding value            バックエンドのエンコーディング（デフォルト：「スラッシュ、無効な UTF-8、ドット」） [$ACD_ENCODING]
   --acd-templink-threshold value  このサイズ以上のファイルは一時リンクを使用してダウンロードされます（デフォルト："9Gi"）[$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               OAuth アクセス トークンを JSON ライチョンとして指定します。 [$ACD_TOKEN]
   --acd-token-url value           トークンサーバーのURL。 [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完全なアップロードに失敗した後、GiB 単位で追加の待機時間を設定します（デフォルト："3m0s"）[$ACD_UPLOAD_WAIT_PER_GB]

   azureblob のオプション

   --azureblob-access-tier value                    Blob のアクセス層: hot、cool、または archive。[$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure ストレージ アカウントの名前。[$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            上書き前にアーカイブティアのブロブを削除します（デフォルト：「false」）[$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     アップロードチャンクサイズ（デフォルト："4Mi"）[$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    証明書ファイルのパスワード（省略可能）。[$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        プライベートキーを含む PEM または PKCS12 証明書ファイルのパス。[$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      使用するクライアントの ID。[$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  サービスプリンシパルのクライアントシークレットのいずれか [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  証明書認証を使用する場合に証明書チェーンを送信する (デフォルト："false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               オブジェクトメタデータに MD5 チェックサムを保存しない（デフォルト："false"）[$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       バックエンドのエンコーディング（デフォルト："Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8"）[$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       サービスのエンドポイント。[$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       ランタイム（環境変数、CLI、または MSI）からの認証情報の読み取り（デフォルト："false"）[$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            ストレージアカウントの共有キー。[$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     ブロブリストのサイズ（デフォルト："5000"）[$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         内部メモリバッファプールをフラッシュする頻度（デフォルト："1m0s"）[$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           内部メモリプールで mmap バッファを使用するかどうか（デフォルト："false"）[$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  使用するユーザー割り当て MSI