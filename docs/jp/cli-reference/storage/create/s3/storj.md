# Storj（S3互換ゲートウェイ）

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 storj - Storj（S3互換ゲートウェイ）

USAGE:
   singularity storage create s3 storj [コマンドオプション] [引数...]

DESCRIPTION:
   --env-auth
      ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ）。

      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力してください。
         | true  | 環境からAWSの認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたはランタイムの認証情報を使用する場合は空にします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイムの認証情報を使用する場合は空にします。

   --endpoint
      Storjゲートウェイのエンドポイント。

      例:
         | gateway.storjshare.io | グローバルホステッドゲートウェイ

   --bucket-acl
      バケット作成時に使用する設定ACLです。

      詳細はhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      このACLはバケット作成時にのみ適用されます。設定されていない場合は "acl" が使用されます。

      "acl" と "bucket_acl" が空の文字列の場合、"X-Amz-Acl" ヘッダは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | 所有者にFULL_CONTROLのアクセス権限があります。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | 所有者にFULL_CONTROLのアクセス権限があります。
         |                    | すべてのユーザーグループにREADアクセス権限があります。
         | public-read-write  | 所有者にFULL_CONTROLのアクセス権限があります。
         |                    | すべてのユーザーグループにREADおよびWRITEアクセス権限があります。
         |                    | バケットにこれを付与することは一般的に推奨されません。
         | authenticated-read | 所有者にFULL_CONTROLのアクセス権限があります。
         |                    | 認証済みユーザーグループにREADアクセス権限があります。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ値です。

      この値を超えるファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクサイズです。

      upload_cutoffを超えるファイルや、サイズが不明なファイル（例：「rclone rcat」からのアップロード、または「rclone mount」またはGoogleフォトまたはGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。

      注意：「--s3-upload-concurrency」のチャンクは、転送ごとにメモリ上にバッファリングされます。

      高速リンクを介して大きなファイルを転送しており、十分なメモリがある場合は、これを増やすことで転送速度を向上させることができます。

      rcloneは、10,000チャンクの制限を下回るように、既知のサイズの大きなファイルをアップロードする場合には、自動的にチャンクサイズを増やします。

      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で10,000個のチャンクができます。したがって、デフォルトでは、ストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグで表示される進行統計の精度が低下します。rcloneは、AWS SDKによってバッファリングされたときにチャンクが送信されたとみなしますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKのバッファと進行状況の報告は真実からかけ離れるようになります。

   --max-upload-parts
      マルチパートアップロードで使用するパーツの最大数です。

      このオプションは、マルチパートアップロード時に使用するマルチパートチャンクの最大数を定義します。

      これは、サービスが10,000チャンクのAWS S3の仕様をサポートしていない場合に便利です。

      rcloneは、既知のサイズの大きなファイルをアップロードする場合、このチャンクサイズを増やして10,000個のチャンク数制限を下回るように自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値です。

      サーバーサイドコピーする必要があるこの値を超えるファイルは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合、カレントユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合は、環境変数 "AWS_PROFILE" または "default" が設定されていません。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクの数を同時にアップロードします。

      高速リンク上で大量の大きなファイルをアップロードし、これらのアップロードによって帯域幅が十分に利用されない場合は、これを増やすと転送速度が向上する場合があります。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。
      falseの場合、rcloneは仮想パススタイルを使用します。詳細は[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、この値をfalseに設定する必要があります - この設定はrcloneはプロバイダの設定に基づいて自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。
      ファイルのバージョンがv4のシグネチャで機能しない場合にのみ、この設定を使用します。

   --list-chunk
      リスト表示のチャンクサイズ（各ListObject S3リクエストごとのレスポンスリスト）。

      このオプションは、AWS S3仕様の "MaxKeys"、"max-items"、または "page-size" とも呼ばれます。
      ほとんどのサービスは、1000オブジェクトを超えても応答リストを切り詰めます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、オプション "rgw list buckets max chunk" でこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動的に変換するための0。

      S3が最初に登場したとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスで、可能な場合は使用する必要があります。

      デフォルト値の0の場合、rcloneはプロバイダの設定に従って呼び出すListObjectsメソッドを推測します。推測が間違っている場合は、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset。

      一部のプロバイダは、ファイル名に制御文字を使用する場合に利用できるURLエンコードリストをサポートしています。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用する内容を選択しますが、ここでrcloneの選択をオーバーライドすることができます。

   --no-check-bucket
      バケットが存在するかチェックし、または作成しようとしません。

      バケットが既に存在することがわかっている場合、このflagを設定することで、rcloneが行うトランザクションの数を最小限に抑えることができます。

      また、使用するユーザーにバケット作成の権限がない場合にも必要です。v1.52.0以前では、これはバグにより静かにパスされていました。

   --no-head
      HEADリクエストを使用してアップロードしたオブジェクトの整合性をチェックしません。

      rcloneは通常、PUTでオブジェクトをアップロードした後、200 OKメッセージを受信した場合、正常にアップロードされたものと見なします。

      特に、次の内容と想定されます。

      - メタデータ（モディファイ時間、ストレージクラス、コンテンツタイプ）がアップロード時と同じであったこと
      - サイズがアップロード時と同じであったこと

      シングルパートのPUTの応答については、以下の項目を読み取ります。

      - MD5SUM
      - アップロード日時

      マルチパートアップロードの場合、これらの項目は読み取られません。

      長さが不明なソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出される可能性が高まります。特に、サイズが正しくない場合などです。通常の運用にはお勧めしませんが、実際のアップロードの失敗の確率は非常に低いです。

   --no-head-object
      GETでオブジェクトを取得する前にHEADを実行しない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールのフラッシュ頻度です。

      ハイブリッドアップロードにはバッファが追加で必要となり、アロケーションのためにメモリプールが使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、S3（特にminio）バックエンドとHTTP/2の問題が解決されていません。AWS S3バックエンドではHTTP/2がデフォルトで有効になっているが、ここで無効にすることができます。この問題が解決されたら、このフラグは削除されます。

      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのためのカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してダウンロードされるデータのegressコストが安価です。このため、通常、CloudFrontのCDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか。

      これはtrue、false、またはデフォルト（プロバイダに基づく）に設定する必要があります。

   --use-presigned-request
      シングルパートアップロードの場合、PresignedリクエストまたはPutObjectを使用するかどうか。

      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョンが1.59未満の場合、シングルパートオブジェクトをアップロードするためにPresignedリクエストを使用し、このフラグをtrueに設定することで、その機能を再有効化することができます。これは、特別な状況やテストの場合にのみ必要です。

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時点のファイルバージョンを表示します。

      パラメータは、日付（ "2006-01-02"）、日時（ "2006-01-02 15:04:05"）または過去の時間の期間（ "100d"または "1h"）である必要があります。

      このオプションを使用する場合、ファイルの書き込み操作は許可されません。
      つまり、ファイルをアップロードまたは削除することはできません。

      有効なフォーマットについては、[時刻オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。

      "Content-Encoding: gzip"が設定されたままS3にオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneは受信したときに "Content-Encoding: gzip" のファイルを展開します。これにより、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルのコンテンツは展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzipに圧縮するかもしれない場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードするときには変更しません。"Content-Encoding: gzip" がアップロードされなかった場合、ダウンロード時にも設定されません。

      しかし、一部のプロバイダは、"Content-Encoding: gzip" でない場合でもオブジェクトをgzipに圧縮する場合があります（たとえば、Cloudflare）。

      これにより、次のようなエラーメッセージが表示されることがあります。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneがContent-Encoding: gzipが設定されたオブジェクトをチャンク化転送エンコードでダウンロードした場合、rcloneはオブジェクトをそのまま解凍します。

      unsetに設定する場合（デフォルト）は、rcloneはプロバイダの設定に従って適用する内容を選択しますが、ここでrcloneの選択をオーバーライドすることができます。

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制します


OPTIONS:
   --access-key-id value      AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --endpoint value           Storjゲートウェイのエンドポイント。 [$ENDPOINT]
   --env-auth                 ランタイムからAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ）。（デフォルト: false） [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケット作成時に使用する設定ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。（デフォルト: "5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値。（デフォルト: "4.656Gi"） [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipでエンコードされたオブジェクトを展開します。（デフォルト: false） [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。（デフォルト: false） [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。（デフォルト: false） [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト: "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。（デフォルト: true） [$FORCE_PATH_STYLE]
   --list-chunk value               リスト表示のチャンクサイズ（各ListObject S3リクエストごとのレスポンスリスト）。 （デフォルト: 1000） [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset （デフォルト: "unset"） [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または自動的に変換するための0。（デフォルト: 0） [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用するパーツの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。（デフォルト: "1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト: false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipに圧縮するかもしれない場合に設定します。 （デフォルト: "unset"） [$MIGHT_GZIP]
   --no-check-bucket                バケットが存在するかチェックしません。（デフォルト: false） [$NO_CHECK_BUCKET]
   --no-head                        HEADリクエストを使用してアップロードしたオブジェクトの整合性をチェックしません。（デフォルト: false） [$NO_HEAD]
   --no-head-object                 GETでオブジェクトを取得する前にHEADを実行しない場合に設定します。（デフォルト: false） [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制します。（デフォルト: false） [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。（デフォルト: 4） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ値。（デフォルト: "200Mi"） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか（デフォルト: "unset"） [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードの場合、PresignedリクエストまたはPutObjectを使用するかどうか（デフォルト: false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。（デフォルト: false） [$V2_AUTH]
   --version-at value               指定した時点のファイルバージョンを表示します。（デフォルト: "off"） [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか（デフォルト: false） [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}