# File Retrieval (Staging)

Since singularity database stores the connections to the original data source, it can also act as a storage provider providing file level retrievals.&#x20;

Note this is not to replace the storage provider because yourself as the client already owns the original data source but it offers a convenient way to offer retrievals over retrieval protocols supported by Filecoin storage providers.

Start by running the content provider

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
