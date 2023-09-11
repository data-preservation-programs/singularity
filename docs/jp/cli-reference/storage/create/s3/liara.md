# Liara オブジェクトストレージ

{% code fullWidth="true" %}
```
名前:
   singularity storage create s3 liara - Liara オブジェクトストレージ

使用法:
   singularity storage create s3 liara [コマンドオプション] [引数...]

説明:
   --env-auth
      AWSの認証情報を実行時に取得します（環境変数または環境変数がない場合はEC2/ECSメタデータから）。

      access_key_idとsecret_access_keyが空の場合のみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたは実行時の認証情報の場合は空にします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたは実行時の認証情報の場合は空にします。

   --endpoint
      LiaraオブジェクトストレージAPIのエンドポイント。

      例:
         | storage.iran.liara.space | デフォルトのエンドポイント
         |                          | イラン

   --acl
      バケットの作成やオブジェクトの保存またはコピー時に使用するCanned ACL。

      このACLはオブジェクトの作成時およびbucket_aclが設定されていない場合も使用されます。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      ソースからのサーバーサイドコピー時にこのACLは適用されますが、S3はコピー元からACLをコピーするのではなく、新しいACLを書き込みます。

      aclが空の場合、X-Amz-Acl:ヘッダは追加されず、デフォルトのプライベートが使用されます。

   --bucket-acl
      バケットの作成時に使用するCanned ACL。

      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。

      このACLはバケットの作成時のみ適用されます。設定されていない場合は「acl」が代わりに使用されます。

      "acl"と"bucket_acl"が空の文字列である場合、ヘッダX-Amz-Acl:は追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーにはFULL_CONTROLが付与されます。
         |                    | 他のユーザーにはアクセス権限がありません（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROLが付与されます。
         |                    | AllUsersグループには読み取りアクセスがあります。
         | public-read-write  | オーナーにはFULL_CONTROLが付与されます。
         |                    | AllUsersグループには読み取りおよび書き込みアクセスがあります。
         |                    | バケットにこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーにはFULL_CONTROLが付与されます。
         |                    | AuthenticatedUsersグループには読み取りアクセスがあります。

   --storage-class
      Liaraで新しいオブジェクトを保存する際に使用するストレージクラス

      例:
         | STANDARD | 標準のストレージクラス

   --upload-cutoff
      チャンク化アップロードに切り替えるためのカットオフ。

      これより大きなファイルはchunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffより大きなファイルやサイズが不明なファイル（「rclone rcat」や「rclone mount」またはGoogleフォトやGoogleドキュメントからアップロードされたファイルなど）は、このチャンクサイズを使用して複数パートのアップロードとしてアップロードされます。

      注："--s3-upload-concurrency"このサイズのチャンクは、
      メモリごとに転送ごとにバッファリングされます。

      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを増やすと転送速度が向上します。

      Rcloneは、既知のサイズの大きなファイルをアップロードするときに10,000チャンクの制限を下回るように自動的にチャンクサイズを増やします。

      未知のサイズのファイルは、設定された
      チャンクサイズでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で
      10,000チャンクの制限があります。したがって、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、チャンクサイズを増やす必要があります。

      チャンクサイズを増やすと、「-P」フラグで表示される進行状況の統計の正確さが低下します。 Rcloneは、AWS SDKによってバッファリングされたチャンクを送信したときにチャンクが送信されたとみなし、実際にはまだアップロード中の場合があります。チャンクサイズが大きいほど、AWS SDKのバッファも大きくなり、真実から逸脱した進行状況の報告が増えます。

   --max-upload-parts
      マルチパートアップロードでのパートの最大数。

      このオプションは、マルチパートアップロード時に使用する最大のマルチパートチャンク数を定義します。

      これは10,000チャンクのAWS S3仕様をサポートしていない場合に役立ちます。

      Rcloneは、既知のサイズの大きなファイルをアップロードするときにチャンクサイズを自動的に増やして、このチャンク数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。

      サーバーサイドコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算して、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまでには長い遅延が生じることがあります。

   --shared-credentials-file
      共有の認証情報ファイルへのパス。

      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは
      "AWS_SHARED_CREDENTIALS_FILE"環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\\.aws\\credentials"

   --profile
      共有の認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有の認証情報ファイルを使用できます。この変数は、そのファイルで使用するプロファイルを制御します。

      空の場合は環境変数"AWS_PROFILE"または
      "default"という環境変数が設定されていない場合にデフォルト値になります。

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。

      同じファイルのチャンク数を同時にアップロードします。

      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合は、これを増やすと転送速度が向上する可能性があります。

   --force-path-style
      trueの場合はパススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルのアクセスを使用します。

      true（デフォルト）の場合、rcloneはパススタイルのアクセスを使用し、falseの場合は仮想パススタイルを使用します。詳細については、[AWS S3
      ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、またはTencent COS）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2の認証を使用します。

      false（デフォルト）の場合、rcloneはv4の認証を使用します。それが設定されている場合、rcloneはv2の認証を使用します。

      v4署名が機能しない場合にのみこれを使用します（たとえば、Jewel/v10 CEPH以前）。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストの応答リストのサイズ）。

      このオプションは、AWS S3の仕様の「MaxKeys」、「max-items」、または「page-size」とも呼ばれます。
      ほとんどのサービスは、要求されるオブジェクト以上のリストを1000個に切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン: 1、2、または0（自動）。

      S3が最初にリリースされた当初は、バケット内のオブジェクトを列挙するためのListObjectsコールのみが提供されました。

      しかし、2016年5月にはListObjectsV2コールが導入されました。これは高いパフォーマンスで、可能な限り使用する必要があります。

      デフォルトの0に設定されている場合、rcloneは提供者に基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が誤っている場合は、ここで手動で設定できます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/未設定

      一部のプロバイダはリストのURLエンコードをサポートし、ファイル名の制御文字を使用する場合にこれがより信頼性のある方法です。これが未設定（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットが存在するか、作成する試みを行わない場合に設定します。

      バケットが既に存在している場合、rcloneが実行するトランザクションの数を最小限に抑えるために有用です。

      バケット作成の権限がない場合も必要です。v1.52.0以前では、これはバグのために黙ってパスされたでしょう。

   --no-head
      アップロードしたオブジェクトのHEADを行って整合性をチェックしません。

      rcloneは、PUT後に200 OKメッセージを受け取った場合、正しくアップロードされたと見なします。

      特に次のことを想定しています。

      - アップロード時のメタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロードされたものと同じであること
      - サイズがアップロードされたものと同じであること

      シングルパートのPUTの場合、次の項目を応答から読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらの項目は読み取られません。

      長さが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出される確率が高まります。特にサイズが正しくない場合です。したがって、通常の操作ではお勧めしません。実際には、このフラグがあっても、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GETの前にHEADを行わない場合に設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、概要の[encodingセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。

      追加バッファが必要なアップロード（マルチパートなど）では、メモリプールが割り当てのために使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2には解決されていない問題があります。S3バックエンドのデフォルトではHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決された場合、このフラグは削除されます。

      参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロードのためのカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してデータをダウンロードすることで、データのダウンロードに対してより安価な輸送方法を提供します。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか

      これはtrue、false、またはデフォルト（プロバイダに応じた）を使用する必要があります。

   --use-presigned-request
      単一パートアップロードに対して署名付きリクエストまたはPutObjectを使用するかどうか

      falseの場合、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン<1.59は、単一パートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定するとその機能が再度有効になります。これは例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時間のファイルバージョンを表示します。

      パラメータは日付（"2006-01-02"）、日時（"2006-01-02
      15:04:05"）、またはその長さにすることができます、例えば "100d" や "1h" です。

      このオプションを使用すると、ファイルの書き込み操作は許可されないため、ファイルをアップロードしたり削除したりすることはできません。

      有効な形式については、[timeオプションドキュメント](/docs/#time-option)を参照してください。

   --decompress
      gzipでエンコードされたオブジェクトを解凍する場合はこのフラグを設定します。

      S3へのアップロードで「Content-Encoding: gzip」が設定されているオブジェクトを、通常rcloneは圧縮されたオブジェクトとしてダウンロードします。

      このフラグを設定すると、rcloneはこれらのオブジェクトを受け取るときに「Content-Encoding: gzip」で解凍します。これにより、rcloneはサイズとハッシュを確認できませんが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzipする可能性がある場合に設定します。

      通常、プロバイダはオブジェクトがダウンロードされるときにオブジェクトを変更しません。 `Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時にはそれも設定されません。

      ただし、いくつかのプロバイダは、 `Content-Encoding: gzip`でアップロードされていないオブジェクトもgzipすることがあります（たとえばCloudflare）。

      これを設定し、rcloneがContent-Encoding: gzipが設定されたオブジェクトとチャンクされた転送エンコードをダウンロードすると、rcloneはオブジェクトをリアルタイムで解凍します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムのメタデータの設定と読み取りを抑制する


