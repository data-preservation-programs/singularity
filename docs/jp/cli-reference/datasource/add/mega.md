# Mega

{% code fullWidth="true" %}
```
名称：
   singularity datasource add mega - Mega

使用方法：
   singularity datasource add mega [コマンドオプション] <データセット名> <ソースパス>

説明：
   --mega-debug
      Megaからのより詳細なデバッグ情報の出力。
      
      このフラグが設定されている場合（-vvとともに）、megaバックエンドからさらにデバッグ情報が出力されます。

   --mega-encoding
      バックエンドのエンコーディング。
      
      より詳細な情報については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --mega-hard-delete
      ファイルをゴミ箱に置くのではなく、永久に削除します。
      
      通常、メガバックエンドは、削除されたすべてのものをゴミ箱に置くため、永久に削除しません。このオプションを指定すると、rcloneはオブジェクトを永久に削除します。

   --mega-pass
      パスワード。

   --mega-use-https
      転送にHTTPSを使用します。
      
      MEGAはデフォルトで平文のHTTP接続を使用しています。一部のISPはHTTP接続を制限するため、転送は非常に遅くなります。これを有効にすると、MEGAはすべての転送にHTTPSを使用するようになります。HTTPSは通常は不要ですが、すべてのデータは既に暗号化されているためです。このオプションを有効にすると、CPU使用率が上がり、ネットワークオーバーヘッドが発生します。

   --mega-user
      ユーザー名。


オプション：
   --help、-h  ヘルプを表示します。

   データの準備オプション：

   --delete-after-export    [危険] データセットをCARファイルにエクスポート後、データセットのファイルを削除します（デフォルト：無効）。
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過した場合、ソースディレクトリを自動的に再スキャンします（デフォルト：無効）。
   --scanning-state value   初期スキャン状態を設定します（デフォルト：ready）。

   mega用オプション：

   --mega-debug value        Megaからのより詳細なデバッグ情報の出力（デフォルト：false）[$MEGA_DEBUG]
   --mega-encoding value     バックエンドのエンコーディング（デフォルト：Slash,InvalidUtf8,Dot）[$MEGA_ENCODING]
   --mega-hard-delete value  ファイルをゴミ箱に置くのではなく、永久に削除します（デフォルト：false）[$MEGA_HARD_DELETE]
   --mega-pass value         パスワード。[$MEGA_PASS]
   --mega-use-https value    転送にHTTPSを使用します（デフォルト：false）[$MEGA_USE_HTTPS]
   --mega-user value         ユーザー名。[$MEGA_USER]

```
{% endcode %}