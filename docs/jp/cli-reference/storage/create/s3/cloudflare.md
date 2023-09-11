# Cloudflare R2 ストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 cloudflare - Cloudflare R2 Storage

USAGE:
   singularity storage create s3 cloudflare [コマンドオプション] [引数...]

DESCRIPTION:
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSのメタデータ）。

      access_key_idとsecret_access_keyが空の場合のみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境変数（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。

      ブランクのままにすると匿名アクセスまたは実行時の認証情報が使用されます。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      ブランクのままにすると匿名アクセスまたは実行時の認証情報が使用されます。

   --region
      接続するリージョン。

      例:
         | auto | R2バケットはCloudflareのデータセンターに自動的に分散され、レイテンシが低くなります。

   --endpoint
      S3 APIのエンドポイント。

      S3クローンを使用する場合には必須です。

   --bucket-acl
      バケットを作成する際に使用する定型のACL。

      詳細については https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。

      このACLはバケットを作成する際にのみ適用されます。
      設定されていない場合は「acl」が使用されます。

      「acl」と「bucket_acl」が空文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーにFULL_CONTROL権限が与えられます。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADアクセス権限があります。
         | public-read-write  | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにはREADおよびWRITEアクセス権限があります。
         |                    | 通常、バケットにこれを設定することはお勧めしません。
         | authenticated-read | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AuthenticatedUsersグループにはREADアクセス権限があります。

   --upload-cutoff
      チャンクアップロードに切り替えるための閾値。

      このサイズより大きなファイルはchunk_sizeで分割してアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（例: "rclone rcat"でアップロードされるファイルや"rclone mount"やGoogleフォトやGoogleドキュメントでアップロードされるファイル）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意: "--s3-upload-concurrency"チャンクのこのサイズは、各転送に対してメモリ内にバッファリングされます。

      高速リンクで大きなファイルを転送し、メモリが十分にある場合は、これを増やすと転送速度が向上します。

      大きなファイルの場合、rcloneは10,000個のチャンク制限以下になるように自動的にチャンクサイズを増やします。

      サイズが不明なファイルは構成されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大で10,000個のチャンクが使用できるため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、進行状況の統計情報（"-P"フラグで表示）の精度が低下します。rcloneは、AWS SDKによってバッファリングされた時点でチャンクを送信したとみなし、まだアップロードが行われているかもしれないということになります。より大きなチャンクサイズは、より大きなAWS SDKバッファおよび進行状況の報告が真実から逸脱することを意味します。

   --max-upload-parts
      マルチパートアップロードで使用する最大パート数。

      このオプションは、マルチパートアップロードを行う際に使用するパートの最大数を定義します。

      もし、サービスがAWS S3の10,000チャンクの仕様をサポートしていない場合は便利です。

      rcloneは、既知のサイズの大きなファイルをアップロードする際に、パート数制限の下でチャンクサイズを自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値。

      このサイズを超えるファイルをサーバーサイドでコピーする場合、このサイズのチャンクでコピーされます。

      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加することがあります。これはデータの整合性チェックには便利ですが、大きなファイルをアップロードする場合にはアップロードの開始に長時間かかることがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を見ます。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。この変数は、そのファイルで使用されるプロファイルを制御します。

      空の場合は、環境変数「AWS_PROFILE」または「default」がデフォルトになります。

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクを同時にアップロードする数です。

      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすと転送速度が向上する可能性があります。

   --force-path-style
      trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。

      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4シグネチャが機能しない場合にのみ使用してください（Jewel/v10 CEPH以前など）。

   --list-chunk
      リストチャンクのサイズ（各リストオブジェクトS3リクエストに対する応答リスト）。

      このオプションは、AWS S3の仕様での「MaxKeys」、「max-items」、または「page-size」とも呼ばれます。
      ほとんどのサービスでは、リクエストされた数を超えても応答リストは1000オブジェクトで切り捨てられます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン: 1,2または0は自動。

      S3が最初にローンチされたとき、バケット内のオブジェクトを列挙するためにのみListObjectsコールが提供されました。

      しかし、2016年5月にListObjectsV2コールが導入されました。これは高性能であり、可能な限り使用するべきです。

      デフォルトの設定である0に設定されている場合、rcloneはプロバイダで設定されたListObjectsメソッドに基づいて推測します。間違って推測すると、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか: true/false/unset

      一部のプロバイダは、リストをURLエンコードすることをサポートし、ファイル名に制御文字を使用する場合にはこれがより信頼性があります。設定が「unset」（デフォルト）になっている場合、rcloneは提供者の設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットの存在をチェックしたり作成したりしない場合は、このフラグを設定します。

      バケットが既に存在することを確認する必要がない場合、rcloneが行うトランザクションの数を最小限にするために役立ちます。

      バケット作成権限がない場合にも必要な場合があります。v1.52.0より前のバージョンでは、これはバグのために黙ってパスされました。

   --no-head
      アップロードされたオブジェクトのHEADリクエストを行わない場合は、このフラグを設定します。

      rcloneはPUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、アップロードが正常に行われたとみなします。

      特に、以下の項目を想定します。

      - アップロード時のメタデータ、更新日時、ストレージクラス、コンテンツタイプがアップロード時と同じであったこと
      - サイズがアップロード時と同じであったこと

      一部のパートのPUTに対しては、次の項目をレスポンスから読み取ります。

      - MD5SUM
      - アップロード日時

      マルチパートアップロードの場合、これらの項目は読み込まれません。

      不明な長さのソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを行います。

      このフラグを設定すると、アップロードの失敗が検出される可能性が高くなります。特にサイズが正しくない場合です。そのため、通常の操作では推奨されません。実際には、このフラグを使用しても検出されないアップロードの失敗の可能性は非常に低いです。

   --no-head-object
      GETの前にHEADを行わない場合に設定します。GET（取得）オブジェクトを取得する場合に必要です。

   --encoding
      バックエンドのエンコーディング。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。

      追加バッファを必要とするアップロード（マルチパートなど）は、割り当てのためにメモリバッファプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2には未解決の問題があります。s3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決された場合、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのためのカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してダウンロードされたデータについて安価な出力を提供します。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      true、false、またはデフォルト（プロバイダのデフォルト）を使用します。

   --use-presigned-request
      シングルパートアップロードにプリサインドリクエストまたはPutObjectを使用するかどうか

      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン< 1.59は、シングルパートオブジェクトをアップロードするために署名付きリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定された時間のファイルバージョンを表示します。

      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはその長い時間前の期間、"100d"または"1h"などです。

      このオプションを使用している場合、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。

      有効なフォーマットについては、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定されている場合、gzipでエンコードされたオブジェクトを解凍します。

      S3に「Content-Encoding: gzip」と設定してオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを受信した際に「Content-Encoding: gzip」と共に解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。オブジェクトが「Content-Encoding: gzip」とアップロードされていない場合、ダウンロード時にもそれは設定されません。

      ただし、一部のプロバイダ（Cloudflareなど）は、「Content-Encoding: gzip」とアップロードされていないオブジェクトをgzip化する場合があります。

      これの症状は、次のようなエラーが発生することです。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneが「Content-Encoding: gzip」が設定され、チャンク化された転送エンコードでオブジェクトをダウンロードすると、rcloneはオブジェクトを逐次解凍します。

      unset（デフォルト）に設定されている場合は、rcloneは提供者の設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する

OPTIONS:
   --access-key-id value      AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --endpoint value           S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                 実行時にAWSの認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSのメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョン。 [$REGION]
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットを作成する際に使用する定型のACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定されている場合、gzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用します。falseの場合、仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか。 (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1,2または0は自動。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用する最大パート数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックしたり作成したりしない場合。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADリクエストを行わない場合。 (default: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わない場合。GET（取得）オブジェクトを取得する場合に必要です。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードにプリサインドリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定された時間のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}