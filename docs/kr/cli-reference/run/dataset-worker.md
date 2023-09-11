# 데이터셋 스캔 및 준비 작업을 처리하는 데이터셋 준비 워커 시작하기

{% code fullWidth="true" %}
```
이름:
   singularity run dataset-worker - 데이터셋 스캔 및 준비 작업을 처리하는 데이터셋 준비 워커 시작하기

사용법:
   singularity run dataset-worker [옵션 명령] [인수...]

옵션:
   --concurrency value  동시 실행할 워커 수 (기본값: 1)
   --enable-scan        데이터셋 스캔 활성화 (기본값: true)
   --enable-pack        CIDs를 계산하여 CAR 파일에 팩킹하는 데이터셋 팩킹 활성화 (기본값: true)
   --enable-dag         데이터셋의 디렉토리 구조를 유지하는 DAG 생성 활성화 (기본값: true)
   --exit-on-complete   더 이상 작업이 없을 때 워커 종료 (기본값: false)
   --exit-on-error      에러가 발생할 때 워커 종료 (기본값: false)
   --help, -h           도움말 표시
```
{% endcode %}