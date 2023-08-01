# 1Fichier

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add fichier - 1Fichier

USAGE:
   singularity datasource add fichier [コマンドオプション] <データセット名> <ソースパス>

DESCRIPTION:
   --fichier-api-key
      APIキーは、https://1fichier.com/console/params.pl から取得できます。

   --fichier-encoding
      バックエンドのエンコーディングです。
      
      詳細については[概要の「エンコーディング」セクション](/overview/#encoding)を参照してください。

   --fichier-file-password
      共有されたパスワードで保護されたファイルをダウンロードする場合、このパラメータを追加します。

   --fichier-folder-password
      共有されたパスワードで保護されたフォルダ内のファイルをリストする場合、このパラメータを追加します。

   --fichier-shared-folder
      共有フォルダをダウンロードする場合、このパラメータを追加します。


OPTIONS:
   --help, -h  ヘルプを表示

   データの準備オプション

   --delete-after-export    [危険] データセットをCARファイルにエクスポートした後、ファイルを削除します。  (デフォルト: false)
   --rescan-interval value  前回のスキャンからこの時間が経過した場合、ソースディレクトリを自動的に再スキャンします。 (デフォルト: 無効)
   --scanning-state value   初期のスキャン状態を設定します。 (デフォルト: ready)

   fichier用オプション

   --fichier-api-key value          APIキーは、https://1fichier.com/console/params.pl から取得できます。 [$FICHIER_API_KEY]
   --fichier-encoding value         バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$FICHIER_ENCODING]
   --fichier-file-password value    共有されたパスワードで保護されたファイルをダウンロードする場合、このパラメータを追加します。 [$FICHIER_FILE_PASSWORD]
   --fichier-folder-password value  共有されたパスワードで保護されたフォルダ内のファイルをリストする場合、このパラメータを追加します。 [$FICHIER_FOLDER_PASSWORD]
   --fichier-shared-folder value    共有フォルダをダウンロードする場合、このパラメータを追加します。 [$FICHIER_SHARED_FOLDER]

```
{% endcode %}