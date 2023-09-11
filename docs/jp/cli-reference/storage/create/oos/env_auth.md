# ランタイム（env）から認証情報を自動的に取得します。最初に認証情報が提供されたものが優先されます。

{% code fullWidth = "true" %}
```
NAME:
   singularity storage create oos env_auth - ランタイム（env）から認証情報を自動的に取得します。最初に認証情報が提供されたものが優先されます

使用法:
   singularity storage create oos env_auth [コマンドオプション] [引数...]

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
      ストレージに新しいオブジェクトを格納する際に使用するストレージクラスです。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | Standardストレージ層、これがデフォルトの層です
         | InfrequentAccess | InfrequentAccessストレージ層
         | Archive          | Archiveストレージ層

   --upload-cutoff
      分割アップロードに切り替えるためのカットオフです。
      
      これより大きいファイルは、chunk_sizeの大きさで分割してアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクのサイズです。
      
      upload_cutoffより大きいファイル、またはサイズが不明なファイル（例："rclone rcat"からのもの、または"rclone mount"またはGoogleフォトまたはGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードを行います。
      
      注意：transferごとに「upload_concurrency」チャンクがメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送して十分なメモリがある場合は、これを増やすことで転送速度が向上します。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする際に、10,000チャンクの制限を下回るようにチャンクサイズを自動的に増やします。
      
      サイズがわからないファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBであり、最大10,000個のチャンクを有することができます。そのため、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、「-P」フラグで表示される進行状況の統計の正確性が低下します。
      

   --upload-concurrency
      マルチパートアップロードの同時実行数です。
      
      同じファイルのチャンクを同時にアップロードします。
      
      高速リンクを介して大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に活用していない場合は、これを増やすことで転送速度が向上するかもしれません。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフです。
      
      サーバーサイドコピーが必要なこのカットオフより大きなファイルは、このサイズで分割してコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --copy-timeout
      コピーのタイムアウトです。
      
      コピーは非同期操作です。コピーが成功するまでの待機時間を指定します。
      

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを格納しないでください。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまでに長時間の遅延が生じることがあります。これはデータの整合性チェックには役立ちますが、大きなファイルのアップロードには長時間かかることがあります。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディング](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      失敗時にアップロードの中止を呼び出さず、すべての正常にアップロードされたパートをS3に残します（手動で回復するため）。

      続きを行う場合はtrueに設定する必要があります。

      警告：未完了のマルチパートアップロードのパーツを保存すると、オブジェクトストレージのスペース使用量に加算され、クリーンアップされない場合は追加料金が発生します。
      

   --no-check-bucket
      バケットの存在をチェックしたり作成したりしようとしない場合は、設定します。
      
      これは、バケットが既に存在することを知っている場合に、rcloneが実行するトランザクションの数を最小限に抑える必要がある場合に便利です。
      
      バケット作成の権限を持っていないユーザーを使用している場合にも必要です。
      

   --sse-customer-key-file
      SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのBase64エンコードされた文字列を含むファイルです。
      sse_customer_key_file、sse_customer_key、sse_kms_key_idのいずれか1つだけ必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化または復号化に使用するBase64エンコードされた256ビット暗号化キーを指定するオプションヘッダです。
      sse_customer_key_file、sse_customer_key、sse_kms_key_idのいずれか1つだけ必要です。詳細については、
      サーバーサイド暗号化の独自のキーの使用
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-Cを使用している場合、暗号化キーのBase64エンコードされたSHA256ハッシュを指定するオプションヘッダです。
      この値は、暗号化キーの整合性をチェックするために使用されます。詳細については、
      サーバーサイド暗号化の独自のキーの使用
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-kms-key-id
      自分のボルトで使用している独自のマスターキーを使用する場合、このヘッダはマスターエンクリプションキーのOCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）を指定します。
      sse_customer_key_file、sse_customer_key、sse_kms_key_idのいずれか1つだけ必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
      SSE-Cを使用する場合、オプションのヘッダである暗号化アルゴリズムとして「AES256」を指定します。
      オブジェクトストレージは「AES256」を暗号化アルゴリズムとしてサポートしています。詳細については、
      サーバーサイド暗号化の独自のキーの使用
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

   Advanced

   --chunk-size value               アップロードに使用するチャンクのサイズ（デフォルト: "5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ（デフォルト: "4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト（デフォルト: "1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを格納しないでください（デフォルト: false） [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング（デフォルト: "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           失敗時にアップロードの中止を呼び出さず、すべての正常にアップロードされたパートをS3に残します（手動で回復するため）（デフォルト: false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                バケットの存在をチェックしたり作成したりしようとしない場合は、設定します（デフォルト: false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-Cを使用する場合、オプションのヘッダである暗号化アルゴリズムとして「AES256」を指定します。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データの暗号化または復号化に使用する256ビット暗号化キーのBase64エンコードされた文字列を指定します。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのBase64エンコードされた文字列を含むファイルです。[$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-Cを使用している場合、暗号化キーのBase64エンコードされたSHA256ハッシュを指定するオプションヘッダです。[$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           自分のボルトで使用している独自のマスターキーを使用する場合、このヘッダはマスターエンクリプションキーのOCID（https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm）を指定します。[$SSE_KMS_KEY_ID]
   --storage-tier value             ストレージに新しいオブジェクトを格納する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト: "Standard"） [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時実行数（デフォルト: 10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分割アップロードに切り替えるためのカットオフ（デフォルト: "200Mi"） [$UPLOAD_CUTOFF]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}