# 로컬 디렉토리로부터 CAR 파일의 폴더 또는 파일 추출하기

{% code fullWidth="true" %}
```
NAME:
   singularity extract-car - CAR 파일 디렉토리로부터 폴더 또는 파일 추출

사용법:
   singularity extract-car [command options] [arguments...]

카테고리:
   유틸리티

옵션:
   --input-dir value, -i value  CAR 파일이 들어있는 입력 디렉토리. 이 디렉토리는 재귀적으로 스캔됩니다.
   --output value, -o value     추출할 출력 디렉토리 또는 파일. 존재하지 않을 경우에는 생성됩니다 (기본값: ".")
   --cid value, -c value        추출할 폴더 또는 파일의 CID
   --help, -h                   도움말 표시
```
{% endcode %}