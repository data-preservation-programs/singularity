# Google Cloud Storage（これはGoogleドライブではありません）

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add gcs - Google Cloud Storage（これはGoogleドライブではありません）

USAGE:
   singularity datasource add gcs [コマンドオプション] <データセット名> <ソースパス>

DESCRIPTION:
   --gcs-anonymous
      認証情報なしでパブリックバケットやオブジェクトにアクセスします。
      
      ファイルをダウンロードするだけで認証情報の設定を行いたくない場合には「true」と設定します。

   --gcs-auth-url
      認証サーバーのURLです。
      
      デフォルトのプロバイダーを使用するため、空白のままにしておきます。

   --gcs-bucket-acl
      新しいバケットのアクセス制御リストです。

      例:
         | authenticatedRead | プロジェクトチームのオーナーには所有者アクセス権が付与されます。
                             | すべての認証済みユーザーにはリーダーアクセス権が付与されます。
         | private           | プロジェクトチームのオーナーには所有者アクセス権が付与されます。
                             | 既定値です。
         | projectPrivate    | プロジェクトチームのメンバーには、役割によってアクセス権が付与されます。
         | publicRead        | プロジェクトチームのオーナーには所有者アクセス権が付与されます。
                             | すべてのユーザーにはリーダーアクセス権が付与されます。
         | publicReadWrite   | プロジェクトチームのオーナーには所有者アクセス権が付与されます。
                             | すべてのユーザーにはライターアクセス権が付与されます。

   --gcs-bucket-policy-only
      バケットレベルのIAMポリシーを使用してアクセス確認を行います。
      
      バケットにBucket Policy Onlyが設定されている場合、オブジェクトのアップロードを行いたい場合にはこれを設定する必要があります。
      
      この設定が有効な場合、rcloneは以下の操作を行います:
      
      - バケットに設定されたACLを無視します
      - オブジェクトに設定されたACLを無視します
      - Bucket Policy Onlyが設定されたバケットを作成します
      
      ドキュメント: https://cloud.google.com/storage/docs/bucket-policy-only
      

   --gcs-client-id
      OAuthのクライアントIDです。
      
      通常は空白のままにしておきます。

   --gcs-client-secret
      OAuthのクライアントシークレットです。
      
      通常は空白のままにしておきます。

   --gcs-decompress
      これを設定すると、gzipでエンコードされたオブジェクトを解凍します。
      
      "Content-Encoding: gzip"が設定された状態でオブジェクトをGCSにアップロードすることができます。
      通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信時に"Content-Encoding: gzip"としてこれらのファイルを解凍します。
      これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --gcs-encoding
      バックエンドのエンコーディングです。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --gcs-endpoint
      サービスのエンドポイントです。
      
      通常は空白のままにしておきます。

   --gcs-env-auth
      ランタイムからGCP IAMの認証情報を取得します（環境変数またはインスタンスメタデータがない場合は環境変数に取ります）。
      
      service_account_fileとservice_account_credentialsが空白の場合にのみ適用されます。

      例:
         | false | 次のステップで認証情報を入力します。
         | true  | ランタイムからGCP IAMの認証情報を取得します（環境変数またはIAM）。

   --gcs-location
      新しく作成されるバケットの場所です。

      例:
         | <unset>                 | デフォルトの場所（US）
         | asia                    | アジアのマルチリージョンの場所
         | eu                      | ヨーロッパのマルチリージョンの場所
         | us                      | アメリカ合衆国のマルチリージョンの場所
         | asia-east1              | 台湾
         | asia-east2              | 香港
         | asia-northeast1         | 東京
         | asia-northeast2         | 大阪
         | asia-northeast3         | ソウル
         | asia-south1             | ムンバイ
         | asia-south2             | デリー
         | asia-southeast1         | シンガポール
         | asia-southeast2         | ジャカルタ
         | australia-southeast1    | シドニー
         | australia-southeast2    | メルボルン
         | europe-north1           | フィンランド
         | europe-west1            | ベルギー
         | europe-west2            | ロンドン
         | europe-west3            | フランクフルト
         | europe-west4            | オランダ
         | europe-west6            | チューリッヒ
         | europe-central2         | ワルシャワ
         | us-central1             | アイオワ
         | us-east1                | 南カロライナ
         | us-east4                | ノーザンバージニア
         | us-west1                | オレゴン
         | us-west2                | カリフォルニア
         | us-west3                | ソルトレイクシティ
         | us-west4                | ラスベガス
         | northamerica-northeast1 | モントリオール
         | northamerica-northeast2 | トロント
         | southamerica-east1      | サンパウロ
         | southamerica-west1      | サンティアゴ
         | asia1                   | デュアルリージョン: asia-northeast1 と asia-northeast2.
         | eur4                    | デュアルリージョン: europe-north1 と europe-west4.
         | nam4                    | デュアルリージョン: us-central1 と us-east1.

   --gcs-no-check-bucket
      設定済みのバケットの存在チェックや作成を行わないようにします。
      
      バケットが既に存在することがわかっている場合に、rcloneが実行するトランザクションの数を最小限に抑えるために便利です。
      

   --gcs-object-acl
      新しいオブジェクトのアクセス制御リストです。

      例:
         | authenticatedRead      | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | すべての認証済みユーザーにはリーダーアクセス権が付与されます。
         | bucketOwnerFullControl | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | プロジェクトチームのオーナーには所有者アクセス権が付与されます。
         | bucketOwnerRead        | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | プロジェクトチームのオーナーにはリーダーアクセス権が付与されます。
         | private                | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | 既定値です。
         | projectPrivate         | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | プロジェクトチームのメンバーは、役割に応じてアクセス権が付与されます。
         | publicRead             | オブジェクトの所有者には所有者アクセス権が付与されます。
                                  | すべてのユーザーにはリーダーアクセス権が付与されます。

   --gcs-project-number
      プロジェクト番号です。
      
      バケットのリスト／作成／削除に必要です。開発者コンソールをご覧ください。

   --gcs-service-account-credentials
      サービスアカウントの認証情報JSONブロブです。
      
      通常は空白のままにしておきます。対話型ログインの代わりにSAを使用する場合にのみ必要です。

   --gcs-service-account-file
      サービスアカウントの認証情報JSONファイルパスです。
      
      通常は空白のままにしておきます。対話型ログインの代わりにSAを使用する場合にのみ必要です。
      
      先頭の `~` はファイル名で展開され、`${RCLONE_CONFIG_DIR}` のような環境変数も展開されます。

   --gcs-storage-class
      Google Cloud Storageにオブジェクトを格納する場合に使用するストレージクラスです。

      例:
         | <unset>                      | デフォルト
         | MULTI_REGIONAL               | マルチリージョンのストレージクラス
         | REGIONAL                     | リージョナルのストレージクラス
         | NEARLINE                     | ニアラインのストレージクラス
         | COLDLINE                     | コールドラインのストレージクラス
         | ARCHIVE                      | アーカイブのストレージクラス
         | DURABLE_REDUCED_AVAILABILITY | 耐久力の低下した可用性のストレージクラス

   --gcs-token
      OAuth Access TokenをJSONブロブとして入力します。

   --gcs-token-url
      トークンサーバーのURLです。
      
      プロバイダーのデフォルトを使用するため、空白のままにしておきます。


