# インスタンスプリンシパルを使用して、APIコールを行うためにインスタンスに認証を許可します。
各インスタンスには独自のアイデンティティがあり、インスタンスメタデータから読み取った証明書を使用して認証します。
https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos instance_principal_auth - インスタンスプリンシパルを使用して、APIコールを行うためにインスタンスに認証を許可します。
                                                            各インスタンスには独自のアイデンティティがあり、インスタンスメタデータから読み取った証明書を使用して認証します。
                                                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

使用法:
   singularity storage create oos instance_principal_auth [command options] [arguments...]

説明:
   --namespace
      オブジェクトストレージの名前空間

   --compartment
      オブジェクトストレージのコンパートメントOCID

   --region
      オブジェクトストレージのリージョン

   --endpoint
      オブジェクトストレージAPIのエンドポイント。
      
      リージョンのデフォルトエンドポイントを使用する場合は空にしてください。

   --storage-tier
      新しいオブジェクトのストレージに使用するストレージクラス。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      例:
         | Standard         | Standardストレージクラス、これがデフォルトのクラスです
         | InfrequentAccess | InfrequentAccessストレージクラス
         | Archive          | Archiveストレージクラス

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフ値。
      
      この値より大きいファイルは、chunk_sizeのサイズでチャンクアップロードされます。
      最小は0、最大は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズ。
      
      upload_cutoffより大きいファイルやサイズ不明のファイル（例：「rclone rcat」でのアップロードや「rclone mount」またはGoogle
      フォトまたはGoogleドキュメントでアップロードされたファイル）は、このチャンクサイズを使用してマルチパートアップロードされます。
      
      注意：「upload_concurrency」個のこのサイズのチャンクが転送ごとにメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送しており、十分なメモリがある場合は、これを増やすことで転送速度が向上します。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、チャンクサイズを自動的に増やして10,000チャンクの制限の下に留まります。
      
      サイズ不明のファイルは設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5MiBで最大10,000のチャンクが可能ですので、デフォルトではストリームアップロードできるファイルの最大サイズは48GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      chunkサイズを増やすと、「-P」フラグで表示される進行状況の統計の精度が低下します。
      

   --upload-concurrency
      マルチパートアップロードの同時実行数。
      
      同じファイルのチャンクの数です。
      
      高速リンクで少数の大きなファイルをアップロードしており、これらのアップロードが帯域幅を完全に利用していない場合は、これを増やすと転送速度が向上する場合があります。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドでコピーする必要があるこの値より大きいファイルは、このサイズのチャンクにコピーされます。
      
      最小は0、最大は5 GiBです。

   --copy-timeout
      コピーのタイムアウト時間。
      
      コピーは非同期操作です。コピーが成功するまでの待機時間を指定します。
      

   --disable-checksum
      オブジェクトメタデータと一緒にMD5チェックサムを保存しません。
      
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算してオブジェクトのメタデータに追加します。これはデータの整合性チェックには便利ですが、大きなファイルのアップロードの開始までに長い遅延が発生する場合があります。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --leave-parts-on-error
      trueの場合、失敗時にアップロードの中断を呼び出さず、S3上のすべての正常にアップロードされたパーツを残します。
      
      異なるセッション間でアップロードの再開を行う場合にtrueに設定する必要があります。
      
      警告：不完全なマルチパートアップロードのパーツを保持することは、オブジェクトストレージ上のスペース使用量に加算され、クリーンアップされない場合に追加のコストが発生します。
      

   --no-check-bucket
      セットされた場合、バケットの存在を確認したり作成しないようにします。
      
      バケットが既に存在することを知っている場合、rcloneのトランザクション数を最小化しようとする場合に便利です。
      
      また、使用しているユーザーにバケットの作成権限がない場合にも必要です。
      

   --sse-customer-key-file
      SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコード文字列を含むファイル。 sse_customer_key_file|sse_customer_key|sse_kms_key_idのどれか1つだけが必要です。

      例:
         | <unset> | None

   --sse-customer-key
      SSE-Cを使用する場合、データの暗号化または復号化に使用するオプションのヘッダーで、base64エンコードされた256ビットの暗号化キーを指定します。s
      sse_customer_key_file|sse_customer_key|sse_kms_key_idのどれか1つだけが必要です。詳細については、サーバーサイド暗号化に独自のキーを使用する
      （https://doi.org/10.5281/zenodo.2521257）を参照してください。

      例:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションのヘッダーです。
      この値は暗号化キーの整合性をチェックするために使用されます。詳細については、サーバーサイド暗号化に独自のキーを使用する
      （https://doi.org/10.5281/zenodo.2521257）を参照してください。

      例:
         | <unset> | None

   --sse-kms-key-id
      ご使用の保管庫で独自のマスターキーを使用する場合、このヘッダーはKey Managementサービスを呼び出してデータ暗号化キーを生成するか、
      データ暗号化キーを暗号化または復号化するために使用するマスターエンクリプションキーのOCID（https://doi.org/10.5281/zenodo.2521257）を
      指定します。sse_customer_key_file|sse_customer_key|sse_kms_key_idのどれか1つだけが必要です。

      例:
         | <unset> | None

   --sse-customer-algorithm
      SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」と指定するオプションのヘッダーです。
      オブジェクトストレージは「AES256」を暗号化アルゴリズムとしてサポートしています。詳細については、サーバーサイド暗号化に独自のキーを使用する
      （https://doi.org/10.5281/zenodo.2521257）を参照してください。

      例:
         | <unset> | None
         | AES256  | AES256


