# 데이터 소스를 위한 DAG 만들기

이 문맥에서 DAG는 데이터 소스의 모든 관련 폴더 정보와 파일이 여러 청크로 분리되는 방법을 포함합니다. 이 DAG의 CAR 파일이 스토리지 제공자에 의해 봉인되었다면, 데이터셋의 단일 루트 CID를 사용하여 파일을 검색할 수 있습니다.

데이터 소스에 대한 DAG 생성 프로세스를 트리거하려면 다음을 실행합니다.

```sh
# 단일 데이터소스가 있다고 가정합니다.
singularity datasource daggen 1
```

이제 작업이 데이터베이스에 기록되었습니다. 데이터셋 워커를 다시 실행하거나 워커가 이미 실행 중인 경우 작업을 가져 오기를 기다려야합니다.

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

완료되면 관련 DAG를 확인할 수 있습니다.

```
singularity datasource inspect dags 1
```

DAG의 CAR 파일은 거래 생성을 위해 자동으로 포함됩니다.

## 다음 단계

[distribute-car-files.md](../content-distribution/distribute-car-files.md "언급")

## 관련 자료

[DAG 생성 트리거](../cli-reference/datasource/daggen.md)

[데이터 소스의 DAG 검사](../cli-reference/datasource/inspect/dags.md)