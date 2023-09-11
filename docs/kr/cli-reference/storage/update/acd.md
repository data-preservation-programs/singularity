# Amazon Drive

{% code fullWidth="true" %}
```
이름:
   singularity storage update acd - Amazon Drive

사용법:
   singularity storage update acd [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 ID입니다.

      일반적으로 비워둡니다.

   --client-secret
      OAuth 클라이언트 시크릿입니다.

      일반적으로 비워둡니다.

   --token
      JSON 블롭으로 된 OAuth 액세스 토큰입니다.

   --auth-url
      인증 서버 URL입니다.

      기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL입니다.

      기본값을 사용하려면 비워둡니다.

   --checkpoint
      내부 폴링에 대한 체크포인트입니다 (디버그 용도).

   --upload-wait-per-gb
      실패한 완료 업로드 후 파일이 나타나는지 확인하기 위해 1 GiB 당 추가로 기다릴 시간입니다.

      때때로 Amazon Drive는 파일이 완전히 업로드되었음에도 불구하고 파일이 잠시 후에 나타나는 경우 오류가 발생합니다.
      이런 현상은 파일 크기가 1 GiB를 초과하는 경우에 종종 발생하며, 파일이 10 GiB보다 큰 경우에는 거의 항상 발생합니다.
      이 매개변수는 rclone이 파일이 나타나기를 기다리는 시간을 제어합니다.

      이 매개변수의 기본값은 1 GiB 당 3분이므로, 기본적으로 각 GiB 업로드별로 3분 동안 파일이 나타나기를 기다립니다.

      이 기능을 비활성화하려면 0으로 설정하십시오. 이 경우 rclone이 실패한 업로드를 다시 시도하지만 파일은 결국 올바르게 나타날 것입니다.

      이 값은 다양한 파일 크기의 많은 업로드를 관찰하여 경험적으로 결정되었습니다.

      이 상황에 대해 rclone이 수행하는 작업에 대해 자세한 정보를 보려면 "-v" 플래그로 업로드하십시오.

   --templink-threshold
      이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다.

      이 크기 이상 또는 이 크기의 파일은 "tempLink"를 통해 다운로드됩니다.
      이는 약 10 GiB보다 큰 파일의 다운로드를 차단하는 Amazon Drive의 문제를 해결하기 위한 것입니다.
      이 매개변수의 기본값은 9 GiB로 변경할 필요가 없을 것입니다.

      이 임계값 이상의 파일을 다운로드하려면 rclone은 내부 S3 스토리지에서 직접 임시 URL을 통해 파일을 다운로드하는 "tempLink"를 요청합니다.

   --encoding
      백엔드에 대한 인코딩입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID입니다. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 시크릿입니다. [$CLIENT_SECRET]
   --help, -h             도움말 표시

   고급

   --auth-url value            인증 서버 URL입니다. [$AUTH_URL]
   --checkpoint value          내부 폴링에 대한 체크포인트입니다 (디버그 용도). [$CHECKPOINT]
   --encoding value            백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --templink-threshold value  이 크기 이상의 파일은 tempLink를 통해 다운로드됩니다. (기본값: "9Gi") [$TEMPLINK_THRESHOLD]
   --token value               JSON 블롭으로 된 OAuth 액세스 토큰입니다. [$TOKEN]
   --token-url value           토큰 서버 URL입니다. [$TOKEN_URL]
   --upload-wait-per-gb value  실패한 완료 업로드 후 파일이 나타나는지 확인하기 위해 1 GiB 당 추가로 기다릴 시간입니다. (기본값: "3m0s") [$UPLOAD_WAIT_PER_GB]

```
{% endcode %}