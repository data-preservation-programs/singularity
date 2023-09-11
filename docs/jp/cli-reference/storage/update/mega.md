# Mega

{% code fullWidth="true" %}
```
名前:
   singularity storage update mega - Mega

使い方:
   singularity storage update mega [コマンドオプション] <名前|ID>

説明:
   --user
      ユーザー名。

   --pass
      パスワード。

   --debug
      Megaからより多くのデバッグ情報を出力します。
      
      このフラグが設定されている場合（vvと共に指定されている場合）、メガバックエンドからさらなるデバッグ情報が表示されます。

   --hard-delete
      ファイルをゴミ箱に入れずに完全に削除します。
      
      通常、メガバックエンドは削除したすべてのオブジェクトをゴミ箱に入れるか、完全に削除するかを選択します。このフラグを指定すると、rcloneはオブジェクトを完全に削除します。

   --use-https
      転送にHTTPSを使用します。
      
      MEGAはデフォルトでプレーンテキストのHTTP接続を使用します。一部のISPはHTTP接続を制限するため、転送が非常に遅くなります。これを有効にすると、MEGAはすべての転送にHTTPSを使用します。HTTPSは通常必要ありません。これはデータがすでに暗号化されているためです。そのため、CPU使用率が増加し、ネットワークのオーバーヘッドが追加されます。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h    ヘルプを表示
   --pass value  パスワード。[$PASS]
   --user value  ユーザー名。[$USER]

   高度なオプション:

   --debug           Megaからより多くのデバッグ情報を出力します。（デフォルト：false）[$DEBUG]
   --encoding value  バックエンドのエンコーディングです。（デフォルト："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --hard-delete     ファイルをゴミ箱に入れずに完全に削除します。（デフォルト：false）[$HARD_DELETE]
   --use-https       転送にHTTPSを使用します。（デフォルト：false）[$USE_HTTPS]

```
{% endcode %}