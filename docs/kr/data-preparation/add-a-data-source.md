---
설명: 준비가 필요한 데이터 소스에 연결합니다.
---

# 데이터 소스 추가

## 로컬 파일 시스템 데이터 소스 추가

가장 일반적인 데이터 소스는 로컬 파일 시스템입니다. 데이터 세트에 폴더를 데이터 소스로 추가하려면 다음과 같이 실행하세요:

```sh
singularity datasource add local my_dataset /mnt/dataset/folder
```

## 공개 S3 데이터 소스 추가

데이터 세트에 S3 데이터 소스를 추가하는 방법을 보여주기 위해 [Foldingathome COVID-19 Datasets](https://registry.opendata.aws/foldingathome-covid19/)라는 공개 데이터 세트를 사용해 보겠습니다.

```
singularity datasource add s3 my_dataset fah-public-data-covid19-cryptic-pocketst 
```

## 다음 단계

[start-dataset-worker.md](start-dataset-worker.md "mention")

## 관련 자료

[모든 데이터 소스 유형](../cli-reference/datasource/add/)