# Alibaba Cloud オブジェクトストレージシステム (OSS)（かつて Aliyun と呼ばれていました）

{% code fullWidth="true" %}
```
名前:
   singularity storage update s3 alibaba - Alibaba Cloud オブジェクトストレージシステム (OSS)（かつて Aliyun と呼ばれていました）

使用法:
   singularity storage update s3 alibaba [command options] <name|id>

説明:
   --env-auth
      実行時に AWS 認証情報（環境変数または EC2/ECS メタデータ）を取得します。
      
      access_key_id と secret_access_key が空の場合にのみ適用されます。

      例:
         | false | 次のステップで AWS 認証情報を入力してください。
         | true  | 環境変数（env vars または IAM ）から AWS 認証情報を取得します。

   --access-key-id
      AWS アクセスキー ID。
      
      匿名アクセスまたは実行時の資格情報の場合は空のままにしてください。

   --secret-access-key
      AWS シークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の資格情報の場合は空のままにしてください。

   --endpoint
      OSS API のエンドポイント。

      例:
         | oss-accelerate.aliyuncs.com          | グローバル アクセラレート
         | oss-accelerate-overseas.aliyuncs.com | グローバル アクセラレート（中国本土外）
         | oss-cn-hangzhou.aliyuncs.com         | 中国東部 1（杭州）
         | oss-cn-shanghai.aliyuncs.com         | 中国東部 2（上海）
         | oss-cn-qingdao.aliyuncs.com          | 中国北部 1（青島）
         | oss-cn-beijing.aliyuncs.com          | 中国北部 2（北京）
         | oss-cn-zhangjiakou.aliyuncs.com      | 中国北部 3（張家口）
         | oss-cn-huhehaote.aliyuncs.com        | 中国北部 5（フフホト）
         | oss-cn-wulanchabu.aliyuncs.com       | 中国北部 6（ウランチャブ）
         | oss-cn-shenzhen.aliyuncs.com         | 中国南部 1（深セン）
         | oss-cn-heyuan.aliyuncs.com           | 中国南部 2（河源）
         | oss-cn-guangzhou.aliyuncs.com        | 中国南部 3（広州）
         | oss-cn-chengdu.aliyuncs.com          | 中国西部 1（成都）
         | oss-cn-hongkong.aliyuncs.com         | 香港（香港）
         | oss-us-west-1.aliyuncs.com           | 米国西部 1（シリコンバレー）
         | oss-us-east-1.aliyuncs.com           | 米国東部 1（バージニア）
         | oss-ap-southeast-1.aliyuncs.com      | 東南アジア 東南 1（シンガポール）
         | oss-ap-southeast-2.aliyuncs.com      | アジア太平洋 南東 2（シドニー）
         | oss-ap-southeast-3.aliyuncs.com      | 東南アジア 東南 3（クアラルンプール）
         | oss-ap-southeast-5.aliyuncs.com      | アジア太平洋 南東 5（ジャカルタ）
         | oss-ap-northeast-1.aliyuncs.com      | アジア太平洋 北東 1（日本）
         | oss-ap-south-1.aliyuncs.com          | アジア太平洋 南 1（ムンバイ）
         | oss-eu-central-1.aliyuncs.com        | 中央ヨーロッパ 1（フランクフルト）
         | oss-eu-west-1.aliyuncs.com           | 西ヨーロッパ（ロンドン）
         | oss-me-east-1.aliyuncs.com           | 中東 1（ドバイ）

   --acl
      バケットの作成やオブジェクトの保存またはコピー時に使用される予約済み ACL。
      
      この ACL はオブジェクトの作成時に使用され、bucket_acl が設定されていない場合はバケットの作成にも使用されます。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      サーバーサイドのオブジェクトのコピー時にこの ACL が適用されることに注意してください。
      S3 はソースから ACL をコピーせず、代わりに新しい ACL を書き込むためです。
      
      もし acl が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト値（非公開）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用される予約済み ACL。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      「acl」と「bucket_acl」が設定されていない場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト値（非公開）が使用されます。

      例:
         | private            | オーナーに FULL_CONTROL の権限があります。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーに FULL_CONTROL の権限があります。
         |                    | AllUsers グループは READ アクセスがあります。
         | public-read-write  | オーナーに FULL_CONTROL の権限があります。
         |                    | AllUsers グループは READ および WRITE アクセスがあります。
         |                    | バケットでこれを許可することは一般的には推奨されません。
         | authenticated-read | オーナーに FULL_CONTROL の権限があります。
         |                    | AuthenticatedUsers グループは READ アクセスがあります。

   --storage-class
      OSS に新しいオブジェクトを保存する際に使用するストレージクラス。

      例:
         | <未設定>     | デフォルト値
         | STANDARD    | 標準ストレージクラス
         | GLACIER     | 過去アクセスストレージモード
         | STANDARD_IA | 低頻度アクセスストレージモード

   --upload-cutoff
      チャンク化アップロードに切り替えるためのファイルのカットオフ値。
      
      この値より大きなファイルは、chunk_size のチャンクでアップロードされます。
      0 以上 5 GiB 以下の範囲を設定できます。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoff より大きなファイルや、サイズが不明なファイル（「rclone rcat」からのもの、
      「rclone mount」でアップロードされたもの、Google フォトまたは Google ドキュメントなど）は、
      このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意: "--s3-upload-concurrency" を使用して、このチャンクサイズのチャンクが転送ごとにメモリ内でバッファリングされます。
      
      高速リンクを介して大きなファイルを転送し、メモリを十分に備えている場合は、これを増やすと転送が高速化されます。
      
      Rclone は、10,000 チャンクの制限を超えることなく、既知のサイズの大きなファイルをアップロードする場合、自動的に
      チャンクサイズを増やします。
      
      サイズが不明なファイルは、設定された chunk_size でアップロードされます。
      デフォルトのチャンクサイズが 5 MiB で最大 10,000 チャンクまであることを考慮すると、
      ストリームアップロードできるファイルの最大サイズは 48 GiB です。ファイルの大きさを超えるストリームアップロードを
      行う場合は、チャンクサイズを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況の統計情報の精度が低下します。Rclone は、AWS SDK によってチャンクがバッファリング
      されるとチャンクが送信されたとみなすため、まだアップロード中の場合でもです。より大きなチャンクサイズは、
      より大きな AWS SDK バッファと進行状況レポートの真実からの逸脱を意味します。
      

   --max-upload-parts
      マルチパートアップロードの最大パート数。
      
      このオプションでは、マルチパートアップロード時に使用されるパートの最大数を設定します。
      
      サービスが AWS S3 の 10,000 チャンク仕様をサポートしていない場合に役立ちます。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする場合、
      このチャンクの数の制限を超えることなく自動的にチャンクサイズを増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのファイルのカットオフ値。
      
      サーバーサイドでコピーする必要があるこの値より大きなファイルは、このサイズのチャンクでコピーされます。
      
      0 以上 5 GiB 以下の範囲を設定できます。

   --disable-checksum
      オブジェクトメタデータに MD5 チェックサムを保存しないでください。
      
      通常、アップロード前に rclone は入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加するため、
      大きなファイルをアップロードするときにはアップロードの開始までに長い遅延が発生します。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を参照します。
      環境変数が空の場合は、現在のユーザーのホームディレクトリがデフォルト値となります。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数は、そのファイルで使用されるプロファイルを制御します。
      
      空にした場合、環境変数 "AWS_PROFILE" または設定されていない場合は "default" がデフォルト値となります。
      

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。
      
      同じファイルのチャンクを並行してアップロードする数です。
      
      高速リンクで大量の大きなファイルを転送し、これらのアップロードが
      帯域幅を完全に活用していない場合は、これを増やすと転送が高速化される可能性があります。

   --force-path-style
      true の場合、パススタイルアクセスを使用します。false の場合、仮想ホスト形式を使用します。
      
      これが true の場合（デフォルト）、rclone はパススタイルアクセスを使用します。
      false の場合、rclone は仮想ホスト形式を使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、または Tencent COS）では、
      false に設定する必要があります（rclone はこれをプロバイダの設定に基づいて自動的に行います）。
      

   --v2-auth
      true の場合、v2 認証を使用します。
      
      false に設定すると（デフォルト）、rclone は v4 認証を使用します。
      設定されると、rclone は v2 認証を使用します。
      
      v4 シグネチャが機能しない場合にのみ使用してください（Jewel/v10 CEPH など）。

   --list-chunk
      リスティングのチャンクサイズ（各 ListObject S3 リクエストごとのレスポンスリストのサイズ）。
      
      このオプションは AWS S3 仕様の MaxKeys、max-items、または page-size としても知られています。
      ほとんどのサービスは、1000 オブジェクトを超えるリクエストでもリストを切り詰めます。
      AWS S3 では、これはグローバルな最大値ですが、詳細については、[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Ceph の場合、オプション "rgw list buckets max chunk" でこれを増やすことができます。
      

   --list-version
      使用する ListObjects のバージョン: 1,2 または 0（自動）。
      
      S3 の最初のリリース時、バケット内のオブジェクトを列挙するための ListObjects 呼び出しが提供されました。
      
      しかし、2016 年 5 月に ListObjectsV2 呼び出しが導入されました。これは非常に高性能であり、可能な限り使用する必要があります。
      
      デフォルト値の 0 の場合、rclone はプロバイダがリストオブジェクトメソッドを呼び出すと推測し、メソッドを自動的に設定します。
      正しく推測されない場合は、ここで手動で設定できます。
      

   --list-url-encode
      リストの URL エンコーディングに使用するかどうか: true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を使用する場合に URL エンコーディングの一貫性があります。
      プロバイダの設定に従って適用する内容を rclone が選択しますが、ここで rclone の選択を上書きできます。
      

   --no-check-bucket
      バケットの存在を確認したり作成しない設定をします。
      
      バケットが既に存在する場合、rclone が実行するトランザクションの数を最小限にするために役立ちます。
      
      バケット作成の権限を持たない場合にも必要になる場合があります。
      v1.52.0 より前のバージョンでは、このバグのために黙ってパスしました。
      

   --no-head
      アップロードしたオブジェクトの HEAD をチェックしない設定をします。
      
      rclone が PUT でオブジェクトをアップロードした後に 200 OK メッセージを受信した場合、正常にアップロードされたと想定します。
      
      特に次のことを想定します:
      
      - メタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロードしたものと同じであること
      - サイズがアップロードしたものと同じであること
      
      シングルパートの PUT の場合、次の項目をレスポンスから読み取ります:
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取りません。
      
      サイズが不明なソースオブジェクトをアップロードする場合、rclone は HEAD リクエストを実行します。
      
      このフラグを設定すると、アップロードの失敗を検出できない可能性が高くなるため、通常の操作ではお勧めしません。
      実際には、このフラグを設定しても、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      オブジェクトを取得する前に HEAD を実行しない設定をします。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされるタイミング。
      
      追加のバッファが必要なアップロード（たとえば、マルチパート）では、アロケーションにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドで http2 の使用を無効にします。
      
      現在、s3 (特に minio) バックエンドと HTTP/2 の未解決の問題があります。
      HTTP/2 は、デフォルトで s3 バックエンドで有効になっていますが、ここでは無効にすることができます。
      問題が解決されると、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロードに使用するカスタムエンドポイント。
      これは通常、AWS S3 が CloudFront ネットワークを介してダウンロードされたデータの配信でより低コストな転送が提供されるように、
      CloudFront CDN の URL に設定されます。

   --use-multipart-etag
      マルチパートアップロード時に ETag を検証に使用するかどうか
      
      true、false、またはそれをプロバイダのデフォルト値に設定フラグに設定して使用します。
      

   --use-presigned-request
      シングルパートアップロードに署名付きリクエストまたは PutObject を使用するかどうか
      
      false の場合、rclone は AWS SDK の PutObject を使用してオブジェクトをアップロードします。
      
      rclone のバージョン 1.59 未満では、単一のパートオブジェクトをアップロードするために署名付きリクエストを使用し、
      このフラグを true に設定するとその機能が再有効になります。これは例外的な状況やテスト以外では必要ありません。
      

   --versions
      ディレクトリリスティングに古いバージョンも含めるかどうか。

   --version-at
      指定した時点でのファイルバージョンを表示します。
      
      パラメータは日付、"2006-01-02"、日時 "2006-01-02 15:04:05"、またはそのような過去の期間、たとえば "100d" や "1h" です。
      
      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルのアップロードや削除はできません。
      
      有効な書式については、[時間オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      この設定を行うと、gzip でエンコードされたオブジェクトが解凍されます。
      
      "Content-Encoding: gzip" が設定されている場合、rclone はこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されていると、rclone は "Content-Encoding: gzip" のファイルを受信時に解凍します。
      つまり、rclone はサイズとハッシュをチェックすることはできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定してください。
      
      通常、プロバイダはダウンロード時にオブジェクトを変更しません。`Content-Encoding: gzip` でアップロードされていない場合、
      ダウンロード時には設定されません。
      
      ただし、一部のプロバイダ（Cloudflare など）は、`Content-Encoding: gzip` でアップロードされていなくてもオブジェクトを gzip 圧縮する場合があります。
      
      これによって次のようなエラーが発生することがあります:
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rclone が `Content-Encoding: gzip` が設定され、チャンク化転送エンコーディングを使用してオブジェクトをダウンロードすると、
      rclone はそのオブジェクトをリアルタイムで解凍します。
      
      これが設定されると、解凍しない場合（デフォルト）に比べてアップロードの失敗が検出される可能性が高くなります。

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します 


オプション:
   --access-key-id value      AWS アクセスキー ID。 [$ACCESS_KEY_ID]
   --acl value                バケットの作成やオブジェクトの保存またはコピー時に使用される予約済み ACL。 [$ACL]
   --endpoint value           OSS API のエンドポイント。 [$ENDPOINT]
   --env-auth                 実行時に AWS 認証情報（環境変数または EC2/ECS メタデータ）を取得します。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示
   --secret-access-key value  AWS シークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --storage-class value      OSS に新しいオブジェクトを保存する際に使用するストレージクラス。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用される予約済み ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのファイルのカットオフ値。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     この設定を行うと、gzip でエンコードされたオブジェクトが解凍されます。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータに MD5 チェックサムを保存しないでください。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドで http2 の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロードに使用するカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合、パススタイルアクセスを使用します。false の場合、仮想ホスト形式を使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リスティングのチャンクサイズ。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストの URL エンコーディングに使用するかどうか。 (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects のバージョン: 1,2 または 0（自動） (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パート数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされるタイミング。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールで mmap バッファを使用するかどうか。 (デフォルト: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトを gzip 圧縮する可能性がある場合に設定します。 (デフォルト: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在を確認したり作成しない設定をします。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --no-head                        アップロードしたオブジェクトの HEAD をチェックしない設定をします。 (デフォルト: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する前に HEAD を実行しない設定をします。 (デフォルト: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (デフォルト: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWS セッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行性。 (デフォルト: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるためのファイルのカットオフ値。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロード時に ETag を検証に使用するかどうか (デフォルト: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名付きリクエストまたは PutObject を使用するかどうか (デフォルト: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true の場合、v2 認証を使用します。 (デフォルト: false) [$V2_AUTH]
   --version-at value               指定した時点でのファイルバージョンを表示します。 (デフォルト: "off") [$VERSION_AT]
   --versions                       ディレクトリリスティングに古いバージョンも含めるかどうか。 (デフォルト: false) [$VERSIONS]

```
{% endcode %}