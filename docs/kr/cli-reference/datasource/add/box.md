# Box

{% code fullWidth="true" %}
```
이름:
   singularity 데이터 소스 추가 box - Box

사용법:
   singularity 데이터 소스 추가 box [command options] <데이터세트_이름> <소스_경로>

설명:
   --box-access-token
      Box 앱 주요 액세스 토큰
      
      보통 비워 둡니다.

   --box-auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 비워 둡니다.

   --box-box-config-file
      Box 앱 config.json 위치
      
      보통 비워 둡니다.
      
      `~`로 시작하면 파일 이름에서 확장됩니다. `${RCLONE_CONFIG_DIR}`와 같은 환경 변수도 확장됩니다.

   --box-box-sub-type
      

      예시:
         | 사용자       | Rclone은 사용자를 대신하여 작동합니다.
         | 기업     | Rclone은 서비스 계정을 대신하여 작동합니다.

   --box-client-id
      OAuth 클라이언트 ID.
      
      보통 비워 둡니다.

   --box-client-secret
      OAuth 클라이언트 비밀번호.
      
      보통 비워 둡니다.

   --box-commit-retries
      멀티파트 파일 커밋을 시도하는 최대 횟수입니다.

   --box-encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --box-list-chunk
      목록 청크의 크기는 1-1000입니다.

   --box-owned-by
      로그인 (이메일 주소)으로 소유한 항목만 표시합니다.

   --box-root-folder-id
      Rclone에게 시작점으로 사용할 특정 폴더를 입력합니다.

   --box-token
      OAuth 액세스 토큰을 JSON blob으로 입력합니다.

   --box-token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 비워 둡니다.

   --box-upload-cutoff
      멀티파트 업로드로 전환할 기준 (>= 50 MiB).


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    데이터세트를 CAR 파일로 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔 이후로 이 간격이 지날 때마다 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비됨)

   Box 옵션

   --box-access-token value     Box 앱 주요 액세스 토큰 [$BOX_ACCESS_TOKEN]
   --box-auth-url value         인증 서버 URL. [$BOX_AUTH_URL]
   --box-box-config-file value  Box 앱 config.json 위치 [$BOX_BOX_CONFIG_FILE]
   --box-box-sub-type value     (기본값: "user") [$BOX_BOX_SUB_TYPE]
   --box-client-id value        OAuth 클라이언트 ID. [$BOX_CLIENT_ID]
   --box-client-secret value    OAuth 클라이언트 비밀번호. [$BOX_CLIENT_SECRET]
   --box-commit-retries value   멀티파트 파일 커밋을 시도하는 최대 횟수입니다. (기본값: "100") [$BOX_COMMIT_RETRIES]
   --box-encoding value         백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,RightSpace,InvalidUtf8,Dot") [$BOX_ENCODING]
   --box-list-chunk value       목록 청크의 크기는 1-1000입니다. (기본값: "1000") [$BOX_LIST_CHUNK]
   --box-owned-by value         로그인 (이메일 주소)으로 소유한 항목만 표시합니다. [$BOX_OWNED_BY]
   --box-root-folder-id value   Rclone에게 시작점으로 사용할 특정 폴더를 입력합니다. (기본값: "0") [$BOX_ROOT_FOLDER_ID]
   --box-token value            OAuth 액세스 토큰을 JSON blob으로 입력합니다. [$BOX_TOKEN]
   --box-token-url value        토큰 서버 URL. [$BOX_TOKEN_URL]
   --box-upload-cutoff value    멀티파트 업로드로 전환할 기준 (>= 50 MiB). (기본값: "50Mi") [$BOX_UPLOAD_CUTOFF]

```
{% endcode %}