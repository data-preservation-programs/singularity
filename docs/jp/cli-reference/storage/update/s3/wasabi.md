# Wasabiオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 wasabi - Wasabiオブジェクトストレージ

USAGE:
   singularity storage update s3 wasabi [コマンドオプション] <名前|ID>

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはenv varsやIAMのEC2/ECSメタデータが使用可能）。
      
      access_key_idとsecret_access_keyが空の場合のみ有効です。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSアクセスキーIDです。
      
      匿名アクセスまたはランタイムの認証情報を使用する場合は空にします。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）です。
      
      匿名アクセスまたはランタイムの認証情報を使用する場合は空にします。

   --region
      接続するリージョンです。
      
      S3クローンを使用しリージョンを持っていない場合は空にします。

      例:
         | <unset>            | 確認できない場合に使用します。
         |                    | v4署名と空のリージョンを使用します。
         | other-v2-signature | v4署名が機能しない場合にのみ使用します。
         |                    | たとえば、Jewel/v10以前のCEPH。

   --endpoint
      S3 APIのエンドポイントです。
      
      S3クローンを使用している場合は必須です。

      例:
         | s3.wasabisys.com                | Wasabi米国東部1（N. バージニア）
         | s3.us-east-2.wasabisys.com      | Wasabi米国東部2（N. バージニア）
         | s3.us-central-1.wasabisys.com   | Wasabi米国中部1（テキサス）
         | s3.us-west-1.wasabisys.com      | Wasabi米国西部1（オレゴン）
         | s3.ca-central-1.wasabisys.com   | Wasabi CA中央1（トロント）
         | s3.eu-central-1.wasabisys.com   | Wasabi EU中央1（アムステルダム）
         | s3.eu-central-2.wasabisys.com   | Wasabi EU中央2（フランクフルト）
         | s3.eu-west-1.wasabisys.com      | Wasabi EU西1（ロンドン）
         | s3.eu-west-2.wasabisys.com      | Wasabi EU西2（パリ）
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP東北1（東京）エンドポイント
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP東北2（大阪）エンドポイント
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP東南1（シンガポール）
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP東南2（シドニー）

   --location-constraint
      リージョンと一致するように設定する場所の制約です。
      
      よくわからない場合は空にします。バケット作成時のみ使用されます。

   --acl
      バケットの作成およびオブジェクトの保存やコピー時に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成時と、bucket_aclが設定されていない場合に使用されます。
      
      詳細については[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3では、サーバー側でオブジェクトをコピーする際、ACLをコピーせずに新たに作成します。
      
      aclが空の文字列の場合、「X-Amz-Acl:」ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。
      
      詳細については[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucketを作成するときのみ、このACLが適用されます。設定されていない場合は「acl」が代わりに使用されます。
      
      「acl」と「bucket_acl」が空の文字列である場合、「X-Amz-Acl:」ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

      例:
         | private            | オーナーにFULL_CONTROLが与えられます。
         |                    | 他にはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADアクセスが与えられます。
         | public-read-write  | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADおよびWRITEアクセスが与えられます。
         |                    | バケットにこれを許可することは一般的にはお勧めできません。
         | authenticated-read | オーナーにFULL_CONTROLが与えられます。
         |                    | AuthenticatedUsersグループにREADアクセスが与えられます。

   --upload-cutoff
      チャンク化アップロードに切り替えるためのキャップです。
      
      これより大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffより大きなファイルまたはサイズが不明なファイル（例：「rclone rcat」からのアップロードや「rclone mount」、google photosやgoogle docsからのアップロード）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      なお、「--s3-upload-concurrency」はこのチャンクサイズの数だけ転送毎にメモリにバッファします。
      
      高速リンク上で大きなファイルを転送してメモリが十分にある場合、チャンクサイズを増やすことで転送速度を向上させることができます。
      
      rcloneは、既知の大きさの大きなファイルをアップロードする場合はチャンクサイズを自動的に増やして、10000チャンクの制限に抑えます。
      
      未知のサイズのファイルは設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBで最大10,000のチャンクがあるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。さらに大きなファイルをストリームアップロードする場合、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、報告される進捗の統計情報の精度が低下します。「-P」フラグとともに使用されるchunkは、AWS SDKによってバッファリングされるまで送信されたとみなされ、実際にはまだアップロードが行われる可能性があります。
      

   --max-upload-parts
      マルチパートアップロードにおけるパートの最大数です。
      
      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。
      
      10,000のパートがサポートしていないサービスの場合に使用することができます。
      
      rcloneは、既知の大きさの大きなファイルをアップロードする場合、パート数の制限内に収まるように自動的にチャンクサイズを増やすことがあります。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのキャップです。
      
      サーバーサイドでコピーする必要があるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneは、アップロードする前に入力のMD5チェックサムを計算してオブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまでに長い遅延が発生する場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルト値となります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合にデフォルト値となります。
      

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロード時の同時実行数です。
      
      同じファイルのチャンクを同時にアップロードします。
      
      高速リンクで大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に活用しきれない場合、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、このオプションをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。設定した場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用してください。たとえば、Jewel/v10以前のCEPHの場合です。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）です。
      
      このオプションは、AWS S3仕様の"MaxKeys"、"max-items"、または"page-size"としても知られています。
      ほとんどのサービスは、リクエストされたより多くのオブジェクトを要求しても、レスポンスリストを1000オブジェクトまでに切り詰めます。
      AWS S3では、これはグローバルに最大値であり、変更できません。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン: 1、2、または0（自動）。
      
      最初にS3がローンチされたとき、バケット内のオブジェクトを列挙するために「ListObjects」呼び出しされていました。
      
      しかし、2016年5月に「ListObjectsV2」呼び出しが導入されました。これははるかに高速であり、可能な限り使用する必要があります。
      
      デフォルトで設定された0であれば、rcloneはプロバイダに応じてどのリストオブジェクトメソッドを呼び出すかを推測します。正しく推測できない場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストをURLエンコードするかどうか: true/false/unset
      
      いくつかのプロバイダは、リストをURLエンコードサポートしており、これを利用できる場合は、ファイル名に制御文字を使用するときにより信頼性が高くなります。unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。
      

   --no-check-bucket
      バケットの存在を確認せずに作成を試みない場合は、このフラグを設定します。
      
      これは、バケットが既に存在することがわかっている場合に、rcloneが実行するトランザクションの数を最小限に抑えるために使用できます。
      
      バケット作成権限がない場合でも必要な場合があります。v1.52.0より前のバージョンでは、バグのためにこれが無視されました。
      

   --no-head
      HEADリクエストを使用してオブジェクトの整合性を確認しない場合は、このフラグを設定します。
      
      rcloneがオブジェクトをPUTでアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたと仮定します。
      
      特に、次のことを仮定します：
      
      - 元のメタデータ、変更時刻、ストレージクラス、コンテンツタイプはアップロード時と同じであること。
      - サイズはアップロード時と同じであること。
      
      以下の項目を1つのパートのPUTのレスポンスから読み込みます：
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードでは、これらの項目は読み込まれません。
      
      サイズ不明のソースオブジェクトをアップロードする場合は、rcloneはHEADリクエストを行います。
      
      このフラグを設定すると、アップロードの際の検出されなかった失敗の可能性が高まるため、通常の操作にはお勧めできません。このフラグを使用しても、アップロードの失敗が検出されない可能性は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前にHEADを行わない場合は、このフラグを設定します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[エンコーディングの概要](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。
      
      追加のバッファが必要なアップロード（たとえばマルチパート）では、メモリプールを使用して割り当てが行われます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      現在、s3バックエンド（特にminio）とHTTP/2には未解決の問題があります。HTTP/2はデフォルトでS3バックエンドで有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロードのためのカスタムエンドポイントです。
      これは通常、AWS S3は、CloudFrontネットワーク経由でダウンロードされたデータに対してより安価なアウトバウンドトラフィックを提供するために、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロードの検証にETagを使用するかどうか
      
      これはtrue、false、または設定がない場合のいずれかである必要があります。
      

   --use-presigned-request
      シングルパートのアップロードに署名済リクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rclone < 1.59のバージョンでは、シングルパートオブジェクトのアップロードに署名済要求を使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは特殊な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはその前の期間、「100d」または「1h」のようにする必要があります。
      
      このオプションを使用すると、ファイルの書き込み操作は許可されないため、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。
      
      "Content-Encoding: gzip"が設定されたままオブジェクトをAWS S3にアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを"Lcontent-Encoding: gzip"が受信された段階で展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。
      

   --might-gzip
      プロバイダがオブジェクトにgzipを使用する可能性がある場合、これを設定します。
      
      通常、プロバイダは、ダウンロード時にオブジェクトを変更しません。`Content-Encoding:gzip`でアップロードされていない場合、ダウンロード時には設定されません。
      
      ただし、一部のプロバイダ（Cloudflareなど）は、`Content-Encoding:gzip`でアップロードされていなくてもオブジェクトをgzipで圧縮できる場合があります。
      
      これを設定すると、rcloneがContent-Encoding:gzipが設定されたチャンク転送エンコードを使用してオブジェクトをダウンロードすると、rcloneがオブジェクトを逐次的に解凍します。
      
      unset（デフォルト）に設定された場合、rcloneはプロバイダの設定に従って適用するかどうかを選択しますが、ここでrcloneの選択を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value        AWSアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存やコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value             S3 APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                   ランタイムからAWSの認証情報を取得します（環境変数またはenv varsやIAMのEC2/ECSメタデータが使用可能）。（デフォルト：false） [$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  リージョンと一致するように設定する場所の制約です。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト： "5MiB") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのキャップです。 (デフォルト： "4.656GiB") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。 (デフォルト： false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト： false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのHTTP/2の使用を無効にします。 (デフォルト： false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト： "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト： true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）です。 (デフォルト： 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか。 (デフォルト： "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1、2、または0（自動）。 (デフォルト： 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードにおけるパートの最大数。 (デフォルト： 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (デフォルト： "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト： false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               プロバイダがオブジェクトにgzipを使用する可能性がある場合、これを設定します。 (デフォルト： "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せずに作成を試みない場合は、このフラグを設定します。 (デフォルト： false) [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してオブジェクトの整合性を確認しない場合は、このフラグを設定します。 (デフォルト： false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADを行わない場合は、このフラグを設定します。 (デフォルト： false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト： false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロード時の同時実行数。 (デフォルト： 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのキャップ。 (デフォルト： "200MiB") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードの検証にETagを使用するかどうか。 (デフォルト： "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名済リクエストまたはPutObjectを使用するかどうか。 (デフォルト： false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (デフォルト： false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (デフォルト： "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (デフォルト： false) [$VERSIONS]

```
{% endcode %}