# IBM COS S3

{% code fullWidth="true" %}
```
名前：
   singularity storage create s3 ibmcos - IBM COS S3

使用法：
   singularity storage create s3 ibmcos [コマンドオプション] [引数...]

説明：
   --env-auth
      AWSの認証情報を実行時に取得します（環境変数または環境変数がない場合のEC2/ECSメタデータ）。
      
      access_key_idとsecret_access_keyが空白の場合のみ適用されます。

      例:
         | false | AWSの認証情報を次のステップで入力します。
         | true  | 環境（環境変数またはIAM）からAWSの認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにします。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにします。

   --region
      接続するリージョン。
      
      S3のクローンを使用していてリージョンがない場合は空白のままにします。

      例:
         | <未設定>            | 迷った場合はこれを使用します。
         |                    | v4シグネチャと空のリージョンを使用します。
         | other-v2-signature | v4シグネチャが動作しない場合のみ使用します。
         |                    | 例: Jewel/v10 CEPH以前。

   --endpoint
      IBM COS S3 APIのエンドポイント。
      
      IBM COS On Premiseを使用している場合に指定します。

      例:
         | s3.us.cloud-object-storage.appdomain.cloud                | US Cross Region エンドポイント
         | s3.dal.us.cloud-object-storage.appdomain.cloud            | US Cross Region Dallas エンドポイント
         | s3.wdc.us.cloud-object-storage.appdomain.cloud            | US Cross Region Washington DC エンドポイント
         | s3.sjc.us.cloud-object-storage.appdomain.cloud            | US Cross Region San Jose エンドポイント
         | s3.private.us.cloud-object-storage.appdomain.cloud        | US Cross Region プライベート エンドポイント
         | s3.private.dal.us.cloud-object-storage.appdomain.cloud    | US Cross Region Dallas プライベート エンドポイント
         | s3.private.wdc.us.cloud-object-storage.appdomain.cloud    | US Cross Region Washington DC プライベート エンドポイント
         | s3.private.sjc.us.cloud-object-storage.appdomain.cloud    | US Cross Region San Jose プライベート エンドポイント
         | s3.us-east.cloud-object-storage.appdomain.cloud           | US Region East エンドポイント
         | s3.private.us-east.cloud-object-storage.appdomain.cloud   | US Region East プライベート エンドポイント
         | s3.us-south.cloud-object-storage.appdomain.cloud          | US Region South エンドポイント
         | s3.private.us-south.cloud-object-storage.appdomain.cloud  | US Region South プライベート エンドポイント
         | s3.eu.cloud-object-storage.appdomain.cloud                | EU Cross Region エンドポイント
         | s3.fra.eu.cloud-object-storage.appdomain.cloud            | EU Cross Region Frankfurt エンドポイント
         | s3.mil.eu.cloud-object-storage.appdomain.cloud            | EU Cross Region Milan エンドポイント
         | s3.ams.eu.cloud-object-storage.appdomain.cloud            | EU Cross Region Amsterdam エンドポイント
         | s3.private.eu.cloud-object-storage.appdomain.cloud        | EU Cross Region プライベート エンドポイント
         | s3.private.fra.eu.cloud-object-storage.appdomain.cloud    | EU Cross Region Frankfurt プライベート エンドポイント
         | s3.private.mil.eu.cloud-object-storage.appdomain.cloud    | EU Cross Region Milan プライベート エンドポイント
         | s3.private.ams.eu.cloud-object-storage.appdomain.cloud    | EU Cross Region Amsterdam プライベート エンドポイント
         | s3.eu-gb.cloud-object-storage.appdomain.cloud             | Great Britain エンドポイント
         | s3.private.eu-gb.cloud-object-storage.appdomain.cloud     | Great Britain プライベート エンドポイント
         | s3.eu-de.cloud-object-storage.appdomain.cloud             | EU Region DE エンドポイント
         | s3.private.eu-de.cloud-object-storage.appdomain.cloud     | EU Region DE プライベート エンドポイント
         | s3.ap.cloud-object-storage.appdomain.cloud                | APAC Cross Regional エンドポイント
         | s3.tok.ap.cloud-object-storage.appdomain.cloud            | APAC Cross Regional Tokyo エンドポイント
         | s3.hkg.ap.cloud-object-storage.appdomain.cloud            | APAC Cross Regional HongKong エンドポイント
         | s3.seo.ap.cloud-object-storage.appdomain.cloud            | APAC Cross Regional Seoul エンドポイント
         | s3.private.ap.cloud-object-storage.appdomain.cloud        | APAC Cross Regional プライベート エンドポイント
         | s3.private.tok.ap.cloud-object-storage.appdomain.cloud    | APAC Cross Regional Tokyo プライベート エンドポイント
         | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud    | APAC Cross Regional HongKong プライベート エンドポイント
         | s3.private.seo.ap.cloud-object-storage.appdomain.cloud    | APAC Cross Regional Seoul プライベート エンドポイント
         | s3.jp-tok.cloud-object-storage.appdomain.cloud            | APAC Region Japan エンドポイント
         | s3.private.jp-tok.cloud-object-storage.appdomain.cloud    | APAC Region Japan プライベート エンドポイント
         | s3.au-syd.cloud-object-storage.appdomain.cloud            | APAC Region Australia エンドポイント
         | s3.private.au-syd.cloud-object-storage.appdomain.cloud    | APAC Region Australia プライベート エンドポイント
         | s3.ams03.cloud-object-storage.appdomain.cloud             | Amsterdam Single Site エンドポイント
         | s3.private.ams03.cloud-object-storage.appdomain.cloud     | Amsterdam Single Site プライベート エンドポイント
         | s3.che01.cloud-object-storage.appdomain.cloud             | Chennai Single Site エンドポイント
         | s3.private.che01.cloud-object-storage.appdomain.cloud     | Chennai Single Site プライベート エンドポイント
         | s3.mel01.cloud-object-storage.appdomain.cloud             | Melbourne Single Site エンドポイント
         | s3.private.mel01.cloud-object-storage.appdomain.cloud     | Melbourne Single Site プライベート エンドポイント
         | s3.osl01.cloud-object-storage.appdomain.cloud             | Oslo Single Site エンドポイント
         | s3.private.osl01.cloud-object-storage.appdomain.cloud     | Oslo Single Site プライベート エンドポイント
         | s3.tor01.cloud-object-storage.appdomain.cloud             | Toronto Single Site エンドポイント
         | s3.private.tor01.cloud-object-storage.appdomain.cloud     | Toronto Single Site プライベート エンドポイント
         | s3.seo01.cloud-object-storage.appdomain.cloud             | Seoul Single Site エンドポイント
         | s3.private.seo01.cloud-object-storage.appdomain.cloud     | Seoul Single Site プライベート エンドポイント
         | s3.mon01.cloud-object-storage.appdomain.cloud             | Montreal Single Site エンドポイント
         | s3.private.mon01.cloud-object-storage.appdomain.cloud     | Montreal Single Site プライベート エンドポイント
         | s3.mex01.cloud-object-storage.appdomain.cloud             | Mexico Single Site エンドポイント
         | s3.private.mex01.cloud-object-storage.appdomain.cloud     | Mexico Single Site プライベート エンドポイント
         | s3.sjc04.cloud-object-storage.appdomain.cloud             | San Jose Single Site エンドポイント
         | s3.private.sjc04.cloud-object-storage.appdomain.cloud     | San Jose Single Site プライベート エンドポイント
         | s3.mil01.cloud-object-storage.appdomain.cloud             | Milan Single Site エンドポイント
         | s3.private.mil01.cloud-object-storage.appdomain.cloud     | Milan Single Site プライベート エンドポイント
         | s3.hkg02.cloud-object-storage.appdomain.cloud             | Hong Kong Single Site エンドポイント
         | s3.private.hkg02.cloud-object-storage.appdomain.cloud     | Hong Kong Single Site プライベート エンドポイント
         | s3.par01.cloud-object-storage.appdomain.cloud             | Paris Single Site エンドポイント
         | s3.private.par01.cloud-object-storage.appdomain.cloud     | Paris Single Site プライベート エンドポイント
         | s3.sng01.cloud-object-storage.appdomain.cloud             | Singapore Single Site エンドポイント
         | s3.private.sng01.cloud-object-storage.appdomain.cloud     | Singapore Single Site プライベート エンドポイント

   --location-constraint
      バケットの場所の制約 - IBM Cloud Publicを使用する場合はエンドポイントと一致する必要があります。
      
      On-prem COSの場合は、このリストから選択しないでください。Enterキーを押します。

      例:
         | us-standard       | US Cross Region Standard
         | us-vault          | US Cross Region Vault
         | us-cold           | US Cross Region Cold
         | us-flex           | US Cross Region Flex
         | us-east-standard  | US East Region Standard
         | us-east-vault     | US East Region Vault
         | us-east-cold      | US East Region Cold
         | us-east-flex      | US East Region Flex
         | us-south-standard | US South Region Standard
         | us-south-vault    | US South Region Vault
         | us-south-cold     | US South Region Cold
         | us-south-flex     | US South Region Flex
         | eu-standard       | EU Cross Region Standard
         | eu-vault          | EU Cross Region Vault
         | eu-cold           | EU Cross Region Cold
         | eu-flex           | EU Cross Region Flex
         | eu-gb-standard    | Great Britain Standard
         | eu-gb-vault       | Great Britain Vault
         | eu-gb-cold        | Great Britain Cold
         | eu-gb-flex        | Great Britain Flex
         | ap-standard       | APAC Standard
         | ap-vault          | APAC Vault
         | ap-cold           | APAC Cold
         | ap-flex           | APAC Flex
         | mel01-standard    | Melbourne Standard
         | mel01-vault       | Melbourne Vault
         | mel01-cold        | Melbourne Cold
         | mel01-flex        | Melbourne Flex
         | tor01-standard    | Toronto Standard
         | tor01-vault       | Toronto Vault
         | tor01-cold        | Toronto Cold
         | tor01-flex        | Toronto Flex

   --acl
      バケットとオブジェクトの作成時に使用するCanned ACL。
      
      このACLはオブジェクトの作成に使用され、bucket_aclが設定されていない場合はバケットの作成にも使用されます。
      
      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      S3では、サーバーサイドでオブジェクトをコピーする場合にACLは適用されません。
      S3は、コピー元からACLをコピーせず、新しいACLを書き込むためです。
      
      aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

      例:
         | private            | オーナーがFULL_CONTROLを持ちます。
         |                    | 他のユーザーにアクセス権限はありません（デフォルト）。
         |                    | このACLはIBM Cloud（Infra）、IBM Cloud（Storage）、On-Premise COSで使用できます。
         | public-read        | オーナーがFULL_CONTROLを持ちます。
         |                    | AllUsersグループにREADアクセスがあります。
         |                    | このACLはIBM Cloud（Infra）、IBM Cloud（Storage）、On-Premise IBM COSで利用できます。
         | public-read-write  | オーナーがFULL_CONTROLを持ちます。
         |                    | AllUsersグループにREADおよびWRITEアクセスがあります。
         |                    | このACLはIBM Cloud（Infra）、On-Premise IBM COSで利用できます。
         | authenticated-read | オーナーがFULL_CONTROLを持ちます。
         |                    | AuthenticatedUsersグループにREADアクセスがあります。
         |                    | Bucketsでサポートされていません。
         |                    | このACLはIBM Cloud（Infra）およびOn-Premise IBM COSで利用できます。

   --bucket-acl
      バケットの作成時に使用するCanned ACL。
      
      詳細は[こちら](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)を参照してください。
      
      このACLはバケット作成時にのみ適用されます。このオプションが設定されていない場合は、代わりに「acl」が使用されます。
      
      "acl"と"bucket_acl"が空の文字列の場合、X-Amz-Acl: ヘッダーは追加されず、デフォルト（プライベート）が使用されます。
      

      例:
         | private            | オーナーがFULL_CONTROLを持ちます。
         |                    | 他のユーザーにアクセス権限はありません（デフォルト）。
         | public-read        | オーナーがFULL_CONTROLを持ちます。
         |                    | AllUsersグループにREADアクセスがあります。
         | public-read-write  | オーナーがFULL_CONTROLを持ちます。
         |                    | AllUsersグループにREADおよびWRITEアクセスがあります。
         |                    | バケットでのこれは一般的に推奨されません。
         | authenticated-read | オーナーがFULL_CONTROLを持ちます。
         |                    | AuthenticatedUsersグループにREADアクセスがあります。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ値。
      
      この値より大きいファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのアップロード、
      「rclone mount」またはGoogleフォト、Googleドキュメントからアップロードされたなど）は、このチャンクサイズを使用して
      マルチパートのアップロードとしてアップロードされます。
      
      "--s3-upload-concurrency" 個のこのサイズのチャンクがトランスファごとにメモリにバッファリングされます。
      
      高速なリンクで大きなファイルを転送しており十分なメモリがある場合は、これを増やすと転送が速くなります。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードすると、10,000のチャンク制限を下回るように
      チャンクサイズを自動的に増やします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのchunk_sizeは5 MiBであり、最大10,000のチャンクがあるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。
      
      チャンクサイズを増やすと、プログレス統計情報の精度が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされたときに
      そのチャンクを送信したと見なし、まだアップロードされている可能性があると見なします。より大きなチャンクサイズは、
      AWS SDKのバッファおよび進行状況の報告の精度を増加させます。

   --max-upload-parts
      マルチパートアップロードの最大パート数。
      
      このオプションは、マルチパートアップロード時に使用するパートの最大数を定義します。
      
      これは、サービスが10,000チャンクのAWS S3仕様をサポートしていない場合に使用できます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードすると、このパート数の制限を下回るように
      チャンクサイズを自動的に増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドでコピーする必要があるこのカットオフ値より大きいファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算してからアップロードするため、オブジェクトのメタデータに追加します。
      これはデータの整合性チェックには適していますが、大きなファイルのアップロードが開始するまでには長い待ち時間が発生する場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」の環境変数を探します。
      環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux / OSX："$ HOME / .aws / credentials"
          Windows：   "% USERPROFILE% \.aws \credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用するプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」または「default」が設定されていない場合はデフォルトになります。
      

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行数。
      
      同じファイルの複数のチャンクを同時にアップロードする数です。
      
      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合、これを増やすと転送が速くなる場合があります。

   --force-path-style
      真の場合、パススタイルアクセスを使用します。偽の場合、仮想ホストスタイルを使用します。
      
      この値がtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用します。
      falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3ドキュメント](https://docs.aws.amazon.com/cli/latest/reference/s3/UsingBucket.html#access-bucket-intro)を参照してください。
      
      一部のプロバイダー（AWS、Aliyun OSS、Netease COS、Tencent COSなど）では、これをfalseに設定する必要があります。
      rcloneは、プロバイダーの設定に基づいてこれを自動的に行います。

   --v2-auth
      真の場合、v2認証を使用します。
      
      これが偽（デフォルト）に設定されている場合、rcloneはv4認証を使用します。
      設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャを使用できない場合にのみ使用してください。例: Jewel/v10 CEPH以前。

   --list-chunk
      リストのチャンクのサイズ（各ListObject S3リクエストの応答リスト）。
      
      このオプションはAWS S3仕様の「MaxKeys」、「max-items」、「page-size」としても知られています。
      ほとんどのサービスは、要求された数よりも多い場合でも応答リストを1000個に切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更することはできません。詳細は[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1,2、または0は自動です。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためにListObjects呼び出しが提供されていました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これははるかに高いパフォーマンスを持ち、可能な限り使用する必要があります。
      
      デフォルトの0に設定されている場合、rcloneはプロバイダー設定に従ってListObjectsメソッドの呼び出し方法を推測します。
      推測が間違っている場合は、手動で設定できます。
      

   --list-url-encode
      リストのURLエンコードの有効化/無効化: true/false/unset
      
      一部のプロバイダは、リストをURLエンコードすることをサポートしており、利用可能な場合はファイル名に制御文字を使用する際により信頼性が高くなります。これがunsetに設定されている場合（デフォルトの場合）、rcloneは、プロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きすることができます。
      

   --no-check-bucket
      バケットの存在をチェックせず、作成しようともしません。
      
      バケットがすでに存在する場合、トランザクションの数を最小限にするために、このオプションが有効になることがあります。
      
      バケット作成の権限がない場合に必要になる場合もあります。バージョン1.52.0以前では、これはエラーなしに通過していましたが、バグのためです。
      

   --no-head
      アップロードされたオブジェクトの整合性を確認するためにHEADを行いません。
      
      rcloneは通常、PUTでオブジェクトをアップロードした後にHEADリクエストを行って整合性を確認します。
      これは大きなファイルの開始まで長い遅延を引き起こすためです。

      このフラグを設定すると、rcloneはPUTでオブジェクトをアップロードした後に200 OKメッセージを受信すると、正しくアップロードされたと見なします。
      
      特に以下を前提とします。
      
      - アップロードされたときのメタデータ、モディファイ時間、ストレージクラス、コンテンツタイプがアップロード時と同じであること
      - サイズがアップロード時と同じであること
      
      それは以下を読み込みます単一部品PUTレスポンスからの項目：
      
      - MD5SUMの値
      - アップロードされた日付
      
      マルチパートアップロードの場合、これらの項目は読み取られません。
      
      サイズが不明なソースオブジェクトがアップロードされる場合、rcloneはHEADリクエストを行います。
      
      このフラグを設定すると、アップロードのエラーを検出する確率が増加します。特にサイズが正しくないという問題が起こる可能性があるため、通常の操作ではお勧めしません。実際、このフラグを設定しても、アップロードエラーが検出される可能性は非常に低いです。

   --no-head-object
      GET前にHEADを行わない場合はtrueを設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      メモリバッファプールをどのくらいの頻度でフラッシュするか。
      
      追加バッファが必要なアップロード（マルチパートなど）はメモリプールを使用して割り当てられます。
      このオプションは、未使用のバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドのhttp2の使用を無効にします。
      
      s3（特にminio）バックエンドとHTTP/2に関する未解決の問題が現在あります。
      HTTP/2はデフォルトでs3バックエンドで有効になっていますが、ここでは無効にできます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      これは通常、AWS S3によるデータのダウンロードにおいて、AWS S3を介したデータのエグレスが安価になるため、CloudFront CDN URLに設定されます。

   --use-multipart-etag
      マルチパートアップロードでETagを使用して検証するかどうか
      
      これはtrue、false、またはプロバイダのデフォルトを使用するために設定されます。
      

   --use-presigned-request
      シングルパートアップロードのために署名済みリクエストまたはPutObjectを使用するかどうか
      
      もしfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rclone < 1.59のバージョンでは、シングルパートオブジェクトのアップロードに対して署名済みリクエストを使用し、このフラグをtrueに設定するとその機能が再度有効になります。これは例外的な状況やテスト以外では必要ありません。

   --versions
      ディレクトリリストに古いバージョンを含めます。

   --version-at
      指定された時間の当時のファイルバージョンを表示します。
      
      パラメータには、日付（「2006-01-02」）、日時（「2006-01-02 15:04:05」）、またはその時間前の期間（例：「100d」または「1h」）を指定します。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されません。つまり、ファイルのアップロードや削除はできません。
      
      有効な形式については、[timeオプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      必要に応じてgzipでエンコードされたオブジェクトを解凍します。
      
      "Content-Encoding: gzip"が設定されている状態でオブジェクトをS3にアップロードすることができます。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを受信時に「Content-Encoding: gzip」で解凍します。これにより、rcloneはサイズとハッシュを確認できませんが、ファイルの内容が解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzipで圧縮する場合に設定します。
      
      通常、プロバイダはオブジェクトをダウンロードする際に変更しません。`Content-Encoding: gzip`でアップロードされていない場合、ダウンロード時にも設定されません。
      
      ただし、いくつかのプロバイダ（例：Cloudflare）は、オブジェクトを`Content-Encoding: gzip`で圧縮する場合があります。
      
      これによる症状は以下のようなエラーです。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定してrcloneが`Content-Encoding: gzip`が設定されたオブジェクトをダウンロードし、チャンクされた転送エンコーディングを受信する場合、rcloneはオブジェクトをフライで解凍します。
      
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここではrcloneの選択を上書きすることができます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制する


オプション：
   --access-key-id value        AWSのアクセスキーID。[$ACCESS_KEY_ID]
   --acl value                  バケットとオブジェクトの作成時に使用するCanned ACL。[$ACL]
   --endpoint value             IBM COS S3 APIのエンドポイント。[$ENDPOINT]
   --env-auth                   実行時にAWSの認証情報を取得します（環境変数または環境変数がない場合のEC2/ECSメタデータ）。 (default: false) [$ENV_AUTH]
   --help, -h                   ヘルプを表示します
   --location-constraint value  バケットの場所の制約 - IBM Cloud Publicを使用する場合はエンドポイントと一致する必要があります。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。[$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               バケットの作成時に使用するCanned ACL。[$BUCKET_ACL]
   --chunk-size value               アップロードに使用するチャンクサイズ。 (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値。 (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     必要に応じてgzipでエンコードされたオブジェクトを解凍します。 (default: false) [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。 (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドのhttp2の使用を無効にします。 (default: false) [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。 (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               真の場合、パススタイルアクセスを使用します。偽の場合、仮想ホストスタイルを使用します。 (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクのサイズ（各ListObject S3リクエストの応答リスト）。 (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードの有効化/無効化：true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1,2、または0は自動です。 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パート数。 (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールのフラッシュ間隔。 (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。 (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzipで圧縮する場合に設定します。 (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しようともしません。 (default: false) [$NO_CHECK_BUCKET]
   --no-head                        HEADを行わない場合はtrueを設定します。 (default: false) [$NO_HEAD]
   --no-head-object                 GET前にHEADを行わない場合はtrueを設定します。 (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行数。 (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ値。 (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       マルチパートアップロードでETagを使用して検証するかどうか (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードのために署名済みリクエストまたはPutObjectを使用するかどうか (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        真の場合、v2認証を使用します。 (default: false) [$V2_AUTH]
   --version-at value               指定された時間の当時のファイルバージョンを表示します。 (default: "off") [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めます。 (default: false) [$VERSIONS]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}