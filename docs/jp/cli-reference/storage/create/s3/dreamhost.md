# Dreamhost DreamObjects

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 dreamhost - Dreamhost DreamObjects

USAGE:
   singularity storage create s3 dreamhost [command options] [arguments...]

DESCRIPTION:
   --env-auth
      [Option] ランタイムからAWS認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSメタデータ）。
      
      アクセスキーIDとシークレットアクセスキーが空である場合のみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力してください。
         | true  | 環境（環境変数またはIAM）からAWS認証情報を取得します。

   --access-key-id
      [オプション] AWSアクセスキーID。
      
      匿名アクセスまたはランタイム認証情報の場合は空にします。

   --secret-access-key
      [オプション] AWSシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたはランタイム認証情報の場合は空にします。

   --region
      [オプション] 接続するリージョン。
      
      S3のクローンを使用していてリージョンを持っていない場合は空にします。

      例:
         | <unset>            | 迷った場合はこれを使用します。
         |                    | v4シグネチャと空のリージョンが使用されます。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                    | 例: Jewel/v10 CEPH以前。

   --endpoint
      [オプション] S3 APIのエンドポイント。
      
      S3クローンを使用している場合は必須です。

      例:
         | objects-us-east-1.dream.io | Dream Objectsのエンドポイント

   --location-constraint
      [オプション] リージョンと一致するように設定する場所の制約。
      
      自信がない場合は空にします。バケットを作成する際にのみ使用されます。

   --acl
      [オプション] バケットの作成およびオブジェクトの保存またはコピー時に使用されるキャニステートACL。
      
      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。
      
      詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3がオブジェクトをサーバー側でコピーする場合、
      S3はソースからACLをコピーせずに新しいACLを書き込みます。
      
      もしaclが空の文字列ならばX-Amz-Acl:ヘッダは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      [オプション] バケットの作成時に使用されるキャニステートACL。
      
      詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      このACLはバケットの作成時のみ適用されます。設定されていない場合は「acl」が代わりに使用されます。
      
      もし「acl」と「bucket_acl」が空の文字列ならばX-Amz-Acl:ヘッダは追加されず、デフォルト（private）が使用されます。
      

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADアクセスを取得します。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADおよびWRITEアクセスを取得します。
         |                    | バケットに対してこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループはREADアクセスを取得します。

   --upload-cutoff
      [オプション] チャンクアップロードに切り替えるためのカットオフ。
      
      このサイズより大きなファイルはchunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5GiBです。

   --chunk-size
      [オプション] アップロードに使用するチャンクサイズ。
      
      upload_cutoffより大きいファイル、またはサイズの不明なファイル（例: "rclone rcat"からのアップロードや"rclone mount"またはGoogleフォトやGoogleドキュメントからアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      "--s3-upload-concurrency"スイッチチャンクのユーザーごとにメモリ内にバッファリングされます。
      
      高速リンク上で大きなファイルを転送していて十分なメモリがある場合、これを増やすことで転送速度を上げることができます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、10,000のチャンク制限を超えないようにします。
      
      サイズの不明なファイルは設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5MiBですが、最大で10,000のチャンクがありますので、デフォルトのストリームアップロード可能なファイルの最大サイズは48GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクのサイズを増やすことにより、進行状況の統計情報が表示されるときの正確さが低下します。Rcloneは、AWS SDKによってチャンクがバッファリングされたときにチャンクを送信したと判断し、まだアップロード中である可能性がありますが、実際には違います。
      

   --max-upload-parts
      [オプション] マルチパートアップロード内のパートの最大数。
      
      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。
      
      これは、サービスがAWS S3の10,000個のチャンクの仕様をサポートしていない場合に役立ちます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。
      

   --copy-cutoff
      [オプション] マルチパートコピーに切り替えるためのカットオフ。
      
      サーバーサイドでコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5GiBです。

   --disable-checksum
      [オプション] オブジェクトメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算して、オブジェクトのメタデータに追加します。これはデータの整合性検査には非常に有効ですが、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      [オプション] 共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトです。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      [オプション] 共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用することができます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数"AWS_PROFILE"または設定されていない場合は"default"がデフォルトです。
      

   --session-token
      [オプション] AWSセッショントークン。

   --upload-concurrency
      [オプション] マルチパートアップロードの同時実行数。
      
      同時にアップロードされる同じファイルのチャンク数です。
      
      高速回線で大量のファイルをアップロードし、これらのアップロードが帯域幅をフルに利用しない場合、これを増やすと転送速度が向上することがあります。

   --force-path-style
      [オプション] trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。
      falseの場合、rcloneは仮想パススタイルを使用します。詳細は[the AWS S3 docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（例: AWS、Aliyun OSS、Netease COS、またはTencent COS）では、falseに設定する必要があります。このプロバイダに基づいてrcloneが自動的に行います。

   --v2-auth
      [オプション] trueの場合はv2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャが機能しない場合のみ使用します。旧Jewel/v10 CEPHなど。

   --list-chunk
      [オプション] リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）。
      
      このオプションは、AWS S3仕様のMaxKeys、max-items、またはpage-sizeとしても知られています。
      大部分のサービスは、リクエストされた以上の1000オブジェクトを切り捨てます。
      AWS S3では、これはグローバルの最大値であり、変更することはできません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、オプション「rgw list buckes max chunk」で増やすことができます。
      

   --list-version
      [オプション] 使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にリリースされた当初は、バケット内のオブジェクトを列挙するためにListObjects呼び出しのみが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高い性能を提供し、できる限り使用するべきです。
      
      デフォルトの設定（0）にすると、rcloneはプロバイダに応じてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が間違っている場合は、ここで手動で設定できます。
      

   --list-url-encode
      [オプション] リストをURLエンコードするかどうか：true / false / unset
      
      一部のプロバイダは、リストをURLエンコードサポートしており、ファイル名に制御文字を使用する場合にはこれがより信頼性があります。この設定がunset（デフォルト）になっている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。
      

   --no-check-bucket
      [オプション] バケットの存在を確認せずに作成またはチェックしようとしない場合に設定します。
      
      クライアントがバケットが既に存在することを知っている場合、rcloneが行うトランザクションの数を最小限にするために便利です。
      
      バケット作成権限を持っていない場合、必要になるかもしれません。v1.52.0より前のバージョンでは、バグのために無条件に合格します。
      

   --no-head
      [オプション] アップロードしたオブジェクトのHEADをチェックして整合性を確認しない場合に設定します。
      
      rcloneはこのフラグが設定されている場合、PUTでオブジェクトをアップロードした後、200 OKのメッセージを受け取った場合は、正しくアップロードされたと見なします。
      
      特に、次のことを想定します：
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプを含む）がアップロード時と同じであること。
      - サイズがアップロード時と同じであること。
      
      単一パートのPUTの応答から読み取る項目は次のとおりです：
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードでは、これらの項目は読み取られません。
      
      不明な長さの元のオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、正常なオペレーションには推奨されませんが、アップロードの失敗を検出できない可能性が増えます。
      特に誤ったサイズの場合ですが、このフラグを使用しない場合でも、アップロード失敗の可能性は非常に小さいです。

   --no-head-object
      [オプション] GETの前にHEADを行わない場合に設定します。

   --encoding
      [オプション] バックエンドのエンコーディング。
      
      詳細は、[エンコーディングの概要](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      [オプション] 内部メモリバッファプールがフラッシュされる頻度。
      
      追加バッファが必要なアップロード（たとえばマルチパート）は、割り当てにメモリプールを使用します。
      このオプションは、使用されなくなったバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      [オプション] 内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      [オプション] S3バックエンドでのhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2の問題が未解決です。s3バックエンドでHTTP/2はデフォルトで有効ですが、ここで無効にすることもできます。問題が解決された場合、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673 、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      [オプション] ダウンロードのカスタムエンドポイント。
      これは通常、AWS S3がCloudFrontネットワーク経由でダウンロードされたデータに対して安価なエグレスを提供するため、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      [オプション] マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルトで使用するかどうかプロバイダに依存します。
      

   --use-presigned-request
      [オプション] シングルパートアップロードのために署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。
      
      rclone < 1.59のバージョンでは、単一のパートオブジェクトをアップロードするために署名済みのリクエストを使用していたため、このフラグをtrueに設定すると、その機能が再度有効になります。これは、異常な状況またはテストの場合を除いては必要ありません。
      

   --versions
      [オプション] ディレクトリリストに古いバージョンを含めます。

   --version-at
      [オプション] 指定した時間のファイルバージョンを表示します。
      
      パラメータは日付（"2006-01-02"）、日時（"2006-01-02 15:04:05"）、その長い前の期間（例: "100d"または"1h"）である必要があります。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      [オプション] これが設定されている場合、gzipエンコードされたオブジェクトを展開します。
      
      S3に「Content-Encoding: gzip」が設定されている状態でオブジェクトをアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信時に「Content-Encoding: gzip」でこれらのファイルを展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は展開されます。
      

   --might-gzip
      [オプション] バックエンドでオブジェクトがgzipになる可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更することはありません。`Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時には設定されません。
      
      ただし、一部のプロバイダ（例: Cloudflare）は、オブジェクトが`Content-Encoding: gzip`でない場合でも圧縮する場合があります。
      
      これに設定することにより、rcloneがContent-Encoding: gzipとチャンク転送エンコーディングが設定されたオブジェクトをダウンロードした場合、rcloneはオブジェクトを即座に展開します。
      
      これがunset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。
      

   --no-system-metadata
      [オプション] システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用されるキャニステートACL。 [$ACL]
   --endpoint value             S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                   ランタイムからAWS認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンと一致するように設定する場所の制約。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。 [$REGION]
   --secret-access-key value    AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるキャニステートACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipエンコードされたオブジェクトを展開します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1,2、または0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロード内のパートの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドでオブジェクトがgzipになる可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せずに作成またはチェックしようとしない場合に設定します。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトのHEADをチェックして整合性を確認しない場合に設定します。 (default: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わない場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードのために署名済みリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}