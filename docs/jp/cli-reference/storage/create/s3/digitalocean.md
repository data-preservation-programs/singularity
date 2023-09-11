# DigitalOcean Spaces

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 digitalocean - DigitalOcean Spaces

USAGE:
   singularity storage create s3 digitalocean [command options] [arguments...]

DESCRIPTION:
   --env-auth
      ランタイムからAWS認証情報を取得します（環境変数またはEC2 / ECSメタデータがない場合）。

      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境からAWS認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSアクセスキーID。

      匿名アクセスまたはランタイム認証情報ではない場合は空にしてください。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。

      匿名アクセスまたはランタイム認証情報ではない場合は空にしてください。

   --region
      接続するリージョン。

      S3クローンを使用している場合でリージョンがない場合は空にしてください。

      例:
         | <unset>            | 判断がつかない場合に使用します。
         |                    | v4署名と空のリージョンが使用されます。
         | other-v2-signature | v4署名が機能しない場合にのみ使用します。
         |                    | たとえば、Jewel/v10 CEPH以前。

   --endpoint
      S3 APIのエンドポイント。

      S3クローンを使用している場合は必須です。

      例:
         | syd1.digitaloceanspaces.com | DigitalOcean Spaces Sydney 1
         | sfo3.digitaloceanspaces.com | DigitalOcean Spaces San Francisco 3
         | fra1.digitaloceanspaces.com | DigitalOcean Spaces Frankfurt 1
         | nyc3.digitaloceanspaces.com | DigitalOcean Spaces New York 3
         | ams3.digitaloceanspaces.com | DigitalOcean Spaces Amsterdam 3
         | sgp1.digitaloceanspaces.com | DigitalOcean Spaces Singapore 1

   --location-constraint
      リージョンと一致する必要がある場所の制約。

      判断がつかない場合は空にしてください。バケットを作成する場合にのみ使用されます。

   --acl
      バケットを作成したりオブジェクトを保存したりコピーするために使用されるかんたんACL。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合にもバケットの作成に使用されます。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      S3はサーバーサイドでオブジェクトをコピーする際にACLをコピーせず、新たに書き込むため、このACLが適用されます。

      aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

   --bucket-acl
      バケットを作成するときに使用されるかんたんACL。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      バケットの作成時にのみ使用され、設定されていない場合は「acl」が代わりに使用されます。

      `acl`と`bucket_acl`が空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | ファイルの所有者がFULL_CONTROLを取得します。
         |                    | 誰もがアクセス権限を持たない（デフォルト）。
         | public-read        | ファイルの所有者がFULL_CONTROLを取得します。
         |                    | AllUsersグループがREADアクセスを取得します。
         | public-read-write  | ファイルの所有者がFULL_CONTROLを取得します。
         |                    | AllUsersグループがREADおよびWRITEアクセスを取得します。
         |                    | バケットでこれを許可することは一般的には推奨されていません。
         | authenticated-read | ファイルの所有者がFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループがREADアクセスを取得します。

   --upload-cutoff
      チャンク化アップロードに切り替えるためのカットオフ。

      これより大きなファイルはchunk_sizeのchunk単位でアップロードされます。
      最小値は0で、最大値は5GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffよりも大きなサイズのファイルや、サイズが不明なファイル（たとえば「rclone rcat」からのアップロードや「rclone mount」またはgoogle photosまたはgoogle docsでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。

      注意："--s3-upload-concurrency"このサイズのチャンクは、個々の転送ごとにメモリにバッファされます。

      高速リンクで大きなファイルを転送していて、十分なメモリがある場合、これを増やすと転送速度が向上します。

      Rcloneは、既知のサイズの大きなファイルをアップロードするときにチャンクサイズを自動的に増やし、10000チャンクの制限を下回るようにします。

      サイズが不明なファイルは設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5MiBで、最大で10000個のチャンクがある場合、デフォルトではストリームアップロードできるファイルの最大サイズは48GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、"-P"フラグで表示される進行状況統計の正確性が低下します。Rcloneは、チャンクがAWS SDKによってバッファされたときにチャンクが送信されたと扱い、まだアップロード中である場合でも、チャンクのサイズが大きくなればなるほど、AWS SDKのバッファが大きくなり、進行状況が真実から外れる可能性があります。

   --max-upload-parts
      マルチパートアップロードの最大パート数。

      このオプションは、マルチパートアップロードを行う際に使用するパートの最大数を定義します。

      サービスがAWS S3の10000パートの仕様をサポートしていない場合に便利です。

      Rcloneは、既知のサイズの大きなファイルをアップロードするときにチャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。

      サーバーサイドでコピーする必要のあるこれより大きなファイルは、このサイズのチャンクでコピーされます。

      最小値は0で、最大値は5GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合は、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクを同時にアップロードする数です。

      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅をフルに活用していない場合、これを増やすと転送を高速化するのに役立ちます。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。

      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。
      falseの場合は仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（たとえばAWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4署名が機能しない場合にのみ使用してください（たとえばJewel/v10 CEPH以前）。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストの応答リスト）。

      このオプションは、AWS S3のMaxKeys、「max-items」、または「page-size」としても知られています。
      大多数のサービスは、要求された以上のリストを最大1000件で切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、またはオプションで0。

      S3は元々、バケット内のオブジェクトを列挙するためのListObjects呼び出しのみを提供していました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは非常に高速であり、可能であれば利用する必要があります。

      デフォルトで0に設定されている場合、rcloneはプロバイダによってセットされたguessからどちらのlist objectsメソッドを呼び出すかを推測します。推測が間違っている場合は、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダでは、ファイル名に制御文字を使用する際にURLエンコードリストをサポートしており、これが利用可能な場合は使用すると信頼性が向上します。これがunset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用することを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットの存在を確認せず、作成しようとしません。

      バケットが既に存在することを知っている場合、rcloneが行うトランザクションの数を最小限にするために役立ちます。

      バケット作成の権限がない場合も必要です。v1.52.0より前のバージョンでは、このエラーが静かにパスされました。

   --no-head
      アップロードされたオブジェクトの整合性を確認するために、HEADリクエストを行いません。

      rcloneは、PUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正常にアップロードされたと見なします。

      特に次のものを前提とします。

      - メタデータ（modtime、ストレージクラス、コンテンツタイプを含む）がアップロード時と同じであること
      - サイズがアップロード時と同じであること

      以下のアイテムを1つのパートPUTの応答から読み取ります。

      - MD5SUM
      - アップロードされた日付

      マルチパートアップロードの場合、これらのアイテムは読み取られません。

      サイズが不明のソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出される可能性が高くなります。特にサイズが不正確な場合ですので、通常の操作ではお勧めできません。実際のアップロードの失敗の可能性は非常に低いです。

   --no-head-object
      オブジェクトを取得する前にHEADリクエストを行いません。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要セクションのエンコーディング](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      インターナルメモリバッファプールがフラッシュされる頻度。

      追加のバッファが必要なアップロード（たとえばマルチパート）では、アロケーションのためにメモリプールが使用されます。このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      インターナルメモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2の問題が解決できていません。s3バックエンドの場合、デフォルトでHTTP/2が有効になっていますが、ここで無効にすることもできます。問題が解決したら、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのためのカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワーク経由でダウンロードされたデータに対してより安価な転送を提供するため、これはCloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか。

      true、false、またはデフォルトを使用します。

   --use-presigned-request
      シングルパートアップロードの場合、署名済みリクエストまたはPutObjectを使用するかどうか。

      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン<1.59は、1つのパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueにするとその機能が再度有効になります。これは例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時間のファイルバージョンを表示します。

      パラメータは日付、「2006-01-02」、日時「2006-01-02 15:04:05」、またはその長い前に戻るための期間、例えば「100d」または「1h」です。

      このオプションを使用する場合、ファイルの書き込み操作は許可されません。つまり、ファイルをアップロードしたり削除したりすることはできません。

      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      Gzipエンコードされたオブジェクトを解凍します。

      S3に「Content-Encoding: gzip」を設定してオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneは「Content-Encoding: gzip」で受信したファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルの内容は展開されます。

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。

      通常、プロバイダはオブジェクトをダウンロードする際にそれらを変更しません。`Content-Encoding: gzip`でアップロードされなかったオブジェクトはダウンロード時にも設定されません。

      ただし、一部のプロバイダでは、`Content-Encoding: gzip`でアップロードされていないオブジェクトもgzip化する場合があります（例：Cloudflare）。

      これを設定し、rcloneがContent-Encoding: gzipが設定されたオブジェクトとチャンクされたトランスファーエンコードを受信すると、rcloneはオブジェクトを動的に展開します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用することを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  バケットを作成したりオブジェクトを保存したりコピーするために使用されるかんたんACL。 [$ACL]
   --endpoint value             S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                   ランタイムからAWS認証情報を取得します（環境変数またはEC2 / ECSメタデータがない場合）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンと一致する必要がある場所の制約。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。 [$REGION]
   --secret-access-key value    AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   アドバンスト

   --bucket-acl value               バケットを作成するときに使用されるかんたんACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     もし設定するとgzipエンコードされたオブジェクトを解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストの応答リスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1,2または自動のための0。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パート数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   インターナルメモリバッファプールがフラッシュされる頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           インターナルメモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、作成しようとしません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトの整合性を確認するために、HEADリクエストを行いません。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前にHEADリクエストを行いません。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            An AWS session token. [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのカットオフ。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードの場合、署名済みリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。 (デフォルト: false) [$VERSIONS]

   一般

   --name value  ストレージの名前（デフォルト: オートジェネレート）
   --path value  ストレージのパス

```
{% endcode %}