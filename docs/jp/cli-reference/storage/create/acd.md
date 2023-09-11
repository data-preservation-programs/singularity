# Amazon Drive

{% code fullWidth="true" %}
```
名前:
   シンギュラリティ ストレージの作成 acd - Amazon Drive

使用法:
   シンギュラリティ ストレージの作成 acd [コマンドオプション] [引数...]

説明:
   --client-id
      OAuth クライアント ID。
      
      通常は空白のままにしておきます。

   --client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにしておきます。

   --token
      JSON 形式の OAuth アクセストークン。

   --auth-url
      認証サーバーの URL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --token-url
      トークンサーバーの URL。
      
      プロバイダのデフォルトを使用する場合は空白のままにしてください。

   --checkpoint
      内部ポーリングのチェックポイント (デバッグ用)。

   --upload-wait-per-gb
      失敗した完全なアップロード後に見かけるまで待機するための GiB ごとの追加時間。
      
      1 GiB より大きいファイルや、特に 10 GiB より大きいファイルでは、
      アップロードが完了したにもかかわらず Amazon Drive でエラーが発生する場合がありますが、
      しばらくするとファイルが表示されることがあります。このパラメータは、
      ファイルが表示されるまで rclone が待機する時間を制御します。
      
      このパラメータのデフォルト値は、1 GiB あたり 3 分ですので、
      デフォルトでは、1 GiB アップロードするたびに 3 分待機します。
      
      この機能を無効にするには、値を 0 に設定します。これにより、
      失敗したアップロードが rclone によって再試行されますが、ファイルは
      最終的に正しく表示されるはずです。
      
      これらの値は、さまざまなファイルサイズの大きなファイルのアップロードを観察することで、
      経験的に決定されました。
      
      この状況での rclone の動作についての詳細情報を表示するには、
      "-v" フラグを使用してアップロードしてください。

   --templink-threshold
      このサイズ以上のファイルは、tempLink を介してダウンロードされます。
      
      この以上のサイズのファイルは、"tempLink" を介してダウンロードされます。これは、
      約 10 GiB を超えるファイルのダウンロードを阻止する問題を解決するためのものです。これについてのデフォルト値は、
      9 GiB ですが、変更する必要はありません。
      
      この閾値を超えるファイルをダウンロードするには、rclone は「tempLink」を要求し、
      基礎となる S3 ストレージから直接一時的な URL を使用してファイルをダウンロードします。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --client-id value      OAuth クライアント ID。[$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示

   ローカル
    
   --auth-url value            認証サーバーの URL。[$AUTH_URL]
   --checkpoint value          内部ポーリングのチェックポイント (デバッグ用)。[$CHECKPOINT]
   --encoding value            バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --templink-threshold value  このサイズ以上のファイルは、tempLink を介してダウンロードされます。 (デフォルト: "9Gi") [$TEMPLINK_THRESHOLD]
   --token value               JSON 形式の OAuth アクセストークン。[$TOKEN]
   --token-url value           トークンサーバーの URL。[$TOKEN_URL]
   --upload-wait-per-gb value  失敗した完全なアップロード後に見かけるまで待機するための GiB ごとの追加時間。 (デフォルト: "3m0s") [$UPLOAD_WAIT_PER_GB]

   一般的

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}