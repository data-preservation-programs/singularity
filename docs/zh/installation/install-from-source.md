# 从源代码安装

要安装 Singularity，您需要先安装一个可用的 [Go 1.19 及以上版本](https://golang.org/dl/)

```sh
wget -c https://golang.org/dl/go1.19.7.linux-amd64.tar.gz -O - \
    | sudo tar -xz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
    >> ~/.bashrc && source ~/.bashrc
```

然后您可以使用以下命令来安装 Singularity：

```sh
go install github.com/data-preservation-programs/singularity@latest
```

如果您无法使用 C 编译器，您也可以选择在没有 [CGO](https://zchee.github.io/golang-wiki/cgo/) 支持的情况下进行构建。不用担心，CGO 依赖仅用于稍快一些的 SQLite 后端，但它不应该在生产工作负载中使用。

```bash
CGO_ENABLED=0 go install github.com/data-preservation-programs/singularity@latest
```