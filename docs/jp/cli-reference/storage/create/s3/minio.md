# Minio オブジェクトストレージ

{% code fullWidth="true" %}
```
名前:
   singularity storage create s3 minio - Minio オブジェクトストレージ

使用法:
   singularity storage create s3 minio [コマンド オプション] [引数...]

説明:
   --env-auth
      実行時に AWS 認証情報を取得します (環境変数または EC2/ECS メタデータから)。
      
      access_key_id と secret_access_key が空の場合にのみ適用されます。

      例:
         | false | AWS 認証情報を次のステップで入力します。
         | true  | 環境から AWS 認証情報を取得します (環境変数または IAM)。

   --access-key-id
      AWS アクセスキー ID。
      
      匿名アクセスまたは実行時の認証情報の場合、空白のままにしてください。

   --secret-access-key
      AWS シークレットアクセスキー (パスワード)。
      
      匿名アクセスまたは実行時の認証情報の場合、空白のままにしてください。

   --region
      接続するリージョン。
      
      S3 互換のストレージを使用し、リージョンがない場合は空白のままにしてください。

      例:
         | <未設定>                | 確定できない場合はこれを使用します。
         |                        | v4 署名と空のリージョンを使用します。
         | other-v2-signature      | v4 署名が機能しない場合のみ使用します。
         |                        | たとえば、Jewel/v10 CEPH 以前。

   --endpoint
      S3 API のエンドポイント。
      
      S3 互換のストレージを使用する場合、必須です。

   --location-constraint
      リージョンに一致させる必要のあるロケーション制約。
      
      確かめることができない場合は空白のままにしてください。バケット作成時にのみ使用されます。

   --acl
      バケットの作成、オブジェクトの保存またはコピー時に使用される一括 ACL。
      
      この ACL はオブジェクトの作成に使用され、bucket_acl が設定されていない場合にはバケットの作成にも使用されます。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      S3 によるオブジェクトのサーバーサイドコピー時にこの ACL が適用されることに注意してください。
      ソースから ACL をコピーせずに新しい ACL を書き込みます。
      
      もし acl が空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト (private) が使用されます。
      

   --bucket-acl
      バケットの作成時に使用される一括 ACL。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      この ACL はバケットの作成時のみ適用されます。設定されていない場合は "acl" が代わりに使用されます。
      
      もし "acl" と "bucket_acl" が空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト (private) が使用されます。
      

      例:
         | private                | オーナーには FULL_CONTROL 権限があります。
         |                        | 他のユーザーにはアクセス権がありません (デフォルト)。
         | public-read            | オーナーには FULL_CONTROL 権限があります。
         |                        | AllUsers グループには READ 権限があります。
         | public-read-write      | オーナーには FULL_CONTROL 権限があります。
         |                        | AllUsers グループには READ および WRITE 権限があります。
         |                        | バケットにこれを許可することは一般的に推奨されません。
         | authenticated-read     | オーナーには FULL_CONTROL 権限があります。
         |                        | AuthenticatedUsers グループには READ 権限があります。

   --server-side-encryption
      このオブジェクトを S3 に保存する際に使用するサーバーサイド暗号化アルゴリズム。

      例:
         | <未設定> | None
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C を使用する場合、このオブジェクトを S3 に保存する際に使用されるサーバーサイド暗号化アルゴリズム。

      例:
         | <未設定> | None
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID を使用する場合、キーの ARN を提供する必要があります。

      例:
         | <未設定>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C を使用する場合、データの暗号化/復号に使用する秘密の暗号キーを提供できます。
      
      代わりに --sse-customer-key-base64 を指定することもできます。

      例:
         | <未設定> | None

   --sse-customer-key-base64
      SSE-C を使用する場合、データの暗号化/復号に使用する秘密の暗号キーを base64 形式でエンコードして提供する必要があります。
      
      代わりに --sse-customer-key を指定することもできます。

      例:
         | <未設定> | None

   --sse-customer-key-md5
      SSE-C を使用する場合、秘密の暗号キーの MD5 チェックサムを提供できます (オプション)。
      
      空のままにすると、sse_customer_key から自動的に計算されます。

      例:
         | <未設定> | None

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ範囲。
      
      これより大きいファイルは chunk_size のチャンクに分割してアップロードされます。
      最小値は 0、最大値は 5 GiB です。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoff よりも大きいファイル、またはサイズが不明なファイル ("rclone rcat" からのアップロードや "rclone mount" や Google フォトまたは Google ドキュメントでアップロードされたファイルなど) は、このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      注意:
      --s3-upload-concurrency は転送ごとにこのサイズのチャンクがメモリ内にバッファリングされます。
      
      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合は、これを高めると転送が高速化します。
      
      Rclone は、10,000 のチャンク数制限を超えないように、既知のサイズの大きなファイルをアップロードする際には自動的にチャンクサイズを増やします。
      
      不明のサイズのファイルは設定された chunk_size でアップロードされます。デフォルトの chunk_size は 5 MiB で、最大で 10,000 のチャンクがあるため、デフォルト設定ではストリームアップロードできるファイルの最大サイズは 48 GiB です。より大きなファイルをストリームアップロードする場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを大きくすると、進捗状況の統計情報の精度が低下します。Rclone は、AWS SDK によってチャンクがバッファリングされたときにチャンクを送信したとみなし、実際にはまだアップロードされている場合でもそうです。大きなチャンクサイズは、AWS SDK のバッファサイズと、真実からの進捗報告のズレを大きくします。
      

   --max-upload-parts
      マルチパートアップロードで使用する最大のパート数。
      
      このオプションは、マルチパートアップロード時に使用する最大のマルチパートチャンク数を定義します。
      
      AWS S3 仕様の 10,000 チャンクをサポートしていないサービスに便利です。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする際には自動的にチャンクサイズを増やして、このチャンク数の制限以下にするようにします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるカットオフ範囲。
      
      コピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小値は 0、最大値は 5 GiB です。

   --disable-checksum
      オブジェクトのメタデータに MD5 チェックサムを格納しない。
      
      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまで長い遅延が発生することがあります。

   --shared-credentials-file
      共有認証情報ファイルのパス。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は環境変数 "AWS_PROFILE" または "default" の値がデフォルトになります。
      

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同じファイルのチャンクが同時にアップロードされます。
      
      高速リンクを介して大きな数の大きなファイルをアップロードし、これらのアップロードが帯域幅を完全に活用していない場合、これを増やすと転送が高速化する場合があります。

   --force-path-style
      true の場合、パス形式のアクセスを使用します。 false の場合、仮想ホスト形式のアクセスを使用します。
      
      この値が true (デフォルト) の場合、rclone はパス形式のアクセスを使用します。
      false の場合、rclone は仮想パス形式のアクセスを使用します。詳細は[AWS S3 のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ (AWS、Aliyun OSS、Netease COS、または Tencent COS など) は、この値を false に設定する必要があります。rclone はプロバイダの設定に基づいて自動的にこれを行います。

   --v2-auth
      v2 認証を使用する場合は true を設定します。
      
      false (デフォルト) の場合、rclone は v4 認証を使用します。設定されている場合は rclone は v2 認証を使用します。
      
      v4 署名が機能しない場合にのみ使用してください。たとえば、Jewel/v10 CEPH 以前の場合。

   --list-chunk
      リストのチャンクサイズ (各 ListObject S3 リクエストの応答リスト)。
      
      このオプションは、AWS S3 仕様の "MaxKeys"、"max-items"、または "page-size" としても知られています。
      大抵のサービスは、要求されたよりも多くを要求してもレスポンスリストを最大 1000 オブジェクトに切り捨てます。
      AWS S3 では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Ceph では、「rgw list buckets max chunk」オプションを使用してこれを増やすことができます。
      

   --list-version
      使用する ListObjects のバージョン: 1、2、または 0 (自動)。
      
      S3 が最初にリリースされた当初、バケット内のオブジェクトを列挙するために ListObjects 呼び出しのみが提供されていました。
      
      しかし、2016 年 5 月に ListObjectsV2 呼び出しが導入されました。これは非常に高速であり、可能であれば使用する必要があります。
      
      デフォルト値の 0 の場合、rclone はプロバイダの設定に基づいてどちらの list objects メソッドを呼び出すか推測します。推測が間違っている場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストの URL エンコードの有無: true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を使用する場合に URL エンコードリストをサポートしています。可能な場合は制御文字を使用する場合は、これが重要です。これが unset に設定されている場合 (デフォルト) 、rclone はプロバイダによる設定に基づいて適用するものを選択しますが、ここで rclone の選択を上書きすることができます。
      

   --no-check-bucket
      バケットの存在を確認せず、作成も行いません。

      バケットが既に存在することを事前に知っている場合、rclone が実行するトランザクション数を最小限にするために役立ちます。
      
      バケット作成権限を持っていない場合にも必要になる場合があります。バージョン 1.52.0 より前では、これは無音で渡されるはずでしたが、バグのためにそれは黙って過ぎ去りました。
      

   --no-head
      アップロードされたオブジェクトの HEAD をチェックしない場合は設定します。

      rclone が PUT 後に 200 OK メッセージを受け取った場合、正しくアップロードされたと仮定します。
      
      特に次のことを仮定します:
      
      - メタデータ (変更日時、ストレージクラス、コンテンツタイプを含む) がアップロードしたものと同じであること
      - サイズがアップロードしたものと同じであること
      
      シングルパートの PUT の場合、rclone は次の項目をレスポンスから読み取ります:
      
      - MD5SUM
      - アップロード日
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      未知の長さを持つ元のオブジェクトがアップロードされた場合、rclone は HEAD リクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗を検出できるチャンスが増えます。特に、正しくないサイズの場合です。そのため、通常の操作では推奨されません。実際にアップロードの失敗が発生する確率は非常に低く、このフラグを使用しても変わりません。
      

   --no-head-object
      GET の前に HEAD を行わない場合は設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      メモリ内バッファプールをフラッシュする頻度。
      
      追加のバッファが必要なアップロード (たとえばマルチパート) では、割り当てにメモリプールが使用されます。
      このオプションは、未使用のバッファをメモリプールから削除する頻度を制御します。

   --memory-pool-use-mmap
      バックエンドの内部メモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。
      
      現在 s3 (特に minio) バックエンドと HTTP/2 の問題が未解決です。HTTP/2 は s3 バックエンドではデフォルトで有効ですが、ここで無効にできます。問題が解決されると、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3 では CloudFront CDN URL に設定され、CloudFront ネットワークを介してダウンロードされたデータに対して低コストの転送が提供されます。

   --use-multipart-etag
      マルチパートアップロードで ETag を使用して検証するかどうか
      
      true、false、またはデフォルトのプロバイダの設定を使用するには true、false、または未設定のいずれかを設定します。
      

   --use-presigned-request
      シングルパートアップロードの場合に署名済みリクエストまたは PutObject を使用するかどうか
      
      false の場合、rclone は AWS SDK の PutObject を使用してオブジェクトをアップロードします。
      
      rclone のバージョンが 1.59 未満では、シングルパートオブジェクトをアップロードするために署名付きリクエストを使用し、このフラグを true に設定すると、その機能が再度有効になります。これは、特殊な状況やテスト以外では不要です。
      

   --versions
      ディレクトリリストに旧バージョンを含めます。

   --version-at
      指定した時刻でのファイルバージョンを表示します。
      
      パラメータは日付 "2006-01-02"、datetime "2006-01-02
      15:04:05"、またはそれからの経過時間、例えば "100d" や "1h" のいずれかにする必要があります。
      
      このオプションを使用すると、ファイルの書き込み操作は許可されないため、ファイルのアップロードや削除はできません。
      
      有効な書式については、[time オプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      Gzip エンコードされたオブジェクトを解凍する場合は設定します。
      
      S3 に "Content-Encoding: gzip"が設定されたファイルをアップロードすることができます。通常、rclone はこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rclone はこれらのファイルを "Content-Encoding: gzip" として受け取るときに解凍します。これにより、rclone はサイズとハッシュを確認できませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトを圧縮する可能性がある場合に設定します。
      
      一般的に、プロバイダはオブジェクトをダウンロードする際には変更しません。"Content-Encoding: gzip" でアップロードされなかった場合、ダウンロード時にも設定されません。
      
      ただし、一部のプロバイダ (Cloudflare など) は、"Content-Encoding: gzip" でアップロードされていないオブジェクトであっても、オブジェクトを gzip 圧縮する場合があります。
      
      次のようなエラーが発生する場合、これが原因です。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rclone が Content-Encoding: gzip が設定され、チャンク転送エンコーディングを使用するようなオブジェクトをダウンロードすると、rclone はオブジェクトをリアルタイムで解凍します。
      
      unset (デフォルト) の場合は、rclone はプロバイダの設定に基づいて適用するものを選択しますが、ここで rclone の選択を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します


オプション:
   --access-key-id value           AWS アクセスキー ID。[$ACCESS_KEY_ID]
   --acl value                     バケットの作成、オブジェクトの保存またはコピー時に使用される一括 ACL。[$ACL]
   --endpoint value                S3 API のエンドポイント。[$ENDPOINT]
   --env-auth                      実行時に AWS 認証情報を取得します (環境変数または EC2/ECS メタデータから)。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示
   --location-constraint value     リージョンに一致させる必要のあるロケーション制約。[$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。[$REGION]
   --secret-access-key value       AWS シークレットアクセスキー (パスワード)。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  このオブジェクトを S3 に保存する際に使用するサーバーサイド暗号化アルゴリズム。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID を使用する場合、キーの ARN を提供する必要があります。[$SSE_KMS_KEY_ID]

   先進的なオプション

   --bucket-acl value               バケットの作成時に使用される一括 ACL。[$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるカットオフ範囲。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     Gzip エンコードされたオブジェクトを解凍する場合は設定します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータに MD5 チェックサムを格納しない。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドでの http2 の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合、パス形式のアクセスを使用します。 false の場合、仮想ホスト形式のアクセスを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストの URL エンコードの有無: true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1,2, または 0 (自動)。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用する最大のパート数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファープールのフラッシュ頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           バックエンドの内部メモリプールで mmap バッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを圧縮する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認せず、作成も行いません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトの HEAD をチェックしない場合は設定します。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 GET の前に HEAD を行わない場合は設定します。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWS セッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルのパス。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C を使用する場合、このオブジェクトを S3 に保存する際に使用されるサーバーサイド暗号化アルゴリズム。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C を使用する場合、データの暗号化/復号に使用する秘密の暗号キーを提供できます。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C を使用する場合、データの暗号化/復号に使用する秘密の暗号キーを base64 形式でエンコードして提供する必要があります。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C を使用する場合、秘密の暗号キーの MD5 チェックサムを提供できます (オプション)。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       マルチパートアップロードの同時実行数。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ範囲。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードで ETag を使用して検証するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードの場合に署名済みリクエストまたは PutObject を使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 認証を使用する場合は true を設定します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時刻でのファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに旧バージョンを含めます。 (デフォルト: false) [$VERSIONS]

   全般

   --name value  ストレージの名前 (デフォルト: 生成されたもの)
   --path value  ストレージのパス

```
{% endcode %}