# Box

{% code fullWidth="true" %}
```
名前:
   singularity storage create box - Box

使用法:
   singularity storage create box [コマンドオプション] [引数...]

説明:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにします。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --token
      JSON形式のOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --root-folder-id
      rcloneが開始点として使用するルートフォルダー以外を指定します。

   --box-config-file
      Box App config.json ファイルの場所
      
      通常は空白のままにします。
      
      先頭の`~`はファイル名で展開され、`${RCLONE_CONFIG_DIR}`のような環境変数も展開されます。

   --access-token
      Box Appのプライマリアクセストークン
      
      通常は空白のままにします。

   --box-sub-type
      

      例:
         | user       | Rcloneはユーザーの代わりに動作します。
         | enterprise | Rcloneはサービスアカウントの代わりに動作します。

   --upload-cutoff
      マルチパートアップロードに切り替えるためのカットオフサイズ (>= 50 MiB)。

   --commit-retries
      マルチパートファイルのコミットを試行する最大回数。

   --list-chunk
      リストのチャンクサイズ 1-1000。

   --owned-by
      ログイン (メールアドレス) が所有するアイテムのみ表示します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --access-token value     Box Appのプライマリアクセストークン [$ACCESS_TOKEN]
   --box-config-file value  Box App config.json ファイルの場所 [$BOX_CONFIG_FILE]
   --box-sub-type value     (デフォルト: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuthクライアントID [$CLIENT_ID]
   --client-secret value    OAuthクライアントシークレット [$CLIENT_SECRET]
   --help, -h               ヘルプを表示する

   Advanced

   --auth-url value        認証サーバーのURL [$AUTH_URL]
   --commit-retries value  マルチパートファイルのコミットを試行する最大回数 (デフォルト: 100) [$COMMIT_RETRIES]
   --encoding value        バックエンドのエンコーディング (デフォルト: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --list-chunk value      リストのチャンクサイズ 1-1000 (デフォルト: 1000) [$LIST_CHUNK]
   --owned-by value        ログイン (メールアドレス) が所有するアイテムのみ表示します [$OWNED_BY]
   --root-folder-id value  rcloneが開始点として使用するルートフォルダー以外を指定します (デフォルト: "0") [$ROOT_FOLDER_ID]
   --token value           JSON形式のOAuthアクセストークン [$TOKEN]
   --token-url value       トークンサーバーのURL [$TOKEN_URL]
   --upload-cutoff value   マルチパートアップロードに切り替えるためのカットオフサイズ (>= 50 MiB) (デフォルト: "50Mi") [$UPLOAD_CUTOFF]

   General

   --name value  ストレージの名前 (デフォルト: Auto generated)
   --path value  ストレージのパス
```
{% endcode %}