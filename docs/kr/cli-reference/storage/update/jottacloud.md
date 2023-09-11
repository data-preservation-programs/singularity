# Jottacloud

{% code fullWidth="true" %}
```
이름:
   singularity storage update jottacloud - Jottacloud

사용법:
   singularity storage update jottacloud [command options] <name|id>

설명:
   --md5-memory-limit
      필요한 경우 MD5를 계산하기 위해 디스크에 캐시될 파일의 크기 제한입니다.

   --trashed-only
      휴지통에 있는 파일만 표시합니다.
      
      이렇게 하면 휴지통에 있는 파일이 원래 디렉토리 구조로 표시됩니다.

   --hard-delete
      파일을 영구적으로 삭제하여 휴지통에 넣지 않습니다.

   --upload-resume-limit
      업로드가 실패한 경우 이 크기보다 큰 파일을 다시 전송할 수 있습니다.

   --no-versions
      파일을 덮어쓰는 대신에 삭제하고 다시 만들어서 서버 측 버전 관리를 피합니다.

   --encoding
      백엔드의 인코딩 설정입니다.
      
      더 많은 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h  도움말 표시

   고급 옵션

   --encoding value             백엔드의 인코딩 설정입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete                파일을 영구적으로 삭제하여 휴지통에 넣지 않습니다. (기본값: false) [$HARD_DELETE]
   --md5-memory-limit value     필요한 경우 MD5를 계산하기 위해 디스크에 캐시될 파일의 크기 제한입니다. (기본값: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                파일을 덮어쓰는 대신에 삭제하고 다시 만들어서 서버 측 버전 관리를 피합니다. (기본값: false) [$NO_VERSIONS]
   --trashed-only               휴지통에 있는 파일만 표시합니다. (기본값: false) [$TRASHED_ONLY]
   --upload-resume-limit value  업로드가 실패한 경우 이 크기보다 큰 파일을 다시 전송할 수 있습니다. (기본값: "10Mi") [$UPLOAD_RESUME_LIMIT]

```
{% endcode %}