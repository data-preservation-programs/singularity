# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage create azureblob - Microsoft Azure Blob Storage

USAGE:
   singularity storage create azureblob [コマンドオプション] [引数...]

DESCRIPTION:
   --account
      Azure ストレージアカウント名。
      
      使用中の Azure ストレージアカウント名を設定します。
      
      SAS URL やエミュレータを使用する場合は空白にしてください。
      
      空白の場合、env_auth が設定されている場合は、環境変数 `AZURE_STORAGE_ACCOUNT_NAME` から読み込まれます。
      

   --env-auth
      ランタイムから認証情報を読み込みます（環境変数、CLI または MSI）。
      
      詳細については、[認証ドキュメント](/azureblob#authentication)を参照してください。

   --key
      ストレージアカウント共有キー。
      
      空白の場合、SAS URL やエミュレータを使用します。

   --sas-url
      コンテナレベルのアクセスに使用する SAS URL。
      
      Account/Key やエミュレータを使用する場合は空白にしてください。

   --tenant
      サービスプリンシパルのテナント ID。ディレクトリ ID とも呼ばれます。
      
      サービス プリンシパル (クライアントシークレットを使用する)、
      サービス プリンシパル (証明書を使用する)、
      ユーザー (ユーザー名とパスワード) を使用する場合に設定します。
      

   --client-id
      使用中のクライアントの ID。
      
      サービス プリンシパル (クライアントシークレットを使用する)、
      サービス プリンシパル (証明書を使用する)、
      ユーザー (ユーザー名とパスワード) を使用する場合に設定します。
      

   --client-secret
      サービスプリンシパルのクライアントシークレットのいずれか。
      
      サービス プリンシパル (クライアントシークレットを使用する) を使用する場合に設定します。
      

   --client-certificate-path
      私有キーを含む PEM または PKCS12 証明書ファイルへのパス。
      
      サービス プリンシパル (証明書を使用する) を使用する場合に設定します。
      

   --client-certificate-password
      証明書ファイルのパスワード (オプション)。
      
      サービス プリンシパル (証明書を使用する) を使用する場合にオプションで設定します。
      

   --client-send-certificate-chain
      証明書認証時に証明書チェーンを送信するかどうかを指定します。
      
      証明書認証を使用する場合、認証リクエストは x5c ヘッダを含めるためにこの値を true に設定します。

      サービス プリンシパル (証明書を使用する) を使用する場合にオプションで設定します。
      

   --username
      ユーザー名 (通常はメールアドレス)
      
      ユーザー (ユーザー名とパスワード) を使用する場合に設定します。
      

   --password
      ユーザーのパスワード
      
      ユーザー (ユーザー名とパスワード) を使用する場合に設定します。
      

   --service-principal-file
      サービス プリンシパルを使用する場合の資格情報を含むファイルへのパス。
      
      通常は空白です。対話型ログインではなくサービス プリンシパルを使用する場合にのみ必要です。
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      詳細については、["Azure サービス プリンシパルの作成"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli) と ["Blob データへのアクセスのための Azure ロールの割り当て"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli) ページを参照してください。
      
      `service_principal_file` を設定する代わりに、資格情報を直接 rclone 設定ファイルに `client_id`、`tenant`、`client_secret` キーの下に入れておく方が便利かもしれません。
      

   --use-msi
      管理サービス ID を使用して認証する (Azure でのみ動作する)。

      true の場合、SAS トークンやアカウントキーの代わりに、[管理サービス ID](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/) を使用して Azure Storage に認証します。
      
      このプログラムが実行されている VM(SS) にシステム割り当て ID がある場合、デフォルトでそれが使用されます。システム割り当てがなく、ユーザー割り当て ID が正確に1つ存在する場合、デフォルトでユーザー割り当て ID が使用されます。リソースに複数のユーザー割り当て ID がある場合は、msi_object_id、msi_client_id、msi_mi_res_id のいずれか 1 つを明示的に指定する必要があります。

   --msi-object-id
      使用するユーザー割り当て MSI のオブジェクト ID。

      もしくは指定する場合は、msi_client_id または msi_mi_res_id を空白にしてください。

   --msi-client-id
      使用するユーザー割り当て MSI のオブジェクト ID。

      もしくは指定する場合は、msi_object_id または msi_mi_res_id を空白にしてください。

   --msi-mi-res-id
      使用するユーザー割り当て MSI の Azure リソース ID。

      もしくは指定する場合は、msi_client_id または msi_object_id を空白にしてください。

   --use-emulator
      提供された場合に、ローカルのストレージ エミュレータを使用します。

      実際の Azure ストレージ エンドポイントを使用する場合は空白にしてください。

   --endpoint
      サービスのエンドポイント。
      
      通常は空白です。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフサイズ (<= 256 MiB)（非推奨）。

   --chunk-size
      アップロードのチャンクサイズ。
      
      これはメモリに保存され、一度にメモリに格納されるチャンクは"--transfers" * "--azureblob-upload-concurrency" までです。

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      これは同じファイルの複数のチャンクを同時にアップロードする数です。
      
      高速リンクを介して大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に活用していない場合、これを増やすことでトランスファーを高速化することができます。
      
      テストで、アップロード速度はアップロード並行性ごとにほぼ直線的に向上します。たとえば、ギガビットパイプを埋めるためには、これを 64 に上げる必要がある場合があります。ただし、これによりメモリが使用されます。
      
      チャンクはメモリ内に格納され、一度に"--transfers" * "--azureblob-upload-concurrency" 個のチャンクがメモリ内に格納される場合があります。

   --list-chunk
      ブロブリストのサイズ。
      
      これは各リストチャンクで要求される blob の数を設定します。デフォルトは最大値 5000 です。"List blobs" のリクエストの完了には 1 MB あたり 2 分かかります。操作が平均して 1 MB あたり 2 分以上かかっている場合は、タイムアウトします (
      [ソース](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)
      )。これは、タイムアウトを回避するためにブロブアイテムの数を制限するために使用できます。

   --access-tier
      ブロブのアクセスティア: hot、cool、または archive。
      
      アーカイブされたブロブは、アクセスティアを hot または cool に設定して復元できます。デフォルトのアクセスティアを使用する場合は空白にしてください。
      
      "アクセスティア" が指定されていない場合、rclone はアクセスティアを適用しません。rclone は、オブジェクトが変更されていない場合には "アクセスティア" への "Set Tier" 操作を実行しません。リモートのブロブが "archive tier" の場合、リモートからのデータ転送操作を行おうとすると許可されません。ユーザーはまず、ブロブを "Hot" または "Cool" にティアリングしてから、復元する必要があります。

   --archive-tier-delete
      上書きする前にアーカイブティアのブロブを削除します。
      
      アーカイブティアのブロブは更新できません。したがって、このフラグを設定しない場合、アーカイブティアのブロブを更新しようとすると、rclone はエラーを返します:
      
          can't update archive tier blob without --azureblob-archive-tier-delete
      
      このフラグが設定されている場合、アーカイブティアのブロブを上書きする前に、既存のブロブを削除してから置き換えをアップロードします。このため、アップロードが失敗した場合にデータの損失の可能性があり、また、アーカイブティアのブロブを早期に削除することで追加の料金が発生する場合があります。
      

   --disable-checksum
      オブジェクトメタデータに MD5 チェックサムを保存しません。
      
      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には時間がかかることがあります。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファ (たとえば、マルチパート) を必要とするアップロードは、アロケーションにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールで mmap バッファを使用するかどうか。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[エンコーディングセクションの概要](/overview/#encoding) を参照してください。

   --public-access
      コンテナのパブリックアクセスレベル: blob または container。

      例:
         | <unset>   | コンテナとそのブロブは、承認されたリクエストでのみアクセスできます。
         |           | デフォルトの値です。
         | blob      | このコンテナ内のブロブデータは匿名リクエストを経由して読み取ることができます。
         | container | コンテナとブロブデータに完全なパブリック読み取りアクセスを許可します。

   --no-check-container
      コンテナの存在を確認せず、作成しようとしません。
      
      コンテナが既に存在する場合、rclone が行うトランザクションの数を最小限にするために役立つ場合があります。
      

   --no-head-object
      GET する際に HEAD を実行しません。


OPTIONS:
   --account value                      Azure ストレージアカウント名。 [$ACCOUNT]
   --client-certificate-password value  証明書ファイルのパスワード (オプション)。 [$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      私有キーを含む PEM または PKCS12 証明書ファイルへのパス。 [$CLIENT_CERTIFICATE_PATH]
   --client-id value                    使用中のクライアントの ID。 [$CLIENT_ID]
   --client-secret value                サービスプリンシパルのクライアントシークレットのいずれか [$CLIENT_SECRET]
   --env-auth                           ランタイムから認証情報を読み込みます（環境変数、CLI または MSI）（デフォルト: false） [$ENV_AUTH]
   --help, -h                           ヘルプを表示します
   --key value                          ストレージアカウント共有キー。 [$KEY]
   --sas-url value                      コンテナレベルのアクセスに使用する SAS URL。 [$SAS_URL]
   --tenant value                       サービスプリンシパルのテナント ID。ディレクトリ ID とも呼ばれます。 [$TENANT]

   Advanced

   --access-tier value              ブロブのアクセスティア: hot、cool、または archive。 [$ACCESS_TIER]
   --archive-tier-delete            上書きする前にアーカイブティアのブロブを削除します (デフォルト: false) [$ARCHIVE_TIER_DELETE]
   --chunk-size value               アップロードのチャンクサイズ (デフォルト: "4Mi") [$CHUNK_SIZE]
   --client-send-certificate-chain  証明書認証時に証明書チェーンを送信するかどうか。 (デフォルト: false) [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               オブジェクトメタデータに MD5 チェックサムを保存しません (デフォルト: false) [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング (デフォルト: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$ENCODING]
   --endpoint value                 サービスのエンドポイント。 [$ENDPOINT]
   --list-chunk value               ブロブリストのサイズ (デフォルト: 5000) [$LIST_CHUNK]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールで mmap バッファを使用するかどうか (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            使用するユーザー割り当て MSI のオブジェクト ID。 [$MSI_CLIENT_ID]
   --msi-mi-res-id value            使用するユーザー割り当て MSI の Azure リソース ID。 [$MSI_MI_RES_ID]
   --msi-object-id value            使用するユーザー割り当て MSI のオブジェクト ID。 [$MSI_OBJECT_ID]
   --no-check-container             コンテナの存在を確認せず、作成しようとしません (デフォルト: false) [$NO_CHECK_CONTAINER]
   --no-head-object                 GET する際に HEAD を実行しません (デフォルト: false) [$NO_HEAD_OBJECT]
   --password value                 ユーザーのパスワード [$PASSWORD]
   --public-access value            コンテナのパブリックアクセスレベル: blob または container。 [$PUBLIC_ACCESS]
   --service-principal-file value   サービス プリンシパルを使用する場合の資格情報を含むファイルへのパス。 [$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数 (デフォルト: 16) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフサイズ (<= 256 MiB) (非推奨) [$UPLOAD_CUTOFF]
   --use-emulator                   提供された場合に、ローカルのストレージ エミュレータを使用する (デフォルト: false) [$USE_EMULATOR]
   --use-msi                        管理サービス ID を使用して認証する (Azure でのみ動作する) (デフォルト: false) [$USE_MSI]
   --username value                 ユーザー名 (通常はメールアドレス) [$USERNAME]

   General

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}