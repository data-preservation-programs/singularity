# ローカルディスク

{% code fullWidth="true" %}
```
NAME:
   シンギュラリティ データソースの追加 local - ローカルディスク

使用方法:
   シンギュラリティ データソースの追加 local [コマンドオプション] <データセット名> <ソースパス>

説明:
   --local-case-insensitive
      ファイルシステムが大文字と小文字を区別しないと報告するようにします。

      通常、ローカルバックエンドはWindows/macOSでは大文字と小文字を区別せず、
      それ以外の場合は大文字と小文字を区別します。 デフォルトの設定を上書きするには、
      このフラグを使用します。

   --local-case-sensitive
      ファイルシステムが大文字と小文字を区別すると報告するようにします。

      通常、ローカルバックエンドはWindows/macOSでは大文字と小文字を区別せず、
      それ以外の場合は大文字と小文字を区別します。 デフォルトの設定を上書きするには、
      このフラグを使用します。

   --local-copy-links
      シンボリックリンクに従い、指し示されたアイテムをコピーします。

   --local-encoding
      バックエンドのエンコーディング。

     詳細については、 [概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --local-links
      シンボリックリンクを通常のファイルに '.rclonelink' 拡張子で変換します。

   --local-no-check-updated
      アップロード中にファイルが変更されたかどうかをチェックしません。

      通常、rcloneはファイルのサイズと変更日時をアップロード中にチェックし、
      ファイルがアップロード中に変更された場合は 「can't copy - source file is being updated」
      というメッセージで中止します。

      しかし、一部のファイルシステムではこの変更日時のチェックが失敗することがあります
      (例: [Glusterfs #2206](https://github.com/rclone/rclone/issues/2206))
      このフラグを使用すると、このチェックを無効にすることができます。

      このフラグが設定されている場合、rcloneはアップデート中のファイルを最善の努力で転送します。
      もしファイルが追記のみ行われている場合(例: ログファイル)、rcloneは最初にそのファイルを
      取得した時のサイズでログファイルを転送します。

      ファイルが常に修正される場合は、ハッシュチェックの失敗が発生する可能性があります。

      詳細については、ファイルの最初のstat()呼び出し後は次の操作が行われます。

      - statが提供したサイズだけを転送します。
      - statが提供したサイズだけをチェックサムします。
      - ファイルのstat情報を更新しません。

   --local-no-preallocate
      転送ファイルのディスクスペースの事前割り当てを無効にします。

      ディスクスペースの事前割り当ては、ファイルシステムの断片化を防ぐのに役立ちます。
      しかし、一部の仮想ファイルシステムレイヤ(例: Google Drive File Stream)
      では、実際のファイルサイズが事前割り当てされたスペースと等しく設定される場合があり、
      チェックサムやファイルサイズのチェックが失敗することがあります。
      事前割り当てを無効にするには、このフラグを使用します。

   --local-no-set-modtime
      モディファイドタイムの設定を無効にします。

      通常、rcloneはファイルのアップロード後に変更時刻を更新します。
      これは、Linuxプラットフォームでrcloneが実行されるユーザーが、
      アップロードされたファイルを所有していない場合(他のユーザーが所有するCIFSマウントに
      コピーする場合など)、パーミッションの問題を引き起こす可能性があります。
      このオプションが有効になっている場合、rcloneはファイルをコピーした後にモディファイドタイムを
      更新しません。

   --local-no-sparse
      マルチスレッドダウンロード時のスパースファイルを無効にします。

      Windowsプラットフォームでは、マルチスレッドダウンロード時にスパースファイルが作成されます。
      これにより、長い待ち時間が発生しないためです。なぜならOSがファイルをゼロで埋めるからです。
      しかし、スパースファイルはディスク断片化の原因になったり、扱いが遅くなる場合があります。

   --local-nounc
      WindowsでのUNC（長いパス名）の変換を無効にします。

      例:
         | true | 長いファイル名を無効にします。

   --local-one-file-system
      ファイルシステムの境界を越えないようにします（UNIX/macOSのみ）。

   --local-skip-links
      スキップされたシンボリックリンクについての警告メッセージを表示しません。

      このフラグは、スキップされるべきシンボリックリンクやジャンクションポイントに対する警告メッセージを無効にします。
      明示的にスキップすることを認識しているためです。

   --local-unicode-normalization
      パスとファイル名にUnicode NFC正規化を適用します。

      このフラグは、ローカルファイルシステムから読み取られるファイル名をUnicode NFC形式に正規化するために使用できます。

      通常、rcloneはファイルシステムから読み取ったファイル名のエンコードを操作しません。

      これは、macOSを使用する場合に便利です。macOSは通常、分解された(NFD)Unicodeを提供し、
      一部のOS（韓国語など）では正しく表示されません。

      rcloneは、同期ルーチンでファイル名を正規化して比較するため、通常はこのフラグを使用しないでください。

   --local-zero-size-links
      リンクのStatサイズをゼロと仮定し（読み取る代わりに）、リンクを読み取ります（非推奨）。

      以前のバージョンのrcloneでは、リンクのStatサイズをリンクサイズとして使用していましたが、
      これは次の場所で失敗します。

      - Windows
      - 一部の仮想ファイルシステム(例: LucidLink)
      - Android

      したがって、rcloneは常にリンクを読み取ります。
      
      
OPTIONS:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval value  前回の正常なスキャンからこのインターバルが経過すると、ソースディレクトリを自動的に再スキャンします（デフォルト: 無効）
   --scanning-state value   初期のスキャンステートを設定します（デフォルト: ready）

   ローカルのオプション

   --local-case-insensitive value       ファイルシステムが大文字と小文字を区別しないようにします（デフォルト: "false"） [$LOCAL_CASE_INSENSITIVE]
   --local-case-sensitive value         ファイルシステムが大文字と小文字を区別すると報告するようにします（デフォルト: "false"） [$LOCAL_CASE_SENSITIVE]
   --local-copy-links value             シンボリックリンクに従い、指し示されたアイテムをコピーします（デフォルト: "false"） [$LOCAL_COPY_LINKS]
   --local-encoding value               バックエンドのエンコーディング（デフォルト: "Slash,Dot"） [$LOCAL_ENCODING]
   --local-links value                  シンボリックリンクを通常のファイルに '.rclonelink' 拡張子で変換します（デフォルト: "false"） [$LOCAL_LINKS]
   --local-no-check-updated value       アップロード中にファイルが変更されたかどうかをチェックしません（デフォルト: "false"） [$LOCAL_NO_CHECK_UPDATED]
   --local-no-preallocate value         転送ファイルのディスクスペースの事前割り当てを無効にします（デフォルト: "false"） [$LOCAL_NO_PREALLOCATE]
   --local-no-set-modtime value         モディファイドタイムの設定を無効にします（デフォルト: "false"） [$LOCAL_NO_SET_MODTIME]
   --local-no-sparse value              マルチスレッドダウンロード時のスパースファイルを無効にします（デフォルト: "false"） [$LOCAL_NO_SPARSE]
   --local-nounc value                  WindowsでのUNC（長いパス名）の変換を無効にします（デフォルト: "false"） [$LOCAL_NOUNC]
   --local-one-file-system value        ファイルシステムの境界を越えないようにします（UNIX/macOSのみ）（デフォルト: "false"） [$LOCAL_ONE_FILE_SYSTEM]
   --local-skip-links value             スキップされたシンボリックリンクについての警告メッセージを表示しません（デフォルト: "false"） [$LOCAL_SKIP_LINKS]
   --local-unicode-normalization value  パスとファイル名にUnicode NFC正規化を適用します（デフォルト: "false"） [$LOCAL_UNICODE_NORMALIZATION]
   --local-zero-size-links value        リンクのStatサイズをゼロと仮定し（読み取る代わりに）、リンクを読み取ります（非推奨）（デフォルト: "false"） [$LOCAL_ZERO_SIZE_LINKS]

```
{% endcode %}