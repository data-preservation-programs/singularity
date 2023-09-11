# Enterprise File Fabric

{% code fullWidth="true" %}
```
名前：
   singularity storage update filefabric - 商業用ファイルファブリック

使用法：
   singularity storage update filefabric [コマンドオプション] <名前|ID>

説明：
   --url
      接続する商業用ファイルファブリックのURL。

      例：
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 接続する商業用ファイルファブリック

   --root-folder-id
      ルートフォルダのID。

      通常は空白のままにします。

      特定のIDのディレクトリでrcloneを開始するために入力します。

   --permanent-token
      永続的な認証トークン。

      永続的な認証トークンは商業用ファイルファブリック内で作成できます。ユーザのダッシュボードの「セキュリティ」の項目にある「マイ認証トークン」というエントリが表示されます。「管理」ボタンをクリックしてトークンを作成します。

      これらのトークンは通常数年間有効です。

      詳細については次を参照してください：https://docs.storagemadeeasy.com/organisationcloud/api-tokens

   --token
      セッショントークン。

      これはrcloneが設定ファイルにキャッシュするセッショントークンです。通常は1時間有効です。

      この値を設定しないでください。rcloneが自動的に設定します。

   --token-expiry
      トークンの有効期限。

      この値を設定しないでください。rcloneが自動的に設定します。

   --version
      ファイルファブリックから読み取られるバージョン。

      この値を設定しないでください。rcloneが自動的に設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h               ヘルプを表示
   --permanent-token value  永続的な認証トークン。 [$PERMANENT_TOKEN]
   --root-folder-id value   ルートフォルダのID。 [$ROOT_FOLDER_ID]
   --url value              接続する商業用ファイルファブリックのURL。 [$URL]

   Advanced

   --encoding value      バックエンドのエンコーディング。 (デフォルト: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         セッショントークン。 [$TOKEN]
   --token-expiry value  トークンの有効期限。 [$TOKEN_EXPIRY]
   --version value       ファイルファブリックから読み取られるバージョン。 [$VERSION]

```
{% endcode %}