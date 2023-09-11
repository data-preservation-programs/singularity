# StackPath Object Storage

{% code fullWidth="true" %}
```
名前:
   シンギュラリティ ストレージの作成 s3 stackpath - StackPath Object Storage

使用法:
   シンギュラリティ ストレージの作成 s3 stackpath [コマンドオプション] [引数...]

説明:
   --env-auth
      ランタイムからAWSの認証情報を取得します (環境変数または環境変数がない場合はEC2/ECSメタデータから)。
      
      access_key_idとsecret_access_keyが空である場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境からAWSの認証情報を取得します (環境変数またはIAM)。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたはランタイムの認証情報の場合は空にしてください。

   --secret-access-key
      AWSのシークレットアクセスキー (パスワード)。
      
      匿名アクセスまたはランタイムの認証情報の場合は空にしてください。

   --region
      接続するリージョンです。
      
      S3のクローンを使用しリージョンがない場合は空にしてください。

      例:
         | <unset>            | よくわからない場合はこれを使用します。
         |                    | v4 Signatureと空のリージョンが使用されます。
         | other-v2-signature | v4 Signatureが機能しない場合にだけ使用します。
         |                    | 例: Jewel/v10 CEPH以降。

   --endpoint
      StackPath Objec Storageのエンドポイントです。

      例:
         | s3.us-east-2.stackpathstorage.com    | 米国東部のエンドポイント
         | s3.us-west-1.stackpathstorage.com    | 米国西部のエンドポイント
         | s3.eu-central-1.stackpathstorage.com | EUのエンドポイント

   --acl
      バケットの作成やオブジェクトの保存、コピー時に使用するCanned ACLです。
      
      オブジェクトの作成に使用され、bucket_aclが設定されていない場合も作成バケットに使用されます。
      
      詳細については[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      サーバー側でオブジェクトをコピーする際にS3はソースからACLをコピーするのではなく、新しいACLを書き込みます。
      
      aclが空の文字列である場合、X-Amz-Acl:ヘッダは追加されず、デフォルト(プライベート)が使用されます。
      

   --bucket-acl
      バケットの作成時に使用するCanned ACLです。
      
      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      このACLは、バケットを作成する際にのみ適用されます。設定されていない場合は代わりに"acl"が使用されます。
      
      "acl"と"bucket_acl"が空の文字列である場合、X-Amz-Acl:ヘッダは追加されず、デフォルト(プライベート)が使用されます。
      

      例:
         | private            | オーナーにFULL_CONTROL権限が与えられます。
         |                    | 他のユーザーはアクセス権限を持ちません (デフォルト)。
         | public-read        | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADアクセスが与えられます。
         | public-read-write  | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADおよびWRITEアクセスが与えられます。
         |                    | バケット上でこれを許可することは一般的には推奨されません。
         | authenticated-read | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AuthenticatedUsersグループにはREADアクセスが与えられます。

   --upload-cutoff
      チャンクアップロードに切り替えるための最小のファイルサイズです。
      
      これを超えるファイルは、chunk_sizeのサイズのチャンクでアップロードされます。
      最小は0、最大は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクのサイズです。
      
      upload_cutoffを超えるサイズのファイル、またはサイズ不明のファイル（たとえば"rclone rcat"からのアップロード、または"rclone mount"やGoogleフォトまたはGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードでアップロードされます。
      
      注意："--s3-upload-concurrency"このサイズのチャンクは、転送ごとにメモリ内にバッファリングされます。
      
      高速リンクで大容量のファイルを転送しており、十分なメモリを使用している場合、この値を増やすと転送速度が向上します。
      
      Rcloneは、最大10000個のチャンクを超えないように、既知のサイズの大きなファイルをアップロードする場合に自動的にチャンクサイズを増やします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10000個のチャンクであるため、デフォルトでストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルのストリームアップロードを行う場合は、chunk_sizeを増やす必要があります。
      
      チャンクのサイズを増やすと、プログレス統計情報の正確性が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされたときにチャンクを送信したものとみなしますが、実際にはまだアップロードされている可能性があります。チャンクサイズが大きいほど、AWS SDKのバッファーが大きくなり、実際の進行状況の報告とのずれも大きくなります。
      

   --max-upload-parts
      マルチパートアップロードでの最大パート数です。
      
      このオプションは、マルチパートアップロード時に使用するパート数の最大数を定義します。
      
      これは、サービスがAWS S3の10000パートの仕様をサポートしていない場合に便利です。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合に、パート数がこの制限に達しないようにチャンクサイズを自動的に増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるための最小のファイルサイズです。
      
      サーバーサイドでコピーする必要のあるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小は0、最大は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータとともにMD5チェックサムを保存しません。
      
      通常、Rcloneはアップロードする前に入力のMD5チェックサムを計算し、それをオブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には時間がかかることがあります。

   --shared-credentials-file
      共有認証情報ファイルのパスです。
      
      env_auth=trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を参照します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_auth=trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は、環境変数"AWS_PROFILE"または"default"が設定されていない場合にデフォルト値となります。
      

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数です。
      
      これは、同じファイルの複数のチャンクを同時にアップロードします。
      
      高速リンクで小数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を完全に活用していない場合、これを増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用し、falseの場合は仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります - rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      falseに設定すると（デフォルト）、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用してください。例: Jewel/v10 CEPH以前のバージョン。

   --list-chunk
      リストのチャンクサイズです（各ListObject S3リクエストごとの応答リスト）。
      
      このオプションは、AWS S3のMaxKeys、max-items、またはpage-sizeとしても知られています。
      大抵のサービスは、リクエストされたよりも多い場合でも応答リストを1000オブジェクトに切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションで増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン: 1、2、または0（自動）。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しのみ提供されました。
      
      しかし、2016年5月にはListObjectsV2呼び出しが導入されました。これは非常に高速であるため、可能な場合は使用するべきです。
      
      デフォルトの0に設定されている場合、rcloneはプロバイダの設定に従ってどのリストオブジェクトメソッドを使用するかを推測します。正しく推測できない場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストをurlエンコードするかどうか:true/false/unset
      
      いくつかのプロバイダは、リストをURLエンコードでき、ファイル名に制御文字を使用する場合にはこれがより信頼性があります。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従ってどのように適用するかを選択しますが、ここでrcloneの選択をオーバーライドすることができます。
      

   --no-check-bucket
      バケットの存在を確認せず、または作成しません。
      
      これは、バケットが既に存在する場合にrcloneが行うトランザクションの数を最小限に抑えるために便利です。
      
      バケット作成の許可がない場合にも必要です。v1.52.0より前のバージョンでは、これはバグのためにサイレントにパスしていました。
      

   --no-head
      アップロードしたオブジェクトのHEADをチェックして整合性を確認しない場合に設定します。
      
      このオプションを設定すると、rcloneはPUT後に200 OKメッセージを受け取った場合、正しくアップロードされたとみなします。
      
      特に、次のことを前提とします。
      
      - アップロードするファイルのメタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロードと同じであること
      - サイズがアップロードと同じであること
      
      一部のPUTメソッドのレスポンスは次のものがあります。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み込まれません。
      
      サイズが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、正常な運用には推奨されないため、アップロードの失敗を検出する可能性が増えます。実際には、このフラグでアップロードの失敗が検出される可能性は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。
      
      追加のバッファが必要なアップロード（例: マルチパート）は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2に関する解決されていない問題があります。HTTP/2はs3バックエンドのデフォルトで有効になっていますが、ここでは無効にすることができます。問題が解決されると、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3を通じたデータのダウンロードにおいて、より安価な回避経路であるCloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルトを使用するため、プロバイダのデフォルト値が適用されます。
      

   --use-presigned-request
      シングルパートアップロード用の署名済みリクエストまたはPutObjectの使用を指定するかどうか
      
      falseに設定すると、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン< 1.59では、シングルパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定するとその機能を再度有効にします。これは例外的な状況やテスト以外では必要ありません。
      

   --versions
      ファイルの古いバージョンをディレクトリリストに含めます。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付、"2006-01-02"、datetime "2006-01-02 15:04:05"、またはその時間前の期間を指定する必要があります。例えば、"100d"または"1h"です。
      
      これを使用すると、ファイルの書き込み操作は許可されないため、ファイルのアップロードや削除はできません。
      
      有効なフォーマットについては、[時間オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      このフラグが設定されている場合、gzipで圧縮されたオブジェクトを解凍します。
      
      S3に"Content-Encoding: gzip"が設定されている状態でオブジェクトをアップロードすることも可能です。通常、rcloneはこれらのファイルを圧縮オブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に"gzip"で解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際には変更しません。"Content-Encoding: gzip"でアップロードされていないオブジェクトは、ダウンロード時に"Content-Encoding: gzip"が設定されません。
      
      ただし、一部のプロバイダ（Cloudflareなど）は、`Content-Encoding: gzip`でアップロードされていないオブジェクトを圧縮する場合があります。
      
      これを設定すると、rcloneがContent-Encoding: gzipが設定されたオブジェクトとチャンク転送エンコーディングでダウンロードした場合、rcloneはオブジェクトをリアルタイムで解凍します。
      
      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従ってどのように適用するかを選択しますが、ここでrcloneの選択をオーバーライドできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


オプション:
   --access-key-id value      AWSのアクセスキーID。[$ACCESS_KEY_ID]
   --acl value                バケットの作成やオブジェクトの保存、コピー時に使用するCanned ACL。[$ACL]
   --endpoint value           StackPath Objec Storageのエンドポイント。[$ENDPOINT]
   --env-auth                 ランタイムからAWSの認証情報を取得します (環境変数または環境変数がない場合はEC2/ECSメタデータから)。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョン。[$REGION]
   --secret-access-key value  AWSのシークレットアクセスキー (パスワード)。[$SECRET_ACCESS_KEY]

   高度なオプション

   --bucket-acl value               バケットの作成時に使用するCanned ACL。[$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクのサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための最小のファイルサイズ。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     このフラグが設定されている場合、gzipで圧縮されたオブジェクトを解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータとともにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをurlエンコードするかどうか:true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1、2、または0（自動）。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パート数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトのHEADをチェックして整合性を確認しない場合に設定します。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADを行わない場合に設定します。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制 (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWSのセッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルのパス。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるための最小のファイルサイズ。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロード用の署名済みリクエストまたはPutObjectの使用を指定するかどうか。 (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ファイルの古いバージョンをディレクトリリストに含めます。 (デフォルト: false) [$VERSIONS]

   一般

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}