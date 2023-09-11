# QingCloudオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage create qingstor - QingCloudオブジェクトストレージ

USAGE:
   singularity storage create qingstor [コマンドオプション] [引数...]

DESCRIPTION:
   --env-auth
      ランタイムからQingStorの認証情報を取得します。

      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでQingStorの認証情報を入力します。
         | true  | 環境変数やIAMからQingStorの認証情報を取得します。

   --access-key-id
      QingStorのアクセスキーIDです。

      匿名アクセスまたは実行時の認証情報の場合は空白のままにしてください。

   --secret-access-key
      QingStorのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたは実行時の認証情報の場合は空白のままにしてください。

   --endpoint
      QingStor APIへの接続に使用するエンドポイントURLを入力してください。

      空白の場合はデフォルト値 "https://qingstor.com:443" が使用されます。

   --zone
      接続するゾーンです。

      デフォルトは "pek3a" です。

      例:
         | pek3a | 北京（中国）第3ゾーン。pek3aの場所制約が必要です。
         |       |
         | sh1a  | 上海（中国）第1ゾーン。sh1aの場所制約が必要です。
         |       |
         | gd2a  | 広東（中国）第2ゾーン。gd2aの場所制約が必要です。

   --connection-retries
      接続の再試行回数です。

   --upload-cutoff
      チャンクアップロードに切り替えるためのカットオフです。

      この値より大きなファイルは、chunk_sizeのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクのサイズです。

      upload_cutoffより大きなファイルは、このチャンクサイズを使用してマルチパートアップロードされます。

      注意: "--qingstor-upload-concurrency"は、転送ごとにこのサイズのチャンクがメモリにバッファリングされます。

      高速リンク経由で大きなファイルを転送し、十分なメモリがある場合は、この値を増やすと転送が高速化されます。

   --upload-concurrency
      マルチパートアップロードの並行性です。

      同じファイルのチャンクの数を同時にアップロードします。

      注意: この値を1より大きく設定すると、マルチパートアップロードのチェックサムが破損します（ただし、アップロード自体は破損しません）。

      少数の大きなファイルを高速リンク経由でアップロードし、これらのアップロードが帯域幅を十分に利用しない場合は、この値を増やすと転送が高速化される場合があります。

   --encoding
      バックエンドのエンコーディングです。

      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --access-key-id value      QingStorのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --endpoint value           QingStor APIへの接続に使用するエンドポイントURLです。 [$ENDPOINT]
   --env-auth                 ランタイムからQingStorの認証情報を取得します。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  QingStorのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --zone value               接続するゾーンです。 [$ZONE]

   Advanced（詳細オプション）

   --chunk-size value          アップロードに使用するチャンクのサイズです。 (デフォルト: "4Mi") [$CHUNK_SIZE]
   --connection-retries value  接続の再試行回数です。 (デフォルト: 3) [$CONNECTION_RETRIES]
   --encoding value            バックエンドのエンコーディングです。 (デフォルト: "Slash,Ctl,InvalidUtf8") [$ENCODING]
   --upload-concurrency value  マルチパートアップロードの並行性です。 (デフォルト: 1) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       チャンクアップロードに切り替えるためのカットオフです。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]

   General（一般オプション）

   --name value  ストレージの名前です（デフォルト: 自動生成）
   --path value  ストレージのパスです

```
{% endcode %}