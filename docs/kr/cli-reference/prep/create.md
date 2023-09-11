# 새로운 준비 생성

{% code fullWidth="true" %}
```
명령어:
   singularity prep create - 새로운 준비 생성

사용법:
   singularity prep create [옵션] [인수들...]

분류:
   준비 관리

옵션들:
   --delete-after-export              CAR 파일로 내보낸 후 소스 파일을 삭제할 지 여부 (기본값: false)
   --help, -h                         도움말 출력
   --max-size value                   단일 CAR 파일의 최대 크기 (기본값: "31.5GiB")
   --name value                       준비 이름 (기본값: 자동 생성)
   --output value [ --output value ]  준비에 사용할 출력 스토리지의 ID 또는 이름
   --piece-size value                 조각 검증을 위한 CAR 파일의 목표 크기 (기본값: --max-size에 의해 결정됨)
   --source value [ --source value ]  준비에 사용할 소스 스토리지의 ID 또는 이름

   로컬 출력 경로로 빠른 생성

   --local-output value [ --local-output value ]  준비에 사용할 로컬 출력 경로. 이 플래그는 제공된 경로로 출력 스토리지를 생성하는 편리한 기능입니다.

   로컬 소스 경로로 빠른 생성

   --local-source value [ --local-source value ]  준비에 사용할 로컬 소스 경로. 이 플래그는 제공된 경로로 소스 스토리지를 생성하는 편리한 기능입니다.
```
{% endcode %}