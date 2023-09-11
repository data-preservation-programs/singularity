# Pcloud

{% code fullWidth="true" %}
```
이름:
   singularity storage update pcloud - Pcloud

사용법:
   singularity storage update pcloud [command options] <이름|ID>

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      보통 비워둡니다.

   --token
      JSON blob 형태의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      제공자의 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      제공자의 기본값을 사용하려면 비워둡니다.

   --encoding
      백엔드의 인코딩 방식.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --root-folder-id
      rclone에서 시작 지점으로 사용할 루트 폴더(ID)를 입력합니다.

   --hostname
      연결할 호스트 이름.
      
      보통 rclone이 초기 oauth 연결을 수행할 때 설정되지만,
      rclone authorize에서 원격 구성을 사용하는 경우 수동으로 설정해야 합니다.
      

      예시:
         | api.pcloud.com  | 원본/미국 지역
         | eapi.pcloud.com | 유럽 지역

   --username
      pcloud 사용자 이름.
            
      cleanup 명령어를 사용하려면 필요합니다. 버그로 인해
      해당 API는 OAuth 인증을 지원하지 않으므로 패스워드 인증을 사용해야 합니다.

   --password
      pcloud 패스워드.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시

   고급

   --auth-url value        인증 서버 URL. [$AUTH_URL]
   --encoding value        백엔드의 인코딩 방식. (기본값: "Slash,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hostname value        연결할 호스트 이름. (기본값: "api.pcloud.com") [$HOSTNAME]
   --password value        pcloud 패스워드. [$PASSWORD]
   --root-folder-id value  rclone에서 시작 지점으로 사용할 루트 폴더(ID)를 입력합니다. (기본값: "d0") [$ROOT_FOLDER_ID]
   --token value           JSON blob 형태의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value       토큰 서버 URL. [$TOKEN_URL]
   --username value        pcloud 사용자 이름. [$USERNAME]
 
```
{% endcode %}