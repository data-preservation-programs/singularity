# 基准测试

EZ准备命令提供了一种简单的方式来执行基准测试。

## 准备测试数据

首先，准备一些用于基准测试的数据。我们使用稀疏文件来避免考虑磁盘IO时间。截至目前，我们不执行CID去重，因此Singularity将其视为随机字节。

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

如果您想将磁盘IO时间作为基准的一部分，可以使用您自己喜欢的方式创建一个随机文件，例如：

```
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## 运行ez-prep

EZ预处理命令是一个简单的命令，它运行一些内部命令来准备一个具有非常少定制设置的本地文件夹。

#### 在线预处理基准测试

在线预处理消除了导出CAR文件并在数据库中保存必要元数据的需求。

```sh
time singularity ez-prep --output-dir '' ./dataset
```

#### 使用内存数据库进行基准测试

为了进一步减少磁盘IO，您还可以选择使用内存数据库。

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

#### 使用多个工作线程进行基准测试

为了充分利用所有CPU核心，您可以为基准测试设置并发标志。请注意，每个工作线程占用大约4个CPU核心，因此您需要正确设置。

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 解释结果

您会看到如下所示的内容

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

`real`表示实际的钟表时间。增加更多的并发工作线程可能会减少此数字。

`user`表示在用户空间上花费的CPU时间。如果将`user`除以`real`，结果大致上表示程序使用了多少个CPU核心。增加更多的并发可能不会对这个数字产生显著影响，因为需要完成的工作并没有改变。

`sys`表示在内核空间中花费的CPU时间，即磁盘IO时间

## 对比

以下测试是在一个随机的8G文件上进行的

<table><thead><tr><th width="290">工具</th><th width="178.33333333333331" data-type="number">钟表时间（秒）</th><th data-type="number">CPU时间（秒）</th><th data-type="number">内存（KB）</th></tr></thead><tbody><tr><td>带内联预处理的Singularity</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>不带内联预处理的Singularity</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>