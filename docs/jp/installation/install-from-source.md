# ソースからのSingularityのインストール

Singularityをインストールする前に、システムに[Go 1.20以上](https://golang.org/dl/)がインストールされているかどうかを確認してください。

## Goのセットアップ

Goを使用するためには、次の手順に従ってください:

1. **Goのバイナリをダウンロードして展開**:
    ```bash
    wget -c https://golang.org/dl/go1.20.7.linux-amd64.tar.gz -O - \
        | sudo tar -xz -C /usr/local
    ```

2. **PATHを更新**:
   Goのバイナリディレクトリとワークスペースディレクトリを`PATH`に追加します:
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
        >> ~/.bashrc && source ~/.bashrc
    ```

## 最新リリースのSingularityのインストール

最新の安定版Singularityの場合:
```bash
go install github.com/data-preservation-programs/singularity@latest
```

## 未リリースの機能を試す

最新の未リリース機能を試したい場合:

1. **Singularityリポジトリをクローンする**:
    ```bash
    git clone https://github.com/data-preservation-programs/singularity.git
    ```
2. **Singularityディレクトリに移動**:
    ```bash
    cd singularity
    ```
3. **ビルドしてインストール**:
    ```bash
    go build -o singularity .
    cp singularity $GOPATH/bin
    ```