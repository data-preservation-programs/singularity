# 1Fichier

{% code fullWidth="true" %}
```
名前:
   singularity storage create fichier - 1Fichier

使用法:
   singularity storage create fichier [コマンドオプション] [引数...]

説明:
   --api-key
      APIキー、取得方法は https://1fichier.com/console/params.pl を参照してください。

   --shared-folder
      共有フォルダをダウンロードする場合、このパラメータを追加してください。

   --file-password
      パスワードで保護された共有ファイルをダウンロードする場合、このパラメータを追加してください。

   --folder-password
      パスワードで保護された共有フォルダのファイルをリストする場合、このパラメータを追加してください。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細は [概要のエンコーディングセクション](/overview/#encoding) を参照してください。


オプション:
   --api-key value  APIキー、取得方法は https://1fichier.com/console/params.pl を参照してください。 [$API_KEY]
   --help, -h       ヘルプを表示する

   上級者向け

   --encoding value         バックエンドのエンコーディングです。 (デフォルト: "Slash,LtGt,DoubleQuote,SingleQuote,BackQuote,Dollar,BackSlash,Del,Ctl,LeftSpace,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --file-password value    パスワードで保護された共有ファイルをダウンロードする場合、このパラメータを追加してください。 [$FILE_PASSWORD]
   --folder-password value  パスワードで保護された共有フォルダのファイルをリストする場合、このパラメータを追加してください。 [$FOLDER_PASSWORD]
   --shared-folder value    共有フォルダをダウンロードする場合、このパラメータを追加してください。 [$SHARED_FOLDER]

   一般

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}