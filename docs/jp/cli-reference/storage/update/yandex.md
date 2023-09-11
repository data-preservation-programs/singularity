# ヤンデックスディスク

{% code fullWidth="true" %}
```
NAME:
   singularity storage update yandex - ヤンデックスディスク

USAGE:
   singularity storage update yandex [コマンドオプション] <名前|ID>

DESCRIPTION:
   --client-id
      OAuth クライアントID。
      
      通常は空白のままにします。

   --client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにします。

   --token
      OAuth アクセストークン (JSON形式)。

   --auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。

   --hard-delete
      ファイルをゴミ箱に入れる代わりに完全に削除します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuth クライアントID。 [$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示する

   高度なオプション

   --auth-url value   認証サーバーのURL。 [$AUTH_URL]
   --encoding value   バックエンドのエンコーディング。 (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete      ファイルをゴミ箱に入れる代わりに完全に削除します。 (default: false) [$HARD_DELETE]
   --token value      OAuth アクセストークン (JSON形式)。 [$TOKEN]
   --token-url value  トークンサーバーのURL。 [$TOKEN_URL]

```
{% endcode %}