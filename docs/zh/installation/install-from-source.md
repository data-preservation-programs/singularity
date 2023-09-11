# 从源代码安装 Singularity

在继续安装 Singularity 之前，请确保您的系统已经安装了 [Go 1.20 或更高版本](https://golang.org/dl/)。

## 配置 Go

按照以下步骤使 Go 正常运行：

1. **下载并解压 Go 二进制文件**：
    ```bash
    wget -c https://golang.org/dl/go1.20.7.linux-amd64.tar.gz -O - \
        | sudo tar -xz -C /usr/local
    ```

2. **更新 PATH 变量**：
   将 Go 的二进制文件和工作目录添加到您的 `PATH` 变量中：
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
        >> ~/.bashrc && source ~/.bashrc
    ```

## 安装最新版本的 Singularity

安装最新稳定版的 Singularity：
```bash
go install github.com/data-preservation-programs/singularity@latest
```

## 尝试未发布的功能

如果您想要探索最新的未发布功能：

1. **克隆 Singularity 仓库**：
    ```bash
    git clone https://github.com/data-preservation-programs/singularity.git
    ```
2. **进入 Singularity 目录**：
    ```bash
    cd singularity
    ```
3. **构建并安装**：
    ```bash
    go build -o singularity .
    cp singularity $GOPATH/bin
    ```