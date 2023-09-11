# 새로운 액세스 그랜트 생성하기

{% code fullWidth="true" %}
```
명령어:
   singularity storage update storj new - 위성 주소, API 키 및 암호구문으로 새로운 액세스 그랜트를 생성합니다.

사용법:
   singularity storage update storj new [command options] <name|id>

설명:
   --satellite-address
      위성 주소.
      
      사용자 지정 위성 주소는 다음 형식과 일치해야 합니다: `<노드id>@<주소>:<포트>`.

      예제:
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key
      API 키.

   --passphrase
      암호구문.
      
      기존 객체에 액세스하려면 업로드에 사용한 암호구문을 입력하세요.


옵션:
   --satellite-address value  위성 주소. (기본값: "us1.storj.io") [$SATELLITE_ADDRESS]
   --api-key value            API 키. [$API_KEY]
   --passphrase value         암호구문. [$PASSPHRASE]
   --help, -h                 도움말 표시
```
{% endcode %}