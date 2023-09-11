# リソースプリンシパルを使用してAPI呼び出しを行う

{% code fullWidth="true" %}
```
名前:
  singularity storage update oos resource_principal_auth - リソースプリンシパルを使用してAPI呼び出しを行う

使用方法:
  singularity storage update oos resource_principal_auth [コマンドオプション] <名称|ID>

説明:
  --namespace
    オブジェクトストレージの名前空間

  --compartment
    オブジェクトストレージのコンパートメントOCID

  --region
    オブジェクトストレージのリージョン

  --endpoint
    オブジェクトストレージAPIのエンドポイント。
    
    リージョンのデフォルトエンドポイントを使用する場合は、空白のままにします。

  --storage-tier
    オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

    例:
      | Standard         | 標準のストレージクラス（デフォルト）
      | InfrequentAccess | 低頻度アクセスのストレージクラス
      | Archive          | アーカイブのストレージクラス

  --upload-cutoff
    チャンクアップロードに切り替えるためのカットオフ。
    
    これより大きいファイルは、chunk_sizeのチャンクでアップロードされます。
    最小値は0、最大値は5 GiBです。

  --chunk-size
    アップロードに使用するチャンクサイズ。
    
    upload_cutoffより大きいファイルやサイズが不明なファイル（「rclone rcat」からのアップロード、
    「rclone mount」またはGoogleフォトまたはGoogleドキュメントでアップロードされたファイルなど）は、
    このチャンクサイズを使用してマルチパートアップロードでアップロードされます。
    
    注意：転送ごとに「upload_concurrency」個のこのサイズのチャンクがメモリ上にバッファリングされます。
    
    高速リンク上で大きなファイルを転送し、十分なメモリがある場合は、これを増やすことで転送速度が向上します。
    
    Rcloneは、既知のサイズの大きなファイルをアップロードする際に自動的にチャンクサイズを増やし、
    10,000チャンクの制限を下回るようにします。
    
    サイズが不明なファイルは、設定されたchunk_sizeでアップロードされます。
    デフォルトのチャンクサイズは5 MiBであり、最大で10,000チャンクを持つことができますので、
    ファイルの最大サイズはデフォルトで48 GiBのストリームアップロードです。
    より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
    
    チャンクサイズを増やすと、「-P」フラグで表示される進行状況の精度が低下します。
    

  --upload-concurrency
    マルチパートアップロードの同時実行数。
    
    同時にアップロードされる同じファイルのチャンクの数です。
    
    高速リンク上で数個の大きなファイルをアップロードしており、これらのアップロードが帯域幅を十分に利用していない場合、
    これを増やすことで転送速度を向上させることができます。

  --copy-cutoff
    マルチパートコピーに切り替えるためのカットオフ。
    
    サーバーサイドでコピーする必要のあるこれより大きなファイルは、
    このサイズのチャンクでコピーされます。
    
    最小値は0、最大値は5 GiBです。

  --copy-timeout
    コピーのタイムアウト。
    
    コピーは非同期操作であり、コピーの成功を待つためのタイムアウトを指定します。
    

  --disable-checksum
    オブジェクトのメタデータにMD5チェックサムを格納しないでください。
    
    通常、rcloneは入力のMD5チェックサムを計算してアップロードする前に、
    オブジェクトのメタデータに追加します。これはデータの整合性チェックには便利ですが、
    大きなファイルをアップロードするときに長い遅延を引き起こす場合があります。

  --encoding
    バックエンドのエンコーディング。
    
    詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

  --leave-parts-on-error
    失敗時にアボートアップロードを呼び出さずに成功したすべてのアップロードパーツをS3に残す場合は、trueに設定してください。
    
    異なるセッション間でアップロードを再開する場合は、これをtrueに設定する必要があります。
    
    警告：未完了のマルチパートアップロードの一部を保存すると、オブジェクトストレージのスペース使用量に含まれ、
    クリーンアップされていない場合は追加のコストが発生します。
    

  --no-check-bucket
    バケットの存在をチェックせず、作成しない場合は設定してください。
    
    バケットが既に存在することを事前に知っている場合、rcloneのトランザクション数を最小限に抑えるために便利です。
    
    バケット作成権限を持たないユーザーを使用している場合にも必要になる場合があります。
    

  --sse-customer-key-file
    SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコード文字列を含むファイルです。
    sse_customer_key_file|sse_customer_key|sse_kms_key_idのいずれか1つのみが必要です。

    例:
      | <未設定> | 無し

  --sse-customer-key
    SSE-Cを使用する場合、データの暗号化または復号化に使用する、
    ベース64エンコードされた256ビット暗号化キーを指定するオプションヘッダです。
    sse_customer_key_file|sse_customer_key|sse_kms_key_idのいずれか1つのみが必要です。
    詳細については、「サーバーサイド暗号化のために固有のキーを使用する」を参照してください。
    (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

    例:
      | <未設定> | 無し

  --sse-customer-key-sha256
    SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションヘッダです。
    暗号化キーの整合性をチェックするために使用される値です。
    サーバーサイド暗号化のために固有のキーを使用する」を参照してください。
    (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

    例:
      | <未設定> | 無し

  --sse-kms-key-id
    使用しているキーマネジメントサービスで独自のマスターキーを使用する場合、
    このヘッダはデータ暗号化キーを生成するためにキーマネジメントサービスを呼び出すために使用される、
    マスター暗号化キーのOCIDを指定します。
    sse_customer_key_file|sse_customer_key|sse_kms_key_idのいずれか1つのみが必要です。

    例:
      | <未設定> | 無し

  --sse-customer-algorithm
    SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダです。
    オブジェクトストレージは、暗号化アルゴリズムとして「AES256」をサポートしています。
    詳細については、「サーバーサイド暗号化のために固有のキーを使用する」を参照してください。
    (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)。

    例:
      | <未設定> | None
      | AES256  | AES256


オプション:
  --compartment 値  オブジェクトストレージのコンパートメントOCID [$COMPARTMENT]
  --endpoint 値     オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
  --help, -h           ヘルプを表示
  --namespace 値    オブジェクトストレージの名前空間 [$NAMESPACE]
  --region 値       オブジェクトストレージのリージョン [$REGION]

  Advanced

  --chunk-size 値               アップロードに使用するチャンクサイズ（デフォルト値: "5Mi"） [$CHUNK_SIZE]
  --copy-cutoff 値              マルチパートコピーに切り替えるためのカットオフ（デフォルト値: "4.656Gi"） [$COPY_CUTOFF]
  --copy-timeout 値             コピーのタイムアウト（デフォルト値: "1m0s"） [$COPY_TIMEOUT]
  --disable-checksum               オブジェクトのメタデータにMD5チェックサムを格納しないでください（デフォルト: false） [$DISABLE_CHECKSUM]
  --encoding 値                 バックエンドのエンコーディング（デフォルト値: "Slash,InvalidUtf8,Dot"） [$ENCODING]
  --leave-parts-on-error           失敗時にアボートアップロードを呼び出さずに成功したすべてのアップロードパーツをS3に残す場合は、trueに設定してください（デフォルト: false） [$LEAVE_PARTS_ON_ERROR]
  --no-check-bucket                バケットの存在をチェックせず、作成しない場合は設定してください（デフォルト: false） [$NO_CHECK_BUCKET]
  --sse-customer-algorithm 値   SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダです。 [$SSE_CUSTOMER_ALGORITHM]
  --sse-customer-key 値         SSE-Cを使用する場合、データの暗号化または復号化に使用する、 [$SSE_CUSTOMER_KEY]
  --sse-customer-key-file 値    SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコード文字列を含むファイルです。 [$SSE_CUSTOMER_KEY_FILE]
  --sse-customer-key-sha256 値  SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションヘッダです。 [$SSE_CUSTOMER_KEY_SHA256]
  --sse-kms-key-id 値           使用しているキーマネジメントサービスで独自のマスターキーを使用する場合、 [$SSE_KMS_KEY_ID]
  --storage-tier 値             オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト値: "Standard"） [$STORAGE_TIER]
  --upload-concurrency 値       マルチパートアップロードの同時実行数（デフォルト値: 10） [$UPLOAD_CONCURRENCY]
  --upload-cutoff 値            チャンクアップロードに切り替えるためのカットオフ（デフォルト値: "200Mi"） [$UPLOAD_CUTOFF]

```
{% endcode %}