# Scalewayオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 scaleway - Scalewayオブジェクトストレージ

USAGE:
   singularity storage create s3 scaleway [command options] [arguments...]

DESCRIPTION:
   --env-auth
      実行時にAWSの認証情報（環境変数またはEC2/ECSメタデータ）から取得します。

      access_key_idとsecret_access_keyの値が空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境変数（またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSアクセスキーIDです。

      匿名アクセスまたは実行時の認証情報を使用する場合は空白のままにしてください。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）です。

      匿名アクセスまたは実行時の認証情報を使用する場合は空白のままにしてください。

   --region
      接続するリージョンです。

      例:
         | nl-ams | オランダ、アムステルダム
         | fr-par | フランス、パリ
         | pl-waw | ポーランド、ワルシャワ

   --endpoint
      Scalewayオブジェクトストレージのエンドポイントです。

      例:
         | s3.nl-ams.scw.cloud | アムステルダムエンドポイント
         | s3.fr-par.scw.cloud | パリエンドポイント
         | s3.pl-waw.scw.cloud | ワルシャワエンドポイント

   --acl
      オブジェクトの作成、バケットの作成、およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。

      このACLはオブジェクトの作成およびバケット_ACLが設定されていない場合にも使用されます。

      詳細については、[Amazon S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3はサーバーサイドでオブジェクトをコピーする際、ソースからACLをコピーするのではなく、新しいACLを書き込むため、このACLが適用されます。

      aclが空の場合、「X-Amz-Acl:」ヘッダーは追加されず、デフォルトの「private」が使用されます。

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。

      詳細については、[Amazon S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      aclとbucket_aclが設定されていない場合、「X-Amz-Acl:」ヘッダーは追加されず、デフォルトの「private」が使用されます。

      例:
         | private            | 所有者にFULL_CONTROLを許可します。
         |                    | 他のユーザーはアクセス権がありません（デフォルト）。
         | public-read        | 所有者にFULL_CONTROLを許可します。
         |                    | AllUsersグループはREADアクセス権があります。
         | public-read-write  | 所有者にFULL_CONTROLを許可します。
         |                    | AllUsersグループはREADおよびWRITEアクセス権があります。
         |                    | バケットでこれを設定することは一般的には推奨されていません。
         | authenticated-read | 所有者にFULL_CONTROLを許可します。
         |                    | AuthenticatedUsersグループはREADアクセス権があります。

   --storage-class
      S3に新しいオブジェクトを保存する際に使用するストレージクラスです。

      例:
         | <未設定> | デフォルト。
         | STANDARD | ストリーミングやCDNなどのオンデマンドコンテンツに適したStandardクラス。
         | GLACIER  | アーカイブストレージ。
         |          | 価格は低いですが、利用するためには復元する必要があります。

   --upload-cutoff
      チャンクアップロードに切り替えるためのサイズ制限です。

      これより大きいファイルは、chunk_sizeのサイズのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。

      upload_cutoffよりも大きいファイルやサイズが不明なファイル（「rclone rcat」からのアップロードや「rclone mount」やGoogleフォトやGoogleドキュメントでのアップロードなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      メモリごとに"--s3-upload-concurrency"個のチャンクがこのサイズでバッファリングされます。

      高速リンクで大量のファイルを転送しており、メモリが十分ある場合は、これを増やすと転送が高速化します。

      10,000個のチャンク制限を下回るように、rcloneは既知の大きさの大きなファイルをアップロードする際に自動的にチャンクサイズを増やします。

      サイズが不明なファイルは、設定されたチャンクサイズでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で10,000個のチャンクがあります。したがって、デフォルトではストリーミングアップロードできるファイルの最大サイズは48 GiBになります。より大きなファイルをストリーミングアップロードする場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを増やすと、進行状況の統計情報が"-P"フラグで表示される精度が低下します。rcloneは、AWS SDKによってバッファリングされるチャンクを送信したときに、実際にはまだアップロード中の場合でもチャンクを送信として扱います。チャンクサイズが大きいほど、AWS SDKのバッファも大きくなり、進捗報告が実際と乖離する可能性があります。

   --max-upload-parts
      マルチパートアップロードで使用する最大パート数です。

      このオプションでは、マルチパートアップロード時に使用する最大チャンク数を定義します。

      これは、一部のサービスがAWS S3の10,000チャンクの仕様をサポートしていない場合に便利です。

      rcloneは、既知のサイズの大きなファイルをアップロードする際に、チャンクのサイズを自動的に増やしてこのチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのサイズ制限です。

      サーバーサイドでコピーする必要があるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算してオブジェクトのメタデータに追加します。これはデータの整合性チェックには便利ですが、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有の認証情報ファイルへのパスです。

      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用することができます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」という環境変数を検索します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイルです。

      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用することができます。この変数はそのファイルで使用するプロファイルを制御します。

      空白の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルト値になります。

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行性です。

      同じファイルのチャンク数を同時にアップロードします。

      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすと転送が高速化するかもしれません。

   --force-path-style
      trueの場合、パス形式のアクセスを使用します。falseの場合、仮想ホステッド形式を使用します。

      これがtrue（デフォルト）の場合、rcloneはパス形式のアクセスを使用します。falseの場合、rcloneは仮想パス形式を使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、Tencent COS）では、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4シグネチャが機能しない場合にのみ使用してください（例：Jewel/v10 CEPHの前）。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストの応答リストのサイズ）です。

      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、または「page-size」としても知られています。
      大部分のサービスでは、要求された以上のリスト数が表示されないように、リストを1000オブジェクトに切り詰めます。
      AWS S3では、これはグローバルな最大値であり、変更することはできません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動の0です。

      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjectsの呼び出しのみが提供されました。

      しかし、2016年5月にListObjectsV2の呼び出しが導入されました。これははるかに高いパフォーマンスを持ち、可能な場合は必ず使用する必要があります。

      デフォルトの設定である0に設定されている場合、rcloneは設定されたプロバイダに応じて、どのリストオブジェクトメソッドを呼び出すかを推測します。誤った推測を行うと、ここで手動で設定することがあります。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダは、ファイル名に制御文字を使用する場合に、URLエンコードリストをサポートしています。使用できる場合、これはファイルのコントロール文字を使用する際に信頼性が高まります。これが設定されていない場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットの存在をチェックせず、作成しようとしません。

      バケットが既に存在することを事前に知っている場合、トランザクションの数を最小限にするためにこのフラグを設定すると便利です。

      ユーザーにバケット作成の権限がない場合も必要になる場合があります。v1.52.0より前では、これはバグにより静かにパスされていました。

   --no-head
      アップロードされたオブジェクトのHEADをチェックしない場合に設定します。

      rcloneは、PUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたと仮定します。

      特に次のことを仮定します:

      - メタデータ（modtime、ストレージクラス、コンテンツタイプを含む）がアップロード時と同じであったこと
      - サイズはアップロード時と同じであったこと

      一部のプロバイダの単一部品PUTの応答から次の項目を読み取ります:

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらの項目は読み取られません。

      サイズが不明なソースオブジェクトがアップロードされると、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗を検出しない可能性が高まります。特に、サイズが不正確な場合です。しかし、このフラグを使用しない場合でも、通常の動作ではアップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      オブジェクトの取得前にHEADを実行しない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。

      追加のバッファが必要なアップロード（マルチパートなど）は、割り当てにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（具体的にはminio）バックエンドとHTTP/2に関する解決していない問題があります。HTTP/2はs3バックエンドのデフォルトで有効になっていますが、ここでは無効にすることができます。問題が解決されると、このフラグは削除されます。

      参考: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      AWS S3は、CloudFrontネットワークを介してダウンロードされたデータに対して低価格の出力を提供するため、通常はCloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      これはtrue、false、またはデフォルトを使用するために設定されます。

   --use-presigned-request
      シングルパートアップロードの場合に署名済みリクエストまたはPutObjectを使用するかどうか

      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョンが1.59未満の場合、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。このことは例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時間時点のファイルバージョンを表示します。

      パラメータは、日付（「2006-01-02」）、日時（「2006-01-02 15:04:05」）またはそれ以前の期間（「100d」または「1h」など）です。

      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。

      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、圧縮されたgzipエンコードされたオブジェクトを展開します。

      S3に「Content-Encoding: gzip」が設定された状態でオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に「Content-Encoding: gzip」で展開します。これにより、rcloneはサイズとハッシュをチェックすることができませんが、ファイルのコンテンツは展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトには設定されません。

      ただし、一部のプロバイダは`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzip化する場合があります（例：Cloudflare）。

      これを設定し、rcloneが`Content-Encoding: gzip`が設定されたオブジェクトとチャンク転送エンコーディングをダウンロードすると、rcloneはオブジェクトをダウンロード時に展開します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用することを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value      AWSアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                オブジェクトの作成、バケットの作成、およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value           Scalewayオブジェクトストレージのエンドポイントです。 [$ENDPOINT]
   --env-auth                 実行時にAWSの認証情報（環境変数またはEC2/ECSメタデータ）から取得します。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョンです。 [$REGION]
   --secret-access-key value  AWSシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --storage-class value      S3に新しいオブジェクトを保存する際に使用するストレージクラスです。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのサイズ制限です。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、圧縮されたgzipエンコードされたオブジェクトを展開します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パス形式のアクセスを使用します。falseの場合、仮想ホステッド形式を使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズです。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または自動の0です。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用する最大パート数です。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度です。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しようとしません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADをチェックしない場合に設定します。 (default: false) [$NO_HEAD]
   --no-head-object                 オブジェクトの取得前にHEADを実行しない場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有の認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性です。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのサイズ制限です。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードの場合に署名済みリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時間時点のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}