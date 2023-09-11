# 새로운 액세스 권한을 위한 위성 주소, API 키 및 암호문으로 액세스 권한 생성하기.

{% code fullWidth="true" %}
```
이름:
   singularity storage create storj new - 위성 주소, API 키 및 암호문을 이용하여 새로운 액세스 권한 생성하기.

사용법:
   singularity storage create storj new [command options] [arguments...]

설명:
   --satellite-address
      위성 주소.
      
      사용자 정의 위성 주소는 다음 형식과 일치해야 합니다: `<nodeid>@<address>:<port>`.

      예시:
         | us1.storj.io | US1
         | eu1.storj.io | EU1
         | ap1.storj.io | AP1

   --api-key
      API 키.

   --passphrase
      암호문.
      
      기존 객체에 액세스하려면 업로드할 때 사용한 암호문을 입력하세요.


옵션:
   --api-key value            API 키. [$API_KEY]
   --help, -h                 도움말 표시
   --passphrase value         암호문. [$PASSPHRASE]
   --satellite-address value  위성 주소. (기본값: "us1.storj.io") [$SATELLITE_ADDRESS]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}