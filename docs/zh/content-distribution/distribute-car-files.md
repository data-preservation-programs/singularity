# 分发CAR文件

现在是将CAR文件分发给存储提供者的时候了，以便他们可以在他们的一侧导入。 首先启动内容提供者服务，并下载我们准备好的数据集中的任何Pieces：

```sh
singularity run content-provider
wget 127.0.0.1:8088/piece/bagaxxxx
```

如果您之前已经指定了导出CAR的输出目录（这会禁用内联预处理），则该CAR文件将直接从那些CAR文件中提供。否则，如果您一直在使用内联预处理或不小心删除了那些CAR文件，则将直接从原始数据源提供。

## 下一步

[deal-making-prerequisite.md](../deal-making/deal-making-prerequisite.md "mention")