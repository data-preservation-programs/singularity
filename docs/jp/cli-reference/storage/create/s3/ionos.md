# IONOS Cloud

{% code fullWidth="true" %}
```
NAME:
   シンギュラリティーストレージの作成 s3 ionos - IONOS Cloud

使用法:
   singularity storage create s3 ionos [command options] [arguments...]

説明:
   --env-auth
      AWS認証情報をランタイムから取得します（環境変数または環境変数がない場合、EC2/ECSメタデータから取得します）。

      アクセスキーIDとシークレットアクセスキーが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境からAWS認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSアクセスキーID。

      匿名アクセスまたはランタイム認証情報の場合は空白にします。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイム認証情報の場合は空白にします。

   --region
      バケットが作成され、データが保存されるリージョン。

      例:
         | de           | ドイツ、フランクフルト
         | eu-central-2 | ドイツ、ベルリン
         | eu-south-2   | スペイン、ログローニョ

   --endpoint
      IONOS S3オブジェクトストレージのエンドポイント。

      同じリージョンのエンドポイントを指定します。

      例:
         | s3-eu-central-1.ionoscloud.com | ドイツ、フランクフルト
         | s3-eu-central-2.ionoscloud.com | ドイツ、ベルリン
         | s3-eu-south-2.ionoscloud.com   | スペイン、ログローニョ

   --acl
      バケットを作成し、オブジェクトを保存またはコピーする際に使用されるCanned ACL。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合にはバケットの作成にも使用されます。

      詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3はサーバーサイドでオブジェクトをコピーする際にACLをコピーせず、新たに作成します。

      aclが空の場合は、X-Amz-Acl:ヘッダが追加されず、デフォルト（プライベート）が使用されます。

   --bucket-acl
      バケットを作成する際に使用されるCanned ACL。

      詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      このACLはバケットの作成時にのみ適用されます。設定されていない場合は「acl」が代わりに使用されます。

      「acl」と「bucket_acl」が空の文字列である場合は、X-Amz-Acl:ヘッダが追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | 他のユーザーはアクセス権がありません（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADアクセスが与えられます。
         | public-read-write  | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADおよびWRITEアクセスが与えられます。
         |                    | バケットへのこの権限を付与することは一般的に推奨されていません。
         | authenticated-read | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AuthenticatedUsersグループにはREADアクセスが与えられます。

   --upload-cutoff
      チャンクアップロードに切り替える閾値。

      このサイズを超えるファイルは、chunk_sizeのサイズのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクサイズ。

      upload_cutoffを超えるファイルやサイズが不明なファイル（"rclone rcat"からのアップロードや"rclone mount"やGoogle PhotosやGoogle Docsでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意: "--s3-upload-concurrency" のチャンクサイズは、転送ごとにメモリにバッファリングされます。

      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合、チャンクサイズを増やすと転送が高速化します。

      Rcloneは、10,000チャンクの制限を下回るように、既知の大きなファイルをアップロードする際に自動的にチャンクサイズを増やします。

      サイズの未知のファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10,000チャンクがあるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、進行状況の統計情報の精度が低下します。Rcloneは、AWS SDKによってバッファリングされた段階でチャンクが送信されたと見なしていますが、まだアップロード途中の場合があります。チャンクサイズが大きいと、AWS SDKのバッファも大きくなり、進捗レポートが実際の状況から大きく逸脱することになります。

   --max-upload-parts
      マルチパートアップロードで使用するパートの最大数。

      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。

      サービスがAWS S3の10,000チャンクの仕様をサポートしていない場合に有用です。

      Rcloneは、既知の大きなファイルをアップロードする際に10,000チャンクの制限を下回るようにチャンクサイズを自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替える閾値。

      サーバーサイドでコピーする必要があるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長時間かかることがあります。

   --shared-credentials-file
      共有認証情報ファイルのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を探します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルト値になります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用されるプロファイルを制御します。

      空の場合、環境変数 "AWS_PROFILE" または "default" が設定されていない場合は環境変数の値がデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行数。

      同じファイルのチャンクを同時にアップロードする数です。

      高速リンクを介して大量の大きなファイルを転送し、これらのアップロードが帯域幅を十分に活用しない場合、この数を増やすことで転送を高速化することができます。

   --force-path-style
      trueの場合、pathスタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      これがtrue（デフォルト）の場合、rcloneはpathスタイルのアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいて、自動的にこれを行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      これがfalse（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4シグネチャが機能しない場合にのみ使用してください（例：Jewel/v10 CEPH以前）。

   --list-chunk
      リストのチャンクのサイズ（各ListObject S3リクエストのレスポンスリスト）。

      このオプションはAWS S3仕様の"MaxKeys"、"max-items"、または"page-size"とも呼ばれます。
      ほとんどのサービスは、リクエストされた数を超えた場合でも、レスポンスリストを1000のオブジェクトに切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、"rgw list buckets max chunk"オプションで増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動（0）。

      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを提供し、可能であれば使用する必要があります。

      デフォルトの0に設定すると、rcloneはプロバイダが設定に基づいて呼び出すリストオブジェクトのメソッドを推測します。推測が間違っている場合は、ここで手動で設定できます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      いくつかのプロバイダは、ファイル名に制御文字を含める場合、URLエンコードされたリストをサポートしています。利用可能な場合、これはファイル名に制御文字を使用する際の信頼性が向上します。
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットの存在を確認せず、作成も試みません。

      バケットが既に存在する場合のトランザクションの数を最小限に抑えることを試みる場合に便利です。

      ユーザーにバケット作成権限がない場合にも必要です。バージョン1.52.0より前では、これはバグによりサイレントにパスされます。

   --no-head
      HEADリクエストを使用してアップロードされたオブジェクトの整合性を確認しません。

      rcloneがPUT後に200 OKメッセージを受け取った場合、オブジェクトが正常にアップロードされたと想定します。

      特に次のことを想定します。

      - メタデータ（変更日時、保存クラス、コンテンツタイプなど）がアップロード時と同じであったこと
      - サイズがアップロード時と同じであったこと

      チャンク1つのPUTのレスポンスから、次のアイテムを読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらのアイテムは読み取られません。

      長さが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が増えます。特に、正しいサイズではない場合です。ただし、このフラグを使用すると、通常の操作では推奨されません。実際の操作では、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GETの前にHEADを実行しません。

   --encoding
      バックエンド用のエンコーディング。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールのフラッシュ間隔。

      追加のバッファ（たとえばマルチパート）が必要なアップロードでは、メモリプールを使用して割り当てます。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（具体的にはminio）バックエンドとHTTP/2に関する未解決の問題があります。S3バックエンドではHTTP/2がデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決したら、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3はCloudFront CDN URLとして設定されるため、データはCloudFrontネットワークを介してダウンロードされることで、より安価になります。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      true、false、または未設定にする必要があります。デフォルトはプロバイダに依存します。

   --use-presigned-request
      シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン1.59より前では、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは特殊な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時刻のファイルバージョンを表示します。

      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはそのような過去に起きた時間の間隔（例： "100d" または "1h"）である必要があります。

      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。

      有効なフォーマットについては、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      gzipでエンコードされたオブジェクトを解凍します。

      AWS S3に「Content-Encoding: gzip」を設定してオブジェクトをアップロードすることもできます。通常、rcloneはこれらのファイルを圧縮オブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを"Content-Encoding: gzip"で受信すると、ファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容が解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際には変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトには、ダウンロード時に設定されません。

      ただし、一部のプロバイダは、`Content-Encoding: gzip`でアップロードされていないオブジェクト（例：Cloudflare）をgzip化する場合があります。

      これにより、次のようなエラーメッセージが表示されることがあります。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、チャンク付き転送エンコーディングでオブジェクトをダウンロードすると、rcloneはオブジェクトをリアルタイムで解凍します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value      AWSアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                バケットを作成し、オブジェクトを保存またはコピーする際に使用されるCanned ACL。 [$ACL]
   --endpoint value           IONOS S3オブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --env-auth                 AWS認証情報をランタイムから取得します（環境変数または環境変数がない場合、EC2/ECSメタデータから取得します）。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --region value             バケットが作成され、データが保存されるリージョン。 [$REGION]
   --secret-access-key value  AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットを作成する際に使用されるCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替える閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンド用のエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、pathスタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクのサイズ（各ListObject S3リクエストのレスポンスリスト）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または自動（0） (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用するパートの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールのフラッシュ間隔。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、作成も試みません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してアップロードされたオブジェクトの整合性を確認しません。 (default: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを実行しません。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替える閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時刻のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}