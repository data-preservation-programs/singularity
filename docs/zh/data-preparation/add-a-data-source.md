---
description: 连接到需要准备的数据源
---

# 添加数据源

## 添加本地文件系统数据源

最常见的数据源是本地文件系统。要将文件夹添加为数据源到数据集中，请执行以下操作：

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

## 添加公共 S3 数据源

为了演示如何将 S3 数据源添加到数据集中，让我们使用一个名为 [Foldingathome COVID-19 Datasets](https://registry.opendata.aws/foldingathome-covid19/) 的公共数据集。

```
singularity datasource add s3 my_dataset fah-public-data-covid19-cryptic-pocketst 
```

## 下一步

[start-dataset-worker.md](start-dataset-worker.md "mention")

## 相关资源

[所有数据源类型](../cli-reference/datasource/add/)