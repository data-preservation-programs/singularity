# Internet Archive

{% code fullWidth="true" %}
```
名前：
   singularity datasource add internetarchive - Internet Archive

使用法：
   singularity datasource add internetarchive [コマンドオプション] <データセット名> <ソースパス>

説明：
   --internetarchive-access-key-id
      IAS3アクセスキーです。
      
      匿名アクセスの場合は空白にしてください。
      ここで見つけることができます: https://archive.org/account/s3.php

   --internetarchive-disable-checksum
      サーバーにrcloneによって計算されたMD5チェックサムをテストするように要求しないでください。
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、サーバーにチェックサムでオブジェクトを確認するように要求します。
      これはデータの整合性チェックには素晴らしいですが、大きなファイルのアップロードの開始に長時間かかる場合があります。

   --internetarchive-encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --internetarchive-endpoint
      IAS3エンドポイントです。
      
      デフォルト値の場合は空白のままにしてください。

   --internetarchive-front-endpoint
      Internet Archiveのフロントエンドのホストです。
      
      デフォルト値の場合は空白のままにしてください。

   --internetarchive-secret-access-key
      IAS3シークレットキー（パスワード）です。
      
      匿名アクセスの場合は空白にしてください。

   --internetarchive-wait-archive
      サーバーの処理タスク（特にarchiveとbook_op）が完了するまでの待機タイムアウトです。
      書き込み操作後に反映されることを保証する必要がある場合にのみ有効にします。
      待機を無効にするには0です。タイムアウトの場合にエラーがスローされません。


オプション：
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをエクスポート後に削除します。 (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過した場合、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   Internet Archiveのオプション

   --internetarchive-access-key-id value      IAS3アクセスキーです。 [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   サーバーにrcloneによって計算されたMD5チェックサムをテストするように要求しないでください。 (デフォルト: "true") [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3エンドポイントです。 (デフォルト: "https://s3.us.archive.org") [$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     Internet Archiveのフロントエンドのホストです。 (デフォルト: "https://archive.org") [$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3シークレットキー（パスワード）です。 [$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       サーバーの処理タスク（特にarchiveとbook_op）が完了するまでの待機タイムアウトです。 (デフォルト: "0s") [$INTERNETARCHIVE_WAIT_ARCHIVE]

```
{% endcode %}