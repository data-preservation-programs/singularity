# RackCorp オブジェクトストレージ

{% code fullWidth="true" %}
```
名称:
   singularity storage update s3 rackcorp - RackCorp オブジェクトストレージ

使用法:
   singularity storage update s3 rackcorp [command options] <名前 | ID>

説明:
   --env-auth
      AWSの認証情報を実行時に取得します（環境変数または環境依存のメタデータ）。
      
      access_key_idとsecret_access_keyが空白の場合のみ適用されます。

      例:
         | false | AWSの認証情報を次のステップで入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --secret-access-key
      AWSの秘密アクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --region
      バケットが作成され、データが保存される場所を指定します。
      

      例:
         | global    | グローバルCDN（すべてのロケーション）
         | au        | オーストラリア（すべての州）
         | au-nsw    | NSW（オーストラリア）
         | au-qld    | QLD（オーストラリア）
         | au-vic    | VIC（オーストラリア）
         | au-wa     | パース（オーストラリア）
         | ph        | マニラ（フィリピン）
         | th        | バンコク（タイ）
         | hk        | 香港
         | mn        | ウランバートル（モンゴル）
         | kg        | ビシュケク（キルギス）
         | id        | ジャカルタ（インドネシア）
         | jp        | 東京（日本）
         | sg        | シンガポール
         | de        | フランクフルト（ドイツ）
         | us        | USA（AnyCast）リージョン
         | us-east-1 | ニューヨーク（USA）リージョン
         | us-west-1 | Freemont（USA）リージョン
         | nz        | オークランド（ニュージーランド）リージョン

   --endpoint
      RackCorp オブジェクトストレージのエンドポイント。

      例:
         | s3.rackcorp.com           | グローバル（AnyCast）エンドポイント
         | au.s3.rackcorp.com        | オーストラリア（AnyCast）エンドポイント
         | au-nsw.s3.rackcorp.com    | シドニー（オーストラリア）エンドポイント
         | au-qld.s3.rackcorp.com    | ブリスベン（オーストラリア）エンドポイント
         | au-vic.s3.rackcorp.com    | メルボルン（オーストラリア）エンドポイント
         | au-wa.s3.rackcorp.com     | パース（オーストラリア）エンドポイント
         | ph.s3.rackcorp.com        | マニラ（フィリピン）エンドポイント
         | th.s3.rackcorp.com        | バンコク（タイ）エンドポイント
         | hk.s3.rackcorp.com        | 香港エンドポイント
         | mn.s3.rackcorp.com        | ウランバートル（モンゴル）エンドポイント
         | kg.s3.rackcorp.com        | ビシュケク（キルギス）エンドポイント
         | id.s3.rackcorp.com        | ジャカルタ（インドネシア）エンドポイント
         | jp.s3.rackcorp.com        | 東京（日本）エンドポイント
         | sg.s3.rackcorp.com        | シンガポールエンドポイント
         | de.s3.rackcorp.com        | フランクフルト（ドイツ）エンドポイント
         | us.s3.rackcorp.com        | USA（AnyCast）エンドポイント
         | us-east-1.s3.rackcorp.com | ニューヨーク（USA）エンドポイント
         | us-west-1.s3.rackcorp.com | Freemont（USA）エンドポイント
         | nz.s3.rackcorp.com        | オークランド（ニュージーランド）エンドポイント

   --location-constraint
      バケットが配置され、データが保存される場所を指定します。
      

      例:
         | global    | グローバルCDNリージョン
         | au        | オーストラリア（すべてのロケーション）
         | au-nsw    | NSW（オーストラリア）
         | au-qld    | QLD（オーストラリア）
         | au-vic    | VIC（オーストラリア）
         | au-wa     | パース（オーストラリア）
         | ph        | マニラ（フィリピン）
         | th        | バンコク（タイ）
         | hk        | 香港
         | mn        | ウランバートル（モンゴル）
         | kg        | ビシュケク（キルギス）
         | id        | ジャカルタ（インドネシア）
         | jp        | 東京（日本）
         | sg        | シンガポール
         | de        | フランクフルト（ドイツ）
         | us        | USA（AnyCast）リージョン
         | us-east-1 | ニューヨーク（USA）リージョン
         | us-west-1 | Freemont（USA）リージョン
         | nz        | オークランド（ニュージーランド）リージョン

   --acl
      オブジェクトおよびバケットの作成時に使用するCanned ACL。
      
      このACLはオブジェクトの作成時にも使用され、bucket_aclが設定されていない場合はバケットの作成時にも使用されます。
      
      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      S3では、サーバーサイドのオブジェクトコピー時にACLはソースからコピーされるのではなく、新しいものが書き込まれるため、このACLが適用されます。
      
      aclが空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトの（プライベート）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用するCanned ACL。
      
      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      このACLはバケットの作成時のみ適用されます。設定されていない場合は「acl」が代わりに使用されます。
      
      "acl"と"bucket_acl"が空の文字列の場合、X-Amz-Acl：ヘッダーは追加されず、デフォルトの（プライベート）が使用されます。
      

      例:
         | private            | オーナーがFULL_CONTROLを取得します。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーがFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADアクセスを取得します。
         | public-read-write  | オーナーがFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADおよびWRITEアクセスを取得します。
         |                    | バケットでこれを許可することは一般的にお勧めできません。
         | authenticated-read | オーナーがFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループはREADアクセスを取得します。

   --upload-cutoff
      チャンク化アップロードに切り替えるための閾値。
      
      この閾値を超えるサイズのファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffを超えるサイズのファイルや、サイズの分からないファイル（「rclone rcat」からのアップロードや「rclone mount」またはGoogle
      フォトやGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      注意：「--s3-upload-concurrency」は、搬送ごとにこのサイズのチャンクがメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送して十分なメモリがある場合は、これを増やすと転送速度が向上します。
      
      rcloneはファイルサイズが分かっている大きなファイルをアップロードする際に自動的にチャンクサイズを増やして、10,000のチャンク制限を下回るようにします。
      
      サイズが分からないファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10,000のチャンクまで存在するため、デフォルトの設定ではストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の統計の精度が低下します。rcloneは、AWS SDKによってバッファリングされているチャンクを送信したときにチャンクを送信したと見なしますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKのバッファサイズが大きくなり、真実から逸脱した進行報告が表示されます。
      

   --max-upload-parts
      マルチパートアップロードの最大パーツ数。
      
      このオプションは、マルチパートアップロード時に使用するマルチパートチャンクの最大数を定義します。
      
      これは、サービスがAWS
      S3の10,000のチャンク仕様をサポートしていない場合に便利です。
      
      rcloneはファイルのサイズが分かっている大きなファイルをアップロードする際に自動的にチャンクサイズを増やして、このチャンク数の制限を下回るようにします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値。
      
      サーバーサイドコピーが必要なこの閾値を超えるサイズのファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加できるようにします。これはデータの整合性チェックには非常に役立ちますが、大きなファイルのアップロードを開始するまでに長い待ち時間が発生する場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を参照します。環境の値が空の場合、現在のユーザーのホームディレクトリがデフォルト値となります。
      
          Linux / OSX： "$ HOME / .aws / credentials"
          Windows： 「% USERPROFILE％\.aws\credentials」
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」がセットされていない場合は、デフォルト値となります。
      

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同様なファイルの複数のチャンクが同時にアップロードされます。
      
      高速リンクで数少ない大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用しない場合、これを増やすことで転送を高速化できます。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合、仮想ホストスタイルを使用します。
      
      これがtrueの場合（デフォルト設定）、rcloneはパススタイルアクセスを使用します。それがfalseの場合、rcloneは仮想パススタイルを使用します。詳細については、AWS
      S3ドキュメントを参照してください。
      
      AWS、Aliyun OSS、Netease
      COS、またはTencent COSなどの一部のプロバイダでは、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用してください。たとえば、Jewel / v10 CEPHより前のバージョン。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストのレスポンスリスト）。
      
      このオプションは、AWS
      S3の仕様では「MaxKeys」、「max-items」、または「page-size」としても知られています。ほとんどのサービスは、要求がそれ以上の場合でもレスポンスリストを1000オブジェクトに切り捨てます。AWS
      S3では、これはグローバルな最大値であり、変更できません。詳細については[こちらを参照](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)してください。Cephでは、「rgw
      list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは大幅に高速化され、可能であれば使用するべきです。
      
      デフォルト設定である0に設定すると、rcloneはプロバイダの設定に基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。誤った推測をした場合、ここで手動で設定できます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true / false / unset
      
      一部のプロバイダでは、リストをURLエンコードすることが可能で、ファイル名に制御文字を使用する際にはより信頼性があります。unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しません。
      
      バケットが既に存在することを知っている場合、rcloneが実行するトランザクションの数を最小限に抑えるためにこれを使用すると便利です。
      
      バケット作成の権限を持たないユーザーの場合にも必要です。 v1.52.0より前のバージョンでは、これはバグのために静かに実行されました。
      

   --no-head
      アップロードしたオブジェクトのHEADを行い、整合性をチェックしません。
      
      rcloneはPUT後に200 OKメッセージを受信した場合、適切にアップロードされたと想定します。このフラグを設定すると、rcloneはPUTでオブジェクトをアップロードした後、正しくアップロードされたと想定します。
      
      特に次の場合、rcloneは次を想定します：
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプを含む）がアップロード時と同じであった
      - サイズがアップロード時と同じであった
      
      単一パートのPUTの応答から、次の項目を読み取ります：
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明なソースオブジェクトがアップロードされている場合、rclone **は** HEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗の機会が増えます（特に誤ったサイズ）。通常の操作では推奨されませんが、実際にはこのフラグを使用しても、アップロードの失敗が検出されない確率は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前にHEADを実行しません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールのフラッシュ間隔。
      
      バッファが必要なアップロード（f.eマルチパート）では、割り当て用にメモリプールが使用されます。
      このオプションは、使用されなくなったバッファーがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2の間に未解決の問題があります。 S3バックエンドではHTTP/2がデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS
      S3がCloudFrontネットワークを介してダウンロードされたデータに対してより安価なイーグレスを提供するためにCloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      true、false、またはデフォルト（未設定）にする必要があります。
      

   --use-presigned-request
      単一パートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS
      SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン1.59以前では、署名済みリクエストを使用して単一のパートオブジェクトをアップロードし、このフラグをtrueに設定すると、その機能を再度有効にすることができます。これは特殊な事情やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付 "2006-01-02"、日時 "2006-01-02
      15:04:05"、またはその前の期間（例：「100d」または「1h」）である必要があります。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されません。したがって、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[時間オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      設定する場合、gzipエンコードされたオブジェクトを解凍します。
      
      S3に「Content-Encoding：gzip」が設定されたままオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding：gzip」として受信すると解凍します。つまり、rcloneはサイズとハッシュをチェックできなくなりますが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際にはオブジェクトを変更しません。 `Content-Encoding：gzip`がアップロード時に設定されなかった場合、ダウンロード時にも設定されません。
      
      ただし、一部のプロバイダ（例：Cloudflare）は、`Content-Encoding：gzip`がアップロード時に設定されていなくてもオブジェクトをgzip圧縮する場合があります。
      
      この場合、Content-Encoding：gzipが設定され、チャンク転送エンコードがある場合に、rcloneはオブジェクトを逐次的に解凍します。
      
      unsetに設定されている場合（デフォルト設定）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここではrcloneの選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value        AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                  オブジェクトの作成時に使用するCanned ACL。 [$ACL]
   --endpoint value             RackCorp オブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --env-auth                   AWSの認証情報を実行時に取得します（環境変数または環境依存のメタデータ）。（デフォルト：false） [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  バケットが配置され、データが保存される場所を指定します。 [$LOCATION_CONSTRAINT]
   --region value               バケットが作成され、データが保存される場所を指定します。 [$REGION]
   --secret-access-key value    AWSの秘密アクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用するCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。（デフォルト："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。（デフォルト："4.656Gi"） [$COPY_CUTOFF]
   --decompress                     設定する場合、gzipエンコードされたオブジェクトを解凍します。（デフォルト：false） [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。（デフォルト：false） [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。（デフォルト：false） [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト："Slash,InvalidUtf8,Dot"） [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合、仮想ホストスタイルを使用します。（デフォルト：true） [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。（デフォルト：1000） [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset（デフォルト："unset"） [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）（デフォルト：0） [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パーツ数。（デフォルト：10000） [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。（デフォルト："1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト：false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。（デフォルト："unset"） [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しません。（デフォルト：false） [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトのHEADを行い、整合性をチェックしません。（デフォルト：false） [$NO_HEAD]
   --no-head-object                 HEADを実行せずにオブジェクトを取得しません。（デフォルト：false） [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します（デフォルト：false） [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。（デフォルト：4） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるための閾値。（デフォルト："200Mi"） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか（デフォルト："unset"） [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか（デフォルト：false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します（デフォルト：false） [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します（デフォルト："off"） [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか（デフォルト：false） [$VERSIONS]

```
{% endcode %}