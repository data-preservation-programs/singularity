# Internet Archive

{% code fullWidth="true" %}
```
名前:
   singularity storage update internetarchive - インターネットアーカイブ

利用方法:
   singularity storage update internetarchive [コマンドオプション] <名前|ID>
   
説明:
   --access-key-id
      IAS3アクセスキー。
      
      匿名アクセスの場合は空白のままにします。
      ここでアクセスキーを入手できます: https://archive.org/account/s3.php

   --secret-access-key
      IAS3シークレットキー（パスワード）。
      
      匿名アクセスの場合は空白のままにします。

   --endpoint
      IAS3のエンドポイント。
      
      デフォルト値のままにします。

   --front-endpoint
      インターネットアーカイブのフロントエンドのホスト。
      
      デフォルト値のままにします。

   --disable-checksum
      サーバーにrcloneによって計算されたMD5チェックサムをテストするように要求しないでください。
      通常、rcloneはアップロードする前に入力のMD5チェックサムを計算し、サーバーにチェックサムを使ってオブジェクトを確認するように要求します。
      これはデータの整合性チェックには最適ですが、大きなファイルのアップロードの開始には長い遅延を引き起こす可能性があります。

   --wait-archive
      サーバーの処理タスク（特にアーカイブとbook_op）の終了を待つタイムアウト。
      書き込み操作後に反映されることが保証される場合にのみ有効にします。
      タイムアウトした場合にエラーが発生しないようにするため、0で無効化します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --access-key-id value      IAS3アクセスキー。 [$ACCESS_KEY_ID]
   --help, -h                 ヘルプを表示
   --secret-access-key value  IAS3シークレットキー（パスワード）。 [$SECRET_ACCESS_KEY]

   高度な設定

   --disable-checksum      サーバーにrcloneによって計算されたMD5チェックサムをテストするように要求しないでください。（デフォルト：true） [$DISABLE_CHECKSUM]
   --encoding value        バックエンドのエンコーディング。（デフォルト："Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot"） [$ENCODING]
   --endpoint value        IAS3のエンドポイント。（デフォルト："https://s3.us.archive.org"） [$ENDPOINT]
   --front-endpoint value  インターネットアーカイブのフロントエンドのホスト。（デフォルト："https://archive.org"） [$FRONT_ENDPOINT]
   --wait-archive value    サーバーの処理タスク（特にアーカイブとbook_op）の終了を待つタイムアウト。（デフォルト："0s"） [$WAIT_ARCHIVE]
```
{% endcode %}