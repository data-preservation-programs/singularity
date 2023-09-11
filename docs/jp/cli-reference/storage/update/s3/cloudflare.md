# Cloudflare R2 ストレージ

{% code fullWidth="true" %}
```
名前:
   singularity storage update s3 cloudflare - Cloudflare R2 ストレージ

使用法:
   singularity storage update s3 cloudflare [コマンドオプション] <名前|ID>

説明:
   --env-auth
      ランタイムからAWS認証情報を取得します（環境変数またはエンタープライズメタデータ）。

      「access_key_id」と「secret_access_key」が空の場合のみ有効です。

      例：
         | false | 次の手順でAWS認証情報を入力します。
         | true  | 環境変数（またはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSアクセスキーID。

      匿名アクセスまたはランタイム認証情報の場合は空のままにします。

   --secret-access-key
      AWS Secret Access Key（パスワード）。

      匿名アクセスまたはランタイム認証情報の場合は空のままにします。

   --region
      接続するリージョン。

      例：
         | auto | R2バケットは低レイテンシのためにCloudflareのデータセンター全体に自動的に分散します。

   --endpoint
      S3 APIのエンドポイント。

      S3クローンを使用する場合は必須です。

   --bucket-acl
      バケットを作成する際に使用するCanned ACL。

      詳細については[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      なお、このACLはバケット作成時にのみ適用されます。設定されていない場合は "acl" が使用されます。

      "acl"および"bucket_acl"が空の文字列の場合、"X-Amz-Acl:"ヘッダは追加されず、デフォルト（private）が使用されます。

      例：
         | private            | オーナーにはFULL_CONTROLが付与されます。
         |                    | 他のユーザーはアクセス権限を持たない（デフォルト）。
         | public-read        | オーナーにはFULL_CONTROLが付与されます。
         |                    | AllUsersグループには読み取りアクセス権限が付与されます。
         | public-read-write  | オーナーにはFULL_CONTROLが付与されます。
         |                    | AllUsersグループには読み取りおよび書き込みアクセス権限が付与されます。
         |                    | バケットでこれを許可することは一般的に推奨されていません。
         | authenticated-read | オーナーにはFULL_CONTROLが付与されます。
         |                    | AuthenticatedUsersグループには読み取りアクセス権限が付与されます。

   --upload-cutoff
      チャンクアップロードに切り替えるためのサイズの閾値。

      このサイズより大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffを超えるサイズのファイル、またはサイズが不明のファイル（「rclone rcat」や「rclone mount」、Google フォトやGoogle ドキュメントからアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。

      注意として、"--s3-upload-concurrency"のチャンクは、転送ごとにメモリ上にバッファされます。

      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合はこの値を増やすと転送速度が向上します。

      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、10,000チャンクの制限を下回るように、チャンクサイズを自動的に増やします。

      サイズ不明のファイルは設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で10,000のチャンクがあるため、ストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、進行状況統計の精度が低下します。Rcloneは、AWS SDKによってバッファリングされたチャンクが送信されたときにチャンクが送信されたとみなし、まだアップロード中である可能性があります。大きなチャンクサイズは、AWS SDKのバッファと進行状況報告が真実から逸脱する可能性があります。

   --max-upload-parts
      マルチパートアップロードでの最大パーツ数。

      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。

      AWS S3の10,000パーツの仕様をサポートしていないサービスに役立ちます。

      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、このパーツの数の制限を下回るようにチャンクサイズを自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのサイズの閾値。

      この閾値を超えるファイルのサーバーサイドコピーは、このサイズのチャンクでコピーされます。

      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加します。これにより、データの整合性チェックが行われますが、大きなファイルのアップロードでは長時間の遅延が発生する可能性があります。

   --shared-credentials-file
      共有認証情報ファイルのパス。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。

      空の場合、環境変数 "AWS_PROFILE" または "default"（環境変数が設定されていない場合）がデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクの数です。

      高速リンクで少数の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用しない場合は、この数値を増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合はパス形式でアクセスします。falseの場合は仮想ホスト形式を使用します。

      これがtrue（デフォルト）の場合、rcloneはパス形式でアクセスします。falseの場合、rcloneは仮想パス形式を使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneは、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合はv2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4署名が機能しない場合にのみ、この設定を使用します（たとえば、Jewel/v10 CEPHより前のバージョン）。

   --list-chunk
      リストチャンクのサイズ（各ListObjects S3リクエストのレスポンスリスト）。

      このオプションはAWS S3の仕様では"MaxKeys"、"max-items"、"page-size"としても知られています。
      ほとんどのサービスでは、リクエストに対して1000個以上要求しても、応答リストが1000個に切り捨てられます。
      AWS S3では、これはグローバルな最大値であり、変更できません。詳細については[こちら](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこの値を増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1, 2または0（自動）。

      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するListObjects呼び出ししか提供されませんでした。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高速であり、できるだけ使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダの設定に基づいてどのリストオブジェクトメソッドを呼び出すかを推測します。正しく推測できない場合は、ここで手動で設定できます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダは、リストをURLエンコードしている場合があります。これが利用可能な場合、ファイル名に制御文字を使用するときは、これがより信頼性があります。これがunsetに設定されている場合（デフォルトの場合）、rcloneはプロバイダの設定に応じて適用する内容を選択しますが、ここでrcloneの選択をオーバーライドすることもできます。

   --no-check-bucket
      バケットの存在をチェックせず、作成を試みません。

      バケットが既に存在する場合、rcloneが行うトランザクションの数を最小限にするために便利です。

      また、使用するユーザーにバケット作成権限がない場合にも必要です。v1.52.0以前では、このバグによりサイレントにパスしたことになりました。

   --no-head
      アップロードしたオブジェクトをHEADリクエストで確認しません。

      rcloneは、PUTでオブジェクトをアップロードした後、200 OKのメッセージを受け取った場合、正しくアップロードされたと想定します。

      特に、次の項目を前提とします。

      - メタデータ（modtime、storage class、コンテンツタイプ）がアップロード時と同じであること
      - サイズがアップロード時と同じであること

      単一パートのPUTの応答から次の項目を読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらの項目は読み取りません。

      サイズ不明のソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを実行します。

      このフラグを設定すると、正規の運用には推奨されないため、アップロードの障害を見逃す可能性が増します。実際には、このフラグがあっても、アップロードの障害が見逃される可能性は非常に小さいです。

   --no-head-object
      GETの前にHEADを行わないでオブジェクトを取得する場合に設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。

      追加バッファ（たとえばマルチパートを必要とするアップロード）は、割り当てにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（特にminio）バックエンドとHTTP/2に関する解決されていない問題があります。s3バックエンドではHTTP/2がデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されたときには、このフラグは削除されます。

      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3ではCloudFront CDNのURLが設定されます。AWS S3は、CloudFrontネットワークを介してダウンロードされたデータの割安なイーグレスを提供します。

   --use-multipart-etag
      マルチパートアップロードでETagを検証に使用するかどうか。

      これはtrue、false、またはデフォルトのプロバイダの設定（空に設定します）にする必要があります。

   --use-presigned-request
      マルチパート以外のアップロードで、署名済みリクエストまたはPutObjectを使用するかどうか。

      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン1.59未満では、単一のパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは特殊な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリスティングに古いバージョンを含めます。

   --version-at
      指定した時間時点のファイルバージョンを表示します。

      パラメータは日付（ "2006-01-02" ）、日時（ "2006-01-02 15:04:05" ）、またはその時間前の期間（ "100d" または "1h" ）である必要があります。

      このオプションでは、ファイルの書き込み操作は許可されません。したがって、ファイルをアップロードしたり削除したりすることはできません。

      有効なフォーマットについては、[timeオプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      これが設定されている場合、gzipで圧縮されたオブジェクトを復号化します。

      S3に "Content-Encoding: gzip" を設定してオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されていると、rcloneはこれらのファイルを受信したときに "Content-Encoding: gzip" で展開します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルのコンテンツは解凍されます。

   --might-gzip
      バックエンドでオブジェクトをgzip圧縮する可能性がある場合にこれを設定します。

      通常、プロバイダはダウンロード時にオブジェクトを変更しません。 "Content-Encoding: gzip" が設定されていない場合、ダウンロード時にセットされません。

      ただし、一部のプロバイダは、明示的に "Content-Encoding: gzip" でないとアップロードされていないにもかかわらず、オブジェクトをgzip形式で保管する場合があります（たとえばCloudflare）。

      これによる症状は、次のようなエラーを受け取ることです。

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      このフラグを設定し、rcloneがContent-Encoding: gzipおよびチャンク転送エンコーディングでオブジェクトをダウンロードした場合、rcloneはオブジェクトを逐次解凍します。

      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に応じて適用する内容を選択しますが、ここでrcloneの選択をオーバーライドすることもできます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する


オプション:
   --access-key-id value      AWSアクセスキーID。 [$ACCESS_KEY_ID]
   --endpoint value           S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                 ランタイムからAWS認証情報を取得します（環境変数またはエンタープライズメタデータ）。（デフォルト：false） [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --region value             接続するリージョン。 [$REGION]
   --secret-access-key value  AWS Secret Access Key（パスワード）。 [$SECRET_ACCESS_KEY]

   詳細オプション

   --bucket-acl value               バケットを作成する際に使用するCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのサイズの閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     これが設定されている場合、gzipで圧縮されたオブジェクトを復号化します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合はパス形式でアクセスします。falseの場合は仮想ホスト形式を使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストチャンクのサイズ。（各ListObjects S3リクエストのレスポンスリスト） (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パーツ数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドでオブジェクトをgzip圧縮する可能性がある場合にこれを設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成を試みません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトをHEADリクエストで確認しません。 (default: false) [$NO_HEAD]
   --no-head-object                 GETの前にHEADを行わないでオブジェクトを取得する場合に設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのサイズの閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを検証に使用するかどうか。 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          マルチパート以外のアップロードで、署名済みリクエストまたはPutObjectを使用するかどうか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合はv2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時間時点のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --decompress                     これが設定されている場合、gzipで圧縮されたオブジェクトを復号化します。 (default: false) [$DECOMPRESS]
   --compression                    trueであれば、圧縮フォーマットの解凍を試みます。 (default: false) [$COMPRESSION]
   --might-compress                 からコピーされたオブジェクトが圧縮形式で保存されている可能性がある場合にこれを設定します。 (default: "unset") [$MIGHT_COMPRESS]
   --no-fast-list                   バケット内のオブジェクトの数を高速に算出せず、内部イテレータを使用する。 [$NO_FAST_LIST]
   --no-list-values                 オブジェクトの値をリスト時に読み込まない。 [$NO_LIST_VALUES]
   --no-noop                        --dry-runに対して "Noop" モードを無効にします。 (default: false) [$NO_NOOP]
   --no-partial-upload              アップロードが失敗したときに、中途で切り捨てたオブジェクトを残さずに削除します。 (default: false) [$NO_PARTIAL_UPLOAD]
   --no-write-headers               ヘッダを書き込まないです。 (default: false) [$NO_WRITE_HEADERS]
   --progress-interval value        プログレスバーの更新間隔。 (default: 500ms) [$PROGRESS_INTERVAL]
```
{% endcode %}