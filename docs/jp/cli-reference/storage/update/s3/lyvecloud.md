# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 lyvecloud - Seagate Lyve Cloud

使用方法:
   singularity storage update s3 lyvecloud [command options] <name|id>

概要:
   Seagate Lyve Cloud を使用してAWSのS3を更新します。

詳細:

   --env-auth
      ランタイム（環境変数か環境特有のメタデータ）からAWSの認証情報を取得します。
      
      access_key_idとsecret_access_keyのどちらかが未設定の場合のみ有効です。

      例:
         | false | 次のステップでAWSの認証情報を入力してください。
         | true  | 環境（環境変数 or IAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスかランタイム認証情報を使用する場合はブランクにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスかランタイム認証情報を使用する場合はブランクにしてください。

   --region
      接続するリージョン。
      
      S3クローンを使用しており、リージョンを持っていない場合はブランクにしてください。

      例:
         | <unset>            | よくわからない場合はこれを選択してください。
         |                    | v4署名およびリージョンなしで接続します。
         | other-v2-signature | v4署名が機能しないときのみ使用してください。
         |                    | 例）Jewel/v10 CEPHの前。

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用している場合は必須です。

      例:
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US East 1（バージニア）
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US West 1（カリフォルニア）
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1（シンガポール）

   --location-constraint
      リージョンに一致するように場所の制約を設定します。
      
      よくわからない場合はブランクにしてください。バケット作成時にのみ使用されます。

   --acl
      バケット作成とオブジェクトの保存またはコピー時に使用されるCanned ACL。
      
      このACLはオブジェクトの作成時およびbucket_aclが設定されていない場合に使用されます。
      
      詳細については、[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      注意点として、S3ではオブジェクトをサーバー側でコピーする際にACLがコピーされず新規に書き込まれます。
      
      aclが空の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      
   --bucket-acl
      バケット作成時に使用されるCanned ACL。
      
      詳細については、[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucket_aclが未設定の場合、代わりに"acl"が使用されます。
      
      "acl"と"bucket_acl"が空の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを持つ。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを持つ。
         |                    | AllUsersグループにREADアクセスが与えられます。
         | public-read-write  | オーナーはFULL_CONTROLを持つ。
         |                    | AllUsersグループにREADおよびWRITEアクセスが与えられます。
         |                    | バケットにこれを設定することは一般的にはお勧めされません。
         | authenticated-read | オーナーはFULL_CONTROLを持つ。
         |                    | AuthenticatedUsersグループにREADアクセスが与えられます。

   --upload-cutoff
      チャンク化アップロードに切り替える閾値。
      
      この閾値より大きいファイルは、chunk_size単位でアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffより大きいファイルやサイズが不明なファイル（例：「rclone rcat」で作成されたものや「rclone mount」またはgoogle photosやgoogle docsでアップロードされたもの）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意: "--s3-upload-concurrency"のチャンクごとにこのサイズがメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送し、十分なメモリがある場合は、この値を増やすと転送速度が向上します。
      
      Rcloneは、ファイルサイズ制限10,000チャンクを下回るように、既知の大きなファイルをアップロードするときに自動的にチャンクサイズを増やします。
      
      未知のサイズのファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBで最大で10,000個のチャンクまであるため、デフォルトの場合、ストリーミングアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリーミングアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況の統計表示の精度が低下します。Rcloneは、AWS SDKでバッファリングされたチャンクが送信されたときにチャンクを送信したと見なし、まだアップロードされている場合でも統計情報を表示します。
      大きなチャンクサイズは、AWS SDKのバッファおよび進行状況の報告が真実から逸脱することを意味します。

   --max-upload-parts
      マルチパートアップロード時のパートの最大数。
      
      このオプションは、マルチパートアップロードを行う際のパートの最大数を定義します。
      
      これは、サービスがAWS S3の10,000チャンクの仕様に対応していない場合に便利です。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、このチャンク数制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値。
      
      サーバーサイドでコピーする必要のあるこの閾値を超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しない。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまで時間がかかる場合があります。

   --shared-credentials-file
      共有クレデンシャルファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有クレデンシャルファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有クレデンシャルファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有クレデンシャルファイルを使用できます。この変数はそのファイルで使用されるプロファイルを制御します。
      
      ブランクの場合、環境変数「AWS_PROFILE」または設定されていない場合は「default」がデフォルトになります。

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロード時の同じファイルのチャンクの並列度。
      
      高速リンクで大量の大きなファイルのアップロードを行い、これらのアップロードが帯域幅を完全に活用していない場合は、この値を増やすと転送速度が向上する場合があります。

   --force-path-style
      Trueの場合、パススタイルアクセスを使用します。Falseの場合、仮想ホストスタイルアクセスを使用します。
      
      True（デフォルト）の場合、rcloneはパススタイルアクセスを使用し、Falseの場合は仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）は、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      Trueの場合、v2認証を使用します。
      
      False（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合のみ使用してください。例）Jewel/v10 CEPHでの利用。

   --list-chunk
      リスティングチャンクのサイズ（各ListObject S3リクエストのレスポンスリスト）。
      
      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、または「page-size」としても知られています。
      ほとんどのサービスは、1000個以上をリクエストしても応答リストを切り捨てます。
      AWS S3ではこれはグローバルな最大値であり、変更することはできません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションで増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または0は自動です。
      
      最初にS3が開始したとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを提供し、可能な限り使用する必要があります。
      
      デフォルトの0に設定されている場合、rcloneはリストオブジェクトのメソッドの呼び出し方法をプロバイダの設定に合わせて推測します。予想が外れた場合は、ここで手動で設定することがあります。

   --list-url-encode
      リストのURLエンコードの有無：true/false/unset
      
      一部のプロバイダはリストのURLエンコードをサポートし、このオプションが利用可能な場合、ファイル名に制御文字を含める場合にこれがより信頼性があります。unset（デフォルト）に設定すると、rcloneはプロバイダの設定に従って適用するようになりますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットの存在をチェックせず、作成しません。
      
      バケットが既に存在する場合、トランザクションの数を最小限に抑えようとする場合に有用です。
      
      バケットの作成権限がない場合にも必要です。バージョン1.52.0より前の場合、これはバグのため無視されます。

   --no-head
      アップロードしたオブジェクトのHEADメソッドで整合性をチェックしません。
      
      rcloneは通常、PUTを使用してオブジェクトをアップロードした後、整合性のためにHEADリクエストを送信しています。
      これはデータの整合性を確認するのに大変役立ちますが、大きなファイルをアップロードする場合は長時間遅延することがあります。

   --no-head-object
      GETを実行する前にHEADを行いません。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファ（分割などを必要とするアップロード）は、割り当てにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
     
      現在、s3（特にminio）バックエンドとHTTP/2について未解決の問題があります。HTTP/2はS3バックエンドでデフォルトで有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードに使用するカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してダウンロードされたデータの転送コストが安いため、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを検証に使用するかどうか
      
      これはtrue、false、または未設定のいずれかである必要があります。

   --use-presigned-request
      シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか
      
      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン1.59未満では、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは特殊な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付、「2006-01-02」、日時「2006-01-02 15:04:05」、およびその時からの期間（「100d」や「1h」など）にすることができます。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されません。
      ファイルのアップロードや削除はできません。
      
      有効なフォーマットについては、[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定すると、gzipでエンコードされたオブジェクトを展開します。
      
      S3に「Content-Encoding: gzip」が設定されているオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮オブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信した「Content-Encoding: gzip」のファイルを展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合、これを設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。`Content-Encoding: gzip`でアップロードされなかったオブジェクトはダウンロードされたときにもそれが設定されません。
      
      ただし、一部のプロバイダは、`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzipで圧縮する場合があります（例：Cloudflare）。
      
      これを設定すると、rcloneは`Content-Encoding: gzip`が設定されたチャンク化転送エンコードのオブジェクトを受信した場合、ファイルを逐次展開します。
      
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するようになりますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制します

オプション:
   --access-key-id value        AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                  バケット作成とオブジェクトの保存またはコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value             S3 APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                   ランタイム（環境変数か環境特有のメタデータ）からAWSの認証情報を取得します。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示します。
   --location-constraint value  リージョンに一致するように場所の制約を設定します。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]

   アドバンス

   --bucket-acl value               バケット作成時に使用されるCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値です。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定すると、gzipでエンコードされたオブジェクトを展開します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードに使用するカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               Trueの場合、パススタイルアクセスを使用します。Falseの場合、仮想ホストスタイルアクセスを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リスティングチャンクのサイズです。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-version value             使用するListObjectsのバージョンです。 (デフォルト: 0) [$LIST_VERSION]
   --list-url-encode value          リストのURLエンコードの有無です。 (デフォルト: "unset") [$LIST_URL_ENCODE]
   --max-upload-parts value         マルチパートアップロード時のパートの最大数です。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度です。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうかです。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合、これを設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトのHEADメソッドで整合性をチェックしません。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GETを実行する前にHEADを行いません。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有クレデンシャルファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有クレデンシャルファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロード時の同じファイルのチャンクの並列度です。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替える閾値です。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを検証に使用するかどうかです。 (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうかです。 (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        Trueの場合、v2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (デフォルト: false) [$VERSIONS]

```
{% endcode %}