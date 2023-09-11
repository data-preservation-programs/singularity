# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage update azureblob - マイクロソフトの Azure Blob ストレージ

使用法:
   singularity storage update azureblob [コマンドオプション] <名前|ID>

説明:
   --account
      Azure ストレージアカウント名。

      このオプションを使用して、使用中の Azure ストレージアカウント名を設定します。

      SAS URL やエミュレーターを使用する場合は空白のままで結構ですが、
      それ以外の場合は設定する必要があります。

      空白であり、かつ env_auth が設定されている場合は、
      環境変数 `AZURE_STORAGE_ACCOUNT_NAME` から読み取られる場合があります。

   --env-auth
      実行時に資格情報を読み込みます（環境変数、CLI 、または MSI ）。

      詳細は、[認証ドキュメント](/azureblob#authentication)を参照してください。

   --key
      ストレージアカウントの共有キー。

      SAS URL やエミュレーターを使用する場合は空白のままで結構です。

   --sas-url
      コンテナレベルのアクセス専用の SAS URL。

      アカウント/キーまたはエミュレーターを使用する場合は、空白のままで結構です。

   --tenant
      サービスプリンシパルのテナントの ID。またはディレクトリ ID とも呼ばれます。

      以下の場合に設定します。
      - クライアントシークレットを使用するサービスプリンシパル
      - 証明書を使用するサービスプリンシパル
      - ユーザー名とパスワードを使用するユーザー

   --client-id
      クライアントの ID。

      以下の場合に設定します。
      - クライアントシークレットを使用するサービスプリンシパル
      - 証明書を使用するサービスプリンシパル
      - ユーザー名とパスワードを使用するユーザー

   --client-secret
      サービスプリンシパルのクライアントシークレットの 1 つ。

      以下の場合に設定します。
      - クライアントシークレットを使用するサービスプリンシパル

   --client-certificate-path
      PEM または PKCS12 証明書ファイルのパス（プライベートキーも含む）。

      以下の場合に設定します。
      - 証明書を使用するサービスプリンシパル

   --client-certificate-password
      証明書ファイルのパスワード（オプション）。

      以下の場合に設定します。
      - 証明書を使用するサービスプリンシパル

      証明書にパスワードがある場合にのみ、設定してください。

   --client-send-certificate-chain
      証明書認証時に証明書チェーンを送信するかどうか。

      証明書に基づくサブジェクト名/発行者認証をサポートするための x5c ヘッダを認証リクエスト
      に含めるかどうかを指定します。true に設定すると、認証リクエストに x5c ヘッダが含まれます。

      以下の場合に設定します。
      - 証明書を使用するサービスプリンシパル

   --username
      ユーザー名（通常はメールアドレス）

      以下の場合に設定します。
      - ユーザー名とパスワードを使用するユーザー

   --password
      ユーザーのパスワード

      以下の場合に設定します。
      - ユーザー名とパスワードを使用するユーザー

   --service-principal-file
      サービスプリンシパルに使用する資格情報が含まれるファイルのパス。

      通常は空白のままです。対話的なログインの代わりにサービスプリンシパルを使用する場合にのみ必要です。

          $ az ad sp create-for-rbac --name "<名前>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<サブスクリプション>/resourceGroups/<リソースグループ>/providers/Microsoft.Storage/storageAccounts/<ストレージアカウント>/blobServices/default/containers/<コンテナ>" \
            > azure-principal.json

      詳細については、「[Azure サービスプリンシパルを作成する](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli)」
      および「[Blob データへのアクセスのための Azure ロールを割り当てる](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli)」ページを参照してください。

      `client_id`、`tenant`、および `client_secret` の代わりに資格情報を rclone の設定ファイルに直接記述する方が便利な場合もあります。

   --use-msi
      管理されたサービス ID を使用して認証します（Azure でのみ機能します）。

      true の場合、SAS トークンやアカウントキーの代わりに、[管理されたサービス ID](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)を使用して、
      Azure Storage に認証します。

      このプログラムが実行されている VM（SS）にシステム割り当ての ID がある場合、デフォルトでそれが使用されます。
      リソースにシステム割り当てがなく、ユーザー割り当て ID が 1 つだけある場合、デフォルトでユーザー割り当て ID が使用されます。
      リソースに複数のユーザー割り当て ID がある場合、使用する ID は msi_object_id 、msi_client_id 、または msi_mi_res_id のいずれか 1 つを明示的に指定する必要があります。

   --msi-object-id
      使用するユーザー割り当て MSI のオブジェクト ID。

      msi_client_id または msi_mi_res_id が指定されている場合は空白のままで結構です。

   --msi-client-id
      使用するユーザー割り当て MSI のオブジェクト ID。

      msi_object_id または msi_mi_res_id が指定されている場合は空白のままで結構です。

   --msi-mi-res-id
      使用するユーザー割り当て MSI の Azure リソース ID。

      msi_client_id または msi_object_id が指定されている場合は空白のままで結構です。

   --use-emulator
      ローカルのストレージエミュレーターを使用します（'true' として指定）。

      実際の Azure ストレージエンドポイントを使用する場合は空白のままで結構です。

   --endpoint
      サービスのエンドポイント。

      通常は空白のままで結構です。

   --upload-cutoff
      チャンクアップロードへの切り替えのカットオフサイズ（<= 256 MiB）（非推奨）。

   --chunk-size
      アップロードのチャンクサイズ。

      メモリ内に保存されるため、"--transfers" * "--azureblob-upload-concurrency" 個のチャンクが同時にメモリに保存されます。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクのアップロード数です。

      高速回線上で大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合、
      この値を増やすことで転送を高速化することができます。

      テストでは、アップロード速度はアップロードの並列性とほぼ線形に増加します。
      たとえば、1 ギガビットのパイプを埋めるには、この値を 64 にする必要があるかもしれません。
      注意してください、これにはより多くのメモリを使用します。

      チャンクはメモリに保存されるため、"--transfers" * "--azureblob-upload-concurrency" 個の
      チャンクが同時にメモリに保存される場合があります。

   --list-chunk
      リストの Chunk サイズ。

      これにより、リストの各チャンクで要求されるブロブの数が設定されます。デフォルトは最大値の 5000 です。
      リストブロブのリクエストは、完了するまでに 1 メガバイトあたり 2 分の時間を持つことが許可されています。
      平均して 1 メガバイトあたり 2 分を超える操作はタイムアウトします
      （[ソース](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)）。
      タイムアウトを回避するために、ブロブのアイテム数を制限するために使用することができます。

   --access-tier
      ブロブのアクセスティア：hot、cool、または archive。

      アーカイブされたブロブは、アクセスティアを hot または cool に設定することで復元できます。
      アクセスティアが指定されていない場合、rclone はどのティアも適用しません。
      rclone はアップロード時に対象のブロブで「Set Tier」操作を実行しますが、オブジェクトが変更されていない場合、
      新しいアクセスティアに「Access Tier」を指定しても効果はありません。
      リモートのブロブが「アーカイブティア」にある場合、リモートからのデータ転送操作を実行することはできません。
      ユーザーはまず、ブロブを「Hot」または「Cool」にティアリングしてから復元する必要があります。

   --archive-tier-delete
      上書きする前にアーカイブティアのブロブを削除します。

      アーカイブティアのブロブは更新できません。したがって、アーカイブティアのブロブを更新しようとする場合、
      rclone は次のエラーを発生させます。

          can't update archive tier blob without --azureblob-archive-tier-delete

      このフラグが設定されている場合、rclone はアーカイブティアのブロブを上書きしようとする前に、
      そのブロブを削除してから置き換えをアップロードします。アップロードに失敗した場合、
      データの損失の可能性があり、また、早期にアーカイブティアのブロブを削除することで
      追加費用が発生する可能性もあります。

   --disable-checksum
      オブジェクトのメタデータに MD5 チェックサムを保存しません。

      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、
      オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するのに長い遅延が生じることがあります。

   --memory-pool-flush-time
      インターナルメモリーバッファプールがフラッシュされる頻度。

      追加のバッファ（たとえばマルチパート）を必要とするアップロードは、割り当てのためにメモリープールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      インターナルメモリープールで mmap バッファを使用するかどうか。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --public-access
      コンテナのパブリックアクセスレベル：blob または container。

      例:
         | <未設定>   | コンテナとそのブロブは、認証されたリクエストのみでアクセスできます。
         |          | デフォルトの値です。
         | blob     | このコンテナ内のブロブデータは、匿名リクエストで読み取ることができます。
         | container| コンテナとブロブデータへの完全なパブリックリードアクセスを許可します。

   --no-check-container
      コンテナの存在チェックや作成を試みません。

      既にコンテナが存在することを知っている場合、
      rclone のトランザクションの数を最小限に抑えるために役立つ場合があります。

   --no-head-object
      GET 操作時に HEAD を実行しません。

オプション:
   --account value                      Azure ストレージアカウント名。[$ACCOUNT]
   --client-certificate-password value  証明書ファイルのパスワード（オプション）。[$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      PEM または PKCS12 証明書ファイルのパス（プライベートキーも含む）。[$CLIENT_CERTIFICATE_PATH]
   --client-id value                    クライアントの ID。[$CLIENT_ID]
   --client-secret value                サービスプリンシパルのクライアントシークレットの 1 つ。[$CLIENT_SECRET]
   --env-auth                           実行時に資格情報を読み込みます（環境変数、CLI 、または MSI ）。 (デフォルト値: false) [$ENV_AUTH]
   --help, -h                           ヘルプを表示します
   --key value                          ストレージアカウントの共有キー。[$KEY]
   --sas-url value                      コンテナレベルのアクセス専用の SAS URL。[$SAS_URL]
   --tenant value                       サービスプリンシパルのテナントの ID。またはディレクトリ ID とも呼ばれます。[$TENANT]

   Advanced

   --access-tier value              ブロブのアクセスティア：hot、cool、または archive。[$ACCESS_TIER]
   --archive-tier-delete            上書きする前にアーカイブティアのブロブを削除します。 (デフォルト値: false) [$ARCHIVE_TIER_DELETE]
   --chunk-size value               アップロードのチャンクサイズ。 (デフォルト値: "4Mi") [$CHUNK_SIZE]
   --client-send-certificate-chain  証明書認証時に証明書チェーンを送信するかどうか。 (デフォルト値: false) [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               オブジェクトのメタデータに MD5 チェックサムを保存しません。 (デフォルト値: false) [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト値: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$ENCODING]
   --endpoint value                 サービスのエンドポイント。[$ENDPOINT]
   --list-chunk value               リストのチャンクサイズ。 (デフォルト値: 5000) [$LIST_CHUNK]
   --memory-pool-flush-time value   インターナルメモリーバッファプールがフラッシュされる頻度。 (デフォルト値: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           インターナルメモリープールで mmap バッファを使用するかどうか。 (デフォルト値: false) [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            使用するユーザー割り当て MSI のオブジェクト ID。[$MSI_CLIENT_ID]
   --msi-mi-res-id value            使用するユーザー割り当て MSI の Azure リソース ID。[$MSI_MI_RES_ID]
   --msi-object-id value            使用するユーザー割り当て MSI のオブジェクト ID。[$MSI_OBJECT_ID]
   --no-check-container             コンテナの存在チェックや作成を試みません。 (デフォルト値: false) [$NO_CHECK_CONTAINER]
   --no-head-object                 GET 操作時に HEAD を実行しません。 (デフォルト値: false) [$NO_HEAD_OBJECT]
   --password value                 ユーザーのパスワード [$PASSWORD]
   --public-access value            コンテナのパブリックアクセスレベル：blob または container。[$PUBLIC_ACCESS]
   --service-principal-file value   サービスプリンシパルに使用する資格情報が含まれるファイルのパス。[$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト値: 16) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードへのカットオフサイズ (<= 256 MiB)（非推奨）。[$UPLOAD_CUTOFF]
   --use-emulator                   ローカルのストレージエミュレーターを使用します（デフォルト値：false）。[$USE_EMULATOR]
   --use-msi                        管理されたサービス ID を使用して認証します（Azure でのみ機能します）（デフォルト値: false）。[$USE_MSI]
   --username value                 ユーザー名（通常はメールアドレス）[$USERNAME]

```
{% endcode %}