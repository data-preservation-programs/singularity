# Googleドライブ

{% code fullWidth="true" %}
```
名前:
   singularity storage create drive - Googleドライブ

使用法:
   singularity storage create drive [コマンドオプション] [引数...]

説明:
   --client-id
      GoogleアプリケーションクライアントID
      自分自身の設定が推奨されています。
      作成方法については、https://rclone.org/drive/#making-your-own-client-idを参照してください。
      空白のままにすると、性能が低い内部キーが使用されます。

   --client-secret
      OAuthクライアントシークレット。
      
      通常は空白のままにします。

   --token
      JSONブロブとしてのOAuthアクセストークン。

   --auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルトを使用する場合は空白のままにします。

   --scope
      rcloneがdriveからアクセスを要求する際に使用するスコープ。

      例:
         | drive                   | 全ファイルへの完全なアクセス。ただしアプリケーションデータフォルダは除外します。
         | drive.readonly          | ファイルのメタデータとファイルの内容への読み取り専用アクセス。
         | drive.file              | rcloneによって作成されたファイルへのアクセスのみ。
         |                         | これらはドライブのウェブサイトで表示されます。
         |                         | ユーザーがアプリを認証解除すると、ファイルの承認も取り消されます。
         | drive.appfolder         | Application Dataフォルダへの読み書きアクセスを許可します。
         |                         | これはドライブのウェブサイトで表示されません。
         | drive.metadata.readonly | ファイルメタデータへの読み取り専用アクセスを許可しますが、
         |                         | ファイルのコンテンツを読み取るかダウンロードすることはできません。

   --root-folder-id
      ルートフォルダのID。
      通常は空白のままにします。
      
      「Computers」というフォルダにアクセスするために入力するか、
      rcloneが開始ポイントとして使用するルートフォルダ以外のフォルダを使用するために入力する必要があります。
      

   --service-account-file
      サービスアカウントの認証情報JSONファイルのパス。
      
      通常は空白のままにします。
      インタラクティブログインの代わりにSAを使用する場合にのみ必要です。
      
      先頭の`〜`は、ファイル名で展開されます。`${RCLONE_CONFIG_DIR}`などの環境変数も同様です。

   --service-account-credentials
      サービスアカウントの認証情報JSONブロブ。
      
      通常は空白のままにします。
      インタラクティブログインの代わりにSAを使用する場合にのみ必要です。

   --team-drive
      共有ドライブ（チームドライブ）のID。

   --auth-owner-only
      認証されたユーザーが所有するファイルのみを考慮する。

   --use-trash
      ファイルを完全に削除する代わりに、ゴミ箱に送信します。
      
      デフォルトではファイルをゴミ箱に送信します。
      ファイルを完全に削除するには、`--drive-use-trash=false`を使用します。

   --copy-shortcut-content
      ショートカットの内容ではなく、サーバーサイドでショートカットの内容をコピーします。
      
      サーバーサイドでのコピーを行う場合、通常、rcloneはショートカットをショートカットとしてコピーします。
      
      このフラグを使用すると、rcloneはサーバーサイドでのコピー時にショートカットの内容をコピーし、
      ショートカットそのものではなくします。

   --skip-gdocs
      すべてのリスト表示でGoogleドキュメントをスキップします。
      
      指定された場合、gdocsはrcloneではほぼ表示されなくなります。

   --skip-checksum-gphotos
      GoogleフォトとビデオのMD5チェックサムをスキップします。
      
      Googleフォトやビデオを転送する際にチェックサムエラーが発生する場合は、このフラグを使用します。
      
      このフラグを設定すると、Googleフォトとビデオは
      空白のMD5チェックサムを返します。
      
      Googleフォトは「photos」スペースにあるため、
      フォトまたはビデオを変更した場合でも、
      チェックサムが破損します。
      
      Googleがイメージ/ビデオを変更したがチェックサムを更新しなかったため、
      チェックサムが破損します。

   --shared-with-me
      私と共有されたファイルのみを表示します。
      
      rcloneに「私と共有された」フォルダ（Googleドライブで他のユーザーが共有したファイルとフォルダにアクセスできる場所）で操作するよう指示します。
      
      これは「list」（lsd、lslなど）および「copy」（copy、syncなど）コマンド、および他のすべてのコマンドでも機能します。

   --trashed-only
      ゴミ箱内のファイルのみを表示します。
      
      これにより、トラッシュされたファイルが元のディレクトリ構造で表示されます。

   --starred-only
      スターがつけられたファイルのみを表示します。

   --formats
      廃止: export_formatsを参照してください。

   --export-formats
      Googleドキュメントをダウンロードするための優先する形式のカンマ区切りリスト。

   --import-formats
      Googleドキュメントをアップロードするための優先する形式のカンマ区切りリスト。

   --allow-import-name-change
      Googleドキュメントをアップロードする際にファイルのタイプの変更を許可します。
      
      たとえば、file.docをfile.docxに変更できます。これにより、同期が混乱し、毎回再アップロードされます。

   --use-created-date
      最終更新日時の代わりにファイル作成日時を使用します。
      
      データのダウンロード中に作成日時を最終更新日時として使用したい場合に便利です。
      
      **警告**: このフラグには予期しない影響がある可能性があります。
      
      ドライブにアップロードすると、ファイルは作成時から変更されていない限り上書きされます。
      ダウンロード中には逆のことが起こります。 この副作用は、
      "--checksum"フラグを使用することで回避できます。
      
      この機能は、Googleフォトが記録した写真のキャプチャ日を保持するために実装されました。
      Googleドライブの設定で「Googleフォトフォルダを作成する」オプションを最初に確認する必要があります。
      その後、写真をローカルにコピーまたは移動し、画像が撮影された（作成された）日付を変更日時として設定できます。

   --use-shared-date
      ファイルが共有された日付ではなく、最終更新日時の代わりにファイルの共有日を使用します。
      
      「--drive-use-created-date」と同様に、
      このフラグにはアップロード/ダウンロード時に予期しない影響がある可能性があります。
      
      このフラグと「--drive-use-created-date」の両方が設定されている場合、作成日時が使用されます。

   --list-chunk
      リストのチャンクサイズ 100〜1000、無効にするには0。

   --impersonate
      サービスアカウントを使用する場合に、このユーザーを模倣します。

   --alternate-export
      廃止: もはや必要ありません。

   --upload-cutoff
      チャンクされたアップロードに切り替えるためのカットオフ値。

   --chunk-size
      アップロードのチャンクサイズ。
      
      2のべき乗であり、256k以上でなければなりません。
      
      これを大きくするとパフォーマンスが向上しますが、各チャンクは1つの転送ごとにメモリにバッファリングされます。
      
      メモリ使用量を減らすがパフォーマンスは低下する。

   --acknowledge-abuse
      cannotDownloadAbusiveFileを返すファイルのダウンロードを許可するように設定します。
      
      ファイルのダウンロードが次のエラーで失敗する場合、「This file has been identified as malware or spam and cannot be downloaded」
      エラーコード「cannotDownloadAbusiveFile」とともに返される場合、
      rcloneにこのフラグを供給して、ファイルのダウンロードのリスクを認識していることを示します。
      
      サービスアカウントを使用している場合は、
      正しく動作するためには「マネージャー」の権限（コンテンツマネージャーではなく）が必要です。 
      SAに適切な権限がない場合、Googleはフラグを無視します。

   --keep-revision-forever
      各ファイルの新しいヘッドリビジョンを永久に保持します。

   --size-as-quota
      サイズを実際のサイズではなくストレージクォータ使用量として表示します。
      
      ファイルのサイズを使用されたストレージのクォータとして表示します。これは
      現在のバージョンに加えて、永久に保持されるように設定されている以前のバージョンの合計です。
      
      **警告**: このフラグには予期しない影響がある可能性があります。
      
      構成でこのフラグを設定することはお勧めしません。推奨される使用方法は、
      rclone ls / lsl / lsf / lsjsonなどを行う場合に `--drive-size-as-quota` フラグを使用することです。
      
      同期（推奨されない）にこのフラグを使用する場合は、`--ignore size` も使用する必要があります。

   --v2-download-min-size
      オブジェクトが大きい場合、drive v2 APIを使用してダウンロードします。

   --pacer-min-sleep
      API呼び出し間の最小スリープ時間。

   --pacer-burst
      スリープせずに許可されるAPIコールの数。

   --server-side-across-configs
      サーバーサイドの操作（コピーなど）を異なるドライブ設定間で動作させることを許可します。
      
      これは、2つの異なるGoogleドライブ間でサーバーサイドのコピーを行いたい場合に便利です。
      これはデフォルトで有効ではありません。なぜなら、それが任意の2つの
      設定で動作するかどうかを判断するのは容易ではないためです。

   --disable-http2
      ドライブでhttp2を使用しないようにします。
      
      現在、GoogleドライブバックエンドとHTTP/2の問題が未解決です。 
      デフォルトではドライブバックエンドでHTTP/2が無効になっていますが、ここで再度有効にできます。 
      問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/3631
      
      

   --stop-on-upload-limit
      アップロード制限エラーを致命的なエラーにします。
      
      現在のところ、1日にGoogleドライブに750 GiBのデータをアップロードすることしかできません（この制限は公式には文書化されていません）。
      この制限に達すると、Googleドライブは若干異なるエラーメッセージを生成します。 このフラグが設定されていると、これらのエラーは致命的なものになります。 これにより、進行中の同期が停止します。
      
      エラーメッセージの文字列を基に検出しているため、これはGoogleが文書化していないため、将来的には動作しなくなる可能性があります。
      
      参照: https://github.com/rclone/rclone/issues/3857
      

   --stop-on-download-limit
      ダウンロード制限エラーを致命的なエラーにします。
      
      現在のところ、Googleドライブから1日に10 TiBのデータをダウンロードすることしかできません（この制限は公式には文書化されていません）。
      この制限に達すると、Googleドライブは若干異なるエラーメッセージを生成します。 このフラグが設定されていると、これらのエラーは致命的なものになります。 これにより、進行中の同期が停止します。
      
      エラーメッセージの文字列を基に検出しているため、これはGoogleが文書化していないため、将来的には動作しなくなる可能性があります。
      

   --skip-shortcuts
      ショートカットファイルをスキップします。
      
      通常、rcloneはショートカットファイルを参照解除し、
      オリジナルのファイルのように表示します（[ショートカットセクション](#shortcuts)を参照）。
      このフラグが設定されている場合、rcloneはショートカットファイルを完全に無視します。
      

   --skip-dangling-shortcuts
      ダングリングショートカットファイルをスキップします。
      
      これが設定されている場合、rcloneはリスト表示にダングリングショートカットを表示しません。
      

   --resource-key
      リンク共有ファイルにアクセスするためのリソースキー。
      
      このようなリンクで共有されたファイルにアクセスする必要がある場合、
      
          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      その場合、「XXX」を「ルートフォルダID」として、「YYY」を「リソースキー」として使用する必要があります。そうしないと、ディレクトリにアクセスしようとすると404エラーが発生します。
      
      参照: https://developers.google.com/drive/api/guides/resource-keys
      
      このリソースキーの要件は、一部の古いファイルにのみ適用されます。
      
      また、（rcloneで認証したユーザーで）ウェブインターフェースでフォルダを1回開くだけで、
      リソースキーは必要ないようです。
      

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --alternate-export            廃止: もはや必要ありません。 (default: false) [$ALTERNATE_EXPORT]
   --client-id value             GoogleアプリケーションクライアントID [$CLIENT_ID]
   --client-secret value         OAuthクライアントシークレット。 [$CLIENT_SECRET]
   --help, -h                    ヘルプを表示
   --scope value                 rcloneがdriveからアクセスを要求する際に使用するスコープ。 [$SCOPE]
   --service-account-file value  サービスアカウントの認証情報JSONファイルのパス。 [$SERVICE_ACCOUNT_FILE]

   Advanced

   --acknowledge-abuse                  cannotDownloadAbusiveFileを返すファイルのダウンロードを許可するように設定します。 (default: false) [$ACKNOWLEDGE_ABUSE]
   --allow-import-name-change           Googleドキュメントをアップロードする際にファイルのタイプの変更を許可します。 (default: false) [$ALLOW_IMPORT_NAME_CHANGE]
   --auth-owner-only                    認証されたユーザーが所有するファイルのみを考慮します。 (default: false) [$AUTH_OWNER_ONLY]
   --auth-url value                     認証サーバーのURL。 [$AUTH_URL]
   --chunk-size value                   アップロードのチャンクサイズ。 (default: "8Mi") [$CHUNK_SIZE]
   --copy-shortcut-content              サーバーサイドでショートカットの内容をコピーします。 (default: false) [$COPY_SHORTCUT_CONTENT]
   --disable-http2                      ドライブでhttp2を使用しないようにします。 (default: true) [$DISABLE_HTTP2]
   --encoding value                     バックエンドのエンコーディング。 (default: "InvalidUtf8") [$ENCODING]
   --export-formats value               Googleドキュメントをダウンロードするための優先する形式のカンマ区切りリスト。 (default: "docx,xlsx,pptx,svg") [$EXPORT_FORMATS]
   --formats value                      廃止: export_formatsを参照してください。 [$FORMATS]
   --impersonate value                  サービスアカウントを使用する場合に、このユーザーを模倣します。 [$IMPERSONATE]
   --import-formats value               Googleドキュメントをアップロードするための優先する形式のカンマ区切りリスト。 [$IMPORT_FORMATS]
   --keep-revision-forever              各ファイルの新しいヘッドリビジョンを永久に保持します。 (default: false) [$KEEP_REVISION_FOREVER]
   --list-chunk value                   リストのチャンクサイズ 100-1000、無効にするには0。 (default: 1000) [$LIST_CHUNK]
   --pacer-burst value                  スリープせずに許可されるAPIコールの数。 (default: 100) [$PACER_BURST]
   --pacer-min-sleep value              API呼び出し間の最小スリープ時間。 (default: "100ms") [$PACER_MIN_SLEEP]
   --resource-key value                 リンク共有ファイルにアクセスするためのリソースキー。 [$RESOURCE_KEY]
   --root-folder-id value               ルートフォルダのID。 [$ROOT_FOLDER_ID]
   --server-side-across-configs         サーバーサイドの操作（コピーなど）を異なるドライブ設定間で動作させることを許可します。 (default: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --service-account-credentials value  サービスアカウントの認証情報JSONブロブ。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --shared-with-me                     私と共有されたファイルのみを表示します。 (default: false) [$SHARED_WITH_ME]
   --size-as-quota                      サイズを実際のサイズではなくストレージクォータ使用量として表示します。 (default: false) [$SIZE_AS_QUOTA]
   --skip-checksum-gphotos              GoogleフォトとビデオのMD5チェックサムをスキップします。 (default: false) [$SKIP_CHECKSUM_GPHOTOS]
   --skip-dangling-shortcuts            ダングリングショートカットファイルをスキップします。 (default: false) [$SKIP_DANGLING_SHORTCUTS]
   --skip-gdocs                         すべてのリスト表示でGoogleドキュメントをスキップします。 (default: false) [$SKIP_GDOCS]
   --skip-shortcuts                     ショートカットファイルをスキップします。 (default: false) [$SKIP_SHORTCUTS]
   --starred-only                       スターがつけられたファイルのみを表示します。 (default: false) [$STARRED_ONLY]
   --stop-on-download-limit             ダウンロード制限エラーを致命的なエラーにします。 (default: false) [$STOP_ON_DOWNLOAD_LIMIT]
   --stop-on-upload-limit               アップロード制限エラーを致命的なエラーにします。 (default: false) [$STOP_ON_UPLOAD_LIMIT]
   --team-drive value                   共有ドライブ（チームドライブ）のID。 [$TEAM_DRIVE]
   --token value                        JSONブロブとしてのOAuthアクセストークン。 [$TOKEN]
   --token-url value                    トークンサーバーのURL。 [$TOKEN_URL]
   --trashed-only                       ゴミ箱内のファイルのみを表示します。 (default: false) [$TRASHED_ONLY]
   --upload-cutoff value                チャンクされたアップロードに切り替えるためのカットオフ値。 (default: "8Mi") [$UPLOAD_CUTOFF]
   --use-created-date                   最終更新日時の代わりにファイル作成日時を使用します。 (default: false) [$USE_CREATED_DATE]
   --use-shared-date                    ファイルが共有された日付ではなく、最終更新日時の代わりにファイルの共有日を使用します。 (default: false) [$USE_SHARED_DATE]
   --use-trash                          ファイルを完全に削除する代わりに、ゴミ箱に送信します。 (default: true) [$USE_TRASH]
   --v2-download-min-size value         オブジェクトが大きい場合、drive v2 APIを使用してダウンロードします。 (default: "off") [$V2_DOWNLOAD_MIN_SIZE]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}