# Install from source

To install singularity, you need a working installation of [Go 1.19 or higher](https://golang.org/dl/)

```sh
wget -c https://golang.org/dl/go1.19.7.linux-amd64.tar.gz -O - \
    | sudo tar -xz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
    >> ~/.bashrc && source ~/.bashrc
```

Then you can install singularity with

```sh
go install github.com/data-preservation-programs/singularity@latest
```
