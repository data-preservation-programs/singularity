# Microsoft OneDrive

{% code fullWidth="true" %}
```
NAME:
   singularity storage update onedrive - Microsoft OneDrive

USAGE:
   singularity storage update onedrive [コマンドオプション] <名前|ID>

DESCRIPTION:
   --client-id
      OAuth クライアントID。
      
      通常は空白のままにします。

   --client-secret
      OAuth クライアントシークレット。
      
      通常は空白のままにします。

   --token
      OAuth アクセストークン(JSON形式)。

   --auth-url
      認証サーバーのURL。
      
      プロバイダのデフォルト値を使用するには空白のままにします。

   --token-url
      トークンサーバーのURL。
      
      プロバイダのデフォルト値を使用するには空白のままにします。

   --region
      OneDriveの国別クラウドリージョンを選択します。

      例:
         | global | Microsoft クラウド グローバル版
         | us     | Microsoft クラウド US政府版
         | de     | Microsoft クラウド ドイツ版
         | cn     | 中国のVnet Groupが運営するAzureとOffice 365

   --chunk-size
      ファイルのアップロードに使用するチャンクサイズ - 320k (327,680 バイト)の倍数でなければなりません。
      
      このサイズを超える場合、ファイルはチャンク分割されます - 320k (327,680 バイト)の倍数でなければなりませんし、
      250M (262,144,000 バイト) を超えてはいけません。そうしないと、
      \"Microsoft.SharePoint.Client.InvalidClientQueryException: The request message is too big.\"のエラーが発生する可能性があります。
      チャンクはメモリにバッファリングされます。

   --drive-id
      使用するドライブのID。

   --drive-type
      ドライブのタイプ (personal | business | documentLibrary)。

   --root-folder-id
      ルートフォルダのID。
      
      通常は必要ありませんが、特殊な状況ではアクセスしたいフォルダのIDが分かっている場合は使用します。
      

   --access-scopes
      rcloneが要求するスコープを設定します。
      
      rcloneがリクエストするすべてのスコープをスペースで区切って手動で入力するか、選択します。
      

      例:
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access | すべてのリソースへの読み取りと書き込みのアクセス
         | Files.Read Files.Read.All Sites.Read.All offline_access                                     | すべてのリソースへの読み取り専用アクセス
         | Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All offline_access                | すべてのリソースへの読み取りと書き込みのアクセス、またはSharePointサイトのブラウズができないアクセス。
         |                                                                                             | disable_site_permissionをtrueに設定した場合と同様です。

   --disable-site-permission
      Sites.Read.All パーミッションのリクエストを無効にします。
      
      trueに設定すると、ドライブIDの設定時にSharePointサイトを検索できなくなります。
      これは、rcloneがSites.Read.Allパーミッションを要求しないためです。
      組織がアプリケーションにSites.Read.Allパーミッションを割り当てていない場合や、組織がユーザーにアプリのパーミッション要求を許可していない場合には、trueに設定します。

   --expose-onenote-files
      OneNoteファイルをディレクトリリストに表示するように設定します。
      
      デフォルトでは、rcloneはディレクトリリストでOneNoteファイルを非表示にします。
      "Open"や"Update"などのオペレーションはそれらには動作しません。
      ただし、この動作はそれらを削除することも妨げることがあります。
      OneNoteファイルを削除したり、ディレクトリリストに表示する場合は、このオプションを設定します。

   --server-side-across-configs
      サーバーサイドの操作 (コピーなど) を異なるonedrive設定間で動作させることを許可します。
      
      これは、2つのOneDrive *パーソナル* ドライブ間でコピーを行い、コピーするファイルが既に共有されている場合にのみ動作します。
      それ以外の場合、rcloneは通常のコピーにフォールバックしますが、若干遅くなります。

   --list-chunk
      リストのチャンクサイズ。

   --no-versions
      修正操作時にすべてのバージョンを削除します。
      
      OneDrive for Businessでは、新しいファイルをアップロードして既存のファイルを上書きすると、バージョンが作成され、
      変更日時を設定するとバージョンが作成されます。
      
      これらのバージョンはクォータから容量を消費します。
      
      このフラグはファイルのアップロードと変更日時の設定後にバージョンを確認し、
      最後のバージョン以外を削除します。
      
      **注意** OneDriveパーソナルでは現在、バージョンを削除することはできませんので、このフラグを使用しないでください。
      

   --link-scope
      linkコマンドで作成されるリンクのスコープを設定します。

      例:
         | anonymous    | リンクを持つ人はサインインする必要なくアクセスできます。
         |              | これには組織外の人も含まれる場合があります。
         |              | 匿名リンクサポートは管理者によって無効にされている場合もあります。
         | organization | 所属する組織（テナント）にサインインしたユーザーがリンクを使用してアクセスできます。
         |              | OneDrive for BusinessおよびSharePointでのみ使用できます。

   --link-type
      linkコマンドで作成されるリンクのタイプを設定します。

      例:
         | view  | アイテムへの読み取り専用リンクを作成します。
         | edit  | アイテムへの読み書きリンクを作成します。
         | embed | アイテムへの埋め込みリンクを作成します。

   --link-password
      linkコマンドで作成されるリンクのパスワードを設定します。
      
      現時点では、これはOneDriveパーソナルの有料アカウントのみで動作します。
      

   --hash-type
      バックエンドで使用されるハッシュを指定します。
      
      このオプションは使用されるハッシュタイプを指定します。"auto"に設定すると、デフォルトのハッシュであるQuickXorHashが使用されます。
      
      rclone 1.62以前では、Onedrive PersonalのデフォルトハッシュはSHA1でした。
      1.62以降、すべてのonedriveのデフォルトはQuickXorHashを使用するようになりました。
      SHA1ハッシュが必要な場合は、このオプションを適切に設定します。
      
      2023年7月から、QuickXorHashがOneDrive for BusinessとOneDriver Personalの唯一の利用可能なハッシュになります。
      
      "none"に設定すると、ハッシュは使用しません。
      
      リクエストされたハッシュがオブジェクトに存在しない場合、空の文字列が返されます。rcloneはこれを欠落したハッシュとして認識します。
      

      例:
         | auto     | Rcloneが最適なハッシュを選択します
         | quickxor | QuickXor
         | sha1     | SHA1
         | sha256   | SHA256
         | crc32    | CRC32
         | none     | なし - ハッシュを使用しません

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --client-id value      OAuth クライアントID。 [$CLIENT_ID]
   --client-secret value  OAuth クライアントシークレット。 [$CLIENT_SECRET]
   --help, -h             ヘルプを表示
   --region value         OneDriveの国別クラウドリージョンを選択します。 (デフォルト: "global") [$REGION]

   Advanced

   --access-scopes value         rcloneが要求するスコープを設定します。 (デフォルト: "Files.Read Files.ReadWrite Files.Read.All Files.ReadWrite.All Sites.Read.All offline_access") [$ACCESS_SCOPES]
   --auth-url value              認証サーバーのURL。 [$AUTH_URL]
   --chunk-size value            ファイルのアップロードに使用するチャンクサイズ。 (デフォルト: "10Mi") [$CHUNK_SIZE]
   --disable-site-permission     Sites.Read.All パーミッションのリクエストを無効にします。 (デフォルト: false) [$DISABLE_SITE_PERMISSION]
   --drive-id value              使用するドライブのID。 [$DRIVE_ID]
   --drive-type value            ドライブのタイプ (personal | business | documentLibrary)。 [$DRIVE_TYPE]
   --encoding value              バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --expose-onenote-files        OneNoteファイルをディレクトリリストに表示するように設定します。 (デフォルト: false) [$EXPOSE_ONENOTE_FILES]
   --hash-type value             バックエンドで使用されるハッシュを指定します。 (デフォルト: "auto") [$HASH_TYPE]
   --link-password value         linkコマンドで作成されたリンクのパスワードを設定します。 [$LINK_PASSWORD]
   --link-scope value            linkコマンドで作成されたリンクのスコープを設定します。 (デフォルト: "anonymous") [$LINK_SCOPE]
   --link-type value             linkコマンドで作成されたリンクのタイプを設定します。 (デフォルト: "view") [$LINK_TYPE]
   --list-chunk value            リストのチャンクサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --no-versions                 修正操作時にすべてのバージョンを削除します。 (デフォルト: false) [$NO_VERSIONS]
   --root-folder-id value        ルートフォルダのID。 [$ROOT_FOLDER_ID]
   --server-side-across-configs  サーバーサイドの操作 (コピーなど) を異なるonedrive設定間で動作させることを許可します。 (デフォルト: false) [$SERVER_SIDE_ACROSS_CONFIGS]
   --token value                 OAuth アクセストークン(JSON形式)。 [$TOKEN]
   --token-url value             トークンサーバーのURL。 [$TOKEN_URL]

```
{% endcode %}