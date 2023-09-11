# その他のS3互換プロバイダ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 other - その他のS3互換プロバイダ

使用法:
   singularity storage create s3 other [command options] [arguments...]

説明:
   --env-auth
      実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ）。

      access_key_idとsecret_access_keyが空の場合にのみ適用されます。

      例:
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境からAWSの認証情報を取得します（環境変数またはIAM）。

   --access-key-id
      AWSのアクセスキーID。

      匿名アクセスまたは実行時の認証情報の場合は空にしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。

      匿名アクセスまたは実行時の認証情報の場合は空にしてください。

   --region
      接続するリージョン。

      S3のクローンを使用しており、リージョンが存在しない場合は空にしてください。

      例:
         | <unset>            | 迷った場合はこれを使用します。
         |                    | v4シグネチャと空のリージョンを使用します。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用します。
         |                    | Jewel/v10以前のCEPHなど。

   --endpoint
      S3 APIのエンドポイント。

      S3のクローンを使用している場合は必須です。

   --location-constraint
      リージョンに一致する場所の制約。

      確証がない場合は空にしてください。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成とオブジェクトの保存またはコピー時に使用されるプリセットACL。

      このACLはオブジェクトの作成時に使用されます。bucket_aclが設定されていない場合も、バケットの作成時に使用されます。

      詳細については、[AWSのACLの概要](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      サーバサイドでオブジェクトをコピーする場合、S3はソースからACLをコピーしないため、新しいACLを書き込みます。

      aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（プライベート）が使用されます。

   --bucket-acl
      バケットの作成時に使用されるプリセットACL。

      詳細については、[AWSのACLの概要](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。

      このACLはバケットの作成時のみに適用されます。設定されていない場合は「acl」が代わりに使用されます。

      aclとbucket_aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（プライベート）が使用されます。

      例:
         | private            | オーナーはFULL_CONTROL権限を取得します。
         |                    | 他のユーザーはアクセス権限を持ちません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROL権限を取得します。
         |                    | AllUsersグループは読み取り権限を持ちます。
         | public-read-write  | オーナーはFULL_CONTROL権限を取得します。
         |                    | AllUsersグループは読み取りおよび書き込み権限を持ちます。
         |                    | バケットでこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーはFULL_CONTROL権限を取得します。
         |                    | AuthenticatedUsersグループは読み取り権限を持ちます。

   --upload-cutoff
      チャンクアップロードに切り替えるための閾値。

      これを超えるサイズのファイルはchunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのアップロードや「rclone mount」やGoogle
      PhotosやGoogleドキュメントでアップロードされたファイルなど）は、このチャンクサイズを使用してマルチパートアップロードされます。

      注意点として、"--s3-upload-concurrency"個のこのサイズのチャンクが、トランスファごとにメモリにバッファリングされます。

      高速なリンクで大きなファイルを転送しており、十分なメモリを持っている場合、これを増やすと転送が高速化されます。

      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、制限の10,000チャンク未満になるようにチャンクサイズを自動的に増やします。

      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBであり、最大で10,000チャンクまで存在できるため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。さらに大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、進行状況の統計情報（"-P"フラグ）の正確性が低下します。Rcloneは、AWS SDKによってバッファリングされたチャンクが送信されたときにチャンクを送信したとみなしますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きいほど、AWS SDKのバッファサイズも大きくなり、真実から離れた進行報告が行われます。

   --max-upload-parts
      マルチパートアップロードの最大パーツ数。

      このオプションは、マルチパートアップロード時に使用するパーツの最大数を定義します。

      サービスがAWS S3の10,000チャンクの仕様に対応していない場合、これが役立ちます。

      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、このチャンクサイズを自動的に増やしてこのチャンクの数の制限を下回るようにします。

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値。

      サーバーサイドでコピーする必要があるこれを超えるサイズのファイルは、このサイズのチャンクでコピーされます。
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しない。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始までに長時間の遅延が発生します。

   --shared-credentials-file
      共有認証情報ファイルへのパス。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用できます。

      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を検索します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      共有認証情報ファイルで使用するプロファイル。

      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。

      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルトになります。

   --session-token
      AWSのセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。

      同じファイルのチャンクがいくつか同時にアップロードされます。

      高速リンクで大量の大きなファイルをアップロードしており、これらのアップロードがバンド幅を十分に利用しない場合は、これを増やして転送速度を向上させることができます。

   --force-path-style
      もしtrueならばパス形式のアクセスを使用し、falseならばバーチャルホスト形式を使用します。

      これがtrue（デフォルト）の場合、rcloneはパス形式のアクセスを使用します。falseの場合はバーチャルパス形式を使用します。詳細については、[AWS
      S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。

      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）では、これをfalseに設定する必要があります。rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      もしtrueならばv2認証を使用します。

      false（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。

      v4シグネチャが機能しない場合にのみ使用してください（例：Jewel/v10以前のCEPHなど）。

   --list-chunk
      リストイングのチャンクサイズ（各ListObject S3リクエストのレスポンスリストのサイズ）。

      このオプションはAWS S3のMaxKeys、max-items、またはpage-sizeとしても知られています。
      大部分のサービスでは、リクエストよりも多くの数を指定してもレスポンスリストを1000個で切り捨てます。
      AWS S3では、これは最大値のため、変更できません。詳細については[AWS
      S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動（0）。

      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されていました。

      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高パフォーマンスであり、可能な限り使用する必要があります。

      デフォルトの0に設定されている場合、rcloneはプロバイダに応じてどのリストオブジェクトメソッドを呼び出すかを推測します。推測が間違っている場合、ここで手動で設定することができます。

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset

      一部のプロバイダでは、リストをURLエンコードする機能があり、ファイル名に制御文字を含める場合にこれがより信頼性があります。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に基づいてどの設定を適用するかを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-check-bucket
      バケットの存在を確認せず、または作成しないようにします。

      バケットが既に存在する場合、rcloneが行うトランザクションの数を最小限にするために、これを使用すると便利です。

      バケット作成の権限を持たないユーザーを使用している場合にも必要です。v1.52.0より前のバージョンでは、これはバグのためにサイレントにパスされることがありました。

   --no-head
      アップロードされたオブジェクトのHEADを行わずに整合性をチェックしません。

      rcloneは、PUTでオブジェクトをアップロードした後、200 OKメッセージを受け取ると、正しくアップロードされたと想定します。

      特に、次の項目を想定しています：

      - アップロード時のメタデータ、モディファイドタイム、ストレージクラス、コンテンツタイプはアップロード時のものである。
      - サイズはアップロード時のものである。

      シングルパートのPUTの場合、次の項目をレスポンスから読み取ります：

      - MD5SUM
      - アップロード日時

      マルチパートのアップロードでは、これらの項目は読み取られません。

      不明な長さのソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを行います。

      このフラグを設定すると、アップロードの失敗が検出されない可能性が高まります。特に、サイズが正しくない場合のアップロードの失敗の確率は非常に低いですが、このフラグを通常の操作には推奨されません。実際には、このフラグを使用しても、アップロードの失敗が検出されない可能性は非常に低いです。

   --no-head-object
      GETする前にHEADリクエストを行わないように設定します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。

      追加バッファを必要とするアップロード（たとえばマルチパート）は、割り当てるためにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。

      現在、s3（具体的にはminio）バックエンドとHTTP/2に関する問題が解決されていません。HTTP/2はS3バックエンドのデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3がCloudFrontネットワークを介してダウンロードされたデータの節約のために、CloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードの検証にETagを使用するかどうか

      これはtrue、false、またはプロバイダのデフォルトを使用するために設定しておく必要があります。

   --use-presigned-request
      マルチパートアップロードではなく、単一のパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか。

      falseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。

      rcloneのバージョン1.59未満では、単一のパートオブジェクトをアップロードするために事前に署名されたリクエストを使用し、このフラグをtrueに設定すると、その機能が再度有効になります。これは、例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含める。

   --version-at
      指定した時刻のファイルのバージョンを表示します。

      パラメータは日付「2006-01-02」、日時「2006-01-02
      15:04:05」、そのような昔からの長さを指定する必要があります（例：「100d」または「1h」）。

      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。

      有効な形式については、[time option docs](/docs/#time-option)を参照してください。

   --decompress
      設定するとgzipでエンコードされたオブジェクトを解凍します。

      "Content-Encoding: gzip"が設定された状態でオブジェクトをS3にアップロードすることは可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。

      このフラグが設定されている場合、rcloneは"Content-Encoding: gzip"でこれらのファイルを受信すると解凍します。これにより、rcloneがサイズとハッシュをチェックできなくなりますが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。

      通常のプロバイダは、オブジェクトがダウンロードされるときにオブジェクトを変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトは、ダウンロード時にも設定されません。

      ただし、一部のプロバイダは`Content-Encoding: gzip`でアップロードされていなくてもオブジェクトをgzip圧縮する場合があります（例：Cloudflare）。

      これに設定してrcloneが`Content-Encoding: gzip`とチャンク転送エンコーディングでオブジェクトをダウンロードした場合、rcloneはオブジェクトをフライで解凍します。

      unsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に基づいて何を適用するかを選択しますが、ここでrcloneの選択を上書きすることができます。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する


OPTIONS:
   --access-key-id value        AWSのアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                  バケットの作成とオブジェクトの保存またはコピー時に使用されるプリセットACL。 [$ACL]
   --endpoint value             S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します（環境変数またはEC2/ECSメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示する
   --location-constraint value  リージョンに一致する場所の制約。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用されるプリセットACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定するとgzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               もしtrueならばパス形式のアクセスを使用し、falseならばバーチャルホスト形式を使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストイングのチャンクサイズ（各ListObject S3リクエストのレスポンスリストのサイズ）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン: 1、2、または自動（0）。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パーツ数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、または作成しないようにします。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADを行わずに整合性をチェックしません。 (default: false) [$NO_HEAD]
   --no-head-object                 GETする前にHEADリクエストを行わないように設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制する (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSのセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードの検証にETagを使用するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          マルチパートアップロードではなく、単一のパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        もしtrueならばv2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時刻のファイルのバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含める。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}