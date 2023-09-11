# Qiniuオブジェクトストレージ（Kodo）

{% code fullWidth="true" %}
```
名前:
   singularity storage update s3 qiniu - Qiniuオブジェクトストレージ（Kodo）

使用方法:
   singularity storage update s3 qiniu [command options] <name|id>

説明:
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数または環境変数が存在しない場合はEC2/ECSメタデータから）。
      
      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力してください。
         | true  | 環境からAWSの認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSのAccess Key IDです。
      
      匿名アクセスまたは実行時の認証情報を使用する場合は空のままにしてください。

   --secret-access-key
      AWSのSecret Access Key（パスワード）です。
      
      匿名アクセスまたは実行時の認証情報を使用する場合は空のままにしてください。

   --region
      接続するリージョンです。

      例:
         | cn-east-1      | デフォルトのエンドポイント-迷った場合の選択肢
         |                | 中国東部リージョン1
         |                | cn-east-1の指定が必要
         | cn-east-2      | 中国東部リージョン2
         |                | cn-east-2の指定が必要
         | cn-north-1     | 中国北部リージョン1
         |                | cn-north-1の指定が必要
         | cn-south-1     | 中国南部リージョン1
         |                | cn-south-1の指定が必要
         | us-north-1     | 北米リージョン
         |                | us-north-1の指定が必要
         | ap-southeast-1 | 東南アジアリージョン1
         |                | ap-southeast-1の指定が必要
         | ap-northeast-1 | 北東アジアリージョン1
         |                | ap-northeast-1の指定が必要

   --endpoint
      Qiniuオブジェクトストレージのエンドポイントです。

      例:
         | s3-cn-east-1.qiniucs.com      | 中国東部エンドポイント1
         | s3-cn-east-2.qiniucs.com      | 中国東部エンドポイント2
         | s3-cn-north-1.qiniucs.com     | 中国北部エンドポイント1
         | s3-cn-south-1.qiniucs.com     | 中国南部エンドポイント1
         | s3-us-north-1.qiniucs.com     | 北米エンドポイント1
         | s3-ap-southeast-1.qiniucs.com | 東南アジアエンドポイント1
         | s3-ap-northeast-1.qiniucs.com | 北東アジアエンドポイント1

   --location-constraint
      リージョンと一致する必要がある場所拘束条件です。
      
      バケットを作成する場合にのみ使用されます。

      例:
         | cn-east-1      | 中国東部リージョン1
         | cn-east-2      | 中国東部リージョン2
         | cn-north-1     | 中国北部リージョン1
         | cn-south-1     | 中国南部リージョン1
         | us-north-1     | 北米リージョン1
         | ap-southeast-1 | 東南アジアリージョン1
         | ap-northeast-1 | 北東アジアリージョン1

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成時に使用され、bucket_aclが設定されていない場合にもバケットの作成に使用されます。
      
      詳細はこちらをご覧ください：[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      S3では、サーバー側でオブジェクトをコピーする際、元のACLをコピーするのではなく、新しいACLを書き込むため、このACLは適用されます。
      
      aclが空の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。
      
      詳細はこちらをご覧ください：[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      このACLはバケット作成時のみ適用されます。設定されていない場合は「acl」が代わりに使用されます。
      
      「acl」と「bucket_acl」が空の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他の誰もアクセス権を持ちません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADアクセスを取得します。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループはREADおよびWRITEアクセスを取得します。
         |                    | バケットでこれを許可することは一般的にはおすすめされません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループはREADアクセスを取得します。

   --storage-class
      Qiniuに新しいオブジェクトを保存する際に使用するストレージクラスです。

      例:
         | STANDARD     | 標準ストレージクラス
         | LINE         | 頻繁なアクセスストレージモード
         | GLACIER      | アーカイブストレージモード
         | DEEP_ARCHIVE | ディープアーカイブストレージモード

   --upload-cutoff
      チャンク化アップロードに切り替えるための上限です。
      
      この上限を超えるファイルは、chunk_sizeのチャンクでアップロードされます。
      最小は0、最大は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffを超えるファイルやサイズが不明なファイル（例：「rclone rcat」からのアップロード、または「rclone mount」やGoogleフォトまたはGoogleドキュメントからのアップロード）は、このチャンクサイズを使用してマルチパートでアップロードされます。
      
      注意：「--s3-upload-concurrency」のチャンクは、このサイズのバッファがメモリ上に転送ごとに存在します。
      
      高速リンクを介して大きなファイルを転送している場合で十分なメモリがある場合、これを増やすと転送速度が向上します。
      
      rcloneは、既知のサイズの大きなファイルをアップロードするときにチャンクサイズを自動的に増やして、最大10,000のチャンク制限を下回るようにします。
      
      サイズの不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBで、最大10,000のチャンクがあるという制約があります。したがって、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、「-P」フラグで表示される進行状況統計の正確性が低下します。rcloneは、AWS SDKによってバッファリングされたチャンクが送信されたときにチャンクが送信されたとみなしますが、実際にはまだアップロード中の可能性があります。チャンクサイズが大きくなると、AWS SDKのバッファも大きくなり、真実から逸脱した進行報告が行われます。

   --max-upload-parts
      マルチパートアップロードの最大パーツ数です。
      
      マルチパートアップロードを実行する際に使用するチャンクの最大数を定義するオプションです。
      
      AWS S3の10,000チャンク仕様をサポートしていないサービスの場合、これが役立つ場合があります。
      
      rcloneは、既知のサイズの大きなファイルをアップロードするときにチャンクサイズを自動的に増やして、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるための上限です。
      
      サーバーサイドでコピーする必要があるこの上限を超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小は0、最大は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しないでください。
      
      通常、rcloneはアップロード前に入力データのMD5チェックサムを計算し、それをオブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有の認証情報ファイルへのパスです。
      
      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有の認証情報ファイルで使用するプロファイルです。
      
      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルトになります。
      

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行数です。
      
      同じファイルのチャンク数としてアップロードされるチャンクの数です。
      
      高速リンクを介して大量のファイルを高速で転送しており、これらの転送を完全に帯域幅を利用していない場合は、これを増やすと転送速度が向上することがあります。

   --force-path-style
      trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      true（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      を参照してください。
      
      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、この設定をfalseにする必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2認証を使用します。
      
      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみこれを使用してください。たとえば、Jewel/v10 CEPHの前。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）のサイズです。
      
      このオプションはAWS S3の仕様での"MaxKeys"、"max-items"、または"page-size"とも知られています。
      ほとんどのサービスでは、要求された数よりも多くのオブジェクトをリストできない場合でも、一部のオブジェクト数を切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0は自動です。
      
      初めにS3がリリースされたとき、バケット内のオブジェクトを列挙するためのListObjectsコールスが提供されました。
      
      しかし、2016年5月にListObjectsV2コールが導入されました。これははるかに高性能で、可能な限り使用する必要があります。
      
      デフォルト値である0に設定すると、rcloneはプロバイダの設定に従ってどのリストオブジェクトメソッドを呼び出すかを推測します。推測が誤っている場合は、ここで手動で設定できます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      数字の名前を含むファイル名を使用する場合、一部のプロバイダではURLエンコードリストがサポートされており、制御文字を使用する場合にはこれがより信頼性があります。これがunsetに設定されている場合（デフォルト）、rcloneは提供者の設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることもできます。
      

   --no-check-bucket
      オブジェクトの存在を確認したり、バケットを作成したりしません。
      
      バケットがすでに存在する場合は、rcloneが実行するトランザクションの数を最小限に抑えるために、これを設定すると便利です。
      
      バケット作成権限を持っていないユーザーを使用する場合も必要です。v1.52.0以前では、これはバグのために無視されていました。
      

   --no-head
      アップロードしたオブジェクトをHEADして整合性を確認しない場合に設定します。
      
      rcloneはすべてのオブジェクトのアップロード後にHEADリクエストを送信して整合性を確認しますが、これを設定するとトランザクションの数を最小限に抑えることができます。
      
      セットすると、rcloneはPUTでオブジェクトをアップロードした後に200 OKメッセージを受け取ると、正しくアップロードされたと想定します。
      
      特に次のことを想定します：
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプなど）がアップロード時と同じであること。
      - サイズがアップロード時と同じであること。
      
      PUTのシングルパートの場合、次のアイテムをレスポンスから読み取ります。
      
      - MD5SUM
      - アップロード日
      
      マルチパートアップロードでは、これらのアイテムは読み取られません。
      
      不明な長さのソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを行います。
      
      このフラグを設定すると、アップロードエラーの機会が増え、特にサイズが不正確な場合には推奨されません。実際には、このフラグを使用しても、アップロードエラーの可能性は非常に低いです。
      

   --no-head-object
      GETする前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコードセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールのフラッシュが行われる頻度です。
      
      追加のバッファを必要とするアップロード（たとえばマルチパート）では、メモリプールが割当に使用されます。
      このオプションは、バッファの未使用の部分がプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効にします。
      
      現在、s3（具体的にはminio）バックエンドとHTTP/2に関する未解決の問題があります。 s3バックエンドでは、HTTP/2がデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決したら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      これは通常、AWS S3でデータをダウンロードするためのエグレスがさらに安価なCloudFront CDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルトを使用するため、true、false、または未設定にする必要があります。
      

   --use-presigned-request
      シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      1.59未満のバージョンのrcloneでは、シングルパートオブジェクトをアップロードするために署名付きリクエストを使用します。これは特殊な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時刻でファイルのバージョンを表示します。
      
      パラメータには、日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはそれ以前の期間「100d」または「1h」を指定します。
      
      ただし、これを使用するとファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      設定すると、gzipでエンコードされたオブジェクトを解凍します。
      
      S3に「Content-Encoding: gzip」が設定されたオブジェクトをアップロードすることも可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信時に「Content-Encoding: gzip」でこれらのファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容が解凍されます。
      

   --might-gzip
      バックエンドでオブジェクトをgzipする可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされる際には変更しないはずです。`Content-Encoding: gzip`でアップロードされていないオブジェクトには設定されません。
      
      ただし、一部のプロバイダは（Cloudflareなど）`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzipする場合があります。
      
      これを設定すると、rcloneがContent-Encoding: gzipが設定されておりチャンク送信であるオブジェクトをダウンロードすると、rcloneはオブジェクトを逐次解凍します。
      
      unsetに設定されている場合（デフォルト）、rcloneは提供者の設定に従って適用するものを選択しますが、ここでrcloneの選択肢を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


オプション:
   --access-key-id value        AWSのAccess Key IDです。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用されるCanned ACLです。 [$ACL]
   --endpoint value             Qiniuオブジェクトストレージのエンドポイントです。 [$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します（環境変数または環境変数が存在しない場合はEC2/ECSメタデータから）（デフォルト：false） [$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  リージョンと一致する必要がある場所拘束条件です。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョンです。 [$REGION]
   --secret-access-key value    AWSのSecret Access Key（パスワード）です。 [$SECRET_ACCESS_KEY]
   --storage-class value        Qiniuに新しいオブジェクトを保存する際に使用するストレージクラスです。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるCanned ACLです。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズです。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための上限です。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定すると、gzipでエンコードされたオブジェクトを解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しないでください。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディングです。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストごとの応答リスト）のサイズです。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0は自動です。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パーツ数です。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールのフラッシュが行われる頻度です。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドでオブジェクトをgzipする可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                オブジェクトの存在を確認したり、バケットを作成したりしません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトをHEADして整合性を確認しない場合に設定します。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADを行わない場合に設定します。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有の認証情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value            AWSセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有の認証情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行数です。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるための上限です。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時刻でファイルのバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (デフォルト: false) [$VERSIONS]

```
{% endcode %}