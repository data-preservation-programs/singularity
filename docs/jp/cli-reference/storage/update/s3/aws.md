# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 aws - Amazon Web サービス (AWS) S3

USAGE:
   singularity storage update s3 aws [command options] <name|id>

DESCRIPTION:
   --env-auth
      実行時に AWS の認証情報を取得します (環境変数または環境に設定されている EC2/ECS のメタデータ)。
      
      access_key_id と secret_access_key が空の場合にのみ適用されます。

      例:
         | false | 次の手順で AWS の認証情報を入力してください。
         | true  | 環境 (環境変数または IAM) から AWS の認証情報を取得します。

   --access-key-id
      AWS Access Key ID。
      
      無記入にすると匿名アクセスまたは実行時の認証情報が使用されます。

   --secret-access-key
      AWS Secret Access Key (パスワード)。
      
      無記入にすると匿名アクセスまたは実行時の認証情報が使用されます。

   --region
      接続するリージョン。

      例:
         | us-east-1      | デフォルトのエンドポイント - 迷ったらこちらを選択してください。
         |                | 米国リージョン、バージニア州北部、または太平洋北西部。
         |                | ロケーション制約は空にしてください。
         | us-east-2      | 米国東部 (オハイオ州) リージョン。
         |                | ロケーション制約は us-east-2 にする必要があります。
         | us-west-1      | 米国西部 (カリフォルニア北部) リージョン。
         |                | ロケーション制約は us-west-1 にする必要があります。
         | us-west-2      | 米国西部 (オレゴン州) リージョン。
         |                | ロケーション制約は us-west-2 にする必要があります。
         | ca-central-1   | カナダ (中部) リージョン。
         |                | ロケーション制約は ca-central-1 にする必要があります。
         | eu-west-1      | EU (アイルランド) リージョン。
         |                | ロケーション制約は EU または eu-west-1 にする必要があります。
         | eu-west-2      | EU (ロンドン) リージョン。
         |                | ロケーション制約は eu-west-2 にする必要があります。
         | eu-west-3      | EU (パリ) リージョン。
         |                | ロケーション制約は eu-west-3 にする必要があります。
         | eu-north-1     | EU (ストックホルム) リージョン。
         |                | ロケーション制約は eu-north-1 にする必要があります。
         | eu-south-1     | EU (ミラノ) リージョン。
         |                | ロケーション制約は eu-south-1 にする必要があります。
         | eu-central-1   | EU (フランクフルト) リージョン。
         |                | ロケーション制約は eu-central-1 にする必要があります。
         | ap-southeast-1 | アジアパシフィック (シンガポール) リージョン。
         |                | ロケーション制約は ap-southeast-1 にする必要があります。
         | ap-southeast-2 | アジアパシフィック (シドニー) リージョン。
         |                | ロケーション制約は ap-southeast-2 にする必要があります。
         | ap-northeast-1 | アジアパシフィック (東京) リージョン。
         |                | ロケーション制約は ap-northeast-1 にする必要があります。
         | ap-northeast-2 | アジアパシフィック (ソウル) リージョン。
         |                | ロケーション制約は ap-northeast-2 にする必要があります。
         | ap-northeast-3 | アジアパシフィック (大阪) ローカル リージョン。
         |                | ロケーション制約は ap-northeast-3 にする必要があります。
         | ap-south-1     | アジアパシフィック (ムンバイ) リージョン。
         |                | ロケーション制約は ap-south-1 にする必要があります。
         | ap-east-1      | アジアパシフィック (香港) リージョン。
         |                | ロケーション制約は ap-east-1 にする必要があります。
         | sa-east-1      | 南アメリカ (サンパウロ) リージョン。
         |                | ロケーション制約は sa-east-1 にする必要があります。
         | me-south-1     | 中東 (バーレーン) リージョン。
         |                | ロケーション制約は me-south-1 にする必要があります。
         | af-south-1     | アフリカ (ケープタウン) リージョン。
         |                | ロケーション制約は af-south-1 にする必要があります。
         | cn-north-1     | 中国 (北京) リージョン。
         |                | ロケーション制約は cn-north-1 にする必要があります。
         | cn-northwest-1 | 中国 (寧夏) リージョン。
         |                | ロケーション制約は cn-northwest-1 にする必要があります。
         | us-gov-east-1  | AWS GovCloud (米国-東) リージョン。
         |                | ロケーション制約は us-gov-east-1 にする必要があります。
         | us-gov-west-1  | AWS GovCloud (米国) リージョン。
         |                | ロケーション制約は us-gov-west-1 にする必要があります。

   --endpoint
      S3 API のエンドポイント。
      
      AWS を使用する場合は空にしてデフォルトのリージョンエンドポイントが使用されます。

   --location-constraint
      ロケーション制約 - リージョンに一致する必要があります。
      
      バケットを作成する場合のみ使用します。

      例:
         | <unset>        | 米国リージョン、バージニア州北部、または太平洋北西部の場合は空
         | us-east-2      | 米国東部 (オハイオ州) リージョンの場合
         | us-west-1      | 米国西部 (カリフォルニア北部) リージョンの場合
         | us-west-2      | 米国西部 (オレゴン州) リージョンの場合
         | ca-central-1   | カナダ (中部) リージョンの場合
         | eu-west-1      | EU (アイルランド) リージョンの場合
         | eu-west-2      | EU (ロンドン) リージョンの場合
         | eu-west-3      | EU (パリ) リージョンの場合
         | eu-north-1     | EU (ストックホルム) リージョンの場合
         | eu-south-1     | EU (ミラノ) リージョンの場合
         | EU             | EU リージョンの場合
         | ap-southeast-1 | アジアパシフィック (シンガポール) リージョンの場合
         | ap-southeast-2 | アジアパシフィック (シドニー) リージョンの場合
         | ap-northeast-1 | アジアパシフィック (東京) リージョンの場合
         | ap-northeast-2 | アジアパシフィック (ソウル) リージョンの場合
         | ap-northeast-3 | アジアパシフィック (大阪) ローカル リージョンの場合
         | ap-south-1     | アジアパシフィック (ムンバイ) リージョンの場合
         | ap-east-1      | アジアパシフィック (香港) リージョンの場合
         | sa-east-1      | 南アメリカ (サンパウロ) リージョンの場合
         | me-south-1     | 中東 (バーレーン) リージョンの場合
         | af-south-1     | アフリカ (ケープタウン) リージョンの場合
         | cn-north-1     | 中国 (北京) リージョンの場合
         | cn-northwest-1 | 中国 (寧夏) リージョンの場合
         | us-gov-east-1  | AWS GovCloud (米国-東) リージョンの場合
         | us-gov-west-1  | AWS GovCloud (米国) リージョンの場合

   --acl
      バケットおよびオブジェクトの作成時に使用される Canned ACL。
      
      オブジェクトの作成時と、bucket_acl が設定されていない場合に使用されます。
      
      詳細についてはこちらを参照してください: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      注意: この ACL は、S3 が元の ACL をコピーするのではなく、新しい ACL を書き込むため、
      サーバーサイドのコピー時に適用されます。
      
      ACL が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルトのプライベートが使用されます。
      

   --bucket-acl
      バケットの作成時に使用される Canned ACL。
      
      詳細についてはこちらを参照してください: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl
      
      バケットを作成する場合のみ適用されます。設定されていない場合は "acl" が代わりに使用されます。
      
      "acl" と "bucket_acl" が空の文字列である場合、X-Amz-Acl: のヘッダーは追加されず、デフォルトのプライベートが使用されます。
      

      例:
         | private            | オーナーに FULL_CONTROL が付与されます。
         |                    | 他の人にアクセス権がありません (デフォルト)。
         | public-read        | オーナーに FULL_CONTROL が付与されます。
         |                    | AllUsers グループに READ アクセスが付与されます。
         | public-read-write  | オーナーに FULL_CONTROL が付与されます。
         |                    | AllUsers グループに READ および WRITE アクセスが付与されます。
         |                    | バケットでこれを設定することは推奨されません。
         | authenticated-read | オーナーに FULL_CONTROL が付与されます。
         |                    | AuthenticatedUsers グループに READ アクセスが付与されます。

   --requester-pays
      S3 バケットとのやり取り時に、リクエスターペイオプションを有効にします。

   --server-side-encryption
      S3 にこのオブジェクトを格納するために使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C を使用する場合、S3 にこのオブジェクトを格納するために使用されるサーバーサイドの暗号化アルゴリズム。

      例:
         | <unset> | なし
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID を使用する場合はキーの ARN を指定する必要があります。

      例:
         | <unset>                 | なし
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C を使用する場合は、データを暗号化/復号化するために使用される秘密の暗号化キーを指定することができます。
      
      代わりに --sse-customer-key-base64 を指定することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-base64
      SSE-C を使用する場合は、データを暗号化/復号化するために使用される秘密の暗号化キーを base64 形式で指定する必要があります。
      
      代わりに --sse-customer-key を指定することもできます。

      例:
         | <unset> | なし

   --sse-customer-key-md5
      SSE-C を使用する場合、秘密の暗号化キーの MD5 チェックサムを指定できます (省略可能)。
      
      空の場合、s3_customer_key から自動的に計算されます。
      

      例:
         | <unset> | なし

   --storage-class
      新しいオブジェクトを S3 に格納する際に使用するストレージクラス。

      例:
         | <unset>             | デフォルト
         | STANDARD            | 標準のストレージクラス
         | REDUCED_REDUNDANCY  | 冗長性の低いストレージクラス
         | STANDARD_IA         | 標準の低頻度アクセス (IA) ストレージクラス
         | ONEZONE_IA          | 1 ゾーンの低頻度アクセス (IA) ストレージクラス
         | GLACIER             | Glacier ストレージクラス
         | DEEP_ARCHIVE        | Deep Archive ストレージクラス
         | INTELLIGENT_TIERING | Intelligent-Tiering ストレージクラス
         | GLACIER_IR          | Glacier Instant Retrieval ストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ値。
      
      これより大きなサイズのファイルは、chunk_size のチャンクでアップロードされます。
      最小値は 0 で、最大値は 5 GiB です。

   --chunk-size
      アップロード時に使用するチャンクサイズ。
      
      upload_cutoff を超えるか、サイズの不明なファイル (たとえば "rclone rcat" からのものや "rclone mount" や Google フォトまたは Google ドキュメントでアップロードされたもの) の場合、このチャンクサイズを使用してチャンクアップロードが行われます。
      
      注意: "--s3-upload-concurrency" ごとにこのチャンクサイズのチャンクがメモリ内にバッファリングされます。
      
      高速リンクで大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送が高速化されます。
      
      Rclone は、サイズのわかっている大きなファイルをアップロードする場合は、10,000 のチャンクの制限を下回るように自動的にチャンクサイズを増やします。
      
      サイズのわからないファイルは設定された chunk_size でアップロードされます。デフォルトの chunk_size が 5 MiB であり、最大で 10,000 のチャンクが存在するため、ストリームアップロードできるファイルの最大サイズは 48 GiB になります。サイズがより大きいファイルをストリームアップロードする場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを増やすと、進行状況統計情報が "-P" フラグで表示される場合に、正確性が低下します。Rclone は、AWS SDK によってバッファリングされたチャンクが送信されたときにチャンクを送信したと見なしますが、まだアップロード中かもしれません。大きなチャンクサイズは、AWS SDK のバッファおよび進行状況報告の精度を低下させます。
      

   --max-upload-parts
      マルチパートアップロード中のパートの最大数。
      
      このオプションは、マルチパートアップロードを行う場合の最大マルチパートの数を定義します。
      
      これは、1 つのサービスが 10,000 のチャンクの AWS S3 仕様をサポートしない場合に使用することができます。
      
      Rclone は、サイズのわかっている大きなファイルをアップロードする場合は、このチャンクの数の制限を下回るように自動的にチャンクサイズを増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドコピーする必要のあるこれより大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小値は 0 で、最大値は 5 GiB です。

   --disable-checksum
      オブジェクトメタデータに MD5 チェックサムを保存しません。
      
      通常、rclone はアップロードする前に入力の MD5 チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまで時間がかかることがあります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rclone は "AWS_SHARED_CREDENTIALS_FILE" 環境変数を参照します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリが使用されます。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = true の場合、rclone は共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合は、環境変数 "AWS_PROFILE" またはデフォルトが設定されていない場合は "default" になります。
      

   --session-token
      AWS セッショントークン。

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同じファイルのチャンクを同時にアップロードします。
      
      ハイスピードリンクで大量の大きなファイルを転送し、これらの転送が帯域幅を十分に利用しない場合、これを増やすと転送速度が向上する可能性があります。

   --force-path-style
      true の場合、パススタイルのアクセスを使用し、false の場合は仮想ホストスタイルを使用します。
      
      これが true (デフォルト) の場合、rclone はパススタイルアクセスを使用し、
      false の場合は仮想パススタイルを使用します。詳細については、[AWS S3 ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) を参照してください。
      
      一部のプロバイダ (例: AWS、Aliyun OSS、Netease COS、または Tencent COS) では、
      false にする必要があります - rclone は、プロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      v2 認証を使用する場合は true を指定します。
      
      これが false (デフォルト) の場合、rclone は v4 認証を使用します。
      設定されている場合、rclone は v2 認証を使用します。
      
      v4 署名が機能しない場合にのみ使用してください。たとえば、Jewel/v10 CEPH より前のバージョンの場合。

   --use-accelerate-endpoint
      true の場合、AWS S3 の高速エンドポイントを使用します。
      
      [AWS S3 転送アクセラレーション](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html) を参照してください。

   --leave-parts-on-error
      true の場合、失敗時にアボートアップロードを呼び出さず、マニュアルで回復できるように S3 に正常にアップロードされたパートがすべて残ります。
      
      異なるセッションでアップロードの再開を行う場合には true に設定する必要があります。
      
      警告: 不完全なマルチパートアップロードのパートを保存すると、S3 上のスペース使用量にカウントされ、クリーンアップしない場合に追加のコストが発生します。
      

   --list-chunk
      リストのチャンクサイズ (各 ListObject S3 リクエストのレスポンスリスト)。
      
      このオプションは、AWS S3 仕様の "MaxKeys"、"max-items"、または "page-size" としても知られています。
      大多数のサービスは、要求した数よりも多くのオブジェクトを含む応答リストを切り捨てます。
      AWS S3 では、これはグローバル最大値であり、変更することはできません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) を参照してください。
      Ceph の場合、"rgw list buckets max chunk" オプションで増やすことができます。
      

   --list-version
      使用する ListObjects のバージョン: 1、2、または auto のいずれかを指定します。
      
      S3 が最初にリリースされた当初、バケット内のオブジェクトを列挙するための ListObjects 呼び出ししか提供されていませんでした。
      
      しかし、2016 年 5 月に ListObjectsV2 呼び出しが導入されました。これははるかに高性能であり、可能な限り使用する必要があります。
      
      0 に設定されている場合、rclone は list objects メソッドを呼び出すためのプロバイダに従って推測します。推測が間違っている場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストを URL エンコードするかどうか: true/false/unset
      
      一部のプロバイダはリストを URL エンコードし、利用できる場合には制御文字をファイル名に使用する場合にこれがより信頼性が高くなります。設定が unset になっている場合 (デフォルト)、rclone はプロバイダの設定に従って適用するため、rclone の選択をここで上書きできます。
      

   --no-check-bucket
      エラーバケットが存在するか、作成を試みないようにします。
      
      バケットが既に存在することを事前にわかっている場合、rclone が行うトランザクションの数を最小限にするために便利です。
      
      また、使用するユーザーにバケット作成権限がない場合にも必要です。v1.52.0 より前のバージョンでは、これは無言で渡されていましたが、バグのために通過してしまいました。
      

   --no-head
      アップロードしたオブジェクトに HEAD で整合性を確認しないようにします。
      
      rclone は PUT 後に 200 OK メッセージを受信した場合、正しくアップロードされたと想定します。
      
      特に、次を想定します:
      
      - アップロードされたときのメタデータ、変更時刻、ストレージクラス、コンテンツタイプがアップロードしたままであったこと
      - サイズがアップロードしたままであったこと
      
      単一パートの PUT の応答から次の項目を読み取ります:
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明のソースオブジェクトがアップロードされる場合、rclone は HEAD リクエストを実行します。
      
      このフラグを設定すると、アップロードの隠れた失敗の確率が増えるため、通常の操作には推奨されません。実際には、このフラグが設定されている場合、アップロードの隠れた失敗の確率は非常に低いです。

   --no-head-object
      オブジェクトを取得する前に HEAD を実行しないようにします。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[エンコーディングセクション](/overview/#encoding)の概要を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファ (たとえばマルチパート) が必要なアップロード時には、メモリプールがアロケーションに使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールで mmap バッファを使用するかどうか。

   --disable-http2
      S3 バックエンドでの http2 の使用を無効にします。
      
      現在、s3 (特に minio) バックエンドと HTTP/2 に関して未解決の問題があります。HTTP/2 は s3 バックエンドのデフォルトで有効になっていますが、ここで無効にすることもできます。問題が解決されると、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常は AWS S3 がクラウドフロント CDN URL に設定されています。
      AWS S3 は、クラウドフロントネットワークを介してダウンロードされたデータに対してより安価な転送を提供します。

   --use-multipart-etag
      マルチパートアップロードで ETag を使用して整合性を確認するかどうか

      true、false、またはデフォルトを使用するために設定してください

   --use-presigned-request
      シングルパートのアップロードに署名済みのリクエストまたは PutObject を使用するかどうか

      false の場合、rclone は AWS SDK の PutObject を使用してオブジェクトをアップロードします。

      rclone のバージョン < 1.59 では、シングルパートオブジェクトをアップロードするための署名済みのリクエストが使用されます。このフラグを true に設定すると、その機能が再度有効になります。これは特別な状況やテストのために必要です。

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時間におけるファイルバージョンを表示します。
      
      パラメータは、"2006-01-02" のような日付、"2006-01-02 15:04:05" のような日時、あるいはそれ以前の時間の持続時間 ("100d" や "1h" など) です。
      
      これを使用する場合、ファイルの書き込み操作は許可されないため、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[time オプションドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      設定すると、gzip でエンコードされたオブジェクトを解凍します。
      
      S3 には "Content-Encoding: gzip" が設定された状態でオブジェクトをアップロードすることができます。通常、rclone はこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグを設定すると、rclone は受信したときに "Content-Encoding: gzip" でこれらのファイルを解凍します。つまり、rclone はサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトを gzip 圧縮する場合に設定してください。
      
      通常、プロバイダはオブジェクトがダウンロードされると変更しません。"Content-Encoding: gzip" が設定されていない場合はダウンロードされません。
      
      ただし、一部のプロバイダ (たとえば Cloudflare) は、"Content-Encoding: gzip" が設定されていない場合でもオブジェクトを gzip 圧縮する場合があります。
      
      これによる症状は、次のようなエラーが発生することです:
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rclone が Content-Encoding: gzip が設定されており、チャンク化された伝送エンコーディングでオブジェクトをダウンロードすると、rclone はオブジェクトをその場で解凍します。
      
      これが unset (デフォルト) に設定されている場合、rclone はプロバイダの設定に従って適用するように選択しますが、ここで rclone の選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します。

   --sts-endpoint
      STS のエンドポイント。
      
      AWS を使用する場合は空にして、リージョンのデフォルトエンドポイントを使用します。


OPTIONS:
   --access-key-id value           AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                     バケットおよびオブジェクトの作成時に使用される Canned ACL。[$ACL]
   --endpoint value                S3 API のエンドポイント。[$ENDPOINT]
   --env-auth                      実行時に AWS の認証情報を取得します (環境変数または環境に設定されている EC2/ECS のメタデータ)。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示
   --location-constraint value     ロケーション制約 - リージョンに一致する必要があります。[$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。[$REGION]
   --secret-access-key value       AWS Secret Access Key (パスワード)。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3 にこのオブジェクトを格納するために使用されるサーバーサイドの暗号化アルゴリズム。[$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID を使用する場合はキーの ARN を指定する必要があります。[$SSE_KMS_KEY_ID]
   --storage-class value           新しいオブジェクトを S3 に格納する際に使用するストレージクラス。[$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用される Canned ACL。[$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定すると、gzip でエンコードされたオブジェクトを解凍します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータに MD5 チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 バックエンドでの http2 の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true の場合、パススタイルのアクセスを使用し、false の場合は仮想ホストスタイルを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --leave-parts-on-error           true の場合、失敗時にアボートアップロードを呼び出さず、マニュアルで回復できるように S3 に正常にアップロードされたパートがすべて残ります。 (デフォルト: false) [$LEAVE_PARTS_ON_ERROR]
   --list-chunk value               リストのチャンクサイズ (各 ListObject S3 リクエストのレスポンスリスト)。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストを URL エンコードするかどうか: true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用する ListObjects の