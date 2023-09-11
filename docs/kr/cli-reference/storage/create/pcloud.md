# Pcloud

{% code fullWidth="true" %}
```
명령:
   singularity storage create pcloud - Pcloud

사용법:
   singularity storage create pcloud [command options] [arguments...]

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워 둡니다.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      일반적으로 비워 둡니다.

   --token
      OAuth 액세스 토큰(JSON 블롭).

   --auth-url
      인증 서버 URL.
      
      사용자 기본 설정을 사용하려면 비워 둡니다.

   --token-url
      토큰 서버 URL.
      
      사용자 기본 설정을 사용하려면 비워 둡니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --root-folder-id
      rclone에서 시작 지점으로 사용할 최상위 폴더를 입력하십시오.

   --hostname
      연결할 호스트명.
      
      일반적으로 rclone이 초기에 OAuth 연결을 수행하지만,
      rclone authorize를 사용하여 원격 구성을 하는 경우 수동으로 설정해야 합니다.
      

      예시:
         | api.pcloud.com  | 원본/미국 지역
         | eapi.pcloud.com | 유럽 지역

   --username
      pcloud 사용자명.
            
      cleanup 명령을 사용하려면 필요합니다. pcloud API에서 필수로 요구되는 API가 OAuth 인증을 지원하지 않으므로, 사용자 비밀번호 인증에 의존해야 합니다.

   --password
      pcloud 비밀번호.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시

   고급

   --auth-url value        인증 서버 URL. [$AUTH_URL]
   --encoding value        백엔드의 인코딩. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hostname value        연결할 호스트명. (기본값: "api.pcloud.com") [$HOSTNAME]
   --password value        pcloud 비밀번호. [$PASSWORD]
   --root-folder-id value  rclone에서 시작 지점으로 사용할 최상위 폴더를 입력하십시오. (기본값: "d0") [$ROOT_FOLDER_ID]
   --token value           OAuth 액세스 토큰(JSON 블롭). [$TOKEN]
   --token-url value       토큰 서버 URL. [$TOKEN_URL]
   --username value        pcloud 사용자명. [$USERNAME]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}