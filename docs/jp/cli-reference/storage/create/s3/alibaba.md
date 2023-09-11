# Alibaba Cloud Object Storage System (OSS) 以前は Aliyun といました

{% code fullWidth="true" %}
```
名前:
   singularity storage create s3 alibaba - Alibaba Cloud Object Storage System (OSS) 以前は Aliyun

使用法:
   singularity storage create s3 alibaba [コマンドオプション] [引数...]

詳細:
   --env-auth
      AWSの認証情報をランタイムから取得します（環境変数やEC2/ECSメタデータが利用できない場合は）。

      アクセスキーIDとシークレットアクセスキーが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境からAWSの認証情報を取得します（環境変数やIAM）。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたはランタイムの認証情報の場合は空のままにしておく。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイムの認証情報の場合は空のままにしておく。

   --endpoint
      OSS APIのエンドポイント。

      例:
         | oss-accelerate.aliyuncs.com          | グローバルアクセラレーション
         | oss-accelerate-overseas.aliyuncs.com | グローバルアクセラレーション（中国本土外）
         | oss-cn-hangzhou.aliyuncs.com         | 华东1（杭州）
         | oss-cn-shanghai.aliyuncs.com         | 华东2（上海）
         | oss-cn-qingdao.aliyuncs.com          | 华北1（青岛）
         | oss-cn-beijing.aliyuncs.com          | 华北2（北京）
         | oss-cn-zhangjiakou.aliyuncs.com      | 华北3（张家口）
         | oss-cn-huhehaote.aliyuncs.com        | 华北5（呼和浩特）
         | oss-cn-wulanchabu.aliyuncs.com       | 华北6（乌兰察布）
         | oss-cn-shenzhen.aliyuncs.com         | 华南1（深圳）
         | oss-cn-heyuan.aliyuncs.com           | 华南2（河源）
         | oss-cn-guangzhou.aliyuncs.com        | 华南3（广州）
         | oss-cn-chengdu.aliyuncs.com          | 西南1（成都）
         | oss-cn-hongkong.aliyuncs.com         | 中国香港（香港）
         | oss-us-west-1.aliyuncs.com           | 美国西部1（硅谷）
         | oss-us-east-1.aliyuncs.com           | 美国东部1（弗吉尼亚）
         | oss-ap-southeast-1.aliyuncs.com      | 东南亚新加坡东南亚1
         | oss-ap-southeast-2.aliyuncs.com      | 亚太地区悉尼亚太地区2
         | oss-ap-southeast-3.aliyuncs.com      | 东南亚吉隆坡东南亚3
         | oss-ap-southeast-5.aliyuncs.com      | 亚太地区雅加达东南亚5
         | oss-ap-northeast-1.aliyuncs.com      | 亚太地区东京东北亚1
         | oss-ap-south-1.aliyuncs.com          | 亚太地区孟买亚太南1
         | oss-eu-central-1.aliyuncs.com        | 欧洲中部法兰克福中欧1
         | oss-eu-west-1.aliyuncs.com           | 欧洲西部伦敦（英国）
         | oss-me-east-1.aliyuncs.com           | 中东迪拜中东1

   --acl
      バケットの作成とオブジェクトの保存またはコピー時に使用するCanned ACL。

      このACLは、オブジェクトの作成時にも使用され、bucket_aclが設定されていない場合にも使用されます。

      詳細については、[AWSのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      サーバー側でオブジェクトをコピーする際は、S3は元のバケットからACLをコピーせず、新しいACLを書き込みます。

      aclが空の場合、X-Amz-Aclヘッダは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用するCanned ACL。

      詳細については、[AWSのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      バケットの作成時にのみ適用されるACLです。設定されていない場合は、aclが代わりに使用されます。

      aclとbucket_aclが空の文字列である場合、X-Amz-Aclヘッダは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADアクセス権限があります。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADおよびWRITEアクセス権限があります。
         |                    | バケットでこれを設定することは一般的にはお勧めしません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループにはREADアクセス権限があります。

   --storage-class
      OSSに新しいオブジェクトを保存するために使用するストレージクラス。

      例:
         | <unset>     | デフォルト
         | STANDARD    | 標準ストレージクラス
         | GLACIER     | アーカイブストレージモード
         | STANDARD_IA | インフリークエントアクセスストレージモード

   --upload-cutoff
      チャンク化アップロードに切り替えるためのカットオフサイズ。

      これより大きなサイズのファイルは、チャンクサイズ単位でアップロードされます。
      最小は0、最大は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのもの、または「rclone mount」やGoogleフォトやGoogleドキュメントでアップロードされたものなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意："--s3-upload-concurrency"個のチャンクが転送ごとにメモリ上でバッファリングされます。

      高速リンクで大きなファイルを転送しており、メモリが十分にある場合は、これを増やすと転送が高速化されます。

      rcloneは、10,000チャンクの制限を下回るように、既知のサイズの大きなファイルをアップロードする場合に自動的にチャンクサイズを増やします。

      不明なサイズのファイルは、設定されたチャンクサイズでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10,000チャンクがあり、これによりストリーミングアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリーミングアップロードする場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグとともに表示される進行状況の統計の精度が低下します。rcloneは、AWS SDKによってバッファリングされたチャンクが送信されたときにチャンクを送信完了とみなしますが、実際にはまだアップロード中の場合があります。大きなチャンクサイズは、AWS SDKのバッファおよび進行状況の報告の信頼性が低下します。

   --max-upload-parts
      マルチパートアップロードの最大パート数。

      このオプションは、マルチパートアップロードを行う際のパート数の最大値を定義します。

      サービスがAWS S3の10,000パート仕様をサポートしていない場合に役立ちます。

      rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、このパート数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフサイズ。

      サーバーサイドのコピーが必要なこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。

      最小は0、最大は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、それをオブジェクトのメタデータに追加することでデータの整合性チェックを行います。これは大きなファイルのアップロードが開始するまで長時間待つことがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth=trueの場合、rcloneは共有認証情報ファイルを使用することができます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を見ます。環境変数の値が空の場合、このデフォルトは現在のユーザーのホームディレクトリです。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth=trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルト値となります。

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並列性。

      同じファイルのチャンクの数を同時にアップロードします。

      ハイスピードリンクを介して大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすと転送が高速化される場合があります。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      trueであれば（デフォルト）、rcloneはパススタイルアクセスを使用し、falseであればrcloneは仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント（英語）](https://docs.aws.amazon.com/cn_zh_cn/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2の認証を使用します。

      これはfalse（デフォルト）であれば、rcloneはv4の認証を使用します。設定された場合、rcloneはv2の認証を使用します。

      v4の署名が機能しない場合にのみこれを使用してください（例：Jewel/v10の前のCEPH）。

   --list-chunk
      リストチャンクのサイズ（各ListObject S3リクエストごとの応答リスト）。

      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、「page-size」としても知られています。
      多くのサービスでは、要求された件数がそれ以上の場合でも、レスポンスリストを1000件に切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3を参照してください](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動的に0。

      S3は最初にリストバケットのためのListObjectsを提供するだけでした。

      しかし、2016年5月にListObjectsV2コールが導入されました。これははるかに高いパフォーマンスを提供し、可能な限り使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダによって設定されたリストオブジェクトのメソッドを呼び出します。正しく予測できない場合は、ここで手動で設定することもできます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダは、ファイル名に制御文字を使用する場合にURLエンコードのリストをサポートしており、これが利用可能な場合は信頼性が高くなります。これが設定されていない場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットの存在をチェックするか作成しない場合に設定します。

      バケットが既に存在することを確認する必要がない場合や、トランザクションの数を最小限に抑えたい場合に役立ちます。

      使用しているユーザーにバケット作成の許可がない場合にも必要です。v1.52.0より前のバージョンでは、バグのためにこれが無視されていました。

   --no-head
      HEADリクエストを使用してアップロードしたオブジェクトの整合性をチェックしない場合に設定します。

      rcloneは、PUTでオブジェクトをアップロードした後に200 OKメッセージを受信した場合、正しくアップロードされたと見なします。

      特に次の場合には、次の項目を前提とします。

      - メタデータ（modtime、ストレージクラス、コンテンツタイプなど）はアップロード時と同じであること
      - サイズがアップロード時と同じであること
      
      シングルパートPUTのレスポンスから次の項目を読み込みます。

      - MD5SUM
      - アップロードされた日時

      マルチパートアップロードの場合、これらのアイテムは読み込まれません。

      サイズが不明なソースオブジェクトをアップロードする場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、誤ったサイズを含むアップロードの失敗が検出される可能性が高くなりますので、通常の操作では推奨されません。実際には、このフラグがある場合でも、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GETする前にHEADを実行しない場合に設定します。

   --encoding
      バックエンドのエンコーディング。

      [概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      インターナルメモリバッファプールがフラッシュされる頻度。

      追加バッファが必要なアップロード（たとえばマルチパート）では、アロケーションにメモリプールが使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      インターナルメモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効化します。

      s3ベースエンドポイントに関しては、現在未解決の問題があります。S3バックエンドのデフォルトでは、HTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決した際には、このフラグは削除されます。

      参照先: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3がCloudFrontネットワークを介してダウンロードデータのエグレスを提供するため、CloudFront CDN URLに設定されます。

   --use-multipart-etag
      MultipartアップロードでETagを検証に使用するかどうか。

      true、false、またはデフォルトを使用するためにunset（未設定）のいずれかで設定します。

   --use-presigned-request
      単一パートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rclone < 1.59のバージョンでは、単一のパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能を再度有効にします。これは特殊な状況やテスト以外では必要ありません。

   --versions
      古いバージョンをディレクトリリストに含める。

   --version-at
      指定した時点のファイルバージョンを表示します。

      パラメータは、日付「2006-01-02」、日時「2006-01-02
      15:04:05」またはその前の時間の期間、「100d」または「1h」など、設定された時間からの期間です。

      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。

      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定されている場合、gzipエンコードされたオブジェクトを解凍します。

      S3に「Content-Encoding: gzip」が設定されたオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮オブジェクトとしてダウンロードします。

      このフラグを設定すると、rcloneは受信した「Content-Encoding: gzip」ファイルを展開します。これにより、rcloneはサイズとハッシュを確認できませんが、ファイルの内容が展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトがダウンロードされる際には変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトには`Content-Encoding: gzip`が設定されません。

      ただし、一部のプロバイダは、`Content-Encoding: gzip`でアップロードされていないオブジェクトを圧縮する場合があります（Cloudflareなど）。

      次のようなエラーを受け取る場合、これを設定してください。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneが`Content-Encoding: gzip`とチャンク転送エンコーディングの設定されたオブジェクトをダウンロードすると、rcloneはオブジェクトを逐次展開します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み込みを抑制する


オプション:
   --access-key-id value      AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                バケットの作成とオブジェクトの保存またはコピー時に使用するCanned ACLです。 [$ACL]
   --endpoint value           OSS APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                 ランタイムからAWSの認証情報を取得します（環境変数やEC2/ECSメタデータが利用できない場合）。（デフォルト：false） [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --storage-class value      OSSに新しいオブジェクトを保存するために使用するストレージクラスです。 [$STORAGE_CLASS]

   より詳細な設定

   --bucket-acl value               バケットの作成時に使用するCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフサイズです。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定されている場合、gzipエンコードされたオブジェクトを解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズです（各ListObject S3リクエストごとの応答リスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョンです：1,2、または自動的に0。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パート数です。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   インターナルメモリバッファプールがフラッシュされる頻度です。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           インターナルメモリプールでmmapバッファを使用するかどうかです。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックしないか作成しません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してアップロードしたオブジェクトの整合性をチェックしません。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを実行しません。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み込みを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並列性です。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのカットオフサイズです。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       MultipartアップロードでETagを検証に使用するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2の認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       古いバージョンをディレクトリリストに含めます。 (デフォルト: false) [$VERSIONS]

   一般的

   --name value  ストレージの名前（デフォルト：自動生成されます）
   --path value  ストレージのパス

```
{% endcode %}