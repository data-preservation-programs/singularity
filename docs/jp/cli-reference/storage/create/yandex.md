# Yandex ディスク

{% code fullWidth="true" %}
```
NAME:
   singularity storage create yandex - Yandex ディスク

USAGE:
   singularity storage create yandex [command options] [arguments...]

DESCRIPTION:
   --client-id
      OAuth クライアント ID。
      
      通常、空白のままにしておきます。

   --client-secret
      OAuth クライアントシークレット。
      
      通常、空白のままにしておきます。

   --token
      JSON Blob 形式の OAuth アクセストークン。

   --auth-url
      認証サーバーの URL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしておきます。

   --token-url
      トークンサーバーの URL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしておきます。

   --hard-delete
      ファイルをゴミ箱に入れるのではなく、完全に削除します。

   --encoding
      バックエンドのエンコード方式。
      
      詳細については [概要の "エンコーディング" セクション](/overview/#encoding) を参照してください。


OPTIONS:
   --client-id value      OAuth クライアント ID。 [$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示する

   Advanced

   --auth-url value   認証サーバーの URL。 [$AUTH_URL]
   --encoding value   バックエンドのエンコード方式。 (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete      ファイルをゴミ箱に入れるのではなく、完全に削除します。 (default: false) [$HARD_DELETE]
   --token value      JSON Blob 形式の OAuth アクセストークン。 [$TOKEN]
   --token-url value  トークンサーバーの URL。 [$TOKEN_URL]

   General

   --name value  ストレージの名前 (default: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}