# Arvan Cloud オブジェクトストレージ（AOS）

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 arvancloud - Arvan Cloud オブジェクトストレージ（AOS）の設定を更新します

USAGE:
   singularity storage update s3 arvancloud [command options] <name|id>

DESCRIPTION:
   --env-auth
      実行時にAWSの認証情報（環境変数またはEC2/ECSメタデータ）を取得します。
      
      access_key_idとsecret_access_keyが空白の場合のみ適用されます。

      例:
         | false | 次の手順でAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーIDです。

      匿名アクセスまたは実行時の認証情報の場合は空白にしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたは実行時の認証情報の場合は空白にしてください。

   --endpoint
      Arvan Cloud オブジェクトストレージ（AOS）APIのエンドポイントです。

      例:
         | s3.ir-thr-at1.arvanstorage.com | デフォルトのエンドポイント - よくわからない場合はこれを選択してください。
         |                                | イラン、テヘラン (Asiatech)
         | s3.ir-tbz-sh1.arvanstorage.com | イラン、タブリーズ (Shahriar)

   --location-constraint
      エンドポイントに対応した位置制約です。

      バケット作成時のみ使用されます。

      例:
         | ir-thr-at1 | イラン、テヘラン (Asiatech)
         | ir-tbz-sh1 | イラン、タブリーズ (Shahriar)

   --acl
      バケットの作成やオブジェクトの保存やコピーに使用する事前設定されたACLです。

      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合、バケットの作成にも使用されます。

      詳細はこちらを参照してください：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      S3はサーバーサイドでオブジェクトのコピーを行う際にACLをコピーしないため、このACLはS3が新しいACLを書き込みます。

      もしACLが空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトの（プライベート）が使用されます。

   --bucket-acl
      バケットの作成時に使用される事前設定されたACLです。

      詳細はこちらを参照してください：https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl

      このACLはバケットの作成時のみ適用されます。設定されていない場合は「acl」が代わりに使用されます。

      もし「acl」と「bucket_acl」が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトの（プライベート）が使用されます。

      例:
         | private            | オーナーにFULL_CONTROL権限が与えられます。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにREADアクセス権限が与えられます。
         | public-read-write  | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AllUsersグループにREADおよびWRITEアクセス権限が与えられます。
         |                    | バケットでこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーにFULL_CONTROL権限が与えられます。
         |                    | AuthenticatedUsersグループにREADアクセス権限が与えられます。

   --storage-class
      ArvanCloudに新しいオブジェクトを保存する際に使用するストレージクラスです。

      例:
         | STANDARD | スタンダードストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるカットオフです。

      これよりも大きなファイルは、chunk_sizeごとにチャンクとしてアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。

      upload_cutoffよりも大きなファイルまたはサイズ不明のファイル（「rclone rcat」で作成されたファイルや「rclone mount」やGoogleフォトやGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      ことに注意してください、"--s3-upload-concurrency"サイズのチャンクは転送ごとにメモリにバッファリングされます。

      高速リンクで大きなファイルを転送しており、メモリが十分にある場合は、これを増やすと転送が高速化されます。

      rcloneは、10,000個のチャンクの制限を超えないように、既知のサイズの大きなファイルをアップロードする場合は自動的にチャンクサイズを増やします。

      サイズが不明なファイルは、構成されたチャンクサイズでアップロードされます。デフォルトのチャンクサイズが5 MiBで最大で10,000個のチャンクがあり、これによりデフォルトでストリームアップロードできるファイルの最大サイズは48 GiBです。もし大きなファイルをストリームアップロードしたい場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを増やすと、進行状況の統計情報の正確性が低下します。rcloneは、チャンクがAWS SDKによってバッファリングされると送信されたと見なし、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKバッファのサイズが大きくなり、進行状況の報告が実際の状況と乖離する場合があります。

   --max-upload-parts
      マルチパートアップロードでのパート（チャンク）の最大数です。

      このオプションは、マルチパートアップロードを行う場合に使用されるマルチパートチャンクの最大数を定義します。

      サービスがAWS S3の10,000個のチャンクの仕様をサポートしない場合に役立ちます。

      rcloneは、既知のサイズの大きなファイルをアップロードする場合は、このチャンク数の制限を超えないように、自動的にチャンクサイズを増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるカットオフです。

      サーバーサイドでコピーする必要があり、これよりも大きいファイルは、このサイズのチャンクでコピーされます。

      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しないでください。

      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には時間がかかる場合があります。

   --shared-credentials-file
      共有の資格情報ファイルへのパスです。

      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルト値になります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有資格情報ファイルで使用するプロファイルです。

      env_auth = trueの場合、rcloneは共有の資格情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。

      値が空の場合、デフォルトでは「AWS_PROFILE」または「default」環境変数が使用されます。

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートアップロードの並行性です。

      同じファイルのチャンクを同時にアップロードする数です。

      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすことで転送を高速化することができます。

   --force-path-style
      trueの場合、パス形式でアクセス。falseの場合、仮想ホスト方式でアクセスします。

      これがtrue（デフォルト）の場合、rcloneはパス形式でアクセスしますが、falseの場合は仮想パス形式を使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、この設定によってfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。

      これがfalse（デフォルト）に設定されている場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4署名が機能しない場合にのみ使用してください。Jewel/v10 CEPH。

   --list-chunk
      リストのサイズ（各ListObject S3リクエストのレスポンスリスト）です。

      このオプションは、AWS S3の仕様では「MaxKeys」「max-items」「page-size」としても知られています。ほとんどのサービスは、要求されたよりも多くの場合でも、レスポンスリストを1000オブジェクトに切り詰めます。AWS S3では、これはグローバルな最大値であり、変更できません。詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。Cephでは、"rgw list buckets max chunk"オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1,2、または0は自動設定です。

      S3が最初に登場したとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されました。

      しかし、2016年5月にはListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを持ち、可能であれば使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダに基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が間違っている場合は、ここで手動で設定することができます。

   --list-url-encode
      リストのURLエンコードを行うかどうか：true/false/unset

      一部のプロバイダは、ファイル名に制御文字を含める場合にURLエンコードリストをサポートしています。これが使用可能な場合、これはファイル名の制御文字を使用する場合により信頼性が向上します。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、rcloneの選択をオーバーライドすることもできます。

   --no-check-bucket
      該当のバケットの存在をチェックせず、作成しようとしない。

      バケットが既に存在する場合、rcloneが行うトランザクションの数を最小限に抑えることができるため、これが有用である場合があります。

      バケットの作成権限がない場合も必要になる場合があります。v1.52.0以下のバージョンでは、バグのためにこれがサイレントにパスされます。

   --no-head
      UPLOADEDオブジェクトをHEADして整合性をチェックしない。

      rcloneは通常、PUTでオブジェクトをアップロードした後に整合性を確認するためにHEADリクエストを行います。

      これはrcloneがPUTをした後に200 OKメッセージを受け取った場合、正しくアップロードされたとみなします。

      特に以下を仮定します。

      - アップロード時のメタデータ、モディファイタイム、ストレージクラス、コンテンツタイプがアップロード時と同じであること
      - サイズがアップロード時と同じであること

      1個のパートPUTの場合、以下のアイテムをレスポンスから読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらのアイテムは読み取られません。

      ソースオブジェクトのサイズが不明な場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの際の検出されていないアップロードエラーの可能性が高まるため、通常の操作には推奨されません。実際には、このフラグを設定しても、アップロードエラーが検出されない可能性は非常に低いです。

   --no-head-object
      GETの前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度です。

      追加のバッファが必要なアップロード（マルチパートなど）では、メモリプールを使用して割り当てられます。このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効化します。

      s3（特にminio）バックエンドとHTTP/2の問題がまだ解決されていません。デフォルトでS3バックエンドではHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決されるまで、このフラグは削除されません。

      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイントです。通常、AWS S3はCloudFrontのCDN URLに設定され、CloudFrontネットワークを介してダウンロードされたデータの転送エコノミーを提供します。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      これはtrue、false、またはデフォルト（プロバイダに設定されている値）のいずれかで設定する必要があります。

   --use-presigned-request
      単一パートのアップロードの場合、署名済みリクエストまたはPutObjectを使用するかどうか。

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン< 1.59では、単一パートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定することでその機能を再有効化することができます。これは、例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定された時点のファイルバージョンを表示します。

      パラメータは、日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはその時間前の期間、「100d」または「1h」などです。

      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードや削除することはできません。

      使用可能な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定する場合、gzipでエンコードされたオブジェクトを展開します。

      S3に「Content-Encoding: gzip」が設定されたファイルをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に「Content-Encoding: gzip」で解凍します。これにより、rcloneはサイズとハッシュを確認できませんが、ファイルの内容が解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzipする可能性がある場合に設定します。

      通常、プロバイダはオブジェクトがダウンロードされると変更しません。オブジェクトが「Content-Encoding: gzip」でアップロードされていない場合、ダウンロード時に設定されません。

      ただし、一部のプロバイダは、`Content-Encoding: gzip`でアップロードされていないオブジェクトをgzipする場合があります（例：Cloudflare）。

      これにより、エラーメッセージが次のように表示されることがあります。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、チャンク転送エンコーディングを使用してオブジェクトをダウンロードした場合、rcloneはオブジェクトをリアルタイムに解凍します。

      unsetに設定された場合（デフォルト）、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択をオーバーライドすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します。


