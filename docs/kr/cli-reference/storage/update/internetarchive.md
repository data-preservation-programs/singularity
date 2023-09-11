# 인터넷 아카이브

{% code fullWidth="true" %}
```
이름:
   singularity storage update internetarchive - 인터넷 아카이브

사용법:
   singularity storage update internetarchive [command options] <name|id>

설명:
   --access-key-id
      IAS3 액세스 키입니다.
      
      익명 액세스를 위해 비워 둡니다.
      여기에서 액세스 키를 찾을 수 있습니다: https://archive.org/account/s3.php

   --secret-access-key
      IAS3 비밀 키(암호)입니다.
      
      익명 액세스를 위해 비워 둡니다.

   --endpoint
      IAS3 엔드포인트입니다.
      
      기본 값을 사용하려면 비워 둡니다.

   --front-endpoint
      인터넷 아카이브 프론트엔드 호스트입니다.
      
      기본 값을 사용하려면 비워 둡니다.

   --disable-checksum
      rclone이 계산한 MD5 체크섬과 서버에 저장된 객체를 테스트하는 것을 요청하지 않습니다.
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 서버에 체크섬을 확인할 수 있습니다.
      이는 데이터 무결성 확인에 유용하지만 큰 파일의 업로드 시작까지 오랜 지연을 초래할 수 있습니다.

   --wait-archive
      서버의 처리 작업(특히 아카이브 및 book_op)이 완료될 때까지 대기하는 시간 제한입니다.
      쓰기 작업 이후에 반영된 결과가 보장되어야 하는 경우에만 사용합니다.
      대기를 비활성화하려면 0을 입력합니다. 시간 초과 시 오류가 발생하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-key-id value      IAS3 액세스 키입니다. [$ACCESS_KEY_ID]
   --help, -h                 도움말 표시
   --secret-access-key value  IAS3 비밀 키(암호)입니다. [$SECRET_ACCESS_KEY]

   고급 옵션:

   --disable-checksum      rclone이 계산한 MD5 체크섬과 서버에 저장된 객체를 테스트하지 않습니다. (기본값: true) [$DISABLE_CHECKSUM]
   --encoding value        백엔드의 인코딩입니다. (기본값: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3 엔드포인트입니다. (기본값: "https://s3.us.archive.org") [$ENDPOINT]
   --front-endpoint value  인터넷 아카이브 프론트엔드 호스트입니다. (기본값: "https://archive.org") [$FRONT_ENDPOINT]
   --wait-archive value    서버의 처리 작업(특히 아카이브 및 book_op)이 완료될 때까지 대기하는 시간 제한입니다. (기본값: "0s") [$WAIT_ARCHIVE]
```
{% endcode %}