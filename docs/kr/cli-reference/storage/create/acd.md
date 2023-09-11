# Amazon Drive

{% code fullWidth="true" %}
```
이름:
   singularity storage create acd - Amazon Drive

사용법:
   singularity storage create acd [command options] [arguments...]

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
      
      제공 업체의 기본 설정을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL.
      
      제공 업체의 기본 설정을 사용하려면 비워두십시오.

   --checkpoint
      내부 폴링용 체크포인트 (디버그).

   --upload-wait-per-gb
      완료된 업로드 실패 후 파일이 표시되는지 확인하기 위해 GiB당 추가 대기 시간.
      
      가끔씩 파일이 완전히 업로드되었음에도 불구하고 Amazon Drive에서 오류가 발생할 수 있으며, 잠시 후에 파일이 표시될 때가 있습니다. 파일 크기가 1 GiB를 초과하는 경우에는 이러한 상황이 종종 발생하며, 파일 크기가 10 GiB를 초과하는 경우에는 거의 항상 발생합니다. 이 매개변수는 rclone이 파일이 표시되기까지 기다리는 시간을 제어합니다.
      
      이 매개변수의 기본값은 GiB당 3분입니다. 따라서 기본적으로 GiB당 3분 동안 파일이 업로드될 때까지 기다립니다.
      
      이 기능을 사용하지 않으려면 값을 0으로 설정하십시오. 이렇게 하면 rclone이 실패한 업로드를 다시 시도하지만 파일이 최종적으로 올바르게 나타날 가능성이 있다는 충돌 오류가 발생할 수 있습니다.
      
      이러한 값은 다양한 파일 크기에 대한 많은 업로드를 관찰함으로써 경험적으로 결정되었습니다.
      
      이 상황에서 rclone이 수행하는 작업에 대한 자세한 정보를 보려면 "-v" 플래그로 업로드하십시오.

   --templink-threshold
      이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다.
      
      이 크기 이상의 파일은 "tempLink"를 통해 다운로드됩니다. 이는 약 10 GiB보다 큰 파일을 다운로드할 수 없는 Amazon Drive의 문제를 해결하기 위한 것입니다. 기본값은 9 GiB이며, 변경할 필요가 없어야 합니다.
      
      이 임계 값을 초과하는 파일을 다운로드하기 위해 rclone은 "tempLink"를 요청하고 기본 S3 스토리지에서 직접 임시 URL을 통해 파일을 다운로드합니다.

   --encoding
      백엔드용 인코딩.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시

   고급 옵션

   --auth-url value            인증 서버 URL. [$AUTH_URL]
   --checkpoint value          내부 폴링용 체크포인트 (디버그). [$CHECKPOINT]
   --encoding value            백엔드용 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --templink-threshold value  이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다. (기본값: "9Gi") [$TEMPLINK_THRESHOLD]
   --token value               JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value           토큰 서버 URL. [$TOKEN_URL]
   --upload-wait-per-gb value  완료된 업로드 실패 후 파일이 표시되는지 확인하기 위한 GiB당 추가 대기 시간. (기본값: "3m0s") [$UPLOAD_WAIT_PER_GB]

   일반 옵션

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}