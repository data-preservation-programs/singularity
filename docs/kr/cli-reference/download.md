# 메타데이터 API에서 CAR 파일 다운로드

{% code fullWidth="true" %}
```
이름:
   singularity download - 메타데이터 API에서 CAR 파일 다운로드

사용법:
   singularity download [command options] PIECE_CID

카테고리:
   유틸리티

옵션:
   일반 옵션

   --api 값을, -j 값을 사용하여, 병렬로 다운로드할 수 있는 파일 개수
   --out-dir 값을, -o 값을 사용하여 CAR 파일을 저장할 디렉터리

acd 옵션

   --acd-auth-url 값을 사용하여, 인증 서버 URL을 설정
   --acd-client-id 값을 사용하여, OAuth 클라이언트 ID를 설정
   --acd-client-secret 값을 사용하여, OAuth 클라이언트 비밀키를 설정

azureblob 옵션

   --azureblob-access-tier 값을 사용하여, 블롭의 접근 수준을 설정
   --azureblob-account 값을 사용하여, Azure 스토리지 계정 이름을 설정

koofr 옵션

   --koofr-encoding 값을 사용하여, 백엔드 인코딩을 설정
   --koofr-endpoint 값을 사용하여, 사용할 Koofr API 엔드포인트를 설정
   --koofr-mountid 값을 사용하여, 사용할 마운트 ID를 설정
   --koofr-password 값을 사용하여, rclone의 비밀번호를 설정
   --koofr-provider 값을 사용하여, 스토리지 제공자를 설정
   --koofr-setmtime 값을 사용하여, 수정 시간을 설정할 수 있는 백엔드를 설정
   --koofr-user 값을 사용하여, 사용자 이름을 설정

onedrive 옵션

   --onedrive-encoding 값을 사용하여, 백엔드 인코딩을 설정
   --onedrive-expose-onenote-files 값을 사용하여, OneNote 파일을 디렉터리 목록에 표시할지 여부를 설정
   --onedrive-use-created-date 값을 사용하여, 수정 날짜 대신 생성 날짜를 사용할지 여부를 설정

webdav 옵션

   --webdav-encoding 값을 사용하여, 백엔드 인코딩을 설정

zoho 옵션

   --zoho-encoding 값을 사용하여, 백엔드 인코딩을 설정
```
{% endcode %}