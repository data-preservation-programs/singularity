# OCIユーザーとAPIキーを使用して認証します。
OCIテナンシーのOCID、ユーザーのOCID、リージョン、APIキーへのパス、フィンガープリントを構成ファイルに入力する必要があります。
https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos user_principal_auth - OCIユーザーとAPIキーを使用して認証します。
                                                        OCIテナンシーのOCID、ユーザーのOCID、リージョン、APIキーへのパス、フィンガープリントを構成ファイルに入力する必要があります。
                                                        https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

USAGE:
   singularity storage create oos user_principal_auth [command options] [arguments...]

DESCRIPTION:
   --namespace
      オブジェクトストレージの名前空間

   --compartment
      オブジェクトストレージのコンパートメントOCID

   --region
      オブジェクトストレージのリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。
      
      リージョンのデフォルトエンドポイントを使用する場合は、空のままにします。

   --config-file
      OCIの構成ファイルへのパス

      例:
         | ~/.oci/config | ociの設定ファイルの場所

   --config-profile
      oci構成ファイル内のプロファイル名

      例:
         | Default | デフォルトのプロファイルを使用する

   --storage-tier
      オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | 標準のストレージクラス。これがデフォルトのクラスです
         | InfrequentAccess | 低頻度アクセスのストレージクラス
         | Archive          | アーカイブのストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ。
      
      これを超えるサイズのファイルは、chunk_sizeごとにチャンクアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffよりも大きなサイズのファイルやサイズのわからない
      ファイル（たとえば「rclone rcat」からのファイルや「rclone mount」でアップロードされたファイル、Google
      フォトやGoogleドキュメント）は、このチャンクサイズを使用してマルチパートアップロードでアップロードされます。
      
      ファイルごとに「upload_concurrency」のチャンクがメモリにバッファされます。
      
      高速リンクで大きなファイルを転送し、十分なメモリがある場合は、
      これを大きくすると転送速度が向上します。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードするときに
      チャンクサイズを自動的に増やして、10,000個のチャンクの制限を下回るようにします。
      
      サイズがわからないファイルは、設定されている
      チャンクサイズでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大で
      48 GiBまでのストリームアップロードのファイルサイズをデフォルトで処理できます。
      
      より大きなファイルをストリームアップロードしたい場合は、チャンクサイズを増やす必要があります。
      
      チャンクサイズを大きくすると、進行状況の
      統計情報が"-P"フラグで表示されるときの精度が低下します。
      

   --upload-concurrency
      マルチパートアップロードの並列実行数。
      
      同じファイルのチャンクを同時にアップロードします。
      
      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に活用していない場合は、
      これを増やすと転送速度が向上する可能性があります。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。
      
      サーバーサイドでコピーする必要があるこのサイズを超えるファイルは、
      このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5 GiBです。

   --copy-timeout
      コピーのタイムアウト。
      
      コピーは非同期操作です。コピーの成功を待つためのタイムアウトを指定してください。
      

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを格納しないでください。
      
      通常、rcloneはアップロード前に入力のMD5ハッシュ値を計算し、オブジェクトのメタデータに追加するため、
      大きなファイルのアップロード開始までに長時間待つ可能性があります。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      失敗時にアップロードを中止せず、すべての正常にアップロードされたパートをS3に残します。
      
      異なるセッション間でアップロードを再開する場合には、trueに設定する必要があります。
      
      警告: 不完全なマルチパートアップロードのパーツを保存すると、オブジェクトストレージのスペース使用量と追加のコストが増えます。
      

   --no-check-bucket
      バケットの存在を確認せず、作成しようとしないでください。
      
      これは、rcloneがトランザクションの数を最小限に抑えようとする場合に便利です
      バケットが既に存在することを知っている場合です。
      
      使用しているユーザーにバケットの作成権限がない場合にも必要です。
      

   --sse-customer-key-file
      SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーの
      Base64エンコードされた文字列を含むファイル。 sse_customer_key_file | sse_customer_key | sse_kms_key_id のいずれかだけが必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-Cを使用する場合、データを暗号化または復号化するために使用する
      Base64エンコードされた256ビット暗号化キーを指定するオプションのヘッダです。 sse_customer_key_file | sse_customer_key | sse_kms_key_id のいずれかだけが必要です。詳細については、[サーバーサイド暗号化の設定](https://docs.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-Cを使用する場合、暗号化キーのBase64エンコードされたSHA256ハッシュを指定するオプションのヘッダです。
      この値は、暗号化キーの整合性を確認するために使用されます。[サーバーサイド暗号化の設定](https://docs.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-kms-key-id
      マスターキーを使用する場合、このヘッダは、Key Managementサービスを呼び出してデータ暗号化キーを生成するか、
      データ暗号化キーを暗号化または復号化するために使用されるマスター暗号化キーのOCIDを指定します。
      sse_customer_key_file | sse_customer_key | sse_kms_key_id のいずれかだけが必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
      SSE-Cを使用する場合、オプションのヘッダである"AES256"を暗号化アルゴリズムとして指定します。
      オブジェクトストレージでは "AES256"を暗号化アルゴリズムとしてサポートしています。
      詳細については、[サーバーサイド暗号化の設定](https://docs.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --compartment value     オブジェクトストレージのコンパートメントOCID [$COMPARTMENT]
   --config-file value     OCIの構成ファイルへのパス（デフォルト： "~/.oci/config"） [$CONFIG_FILE]
   --config-profile value  oci構成ファイル内のプロファイル名（デフォルト： "Default"） [$CONFIG_PROFILE]
   --endpoint value        オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
   --help, -h              ヘルプを表示します
   --namespace value       オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value          オブジェクトストレージのリージョン [$REGION]

   Advanced

   --chunk-size value               アップロードに使用するチャンクサイズ（デフォルト："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ（デフォルト："4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト（デフォルト："1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを格納しないでください（デフォルト：false） [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング（デフォルト："Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           失敗時にアップロードを中止せず、すべての正常にアップロードされたパートをS3に残します（デフォルト：false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                バケットの存在を確認せず、作成しようとしないでください（デフォルト：false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-Cを使用する場合、暗号化アルゴリズムとして" AES256 "を指定するオプションのヘッダです。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データを暗号化または復号化するために使用する256ビットのBase64エンコードされた暗号化キーです。 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのBase64エンコードされた文字列を含むファイルです。 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのBase64エンコードされたSHA256ハッシュを指定するオプションのヘッダです。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           マスターキーを使用する場合、このヘッダは、Key Managementサービスを呼び出してデータ暗号化キーを生成するか、データ暗号化キーを暗号化または復号化するために使用されるマスター暗号化キーのOCIDを指定します。 [$SSE_KMS_KEY_ID]
   --storage-tier value             オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト："Standard"） [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの並列実行数（デフォルト：10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ（デフォルト："200Mi"） [$UPLOAD_CUTOFF]

   General

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}