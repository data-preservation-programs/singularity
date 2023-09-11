# 데이터셋 준비 생성 및 관리

{% code fullWidth="true" %}
```
이름:
   singularity prep - 데이터셋 준비 생성 및 관리

사용법:
   singularity prep command [command options] [arguments...]

COMMANDS:
   create         새 준비 생성
   list           모든 준비 목록 보기
   status         준비의 준비 작업 상태 가져오기
   attach-source  준비에 소스 스토리지 연결
   attach-output  준비에 출력 스토리지 연결
   detach-output  준비의 출력 스토리지 연결 해제
   start-scan     소스 스토리지 스캔 시작
   pause-scan     스캔 작업 일시 중지
   start-pack     모든 팩 작업 또는 특정 팩 작업 시작/재시작
   pause-pack     모든 팩 작업 또는 특정 팩 작업 일시 중지
   start-daggen   모든 폴더 구조의 스냅샷을 생성하는 DAG 생성 시작
   pause-daggen   DAG 생성 작업 일시 중지
   list-pieces    준비에 대해 생성된 모든 조각 목록 보기
   add-piece      수동으로 준비에 조각 정보 추가. 외부 도구로 준비한 조각에 유용합니다.
   explore        경로로 준비된 소스 탐색
   attach-wallet  준비에 지갑 연결
   list-wallets   준비에 연결된 지갑 목록 보기
   detach-wallet  준비의 지갑 연결 해제
   help, h        명령어 목록 표시 또는 특정 명령어에 대한 도움말 표시

옵션:
   --help, -h  도움말 표시
```
{% endcode %}