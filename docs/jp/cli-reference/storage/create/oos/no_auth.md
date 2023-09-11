# 認証は不要です。一般的には公開されたバケットを読むために使用されます。

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos no_auth - 認証は不要です。一般的には公開されたバケットを読むために使用されます。

USAGE:
   singularity storage create oos no_auth [command options] [arguments...]

DESCRIPTION:
   --namespace
      オブジェクトストレージの名前空間

   --region
      オブジェクトストレージリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。
      
      リージョンのデフォルトエンドポイントを使用する場合は、空白のままにしておきます。

   --storage-tier
      新しいオブジェクトを保存するために使用するストレージクラス。
      https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | 標準のストレージクラス。デフォルトのクラスです。
         | InfrequentAccess | インフリークエントアクセスのストレージクラス
         | Archive          | アーカイブのストレージクラス

   --upload-cutoff
      分割アップロードに切り替えるためのカットオフ値。
      
      この値以上のファイルは、chunk_sizeの大きさで分割アップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクのサイズ。
      
      upload_cutoffより大きいファイルや、サイズ不明のファイル（例: "rclone rcat" でアップロードしたり、
      "rclone mount" や Google フォト/Google ドキュメントでアップロードしたりするファイル）は、
      このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意: "upload_concurrency" 個のこのサイズのチャンクが転送ごとにメモリ内にバッファされます。
      
      高速なリンクを介して大きなファイルを転送しており、十分なメモリを持っている場合は、
      この値を増やして転送を高速化することができます。
      
      Rclone は、既知のサイズの大きなファイルをアップロードする際に、チャンクのサイズを自動的に増やして、
      最大10,000のチャンクの制限を下回るようにします。
      
      サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。
      デフォルトのチャンクサイズは 5 MiB であり、最大で 10,000 個のチャンクがあるため、
      デフォルトではストリームアップロード可能なファイルの最大サイズは 48 GiB です。
      もしストリームアップロード可能なより大きなファイルをアップロードしたい場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを増やすことにより、プログレス統計情報の表示精度が低下します（-Pフラグ使用時）。
      

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同一ファイルの複数のチャンクが同時にアップロードされます。
      
      高速なリンクで少数の大きなファイルをアップロードしており、
      これらのアップロードが帯域幅を完全に利用していない場合は、これを増やすことで転送を高速化できる場合があります。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドでコピーする必要があるこの値より大きなファイルは、
      このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5 GiBです。

   --copy-timeout
      コピーのタイムアウト時間。
      
      コピーは非同期操作です。コピーの成功を待つためにタイムアウト時間を指定します。
      

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しません。
      
      通常、rclone はアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加します。
      これによりデータの整合性チェックができますが、大きなファイルのアップロードの開始までに長時間待つことがあります。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      失敗時にアップロードの中止を回避し、すべての正常にアップロードされたパートを手動で回収するために
      S3 に保存します（デフォルト: false）。
      
      異なるセッション間でのアップロードの再開には true に設定する必要があります。
      
      注意: 不完全なマルチパートアップロードのパートを保存すると、オブジェクトストレージでのスペース使用量にカウントされ、
      クリーンアップされない場合は追加費用が発生します。
      

   --no-check-bucket
      バケットの存在をチェックしたり、作成しようとしないでください（デフォルト: false）。
      
      バケットが既に存在することを知っている場合、rclone が実行するトランザクションの数を最小限に抑える必要がある場合に便利です。
      
      バケット作成の権限を持っていない場合も必要になる場合があります。
      

   --sse-customer-key-file
       SSE-C を使用するために、オブジェクトに関連付けられた AES-256 暗号化キーの base64 エンコードされた文字列を含む
      ファイルを指定します。
       sse_customer_key_file|sse_customer_key|sse_kms_key_id のうちのいずれか1つのみが必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-C を使用するために、データの暗号化または復号化に使用する base64 エンコードされた 256 ビットの暗号化キーを指定するオプションヘッダ。
      sse_customer_key_file|sse_customer_key|sse_kms_key_id のうちのいずれか1つのみが必要です。
      詳細については、以下を参照してください。
      [自分自身のキーを使用したサーバーサイド暗号化の使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-C を使用する場合、暗号化キーの base64 エンコードされた SHA256 ハッシュを指定するオプションヘッダ。
      この値は暗号化キーの整合性を確認するために使用されます。詳細については、以下を参照してください。
      [自分自身のキーを使用したサーバーサイド暗号化の使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      例:
         | <unset> | None

   --sse-kms-key-id
      自分自身のマスターキーを使っている場合、このヘッダは、データ暗号化キーを生成するために Key Management サービスを呼び出すために使用する
      マスター暗号化キーの OCID を指定します（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)。
      sse_customer_key_file|sse_customer_key|sse_kms_key_id のうちのいずれか1つのみが必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
       SSE-C を使用する場合、暗号化アルゴリズムとして "AES256" を指定するオプションヘッダです。
       オブジェクトストレージは "AES256" を暗号化アルゴリズムとしてサポートしています。
       詳細については、以下を参照してください。
      [自分自身のキーを使用したサーバーサイド暗号化の使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      例:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --endpoint value   オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
   --help, -h         ヘルプを表示
   --namespace value  オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value     オブジェクトストレージリージョン [$REGION]

   アドバンスド

   --chunk-size value               アップロードに使用するチャンクのサイズ (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト時間 (デフォルト: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しません (デフォルト: false) [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           失敗時にアップロードの中止を回避し、すべての正常にアップロードされたパートを手動で回収するためのS3に保存します (デフォルト: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                バケットの存在をチェックしたり、作成しようとしないでください (デフォルト: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C を使用する場合、暗号化アルゴリズムとして "AES256" を指定するオプションヘッダ [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C を使用するために、暗号化に使用するbase64エンコードされた256ビットの暗号化キーを指定します [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C を使用するために、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコードされた文字列を含むファイルを指定します [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C を使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションヘッダ [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           自分自身のマスターキーを使っている場合、キーコードを生成するために呼び出すためにKMSサービスを使用するマスター暗号化キーのOCIDを指定します [$SSE_KMS_KEY_ID]
   --storage-tier value             新しいオブジェクトを保存するために使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (デフォルト: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時実行数 (デフォルト: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分割アップロードに切り替えるためのカットオフ値 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]

   一般的

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}