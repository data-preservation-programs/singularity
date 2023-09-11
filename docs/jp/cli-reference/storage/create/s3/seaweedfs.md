# SeaweedFS S3

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 seaweedfs - SeaweedFSのS3

USAGE:
   singularity storage create s3 seaweedfs [コマンドオプション] [引数...]

DESCRIPTION:
   --env-auth
      AWSの認証情報を実行時に取得します（環境変数または環境変数が存在しない場合はEC2/ECSのメタデータ）。
      
      access_key_idおよびsecret_access_keyが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報の場合は空白のままにします。

   --region
      接続するリージョン。
      
      S3クローンを使用しており、リージョンがない場合は空白のままにします。

      例:
         | <未設定>            | 不確かな場合はこのままにします。
         |                    | v4の署名と空のリージョンが使用されます。
         | other-v2-signature | v4の署名が機能しない場合のみ使用します。
         |                    | 例: Jewel/v10 CEPHより前のバージョン。

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用している場合は必須です。

      例:
         | localhost:8333 | SeaweedFS S3のlocalhost

   --location-constraint
      リージョンに一致する場所の制約です。
      
      わからない場合は空白のままにします。バケット作成時にのみ使用されます。

   --acl
      バケットの作成、およびオブジェクトの保存またはコピー時に使用されるCanned ACL。
      
      このACLはオブジェクトの作成にも使用され、bucket_aclが設定されていない場合にも使用されます。
      
      詳細については、[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3はオブジェクトをサーバーサイドでコピーする際に、ソースからACLをコピーするのではなく、新しいACLを書き込みます。
      
      aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト値（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACL。
      
      詳細については、[こちらのドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      このACLはバケットの作成時にのみ使用されます。設定されていない場合は「acl」が使用されます。
      
      aclとbucket_aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト値（private）が使用されます。
      

      例:
         | private            | 所有者にFULL_CONTROLが与えられます。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | 所有者にFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREAD権限が与えられます。
         | public-read-write  | 所有者にFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADとWRITE権限が与えられます。
         |                    | バケットでこれを許可することは一般的にお勧めしません。
         | authenticated-read | 所有者にFULL_CONTROLが与えられます。
         |                    | AuthenticatedUsersグループにREAD権限が与えられます。

   --upload-cutoff
      チャンクアップロードに切り替えるためのCutoff値です。
      
      この値より大きいファイルはchunk_sizeごとにチャンクとしてアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクのサイズです。
      
      upload_cutoffよりも大きなサイズのファイルや、サイズが不明なファイル（例: "rclone rcat"からアップロードされたファイルや"rclone mount"やGoogleフォトやGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意: "--s3-upload-concurrency"のチャンクごとのこのサイズのバッファがメモリに保存されます。
      
      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを大きくすることで転送を高速化できます。
      
      rcloneは、10,000個のチャンクの制限を下回るように、既知の大きさの大きなファイルをアップロードする際に自動的にチャンクサイズを増やします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeが5 MiBであり、最大で10,000個のチャンクがあることを考慮すると、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、"-P"フラグと共に表示される進捗状況の精度が低下します。rcloneは、AWS SDKによってバッファリングされたチャンクを送信したときにチャンクが送信されたと処理するため、実際にはまだアップロード中かもしれないです。
      
      チャンクサイズが大きいほど、AWS SDKのバッファが大きくなり、実際の進行状況とはかけ離れた進行状況が報告されます。
      

   --max-upload-parts
      マルチパートアップロードで使用するパートの最大数です。
      
      このオプションは、マルチパートアップロード時に使用するマルチパートチャンクの最大数を定義します。
      
      これは、サービスがAWS S3仕様の10,000個のチャンクをサポートしていない場合に便利です。
      
      rcloneは、既知の大きさの大きなファイルをアップロードする際に自動的にチャンクサイズを増やしてこのチャンク数の制限以下に保ちます。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのCutoff値です。
      
      これより大きいサイズのファイルをサーバーサイドでコピーする必要がある場合は、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルの場合は開始までに長時間遅延することがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」の環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は、環境変数「AWS_PROFILE」または「default」が設定されます。
      

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの同時実行数です。
      
      同じファイルのチャンクの数を同時にアップロードします。
      
      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすと転送速度が向上する場合があります。

   --force-path-style
      trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式のアクセスを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパス形式のアクセスを使用し、falseの場合、rcloneは仮想パス形式を使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（例: AWS、Aliyun OSS、Netease COS、Tencent COS）は、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2の認証を使用します。
      
      これがfalse（デフォルト）に設定されている場合、rcloneはv4の認証を使用します。設定されている場合、rcloneはv2の認証を使用します。
      
      v4の署名が機能しない場合のみ使用してください。例: Jewel/v10 CEPHより前のバージョン。

   --list-chunk
      リストのチャンクのサイズ（各ListObject S3リクエストの応答リスト）です。
      
      このオプションは、AWS S3のMaxKeys、max-items、またはpage-sizeとしても知られています。
      ほとんどのサービスは、要求されたオブジェクトが1000を超えていても、応答リストを切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。詳細は[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン: 1、2、または0（自動）。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しのみが提供されていました。
      
      しかし、2016年5月にはListObjectsV2呼び出しが導入されました。これははるかに高性能であり、可能な限り使用すべきです。
      
      デフォルトの0に設定すると、rcloneはプロバイダの設定に基づいてどのリストオブジェクトメソッドを呼び出すか推測します。推測が誤っている場合は、ここで手動で設定できます。
      

   --list-url-encode
      リストをURLエンコードするかどうか: true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を使用するときにURLエンコードリストをサポートしています。これが利用可能な場合には、これがファイルの名前に制御文字を含めた場合により信頼性が高くなります。unsetに設定されている場合（デフォルト）、rcloneは、プロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。
      

   --no-check-bucket
      バケットの存在を確認せず、作成しようとしない場合に設定します。
      
      バケットが既に存在することを知っている場合、rcloneが行うトランザクションの数を最小限に抑えるためにこのオプションを使用できます。
      
      バケット作成の権限を持っていない場合も必要になる場合があります。v1.52.0より前のバージョンでは、バグのためこの設定が無視されていましたが、これによりエラーが表示されるようになりました。
      

   --no-head
      チェックダムとしてオブジェクトをHEADせずにアップロードします。
      
      rcloneは、PUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正しくアップロードされたと想定します。
      
      特に、以下が想定されます。
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロード時と同じであること
      - サイズがアップロード時と同じであること
      
      シングルパートのPUTのレスポンスからは次の項目が読み取られます。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      不明な長さのソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの際にアップロード失敗の可能性が高まるため、通常の操作では推奨されません。実際には、このフラグを設定してもアップロード失敗の可能性は非常に低いです。
      

   --no-head-object
      GETする前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細は、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度です。
      
      追加のバッファを必要とするアップロード（たとえばマルチパート）は、アロケーションにメモリプールを使用します。
      このオプションは、未使用のバッファをプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      現在、S3（特にminio）バックエンドとHTTP/2の問題が解決されていません。S3バックエンドのHTTP/2はデフォルトで有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロードのカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してダウンロードされたデータに対して安価な転送費用を提供しています。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      true、false、またはデフォルトプロバイダの設定を使用する場合に設定します。
      

   --use-presigned-request
      シングルパートアップロードの場合に署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン< 1.59では、シングルパートオブジェクトをアップロードするために署名付きリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは特殊な状況やテストのためにのみ必要です。

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはそれ以前の期間（例: "100d"や"1h"）で指定できます。
      
      このオプションを使用すると、書き込み操作が許可されないため、ファイルをアップロードまたは削除することはできません。
      
      有効な形式については[時刻オプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      オブジェクトがgzipでエンコードされている場合、解凍します。
      
      "Content-Encoding: gzip"が設定されたままS3にオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを受け取ったときに「Content-Encoding: gzip」で解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際にオブジェクトを変更しません。オブジェクトが `Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時に設定されません。
      
      ただし、いくつかのプロバイダは、`Content-Encoding: gzip`でアップロードされていない場合でもオブジェクトをgzip化する場合があります（たとえばCloudflare）。
      
      これは次のようなエラーが表示されることの症状です。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneが`Content-Encoding: gzip`が設定され、チャンク化された転送エンコーディングを使用してオブジェクトをダウンロードした場合、rcloneはオブジェクトを逐次的に解凍します。
      
      unsetに設定する場合（デフォルト）は、rcloneはプロバイダの設定に従って適用するかどうかを選択しますが、ここでrcloneの選択を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


OPTIONS:
   --access-key-id value        AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成とオブジェクトの保存またはコピー時に使用されるCanned ACL。 [$ACL]
   --endpoint value             S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                   AWSの認証情報を実行時に取得します（環境変数または環境変数が存在しない場合はEC2/ECSのメタデータ）。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンに一致する場所の制約。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクのサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのCutoff値。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     オブジェクトがgzipでエンコードされている場合、解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式のアクセスを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクのサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか。 (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1、2、または0（自動）。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用するパートの最大数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、作成しようとしない場合に設定します。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        チェックダムとしてオブジェクトをHEADせずにアップロードします。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを行わない場合に設定します。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのCutoff値。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードの場合に署名済みリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2の認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。 (デフォルト: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}