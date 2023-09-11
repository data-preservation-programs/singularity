# Cephオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 ceph - Ceph Object Storage

USAGE:
   singularity storage update s3 ceph [コマンドオプション] <名前|ID>

DESCRIPTION:
   --env-auth
      実行時にAWS認証情報を取得します（環境変数または環境によらずエンドポイントにアクセスする場合）。
      
      access_key_idとsecret_access_keyが空である場合にのみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境からAWS認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにします。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにします。

   --region
      接続するリージョン。
      
      S3クローンを使用している場合でリージョンを持っていない場合は空白のままにします。

      例:
         | <未設定>            | 分からない場合はこのままにします。
         |                    | v4シグネチャと空白のリージョンを使用します。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                    | 例：Jewel/v10 CEPHよりも前。

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用している場合に必要です。

   --location-constraint
      リージョンに合わせて設定する必要がある場所拘束条件。
      
      分からない場合は空白のままにします。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成やオブジェクトの保存またはコピー時に使用される共有ACL。
      
      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。
      
      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3ではサーバーサイドでオブジェクトをコピーする際にACLはソースからコピーされず、代わりに新しいACLが書き込まれます。
      
      aclが空の文字列の場合、X-Amz-Acl:ヘッダは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用される共有ACL。
      
      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucket_aclが設定されていない場合にのみ、バケットの作成時に使用されます。
      
      aclとbucket_aclが空の文字列の場合は、X-Amz-Acl:ヘッダは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループは読み取りアクセス権を取得します。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループは読み取りと書き込みのアクセス権を取得します。
         |                    | バケットでこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループは読み取りアクセス権を取得します。

   --server-side-encryption
      S3でこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <未設定> | サーバーサイドの暗号化なし
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-Cを使用する場合、S3でこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <未設定> | サーバーサイドの暗号化なし
         | AES256  | AES256

   --sse-kms-key-id
      KMS IDを使用する場合、キーのARNを指定する必要があります。

      例:
         | <未設定>                 | キーなし
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化/復号に使用する秘密鍵を指定することができます。
      
      または、--sse-customer-key-base64を指定することもできます。

      例:
         | <未設定> | 秘密鍵なし

   --sse-customer-key-base64
      SSE-Cを使用する場合、データの暗号化/復号に使用する秘密鍵をBase64形式で指定する必要があります。
      
      または、--sse-customer-keyを指定することもできます。

      例:
         | <未設定> | 秘密鍵なし

   --sse-customer-key-md5
      SSE-Cを使用する場合、秘密鍵のMD5チェックサム（省略可）を指定することができます。
      
      空白の場合は、sse_customer_keyから自動的に計算されます。

      例:
         | <未設定> | チェックサムなし

   --upload-cutoff
      切り替えてチャンクアップロードするためのカットオフ。
      
      これより大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      このchunk_sizeを使用して、upload_cutoffよりも大きなファイルまたはサイズ不明のファイル
      （例: "rclone rcat"または"rclone mount"でアップロードまたはGoogleフォトまたはGoogleドキュメントでアップロードされたファイル）が
      マルチパートアップロードとしてアップロードされます。
      
      注意："--s3-upload-concurrency"はこのchunk_sizeのチャンクがトランスファごとにメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送し、十分なメモリがある場合は、
      チャンクサイズを増やすことで転送が高速化します。
      
      rcloneは、既知のサイズの大きなファイルをアップロードする場合にはチャンクサイズを自動的に増やし、
      10,000チャンクの制限を超えないようにします。
      
      サイズ不明のファイルは、構成されたchunk_sizeでアップロードされます。
      デフォルトのチャンクサイズが5 MiBで10,000チャンクの最大サイズであるため、
      ストリームアップロードできるファイルの最大サイズは48 GiBです。
      より大きなファイルをストリームアップロードする場合は、チャンクサイズを増やす必要があります。
      
      チャンクサイズを増やすと、プログレス統計の精度が低下します。
      rcloneはチャンクごとにAWS SDKによってバッファされると、送信済みの部分として処理しますが、
      実際にはまだアップロード中の場合もあります。
      チャンクサイズが大きいほど、AWS SDKのバッファも大きくなり、進捗報告が真実からずれることがあります。

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数。
      
      マルチパートアップロードを行う際に使用するパーツの最大数を定義するオプションです。
      
      サービスが10,000パーツ以上のAWS S3仕様をサポートしていない場合に役立ちます。
      
      rcloneは既知のサイズの大きなファイルをアップロードする場合、
      このパーツ数の制限を超えないようにチャンクサイズを自動的に増やします。
      
   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。
      
      サーバーサイドでコピーする必要があるこれより大きなファイルは、
      このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、
      オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するのに長い遅延が発生する場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を検索します。
      環境値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は環境変数"AWS_PROFILE"または"default"が設定されていない場合にデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行数。
      
      同じファイルのチャンクが並行してアップロードされます。
      
      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用しない場合、
      並行性を増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用し、
      falseの場合は仮想パススタイルを使用します。
      詳細については[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）は、これをfalseに設定する必要があります。
      rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。
      設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャが機能しない場合のみ使用してください。例：Jewel/v10 CEPH。

   --list-chunk
      リストチャンクのサイズ（各ListObject S3リクエストのレスポンスリスト）。
      
      このオプションはAWS S3仕様の「MaxKeys」、「max-items」、または「page-size」とも呼ばれます。
      ほとんどのサービスでは、リクエストされたオブジェクトのリストを1000オブジェクトに切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、この値を「rgw list buckets max chunk」オプションで増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にローンチされた時、バケット内のオブジェクトを列挙するためにListObjects呼び出しだけが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高速であり、可能な場合は使用する必要があります。
      
      デフォルトの0に設定されている場合、rcloneはプロバイダによってリストオブジェクトメソッドの呼び出しを推測します。
      誤った推測をした場合は、ここで手動で設定することもできます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を使用する場合、URLエンコードリストをサポートしており、このオプションを選択すると確実性が高まります。
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることもできます。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しない場合は設定します。

      バケットが既に存在することを事前に知っている場合、rcloneが行うトランザクションの数を最小限に抑えるために便利です。
      
      バケットの作成権限がない場合は、必要になる場合があります。
      バージョン1.52.0より前では、このバグのために無言で渡されますが。
      

   --no-head
      アップロードしたオブジェクトをHEADして整合性を確認しない場合は設定します。
      
      rcloneはアップロード中にPUT後に200 OKメッセージを受信した場合、正しくアップロードされたと想定します。

      特に、次の項目を以下のように想定します。

      - MB5SUM
      - アップロード日

      確定的なファイルのサイズが不明の場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が増えます。
      
      特に、正しいサイズではない場合、このフラグは通常の操作ではお勧めしません。
      
      実際には、このフラグを使用しても、アップロードの失敗が検出される可能性は非常に少ないです。

   --no-head-object
      オブジェクトを取得する前にHEADを実行しない場合は設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      本来は追加バッファ（例: マルチパート）が必要なアップロードは、メモリプールが割り当て用に使用されます。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドのhttp2の使用を無効にします。
      
      s3（具体的にはminio）バックエンドとHTTP/2の現在の未解決の問題があります。
      S3バックエンドのHTTP/2はデフォルトで有効になっていますが、ここで無効にできます。
      問題が解決されたときには、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3が提供するよりも安価なユーザー間ネットワークを介してデータをダウンロードするために、CloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      これはtrue、false、またはデフォルトのプロバイダーに使用される値を指定する必要があります。

   --use-presigned-request
      マルチパートアップロード以外のシングルパートアップロードに署名付きリクエストまたはPutObjectを使用するかどうか

      これがfalseの場合、rcloneはオブジェクトをアップロードするためにAWS SDKからPutObjectを使用します。

      rcloneのバージョン1.59より前では、署名付きのリクエストを使用して単一パーツオブジェクトをアップロードし、
      このフラグをtrueに設定すると、その機能が再有効化されます。
      
      逆の特別な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時点でのファイルバージョンを表示する。
      
      パラメータは日付（"2006-01-02"）、日時（"2006-01-02 15:04:05"）またはそれより前の期間（"100d"または"1h"）である必要があります。
      
      このオプションの使用中は、ファイルの書き込み操作は許可されません。
      したがって、ファイルのアップロードや削除はできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      gzipでエンコードされたオブジェクトを解凍する場合は、これを設定します。
      
      S3に「Content-Encoding: gzip」が設定されたオブジェクトをアップロードすることもできます。
      通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding: gzip」として受け取ったときに解凍します。
      これにより、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルの内容が解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzipする可能性がある場合は、これを設定します。
      
      通常のプロバイダは、オブジェクトがダウンロードされたときに変更しないでしょう。
      `Content-Encoding: gzip`が設定されていない場合、ダウンロード時にはそれも設定されません。
      
      ただし、いくつかのプロバイダは、それらが`Content-Encoding: gzip`でアップロードされていなくてもオブジェクトをgzip圧縮する場合があります（例: Cloudflare）。

      これを設定してrcloneが`Content-Encoding: gzip`が設定されたチャンク化転送エンコードを受信すると、
      rcloneはオブジェクトを逐次解凍します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることもできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value           AWSアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                     バケットの作成やオブジェクトの保存またはコピー時に使用される共有ACL。 [$ACL]
   --endpoint value                S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                      実行時にAWS認証情報を取得します（環境変数またはエンドポイントへのアクセスに依存しません）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示します
   --location-constraint value     リージョンに合わせて設定する必要がある場所拘束条件。 [$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。 [$REGION]
   --secret-access-key value       AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3でこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。 [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS IDを使用する場合、キーのARNを指定する必要があります。 [$SSE_KMS_KEY_ID]

   高度な設定

   --bucket-acl value               バケットの作成時に使用される共有ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzipでエンコードされたオブジェクトを解凍する場合は、これを設定します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズ（各ListObject S3リクエストのレスポンスリスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動） (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipする可能性がある場合は、これを設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しない場合は設定します。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトをHEADして整合性を確認しない場合は設定します。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADを実行しない場合は設定します。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-Cを使用する場合、S3でこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データの暗号化/復号に使用する秘密鍵を指定することができます。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-Cを使用する場合、データの暗号化/復号に使用する秘密鍵をBase64形式で指定する必要があります。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-Cを使用する場合、秘密鍵のMD5チェックサム（省略可）を指定することができます。 [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       マルチパートアップロードの並行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            切り替えてチャンクアップロードするためのカットオフ。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          マルチパートアップロード以外のシングルパートアップロードに署名付きリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示する。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (デフォルト: false) [$VERSIONS]

```
{% endcode %}