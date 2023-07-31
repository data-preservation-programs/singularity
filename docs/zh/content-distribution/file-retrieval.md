# 文件检索

由于Singularity数据库存储了与原始数据源的连接，它还可以作为存储提供者，提供文件级的检索功能。首先运行内容提供者：

```sh
singularitu run content-provider
```

使用HTTP协议检索文件：

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

使用Bitswap协议检索文件：

```sh
ipfs daemon
ipfs swarm connect <singularity中显示的multi_addr>
ipfs get bafyxxxx
```