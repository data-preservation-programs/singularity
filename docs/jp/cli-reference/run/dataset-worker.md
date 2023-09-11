# データセット準備ワーカーを起動して、データセットのスキャンと準備タスクを処理します

{% code fullWidth="true" %}
```
NAME:
   singularity run dataset-worker - データセット準備ワーカーを起動して、データセットのスキャンと準備タスクを処理します

使用法:
   singularity run dataset-worker [コマンドオプション] [引数...]

オプション:
   --concurrency value  同時実行するワーカーの数 (デフォルト: 1)
   --enable-scan        データセットのスキャンを有効にする (デフォルト: true)
   --enable-pack        CIDsを計算し、CARファイルにパックしてデータセットをパックするのを有効にする (デフォルト: true)
   --enable-dag         データセットのディレクトリ構造を保持するダグの生成を有効にする (デフォルト: true)
   --exit-on-complete   作業がなくなった時にワーカーを終了する (デフォルト: false)
   --exit-on-error      エラーが発生した場合にワーカーを終了する (デフォルト: false)
   --help, -h           ヘルプを表示
```
{% endcode %}