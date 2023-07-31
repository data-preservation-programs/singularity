# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
名前:
   singularity datasource add azureblob - Microsoft Azure Blob Storage

使用法:
   singularity datasource add azureblob [コマンドオプション] <データセット名> <ソースパス>

説明:
   --azureblob-access-tier
      ブロブのアクセスレベル: hot, cool, archive のどれか。

      アーカイブされたブロブは、アクセスレベルを hot や cool に設定することで復元できます。
      アクセスレベルが指定されていない場合は、アカウントレベルでデフォルトのアクセスレベルが設定されます。
      
      "アクセスレベル" が指定されていない場合、rclone はエクスポート先に新しい "アクセスレベル" を追加しません。
      rclone は、アップロード時にブロブに "アクセスレベル" を設定する "Set Tier" の操作を実行します。
      オブジェクトが変更されていない場合、新しい "アクセスレベル" は影響を与えません。
      リモートで "アーカイブレベル" にあるブロブからのデータ転送操作は許可されません。
      ユーザーは、まずブロブを "Hot" または "Cool" のアクセスレベルに移行してから、データ転送操作を実行する必要があります。

   --azureblob-account
      Azure ストレージアカウント名。

      使用している Azure ストレージアカウント名を設定します。

      SAS URL またはエミュレータを使用する場合は空白のままにしますが、それ以外の場合は設定する必要があります。

      この項目が空白であり、env_auth も設定されている場合、環境変数 `AZURE_STORAGE_ACCOUNT_NAME` から読み取ります。

   --azureblob-archive-tier-delete
      アーカイブレベルのブロブを上書きする前にアーカイブレベルのブロブを削除します。

      アーカイブレベルのブロブは更新できないため、このフラグを指定せずにアーカイブレベルのブロブを更新しようとすると、
      rclone は次のエラーメッセージを生成します。

          --azureblob-archive-tier-delete を指定しないとアーカイブレベルのブロブを更新できません。

      このフラグを設定すると、アーカイブレベルのブロブを上書きする前に既存のブロブを削除し、
      上書きするブロブをアップロードします。アップロードに失敗した場合、データの損失が発生する可能性があり、
      アーカイブレベルのブロブの削除は、早期に削除すると費用が発生する可能性があるため、コストもかかる場合もあります。

   --azureblob-chunk-size
      アップロードするチャンクのサイズ。

      これはメモリに格納され、メモリ内には最大で "--transfers" * "--azureblob-upload-concurrency" 個の
      チャンクが格納される可能性があります。

   --azureblob-client-certificate-password
      クライアント証明書ファイルのパスワード (省略可能)。

      クライアント証明書を使用する場合にオプションで設定します。
      - Service principal with certificate
      
      また、証明書にパスワードが設定されている場合にも、これを設定できます。

   --azureblob-client-certificate-path
      PEM または PKCS12 証明書ファイルのパス、プライベートキーを含んでいます。

      クライアント証明書を使用する場合は、これを設定します。
      - Service principal with certificate

   --azureblob-client-id
      使用中のクライアントの ID。

      これを設定すると、次のものを使用できます。
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password

   --azureblob-client-secret
      サービスプリンシパルのクライアントシークレットの1つ。

      これを設定すると、次のものを使用できます。
      - Service principal with client secret

   --azureblob-client-send-certificate-chain
      証明書認証時に証明書チェーンを送信します。

      認証リクエストに x5c ヘッダを含めるかどうかを指定します。
      true に設定すると、認証リクエストに x5c ヘッダが含まれます。
      
      オプションでこれを設定すると、証明書を使用する場合に便利です。
      - Service principal with certificate

   --azureblob-disable-checksum
      オブジェクトのメタデータに MD5 チェックサムを保存しない。

      通常、rclone はアップロード前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加します。
      これはデータの整合性チェックには最適ですが、大きなファイルのアップロードでは長い遅延が発生する可能性があります。

   --azureblob-encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --azureblob-endpoint
      サービスのエンドポイント。

      通常は空白のままにします。

   --azureblob-env-auth
      ランタイムから認証情報を読み取ります (環境変数、CLI、または MSI)。

      詳細については、[認証ドキュメント](/azureblob#authentication)を参照してください。

   --azureblob-key
      ストレージアカウントの共有キー。

      SAS URL またはエミュレータを使用する場合は空白のままにします。

   --azureblob-list-chunk
      ブロブリストのサイズ。

      これは一度のリスト取得で要求されるブロブの数を指定します。デフォルトは最大の 5000 です。
      "リストブロブ" リクエストは、完了までにメガバイト当たり2分のリクエスト時間を許可されます。
      平均して、2分のリクエスト時間当たりの処理が長時間になる場合は、タイムアウトします。
      ([ソース](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval))
      タイムアウトを回避するために、返すブロブアイテムの数を制限するために使用できます。

   --azureblob-memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。

      追加のバッファ (マルチパートなど) を必要とするアップロードは、メモリプールを使用して割り当てられます。
      このオプションは、使用されていないバッファがプールから削除される頻度を制御します。

   --azureblob-memory-pool-use-mmap
      インターナルメモリプールで mmap バッファを使用するかどうか。

   --azureblob-msi-client-id
      使用するユーザに割り当てられた MSI のオブジェクト ID。

      msi_object_id または msi_mi_res_id が指定されている場合は空白のままにします。

   --azureblob-msi-mi-res-id
      使用するユーザに割り当てられた MSI の Azure リソース ID。

      msi_client_id または msi_object_id が指定されている場合は空白のままにします。

   --azureblob-msi-object-id
      使用するユーザに割り当てられた MSI のオブジェクト ID。

      msi_client_id または msi_mi_res_id が指定されている場合は空白のままにします。

   --azureblob-no-check-container
      コンテナが存在するかどうかを確認せず、作成しないようにします。

      これは、コンテナが既に存在することを知っている場合に、rclone が実行するトランザクション数を最小限にするために便利です。

   --azureblob-no-head-object
      GET 操作を実行する前に HEAD 操作を行わないようにします。

   --azureblob-password
      ユーザのパスワード

      これを設定すると、次のものを使用できます。
      - User with username and password

   --azureblob-public-access
      コンテナの公開アクセスレベル: blob または container。

      例:
         | <unset>   | コンテナとそのブロブは認証されたリクエストのみアクセスできます。
                     | デフォルト値です。
         | blob      | このコンテナ内の Blob データは匿名リクエストで読み取ることができます。
         | container | コンテナと Blob データの完全なパブリック読み取りアクセスを許可します。

   --azureblob-sas-url
      コンテナレベルのアクセスのための SAS URL。

      アカウント/キーまたはエミュレータを使用する場合は空白のままにします。

   --azureblob-service-principal-file
      サービスプリンシパルを使用するための資格情報を含むファイルへのパス。

      通常は空白のままにします。対話形式のログインではなく、サービスプリンシパルを使用する場合にのみ必要です。

          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      詳細については、["Create an Azure service principal"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) および
      ["Assign an Azure role for access to blob data"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli) ページを参照してください。
      
      `client_id`、`tenant`、および `client_secret` のキーではなく、資格情報を直接 rclone の構成ファイルに入れる方が便利かもしれません。
      `service_principal_file` を設定する代わりに、`client_id`、`tenant`、`client_secret` のキーに直接資格情報を設定できます。

   --azureblob-tenant
      サービスプリンシパルのテナントの ID。ディレクトリ ID とも呼ばれます。

      これを設定すると、次のものを使用できます。
      - Service principal with client secret
      - Service principal with certificate
      - User with username and password

   --azureblob-upload-concurrency
      マルチパートアップロードの並行性。

      これは同じファイルのチャンクの数を指定します。

      高速リンクで小数の大きなファイルをアップロードし、これらのアップロードが帯域幅を完全に使用しきれない場合は、
      この値を増やすと転送が高速化する場合があります。

      テストでは、アップロード速度はアップロードの並行性にほぼ比例して増加します。
      たとえば、ギガビットパイプを埋めるには、これを 64 に増やす必要があるかもしれません。
      ただし、これによりメモリをより多く使用します。

      チャンクはメモリに格納され、最大で "--transfers" * "--azureblob-upload-concurrency" 個の
      チャンクが一度にメモリ内に格納される可能性があります。

   --azureblob-upload-cutoff
      チャンク化アップロードに切り替える基準 (<= 256 MiB) (非推奨)。

   --azureblob-use-emulator
      提供された場合にローカルストレージエミュレータを使用します。

      実際の Azure ストレージエンドポイントを使用する場合は空白のままにします。

   --azureblob-use-msi
      管理サービス ID を使用して認証します (Azure のみで動作します)。

      true の場合、[管理サービス ID](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)
      を使用して Azure Storage に認証し、SAS トークンまたはアカウントキーを使用しません。

      プログラムが実行されている VM (SS) にシステム割り当て ID がある場合、デフォルトでそれが使用されます。
      システム割り当てが存在しない場合で、ユーザ割り当て ID が 1 つだけある場合、デフォルトでユーザ割り当て ID が使用されます。
      ユーザ割り当て ID が複数ある場合、msi_object_id、msi_client_id、または msi_mi_res_id のいずれかを
      明示的に指定する必要があります。

   --azureblob-username
      ユーザ名 (通常メールアドレス)

      これを設定すると、次のものを使用できます。
      - User with username and password

オプション:
   --help、 -h  ヘルプを表示

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルを CAR ファイルにエクスポート後に削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   azureblob のオプション

   --azureblob-access-tier value                    ブロブのアクセスレベル: hot, cool, archive のどれか。 [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure ストレージアカウント名。 [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            上書きする前にアーカイブレベルのブロブを削除します。 (デフォルト: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     アップロードするチャンクのサイズ。 (デフォルト: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    クライアント証明書ファイルのパスワード (省略可能)。 [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        PEM または PKCS12 証明書ファイルのパス、プライベートキーを含んでいます。 [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      使用中のクライアントの ID。 [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  サービスプリンシパルのクライアントシークレットの1つ [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  証明書認証時に証明書チェーンを送信します。 (デフォルト: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               オブジェクトのメタデータに MD5 チェックサムを保存しません。 (デフォルト: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       バックエンドのエンコーディング。 (デフォルト: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       サービスのエンドポイント。 [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       ランタイムから認証情報を読み取ります (環境変数、CLI、または MSI)。 (デフォルト: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            ストレージアカウントの共有キー。 [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     ブロブリストのサイズ。 (デフォルト: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         内部メモリバッファプールをフラッシュする頻度。 (デフォルト: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           インターナルメモリプールで mmap バッファを使用するかどうか。 (デフォルト: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  使用するユーザに割り当てられた MSI のオブジェクト ID。 [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  使用するユーザに割り当てられた MSI の Azure リソース ID。 [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  使用するユーザに割り当てられた MSI のオブジェクト ID。 [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             コンテナが存在するか確認し、作成しないようにします。 (デフォルト: "false") [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 GET 操作を実行する前に HEAD 操作を行わないようにします。 (デフォルト: "false") [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       ユーザのパスワード [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  コンテナの公開アクセスレベル: blob または container。 [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        コンテナレベルのアクセスのための SAS URL。 [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         サービスプリンシパルを使用するための資格情報を含むファイルへのパス。 [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         サービスプリンシパルのテナントの ID。ディレクトリ ID とも呼ばれます。 [$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             マルチパートアップロードの並行性。 (デフォルト: "16") [$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  チャンク化アップロードに切り替える基準 (<= 256 MiB) (非推奨) [$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulator value                   提供された場合にローカルストレージエミュレータを使用します。 (デフォルト: "false") [$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        管理サービス ID を使用して認証します (Azure のみで動作します)。 (デフォルト: "false") [$AZUREBLOB_USE_MSI]
   --azureblob-username value                       ユーザ名 (通常メールアドレス) [$AZUREBLOB_USERNAME]

```
{% endcode %}