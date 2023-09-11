# 新しいアクセス権を作成する - サテライトアドレス、APIキー、パスフレーズを使用

{% code fullWidth="true" %}
```
NAME（名前）:
   singularity storage create storj new - サテライトアドレス、APIキー、およびパスフレーズから新しいアクセス権を作成します。

USAGE（使用法）:
   singularity storage create storj new [command options] [arguments...]

DESCRIPTION（説明）:
   --satellite-address（サテライトアドレス）
      サテライトアドレスを指定します。

      カスタムサテライトアドレスは次の形式に合わせる必要があります：`<nodeid>@<address>:<port>`。

      例：
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key（APIキー）
      APIキーを指定します。

   --passphrase（パスフレーズ）
      暗号化のためのパスフレーズを指定します。
      
      既存のオブジェクトにアクセスする場合は、アップロード時に使用したパスフレーズを入力してください。


OPTIONS（オプション）:
   --api-key（APIキー） value（値）            APIキーを指定します。 [$API_KEY]
   --help（ヘルプ）、 -h                     ヘルプを表示します
   --passphrase（パスフレーズ） value（値）     暗号化のためのパスフレーズを指定します。 [$PASSPHRASE]
   --satellite-address（サテライトアドレス） value（値）  サテライトアドレスを指定します。（デフォルト値： "us1.storj.io"） [$SATELLITE_ADDRESS]

   一般的なオプション

   --name（名前） value（値）  ストレージの名前（デフォルト：自動生成）
   --path（パス） value（値）  ストレージのパス

```
{% endcode %}