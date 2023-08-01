# Amazon Drive

{% code fullWidth="true" %}
```
이름:
   singularity datasource add acd - Amazon Drive

사용법:
   singularity datasource add acd [command options] <데이터셋_이름> <소스_경로>

설명:
   --acd-auth-url
      인증 서버 URL.
      
      공란으로 두면 공급자의 기본값을 사용합니다.

   --acd-checkpoint
      내부 폴링을 위한 체크포인트 (디버그).

   --acd-client-id
      OAuth 클라이언트 ID.
      
      보통 공란으로 둡니다.

   --acd-client-secret
      OAuth 클라이언트 시크릿.
      
      보통 공란으로 둡니다.

   --acd-encoding
      백엔드의 인코딩.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --acd-templink-threshold
      이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다.
      
      이 크기 이상의 파일은 "tempLink"를 통해 다운로드됩니다.
      이는 약 10 GiB보다 큰 파일의 다운로드가 차단되는 문제를 해결하기 위한 것입니다.
      이 매개변수의 기본값은 9 GiB이며, 변경할 필요가 없을 것입니다.
      
      이 임계값을 초과하는 파일을 다운로드하기 위해 rclone은 백엔드의 S3 저장소로부터
      임시 URL을 통해 파일을 직접 다운로드하기 위해 "tempLink"를 요청합니다.

   --acd-token
      JSON 블롭 형식의 OAuth 액세스 토큰.

   --acd-token-url
      토큰 서버 URL.
      
      공란으로 두면 공급자의 기본값을 사용합니다.

   --acd-upload-wait-per-gb
      완료된 업로드 후에 파일이 표시될 때까지 GiB당 추가 대기 시간.
      
      때때로 Amazon Drive는 파일을 완전히 업로드한 후에 오류가 발생할 수 있지만,
      파일이 잠시 후에 나타날 수도 있습니다. 이런 경우는 파일의 크기가 1 GiB 이상일 때에
      종종 발생하며, 파일의 크기가 10 GiB보다 큰 경우에는 항상 발생합니다.
      이 매개변수는 파일이 나타날 때까지 rclone이 대기하는 시간을 제어합니다.
      
      이 매개변수의 기본값은 GiB당 3분이며, 기본적으로 1 GiB당 3분씩 대기해서
      파일이 나타나는지 확인합니다.
      
      이 기능을 비활성화하려면 0으로 설정하세요.
      이렇게 하면 rclone이 실패한 업로드를 다시 시도하지만,
      파일은 결국 제대로 나타날 것입니다.
      
      이 값을 결정하는 데에는 다양한 파일 크기로 업로드를 여러 번 시도하여
      경험적으로 결정되었습니다.
      
      이 상황에 대한 더 자세한 정보를 보려면 "-v" 플래그를 사용하여 rclone이
      수행하는 작업에 대한 정보를 확인하세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후에 데이터셋의 파일 삭제 (기본값: false)
   --rescan-interval value  이 시간 간격 이후에 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비)

   acd 옵션

   --acd-auth-url value            인증 서버 URL. [$ACD_AUTH_URL]
   --acd-client-id value           OAuth 클라이언트 ID. [$ACD_CLIENT_ID]
   --acd-client-secret value       OAuth 클라이언트 시크릿. [$ACD_CLIENT_SECRET]
   --acd-encoding value            백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ACD_ENCODING]
   --acd-templink-threshold value  이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다. (기본값: "9Gi") [$ACD_TEMPLINK_THRESHOLD]
   --acd-token value               JSON 블롭 형식의 OAuth 액세스 토큰. [$ACD_TOKEN]
   --acd-token-url value           토큰 서버 URL. [$ACD_TOKEN_URL]
   --acd-upload-wait-per-gb value  완료된 업로드 후에 파일이 표시될 때까지 GiB당 추가 대기 시간. (기본값: "3m0s") [$ACD_UPLOAD_WAIT_PER_GB]
```
{% endcode %}