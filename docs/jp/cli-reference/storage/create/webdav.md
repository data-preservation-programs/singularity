# WebDAV

{% code fullWidth="true" %}
```
名前:
   singularity storage create webdav - WebDAV

使用法:
   singularity storage create webdav [コマンドオプション] [引数...]

説明:
   --url
      接続するhttpホストのURLです。
      
      例: https://example.com.

   --vendor
      使用しているWebDAVサイト/サービス/ソフトウェアの名前です。

      例:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online,Microsoftアカウントで認証されます
         | sharepoint-ntlm | Sharepoint,NTLM認証で通常は自己ホストまたはオンプレミス
         | other           | その他のサイト/サービスまたはソフトウェア

   --user
      ユーザー名です。

      NTLM認証を使用する場合、ユーザー名は 'ドメイン\ユーザー' の形式である必要があります。

   --pass
      パスワードです。

   --bearer-token
      ユーザー/パスワードの代わりにベアラトークン（Macaroonなど）を使用します。

   --bearer-token-command
      ベアラトークンを取得するコマンドです。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

      sharepoint-ntlmの場合は、デフォルトのエンコーディングはSlash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8です。
      それ以外の場合はidentityです。

   --headers
      すべてのトランザクションに対してHTTPヘッダーを設定します。

      これを使用してすべてのトランザクションに対して追加のHTTPヘッダーを設定します。

      入力形式は、キーと値のペアのコンマ区切りのリストです。標準の
      [CSVエンコーディング](https://godoc.org/encoding/csv)が使用できます。

      たとえば、Cookieを設定する場合は 'Cookie,name=value' または '"Cookie","name=value"' を使用します。

      複数のヘッダーを設定することもできます。例: '"Cookie","name=value","Authorization","xxx"'。
      

オプション:
   --bearer-token value  ユーザー/パスワードの代わりにベアラトークン（Macaroonなど）を使用します。[$BEARER_TOKEN]
   --help, -h            ヘルプを表示します
   --pass value          パスワード。[$PASS]
   --url value           接続するhttpホストのURL。[$URL]
   --user value          ユーザー名。[$USER]
   --vendor value        使用しているWebDAVサイト/サービス/ソフトウェアの名前。[$VENDOR]

   Advanced

   --bearer-token-command value  ベアラトークンを取得するコマンド。[$BEARER_TOKEN_COMMAND]
   --encoding value              バックエンドのエンコーディング。[$ENCODING]
   --headers value               すべてのトランザクションに対してHTTPヘッダーを設定。[$HEADERS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}