OPTIONS:
   --access-key-id value        AWSのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --bucket-acl value           バケットの作成時に使用される事前設定されたACLです。 [$BUCKET_ACL]
   --endpoint value             Arvan Cloud オブジェクトストレージ（AOS）APIのエンドポイントです。 [$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します。 (default: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  エンドポイントに対応した位置制約です。 [$LOCATION_CONSTRAINT]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --storage-class value        ArvanCloudに新しいオブジェクトを保存する際に使用するストレージクラスです。 [$STORAGE_CLASS]

   Advanced

   --acl value                  バケットの作成やオブジェクトの保存やコピーに使用する事前設定されたACLです。 [$ACL]
   --chunk-size value           アップロードに使用するチャンクサイズです。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value          マルチパートコピーに切り替えるカットオフです。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                 設定する場合、gzipでエンコードされたオブジェクトを展開します。 (default: false) [$DECOMPRESS]
   --disable-checksum           オブジェクトのメタデータにMD5チェックサムを保存しないでください。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2              S3バックエンドでhttp2の使用を無効化します。 (default: false) [$DISABLE_HTTP2]
   --download-url value         ダウンロード用のカスタムエンドポイントです。 [$DOWNLOAD_URL]
   --encoding value             バックエンドのエンコーディングです。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style           trueの場合、パス形式でアクセス。falseの場合、仮想ホスト方式でアクセスします。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value           リストのサイズ（各ListObject S3リクエストのレスポンスリスト）です。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value      リストのURLエンコードを行うかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value         使用するListObjectsのバージョン：1,2、または0は自動設定です。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value     マルチパートアップロードでのパート（チャンク）の最大数です。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value  内部メモリバッファプールがフラッシュされる頻度です。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap       内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value           バックエンドがオブジェクトをgzipする可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket            該当のバケットの存在をチェックせず、作成しようとしない。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                    UPLOADEDオブジェクトをHEADして整合性をチェックしない。 (default: false) [$NO_HEAD]
   --no-head-object             GETの前にHEADを行わない場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata         システムメタデータの設定と読み取りを抑制します。 (default: false) [$NO_SYSTEM_METADATA]
   --profile value              共有資格情報ファイルで使用するプロファイルです。 [$PROFILE]
   --session-token value        AWSセッショントークンです。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有の資格情報ファイルへのパスです。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value   マルチパートアップロードの並行性です。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value        チャンクアップロードに切り替えるカットオフです。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value   マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request      単一パートのアップロードの場合、署名済みリクエストまたはPutObjectを使用するかどうか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                    trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value           指定された時点のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                   ディレクトリリストに古いバージョンを含める。 (default: false) [$VERSIONS]

```
{% endcode %}