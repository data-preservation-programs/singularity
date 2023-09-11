# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
NAME:
   singularity storage create koofr digistorage - Digi Storage、https://storage.rcs-rds.ro/

USAGE:
   singularity storage create koofr digistorage [command options] [arguments...]

DESCRIPTION:
   --mountid
      使用するマウントのマウントID。
      
      省略すると、プライマリマウントが使用されます。

   --setmtime
      バックエンドは修正日時の設定をサポートしていますか。
      
      DropboxやAmazon Driveバックエンドを指すマウントIDを使用している場合、「false」に設定します。

   --user
      ユーザー名。

   --password
      rcloneのパスワード（[ここで生成してください](https://storage.rcs-rds.ro/app/admin/preferences/password)）。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h        ヘルプを表示します
   --password value  rcloneのパスワード（[ここで生成してください](https://storage.rcs-rds.ro/app/admin/preferences/password)）。 [$PASSWORD]
   --user value      ユーザー名。 [$USER]

   Advanced

   --encoding value  バックエンドのエンコーディング。 (デフォルト値: "スラッシュ、バックスラッシュ、削除、制御、無効なUTF8、ドット") [$ENCODING]
   --mountid value   使用するマウントのマウントID。 [$MOUNTID]
   --setmtime        バックエンドは修正日時の設定をサポートしていますか。 (デフォルト値: true) [$SETMTIME]

   General

   --name value  ストレージの名前（デフォルト値: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}