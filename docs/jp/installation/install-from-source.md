# ソースからのインストール

Singularityをインストールするためには、[Go 1.19以上の動作するインストール](https://golang.org/dl/)が必要です。

```sh
wget -c https://golang.org/dl/go1.19.7.linux-amd64.tar.gz -O - \
    | sudo tar -xz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
    >> ~/.bashrc && source ~/.bashrc
```

その後、以下のコマンドを使用してSingularityをインストールできます。

```sh
go install github.com/data-preservation-programs/singularity@latest
```

もしCコンパイラが動作しない場合は、[CGO](https://zchee.github.io/golang-wiki/cgo/)サポートなしでビルドすることもできます。心配しないでください、CGOの依存性は僅かに高速なSQLiteバックエンドにのみ影響するものであり、本番のワークロードで使用する必要はありません。

```bash
CGO_ENABLED=0 go install github.com/data-preservation-programs/singularity@latest
```