# StackPathオブジェクトストレージ

## 使用法
```
NAME:
   singularity storage update s3 stackpath - StackPathオブジェクトストレージ

USAGE:
   singularity storage update s3 stackpath [command options] <name|id>

DESCRIPTION:
   --env-auth
      ランタイムからAWS認証情報を取得する（環境変数やEC2/ECSメタデータを参照）。

      以下の場合のみ有効:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境からAWS認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたはランタイムの認証情報を利用する場合は空白のままにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイムの認証情報を利用する場合は空白のままにしてください。

   --region
      接続するリージョン。

      S3クローンを使用しており、リージョンが存在しない場合は空白のままにしてください。

      例:
         | <未設定>           | わからない場合はこれを利用します。
         |                   | v4シグネチャと空のリージョンを使用します。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                   | 例: Jewel/v10 CEPH以前のバージョン。

   --endpoint
      StackPathオブジェクトストレージのエンドポイント。

      例:
         | s3.us-east-2.stackpathstorage.com    | 米国東部エンドポイント
         | s3.us-west-1.stackpathstorage.com    | 米国西部エンドポイント
         | s3.eu-central-1.stackpathstorage.com | EUエンドポイント

   --acl
      バケットの作成とオブジェクトの保存またはコピー時に使用するプリセットACL。

      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。

      詳細については以下を参照してください: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意: S3はオブジェクトをサーバ側でコピーする際にACLをコピーせずに新しいACLを書き込むため、このACLが適用されます。

      ACLが空の文字列の場合、X-Amz-Aclヘッダが追加されず、デフォルトの設定（プライベート）が使用されます。

   --bucket-acl
      バケットの作成時に使用するプリセットACL。

      詳細については以下を参照してください: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      注意: バケットを作成する場合にのみこのACLが適用されます。設定されていない場合は「acl」が代わりに使用されます。

      ACLとbucket_aclが空の文字列の場合、X-Amz-Aclヘッダは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | 所有者がFULL_CONTROL権限を持ちます。
         |                   | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | 所有者がFULL_CONTROL権限を持ちます。
         |                   | AllUsersグループに読み取りアクセスがあります。
         | public-read-write  | 所有者がFULL_CONTROL権限を持ちます。
         |                   | AllUsersグループに読み取りおよび書き込みアクセスがあります。
         |                   | バケットでこれを設定することは一般的に推奨されません。
         | authenticated-read | 所有者がFULL_CONTROL権限を持ちます。
         |                   | AuthenticatedUsersグループに読み取りアクセスがあります。

   --upload-cutoff
      チャンクアップロードに切り替えるための上限。

      このサイズより大きいファイルは、chunk_sizeごとにチャンクとしてアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffより大きいファイルや、サイズ不明のファイル（"rclone rcat"からのアップロード、または"rclone mount"やGoogleフォトやGoogleドキュメントからのアップロードなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意: "--s3-upload-concurrency"は、このチャンクサイズごとに転送ごとにメモリ内にバッファされます。

      高速リンクで大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送速度が向上します。

      rcloneは、既知のサイズの大きなファイルをアップロードする際には自動的にチャンクサイズを増やして、10000チャンクの制限を下回るようにします。

      サイズが不明のファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で10000のチャンクがあるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の統計の精度が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされた時点でチャンクが送信されたとみなすため、実際にはまだアップロード中かもしれませんが、チャンクサイズが大きいと、AWS SDKのバッファが大きくなり、進行状況の報告が真実から大きく逸脱します。

   --max-upload-parts
      マルチパートアップロード中のパートの最大数。

      このオプションは、マルチパートアップロードを行う際の最大のマルチパートチャンクの数を定義します。

      これは、サービスが10,000のチャンク仕様をサポートしていない場合に便利です。

      rcloneは、既知のサイズの大きなファイルをアップロードする際には自動的にチャンクサイズを増やして、このチャンクの制限数を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるための上限。

      サーバーサイドでコピーが必要なこのサイズより大きいファイルは、このサイズのチャンクでコピーされます。
 
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しない。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始までに長い遅延が発生することがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリをデフォルト値とします。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロフィールを制御します。

      空の場合、環境変数 "AWS_PROFILE"または "default"がデフォルト値となります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクを同時にアップロードする回数です。

      高速リンクで小数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に活用していない場合、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用し、falseの場合は仮想パススタイルを使用します。詳細については[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります - rcloneはプロバイダの設定に基づいてこれを自動的に実行します。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合は、rcloneはv2認証を使用します。

      v4の署名が機能しない場合にのみ使用してください。例: Jewel/v10 CEPH以前のバージョン。

   --list-chunk
      リストのチャンクのサイズ（各ListObject S3リクエストごとの応答リスト）。

      このオプションは、AWS S3のMaxKeys、max-items、page-sizeなどとしても知られています。
      要求された数よりも多いオブジェクトは、ほとんどのサービスで1000個のリストに切り詰められます。AWS S3では、これはグローバルな最大値であり、変更することはできません。詳細は[こちら](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、"rgw list buckets max chunk"オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または0（auto）。

      S3が最初に提供されたときは、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されていました。

      しかし、2016年5月にListObjectsV2コールが導入されました。これははるかに高速なパフォーマンスがあり、できる限り使用するべきです。

      デフォルトで0に設定されている場合、rcloneはプロバイダの設定に従ってどのリストオブジェクトメソッドを呼び出すか推測します。推測が誤った場合は、ここで手動で設定することもできます。

   --list-url-encode
      リストをURLエンコードするかどうか: true/false/unset

      一部のプロバイダでは、ファイル名に制御文字を使用する場合により信頼性があるURLエンコードリストがサポートされています。これが未設定（デフォルト）の場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、rcloneの選択肢をここで上書きすることもできます。

   --no-check-bucket
      バケットが存在するか、または作成する試行を行わないように設定します。

      既に存在するバケットを知っている場合、rcloneのトランザクション数を最小限に抑えるためにこのフラグを使用することができます。

      バケットの作成権限がない場合に必要になる場合もあります。 v1.52.0より前のバージョンでは、これは無音でパスされるはずでしたが、バグによりそれが発生しました。

   --no-head
      アップロードしたオブジェクトの整合性をチェックするためにHEADリクエストを行わないように設定します。

      rcloneは通常、PUTでオブジェクトをアップロードした後にHEADリクエストを行って、正しくアップロードされたと仮定します。

      特に次のことを仮定します:
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）はアップロード時と同じであると仮定します。
      - サイズはアップロード時と同じであると仮定します。

      以下の項目は、PUTリクエストでのシングルパートの場合に応答から読み込まれます:

      - MD5SUM
      - アップロードされた日付

      マルチパートアップロードでは、これらのアイテムは読み込まれません。

      ソースオブジェクトのサイズが不明な場合、rcloneはHEADリクエストを行います。

      このフラグを設定すると、アップロード後にrcloneが200 OKメッセージを受信する場合、正常にアップロードされたと仮定します。

      特に次のことを仮定します:
      
      - メタデータ、モディファイデート、ストレージクラス、コンテンツタイプがアップロード時と同じであると仮定します。
      - サイズがアップロード時と同じであると仮定します。

      未知の長さのソースオブジェクトをアップロードする場合、rclone **は**HEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出される可能性が高くなります。特に、サイズが間違っている場合のチャンスが増えるため、通常の操作には推奨されません。実際には、このフラグを設定しても、アップロードの失敗を検出する確率は非常に低いです。

   --no-head-object
      GETの前にHEADを行わないように設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールのフラッシュ間隔。

      追加のバッファ（たとえばマルチパートが必要なアップロード）では、メモリプールを使用して割り当てられます。
      このオプションは、使用されなくなったバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2に関連する問題が解決されていません。 S3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。

      通常、AWS S3はCloudFront CDN URLに設定されます。これにより、CloudFrontネットワークを経由してダウンロードされるデータのアウトバウンドコストが削減されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか。

      true、false、またはデフォルトのプロバイダーの設定を使用します。

   --use-presigned-request
      シングルパートのアップロードに署名リクエストを使用するか、PutObjectを使用するか。

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン<1.59は、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、特別な事情やテストを除いては必要ありません。

   --versions
      ディレクトリリスティングに古いバージョンを含める。

   --version-at
      指定した時刻のファイルバージョンを表示します。

      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはそれに対する時間が長い期間、"100d"または"1h"など、です。

      このオプションを使用する場合は、ファイルの書き込み操作は許可されません。したがって、ファイルをアップロードしたり削除したりすることはできません。

      有効な形式については、[time optionドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定された場合、gzipでエンコードされたオブジェクトを解凍します。

      "Content-Encoding: gzip"が設定された状態でオブジェクトをS3にアップロードすることができます。通常、rcloneはこれらのファイルを圧縮オブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneは「Content-Encoding: gzip」で受信されるこれらのファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルのコンテンツは解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。

      通常、プロバイダはダウンロード時にオブジェクトを変更しません。 `Content-Encoding: gzip`でアップロードされなかった場合、ダウンロード時には設定されません。

      ただし、一部のプロバイダ（Cloudflareなど）は、 `Content-Encoding: gzip`でアップロードされていないオブジェクトをgzip圧縮する場合があります。

      これを設定すると、rcloneがContent-Encoding: gzipが設定され、チャンクされた転送エンコーディングがある状態でオブジェクトをダウンロードすると、rcloneはオブジェクトを逐次解凍します。

      これが未設定（デフォルト）の場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択肢を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み込みの抑制

OPTIONS:
   --access-key-id value      AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                バケットの作成とオブジェクトの保存またはコピー時に使用するプリセットACL。 [$ACL]
   --endpoint value           StackPathオブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --env-auth                 ランタイムからAWS認証情報を取得する（環境変数やEC2/ECSメタデータを参照）。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョン。 [$REGION]
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   高度なオプション

   --bucket-acl value               バケットの作成時に使用するプリセットACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための上限。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定された場合、gzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しない。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクのサイズ（各ListObject S3リクエストごとの応答リスト）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（auto）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロード中のパートの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールのフラッシュ間隔。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットが存在するか、または作成する試行を行わないように設定します。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトの整合性をチェックするためにHEADリクエストを行わないように設定します。 (default: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わないように設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み込みの抑制 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるための上限。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名リクエストを使用するか、PutObjectを使用するか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時刻のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリスティングに古いバージョンを含める。 (default: false) [$VERSIONS]
```
