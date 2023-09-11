# その他のKoofr API互換ストレージサービス

{% code fullWidth="true" %}
```
NAME:
   singularity storage update koofr other - その他のKoofr API互換ストレージサービス

USAGE:
   singularity storage update koofr other [コマンドオプション] <名前|ID>

DESCRIPTION:
   --endpoint
      使用するKoofr APIのエンドポイントです。

   --mountid
      使用するマウントのマウントIDです。
      
      省略すると、プライマリマウントが使用されます。

   --setmtime
      バックエンドが変更日時の設定をサポートしているかどうかです。
      
      DropboxやAmazon Driveのバックエンドを指すマウントIDを使用する場合は、これをfalseに設定してください。

   --user
      ユーザー名です。

   --password
      rclone用のパスワードです（サービスの設定ページで生成してください）。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --endpoint value  使用するKoofr APIのエンドポイントです。 [$ENDPOINT]
   --help, -h        ヘルプを表示します
   --password value  rclone用のパスワードです（サービスの設定ページで生成してください）。 [$PASSWORD]
   --user value      ユーザー名です。 [$USER]

   Advanced

   --encoding value  バックエンドのエンコーディングです。 (デフォルト: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   使用するマウントのマウントIDです。 [$MOUNTID]
   --setmtime        バックエンドが変更日時の設定をサポートしているかどうかです。 (デフォルト: true) [$SETMTIME]

```
{% endcode %}