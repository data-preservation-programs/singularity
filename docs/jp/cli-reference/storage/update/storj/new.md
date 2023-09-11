＃サテライトアドレス、APIキー、およびパスフレーズから新しいアクセス許可を作成します。

{% code fullWidth="true" %}
```
NAME:
   singularity storage update storj new - サテライトアドレス、APIキー、およびパスフレーズから新しいアクセス許可を作成します。

使用法:
   singularity storage update storj new [コマンドオプション] <名前|ID>

説明:
   --satellite-address
      サテライトアドレス。

      カスタムサテライトアドレスは次の形式に一致する必要があります：`<ノードID>@<アドレス>：<ポート>`。

      例：
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key
      APIキー。

   --passphrase
      暗号化のパスフレーズ。

      既存のオブジェクトにアクセスするには、アップロードに使用したパスフレーズを入力してください。


オプション：
   --satellite-address value  サテライトアドレス（デフォルト値： "us1.storj.io"）[$SATELLITE_ADDRESS]
   --api-key value            APIキー[$API_KEY]
   --passphrase value         暗号化のパスフレーズ[$PASSPHRASE]
   --help, -h                 ヘルプを表示
```
{% endcode %}