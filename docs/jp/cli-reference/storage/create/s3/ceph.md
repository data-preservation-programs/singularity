# Cephオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 ceph - Cephオブジェクトストレージ

USAGE:
   singularity storage create s3 ceph [command options] [arguments...]

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSのメタデータ）。
      
      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | AWSの認証情報を次のステップで入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたはランタイムの認証情報の場合は空にしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたはランタイムの認証情報の場合は空にしてください。

   --region
      接続するリージョン。
      
      S3クローンを使用している場合でリージョンが必要ない場合は空にしてください。

      例:
         | <unset>            | 不確かな場合はこれを使用します。
         |                    | v4シグネチャと空のリージョンが使用されます。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                    | 例：旧バージョンのCEPH。

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用している場合に必要です。

   --location-constraint
      リージョンに一致するロケーション制約。
      
      よくわからない場合は空にしてください。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成およびオブジェクトの保存やコピー時に使用されるCanned ACL。
      
      このACLはオブジェクトの作成時およびbucket_aclが設定されていない場合にも使用されます。
      
      詳細は以下を参照してください：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      S3はソースからACLをコピーせず、新しいACLを書き込むため、サーバーサイドでオブジェクトをコピーする際にこのACLが適用されます。
      
      aclが空の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用されるCanned ACL。
      
      詳細は以下を参照してください：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      このACLはバケットの作成時にのみ適用されます。設定されていない場合は、aclが代わりに使用されます。
      
      aclとbucket_aclが空の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーにFULL_CONTROLが付与されます。
         |                    | 他のユーザーにアクセス権限はありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLが付与されます。
         |                    | AllUsersグループに読み取りアクセスが付与されます。
         | public-read-write  | オーナーにFULL_CONTROLが付与されます。
         |                    | AllUsersグループに読み取りと書き込みのアクセスが付与されます。
         |                    | バケットに対してこれを許可することは一般的にお勧めしません。
         | authenticated-read | オーナーにFULL_CONTROLが付与されます。
         |                    | AuthenticatedUsersグループに読み取りアクセスが付与されます。

   --server-side-encryption
      S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-Cを使用する場合、S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-kms-key-id
      KMS IDを使用する場合は、キーのARNを指定する必要があります。

      例:
         | <unset>                 | なし
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-Cを使用する場合、データを暗号化/復号化するために使用される秘密の暗号化キーを指定できます。
      
      代わりに、--sse-customer-key-base64で指定することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-base64
      SSE-Cを使用する場合、データを暗号化/復号化するために使用される秘密の暗号化キーをBase64形式で指定できます。
      
      代わりに、--sse-customer-keyを指定することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-md5
      SSE-Cを使用する場合は、秘密の暗号化キーのMD5チェックサムを指定できます（任意）。
      
      空の場合、sse_customer_keyから自動的に計算されます。
      

      例:
         | <unset> | なし

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフサイズ。
      
      これを超えるサイズのファイルは、chunk_sizeごとにチャンクアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（"rclone rcat"でアップロードされたものや"rclone mount"やGoogle PhotosやGoogle Docsでアップロードされたものなど）は、このチャンクサイズを使用してマルチパートのアップロードとしてアップロードされます。
      
      "--s3-upload-concurrency"個のこのサイズのチャンクが、転送ごとにメモリ内にバッファリングされます。
      
      高速リンク上で大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送が高速化します。
      
      Rcloneは、10,000のチャンク制限を超えないようにするため、既知のサイズの大きなファイルのアップロード時に自動的にチャンクサイズを増やします。
      
      サイズが不明なファイルは、構成されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大10,000のチャンクまで存在するため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の統計の正確性が低下します。Rcloneは、AWS SDKがバッファリングされたチャンクを送信したときにチャンクを送信済みと見なし、実際にはまだアップロード中かもしれないためです。
      チャンクサイズが大きいほど、AWS SDKのバッファーサイズが大きくなり、進行状況報告が真実から外れる可能性があります。
      

   --max-upload-parts
      マルチパートのアップロードで使用するパートの最大数を定義します。
      
      このオプションは、マルチパートのアップロード時に使用するマルチパートチャンクの最大数を定義します。
      
      サービスが10,000チャンクのAWS S3仕様をサポートしていない場合に役立ちます。
      
      Rcloneは、既知のサイズの大きなファイルのアップロード時に自動的にチャンクサイズを増やして、このチャンク数の制限を下回るようにします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフサイズ。
      
      サーバーサイドコピーする必要のあるこれを超えるサイズのファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しない。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまでに長時間待たされることがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は、環境変数"AWS_PROFILE"または"デフォルト"が設定されていない場合にデフォルトになります。
      

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートのアップロードの並行数。
      
      同じファイルのチャンクが並行してアップロードされる数です。
      
      帯域幅を十分に使用していない高速リンク上で少数の大きなファイルをアップロードしている場合、これを増やすと転送が高速化する可能性があります。

   --force-path-style
      trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      これがtrueの場合（デフォルト）、rcloneはパススタイルのアクセスを使用し、falseの場合は仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）は、これがfalseに設定されなければなりません。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      これがfalseの場合（デフォルト）、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャが機能しない場合にのみ使用してください。例：Jewel/v10以前のCEPH。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストに対する応答リスト）のサイズ。
      
      このオプションは、AWS S3仕様のMaxKeys、max-items、またはpage-sizeとしても知られています。
      多くのサービスはリクエストよりも多数のリストを使用する場合でも、応答リストを1000個に切り詰めます。
      AWS S3では、これはグローバルな最大値であり、変更することはできません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、 "rgw list buckets max chunk"オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にリリースされた当初、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されていました。
      
      しかし、2016年5月、ListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを提供し、可能であれば使用する必要があります。
      
      デフォルトの設定（0）の場合、rcloneはプロバイダがどのリストオブジェクトのメソッドを呼び出すかを推測します。推測が間違っている場合は、ここで手動で設定できます。
      

   --list-url-encode
      リストのURLエンコードを行うかどうか: true/false/unset
      
      一部のプロバイダは、リストをURLエンコードできる場合があります。利用可能な場合、制御文字をファイル名に使用する場合にこれがより信頼性があります。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って選択しますが、ここでrcloneの選択を上書きできます。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しようとしません。
      
      バケットがすでに存在することを知っている場合に、rcloneが実行するトランザクションの数を最小限に抑えるために便利です。
      
      また、使用しているユーザーにバケット作成権限がない場合にも必要です。v1.52.0より前では、バグのため、これは無視されていました。
      

   --no-head
      アップロードしたオブジェクトの整合性を確認するためにHEADを行いません。
      
      rcloneがPUT後に200 OKメッセージを受信した場合、正しくアップロードされたと仮定します。
      
      特に次のものの場合、次の内容を仮定します。
      
      - アップロード時のメタデータ、モディファイド日時、ストレージクラス、コンテンツタイプはアップロードしたものと同じ
      - サイズはアップロードしたものと同じ
      
      シングルパートPUTの応答から次の項目を読み取ります。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明のソースオブジェクトがアップロードされると、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗が検出される可能性が増します。特に、誤ったサイズの場合ですので、通常の操作ではお勧めしません。実際には、このフラグを設定しても、アップロードの失敗が検出される可能性は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は、概要の[encodingセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファ（マルチパートなどが必要なアップロード）は、割り当て用にメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドのhttp2の使用を無効にします。
      
      現在、s3（具体的にはminio）バックエンドとHTTP/2に関する未解決の問題があります。HTTP/2はs3バックエンドのデフォルトで有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3はCloudFrontネットワークを介してダウンロードされたデータに対してより安価な出口を提供するため、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルト（プロバイダの設定）のいずれかに設定する必要があります。
      

   --use-presigned-request
      シングルパートのアップロードに署名付きリクエストを使用するか、PutObjectを使用するか指定します。
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン1.59未満では、シングルパートオブジェクトのアップロードに署名付きリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。特別な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリスティングに古いバージョンを含めます。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付、「2006-01-02」、日時「2006-01-02 15:04:05」、長い時間のための期間、例えば「100d」または「1h」である必要があります。
      
      このオプションではファイルの書き込み操作は許可されていませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。
      
      S3に「Content-Encoding: gzip」が設定された状態でオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは「Content-Encoding: gzip」として受信したファイルを展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。 `Content-Encoding: gzip`がアップロード時に設定されていない場合、ダウンロード時に設定されません。
      
      ただし、一部のプロバイダはオブジェクトをgzip圧縮する場合があります（例：Cloudflare）。
      
      これにより、次のようなエラーが発生することがあります。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、チャンク化転送エンコードが設定されたオブジェクトをダウンロードする場合、rcloneはオブジェクトをリアルタイムで展開します。
      
      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って選択した内容を適用しますが、ここでrcloneの選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定と読み込みを抑制します


OPTIONS:
   --access-key-id value           AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                     バケットの作成およびオブジェクトの保存やコピー時に使用されるCanned ACL。 [$ACL]
   --endpoint value                S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                      ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSのメタデータ）（デフォルト：false） [$ENV_AUTH]
   --help, -h                      ヘルプを表示
   --location-constraint value     リージョンに一致するロケーション制約。 [$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。 [$REGION]
   --secret-access-key value       AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。 [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS IDを使用する場合は、キーのARNを指定する必要があります。 [$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。（デフォルト： "5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフサイズ。（デフォルト： "4.656Gi"） [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。（デフォルト： false） [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しない。（デフォルト： false） [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。（デフォルト： false） [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト： "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --force-path-style               trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。（デフォルト： true） [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストに対する応答リスト）のサイズ。（デフォルト： 1000） [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードを行うかどうか：true/false/unset（デフォルト："unset"） [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。（デフォルト： 0） [$LIST_VERSION]
   --max-upload-parts value         マルチパートのアップロードで使用するパートの最大数。（デフォルト： 10000） [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。（デフォルト： "1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト： false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。（デフォルト："unset"） [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しようとしません。（デフォルト： false） [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトの整合性をチェックするためにHEADを行いません。（デフォルト： false） [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADを行わない場合に設定します。（デフォルト： false） [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み込みを抑制します（デフォルト： false） [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-Cを使用する場合、S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データを暗号化/復号化するために使用される秘密の暗号化キーを指定できます。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-Cを使用する場合、データを暗号化/復号化するために使用される秘密の暗号化キーをBase64形式で指定できます。 [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-Cを使用する場合は、秘密の暗号化キーのMD5チェックサムを指定できます（任意）。 [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       マルチパートのアップロードの並行数。（デフォルト： 4） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフサイズ。（デフォルト： "200Mi"） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか（デフォルト："unset"） [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名付きリクエストを使用するか、PutObjectを使用するか指定します（デフォルト： false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します（デフォルト： false） [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します（デフォルト："off"） [$VERSION_AT]
   --versions                       ディレクトリリスティングに古いバージョンを含めます（デフォルト： false） [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}