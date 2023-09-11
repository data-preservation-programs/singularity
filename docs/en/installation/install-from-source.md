# Installing Singularity from Source

Before you proceed with installing Singularity, make sure you have [Go 1.20 or higher](https://golang.org/dl/) installed on your system.

## Setting Up Go

To get Go up and running, follow these steps:

1. **Download and Extract Go Binaries**:
    ```bash
    wget -c https://golang.org/dl/go1.20.7.linux-amd64.tar.gz -O - \
        | sudo tar -xz -C /usr/local
    ```

2. **Update PATH**:
   Add Go's binary and workspace directories to your `PATH`:
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
        >> ~/.bashrc && source ~/.bashrc
    ```

## Installing the Latest Release of Singularity

For the latest stable release of Singularity:
```bash
go install github.com/data-preservation-programs/singularity@latest
```

## Trying Out Unreleased Features

If you're keen on exploring the latest, yet-to-be-released features of Singularity:

1. **Clone the Singularity Repository**:
    ```bash
    git clone https://github.com/data-preservation-programs/singularity.git
    ```
2. **Navigate to the Singularity Directory**:
    ```bash
    cd singularity
    ```
3. **Build and Install**:
    ```bash
    go build -o singularity .
    cp singularity $GOPATH/bin
    ```
