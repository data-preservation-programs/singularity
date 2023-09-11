# Netease Object Storage（NOS）

{% code fullWidth="true" %}
```
NAME:
    singularity storage update s3 netease - Netease Object Storage（NOS）

USAGE:
    singularity storage update s3 netease [command options] <name | id>

DESCRIPTION:
    --env-auth
        実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSのメタデータ）。

        access_key_idとsecret_access_keyが空白の場合のみ適用されます。

        例:
            | false | 次のステップでAWSの認証情報を入力します。
            | true  | 環境変数（環境変数またはIAM）からAWSの認証情報を取得します。

    --access-key-id
        AWSのアクセスキーIDです。

        匿名アクセスまたは実行時の資格情報にしたい場合は空白のままにしてください。

    --secret-access-key
        AWSのシークレットアクセスキー（パスワード）です。

        匿名アクセスまたは実行時の資格情報にしたい場合は空白のままにしてください。

    --region
        接続するリージョンです。

        S3クローンを使用している場合でリージョンが不明な場合は空白のままにしてください。

        例:
            | <unset>            | 分からない場合はこれを使用します。
            |                    | v4シグネチャと空のリージョンを使用します。
            | other-v2-signature | v4シグネチャが動作しない場合にのみ使用します。
            |                    | 例：Jewel/v10 CEPH以前。

    --endpoint
        S3 APIのエンドポイントです。

        S3クローンを使用する場合は必須です。

    --location-constraint
        リージョンと一致させる必要のあるロケーション制約です。

        よく分からない場合は空白のままにしてください。バケットを作成する際にのみ使用されます。

    --acl
        バケットの作成、オブジェクトの保存またはコピー時に使用するCanned ACLです。

        このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。

        詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。

        サーバーサイドでオブジェクトをコピーする場合、このACLはS3によって適用されます。
        S3はソースからACLをコピーするのではなく、新しいACLを書き込みます。

        aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。

    --bucket-acl
        バケットの作成時に使用するCanned ACLです。

        詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。

        このACLはバケットの作成時にのみ適用されます。設定されていない場合は、「acl」が代わりに使用されます。

        aclとbucket_aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。

        例:
            | private            | オーナーはFULL_CONTROL権限を持ちます。
            |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
            | public-read        | オーナーはFULL_CONTROL権限を持ちます。
            |                    | AllUsersグループには読み取りアクセス権限があります。
            | public-read-write  | オーナーはFULL_CONTROL権限を持ちます。
            |                    | AllUsersグループには読み取りおよび書き込みアクセス権限があります。
            |                    | バケットへの付与は一般的に推奨されません。
            | authenticated-read | オーナーはFULL_CONTROL権限を持ちます。
            |                    | AuthenticatedUsersグループには読み取りアクセス権限があります。

    --upload-cutoff
        分割アップロードに切り替えるための閾値です。

        これより大きいサイズのファイルは、chunk_size単位で分割してアップロードされます。
        最小は0、最大は5 GiBです。

    --chunk-size
        アップロードに使用するチャンクサイズです。

        upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのものや「rclone mount」またはGoogleフォトやGoogleドキュメントでアップロードされたものなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

        メモ: "--s3-upload-concurrency"のチャンクサイズのチャンクは、転送ごとにメモリ上でバッファリングされます。

        高速リンクを介して大きなファイルを転送しており、十分なメモリを持っている場合、これを増やすと転送が高速化されます。

        Rcloneは、すでに知られている大きなサイズのファイルをアップロードするときに、10,000チャンクの制限以下になるように、自動的にチャンクサイズを増やします。

        サイズが不明なファイルは設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBで、最大10,000のチャンクが存在できるため、デフォルトでは48 GiBまでのファイルサイズをストリームアップロードすることができます。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

        チャンクサイズを増やすと、進捗状況の統計情報の精度が低下します。Rcloneは、AWS SDKによってバッファリングされたチャンクが送信されたときにチャンクとして処理しますが、実際にはまだアップロードされている場合もあります。チャンクサイズが大きくなると、AWS SDKのバッファと進捗状況の報告の信頼性が低下します。

    --max-upload-parts
        マルチパートアップロードで使用するパートの最大数です。

        このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。

        10,000のチャンクの仕様に対応していないサービスがある場合に便利です。

        Rcloneは、既知のサイズの大きなファイルをアップロードする場合、このチャンクサイズを自動的に増やして、このチャンク数の制限以下になるようにします。

    --copy-cutoff
        サーバーサイドコピーする必要のあるこれより大きなサイズのファイルは、このサイズのチャンクでコピーされます。

        最小は0、最大は5 GiBです。

    --disable-checksum
        オブジェクトのメタデータにMD5チェックサムを保存しません。

        通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長い遅延が発生する場合があります。

    --shared-credentials-file
        共有認証情報ファイルへのパスです。

        env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

        この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。

            Linux / OSX: "$HOME/.aws/credentials"
            Windows: "%USERPROFILE%\.aws\credentials"

    --profile
        共有認証情報ファイルで使用するプロファイルです。

        env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。

        空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合にデフォルトになります。

    --session-token
        AWSのセッショントークンです。

    --upload-concurrency
        マルチパートアップロードに使用する並行性です。

        同じファイルのチャンクを同時にアップロードする数です。

        高速リンクを介して少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合、これを増やすと転送が高速化されるかもしれません。

    --force-path-style
        trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

        これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

        一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります。Rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

    --v2-auth
        trueの場合、v2認証を使用します。

        これがfalse（デフォルト）に設定されている場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

        v4シグネチャが動作しない場合にのみ使用してください。例：Jewel/v10 CEPH以前。

    --list-chunk
        リストチャンクのサイズ（各ListObject S3リクエストの応答リスト）です。

        このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、または「page-size」としても知られています。
        ほとんどのサービスは、要求されたより多くのリストを1000オブジェクトに切り捨てます。
        AWS S3ではこれがグローバルな最大値であり、変更することはできません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
        Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

    --list-version
        使用するListObjectsのバージョン：1、2、または0（自動）です。

        S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されていました。

        ただし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを提供し、可能な限り使用する必要があります。

        デフォルトの0に設定されている場合、rcloneはプロバイダの設定に基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が間違っている場合は、ここで手動で設定できます。

    --list-url-encode
        リストをURLエンコードするかどうか：true/false/unsetです。

        一部のプロバイダは、ファイル名で制御文字を使用する場合、URLエンコードされたリストをサポートしています。利用可能な場合、これは制御文字を使用する際のより信頼性のある方法です。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するように選択しますが、ここでrcloneの選択を上書きすることもできます。

    --no-check-bucket
        バケットの存在をチェックしたり作成したりしようとしない場合は設定してください。

        バケットが既に存在する場合、rcloneが行うトランザクションの数を最小限にするため、これは便利です。

        バケット作成権限を持たない場合にも必要になる場合があります。v1.52.0より前では、バグのため、これは無音で渡されます。

    --no-head
        アップロードされたオブジェクトのHEADを行って整合性をチェックしない場合は設定してください。

        rcloneは、PUTによるオブジェクトのアップロード後に200 OKメッセージを受信した場合、正しくアップロードされたと想定します。

        特に、次の項目を想定します。

        - メタデータ（modtime、ストレージクラス、コンテンツタイプ）はアップロード時のものと同じであったこと
        - サイズはアップロード時のものであったこと

        単一パートPUTの応答から以下の項目を読み込みます。

        - MD5SUM
        - アップロード日

        マルチパートアップロードの場合、これらの項目は読み込まれません。

        サイズが不明なソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを行います。

        このフラグを設定すると、認識されていないアップロードの失敗の可能性が増加します。
        特に、誤ったサイズとするアップロードの失敗の確立が高くなるため、通常の操作では推奨されません。実際には、このフラグがあっても、アップロードの失敗が検出されない確率は非常に低いです。

    --no-head-object
        GETの前にHEADを実行しない場合は設定してください。

    --encoding
        バックエンドの文字コードです。

        詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

    --memory-pool-flush-time
        内部メモリバッファープールのフラッシュ頻度です。

        追加のバッファ（たとえばマルチパート）が必要なアップロードでは、メモリプールを使用して割り当てられます。
        このオプションは、未使用のバッファがプールから削除される頻度を制御します。

    --memory-pool-use-mmap
        内部メモリプールでmmapバッファを使用するかどうか。

    --disable-http2
        S3バックエンドのhttp2の使用を無効にします。

        現在、s3（特にminio）バックエンドとHTTP/2の問題が未解決です。
        S3バックエンドではHTTP/2がデフォルトで有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。

        参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

    --download-url
        ダウンロード用のカスタムエンドポイントです。
        通常、AWS S3はCloudFront CDN URLに設定され、CloudFrontネットワークを介してダウンロードされるデータの出力が安価になります。

    --use-multipart-etag
        マルチパートアップロードでETagを検証するかどうか。

        true、false、またはデフォルトを使用するために設定してください。

    --use-presigned-request
        シングルパートのアップロードに署名済みのリクエストまたはPutObjectを使用するかどうか。

        これがfalseの場合、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。

        rcloneのバージョン < 1.59では、単一のパートオブジェクトをアップロードするために署名済みのリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、例外的な状況やテスト以外では必要ありません。

    --versions
        ディレクトリリストに古いバージョンを含めます。

    --version-at
        指定された時点でのファイルバージョンを表示します。

        パラメータは、日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはその時間前の期間「100d」または「1h」のいずれかにする必要があります。

        このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。

        有効な形式については、[時刻オプションドキュメント](/docs/#time-option)を参照してください。

    --decompress
        設定すると、gzipでエンコードされたオブジェクトを解凍します。

        S3へのアップロード時に「Content-Encoding: gzip」が設定されたオブジェクトを通常、rcloneは圧縮されたオブジェクトとしてダウンロードします。

        このフラグが設定されている場合、rcloneは「Content-Encoding: gzip」で受信したファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。

    --might-gzip
        バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定してください。

        通常、プロバイダはオブジェクトをダウンロードする際には変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトの場合、ダウンロード時には設定されません。

        ただし、一部のプロバイダ（たとえばCloudflare）は、`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzipで圧縮する場合があります。

        これが設定されている場合、rcloneがContent-Encoding: gzipが設定されており、チャンク化された転送エンコーディングを使用してオブジェクトをダウンロードすると、rcloneはオブジェクトをオンザフライで解凍します。

        unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するように選択しますが、ここでrcloneの選択を上書きすることもできます。

    --no-system-metadata
        システムメタデータの設定と読み取りを抑制します


