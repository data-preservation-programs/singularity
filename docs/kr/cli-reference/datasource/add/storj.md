# Storj 분산형 클라우드 스토리지

{% code fullWidth="true" %}
```
이름:
   singularity datasource add storj - Storj 분산형 클라우드 스토리지

사용법:
   singularity datasource add storj [command options] <데이터셋_이름> <소스_경로>

설명:
   --storj-access-grant
      [공급자] - 기존
         액세스 권한 부여.

   --storj-api-key
      [공급자] - 새로운
         API 키.

   --storj-passphrase
      [공급자] - 새로운
         암호화 암호.
         
         기존 객체에 액세스하려면 업로드 시 사용된 암호화 암호를 입력하십시오.

   --storj-provider
      인증 방법을 선택하십시오.

      예시:
         | 기존     | 기존 액세스 권한을 사용합니다.
         | 새로운   | 위성 주소, API 키 및 암호화 암호로 새 액세스 권한을 생성합니다.

   --storj-satellite-address
      [공급자] - 새로운
         위성 주소.
         
         사용자 정의 위성 주소는 다음 형식과 일치해야 합니다: `<노드ID>@<주소>:<포트>`.

         예시:
            | us1.storj.io | 미국
            | eu1.storj.io | 유럽
            | ap1.storj.io | 아시아 태평양


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막으로 성공적으로 검사된 시간부터 설정된 시간 간격이 지나면 소스 디렉터리를 자동으로 검사합니다. (기본값: 사용 안 함)
   --scanning-state value   초기 검사 상태를 설정합니다. (기본값: 준비 완료)

   storj를 위한 옵션

   --storj-access-grant value       액세스 권한 부여. [$STORJ_ACCESS_GRANT]
   --storj-api-key value            API 키. [$STORJ_API_KEY]
   --storj-passphrase value         암호화 암호. [$STORJ_PASSPHRASE]
   --storj-provider value           인증 방법을 선택하십시오. (기본값: "기존") [$STORJ_PROVIDER]
   --storj-satellite-address value  위성 주소. (기본값: "us1.storj.io") [$STORJ_SATELLITE_ADDRESS]
```
{% endcode %}