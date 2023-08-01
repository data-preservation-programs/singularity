# 데이터셋 워커 시작하기

데이터셋을 준비하기 위해 데이터셋 워커를 시작하려면 다음 명령을 실행하십시오.

```sh
singularity run dataset-worker --exit-on-complete --exit-on-error
```

기본적으로 데이터셋을 스캔, 패킹 및 dagnify 작업을 수행하는 단일 워커 스레드를 생성합니다. 작업이 완료되거나 오류가 발생하면 프로세스가 종료됩니다. 배포 시에는 계속 실행되도록 설정해야 합니다.

`--concurrency value` 플래그로 일부 동시성 값을 구성할 수도 있습니다.

준비가 완료되면 다음 명령 중 일부를 사용하여 준비된 데이터를 검사할 수 있습니다.

```sh
# 추가된 모든 데이터 소스 나열
singularity datasource list

# 스캔 및 패킹 결과 개괄 제공
singularity datasource status 1

# 루트 폴더의 각 파일의 CID 확인
singularity datasource inspect dir 1

# 생성된 모든 CAR 파일 확인
singularity datasource inspect chunks

# 준비된 모든 항목 확인
singularity datasource inspect items
```

## 다음 단계

[데이터 소스를 위한 DAG 생성하기](create-dag-for-the-data-source.md "mention")

## 관련 자료

[모든 데이터소스 나열하기](../cli-reference/datasource/list.md)

[데이터소스 준비 상태 확인하기](../cli-reference/datasource/status.md)

[데이터소스의 모든 항목 확인하기](../cli-reference/datasource/inspect/items.md)

[데이터소스의 모든 청크 확인하기](../cli-reference/datasource/inspect/chunks.md)

[데이터소스의 디렉토리 확인하기](../cli-reference/datasource/inspect/dir.md)