# Huawei Object Storage Service

```
NAME:
   singularity storage create s3 huaweiobs - Huawei Object Storage Service

USAGE:
   singularity storage create s3 huaweiobs [command options] [arguments...]

DESCRIPTION:
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ）。

      access_key_idとsecret_access_keyが空の場合に適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーIDです。

      匿名アクセスまたは実行時の認証情報にするには、空のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたは実行時の認証情報にするには、空のままにします。

   --region
      接続する地域です。- バケットが作成され、データが保存される場所です。エンドポイントと同じである必要があります。

      例:
         | af-south-1     | AF-Johannesburg
         | ap-southeast-2 | AP-Bangkok
         | ap-southeast-3 | AP-Singapore
         | cn-east-3      | CN East-Shanghai1
         | cn-east-2      | CN East-Shanghai2
         | cn-north-1     | CN North-Beijing1
         | cn-north-4     | CN North-Beijing4
         | cn-south-1     | CN South-Guangzhou
         | ap-southeast-1 | CN-Hong Kong
         | sa-argentina-1 | LA-Buenos Aires1
         | sa-peru-1      | LA-Lima1
         | na-mexico-1    | LA-Mexico City1
         | sa-chile-1     | LA-Santiago2
         | sa-brazil-1    | LA-Sao Paulo1
         | ru-northwest-2 | RU-Moscow2

   --endpoint
      OBS APIのエンドポイントです。

      例:
         | obs.af-south-1.myhuaweicloud.com     | AF-Johannesburg
         | obs.ap-southeast-2.myhuaweicloud.com | AP-Bangkok
         | obs.ap-southeast-3.myhuaweicloud.com | AP-Singapore
         | obs.cn-east-3.myhuaweicloud.com      | CN East-Shanghai1
         | obs.cn-east-2.myhuaweicloud.com      | CN East-Shanghai2
         | obs.cn-north-1.myhuaweicloud.com     | CN North-Beijing1
         | obs.cn-north-4.myhuaweicloud.com     | CN North-Beijing4
         | obs.cn-south-1.myhuaweicloud.com     | CN South-Guangzhou
         | obs.ap-southeast-1.myhuaweicloud.com | CN-Hong Kong
         | obs.sa-argentina-1.myhuaweicloud.com | LA-Buenos Aires1
         | obs.sa-peru-1.myhuaweicloud.com      | LA-Lima1
         | obs.na-mexico-1.myhuaweicloud.com    | LA-Mexico City1
         | obs.sa-chile-1.myhuaweicloud.com     | LA-Santiago2
         | obs.sa-brazil-1.myhuaweicloud.com    | LA-Sao Paulo1
         | obs.ru-northwest-2.myhuaweicloud.com | RU-Moscow2

   --acl
      オブジェクトを作成し、格納またはコピーする際に使用するCanned ACLです。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。

      詳細については、[Amazon S3開発者ガイドのACLの概要](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      S3では、サーバー側でオブジェクトをコピーする際にACLをコピーするのではなく、新たに作成します。

      aclが空の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットの作成に使用するCanned ACLです。

      詳細については、[Amazon S3開発者ガイドのACLの概要](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      aclが設定されていない場合は、バケットの作成時のみに適用されます。

      aclとbucket_aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループには読み取りアクセス権限があります。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループには読み取りおよび書き込みアクセス権限があります。
         |                    | バケット上でこれを付与することは一般的に推奨されません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループには読み取りアクセス権限があります。

   --upload-cutoff
      チャンク化アップロードに切り替えるためのカットオフです。

      これを超えるサイズのファイルは、チャンクサイズで分割してアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。

      upload_cutoffより大きなサイズのファイル、またはサイズが不明なファイル（「rclone rcat」からのアップロードや「rclone mount」やGoogleフォトまたはGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意："--s3-upload-concurrency"は、このチャンクサイズごとに転送ごとにメモリにバッファリングされます。

      高速リンク上で大容量のファイルを転送し、メモリが十分な場合は、これを増やすと転送が高速化されます。

      Rcloneは、10,000チャンクの制限を下回るように、既知の大きさの大ファイルをアップロードする場合には自動的にチャンクサイズを増加させます。

      不明なサイズのファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10,000チャンクまであることを考慮すると、ストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードするには、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の統計の正確性が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされたときにチャンクが送信されたと扱いますが、実際にはまだアップロード中です。チャンクサイズが大きいほど、AWS SDKのバッファも大きくなり、真実から逸脱した進捗報告が表示されます。

   --max-upload-parts
      マルチパートアップロードでの最大パート数です。

      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。

      10,000パートのAWS S3仕様をサポートしていないサービスに有用です。

      Rcloneは、既知のサイズの大きなファイルをアップロードする場合には、このチャンクサイズをパート数の上限以下に保つために、自動的にチャンクサイズを増加させます。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフです。

      これを超えるサイズのファイルは、このサイズのチャンクでコピーされます。
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算して、オブジェクトのメタデータに追加するため、大きなファイルのアップロード開始には時間がかかります。

   --shared-credentials-file
      共有資格情報ファイルへのパスです。

      env_auth = trueの場合、rcloneは共有資格情報ファイルを使用できます。

      この変数が空の場合、rcloneは環境変数"AWS_SHARED_CREDENTIALS_FILE"を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有資格情報ファイルで使用するプロファイルです。

      env_auth = trueの場合、rcloneは共有資格情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合、環境変数"AWS_PROFILE"または"default"が設定されていない場合のデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行度。

      同一ファイルのチャンクを同時にアップロードします。

      大量の大きなファイルを高速リンクでアップロードし、これらのアップロードが帯域幅を完全に利用しない場合は、これを増やすことで転送を高速化することができます。

   --force-path-style
      trueの場合、パス形式のアクセスを使用し、falseの場合は虚構ホストスタイルを使用します。

      trueの場合（デフォルト）、rcloneはパス形式のアクセスを使用します。falseの場合、rcloneはバーチャルパス形式を使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に応じてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4署名が機能しない場合にのみ使用します（例：Jewel/v10 CEPH以前）。

   --list-chunk
      リスト操作のためのチャンクサイズ（各ListObject S3リクエストのレスポンスリスト）。

      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、「page-size」としても知られています。

      大部分のサービスはリクエスト数が1000を超えてもリストを切り詰めます。

      AWS S3では、これは最大値であり変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。

      Cephでは、この「rgw list buckets max chunk」オプションで増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。

      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは非常に高速であり、可能な限り使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダの設定に応じて呼び出すリストオブジェクトメソッドを推測します。推測が誤っている場合は、ここで手動で設定できます。

   --list-url-encode
      リストのURLエンコードを行うかどうか：true/false/unset。

      一部のプロバイダはリストのURLエンコードをサポートしており、ファイル名に制御文字を使用する場合には、これが利用可能な場合に信頼性が向上します。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に応じて適用される内容を選択します。

   --no-check-bucket
      バケットの存在をチェックせず、作成しません。

      バケットが既に存在する場合にトランザクション数を最小限に抑えるために役立ちます。

      バケット作成の権限がない場合にも必要です。v1.52.0以前では、これはバグのためにサイレントに渡されました。

   --no-head
      HEADリクエストを使用してアップロード済みオブジェクトの整合性を確認しない場合に設定します。

      rcloneは、PUT後に200 OKメッセージを受け取った場合、適切にアップロードされたものとみなします。

      特に、次の項目を想定しています。

      - メタデータ（ModTime、ストレージクラス、コンテンツタイプ）はアップロードと同じであること
      - サイズはアップロードと同じであること

      単一パートのPUTの応答から以下の項目を読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートのアップロードでは、これらの項目は読み取られません。

      サイズ不明のソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まるため、通常の操作には推奨されません。実際には、このフラグでもアップロードの失敗のチャンスは非常に小さいです。

   --no-head-object
      GETする前にHEADを実行しない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度です。

      追加のバッファを必要とするアップロード（マルチパートなど）では、割り当てにメモリプールが使用されます。
      このオプションは、使用済みバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでのmmapバッファの使用有無。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2には未解決の問題があります。S3バックエンドのHTTP/2はデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3は
      CloudFrontネットワークを介してデータをダウンロードする場合に、より安価なエグレスを提供します。

   --use-multipart-etag
      検証のためにマルチパートアップロードでETagを使用するかどうか

      true、false、または未設定にすることができます。デフォルトはプロバイダによる設定になります。

   --use-presigned-request
      シングルパートアップロードに対してpresignされたリクエストまたはPutObjectを使用するかどうか

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン1.59未満では、シングルパートオブジェクトをアップロードするためにpresignedリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、特殊な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時点でのファイルバージョンを表示します。

      パラメータには、日付 "2006-01-02"、日時 "2006-01-02 15:04:05"、またはそれより前の期間、例えば "100d" または "1h" を指定できます。

      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。

      有効なフォーマットについては、[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトを解凍します。

      "Content-Encoding: gzip"が設定された状態でオブジェクトをAWS S3にアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に「Content-Encoding: gzip」で解凍します。そのため、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドによっては、オブジェクトをgzipで圧縮する可能性があるため、これを設定します。

      通常のプロバイダでは、ダウンロード時にオブジェクトを変更しません。オブジェクトが「Content-Encoding: gzip」なしでアップロードされなかった場合は、ダウンロード時に設定されません。

      ただし、一部のプロバイダは、「Content-Encoding: gzip」でなくてもオブジェクトをgzipで圧縮する場合があります（例: Cloudflare）。

      次のようなエラーが表示される場合、これを設定します。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグが設定されており、rcloneがContent-Encoding: gzipが設定され、チャンクの転送エンコーディングの場合、rcloneはオブジェクトをスムーズに解凍します。

      無設定（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用される内容を選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します

OPTIONS:
   --access-key-id value      AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                オブジェクトの作成と格納またはコピー時に使用するCanned ACL。 [$ACL]
   --endpoint value           OBS APIのエンドポイント。 [$ENDPOINT]
   --env-auth                 実行時にAWSの認証情報を取得。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続する地域。バケットが作成され、データが保存される場所です。 [$REGION]
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成に使用するCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5MiB") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (default: "4.656GiB") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのHTTP/2の使用を無効にする。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パス形式のアクセスを使用し、falseの場合は虚構ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リスト操作のためのチャンクサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードを行うかどうか。 (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パート数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでのmmapバッファの使用有無。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドによっては、オブジェクトをgzipで圧縮する可能性があるため、これを設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してアップロード済みオブジェクトの整合性を確認しない場合に設定します。 (default: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを実行しない場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有資格情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有資格情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行度。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのカットオフ。 (default: "200MiB") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       検証のためにマルチパートアップロードでETagを使用するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに対してpresignされたリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（自動生成）
   --path value  ストレージのパス
```