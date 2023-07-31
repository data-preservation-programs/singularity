# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add dropbox - Dropbox

USAGE:
   singularity datasource add dropbox [command options] <dataset_name> <source_path>

DESCRIPTION:
   --dropbox-auth-url
      認証サーバーのURL。
      
      デフォルトのプロバイダーの設定を使用する場合は空白のままにしてください。

   --dropbox-batch-commit-timeout
      バッチの終了を待つ最大時間

   --dropbox-batch-mode
      ファイルのバッチ同期のアップロードモードを設定します（sync|async|off）。
      
      これはrcloneで使用するバッチモードを設定します。
      
      詳細については[メインのドキュメント](https://rclone.org/dropbox/#batch-mode)を参照してください。
      
      3つの値が指定できます。
      
      - off - バッチ処理なし
      - sync - バッチでのアップロードと完了チェック（デフォルト）
      - async - バッチでのアップロードと完了チェックなし
      
      rcloneは終了時に未完成のバッチをクローズしますが、これにより終了に遅延が発生する可能性があります。
      

   --dropbox-batch-size
      アップロードバッチ内の最大ファイル数。
      
      これはアップロードするファイルのバッチサイズを設定します。1000未満である必要があります。
      
      デフォルトでは、これは0であり、rcloneはbatch_modeの設定に基づいてバッチサイズを計算します。
      
      - batch_mode: async - デフォルトのbatch_sizeは100
      - batch_mode: sync - デフォルトのbatch_sizeは--transfersと同じ
      - batch_mode: off - 使用しない
      
      rcloneは終了時に未完成のバッチをクローズしますが、これにより終了に遅延が発生する可能性があります。
      
      これは、多くの小さなファイルをアップロードする場合には非常に良いアイデアです。--transfers 32を使用するとスループットを最大にすることができます。
      

   --dropbox-batch-timeout
      アップロード前のアップロードバッチのアイドル時間の最大許容時間。
      
      アップロードバッチがこの時間よりも長い間アイドル状態が続くと、アップロードされます。
      
      デフォルトは0であり、rcloneは使用中のbatch_modeに基づいて適切なデフォルトを選択します。
      
      - batch_mode: async - デフォルトのbatch_timeoutは500ミリ秒です
      - batch_mode: sync - デフォルトのbatch_timeoutは10秒です
      - batch_mode: off - 使用しない
      

   --dropbox-chunk-size
      アップロードチャンクのサイズ（< 150Mi）。
      
      このサイズを超えるファイルは、このサイズのチャンクでアップロードされます。
      
      チャンクはメモリ内でバッファリングされます（1つずつ）、そのためrcloneはリトライ時に対応できます。これを大きく設定すると、スピードがわずかに向上します（128 MiBの場合、最大10%のテスト結果）が、メモリ使用量が増えます。メモリに制約がある場合は、これをより小さく設定することができます。

   --dropbox-client-id
      OAuthクライアントID。
      
      通常は空白のままにしてください。

   --dropbox-client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにしてください。

   --dropbox-encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --dropbox-impersonate
      ビジネスアカウントを使用する際に、このユーザーになりすます。
      
      使用する場合、"rclone config"を実行する際にこのフラグが設定されていることを確認する必要があります。これにより、rcloneが通常はリクエストしない"members.read"スコープを要求します。これはdropboxがAPIで使用する内部IDにメンバーのメールアドレスを検索するために必要です。
      
      "members.read"スコープを使用するには、OAuthフロー中にDropbox Team Adminの承認が必要です。
      
      このオプションを使用するには、独自のアプリ（独自のclient_idとclient_secretを設定する）を使用する必要があります。現在のrcloneのデフォルトの権限セットに"members.read"が含まれていないため、v1.55以降がどこでも使用されるようになるまで、これを追加することができます。

   --dropbox-shared-files
      個別の共有ファイルで作業するようにrcloneに指示します。
      
      このモードでは、rcloneの機能は非常に制限されます - リスト（ls、lslなど）操作および読み取り操作（ダウンロードなど）のみがこのモードでサポートされます。他のすべての操作は無効になります。

   --dropbox-shared-folders
      共有フォルダで作業するようにrcloneに指示します。
            
      このフラグをパスなしで使用すると、リスト操作のみがサポートされ、利用可能なすべての共有フォルダがリストされます。パスを指定すると、最初の部分は共有フォルダの名前として解釈されます。rcloneはこの共有フォルダをルート名前空間にマウントしようとします。成功した場合、共有フォルダは通常のフォルダとほぼ同じですが、通常のすべての操作がサポートされます。
      
      使用後に共有フォルダがアンマウントされないことに注意してくださいので、特定の共有フォルダの最初の使用の後に--dropbox-shared-foldersは省略できます。

   --dropbox-token
      JSON blobとしてのOAuthアクセストークン。

   --dropbox-token-url
      トークンサーバーのURL。
      
      デフォルトのプロバイダーの設定を使用する場合は空白のままにしてください。


OPTIONS:
   --help, -h  ヘルプを表示します

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。 (デフォルト: false)
   --rescan-interval value  前回のスキャンからこの間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   dropboxのオプション

   --dropbox-auth-url value              認証サーバーのURL。 [$DROPBOX_AUTH_URL]
   --dropbox-batch-commit-timeout value  バッチの終了を待つ最大時間 (デフォルト: "10m0s") [$DROPBOX_BATCH_COMMIT_TIMEOUT]
   --dropbox-batch-mode value            ファイルのバッチ同期のアップロードモードを設定します（sync|async|off） (デフォルト: "sync") [$DROPBOX_BATCH_MODE]
   --dropbox-batch-size value            アップロードバッチ内の最大ファイル数 (デフォルト: "0") [$DROPBOX_BATCH_SIZE]
   --dropbox-batch-timeout value         アップロード前のアップロードバッチのアイドル時間の最大許容時間 (デフォルト: "0s") [$DROPBOX_BATCH_TIMEOUT]
   --dropbox-chunk-size value            アップロードチャンクのサイズ（< 150Mi） (デフォルト: "48Mi") [$DROPBOX_CHUNK_SIZE]
   --dropbox-client-id value             OAuthクライアントID。 [$DROPBOX_CLIENT_ID]
   --dropbox-client-secret value         OAuthクライアントシークレット。 [$DROPBOX_CLIENT_SECRET]
   --dropbox-encoding value              バックエンドのエンコーディング (デフォルト: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$DROPBOX_ENCODING]
   --dropbox-impersonate value           ビジネスアカウントを使用する際に、このユーザーになりすます。 [$DROPBOX_IMPERSONATE]
   --dropbox-shared-files value          個別の共有ファイルで作業するようにrcloneに指示します (デフォルト: "false") [$DROPBOX_SHARED_FILES]
   --dropbox-shared-folders value        共有フォルダで作業するようにrcloneに指示します (デフォルト: "false") [$DROPBOX_SHARED_FOLDERS]
   --dropbox-token value                 JSON blobとしてのOAuthアクセストークン。 [$DROPBOX_TOKEN]
   --dropbox-token-url value             トークンサーバーのURL。 [$DROPBOX_TOKEN_URL]
```
{% endcode %}