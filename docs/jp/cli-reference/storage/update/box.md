# Box

{% code fullWidth="true" %}
```
名前:
   singularity storage update box - Box

使用法:
   singularity storage update box [コマンドオプション] <名前|ID>

説明:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにしておきます。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにしておきます。

   --token
      OAuthアクセストークンをJSONのブロブとして指定します。

   --auth-url
      認証サーバのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしておきます。

   --token-url
      トークンサーバのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしておきます。

   --root-folder-id
      rcloneが開始地点として非ルートフォルダを使用する場合に入力します。

   --box-config-file
      Boxアプリのconfig.jsonの場所
      
      通常は空白のままにしておきます。
      
      先頭の`~`はファイル名として展開され、`${RCLONE_CONFIG_DIR}`などの環境変数も展開されます。

   --access-token
      Boxアプリのプライマリアクセストークン
      
      通常は空白のままにしておきます。

   --box-sub-type
      

      例:
         | user       | Rcloneはユーザーを代表して操作します。
         | enterprise | Rcloneはサービスアカウントを代表して操作します。

   --upload-cutoff
      マルチパートアップロードに切り替えるための上限値（>= 50 MiB）。

   --commit-retries
      マルチパートファイルのコミットを試行する最大回数。

   --list-chunk
      リストのチャンクサイズ1-1000。

   --owned-by
      ログイン（メールアドレス）によって所有されているアイテムのみを表示します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

オプション:
   --access-token value     Boxアプリのプライマリアクセストークン [$ACCESS_TOKEN]
   --box-config-file value  Boxアプリのconfig.jsonの場所 [$BOX_CONFIG_FILE]
   --box-sub-type value     (デフォルト: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuthクライアントID [$CLIENT_ID]
   --client-secret value    OAuthクライアントシークレット [$CLIENT_SECRET]
   --help, -h               ヘルプを表示

   高度なオプション

   --auth-url value        認証サーバのURL [$AUTH_URL]
   --commit-retries value  マルチパートファイルのコミットを試行する最大回数（デフォルト: 100） [$COMMIT_RETRIES]
   --encoding value        バックエンドのエンコーディング（デフォルト: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot"） [$ENCODING]
   --list-chunk value      リストのチャンクサイズ1-1000（デフォルト: 1000） [$LIST_CHUNK]
   --owned-by value        ログイン（メールアドレス）によって所有されているアイテムのみを表示します [$OWNED_BY]
   --root-folder-id value  rcloneが開始地点として非ルートフォルダを使用する場合の値（デフォルト: "0"） [$ROOT_FOLDER_ID]
   --token value           OAuthアクセストークンをJSONのブロブとして指定します [$TOKEN]
   --token-url value       トークンサーバのURL [$TOKEN_URL]
   --upload-cutoff value   マルチパートアップロードに切り替えるための上限値（>= 50 MiB）（デフォルト: "50Mi"） [$UPLOAD_CUTOFF]

```
{% endcode %}