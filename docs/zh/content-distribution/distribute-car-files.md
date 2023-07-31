# 分发CAR文件

现在是时候将CAR文件分发给存储提供者，以便他们可以在他们的一侧进行导入。首先，运行内容提供者服务并下载我们准备好的数据集的任何片段：

```sh
singularity run content-provider
wget 127.0.0.1:8088/piece/bagaxxxx
```

如果您之前已指定导出CAR文件的输出目录（禁用内联准备），则该CAR文件将直接从那些CAR文件中提供。否则，如果您一直在使用内联准备或者意外删除了那些CAR文件，它将直接从原始数据源提供。

## 下一步

[deal-making-prerequisite.md](../deal-making/deal-making-prerequisite.md "mention")