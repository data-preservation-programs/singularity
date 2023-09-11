# Pcloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create pcloud - Pcloud

USAGE:
   singularity storage create pcloud [command options] [arguments...]

DESCRIPTION:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままでかまいません。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままでかまいません。

   --token
      JSONブロブ形式でのOAuthアクセストークン。

   --auth-url
      AuthサーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --root-folder-id
      rcloneが非ルートフォルダを起点として使用するために入力してください。

   --hostname
      接続するホスト名。
      
      通常、rcloneが最初にOAuth接続を行ったときに設定されますが、
      rclone authorizeを使用してリモート構成を使用している場合は手動で設定する必要があります。

      例:
         | api.pcloud.com  | オリジナル/USリージョン
         | eapi.pcloud.com | EUリージョン

   --username
      pcloudのユーザー名。
            
      cleanupコマンドを使用したい場合にのみ必要です。pcloud APIのバグにより、
      必要なAPIはOAuth認証をサポートしていないため、ユーザーパスワード認証に依存する必要があります。

   --password
      pcloudのパスワード。


OPTIONS:
   --client-id value      OAuthクライアントID。[$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示します

   Advanced

   --auth-url value        AuthサーバーのURL。[$AUTH_URL]
   --encoding value        バックエンドのエンコーディング。 (デフォルト値: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hostname value        接続するホスト名。 (デフォルト値: "api.pcloud.com") [$HOSTNAME]
   --password value        pcloudのパスワード。[$PASSWORD]
   --root-folder-id value  rcloneが非ルートフォルダを起点として使用するために入力してください。 (デフォルト値: "d0") [$ROOT_FOLDER_ID]
   --token value           JSONブロブ形式でのOAuthアクセストークン。[$TOKEN]
   --token-url value       トークンサーバーのURL。[$TOKEN_URL]
   --username value        pcloudのユーザー名。[$USERNAME]

   General

   --name value  ストレージの名前 (デフォルト値: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}