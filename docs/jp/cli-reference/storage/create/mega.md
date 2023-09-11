# Mega

{% code fullWidth="true" %}
```
名前:
   singularity storage create mega - Mega

使用方法:
   singularity storage create mega [コマンドオプション] [引数...]

説明:
   --user
      ユーザー名。

   --pass
      パスワード。

   --debug
      Megaからのより詳細なデバッグ情報を出力します。

      このフラグが設定されている場合（-vvと共に設定されている場合）、
      megaのバックエンドからさらに詳細なデバッグ情報が表示されます。

   --hard-delete
      ファイルをごみ箱ではなく完全に削除します。

      通常、Megaバックエンドはすべての削除をごみ箱に入れるため、
      永久削除する代わりに削除します。
      これを指定すると、rcloneはオブジェクトを完全に削除します。

   --use-https
      転送にHTTPSを使用します。

      MEGAはデフォルトで平文のHTTP接続を使用します。
      一部のISPはHTTP接続を制限するため、転送が非常に遅くなることがあります。
      これを有効にすると、MEGAはすべての転送のためにHTTPSを使用します。
      HTTPSは通常は必要ありませんが、すでにすべてのデータが暗号化されているためです。
      これを有効にすると、CPU使用率が増加し、ネットワークオーバーヘッドが追加されます。

   --encoding
      バックエンドのエンコーディング。

      詳細については [エンコーディングセクションの概要](/overview/#encoding) を参照してください。


オプション:
   --help, -h    ヘルプを表示
   --pass value  パスワード。 [$PASS]
   --user value  ユーザー名。 [$USER]

   Advanced（拡張オプション）

   --debug           Megaからのより詳細なデバッグ情報を出力します。 (デフォルト: false) [$DEBUG]
   --encoding value  バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete     ファイルをごみ箱ではなく完全に削除します。 (デフォルト: false) [$HARD_DELETE]
   --use-https       転送にHTTPSを使用します。 (デフォルト: false) [$USE_HTTPS]

   General（一般オプション）

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}