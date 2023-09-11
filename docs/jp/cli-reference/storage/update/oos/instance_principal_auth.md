# インスタンスプリンシパルを使用して、API呼び出しを行うインスタンスを承認するために使用します。
各インスタンスは独自のアイデンティティを持ち、インスタンスメタデータから読み取られる証明書を使用して認証されます。
[ここ](https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm)を参照してください。

## 使用法
```sh
singularity storage update oos instance_principal_auth [コマンドオプション] <name|id>
```

## 説明
このコマンドでは以下のオプションが使用できます。

- `--namespace`
  オブジェクトストレージの名前空間です。

- `--compartment`
  オブジェクトストレージのコンパートメントOCIDです。

- `--region`
  オブジェクトストレージのリージョンです。

- `--endpoint`
  オブジェクトストレージAPIのエンドポイントです。
  
  リージョンのデフォルトエンドポイントを使用する場合は、空白のままでおいてください。

- `--storage-tier`
  オブジェクトストレージに新しいオブジェクトを保存するときに使用するストレージクラスです。
  [ここ](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)を参照してください。

  例:
  - `Standard` - Standardストレージクラス（デフォルト）
  - `InfrequentAccess` - InfrequentAccessストレージクラス
  - `Archive` - Archiveストレージクラス

- `--upload-cutoff`
  チャンクアップロードに切り替えるためのファイルのカットオフです。
  
  このサイズ以上のファイルは、chunk_sizeごとにチャンク分割してアップロードされます。
  最小値は0、最大値は5 GiBです。

- `--chunk-size`
  アップロードに使用するチャンクサイズです。
  
  upload_cutoffを超えるサイズのファイルや、サイズが不明なファイル（「rclone rcat」からのアップロード、
  「rclone mount」またはGoogleフォトやGoogleドキュメントからのアップロード）は、このチャンクサイズを使用して
  マルチパートアップロードされます。
  
  transferごとに「upload_concurrency」個のこのサイズのチャンクがメモリにバッファされます。
  
  高速リンクで大きなファイルを転送しており、メモリが十分にある場合は、これを増やすと転送速度が向上します。
  
  Rcloneは、既知のサイズの大きなファイルをアップロードするときに、チャンクサイズを自動的に増やして、最大で
  10,000のチャンク制限を下回るようにします。
  
  サイズが不明なファイルは、構成されたchunk_sizeでアップロードされます。
  デフォルトのチャンクサイズは5 MiBで、最大で10,000のチャンクがあります。したがって、デフォルトでは48 GiBまでの
  ファイルサイズをストリームアップロードできます。より大きなファイルをストリームアップロードする場合は、
  chunk_sizeを増やす必要があります。
  
  チャンクサイズを増やすと、「-P」フラグで表示される進行状況の統計の正確性が低下します。

- `--upload-concurrency`
  マルチパートアップロードの並行性です。
  
  同じファイルの複数のチャンクを同時にアップロードします。
  
  高速リンクで大量の大きなファイルをアップロードし、これらのアップロードで帯域幅を十分に利用しきれない場合は、
  これを増やすことで転送速度を向上させることができます。

- `--copy-cutoff`
  マルチパートコピーに切り替えるためのファイルのカットオフです。
  
  このサイズより大きいファイルをサーバーサイドでコピーする必要がある場合は、このサイズごとにコピーが行われます。
  
  最小値は0、最大値は5 GiBです。

- `--copy-timeout`
  コピーのタイムアウトです。
  
  コピーは非同期の操作です。コピーが成功するまでのタイムアウトを指定してください。

- `--disable-checksum`
  オブジェクトメタデータにMD5チェックサムを保存しないでください。
  
  通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加します。
  データの整合性チェックには役立ちますが、大きなファイルのアップロードの開始には長時間かかる場合があります。

- `--encoding`
  バックエンドのエンコード方法です。
  
  詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

- `--leave-parts-on-error`
  失敗時にアップロードの中断を回避し、すべての成功したアップロードパートを手動で
  S3上に残します。セッション間でアップロードを再開する場合には、これを`true`に設定してください。
  
  注意: 不完全なマルチパートアップロードのパーツを保持すると、オブジェクトストレージの使用スペースにカウントされ、
  クリーンアップされない場合には追加費用が発生します。

- `--no-check-bucket`
  バケットの存在を確認せず、または作成しようとしないようにする場合は、これを設定してください。

  バケットが既に存在する場合にトランザクションの数を最小限に抑える必要がある場合に便利です。
  
  使用するユーザーにバケット作成の権限がない場合にも必要になる場合があります。

- `--sse-customer-key-file`
  SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64でエンコードされた文字列を
  含むファイルです。`sse_customer_key_file|sse_customer_key|sse_kms_key_id`のいずれか一つが必要です。

  例:
  - `<unset>` - なし

