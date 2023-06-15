# File Retrieval

Since singularity database stores the connections to the original data source, it can also act as a storage provider providing file level retrievals. Start by running the content provider

```sh
singularitu run content-provider
```

Retrieve a file using HTTP protocol

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

Retrieve a file using Bitswap protocol

```sh
ipfs daemon
ipfs swarm connect <multi_addr_shown_in_singularity>
ipfs get bafyxxxx
```
