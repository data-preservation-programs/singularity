# 使用 Singularity 进行基准测试

Singularity 中的 `ez-prep` 命令提供了一种简化的基准测试方法。

## 准备测试数据

首先，您需要生成用于基准测试的数据。这里使用稀疏文件来消除基准测试中的磁盘 IO 时间。由于 Singularity 目前不执行 CID 去重，因此它将这些文件处理为随机字节。

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

如果您希望在基准测试中包含磁盘 IO 时间，请使用以下方法创建一个随机文件：

```sh
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## 使用 `ez-prep`
`ez-prep` 命令通过最少的可配置选项从本地文件夹中简化数据准备过程。

### 在线准备基准测试数据
在线准备可以避免导出 CAR 文件，直接将元数据保存到数据库中：

```sh
time singularity ez-prep --output-dir '' ./dataset
```

### 使用内存数据库进行基准测试

为了最小化磁盘 IO，可以选择使用内存数据库：

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

### 使用多个工作线程进行基准测试

为了最大限度地利用 CPU 核心，可以设置基准测试的并发性。注意：每个工作线程使用大约 4 个 CPU 核心：

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 解读结果

典型的输出将类似于以下内容：

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

* `real`：实际经过的时间。使用更多的工作线程应该会减少这个时间。
* `user`：在用户空间使用的 CPU 时间。将 `user` 除以 `real` 可以近似计算出使用的 CPU 核心数。
* `sys`：在内核空间使用的 CPU 时间（表示磁盘 IO）。

## 比较

以下基准测试是在一个随机的 8G 文件上进行的：

| 工具                        | 时钟时间（秒） | CPU 时间（秒） | 内存（KB） |
| --------------------------- | -------------- | ------------ | --------- |
| Singularity（内联准备） | 15.66          | 51.82        | 99        |
| Singularity（非内联准备）   | 19.13          | 51.51        | 99        |
| go-fil-dataprep             | 16.39          | 43.94        | 83        |
| generate-car                | 42.6           | 56.08        | 44        |
| go-car + stream-commp       | 70.21          | 139.01       | 42        |