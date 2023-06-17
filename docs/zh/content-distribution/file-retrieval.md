# 文件检索

由于 Singularity 数据库存储了与原始数据源的连接，因此它也可以作为存储提供者提供文件级别的检索。首先运行内容提供者。

```sh
singularity run content-provider
```

使用 HTTP 协议检索文件：

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

使用 BitSwap 协议检索文件：

```sh
ipfs daemon
ipfs swarm connect <multi_addr_shown_in_singularity>
ipfs get bafyxxxx
```
注意：不要翻译代码块中的命令和参数用法，只翻译注释或描述文本。