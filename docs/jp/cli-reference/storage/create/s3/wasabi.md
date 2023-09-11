# Wasabiオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 wasabi - Wasabiオブジェクトストレージ

使い方：
   singularity storage create s3 wasabi [コマンドのオプション] [引数...]

説明：
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数またはenv vars、またはIAMによるEC2/ECSメタデータ）。
      
      access_key_idとsecret_access_keyが空白の場合に適用されます。

      例：
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーIDです。
      
      匿名アクセスまたは実行時の認証情報を使用する場合は空白のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。
      
      匿名アクセスまたは実行時の認証情報を使用する場合は空白のままにします。

   --region
      接続するリージョンです。
      
      S3のクローンを使用していてリージョンが不要な場合は空白のままにします。

      例：
         | <未設定>            | 迷った場合はこのオプションを使用します。
         |                    | v4署名と空のリージョンを使用します。
         | other-v2-signature | v4署名が機能しない場合にのみ使用します。
         |                    | 例：Jewel/v10 CEPHよりも前のバージョン。

   --endpoint
      S3 APIのエンドポイントです。
      
      S3のクローンを使用している場合は必須です。

      例：
         | s3.wasabisys.com                | Wasabi US East 1（北バージニア）
         | s3.us-east-2.wasabisys.com      | Wasabi US East 2（北バージニア）
         | s3.us-central-1.wasabisys.com   | Wasabi US Central 1（テキサス）
         | s3.us-west-1.wasabisys.com      | Wasabi US West 1（オレゴン）
         | s3.ca-central-1.wasabisys.com   | Wasabi CA Central 1（トロント）
         | s3.eu-central-1.wasabisys.com   | Wasabi EU Central 1（アムステルダム）
         | s3.eu-central-2.wasabisys.com   | Wasabi EU Central 2（フランクフルト）
         | s3.eu-west-1.wasabisys.com      | Wasabi EU West 1（ロンドン）
         | s3.eu-west-2.wasabisys.com      | Wasabi EU West 2（パリ）
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP Northeast 1（東京）のエンドポイント
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP Northeast 2（大阪）のエンドポイント
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP Southeast 1（シンガポール）
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP Southeast 2（シドニー）

   --location-constraint
      リージョンと一致するように設定されたロケーション制約です。
      
      不確かな場合は空白のままにします。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成とオブジェクトの保存またはコピー時に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合もバケットの作成に使用されます。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      ソースからのサーバーサイドコピーでは、ソースからACLをコピーせずに新しく書き込みます。
      
      aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      バケットの作成時のみこのACLが適用されます。設定されていない場合はaclが代わりに使用されます。
      
      aclおよびbucket_aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

      例：
         | private            | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループには読み取り権限が与えられます。
         | public-read-write  | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループには読み取りおよび書き込み権限が与えられます。
         |                    | バケット上でこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーにはFULL_CONTROL権限が与えられます。
         |                    | AuthenticatedUsersグループには読み取り権限が与えられます。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフです。
      
      このサイズよりも大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小は0で、最大は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのアップロードや「rclone mount」またはGoogleフォトやGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートのアップロードとしてアップロードされます。
      
      注意：「--s3-upload-concurrency」は、このチャンクサイズのチャンクが個々の転送ごとにメモリにバッファリングされます。
      
      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送が高速化します。
      
      rcloneは、10,000のチャンク制限を下回るため、既知のサイズの大きなファイルをアップロードする場合は、自動的にチャンクサイズを増やします。
      
      サイズの不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBで、最大10,000のチャンクがあるため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、「-P」フラグで表示される進行状況統計の正確さが低下することに注意してください。Rcloneは、AWS SDKによってバッファリングされたときにチャンクが送信されたと判断し、実際にはまだアップロード中である場合もあります。チャンクサイズが大きいほど、AWS SDKのバッファサイズと進捗報告が真実から逸脱するためです。
      

   --max-upload-parts
      マルチパートアップロード内のパートの最大数です。
      
      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。
      
      AWS S3の10,000のチャンク仕様をサポートしていない場合に便利です。
      
      rcloneは、既知のサイズの大きなファイルをアップロードする場合、このチャンクサイズを増やして、このチャンク数の制限を超えないようにします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフです。
      
      サーバーサイドでコピーする必要のあるこのサイズよりも大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小は0で、最大は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまでに長い遅延が発生することがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合にデフォルトになります。
      

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの同時実行数です。
      
      同時にアップロードされる同じファイルのチャンク数です。
      
      高速リンクを介して大量の大きなファイルを転送し、これらの転送が帯域幅を十分に利用していない場合、これを増やすと転送が高速化するかもしれません。

   --force-path-style
      trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。
      falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント]（https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro）を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用します。例：Jewel/v10 CEPHよりも前のバージョン。

   --list-chunk
      リストチャンクのサイズ（各ListObject S3リクエストごとのレスポンスリスト）です。
      
      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」または「page-size」とも呼ばれます。
      大多数のサービスは、リクエストされたオブジェクトが1000個を超えていても、応答リストを切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）です。
      
      S3が最初にリリースされた時、バケット内のオブジェクトを列挙するためのListObjectsが提供されていました。
      
      ただし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高速であり、できるだけ使用する必要があります。
      
      デフォルトの0に設定すると、rcloneはプロバイダでどのリストオブジェクトのメソッドを呼び出すかを推測します。推測が誤っている場合は、このオプションで手動で設定できます。
      

   --list-url-encode
      リストのURLエンコードの有効/無効を設定します：true/false/unset
      
      一部のプロバイダは、名前に制御文字を含める場合にURLエンコードリストをサポートし、使用できる場合はこれが使用されます。unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従い、適用する内容を選択します。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しようとしません。

      バケットが既に存在する場合、トランザクション数を最小限に抑えるときに便利です。
      
      バケット作成の権限を持っていない場合にも必要になる場合があります。v1.52.0より前のバージョンでは、これはバグのために無条件に渡されていました。
      

   --no-head
      アップロード済みのオブジェクトの完全性を確認するためにHEADリクエストを行わないようにします。

      rcloneがPUTでオブジェクトをアップロードした後に200 OKメッセージを受信した場合、正常にアップロードされたものと見なします。
      
      特に、次のものと見なします：
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）は、アップロード時と同じであったこと
      - サイズはアップロード時と同じであったこと
      
      単一パートPUTの応答から読み取られる項目は以下のとおりです：
      
      - MD5SUM
      - アップロード日
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズの不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が高くなります。特に、正確なサイズが不正な場合なので、通常の操作にはお勧めしません。実際には、このフラグを使用しても、アップロードの失敗が検出されない確率は非常に低いです。
      

   --no-head-object
      GETする前にHEADを行わないようにします。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要セクションのエンコーディング]（/overview/#encoding）を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする時間間隔です。
      
      追加バッファが必要なアップロードにはメモリプールを使用します（マルチパートなど）。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2には未解決の問題があります。S3バックエンドではHTTP/2がデフォルトで有効になっていますが、ここで無効にできます。問題が解決されたら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのためのカスタムエンドポイントです。
      これは通常、AWS S3がクラウドフロントCDNのURLに設定されているため、
      クラウドフロントネットワークを介してダウンロードされたデータに対してAWS S3がより安価なエグレスを提供します。

   --use-multipart-etag
      マルチパートアップロードの確認にETagを使用するかどうか
      
      これは、true、false、またはデフォルトの場合のいずれかを指定します。
      

   --use-presigned-request
      シングルパートアップロードに署名付きリクエストまたはPutObjectを使用するかどうか
      
      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョンが1.59未満の場合、署名付きリクエストを使用してシングルパートオブジェクトをアップロードし、このフラグをtrueに設定すると、その機能が再度有効になります。これは、特殊な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはそれより以前の期間、例えば "100d" または "1h" でなければなりません。
      
      これを使用すると、ファイルの書き込み操作は許可されないため、ファイルをアップロードしたり削除したりすることはできません。
      
      有効なフォーマットについては、[時間オプションのドキュメント]（/docs/#time-option）を参照してください。
      

   --decompress
      このフラグが設定されている場合、gzipでエンコードされたオブジェクトを展開します。
      
      "Content-Encoding: gzip"が設定された状態でオブジェクトをS3にアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信したオブジェクトを"Content-Encoding: gzip"で展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzipに圧縮する場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされるときにはオブジェクトを変更しません。オブジェクトが 'Content-Encoding: gzip' でアップロードされていない場合、ダウンロード時には設定されません。
      
      ただし、一部のプロバイダは、 'Content-Encoding: gzip' が設定されていないにもかかわらずオブジェクトを圧縮する場合があります（たとえば、Cloudflare）。
      
      これが設定されており、rcloneが 'Content-Encoding: gzip' が設定され、チャンク送信エンコードの場合、rcloneはオブジェクトを逐次圧縮します。
      
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従い、適用する内容を選択します。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


