# ローカルディスク

{% code fullWidth="true" %}
```
NAME:
   singularity storage create local - ローカルディスク

USAGE:
   singularity storage create local [command options] [arguments...]

DESCRIPTION:
   --nounc
      WindowsのUNC（長いパス名）変換を無効にします。

      例:
         | true | 長いファイル名が無効になります。

   --copy-links
      シンボリックリンクに従って、リンク先のアイテムをコピーします。

   --links
      シンボリックリンクを拡張子'.rclonelink'を使用して通常のファイルに変換/変換します。

   --skip-links
      スキップされたシンボリックリンクに関する警告メッセージを表示しません。

      このフラグを設定すると、明示的にスキップする必要があると認識されるため、
      スキップされたシンボリックリンクやジャンクションポイントに関する警告メッセージは無効になります。

   --zero-size-links
      リンクのStatサイズがゼロであると仮定します（それらを読み取ります）（非推奨）。

      Rcloneは、リンクのStatサイズをリンクのサイズとして使用しましたが、これは次の場所で失敗します：

      - Windows
      - 一部の仮想ファイルシステム（LucidLinkなど）
      - Android

      したがって、rcloneは常にリンクを読み取るようになりました。

   --unicode-normalization
      パスとファイル名にUnicode NFC正規化を適用します。

      このフラグは、ローカルファイルシステムから読み込まれる
      Unicode NFC形式のファイル名を正規化するために使用できます。

      Rcloneは通常、ファイルシステムから読み取ったファイル名のエンコーディングには触れません。

      これは、macOSを使用している場合に便利です。macOSは通常、分解された（NFDの）Unicodeを提供します。
      そのため、一部の言語（例：韓国語）は一部のOSで正しく表示されません。

      rcloneでは、同期ルーチン中にファイル名をUnicode正規化で比較するため、通常はこのフラグを使用しないでください。

   --no-check-updated
      ファイルがアップロード中に変更されたかどうかをチェックしません。

      通常、rcloneはファイルをアップロードしている間に、そのサイズと変更時刻をチェックして、
      ファイルが変更された場合にはじまるメッセージ "can't copy - source file is being updated" で中止します。

      ただし、一部のファイルシステムでは、この変更時刻のチェックが失敗する場合があります（例：[Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)）ので、
      このチェックはこのフラグで無効にすることができます。

      このフラグが設定されている場合、rcloneはアップデートされているファイルを転送するために最善の努力をします。
      ファイルに何かが追加されている場合（例：ログファイル）、rcloneは最初に見たときのサイズでログファイルを転送します。

      ファイルが常に変更されている場合（単に追記されているだけでない場合）、転送はハッシュチェックエラーで失敗する場合があります。

      詳細については、ファイルに最初にstat（）が呼び出された後、次を実行します：

      - statが与えたサイズのみを転送
      - statが与えたサイズのみをチェックサム
      - ファイルのstat情報を更新しない

   --one-file-system
      ファイルシステムの境界を越えません（unix/macOSのみ）。

   --case-sensitive
      ファイルシステムを大文字小文字を区別すると報告するように強制します。

      通常、ローカルバックエンドはWindows/macOSでは大文字小文字を区別しないと宣言し、他のすべてでは大文字小文字を区別するようになっています。
      このフラグを使用してデフォルトの選択肢をオーバーライドします。

   --case-insensitive
      ファイルシステムを大文字小文字を区別しないと報告するように強制します。

      通常、ローカルバックエンドはWindows/macOSでは大文字小文字を区別しないと宣言し、他のすべてでは大文字小文字を区別するようになっています。
      このフラグを使用してデフォルトの選択肢をオーバーライドします。

   --no-preallocate
      転送されたファイルのディスクスペースの事前確保を無効にします。

      ディスクスペースの事前確保は、ファイルシステムの断片化を防ぐのに役立ちます。
      ただし、一部の仮想ファイルシステムレイヤ（Google Drive File Streamなど）は、
      実際のファイルサイズを事前に確保されたスペースと同じに設定する場合があり、
      チェックサムとファイルサイズのチェックに失敗することがあります。
      このフラグを使用して、事前確保を無効にします。

   --no-sparse
      マルチスレッドダウンロードのスパースファイルを無効にします。

      Windowsプラットフォームでは、マルチスレッドダウンロード時にスパースファイルを作成します。
      これにより、OSがファイルのゼロ化を行う大きなファイルで長い一時停止が回避されます。
      ただし、スパースファイルはディスクの断片化を引き起こし、操作が遅くなる場合があります。

   --no-set-modtime
      modtimeの設定を無効にします。

      通常、rcloneはファイルのアップロード後に修正日時を更新します。
      これにより、Linuxプラットフォームでのアクセス権の問題が発生する場合があります。
      たとえば、別のユーザーが所有するCIFSマウントにコピーする場合などです。

      このオプションが有効になっている場合、rcloneはファイルをコピーした後、modtimeを更新しなくなります。

   --encoding
      バックエンドのエンコーディング。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h  ヘルプを表示する

   Advanced

   --case-insensitive       ファイルシステムを大文字小文字を区別しないと報告するように強制します。 (default: false) [$CASE_INSENSITIVE]
   --case-sensitive         ファイルシステムを大文字小文字を区別すると報告するように強制します。 (default: false) [$CASE_SENSITIVE]
   --copy-links, -L         シンボリックリンクに従って、リンク先のアイテムをコピーします。 (default: false) [$COPY_LINKS]
   --encoding value         バックエンドのエンコーディング。 (default: "Slash,Dot") [$ENCODING]
   --links, -l              シンボリックリンクを拡張子'.rclonelink'を使用して通常のファイルに変換/変換します。 (default: false) [$LINKS]
   --no-check-updated       ファイルがアップロード中に変更されたかどうかをチェックしません。 (default: false) [$NO_CHECK_UPDATED]
   --no-preallocate         転送されたファイルのディスクスペースの事前確保を無効にします。 (default: false) [$NO_PREALLOCATE]
   --no-set-modtime         modtimeの設定を無効にします。 (default: false) [$NO_SET_MODTIME]
   --no-sparse              マルチスレッドダウンロードのスパースファイルを無効にします。 (default: false) [$NO_SPARSE]
   --nounc                  WindowsのUNC（長いパス名）変換を無効にします。 (default: false) [$NOUNC]
   --one-file-system, -x    ファイルシステムの境界を越えません（unix/macOSのみ）。 (default: false) [$ONE_FILE_SYSTEM]
   --skip-links             スキップされたシンボリックリンクに関する警告メッセージを表示しません。 (default: false) [$SKIP_LINKS]
   --unicode-normalization  パスとファイル名にUnicode NFC正規化を適用します。 (default: false) [$UNICODE_NORMALIZATION]
   --zero-size-links        リンクのStatサイズがゼロであると仮定します（それらを読み取ります）（非推奨）。 (default: false) [$ZERO_SIZE_LINKS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}