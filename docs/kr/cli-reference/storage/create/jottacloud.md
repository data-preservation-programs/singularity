# Jottacloud

{% code fullWidth="true" %}
```
이름:
   singularity storage create jottacloud - Jottacloud

사용법:
   singularity storage create jottacloud [command options] [arguments...]

설명:
   --md5-memory-limit
      이 크기보다 큰 파일은 필요한 경우 MD5를 계산하기 위해 디스크에 캐시됩니다.

   --trashed-only
      휴지통에 있는 파일만 표시합니다.
      
      이는 휴지통에 있는 파일을 원래 디렉터리 구조로 표시합니다.

   --hard-delete
      파일을 휴지통에 넣지 않고 영구적으로 삭제합니다.

   --upload-resume-limit
      이 크기보다 큰 파일은 업로드 실패 시 재개될 수 있습니다.
   
   --no-versions
      파일을 덮어쓰는 대신 파일을 삭제하고 다시 생성하여 서버 측 버전 관리를 피합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --help, -h  도움말 표시

   고급

   --encoding value             백엔드의 인코딩입니다. (기본값: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete                파일을 휴지통에 넣지 않고 영구적으로 삭제합니다. (기본값: false) [$HARD_DELETE]
   --md5-memory-limit value     이 크기보다 큰 파일은 필요한 경우 MD5를 계산하기 위해 디스크에 캐시됩니다. (기본값: "10Mi") [$MD5_MEMORY_LIMIT]
   --no-versions                파일을 덮어쓰는 대신 파일을 삭제하고 다시 생성하여 서버 측 버전 관리를 피합니다. (기본값: false) [$NO_VERSIONS]
   --trashed-only               휴지통에 있는 파일만 표시합니다. (기본값: false) [$TRASHED_ONLY]
   --upload-resume-limit value  이 크기보다 큰 파일은 업로드 실패 시 재개될 수 있습니다. (기본값: "10Mi") [$UPLOAD_RESUME_LIMIT]

   일반

   --name value  저장소 이름 (기본값: 자동 생성)
   --path value  저장소 경로

```
{% endcode %}