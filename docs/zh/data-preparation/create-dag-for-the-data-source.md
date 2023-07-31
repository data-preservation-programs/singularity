# 创建数据源的DAG

在这个上下文中，DAG包含与数据源相关的所有文件夹信息，以及文件如何分割成多个块。如果此DAG的CAR文件由存储提供者保密，您将能够使用数据集的单个Root CID在unixfs路径下查找文件。

要触发数据源的DAG生成过程，请执行以下操作：

```sh
# 假设有一个单一的数据源
singularity datasource daggen 1
```

现在作业已经记录在数据库中，您需要重新运行数据集工作进程，或者如果工作进程已经在运行中，则等待其接收作业。

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

完成后，您可以检查相关的DAG。

```
singularity datasource inspect dags 1
```

DAG的CAR文件将自动包含用于交易。

## 下一步

[分发CAR文件](../content-distribution/distribute-car-files.md)

## 相关资源

[触发DAG生成](../cli-reference/datasource/daggen.md)

[检查数据源的DAG](../cli-reference/datasource/inspect/dags.md)