# 中国移动Ecloud弹性对象存储（EOS）

{% code fullWidth="true" %}
```
名前：
   singularity storage update s3 chinamobile- 中国移動Ecloud弾性オブジェクトストレージ（EOS）のアップデート

使用方法：
   singularity storage update s3 chinamobile [コマンドオプション] <名前|ID>

説明：
   --env-auth
      実行時にAWSクレデンシャル（環境変数または環境変数がない場合のEC2/ECSメタデータ）を取得します。
      
      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例：
         | false | 次のステップでAWSの認証情報を入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWS Access Key IDです。
      
      無名アクセスまたはランタイムの認証情報の場合は空白のままにしてください。

   --secret-access-key
      AWS Secret Access Key（パスワード）です。
      
      無名アクセスまたはランタイムの認証情報の場合は空白のままにしてください。

   --endpoint
      中国移動Ecloud弾性オブジェクトストレージ（EOS）APIのエンドポイントです。

      例：
         | eos-wuxi-1.cmecloud.cn       | デフォルトのエンドポイント - 迷った場合の良い選択です。
         |                              | 東中国（蘇州）
         | eos-jinan-1.cmecloud.cn      | 東中国（蘇州）
         | eos-ningbo-1.cmecloud.cn     | 東中国（杭州）
         | eos-shanghai-1.cmecloud.cn   | 東中国（上海-1）
         | eos-zhengzhou-1.cmecloud.cn  | 中央中国（鄭州）
         | eos-hunan-1.cmecloud.cn      | 中央中国（長沙-1）
         | eos-zhuzhou-1.cmecloud.cn    | 中央中国（長沙-2）
         | eos-guangzhou-1.cmecloud.cn  | 南中国（広州-2）
         | eos-dongguan-1.cmecloud.cn   | 南中国（広州-3）
         | eos-beijing-1.cmecloud.cn    | 北中国（北京-1）
         | eos-beijing-2.cmecloud.cn    | 北中国（北京-2）
         | eos-beijing-4.cmecloud.cn    | 北中国（北京-3）
         | eos-huhehaote-1.cmecloud.cn  | 北中国（フフホト）
         | eos-chengdu-1.cmecloud.cn    | 西南中国（成都）
         | eos-chongqing-1.cmecloud.cn  | 西南中国（重慶）
         | eos-guiyang-1.cmecloud.cn    | 西南中国（貴陽）
         | eos-xian-1.cmecloud.cn       | 西南中国（西安）
         | eos-yunnan.cmecloud.cn       | 云南中国（昆明）
         | eos-yunnan-2.cmecloud.cn     | 云南中国（昆明-2）
         | eos-tianjin-1.cmecloud.cn    | 天津中国（天津）
         | eos-jilin-1.cmecloud.cn      | 吉林中国（長春）
         | eos-hubei-1.cmecloud.cn      | 湖北中国（湘陰）
         | eos-jiangxi-1.cmecloud.cn    | 江西中国（南昌）
         | eos-gansu-1.cmecloud.cn      | 甘粛中国（蘭州）
         | eos-shanxi-1.cmecloud.cn     | 山西中国（太原）
         | eos-liaoning-1.cmecloud.cn   | 遼寧中国（瀋陽）
         | eos-hebei-1.cmecloud.cn      | 河北中国（石家庄）
         | eos-fujian-1.cmecloud.cn     | 福建中国（厦門）
         | eos-guangxi-1.cmecloud.cn    | 広西中国（南寧）
         | eos-anhui-1.cmecloud.cn      | 安徽中国（懐南）

   --location-constraint
      エンドポイントと一致する場所の制約です。
      
      バケットの作成時にのみ使用されます。

      例：
         | wuxi1      | 東中国（蘇州）
         | jinan1     | 東中国（蘇州）
         | ningbo1    | 東中国（杭州）
         | shanghai1  | 東中国（上海-1）
         | zhengzhou1 | 中央中国（鄭州）
         | hunan1     | 中央中国（長沙-1）
         | zhuzhou1   | 中央中国（長沙-2）
         | guangzhou1 | 南中国（広州-2）
         | dongguan1  | 南中国（広州-3）
         | beijing1   | 北中国（北京-1）
         | beijing2   | 北中国（北京-2）
         | beijing4   | 北中国（北京-3）
         | huhehaote1 | 北中国（フフホト）
         | chengdu1   | 西南中国（成都）
         | chongqing1 | 西南中国（重慶）
         | guiyang1   | 西南中国（貴陽）
         | xian1      | 西南中国（西安）
         | yunnan     | 云南中国（昆明）
         | yunnan2    | 云南中国（昆明-2）
         | tianjin1   | 天津中国（天津）
         | jilin1     | 吉林中国（長春）
         | hubei1     | 湖北中国（湘陰）
         | jiangxi1   | 江西中国（南昌）
         | gansu1     | 甘粛中国（蘭州）
         | shanxi1    | 山西中国（太原）
         | liaoning1  | 遼寧中国（瀋陽）
         | hebei1     | 河北中国（石家庄）
         | fujian1    | 福建中国（厦門）
         | guangxi1   | 広西中国（南寧）
         | anhui1     | 安徽中国（懐南）

   --acl
      オブジェクトを作成し、または保存、コピーする際に使用されるCanned ACLです。
      
      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。
      
      詳細については、[Amazon S3の公式ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3では、サーバーサイドでオブジェクトをコピーする際にACLはコピーされず、代わりに新しいACLが書き込まれます。
      
      aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACLです。
      
      詳細については、[Amazon S3の公式ドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      bucket_aclが設定されていない場合にのみ、バケットの作成時に使用されます。
      
      aclとbucket_aclの両方が空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。
      

      例：
         | private            | オーナーにFULL_CONTROLの権限があります。
         |                    | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLの権限があります。
         |                    | AllUsersグループにREAD権限があります。
         | public-read-write  | オーナーにFULL_CONTROLの権限があります。
         |                    | AllUsersグループにREADとWRITEの権限があります。
         |                    | バケットでこの設定を使用することは一般的にはお勧めできません。
         | authenticated-read | オーナーにFULL_CONTROLの権限があります。
         |                    | AuthenticatedUsersグループにREAD権限があります。

   --server-side-encryption
      S3にこのオブジェクトを保存する際に使用するサーバーサイドの暗号化アルゴリズムです。

      例：
         | <unset> | 無し
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-Cを使用する場合、S3にこのオブジェクトを保存する際に使用するサーバーサイドの暗号化アルゴリズムです。

      例：
         | <unset> | 無し
         | AES256  | AES256

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化/復号化に使用する秘密の暗号化キーを提供することができます。
      
      代わりに--sse-customer-key-base64を指定することもできます。

      例：
         | <unset> | 無し

   --sse-customer-key-base64
      SSE-Cを使用する場合、データの暗号化/復号化に使用する秘密の暗号化キーをBase64形式で提供する必要があります。
      
      代わりに--sse-customer-keyを指定することもできます。

      例：
         | <unset> | 無し

   --sse-customer-key-md5
      SSE-Cを使用する場合、秘密の暗号化キーのMD5チェックサムを指定することができます（オプション）。
      
      空白の場合、sse_customer_keyから自動的に計算されます。
      

      例：
         | <unset> | 無し

   --storage-class
      中国移動の新しいオブジェクトを保存する際に使用するストレージクラスです。

      例：
         | <unset>     | デフォルト
         | STANDARD    | 標準ストレージクラス
         | GLACIER     | アーカイブストレージモード
         | STANDARD_IA | 低頻度アクセスストレージモード

   --upload-cutoff
      チャンク化アップロードに切り替えるための閾値です。
      
      この閾値よりも大きなサイズのファイルはchunk_sizeの大きさのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクのサイズです。
      
      upload_cutoffを超えるサイズのファイル、またはサイズが不明なファイル（例：「rclone rcat」または「rclone mount」、Googleフォト、Googleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートのアップロードが行われます。
      
      "--s3-upload-concurrency"チャンクにつき、このサイズのバッファがメモリに格納されます。
      
      高速リンクで大きなファイルを転送し、十分なメモリがある場合、これを増やすと転送が高速化されます。
      
      キャッシュ制限を回避するため、rcloneは既知のサイズの大きなファイルのアップロード時にチャンクサイズを自動的に増やすこともできます。最大10,000のチャンク制限を超えないようにします。
      
      不明なサイズのファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大10,000のチャンクがあるため、デフォルトのストリームアップロード可能なファイルサイズの最大値は48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクのサイズを増やすと、進行状況統計情報の精度が低下します。rcloneは、AWS SDKによってチャンクがバッファに格納されたときにチャンクを送信したものとして処理しますが、実際にはまだアップロード中の場合があります。
      

   --max-upload-parts
      マルチパートのアップロードで使用するパートの最大数を定義します。
      
      マルチパートのアップロードを実行する際に、このオプションは使用するマルチパートチャンクの最大数を定義します。
      
      プロバイダがAWS S3の10,000チャンクの仕様をサポートしていない場合、これが役立つ場合があります。
      
      rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やし、このチャンク数の制限を下回るようにします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるための閾値です。
      
      サーバーサイドでコピーする必要があるこの閾値よりも大きなファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードの開始には長時間かかる場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパスです。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」という環境変数を探します。環境変数の値が空の場合は、カレントユーザーのホームディレクトリがデフォルトになります。
      
          Linux/OSX："$HOME/.aws/credentials"
          Windows: "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイルです。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用することができます。この変数はそのファイルで使用されるプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」がデフォルトになります。
      

   --session-token
      AWSセッショントークンです。

   --upload-concurrency
      マルチパートのアップロードの並列性です。
      
      同じファイルのマルチパートチャンクを同時にアップロードする数です。
      
      高速リンクを介して大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に利用しない場合は、これを増やすと転送が高速化される場合があります。

   --force-path-style
      trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式を使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパス形式のアクセスを使用します。falseの場合、rcloneは仮想パス形式を使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、Tencent COSなど）では、この設定が必要ですが、rcloneはプロバイダの設定に基づいて自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      これがfalse（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャが機能しない場合にのみ、v2認証を使用してください（たとえば、Jewel/v10 CEPHの場合）。

   --list-chunk
      リスティングのチャンクのサイズ（各ListObject S3リクエストの応答リスト）です。
      
      このオプションは、AWS S3のMaxKeys、max-items、またはpage-sizeとしても知られています。
      大抵のサービスは、リクエストされたよりも多くのオブジェクトを含む応答リストを切り捨てます。AWS S3では、これはグローバルな制限であり、変更できません。詳しくは[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションで増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）です。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高速であり、可能な限り使用する必要があります。
      
      デフォルトの0に設定されている場合、rcloneは設定されているプロバイダに基づいてどのリストオブジェクトのメソッドを呼び出すかを推測します。誤った推測をする場合は、ここで手動で設定することがあります。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダは、リストをURLエンコードすることをサポートし、制御文字をファイル名に含める場合にこれがより信頼性のあるオプションです。unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択します。
      

   --no-check-bucket
      エラーメッセージ「permission denied」を防ぐため、バケットの存在を確認せず、または作成しないようにします。
      
      バケットが既に存在する場合、rcloneのトランザクション数を最小限に抑えるため、このオプションは役立ちます。
      
      バケットの作成権限がない場合、必要になる場合があります。v1.52.0より前のバージョンでは、これはバグにより静かにパスしていました。
      

   --no-head
      オブジェクトのアップロード確認のためにアップロードしたオブジェクトをHEADリクエストしません。
      
      rcloneは100 OKメッセージを受信すると、PUTでオブジェクトが正常にアップロードされたと想定します。
      
      特に次のことが想定されます。
      
      - アップロードされたときのメタデータ（modtime、ストレージクラス、コンテンツタイプ）がアップロードされた内容と同じであること
      - サイズがアップロードされた内容と同じであること
      
      シングルパートのPUTの応答ですべてのアップロードを読み込みます。
      
      - MD5SUM
      - アップロード日付
      
      マルチパートのアップロードでは、これらの項目は読み込まれません。
      
      サイズが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを実行します。
      
      このフラグを設定すると、正常なアップロードの失敗の確率が高まるため、通常の操作には推奨されません。実際には、このフラグを設定してもアップロードの失敗が検出される確率は非常に低いです。
      

   --no-head-object
      オブジェクトを取得する際に、HEADの前にGETを行わないようにします。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度です。
      
      追加のバッファ（マルチパートなど）を必要とするアップロードでは、メモリプールを使用して割り当てを行います。
      このオプションは、未使用のバッファがプールから削除されるタイミングを制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうかです。

   --disable-http2
      S3バックエンドのhttp2の利用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2の問題が解決されていません。s3バックエンドのデフォルトではHTTP/2が有効になっていますが、ここでは無効にすることができます。問題が解決したら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673、https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイントです。
      通常、AWS S3では、CloudFrontネットワークを介してダウンロードされたデータはより安価な流出となりますので、クラウドフロントCDNのURLに設定されます。

   --use-multipart-etag
      マルチパートアップロード時にETagを使用して検証するかどうか
      
      これはtrue、false、またはデフォルトのプロバイダを使用して設定されていない場合のいずれかである必要があります。
      

   --use-presigned-request
      シングルパートのアップロードに署名付きリクエストまたはPutObjectを使用するかどうか
      
      falseに設定すると、rcloneはAWS SDKからPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン< 1.59では、署名付きリクエストを使用してシングルパートオブジェクトをアップロードし、このフラグをtrueに設定すると、この機能が再度有効になります。これは例外的な状況やテスト目的以外では必要ありません。
      

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付「2006-01-02」、日時「2006-01-02 15:04:05」、またはその時間前の期間、「100d」または「1h」などで指定することができます。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      設定すると、gzip形式でエンコードされたオブジェクトを解凍します。
      
      "Content-Encoding: gzip"が設定されている状態でオブジェクトをS3にアップロードすることが可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを受け取った通りに"Content-Encoding: gzip"で解凍します。これにより、rcloneはサイズとハッシュをチェックすることはできませんが、ファイルの内容が解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzip形式で圧縮する場合に設定してください。
      
      通常のプロバイダは、オブジェクトをダウンロードする際に変更しません。オブジェクトが`Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時にも設定されません。
      
      ただし、一部のプロバイダは、`Content-Encoding: gzip`でアップロードされていないにもかかわらずオブジェクトをgzip形式で圧縮する場合があります（例：Cloudflare）。
      
      これが設定されており、rcloneがContent-Encoding: gzipが設定されたオブジェクトをチャンク化転送エンコーディングでダウンロードする場合、rcloneはオブジェクトをリアルタイムで解凍します。
      
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択します。
      

   --no-system-metadata
      システムメタデータの設定と読み込みを抑制します。


