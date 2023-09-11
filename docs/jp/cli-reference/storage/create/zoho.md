# Zoho

{% code fullWidth="true" %}
```
名前:
   singularity storage create zoho - Zoho

使い方:
   singularity storage create zoho [コマンドオプション] [引数...]

説明:
   --client-id
      OAuthクライアントID。
      
      通常、空白のままにしてください。

   --client-secret
      OAuthクライアントシークレット。
      
      通常、空白のままにしてください。

   --token
      JSONデータとしてのOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は、空白のままにしてください。

   --token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は、空白のままにしてください。

   --region
      接続するZohoリージョン。
      
      組織が登録されているリージョンを使用する必要があります。自信がない場合は、ブラウザで接続するトップレベルドメインと同じものを使用してください。

      例:
         | com    | アメリカ合衆国 / グローバル
         | eu     | ヨーロッパ
         | in     | インド
         | jp     | 日本
         | com.cn | 中国
         | com.au | オーストラリア

   --encoding
      バックエンドのエンコーディング。
      
      詳細は、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --client-id value      OAuthクライアントID。[$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --region value         接続するZohoリージョン。[$REGION]

   高度

   --auth-url value   認証サーバーのURL。[$AUTH_URL]
   --encoding value   バックエンドのエンコーディング。(デフォルト: "Del,Ctl,InvalidUtf8")[$ENCODING]
   --token value      JSONデータとしてのOAuthアクセストークン。[$TOKEN]
   --token-url value  トークンサーバーのURL。[$TOKEN_URL]

   一般

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス
```
{% endcode %}