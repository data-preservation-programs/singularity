# Dreamhost DreamObjects

```
NAME:
   singularity storage update s3 dreamhost - Dreamhost DreamObjects

使用法:
   singularity storage update s3 dreamhost [コマンドオプション] <名前|ID>

説明:
   --env-auth
      実行時にAWS認証情報（環境変数または環境変数がない場合はEC2 / ECSメタデータ）から取得する。
      
      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境（env varsまたはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたは実行時の認証情報にする場合は空白のままにしてください。

   --region
      接続するリージョン。
      
      S3クローンを使用し、リージョンがない場合は空白のままにしてください。

      例:
         | <unset>            | 確定がない場合に使用します。
         |                    | V4シグネチャと空のリージョンが使用されます。
         | other-v2-signature | V4シグネチャが機能しない場合にのみ使用します。
         |                    | 例：Jewel/v10 CEPHの前。

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用する場合は必須です。

      例:
         | objects-us-east-1.dream.io | Dream Objectsのエンドポイント

   --location-constraint
      リージョンに合わせて設定する位置制約。
      
      確定がない場合は空白のままにしてください。バケットを作成する場合にのみ使用されます。

   --acl
      バケットを作成したり、オブジェクトを保存またはコピーする際に使用するキャンドACL。
      
      このACLは、オブジェクトの作成時およびbucket_aclが設定されていない場合にも使用されます。
      
      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      なお、S3はソースからACLをコピーするのではなく、新しいACLを書き込むため、このACLが適用されます。
      
      aclが空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。

   --bucket-acl
      バケットを作成する際に使用するキャンドACL。
      
      詳細についてはhttps://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-aclを参照してください。
      
      このACLは、バケット作成時にのみ適用されます。設定されていない場合は「acl」が代わりに使用されます。
      
      「acl」と「bucket_acl」の両方が空の文字列の場合、X-Amz-Acl:ヘッダーは追加されず、デフォルト（private）が使用されます。

      例:
         | private            | オーナーにFULL_CONTROLが与えられます。
         |                    | 他の誰もアクセス権限はありません（デフォルト）。
         | public-read        | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループにREADアクセスが与えられます。
         | public-read-write  | オーナーにFULL_CONTROLが与えられます。
         |                    | AllUsersグループがREADおよびWRITEアクセスを取得します。
         |                    | バケットでこれを設定することは一般的に推奨されません。
         | authenticated-read | オーナーにFULL_CONTROLが与えられます。
         |                    | AuthenticatedUsersグループがREADアクセスを取得します。

   --upload-cutoff
      チャンク化アップロードへの切り替えのためのカットオフ値。
      
      この値よりも大きいファイルは、chunk_sizeのチャンクとしてアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロード時に使用するチャンクサイズ。
      
      upload_cutoffよりも大きいファイル、またはサイズがわからないファイル（「rclone rcat」からのアップロードや「rclone mount」やGoogleフォトやGoogleドキュメントからのアップロードなど）は、このチャンクサイズを使用してマルチパートのアップロードとなります。
      
      注意：「--s3-upload-concurrency」は、このサイズのチャンクを転送ごとにメモリ上にバッファリングします。
      
      高速リンクで大きなファイルを転送して十分なメモリがある場合は、これを増やすと転送速度が向上します。
      
      Rcloneは、既知のサイズで大きなファイルをアップロードする場合、10,000のチャンクリミットを下回るために自動的にチャンクサイズを増やします。
      
      サイズがわからないファイルは、設定されたchunk_sizeでアップロードされます。デフォルトのチャンクサイズは5 MiBで、最大で10,000のチャンクまでありますので、デフォルトではストリームアップロードできるファイルの最大サイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、chunk_sizeを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況に関する「-P」フラグに表示される進行統計情報の正確性が低下します。Rcloneは、チャンクがAWS SDKによってバッファリングされた時点でチャンクが送信されたと扱いますが、実際にはまだアップロード中かもしれません。チャンクサイズが大きくなると、AWS SDKのバッファと進行状況の報告の正確性が乖離するため、進行度統計情報の正確性が低下します。

   --max-upload-parts
      マルチパートアップロードでの最大パート数。
      
      このオプションは、マルチパートアップロードを行う際に使用する最大マルチパートチャンク数を定義します。
      
      サービスが10,000のチャンク仕様をサポートしていない場合に役立ちます。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、このチャンクサイズを増やすことで10,000チャンクの制限を下回るように自動的にチャンクサイズを増やします。

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフ値。
      
      サーバーサイドコピーが必要なこのサイズを超えるファイルは、このサイズのチャンクでコピーされます。
      
      最小値は0、最大値は5 GiBです。

   --disable-checksum
      オブジェクトメタデータにMD5チェックサムを保存しない。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加します。これにより、大きなファイルのアップロードを開始するために長時間待つことができます。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      空の場合、rcloneは「AWS_SHARED_CREDENTIALS_FILE」環境変数を探します。環境変数の値が空の場合、現在のユーザーのホームディレクトリがデフォルトになります。
      
          Linux / OSX：「$HOME / .aws / credentials」
          Windows：「%USERPROFILE% \ .aws \ credentials」

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_auth = trueの場合、rcloneは共有認証情報ファイルを使用できます。この変数はそのファイルで使用されるプロファイルを制御します。
      
      空の場合、環境変数「AWS_PROFILE」と「default」の環境変数が設定されていない場合にはデフォルトになります。

   --session-token
      AWSセッショントークン。

   --upload-concurrency
      マルチパートアップロードの並行度。
      
      同じファイルのチャンクの並行アップロード数です。
      
      ハイスピードリンクで大量の大きなファイルを転送しており、これらの転送が帯域幅を十分に利用していない場合は、これを増やすと転送速度が向上するかもしれません。

   --force-path-style
      もし真ならパス形式アクセスを使い、もし偽なら仮想ホスト形式アクセスを使う。
      
      これが真（デフォルト）なら、rcloneはパス形式アクセスを使用します。偽なら、rcloneは仮想パススタイルを使用します。詳細は[the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      を参照してください。
      
      一部のプロバイダ（AWS、Aliyun OSS、Netease COS、またはTencent COSなど）は、この値を設定する必要がありますが、rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      もし真ならv2認証を使用する。
      
      偽（デフォルト）ならrcloneはv4認証を使用します。セットされている場合、rcloneはv2認証を使用します。
      
      v4署名が機能しない場合にのみ使用してください。例：Jewel/v10 CEPHの前。

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストに対する応答リストのサイズ）。
      
      このオプションは「MaxKeys」、「max-items」または「page-size」とも呼ばれ、AWS S3の仕様からも取得できます。
      多くのサービスは、それ以上を要求してもリクエストリストを1000オブジェクトに切り詰めます。
      AWS S3では、これはグローバルな最大であり、変更することはできません。[AWS S３](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。

   --list-version
      使用するListObjectsのバージョン：1、2、または自動の場合は0。
      
      S3が最初にリリースされた当初、バケット内のオブジェクトを列挙するためにListObjects呼び出しだけが提供されました。
      
      ただし、2016年5月にはListObjectsV2呼び出しが導入されました。これははるかに高性能で、可能な限り使用する必要があります。
      
      デフォルトで設定されている0の場合、rcloneはプロバイダの設定に応じて呼び出すリストオブジェクトのメソッドを推測します。推測が間違っている場合は、ここで手動で設定できます。

   --list-url-encode
      リストのURLエンコードを行うかどうか：true/false/unset
      
      一部のプロバイダは、ファイル名に制御文字を使用する場合にURLエンコードリストをサポートしており、利用可能な場合はこれがより信頼性が高いです。これがunsetに設定されている場合（デフォルト）、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-check-bucket
      バケットの存在をチェックせず、作成しようとしません。
      
      バケットが既に存在する場合にrcloneが行うトランザクションの数を最小限にする場合に役立ちます。
      
      また、使用するユーザーにバケット作成の権限がない場合にも必要です。v1.52.0より前のバージョンでは、これはバグのために静かにパスされました。

   --no-head
      アップロードされたオブジェクトのHEADをチェックしない場合に設定します。
      
      rcloneは、PUTでオブジェクトをアップロードした後に200 OKのメッセージを受け取った場合、正しくアップロードされたと見なします。

      特に次のことを仮定します。

      - メタデータ（モディファイ時間、ストレージクラス、コンテンツタイプを含む）はアップロード時と同じであった
      - サイズはアップロードされたものと同じであった
      
      以下の項目を単一のパートPUTのレスポンスから読み取ります。

      - MD5SUM
      - アップロード日

      マルチパートアップロードの場合、これらの項目は読み取られません。

      不明な長さのソースオブジェクトがアップロードされた場合、rclone **は** HEADリクエストを実行します。

      このフラグを設定すると、アップロードの失敗が検出される可能性が増えます。特にサイズが正しくない場合ですので、通常の動作には推奨されません。実際には、このフラグを使用しても、アップロードの失敗が検出される可能性は非常に低いです。

   --no-head-object
      GETする前にHEADを実行しない場合に設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、概要の[エンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールがフラッシュされる頻度。
      
      追加のバッファを必要とするアップロード（たとえばマルチパート）は、割り当てにメモリプールを使用します。
      このオプションは、未使用のバッファがプールから取り除かれる頻度を制御します。

   --memory-pool-use-mmap
      メモリプール内のmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでのhttp2の使用を無効にします。
      
      S3（具体的にはminio）バックエンドとHTTP/2の問題が現在解決されていません。S3バックエンドではデフォルトでHTTP/2が有効になっていますが、ここで無効にすることができます。問題が解決されたら、このフラグは削除されます。
      
      参照: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してデータをダウンロードすることでegress費用が削減されます。

   --use-multipart-etag
      検証のためにマルチパートアップロードでETagを使用するかどうか
      
      これはtrue、false、またはデフォルトのために設定されていない状態である必要があります。

   --use-presigned-request
      シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン< 1.59では、単一のパートオブジェクトをアップロードするために署名済みリクエストを使用し、このフラグをtrueに設定するとこの機能を再度有効にすることができます。これは特殊な事情やテスト以外では必要ありません。

   --versions
      ディレクトリリスティングに古いバージョンを含める。

   --version-at
      指定した時間の存在した時点のファイルバージョンを表示します。
      
      パラメータは日付、「2006-01-02」、datetime「2006-01-02 15:04:05」、またはそのように長い間前の期間、「100d」や「1h」などです。
      
      このオプションを使用する場合、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な書式については、[時間オプションのドキュメント](/docs/#time-option)を参照してください。

   --decompress
      設定されている場合、gzipエンコードされたオブジェクトを解凍します。
      
      S3に「Content-Encoding: gzip」が設定されたファイルをアップロードすることも可能です。通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneは受信時に「Content-Encoding: gzip」と共にこれらのファイルを解凍します。これにより、rcloneはサイズとハッシュをチェックできなくなりますが、ファイルの内容は解凍されます。

   --might-gzip
      バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされる際に変更しません。`Content-Encoding: gzip`でアップロードされていないオブジェクトには設定されません。
      
      ただし、一部のプロバイダ（Cloudflareなど）は、`Content-Encoding: gzip`でアップロードされていなくてもオブジェクトをgzip圧縮する場合があります。
      
      これによって次のようなエラーが発生する場合があります。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      このフラグを設定し、rcloneがContent-Encoding: gzipが設定され、チャンク付き転送エンコードでオブジェクトをダウンロードすると、rcloneはオブジェクトを逐次解凍します。
      
      unset（デフォルト）に設定されている場合、rcloneはプロバイダの設定に従って適用するものを選択しますが、ここでrcloneの選択を上書きできます。

   --no-system-metadata
      システムメタデータの設定および読み取りの抑制


オプション:
   --access-key-id value        AWSのアクセスキーID。[$ACCESS_KEY_ID]
   --acl value                  バケットを作成したり、オブジェクトを保存またはコピーする際に使用するキャンドACL。[$ACL]
   --endpoint value             S3 APIのエンドポイント。[$ENDPOINT]
   --env-auth                   実行時にAWS認証情報（環境変数または環境変数がない場合はEC2 / ECSメタデータ）から取得する。[デフォルト: false] [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンに合わせて設定する位置制約。[$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。[$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。[$SECRET_ACCESS_KEY]

   高度なオプション

   --bucket-acl value               バケットを作成する際に使用するキャンドACL。[$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。[デフォルト: "5Mi"] [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフ値。[デフォルト: "4.656Gi"] [$COPY_CUTOFF]
   --decompress                     設定されている場合、gzipエンコードされたオブジェクトを解凍します。[デフォルト: false] [$DECOMPRESS]
   --disable-checksum               オブジェクトメタデータにMD5チェックサムを保存しない。[デフォルト: false] [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでのhttp2の使用を無効にします。[デフォルト: false] [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。[$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。[デフォルト: "Slash,InvalidUtf8,Dot"] [$ENCODING]
   --force-path-style               もし真ならパス形式アクセスを使い、もし偽なら仮想ホスト形式アクセスを使う。[デフォルト: true] [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストに対する応答リストのサイズ）。[デフォルト: 1000] [$LIST_CHUNK]
   --list-url-encode value          リストのURLエンコードを行うかどうか：true/false/unset [デフォルト: "unset"] [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または自動の場合は0。[デフォルト: 0] [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードでの最大パート数。[デフォルト: 10000] [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールがフラッシュされる頻度。[デフォルト: "1m0s"] [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           メモリプール内のmmapバッファを使用するかどうか。[デフォルト: false] [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip圧縮する可能性がある場合に設定します。[デフォルト: "unset"] [$MIGHT_GZIP]
   --no-check-bucket                バケットの存在をチェックせず、作成しようとしません。[デフォルト: false] [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADをチェックしない場合に設定します。[デフォルト: false] [$NO_HEAD]
   --no-head-object                 GETする前にHEADを実行しない場合に設定します。[デフォルト: false] [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定および読み取りの抑制[デフォルト: false] [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。[$PROFILE]
   --session-token value            AWSのセッショントークン。[$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。[$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行度。[デフォルト: 4] [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンク化アップロードへの切り替えのためのカットオフ値。[デフォルト: "200Mi"] [$UPLOAD_CUTOFF]
   --use-multipart-etag value       検証のためにマルチパートアップロードでETagを使用するかどうか [デフォルト: "unset"] [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロードに署名済みリクエストまたはPutObjectを使用するかどうか[デフォルト: false] [$USE_PRESIGNED_REQUEST]
   --v2-auth                        もし真ならv2認証を使用する。[デフォルト: false] [$V2_AUTH]
   --version-at value               指定した時間の存在した時点のファイルバージョンを表示します。[デフォルト: "off"] [$VERSION_AT]
   --versions                       ディレクトリリスティングに古いバージョンを含める。[デフォルト: false] [$VERSIONS]

```