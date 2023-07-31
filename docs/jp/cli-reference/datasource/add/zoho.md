# Zoho

{% code fullWidth="true" %}
```
名前：
   "singularity datasource add zoho" - Zoho

使用法：
   singularity datasource add zoho [コマンドオプション] <データセット名> <ソースパス>

説明：
   --zoho-auth-url
      認証サーバーのURLです。
      
      デフォルトのプロバイダーを使用する場合は空白にしてください。

   --zoho-client-id
      OAuthクライアントIDです。
      
      通常は空白のままにしてください。

   --zoho-client-secret
      OAuthクライアントシークレットです。
      
      通常は空白のままにしてください。

   --zoho-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --zoho-region
      接続するZohoリージョンです。
      
      所属する組織が登録されているリージョンを使用する必要があります。確認できない場合は、通常ブラウザで接続するトップレベルドメインを使用してください。

      例：
         | com    | アメリカ合衆国 / グローバル
         | eu     | ヨーロッパ
         | in     | インド
         | jp     | 日本
         | com.cn | 中国
         | com.au | オーストラリア

   --zoho-token
      OAuthアクセストークンをJSON blobとして指定します。

   --zoho-token-url
      トークンサーバーのURLです。
      
      デフォルトのプロバイダーを使用する場合は空白にしてください。


オプション：
   --help, -h  ヘルプを表示

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  前回のスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   Zoho用のオプション

   --zoho-auth-url value       認証サーバーのURLです。[$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuthクライアントIDです。[$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuthクライアントシークレットです。[$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       バックエンドのエンコーディングです。 (デフォルト: "Del,Ctl,InvalidUtf8") [$ZOHO_ENCODING]
   --zoho-region value         接続するZohoリージョンです。[$ZOHO_REGION]
   --zoho-token value          OAuthアクセストークンをJSON blobとして指定します。[$ZOHO_TOKEN]
   --zoho-token-url value      トークンサーバーのURLです。[$ZOHO_TOKEN_URL]

```
{% endcode %}