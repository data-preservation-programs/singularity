# Oracle Cloud Infrastructure オブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add oos - Oracle Cloud Infrastructure オブジェクトストレージ

USAGE:
   singularity datasource add oos [command options] <dataset_name> <source_path>

DESCRIPTION:
   --oos-chunk-size
      アップロードに使用されるチャンクのサイズ。

      アップロードの切り替え基準となるファイルサイズupload_cutoffを超える
      ファイルまたはサイズが不明なファイル（例：「rclone rcat」で作成されたり、「rclone mount」でアップロードされたりするファイル、GoogleフォトまたはGoogleドキュメント）は、
      このチャンクサイズを使用してマルチパートアップロードとしてアップロードされます。
      
      ファイルサイズが大きく、高速リンク経由で転送して十分なメモリがある場合、
      このサイズを増やして転送を高速化することができます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際にチャンクサイズを自動的に増やして、
      10,000チャンクの制限を下回るようにします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。
      デフォルトのchunk_sizeは5 MiBであり、最大で10,000チャンクまで可能です。
      そのため、デフォルトの場合、ストリームアップロードできるファイルの最大サイズは48 GiBです。
      より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況の統計情報（"-P"フラグ）の精度が低下します。
      

   --oos-compartment
      [Provider] - user_principal_auth
         オブジェクトストレージのコンパートメントOCID

   --oos-config-file
      [Provider] - user_principal_auth
         OCIコンフィグファイルへのパス

         例:
            | ~/.oci/config | ociの設定ファイルの場所

   --oos-config-profile
      [Provider] - user_principal_auth
         ociコンフィグファイル内のプロファイル名

         例:
            | Default | デフォルトのプロファイルを使用

   --oos-copy-cutoff
      マルチパートコピーへの切り替えのためのカットオフ。

      このサイズより大きなファイルに対して、サーバーサイドでコピーを行います。
      
      最小値は0、最大値は5 GiBです。

   --oos-copy-timeout
      コピーのタイムアウト。

      コピーは非同期操作です。コピーが成功するまでのタイムアウトを指定してください。
      

   --oos-disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しないようにします。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算して、
      オブジェクトのメタデータに追加するため、大きなファイルのアップロードが開始されるまで長時間待機する可能性があります。

   --oos-encoding
      バックエンドのエンコーディング。

      詳細は、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --oos-endpoint
      オブジェクトストレージAPIのエンドポイント。

      リージョンのデフォルトエンドポイントを使用するには、空白のままにします。

   --oos-leave-parts-on-error
      Trueの場合、失敗時にアップロードを中止せず、すべての正常にアップロードされたパーツをS3に残します。
      
      異なるセッション間でアップロードを再開する場合は、
      trueに設定する必要があります。
      
      注意: 不完全なマルチパートアップロードのパーツの保存は、
      オブジェクトストレージ上のスペースの使用量にカウントされ、
      クリーンアップされない場合は追加のコストが発生します。
      

   --oos-namespace
      オブジェクトストレージの名前空間

   --oos-no-check-bucket
      バケットの存在チェックや作成を試みないようにします。
      
      渡されたバケットが既に存在する場合に、rcloneのトランザクション数を
      最小限に抑えるために便利です。
      
      使用するユーザーがバケットを作成する権限がない場合にも必要です。
      

   --oos-provider
      Authプロバイダを選択します。

      例:
         | env_auth                | 実行時（環境）から認証情報を自動的に使用します。最初に認証情報が提供されたものを使用します。
         | user_principal_auth     | OCIユーザーとAPIキーを使用した認証。
                                   | ociの設定ファイルにかなりのOCID、ユーザーOCID、リージョン、パス、APIキーへのフィンガープリントを記述する必要があります。
                                   | https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
         | instance_principal_auth | インスタンスプリンシパルを使用してAPI呼び出しを許可します。
                                   | 各インスタンスには固有のIDがあり、インスタンスメタデータから読み込まれた証明書を使用して認証します。
                                   | https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
         | resource_principal_auth | リソースプリンシパルを使用してAPI呼び出しを行う
         | no_auth                 | 認証情報は不要です。これは通常、パブリックバケットの読み取りに使用されます。

   --oos-region
      オブジェクトストレージリージョン

   --oos-sse-customer-algorithm
      SSE-Cを使用する場合、暗号化アルゴリズムとしてオプションのヘッダー"AES256"を指定します。
      オブジェクトストレージは、“AES256”をデフォルトとする暗号化アルゴリズムをサポートしています。
      詳細については、[Using Your Own Keys for Server-Side Encryption](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None
         | AES256  | AES256

   --oos-sse-customer-key
      SSE-Cを使用するには、オプションのヘッダーとして使用するベース64エンコードされた256ビット暗号化キーを指定します

      例:
         | <unset> | None

   --oos-sse-customer-key-file
      SSE-Cを使用するには、オブジェクトに関連付けられたAES-256暗号化キーのベース64エンコードされた文字列をファイルに指定します

      例:
         | <unset> | None

   --oos-sse-customer-key-sha256
      SSE-Cを使用する場合、暗号化キーのベース64エンコードされたSHA256ハッシュを指定するオプションヘッダーです。
      この値は、暗号化キーの整合性を確認するために使用されます。
      詳細は、[Using Your Own Keys for Server-Side Encryption](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --oos-sse-kms-key-id
      自分自身のマスターキーを使用する場合、このヘッダーはMaster Encryption KeyのOCIDを指定します。
      データの暗号化キーの生成、暗号化、または復号化にOracle Key Managementサービスを呼び出すために使用されます。
      sse_customer_key_file|sse_customer_key|sse_kms_key_idのうち、いずれか1つだけが必要です。

      例:
         | <unset> | None

   --oos-storage-tier
      ストレージに新しいオブジェクトを格納するために使用するストレージクラス。
      [ストレージティアの理解](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)を参照してください。

      例:
         | Standard         | 標準のストレージティア、これがデフォルトです
         | InfrequentAccess | 低アクセスのストレージティア
         | Archive          | アーカイブのストレージティア

   --oos-upload-concurrency
      マルチパートアップロードの並行性。

      同じファイルのチャンクのアップロードの並行数です。
      
      高速リンク経由で少数の大きなファイルをアップロードしており、
      これらのアップロードが帯域幅を十分に使用していない場合、
      これを増やして転送を高速化できるかもしれません。

   --oos-upload-cutoff
      チャンクアップロードへの切り替えのためのカットオフ。

      このサイズより大きいファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。


OPTIONS:
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。 (デフォルト: false)
   --rescan-interval value  最後のスキャンからこの間隔が経過したら、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   --oos-chunk-size value               アップロードに使用されるチャンクのサイズです。 (デフォルト: "5Mi") [$OOS_CHUNK_SIZE]
   --oos-compartment value              オブジェクトストレージのコンパートメントOCIDです [$OOS_COMPARTMENT]
   --oos-config-file value              OCIコンフィグファイルへのパスです (デフォルト: "~/.oci/config") [$OOS_CONFIG_FILE]
   --oos-config-profile value           ociコンフィグファイル内のプロファイル名です (デフォルト: "Default") [$OOS_CONFIG_PROFILE]
   --oos-copy-cutoff value              マルチパートコピーへの切り替えのためのカットオフです (デフォルト: "4.656Gi") [$OOS_COPY_CUTOFF]
   --oos-copy-timeout value             コピーのタイムアウトです (デフォルト: "1m0s") [$OOS_COPY_TIMEOUT]
   --oos-disable-checksum value         オブジェクトのメタデータにMD5チェックサムを保存しないようにします (デフォルト: "false") [$OOS_DISABLE_CHECKSUM]
   --oos-encoding value                 バックエンドのエンコーディングです (デフォルト: "Slash,InvalidUtf8,Dot") [$OOS_ENCODING]
   --oos-endpoint value                 オブジェクトストレージAPIのエンドポイントです [$OOS_ENDPOINT]
   --oos-leave-parts-on-error value     Trueの場合、失敗時にアップロードを中止せず、すべての正常にアップロードされたパーツをS3に残します (デフォルト: "false") [$OOS_LEAVE_PARTS_ON_ERROR]
   --oos-namespace value                オブジェクトストレージの名前空間です [$OOS_NAMESPACE]
   --oos-no-check-bucket value          設定されている場合、バケットの存在を確認するか作成しません (デフォルト: "false") [$OOS_NO_CHECK_BUCKET]
   --oos-provider value                 Authプロバイダを選択します (デフォルト: "env_auth") [$OOS_PROVIDER]
   --oos-region value                   オブジェクトストレージリージョンです [$OOS_REGION]
   --oos-sse-customer-algorithm value   SSE-Cを使用する場合、オプションのヘッダーとして"AES256"を指定します [$OOS_SSE_CUSTOMER_ALGORITHM]
   --oos-sse-customer-key value         SSE-Cを使用するには、オプションのヘッダーとしてベース64でエンコードされた256ビット暗号化キーを指定します [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    SSE-Cを使用するには、オブジェクトに関連するAES-256暗号化キーのベース64エンコードされた文字列を含むファイルを指定します [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのベース64エンコードされたSHA256ハッシュを指定するオプションヘッダです [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           自分自身のマスターキーを使用する場合、OCIDを指定します [$OOS_SSE_KMS_KEY_ID]
   --oos-storage-tier value             ストレージに新しいオブジェクトを格納するために使用するストレージクラスです。[ストレージティアの理解](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)を参照してください (デフォルト: "Standard") [$OOS_STORAGE_TIER]
   --oos-upload-concurrency value       マルチパートアップロードの並行性です (デフォルト: "10") [$OOS_UPLOAD_CONCURRENCY]
   --oos-upload-cutoff value            チャンクアップロードへの切り替えのためのカットオフです (デフォルト: "200Mi") [$OOS_UPLOAD_CUTOFF]

```
{% endcode %}