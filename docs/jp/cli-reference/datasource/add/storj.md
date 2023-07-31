# Storj分散型クラウドストレージ

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add storj - Storj分散型クラウドストレージ

USAGE:
   singularity datasource add storj [command options] <dataset_name> <source_path>

DESCRIPTION:
   --storj-access-grant
      [プロバイダ] - 存在する
         アクセス権限。

   --storj-api-key
      [プロバイダ] - 新しい
         APIキー。

   --storj-passphrase
      [プロバイダ] - 新しい
         暗号化パスフレーズ。

         既存のオブジェクトにアクセスする場合は、アップロード時に使用したパスフレーズを入力してください。

   --storj-provider
      認証方法を選択します。

      例:
         | existing | 既存のアクセス権限を使用します。
         | new      | サテライトアドレス、APIキー、およびパスフレーズから新しいアクセス権限を作成します。

   --storj-satellite-address
      [プロバイダ] - 新しい
         サテライトアドレス。

         カスタムのサテライトアドレスは、次のフォーマットに一致する必要があります: `<nodeid>@<address>:<port>`。

         例:
            | us1.storj.io | US1
            | eu1.storj.io | EU1
            | ap1.storj.io | AP1


OPTIONS:
   --help, -h  ヘルプを表示します

   データの準備オプション

   --delete-after-export    [注意] データセットをCARファイルにエクスポート後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過したときに、ソースディレクトリを自動的に再スキャンします。(デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。(デフォルト: ready)

   storjのオプション

   --storj-access-grant value       アクセス権限。[$STORJ_ACCESS_GRANT]
   --storj-api-key value            APIキー。[$STORJ_API_KEY]
   --storj-passphrase value         暗号化パスフレーズ。[$STORJ_PASSPHRASE]
   --storj-provider value           認証方法を選択します。(デフォルト: "existing") [$STORJ_PROVIDER]
   --storj-satellite-address value  サテライトアドレス。(デフォルト: "us1.storj.io")[$STORJ_SATELLITE_ADDRESS]
```
{% endcode %}