# Internet Archive

{% code fullWidth="true" %}
```
名前:
   singularity storage create internetarchive - Internet Archive

使用方法:
   singularity storage create internetarchive [コマンドオプション] [引数...]

説明:
   --access-key-id
      IAS3のアクセスキーです。
      
      匿名アクセスの場合は空白で入力します。
      こちらで取得できます: https://archive.org/account/s3.php

   --secret-access-key
      IAS3のシークレットキー（パスワード）です。
      
      匿名アクセスの場合は空白で入力します。

   --endpoint
      IAS3のエンドポイントです。
      
      デフォルト値のままで使用する場合は空白で入力します。

   --front-endpoint
      Internet Archiveフロントエンドのホストです。
      
      デフォルト値のままで使用する場合は空白で入力します。

   --disable-checksum
      サーバーにrcloneによって計算されたMD5チェックサムのテストを要求しません。
      通常、rcloneはアップロード前に入力のMD5チェックサムを計算し、サーバーにオブジェクトのチェックサムを要求します。
      これはデータの整合性チェックには最適ですが、大きなファイルのアップロードの際に長時間の遅延を引き起こす可能性があります。

   --wait-archive
      サーバーの処理タスク（特にアーカイブとbook_op）が終了するまでのタイムアウト時間です。
      書き込み操作後に反映されることを保証する必要がある場合にのみ有効にします。
      タイムアウトの場合はエラーが発生せず、処理が続行されます。
      タイムアウトを無効にする場合は0と入力します。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --access-key-id value      IAS3のアクセスキーです。[$ACCESS_KEY_ID]
   --help, -h                 ヘルプを表示します
   --secret-access-key value  IAS3のシークレットキー（パスワード）です。[$SECRET_ACCESS_KEY]

   上級者向け

   --disable-checksum      サーバーにrcloneによって計算されたMD5チェックサムのテストを要求しません。 (デフォルト: true) [$DISABLE_CHECKSUM]
   --encoding value        バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3のエンドポイントです。 (デフォルト: "https://s3.us.archive.org") [$ENDPOINT]
   --front-endpoint value  Internet Archiveフロントエンドのホストです。 (デフォルト: "https://archive.org") [$FRONT_ENDPOINT]
   --wait-archive value    サーバーの処理タスク（特にアーカイブとbook_op）が終了するまでのタイムアウト時間です。 (デフォルト: "0s") [$WAIT_ARCHIVE]

   一般

   --name value  ストレージの名前です（デフォルト: 自動生成）
   --path value  ストレージのパスです
```
{% endcode %}