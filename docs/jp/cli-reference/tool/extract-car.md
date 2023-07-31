# フォルダ内のCARファイルからフォルダやファイルをローカルディレクトリに抽出する

{% code fullWidth="true" %}
```
NAME:
   singularity tool extract-car - フォルダ内のCARファイルからフォルダやファイルをローカルディレクトリに抽出する

使用法:
   singularity tool extract-car [コマンドオプション] [引数...]

オプション:
   --input-dir value, -i value  CARファイルを含むディレクトリを指定します。このディレクトリは再帰的にスキャンされます
   --output value, -o value     抽出先のディレクトリまたはファイルを指定します。存在しない場合は作成されます (デフォルト: ".")
   --cid value, -c value        抽出するフォルダやファイルのCIDを指定します
   --help, -h                   ヘルプを表示する
```
{% endcode %}