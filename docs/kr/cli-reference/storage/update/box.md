# Box

{% code fullWidth="true" %}
```
이름:
   singularity storage update box - Box

사용법:
   singularity storage update box [command options] <이름|아이디>

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워두십시오.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      일반적으로 비워두십시오.

   --token
      OAuth 액세스 토큰(JSON 형식)입니다.

   --auth-url
      인증 서버 URL.
      
      제공자 기본값을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL.
      
      제공자 기본값을 사용하려면 비워두십시오.

   --root-folder-id
      rclone이 시작 위치로 사용하는 루트 폴더 이외의 폴더를 지정합니다.

   --box-config-file
      Box 앱 config.json 파일 위치입니다.
      
      일반적으로 비워두십시오.
      
      `~`로 시작하는 경우 파일 이름과 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수가 확장됩니다.

   --access-token
      Box 앱 주 요청 토큰입니다.
      
      일반적으로 비워두십시오.

   --box-sub-type
      

      예시:
         | user       | Rclone은 사용자를 대신하여 작동해야 합니다.
         | enterprise | Rclone은 서비스 계정을 대신하여 작동해야 합니다.

   --upload-cutoff
      대용량 업로드로 전환할 기준 크기입니다(50 MiB 이상).

   --commit-retries
      멀티파트 파일을 커밋하는 시도 횟수의 최대값입니다.

   --list-chunk
      목록 분할 크기입니다(1-1000).

   --owned-by
      이메일 주소로 전달된 로그인(이메일 주소)이 소유한 항목만 표시합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      더 많은 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-token value     Box 앱 주 요청 토큰 [$ACCESS_TOKEN]
   --box-config-file value  Box 앱 config.json 파일 위치 [$BOX_CONFIG_FILE]
   --box-sub-type value     (기본값: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuth 클라이언트 ID [$CLIENT_ID]
   --client-secret value    OAuth 클라이언트 시크릿 [$CLIENT_SECRET]
   --help, -h               도움말 표시

   고급 옵션

   --auth-url value        인증 서버 URL [$AUTH_URL]
   --commit-retries value  멀티파트 파일을 커밋하는 시도 횟수의 최대값(기본값: 100) [$COMMIT_RETRIES]
   --encoding value        백엔드의 인코딩(기본값: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --list-chunk value      목록 분할 크기 1-1000(기본값: 1000) [$LIST_CHUNK]
   --owned-by value        이메일 주소로 전달된 로그인(이메일 주소)이 소유한 항목만 표시합니다. [$OWNED_BY]
   --root-folder-id value  rclone이 시작 위치로 사용하는 루트 폴더 이외의 폴더를 지정합니다(기본값: "0") [$ROOT_FOLDER_ID]
   --token value           OAuth 액세스 토큰(JSON 형식) [$TOKEN]
   --token-url value       토큰 서버 URL [$TOKEN_URL]
   --upload-cutoff value   대용량 업로드로 전환할 기준 크기(50 MiB 이상)(기본값: "50Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}