# 源代码安装

要安装 singularity，您需要已安装 [Go 1.19 或更高版本](https://golang.org/dl/)。

```sh
wget -c https://golang.org/dl/go1.19.7.linux-amd64.tar.gz -O - \
    | sudo tar -xz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
    >> ~/.bashrc && source ~/.bashrc
```

然后您可以使用以下命令安装 singularity：

```sh
go install github.com/data-preservation-programs/singularity@latest
```

如果您的 C 编译器无法正常工作，也可以在没有 [CGO](https://zchee.github.io/golang-wiki/cgo/) 支持的情况下进行构建。不用担心，CGO 依赖仅用于略微更快的 SQLite 后端，不应该用于生产工作负载。

```bash
CGO_ENABLED=0 go install github.com/data-preservation-programs/singularity@latest
```