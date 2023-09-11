# Arvan Cloud Object Storage（AOS）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 arvancloud - Arvan Cloud Object Storage（AOS）の作成

USAGE:
   singularity storage create s3 arvancloud [command options] [arguments...]

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはenv varsなしの場合はEC2 / ECSメタデータ）。

      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例：
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境からAWSの認証情報を取得します（env varsまたはIAM）。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたはランタイム認証情報の場合は空白のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイム認証情報の場合は空白のままにします。

   --endpoint
      Arvan Cloud Object Storage（AOS）APIのエンドポイント。

      例：
         | s3.ir-thr-at1.arvanstorage.com | デフォルトのエンドポイント - よくわからない場合の良い選択肢です。
         |                                | イラン、テヘラン（アジアテック）
         | s3.ir-tbz-sh1.arvanstorage.com | イラン、タブリーズ（シャリアール）

   --location-constraint
      エンドポイントと一致する場所制約。

      バケットを作成する場合にのみ使用されます。

      例：
         | ir-thr-at1 | イラン、テヘラン（アジアテック）
         | ir-tbz-sh1 | イラン、タブリーズ（シャリアール）

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用される事前設定のACL。

      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合はバケットの作成時にも使用されます。

      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      サーバーサイドでオブジェクトをコピーする場合、S3はソースからACLをコピーするのではなく、新しく書き込みます。

      aclが空の文字列の場合、X-Amz-Acl：ヘッダーは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用される事前設定のACL。

      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      このACLはバケットの作成時にのみ適用されます。設定されていない場合は「acl」が代わりに使用されます。

      aclとbucket_aclが空の文字列の場合、X-Amz-Acl：ヘッダーは追加されず、デフォルト（private）が使用されます。

      例：
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADアクセスがあります。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADおよびWRITEアクセスがあります。
         |                    | バケットにこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループにはREADアクセスがあります。

   --storage-class
      新しいオブジェクトを保存するときに使用するストレージクラス。

      例：
         | STANDARD | 標準のストレージクラス

   --upload-cutoff
      チャンク化アップロードに切り替える閾値。

      この閾値より大きいファイルは、chunk_sizeごとにチャンク単位でアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクサイズ。

      upload_cutoffより大きいサイズのファイル、またはサイズが不明なファイル（「rclone rcat」でのアップロード、または「rclone mount」またはGoogleフォトやGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意："--s3-upload-concurrency"は、このチャンクサイズごとに転送ごとにメモリにバッファリングされます。

      高速回線を介して大きなファイルを転送し、十分なメモリがある場合は、これを増やすことで転送速度を向上させることができます。

      Rcloneは、10,000のチャンク制限を超えないように、既知のサイズの大きなファイルをアップロードするときに自動的にチャンクサイズを増やします。

      サイズが不明なファイルは、設定されたチャンクサイズでアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大10,000個のチャンクがあります。したがって、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードしたい場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを増やすと、進捗状況の統計情報の精度が低下します。Rcloneは、チャンクがAWS SDKによってバッファに送信されたときにチャンクとして送信したと見なしますが、まだアップロード中である場合があります。チャンクサイズが大きくなるほど、AWS SDKバッファが大きくなり、進捗報告が真実から逸脱することになります。

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数。

      このオプションは、マルチパートアップロード時に使用する最大マルチパートチャンク数を定義します。

      AWS S3のspecificationである10,000のチャンクをサポートしていない場合に使用することができます。

      Rcloneは、既知のサイズの大きなファイルをアップロードするときに10,000個のチャンク制限を超えないように、自動的にチャンクサイズを増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。

      サーバーサイドでコピーする必要があるこのサイズより大きいファイルは、このサイズのチャンクでコピーされます。

      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しない。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加してアップロードします。これはデータの整合性チェックに優れていますが、大きなファイルをアップロードするときには長時間の遅延を引き起こす可能性があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を探します。環境変数の値が空の場合、デフォルトは現在のユーザーのホームディレクトリです。

          Linux / OSX："$ HOME / .aws / credentials"
          Windows："% USERPROFILE% \ .aws \ credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用されるプロファイルを制御します。

      空の場合、環境変数"AWS_PROFILE"または設定されていない場合は「default」をデフォルトとします。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクを同時にアップロードします。

      高速リンク上で少数の大きなファイルを転送し、これらのアップロードがバンド幅を十分に活用しない場合は、これを増やすと転送が高速化される場合があります。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルアクセスを使用します。

      これがtrueの場合（デフォルト）、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細は[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)をご覧ください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります- rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2認証を使用します。

      真の場合（デフォルトではfalse）、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4シグネチャが機能しない場合にのみ使用します。例：Jewel / v10 CEPHの前

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストごとのレスポンスリスト）。

      このオプションは、AWS S3の仕様の「MaxKeys」、「max-items」、「page-size」としても知られています。
      多くのサービスは、1000を超えるオブジェクトを要求しても、この値に切り捨てられます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1,2、または自動（0）。

      S3が最初に提供された時、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは高性能ですので、できるだけ使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダに基づいて呼び出すリストオブジェクトのメソッドを推測します。推測が誤っている場合は、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダは、ファイル名に制御文字を使用する場合、URLエンコードリストをサポートしており、これはファイル名の制御文字を使用する場合の信頼性が高いです。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットの存在を確認せず、作成もしません。

      バケットが既に存在する場合、rcloneが実行するトランザクションの数を最小限にするために便利です。

      バケット作成のアクセス権限がない場合にも必要です。バージョン1.52.0より前では、バグのためにこれが黙ってパスされたはずです。

   --no-head
      HEADERリクエストなしでアップロード済みのオブジェクトの整合性を確認しません。

      rcloneはアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたものと見なします。

      特に、次のことを想定しています。

      - メタデータ（modtime、ストレージクラス、コンテンツタイプを含む）がアップロード時と同じであったこと
      - サイズがアップロード時と同じであったこと

      アップロードしたパートの単一のPUTリクエストの場合、以下の項目をレスポンスから読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートのアップロードの場合、これらの項目は読み取られません。

      長さ不明のソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを行います。

      このフラグを設定すると、アップロードの失敗が検出される確率が増えます。特に、サイズが間違っている場合など、アップロードの失敗が検出されない確率が高くなります。通常の操作では推奨されません。実際には、このフラグが設定されていても、アップロードの失敗が検出される確率は非常に小さいです。

   --no-head-object
      GETを実行する前にHEADを実行しない。

   --encoding
      バックエンドのエンコーディング。

      詳細は[概要（/overview/#encoding）のエンコーディングセクションを参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする間隔。

      追加バッファを必要とするアップロード（マルチパートなど）は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（具体的にはminio）バックエンドとHTTP/2に関して未解決の問題があります。S3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。

      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3はCloudFrontネットワークを介してダウンロードするデータに対してより安価な出口を提供しているため、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      検証のためにマルチパートアップロードでETagを使用するかどうか。

      true、false、またはデフォルト値を使用するか、未設定のままにします。

   --use-presigned-request
      1つのパートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン<1.59は、単一のパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定するとその機能が再有効になります。これは、例外的な状況またはテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時点でのファイルのバージョンを表示します。

      パラメータは日付、「2006-01-02」、日付時刻「2006-01-02 15:04:05」、またはその時間よりも前の期間、「100d」または「1h」などです。

      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードまたは削除することはできません。

      有効な形式については、[time option docs](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。

      S3には「Content-Encoding：gzip」が設定されている状態でオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding：gzip」で受信したときに展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定してください。

      通常、プロバイダはダウンロード時にオブジェクトを変更しません。`Content-Encoding：gzip`でアップロードされていない場合、ダウンロード時には設定されません。

      ただし、一部のプロバイダは、`Content-Encoding：gzip`でないオブジェクトをgzip化する場合があります（例：Cloudflare）。

      これを設定し、rcloneがContent-Encoding：gzipとチャンク転送エンコーディングでオブジェクトをダウンロードすると、rcloneはオブジェクトをリアルタイムで展開します。

      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する


OPTIONS:
   --access-key-id value        AWSのアクセスキーID。[$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用される事前設定のACL。[$ACL]
   --endpoint value             Arvan Cloud Object Storage（AOS）APIのエンドポイント。[$ENDPOINT]
   --env-auth                   ランタイムからAWSの認証情報を取得します（環境変数またはenv varsなしの場合はEC2 / ECSメタデータ）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  エンドポイントと一致する場所制約。[$LOCATION_CONSTRAINT]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]
   --storage-class value        新しいオブジェクトを保存するときに使用するストレージクラス。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用される事前設定のACL。[$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルアクセスを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストごとのレスポンスリスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1,2または0（自動）。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする間隔。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定してください。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、作成もしません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADERリクエストなしでアップロード済みのオブジェクトの整合性を確認しません。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GETを実行する前にHEADを実行しない。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替える閾値。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       検証のためにマルチパートアップロードでETagを使用するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          1つのパートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。 (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルのバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (デフォルト: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}