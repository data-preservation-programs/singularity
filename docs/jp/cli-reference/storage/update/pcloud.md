# Pcloud

{% code fullWidth="true" %}
```
名前:
   singularity storage update pcloud - Pcloud

使用法:
   singularity storage update pcloud [コマンドオプション] <名前|ID>

説明:
   --client-id
      OAuth クライアント ID。
      
      通常は空白のままにします。

   --client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにします。

   --token
      OAuth アクセストークン（JSON ブロブ）。

   --auth-url
      認証サーバの URL。
      
      プロバイダのデフォルトを使用するには空白のままにします。

   --token-url
      トークンサーバの URL。
      
      プロバイダのデフォルトを使用するには空白のままにします。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --root-folder-id
      rclone が開始点として非ルートフォルダを使用するために入力します。

   --hostname
      接続するホスト名。
      
      通常、rclone が最初に OAuth 接続を行う際に設定されますが、
      rclone authorize を使ってリモート設定を行っている場合は手動で設定する必要があります。
      

      例:
         | api.pcloud.com  | オリジナル/US リージョン
         | eapi.pcloud.com | EU リージョン

   --username
      pcloud のユーザ名。
            
      cleanup コマンドを使用する場合にのみ必要です。Pcloud API の仕様上サポートされていないため、
      OAuth 認証の代わりにユーザパスワード認証を利用する必要があります。

   --password
      pcloud のパスワード。


オプション:
   --client-id value      OAuth クライアント ID。[$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示する

   高度なオプション

   --auth-url value        認証サーバの URL。[$AUTH_URL]
   --encoding value        バックエンドのエンコーディング。（デフォルト： "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）[$ENCODING]
   --hostname value        接続するホスト名。（デフォルト： "api.pcloud.com"）[$HOSTNAME]
   --password value        pcloud のパスワード。[$PASSWORD]
   --root-folder-id value  rclone が開始点として非ルートフォルダを使用するために入力します。（デフォルト： "d0"）[$ROOT_FOLDER_ID]
   --token value           OAuth アクセストークン（JSON ブロブ）。[$TOKEN]
   --token-url value       トークンサーバの URL。[$TOKEN_URL]
   --username value        pcloud のユーザ名。[$USERNAME]

```
{% endcode %}