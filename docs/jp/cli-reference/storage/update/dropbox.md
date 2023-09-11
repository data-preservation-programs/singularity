# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity storage update dropbox - Dropbox

USAGE:
   singularity storage update dropbox [コマンドオプション] <名前|ID>

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
      
      プロバイダのデフォルトを使用するには空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用するには空白のままにします。

   --chunk-size
      アップロードのチャンクサイズ（150Mi未満）。
      
      このサイズを上回るファイルは、このサイズのチャンクでアップロードされます。
      
      チャンクはメモリ上でバッファリングされます（1つずつ）ので、rcloneはリトライを処理できます。
      この値を大きくすると、速度がわずかに向上します（テストにおいて最大10%、128MiBの場合）が、
      メモリの使用量が増加します。メモリが不足している場合は、この値を小さく設定できます。

   --impersonate
      ビジネスアカウントを使用する場合に、このユーザーの権限で実行します。
      
      注意: "rclone config"を実行する際に、このフラグが設定されていることを確認する必要があります。
      これにより、rcloneは通常要求しない"members.read"スコープを要求します。
      これが必要です。このスコープを使用するには、Dropbox Team Adminの承認が必要です。
      
      特定の共有フォルダの最初の使用以降、--dropbox-shared-foldersフラグは省略できます。

   --shared-files
      個々の共有ファイルでrcloneを動作させます。
      
      このモードでは、rcloneの機能は非常に限定的です。
      一覧表示（ls、lslなど）および読み取り操作（ダウンロードなど）のみがこのモードでサポートされます。
      その他の操作は無効になります。

   --shared-folders
      共有フォルダでrcloneを動作させます。
            
      このフラグを使用すると、パスが指定されていない場合はリスト操作のみサポートされ、
      利用可能なすべての共有フォルダがリストされます。
      パスを指定する場合、最初の部分は共有フォルダの名前と解釈されます。
      rcloneはこの共有フォルダをルート名前空間にマウントしようとします。マウントが成功すると、
      共有フォルダはほぼ通常のフォルダとなり、通常のすべての操作がサポートされます。
      
      共有フォルダはその後もアンマウントされないため、特定の共有フォルダの最初の使用後に
      --dropbox-shared-foldersを省略できます。

   --batch-mode
      ファイルのバッチアップロードモードを設定します。
      sync|async|offを指定できます。
      
      詳細については、[メインドキュメント](https://rclone.org/dropbox/#batch-mode)を参照してください。
      
      3つの可能な値があります。
      
      - off - バッチ処理なし
      - sync - バッチアップロードと完了確認（デフォルト）
      - async - バッチアップロードと完了確認なし
      
      Rcloneは終了時に保留中のバッチを閉じるため、終了に遅延が生じる場合があります。
      

   --batch-size
      アップロードバッチ内の最大ファイル数。
      
      この設定によってアップロードするファイルのバッチサイズが設定されます。1000未満である必要があります。
      
      デフォルトは0で、バッチモードの設定に応じてバッチサイズが計算されます。
      
      - batch_mode: async - デフォルトのbatch_sizeは100です。
      - batch_mode: sync - デフォルトのbatch_sizeは--transfersと同じです。
      - batch_mode: off - 使用されません。
      
      Rcloneは終了時に保留中のバッチを閉じるため、終了に遅延が生じる場合があります。
      
      小さいファイルを大量にアップロードする場合、
      この設定を行うことは非常に良いアイデアです。これにより、アップロードが大幅に高速化されます。
      最大スループットを実現するには、--transfers 32を使用できます。
      

   --batch-timeout
      アップロードバッチがアップロードする前にアイドル状態であることを許可する最大時間。
      
      アップロードバッチがこのより長い時間アイドル状態の場合、アップロードが開始されます。
      
      デフォルトは0で、rcloneは使用中のbatch_modeに基づいて適切なデフォルト値を選択します。
      
      - batch_mode: async - デフォルトのbatch_timeoutは500msです。
      - batch_mode: sync - デフォルトのbatch_timeoutは10sです。
      - batch_mode: off - 使用されません。
      

   --batch-commit-timeout
      バッチの終了を待つ最大時間

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuthクライアントID。[$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示

   アドバンスオプション

   --auth-url value              認証サーバーのURL。[$AUTH_URL]
   --batch-commit-timeout value  バッチの終了を待つ最大時間（デフォルト: "10m0s"）[$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            ファイルのバッチアップロードモードを設定します。sync|async|off（デフォルト: "sync"）[$BATCH_MODE]
   --batch-size value            アップロードバッチ内の最大ファイル数（デフォルト: 0）[$BATCH_SIZE]
   --batch-timeout value         アップロードバッチがアップロードする前にアイドル状態であることを許可する最大時間（デフォルト: "0s"）[$BATCH_TIMEOUT]
   --chunk-size value            アップロードのチャンクサイズ（150Mi未満）（デフォルト: "48Mi"）[$CHUNK_SIZE]
   --encoding value              バックエンドのエンコーディング（デフォルト: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"）[$ENCODING]
   --impersonate value           ビジネスアカウントを使用する場合に、このユーザーの権限で実行します。[$IMPERSONATE]
   --shared-files                個々の共有ファイルでrcloneを動作させます。（デフォルト: false）[$SHARED_FILES]
   --shared-folders              共有フォルダでrcloneを動作させます。（デフォルト: false）[$SHARED_FOLDERS]
   --token value                 JSON形式のOAuthアクセストークン。[$TOKEN]
   --token-url value             トークンサーバーのURL。[$TOKEN_URL]

```
{% endcode %}