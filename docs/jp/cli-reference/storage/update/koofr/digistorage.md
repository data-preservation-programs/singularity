# Digi Storage, https://storage.rcs-rds.ro/

{% code fullWidth="true" %}
```
NAME:
   singularity storage update koofr digistorage - Digiのストレージ、https://storage.rcs-rds.ro/

使用法:
   singularity storage update koofr digistorage [コマンドオプション] <名前|ID>

説明:
   --mountid
      使用するマウントのマウントIDです。
      
      省略した場合は、プライマリマウントが使用されます。

   --setmtime
      バックエンドは変更日時の設定をサポートしていますか。
      
      DropboxまたはAmazon Driveバックエンドを指すマウントIDを使用する場合は、これをfalseに設定します。

   --user
      ユーザー名です。

   --password
      rcloneのパスワードです（https://storage.rcs-rds.ro/app/admin/preferences/passwordで生成できます）。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h        ヘルプを表示します
   --password value  rcloneのパスワードです（https://storage.rcs-rds.ro/app/admin/preferences/passwordで生成できます）。 [$PASSWORD]
   --user value      ユーザー名です。 [$USER]

   Advanced

   --encoding value  バックエンドのエンコーディングです。 (デフォルト: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --mountid value   使用するマウントのマウントIDです。 [$MOUNTID]
   --setmtime        バックエンドは変更日時の設定をサポートしていますか。 (デフォルト: true) [$SETMTIME]

```
{% endcode %}