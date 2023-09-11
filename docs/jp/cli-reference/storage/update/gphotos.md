# Google Photos

{% code fullWidth="true" %}
```
NAME:
   singularity storage update gphotos - Google Photos

USAGE:
   singularity storage update gphotos [command options] <name|id>

DESCRIPTION:
   --client-id
      OAuthクライアントIDです。
      
      通常は空白のままにしてください。

   --client-secret
      OAuthクライアントシークレットです。
      
      通常は空白のままにしてください。

   --token
      JSON形式のOAuthアクセストークンです。

   --auth-url
      認証サーバーのURLです。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --token-url
      トークンサーバーのURLです。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --read-only
      Google Photosバックエンドを読み取り専用に設定します。
      
      読み取り専用を選択すると、rcloneは写真への読み取り専用アクセスを要求します。それ以外の場合はフルアクセスを要求します。

   --read-size
      メディアアイテムのサイズを読み取るように設定します。
      
      通常、rcloneはメディアアイテムのサイズを読み取りません。これは別のトランザクションを必要とするためです。同期には必要ありません。ただし、rclone mountは、読み取る前にファイルのサイズを事前に知る必要があるため、rclone mountを使用する場合は、このフラグを設定することをお勧めします。

   --start-year
      アップロードされた年以降の写真のみをダウンロードするように制限します。

   --include-archived
      アーカイブされたメディアも表示およびダウンロードします。
      
      デフォルトでは、rcloneはアーカイブされたメディアをリクエストしません。したがって、同期する場合、アーカイブされたメディアはディレクトリリストや転送時に表示されません。
      
      アルバム内のメディアは、アーカイブステータスに関係なく常に表示および同期されます。
      
      このフラグを使用すると、アーカイブされたメディアは常にディレクトリリストに表示され、転送されます。
      
      このフラグを使用しない場合、アーカイブされたメディアはディレクトリリストに表示されず、転送されません。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuthクライアントIDです。 [$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレットです。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示します
   --read-only            Google Photosバックエンドを読み取り専用に設定します。 (デフォルト: false) [$READ_ONLY]

   Advanced

   --auth-url value    認証サーバーのURLです。 [$AUTH_URL]
   --encoding value    バックエンドのエンコーディングです。 (デフォルト: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --include-archived  アーカイブされたメディアも表示およびダウンロードします。 (デフォルト: false) [$INCLUDE_ARCHIVED]
   --read-size         メディアアイテムのサイズを読み取るように設定します。 (デフォルト: false) [$READ_SIZE]
   --start-year value  アップロードされた年以降の写真のみをダウンロードするように制限します。 (デフォルト: 2000) [$START_YEAR]
   --token value       JSON形式のOAuthアクセストークンです。 [$TOKEN]
   --token-url value   トークンサーバーのURLです。 [$TOKEN_URL]

```
{% endcode %}