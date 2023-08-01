# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularityデータソースのWebDAVの追加 - WebDAV

使用法:
   singularityデータソースのWebDAVの追加 [コマンドオプション] <データセット名> <ソースパス>

説明:
   --webdav-bearer-token
      ユーザー/パスワードの代わりにベアラートークンを使用します（例：Macaroon）。

   --webdav-bearer-token-command
      ベアラートークンを取得するために実行するコマンド。

   --webdav-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。
      
      デフォルトのエンコーディングは、sharepoint-ntlmまたはidentityの場合はSlash、LtGt、DoubleQuote、Colon、Question、Asterisk、Pipe、Hash、Percent、BackSlash、Del、Ctl、LeftSpace、LeftTilde、RightSpace、RightPeriod、InvalidUtf8です。

   --webdav-headers
      すべてのトランザクションに対してHTTPヘッダーを設定します。
      
      これを使用して、すべてのトランザクションに対して追加のHTTPヘッダーを設定できます。
      
      入力形式は、キー、値のペアのコンマ区切りリストです。 標準の[CSVエンコーディング](https://godoc.org/encoding/csv)が使用できます。
      
      たとえば、Cookieを設定する場合は 'Cookie,name=value' や '"Cookie","name=value"' を使用します。
      
      複数のヘッダーを設定することもできます。 例: '"Cookie","name=value","Authorization","xxx"'。

   --webdav-pass
      パスワード。

   --webdav-url
      HTTPホストのURLに接続します。
      
      例: https://example.com。

   --webdav-user
      ユーザー名。
      
      NTLM認証を使用する場合、ユーザー名は 'Domain\User' の形式である必要があります。

   --webdav-vendor
      使用しているWebDAVサイト/サービス/ソフトウェアの名前。

      例:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online、Microsoftアカウントで認証される
         | sharepoint-ntlm | Sharepoint、通常は自己ホストまたはオンプレミスでNTLM認証を使用する
         | other           | その他のサイト/サービスまたはソフトウェア


オプション:
   --help、-h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval 値  最後の正常なスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state 値   初期のスキャン状態を設定します (デフォルト: ready)

   WebDAVのオプション

   --webdav-bearer-token 値          ユーザー/パスワードの代わりにベアラートークンを使用します（例：Macaroon）。 [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command 値  ベアラートークンを取得するために実行するコマンド。 [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding 値              バックエンドのエンコーディング。 [$WEBDAV_ENCODING]
   --webdav-headers 値               すべてのトランザクションに対してHTTPヘッダーを設定します。 [$WEBDAV_HEADERS]
   --webdav-pass 値                  パスワード。 [$WEBDAV_PASS]
   --webdav-url 値                   HTTPホストのURLに接続します。 [$WEBDAV_URL]
   --webdav-user 値                  ユーザー名。 [$WEBDAV_USER]
   --webdav-vendor 値                使用しているWebDAVサイト/サービス/ソフトウェアの名前。 [$WEBDAV_VENDOR]

```
{% endcode %}