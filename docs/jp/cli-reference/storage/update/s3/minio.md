# Minio オブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 minio - Minio オブジェクトストレージ

USAGE:
   singularity storage update s3 minio [command options] <name|id>

DESCRIPTION:
   --env-auth
      ランタイムから AWS 認証情報を取得します（環境変数またはランタイムの認証情報がない場合は EC2/ECS メタデータから取得）。
      
      access_key_id と secret_access_key が空白の場合にのみ適用されます。

      例:
         | false | 次のステップで AWS 認証情報を入力します。
         | true  | 環境（環境変数または IAM）から AWS 認証情報を取得します。

   --access-key-id
      AWS アクセスキー ID。
      
      匿名アクセスまたは実行時認証情報の場合は空白のままにします。

   --secret-access-key
      AWS シークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時認証情報の場合は空白のままにします。

   --region
      接続するリージョンを指定します。
      
      S3 のクローンを使用しており、リージョンが存在しない場合は空白のままにします。

      例:
         | <unset>            | わからない場合はこれを使用します。
         |                    | v4 署名と空のリージョンを使用します。
         | other-v2-signature | v4 署名が機能しない場合にのみ使用します。
         |                    | たとえば、Jewel/v10 CEPH 以前のバージョン。

   --endpoint
      S3 API のエンドポイントです。
      
      S3 のクローンを使用する場合は必須です。

   --location-constraint
      リージョンと一致するように設定する場所制約です。
      
      わからない場合は空白のままにします。バケットの作成時にのみ使用されます。

   --acl
      バケットを作成したり、オブジェクトを保存したり、コピーしたりするときに使用する事前設定の ACL です。
      
      この ACL はオブジェクトの作成にも使用され、bucket_acl が設定されていない場合も使用されます。
      
      詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3 はソースから ACL をコピーするのではなく、新しい ACL を書き込むため、この ACL はサーバーサイドでオブジェクトをコピーする際に適用されます。
      
      acl が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトの ACL（プライベート）が使用されます。
      

   --bucket-acl
      バケットを作成する際に使用される事前設定の ACL です。
      
      詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      この ACL はバケットを作成するときにのみ適用されます。設定されていない場合は、"acl" が代わりに使用されます。
      
      "acl" と "bucket_acl" が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトの ACL（プライベート）が使用されます。
      

      例:
         | private            | オーナーは FULL_CONTROL 権限を持ちます。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーは FULL_CONTROL 権限を持ちます。
         |                    | AllUsers グループは READ アクセス権限を持ちます。
         | public-read-write  | オーナーは FULL_CONTROL 権限を持ちます。
         |                    | AllUsers グループは READ および WRITE アクセス権限を持ちます。
         |                    | バケットでこの権限を付与することは一般に推奨されていません。
         | authenticated-read | オーナーは FULL_CONTROL 権限を持ちます。
         |                    | AuthenticatedUsers グループは READ アクセス権限を持ちます。

   --server-side-encryption
      S3 にこのオブジェクトを保存する際に使用するサーバーサイドの暗号化アルゴリズムです。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C を使用する場合、このオブジェクトを S3 に保存する際に使用するサーバー側の暗号化アルゴリズムです。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID を使用する場合、鍵の ARN を指定する必要があります。

      例:
         | <unset>                 | なし
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C を使用する場合、データの暗号化/複合に使用するシークレット暗号化キーを指定できます。
      
      --sse-customer-key-base64 を使用することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-base64
      SSE-C を使用する場合、データの暗号化/複合に使用するシークレット暗号化キーを Base64 形式でエンコードして指定する必要があります。
      
      --sse-customer-key を使用することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-md5
      SSE-C を使用する場合、シークレット暗号化キーの MD5 チェックサムを指定できます（オプション）。
      
      空白のままにすると、s3_customer_key から自動的に計算されます。
      

      例:
         | <unset> | なし

   --upload-cutoff
      チャンクアップロードに切り替えるための閾値です。
      
      これを超えるサイズのファイルは、chunk_size 単位でアップロードされます。
      最小値は 0、最大値は 5 GiB です。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoff より大きなサイズのファイルや、サイズが不明なファイル（"rclone rcat" からのアップロード、または "rclone mount" や Google フォトや Google ドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      注意: "--s3-upload-concurrency" byte のチャンクは転送ごとにメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送して十分なメモリを持っている場合、チャンクサイズを増やすことで転送速度を向上させることができます。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする場合、10,000 個のチャンクの制限を下回るように自動的にチャンクサイズを増加させます。
      
      サイズが不明なファイルは、設定された chunk_size でアップロードされます。デフォルトのチャンクサイズは 5 MiB であり、最大 10,000 チャンクまで存在できるため、デフォルトではストリームアップロードできるファイルの最大サイズは 48 GiB です。より大きなファイルをストリームアップロードしたい場合は、chunk_size のサイズを大きくする必要があります。
      
      チャンクサイズを増やすと、"-P" フラグで表示される進行状況の統計の正確性が低下します。Rclone は、AWS SDK によってバッファリングされたチャンクを送信したと見なし、実際にはまだアップロード中の場合でもそのチャンクを送信したと見なすためです。
      チャンクサイズが大きいほど、AWS SDK のバッファも大きくなるため、進行状況の報告が真実からずれることがあります。
      

   --max-upload-parts
      マルチパートアップロードでアップロードするパートの最大数です。
      
      このオプションは、マルチパートアップロードを実行する際に使用するマルチパートの最大数を定義します。
      
      サービスが AWS S3 の 10,000 チャンクの仕様をサポートしていない場合に役立ちます。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする場合、このチャンク数の制限を下回るように自動的にチャンクサイズを増加させます。
      

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値です。
      
      サーバーサイドでコピーする必要があるこれより大きなサイズのファイルは、このサイズのチャンクでコピーされます。
      
      最小値は 0、最大値は 5 GiB です。

   --disable-checksum
      オブジェクトのメタデータとともに MD5 チェックサムを保存しません。
      
      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算して、オブジェクトのメタデータに追加するため、大きなファイルをアップロードする場合には長い遅延が発生する可能性があります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rclone は"AWS_SHARED_CREDENTIALS_FILE" 環境変数を検索します。環境変数が空の場合、現在のユーザーのホームディレクトリがデフォルト値として使用されます。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数は、そのファイル内で使用されるプロファイルを制御します。
      
      空の場合、環境変数 "AWS_PROFILE" または "default" が設定されていない場合にデフォルト値として使用されます。
      

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数です。
      
      同じファイルの複数のチャンクを同時にアップロードします。
      
      高速リンクで少数の大きなファイルをアップロードして、これらのアップロードが帯域幅を十分に利用しない場合、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      true の場合、パス形式でアクセスします。false の場合、仮想ホスト形式でアクセスします。
      
      true（デフォルト）の場合、rclone はパス形式アクセスを使用します。
      false の場合、rclone は仮想パス形式を使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、または Tencent COS など）は、これを false に設定する必要があります。rclone はプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      true の場合、v2 認証を使用します。
      
      false（デフォルト）の場合、rclone は v4 認証を使用します。
      設定されている場合、rclone は v2 認証を使用します。
      
      v4 署名が機能しない場合にのみ使用します。たとえば、Jewel/v10 CEPH 以前のバージョンで使用します。

   --list-chunk
      リストのチャンクサイズです（各 ListObject S3 リクエストごとに応答リストのサイズ）。
      
      このオプションは、AWS S3 の仕様では「MaxKeys」、「max-items」、または「page-size」としても知られています。
      ほとんどのサービスは、リクエストされた数よりも多い場合でも、レスポンスリストを 1000 オブジェクトに切り詰めます。
      AWS S3 ではこれはグローバルな最大値であり、変更することはできません。詳細については [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) を参照してください。
      Ceph では、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用する ListObjects のバージョン: 1、2、または 0（自動）。
      
      S3 が最初にリリースされたとき、バケット内のオブジェクトを列挙するための ListObjects 呼び出しばかりが提供されていました。
      
      しかし、2016 年 5 月に ListObjectsV2 呼び出しが導入されました。これははるかに高性能であり、可能な場合は使用する必要があります。
      
      デフォルトの設定 0 の場合、rclone はプロバイダで設定に応じてどのリストオブジェクトメソッドを呼び出すか推測します。推測が誤っている場合、ここで手動で設定できます。
      

   --list-url-encode
      リストを URL エンコードするかどうか: true/false/unset
      
      一部のプロバイダは、ファイル名で制御文字を使用する場合に URL エンコードリストをサポートしています。利用可能な場合、これはファイルの制御文字を使用すると信頼性が向上します。設定が unset の場合（デフォルト）、rclone はプロバイダの設定に従って適用するものを選択しますが、ここで rclone の選択を上書きできます。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しないように設定します。
      
      バケットが既に存在する場合、rclone が実行するトランザクションの回数を最小限に抑えるために便利です。
      
      また、使用するユーザーにバケット作成の権限がない場合にも必要です。v1.52.0 以前では、このバグのために無音で渡されました。
      

   --no-head
      アップロードしたオブジェクトの HEAD を行って整合性をチェックしません。
      
      rclone が PUT 後に 200 OK メッセージを受信した場合、正常にアップロードされたものと見なします。
      
      特に次の項目を前提とします。
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプなど）がアップロードと同じであること
      - サイズがアップロードと同じであること
      
      単一部品の PUT の場合、次の項目をレスポンスから読み取ります。
      
      - MD5SUM
      - アップロード日
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明なソースオブジェクトがアップロードされる場合、rclone は **HEAD リクエストを行います。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が高くなるため、通常の操作では推奨されません。実際には、このフラグを設定してもアップロードの失敗が検出される可能性は非常に低いです。
      

   --no-head-object
      オブジェクトの取得前に HEAD を行いません。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがどのくらいの頻度でフラッシュされるかを指定します。
      
      追加のバッファが必要なアップロード（例：マルチパート）では、メモリプールが割り当てに使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      インターナルメモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。
      
      現在、s3（特に Minio）バックエンドと HTTP/2 の問題が未解決です。S3 バックエンドの HTTP/2 はデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      通常は CloudFront CDN の URL に設定されます。AWS S3 では、CloudFront ネットワークを介したデータのダウンロードに低コストのエグレスを提供しています。

   --use-multipart-etag
      マルチパートアップロードで ETag を使用して検証するかどうか
      
      true、false、またはデフォルト（unset）を設定して使用するプロバイダのデフォルト値を使用します。
      

   --use-presigned-request
      単一部品のアップロードに予め署名済みリクエストまたは PutObject を使用するかどうか
      
      false の場合、rclone は AWS SDK の PutObject を使用してオブジェクトをアップロードします。
      
      rclone < 1.59 のバージョンでは、単一部品のオブジェクトをアップロードするために予め署名済みリクエストを使用し、
      このフラグを true に設定すると、その機能が再度有効になります。これは例外的な事例またはテストを除いては必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時点でのファイルのバージョンを表示します。
      
      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはその以前の期間、「100d」や「1h」などです。
      
      このオプションを使用する場合、ファイルへの書き込み操作は許可されないため、ファイルをアップロードまたは削除できません。
      
      有効な形式については、[time オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      埋め込まれた gzip エンコードされたオブジェクトを解凍します。
      
      S3 に "Content-Encoding: gzip" が設定された状態でオブジェクトをアップロードすることができます。通常、rclone はこれらのファイルを圧縮オブジェクトとしてダウンロードします。
      
      このフラグを設定すると、rclone は受信した際に "Content-Encoding: gzip" でこれらのファイルを解凍します。つまり、rclone はサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトを圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。`Content-Encoding: gzip` でアップロードされなかった場合、ダウンロード時に設定されません。
      
      ただし、一部のプロバイダーは、`Content-Encoding: gzip` でアップロードされていないオブジェクトを圧縮する場合があります（例：Cloudflare）。
      
      これにより、次のようなエラーが発生します。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      もしこのフラグを設定し、rclone が `Content-Encoding: gzip` が設定され、チャンク化された転送エンコードを使用してオブジェクトをダウンロードした場合、rclone はオブジェクトを逐次解凍します。
      
      unset に設定されている場合（デフォルト）、rclone はプロバイダの設定に従って適用するものを選択しますが、ここで rclone の選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value           AWS アクセスキー ID。 [$ACCESS_KEY_ID]
   --acl value                     バケットを作成したり、オブジェクトを保存したり、コピーしたりするときに使用する事前設定の ACL。 [$ACL]
   --endpoint value                S3 API のエンドポイント。 [$ENDPOINT]
   --env-auth                      ランタイムから AWS 認証情報を取得します（環境変数またはランタイムの認証情報がない場合は EC2/ECS メタデータから取得）。 (default: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示
   --location-constraint value     リージョンと一致するように設定する場所制約。 [$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。 [$REGION]
   --secret-access-key value       AWS シークレットアクセスキー（パースワード）。 [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3 にこのオブジェクトを保存する際に使用するサーバーサイドの暗号化アルゴリズム。 [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID を使用する場合、鍵の ARN を指定する必要があります。 [$SSE_KMS_KEY_ID]

   バックエンドの詳細

   --bucket-acl value               バケットを作成する際に使用される事前設定の ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     埋め込まれた gzip エンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータとともに MD5 チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドの HTTP/2 の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合、パス形式でアクセスします。false の場合、仮想ホスト形式でアクセスします。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストを URL エンコードするかどうか: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1、2、または 0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでアップロードするパートの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがどのくらいの頻度でフラッシュされるか。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           インターナルメモリプールで mmap バッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しないように設定します。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトの HEAD を行って整合性をチェックしません。 (default: false) [$NO_HEAD]
   --no-head-object                 オブジェクトの取得前に HEAD を行いません。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWS セッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C を使用する場合、このオブジェクトを S3 に保存する際に使用するサーバー側の暗号化アルゴリズム。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C を使用する場合、データの暗号化/複合に使用するシークレット暗号化キー。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C を使用する場合、データの暗号化/複合に使用するシークレット暗号化キーを Base64 形式でエンコードします。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C を使用する場合、シークレット暗号化キーの MD5 チェックサムが計算されずに指定できます（オプション）。 [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードで ETag を使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一部品のアップロードに予め署名済みリクエストまたは PutObject を使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true の場合、v2 認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルのバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

```
{% endcode %}