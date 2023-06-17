# 基准测试

EZ preparation命令提供了一种简单的方法来执行基准测试。

## 准备测试数据

首先，为基准测试准备一些数据。 我们使用了稀疏文件来避免考虑磁盘IO时间。 我们目前没有进行CID重复数据删除，因此Singularity处理它们时会将其视为随机字节。

```sh
mkdir dataset
truncate -s 1024G dataset/1T.bin
```

如果要将磁盘IO时间作为基准测试的一部分，您可以使用自己喜欢的方法创建一个随机文件，例如：

```
dd if=/dev/urandom of=dataset/8G.bin bs=1M count=8192
```

## 运行ez-prep

EZ prep命令是一个简单的命令，运行了几个内部命令，以非常少的可自定义设置准备本地文件夹。

#### 在线准备基准测试

Inline准备消除了导出CAR文件并将必要的元数据保存在数据库中的需要

```sh
time singularity ez-prep --output-dir '' ./dataset
```

#### 使用内存中的数据库进行基准测试

为了进一步减少磁盘IO，您还可以选择使用内存中的数据库

```sh
time singularity ez-prep --output-dir '' --database-file '' ./dataset
```

#### 使用多个工作线程进行基准测试

为了利用所有CPU核心，可以为基准测试设置并发标志。 请注意，每个工作进程需要约4个CPU核心，因此您需要正确设置它。

```sh
time singularity ez-prep --output-dir '' -j $(($(nproc) / 4 + 1)) ./dataset
```

## 解释结果

您将看到如下所示的内容

```
real    0m20.379s
user    0m44.937s
sys     0m8.981s
```

`real`表示实际的时钟时间。 使用更多的工作进程并行将可能减少此数字。

`user`表示用户空间消耗的CPU时间。 如果您将`user`除以`real`，它大约是程序使用了多少CPU核心。 使用更多并发可能不会对此数字产生重大影响，因为需要完成的工作不会改变。

`sys`表示在内核空间中用于磁盘IO的CPU时间。

## 对比

以下测试是在随机8G文件上执行的

<table><thead><tr><th width="290">工具</th><th width="178.33333333333331" data-type="number">时钟时间 (秒)</th><th data-type="number">CPU时间 (秒)</th><th data-type="number">内存 (KB)</th></tr></thead><tbody><tr><td>Singularity w/ inline prep</td><td>15.66</td><td>51.82</td><td>99</td></tr><tr><td>Singularity w/o inline prep</td><td>19.13</td><td>51.51</td><td>99</td></tr><tr><td>go-fil-dataprep</td><td>16.39</td><td>43.94</td><td>83</td></tr><tr><td>generate-car</td><td>42.6</td><td>56.08</td><td>44</td></tr><tr><td>go-car + stream-commp</td><td>70.21</td><td>139.01</td><td>42</td></tr></tbody></table>