OPTIONS:
    --access-key-id value        AWSのアクセスキーIDです。[$ACCESS_KEY_ID]
    --acl value                  バケットの作成、オブジェクトの保存またはコピー時に使用するCanned ACLです。[$ACL]
    --endpoint value             S3 APIのエンドポイントです。[$ENDPOINT]
    --env-auth                   実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSのメタデータ）。（デフォルト：false）[$ENV_AUTH]
    --help, -h                   ヘルプを表示する
    --location-constraint value  リージョンと一致させる必要のあるロケーション制約です。[$LOCATION_CONSTRAINT]
    --region value               接続するリージョンです。[$REGION]
    --secret-access-key value    AWSのシークレットアクセスキー（パスワード）です。[$SECRET_ACCESS_KEY]

    Advanced

    --bucket-acl value               バケットの作成時に使用するCanned ACLです。[$BUCKET_ACL]
    --chunk-size value               アップロードに使用するチャンクサイズです。（デフォルト： "5Mi"）[$CHUNK_SIZE]
    --copy-cutoff value              サーバーサイドコピーする必要のあるこれより大きなサイズのファイルは、このサイズのチャンクでコピーされます。（デフォルト： "4.656Gi"）[$COPY_CUTOFF]
    --decompress                     設定すると、gzipでエンコードされたオブジェクトを解凍します。（デフォルト：false）[$DECOMPRESS]
    --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。（デフォルト：false）[$DISABLE_CHECKSUM]
    --disable-http2                  S3バックエンドのhttp2の使用を無効にします。（デフォルト：false）[$DISABLE_HTTP2]
    --download-url value             ダウンロード用のカスタムエンドポイントです。[$DOWNLOAD_URL]
    --encoding value                 バックエンドの文字コードです。（デフォルト： "Slash,InvalidUtf8,Dot"）[$ENCODING]
    --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。（デフォルト：true）[$FORCE_PATH_STYLE]
    --list-chunk value               リストチャンクのサイズ（各ListObject S3リクエストの応答リスト）です。（デフォルト：1000）[$LIST_CHUNK]
    --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset（デフォルト："unset"）[$LIST_URL_ENCODE]
    --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）（デフォルト：0）[$LIST_VERSION]
    --max-upload-parts value         マルチパートアップロードで使用するパートの最大数（デフォルト：10000）[$MAX_UPLOAD_PARTS]
    --memory-pool-flush-time value   内部メモリバッファープールのフラッシュ頻度（デフォルト："1m0s"）[$MEMORY_POOL_FLUSH_TIME]
    --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト：false）[$MEMORY_POOL_USE_MMAP]
    --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定してください。（デフォルト："unset"）[$MIGHT_GZIP]
    --no-check-bucket                バケットの存在をチェックしたり作成したりしようとしない場合は設定してください。（デフォルト：false）[$NO_CHECK_BUCKET]
    --no-head                        アップロードされたオブジェクトのHEADを行って整合性をチェックしない場合は設定してください。（デフォルト：false）[$NO_HEAD]
    --no-head-object                 GETの前にHEADを実行しない場合は設定してください。（デフォルト：false）[$NO_HEAD_OBJECT]
    --no-system-metadata             システムメタデータの設定と読み取りを抑制します（デフォルト：false）[$NO_SYSTEM_METADATA]
    --profile value                  共有認証情報ファイルで使用するプロファイルです。[$PROFILE]
    --session-token value            AWSのセッショントークンです。[$SESSION_TOKEN]
    --shared-credentials-file value  共有認証情報ファイルへのパスです。[$SHARED_CREDENTIALS_FILE]
    --upload-concurrency value       マルチパートアップロードに使用する並行性（デフォルト：4）[$UPLOAD_CONCURRENCY]
    --upload-cutoff value            分割アップロードに切り替えるための閾値（デフォルト："200Mi"）[$UPLOAD_CUTOFF]
    --use-multipart-etag value       マルチパートアップロードでETagを検証するかどうか（デフォルト："unset"）[$USE_MULTIPART_ETAG]
    --use-presigned-request          シングルパートのアップロードに署名済みのリクエストまたはPutObjectを使用するかどうか（デフォルト：false）[$USE_PRESIGNED_REQUEST]
    --v2-auth                        trueの場合、v2認証を使用します（デフォルト：false）[$V2_AUTH]
    --version-at value               指定された時点でのファイルバージョンを表示します（デフォルト："off"）[$VERSION_AT]
    --versions                       ディレクトリリストに古いバージョンを含めます（デフォルト：false）[$VERSIONS]

```
{% endcode %}