オプション:
   --access-key-id value      AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                バケットの作成やオブジェクトの保存またはコピー時に使用するCanned ACL。 [$ACL]
   --endpoint value           LiaraオブジェクトストレージAPIのエンドポイント。 [$ENDPOINT]
   --env-auth                 AWSの認証情報を実行時に取得します（環境変数または環境変数がない場合はEC2/ECSメタデータから）。 （デフォルト：false） [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --storage-class value      Liaraで新しいオブジェクトを保存する際に使用するストレージクラス [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用するCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (デフォルト："5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (デフォルト："4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzipでエンコードされたオブジェクトを解凍する場合はこのフラグを設定します。 (デフォルト：false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト：false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。 (デフォルト：false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードのためのカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト："Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルのアクセスを使用します。 (デフォルト：true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストの応答リストのサイズ）。 (デフォルト：1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/未設定 (デフォルト："unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1,2または0（自動）。 (デフォルト：0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでのパートの最大数。 (デフォルト：10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (デフォルト："1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (デフォルト：false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipする可能性がある場合に設定します。 (デフォルト："unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットが存在するか、作成する試みを行わない場合に設定します。 (デフォルト：false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトのHEADを行って整合性をチェックしません。 (デフォルト：false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わない場合に設定します。 (デフォルト：false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムのメタデータの設定と読み取りを抑制する (デフォルト：false) [$NO_SYSTEM_METADATA]
   --profile value                  共有の認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有の認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。 (デフォルト：4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのカットオフ。 (デフォルト："200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (デフォルト："unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          単一パートアップロードに対して署名付きリクエストまたはPutObjectを使用するかどうか (デフォルト：false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2の認証を使用します。 (デフォルト：false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (デフォルト："off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (デフォルト：false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}