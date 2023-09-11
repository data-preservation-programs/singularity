# 他のKoofr API互換のストレージサービス

{% code fullWidth="true" %}
```
NAME:
   singularity storage create koofr other - 他のKoofr API互換のストレージサービス

USAGE:
   singularity storage create koofr other [command options] [arguments...]

DESCRIPTION:
   --endpoint
      使用するKoofr APIのエンドポイント。

   --mountid
      使用するマウントのマウントID。
      
      省略すると、プライマリマウントが使用されます。 

   --setmtime
      バックエンドが変更時刻の設定をサポートしているかどうか。

      DropboxやAmazon Driveのバックエンドを指すマウントIDを使用する場合、これをfalseに設定します。

   --user
      ユーザー名。

   --password
      rcloneのパスワード（サービスの設定ページで生成します）。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

OPTIONS:
   --endpoint value  使用するKoofr APIのエンドポイント。 [$ENDPOINT]
   --help, -h        ヘルプを表示します
   --password value  rcloneのパスワード（サービスの設定ページで生成します）。 [$PASSWORD]
   --user value      ユーザー名。 [$USER]

   Advanced

   --encoding value  バックエンドのエンコーディング（デフォルト: "スラッシュ、バックスラッシュ、削除、制御、無効なUTF8、ドット"） [$ENCODING]
   --mountid value   使用するマウントのマウントID。 [$MOUNTID]
   --setmtime        バックエンドが変更時刻の設定をサポートしているかどうか（デフォルト: true） [$SETMTIME]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}