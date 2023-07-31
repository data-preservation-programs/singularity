# QingCloud オブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add qingstor - QingCloudオブジェクトストレージ

使い方:
   singularity datasource add qingstor [コマンドオプション] <データセット名> <ソースパス>

説明:
   --qingstor-access-key-id
      QingStorのアクセスキーID。
      
      匿名アクセスまたはランタイム認証情報の場合は空白のままでください。

   --qingstor-chunk-size
      アップロードに使用するチャンクサイズ。
      
      アップロード切り替えサイズ以上のファイルは、このチャンクサイズを使用して
      マルチパートアップロードされます。
      
      注意: "--qingstor-upload-concurrency" のチャンクサイズは転送ごとにメモリ上で
      バッファリングされます。
      
      高速リンク経由で大きなファイルを転送しており、メモリが十分にある場合は、
      これを増やすことで転送速度が向上します。

   --qingstor-connection-retries
      接続のリトライ回数。

   --qingstor-encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --qingstor-endpoint
      QingStor APIに接続するためのエンドポイントURLを入力してください。
      
      空白の場合はデフォルト値 "https://qingstor.com:443" が使用されます。

   --qingstor-env-auth
      ランタイムからQingStorの認証情報を取得します。
      
      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでQingStorの認証情報を入力します。
         | true  | 環境変数またはIAMからQingStorの認証情報を取得します。

   --qingstor-secret-access-key
      QingStorのシークレットアクセスキー（パスワード）。
      
      匿名アクセスまたはランタイム認証情報の場合は空白のままでください。

   --qingstor-upload-concurrency
      マルチパートアップロードの同時実行数。
      
      これは同じファイルのチャンクの並行してアップロードされる数です。
      
      注意: これを1より大きな値に設定すると、マルチパートアップロードの
      チェックサムが破損してしまいますが（アップロード自体は破損しません）。
      
      高速リンク経由で大量の大きなファイルを転送しており、これらのアップロードが
      帯域幅を十分に利用していない場合、これを増やすことで転送速度を向上させることができます。

   --qingstor-upload-cutoff
      チャンクアップロードに切り替えるためのカットオフサイズ。
      
      このサイズより大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --qingstor-zone
      接続するゾーン。
      
      デフォルト値は "pek3a" です。

      例:
         | pek3a | 北京（中国）第3ゾーン。
                 | ロケーション制約pek3aが必要です。
         | sh1a  | 上海（中国）第1ゾーン。
                 | ロケーション制約sh1aが必要です。
         | gd2a  | 広東（中国）第2ゾーン。
                 | ロケーション制約gd2aが必要です。


オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export    [危険] データセットのファイルをCARファイルにエクスポートした後、削除します。  (デフォルト: false)
   --rescan-interval value  最後のスキャンからこの間隔が経過すると、自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   QingStor向けオプション

   --qingstor-access-key-id value       QingStorのアクセスキーID。 [$QINGSTOR_ACCESS_KEY_ID]
   --qingstor-chunk-size value          アップロードに使用するチャンクサイズ。 (デフォルト: "4Mi") [$QINGSTOR_CHUNK_SIZE]
   --qingstor-connection-retries value  接続のリトライ回数。 (デフォルト: "3") [$QINGSTOR_CONNECTION_RETRIES]
   --qingstor-encoding value            バックエンドのエンコーディング。 (デフォルト: "Slash,Ctl,InvalidUtf8") [$QINGSTOR_ENCODING]
   --qingstor-endpoint value            QingStor APIに接続するためのエンドポイントURLを入力してください。 [$QINGSTOR_ENDPOINT]
   --qingstor-env-auth value            ランタイムからQingStorの認証情報を取得します。 (デフォルト: "false") [$QINGSTOR_ENV_AUTH]
   --qingstor-secret-access-key value   QingStorのシークレットアクセスキー（パスワード）。 [$QINGSTOR_SECRET_ACCESS_KEY]
   --qingstor-upload-concurrency value  マルチパートアップロードの同時実行数。 (デフォルト: "1") [$QINGSTOR_UPLOAD_CONCURRENCY]
   --qingstor-upload-cutoff value       チャンクアップロードに切り替えるためのカットオフサイズ。 (デフォルト: "200Mi") [$QINGSTOR_UPLOAD_CUTOFF]
   --qingstor-zone value                接続するゾーン。 [$QINGSTOR_ZONE]

```
{% endcode %}