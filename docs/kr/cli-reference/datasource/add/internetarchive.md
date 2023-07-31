# 인터넷 아카이브

{% code fullWidth="true" %}
```
이름:
   singularity datasource add internetarchive - 인터넷 아카이브

사용법:
   singularity datasource add internetarchive [command options] <dataset_name> <source_path>

설명:
   --internetarchive-access-key-id
      IAS3 액세스 키입니다.
      
      익명 액세스를 위해 비워 둡니다.
      여기에서 액세스 키를 얻을 수 있습니다: https://archive.org/account/s3.php

   --internetarchive-disable-checksum
      서버에 rclone이 계산한 MD5 체크섬으로 테스트를 요청하지 마십시오.
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 서버에 체크섬을 확인하도록 요청합니다.
      이는 데이터 무결성 확인에는 훌륭하지만 대용량 파일의 업로드 시작까지 긴 대기 시간을 야기할 수 있습니다.

   --internetarchive-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 확인하십시오.

   --internetarchive-endpoint
      IAS3 엔드포인트입니다.
      
      기본값으로 비워 둡니다.

   --internetarchive-front-endpoint
      인터넷 아카이브 프론트엔드의 호스트입니다.
      
      기본값으로 비워 둡니다.

   --internetarchive-secret-access-key
      IAS3 비밀 키(암호)입니다.
      
      익명 액세스를 위해 비워 둡니다.

   --internetarchive-wait-archive
      서버의 처리 작업(특히 아카이브 및 book_op)이 완료될 때까지 대기하는 시간 제한입니다.
      쓰기 작업 이후에 반영이 보장되어야 하는 경우에만 사용하세요.
      대기를 비활성화하려면 0입니다. 제한 시간 초과시 오류 없이 실행됩니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터 세트의 파일을 CAR 파일로 내보낸 후 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막으로 성공적인 스캔 이후 지난 시간이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비됨)

   인터넷 아카이브 옵션

   --internetarchive-access-key-id value      IAS3 액세스 키입니다. [$INTERNETARCHIVE_ACCESS_KEY_ID]
   --internetarchive-disable-checksum value   서버에 rclone이 계산한 MD5 체크섬으로 테스트를 요청하지 마십시오. (기본값: "true") [$INTERNETARCHIVE_DISABLE_CHECKSUM]
   --internetarchive-encoding value           백엔드의 인코딩입니다. (기본값: "Slash,LtGt,CrLf,Del,Ctl,InvalidUtf8,Dot") [$INTERNETARCHIVE_ENCODING]
   --internetarchive-endpoint value           IAS3 엔드포인트입니다. (기본값: "https://s3.us.archive.org") [$INTERNETARCHIVE_ENDPOINT]
   --internetarchive-front-endpoint value     인터넷 아카이브 프론트엔드의 호스트입니다. (기본값: "https://archive.org") [$INTERNETARCHIVE_FRONT_ENDPOINT]
   --internetarchive-secret-access-key value  IAS3 비밀 키(암호)입니다. [$INTERNETARCHIVE_SECRET_ACCESS_KEY]
   --internetarchive-wait-archive value       서버의 처리 작업(특히 아카이브 및 book_op)이 완료될 때까지 대기하는 시간 제한입니다. (기본값: "0s") [$INTERNETARCHIVE_WAIT_ARCHIVE]

```
{% endcode %}