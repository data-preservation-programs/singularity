# 启动数据集工作器

要启动一个数据集工作器来准备数据集，请运行以下命令

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

默认情况下，它将生成一个单线程的工作器，用于扫描、打包和探查数据集。该过程将在完成或出现任何错误时退出。在生产环境中，您希望它保持运行。

您还可以使用标志 `--concurrency value` 配置一些并发值。

准备完成后，您可以使用以下某些命令检查准备好的数据：

```sh
# 列出已添加的所有数据源
singularity datasource list

# 给出扫描和打包结果的概述
singularity datasource status 1

# 检查根文件夹中每个文件的CID
singularity datasource inspect dir 1

# 检查所有生成的CAR文件
singularity datasource inspect chunks

# 检查准备好的所有项目
singularity datasource inspect items
```

## 下一步

[创建数据源的DAG](create-dag-for-the-data-source.md "mention")

## 相关资源

[列出所有数据源](../cli-reference/datasource/list.md)

[检查数据源准备状态](../cli-reference/datasource/status.md)

[检查数据源的所有项目](../cli-reference/datasource/inspect/items.md)

[检查数据源的所有块](../cli-reference/datasource/inspect/chunks.md)

[检查数据源的目录](../cli-reference/datasource/inspect/dir.md)