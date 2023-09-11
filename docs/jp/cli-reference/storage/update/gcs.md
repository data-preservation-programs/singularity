# Google Cloud Storage（これはGoogle Driveではありません）

{% code fullWidth="true" %}
```
名称：
   singularity storage update gcs - Google Cloud Storage（これはGoogle Driveではありません）

使用方法：
   singularity storage update gcs [コマンドオプション] <name|id>

説明：
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにしておきます。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにしておきます。

   --token
      OAuthアクセストークン（JSON形式）。

   --auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにしておきます。

   --token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにしておきます。

   --project-number
      プロジェクト番号。
      
      バケットのリスト/作成/削除にのみ必要です。開発者コンソールを参照してください。

   --service-account-file
      サービスアカウントの資格情報JSONファイルパス。
      
      通常は空白のままにしておきます。
      対話型ログインの代わりにSAを使用する場合のみ必要です。
      
      ファイル名には先頭の`〜`や、`${RCLONE_CONFIG_DIR}`などの環境変数が展開されます。

   --service-account-credentials
      サービスアカウントの資格情報JSONブロブ。
      
      通常は空白のままにしておきます。
      対話型ログインの代わりにSAを使用する場合のみ必要です。

   --anonymous
      資格情報なしでパブリックなバケットやオブジェクトにアクセスします。
      
      ファイルをダウンロードするだけで資格情報を設定しない場合は、'true'に設定します。

   --object-acl
      新しいオブジェクトのアクセスコントロールリスト。

      例：
         | authenticatedRead      | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | すべての認証済みユーザーにREADERアクセスが与えられます。
         | bucketOwnerFullControl | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | プロジェクトチームの所有者にOWNERアクセスが与えられます。
         | bucketOwnerRead        | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | プロジェクトチームの所有者にREADERアクセスが与えられます。
         | private                | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | 上記のいずれも指定されていない場合のデフォルトです。
         | projectPrivate         | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | プロジェクトチームのメンバーは、役割に応じたアクセスが与えられます。
         | publicRead             | オブジェクトの所有者にOWNERアクセスが与えられます。
         |                        | すべてのユーザーにREADERアクセスが与えられます。

   --bucket-acl
      新しいバケットのアクセスコントロールリスト。

      例：
         | authenticatedRead | プロジェクトチームの所有者はOWNERアクセスが与えられます。
         |                   | すべての認証済みユーザーにREADERアクセスが与えられます。
         | private           | プロジェクトチームの所有者はOWNERアクセスが与えられます。
         |                   | 上記のいずれも指定されていない場合のデフォルトです。
         | projectPrivate    | プロジェクトチームのメンバーは、役割に応じたアクセスが与えられます。
         | publicRead        | プロジェクトチームの所有者はOWNERアクセスが与えられます。
         |                   | すべてのユーザーにREADERアクセスが与えられます。
         | publicReadWrite   | プロジェクトチームの所有者はOWNERアクセスが与えられます。
         |                   | すべてのユーザーにWRITERアクセスが与えられます。

   --bucket-policy-only
      アクセスチェックはバケットレベルのIAMポリシーを使用する必要があります。
      
      Bucket Policy Onlyが設定されたバケットにオブジェクトをアップロードする場合は、これを設定する必要があります。
      
      設定されている場合、rcloneは次のように動作します：
      
      - バケットに設定されたACLを無視します
      - オブジェクトに設定されたACLを無視します
      - Bucket Policy Onlyが設定されたバケットを作成します
      
      ドキュメント：[https://cloud.google.com/storage/docs/bucket-policy-only](https://cloud.google.com/storage/docs/bucket-policy-only)
      

   --location
      新しく作成されるバケットの場所。

      例：
         | <unset>                 | デフォルトの場所（米国）
         | asia                    | アジアのためのマルチリージョンの場所
         | eu                      | ヨーロッパのためのマルチリージョンの場所
         | us                      | アメリカ合衆国のためのマルチリージョンの場所
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
         | southamerica-west1      | サンチアゴ
         | asia1                   | デュアルリージョン：アジア東北1とアジア東北2
         | eur4                    | デュアルリージョン：ヨーロッパ北1とヨーロッパ西4
         | nam4                    | デュアルリージョン：米国中央1と米国東1

   --storage-class
      Google Cloud Storageでオブジェクトを保存する際に使用するストレージクラス。

      例：
         | <unset>                      | デフォルト
         | MULTI_REGIONAL               | マルチリージョンのストレージクラス
         | REGIONAL                     | リージョナルのストレージクラス
         | NEARLINE                     | Nearlineのストレージクラス
         | COLDLINE                     | Coldlineのストレージクラス
         | ARCHIVE                      | Archiveのストレージクラス
         | DURABLE_REDUCED_AVAILABILITY | Durable reduced availabilityのストレージクラス

   --no-check-bucket
      バケットの存在を確認せず、作成しようとしない場合は設定します。
      
      バケットが既に存在することを知っている場合、rcloneが実行するトランザクションの数を最小限にするために役立ちます。
      

   --decompress
      設定するとgzipでエンコードされたオブジェクトが展開されます。
      
      "Content-Encoding: gzip"が設定されたオブジェクトをGCSにアップロードできます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは"Content-Encoding: gzip"で受け取ったファイルを展開します。これにより、rcloneはサイズとハッシュをチェックできなくなりますが、ファイルの内容は展開されます。
      

   --endpoint
      サービスのエンドポイント。
      
      通常は空白のままにしておきます。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --env-auth
      実行時にGCP IAM資格情報を取得します（環境変数または環境変数がない場合はインスタンスメタデータ）。
      
      service_account_fileとservice_account_credentialsが空白の場合のみ適用されます。

      例：
         | false | 次のステップで資格情報を入力します。
         | true  | 環境からGCP IAMの資格情報を取得します（環境変数またはIAM）。


オプション：
   --anonymous                          資格情報なしでパブリックなバケットやオブジェクトにアクセスします。 (デフォルト: false) [$ANONYMOUS]
   --bucket-acl value                   新しいバケットのアクセスコントロールリスト。 [$BUCKET_ACL]
   --bucket-policy-only                 アクセスチェックはバケットレベルのIAMポリシーを使用する必要があります。 (デフォルト: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value                OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --env-auth                           実行時にGCP IAM資格情報を取得します（環境変数または環境変数がない場合はインスタンスメタデータ）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                           ヘルプを表示します
   --location value                     新しく作成されるバケットの場所。 [$LOCATION]
   --object-acl value                   新しいオブジェクトのアクセスコントロールリスト。 [$OBJECT_ACL]
   --project-number value               プロジェクト番号。 [$PROJECT_NUMBER]
   --service-account-credentials value  サービスアカウントの資格情報JSONブロブ。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         サービスアカウントの資格情報JSONファイルパス。 [$SERVICE_ACCOUNT_FILE]
   --storage-class value                Google Cloud Storageでオブジェクトを保存する際に使用するストレージクラス。 [$STORAGE_CLASS]

   Advanced

   --auth-url value   認証サーバーのURL。 [$AUTH_URL]
   --decompress       設定するとgzipでエンコードされたオブジェクトが展開されます。 (デフォルト: false) [$DECOMPRESS]
   --encoding value   バックエンドのエンコーディング。 (デフォルト: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   サービスのエンドポイント。 [$ENDPOINT]
   --no-check-bucket  バケットの存在を確認せず、作成しようとしない場合は設定します。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --token value      OAuthアクセストークン（JSON形式）。 [$TOKEN]
   --token-url value  トークンサーバーのURL。 [$TOKEN_URL]

```
{% endcode %}