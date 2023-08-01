# Jottacloud

{% code fullWidth="true" %}
```
명령어:
   singularity datasource add jottacloud - Jottacloud

사용법:
   singularity datasource add jottacloud [추가 옵션] <데이터셋_이름> <소스_경로>

설명:
   --jottacloud-encoding
      백엔드의 인코딩 설정입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --jottacloud-hard-delete
      파일을 휴지통으로 보내지 않고 재차 삭제합니다.

   --jottacloud-md5-memory-limit
      이 크기보다 큰 파일은 필요한 경우 MD5를 계산하기 위해 디스크에 캐시됩니다.

   --jottacloud-no-versions
      파일을 덮어쓰는 대신 삭제하고 다시 생성함으로써 서버 측 버전관리를 피합니다.

   --jottacloud-trashed-only
      휴지통에 있는 파일만 표시합니다.
      
      원래의 디렉토리 구조에 있는 휴지통 파일이 표시됩니다.

   --jottacloud-upload-resume-limit
      이 크기보다 큰 파일은 업로드가 실패한 경우 재개될 수 있습니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  소스 디렉토리를 자동으로 재스캔하는 간격을 설정합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   Jottacloud 옵션

   --jottacloud-encoding value             백엔드의 인코딩 설정입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$JOTTACLOUD_ENCODING]
   --jottacloud-hard-delete value          파일을 휴지통으로 보내지 않고 재차 삭제할지 설정합니다. (기본값: "false") [$JOTTACLOUD_HARD_DELETE]
   --jottacloud-md5-memory-limit value     이 크기보다 큰 파일은 필요한 경우 MD5를 계산하기 위해 디스크에 캐시됩니다. (기본값: "10Mi") [$JOTTACLOUD_MD5_MEMORY_LIMIT]
   --jottacloud-no-versions value          파일을 덮어쓰는 대신 삭제하고 다시 생성함으로써 서버 측 버전관리를 피할지 설정합니다. (기본값: "false") [$JOTTACLOUD_NO_VERSIONS]
   --jottacloud-trashed-only value         휴지통에 있는 파일만 표시할지 설정합니다. (기본값: "false") [$JOTTACLOUD_TRASHED_ONLY]
   --jottacloud-upload-resume-limit value  이 크기보다 큰 파일은 업로드가 실패한 경우 재개됩니다. (기본값: "10Mi") [$JOTTACLOUD_UPLOAD_RESUME_LIMIT]

```
{% endcode %}