OPTIONS:
   --help, -h  ヘルプを表示します

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval value  ラストスキャンからこのインターバルが経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   イニシャルスキャンの状態を設定します (デフォルト: ready)

   gcsのオプション

   --gcs-anonymous value             認証情報なしでパブリックバケットやオブジェクトにアクセスします。 (デフォルト: "false") [$GCS_ANONYMOUS]
   --gcs-auth-url value              認証サーバーのURLです。 [$GCS_AUTH_URL]
   --gcs-bucket-acl value            新しいバケットのアクセス制御リストです。 [$GCS_BUCKET_ACL]
   --gcs-bucket-policy-only value    バケットレベルのIAMポリシーを使用してアクセス確認を行います。 (デフォルト: "false") [$GCS_BUCKET_POLICY_ONLY]
   --gcs-client-id value             OAuthのクライアントIDです。 [$GCS_CLIENT_ID]
   --gcs-client-secret value         OAuthのクライアントシークレットです。 [$GCS_CLIENT_SECRET]
   --gcs-decompress value            これを設定すると、gzipでエンコードされたオブジェクトを解凍します。 (デフォルト: "false") [$GCS_DECOMPRESS]
   --gcs-encoding value              バックエンドのエンコーディングです。 (デフォルト: "Slash,CrLf,InvalidUtf8,Dot") [$GCS_ENCODING]
   --gcs-endpoint value              サービスのエンドポイントです。 [$GCS_ENDPOINT]
   --gcs-env-auth value              ランタイムからGCP IAMの認証情報を取得します（環境変数またはインスタンスメタデータがない場合は環境変数に取ります）。 (デフォルト: "false") [$GCS_ENV_AUTH]
   --gcs-location value              新しく作成されるバケットの場所です。 [$GCS_LOCATION]
   --gcs-no-check-bucket value       設定済みのバケットの存在チェックや作成を行わないようにします。 (デフォルト: "false") [$GCS_NO_CHECK_BUCKET]
   --gcs-object-acl value            新しいオブジェクトのアクセス制御リストです。 [$GCS_OBJECT_ACL]
   --gcs-project-number value        プロジェクト番号です。 [$GCS_PROJECT_NUMBER]
   --gcs-service-account-file value  サービスアカウントの認証情報JSONファイルパスです。 [$GCS_SERVICE_ACCOUNT_FILE]
   --gcs-storage-class value         Google Cloud Storageにオブジェクトを格納する場合に使用するストレージクラスです。 [$GCS_STORAGE_CLASS]
   --gcs-token value                 OAuth Access TokenをJSONブロブとして入力します。 [$GCS_TOKEN]
   --gcs-token-url value             トークンサーバーのURLです。 [$GCS_TOKEN_URL]

```
{% endcode %}