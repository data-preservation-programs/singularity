# 로컬 디렉토리로부터 CAR 파일의 폴더 또는 파일 추출하기

{% code fullWidth="true" %}
```
이름:
   singularity tool extract-car - CAR 파일 디렉토리에서 폴더 또는 파일을 로컬 디렉토리로 추출합니다.

사용법:
   singularity tool extract-car [command 옵션] [인수...]

옵션:
   --input-dir value, -i value  CAR 파일을 포함하는 입력 디렉토리입니다. 이 디렉토리는 재귀적으로 스캔됩니다.
   --output value, -o value     추출할 출력 디렉토리 또는 파일입니다. 존재하지 않으면 생성됩니다 (기본값: ".")
   --cid value, -c value        추출할 폴더 또는 파일의 CID입니다.
   --help, -h                   도움말 표시
```
{% endcode %}