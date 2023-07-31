# Amazon Drive

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add acd - Amazon Drive

使用方法:
   singularity datasource add acd [コマンドオプション] <データセット名> <ソースパス>

説明:
   --acd-auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --acd-checkpoint
      内部ポーリングのチェックポイント（デバッグ用）。

   --acd-client-id
      OAuthクライアントID。
      
      通常は空白のままにします。

   --acd-client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --acd-encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --acd-templink-threshold
      このサイズ以上のファイルはtempLinkを使用してダウンロードされます。
      
      このサイズ以上のファイルは「tempLink」を使用してダウンロードされます。
      これは、Amazon Driveが約10 GiBより大きいファイルのダウンロードをブロックする問題を回避するためです。
      このデフォルト値は9 GiBで、変更する必要はありません。
      
      このしきい値以上のファイルをダウンロードするために、rcloneは「tempLink」を要求し、基礎となるS3ストレージから一時URLを介してファイルをダウンロードします。

   --acd-token
      JSON形式のOAuthアクセストークン。

   --acd-token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --acd-upload-wait-per-gb
      完全なアップロードが失敗した後、GiBごとに待機してから表示されるまでの追加時間。
      
      時々、ファイルが完全にアップロードされているのにもかかわらず、Amazon Driveでエラーが発生する場合がありますが、しばらくするとファイルが表示されます。
      これは、サイズが1 GiBを超えるファイルの場合には時々発生し、10 GiBより大きいファイルの場合はほぼ毎回発生します。
      このパラメータは、ファイルが表示されるまでrcloneが待機する時間を制御します。
      
      このパラメータのデフォルト値はGiBごとに3分ですので、デフォルトでは、GiBごとに3分間待機して、ファイルが表示されるかどうかを確認します。
      
      これを無効にするには、0に設定します。これにより、rcloneは失敗したアップロードを再試行しますが、ファイルは最終的に正しく表示される可能性があります。
      
      これらの値は、さまざまなファイルサイズの大きなファイルのアップロードを観察して経験的に決定されました。
      
      この状況でrcloneが何を行っているかについての詳細情報を表示するには、"-v"フラグを使用してアップロードしてください。


オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データを削除します。  (デフォルト: false)
   --rescan-interval value  前回の成功したスキャンからこの間隔が経過したときに自動的にソースディレクトリを再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   acd用のオプション

   --acd-auth-url value            認証サーバーのURL。 [$ACD_AUTH_URL]
   --acd-client-id value           OAuthクライアントID。 [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuthクライアントシークレット。 [$ACD_CLIENT_SECRET]
   --acd-encoding value            バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  このサイズ以上のファイルはtempLinkを使用してダウンロードされます。 (デフォルト: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               JSON形式のOAuthアクセストークン。 [$ACD_TOKEN]
   --acd-token-url value           トークンサーバーのURL。 [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  完全なアップロードが失敗した後、GiBごとに待機してから表示されるまでの追加時間。 (デフォルト: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]

```
{% endcode %}