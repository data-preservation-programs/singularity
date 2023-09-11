# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
名称：
   singularity storage create s3 aws - Amazon Web Services (AWS) S3

使用方法：
   singularity storage create s3 aws [コマンドオプション] [引数...]

説明：
   --env-auth
      実行時（環境変数またはEC2/ECSメタデータが利用できない場合）からAWSの認証情報を取得します。
      
      access_key_idおよびsecret_access_keyが空の場合にのみ適用されます。
      
      利用例:
         | false | 次の手順でAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報を利用する場合は、空白のままにしてください。

   --secret-access-key
      AWSシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報を利用する場合は、空白のままにしてください。

   --region
      接続するリージョン。

      利用例:
         | us-east-1      | デフォルトのエンドポイント - 迷った場合の良い選択肢です。
         |                | 米国リージョン、北バージニアまたは太平洋北西部。
         |                | 場所の制約は空白のままにしてください。
         | us-east-2      | 米国東部（オハイオ）リージョン。
         |                | 場所の制約はus-east-2にする必要があります。
         | us-west-1      | 米国西部（北カリフォルニア）リージョン。
         |                | 場所の制約はus-west-1にする必要があります。
         | us-west-2      | 米国西部（オレゴン）リージョン。
         |                | 場所の制約はus-west-2にする必要があります。
         | ca-central-1   | カナダ（中央）リージョン。
         |                | 場所の制約はca-central-1にする必要があります。
         | eu-west-1      | EU（アイルランド）リージョン。
         |                | 場所の制約はEUまたはeu-west-1にする必要があります。
         | eu-west-2      | EU（ロンドン）リージョン。
         |                | 場所の制約はeu-west-2にする必要があります。
         | eu-west-3      | EU（パリ）リージョン。
         |                | 場所の制約はeu-west-3にする必要があります。
         | eu-north-1     | EU（ストックホルム）リージョン。
         |                | 場所の制約はeu-north-1にする必要があります。
         | eu-south-1     | EU（ミラノ）リージョン。
         |                | 場所の制約はeu-south-1にする必要があります。
         | eu-central-1   | EU（フランクフルト）リージョン。
         |                | 場所の制約はeu-central-1にする必要があります。
         | ap-southeast-1 | アジアパシフィック（シンガポール）リージョン。
         |                | 場所の制約はap-southeast-1にする必要があります。
         | ap-southeast-2 | アジアパシフィック（シドニー）リージョン。
         |                | 場所の制約はap-southeast-2にする必要があります。
         | ap-northeast-1 | アジアパシフィック（東京）リージョン。
         |                | 場所の制約はap-northeast-1にする必要があります。
         | ap-northeast-2 | アジアパシフィック（ソウル）リージョン。
         |                | 場所の制約はap-northeast-2にする必要があります。
         | ap-northeast-3 | アジアパシフィック（大阪-ローカル）リージョン。
         |                | 場所の制約はap-northeast-3にする必要があります。
         | ap-south-1     | アジアパシフィック（ムンバイ）リージョン。
         |                | 場所の制約はap-south-1にする必要があります。
         | ap-east-1      | アジアパシフィック（香港）リージョン。
         |                | 場所の制約はap-east-1にする必要があります。
         | sa-east-1      | 南米（サンパウロ）リージョン。
         |                | 場所の制約はsa-east-1にする必要があります。
         | me-south-1     | 中東（バーレーン）リージョン。
         |                | 場所の制約はme-south-1にする必要があります。
         | af-south-1     | アフリカ（ケープタウン）リージョン。
         |                | 場所の制約はaf-south-1にする必要があります。
         | cn-north-1     | 中国（北京）リージョン。
         |                | 場所の制約はcn-north-1にする必要があります。
         | cn-northwest-1 | 中国（寧夏）リージョン。
         |                | 場所の制約はcn-northwest-1にする必要があります。
         | us-gov-east-1  | AWS GovCloud（米国-東部）リージョン。
         |                | 場所の制約はus-gov-east-1にする必要があります。
         | us-gov-west-1  | AWS GovCloud（米国）リージョン。
         |                | 場所の制約はus-gov-west-1にする必要があります。

   --endpoint
      S3 APIのエンドポイント。
      
      AWSを使用する場合は、リージョンのデフォルトエンドポイントを使用する場合は空白のままにしてください。

   --location-constraint
      リージョンに一致する場所の制約。
      
      バケットを作成する場合にのみ使用されます。

      利用例:
         | <unset>        | 米国リージョン、北バージニア、または太平洋北西部は空白です
         | us-east-2      | 米国東部（オハイオ）リージョン
         | us-west-1      | 米国西部（北カリフォルニア）リージョン
         | us-west-2      | 米国西部（オレゴン）リージョン
         | ca-central-1   | カナダ（中央）リージョン
         | eu-west-1      | EU（アイルランド）リージョン
         | eu-west-2      | EU（ロンドン）リージョン
         | eu-west-3      | EU（パリ）リージョン
         | eu-north-1     | EU（ストックホルム）リージョン
         | eu-south-1     | EU（ミラノ）リージョン
         | EU             | EUリージョン
         | ap-southeast-1 | アジアパシフィック（シンガポール）リージョン
         | ap-southeast-2 | アジアパシフィック（シドニー）リージョン
         | ap-northeast-1 | アジアパシフィック（東京）リージョン
         | ap-northeast-2 | アジアパシフィック（ソウル）リージョン
         | ap-northeast-3 | アジアパシフィック（大阪-ローカル）リージョン
         | ap-south-1     | アジアパシフィック（ムンバイ）リージョン
         | ap-east-1      | アジアパシフィック（香港）リージョン
         | sa-east-1      | 南米（サンパウロ）リージョン
         | me-south-1     | 中東（バーレーン）リージョン
         | af-south-1     | アフリカ（ケープタウン）リージョン
         | cn-north-1     | 中国（北京）リージョン
         | cn-northwest-1 | 中国（寧夏）リージョン
         | us-gov-east-1  | AWS GovCloud（米国-東部）リージョン
         | us-gov-west-1  | AWS GovCloud（米国）リージョン

   --acl
      バケットの作成およびオブジェクトの保存またはコピー時に使用される事前定義のACL。
      
      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      S3では、サーバーサイドでオブジェクトをコピーする際にACLはソースからコピーされず、新しいACLが書き込まれます。
      
      Aclが空の文字列の場合、X-Amz-Acl:ヘッダは追加されず、デフォルト（プライベート）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用される事前定義のACL。
      
      詳細については、https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl を参照してください。
      
      このACLはバケットの作成時にのみ適用されます。設定されていない場合は、"acl"が代わりに使用されます。
      
      「acl」と「bucket_acl」が空の文字列の場合、X-Amz-Acl:ヘッダは追加されず、デフォルト（プライベート）が使用されます。
      

      利用例:
         | private            | オーナーにFULL_CONTROLが与えられます。
         |                    | 他の誰もアクセス権限はありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADアクセスが与えられます。
         | public-read-write  | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADおよびWRITEアクセスが与えられます。
         |                    | バケットでこれを許可することは一般的に推奨されません。
         | authenticated-read | オーナーにFULL_CONTROLが与えられます。
         |                    | AuthenticatedUsersグループにREADアクセスが与えられます。

   --requester-pays
      S3バケットとのやり取り時にリクエスト元が支払うオプションを有効にします。

   --server-side-encryption
      S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      利用例:
         | <unset> | なし
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-Cを使用している場合、S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。

      利用例:
         | <unset> | なし
         | AES256  | AES256

   --sse-kms-key-id
      KMS IDを使用している場合は、キーのARNを指定する必要があります。

      利用例:
         | <unset>                 | なし
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化/復号に使用される秘密の暗号化キーを指定できます。
      
      代わりに--sse-customer-key-base64を指定することもできます。

      利用例:
         | <unset> | なし

   --sse-customer-key-base64
      SSE-Cを使用する場合、データの暗号化/復号に使用する秘密の暗号化キーをBase64形式で指定する必要があります。
      
      代わりに--sse-customer-keyを指定することもできます。

      利用例:
         | <unset> | なし

   --sse-customer-key-md5
      SSE-Cを使用している場合、秘密の暗号化キーのMD5チェックサムを指定できます（オプション）。
      
      空白の場合は、sse_customer_keyから自動的に計算されます。
      

      利用例:
         | <unset> | なし

   --storage-class
      新しいオブジェクトをS3に保存する際に使用するストレージクラス。

      利用例:
         | <unset>             | デフォルト
         | STANDARD            | 標準のストレージクラス
         | REDUCED_REDUNDANCY  | 減少冗長性のストレージクラス
         | STANDARD_IA         | 標準の低頻度アクセスのストレージクラス
         | ONEZONE_IA          | 1か所の低頻度アクセスのストレージクラス
         | GLACIER             | Glacierのストレージクラス
         | DEEP_ARCHIVE        | Glacier Deep Archiveのストレージクラス
         | INTELLIGENT_TIERING | Intelligent-Tieringのストレージクラス
         | GLACIER_IR          | Glacier Instant Retrievalのストレージクラス

   --upload-cutoff
      チャンクアップロードへの切り替えのためのカットオフ。
      
      このサイズより大きなファイルは、チャンクサイズの大きさのチャンクでアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロード時のチャンクサイズ。
      
      upload_cutoffより大きなファイルや、サイズが不明なファイル（たとえば「rclone rcat」からのアップロード、
      「rclone mount」でアップロードされたファイル、GoogleフォトやGoogleドキュメントなど）をアップロードする場合、
      このチャンクサイズを使用して複数パートのアップロードを行います。
      
      注意："--s3-upload-concurrency"は1つの転送ごとにこのチャンクサイズのチャンクをバッファリングします。
      
      高速リンクを介して大きなファイルを転送し、メモリが十分にある場合は、チャンクサイズを増やすと転送速度が速くなります。
      
      rcloneは、既知のサイズの大きなファイルをアップロードする場合、10,000パートの制限を下回るように自動的にチャンクサイズを増やします。
      
      未知のサイズのファイルは、設定されたchunk_sizeでアップロードされます。
      デフォルトのchunk_sizeは5 MiBで、最大で10,000のチャンクがあるため、「rclone rcat」または「rclone mount」またはGoogleフォトまたはGoogleドキュメントからアップロードされる
      サイズのわからないファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況に関する統計情報の正確性が低下します。
      rcloneは、AWS SDKでバッファリングされるチャンクを送信した時点でチャンクを送信したものとみなしますが、実際はまだアップロード中です。
      チャンクサイズが大きくなると、AWS SDKのバッファが大きくなり、進行状況が真実からずれた報告が行われます。
      

   --max-upload-parts
      マルチパートアップロードで使用するパートの最大数。
      
      このオプションは、マルチパートアップロード時の使用するパートの最大数を定義します。
      
      これはAWS S3の10,000パートの仕様をサポートしていないサービスがある場合に役立ちます。
      
      rcloneは、既知のサイズの大きなファイルをアップロードする場合、このパート数の制限を下回るように自動的にチャンクサイズを増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。
      
      サーバーサイドでコピーする必要があるこのサイズより大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、
      大きなファイルのアップロードの開始には長い遅延が発生する可能性があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      もしenv_auth = trueなら、rcloneは共有認証情報ファイルを使用することができます。
      
      空の場合、rcloneは"AWS_SHARED_CREDENTIALS_FILE"環境変数を参照します。
      環境変数の値が空の場合は、デフォルトで現在のユーザーのホームディレクトリが使用されます。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      もしenv_auth = trueなら、rcloneは共有認証情報ファイルを使用することができます。
      この変数は、そのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数"AWS_PROFILE"または"default"が設定されていない場合にデフォルトで使用されます。
      

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行性。
      
      これにより、同じファイルの複数のチャンクが同時にアップロードされます。
      
      高速リンク上で大量のファイルを高速リンク上でアップロードし、これらのアップロードが帯域幅を十分に使用していない場合は、この数値を増やすと転送速度が向上する場合があります。

   --force-path-style
      trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルのアクセスを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパススタイルのアクセスを使用します。
      これがfalseの場合、rcloneは仮想パススタイルを使用します。詳細は[AWS S3ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、Tencent COSなど）は、この設定に基づいて自動的に仮想パススタイルを設定する必要があります。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      これがfalse（デフォルト）の場合、rcloneはv4認証を使用します。
      これが設定されている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合のみ、つまりJewel/v10 CEPH以前の場合にのみ使用してください。

   --use-accelerate-endpoint
      trueの場合、AWS S3アクセラレートエンドポイントを使用します。
      
      参照：[AWS S3転送高速化](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)

   --leave-parts-on-error
      trueの場合、失敗時にアップロードを中止せず、すべての正常にアップロードされたパートをS3に残して手動で回復できます。
      
      異なるセッション間でアップロードを再開する場合はtrueに設定する必要があります。
      
      警告：未完全なマルチパートアップロードの一部を保存すると、S3のスペース使用量にカウントされ、クリーンアップされない場合に追加のコストが発生します。
      

   --list-chunk
      リストチャンクのサイズ（各ListObject S3リクエストの応答リスト）。
      
      このオプションは、AWS S3の仕様である「MaxKeys」、「max-items」、または「page-size」としても知られています。
      ほとんどのサービスは、リクエスト数が1000を超えても、1000オブジェクトまでしか応答リストを切り詰めません。
      AWS S3では、これはグローバルな制限であり、変更することはできません。詳細については[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、これを「rgw list buckets max chunk」オプションで増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または自動で0。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しが提供されました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高速な性能であり、可能な限り使用する必要があります。
      
      デフォルトで設定された0の場合、rcloneは設定されたプロバイダに従って、呼び出すリストオブジェクトメソッドを推測します。推測が誤っている場合は、ここで手動で設定することができます。
      

   --list-url-encode
      リストのURLエンコードの有無：true/false/unset
      
      一部のプロバイダでは、URLエンコードリストをサポートしており、利用可能な場合は、ファイル名に制御文字を使用する際にこれがより信頼性があります。これがunset（デフォルト）の場合、rcloneはプロバイダの設定に従って適用されるものを選択しますが、ここでrcloneの選択をオーバーライドできます。
      

   --no-check-bucket
      該当する場合、バケットの存在を確認せず、作成しようとしません。
      
      バケットが既に存在する場合や、トランザクション数を最小限に抑えようとしている場合に便利です。
      
      ユーザーがバケットの作成権限を持たない場合は、必要です。v1.52.0より前では、このバグのために黙ってパスしました。
      

   --no-head
      アップロードしたオブジェクトのヘッドリクエストを行って整合性をチェックしません。
      
      rcloneは、PUTでオブジェクトをアップロードした後、200 OKメッセージを受け取った場合、正しくアップロードされたものとみなすため、このフラグが設定されていると、
      オブジェクトが正しくアップロードされたと想定します。
      
      特に次のものを前提としています:
      
      - metadata、modtime、storage class、およびcontent typeがアップロード時と同じであること
      - サイズがアップロード時のものであること
      
      以下の項目を1つの部分PUTの応答から読み取ります:
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズのわからないソースオブジェクトがアップロードされる場合でも、rcloneはHEADリクエストを行います。
      
      このフラグを設定すると、アップロードの失敗が検出されない可能性が高くなります。特に、正しいサイズではない場合などですが、通常の操作では推奨されません。
      実際のところ、このフラグを設定していても、アップロードの失敗が検出される可能性は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する前にHEADを行いません。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファが必要なアップロード（たとえばマルチパート）では、割り当てにメモリプールが使用されます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでのmmapバッファの使用有無。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      現在、s3（具体的にはminio）バックエンドとHTTP/2の問題が未解決です。S3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。
      問題が解決されると、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3からのデータのダウンロードでは安価なデータ接出が提供されているため、CloudFront CDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを検証に使用するかどうか
      
      true、false、またはデフォルトを使用するためのアンセットに設定してください。
      

   --use-presigned-request
      シングルパートアップロードに対して署名済みのリクエストまたはPutObjectを使用するかどうか。
      
      これがfalseの場合、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン1.59未満では、シングルパートオブジェクトをアップロードするために署名付きリクエストを使用し、このフラグをtrueに設定すると、
      この機能が再度有効になります。これは特別な状況やテスト以外では必要ありません。
      

   --versions
      古いバージョンをディレクトリリストに含める。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付 "2006-01-02"、日時 "2006-01-02 15:04:05"、またはそれ以前の時間の期間（たとえば "100d" または "1h"）にする必要があります。
      
      ただし、これを使用する場合、ファイルの書き込み操作は許可されません。つまり、ファイルのアップロードや削除はできません。
      
      有効な形式については[時刻オプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      設定するとgzipでエンコードされたオブジェクトを復号します。
      
      S3には「Content-Encoding: gzip」が設定されたままオブジェクトをアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを受け取る際に「Content-Encoding: gzip」で解凍します。これにより、rcloneはサイズとハッシュをチェックできませんが、ファイルの内容は解凍されます。
      

   --might-gzip
      バックエンドによってオブジェクトがgzipで圧縮される可能性がある場合に設定します。
      
      通常、プロバイダはダウンロード時にオブジェクトを変更しません。`Content-Encoding: gzip`でアップロードされなかったオブジェクトは、ダウンロード時には設定されません。
      
      ただし、一部のプロバイダは、gipzで圧縮されていなくてもオブジェクトをgzipで圧縮する場合があります（例：Cloudflare）。
      
      これの症状は次のようなエラーの受信です。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、チャンク分割転送エンコードでオブジェクトをダウンロードする場合、rcloneはオブジェクトを逐次解凍します。
      
      アンセット（デフォルト）に設定する場合、rcloneはプロバイダの設定に従って適用されるものを選択しますが、ここでrcloneの選択をオーバーライドできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します。

   --sts-endpoint
      STSのエンドポイント。
      
      AWSを使用する場合は、リージョンのデフォルトエンドポイントを使用する場合は空白のままにしてください。

オプション：
   --access-key-id value           AWSアクセスキーID。 [$ACCESS_KEY_ID]
   --acl value                     バケットの作成およびオブジェクトの保存またはコピー時に使用される事前定義のACL。 [$ACL]
   --endpoint value                S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                      実行時（環境変数またはEC2/ECSメタデータが利用できない場合）からAWSの認証情報を取得します。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示します
   --location-constraint value     リージョンに一致する場所の制約。 [$LOCATION_CONSTRAINT]
   --region value                  接続するリージョン。 [$REGION]
   --secret-access-key value       AWSシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。 [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS IDを使用している場合は、キーのARNを指定する必要があります。 [$SSE_KMS_KEY_ID]
   --storage-class value           新しいオブジェクトをS3に保存する際に使用するストレージクラス。 [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               バケットの作成時に使用される事前定義のACL。 [$BUCKET_ACL]
   --chunk-size value               アップロード時のチャンクサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定するとgzipでエンコードされたオブジェクトを復号します。 (デフォルト: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。 (デフォルト: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パススタイルのアクセスを使用し、falseの場合は仮想ホストスタイルのアクセスを使用します。 (デフォルト: true) [$FORCE_PATH_STYLE]
   --leave-parts-on-error           trueの場合、失敗時にアップロードを中止せず、すべての正常にアップロードされたパートをS3に残して手動で回復できます。 (デフォルト: false) [$LEAVE_PARTS_ON_ERROR]
   --list-chunk value               リストチャンクのサイズ（各ListObject S3リクエストの応答リスト）。 (デフォルト: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードの有無：true/false/unset (デフォルト: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2または自動で0。 (デフォルト: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードで使用するパートの最大数。 (デフォルト: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。 (デフォルト: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでのmmapバッファの使用有無。 (デフォルト: false)