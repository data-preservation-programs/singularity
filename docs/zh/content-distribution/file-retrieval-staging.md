# 文件提取（分期）

由于 Singularity 数据库存储了与原始数据源的连接，因此它也可以作为存储提供者提供文件级别的检索。 

请注意，这不是替代存储提供者的服务，因为您作为客户端已经拥有原始数据源，但它提供了一种方便的方式来通过 Filecoin 存储提供者支持的检索协议提供检索。

首先启动内容提供者

```sh
singularitu run content-provider
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