オプション：
   --access-key-id value        AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成とオブジェクトの保存またはコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value             S3 APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します（環境変数またはenv vars、またはIAMによるEC2/ECSメタデータ）。 (デフォルト：false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンと一致するように設定されたロケーション制約です。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフです。 (デフォルト："4.656Gi") [$COPY_CUTOFF]
   --decompress                     このフラグが設定されている場合、gzipでエンコードされたオブジェクトを展開します。 (デフォルト：false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト：false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト：false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。 (デフォルト：true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズ（各ListObject S3リクエストごとのレスポンスリスト）。 (デフォルト：1000) [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードの有効/無効を設定します。 (デフォルト："unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。 (デフォルト：0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロード内のパートの最大数。 (デフォルト：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする時間間隔。 (デフォルト："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipに圧縮する場合に設定します。 (デフォルト："unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しようとしません。 (デフォルト：false) [$NO_CHECK_BUCKET]
   --no-head                        アップロード済みのオブジェクトの完全性を確認するためにHEADリクエストを行わないようにします。 (デフォルト：false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを行わないようにします。 (デフォルト：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト：false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ。 (デフォルト："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードの確認にETagを使用するかどうか (デフォルト："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名付きリクエストまたはPutObjectを使用するかどうか (デフォルト：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (デフォルト：false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (デフォルト："off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。 (デフォルト：false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}