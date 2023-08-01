# データソースのためのDAGを作成する

このコンテキストでのDAGには、データソースに関するすべての関連フォルダ情報と、ファイルが複数のチャンクに分割される方法が含まれています。もしDAGのためのCARファイルがストレージプロバイダによってシールされていれば、データセットの単一のRoot CIDを使用してunixfsパスでファイルを参照することができます。

データソースのDAG生成プロセスをトリガーするには

```sh
# 単一のデータソースがあると仮定する
singularity datasource daggen 1
```

ジョブがデータベースに記録されたら、データセットワーカーを再実行するか、ワーカーが既に実行中の場合はジョブを処理するまで待機する必要があります。

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

完了したら、関連するDAGを確認できます。

```
singularity datasource inspect dags 1
```

DAGのためのCARファイルは、自動的にディールメイキングのために含まれます。

## 次のステップ

[distribute-car-files.md](../content-distribution/distribute-car-files.md "記載")

## 関連するリソース

[DAG生成をトリガーする](../cli-reference/datasource/daggen.md)

[データソースのDAGを確認する](../cli-reference/datasource/inspect/dags.md)