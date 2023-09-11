# 인터넷 아카이브

{% code fullWidth="true" %}
```
이름:
   singularity storage create internetarchive - 인터넷 아카이브

사용법:
   singularity storage create internetarchive [옵션] [인수...]

설명:
   --access-key-id
      IAS3 액세스 키입니다.
      
      익명 액세스를 위해 비워 두세요.
      여기에서 액세스 키를 찾을 수 있습니다: https://archive.org/account/s3.php

   --secret-access-key
      IAS3 시크릿 키(암호)입니다.
      
      익명 액세스를 위해 비워 두세요.

   --endpoint
      IAS3 엔드포인트입니다.
      
      기본값을 사용하기 위해 비워 두세요.

   --front-endpoint
      인터넷 아카이브 프론트엔드 호스트입니다.
      
      기본값을 사용하기 위해 비워 두세요.

   --disable-checksum
      rclone이 계산한 MD5 체크섬과 서버에게 객체의 체크섬을 확인하도록 요청하지 않습니다.
      보통 rclone은 업로드하기 전에 입력 데이터의 MD5 체크섬을 계산해 서버에게 체크섬 확인을 요청합니다.
      이는 데이터 무결성 확인에 유용하지만 큰 파일의 업로드 시작에 장시간 대기 시간을 초래할 수 있습니다.

   --wait-archive
      서버의 처리 작업(특히 아카이브와 book_op 작업)이 완료될 때까지 대기하는 시간 제한입니다.
      쓰기 작업 이후에 반드시 반영되어야 하는 경우에만 활성화하세요.
      대기를 비활성화하려면 0을 입력하세요. 시간 초과 시 오류가 발생하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-key-id value      IAS3 액세스 키입니다. [$ACCESS_KEY_ID]
   --help, -h                 도움말 표시
   --secret-access-key value  IAS3 시크릿 키(암호)입니다. [$SECRET_ACCESS_KEY]

   고급

   --disable-checksum      rclone이 계산한 MD5 체크섬과 서버에게 확인을 요청하지 않습니다. (기본값: true) [$DISABLE_CHECKSUM]
   --encoding value        백엔드의 인코딩입니다. (기본값: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value        IAS3 엔드포인트입니다. (기본값: "https://s3.us.archive.org") [$ENDPOINT]
   --front-endpoint value  인터넷 아카이브 프론트엔드 호스트입니다. (기본값: "https://archive.org") [$FRONT_ENDPOINT]
   --wait-archive value    서버의 처리 작업(특히 아카이브와 book_op 작업)이 완료될 때까지 대기하는 시간 제한입니다. (기본값: "0s") [$WAIT_ARCHIVE]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}