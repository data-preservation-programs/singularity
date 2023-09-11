# ローカルディスク

{% code fullWidth="true" %}
```
NAME:
   singularity storage update local - ローカルディスク

USAGE:
   singularity storage update local [command options] <name|id>

DESCRIPTION:
   --nounc
      Windows上でUNC（長いパス名）の変換を無効にします。

      例:
         | true | ファイル名が長い場合、無効にします。

   --copy-links
      シンボリックリンクに従い、指し示されたアイテムをコピーします。

   --links
      シンボリックリンクを '.rclonelink' 拡張子を持つ通常のファイルへ変換します。

   --skip-links
      スキップされるシンボリックリンクについて警告しないようにします。
      
      このフラグは、スキップされることが明示的に認識されているため、
      スキップされるシンボリックリンクまたはジャンクションポイントに関する警告メッセージを無効にします。

   --zero-size-links
      リンクの Stat サイズをゼロとして扱い（リンクを読み取る）、
      リンクの Stat サイズをゼロとして扱う（廃止予定）。
      
      rclone は、リンクの Stat サイズをリンクのサイズとして使用していましたが、
      これは次の場所で失敗します：
      
      - Windows
      - 一部の仮想ファイルシステム（LucidLinkなど）
      - Android
      
      したがって、rclone は常にリンクを読み取ります。
      

   --unicode-normalization
      パスとファイル名に unicode NFC 正規化を適用します。
      
      このフラグは、ローカルファイルシステムから読み取られるファイル名を
      Unicode NFC 形式に正規化するために使用できます。
      
      rclone は通常、ファイルシステムから読み取ったファイル名のエンコーディングを変更しません。
      
      これは、macOS を使用している場合に便利です。macOS は通常、分解された（NFD）の Unicode を提供し、
      いくつかの OS 上で適切に表示されない場合があります（韓国語など）。
      
      rclone は、同期処理でファイル名を unicode 正規化で比較するため、
      通常はこのフラグを使用しないでください。

   --no-check-updated
      アップロード中のファイルが変更されたかどうかをチェックしません。
      
      通常、rclone はファイルのサイズと変更日時をチェックし、
      アップロード中にファイルが変更された場合、"can't copy - source file is being updated" で始まるメッセージで中止します。
      
      ただし、一部のファイルシステムでは、この変更日時のチェックが失敗する場合があります（例：[Glusterfs #2206](https://github.com/rclone/rclone/issues/2206)）ので、
      このチェックを無効にできるようになっています。
      
      このフラグが設定されている場合、rclone はファイルのアップデート中に最善の努力を行い、
      ファイルに追加されるもの（ログなど）に関しては、rclone が最初にそのファイルを見たときのサイズでログファイルを転送します。
      
      ファイルが終始変更される場合は、ハッシュチェックの失敗で転送が失敗する場合があります。
      
      詳しくは、最初に stat() がファイルに対して呼び出された後、
      次のようにします：
      
      - stat が提供したサイズのみ転送します
      - stat が提供したサイズのみチェックサムします
      - ファイルの stat 情報を更新しません

   --one-file-system
      ファイルシステム境界を越えないでください（unix/macOS のみ）。

   --case-sensitive
      ファイルシステム自体を大文字・小文字を区別すると報告するようにします。
      
      通常、ローカルバックエンドは Windows/macOS では大文字小文字を区別せず、
      それ以外のすべての場合には大文字小文字を区別するように宣言します。既定値を上書きするには、
      このフラグを使用します。

   --case-insensitive
      ファイルシステム自体を大文字小文字を区別しないと報告するようにします。
      
      通常、ローカルバックエンドは Windows/macOS では大文字小文字を区別せず、
      それ以外のすべての場合には大文字小文字を区別するように宣言します。既定値を上書きするには、
      このフラグを使用します。

   --no-preallocate
      転送されるファイルのディスクスペースの事前割り当てを無効にします。
      
      ディスクスペースの事前割り当ては、ファイルシステムの断片化を防ぐのに役立ちます。
      ただし、Google Drive File Stream などの一部の仮想ファイルシステムレイヤは、
      実際のファイルサイズを事前に割り当てられたスペースと同じに設定する場合があり、
      チェックサムチェックとファイルサイズチェックが失敗する場合があります。
      事前割り当てを無効にするには、このフラグを使用します。

   --no-sparse
      マルチスレッドダウンロード用のスパースファイルを無効にします。
      
      Windows プラットフォームでは、rclone はマルチスレッドダウンロード時にスパースファイルを作成します。
      これにより、OSがファイルをゼロ化することによる長い一時停止が回避されます。
      ただし、スパースファイルはディスクの断片化を引き起こす可能性があり、
      操作速度が低下する場合があります。

   --no-set-modtime
      modtime の設定を無効にします。
      
      通常、rclone はファイルのアップロードが完了した後に、ファイルの更新日時を更新します。
      Linux プラットフォームで rclone が実行されているユーザーがアップロードしたファイルの所有者ではない場合、
      この操作により許可エラーが発生する可能性があります。
      このオプションが有効にされている場合、rclone はファイルをコピーした後も modtime を更新しません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h  ヘルプを表示する

   Advanced

   --case-insensitive       ファイルシステム自体を大文字小文字を区別しないと報告する。
                            (既定値: false) [$CASE_INSENSITIVE]
   --case-sensitive         ファイルシステム自体を大文字小文字を区別すると報告する。
                            (既定値: false) [$CASE_SENSITIVE]
   --copy-links, -L         シンボリックリンクに従い、指し示されたアイテムをコピーする。
                            (既定値: false) [$COPY_LINKS]
   --encoding value         バックエンドのエンコーディング。
                            (既定値: "Slash,Dot") [$ENCODING]
   --links, -l              シンボリックリンクを '.rclonelink' 拡張子を持つ通常のファイルへ変換する。
                            (既定値: false) [$LINKS]
   --no-check-updated       アップロード中のファイルが変更されたかどうかをチェックしない。
                            (既定値: false) [$NO_CHECK_UPDATED]
   --no-preallocate         転送されるファイルのディスクスペースの事前割り当てを無効にする。
                            (既定値: false) [$NO_PREALLOCATE]
   --no-set-modtime         modtime の設定を無効にする。
                            (既定値: false) [$NO_SET_MODTIME]
   --no-sparse              マルチスレッドダウンロード用のスパースファイルを無効にする。
                            (既定値: false) [$NO_SPARSE]
   --nounc                  Windows上でUNC（長いパス名）の変換を無効にする。
                            (既定値: false) [$NOUNC]
   --one-file-system, -x    ファイルシステム境界を越えないでください（unix/macOS のみ）。
                            (既定値: false) [$ONE_FILE_SYSTEM]
   --skip-links             スキップされるシンボリックリンクについて警告しない。
                            (既定値: false) [$SKIP_LINKS]
   --unicode-normalization  パスとファイル名に unicode NFC 正規化を適用する。
                            (既定値: false) [$UNICODE_NORMALIZATION]
   --zero-size-links        リンクの Stat サイズをゼロとして扱う（リンクを読み取る）、
                            リンクの Stat サイズをゼロとして扱う（廃止予定）。
                            (既定値: false) [$ZERO_SIZE_LINKS]

```
{% endcode %}