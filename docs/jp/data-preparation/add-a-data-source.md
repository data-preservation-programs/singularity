---
description: データの準備が必要なデータソースに接続します
---

# データソースの追加

## ローカルファイルシステムデータソースの追加

最も一般的なデータソースはローカルファイルシステムです。データセットにフォルダをデータソースとして追加するには、次のコマンドを実行してください：

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

## 公開S3データソースの追加

データセットにS3データソースを追加する方法を示すために、[Foldingathome COVID-19 Datasets](https://registry.opendata.aws/foldingathome-covid19/)という公開データセットを使用します。

```
singularity datasource add s3 my_dataset fah-public-data-covid19-cryptic-pocketst
```

## 次のステップ

[start-dataset-worker.md](start-dataset-worker.md "mention")

## 関連リソース

[すべてのデータソースのタイプ](../cli-reference/datasource/add/)