# OCIユーザーとAPIキーを使用した認証を使用します。
テナンシーOCID、ユーザーOCID、リージョン、パス、APIキーのフィンガープリントを構成ファイルに入力する必要があります。
https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

{% code fullWidth="true" %}
```
名前:
   singularity storage update oos user_principal_auth - OCIユーザーとAPIキーを使用した認証を使用します。
                                                        テナンシーOCID、ユーザーOCID、リージョン、パス、APIキーのフィンガープリントを構成ファイルに入力する必要があります。
                                                        https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

使用法:
   singularity storage update oos user_principal_auth [コマンドオプション] <name|id>

概要:
   --namespace
      オブジェクトストレージの名前空間

   --compartment
      オブジェクトストレージのコンパートメントOCID

   --region
      オブジェクトストレージのリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。

      リージョンのデフォルトエンドポイントを使用する場合は空白のままにします。

   --config-file
      OCI構成ファイルのパス

      例:
         | ~/.oci/config | oci構成ファイルの場所

   --config-profile
      OCI構成ファイル内のプロファイル名

      例:
         | Default | デフォルトプロファイルを使用する

   --storage-tier
      ストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | 標準ストレージクラス、これがデフォルトのクラスです
         | InfrequentAccess | 低頻度アクセスストレージクラス
         | Archive          | アーカイブストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ。

      これより大きなファイルは、chunk_sizeでチャンク分割してアップロードされます。
      最小値は0で、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。

      upload_cutoffよりも大きいファイルやサイズのわからないファイル（「rclone rcat」からのアップロード、または「rclone mount」やGoogleフォト、Googleドキュメントからアップロードされたファイルなど）は、このチャンクサイズでマルチパートアップロードを使用してアップロードされます。

      注意：トランスファごとに「upload_concurrency」チャンクがこのサイズでメモリ内にバッファリングされます。

      高速リンクで大きなファイルを転送してメモリが十分にある場合は、この値を増やすことで転送を高速化することができます。

      Rcloneは、10,000個のチャンクの制限を下回るように、既知のサイズの大きなファイルをアップロードする場合に自動的にチャンクサイズを増やします。

      サイズのわからないファイルは、構成済みのchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大10,000チャンクまで許容されるため、デフォルトではストリームアップロード可能なファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。

      チャンクサイズを増やすと、「-P」フラグと共に表示される進行状況の精度が低下します。
      

   --upload-concurrency
      マルチパートアップロードの同時実行数。

      同じファイルのチャンクのうち、同時にアップロードされる数です。

      数が少ない大きなファイルを高速リンクでアップロードし、これらのアップロードが帯域幅を十分に利用しない場合、これを増やすことで転送を高速化することができます。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ。

      サーバーサイドでコピーする必要があるこのカットオフより大きなファイルは、このサイズのチャンクにコピーされます。

      最小値は0で、最大値は5 GiBです。

   --copy-timeout
      コピーのタイムアウト。

      コピーは非同期操作です。コピーが成功するまでの待機時間を指定します。
      

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しない。

      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、大きなファイルのアップロードを開始するまで時間がかかることがあります。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      失敗時にアボートアップロードを呼び出さず、S3上のすべての正常にアップロードされたパーツを手動で回復のために残します。

      異なるセッション間でアップロードを再開する場合、これをtrueに設定する必要があります。

      警告：不完全なマルチパートアップロードのパーツをストアすると、オブジェクトストレージ上のスペース使用量にカウントされ、クリーンアップされない場合は追加のコストがかかります。
      

   --no-check-bucket
      設定されている場合、バケットの存在をチェックせず、作成しません。

      バケットが既に存在する場合に、rcloneが実行するトランザクション数を最小限に抑える必要がある場合に有用です。

      使用するユーザーにバケットの作成権限がない場合にも必要となる場合があります。
      

   --sse-customer-key-file
      SSE-Cを使用する場合、オブジェクトに関連付けられているAES-256暗号化キーのbase64エンコード文字列を含むファイル。sse_customer_key_file|sse_customer_key|sse_kms_key_idのうち、1つだけが必要です。

      例:
         | <unset> | なし

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化または復号化に使用する、base64エンコードされた256ビット暗号化キーを指定するオプションヘッダー。sse_customer_key_file|sse_customer_key|sse_kms_key_idのうち、1つだけが必要です。詳細については、「独自の暗号化キーを使用したサーバーサイド暗号化の使用」を参照してください。
      
      例:
         | <unset> | なし

   --sse-customer-key-sha256
      SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションヘッダー。この値は暗号化キーの整合性を確認するために使用されます。詳細については、「独自の暗号化キーを使用したサーバーサイド暗号化の使用」を参照してください。

      例:
         | <unset> | なし

   --sse-kms-key-id
      独自のマスターキーを使用する場合、このヘッダーは、暗号化キーを生成するためキーマネジメントサービスを呼び出すために使用するマスター暗号化キーのOCIDを指定します。sse_customer_key_file|sse_customer_key|sse_kms_key_idのうち、1つだけが必要です。

      例:
         | <unset> | なし

   --sse-customer-algorithm
      SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダー。

      オブジェクトストレージは、「AES256」を暗号化アルゴリズムとしてサポートしています。詳細については、「独自の暗号化キーを使用したサーバーサイド暗号化の使用」を参照してください。

      例:
         | <unset> | なし
         | AES256  | AES256


オプション:
   --compartment value     オブジェクトストレージのコンパートメントOCID [$COMPARTMENT]
   --config-file value     OCI構成ファイルのパス（デフォルト: "~/.oci/config"） [$CONFIG_FILE]
   --config-profile value  OCI構成ファイル内のプロファイル名（デフォルト: "Default"） [$CONFIG_PROFILE]
   --endpoint value        オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
   --help, -h              ヘルプを表示
   --namespace value       オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value          オブジェクトストレージのリージョン [$REGION]

   高度な設定

   --chunk-size value               アップロードに使用するチャンクサイズ（デフォルト: "5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ（デフォルト: "4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト（デフォルト: "1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない（デフォルト: false） [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング（デフォルト: "Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           失敗時にアボートアップロードを呼び出さず、S3上のすべての正常にアップロードされたパーツを手動で回復のために残す（デフォルト: false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                設定されている場合、バケットの存在をチェックせず、作成しません（デフォルト: false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」を指定するオプションヘッダー [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データの暗号化または復号化に使用する、base64エンコードされた256ビット暗号化キー [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-Cを使用する場合、オブジェクトに関連付けられているAES-256暗号化キーのbase64エンコード文字列を含むファイル [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションヘッダ [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           独自のマスターキーを使用する場合、OCIDを指定するこのヘッダーは、データ暗号化キーを生成するためにキーマネジメントサービスを呼び出すために使用されます [$SSE_KMS_KEY_ID]
   --storage-tier value             ストレージに新しいオブジェクトを保存する際に使用するストレージクラス。https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト: "Standard"） [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時実行数（デフォルト: 10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ（デフォルト: "200Mi"） [$UPLOAD_CUTOFF]

```
{% endcode %}