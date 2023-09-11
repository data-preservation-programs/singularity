# IONOS Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 ionos - IONOS クラウド

USAGE:
   singularity storage update s3 ionos [コマンドオプション] <名前|ID>

DESCRIPTION:
   --env-auth
      実行時に AWS 認証情報を取得します（環境変数または環境変数が存在しない場合には EC2/ECS のメタデータから取得）。
      
      アクセスキーとシークレットアクセスキーが空白の場合にのみ適用されます。

      例:
         | false | 次の手順で AWS 認証情報を入力します。
         | true  | 環境（環境変数または IAM）から AWS 認証情報を取得します。

   --access-key-id
      AWS アクセスキーID。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --secret-access-key
      AWS シークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --region
      バケットの作成およびデータの保存に使用するリージョン。
      

      例:
         | de           | ドイツ、フランクフルト
         | eu-central-2 | ドイツ、ベルリン
         | eu-south-2   | スペイン、ログローニョ

   --endpoint
      IONOS S3 オブジェクトストレージのエンドポイント。
      
      同じリージョンのエンドポイントを指定します。

      例:
         | s3-eu-central-1.ionoscloud.com | ドイツ、フランクフルト
         | s3-eu-central-2.ionoscloud.com | ドイツ、ベルリン
         | s3-eu-south-2.ionoscloud.com   | スペイン、ログローニョ

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用される定義済み ACL。
      
      この ACL は、オブジェクトの作成時およびバケット acl が設定されていない場合にバケットの作成時にも使用されます。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      サーバーサイドでオブジェクトをコピーする場合、ターゲットから ACL をコピーせずに新しい ACL を書き込むため、この ACL が適用されます。
      
      acl が空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用される定義済み ACL。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      この ACL は、バケットの作成時にのみ適用されます。設定されていない場合は "acl" が代わりに使用されます。
      
      "acl" および "bucket_acl" のいずれも空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。
      

      例:
         | private            | 所有者に FULL_CONTROL 権限があります。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | 所有者に FULL_CONTROL 権限があります。
         |                    | AllUsers グループには READ アクセス権限があります。
         | public-read-write  | 所有者に FULL_CONTROL 権限があります。
         |                    | AllUsers グループには READ および WRITE アクセス権限があります。
         |                    | バケット上でこれを許可することは推奨されていません。
         | authenticated-read | 所有者に FULL_CONTROL 権限があります。
         |                    | AuthenticatedUsers グループには READ アクセス権限があります。

   --upload-cutoff
      チャンク化アップロードに切り替えるためのカットオフ。
      
      このサイズより大きなファイルは、chunk_size のチャンクでアップロードされます。
      最小は0、最大は5 GiB です。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoff よりも大きなファイル、またはサイズが不明なファイル（たとえば "rclone rcat" でアップロードされたり、"rclone mount" や google フォトや google ドキュメントでアップロードされたりしたファイル）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      メモリあたりの転送ごとに "--s3-upload-concurrency" チャンクのサイズがバッファリングされます。
      
      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを増やすと転送速度が向上します。
      
      Rclone は、10,000 チャンクの制限以下になるように、既知のサイズの大きなファイルをアップロードする場合に自動的にチャンクサイズを増やします。
      
      サイズが不明なファイルは、設定された chunk_size でアップロードされます。
      デフォルトのチャンクサイズが5 MiB であり、最大で 10,000 個のチャンクがあるため、デフォルトではストリームアップロードできるファイルの最大サイズは 48 GiB です。 より大きなファイルのストリームアップロードを行いたい場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを増やすと、進行状況の統計情報が "-P" フラグで表示されるときの正確性が低下します。 Rclone は、AWS SDK によってチャンクがバッファリングされたときにチャンクが送信されたと見なすため、実際にはまだアップロードされている場合でもそのように報告します。
      チャンクサイズが大きいほど、AWS SDK のバッファサイズが大きくなり、進行状況は真実から逸脱します。
      

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数。
      
      このオプションは、マルチパートアップロード時に使用するパーツの最大数を定義します。
      
      サービスが AWS S3 の 10,000 チャンクという仕様をサポートしていない場合に使用することができます。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする場合に、このチャンクサイズを増やして制限数のチャンク数未満になるように自動的にチャンクサイズを増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。
      
      サーバーサイドでコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小は0、最大は5 GiB です。

   --disable-checksum
      オブジェクトメタデータに MD5 チェックサムを保存しません。
      
      通常、rclone はアップロード前に入力の MD5 チェックサムを計算して、オブジェクトのメタデータに追加します。これはデータの整合性チェックには便利ですが、大きなファイルのアップロードの開始に長い遅延が発生することがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を検索します。環境変数の値が空の場合は、デフォルトで現在のユーザーのホームディレクトリが使用されます。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は、環境変数 "AWS_PROFILE" または "default" が設定されていない場合はデフォルトになります。
      

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。
      
      同じファイルの複数のチャンクが同時にアップロードされます。
      
      高速リンクで大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      「true」の場合、パススタイルアクセスを使用します。「false」の場合、仮想ホストスタイルを使用します。
      
      これが「true」（デフォルト）の場合、rclone はパススタイルアクセスを使用します。これが「false」の場合、rclone は仮想パススタイルを使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（例: AWS、Aliyun OSS、Netease COS、または Tencent COS）は、この設定に応じて自動的にこちらを設定する必要があります。

   --v2-auth
      「true」の場合、v2 認証を使用します。
      
      これが「false」（デフォルト）の場合、rclone は v4 認証を使用します。設定されている場合、rclone は v2 認証を使用します。
      
      v4 署名が機能しない場合にのみ使用してください（Jewel/v10 CEPH の場合など）。

   --list-chunk
      リストチャンクのサイズ（各 ListObject S3 リクエストのレスポンスリスト）。
      
      このオプションは、AWS S3 仕様の「MaxKeys」、「max-items」、または「page-size」としても知られています。
      大部分のサービスは、リクエスト数が 1000 を超えてもレスポンスリストを切り捨てます。
      AWS S3 では、これはグローバルな最大値であり、変更できません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Ceph の場合、これは「rgw list buckets max chunk」オプションで増やすことができます。

   --list-version
      使用する ListObjects のバージョン: 1、2、または 0（自動）。
      
      S3 の最初のリリース時、バケット内のオブジェクトを列挙するために ListObjects 呼び出しが提供されました。
      
      しかし、2016 年 5 月に ListObjectsV2 呼び出しが導入されました。これははるかに高性能であり、可能であれば使用する必要があります。
      
      デフォルト（0）に設定されている場合、rclone はプロバイダ設定に基づいてどのリストオブジェクトメソッドを呼び出すか推測します。推測が間違っている場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストを URL エンコードするかどうか: true/false/unset
      
      一部のプロバイダは、リストを URL エンコードすることをサポートし、ファイル名に制御文字を使用する場合は、これが使用できます。これが「unset」（デフォルト）に設定されている場合、rclone は設定に従ってプロバイダの設定に適用するものを選択しますが、ここで rclone の選択を上書きすることもできます。
      

   --no-check-bucket
      バケットの存在チェックや作成を試みないようにします。
      
      バケットが既に存在することを知っている場合に、rclone が行うトランザクション数を最小限にする必要がある場合に有用です。
      
      バケットの作成権限がない場合も必要です。v1.52.0 より前のバージョンでは、このバグのため、これは黙って通過していました。
      

   --no-head
      アップロードされたオブジェクトの HEAD リクエストを行って整合性をチェックしません。
      
      rclone が PUT 後に 200 OK メッセージを受け取った場合、正しくアップロードされたと想定します（実行時チェックを行いません）。
      
      特に次のことを想定します。
      
      - アップロード時のメタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロード時と同じであること
      - サイズがアップロード時と同じであること
      
      単一パート PUT のレスポンスから、次の項目を読み取ります。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明なソースオブジェクトがアップロードされる場合、rclone **は** HEAD リクエストを行います。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まります。特に、サイズが正しくない場合などです。そのため、通常の操作には推奨されません。実際には、このフラグを設定しても、アップロードの失敗が検出されない確率は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前に HEAD を行わない場合に設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。
      
      追加バッファが必要なアップロード（たとえばマルチパート）は、アロケーションにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから取り除かれる頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールに mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。
      
      現在、s3（特に minio）バックエンドと HTTP/2 に関する未解決の問題があります。S3 バックエンドの HTTP/2 はデフォルトで有効になっていますが、ここで無効にすることもできます。この問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      通常、AWS S3 は CloudFront ネットワーク経由でダウンロードされたデータの Egress よりも安価となる CloudFront CDN URL に設定されます。

   --use-multipart-etag
      マルチパートアップロードでは ETag を使用して検証しますか
      
      これは true、false、またはデフォルト（プロバイダによる）のいずれかである必要があります。
      

   --use-presigned-request
      単一パートアップロードに署名済みリクエストまたは PutObject を使用するかどうか
      
      これが false の場合、rclone はオブジェクトをアップロードするために AWS SDK の PutObject を使用します。
      
      rclone < 1.59 のバージョンでは、署名済みリクエストを使用して単一パートオブジェクトをアップロードし、このフラグを true に設定すると、その機能が再度有効になります。これは、特別な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した日時のファイルバージョンを表示します。
      
      パラメータは、日付 "2006-01-02"、日時 "2006-01-02 15:04:05"、またはその長い時間前の期間 "100d" や "1h" などである必要があります。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。
      
      有効な形式については、[time オプションのドキュメント](/docs/#time-option) を参照してください。
      

   --decompress
      この場合、gzip エンコードされたオブジェクトを展開します。
      
      "Content-Encoding: gzip" が設定された状態で S3 にオブジェクトをアップロードすることが可能です。通常、rclone はこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rclone はこれらのファイルを受信時に "Content-Encoding: gzip" で展開するようになります。これにより、rclone はサイズとハッシュを確認できませんが、ファイルのコンテンツは展開されます。
      

   --might-gzip
      バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはダウンロード時にオブジェクトを変更しません。`Content-Encoding: gzip` でアップロードされていない場合、ダウンロード時にも設定されません。
      
      ただし、一部のプロバイダは、`Content-Encoding: gzip` でアップロードされていなくてもオブジェクトを gzip 圧縮する場合があります（例: Cloudflare）。
      
      これを設定すると、rclone が `Content-Encoding: gzip` が設定され、チャンク化転送エンコードがあるオブジェクトをダウンロードした場合、rclone はオブジェクトをリアルタイムで展開します。
      
      これが未設定（デフォルト）に設定されている場合、rclone は設定によってプロバイダの設定に従って適用する内容を選択しますが、ここで rclone の選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制


OPTIONS:
   --access-key-id value      AWS アクセスキー ID。[$ACCESS_KEY_ID]
   --acl value                バケットの作成および保存またはコピー時に使用される定義済み ACL。[$ACL]
   --endpoint value           IONOS S3 オブジェクトストレージのエンドポイント。[$ENDPOINT]
   --env-auth                 実行時に AWS 認証情報を取得します（環境変数または環境変数が存在しない場合には EC2/ECS のメタデータから取得）。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             バケットの作成およびデータの保存に使用するリージョン。[$REGION]
   --secret-access-key value  AWS シークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用される定義済み ACL。[$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     この場合、gzip エンコードされたオブジェクトを展開します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータに MD5 チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドでの http2 の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               「true」の場合、パススタイルアクセスを使用します。「false」の場合、仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストを URL エンコードするかどうか。 (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1,2, または 0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールに mmap バッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在チェックや作成を試みないようにします。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトの HEAD リクエストを行って整合性をチェックしません。 (default: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前に HEAD を行わない場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWS セッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのカットオフ。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでは ETag を使用して検証しますか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パートアップロードに署名済みリクエストまたは PutObject を使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        「true」の場合、v2 認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した日時のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (default: false) [$VERSIONS]

```
{% endcode %}