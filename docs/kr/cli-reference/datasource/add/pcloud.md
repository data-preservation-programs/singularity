# Pcloud

{% code fullWidth="true" %}
```
이름:
   singularity datasource add pcloud - Pcloud

사용법:
   singularity datasource add pcloud [command options] <dataset_name> <source_path>

설명:
   --pcloud-auth-url
      인증 서버 URL.
      
      이 값을 비워두면 제공자의 기본값을 사용합니다.

   --pcloud-client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워둡니다.

   --pcloud-client-secret
      OAuth 클라이언트 비밀.
      
      일반적으로 비워둡니다.

   --pcloud-encoding
      백엔드의 인코딩 방식.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --pcloud-hostname
      연결할 호스트 이름.
      
      보통 rclone이 oauth 연결을 처음 수행할 때 설정되지만,
      rclone authorize와 함께 원격 구성을 사용하는 경우 직접 설정해야 합니다.
      

      예시:
         | api.pcloud.com  | 기본/미국 리전
         | eapi.pcloud.com | EU 리전

   --pcloud-password
      pcloud 비밀번호.

   --pcloud-root-folder-id
      rclone이 시작 위치로 사용할 루트가 아닌 폴더를 지정하세요.

   --pcloud-token
      JSON blob으로 표현된 OAuth 액세스 토큰.

   --pcloud-token-url
      토큰 서버 URL.
      
      이 값을 비워두면 제공자의 기본값을 사용합니다.

   --pcloud-username
      pcloud 사용자명.
            
      cleanup 명령을 사용하려면 필요합니다. pcloud API에 버그로 인해
      필요한 API가 OAuth 인증을 지원하지 않으므로 사용자 암호 인증에 의존해야 합니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export     [주의] CAR 파일로 데이터셋을 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value   소스 디렉토리를 자동으로 다시 스캔하는 간격입니다. (기본값: 비활성화)
   --scanning-state value    초기 스캔 상태를 설정합니다. (기본값: 준비 완료)

   pcloud용 옵션

   --pcloud-auth-url value         인증 서버 URL. [$PCLOUD_AUTH_URL]
   --pcloud-client-id value        OAuth 클라이언트 ID. [$PCLOUD_CLIENT_ID]
   --pcloud-client-secret value    OAuth 클라이언트 비밀. [$PCLOUD_CLIENT_SECRET]
   --pcloud-encoding value         백엔드의 인코딩 방식. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$PCLOUD_ENCODING]
   --pcloud-hostname value         연결할 호스트 이름. (기본값: "api.pcloud.com") [$PCLOUD_HOSTNAME]
   --pcloud-password value         pcloud 비밀번호. [$PCLOUD_PASSWORD]
   --pcloud-root-folder-id value   rclone이 시작 위치로 사용할 루트가 아닌 폴더를 지정하세요. (기본값: "d0") [$PCLOUD_ROOT_FOLDER_ID]
   --pcloud-token value            JSON blob으로 표현된 OAuth 액세스 토큰. [$PCLOUD_TOKEN]
   --pcloud-token-url value        토큰 서버 URL. [$PCLOUD_TOKEN_URL]
   --pcloud-username value         pcloud 사용자명. [$PCLOUD_USERNAME]

```
{% endcode %}