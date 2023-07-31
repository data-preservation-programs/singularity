# Google ドライブ

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add drive - Google ドライブ

USAGE:
   singularity datasource add drive [command options] <データセット名> <ソースパス>

DESCRIPTION:
   --drive-acknowledge-abuse
      エラー "このファイルは、マルウェアやスパムとして識別されたためダウンロードできません" 
      が "cannotDownloadAbusiveFile" エラーコードと共に返された場合、
      このフラグを rclone に提供し、それによってファイルのダウンロードが許可されるようにします。
      
      マネージャーパーミッションが必要です。サービスアカウントを使用している場合、
      このフラグを使用するには、マネージャーパーミッション（コンテンツマネージャーではありません）が必要です。 
      SA に正しい権限がない場合、Google はこのフラグを無視します。
      
      詳細については、https://rclone.org/drive/#making-your-own-client-id を参照してください。

   --drive-allow-import-name-change
      Google ドキュメントをアップロードする際に、ファイルの形式が変更されることを許可します。
      
      例：file.doc から file.docx へ。これにより、同期が混乱し、毎回再アップロードされます。

   --drive-auth-owner-only
      認証されたユーザーが所有するファイルのみを考慮します。

   --drive-auth-url
      認証サーバーの URL。
      
      デフォルトのプロバイダを使用するには、空白のままにしてください。

   --drive-chunk-size
      アップロードのチャンクサイズ。
      
      2 のべき乗（256k 以上）でなければなりません。
      
      これを大きくするとパフォーマンスが向上しますが、各チャンクは1回の転送ごとにメモリにバッファリングされます。
      
      これを減らすとメモリ使用量は減りますが、パフォーマンスは低下します。

   --drive-client-id
      Google アプリケーションのクライアント ID
      自分の ID を設定することをお勧めします。
      独自のクライアント ID を作成する方法については、https://rclone.org/drive/#making-your-own-client-id を参照してください。
      空白の場合、低パフォーマンスの内部キーが使用されます。

   --drive-client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままです。

   --drive-copy-shortcut-content
      ショートカットの内容をサーバーサイドでコピーします（ショートカット自体ではありません）。
      
      サーバーサイドのコピーを実行する場合、通常 rclone はショートカットをショートカットのままコピーします。
      
      このフラグを使用すると、rclone はサーバーサイドのコピー時にショートカットの内容をコピーします。

   --drive-disable-http2
      http2 で Google ドライブを無効にします。
      
      現在、Google ドライブバックエンドと HTTP/2 の問題が解決されていません。
      そのため、HTTP/2 はドライブバックエンドではデフォルトで無効になっていますが、ここで再有効化することができます。
      この問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/3631

   --drive-encoding
      バックエンドのエンコーディング。
      
      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --drive-export-formats
      Google ドキュメントのダウンロードに対して優先する形式のコンマ区切りのリスト。

   --drive-impersonate
      サービスアカウントを使用している場合に、このユーザーを偽装します。

   --drive-import-formats
      Google ドキュメントのアップロードに対して優先する形式のコンマ区切りのリスト。

   --drive-keep-revision-forever
      各ファイルの新しいヘッドリビジョンを永久に保持します。

   --drive-list-chunk
      リストのチャンクサイズ 100-1000、0 で無効化。

   --drive-pacer-burst
      スリープなしで許可される API コールの数。

   --drive-pacer-min-sleep
      API コール間の最小スリープ時間。

   --drive-resource-key
      リンクで共有されたファイルにアクセスするためのリソースキー。

      https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      のようなリンクで共有されたファイルにアクセスする必要がある場合、
      サービスアカウントを認証したユーザーの最初の部分 "XXX" を "root_folder_id" として使用し、
      2 番目の部分 "YYY" を "resource_key" として使用する必要があります。
      そうしないと、ディレクトリへのアクセス時に 404 エラーが発生します。
      
      参照: https://developers.google.com/drive/api/guides/resource-keys
      
      このリソースキーの要件は、一部の古いファイルにのみ適用されます。
      
      リソースキーは、rclone で認証済みのユーザーが一度 Web インターフェースでフォルダを開くことで、必要なくなるようです。

   --drive-root-folder-id
      ルートフォルダの ID。
      通常は空白のままです。
      
      "Computers" フォルダにアクセスする場合（ドキュメントを参照）、または rclone が開始地点としてルートフォルダ以外を使用する場合に入力します。

   --drive-scope
      drive がアクセスをリクエストする際に使用するスコープ。

      例：
         | drive                   | フルアクセス。Application Data フォルダを除くすべてのファイルにアクセスします。
         | drive.readonly          | ファイルメタデータとファイルコンテンツの読み取り専用アクセス。
         | drive.file              | rclone によって作成されたファイルへのアクセスのみを許可します。
                                   | これらはドライブのウェブサイトで表示されます。
                                   | ファイルの承認の取り消しはユーザーがアプリを認可解除したときに行われます。
         | drive.appfolder         | Application Data フォルダへの読み書きアクセスを許可します。
                                   | これはドライブのウェブサイトでは表示されません。
         | drive.metadata.readonly | ファイルメタデータへの読み取り専用アクセスのみを許可しますが、
                                   | ファイルコンテンツの読み取りまたはダウンロードは許可しません。

   --drive-server-side-across-configs
      サーバーサイドの操作（コピーなど）を異なるドライブの構成間で動作させることを許可します。
      
      これは、2 つの異なる Google ドライブ間でサーバーサイドコピーを行いたい場合に便利です。
      この機能はデフォルトでは有効になっていないため、任意の 2 つの構成間で機能するかどうかは判断しづらいためです。

   --drive-service-account-credentials
      サービスアカウントの認証情報 JSON ブロブ。
      
      通常は空白のままです。
      対話型ログインの代わりに SA を使用する場合にのみ必要です。

   --drive-service-account-file
      サービスアカウントの認証情報 JSON ファイルパス。
      
      通常は空白のままです。
      対話型ログインの代わりに SA を使用する場合にのみ必要です。
      
      ファイル名内の `~` や `${RCLONE_CONFIG_DIR}` 等の環境変数も展開されます。

   --drive-shared-with-me
      共有されたファイルのみを表示します。
      
      "共有されたファイル" フォルダ（Google ドライブで他のユーザーが共有したファイルとフォルダにアクセスできる場所）で rclone が操作するように指示します。
      
      これは、"list"（lsd、lsl など）および "copy"（copy、sync など）コマンド、その他のすべてのコマンドでも機能します。

   --drive-size-as-quota
      サイズを実際のサイズではなく、ストレージクオータ使用量として表示します。
      
      ファイルのサイズをストレージクオータ使用量として表示します。
      これは、現在のバージョンと、永久に保持されている古いバージョンの合計です。
      
      **注意**: このフラグには予期しない結果が生じる場合があります。
      
      このフラグを設定することはお勧めしません。推奨される使用法は、rclone ls/lsl/lsf/lsjson などの実行時に --drive-size-as-quota のフラグ形式を使用することです。
      
      同期にこのフラグを使用する場合（お勧めしません）は、--ignore size も使用する必要があります。

   --drive-skip-checksum-gphotos
      Google フォトとビデオの MD5 チェックサムをスキップします。
      
      Google フォトやビデオの転送時にチェックサムエラーが発生する場合に使用します。
      
      このフラグを設定すると、Google フォトとビデオは空の MD5 チェックサムを返します。
      
      Google フォトは「photos」スペースに存在します。
      
      チェックサムの破損は、Google が画像/ビデオを変更してもチェックサムを更新しないことによって引き起こされます。

   --drive-skip-dangling-shortcuts
      ダングリングショートカットファイルをスキップします。

      これが設定されている場合、rclone はリスト表示でダングリングショートカットを表示しません。

   --drive-skip-gdocs
      すべてのリスト表示で Google ドキュメントをスキップします。

      指定された場合、rclone では gdocs はほぼ見えなくなります。

   --drive-skip-shortcuts
      ショートカットファイルをスキップします。

      通常、rclone はショートカットファイルを参照先のファイルとして表示します（ショートカットについての詳細は、[ショートカットセクション](#shortcuts)を参照してください）。
      このフラグが設定されている場合、rclone はショートカットファイルを完全に無視します。

   --drive-starred-only
      スター付きのファイルのみを表示します。

   --drive-stop-on-download-limit
      ダウンロード制限エラーを致命的なエラーとして扱います。
      
      現在、一日あたり Google ドライブから 10 TiB のデータをダウンロードすることが可能です（これは公式には記載されていない制限です）。
      この制限を超えると、Google ドライブは若干異なるエラーメッセージを発生させます。
      このフラグを設定すると、これらのエラーが致命的なエラーとなります。操作中の同期が停止します。

   --drive-stop-on-upload-limit
      アップロード制限エラーを致命的なエラーとして扱います。
      
      現在、一日あたり Google ドライブへ 750 GiB のデータをアップロードすることが可能です（これは公式には記載されていない制限です）。
      この制限を超えると、Google ドライブは若干異なるエラーメッセージを発生させます。
      このフラグを設定すると、これらのエラーが致命的なエラーとなります。操作中の同期が停止します。
      
      参照: https://github.com/rclone/rclone/issues/3857

   --drive-team-drive
      共有ドライブ（チームドライブ）の ID。

   --drive-trashed-only
      ゴミ箱にあるファイルのみを表示します。
      
      これにより、ゴミ箱内のファイルが元のディレクトリ構造で表示されます。

   --drive-upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ。

   --drive-use-created-date
      変更日時の代わりにファイル作成日時を使用します。
      
      データをダウンロードし、変更日時の代わりに作成日時を使用したい場合に便利です。
      
      **注意**: このフラグは予期しない結果が生じる場合があります。
      
      ドライブにアップロードする場合、ファイルは更新されていない限りすべて上書きされます。逆の場合も同様です。
      この副作用は、"--checksum" フラグを使用することで回避できます。
      
      この機能は、Google フォトで記録された写真のキャプチャ日時を保持するために実装されました。
      Google ドライブの設定で「Google フォトフォルダを作成する」オプションをチェックする必要があります。
      その後、フォルダをローカルにコピーまたは移動し、画像が撮影された（作成された）日時を変更日時として設定できます。

   --drive-use-shared-date
      変更日時の代わりにファイルが共有された日時を使用します。
      
      "--drive-use-created-date" と同様に、このフラグには予期しない結果が生じる場合があります。

      両方のフラグが設定されている場合、作成日時が使用されます。

   --drive-use-trash
      ファイルを完全に削除する代わりにゴミ箱に送信します。
      
      デフォルトでは true、つまりファイルはゴミ箱に送信されます。
      ファイルを完全に削除する場合は、`--drive-use-trash=false` を使用します。

   --drive-v2-download-min-size
      オブジェクトが大きい場合、drive v2 API を使用してダウンロードします。

OPTIONS:
   --help, -h  ヘルプを表示します。

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。  (デフォルト: false)
   --rescan-interval value  最後の正常なスキャンから指定の間隔が経過すると、自動的にソースディレクトリを再スキャンします。  (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。  (デフォルト: ready)

   ドライブのオプション

   --drive-acknowledge-abuse value            エラーファイルをダウンロードできるようにします。 (デフォルト: "false") [$DRIVE_ACKNOWLEDGE_ABUSE]
   --drive-allow-import-name-change value     Google ドキュメントをアップロードする際に、ファイルの形式が変更されることを許可します。 (デフォルト: "false") [$DRIVE_ALLOW_IMPORT_NAME_CHANGE]
   --drive-auth-owner-only value              認証されたユーザーが所有するファイルのみを考慮します。 (デフォルト: "false") [$DRIVE_AUTH_OWNER_ONLY]
   --drive-auth-url value                     認証サーバーの URL。 [$DRIVE_AUTH_URL]
   --drive-chunk-size value                   アップロードのチャンクサイズ。 (デフォルト: "8Mi") [$DRIVE_CHUNK_SIZE]
   --drive-client-id value                    Google アプリケーションのクライアント ID [$DRIVE_CLIENT_ID]
   --drive-client-secret value                OAuth クライアントシークレット。 [$DRIVE_CLIENT_SECRET]
   --drive-copy-shortcut-content value        ショートカットの内容をサーバーサイドでコピーします（ショートカット自体ではありません）。 (デフォルト: "false") [$DRIVE_COPY_SHORTCUT_CONTENT]
   --drive-disable-http2 value                http2 を無効にします。 (デフォルト: "true") [$DRIVE_DISABLE_HTTP2]
   --drive-encoding value                     バックエンドのエンコーディング。 (デフォルト: "InvalidUtf8") [$DRIVE_ENCODING]
   --drive-export-formats value               Google ドキュメントのダウンロードに対して優先する形式のコンマ区切りのリスト。 (デフォルト: "docx,xlsx,pptx,svg") [$DRIVE_EXPORT_FORMATS]
   --drive-formats value                      非推奨: export_formats を参照してください。 [$DRIVE_FORMATS]
   --drive-impersonate value                  サービスアカウントを使用している場合に、このユーザーを偽装します。 [$DRIVE_IMPERSONATE]
   --drive-import-formats value               Google ドキュメントのアップロードに対して優先する形式のコンマ区切りのリスト。 [$DRIVE_IMPORT_FORMATS]
   --drive-keep-revision-forever value        各ファイルの新しいヘッドリビジョンを永久に保持します。 (デフォルト: "false") [$DRIVE_KEEP_REVISION_FOREVER]
   --drive-list-chunk value                   リストのチャンクサイズ 100-1000、0 で無効化。 (デフォルト: "1000") [$DRIVE_LIST_CHUNK]
   --drive-pacer-burst value                  スリープなしで許可される API コールの数。 (デフォルト: "100") [$DRIVE_PACER_BURST]
   --drive-pacer-min-sleep value              API コール間の最小スリープ時間。 (デフォルト: "100ms") [$DRIVE_PACER_MIN_SLEEP]
   --drive-resource-key value                 リンクで共有されたファイルにアクセスするためのリソースキー。 [$DRIVE_RESOURCE_KEY]
   --drive-root-folder-id value               ルートフォルダの ID。 [$DRIVE_ROOT_FOLDER_ID]
   --drive-scope value                        drive がアクセスをリクエストする際に使用するスコープ。 [$DRIVE_SCOPE]
   --drive-server-side-across-configs value   サーバーサイドの操作（コピーなど）を異なるドライブの構成間で動作させることを許可します。 (デフォルト: "false") [$DRIVE_SERVER_SIDE_ACROSS_CONFIGS]
   --drive-service-account-credentials value  サービスアカウントの認証情報 JSON ブロブ。 [$DRIVE_SERVICE_ACCOUNT_CREDENTIALS]
   --drive-service-account-file value         サービスアカウントの認証情報 JSON ファイルパス。 [$DRIVE_SERVICE_ACCOUNT_FILE]
   --drive-shared-with-me value               共有されたファイルのみを表示します。 (デフォルト: "false") [$DRIVE_SHARED_WITH_ME]
   --drive-size-as-quota value                サイズを実際のサイズではなく、ストレージクオータ使用量として表示します。 (デフォルト: "false") [$DRIVE_SIZE_AS_QUOTA]
   --drive-skip-checksum-gphotos value        Google フォトとビデオの MD5 チェックサムをスキップします。 (デフォルト: "false") [$DRIVE_SKIP_CHECKSUM_GPHOTOS]
   --drive-skip-dangling-shortcuts value      ダングリングショートカットファイルをスキップします。 (デフォルト: "false") [$DRIVE_SKIP_DANGLING_SHORTCUTS]
   --drive-skip-gdocs value                   すべてのリスト表示で Google ドキュメントをスキップします。 (デフォルト: "false") [$DRIVE_SKIP_GDOCS]
   --drive-skip-shortcuts value               ショートカットファイルをスキップします。 (デフォルト: "false") [$DRIVE_SKIP_SHORTCUTS]
   --drive-starred-only value                 スター付きのファイルのみを表示します。 (デフォルト: "false") [$DRIVE_STARRED_ONLY]
   --drive-stop-on-download-limit value       ダウンロード制限エラーを致命的なエラーとして扱います。 (デフォルト: "false") [$DRIVE_STOP_ON_DOWNLOAD_LIMIT]
   --drive-stop-on-upload-limit value         アップロード制限エラーを致命的なエラーとして扱います。 (デフォルト: "false") [$DRIVE_STOP_ON_UPLOAD_LIMIT]
   --drive-team-drive value                   共有ドライブ（チームドライブ）の ID。 [$DRIVE_TEAM_DRIVE]
   --drive-token value                        OAuth アクセストークンを JSON ブロブとして指定します。 [$DRIVE_TOKEN]
   --drive-token-url value                    トークンサーバーの URL。 [$DRIVE_TOKEN_URL]
   --drive-trashed-only value                 ゴミ箱にあるファイルのみを表示します。 (デフォルト: "false") [$DRIVE_TRASHED_ONLY]
   --drive-upload-cutoff value                チャンクアップロードに切り替えるためのカットオフ。 (デフォルト: "8Mi") [$DRIVE_UPLOAD_CUTOFF]
   --drive-use-created-date value             変更日時の代わりにファイル作成日時を使用します。 (デフォルト: "false") [$DRIVE_USE_CREATED_DATE]
   --drive-use-shared-date value              変更日時の代わりにファイルが共有された日時を使用します。 (デフォルト: "false") [$DRIVE_USE_SHARED_DATE]
   --drive-use-trash value                    ファイルを完全に削除する代わりにゴミ箱に送信します。 (デフォルト: "true") [$DRIVE_USE_TRASH]
   --drive-v2-download-min-size value         オブジェクトが大きい場合、drive v2 API を使用してダウンロードします。 (デフォルト: "off") [$DRIVE_V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}