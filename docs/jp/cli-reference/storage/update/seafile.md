# seafile

{% code fullWidth="true" %}
```
NAME:
   singularity storage update seafile - seafile

USAGE:
   singularity storage update seafile [command options] <name|id>

DESCRIPTION:
   --url
      接続先のseafileホストのURL。

      例:
         | https://cloud.seafile.com/ | cloud.seafile.comに接続します。

   --user
      ユーザー名（通常はメールアドレス）。

   --pass
      パスワード。

   --2fa
      二要素認証（2FAが有効な場合は'true'）。

   --library
      ライブラリの名称。
      
      空白の場合、暗号化されていないすべてのライブラリにアクセスします。

   --library-key
      ライブラリのパスワード（暗号化されたライブラリの場合）。
      
      コマンドラインで指定する場合は空白のままにします。

   --create-library
      ライブラリが存在しない場合にrcloneがライブラリを作成するかどうか。

   --auth-token
      認証トークン。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --2fa                二要素認証（2FAが有効な場合は'true'）。 (default: false) [$2FA]
   --auth-token value   認証トークン。 [$AUTH_TOKEN]
   --help, -h           ヘルプを表示
   --library value      ライブラリの名称。 [$LIBRARY]
   --library-key value  ライブラリのパスワード（暗号化されたライブラリの場合）。 [$LIBRARY_KEY]
   --pass value         パスワード。 [$PASS]
   --url value          接続先のseafileホストのURL。 [$URL]
   --user value         ユーザー名（通常はメールアドレス）。 [$USER]

   Advanced

   --create-library  ライブラリが存在しない場合にrcloneがライブラリを作成するかどうか。 (default: false) [$CREATE_LIBRARY]
   --encoding value  バックエンドのエンコーディング。 (default: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$ENCODING]

```
{% endcode %}