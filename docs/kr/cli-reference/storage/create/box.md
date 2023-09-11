# Box

{% code fullWidth="true" %}
```
이름:
   singularity storage create box - Box

사용법:
   singularity storage create box [command options] [arguments...]

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워두십시오.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      일반적으로 비워두십시오.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      제공자 기본값을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL.
      
      제공자 기본값을 사용하려면 비워두십시오.

   --root-folder-id
      rclone이 시작점으로 사용할 루트 폴더가 있는 경우 작성합니다.

   --box-config-file
      Box 앱 config.json 파일 위치
      
      보통 비워두십시오.
      
      `~`는 파일 이름에서 확장되며, `${RCLONE_CONFIG_DIR}`와 같은 환경 변수도 확장됩니다.

   --access-token
      Box 앱 기본 액세스 토큰
      
      일반적으로 비워두십시오.

   --box-sub-type
      

      예시:
         | user       | Rclone은 사용자를 대신하여 작업합니다.
         | enterprise | Rclone은 서비스 계정을 대신하여 작업합니다.

   --upload-cutoff
      (>= 50 MiB)일 경우 멀티파트 업로드로 전환하는 임계값입니다.

   --commit-retries
      멀티파트 파일을 커밋하는 데 시도 할 최대 횟수입니다.

   --list-chunk
      리스트 청크의 크기 (1-1000).

   --owned-by
      로그인한 사용자(e-mail 주소)가 소유한 항목만 표시합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --access-token value     Box 앱 기본 액세스 토큰 [$ACCESS_TOKEN]
   --box-config-file value  Box 앱 config.json 파일 위치 [$BOX_CONFIG_FILE]
   --box-sub-type value     (기본값: "user") [$BOX_SUB_TYPE]
   --client-id value        OAuth 클라이언트 ID [$CLIENT_ID]
   --client-secret value    OAuth 클라이언트 비밀번호 [$CLIENT_SECRET]
   --help, -h               도움말 표시

   고급

   --auth-url value        인증 서버 URL [$AUTH_URL]
   --commit-retries value  멀티파트 파일을 커밋하는 데 시도 할 최대 횟수 (기본값: 100) [$COMMIT_RETRIES]
   --encoding value        백엔드의 인코딩 (기본값: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --list-chunk value      리스트 청크의 크기 1-1000 (기본값: 1000) [$LIST_CHUNK]
   --owned-by value        로그인한 사용자(e-mail 주소)가 소유한 항목만 표시합니다. [$OWNED_BY]
   --root-folder-id value  rclone이 시작점으로 사용할 루트 폴더 (기본값: "0") [$ROOT_FOLDER_ID]
   --token value           JSON blob 형식의 OAuth 액세스 토큰 [$TOKEN]
   --token-url value       토큰 서버 URL [$TOKEN_URL]
   --upload-cutoff value   멀티파트 업로드로 전환하는 임계값 (>= 50 MiB) (기본값: "50Mi") [$UPLOAD_CUTOFF]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}