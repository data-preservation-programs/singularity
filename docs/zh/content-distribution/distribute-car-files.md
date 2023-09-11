# 使用 Singularity 分发 CAR 文件

为确保存储供应商可以轻松访问您的数据，您需要有效地分发 CAR（内容可寻址存档）文件。

## 1. 启动内容提供商服务

首先启动内容提供商服务。此服务有助于从您准备的数据集中下载数据片段。

```sh
singularity run content-provider
```

## 2. CAR 文件下载方法

存储供应商可以使用多种方法下载 CAR 文件：

### 直接 HTTP 下载

供应商可以利用内容提供商服务提供的 HTTP API 直接下载 CAR 文件：

```shell
wget http://127.0.0.1:7777/piece/bagaxxxxxxxxxxx
```
如果在准备过程中指定了输出目录，则 CAR 文件将直接从该目录提取。但是，如果您使用了内联准备或意外删除了 CAR 文件，则该服务将从原始数据源检索内容并提供服务。

### Singularity 下载实用程序

对于寻求备选下载方法的供应商，尤其是在处理像 S3 或 FTP 这样的远程数据源时，Singularity 提供了专用的下载实用程序：

```shell
singularity download bagaxxxxxxxxxxx
```
该实用程序与内容提供商服务通信，获取有关数据片段的元数据。获取到元数据后，它使用该元数据直接从原始数据源重构数据片段。