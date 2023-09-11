# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 lyvecloud - Seagate Lyve Cloud

USAGE:
   singularity storage create s3 lyvecloud [コマンドオプション] [引数...]

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはランタイムクレデンシャルがない場合はEC2/ECSメタデータ）。
      
      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境変数（env varsまたはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーIDです。
      
      無名アクセスまたはランタイムクレデンシャルの場合は空白のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。
      
      無名アクセスまたはランタイムクレデンシャルの場合は空白のままにします。

   --region
      接続するリージョンです。
      
      S3のクローンを使用しており、リージョンがない場合は空白のままにします。

      例:
         | <unset>            | よくわからない場合はこのままでOKです。
         |                    | v4の署名と、空のリージョンを使用します。
         | other-v2-signature | v4の署名が機能しない場合のみ使用します。
         |                    | 例：Jewel/v10 CEPH以前。

   --endpoint
      S3 APIのエンドポイントです。
      
      S3クローンを使用している場合は必須です。

      例:
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US East 1（バージニア）
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US West 1（カリフォルニア）
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1（シンガポール）

   --location-constraint
      リージョンと一致するように場所の制限を設定します。
      
      よくわからない場合は空白のままにします。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合にはバケットの作成時にも使用されます。
      
      詳細については、[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3のサーバーサイドでオブジェクトをコピーする際には、このACLが適用されます。
      S3はソースからACLをコピーするのではなく、新しいACLを書き込みます。
      
      aclが空の文字列の場合はX-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用するCanned ACLです。
      
      詳細については、[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucket_aclが設定されていない場合、バケットの作成時のみこのACLが適用されます。
      
      aclとbucket_aclが空の文字列の場合はX-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

      例:
         | private            | オーナーがFULL_CONTROLを取得します。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーがFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADアクセス権があります。
         | public-read-write  | オーナーがFULL_CONTROLを取得します。
         |                    | AllUsersグループにはREADおよびWRITEアクセス権があります。
         |                    | バケットでこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーがFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループにはREADアクセス権があります。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフです。
      
      これを超えるファイルは、chunk_sizeの大きさのチャンクでアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffを超えるファイルやサイズの不明なファイル（たとえば「rclone rcat」でアップロードされたファイルや「rclone mount」やGoogleフォトやGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意：「--s3-upload-concurrency」のチャンクは、転送ごとにメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを増やすと転送が高速化されます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際に、10,000チャンクの制限以下にするために、自動的にチャンクサイズを増やします。
      
      サイズの不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBであり、最大10,000個のチャンクがあるため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合には、chunk_sizeを増やす必要があります。
      
      chunkサイズを増やすと、「-P」フラグで表示される進行状況統計の正確性が低下します。Rcloneは、chunkがAWS SDKによってバッファリングされたときに、chunkを送信したとみなしますが、実際にはまだアップロード中かもしれません。chunkサイズが大きいと、AWS SDKのバッファサイズが大きくなり、進行状況の報告が真実から逸脱する可能性があります。
      

   --max-upload-parts
      マルチパートアップロードでの最大パート数です。
      
      このオプションは、マルチパートアップロードを行う場合に使用する最大のマルチ部分チャンク数を定義します。
      
      これは、サービスがAWS S3の10,000チャンクの仕様をサポートしていない場合に便利です。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際に、[このチャンク数の制限を下回るように]自動的にチャンクサイズを増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフです。
      
      サーバーサイドでコピーする必要のあるこれを超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始するまで時間がかかる場合があります。

   --shared-credentials-file
      共有クレデンシャルファイルへのパスです。
      
      env_auth = trueの場合、rcloneは共有クレデンシャルファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有クレデンシャルファイルで使用するプロファイルです。
      
      env_auth = trueの場合、rcloneは共有クレデンシャルファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」がデフォルト値になります。
      

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行数です。
      
      同じファイルのチャンクを同時にアップロードする数です。
      
      高速リンクで大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用しない場合は、これを増やして転送を高速化できます。

   --force-path-style
      trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、この値をfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      falseの場合（デフォルト）、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用してください。たとえば、Jewel/v10 CEPH以前。

   --list-chunk
      リストリングのチャンクサイズ（各ListObject S3リクエストごとのレスポンスリストのサイズ）です。
      
      このオプションは、AWS S3の仕様における「MaxKeys」、「max-items」、「page-size」とも呼ばれます。
      ほとんどのサービスは、要求されたオブジェクト数が1000を超えてもリストを切り捨てます。
      AWS S3では、これはグローバルな制限であり、変更することはできません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）です。
      
      S3が最初に開始したとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しのみが提供されました。
      
      しかし、2016年5月にListObjectsV2が導入されました。これははるかに高いパフォーマンスですので、可能な限り使用するべきです。
      
      デフォルトで設定されている0の場合、rcloneはプロバイダの設定に従ってどのリストオブジェクトメソッドを呼び出すかを推測します。間違った推測をした場合、ここで手動で設定することもできます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダは、リストをURLエンコードすることをサポートし、コントロール文字をファイル名で使用する際にはこれがより信頼性があります。これが未設定（デフォルト）になっている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択をオーバーライドできます。
      

   --no-check-bucket
      バケットの存在を確認したり作成したりしようとしない場合に設定します。
      
      バケットが既に存在することを知っている場合、rcloneのトランザクションの数を最小限に抑える必要がある場合に役立ちます。
      
      ユーザーがバケットの作成権限を持っていない場合も必要です。v1.52.0より前では、このバグのためにサイレントにパスされました。
      

   --no-head
      アップロード済みのオブジェクトのHEADリクエストを行って整合性をチェックしません。
      
      rcloneがPUTでオブジェクトをアップロードした後、200 OKメッセージを受け取った場合、正しくアップロードされたとみなします。
      
      特に、次のことを前提とします：
      
      - アップロード時のメタデータ（modtime、保存クラス、コンテンツタイプ）がアップロードと同じだったこと
      - サイズがアップロード時と同じかどうか
      
      以下の項目を、単一の部分PUTのレスポンスから読み取ります。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズの不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの障害が検出される可能性が増え、特にサイズが正しくない場合など、通常の操作では推奨されません。実際には、このフラグを使用しても、アップロードの障害が検出される可能性は非常に低いです。
      

   --no-head-object
      GETの前にHEADを行わないようにします。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。
      
      バッファが必要なアップロード（たとえばマルチパート）は、割り当てにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2に関連する解決できない問題があります。デフォルトでは、s3バックエンドではHTTP/2が有効になっていますが、ここで無効にすることもできます。問題が解決した場合、このフラグは削除されます。
      
      参照:[https://github.com/rclone/rclone/issues/4673](https://github.com/rclone/rclone/issues/4673), [https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)
      

   --download-url
      ダウンロードのカスタムエンドポイントです。
      これは通常、AWS S3がCloudFrontネットワークを介してダウンロードされたデータに対してエグレスが安価であるため、CloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルト値（プロバイダによる）を使用する必要があります。
      

   --use-presigned-request
      シングルパートアップロードに対して署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはオブジェクトをアップロードするためにAWS SDKのPutObjectを使用します。
      
      rcloneのバージョン1.59未満では、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは例外的な場合やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは日付、「2006-01-02」、日時「2006-01-02 15:04:05」、そのように昔からの時間の期間「100d」や「1h」です。
      
      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[時間オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      これを設定すると、gzipでエンコードされたオブジェクトを解凍します。
      
      S3へのアップロード時に「Content-Encoding: gzip」が設定されたオブジェクトを、通常は圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは「Content-Encoding: gzip」で受け取ったファイルを解凍します。つまり、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドでオブジェクトをgzipで圧縮する可能性がある場合にこれを設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。ファイルが`Content-Encoding: gzip`でアップロードされていない場合は、ダウンロード時にも設定されません。
      
      しかし、一部のプロバイダ（例：Cloudflare）は、`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzipで圧縮することがあります。
      
      これには、次のようなエラーが表示されることがあります：
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneがContent-Encoding: gzipで設定され、チャンク転送エンコーディングを使用してオブジェクトをダウンロードした場合、rcloneはオブジェクトをリアルタイムで解凍します。
      
      これが未設定（デフォルト）にされている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択をオーバーライドできます。
      

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制します


OPTIONS:
   --access-key-id value        AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value             S3 APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                   ランタイムからAWSの認証情報を取得します（環境変数またはランタイムクレデンシャルがない場合はEC2/ECSメタデータ）。 (デフォルト値: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示する
   --location-constraint value  リージョンと一致するように場所の制限を設定します。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用するCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト値: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフです。 (デフォルト値: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzipでエンコードされたオブジェクトを解凍する場合に設定します。 (デフォルト値: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト値: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト値: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト値: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。 (デフォルト値: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストリングのチャンクサイズです。 (デフォルト値: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか。 (デフォルト値: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。 (デフォルト値: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パート数です。 (デフォルト値: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度です。 (デフォルト値: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト値: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドでオブジェクトをgzipで圧縮する可能性がある場合に設定します。 (デフォルト値: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認したり作成したりしようとしない場合に設定します。 (デフォルト値: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロード済みのオブジェクトのHEADリクエストを行って整合性をチェックしません。 (デフォルト値: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わないようにします。 (デフォルト値: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制します (デフォルト値: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有クレデンシャルファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSのセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有クレデンシャルファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行数です。 (デフォルト値: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフです。 (デフォルト値: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか。 (デフォルト値: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに対して署名済みリクエストまたはPutObjectを使用するかどうか。 (デフォルト値: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (デフォルト値: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (デフォルト値: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (デフォルト値: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}