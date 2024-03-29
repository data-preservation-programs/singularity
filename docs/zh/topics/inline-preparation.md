# 在线数据准备

## 概述

传统的数据准备方法是将原始数据源（通常是本地文件系统上的一个文件夹）转换为小于32GB的CAR文件集合。这种方法要求数据准备者拥有两倍的存储容量，这可能非常昂贵。例如，准备1 PB的数据集还需要另外1 PB的存储空间来存储CAR文件，总共需要2 PB的存储空间。

通过在线数据准备，可以将CAR文件的块映射回原始数据源，因此不需要存储导出的CAR文件。

<div align="center">

<img src="https://github.com/data-preservation-programs/singularity/assets/12418265/4292faf1-9f01-4b7c-b79f-67b0bc1e2acc" alt="传统数据准备图示" width="500">

 

<img src="https://github.com/data-preservation-programs/singularity/assets/12418265/f5cfc209-5e38-4bb9-8cd9-f1aeffaf284d" alt="在线数据准备图示" width="500">

</div>

## CAR获取的工作原理

通过在线数据准备，可以使用元数据数据库和原始数据源通过HTTP提供CAR文件，因为它知道如何将CAR文件的字节范围映射回原始数据源。

要通过HTTP提供CAR文件，只需启动内容提供程序

```sh
singularity run content-provider
```

> 注意：该命令将运行一个本地的HTTP服务器。如果你打算使其通过互联网访问，你可能希望将其放在一个反向代理后面，例如nginx。

如果数据源已经是一个远程存储系统（即S3或FTP），那么这将创建一个潜在的瓶颈，因为文件内容将通过Singularity内容提供程序代理到存储提供程序。

我们有一个解决这个挑战的方法，使用Singularity元数据API和Singularity下载工具。

要运行Singularity元数据API：

```sh
singularity run api
```

然后，使用Singularity下载工具（在存储提供程序上）：

```sh
singularity download <piece_cid>
```

Singularity元数据API将返回一个从原始数据源组装CAR文件的计划，Singularity下载工具将解释这个计划，并将数据从原始数据流式传输到本地CAR文件中。中间不会进行任何转换或组装，一切都以流的形式工作。

元数据API不会返回访问原始数据源所需的任何凭据。存储提供程序需要获取他们自己的访问数据源的凭据，并将这些凭据提供给`singularity download`命令。

## 额外开销

在线数据准备引入了一些极小的开销，主要是存储空间的需求。此外，计算和带宽开销也非常小。

每个数据块的元数据以数据库行的形式存储，对于每1 MiB的准备数据，需要100字节来存储映射元数据。对于1 PiB的数据集，需要10 TiB的磁盘空间来存储映射元数据。虽然这通常不是问题，但是具有大量小文件的数据集可能会导致显著的磁盘开销。

当需要从原始数据源动态重新生成CAR文件时，需要在数据库中交叉引用这些映射。然而，这通常不是一个问题。1 GB/sec的带宽等于1,000次数据库条目查找，远远没有所有支持的数据库后端的能力极限。此外，未来的优化可能进一步减少此开销。

## 启用在线数据准备

对于不需要加密的数据集，将自动启用在线数据准备。在创建数据集时，指定输出目录后，CAR文件将导出到该位置。CAR文件检索请求将优先考虑这些目录。如果用户删除了CAR文件，系统将回退到从原始数据源获取。