オプション：
   --access-key-id value           AWS Access Key ID。[$ACCESS_KEY_ID]
   --acl value                     オブジェクトを作成し、または保存、コピーする際に使用されるCanned ACL。[$ACL]
   --endpoint value                中国移動Ecloud弾性オブジェクトストレージ（EOS）APIのエンドポイント。[$ENDPOINT]
   --env-auth                      実行時にAWSクレデンシャル（環境変数または環境変数がない場合のEC2/ECSメタデータ）を取得します。 (default: false) [$ENV_AUTH]
   --help, -h                      ヘルプを表示します。
   --location-constraint value     エンドポイントと一致する場所の制約。[$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS Secret Access Key（パスワード）。[$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3にこのオブジェクトを保存する際に使用するサーバーサイドの暗号化アルゴリズム。[$SERVER_SIDE_ENCRYPTION]
   --storage-class value           中国移動の新しいオブジェクトを保存する際に使用するストレージクラス。[$STORAGE_CLASS]

   詳細

   --bucket-acl value               バケットの作成時に使用されるCanned ACL。[$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクのサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるための閾値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     設定すると、gzip形式でエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の利用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               trueの場合、パス形式のアクセスを使用し、falseの場合は仮想ホスト形式を使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リスティングのチャンクのサイズ。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1,2,または0（自動） (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートのアップロードで使用するパートの最大数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip形式で圧縮する場合に設定してください。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                エラーメッセージ「permission denied」を防ぐため、バケットの存在を確認せず、または作成しないようにします。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        オブジェクトのアップロード確認のためにアップロードしたオブジェクトをHEADリクエストしません。 (default: false) [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する際に、HEADの前にGETを行わないようにします。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み込みを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWSセッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。[$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-Cを使用する場合、S3にこのオブジェクトを保存する際に使用されるサーバーサイドの暗号化アルゴリズム。[$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データの暗号化/復号化に使用する秘密の暗号化キーを提供することができます。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-Cを使用する場合、データの暗号化/復号化に使用する秘密の暗号化キーをBase64形式で提供する必要があります。[$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-Cを使用する場合、秘密の暗号化キーのMD5チェックサムを指定することができます（オプション）。[$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       マルチパートのアップロードの並列性。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードに切り替えるための閾値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロード時にETagを使用して検証するかどうか。 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートのアップロードに署名付きリクエストまたはPutObjectを使用するかどうか。 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

```
{% endcode %}