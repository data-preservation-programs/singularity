# Koofr、https://app.koofr.net/

{% code fullWidth="true" %}
```
NAME:
   singularity storage create koofr koofr - Koofr、https://app.koofr.net/

使用法:
   singularity storage create koofr koofr [コマンドオプション] [引数...]

説明:
   --mountid
      使用するマウントのマウントID。
      
      省略すると、プライマリマウントが使用されます。

   --setmtime
      バックエンドでの変更時刻の設定をサポートしているかどうか。
      
      DropboxやAmazon Driveのバックエンドを指すマウントIDを使用する場合は、これをfalseに設定します。

   --user
      ユーザー名。

   --password
      rcloneのパスワード（https://app.koofr.net/app/admin/preferences/passwordで生成）。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h        ヘルプを表示
   --password value  rcloneのパスワード（https://app.koofr.net/app/admin/preferences/passwordで生成）。 [$PASSWORD]
   --user value      ユーザー名。 [$USER]

   高度な設定

   --encoding value  バックエンドのエンコーディング（デフォルト："Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot"）。 [$ENCODING]
   --mountid value   使用するマウントのマウントID。 [$MOUNTID]
   --setmtime        バックエンドでの変更時刻の設定をサポートしているかどうか（デフォルト：true）。 [$SETMTIME]

   一般設定

   --name value  ストレージの名前（デフォルト：自動生成）
   --path value  ストレージのパス

```
{% endcode %}