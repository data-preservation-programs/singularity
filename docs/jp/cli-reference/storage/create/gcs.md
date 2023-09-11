# Google Cloud Storage（これはGoogle Driveではありません）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create gcs - Google Cloud Storage（これはGoogle Driveではありません）

使用法:
   singularity storage create gcs [command options] [arguments...]

説明:
   --client-id
      OAuthクライアントID。
      
      通常、空白のままにします。

   --client-secret
      OAuthクライアントシークレット。
      
      通常、空白のままにします。

   --token
      OAuthアクセストークン（JSONブロブ）。

   --auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルトを使用する場合、空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用する場合、空白のままにします。

   --project-number
      プロジェクト番号。
      
      オプション - バケットのリスト/作成/削除にのみ必要です。開発者コンソールを参照してください。

   --service-account-file
      サービスアカウントの認証情報JSONファイルのパス。
      
      通常、空白のままにします。
      対話型ログインの代わりにSAを使用する場合にのみ必要です。
      
      先頭の`~`はファイル名で展開されます。`${RCLONE_CONFIG_DIR}`などの環境変数も展開されます。

   --service-account-credentials
      サービスアカウントの認証情報JSONブロブ。
      
      通常、空白のままにします。
      対話型ログインの代わりにSAを使用する場合にのみ必要です。

   --anonymous
      認証情報なしでパブリックバケットとオブジェクトにアクセスします。
      
      ファイルをダウンロードするだけで、認証情報を構成しない場合は「true」に設定します。

   --object-acl
      新しいオブジェクトのアクセス制御リスト。

      例:
         | authenticatedRead      | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | 認証済みユーザー全員にREADERアクセス権が与えられます。
         | bucketOwnerFullControl | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | プロジェクトチームの所有者にOWNERアクセス権が与えられます。
         | bucketOwnerRead        | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | プロジェクトチームの所有者にREADERアクセス権が与えられます。
         | private                | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | デフォルト（空白の場合）。
         | projectPrivate         | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | プロジェクトチームメンバーは役割に基づいてアクセス権が与えられます。
         | publicRead             | オブジェクトの所有者にOWNERアクセス権が与えられます。
         |                        | 全ユーザーにREADERアクセス権が与えられます。

   --bucket-acl
      新しいバケットのアクセス制御リスト。

      例:
         | authenticatedRead | プロジェクトチームの所有者にOWNERアクセス権が与えられます。
         |                   | 認証済みユーザー全員にREADERアクセス権が与えられます。
         | private           | プロジェクトチームの所有者にOWNERアクセス権が与えられます。
         |                   | デフォルト（空白の場合）。
         | projectPrivate    | プロジェクトチームメンバーは役割に基づいてアクセス権が与えられます。
         | publicRead        | プロジェクトチームの所有者にOWNERアクセス権が与えられます。
         |                   | 全ユーザーにREADERアクセス権が与えられます。
         | publicReadWrite   | プロジェクトチームの所有者にOWNERアクセス権が与えられます。
         |                   | 全ユーザーにWRITERアクセス権が与えられます。

   --bucket-policy-only
      アクセスチェックはバケットレベルのIAMポリシーを使用する必要があります。
      
      Bucket Policy Onlyが設定されたバケットにオブジェクトをアップロードする場合、この設定が必要です。
      
      この設定を行うと、rcloneは次のように動作します。
      
      - バケットに設定されたACLを無視します。
      - オブジェクトに設定されたACLを無視します。
      - Bucket Policy Onlyが設定されているバケットを作成します。
      
      ドキュメント：https://cloud.google.com/storage/docs/bucket-policy-only
      

   --location
      新しく作成されるバケットの場所。

      例:
         | <unset>                 | デフォルトの場所（米国）
         | asia                    | アジア向けのマルチリージョンの場所
         | eu                      | ヨーロッパ向けのマルチリージョンの場所
         | us                      | アメリカ合衆国向けのマルチリージョンの場所
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
         | us-east1                | 南カロライナ州
         | us-east4                | バージニア北部
         | us-west1                | オレゴン
         | us-west2                | カリフォルニア
         | us-west3                | ソルトレイクシティ
         | us-west4                | ラスベガス
         | northamerica-northeast1 | モントリオール
         | northamerica-northeast2 | トロント
         | southamerica-east1      | サンパウロ
         | southamerica-west1      | サンティアゴ
         | asia1                   | デュアルリージョン：Asia-northeast1およびAsia-northeast2
         | eur4                    | デュアルリージョン：Europe-north1およびEurope-west4
         | nam4                    | デュアルリージョン：US-central1およびUS-east1

   --storage-class
      Google Cloud Storageにオブジェクトを格納する際に使用するストレージクラス。

      例:
         | <unset>                      | デフォルト
         | MULTI_REGIONAL               | マルチリージョンストレージクラス
         | REGIONAL                     | リージョナルストレージクラス
         | NEARLINE                     | ニアラインストレージクラス
         | COLDLINE                     | コールドラインストレージクラス
         | ARCHIVE                      | アーカイブストレージクラス
         | DURABLE_REDUCED_AVAILABILITY | 耐久性のある低い可用性ストレージクラス

   --no-check-bucket
      バケットの存在を確認せず、作成も試みません。
      
      これは、バケットが既に存在する場合や、トランザクションの回数を最小限に抑えたい場合に便利です。
      

   --decompress
      これを設定すると、gzipエンコードされたオブジェクトが解凍されます。
      
      GCSに「Content-Encoding: gzip」が設定されたファイルをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信した"Content-Encoding: gzip"のファイルを解凍します。これにより、rcloneはサイズやハッシュを確認することはできませんが、ファイルの内容は解凍されます。
      

   --endpoint
      サービスのエンドポイント。
      
      通常、空白のままにします。

   --encoding
      バックエンドのエンコーディング。
      
      [概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --env-auth
      ランタイムからGCP IAMの認証情報を取得します（環境変数または環境変数がない場合はインスタンスメタデータ）。
      
      サービスアカウントファイルとサービスアカウントの認証情報が空白の場合のみ適用されます。

      例:
         | false | 次のステップで認証情報を入力します。
         | true  | 環境からGCP IAMの認証情報を取得します（環境変数またはIAM）。


オプション:
   --anonymous                          認証情報なしでパブリックバケットとオブジェクトにアクセスします。 (default: false) [$ANONYMOUS]
   --bucket-acl value                   新しいバケットのアクセス制御リスト。 [$BUCKET_ACL]
   --bucket-policy-only                 アクセスチェックはバケットレベルのIAMポリシーを使用する必要があります。 (default: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value                OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --env-auth                           ランタイムからGCP IAMの認証情報を取得します（環境変数または環境変数がない場合はインスタンスメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                           ヘルプを表示
   --location value                     新しく作成されるバケットの場所。 [$LOCATION]
   --object-acl value                   新しいオブジェクトのアクセス制御リスト。 [$OBJECT_ACL]
   --project-number value               プロジェクト番号。 [$PROJECT_NUMBER]
   --service-account-credentials value  サービスアカウントの認証情報JSONブロブ。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         サービスアカウントの認証情報JSONファイルのパス。 [$SERVICE_ACCOUNT_FILE]
   --storage-class value                Google Cloud Storageにオブジェクトを格納する際に使用するストレージクラス。 [$STORAGE_CLASS]

   Advanced

   --auth-url value   認証サーバーのURL。 [$AUTH_URL]
   --decompress       これを設定すると、gzipエンコードされたオブジェクトが解凍されます。 (default: false) [$DECOMPRESS]
   --encoding value   バックエンドのエンコーディング。 (default: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   サービスのエンドポイント。 [$ENDPOINT]
   --no-check-bucket  バケットの存在を確認せず、作成も試みません。 (default: false) [$NO_CHECK_BUCKET]
   --token value      OAuthアクセストークン（JSONブロブ）。 [$TOKEN]
   --token-url value  トークンサーバーのURL。 [$TOKEN_URL]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}