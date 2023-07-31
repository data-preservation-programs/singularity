# Koofr, Digi Storage、およびその他のKoofr互換のストレージプロバイダー

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add koofr - Koofr、Digi Storage、およびその他のKoofr互換のストレージプロバイダー

使用法:
   singularity datasource add koofr [コマンドオプション] <データセット名> <ソースパス>

説明:
   --koofr-encoding
      バックエンドのエンコーディングです。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --koofr-endpoint
      [プロバイダー] - その他
         利用するKoofr APIのエンドポイントです。

   --koofr-mountid
      使用するマウントのマウントIDです。

      省略すると、プライマリマウントが使用されます。

   --koofr-password
      [プロバイダー] - koofr
         rcloneのパスワードです（https://app.koofr.net/app/admin/preferences/passwordで生成できます）。

      [プロバイダー] - digistorage
         rcloneのパスワードです（https://storage.rcs-rds.ro/app/admin/preferences/passwordで生成できます）。

      [プロバイダー] - その他
         rcloneのパスワードです（各サービスの設定ページで生成できます）。

   --koofr-provider
      ストレージプロバイダーを選択します。

      例:
         | koofr       | Koofr, https://app.koofr.net/
         | digistorage | Digi Storage, https://storage.rcs-rds.ro/
         | other       | その他のKoofr API互換のストレージサービス

   --koofr-setmtime
      バックエンドが変更日時の設定をサポートしているかどうかです。

      DropboxやAmazon Driveバックエンドを指すマウントIDを使用する場合は、これをfalseに設定します。

   --koofr-user
      ユーザー名です。


オプション:
   --help, -h  ヘルプを表示します

   データ準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、データセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  前回のスキャンから指定した間隔が経過すると、ソースディレクトリを自動的に再スキャンします (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します (デフォルト: ready)

   koofr用オプション

   --koofr-encoding value  バックエンドのエンコーディングです。 (デフォルト: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$KOOFR_ENCODING]
   --koofr-endpoint value  利用するKoofr APIのエンドポイントです。 [$KOOFR_ENDPOINT]
   --koofr-mountid value   使用するマウントのマウントIDです。 [$KOOFR_MOUNTID]
   --koofr-password value  rcloneのパスワードです（https://app.koofr.net/app/admin/preferences/passwordで生成できます）。 [$KOOFR_PASSWORD]
   --koofr-provider value  ストレージプロバイダーを選択します。 [$KOOFR_PROVIDER]
   --koofr-setmtime value  バックエンドが変更日時の設定をサポートしているかどうかです。 (デフォルト: "true") [$KOOFR_SETMTIME]
   --koofr-user value      ユーザー名です。 [$KOOFR_USER]

```
{% endcode %}