# QingCloud 객체 스토리지

{% code fullWidth="true" %}
```
이름:
   singularity storage update qingstor - QingCloud 객체 스토리지

사용법:
   singularity storage update qingstor [command options] <name|id>

설명:
   --env-auth
      런타임에서 QingStor 자격 증명 가져오기.
      
      access_key_id와 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 QingStor 자격 증명을 입력하세요.
         | true  | 환경에서 QingStor 자격 증명 가져오기(환경 변수 또는 IAM).

   --access-key-id
      QingStor 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 두세요.

   --secret-access-key
      QingStor 비밀 액세스 키(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 두세요.

   --endpoint
      QingStor API에 연결할 엔드포인트 URL을 입력하세요.
      
      비워 둘 경우 기본 값인 "https://qingstor.com:443"을 사용합니다.

   --zone
      연결할 존을 입력하세요.
      
      기본값은 "pek3a"입니다.

      예시:
         | pek3a | 베이징(중국) 3존입니다.
         |       | 위치 제약 조건 pek3a가 필요합니다.
         | sh1a  | 상해(중국) 1존입니다.
         |       | 위치 제약 조건 sh1a가 필요합니다.
         | gd2a  | 광동(중국) 2존입니다.
         |       | 위치 제약 조건 gd2a가 필요합니다.

   --connection-retries
      연결 재시도 횟수입니다.

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값입니다.
      
      이 값보다 큰 파일은 chunk_size의 크기로 청크별 업로드됩니다.
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      "--qingstor-upload-concurrency" 크기의 청크는 전송 당 메모리에 버퍼링됩니다.
      
      대역폭을 충분히 사용할 수 있고 고속 링크로 큰 파일을 전송하는 경우에는
      이 값을 증가시키면 전송 속도가 빨라집니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      주의: 이 값을 1보다 크게 설정하면 멀티파트 업로드의 체크섬이 손상됩니다
      (그러나 업로드 자체는 손상되지 않습니다).
      
      대역폭을 충분히 사용하지 않는 고속 링크를 통해 소수의 큰 파일을 업로드하고 있는 경우에는
      이 값을 증가시켜 전송 속도를 개선할 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --access-key-id value      QingStor 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --endpoint value           QingStor API에 연결할 엔드포인트 URL입니다. [$ENDPOINT]
   --env-auth                 런타임에서 QingStor 자격 증명 가져오기 (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  QingStor 비밀 액세스 키(비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --zone value               연결할 존입니다. [$ZONE]

   고급

   --chunk-size value          업로드에 사용할 청크 크기입니다. (기본값: "4Mi") [$CHUNK_SIZE]
   --connection-retries value  연결 재시도 횟수입니다. (기본값: 3) [$CONNECTION_RETRIES]
   --encoding value            백엔드의 인코딩입니다. (기본값: "Slash,Ctl,InvalidUtf8") [$ENCODING]
   --upload-concurrency value  멀티파트 업로드에 대한 동시성입니다. (기본값: 1) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       청크 업로드로 전환하기 위한 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}