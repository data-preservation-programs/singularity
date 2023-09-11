# Qiniu オブジェクトストレージ（Kodo）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 qiniu - Qiniu オブジェクトストレージ（Kodo）

USAGE:
   singularity storage create s3 qiniu [command options] [arguments...]

DESCRIPTION:
   --env-auth
      ランタイムからAWS認証情報を取得します（環境変数または、env変数の場合はEC2 / ECSメタデータ）。

      access_key_idとsecret_access_keyが空白の場合のみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境（env_varsまたはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSのアクセスキーIDです。

      匿名アクセスまたはランタイム認証情報をご利用になる場合は空白にしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたはランタイム認証情報をご利用になる場合は空白にしてください。

   --region
      接続するリージョンです。

      例:
         | cn-east-1      | デフォルトのエンドポイント - 分からない場合はこのリージョンを選択してください
         |                | 中国東部リージョン1
         |                | Location constraintはcn-east-1を設定する必要があります。
         | cn-east-2      | 中国東部リージョン2
         |                | Location constraintはcn-east-2を設定する必要があります。
         | cn-north-1     | 中国北部リージョン1
         |                | Location constraintはcn-north-1を設定する必要があります。
         | cn-south-1     | 中国南部リージョン1
         |                | Location constraintはcn-south-1を設定する必要があります。
         | us-north-1     | 北アメリカリージョン
         |                | Location constraintはus-north-1を設定する必要があります。
         | ap-southeast-1 | 東南アジアリージョン1
         |                | Location constraintはap-southeast-1を設定する必要があります。
         | ap-northeast-1 | 北東アジアリージョン1
         |                | Location constraintはap-northeast-1を設定する必要があります。

   --endpoint
      Qiniuオブジェクトストレージのエンドポイントです。

      例:
         | s3-cn-east-1.qiniucs.com      | 中国東部エンドポイント1
         | s3-cn-east-2.qiniucs.com      | 中国東部エンドポイント2
         | s3-cn-north-1.qiniucs.com     | 中国北部エンドポイント1
         | s3-cn-south-1.qiniucs.com     | 中国南部エンドポイント1
         | s3-us-north-1.qiniucs.com     | 北アメリカエンドポイント1
         | s3-ap-southeast-1.qiniucs.com | 東南アジアエンドポイント1
         | s3-ap-northeast-1.qiniucs.com | 北東アジアエンドポイント1

   --location-constraint
      リージョンに一致させる必要があるロケーション制約です。

      バケットを作成するときにのみ使用します。

      例:
         | cn-east-1      | 中国東部リージョン1
         | cn-east-2      | 中国東部リージョン2
         | cn-north-1     | 中国北部リージョン1
         | cn-south-1     | 中国南部リージョン1
         | us-north-1     | 北アメリカリージョン1
         | ap-southeast-1 | 東南アジアリージョン1
         | ap-northeast-1 | 北東アジアリージョン1

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用されるプリセットACLです。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合（バケットの作成時にも）、バケットの作成にも使用されます。

      詳細については、[公式ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3では、サーバ側でオブジェクトをコピーする際にACLは元のACLはコピーされず、新しいACLが適用されることに注意してください。

      もしaclが空の文字列である場合はX-Amz-Aclヘッダーは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用されるプリセットACLです。

      詳細については、[公式ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      バケットの作成時にのみ使用されます。設定されていない場合は「acl」が代わりに使用されます。

      もし「acl」と「bucket_acl」が空の文字列である場合はX-Amz-Aclヘッダーは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーにアクセス権限はありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループへのREADアクセスが設定されます。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループにREADおよびWRITEアクセスが設定されます。
         |                    | バケットでこれを許可することは一般的にはお勧めできません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループにREADアクセスが設定されます。

   --storage-class
      Qiniuで新しいオブジェクトを保存するためのストレージクラスです。

      例:
         | STANDARD     | 標準ストレージクラス
         | LINE         | 低頻度アクセスストレージモード
         | GLACIER      | アーカイブストレージモード
         | DEEP_ARCHIVE | ディープアーカイブストレージモード

   --upload-cutoff
      チャンク化アップロードに切り替えるための閾値です。

      この閾値よりも大きなサイズのファイルはchunk_sizeのサイズに分割してアップロードされます。
      最小は0、最大は5GiBです。

   --chunk-size
      アップロード時に使用するチャンクのサイズです。

      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのアップロード、または「rclone mount」またはGoogleフォトやGoogleドキュメントからアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートのアップロードとしてアップロードされます。

      注意点として、"--s3-upload-concurrency"のチャンク（chunk_sizeのサイズ）は、転送ごとにメモリ中にバッファリングされます。

      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを増やすと転送が高速化されます。

      Rcloneは、既知のサイズの大きなファイルを10,000チャンクの制限以下に保つため、チャンクサイズを自動的に増やします。

      不明なサイズのファイルの場合は、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5MiBであり、最大で10,000のチャンクが存在できるため、デフォルトでは最大48GiBのファイルサイズのストリームアップロードが可能です。それ以上のファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      chunkサイズを大きくすることで、「-P」フラグで表示される進捗統計の精度が低下します。Rcloneは、AWS SDKによってバッファリングされたチャンクが送信された時点で、チャンクが送信されたと見なしていますが、実際にはまだアップロード中の場合があります。より大きなチャンクサイズは、AWS SDKのバッファーサイズが大きくなり、真実から逸脱した進捗レポートが発生します。

   --max-upload-parts
      マルチパートアップロードの最大パート数です。

      このオプションは、マルチパートアップロードを行う際に使用するパートの最大数を定義します。

      サービスがAWS S3の10,000パートの仕様をサポートしていない場合、これが役立ちます。

      Rcloneは、既知のサイズの大きなファイルを10,000チャンクの制限以下に保つため、チャンクサイズを自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値です。

      サーバ側でコピーする必要がある、この閾値よりも大きなサイズのファイルはこのサイズのチャンクでコピーされます。
      最小は0、最大は5GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しないでください。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には時間がかかる場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。

      変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」の環境変数を探します。環境変数の値が空の場合は、現在のユーザのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイルです。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。この変数でファイル内のどのプロファイルを使用するかを制御します。

      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていなければデフォルトになります。

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの同時実行数です。

      同じファイルの複数のチャンクが同時にアップロードされます。

      あなたが高速リンクを使用して大きなファイルの小数を転送しており、この転送があなたの帯域幅を十分に利用していない場合、これを増やすことで転送を高速化するのに役立つ場合があります。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      trueの場合（デフォルト）、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、falseに設定する必要があります - rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2の認証を使用します。

      false（デフォルト）の場合、rcloneはv4の認証を使用します。設定されている場合、rcloneはv2の認証を使用します。

      v4シグネチャが機能しない場合にのみ使用してください。例：jewel/v10 CEPH以前。

   --list-chunk
      リストのチャンクのサイズ（各ListObject S3リクエストの応答リスト）。

      このオプションは、AWS S3の仕様では「MaxKeys」ともしています。大抵のサービスは、リクエストされたよりもそれ以上の場合でも、応答リストを1000オブジェクトに切り詰めます。AWS S3ではこれはグローバルな最大値であり、変更することはできません。「[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)」を参照してください。 Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン: 1、2、または自動設定の場合は0です。

      S3が最初にリリースされたとき、バケット内のオブジェクトの列挙を行うためのListObjects呼び出しのみ提供されていました。

      しかし2016年5月、ListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを持ち、可能な限り使用するべきです。

      デフォルトの0に設定すると、rcloneはプロバイダの設定に応じてどのlist objectsメソッドを呼び出すかを推測します。もし推測が間違っていた場合は、ここで手動で設定することもできます。

   --list-url-encode
      リストをURLエンコードするかどうか: true/false/unset

      一部のプロバイダは、ファイル名に制御文字を使用する場合など、URLエンコードされたリストをサポートしています。利用可能な場合、これはファイル名で制御文字を使用するとより信頼性が高くなります。unsetに設定されている場合（デフォルト）、rcloneはどの選択を適用するかに応じてプロバイダの設定を選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      カゴの存在を確認するか、または作成しようとしません。

      バケットが既に存在することを知っている場合、rcloneが行うトランザクションの数を最小限に抑えるためにこれが有用な場合があります。

      バケット作成権限を持っていないユーザーを使用している場合は、このパーミッションが必要になる場合もあります。v1.52.0以前では、このバグのためにこの操作は静かにパスされました。

   --no-head
      アップロードされたオブジェクトのHEADリクエストを使用して整合性をチェックしません。

      rcloneがPUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたと想定します。

      特に、以下のものが想定されます：

      - メタデータ（modtime、storage class、content type）はアップロード時と同じである
      - サイズはアップロード時と同じである

      このフラグは、以下の単一パートPUTの応答から読み取ります：

      - MD5SUM
      - アップロード日

      これらはマルチパートアップロードでは読み取られません。

      不明な長さのソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを送信します。

      このフラグを設定すると、アップロードの途中でのアップロードエラーの可能性が増えます。特にサイズが正しくない場合などですから、通常の操作にはお勧めしません。実際には、このフラグを使用してもアップロードエラーの可能性は非常に低いです。

   --no-head-object
      GETする前にHEADを行わないようにします。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをどのくらいの頻度でフラッシュするかです。

      追加のバッファを必要とするアップロード（たとえばマルチパート）は、割り当てのためにメモリプールを使用します。このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうかです。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2の問題が解決できていません。 HTTP/2は、s3バックエンドのデフォルトで有効ですが、ここで無効にすることができます。 問題が解決されると、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      通常、AWS S3はCloudFrontネットワークを介してデータをダウンロードするとより安価な転送費用がかかるため、通常はCloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      true、false、または設定しないでデフォルトを使用します。

   --use-presigned-request
      シングルパートのアップロードで署名済みリクエストまたはPutObjectを使用するかどうか。

      falseの場合、rcloneはPutObject（AWS SDK）を使用してオブジェクトをアップロードします。

      rclone < 1.59のバージョンでは、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。 これは例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定された時点でのファイルバージョンを表示します。

      パラメータは日付、「2006-01-02」、日時「2006-01-02 15:04:05」、またはその前の時間に対する期間、「100d」または「1h」などです。

      このモードでは、ファイルの書き込み操作は許可されませんので、ファイルのアップロードまたは削除はできません。

      有効な形式については、[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトをデコンプレスします。

      S3に「Content-Encoding: gzip」が設定されたままオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding: gzip」が指定されていると受け取ったまま解凍します。これにより、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがobjejectをgzipで圧縮する可能性がある場合に設定してください。

      通常、プロバイダはダウンロード時にオブジェクトを変更しません。`Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時にもそれが設定されません。

      ただし、一部のプロバイダ（Cloudflareなど）は、`Content-Encoding: gzip`でない場合でもオブジェクトをgzip圧縮する場合があります。

      これは、次のようなエラーメッセージを受け取ることが原因です。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、chunked transfer encodingが設定されているオブジェクトをダウンロードする場合、rcloneはオブジェクトをリアルタイムで解凍します。

      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここではrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する


OPTIONS:
   --access-key-id value        AWSアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用されるプリセットACLです。 [$ACL]
   --endpoint value             Qiniuオブジェクトストレージのエンドポイントです。 [$ENDPOINT]
   --env-auth                   ランタイムからAWS認証情報を取得します（環境変数または、env変数の場合はEC2 / ECSメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  リージョンに一致させる必要があるロケーション制約です。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --storage-class value        Qiniuで新しいオブジェクトを保存するためのストレージクラスです。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるバケットのプリセットACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクのサイズです。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値です。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトをデコンプレスします。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しないでください。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクのサイズ（各ListObject S3リクエストの応答リスト）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1、2、または0 for auto。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パート数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをどのくらいの頻度でフラッシュするか。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがobjejectをgzipで圧縮する可能性がある場合に設定してください。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                カゴの存在を確認するか、または作成しようとしません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADリクエストを使用して整合性をチェックしません。 (default: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを行わないようにします。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]							
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードで署名済みリクエストまたはPutObjectを使用するかどうか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2の認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定された時点でのファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（自動生成されます）
   --path value  ストレージのパス

```
{% endcode %}