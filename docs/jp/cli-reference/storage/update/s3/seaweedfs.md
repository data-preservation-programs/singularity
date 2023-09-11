# SeaweedFS S3

{% code fullWidth="true" %}
```
名前：
   singularity storage update s3 seaweedfs - SeaweedFS S3

使用法：
   singularity storage update s3 seaweedfs [コマンドオプション] <名前|ID>

説明：
   --env-auth
      実行時にAWS認証情報を取得します（環境変数または.envファイルがない場合はEC2/ECSメタデータから取得）。
      
      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例：
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報は空白のままにしておいてください。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報は空白のままにしておいてください。

   --region
      接続するリージョンです。
      
      S3クローンを使用し、リージョンが指定されていない場合は空白のままにしてください。

      例：
         | <未設定>                 | 迷った場合はこれを使用します。
         |                         | v4署名と空のリージョンを使用します。
         | other-v2-signature      | v4署名が機能しない場合にのみ使用します。
         |                         | たとえば、ジュエル/v10 CEPH以前のバージョン。

   --endpoint
      S3 APIのエンドポイントです。
      
      S3クローンを使用する場合は必須です。

      例：
         | localhost:8333 | SeaweedFS S3のlocalhost

   --location-constraint
      リージョンに一致する場所の制約です。
      
      わからない場合は空白のままにしておいてください。バケットを作成する場合にのみ使用されます。

   --acl
      バケットの作成、オブジェクトの保存またはコピー時に使用されるACL（アクセス制御リスト）です。
      
      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合にはバケットの作成にも使用されます。
      
      詳細については[Amazon S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3では、サーバーサイドでオブジェクトをコピーする際にACLがコピーされるのではなく、新しく書き込まれることに注意してください。
      
      もしaclが空の文字列の場合、X-Amz-Acl: header は追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成時に使用されるACL（アクセス制御リスト）です。
      
      詳細については[Amazon S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucket_aclが未設定の場合、"acl"が代わりに使用されます。
      
      もし "acl" と "bucket_acl" が空の文字列の場合、X-Amz-Acl:header は追加されず、デフォルト（private）が使用されます。

      例：
         | private            | オーナーにはFULL_CONTROLが与えられます。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROLが与えられます。
         |                    | AllUsersグループにはREAD権限が与えられます。
         | public-read-write  | オーナーにはFULL_CONTROLが与えられます。
         |                    | AllUsersグループにはREADおよびWRITE権限が与えられます。
         |                    | バケットにこれを付与することは一般的には推奨されません。
         | authenticated-read | オーナーにはFULL_CONTROLが与えられます。
         |                    | 認証済みユーザーグループにはREAD権限が与えられます。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフです。
      
      このサイズより大きいファイルは、chunk_sizeで指定されたサイズのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffより大きいファイルや、サイズが不明なファイル（"rclone rcat"によるアップロードや"rclone mount"、Google Photos、Google Docsでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意："--s3-upload-concurrency"のチャンクは、各転送ごとにメモリ上にバッファリングされます。
      
      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、この値を増やすことで転送速度を向上させることができます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合に、チャンクサイズを自動的に増やし、10,000チャンクの制限を下回るようにします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBで、最大で10,000のチャンクがあるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の精度が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされたときにチャンクを送信したとみなしますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKのバッファサイズも大きくなり、進行状況の報告が真実とは異なる場合があります。

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数です。
      
      このオプションは、マルチパートアップロードを行う際に使用するマルチパートチャンクの最大数を定義します。
      
      サービスがAWS S3の「10,000チャンク」の仕様をサポートしていない場合に役立ちます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合に、チャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフです。
      
      サーバーサイドでコピーする必要があるこのサイズより大きいファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加します。これはデータの整合性チェックに役立ちますが、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合は、デフォルトで現在のユーザーのホームディレクトリを使用します。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空白の場合、環境変数「AWS_PROFILE」または「default」にデフォルトします。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性です。
      
      同じファイルのチャンクを同時にアップロードする数です。
      
      これは、高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合に、転送速度を向上させるのに役立つ場合があります。

   --force-path-style
      Trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパススタイルのアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細は[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、この設定をfalseにする必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      Trueの場合、v2認証を使用します。
      
      これがfalse（デフォルト）に設定されている場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      使用するのはv4署名が機能しない場合のみです。たとえば、前のJewel/v10 CEPH。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストのレスポンスリストのサイズ）です。
      
      このオプションはAWS S3のMaxKeys、max-items、またはpage-sizeとしても知られています。
      ほとんどのサービスでは、リクエストされた以上のオブジェクトをリストするとリストが1000件に切り捨てられます。
      AWS S3では、これはグローバルの最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2または0（自動）。
      
      S3の初期リリースでは、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは非常に高性能であり、可能であれば使用する必要があります。
      
      デフォルトの0に設定されている場合、rcloneはプロバイダの設定に基づいてListObjectsメソッドを呼び出すことを推測します。誤った推測をした場合は、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を含む場合にURLエンコードリストをサポートしています。利用可能な場合、これはより信頼性の高い方法です。rcloneはプロバイダの設定に従って適用することをデフォルト（unset）として選択しますが、ここでrcloneの選択をオーバーライドすることができます。

   --no-check-bucket
      バケットの存在を確認したり作成したりしません。

      バケットが既に存在することを知っている場合、rcloneが実行するトランザクションの数を最小限にするため、これを設定すると便利です。
      
      ユーザーにバケット作成権限がない場合にも必要になる場合があります。v1.52.0より前のバージョンでは、バグのために無視されてしまいました。

   --no-head
      アップロード済みのオブジェクトの整合性を確認するためにHEADリクエストを行いません。

      rcloneは、PUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたと仮定します。
      
      特に次の項目を仮定します：
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）はアップロードと同じであったか
      - サイズはアップロードと同じであったか
      
      単一パートのPUTの場合、次の項目をレスポンスから読み取ります：
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      長さがわからないソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、正常な操作には非推奨ですが、アップロードの失敗が検出されない確率が低くなります。実際には、このフラグがあっても、アップロードの失敗が検出される確率は非常に低いです。

   --no-head-object
      GET実行前にHEADリクエストを行いません。
      
   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要セクションのエンコーディング](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      インターナルメモリバッファプールをフラッシュする頻度です。
      
      追加のバッファが必要なアップロード（マルチパートなど）は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファがどの頻度でプールから削除されるかを制御します。

   --memory-pool-use-mmap
      インターナルメモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドに対してhttp2の使用を無効にします。
      
      s3（特にminio）バックエンドとHTTP/2に関する問題が現在未解決です。S3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワーク経由でダウンロードされるデータの出口が安価です。

   --use-multipart-etag
      バリデーションのためにマルチパートアップロードでETagを使用するかどうか
      
      true、false、またはデフォルトではない場合は、プロバイダのデフォルトを使用します。

   --use-presigned-request
      シングルパートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか
      
      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン < 1.59では、シングルパートのオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンも含めるかどうか。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは、日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはその期間、例えば「100d」や「1h」のいずれかである必要があります。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。
      
      有効な形式については[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      このフラグが設定されている場合、gzipでエンコードされたオブジェクトを展開します。
      
      S3に「Content-Encoding: gzip」が設定されたままオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding: gzip」で受け取るときに解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされる際には変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトは、ダウンロード時にそれが設定されません。
      
      ただし、一部のプロバイダは、gzipが設定されていないにもかかわらずオブジェクトをgzip圧縮する場合があります（例：Cloudflare）。
      
      これによる症状は、次のようなエラーメッセージの受け取りです：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneがContent-Encoding: gzipが設定されたオブジェクトとチャンク転送エンコーディングでダウンロードすると、rcloneはオブジェクトを動的に解凍します。
      
      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するため、ここでrcloneの選択をオーバーライドすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します。

オプション：
   --access-key-id value        AWSアクセスキーID。[$ACCESS_KEY_ID]
   --acl value                  バケットの作成とオブジェクトの保存またはコピー時に使用されるACL。[$ACL]
   --endpoint value             S3 APIのエンドポイント。[$ENDPOINT]
   --env-auth                   実行時にAWS認証情報を取得します（環境変数または.envファイルがない場合はEC2/ECSメタデータから取得）。（デフォルト：false）[$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  リージョンに一致する場所の制約。[$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。[$REGION]
   --secret-access-key value    AWSシークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]

   高度な設定

   --bucket-acl value               バケットの作成時に使用されるACL。[$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。（デフォルト： "5Mi"）[$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。（デフォルト： "4.656Gi"）[$COPY_CUTOFF]
   --decompress                     このフラグが設定されている場合、gzipでエンコードされたオブジェクトを展開します。（デフォルト：false）[$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。（デフォルト：false）[$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドに対してhttp2の使用を無効にします。（デフォルト：false）[$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト："Slash,InvalidUtf8,Dot"）[$ENCODING]
   --force-path-style               Trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。（デフォルト：true）[$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストのレスポンスリストのサイズ）。（デフォルト：1000）[$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset。（デフォルト："unset"）[$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン： 1、2または0（自動）。（デフォルト：0）[$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 （デフォルト： 10000 ）[$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   インターナルメモリバッファプールをフラッシュする頻度。（デフォルト： "1m0s"）[$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           インターナルメモリプールでmmapバッファを使用するかどうか。（デフォルト：false）[$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。（デフォルト： "unset"）[$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認したり作成したりしません。（デフォルト： false ）[$NO_CHECK_BUCKET]
   --no-head                        アップロード済みのオブジェクトの整合性を確認するためにHEADリクエストを行いません。（デフォルト： false ）[$NO_HEAD]
   --no-head-object                 GET実行前にHEADリクエストを行いません。（デフォルト： false ）[$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します。（デフォルト：false）[$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。（デフォルト： 4 ）[$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ。（デフォルト： "200Mi"）[$UPLOAD_CUTOFF]
   --use-multipart-etag value       バリデーションのためにマルチパートアップロードでETagを使用するかどうか（デフォルト："unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。（デフォルト：false）[$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2認証を使用するかどうか。（デフォルト：false）[$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。（デフォルト："off"）[$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンも含めるかどうか。（デフォルト：false）[$VERSIONS]

```
{% endcode %}