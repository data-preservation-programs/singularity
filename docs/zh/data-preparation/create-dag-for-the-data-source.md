# 为数据源创建DAG

在这种情况下，DAG包含数据源的所有相关文件夹信息以及文件如何分割成多个块。如果此DAG的CAR文件由存储提供商封装，则可以使用单个数据集的根CID查找文件的unixfs路径。

要触发数据源DAG生成过程，请执行以下操作：

```sh
# 假设有一个数据源：
singularity datasource daggen 1
```

现在该任务已记录在数据库中，您需要重新运行数据集worker，或者如果worker已经在运行中，则等待其处理该任务。

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

完成后，您可以检查相关的DAG

```
singularity datasource inspect dags 1
```

DAG的CAR文件将自动包含在交易中

## 下一步

[distribute-car-files.md](../content-distribution/distribute-car-files.md "mention")

## 相关资源

[触发DAG生成](../cli-reference/datasource/daggen.md)

[检查数据源的DAG](../cli-reference/datasource/inspect/dags.md)