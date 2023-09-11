# QingCloudオブジェクトストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity storage update qingstor - QingCloudオブジェクトストレージ

USAGE:
   singularity storage update qingstor [コマンドオプション] <名前またはID>

DESCRIPTION:
   --env-auth
      ランタイムからQingStorの認証情報を取得します。

      access_key_idとsecret_access_keyが空白の場合にのみ適用されます。

      例:
         | false | 次のステップでQingStorの認証情報を入力します。
         | true  | 環境変数またはIAMからQingStorの認証情報を取得します。

   --access-key-id
      QingStorのアクセスキーIDです。

      匿名アクセスまたはランタイム認証情報の場合は空白のままでください。

   --secret-access-key
      QingStorのシークレットアクセスキー（パスワード）です。

      匿名アクセスまたはランタイム認証情報の場合は空白のままでください。

   --endpoint
      接続するためのエンドポイントURLを入力してください。

      空白の場合はデフォルト値 "https://qingstor.com:443" を使用します。

   --zone
      接続するゾーンです。

      デフォルトは "pek3a" です。

      例:
         | pek3a | 北京（中国）の第3ゾーンです。
         |       | location constraint pek3aが必要です。
         | sh1a  | 上海（中国）の第1ゾーンです。
         |       | location constraint sh1aが必要です。
         | gd2a  | 広東（中国）の第2ゾーンです。
         |       | location constraint gd2aが必要です。

   --connection-retries
      接続のリトライ回数です。

   --upload-cutoff
      チャンク化アップロードに切り替えるカットオフサイズです。

      これより大きいファイルはchunk_sizeで指定したサイズのチャンクでアップロードされます。
      最小値は0、最大値は5 GiBです。

   --chunk-size
      アップロードに使用するチャンクサイズです。

      upload_cutoffより大きいファイルはmultipartアップロードとして、このチャンクサイズを使用してアップロードされます。

      注意："--qingstor-upload-concurrency"は転送ごとのこのサイズのチャンクをメモリにバッファリングします。

      高速リンクを介して大きなファイルを転送し、十分なメモリがある場合は、これを増やすと転送が高速化されます。

   --upload-concurrency
      multipartアップロードの並行性です。

      同じファイルのチャンクを同時にアップロードする数です。

      注：これを1より大きく設定すると、multipartアップロードのチェックサムが破損します（アップロード自体は破損しません）。

      高速リンクを介して少数の大きなファイルを転送し、これらの転送が帯域幅を十分に利用しない場合は、これを増やすことで転送の高速化に役立つ場合があります。

   --encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --access-key-id value      QingStorのアクセスキーIDです。 [$ACCESS_KEY_ID]
   --endpoint value           接続するためのエンドポイントURLを入力してください。 [$ENDPOINT]
   --env-auth                 ランタイムからQingStorの認証情報を取得します。 (デフォルト: false) [$ENV_AUTH]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  QingStorのシークレットアクセスキー（パスワード）です。 [$SECRET_ACCESS_KEY]
   --zone value               接続するゾーンです。 [$ZONE]

   Advanced

   --chunk-size value          アップロードに使用するチャンクサイズです。 (デフォルト: "4Mi") [$CHUNK_SIZE]
   --connection-retries value  接続のリトライ回数です。 (デフォルト: 3) [$CONNECTION_RETRIES]
   --encoding value            バックエンドのエンコーディングです。 (デフォルト: "Slash,Ctl,InvalidUtf8") [$ENCODING]
   --upload-concurrency value  multipartアップロードの並行性です。 (デフォルト: 1) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       チャンク化アップロードに切り替えるカットオフサイズです。 (デフォルト: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}