# 認証情報は不要で、通常はパブリックなバケットを読むためです。

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos no_auth - 認証情報は不要で、通常はパブリックなバケットを読むためです

USAGE:
   singularity storage update oos no_auth [command options] <名前|ID>

DESCRIPTION:
   --namespace
      オブジェクトストレージの名前空間

   --region
      オブジェクトストレージのリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。
      
      リージョンのデフォルトエンドポイントを使用する場合は空白のままにします。

   --storage-tier
      ストレージに新しいオブジェクトを保存する場合に使用するストレージクラス。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | 標準のストレージクラス、これがデフォルトのクラスです
         | InfrequentAccess | インフリークエントアクセスのストレージクラス
         | Archive          | アーカイブのストレージクラス

   --upload-cutoff
      分割アップロードに切り替えるためのカットオフ値。
      
      この値より大きいファイルは、chunk_sizeのチャンクでアップロードされます。
      最小は 0、最大は 5 GiBです。

   --chunk-size
      アップロードに使用するチャンクのサイズ。
      
      upload_cutoff より大きいファイル、またはサイズが不明なファイル（例：「rclone rcat」からのファイル、または「rclone mount」やGoogleフォトやGoogleドキュメントでアップロードされたファイル）をアップロードする場合は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      プロセスごとに「upload_concurrency」個のこのサイズのチャンクがメモリにバッファされます。
      
      高速リンクで大きなファイルを転送しており、メモリが十分にある場合は、これを増やすことで転送を高速化できます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合は、チャンクサイズを 10,000 チャンクの制限以下に保つために自動的にチャンクサイズを増やします。
      
      サイズ不明のファイルは、設定されたチャンクサイズでアップロードされます。デフォルトのチャンクサイズは 5 MiB で、最大 10,000 チャンクありますので、デフォルトではストリームアップロードできるファイルの最大サイズは 48 GiB です。より大きなファイルをストリームアップロードする場合は、chunk_size を増やす必要があります。
      
      チャンクサイズを増やすと、"-P" フラグで表示される進行状況の統計情報の正確性が低下します。
      

   --upload-concurrency
      マルチパートアップロードの同時性。
      
      同じファイルのチャンクが同時にアップロードされる数です。
      
      高速リンクで少数の大きなファイルを転送しており、これらのアップロードが帯域幅を完全に活用していない場合は、これを増やすことで転送を高速化できます。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドでコピーする必要のあるこのサイズより大きいファイルは、このサイズのチャンクでコピーされます。
      
      最小は 0、最大は 5 GiBです。

   --copy-timeout
      コピーのタイムアウト。
      
      コピーは非同期操作です。コピーが成功するまで待機するにはタイムアウトを指定します。
      

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しないでください。
      
      通常、rclone はアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまでに長い遅延が発生することがあります。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      失敗時にアップロードを中止せず、S3 上のすべての正常にアップロードされたパーツを手動でリカバリするためにメモリ上に残します。

      異なるセッション間でアップロードを再開する場合は、これを true に設定する必要があります。

      注意: 不完全なマルチパートアップロードの一部を格納することは、オブジェクトストレージのスペース使用量にカウントされ、後処理が行われない場合に追加のコストが発生します。
      

   --no-check-bucket
      バケットの存在を確認せず、作成しようとしません。
      
      バケットが既に存在することを知っている場合、rclone のトランザクション数を最小限にするためにこのオプションを使用すると便利です。
      
      また、使用しているユーザーにバケットの作成権限がない場合も必要になることがあります。
      

   --sse-customer-key-file
      SSE-C を使用する場合、オブジェクトと関連付けられた AES-256 暗号化キーの base64 エンコード済み文字列が含まれるファイル。 sse_customer_key_file|sse_customer_key|sse_kms_key_id のいずれかが必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-C を使用する場合、データを暗号化または復号化するために使用される、任意のヘッダーである base64 エンコードされた 256 ビットの暗号化キーを指定します。 sse_customer_key_file|sse_customer_key|sse_kms_key_id のいずれかが必要です。詳細については、[サーバーサイド暗号化の独自のキーの使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-C を使用する場合、暗号化キーの base64 エンコードされた SHA256 ハッシュを指定する任意のヘッダーです。この値は暗号化キーの整合性を確認するために使用されます。[サーバーサイド暗号化の独自のキーの使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None

   --sse-kms-key-id
      ボールトの独自のマスターキーを使用する場合、このヘッダーは、データ暗号化キーを生成するか、データ暗号化キーを暗号化または復号化するためにキーマネジメントサービスを呼び出すために使用されるマスター暗号化キーの OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) を指定します。 sse_customer_key_file|sse_customer_key|sse_kms_key_id のいずれかが必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
      SSE-C を使用する場合、「AES256」という暗号化アルゴリズムを指定する任意のヘッダーです。
      オブジェクトストレージは、暗号化アルゴリズムとして「AES256」をサポートしています。詳細については、[サーバーサイド暗号化の独自のキーの使用](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

      例:
         | <unset> | None
         | AES256  | AES256


OPTIONS:
   --endpoint value   オブジェクトストレージのエンドポイント。 [$ENDPOINT]
   --help, -h         ヘルプを表示
   --namespace value  オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value     オブジェクトストレージのリージョン [$REGION]

   Advanced

   --chunk-size value               アップロードに使用するチャンクのサイズ。 (デフォルト: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト。 (デフォルト: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しないでください。 (デフォルト: false) [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           失敗時にアップロードを中止せず、S3 上のすべての正常にアップロードされたパーツを手動でリカバリするためにメモリ上に残します。 (デフォルト: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                バケットの存在を確認せず、作成しようとしません。 (デフォルト: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C を使用する場合、「AES256」という暗号化アルゴリズムを指定する任意のヘッダーです。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C を使用する場合、データを暗号化または復号化するために使用される、任意のヘッダーであるbase64エンコードされた256ビットの暗号化キーを指定します。[$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C を使用する場合、オブジェクトと関連付けられたAES-256暗号化キーのbase64エンコード済み文字列が含まれるファイル。[$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C を使用する場合、暗号化キーのbase64エンコード済みSHA256ハッシュを指定する任意のヘッダーです。[$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           ボールトの独自のマスターキーを使用する場合、OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) でマスター暗号化キーを指定します。上記のいずれかが必要です。[$SSE_KMS_KEY_ID]
   --storage-tier value             ストレージに新しいオブジェクトを保存する場合に使用するストレージクラス。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm(デフォルト:"Standard") [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時性。 (デフォルト: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            分割アップロードに切り替えるためのカットオフ値。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}