# 1Fichier

```
名前:
   singularity storage update fichier - 1Fichier

使用法:
   singularity storage update fichier [コマンドオプション] <名前|ID>

説明:
   --api-key
      APIキーは、https://1fichier.com/console/params.pl から取得してください。

   --shared-folder
      共有フォルダをダウンロードする場合は、このパラメータを追加してください。

   --file-password
      パスワードで保護された共有ファイルをダウンロードする場合は、このパラメータを追加してください。

   --folder-password
      パスワードで保護された共有フォルダ内のファイルをリストする場合は、このパラメータを追加してください。

   --encoding
      バックエンドのエンコーディング。
      
      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --api-key value  APIキーは、https://1fichier.com/console/params.pl から取得してください。 [$API_KEY]
   --help, -h       ヘルプを表示

   高度なオプション:

   --encoding value         バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    パスワードで保護された共有ファイルをダウンロードする場合は、このパラメータを追加してください。 [$FILE_PASSWORD]
   --folder-password value  パスワードで保護された共有フォルダ内のファイルをリストする場合は、このパラメータを追加してください。 [$FOLDER_PASSWORD]
   --shared-folder value    共有フォルダをダウンロードする場合は、このパラメータを追加してください。 [$SHARED_FOLDER]

```
