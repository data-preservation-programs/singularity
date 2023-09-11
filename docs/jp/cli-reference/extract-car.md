# フォルダーまたはファイルをCARファイルのフォルダーからローカルディレクトリに抽出する

{% code fullWidth="true" %}
```
NAME:
   singularity extract-car - フォルダーまたはファイルをCARファイルのフォルダーからローカルディレクトリに抽出します

使用法:
   singularity extract-car [コマンドオプション] [引数...]

カテゴリー:
   ユーティリティ

オプション:
   --input-dir value, -i value  CARファイルを含む入力ディレクトリ。このディレクトリは再帰的にスキャンされます
   --output value, -o value     抽出先の出力ディレクトリまたはファイル。存在しない場合は作成されます (デフォルト値: ".")
   --cid value, -c value        抽出するフォルダーまたはファイルのCID
   --help, -h                   ヘルプを表示します
```
{% endcode %}