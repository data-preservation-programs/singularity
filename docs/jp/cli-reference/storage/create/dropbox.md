# Dropbox

{% code fullWidth="true" %}
```
NAME:
   singularity storage create dropbox - Dropbox

使用方法:
   singularity storage create dropbox [コマンドオプション] [引数...]

説明:
   --client-id
      OAuthクライアントID。
      
      通常は空白のままにしてください。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにしてください。

   --token
      OAuthアクセストークン（JSON形式）。

   --auth-url
      認証サーバーのURL。
      
      デフォルトのプロバイダーの設定を使用する場合は空白のままにしてください。

   --token-url
      トークンサーバーのURL。
      
      デフォルトのプロバイダーの設定を使用する場合は空白のままにしてください。

   --chunk-size
      アップロードチャンクサイズ（< 150Mi）。
      
      このサイズより大きいファイルは、このサイズに分割してアップロードされます。
      
      チャンクはメモリ上でバッファに格納されます（1つずつ）ので、rcloneがリトライ処理を行います。
      サイズを大きくすると速度がわずかに向上します（テストで128MiBの場合、最大で10%程度）が、メモリ使用量が増えます。
      メモリに制約がある場合は、サイズを小さく設定することもできます。

   --impersonate
      ビジネスアカウントを使用する際に、このユーザーを個人をなりすます。
      
      なりすますには、"rclone config"を実行する際にこのフラグが設定されていることを確認する必要があります。
      これにより、rcloneは通常はリクエストしない"members.read"スコープを要求します。
      "members.read"スコープは、Dropbox APIでメンバーのメールアドレスを内部IDに変換するために必要です。
      
      "members.read"スコープを使用するには、Dropboxチームの管理者の承認が必要です。
      
      このオプションを使用するには、独自のAppを使用する必要があります（client_idとclient_secretを自分のものに設定）。
      これは、rcloneのデフォルトの権限セットに"members.read"が含まれていないためです。
      これは、v1.55以降がすべての場所で使用されるようになった後に追加することができます。
      

   --shared-files
      個々の共有ファイルで作業するようにrcloneに指示します。
      
      このモードでは、rcloneの機能は非常に限定されます。list（ls、lslなど）操作と読み取り操作（例：ダウンロード）のみがサポートされます。
      その他の操作はすべて無効になります。

   --shared-folders
      共有フォルダで作業するようにrcloneに指示します。
            
      パスを指定せずにこのフラグを使用すると、リスト操作のみがサポートされ、すべての利用可能な共有フォルダがリストされます。
      パスを指定する場合、最初の部分が共有フォルダの名前として解釈されます。それにより、rcloneはこの共有フォルダをルート名前空間にマウントしようとします。
      成功すると、共有フォルダは通常のフォルダとほぼ同様に扱われ、すべての通常の操作がサポートされます。
      
      なお、共有フォルダはその後アンマウントされないため、特定の共有フォルダの最初の使用後に--dropbox-shared-foldersを省略することができます。

   --batch-mode
      ファイルのバッチアップロードモードをsync|async|offに設定します。
      
      これは、rcloneが使用するバッチモードを設定します。
      
      詳細については、[メインのドキュメント](https://rclone.org/dropbox/#batch-mode)を参照してください。
      
      3つの可能な値があります。
      
      - off - バッチ処理なし
      - sync - バッチアップロードを行い、完了をチェックします（デフォルト）
      - async - バッチアップロードを行い、完了をチェックしません
      
      rcloneは終了時に未処理のバッチをすべて閉じますが、これにより終了時に遅延が生じる場合があります。
      

   --batch-size
      アップロードバッチ内の最大ファイル数。
      
      アップロードするファイルのバッチサイズを設定します。値は1000未満である必要があります。
      
      デフォルトでは、これは0であり、バッチモードの設定に基づいてバッチサイズが計算されます。
      
      - batch_mode: async - デフォルトのバッチサイズは100です
      - batch_mode: sync - デフォルトのバッチサイズは--transfersと同じです
      - batch_mode: off - 使用されていません
      
      rcloneは終了時に未処理のバッチをすべて閉じますが、これにより終了時に遅延が生じる場合があります。
      
      小さなファイルをたくさんアップロードする場合、これを設定するとアップロード速度が向上します。--transfers 32を使用するとスループットを最大化できます。
      

   --batch-timeout
      アップロードバッチがアップロード前にアイドル状態となる最大時間。
      
      アップロードバッチがこの時間以上アイドル状態が続く場合、アップロードされます。
      
      これのデフォルト値は0であり、rcloneは使用中のバッチモードに基づいて適切なデフォルト値を選択します。
      
      - batch_mode: async - デフォルトのバッチタイムアウトは500msです
      - batch_mode: sync - デフォルトのバッチタイムアウトは10sです
      - batch_mode: off - 使用されていません
      

   --batch-commit-timeout
      バッチの完了待ちに最大で待機する時間

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --client-id value      OAuthクライアントID。 [$CLIENT_ID]
   --client-secret value  OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示

   Advanced

   --auth-url value              認証サーバーのURL。 [$AUTH_URL]
   --batch-commit-timeout value  バッチの完了待ちに最大で待機する時間です（デフォルト値: "10m0s"） [$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            ファイルのバッチアップロードモードをsync|async|offに設定します（デフォルト値: "sync"） [$BATCH_MODE]
   --batch-size value            アップロードバッチ内の最大ファイル数です（デフォルト値: 0） [$BATCH_SIZE]
   --batch-timeout value         アップロードバッチがアップロード前にアイドル状態となる最大時間です（デフォルト値: "0s"） [$BATCH_TIMEOUT]
   --chunk-size value            アップロードチャンクサイズ（< 150Mi）です（デフォルト値: "48Mi"） [$CHUNK_SIZE]
   --encoding value              バックエンドのエンコーディングです（デフォルト値: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot"） [$ENCODING]
   --impersonate value           ビジネスアカウントを使用する際に、このユーザーを個人をなりすますです。 [$IMPERSONATE]
   --shared-files                個々の共有ファイルで作業するようにrcloneに指示します（デフォルト値: false） [$SHARED_FILES]
   --shared-folders              共有フォルダで作業するようにrcloneに指示します（デフォルト値: false） [$SHARED_FOLDERS]
   --token value                 OAuthアクセストークン（JSON形式）です [$TOKEN]
   --token-url value             トークンサーバーのURLです [$TOKEN_URL]

   General

   --name value  ストレージの名前（デフォルト値: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}