# 소스에서 설치하기

Singularity를 설치하려면 [Go 1.19 이상](https://golang.org/dl/)이 설치되어 있어야 합니다.

```sh
wget -c https://golang.org/dl/go1.19.7.linux-amd64.tar.gz -O - \
    | sudo tar -xz -C /usr/local
echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
    >> ~/.bashrc && source ~/.bashrc
```

그런 다음 아래 명령어로 Singularity를 설치할 수 있습니다.

```sh
go install github.com/data-preservation-programs/singularity@latest
```

만약 C 컴파일러를 사용할 수 없다면 [CGO](https://zchee.github.io/golang-wiki/cgo/)를 지원하지 않는 방식으로 빌드할 수도 있습니다. 걱정하지 마세요. CGO는 제품용 워크로드에서는 사용되지 않아도 되는 약간 더 빠른 SQLite 백엔드에 필요합니다.

```bash
CGO_ENABLED=0 go install github.com/data-preservation-programs/singularity@latest
```