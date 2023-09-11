# ローカルパスからデータセットを準備する

{% code fullWidth="true" %}
```
NAME:
   singularity ez-prep - ローカルパスからデータセットを準備する

USAGE:
   singularity ez-prep [command options] <path>

CATEGORY:
   ユーティリティ

DESCRIPTION:
   このコマンドは、最小限の設定可能なパラメータを使用して、ローカルパスからデータセットを準備するために使用できます。
   より高度な使用法については、`storage` と `data-prep` のサブコマンドを使用してください。
   また、このコマンドは、インメモリデータベースとインライン準備によるベンチマークにも使用できます。
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

OPTIONS:
   --max-size value, -M value       作成されるCARファイルの最大サイズ（デフォルト： "31.5GiB"）
   --output-dir value, -o value     CARファイルの出力ディレクトリ。インライン準備を使用する場合は、空の文字列を使用します（デフォルト： "./cars"）
   --concurrency value, -j value    パッキングの並行性（デフォルト： 1）
   --database-file value, -f value  メタデータを保存するデータベースファイル。インメモリデータベースを使用する場合は、空の文字列を使用します（デフォルト：./ezprep-<name>.db）
   --help, -h                       ヘルプを表示
```
{% endcode %}