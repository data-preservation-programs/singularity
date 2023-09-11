# Enterprise File Fabric

{% code fullWidth="true" %}
```
名前:
   シンギュラリティストレージ作成: ファイルファブリック - Enterprise File Fabric

使用方法:
   singularity storage create filefabric [コマンドオプション] [引数...]

説明:
   --url
      接続するEnterprise File FabricのURL。

      例:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | あなたのEnterprise File Fabricに接続

   --root-folder-id
      ルートフォルダのID。
    
      通常は空白のままにします。

      指定すると、rcloneは指定されたIDのディレクトリで開始します。
      

   --permanent-token
      永続認証トークン。

      Enterprise File Fabricで作成できる認証トークンです。ユーザーダッシュボードのセキュリティの下に「マイ認証トークン」という項目があります。作成するには"Manage"ボタンをクリックします。

      これらのトークンは通常何年も有効です。

      詳細についてはこちらを参照してください: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      セッショントークン。

      これはrcloneが設定ファイルにキャッシュするセッショントークンです。通常、有効期限は1時間です。

      この値を設定しないでください - rcloneが自動的に設定します。
      

   --token-expiry
      トークンの有効期限。

      この値を設定しないでください - rcloneが自動的に設定します。
      

   --version
      ファイルファブリックから読み込まれたバージョン。

      この値を設定しないでください - rcloneが自動的に設定します。
      

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h               ヘルプを表示
   --permanent-token value  永続認証トークン。 [$PERMANENT_TOKEN]
   --root-folder-id value   ルートフォルダのID。 [$ROOT_FOLDER_ID]
   --url value              接続するEnterprise File FabricのURL。 [$URL]

   高度なオプション

   --encoding value      バックエンドのエンコーディング。 (デフォルト値: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         セッショントークン。 [$TOKEN]
   --token-expiry value  トークンの有効期限。 [$TOKEN_EXPIRY]
   --version value       ファイルファブリックから読み込まれたバージョン。 [$VERSION]

   一般

   --name value  ストレージの名前 (デフォルト値: 自動生成)
   --path value  ストレージのパス
```
{% endcode %}