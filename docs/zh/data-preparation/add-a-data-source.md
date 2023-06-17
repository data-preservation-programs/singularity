---
description: 连接需要准备的数据源
---

# 添加数据源

## 添加本地文件系统数据源

最常见的数据源是本地文件系统。要将文件夹添加为数据源到数据集中：

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

## 添加公共的S3数据源

为了演示如何将S3数据源添加到数据集中，让我们使用名为[Foldingathome COVID-19 Datasets](https://registry.opendata.aws/foldingathome-covid19/)的公共数据集

```
singularity datasource add s3 my_dataset fah-public-data-covid19-cryptic-pocketst 
```

## 下一步

[开始数据集工作](start-dataset-worker.md)

## 相关资源

[所有数据源类型](../cli-reference/datasource/add/)