# その他のS3互換プロバイダ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 other - その他のS3互換プロバイダ

USAGE:
   singularity storage update s3 other [command options] <name|id>

DESCRIPTION:
   
   --env-auth
      実行時にAWS認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSのメタデータから取得）。
      
      access_key_idとsecret_access_keyが空白の場合に適用されます。

      例:
         | false | 次のステップでAWS認証情報を入力します。
         | true  | 環境変数（env varsまたはIAM）からAWS認証情報を取得します。

   --access-key-id
      AWSのアクセスキーID。
      
      匿名アクセスまたはランタイム認証情報を使用する場合は、空白のままにしてください。

   --secret-access-key
      AWSのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたはランタイム認証情報を使用する場合は、空白のままにしてください。

   --region
      接続するリージョン。
      
      S3クローンを使用していて、リージョンが不要な場合は、空白にしてください。

      例:
         | <未設定>               | 未確定の場合はこれを使用してください。
         |                     | v4シグネチャと空のリージョンを使用します。
         | other-v2-signature | v4シグネチャが機能しない場合にのみ使用してください。
         |                     | 例: pre Jewel/v10 CEPH.

   --endpoint
      S3 APIのエンドポイント。
      
      S3クローンを使用している場合は必須です。

   --location-constraint
      リージョンと一致するように設定する場所の制約。
      
      わからない場合は空白のままにしてください。バケットの作成時にのみ使用されます。

   --acl
      バケットの作成、オブジェクトの保存またはコピー時に使用されるCanned ACL。
      
      このACLはオブジェクトの作成時にも使用されます。bucket_aclが設定されていない場合はデフォルトで使用されます。
      
      詳細については、こちらを参照してください。[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      S3はサーバーサイドでオブジェクトをコピーする際にACLをコピーしないため、
      ソースからACLをコピーするのではなく、新しいACLを書き込みます。
      
      aclが空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。
      

   --bucket-acl
      バケットの作成時に使用されるCanned ACL。
      
      詳細については、こちらを参照してください。[https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      bucket_aclが設定されていない場合は「acl」が代わりに使用されます。
      
      「acl」と「bucket_acl」が空の文字列の場合、X-Amz-Acl: ヘッダは追加されず、デフォルト（private）が使用されます。
      

      例:
         | private               | オーナーはFULL_CONTROLを取得します。
         |                     | 他のユーザーにはアクセス権がありません（デフォルト）。
         | public-read        | オーナーはFULL_CONTROLを取得します。
         |                     | AllUsersグループは読み取りアクセスが可能です。
         | public-read-write  | オーナーはFULL_CONTROLを取得します。
         |                     | AllUsersグループは読み書きアクセスが可能です。
         |                     | バケットでこれを許可することは一般的には推奨されていません。
         | authenticated-read | オーナーはFULL_CONTROLを取得します。
         |                     | AuthenticatedUsersグループは読み取りアクセスが可能です。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフサイズ。
      
      これを超えるファイルは、chunk_sizeごとにチャンクでアップロードされます。
      最小値は0で、最大値は5GiBです。

   --chunk-size
      アップロード時に使用するチャンクサイズ。
      
      upload_cutoffを超えるファイルやサイズが不明なファイル（「rclone rcat」や「rclone mount」、GoogleフォトやGoogleドキュメントなど）は、
      このチャンクサイズを使用して、マルチパートアップロードとしてアップロードされます。
      
      注意："--s3-upload-concurrency"は、転送ごとにこのサイズのチャンクがメモリにバッファリングされます。
      
      高速リンクで大きなファイルを転送し、メモリが十分にある場合は、これを増やすと転送速度が向上します。
      
      Rcloneは、既知の大きさの大きなファイルをアップロードする場合、
      10,000個のチャンクの制限を下回るように自動的にチャンクサイズを増やします。
      
      不明なサイズのファイルは、設定された
      チャンクサイズでアップロードされます。デフォルトのチャンクサイズが5 MiBであり、最大で
      10,000のチャンクがあるため、デフォルトでは最大サイズが
      ストリームアップロードできるファイルのサイズは48 GiBです。より大きなファイルをストリームアップロードする場合は、
      チャンクサイズを増やす必要があります。
      
      チャンクサイズを増やすと、進行状況
      統計情報が "-P" フラグで表示される精度が低下します。Rcloneは、チャンクが送信されたとみなされる時点で、
      AWS SDKによってバッファリングされたときにチャンクとして処理しますが、実際にはまだアップロード中の場合があります。
      チャンクサイズが大きいほど、AWS SDKのバッファサイズが大きくなり、進行状況の報告の
      真実からのずれが大きくなります。
      

   --max-upload-parts
      マルチパートアップロードの最大パーツ数。
      
      このオプションは、マルチパートアップロード時に使用するマルチパートチャンクの最大数を定義します。
      
      これは、サービスがAWS S3の仕様である10,000のチャンクをサポートしていない場合に便利です。
      
      Rcloneは、既知のサイズの大きなファイルをアップロードする場合、
      チャンクサイズを増やしてこのチャンク数の制限を下回るように自動的に増やします。
      

   --copy-cutoff
      マルチパートコピーに切り替えるためのカットオフサイズ。
      
      サーバーサイドでコピーする必要のあるこれより大きなファイルは、
      このサイズのチャンクでコピーされます。
      
      最小値は0で、最大値は5GiBです。

   --disable-checksum
      オブジェクトのメタデータにMD5チェックサムを保存しません。
      
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、オブジェクトのメタデータに追加するため、
      大きなファイルのアップロードの開始までには長い遅延が発生する場合があります。

   --shared-credentials-file
      共有認証情報ファイルへのパス。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用できます。
      
      この変数が空の場合、rcloneは
      "AWS_SHARED_CREDENTIALS_FILE" の環境変数を探します。環境変数の値が空の場合は、現在のユーザーのホームディレクトリがデフォルトです。
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      共有認証情報ファイルで使用するプロファイル。
      
      env_authがtrueの場合、rcloneは共有認証情報ファイルを使用できます。この
      変数は、そのファイルで使用されるプロファイルを制御します。
      
      空の場合、環境変数 "AWS_PROFILE" または
      "default" が設定されていない場合にデフォルトで使用されます。
      

   --session-token
      AWSセッショントークン.

   --upload-concurrency
      マルチパートアップロードの並行数。
      
      同じファイルのチャンクを同時にアップロードする数。

      高速なリンクで大量の大きなファイルをアップロードし、これらのアップロードが帯域幅を十分に活用していない場合、
      これを増やすことで転送速度を向上させることができます。

   --force-path-style
      trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。
      
      これがtrue（デフォルト）の場合、rcloneはパススタイルアクセスを使用しますが、
      falseの場合、rcloneは仮想パススタイルを使用します。詳細については、[AWS S3のドキュメント](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      を参照してください。
      
      一部のプロバイダ（例：AWS、Aliyun OSS、Netease COS、Tencent COS）では、これを
      falseに設定する必要があります - Rcloneはプロバイダの設定に基づいてこれを自動的に行います。

   --v2-auth
      trueの場合、v2認証を使用します。
      
      これがfalse（デフォルト）の場合、rcloneはv4認証を使用します。設定されている場合、rcloneはv2認証を使用します。
      
      v4シグネチャが機能しない場合にのみ使用します。例：pre Jewel/v10 CEPH.

   --list-chunk
      リストのチャンクサイズ（各ListObject S3リクエストごとの応答リストのサイズ）。
      
      このオプションは、AWS S3の仕様の「MaxKeys」、「max-items」、「page-size」とも呼ばれます。
      大多数のサービスは、要求されたものよりも多くのリストを返さないため、このサイズを1000オブジェクトに切り捨てます。
      AWS S3では、これはグローバルな最大値であり、変更できません。[AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)を参照してください。
      Cephでは、「rgw list buckets max chunk」オプションでこれを増やすことができます。
      

   --list-version
      使用するListObjectsのバージョン：1、2、または0（自動）。
      
      S3が最初にリリースされたとき、バケット内のオブジェクトを列挙するためのListObjects呼び出しのみが提供されました。
      
      しかし、2016年5月にListObjectsV2呼び出しが導入されました。これは大幅に高速化され、できるだけ使用する必要があります。
      
      デフォルトで0に設定されている場合、rcloneはプロバイダが使用するリストオブジェクトメソッドを推測します。
      誤った推測をした場合は、ここで手動で設定できます。
      

   --list-url-encode
      リストをURLエンコードするかどうか：true/false/unset
      
      一部のプロバイダはリストをURLエンコードすることをサポートしており、
      ファイル名に制御文字を使用する場合にはこれが信頼性があります。これがunsetに設定されている場合（デフォルト）は、
      rcloneはプロバイダの設定に応じて適用するものを選択しますが、ここでrcloneの選択を上書きできます。
      

   --no-check-bucket
      該当バケットの存在を確認せずに作成しようとしないように設定します。
      
      バケットが既に存在する場合、
      rcloneのトランザクション数を最小限にするために使用することができます。
      
      ユーザーがバケット作成の権限を持っていない場合にも必要です。v1.52.0以前の場合、これはバグのために
      静かに合格してしまいます。

   --no-head
      アップロードされたオブジェクトのHEADリクエストを送信して整合性をチェックしません。
      
      rcloneは、PUTでオブジェクトのアップロード後に200 OKメッセージを受け取ると、正しくアップロードされたと想定します。
      
      特に次のことを想定します。
      
      - アップロード時のメタデータ、モディファイ時間、ストレージクラス、コンテンツタイプがアップロード時と同じであること
      - サイズがアップロード時と同じであること
      
      PUTの場合、以下の項目を応答から読み取ります。
      
      - MD5SUM
      - アップロード日時
      
      マルチパートアップロードの場合、これらの項目は読み取りません。
      
      不明な長さのソースオブジェクトがアップロードされた場合、rcloneはHEADリクエストを確実に実行します。
      
      このフラグを設定すると、アップロードの失敗を検出できる可能性が高くなります。
      特にサイズが正しくない場合ですが、通常の操作では推奨されません。実際には、このフラグを使用しても、
      アップロードの失敗が検出できる確率は非常に低いです。

   --no-head-object
      オブジェクトを取得する際にHEADを実行しないように設定します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --memory-pool-flush-time
      内部メモリバッファプールをフラッシュする頻度。
      
      追加のバッファを必要とするアップロード（マルチパートなど）は、割り当てにメモリプールを使用します。
      このオプションは、使用されなくなったバッファがプールから削除される頻度を制御します。

   --memory-pool-use-mmap
      内部メモリプールでmmapバッファを使用するかどうか。

   --disable-http2
      S3バックエンドでhttp2の使用を無効にします。
      
      現在、s3（特にminio）バックエンドとHTTP/2に関する未解決の問題があります。S3バックエンドではデフォルトで
      HTTP/2が有効ですが、ここで無効にすることもできます。問題が解決されたら、このフラグは削除されます。
      
      参照：https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      ダウンロード用のカスタムエンドポイント。
      通常、AWS S3はCloudFrontネットワークを介してデータをダウンロードするときにより安価な転送料金を提供します。

   --use-multipart-etag
      バリデーションのためにマルチパートアップロードでETagを使用するかどうか
      
      これは、true、false、またはプロバイダのデフォルトを使用するように設定します。
      

   --use-presigned-request
      シングルパートアップロード用の署名済みリクエストまたはPutObjectを使用するかどうか
      
      これがfalseの場合、rcloneはAWS SDKのPutObjectを使用してオブジェクトをアップロードします。
      
      rcloneのバージョン1.59未満では、単一のパートオブジェクトをアップロードするために署名済みリクエストを使用し、
      このフラグをtrueに設定すると、その機能が再有効になります。これは例外的な状況またはテスト以外で必要ではないはずです。
      

   --versions
      ディレクトリリストに古いバージョンを含めるかどうか。

   --version-at
      指定した時間のファイルバージョンを表示します。
      
      パラメータは日付、「2006-01-02」、日時「2006-01-02
      15:04:05」、またはその時間前の期間、例えば「100d」または「1h」です。
      
      このオプションを使用すると、ファイルの書き込み操作は許可されませんので、ファイルをアップロードしたり削除したりすることはできません。
      
      有効な形式については[時間オプションのドキュメント](/docs/#time-option)を参照してください。
      

   --decompress
      このフラグが設定されている場合、gzipでエンコードされたオブジェクトの解凍を行います。
      
      S3へのアップロード時に「Content-Encoding: gzip」が設定されている可能性があります。
      通常、rcloneはこれらのファイルを圧縮されたオブジェクトとしてダウンロードします。
      
      このフラグが設定されている場合、rcloneはこれらのファイルを「Content-Encoding: gzip」で受信されたままの状態で解凍します。
      これにより、rcloneはサイズとハッシュを確認することはできませんが、
      ファイル内容は解凍されます。
      

   --might-gzip
      バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。
      
      通常、プロバイダはオブジェクトがダウンロードされる際にオブジェクトを変更しません。
      「Content-Encoding: gzip」でアップロードされていないオブジェクトには設定されません。
      
      ただし、一部のプロバイダ（例：Cloudflare）は、「Content-Encoding: gzip」でアップロードされていなくても
      オブジェクトをgzip化する場合があります。
      
      これによって受信エラーが発生することがあります。
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      もし、このフラグを設定し、rcloneが「Content-Encoding: gzip」が設定されたオブジェクトを
      チャンク転送エンコードでダウンロードした場合、rcloneはオブジェクトを逐次解凍します。
      
      unsetに設定されている場合（デフォルト）は、rcloneはプロバイダの設定に応じて適用することを選択しますが、
      ここでrcloneの選択を上書きできます。
      

   --no-system-metadata
      システムメタデータの設定と読み取りを抑制します。

OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  バケットの作成およびオブジェクトの保存またはコピー時に使用されるCanned ACL。 [$ACL]
   --endpoint value             S3 APIのエンドポイント。 [$ENDPOINT]
   --env-auth                   実行時にAWS認証情報を取得します（環境変数または環境変数がない場合はEC2/ECSのメタデータから取得）。（デフォルト：false） [$ENV_AUTH]
   --help, -h                   ヘルプを表示
   --location-constraint value  リージョンと一致するように設定する場所の制約。 [$LOCATION_CONSTRAINT]
   --region value               接続するリージョン。 [$REGION]
   --secret-access-key value    AWSのシークレットアクセスキー（パスワード）。 [$SECRET_ACCESS_KEY]

   高度なオプション

   --bucket-acl value               バケットの作成時に使用されるCanned ACL。 [$BUCKET_ACL]
   --chunk-size value               アップロード時に使用するチャンクサイズ。（デフォルト："5Mi"） [$CHUNK_SIZE]
   --copy-cutoff value              マルチパートコピーに切り替えるためのカットオフサイズ。（デフォルト："4.656Gi"） [$COPY_CUTOFF]
   --decompress                     このフラグが設定されている場合、gzipでエンコードされたオブジェクトを解凍します。（デフォルト：false） [$DECOMPRESS]
   --disable-checksum               オブジェクトのメタデータにMD5チェックサムを保存しません。（デフォルト：false） [$DISABLE_CHECKSUM]
   --disable-http2                  S3バックエンドでhttp2の使用を無効にします。（デフォルト：false） [$DISABLE_HTTP2]
   --download-url value             ダウンロード用のカスタムエンドポイント。 [$DOWNLOAD_URL]
   --encoding value                 バックエンドのエンコーディング。（デフォルト：「Slash,InvalidUtf8,Dot」） [$ENCODING]
   --force-path-style               trueの場合、パススタイルアクセスを使用し、falseの場合は仮想ホストスタイルを使用します。（デフォルト：true） [$FORCE_PATH_STYLE]
   --list-chunk value               リストのチャンクサイズ（各ListObject S3リクエストごとの応答リストのサイズ）。 （デフォルト：1000） [$LIST_CHUNK]
   --list-url-encode value          リストをURLエンコードするかどうか：true/false/unset（デフォルト："unset"） [$LIST_URL_ENCODE]
   --list-version value             使用するListObjectsのバージョン：1、2、または0（自動）。（デフォルト：0） [$LIST_VERSION]
   --max-upload-parts value         マルチパートアップロードの最大パーツ数。（デフォルト：10000） [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   内部メモリバッファプールをフラッシュする頻度。（デフォルト："1m0s"） [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           内部メモリプールでmmapバッファを使用するかどうか。（デフォルト：false） [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               バックエンドがオブジェクトをgzip化する可能性がある場合に設定します。（デフォルト："unset"） [$MIGHT_GZIP]
   --no-check-bucket                該当バケットの存在を確認せずに作成しようとしないように設定します。（デフォルト：false） [$NO_CHECK_BUCKET]
   --no-head                        アップロードされたオブジェクトのHEADリクエストを送信して整合性をチェックしません。（デフォルト：false） [$NO_HEAD]
   --no-head-object                 オブジェクトを取得する際にHEADを実行しないように設定します。（デフォルト：false） [$NO_HEAD_OBJECT]
   --no-system-metadata             システムメタデータの設定と読み取りを抑制します。（デフォルト：false） [$NO_SYSTEM_METADATA]
   --profile value                  共有認証情報ファイルで使用するプロファイル。 [$PROFILE]
   --session-token value            AWSセッショントークン。 [$SESSION_TOKEN]
   --shared-credentials-file value  共有認証情報ファイルへのパス。 [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       マルチパートアップロードの並行数。（デフォルト：4） [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            チャンクアップロードに切り替えるためのカットオフサイズ。（デフォルト："200Mi"） [$UPLOAD_CUTOFF]
   --use-multipart-etag value       バリデーションのためにマルチパートアップロードでETagを使用するかどうか（デフォルト："unset"） [$USE_MULTIPART_ETAG]
   --use-presigned-request          シングルパートアップロード用の署名済みリクエストまたはPutObjectを使用するかどうか。（デフォルト：false） [$USE_PRESIGNED_REQUEST]
   --v2-auth                        trueの場合、v2認証を使用します。（デフォルト：false） [$V2_AUTH]
   --version-at value               指定した時間のファイルバージョンを表示します。（デフォルト："off"） [$VERSION_AT]
   --versions                       ディレクトリリストに古いバージョンを含めるかどうか。（デフォルト：false） [$VERSIONS]

```
{% endcode %}