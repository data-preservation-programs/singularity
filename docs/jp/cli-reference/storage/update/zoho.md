# Zoho

{% code fullWidth="true" %}
```
NAME:
   singularity storage update zoho - Zoho

USAGE:
   singularity storage update zoho [コマンドオプション] <名前|ID>

DESCRIPTION:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにします。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --token
      JSONブロブとしてのOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用するには空白にします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用するには空白にします。

   --region
      接続するZohoのリージョン。
      
      お使いの組織が登録されているリージョンを使用する必要があります。自信がない場合は、ブラウザで接続するときと同じトップレベルドメインを使用してください。

      例:
         | com    | アメリカ / グローバル
         | eu     | ヨーロッパ
         | in     | インド
         | jp     | 日本
         | com.cn | 中国
         | com.au | オーストラリア

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプの表示
   --region value         接続するZohoのリージョン。 [$REGION]

   Advanced

   --auth-url value   認証サーバーのURL。 [$AUTH_URL]
   --encoding value   バックエンドのエンコーディング。 (デフォルト: "Del,Ctl,InvalidUtf8") [$ENCODING]
   --token value      JSONブロブとしてのOAuthアクセストークン。 [$TOKEN]
   --token-url value  トークンサーバーのURL。 [$TOKEN_URL]

```
{% endcode %}