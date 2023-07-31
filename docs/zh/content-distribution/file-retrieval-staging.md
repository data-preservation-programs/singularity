# 文件检索（暂存）

由于 Singularity 数据库存储了与原始数据源的连接，因此它也可以作为一个存储提供者，提供文件级别的检索。

请注意，这不能替代存储提供者，因为您作为客户已经拥有原始数据源，但它提供了一种通过 Filecoin 存储提供者支持的检索协议来进行检索的便捷方式。

首先运行内容提供者

```sh
singularity run content-provider
```

使用 HTTP 协议检索文件

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

使用 Bitswap 协议检索文件

```sh
ipfs daemon
ipfs swarm connect <multi_addr_shown_in_singularity>
ipfs get bafyxxxx
```