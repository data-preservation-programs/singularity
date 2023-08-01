# ローカルパスからデータセットを準備する

{% code fullWidth="true" %}
```
NAME:
   singularity ez-prep - ローカルパスからデータセットを準備する

使用法:
   singularity ez-prep [コマンドオプション] <path>

カテゴリー:
   簡単なコマンド

説明:
   このコマンドは、最小限の設定可能パラメータを使用してローカルパスからデータセットを準備するために使用できます。
   より高度な使用法については、「dataset」と「datasource」のサブコマンドを使用してください。
   また、このコマンドはインメモリデータベースとインライン準備を使用してベンチマーキングにも利用できます。
     mkdir dataset
     truncate -s 1024G dataset/1T.bin
     singularity ez-prep --output-dir '' --database-file '' -j $(($(nproc) / 4 + 1)) ./dataset

オプション:
   --max-size value, -M value     作成されるCARファイルの最大サイズ (デフォルト: "31.5GiB")
   --output-dir value, -o value   CARファイルの出力ディレクトリ。インライン準備を使用する場合は空の文字列を指定します (デフォルト: "./cars")
   --concurrency value, -j value  パッキングの並列性 (デフォルト: 1)
   --database-file value          メタデータを格納するデータベースファイル。インメモリデータベースを使用する場合は空の文字列を指定します (デフォルト: ./ezprep-<name>.db)
   --help, -h                     ヘルプを表示
```
{% endcode %}