# QingCloud Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create qingstor - 고려클라우드 객체 스토리지

사용법:
   singularity storage create qingstor [command options] [arguments...]

설명:
   --env-auth
      런타임에서 QingStor 자격 증명 가져오기.
      
      access_key_id와 secret_access_key가 비어있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 QingStor 자격 증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 QingStor 자격 증명 가져오기.

   --access-key-id
      QingStor Access Key ID입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 두세요.

   --secret-access-key
      QingStor Secret Access Key (비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 두세요.

   --endpoint
      연결할 QingStor API의 엔드포인트 URL을 입력하세요.
      
      비워 두면 기본값인 "https://qingstor.com:443"을 사용합니다.

   --zone
      연결할 가상 데이터 위치입니다.
      
      기본값은 "pek3a"입니다.

      예시:
         | pek3a | 베이징(중국) 3존입니다.
         |       | 위치 제약 조건은 "pek3a"가 필요합니다.
         | sh1a  | 상하이(중국) 1존입니다.
         |       | 위치 제약 조건은 "sh1a"가 필요합니다.
         | gd2a  | 광동(중국) 2존입니다.
         |       | 위치 제약 조건은 "gd2a"가 필요합니다.

   --connection-retries
      연결 재시도 횟수입니다.

   --upload-cutoff
      청킹 업로드로 전환하는 커트 오프입니다.
      
      이보다 큰 파일은 chunk_size의 크기로 청크 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드할 때 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일은 이 청크 크기를 사용하여 다중 파트 업로드됩니다.
      
      주의: "--qingstor-upload-concurrency"에서 이 크기만큼의 청크가 전송마다 메모리에 버퍼링됩니다.
      
      대량의 파일을 고속 링크로 전송하고 충분한 메모리가 있는 경우, 이 크기를 늘리면 전송 속도가 향상됩니다.

   --upload-concurrency
      다중 파트 업로드의 동시성입니다.
      
      파일의 동일한 청크 수만큼 동시에 업로드됩니다.
      
      참고: 이 값을 1보다 크게 설정하면 다중 파트 업로드의 체크섬이 손상됩니다(물론 업로드 자체는 손상되지 않습니다).
      
      고속 링크를 통해 작은 수의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 확인하세요.


옵션:
   --access-key-id value      QingStor Access Key ID입니다. [$ACCESS_KEY_ID]
   --endpoint value           연결할 QingStor API의 엔드포인트 URL을 입력하세요. [$ENDPOINT]
   --env-auth                 런타임에서 QingStor 자격 증명 가져오기(기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  QingStor Secret Access Key (비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --zone value               연결할 가상 데이터 위치입니다. [$ZONE]

   고급

   --chunk-size value          업로드할 때 사용할 청크 크기입니다(기본값: "4Mi") [$CHUNK_SIZE]
   --connection-retries value  연결 재시도 횟수입니다(기본값: 3) [$CONNECTION_RETRIES]
   --encoding value            백엔드의 인코딩입니다(기본값: "Slash,Ctl,InvalidUtf8") [$ENCODING]
   --upload-concurrency value  다중 파트 업로드의 동시성입니다(기본값: 1) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       청킹 업로드로 전환하는 커트 오프입니다(기본값: "200Mi") [$UPLOAD_CUTOFF]

   일반

   --name value  스토리지의 이름(기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}