# Amazon Drive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update acd - Amazon Drive

USAGE:
   singularity storage update acd [command options] <名前|ID>

DESCRIPTION:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにします。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --token
      JSON形式のOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダーのデフォルトを使用する場合は空白のままにします。

   --checkpoint
      内部ポーリングのためのチェックポイント（デバッグ用）。

   --upload-wait-per-gb
      失敗した完全なアップロードの後に、ファイルが表示されるか確認するために、1 GiBごとに追加する待機時間。
      
      Amazon Driveでは、ファイルが完全にアップロードされているにもかかわらず、少し時間が経ってからファイルが表示されることがあります。
      ファイルサイズが1 GiBを超える場合には、時々エラーが発生します。ファイルサイズが10 GiBを超える場合はほぼ必ずエラーが発生します。
      このパラメータは、ファイルが表示されるまでrcloneが待機する時間を制御します。
      
      このパラメータのデフォルト値は、1 GiBごとに3分です。つまり、デフォルトでは1 GiBごとに3分待機して、ファイルが表示されるか確認します。
      
      この機能を無効にするには、値を0に設定します。ただし、これによりrcloneが失敗したアップロードを再試行するため、
      衝突エラーが発生する可能性がありますが、ファイルは最終的に正しく表示されるはずです。
      
      これらの値は、さまざまなファイルサイズの大きなファイルを多数アップロードすることによって経験的に決定されました。
      
      この状況でrcloneが行っている操作についての詳細情報を表示するには、「-v」フラグを使用してアップロードします。

   --templink-threshold
      このサイズ以上のファイルは、tempLinkを使用してダウンロードされます。
      
      このサイズ以上のファイルは、"tempLink"を使用してダウンロードされます。これは、
      約10 GiBより大きいファイルのダウンロードをブロックするAmazon Driveの問題を回避するためです。
      この値のデフォルトは9 GiBで、変更する必要はありません。
      
      このしきい値を超えるファイルをダウンロードするために、rcloneは「tempLink」をリクエストし、
      サポートされているS3ストレージから一時的なURLを介してファイルを直接ダウンロードします。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示

   Advanced

   --auth-url value            認証サーバーのURL。 [$AUTH_URL]
   --checkpoint value          内部ポーリングのためのチェックポイント（デバッグ用）。 [$CHECKPOINT]
   --encoding value            バックエンドのエンコーディング（デフォルト："Slash,InvalidUtf8,Dot"）。 [$ENCODING]
   --templink-threshold value  このサイズ以上のファイルは、tempLinkを使用してダウンロードされます（デフォルト："9Gi"）。 [$TEMPLINK_THRESHOLD]
   --token value               JSON形式のOAuthアクセストークン。 [$TOKEN]
   --token-url value           トークンサーバーのURL。 [$TOKEN_URL]
   --upload-wait-per-gb value  失敗した完全なアップロードの後に、1 GiBごとに追加する待機時間（デフォルト："3m0s"）。 [$UPLOAD_WAIT_PER_GB]
```
{% endcode %}