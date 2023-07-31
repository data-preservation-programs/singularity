# Pcloud

{% code fullWidth="true" %}
```
NAME:
   singularityデータソースの追加 pcloud - Pcloud

使い方:
   singularityデータソースの追加 pcloud [コマンドオプション] <データセット名> <ソースパス>

説明:
   --pcloud-auth-url
      認証サーバーのURLです。
      
      プロバイダーのデフォルトを使用するには空白のままにします。

   --pcloud-client-id
      OAuthクライアントIDです。
      
      通常は空白のままにします。

   --pcloud-client-secret
      OAuthクライアントシークレットです。
      
      通常は空白のままにします。

   --pcloud-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --pcloud-hostname
      接続するホスト名です。
      
      通常、rcloneが最初にOAuth接続を行うときに設定されますが、
      rclone authorizeを使用してリモート設定を使用している場合は手動で設定する必要があります。
      

      例:
         | api.pcloud.com  | 元の/米国地域
         | eapi.pcloud.com | EU地域

   --pcloud-password
      pcloudのパスワードです。

   --pcloud-root-folder-id
      rcloneがルートフォルダ以外のフォルダを開始点として使用するために入力してください。

   --pcloud-token
      OAuthアクセストークン(JSON形式)です。

   --pcloud-token-url
      トークンサーバーのURLです。
      
      プライダーのデフォルトを使用するには空白のままにします。

   --pcloud-username
      pcloudのユーザー名です。
            
      クリーンアップコマンドを使用する場合にのみ必要です。pcloud API のバグのため、このコマンドはOAuth認証をサポートしていないため、ユーザーパスワード認証に頼る必要があります。


オプション:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。  (default: false)
   --rescan-interval value  最後の正常なスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (default: disabled)
   --scanning-state value   初期スキャン状態を設定します (default: ready)

   pcloud用オプション

   --pcloud-auth-url value        認証サーバーのURLです。[$PCLOUD_AUTH_URL]
   --pcloud-client-id value       OAuthクライアントIDです。[$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value   OAuthクライアントシークレットです。[$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value        バックエンドのエンコーディングです。 (default: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PCLOUD_ENCODING]
   --pcloud-hostname value        接続するホスト名です。 (default: "api.pcloud.com") [$PCLOUD_HOSTNAME]
   --pcloud-password value        pcloudのパスワードです。[$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value  rcloneがルートフォルダ以外のフォルダを開始点として使用するために入力してください。 (default: "d0") [$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value           OAuthアクセストークン(JSON形式)です。[$PCLOUD_TOKEN]
   --pcloud-token-url value       トークンサーバーのURLです。[$PCLOUD_TOKEN_URL]
   --pcloud-username value        pcloudのユーザー名です。[$PCLOUD_USERNAME]

```
{% endcode %}