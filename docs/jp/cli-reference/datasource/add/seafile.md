# seafile

{% code fullWidth="true" %}
```
名称:
   singularity datasource add seafile - seafile

使用方法:
   singularity datasource add seafile [コマンドオプション] <データセット名> <ソースパス>

説明:
   --seafile-2fa
      二要素認証（アカウントが2FAを使用している場合に 'true' を設定）。

   --seafile-auth-token
      認証トークン。

   --seafile-create-library
      ライブラリが存在しない場合に作成するかどうか。

   --seafile-encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --seafile-library
      ライブラリの名前。
      
      暗号化されていないライブラリ全体にアクセスするには、空白にしてください。

   --seafile-library-key
      ライブラリのパスワード（暗号化されたライブラリのみ）。
      
      コマンドラインでパスワードを渡す場合は、空白にしてください。

   --seafile-pass
      パスワード。

   --seafile-url
      接続するSeafileホストのURL。

      例:
         | https://cloud.seafile.com/ | cloud.seafile.comに接続します。

   --seafile-user
      ユーザー名（通常、メールアドレス）。

オプション:
   --help, -h  ヘルプを表示

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後のスキャンから指定したインターバルが経過した場合、自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期スキャン状態を設定します (デフォルト: ready)

   Seafileのオプション

   --seafile-2fa value             二要素認証（アカウントが2FAを使用している場合に 'true' を設定）(デフォルト: "false") [$SEAFILE_2FA]
   --seafile-create-library value  ライブラリが存在しない場合に作成するかどうか (デフォルト: "false") [$SEAFILE_CREATE_LIBRARY]
   --seafile-encoding value        バックエンドのエンコーディング (デフォルト: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$SEAFILE_ENCODING]
   --seafile-library value         ライブラリの名前 [$SEAFILE_LIBRARY]
   --seafile-library-key value     ライブラリのパスワード（暗号化されたライブラリのみ） [$SEAFILE_LIBRARY_KEY]
   --seafile-pass value            パスワード [$SEAFILE_PASS]
   --seafile-url value             SeafileホストのURL [$SEAFILE_URL]
   --seafile-user value            ユーザー名（通常、メールアドレス） [$SEAFILE_USER]
```
{% endcode %}