- `--sse-customer-key`
  SSE-Cを使用する場合、データの暗号化または復号化に使用する、base64でエンコードされた256ビットの
  暗号化キーを指定するオプションヘッダです。`sse_customer_key_file|sse_customer_key|sse_kms_key_id`のいずれか一つが必要です。
  
  詳細については、[Using Your Own Keys for Server-Side Encryption](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

  例:
  - `<unset>` - なし

- `--sse-customer-key-sha256`
  SSE-Cを使用する場合、暗号化キーのbase64でエンコードされたSHA256ハッシュ値を指定するオプションヘッダです。
  この値は、暗号化キーの整合性を確認するために使用されます。[Using Your Own Keys for Server-Side Encryption](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

  例:
  - `<unset>` - なし

- `--sse-kms-key-id`
  ボールト内の独自のマスターキーを使用する場合、このヘッダは
  Key Managementサービスを呼び出してデータ暗号化キーを生成するか、
  データ暗号化キーの暗号化または復号化をするために使用するマスター暗号化キーのOCIDを指定します。
  `sse_customer_key_file|sse_customer_key|sse_kms_key_id`のいずれか一つが必要です。

  例:
  - `<unset>` - なし

- `--sse-customer-algorithm`
  SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダです。
  オブジェクトストレージは「AES256」を暗号化アルゴリズムとしてサポートしています。
  詳細については、[Using Your Own Keys for Server-Side Encryption](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)を参照してください。

  例:
  - `<unset>` - なし
  - `AES256` - AES256

## オプション
以下のオプションが使用できます。

- `--compartment value`
  オブジェクトストレージのコンパートメントOCIDです。 [$COMPARTMENT]

- `--endpoint value`
  オブジェクトストレージAPIのエンドポイントです。 [$ENDPOINT]

- `--help, -h`
  ヘルプを表示します。

- `--namespace value`
  オブジェクトストレージの名前空間です。 [$NAMESPACE]

- `--region value`
  オブジェクトストレージのリージョンです。 [$REGION]

- Advanced

- `--chunk-size value`
  アップロードに使用するチャンクサイズです。 (デフォルト: "5Mi") [$CHUNK_SIZE]

- `--copy-cutoff value`
  マルチパートコピーに切り替えるためのファイルのカットオフです。 (デフォルト: "4.656Gi") [$COPY_CUTOFF]

- `--copy-timeout value`
  コピーのタイムアウトです。 (デフォルト: "1m0s") [$COPY_TIMEOUT]

- `--disable-checksum`
  オブジェクトメタデータにMD5チェックサムを保存しないでください。 (デフォルト: false) [$DISABLE_CHECKSUM]

- `--encoding value`
  バックエンドのエンコーディングです。 (デフォルト: "Slash,InvalidUtf8,Dot") [$ENCODING]

- `--leave-parts-on-error`
  失敗時にアップロードの中断を回避し、すべての成功したアップロードパートを手動で
  S3上に残します。 (デフォルト: false) [$LEAVE_PARTS_ON_ERROR]

- `--no-check-bucket`
  バケットの存在を確認せず、または作成しようとしないようにします。 (デフォルト: false) [$NO_CHECK_BUCKET]

- `--sse-customer-algorithm value`
  SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダです。 [$SSE_CUSTOMER_ALGORITHM]

- `--sse-customer-key value`
  SSE-Cを使用する場合、データの暗号化または復号化に使用する、base64でエンコードされた256ビットの暗号化キーを指定します。 [$SSE_CUSTOMER_KEY]

- `--sse-customer-key-file value`
  SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64でエンコードされた文字列を含むファイルです。 [$SSE_CUSTOMER_KEY_FILE]

- `--sse-customer-key-sha256 value`
  SSE-Cを使用する場合、暗号化キーのbase64でエンコードされたSHA256ハッシュ値を指定します。 [$SSE_CUSTOMER_KEY_SHA256]

- `--sse-kms-key-id value`
  ボールト内の独自のマスターキーを使用する場合、このヘッダは
  Key Managementサービスを呼び出してデータ暗号化キーを生成するか、
  データ暗号化キーの暗号化または復号化をするために使用するマスター暗号化キーのOCIDを指定します。 [$SSE_KMS_KEY_ID]

- `--storage-tier value`
  オブジェクトストレージに新しいオブジェクトを保存するときに使用するストレージクラスです。
  [ここ](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)を参照してください。
  (デフォルト: "Standard") [$STORAGE_TIER]

- `--upload-concurrency value`
  マルチパートアップロードの並行性です。 (デフォルト: 10) [$UPLOAD_CONCURRENCY]

- `--upload-cutoff value`
  チャンクアップロードに切り替えるためのファイルのカットオフです。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]