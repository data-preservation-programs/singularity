# Koofr, https://app.koofr.net/

{% code fullWidth="true" %}
```
NAME:
   singularity storage update koofr koofr - Koofr, https://app.koofr.net/

USAGE:
   singularity storage update koofr koofr [コマンドオプション] <名前|ID>

DESCRIPTION:
   --mountid
      使用するマウントのマウントIDです。
      
      省略された場合、プライマリマウントが使用されます。

   --setmtime
      バックエンドは変更日時の設定をサポートしていますか。
      
      DropboxやAmazon Driveのバックエンドを指すマウントIDを使用している場合は、これをfalseに設定してください。

   --user
      ユーザー名です。

   --password
      rclone用のパスワードです（https://app.koofr.net/app/admin/preferences/passwordで生成できます）。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --help, -h        ヘルプを表示します
   --password value  rclone用のパスワードです（https://app.koofr.net/app/admin/preferences/passwordで生成できます）。 [$PASSWORD]
   --user value      ユーザー名です。 [$USER]

   Advanced

   --encoding value  バックエンドのエンコーディングです（デフォルト: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"） [$ENCODING]
   --mountid value   使用するマウントのマウントIDです。 [$MOUNTID]
   --setmtime        バックエンドは変更日時の設定をサポートしていますか（デフォルト: true） [$SETMTIME]

```
{% endcode %}