オプション:
   --compartment value  オブジェクトストレージのコンパートメントOCID [$COMPARTMENT]
   --endpoint value     オブジェクトストレージAPIのエンドポイント [$ENDPOINT]
   --help, -h           ヘルプを表示
   --namespace value    オブジェクトストレージの名前空間 [$NAMESPACE]
   --region value       オブジェクトストレージのリージョン [$REGION]

   アドバンス

   --chunk-size value               アップロードに使用するチャンクサイズ（デフォルト："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値（デフォルト："4.656Gi"） [$COPY_CUTOFF]
   --copy-timeout value             コピーのタイムアウト時間（デフォルト："1m0s"） [$COPY_TIMEOUT]
   --disable-checksum               オブジェクトメタデータと一緒にMD5チェックサムを保存しない（デフォルト：false） [$DISABLE_CHECKSUM]
   --encoding value                 バックエンドのエンコーディング（デフォルト："Slash,InvalidUtf8,Dot"） [$ENCODING]
   --leave-parts-on-error           trueの場合、失敗時にアップロードの中断を呼び出さず、S3上のすべての正常にアップロードされたパーツを残します（デフォルト：false） [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                セットされた場合、バケットの存在を確認したり作成しない（デフォルト：false） [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-Cを使用する場合、暗号化アルゴリズムとして「AES256」と指定するオプションのヘッダーです。 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-Cを使用する場合、データの暗号化または復号化に使用するオプションのヘッダーで、base64エンコードされた256ビットの暗号化キー [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-Cを使用する場合、オブジェクトに関連付けられたAES-256暗号化キーのbase64エンコード文字列を含むファイル [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-Cを使用する場合、暗号化キーのbase64エンコードされたSHA256ハッシュを指定するオプションのヘッダーです。 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           ご使用の保管庫で独自のマスターキーを使用する場合、このヘッダーはOCIDを指定します。 [$SSE_KMS_KEY_ID]
   --storage-tier value             新しいオブジェクトのストレージに使用するストレージクラス。 https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm（デフォルト："Standard"） [$STORAGE_TIER]
   --upload-concurrency value       マルチパートアップロードの同時実行数（デフォルト：10） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフ値（デフォルト："200Mi"） [$UPLOAD_CUTOFF]

   一般

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}