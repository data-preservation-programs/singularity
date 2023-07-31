# Box

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add box - Box

USAGE:
   singularity datasource add box [command options] <dataset_name> <source_path>

DESCRIPTION:
   --box-access-token
      Box アプリのプライマリアクセストークン
      
ふつうは空にしておきます。

   --box-auth-url
      認証サーバーのURL
      
プロバイダのデフォルトを使用する場合は空にしておきます。

   --box-box-config-file
      Box アプリ config.json の場所
      
ふつうは空にしておきます。
      
ファイル名には `~` が含まれたり `${RCLONE_CONFIG_DIR}` のような環境変数が展開されます。

   --box-box-sub-type
      タイプの例:
         | user       | ユーザーの代わりに Rclone が処理を実行します。
         | enterprise | サービスアカウントの代わりに Rclone が処理を実行します。

   --box-client-id
      OAuth クライアントID
      
ふつうは空にしておきます。

   --box-client-secret
      OAuth クライアントシークレット
      
ふつうは空にしておきます。

   --box-commit-retries
      マルチパートファイルをコミットしようとする回数の最大値

   --box-encoding
      バックエンドのエンコーディング
      
詳細につきましては、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --box-list-chunk
      リストチャンクのサイズ 1〜1000

   --box-owned-by
      ログイン（メールアドレス）が所有するアイテムのみ表示します。

   --box-root-folder-id
      rclone が開始位置としてノンルートフォルダを使用する場合に入力します。

   --box-token
      JSON形式のOAuthアクセストークン

   --box-token-url
      トークンサーバーURL
      
プロバイダのデフォルトを使用する場合は空にしておきます。

   --box-upload-cutoff
      マルチパートアップロードに切り替えるためのカットオフサイズ (>= 50 MiB)


OPTIONS:
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [注意] データセットのファイルをエクスポートした後に削除します。  (デフォルト: false)
   --rescan-interval value  最後のスキャンからこの間隔が経過したら、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   box用オプション

   --box-access-token value     Box アプリのプライマリアクセストークン [$BOX_ACCESS_TOKEN]
   --box-auth-url value         認証サーバーのURL [$BOX_AUTH_URL]
   --box-box-config-file value  Box アプリ config.json の場所 [$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value     (デフォルト: "user") [$BOX_BOX_SUB_TYPE]
   --box-client-id value        OAuth クライアントID [$BOX_CLIENT_ID]
   --box-client-secret value    OAuth クライアントシークレット [$BOX_CLIENT_SECRET]
   --box-commit-retries value   マルチパートファイルをコミットしようとする回数の最大値 (デフォルト: "100") [$BOX_COMMIT_RETRIES]
   --box-encoding value         バックエンドのエンコーディング (デフォルト: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$BOX_ENCODING]
   --box-list-chunk value       リストチャンクのサイズ 1〜1000 (デフォルト: "1000") [$BOX_LIST_CHUNK]
   --box-owned-by value         ログイン（メールアドレス）が所有するアイテムのみ表示します。 [$BOX_OWNED_BY]
   --box-root-folder-id value   rclone が開始位置としてノンルートフォルダを使用する場合のID (デフォルト: "0") [$BOX_ROOT_FOLDER_ID]
   --box-token value            JSON形式のOAuthアクセストークン [$BOX_TOKEN]
   --box-token-url value        トークンサーバーURL [$BOX_TOKEN_URL]
   --box-upload-cutoff value    マルチパートアップロードに切り替えるためのカットオフサイズ (>= 50 MiB) (デフォルト: "50Mi") [$BOX_UPLOAD_CUTOFF]

```
{% endcode %}