# データセットのスキャンと準備タスクを処理するためのデータセット準備ワーカーを開始する

{% code fullWidth="true" %}
```
NAME:
   singularity run dataset-worker - データセットのスキャンと準備タスクを処理するためのデータセット準備ワーカーを開始する

使用方法:
   singularity run dataset-worker [コマンドオプション] [引数...]

オプション:
   --concurrency value  同時実行するワーカーの数（デフォルト：1）[$DATASET_WORKER_CONCURRENCY]
   --enable-scan        データセットのスキャンを有効にする（デフォルト：true）[$DATASET_WORKER_ENABLE_SCAN]
   --enable-pack        CIDsを計算し、データセットをCARファイルにパックするためのパックを有効にする（デフォルト：true）[$DATASET_WORKER_ENABLE_PACK]
   --enable-dag         データセットのディレクトリ構造を維持するためのDAGの生成を有効にする（デフォルト：true）[$DATASET_WORKER_ENABLE_DAG]
   --exit-on-complete   もう作業がない場合にワーカーを終了する（デフォルト：false）[$DATASET_WORKER_EXIT_ON_COMPLETE]
   --exit-on-error      何らかのエラーが発生した場合にワーカーを終了する（デフォルト：false）[$DATASET_WORKER_EXIT_ON_ERROR]
   --help, -h           ヘルプを表示する
```
{% endcode %}