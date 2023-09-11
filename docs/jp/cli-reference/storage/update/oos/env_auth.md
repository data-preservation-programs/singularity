# ランタイム（環境）から自動的に認証情報を取得し、最初に認証情報を提供したものが優先されます。

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos env_auth - ランタイム（環境）から自動的に認証情報を取得し、最初に認証情報を提供したものが優先されます

使用法:
   singularity storage update oos env_auth [command options] <name|id>

説明:
   --namespace
      オブジェクトストレージの名前空間

   --compartment
      オブジェクトストレージのコンパートメントOCID

   --region
      オブジェクトストレージのリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。
      
      リージョンのデフォルトエンドポイントを使用する場合は、空白のままにしてください。

   --storage-tier
      オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | 標準のストレージクラス（デフォルトのクラス）
         | InfrequentAccess | 低頻度アクセスストレージクラス
         | Archive          | アーカイブストレージクラス

   --upload-cutoff
      分割アップロードに切り替えるためのカットオフ値。
      
      この値より大きいファイルは、chunk_sizeのサイズで分割アップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffよりも大きいファイル、またはサイズの分からないファイル
      （例："rclone rcat"または "rclone mount"またはGoogleフォトまたはGoogleドキュメントでアップロード）は、
      このチャンクサイズを使用してマルチパートでアップロードされます。
      
      注意：1回の転送ごとに「upload_concurrency」チャンクのバッファがメモリ上に保持されます。
      
      高速リンク上で大きなファイルを転送し、メモリが十分ある場合は、
      この値を増やすことで転送速度を向上させることができます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、
      バージョン上限の10,000個のチャンクを超えないように、自動的にチャンクサイズを増やします。
      
      サイズの不明なファイルは、設定済みのチャンクサイズでアップロードされます。
      デフォルトのチャンクサイズは5 MiBであり、最大で10,000個のチャンクまで存在できます。
      つまり、デフォルト設定では、ストリームアップロード可能なファイルの最大サイズは48 GiBです。
      より大きなサイズのファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、「-P」フラグで表示される進行状況の精度が低下します。
      

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同一ファイルの複数のチャンクを同時にアップロードします。
      
      高速リンク上で小数の大きなファイルをアップロードする場合、これを増やすことで
      転送を高速化することができる場合があります。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドでコピーする必要がある、この値より大きいファイルは
      このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --copy-timeout
      コピーのタイムアウト。
      
      コピーは非同期操作ですが、成功するまでの待機時間を指定します。
      

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、
      オブジェクトのメタデータに追加します。これによりデータの整合性チェックが可能になりますが、
      大きなファイルのアップロードが開始されるまでに長い遅延が発生する可能性があります。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      trueの場合、失敗時にアップロードの中止を呼び出さず、
      S3上に正常にアップロードされたすべてのパートを手動で修復するために残します。
      
      異なるセッション間でアップロードを再開する場合は、この設定をtrueにする必要があります。
      
      注意: 不完全なマルチパートアップロードの一部を保存すると、
      オブジェクトストレージのスペース使用量にカウントされ、削除しない場合は追加の費用が発生します。
      

   --no-check-bucket
      設定すると、バケットの存在をチェックせず、作成しません。
      
      バケットが既に存在する場合に、rcloneが行うトランザクションの数を最小限に抑えるために使用できます。
      
      バケット作成の権限がない場合にも必要になる場合があります。
      

   --sse-customer-key-file
      SSE-Cを使用するための、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコードされた文字列を含む
      ファイルです。sse_customer_key_file|sse_customer_key|sse_kms_key_idのうちの1つが必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-Cを使用するための、オブジェクトの暗号化または復号化に使用する、
      base64エンコードされた256ビット暗号化キーを指定するオプションヘッダーです。
      sse_customer_key_file|sse_customer_key|sse_kms_key_idのうちの1つが必要です。
      詳細については、
      自分自身のキーを使用したサーバーサイド暗号化の使用
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定する、
      オプションのヘッダーです。この値は暗号化キーの整合性をチェックするために使用されます。
      自分自身のキーを使用したサーバーサイド暗号化の使用
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-kms-key-id
      バルトで独自のマスターキーを使用する場合、このヘッダーは
      キー管理サービスを呼び出してデータ暗号化キーを生成するか、データ暗号化キーを
      暗号化または復号化するために使用したマスター暗号キーのOCID
      (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)を指定します。
      sse_customer_key_file|sse_customer_key|sse_kms_key_idのうちの1つが必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
      SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定する、
      オプションのヘッダーです。
      オブジェクトストレージは「AES256」を暗号化アルゴリズムとしてサポートしています。
      詳細については、
      自分自身のキーを使用したサーバーサイド暗号化の使用
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None
         | AES256  | AES256


オプション:
   --compartment value  オブジェクトストレージのコンパートメントOCID [$COMPARTMENT]
   --endpoint value     オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
   --help, -h           ヘルプを表示
   --namespace value    オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value       オブジェクトストレージのリージョン [$REGION]

   高度なオプション

   --chunk-size value               アップロードに使用するチャンクサイズ（デフォルト:"5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値（デフォルト:"4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト（デフォルト:"1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しない（デフォルト:false） [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング（デフォルト:"Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           失敗時にアップロードの中止を呼び出さず、
   S3上に正常にアップロードされたすべてのパートを手動で修復するために残す場合はtrue（デフォルト:false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                設定された場合、バケットの存在をチェックせず、作成しません（デフォルト:false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションのヘッダーです。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用するための、オブジェクトの暗号化または復号化に使用する、[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-Cを使用するための、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコードされた文字列を含む [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションのヘッダーです。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           バルトで独自のマスターキーを使用する場合、このヘッダーは [$SSE_KMS_KEY_ID]
   --storage-tier value             オブジェクトストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト:"Standard"） [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時実行数（デフォルト:10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分割アップロードに切り替えるためのカットオフ値（デフォルト:"200Mi"） [$UPLOAD_CUTOFF]

```
{% endcode %}