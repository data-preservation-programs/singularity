# Singularity 시작하기

다음 단계를 따라 Singularity를 설정하고 사용을 시작하세요.

## 1. 데이터베이스 초기화

만약 처음으로 Singularity를 사용하는 경우, 데이터베이스를 초기화해야 합니다. 이 단계는 한 번만 필요합니다.

```sh
singularity admin init
```

## 2. 스토리지 시스템에 연결하기

Singularity는 40여 가지 다른 스토리지 시스템과 원활한 통합을 제공하기 위해 RClone과 협력합니다. 이러한 스토리지 시스템은 주로 두 가지 역할을 수행할 수 있습니다:
* **소스 스토리지**: 데이터셋이 현재 저장되어 있는 곳이며, Singularity가 데이터 준비를 위해 원본 데이터를 가져올 위치입니다.
* **출력 스토리지**: Singularity가 처리 후 CAR(Content Addressable Archive) 파일을 저장할 목적지입니다.
필요에 맞는 스토리지 시스템을 선택하고 Singularity와 연결하여 데이터셋 준비를 시작하세요.

### 2a. 로컬 파일 시스템 추가하기

가장 일반적인 스토리지 시스템은 로컬 파일 시스템입니다. 아래 명령어를 사용하여 폴더를 소스 스토리지로 추가할 수 있습니다:

```sh
singularity storage create local --name "my-source" --path "/mnt/dataset/folder"
```

### 2b. S3 데이터 소스 추가하기

AWS S3, MinIO 등 S3 호환 스토리지 시스템을 포함한 어떤 S3 호환 스토리지 시스템이든 사용할 수 있습니다. 아래는 공개 데이터셋의 예입니다:

```sh
singularity storage create s3 aws --name "my-source" --path "public-dataset-test"
```

## 3. 데이터 준비 생성하기
```sh
singularity prep create --source "my-source" --name "my-prep"
```

## 4. 데이터 준비 진행하기
```sh
singularity prep start-scan my-prep my-source
singularity run dataset-worker
```

## 5. 데이터 준비 상태와 결과 확인하기
```sh
singularity prep status my-prep
singularity prep list-pieces my-prep
```