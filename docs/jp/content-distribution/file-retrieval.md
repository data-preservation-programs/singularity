# ファイルの取得

Singularityデータベースは、元のデータソースへの接続情報を格納しているため、ファイルレベルの取得を提供するストレージプロバイダとしても機能します。まず、コンテンツプロバイダを実行してください。

```sh
singularitu run content-provider
```

HTTPプロトコルを使用してファイルを取得する場合は、

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

BitSwapプロトコルを使用してファイルを取得する場合は、

```sh
ipfs daemon
ipfs swarm connect <singularityで表示されたマルチアドレス>
ipfs get bafyxxxx
```