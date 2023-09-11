# Google Drive

{% code fullWidth="true" %}
```
名前:
   singularity storage update drive - Google Drive

使用法:
   singularity storage update drive [コマンドオプション] <名前|ID>

説明:
   --client-id
      GoogleのアプリケーションクライアントID
      自分自身のものを設定することをおすすめします。
      自分自身で作成する方法については、https://rclone.org/drive/#making-your-own-client-idを参照してください。
      空白のままにすると、パフォーマンスが低下する内部キーが使用されます。

   --client-secret
      OAuthクライアントシークレット。

      通常は空白のままにしておきます。

   --token
      OAuthアクセストークン（JSON形式）。

   --auth-url
      認証サーバーのURL。

      プロバイダのデフォルトを使用するには空白のままにしておきます。

   --token-url
      トークンサーバーのURL。

      プロバイダのデフォルトを使用するには空白のままにしておきます。

   --scope
      ドライブへのアクセス時にrcloneが使用するスコープ。

      例:
         | drive                   | 全ファイルへの完全アクセス（アプリケーションデータフォルダを除く）。
         | drive.readonly          | ファイルのメタデータとファイルの内容への読み取り専用アクセス。
         | drive.file              | rcloneによって作成されたファイルへのアクセスのみ。
         |                         | これらはドライブのウェブサイトで表示されます。
         |                         | ファイルの承認はユーザーがアプリを承認解除すると取り消されます。
         | drive.appfolder         | Application Dataフォルダへの読み取りと書き込みのアクセスを許可。
         |                         | これはドライブのウェブサイトで表示されません。
         | drive.metadata.readonly | ファイルのメタデータへの読み取り専用アクセスですが、
         |                         | ファイルの内容を読み取ることやダウンロードすることはできません。

   --root-folder-id
      ルートフォルダのID。
      通常は空白のままにしておきます。

      ドキュメントを参照）またはrcloneが使用する非ルートフォルダを
      開始点として使用するために入力します。


   --service-account-file
      サービスアカウントの認証情報のJSONファイルパス。

      通常は空白のままにしておきます。
      インタラクティブログインの代わりにSAを使用する場合にのみ必要です。

      先頭の`~`はファイル名の中で展開されます。
      `${RCLONE_CONFIG_DIR}`などの環境変数も同様です。

   --service-account-credentials
      サービスアカウント認証情報のJSON blob。

      通常は空白のままにしておきます。
      インタラクティブログインの代わりにSAを使用する場合にのみ必要です。

   --team-drive
      共有ドライブ（チームドライブ）のID。

   --auth-owner-only
      認証されたユーザーによって所有されているファイルのみを考慮する。

   --use-trash
      ファイルを完全に削除する代わりにゴミ箱に送信。

      既定では、ファイルはゴミ箱に送信されます。
      ファイルを完全に削除する場合は「--drive-use-trash=false」とします。

   --copy-shortcut-content
      ショートカットの代わりにサーバーサイドでショートカットの内容をコピーする。

      サーバーサイドコピーを実行すると、通常rcloneはショートカットをショートカットのままコピーします。

      このフラグを使用すると、サーバーサイドコピーを実行する際に、rcloneはショートカットの内容をコピーします。

   --skip-gdocs
      すべてのリストでGoogleドキュメントをスキップする。

      与えられた場合、gdocsはrcloneにとってほとんど見えなくなります。

   --skip-checksum-gphotos
      GoogleフォトおよびビデオのMD5チェックサムをスキップする。

      Googleフォトやビデオの転送時にチェックサムエラーが発生する場合に使用します。

      このフラグを設定すると、Googleフォトおよびビデオは空のMD5チェックサムを返します。

      なお、Googleフォトは「photos」スペースにあることで識別されます。

      破損したチェックサムは、Googleがイメージ/ビデオを変更したが
      チェックサムを更新しなかったために発生します。

   --shared-with-me
      共有されたファイルのみを表示する。

      rcloneが「Shared with me」フォルダで操作するように指示します（
      ここではGoogleドライブを介して他のユーザーが共有したファイルとフォルダにアクセスできます）。

      これは「list」（lsd、lslなど）および「copy」（copy、syncなど）コマンドだけでなく、
      他のすべてのコマンドでも機能します。

   --trashed-only
      ゴミ箱にあるファイルのみを表示する。

      これにより、元のディレクトリ構造のまま、ゴミ箱の中のファイルが表示されます。

   --starred-only
      スターがついているファイルのみを表示する。

   --formats
      廃止予定: export_formatsを参照してください。

   --export-formats
      Googleドキュメントをダウンロードする際の優先フォーマットのカンマ区切りリスト。

   --import-formats
      Googleドキュメントをアップロードする際の優先フォーマットのカンマ区切りリスト。

   --allow-import-name-change
      Googleドキュメントをアップロードする際にファイルタイプの変更を許可する。

      たとえば、file.docからfile.docxに変わります。
      これにより、同期が混乱し、毎回再アップロードされます。

   --use-created-date
      更新日ではなく作成日を使用する。

      データのダウンロード時に作成日を計算したい場合に便利です。

      **警告**：このフラグにはいくつかの予期しない影響がある場合があります。

      ドライブにアップロードすると、**すべてのファイルが上書き**されます
      （作成以降に変更されていない場合を除く）。
      ダウンロード時にはその逆が起こります。
      この副作用を回避するには、「--checksum」フラグを使用します。

      この機能は、Googleフォトによって記録された写真の撮影日を維持するために実装されました。
      Googleドライブの設定で「Google Photosフォルダを作成する」オプションをチェックする必要があります。
      その後、写真をローカルにコピーまたは移動し、
      イメージの撮影日（作成日）を変更日として設定できます。

   --use-shared-date
      共有された日付（作成日）の代わりに更新日を使用する。

      注：「--drive-use-created-date」と同様に、このフラグには
      アップロード/ダウンロード時のファイルに予期しない影響がある場合があります。

      このフラグと「--drive-use-created-date」の両方が設定されている場合、作成日が使用されます。

   --list-chunk
      リストのチャンクサイズ 100-1000、0で無効。

   --impersonate
      サービスアカウントを使用する場合に、このユーザーをなりすます。

   --alternate-export
      廃止予定: もはや必要ありません。

   --upload-cutoff
      チャンクアップロードに切り替えるための閾値。

   --chunk-size
      アップロードチャンクのサイズ。

      2のn乗であり、256k以上である必要があります。

      これを大きくするとパフォーマンスが向上しますが、
      各チャンクは転送ごとにメモリに1つずつバッファリングされます。

      これを減らすとメモリ使用量が減少しますが、パフォーマンスが低下します。

   --acknowledge-abuse
      "cannotDownloadAbusiveFile"エラーコードで「このファイルは悪意のあるコンテンツまたはスパムとして識別され、ダウンロードできません」というエラーメッセージが返されるファイルをダウンロードできるように設定する。

      ファイルのダウンロード時に「cannotDownloadAbusiveFile」というエラーコードが返される場合、このフラグをrcloneに指定して、
      ダウンロードするファイルのリスクを承知したことを示します。
      rcloneはそれでもファイルをダウンロードします。

      なお、サービス アカウントを使用している場合、このフラグが機能するためには Manager 権限（Content Manager ではありません）が必要です。
      SAに適切な権限がない場合、Googleはこのフラグを無視します。

   --keep-revision-forever
      各ファイルの新しいヘッドリビジョンを永久に保持する。

   --size-as-quota
      サイズをストレージとしての使用クオータ表示する。

      ファイルのサイズを使用したストレージクオータとして表示します。
      これは現行バージョンと、永久に保持するように設定された古いバージョンを合算します。

      **警告**：このフラグにはいくつかの予期しない影響がある場合があります。

      構成ファイルでこのフラグを設定することはお勧めしません。
      推奨されるのは、rclone ls/lsl/lsf/lsjsonなどを実行する際に--drive-size-as-quotaフラグを使用することです。

      同期にこのフラグを使用する場合は --ignore size も使用する必要があります。

   --v2-download-min-size
      オブジェクトがこれよりも大きい場合は、v2 APIを使用してダウンロードします。

   --pacer-min-sleep
      API呼び出し間の最小スリープ時間。

   --pacer-burst
      スリープせずに許可されるAPI呼び出しの数。

   --server-side-across-configs
      異なるドライブ構成間でサーバーサイドの操作（例：コピー）を実行することを許可する。

      これは、2つの異なるGoogleドライブ間でサーバーサイドのコピーを行いたい場合に便利です。
      すべての構成の間で動作するかどうかは簡単に判断できないため、デフォルトでは無効になっています。

   --disable-http2
      Driveがhttp2を使用しないように設定する。

      現在、GoogleドライブバックエンドとHTTP/2の間に解決できない問題があります。
      ドライブバックエンドではHTTP/2はデフォルトで無効になっていますが、ここで再度有効にできます。
      この問題が解決されると、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/3631

   --stop-on-upload-limit
      アップロード制限エラーを致命的にする。

      現時点では、Googleドライブに1日あたり750 GiBのデータをアップロードすることが可能です（これは非公開の制限です）。
      この制限に達すると、Googleドライブはわずかに異なるエラーメッセージを生成します。
      このフラグが設定されていると、これらのエラーは致命的となります。
      これにより、進行中の同期が中止されます。

      Googleが文書化していないエラーメッセージ文字列に依存しているため、この検出は将来的には壊れる可能性があります。

      参照: https://github.com/rclone/rclone/issues/3857

   --stop-on-download-limit
      ダウンロード制限エラーを致命的にする。

      現時点では、Googleドライブから1日あたり10 TiBのデータをダウンロードすることが可能です（これは非公開の制限です）。
      この制限に達すると、Googleドライブはわずかに異なるエラーメッセージを生成します。
      このフラグが設定されていると、これらのエラーは致命的となります。
      これにより、進行中の同期が中止されます。

      Googleが文書化していないエラーメッセージ文字列に依存しているため、この検出は将来的には壊れる可能性があります。

   --skip-shortcuts
      ショートカットファイルをスキップする場合は、設定する。

      通常、rcloneではショートカットファイルを参照先のオリジナルファイルとして処理します（ショートカットについては、[ショートカットセクション](#shortcuts)を参照）。
      このフラグが設定されている場合、rcloneはショートカットファイルを完全に無視します。

   --skip-dangling-shortcuts
      ダングリングショートカットファイルをスキップする場合は、設定する。

      このフラグが設定されている場合、rcloneはリストにダングリングショートカットを表示しません。

   --resource-key
      リンク共有されたファイルにアクセスするためのリソースキー。

      次のようなリンクで共有されたファイルにアクセスする必要がある場合は、

          https://drive.google.com/drive/folders/XXX?resourcekey=YYY&usp=sharing
      
      "XXX"を"root_folder_id"、"YYY"を"resource_key"として使用する必要があります。
      そうしないと、ディレクトリにアクセスしようとすると、404 not found エラーが表示されます。

      参照: https://developers.google.com/drive/api/guides/resource-keys

      このリソースキーの要件は、一部の古いファイルにのみ適用されます。

      また、（利用したいユーザーでrcloneが認証されている場合）、
      1回ウェブインターフェースでフォルダを開くだけで、リソースキーを必要としないようです。

   --encoding
      バックエンドのエンコーディング。

      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --alternate-export            廃止予定: もはや必要ありません。(default: false) [$ALTERNATE_EXPORT]
   --client-id value             GoogleのアプリケーションクライアントID [$CLIENT_ID]
   --client-secret value         OAuthクライアントシークレット [$CLIENT_SECRET]
   --help, -h                    ヘルプを表示する
   --scope value                 ドライブへのアクセス時にrcloneが使用するスコープ [$SCOPE]
   --service-account-file value  サービスアカウントの認証情報のJSONファイルパス [$SERVICE_ACCOUNT_FILE]

   高度なオプション

   --acknowledge-abuse                  "cannotDownloadAbusiveFile"エラーコードで「このファイルは悪意のあるコンテンツまたはスパムとして識別され、ダウンロードできません」というエラーメッセージが返されるファイルをダウンロードできるように設定する。(default: false) [$ACKNOWLEDGE_ABUSE]
   --allow-import-name-change           Googleドキュメントをアップロードする際にファイルタイプの変更を許可する。(default: false) [$ALLOW_IMPORT_NAME_CHANGE]
   --auth-owner-only                    認証されたユーザーによって所有されているファイルのみを考慮する。(default: false) [$AUTH_OWNER_ONLY]
   --auth-url value                     認証サーバーのURL [$AUTH_URL]
   --chunk-size value                   アップロードチャンクのサイズ。(default: "8Mi") [$CHUNK_SIZE]
   --copy-shortcut-content              ショートカットの代わりにサーバーサイドでショートカットの内容をコピーする。(default: false) [$COPY_SHORTCUT_CONTENT]
   --disable-http2                      Driveがhttp2を使用しないように設定する。(default: true) [$DISABLE_HTTP2]
   --encoding value                     バックエンドのエンコーディング。(default: "InvalidUtf8") [$ENCODING]
   --export-formats value               Googleドキュメントをダウンロードする優先フォーマットのカンマ区切りリスト。(default: "docx,xlsx,pptx,svg") [$EXPORT_FORMATS]
   --formats value                      廃止予定: export_formatsを参照してください。 [$FORMATS]
   --impersonate value                  サービスアカウントを使用する場合に、このユーザーをなりすます。 [$IMPERSONATE]
   --import-formats value               Googleドキュメントをアップロードする優先フォーマットのカンマ区切りリスト。 [$IMPORT_FORMATS]
   --keep-revision-forever              各ファイルの新しいヘッドリビジョンを永久に保持する。(default: false) [$KEEP_REVISION_FOREVER]
   --list-chunk value                   リストのチャンクサイズ 100-1000、0で無効。 (default: 1000) [$LIST_CHUNK]
   --pacer-burst value                  スリープせずに許可されるAPI呼び出しの数。 (default: 100) [$PACER_BURST]
   --pacer-min-sleep value              API呼び出し間の最小スリープ時間。 (default: "100ms") [$PACER_MIN_SLEEP]
   --resource-key value                 リンク共有されたファイルにアクセスするためのリソースキー。 [$RESOURCE_KEY]
   --root-folder-id value               ルートフォルダのID。 [$ROOT_FOLDER_ID]
   --server-side-across-configs         異なるドライブ構成間でサーバーサイドの操作（例：コピー）を実行することを許可する。(default: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --service-account-credentials value  サービスアカウント認証情報のJSON blob。 [$SERVICE_ACCOUNT_CREDENTIALS]
   --shared-with-me                     共有されたファイルのみを表示する。(default: false) [$SHARED_WITH_ME]
   --size-as-quota                      サイズをストレージとしての使用クオータ表示する。(default: false) [$SIZE_AS_QUOTA]
   --skip-checksum-gphotos              GoogleフォトおよびビデオのMD5チェックサムをスキップする。(default: false) [$SKIP_CHECKSUM_GPHOTOS]
   --skip-dangling-shortcuts            ダングリングショートカットファイルをスキップする場合は、設定する。(default: false) [$SKIP_DANGLING_SHORTCUTS]
   --skip-gdocs                         すべてのリストでGoogleドキュメントをスキップする。(default: false) [$SKIP_GDOCS]
   --skip-shortcuts                     ショートカットファイルをスキップする場合は、設定する。(default: false) [$SKIP_SHORTCUTS]
   --starred-only                       スターがついているファイルのみを表示する。(default: false) [$STARRED_ONLY]
   --stop-on-download-limit             ダウンロード制限エラーを致命的にする。(default: false) [$STOP_ON_DOWNLOAD_LIMIT]
   --stop-on-upload-limit               アップロード制限エラーを致命的にする。(default: false) [$STOP_ON_UPLOAD_LIMIT]
   --team-drive value                   共有ドライブ（チームドライブ）のID。 [$TEAM_DRIVE]
   --token value                        OAuthアクセストークン（JSON形式）。 [$TOKEN]
   --token-url value                    トークンサーバーのURL。 [$TOKEN_URL]
   --trashed-only                       ゴミ箱にあるファイルのみを表示する。(default: false) [$TRASHED_ONLY]
   --upload-cutoff value                チャンクアップロードに切り替えるための閾値。(default: "8Mi") [$UPLOAD_CUTOFF]
   --use-created-date                   更新日ではなく作成日を使用する。(default: false) [$USE_CREATED_DATE]
   --use-shared-date                    共有された日付（作成日）の代わりに更新日を使用する。(default: false) [$USE_SHARED_DATE]
   --use-trash                          ファイルを完全に削除する代わりにゴミ箱に送信する。(default: true) [$USE_TRASH]
   --v2-download-min-size value         オブジェクトがこれよりも大きい場合は、v2 APIを使用してダウンロードします。(default: "off") [$V2_DOWNLOAD_MIN_SIZE]

```
{% endcode %}