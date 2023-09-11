# Scalewayオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 scaleway - Scalewayオブジェクトストレージ

USAGE:
   singularity storage update s3 scaleway [command options] <name|id>

DESCRIPTION:
   --env-auth
      実行時にAWS認証情報を取得します（環境変数またはEC2/ECSメタデータのいずれか）。

      access_key_idとsecret_access_keyがブランクの場合のみ適用されます。

      例:
         | false | AWS認証情報を次のステップで入力します。
         | true  | 環境（環境変数またはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSアクセスキーIDです。

      匿名アクセスまたは実行時の認証情報を使用する場合は、空白のままにしておきます。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）です。

      匿名アクセスまたは実行時の認証情報を使用する場合は、空白のままにしておきます。

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
      オブジェクトを作成および保存、コピーする際に使用される事前に設定されたACLです。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      サーバーサイドでオブジェクトをコピーする場合、このACLはソースからACLをコピーするのではなく、新しいACLを書き込むため、適用されます。

      aclが空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

   --bucket-acl
      バケットを作成する際に使用する事前に設定されたACLです。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      bucket_aclが設定されていない場合、バケットの作成時にのみこのACLが適用されます。

      aclとbucket_aclの両方が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーにFULL_CONTROLが付与されます。
         |                    | 他の誰もアクセス権がありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLが付与されます。
         |                    | AllUsersグループにREADアクセスが付与されます。
         | public-read-write  | オーナーにFULL_CONTROLが付与されます。
         |                    | AllUsersグループにREADおよびWRITEアクセスが付与されます。
         |                    | バケットに対してこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーにFULL_CONTROLが付与されます。
         |                    | AuthenticatedUsersグループにREADアクセスが付与されます。

   --storage-class
      S3に新しいオブジェクトを保存する際に使用するストレージクラスです。

      例:
         | <unset>  | デフォルトです。
         | STANDARD | 需要に応じたコンテンツ（ストリーミングやCDNなど）に適したStandardクラスです。
         |          | ストレージされたオブジェクト。
         |          | 価格は安くなりますが、アクセスするために復元する必要があります。

   --upload-cutoff
      チャンク化アップロードに切り替えるための閾値です。

      この閾値を超えるファイルは、chunk_sizeごとにチャンク化してアップロードされます。
      0から5 GiBまでの値を設定できます。

   --chunk-size
      アップロードに使用するチャンクのサイズです。

      upload_cutoffよりも大きいファイルや、サイズが不明なファイル（「rclone rcat」や「rclone mount」、GoogleフォトやGoogleドキュメントでアップロードされたファイルなど）の場合、このチャンクサイズを使用してマルチパートでアップロードされます。

      注意：「--s3-upload-concurrency」は、このサイズのチャンクが転送ごとにメモリにバッファリングされます。

      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、この値を増やすことで転送速度が向上します。

      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やして10,000のチャンク制限を下回るようにします。

      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で10,000のチャンクを持つことができます。ストリームアップロードできるファイルの最大サイズを48 GiBに設定しています。より大きなサイズのファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクのサイズを増やすと、進行状況統計の精度が低下します。Rcloneは、AWS SDKにバッファリングされているチャンクが送信されたときにチャンクを送信済みとして扱いますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKのバッファリングが大きくなり、真実とは異なる進行状況が報告されます。

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数です。

      このオプションは、マルチパートアップロードを行う際に使用するパーツの最大数を定義します。

      サービスがAWS S3の10,000のチャンク仕様をサポートしていない場合に有用です。

      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値です。

      サーバーサイドコピーが必要なこの閾値よりも大きなファイルは、このサイズのチャンクでコピーされます。
      0から5 GiBまでの値を設定できます。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータにそれを追加するため、大きなファイルのアップロードには時間がかかることがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を探します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトです。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイルです。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。

      空の場合、デフォルトは環境変数 "AWS_PROFILE" または "default" です。

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行性です。

      同じファイルのチャンクの数を同時にアップロードします。

      高速リンクで数が少ない大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用し、falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に実行します。

   --v2-auth
      trueの場合はv2の認証を使用します。

      これがfalse（デフォルト）に設定されている場合、rcloneはv4の認証を使用します。設定されている場合、rcloneはv2の認証を使用します。

      v4署名が機能しない場合だけ、この設定を使用してください。例えば、Jewel/v10 CEPH以前のバージョンです。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストの応答リストのサイズ）です。

      このオプションは、AWS S3仕様の "MaxKeys"、 "max-items"、または "page-size" としても知られています。
      ほとんどのサービスは、要求された以上の数にリストを切り取ります。
      AWS S3ではこれはグローバルな最大値であり、変更することはできません。詳細については[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションで増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）です。

      S3が最初に開始されたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されていました。

      しかし、2016年5月にはListObjectsV2呼び出しが導入されました。これははるかに高性能であり、可能であれば使用する必要があります。

      デフォルトで設定されている0の場合、rcloneはプロバイダに応じてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が誤っている場合は、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダはリストをURLエンコードしています。利用可能な場合、これはファイル名に制御文字を含めるときにより信頼性が高いです。unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットの存在をチェックせず、作成しようとも試行しません。

      バケットが既に存在する場合、rcloneが実行するトランザクションの数を最小限にする必要がある場合に便利です。

      バケット作成の権限を持たないユーザーを使用する場合にも必要になる場合があります。バージョン1.52.0より前では、このバグのために無駄にパスしていました。

   --no-head
      HEADリクエストを使用してアップロードされたオブジェクトの整合性をチェックしません。

      rcloneは通常、PUTを使用してオブジェクトをアップロードした後、200 OKメッセージを受け取った場合に正しくアップロードされたと想定します。

      特に次のように想定されます：

      - アップロードされたときのメタデータ、修正日時、ストレージクラス、コンテンツタイプがアップロードされたものと同じであること
      - サイズがアップロードされたものと同じであること

      単一パーツのPUTの場合、次のリクエストから次の項目を読み取ります：

      - MD5SUM
      - アップロードされた日付

      マルチパートアップロードの場合、これらの項目は読み取られません。

      サイズが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを送信します。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まります。特に、サイズが正しくない場合などのです。通常の操作では推奨されません。実際、このフラグを使用しても、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GET前にHEADを実行しないようにします。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。

      追加のバッファ（マルチパートなどを必要とするアップロード）を必要とするアップロードは、メモリプールを使用して割り当てを行います。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2の問題が未解決です。s3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決され次第、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631、

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      通常はAWS S3を通じてデータをダウンロードすることで、AWS S3からのダウンロード時の出口が割安になります。

   --use-multipart-etag
      マルチパートアップロード時にETagを使用して検証するかどうか

      true、false、または未設定のいずれかを指定します。プロバイダのデフォルトを使用します。

   --use-presigned-request
      単一パーツのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか

      falseに設定すると、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン<1.59では、単一パーツオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、この機能が再度有効になります。これは、特殊な事情やテスト以外では必要ありません。

   --versions
      ディレクトリリスティングに古いバージョンを含める。

   --version-at
      指定した時間のファイルバージョンを表示します。

      パラメータは日付（ "2006-01-02" ）、datetime（ "2006-01-02
      15:04:05" ）またはその時からの期間、例えば "100d" または "1h" である必要があります。

      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。

      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      オブジェクトがgzipで圧縮されている場合、これを解凍します。

      "Content-Encoding: gzip"が設定されたままS3にオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に "Content-Encoding: gzip" で解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容が解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzipする可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際には変更しません。オブジェクトが「Content-Encoding: gzip」でアップロードされていない場合、ダウンロード時に設定されません。

      ただし、一部のプロバイダ（Cloudflareなど）は「Content-Encoding: gzip」でアップロードされていないオブジェクトに対してもgzip圧縮を適用する場合があります。

      これを設定すると、rcloneがContent-Encoding: gzipが設定され、チャンク化された転送エンコーディングを受け取ると、rcloneはオブジェクトをリアルタイムで解凍します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制します


OPTIONS:
   --access-key-id value      AWSアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                オブジェクトの作成時に使用される事前に設定されたACL。 [$ACL]
   --endpoint value           Scalewayオブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --env-auth                 実行時にAWS認証情報を取得します（環境変数またはEC2/ECSメタデータのいずれか）。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョン。 [$REGION]
   --secret-access-key value  AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --storage-class value      S3に新しいオブジェクトを保存する際に使用するストレージクラス。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用する事前に設定されたACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクのサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzipで圧縮されたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipする可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しようとも試行しません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してアップロードされたオブジェクトの整合性をチェックしません。 (default: false) [$NO_HEAD]
   --no-head-object                 GET前にHEADを実行しないようにします。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロード時にETagを使用して検証するかどうか。 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パーツのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2の認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリスティングに古いバージョンを含める。 (default: false) [$VERSIONS]

```
{% endcode %}