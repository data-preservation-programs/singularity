# データセットワーカーを起動する

データセットを準備するためにデータセットワーカーを起動するには、次のコマンドを実行します。

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

デフォルトでは、データセットのスキャン、パック、およびダグナイジングを行うシングルワーカースレッドが作成されます。処理は完了するかエラーが発生すると終了します。本番環境では、ワーカーを継続的に実行する必要があります。

フラグ`--concurrency value`を使用していくつかの同時実行値を設定することもできます。

準備が完了した後、次のコマンドを使用して準備されたデータを検査することができます。

```sh
# 追加されたすべてのデータソースの一覧を表示する
singularity datasource list

# スキャンおよびパックの結果を一覧表示する
singularity datasource status 1

# ルートフォルダの各ファイルのCIDを確認する
singularity datasource inspect dir 1

# 生成されたすべてのCARファイルを確認する
singularity datasource inspect chunks

# 準備されたすべてのアイテムを確認する
singularity datasource inspect items
```

## 次のステップ

[データソースのためのDAGを作成する](create-dag-for-the-data-source.md "mention")

## 関連するリソース

[すべてのデータソースをリストアップする](../cli-reference/datasource/list.md)

[データソースの準備状況を確認する](../cli-reference/datasource/status.md)

[データソースのすべてのアイテムを確認する](../cli-reference/datasource/inspect/items.md)

[データソースのすべてのチャンクを確認する](../cli-reference/datasource/inspect/chunks.md)

[データソースのディレクトリを確認する](../cli-reference/datasource/inspect/dir.md)