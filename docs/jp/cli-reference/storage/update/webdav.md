# WebDAV

{% code fullWidth="true" %}
```
名前:
   singularity storage update webdav - WebDAV

使用法:
   singularity storage update webdav [コマンドオプション] <名前|ID>

説明:
   --url
      接続するHTTPホストのURL。
      
      例: https://example.com.

   --vendor
      使用しているWebDAVサイト/サービス/ソフトウェアの名前。

      例:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, Microsoftアカウントにより認証される
         | sharepoint-ntlm | NTLM認証を使用するSharepoint、通常は自己ホストまたはオンプレミス
         | other           | その他のサイト/サービスまたはソフトウェア

   --user
      ユーザー名。
      
      NTLM認証を使用する場合、ユーザー名は 'ドメイン\ユーザー' の形式である必要があります。

   --pass
      パスワード。

   --bearer-token
      ユーザー名とパスワードの代わりにベアラートークンを使用します（例: Macaroon）。

   --bearer-token-command
      ベアラートークンを取得するためのコマンドの実行方法。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。
      
      sharepoint-ntlmの場合は、デフォルトのエンコーディングはSlash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8です。

   --headers
      すべてのトランザクションに対してHTTPヘッダーを設定します。
      
      これを使用して、すべてのトランザクションに対して追加のHTTPヘッダーを設定します。
      
      入力形式は、キー、値のペアのカンマ区切りリストです。[CSVエンコーディング](https://godoc.org/encoding/csv)も使用できます。
      
      たとえば、Cookieを設定する場合は 'Cookie,name=value' または '"Cookie","name=value"' を使用します。
      
      複数のヘッダーを設定することもできます。例: '"Cookie","name=value","Authorization","xxx"'。

オプション:
   --bearer-token value  ユーザー名とパスワードの代わりにベアラートークンを使用します（例: Macaroon）。[$BEARER_TOKEN]
   --help, -h            ヘルプを表示します
   --pass value          パスワード。[$PASS]
   --url value           接続するHTTPホストのURL。[$URL]
   --user value          ユーザー名。[$USER]
   --vendor value        使用しているWebDAVサイト/サービス/ソフトウェアの名前。[$VENDOR]

   高度なオプション

   --bearer-token-command value  ベアラートークンを取得するためのコマンドの実行方法。[$BEARER_TOKEN_COMMAND]
   --encoding value              バックエンドのエンコーディング。[$ENCODING]
   --headers value               すべてのトランザクションに対してHTTPヘッダーを設定します。[$HEADERS]

```
{% endcode %}