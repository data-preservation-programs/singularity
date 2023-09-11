# RackCorp オブジェクトストレージ

{% code fullWidth="true" %}
```
名称:
   singularity storage create s3 rackcorp - RackCorp オブジェクトストレージ

使用法:
   singularity storage create s3 rackcorp [コマンドオプション] [引数...]

説明:
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータから取得）。
      
      access_key_idとsecret_access_keyが未入力（空）の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力してください。
         | true  | 環境からAWSの認証情報（環境変数またはIAM）を取得します。

   --access-key-id
      AWSのアクセスキーIDです。
      
      匿名アクセスや実行時の認証情報を使用する場合は、空白のままにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。
      
      匿名アクセスや実行時の認証情報を使用する場合は、空白のままにしてください。

   --region
      バケットを作成しデータを格納する場所です。
      

      例:
         | global    | グローバルCDN（すべてのロケーション）リージョン
         | au        | オーストラリア（すべての地域）
         | au-nsw    | NSW（オーストラリア）リージョン
         | au-qld    | QLD（オーストラリア）リージョン
         | au-vic    | VIC（オーストラリア）リージョン
         | au-wa     | パース（オーストラリア）リージョン
         | ph        | マニラ（フィリピン）リージョン
         | th        | バンコク（タイ）リージョン
         | hk        | 香港（香港）リージョン
         | mn        | ウランバートル（モンゴル）リージョン
         | kg        | ビシュケク（キルギス）リージョン
         | id        | ジャカルタ（インドネシア）リージョン
         | jp        | 東京（日本）リージョン
         | sg        | SG（シンガポール）リージョン
         | de        | フランクフルト（ドイツ）リージョン
         | us        | USA（AnyCast）リージョン
         | us-east-1 | ニューヨーク（アメリカ）リージョン
         | us-west-1 | フリーモント（アメリカ）リージョン
         | nz        | オークランド（ニュージーランド）リージョン

   --endpoint
      RackCorpオブジェクトストレージのエンドポイントです。

      例:
         | s3.rackcorp.com           | グローバル（AnyCast）エンドポイント
         | au.s3.rackcorp.com        | オーストラリア（Anycast）エンドポイント
         | au-nsw.s3.rackcorp.com    | シドニー（オーストラリア）エンドポイント
         | au-qld.s3.rackcorp.com    | ブリスベン（オーストラリア）エンドポイント
         | au-vic.s3.rackcorp.com    | メルボルン（オーストラリア）エンドポイント
         | au-wa.s3.rackcorp.com     | パース（オーストラリア）エンドポイント
         | ph.s3.rackcorp.com        | マニラ（フィリピン）エンドポイント
         | th.s3.rackcorp.com        | バンコク（タイ）エンドポイント
         | hk.s3.rackcorp.com        | 香港（香港）エンドポイント
         | mn.s3.rackcorp.com        | ウランバートル（モンゴル）エンドポイント
         | kg.s3.rackcorp.com        | ビシュケク（キルギス）エンドポイント
         | id.s3.rackcorp.com        | ジャカルタ（インドネシア）エンドポイント
         | jp.s3.rackcorp.com        | 東京（日本）エンドポイント
         | sg.s3.rackcorp.com        | SG（シンガポール）エンドポイント
         | de.s3.rackcorp.com        | フランクフルト（ドイツ）エンドポイント
         | us.s3.rackcorp.com        | USA（AnyCast）エンドポイント
         | us-east-1.s3.rackcorp.com | ニューヨーク（アメリカ）エンドポイント
         | us-west-1.s3.rackcorp.com | フリーモント（アメリカ）エンドポイント
         | nz.s3.rackcorp.com        | オークランド（ニュージーランド）エンドポイント

   --location-constraint
      バケットが配置され、データが格納される場所です。
      

      例:
         | global    | グローバルCDNリージョン
         | au        | オーストラリア（すべての地域）
         | au-nsw    | NSW（オーストラリア）リージョン
         | au-qld    | QLD（オーストラリア）リージョン
         | au-vic    | VIC（オーストラリア）リージョン
         | au-wa     | パース（オーストラリア）リージョン
         | ph        | マニラ（フィリピン）リージョン
         | th        | バンコク（タイ）リージョン
         | hk        | 香港（香港）リージョン
         | mn        | ウランバートル（モンゴル）リージョン
         | kg        | ビシュケク（キルギス）リージョン
         | id        | ジャカルタ（インドネシア）リージョン
         | jp        | 東京（日本）リージョン
         | sg        | SG（シンガポール）リージョン
         | de        | フランクフルト（ドイツ）リージョン
         | us        | USA（AnyCast）リージョン
         | us-east-1 | ニューヨーク（アメリカ）リージョン
         | us-west-1 | フリーモント（アメリカ）リージョン
         | nz        | オークランド（ニュージーランド）リージョン

   --acl
      オブジェクトとバケットの作成時に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成時にのみ使用されます。bucket_aclが設定されていない場合も適用されます。
      
      詳細については、[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      注意として、このACLはS3のサーバー側でオブジェクトをコピーする場合に適用されます。
      S3はソースのACLをコピーせず、代わりに新しいACLを書き込みます。
      
      ACLが空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。
      
      詳細については、[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      注意として、このACLはバケットの作成時にのみ適用されます。設定されていない場合は"acl"が代わりに使用されます。
      
      この「acl」と「bucket_acl」が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

      例:
         | private            | オーナーはFULL_CONTROLを取得します。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループは読み取り権限を取得します。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                    | AllUsersグループは読み取りと書き込みの権限を取得します。
         |                    | バケットでこの設定を行うことは一般的には推奨されません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                    | AuthenticatedUsersグループは読み取り権限を取得します。

   --upload-cutoff
      チャンクアップロードに切り替えるためのサイズの上限です。
      
      このサイズを超えるファイルは、chunk_sizeのチャンク単位でアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。
      
      upload_cutoffよりも大きなサイズのファイルや、サイズが不明なファイル（たとえば、"rclone rcat"でのアップロードや、"rclone mount"やGoogleフォトまたはGoogleドキュメントでのアップロード）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      注意してくださいが、"--s3-upload-concurrency"はこのサイズのチャンクが転送ごとにメモリ内にバッファリングされます。
      
      高速リンクで大きなファイルを転送している場合に十分なメモリがある場合、これを増やすと転送速度が向上します。
      
      Rcloneは、すでに知られている大きなサイズのファイルをアップロードする際に、10,000個のチャンク制限を下回るように自動的にチャンクサイズを増やします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で10,000個のチャンクがあることから、デフォルトでは最大サイズが48 GiBのファイルをストリームアップロードできます。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、"-P"フラグで表示される進行状況の統計の正確さが低下します。RcloneはチャンクがAWS SDKでバッファリングされた時点でチャンクを送信したとみなし、実際にはまだアップロード中である場合でも進捗情報に反映されます。チャンクサイズが大きいほど、AWS SDKのバッファと進行状況の報告は真実からさらに逸脱します。
      

   --max-upload-parts
      マルチパートアップロードでの最大パート数です。
      
      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。
      
      これは、サービスがAWS S3の10,000チャンクの仕様をサポートしていない場合に役立ちます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際に、このチャンクのサイズを自動的に増やして制限数のチャンクよりも小さい範囲に保ちます。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのサイズの上限です。
      
      サーバーサイドでコピーする必要のあるこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを格納しないようにします。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまで時間がかかります。

   --shared-credentials-file
      共有の認証情報ファイルへのパスです。
      
      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を検索します。環境変数の値が空の場合、デフォルトは現在のユーザーのホームディレクトリになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、デフォルトは環境変数「AWS_PROFILE」または「default」が設定されていない場合になります。
      

   --session-token
      AWSのセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行性です。
      
      これは、同じファイルのチャンクの数を同時にアップロードすることを示します。
      
      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合には、これを増やすと転送速度が向上するかもしれません。

   --force-path-style
      trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式のアクセスを使用します。
      
      true（デフォルト）の場合、rcloneはパス形式のアクセスを使用し、falseの場合は仮想パス形式を使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、
      falseに設定する必要があります。rcloneはプロバイダの設定に基づいて自動的にこれを行います。

   --v2-auth
      trueの場合、v2の認証を使用します。
      
      false（デフォルト）の場合は、v4の認証を使用します。指定された場合、rcloneはv2の認証を使用します。
      
      v4の署名が機能しない場合にのみ、このフラグを使用します。例えば、Jewel/v10 CEPHより前のバージョンです。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストの応答リスト）です。
      
      このオプションは、AWS S3仕様の「MaxKeys」、「max-items」、「page-size」ともして知られています。
      ほとんどのサービスは、リクエスト数がそれ以上の場合でも応答リストを1000オブジェクトに切り詰めます。
      AWS S3では、これはグローバルな制限であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されました。
      
      しかし、2016年5月には、ListObjectsV2呼び出しが導入されました。これははるかに高速であり、可能であれば使用する必要があります。
      
      デフォルトで設定されている0に設定すると、rcloneは提供者の設定に基づいて呼び出すべきlistオブジェクトのメソッドを推測します。推測が間違っている場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダはリストのURLエンコードをサポートしており、これが使用可能な場合、ファイル名に制御文字を使用する場合により信頼性が高まります。これがunsetに設定されている場合（デフォルト）は、プロバイダの設定に基づいてrcloneが適用する内容を選択します。

   --no-check-bucket
      バケットの存在を確認せず、または作成しようとしない場合に設定します。
      
      バケットがすでに存在することを知っている場合、rcloneが行うトランザクション数を最小限に抑えるために有用です。
      
      バケット作成権限を持たないユーザーを使用している場合にも必要になる場合があります。v1.52.0より前のバージョンでは、これはバグのために静かにパスされました。
      

   --no-head
      アップロードしたオブジェクトをHEADして整合性を確認しない場合に設定します。
      
      rcloneがPUTでオブジェクトをアップロードした後に200 OKメッセージを受け取った場合、正常にアップロードされたとみなされます。
      
      特に、次の条件が満たされると仮定します：
      
      - metadata（modtime、storage class、およびcontent typeを含む）はアップロード時のものである。
      - サイズはアップロード時のものである。
      
      単一のパートPUTの応答から次のアイテムを読み取ります：
      
      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合は、これらのアイテムは読み取られません。
      
      サイズの不明なソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が増えます。特に、サイズが正しくない場合です。しかし、このフラグを使用しない場合でも、実際には非常に小さいアップロードの失敗の可能性があります。
      

   --no-head-object
      GETする前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。
      
      追加のバッファを必要とするアップロード（たとえばマルチパート）は、割り当てのためにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドのhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2の問題が解決されていません。HTTP/2はs3バックエンドのデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決された場合、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3はクラウドフロントCDNのURLに設定されます。CloudFrontネットワークを介してデータをダウンロードすると、AWS S3はより安価な出口になります。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルト（未設定）に設定する必要があります。
      

   --use-presigned-request
      シングルパートのアップロードの場合に署名済みのリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rclone 1.59未満のバージョンでは、シングルパートオブジェクトのアップロードに署名付きリクエストを使用し、このフラグをtrueに設定すると、この機能が再び有効になります。これは特殊な状況またはテスト以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは、「2006-01-02」の日付、「2006-01-02 15:04:05」のdatetime、またはその前の期間、「100d」または「1h」のような期間です。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードまたは削除することはできません。
      
      有効な形式については、[時刻オプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      これを設定すると、gzipでエンコードされたオブジェクトを解凍します。
      
      S3へのアップロードでは「Content-Encoding: gzip」が設定されたオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信したときに「Content-Encoding: gzip」を持つこれらのファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。`Content-Encoding: gzip`でアップロードされなかったオブジェクトには設定されません。
      
      ただし、一部のプロバイダは`Content-Encoding: gzip`でアップロードされていない場合でもオブジェクトをgzipで圧縮する場合があります（たとえばCloudflare）。
      
      これにより、次のようなエラーが発生することがあります。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneが`Content-Encoding: gzip`を設定し、チャンクされた転送エンコードを受信すると、rcloneはオブジェクトを逐次解凍します。
      
      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に基づいてどの適用内容を選択するかを決定するが、この設定でrcloneの選択を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定および読み取りを抑制します


オプション:
   --access-key-id value        AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                  オブジェクトとバケットの作成時に使用されるCanned ACL。 [$ACL]
   --endpoint value             RackCorpオブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータから取得）。（デフォルト: false） [$ENV_AUTH]
   --help, -h                   ヘルプを表示する
   --location-constraint value  バケットが配置され、データが格納される場所。 [$LOCATION_CONSTRAINT]
   --region value               バケットを作成しデータを格納する場所。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   アドバンスト

   --bucket-acl value               バケットの作成時に使用されるCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。（デフォルト："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのサイズの上限。（デフォルト："4.656Gi"） [$COPY_CUTOFF]
   --decompress                     これを設定すると、gzipでエンコードされたオブジェクトを解凍します。（デフォルト: false） [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを格納しないようにします。（デフォルト: false） [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。（デフォルト: false） [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト: "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --force-path-style               trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式のアクセスを使用します。（デフォルト: true） [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。（デフォルト: 1000） [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true / false /未設定（デフォルト: "unset"）[$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）（デフォルト: 0） [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パート数：10000（デフォルト）） [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。（デフォルト："1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト: false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する可能性がある場合に設定します。（デフォルト: "unset"） [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しようとしない場合に設定します。（デフォルト: false） [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトをHEADして整合性を確認しない場合に設定します。（デフォルト: false） [$NO_HEAD]
   --no-head-object                 GETする前にHEADを行わない場合に設定します。（デフォルト: false） [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りを抑制します。（デフォルト: false） [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有の認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。(デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのサイズの上限。（デフォルト："200Mi"） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか（デフォルト: "unset"）[$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードの場合に署名済みのリクエストまたはPutObjectを使用するかどうか。（デフォルト: false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2の認証を使用します。（デフォルト: false）[$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。（デフォルト: "off"） [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。（デフォルト: false） [$VERSIONS]

   一般

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}