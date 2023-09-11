# 소스로부터 Singularity 설치하기

Singularity를 설치하기 전에 시스템에 [Go 1.20 이상](https://golang.org/dl/)을 설치했는지 확인해주세요.

## Go 설정하기

Go를 설치하고 실행하는 방법은 다음과 같습니다:

1. **Go 바이너리 다운로드 및 압축 해제**:
    ```bash
    wget -c https://golang.org/dl/go1.20.7.linux-amd64.tar.gz -O - \
        | sudo tar -xz -C /usr/local
    ```

2. **PATH 업데이트**:
   Go의 바이너리와 작업 디렉토리를 `PATH`에 추가하세요:
    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin:$(/usr/local/go/bin/go env GOPATH)/bin' \
        >> ~/.bashrc && source ~/.bashrc
    ```

## 최신 릴리즈 Singularity 설치하기

최신 안정 버전의 Singularity를 설치하기 위해 다음을 실행하세요:
```bash
go install github.com/data-preservation-programs/singularity@latest
```

## 미출시 기능 사용해보기

만약 Singularity의 최신 미출시 기능을 탐색하고 싶다면 다음과 같이 진행하세요:

1. **Singularity 저장소를 클론하기**:
    ```bash
    git clone https://github.com/data-preservation-programs/singularity.git
    ```
2. **Singularity 디렉토리로 이동하기**:
    ```bash
    cd singularity
    ```
3. **빌드 및 설치**:
    ```bash
    go build -o singularity .
    cp singularity $GOPATH/bin
    ```