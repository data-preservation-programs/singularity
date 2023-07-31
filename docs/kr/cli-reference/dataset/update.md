# 기존 데이터셋 업데이트

{% code fullWidth="true" %}
```
이름:
   singularity dataset update - 기존 데이터셋 업데이트

사용법:
   singularity dataset update [옵션] <데이터셋_이름>

옵션:
   --help, -h  도움말 보기

   암호화

   --encryption-recipient value [ --encryption-recipient value ]  암호화 수신자의 공개 키
   --encryption-script value                                      사용자 정의 암호화를 위해 실행할 EncryptionScript 명령

   인라인 준비

   --output-dir value, -o value [ --output-dir value, -o value ]  CAR 파일의 출력 디렉터리 (기본값: 필요하지 않음)

   준비 매개변수

   --max-size value, -M value    생성될 CAR 파일의 최대 크기 (기본값: "30GiB")
   --piece-size value, -s value  조각 개시 계산에 사용되는 CAR 파일의 목표 조각 크기 (기본값: 추정함)

```
{% endcode %}