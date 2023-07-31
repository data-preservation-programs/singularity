# Yandex Disk

{% code fullWidth="true" %}
```
名称:
   singularity データソースの追加 yandex - Yandex Disk

使用法:
   singularity データソースの追加 yandex [コマンドオプション] <データセット名> <ソースパス>

説明:
   --yandex-auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。

   --yandex-client-id
      OAuthクライアントID。
      
      通常は空白のままにします。

   --yandex-client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --yandex-encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --yandex-hard-delete
      ファイルをごみ箱に入れる代わりに完全に削除します。

   --yandex-token
      OAuthアクセストークン（JSON blob形式）。

   --yandex-token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。


オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後のスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期スキャン状態を設定します (デフォルト: ready)

   yandex用オプション

   --yandex-auth-url value       認証サーバーのURL。[$YANDEX_AUTH_URL]
   --yandex-client-id value      OAuthクライアントID。[$YANDEX_CLIENT_ID]
   --yandex-client-secret value  OAuthクライアントシークレット。[$YANDEX_CLIENT_SECRET]
   --yandex-encoding value       バックエンドのエンコーディング。 (デフォルト: "Slash,Del,Ctl,InvalidUtf8,Dot") [$YANDEX_ENCODING]
   --yandex-hard-delete value    ファイルをごみ箱に入れる代わりに完全に削除します。 (デフォルト: "false") [$YANDEX_HARD_DELETE]
   --yandex-token value          OAuthアクセストークン（JSON blob形式）。[$YANDEX_TOKEN]
   --yandex-token-url value      トークンサーバーのURL。[$YANDEX_TOKEN_URL]
```
{% endcode %}