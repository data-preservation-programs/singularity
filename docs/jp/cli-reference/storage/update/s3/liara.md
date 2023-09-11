# Liara オブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 liara - Liara オブジェクトストレージ

USAGE:
   singularity storage update s3 liara [command options] <name|id>

DESCRIPTION:
   --env-auth
      ランタイムから AWS 認証情報を取得します（環境変数または環境変数がない場合は EC2/ECS メタデータから取得します）。

      access_key_id と secret_access_key が空の場合のみ適用されます。

      例:
         | false | 次のステップで AWS 認証情報を入力します。
         | true  | 環境から AWS 認証情報を取得します（環境変数または IAM）。

   --access-key-id
      AWS アクセスキー ID。

      匿名アクセスまたは ランタイム認証情報を使用する場合は空のままにしてください。

   --secret-access-key
      AWS シークレットアクセスキー（パスワード）。

      匿名アクセスまたは ランタイム認証情報を使用する場合は空のままにしてください。

   --endpoint
      Liara オブジェクトストレージ API のエンドポイント。

      例:
         | storage.iran.liara.space | デフォルトのエンドポイント
         |                          | イラン

   --acl
      オブジェクトを作成するときと、オブジェクトを保管またはコピーするときに使用する固定 ACL。

      この ACL はオブジェクトの作成に使用され、bucket_acl が設定されていない場合、バケットの作成にも使用されます。

      詳細については、[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3 はソースから ACL をコピーするのではなく、新しい ACL を書き込むため、この ACL はサーバー側でオブジェクトをコピーするときに適用されます。

      ACL が空の文字列の場合、X-Amz-Acl ヘッダは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットを作成するときに使用される固定 ACL。

      詳細については、[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      この ACL はバケットを作成するときにのみ適用されます。設定されていない場合は「acl」が代わりに使用されます。

      「acl」と「bucket_acl」が空の文字列の場合、X-Amz-Acl ヘッダは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | 所有者に FULL_CONTROL 権限があります。
         |                    | 他のユーザーはアクセス権がありません（デフォルト）。
         | public-read        | 所有者に FULL_CONTROL 権限があります。
         |                    | AllUsers グループは READ アクセス権を持ちます。
         | public-read-write  | 所有者に FULL_CONTROL 権限があります。
         |                    | AllUsers グループは READ および WRITE アクセス権を持ちます。
         |                    | バケットにこれを許可することは一般的に推奨されません。
         | authenticated-read | 所有者に FULL_CONTROL 権限があります。
         |                    | AuthenticatedUsers グループは READ アクセス権を持ちます。

   --storage-class
      新しいオブジェクトを保管する際に使用するストレージクラス。

      例:
         | STANDARD | 標準ストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ。

      このサイズを超えるファイルは、chunk_size のチャンクでアップロードされます。
      最小値は 0、最大値は 5 GiB です。

   --chunk-size
      アップロードに使用されるチャンクサイズ。

      upload_cutoff より大きなサイズのファイル、またはサイズが不明なファイル（例：「rclone rcat」でのアップロード、または「rclone mount」や Google フォトや Google ドキュメントからのアップロード）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。

      注意：
      "--s3-upload-concurrency" チャンクは、このサイズごとに転送ごとにメモリ内でバッファリングされます。

      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送が高速化されます。

      Rclone は、既知のサイズの大きなファイルをアップロードする際に、10,000 チャンクの制限を下回るようにチャンクサイズを自動的に増やします。

      サイズの不明なファイルは、設定済みの chunk_size でアップロードされます。デフォルトの chunk_size は 5 MiB であり、最大 10,000 のチャンクが可能です。したがって、デフォルトでは最大 48 GiB のファイルをストリーミングアップロードすることができます。より大きなファイルのストリーミングアップロードを行う場合は、chunk_size を増やす必要があります。

      チャンクサイズを増やすと、「-P」フラグで表示される進行状況統計の正確さが低下します。Rclone は、チャンクを AWS SDK にバッファリングしたときにチャンクを送信したとみなし、まだアップロードされている場合があります。チャンクサイズが大きいほど、AWS SDK のバッファと進行状況の報告が真実から逸脱してしまいます。

   --max-upload-parts
      マルチパートアップロードでのパートの最大数。

      このオプションは、マルチパートアップロードの際に使用するための multipart チャンクの最大数を定義します。

      10,000 チャンクの AWS S3 仕様をサポートしていないサービスの場合、このオプションが役立ちます。

      Rclone は、既知のサイズの大きなファイルをアップロードする際に、このチャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。

      サーバーサイドでコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。

      最小値は 0、最大値は 5 GiB です。

   --disable-checksum
      オブジェクトメタデータに MD5 チェックサムを保存しません。

      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加します。これはデータの整合性チェックには非常に役立ちますが、大きなファイルのアップロードの開始には長い遅延を引き起こすことがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。

      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を検索します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合、環境変数 "AWS_PROFILE" または "default" が設定されていない場合は、デフォルトになります。

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。

      同じファイルのチャンクのアップロードの並行数です。

      高速リンクを介して小数の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用していない場合、これを増やすことは転送の高速化に役立つかもしれません。

   --force-path-style
      true の場合、パス形式のアクセスを使用し、false の場合は仮想ホスト形式のアクセスを使用します。

      true（デフォルト）の場合、rclone はパス形式のアクセスを使用します。false の場合、rclone は仮想パス形式を使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、または Tencent COS など）は、これを false に設定する必要があります。rclone はプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      true の場合、v2 認証を使用します。

      false（デフォルト）の場合、rclone は v4 認証を使用します。設定されている場合、rclone は v2 認証を使用します。

      v4 シグネチャが機能しない場合にのみ使用してください。例：Jewel/v10 CEPH より前のバージョン。

   --list-chunk
      リストのチャンクサイズ（各 ListObject S3 リクエストのレスポンスリスト）。

      このオプションは、AWS S3 仕様の「MaxKeys」、「max-items」、「page-size」とも呼ばれます。
      多くのサービスは、要求した以上のオブジェクトのリストを 1000 に切り捨てます。
      AWS S3 では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」というオプションでこれを増やすことができます。

   --list-version
      使用する ListObjects のバージョン: 1、2、または 0（自動化）。

      S3 の最初のリリース時には、バケットのオブジェクトを列挙するために ListObjects コールしか提供されませんでした。

      ただし、2016 年 5 月に ListObjectsV2 コールが導入されました。これは非常に高速であるため、できる限り使用する必要があります。

      デフォルトの 0 の場合、rclone は、設定されたプロバイダに基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が誤っている場合は、ここで手動で設定することもできます。

   --list-url-encode
      リストを URL エンコードするかどうか：true/false/unset

      一部のプロバイダは、ファイル名に制御文字を使用する場合、リストを URL エンコードする機能をサポートしています。使用しない場合、rclone はプロバイダの設定に従って適用するものを選択します。ただし、ここで rclone の選択を上書きすることもできます。

   --no-check-bucket
      バケットの存在チェックまたは作成の試行を行わない場合、設定します。

      バケットがすでに存在する場合のトランザクションの数を最小限にすることを目指して、このフラグを設定すると便利です。

      ユーザーにバケット作成の権限がない場合にも必要になる場合があります。v1.52.0 以前では、このバグのために無音でパスされます。

   --no-head
      アップロードしたオブジェクトを HEAD リクエストして整合性を確認しない場合、設定します。

      rclone が PUT 後に 200 OK メッセージを受信した場合、正しくアップロードされたと想定します。

      特に、次の項目を想定します。

      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロードと同じであること
      - サイズがアップロードと同じであること

      シングルパート PUT の応答から次の項目を読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらの項目は読み取られません。

      ソースオブジェクトのサイズが不明の場合は HEAD リクエストを行います。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が増えます。特に、サイズの誤りの場合には非常に小さいため、通常の操作にはお勧めできません。実際には、このフラグを設定しても、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GET でオブジェクトを取得する前に HEAD を実行しない場合、設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる間隔。

      追加バッファが必要なアップロード（たとえば、マルチパート）は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。

      現在、s3（特に minio）バックエンドと HTTP/2 の問題が未解決です。s3 バックエンドではデフォルトで HTTP/2 が有効になっていますが、ここでは無効にできます。問題が解決したら、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3 が CloudFront ネットワークを介してダウンロードされたデータのために安価なエグレスを提供するため、CloudFront CDN URL に設定されます。

   --use-multipart-etag
      マルチパートアップロードで ETag を使用して検証するかどうか

      これは true、false、またはデフォルト（プロバイダの場合）に設定する必要があります。

   --use-presigned-request
      1 パートのアップロードにプリサイン済みリクエストまたは PutObject を使用するかどうか。

      false の場合、rclone はオブジェクトをアップロードするために AWS SDK の PutObject を使用します。

      rclone のバージョン < 1.59 では、1 パートのオブジェクトのアップロードにプリサインされたリクエストを使用し、このフラグを true に設定すると、その機能が再度有効になります。これは特別な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時間のファイルバージョンを表示します。

      パラメータは、日付「2006-01-02」、日時「2006-01-02 15:04:05」、そのときからの経過時間「100d」または「1h」である必要があります。

      ただし、このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり、削除したりすることはできません。

      有効な形式については、[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定された場合、gzip エンコードされたオブジェクトを解凍します。

      AWS S3 に "Content-Encoding: gzip" が設定された状態でオブジェクトをアップロードすることができます。通常、rclone はこれらのファイルを圧縮オブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rclone は受信したときに "Content-Encoding: gzip" のファイルを解凍します。これにより、rclone はサイズとハッシュを確認することはできませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトを gzip 圧縮する可能性がある場合、これを設定します。

      通常、プロバイダはオブジェクトをダウンロードする際には変更しません。`Content-Encoding: gzip` がアップロードされなかった場合、ダウンロード時には設定されません。

      ただし、一部のプロバイダは、`Content-Encoding: gzip` がアップロードされていない場合でもオブジェクトを gzip 圧縮する場合があります（例：Cloudflare）。

      次のようなエラーメッセージが表示される場合、これが原因でしまうかもしれません。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rclone が `Content-Encoding: gzip` とチャンク転送エンコーディングが設定されたオブジェクトをダウンロードする場合、rclone はオブジェクトを逐次解凍します。

      これが設定されている場合 unset（デフォルト）に設定されている場合、rclone はプロバイダの設定に従って適用するものを選択しますが、ここで rclone の選択を上書きすることもできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する

OPTIONS:
   --access-key-id value      AWS アクセスキー ID [$ACCESS_KEY_ID]
   --acl value                オブジェクトを作成するときと、オブジェクトを保管またはコピーするときに使用する固定 ACL [$ACL]
   --endpoint value           Liara オブジェクトストレージ API のエンドポイント [$ENDPOINT]
   --env-auth                 ランタイムから AWS 認証情報を取得します（環境変数または環境変数がない場合は EC2/ECS メタデータから取得します） (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --secret-access-key value  AWS シークレットアクセスキー（パスワード） [$SECRET_ACCESS_KEY]
   --storage-class value      新しいオブジェクトを保管する際に使用するストレージクラス [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットを作成するときに使用される固定 ACL [$BUCKET_ACL]
   --chunk-size value               アップロードに使用されるチャンクサイズ (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定された場合、gzip エンコードされたオブジェクトを解凍します (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータに MD5 チェックサムを保存しません (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドでの http2 の使用を無効にします (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合、パス形式のアクセスを使用し、false の場合は仮想ホスト形式のアクセスを使用します (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストを URL エンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1、2、または 0（自動） (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでのパートの最大数 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる間隔 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールで mmap バッファを使用するかどうか (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを gzip 圧縮する可能性がある場合、これを設定します (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在チェックまたは作成の試行を行わない場合、設定します (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトを HEAD リクエストして整合性を確認しない場合、設定します (default: false) [$NO_HEAD]
   --no-head-object                 GET でオブジェクトを取得する前に HEAD を実行しない場合 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル [$PROFILE]
   --session-token value            AWS セッショントークン [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードで ETag を使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          1 パートのアップロードにプリサイン済みリクエストまたは PutObject を使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true の場合、v2 認証を使用します (default: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める (default: false) [$VERSIONS]

```
{% endcode %}