# 데이터셋 스캔 및 준비 작업을 처리하기 위해 데이터셋 준비 워커 시작

{% code fullWidth="true" %}
```
이름:
   singularity run dataset-worker - 데이터셋 스캔 및 준비 작업을 처리하기 위해 데이터셋 준비 워커 시작

사용법:
   singularity run dataset-worker [옵션] [인수...]

옵션:
   --concurrency 값     동시 작업하는 워커의 개수 (기본값: 1) [$DATASET_WORKER_CONCURRENCY]
   --enable-scan         데이터셋 스캔 가능 여부 (기본값: true) [$DATASET_WORKER_ENABLE_SCAN]
   --enable-pack         데이터셋 패킹 가능 여부 (CID 계산 및 CAR 파일로 패킹) (기본값: true) [$DATASET_WORKER_ENABLE_PACK]
   --enable-dag          데이터셋의 디렉터리 구조를 유지하면서 DAG 생성 가능 여부 (기본값: true) [$DATASET_WORKER_ENABLE_DAG]
   --exit-on-complete    더 이상 작업이 없을 때 워커 종료 여부 (기본값: false) [$DATASET_WORKER_EXIT_ON_COMPLETE]
   --exit-on-error       에러 발생 시 워커 종료 여부 (기본값: false) [$DATASET_WORKER_EXIT_ON_ERROR]
   --help, -h            도움말 표시
```
{% endcode %}