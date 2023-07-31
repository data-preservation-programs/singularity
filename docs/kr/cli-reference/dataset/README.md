# 데이터셋 관리

{% code fullWidth="true" %}
```
이름:
   singularity dataset - 데이터셋 관리

사용법:
   singularity dataset command [command options] [arguments...]

COMMANDS:
   create         새로운 데이터셋 생성
   list           모든 데이터셋 나열
   update         기존 데이터셋 업데이트
   remove         특정 데이터셋 제거. 이 작업은 CAR 파일을 제거하지 않습니다.
   add-wallet     데이터셋에 지갑 연결. 지갑은 `singularity wallet import` 명령을 사용하여 먼저 가져와야 합니다.
   list-wallet    데이터셋에 연결된 모든 지갑 나열
   remove-wallet  데이터셋에서 연결된 지갑 제거
   add-piece      거래를 위해 데이터셋에 수동으로 피스(CAR 파일) 등록
   list-pieces    거래 가능한 데이터셋의 모든 피스 나열
   help, h        명령어 목록 또는 특정 명령어에 대한 도움말 표시

OPTIONS:
   --help, -h  도움말 표시
```
{% endcode %}