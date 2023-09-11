# seafile

{% code fullWidth="true" %}
```
名前:
   singularity storage create seafile - seafile

使用法:
   singularity storage create seafile [コマンドオプション] [引数...]

説明:
   --url
      接続するSeafileホストのURL。

      例:
         | https://cloud.seafile.com/ | cloud.seafile.comに接続する。

   --user
      ユーザー名（通常はメールアドレス）。

   --pass
      パスワード。

   --2fa
      2要素認証（アカウントに2FAが有効なら 'true'）。

   --library
      ライブラリの名前。

      暗号化されていないライブラリにアクセスするには、空白のままにしておきます。

   --library-key
      ライブラリのパスワード（暗号化されたライブラリの場合）。

      コマンドラインで渡す場合は、空白のままにしておきます。

   --create-library
      ライブラリが存在しない場合、rcloneがライブラリを作成するかどうか。

   --auth-token
      認証トークン。

   --encoding
      バックエンドのエンコーディング。

      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

オプション:
   --2fa                2要素認証（アカウントに2FAが有効なら 'true'）。（デフォルト: false） [$2FA]
   --auth-token value   認証トークン。 [$AUTH_TOKEN]
   --help, -h           ヘルプを表示
   --library value      ライブラリの名前。 [$LIBRARY]
   --library-key value  ライブラリのパスワード（暗号化されたライブラリの場合）。[$LIBRARY_KEY]
   --pass value         パスワード。 [$PASS]
   --url value          接続するSeafileホストのURL。 [$URL]
   --user value         ユーザー名（通常はメールアドレス）。[$USER]

   Advanced（詳細）

   --create-library  ライブラリが存在しない場合、rcloneがライブラリを作成するかどうか。（デフォルト: false） [$CREATE_LIBRARY]
   --encoding value  バックエンドのエンコーディング。（デフォルト: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8"）[$ENCODING]

   General（一般）

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}