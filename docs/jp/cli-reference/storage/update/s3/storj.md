# Storj (S3 互換のゲートウェイ)

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 storj - Storj (S3 互換のゲートウェイ)

USAGE:
   singularity storage update s3 storj [command options] <name|id>

DESCRIPTION:
   --env-auth
      実行時に AWS 認証情報（環境変数または環境の EC2 / ECS メタデータ）を使用します。
      
      (access_key_id と secret_access_key が空白の場合のみ適用されます)
      
      例:
         | false | 次の設定で AWS 認証情報を入力します。
         | true  | 環境から AWS 認証情報（環境変数または IAM）を取得します。


   --access-key-id
      AWS アクセスキー ID。
      
      匿名アクセスまたは実行時の認証情報で空白のままにします。

   --secret-access-key
      AWS シークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報で空白のままにします。

   --endpoint
      Storj ゲートウェイのエンドポイント。

      例:
         | gateway.storjshare.io | グローバルホステッドゲートウェイ

   --bucket-acl
      バケット作成時に使用するキャニスド ACL。
      
      詳細については、[Amazon S3 ACL の概要](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      この ACL は、バケット作成時のみ適用されます。設定がされていない場合は、"acl" が代わりに使用されます。
      
      "acl" と "bucket_acl" の両方が空の文字列である場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト値（private）が使用されます。

      例:
         | private            | オーナーには FULL_CONTROL の権限が与えられます。
         |                    | 他のユーザーはアクセス権がありません（デフォルト）。
         | public-read        | オーナーには FULL_CONTROL の権限が与えられます。
         |                    | AllUsers グループには読み取り権限が与えられます。
         | public-read-write  | オーナーには FULL_CONTROL の権限が与えられます。
         |                    | AllUsers グループには読み取りおよび書き込み権限が与えられます。
         |                    | バケットでこれを許可することは推奨されていません。
         | authenticated-read | オーナーには FULL_CONTROL の権限が与えられます。
         |                    | AuthenticatedUsers グループには読み取り権限が与えられます。

   --upload-cutoff
      チャンクアップロードに切り替えるためのサイズカットオフ。
      
      このサイズより大きなファイルは、chunk_size のサイズでチャンクアップロードされます。
      最小値は 0 で、最大値は 5 GiB です。

   --chunk-size
      アップロードに使用するチャンクのサイズ。
      
      upload_cutoff より大きなサイズのファイルや、
      サイズが不明なファイル（例: "rclone rcat" で作成されたファイル、"rclone mount" でアップロードされたファイル、Google フォトや Google ドキュメントのアップロードなど）は、
      chunk_size を使用してマルチパートアップロードされます。
      
      注意:
      "--s3-upload-concurrency" で指定された数の chunk は、それらのサイズごとにメモリ内でバッファリングされます。
      
      高速リンクを介して大きなファイルを転送している場合、十分なメモリを利用できる場合は、チャンクサイズを増やすことで転送速度を向上させることができます。
      
      ロークライオンは、既知のサイズの大きなファイルをアップロードする場合、10,000 チャンクの制限を下回るように自動的にチャンクサイズを増やします。
      
      未知のサイズのファイルは、構成済みの chunk_size でアップロードされます。デフォルトの chunk_size が 5 MiB であり、最大 10,000 チャンクまで存在できるため、
      ファイルをストリームアップロードする場合の最大サイズは 48 GiB です。より大きなファイルをストリームアップロードする場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを増やすと、"-P" フラグで表示される進行状況の統計の正確性が低下します。ロークライオンは、AWS SDK によってチャンクが送信されたと判断されると、
      チャンクを送信したとみなしていますが、実際にはまだアップロード中の場合があります。チャンクサイズが大きいほど、
      AWS SDK のバッファも大きくなり、真実とは異なる進行状況が報告されます。

   --max-upload-parts
      マルチパートアップロードの最大部分の数。
      
      このオプションは、マルチパートアップロードを行う際の最大部分の数を定義します。
      
      AWS S3 の 10,000 チャンクの仕様をサポートしていない場合に有用です。
      
      ロークライオンは、既知のサイズの大きなファイルをアップロードする場合、チャンクサイズを増やすことでこのチャンク数の制限以下になるように自動的にチャンクサイズを増やします。

   --copy-cutoff
      パーティションごとのコピーに切り替えるためのカットオフ。
      
      サーバサイドでコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクにコピーされます。
      
      最小値は 0 で、最大値は 5 GiB です。

   --disable-checksum
      オブジェクトのメタデータに MD5 チェックサムを保存しません。
      
      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加します。これにより、データの整合性チェックが行われますが、
      大きなファイルのアップロードが開始されるまでに長時間待機する場合があります。

   --shared-credentials-file
      共有クレデンシャルファイルへのパス。
      
      env_auth が true の場合、rclone は共有クレデンシャルファイルを使用できます。
      
      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を探します。環境変数の値が空の場合は、現在のユーザのホームディレクトリがデフォルト値になります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有クレデンシャルファイルで使用するプロファイル。
      
      env_auth が true の場合、rclone は共有クレデンシャルファイルを使用します。この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、「AWS_PROFILE」環境変数または「default」がデフォルト値になります。

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同じファイルのチャンクの数を同時にアップロードします。
      
      高速リンクを介して大量の大容量ファイルをアップロードしており、これらのアップロードが帯域幅を完全に利用していない場合は、この数を増やすことで転送速度を向上させることができます。

   --force-path-style
      true の場合はパススタイルアクセスを使用し、false の場合は仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rclone はパススタイルアクセスを使用します。
      false の場合、rclone は仮想パススタイルを使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

   --v2-auth
      true の場合は v2 認証を使用します。
      
      false（デフォルト）の場合、rclone は v4 認証を使用します。セットされた場合は v2 認証を使用します。
      
      v4 シグネチャが機能しない場合にのみ使用してください。たとえば、Jewel/v10 CEPH 以前のバージョンで使用する場合などです。

   --list-chunk
      リストのチャンクサイズ（各 ListObject S3 リクエストの応答リストのサイズ）。
      
      このオプションは、AWS S3 仕様の "MaxKeys"、"max-items"、または "page-size" としても知られています。
      ほとんどのサービスは、1000 のオブジェクトをリクエストしても、応答リストを切り捨てます。
      AWS S3 ではこれがグローバルな最大値であり、変更することはできません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Ceph では、"rgw list buckets max chunk" オプションでこれを増やすことができます。

   --list-version
      使用する ListObjects のバージョン: 1、2、または 0（自動）。
      
      S3 の初版では、バケット内のオブジェクトを列挙するための ListObjects 呼び出しが提供されていました。
      
      しかし、2016 年 5 月に ListObjectsV2 の呼び出しが導入されました。これははるかに高性能であり、可能な限り使用する必要があります。
      
      デフォルトの値 0 の場合、rclone はプロバイダによって設定されたリストオブジェクトのメソッドを予想し、呼び出します。予想が誤っている場合は、ここで手動で設定することもあります。

   --list-url-encode
      リストを URL エンコードするかどうか: true/false/unset
      
      いくつかのプロバイダは、リストを URL エンコードすることをサポートしており、ファイル名で制御文字を使用する場合は、これが利用可能な場合、より信頼性が高いです。これが未設定（デフォルト）の場合、
      rclone は、プロバイダの設定に従って適用するものを選択しますが、ここで rclone の選択を上書きすることもできます。

   --no-check-bucket
      バケットの存在を確認せず、または作成しようとしないようにします。
      
      バケットがすでに存在する場合、rclone のトランザクション数を最小限に抑える必要がある場合に便利です。
      
      バケット作成の権限がない場合にも必要になることがあります。v1.52.0 より前ではバグにより、これは結果なしに渡されます。

   --no-head
      アップロード済みオブジェクトの確認のために HEAD を行わないようにします。
      
      rclone が PUT 後に 200 OK メッセージを受信した場合、正しくアップロードされたと仮定します。
      
      特に、以下を仮定します。
      
      - アップロード時のメタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロード時と同じであったこと
      - サイズがアップロード時と同じであったこと
      
      単一パートの PUT の応答から以下の項目を読み取ります。
      
      - MD5SUM
      - アップロード日
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズの不明なソースオブジェクトがアップロードされた場合、rclone は HEAD リクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まります。特にサイズが正しくない場合のチャンスが増えるため、通常の操作では推奨されません。
      実際には、このフラグを使用しても、アップロードの失敗が検出されない可能性は非常に小さいです。

   --no-head-object
      オブジェクトを取得する前に HEAD を行いません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。
      
      追加バッファが必要なアップロード (マルチパートなど) は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。
      
      現在、s3（具体的には minio）バックエンドでは、http2 の問題が解決されていません。s3 バックエンドではデフォルトで HTTP/2 が有効になっていますが、
      ここで無効にすることができます。問題が解決された際には、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロードするためのカスタムエンドポイント。
      通常、AWS S3 はクラウドフロント CDN URL に設定されます。これにより、
      クラウドフロントネットワークを介してデータをダウンロードすると、AWS S3 からのエグレスが安価になります。

   --use-multipart-etag
      検証のためにマルチパートアップロードで ETag を使用するかどうか
      
      これは true、false、またはデフォルト値（プロバイダの設定）を使用する必要があります。

   --use-presigned-request
      シングルパートのアップロードに署名付きリクエストまたは PutObject を使用するかどうか
      
      false の場合、rclone は AWS SDK の PutObject を使用してオブジェクトをアップロードします。
      
      rclone のバージョン 1.59 より前では、署名付きリクエストを使用して単一のパートオブジェクトをアップロードし、
      このフラグを true に設定すると、その機能が再度有効になります。これは、特殊な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、その他の時間の経過を示す "100d" や "1h" のような形式です。
      
      この機能を使用すると、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。
      
      有効な形式については、[time オプションドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定すると、gzip エンコードされたオブジェクトを展開します。
      
      S3 に "Content-Encoding: gzip" を設定してオブジェクトをアップロードすることが可能です。通常、rclone はこれらのファイルを圧縮オブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rclone は受信した "Content-Encoding: gzip" でこれらのファイルを展開します。これにより、
      rclone はサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。

   --might-gzip
      バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされる際にそれらを変更しません。"Content-Encoding: gzip" でアップロードされていないオブジェクトには
      ダウンロード時に設定されないため、この設定は有効ではありません（Cloudflare など）。
      
      これを設定して rclone が Content-Encoding: gzip とチャンク転送エンコーディングでオブジェクトをダウンロードすると、rclone はオブジェクトをリアルタイムで展開します。
      
      これが未設定（デフォルト）の場合、rclone はプロバイダの設定に従って適用するものを選択しますが、ここで rclone の選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制する


OPTIONS:
   --access-key-id value      AWS アクセスキー ID。[$ACCESS_KEY_ID]
   --endpoint value           Storj ゲートウェイのエンドポイント。[$ENDPOINT]
   --env-auth                 実行時に AWS 認証情報（環境変数または環境の EC2 / ECS メタデータ）を使用します。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --secret-access-key value  AWS シークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケット作成時に使用するキャニスド ACL。[$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクのサイズ。(default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              パーティションごとのコピーに切り替えるためのカットオフ。(default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定すると、gzip エンコードされたオブジェクトを展開します。(default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータに MD5 チェックサムを保存しません。(default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドでの http2 の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードするためのカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。(default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合はパススタイルアクセスを使用し、false の場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各 ListObject S3 リクエストの応答リストのサイズ）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストを URL エンコードするかどうか: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1、2、または 0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大部分の数。(default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールで mmap バッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しようとしないようにします。(default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロード済みオブジェクトの確認のために HEAD を行わないようにします。 (default: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前に HEAD を行いません。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有クレデンシャルファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWS セッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有クレデンシャルファイルへのパス。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。(default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのサイズカットオフ。(default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       検証のためにマルチパートアップロードで ETag を使用するかどうか(default: "unset")。[$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名付きリクエストまたは PutObject を使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true の場合は v2 認証を使用します。(default: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。(default: false) [$VERSIONS]

```
{% endcode %}