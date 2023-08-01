# ファイルの取得（ステージング）

Singularityデータベースは元のデータソースへの接続情報を保持しているため、ファイルレベルの取得も提供するストレージプロバイダとして機能することができます。これにより、Filecoinのストレージプロバイダがサポートする取得プロトコルを介して取得を提供する便利な方法が提供されます。

まず、コンテンツプロバイダを実行します。

```sh
singularitu run content-provider
```

HTTPプロトコルを使用してファイルを取得する

```
wget 127.0.0.1:8088/ipfs/bafyxxxx
```

Bitswapプロトコルを使用してファイルを取得する

```sh
ipfs daemon
ipfs swarm connect <singularityに表示されるマルチアドレス>
ipfs get bafyxxxx
```