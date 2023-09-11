# IDrive e2

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 idrive - IDrive e2

USAGE:
   singularity storage update s3 idrive [command options] <name|id>

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ（環境変数が存在しない場合）から）。

      `access_key_id` と `secret_access_key` が空白の場合のみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | ランタイムからAWSの認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSのアクセスキーIDです。

      匿名アクセスまたはランタイム認証情報を使用する場合は空白のままにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたはランタイム認証情報を使用する場合は空白のままにしてください。

   --acl
      バケットの作成やオブジェクトの保存、コピー時に使用する既存のACLです。

      このACLはオブジェクトの作成時に使用されるだけでなく、`bucket_acl` が設定されていない場合にはバケットの作成時にも使用されます。

      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3はオブジェクトをサーバー側でコピーする際にACLをコピーしないため、このACLはサーバー側でオブジェクトをコピーする際に適用されます。

      ACLが空の文字列の場合、X-Amz-Aclヘッダは追加されず、デフォルトのACL（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用する既存のACLです。

      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      このACLはバケットの作成時にのみ適用されます。もし設定されていない場合は `acl` が代わりに使用されます。

      `acl` と `bucket_acl` が空の文字列の場合、X-Amz-Aclヘッダは追加されず、デフォルトのACL（private）が使用されます。

      例:
         | private            | オーナーにはFULL_CONTROL権限を付与します。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROL権限を付与します。
         |                    | AllUsersグループにはREAD権限を付与します。
         | public-read-write  | オーナーにはFULL_CONTROL権限を付与します。
         |                    | AllUsersグループにはREADおよびWRITE権限を付与します。
         |                    | バケットに対してこれを設定することは一般的には推奨されません。
         | authenticated-read | オーナーにはFULL_CONTROL権限を付与します。
         |                    | AuthenticatedUsersグループにはREAD権限を付与します。

   --upload-cutoff
      チャンクアップロードに切り替えるサイズの閾値です。

      このサイズを超えるファイルは、chunk_sizeのサイズでチャンクアップロードされます。
      最小値は0、最大値は5ギガバイトです。

   --chunk-size
      アップロード時に使用するチャンクのサイズです。

      `upload_cutoff` を超える大きなファイルや、サイズが不明なファイル（例：`rclone rcat` や "rclone mount" でアップロードされるファイル、GoogleフォトやGoogleドキュメントなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意： "--s3-upload-concurrency"のチャンクサイズは、転送ごとにメモリ上にバッファリングされます。

      高速なリンクで大きなファイルを転送しており、メモリが十分にある場合は、この値を上げることで転送速度を向上させることができます。

      rcloneは、既知のサイズの大きなファイルをアップロードする場合、チャンクサイズを自動的に増やして、10000個のチャンク制限を下回るようにします。

      不明なサイズのファイルは、設定された `chunk_size` でアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大10,000個のチャンクが可能です。そのため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを大きくすると、進捗統計情報（"-P"フラグ）の正確性が低下します。rcloneは、チャンクがAWS SDKによってバッファリングされた場合にチャンクが送信されたと扱いますが、実際にはまだアップロード中の場合があります。大きなチャンクサイズは、AWS SDKバッファの大きさと進捗報告の正確性の誤差を増加させます。

   --max-upload-parts
      マルチパートアップロードで使用するパートの最大数です。

      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。

      これは、サービスがAWS S3の10,000個のチャンク仕様をサポートしていない場合に役立ちます。

      rcloneは、既知のサイズの大きなファイルをアップロードする場合、チャンクサイズを自動的に増やして、このチャンク数制限を下回るようにします。

   --copy-cutoff
      サーバーサイドでコピーする必要があるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5ギガバイトです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまでには長い遅延が発生することがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。

      `env_auth` がtrueの場合、rcloneは共有認証情報ファイルを使用することができます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」という環境変数を探します。環境変数の値が空の場合、カレントユーザーのホームディレクトリがデフォルト値になります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイルです。

      `env_auth` がtrueの場合、rcloneは共有認証情報ファイルを使用します。この変数はそのファイルで使用するプロファイルを制御します。

      何も設定されていない場合は、環境変数「AWS_PROFILE」または「default」がデフォルト値になります。

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行実行数です。

      同じファイルのチャンクが並行してアップロードされます。

      ハイスピードリンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、この数を増やすと転送速度が向上することがあります。

   --force-path-style
      trueの場合、パススタイルのアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。

      これがtrue（デフォルト）に設定されていると、rcloneはパススタイルのアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細は[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要がありますが、プロバイダの設定に基づいてrcloneが自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4署名が機能しない場合にのみ使用してください。例えば、Jewel/v10 CEPH以前のバージョン。

   --list-chunk
      リストレスポンス（ListObject S3リクエストごとの応答リスト）のサイズです。

      このオプションは、AWS S3の仕様の「MaxKeys」、「max-items」、「page-size」とも呼ばれます。
      大抵のサービスは、リクエストされたより多いアイテムが存在する場合でも、レスポンスリストを1000オブジェクトに切り捨てます。
      AWS S3ではこれがグローバル最大値であり、変更できません。詳細は[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションで増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1,2または0で自動です。

      S3が最初にリリースされた当初、バケット内のオブジェクトを列挙するためにListObjects呼び出しのみが提供されていました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高性能であり、可能な限り使用する必要があります。

      デフォルト値である0に設定すると、rcloneはプロバイダ設定に基づいてどのリストオブジェクトメソッドを呼び出すか推測します。推測が間違っている場合は、ここで手動で設定できます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダは、リストをURLエンコードすることをサポートしており、ファイル名に制御文字を使用する際にはこれがより信頼性のある方法です。この値が設定されていない場合（デフォルト）は、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットが存在するか、または作成することを試みないように設定します。

      バケットが既に存在することを確認する必要がない場合に便利です。

      バケットの作成権限がない場合にも必要になることがあります。v1.52.0以前では、これはエラーを無視して通過してしまう不具合がありました。

   --no-head
      アップロード済みのオブジェクトの整合性を確認するためにHEADリクエストを行わないように設定します。

      rcloneは、PUTでオブジェクトをアップロードした後に200 OKのメッセージを受け取った場合、正常にアップロードされたと見なします。

      特に以下のことを前提とします：

      - metadata（modtime、ストレージクラス、コンテンツタイプなど）がアップロード時と同じであること
      - サイズがアップロード時と同じであること

      単一パートPUTの場合、以下の項目をレスポンスから読み込みます：

      - MD5SUM
      - アップロード日時

      マルチパートアップロードの場合、これらの項目は読み込まれません。

      不明な長さのソースオブジェクトがアップロードされると、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まるため、通常の操作では推奨されません。このフラグを設定しても、アップロードの失敗が検出されない可能性は非常に低いです。

   --no-head-object
      オブジェクトを取得する前にHEADを行わないようにします。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部のメモリバッファプールのフラッシュ頻度です。

      追加のバッファが必要なアップロード（たとえばマルチパート）は、割り当てにメモリプールを使用します。
      このオプションは、使用されていないバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2の問題が解決されていません。HTTP/2はs3バックエンドでデフォルトで有効になっていますが、ここで無効にすることができます。問題が解決された場合、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのカスタムエンドポイントです。
      通常、AWS S3は、CloudFrontネットワークを介してダウンロードされたデータに対してエグレス割引が提供されるため、このフィールドにはCloudFront CDNのURLが設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      これはtrue、false、またはデフォルト値を使用するために未設定のいずれかです。

   --use-presigned-request
      単一パートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか

      これがfalseの場合、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョンによっては、単一パートオブジェクトのアップロードに署名済みリクエストを使用し、このフラグをtrueに設定するとその機能が再度有効化されます。これは特殊な場合やテストにのみ必要です。

   --versions
      ディレクトリリスティングに古いバージョンを含める

   --version-at
      指定した時間のファイルバージョンを表示します。

      パラメータは日付（「2006-01-02」）、日時（「2006-01-02 15:04:05」）、またはその前に行われた時間の期間、「100d」または「1h」などです。

      このオプションを使用すると、ファイルへの書き込み操作は許可されませんので、ファイルをアップロードすることや削除することはできません。

      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトの展開を行います。

      S3に「Content-Encoding: gzip」を設定してオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneは受信したデータを「Content-Encoding: gzip」で受信したまま展開します。つまり、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開された状態になります。

   --might-gzip
      バックエンドがオブジェクトにgzipを適用する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。 `Content-Encoding: gzip` でアップロードされなかったオブジェクトは、ダウンロード時にも設定されません。

      ただし、いくつかのプロバイダ（Cloudflareなど）は、`Content-Encoding: gzip` でアップロードされていないオブジェクトに対してもgzipを適用する場合があります。

      これに問題がある場合、次のようなエラーが発生します。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneが `Content-Encoding: gzip` が設定されたオブジェクトをチャンク転送エンコードでダウンロードする場合、rcloneはオブジェクトを動的に展開します。

      これがunset（デフォルト）に設定されている場合は、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する

OPTIONS:
   --access-key-id value      AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                バケットの作成やオブジェクトの保存やコピー時に使用する既存のACLです。 [$ACL]
   --env-auth                 ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータがない場合）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]

   上級者向け

   --bucket-acl value               バケットの作成時に使用する既存のACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクのサイズです。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるサイズの閾値です。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトの展開を行います。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルのアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズです（S3リクエストごとのレスポンスリスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動） (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用するパートの最大数です。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度です。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトにgzipを適用する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットが存在するか、または作成することを試みないように設定します。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロード済みのオブジェクトの整合性を確認するためにHEADリクエストを行わないように設定します。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADリクエストを行わないようにします。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行実行数です。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるサイズの閾値です。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトの展開を行います。 (デフォルト: false) [$DECOMPRESS]
   --might-gzip value               バックエンドがオブジェクトにgzipを適用する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (デフォルト: false) [$NO_SYSTEM_METADATA]

```