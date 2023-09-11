# Microsoft OneDrive

{% code fullWidth="true" %}
```
名前:
   singularity storage create onedrive - Microsoft OneDrive

使用法:
   singularity storage create onedrive [コマンドオプション] [引数...]

説明:
   --client-id
      OAuth クライアント ID。
      
      通常は空白のままにします。

   --client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにします。

   --token
      JSON blob 形式の OAuth アクセストークン。

   --auth-url
      認証サーバーの URL。
      
      提供元のデフォルトを使用する場合は空白のままにします。

   --token-url
      トークンサーバーの URL。
      
      提供元のデフォルトを使用する場合は空白のままにします。

   --region
      OneDrive の国別クラウドリージョンを選択します。

      例:
         | global | Microsoft Cloud Global
         | us     | Microsoft Cloud for US Government
         | de     | Microsoft Cloud Germany
         | cn     | 中国で運用される Azure および Office 365

   --chunk-size
      ファイルアップロード時のチャンクサイズ - 320k (327,680 バイト) の倍数である必要があります。
      
      このサイズ以上のファイルはチャンク分割されます - 320k (327,680 バイト) の倍数である必要があり、
      250M (262,144,000 バイト) を超えないようにしてください。そうしないと「Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.」のエラーが発生する可能性があります。
      チャンクはメモリにバッファリングされます。

   --drive-id
      利用するドライブの ID。

   --drive-type
      ドライブのタイプ (personal | business | documentLibrary)。

   --root-folder-id
      ルートフォルダーの ID。
      
      通常は必要ありませんが、特殊な状況でアクセスしたいフォルダーの ID を知っているが、パストラバーサルではアクセスできない場合に使用します。
      

   --access-scopes
      rclone が要求するスコープを設定します。
      
      rclone が要求するすべてのスコープのカスタムスペース区切りのリストを選択または手動で入力します。
      

      例:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | すべてのリソースへの読み取りと書き込みアクセス
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | すべてのリソースへの読み取り専用アクセス
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | すべてのリソースへの読み取りと書き込みアクセス。ただし、SharePoint サイトの参照はできません。
         |                                                                                             | disable_site_permission が true に設定されている場合と同様です

   --disable-site-permission
      Sites.Read.All パーミッションの要求を無効にします。
      
      このオプションを true に設定すると、ドライブ ID を構成する際に SharePoint サイトを検索することができなくなります。なぜなら、rclone は Sites.Read.All パーミッションを要求しなくなるからです。
      組織がアプリケーションに Sites.Read.All パーミッションを割り当てていない場合や、ユーザーが独自にアプリケーションのパーミッション要求を承認できない場合に true に設定してください。

   --expose-onenote-files
      OneNote ファイルをディレクトリリストに表示するように設定します。
      
      デフォルトでは、rclone はディレクトリリストで OneNote ファイルを非表示にします。なぜなら「開く」や「更新」といった操作が機能しないためです。ただし、これによりファイルの削除もできなくなる場合があります。OneNote ファイルを削除したい場合やディレクトリリストに表示したい場合は、このオプションを設定してください。

   --server-side-across-configs
      サーバーサイドの操作（例: コピー）を異なる OneDrive の設定間で有効にします。
      
      これは、2 つの OneDrive *個人用* ドライブ間でコピーを実行し、コピーするファイルがすでに共有されている場合のみ機能します。それ以外の場合、rclone は通常のコピーにフォールバックします（やや遅くなります）。

   --list-chunk
      リストのチャンクサイズ。

   --no-versions
      修正操作時にすべてのバージョンを削除します。
      
      Onedrive for business は、新しいファイルをアップロードして既存のファイルを上書きしたときや、更新時にバージョンを作成します。
      
      これらのバージョンはクオータからスペースを消費します。
      
      このフラグは、ファイルのアップロードや更新時にバージョンをチェックし、最後のバージョン以外をすべて削除します。
      
      **注意** Onedrive personal では現在、バージョンを削除することができませんので、このフラグは使用しないでください。
      

   --link-scope
      link コマンドで作成されるリンクのスコープを設定します。

      例:
         | anonymous    | リンクを持っている人はサインインする必要なくアクセスできます。
         |              | これには、組織外の人々も含まれる場合があります。
         |              | 管理者によって匿名リンクのサポートが無効にされている場合があります。
         | organization | 組織（テナント）にサインインしたユーザーはリンクを使用してアクセスできます。
         |              | OneDrive for Business と SharePoint でのみ使用できます。

   --link-type
      link コマンドで作成されるリンクのタイプを設定します。

      例:
         | view  | アイテムへの読み取り専用リンクを作成します。
         | edit  | アイテムへの読み取りと書き込みのリンクを作成します。
         | embed | アイテムへの埋め込みリンクを作成します。

   --link-password
      link コマンドで作成されるリンクのパスワードを設定します。
      
      現時点では、これは OneDrive 個人の有料アカウントのみで機能します。

   --hash-type
      バックエンドで使用するハッシュを指定します。
      
      これは使用するハッシュタイプを指定します。"auto" に設定すると、既定のハッシュタイプである QuickXorHash が使用されます。
      
      rclone 1.62 より前のバージョンでは、OneDrive Personal の既定のハッシュとして SHA1 ハッシュが使用されていました。1.62 以降では、すべての OneDrive タイプでデフォルトとして QuickXorHash を使用するようになりました。SHA1 ハッシュを使用する場合は、このオプションを適切に設定してください。
      
      2023 年 7 月以降、QuickXorHash が OneDrive for Business および OneDriver Personal の唯一の利用可能なハッシュになります。
      
      "none" に設定すると、ハッシュを使用しないようになります。
      
      要求されたハッシュがオブジェクトに存在しない場合、空の文字列として返され、rclone によって欠落したハッシュとして扱われます。
      

      例:
         | auto     | Rclone が最適なハッシュを選択します。
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | ハッシュを使用しない

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --client-id value      OAuth クライアント ID。[$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。[$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --region value         OneDrive の国別クラウドリージョンを選択します。 (デフォルト: "global") [$REGION]

   Advanced

   --access-scopes value         rclone が要求するスコープを設定します。 (デフォルト: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              認証サーバーの URL。[$AUTH_URL]
   --chunk-size value            ファイルアップロード時のチャンクサイズ - 320k (327,680 バイト) の倍数である必要があります。 (デフォルト: "10Mi") [$CHUNK_SIZE]
   --disable-site-permission     Sites.Read.All パーミッションの要求を無効にします。 (デフォルト: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              利用するドライブの ID。[$DRIVE_ID]
   --drive-type value            ドライブのタイプ (personal | business | documentLibrary)。[$DRIVE_TYPE]
   --encoding value              バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        OneNote ファイルをディレクトリリストに表示するように設定します。 (デフォルト: false) [$EXPOSE_ONENOTE_FILES]
   --hash-type value             バックエンドで使用するハッシュを指定します。 (デフォルト: "auto") [$HASH_TYPE]
   --link-password value         link コマンドで作成されるリンクのパスワードを設定します。[$LINK_PASSWORD]
   --link-scope value            link コマンドで作成されるリンクのスコープを設定します。 (デフォルト: "anonymous") [$LINK_SCOPE]
   --link-type value             link コマンドで作成されるリンクのタイプを設定します。 (デフォルト: "view") [$LINK_TYPE]
   --list-chunk value            リストのチャンクサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --no-versions                 修正操作時にすべてのバージョンを削除します。 (デフォルト: false) [$NO_VERSIONS]
   --root-folder-id value        ルートフォルダーの ID。[$ROOT_FOLDER_ID]
   --server-side-across-configs  サーバーサイドの操作（例: コピー）を異なる OneDrive の設定間で有効にします。 (デフォルト: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 JSON blob 形式の OAuth アクセストークン。[$TOKEN]
   --token-url value             トークンサーバーの URL。[$TOKEN_URL]

   General

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}