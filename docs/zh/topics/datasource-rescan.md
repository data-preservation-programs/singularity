# 数据源重新扫描

默认情况下，数据源仅被扫描一次，但 Singularity 提供了一个选项，可以在一定时间间隔后重新扫描数据源。

```sh
singularity datasource add <type> --rescan-interval value
```

您也可以手动触发重新扫描。

```sh
singularity datasource rescan datasource_id
```

在重新扫描期间，所有新文件都将排队准备。所有删除的文件都将被忽略。

#### 文件版本控制

对于已更改的文件，将排队准备文件的新版本，并更新目录 CID 以使用文件的最新版本。

文件是否已更改的逻辑由以下步骤确定：

1. 如果数据源提供文件的哈希值（例如 Etag），则如果哈希值已更改，将创建新版本
2. 否则，如果数据源提供文件的最后修改时间，则如果该值已更改或文件大小已更改，则将创建新版本
3. 否则，使用文件大小来确定是否应创建新版本

如果在重新扫描之间多次覆盖相同的文件，则仍然可能会错过某些文件版本。为确保捕获所有文件版本，用户应使用 [push-and-upload.md](push-and-upload.md "mention")，让 singularity 知道每次更新文件的时间。