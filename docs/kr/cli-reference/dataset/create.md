# 새 데이터셋 생성하기

{% code fullWidth="true" %}
```
명령어:
   singularity dataset create - 새 데이터셋 생성

사용법:
   singularity dataset create [command options] <dataset_name>

설명:
   <dataset_name>은 데이터셋의 고유한 식별자여야 합니다.
   데이터셋은 서로 다른 데이터셋을 구분하기 위한 최상위 객체입니다.

옵션:
   --help, -h  도움말 보기

   암호화

   --encryption-recipient value [ --encryption-recipient value ]  암호화 수취인의 공개키
   --encryption-script value                                      [WIP] 사용자 정의 암호화를 위해 실행할 암호화 스크립트 명령어

   인라인 준비

   --output-dir value, -o value [ --output-dir value, -o value ]  CAR 파일을 위한 출력 디렉터리 (기본값: 필요하지 않음)

   준비 매개변수

   --max-size value, -M value    생성될 CAR 파일의 최대 크기 (기본값: "31.5GiB")
   --piece-size value, -s value  조각 커밋 계산에 사용되는 CAR 파일의 대상 조각 크기 (기본값: 추론됨)

```
{% endcode %}