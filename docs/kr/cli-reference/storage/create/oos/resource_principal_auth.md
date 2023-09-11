# API 호출을 위해 리소스 주체 사용

{% code fullWidth="true" %}
```
이름:
   singularity storage create oos resource_principal_auth - API 호출을 위해 리소스 주체 사용

사용법:
   singularity storage create oos resource_principal_auth [command options] [arguments...]

설명:
   --namespace
      객체 저장소 네임스페이스

   --compartment
      객체 저장소 구획 OCID

   --region
      객체 저장소 리전

   --endpoint
      객체 저장소 API의 엔드포인트.
      
      리전의 기본 엔드포인트를 사용하려면 공란으로 둡니다.

   --storage-tier
      새 객체를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예:
       | Standard         | 표준 스토리지 계층, 기본 계층입니다.
       | InfrequentAccess | 자주 사용되지 않는 스토리지 계층
       | Archive          | 아카이브 스토리지 계층

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계치.
      
      이보다 큰 파일은 chunk_size 단위로 청크를 업로드합니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일(예: "rclone rcat"에서 온 파일이나 "rclone mount"로 업로드된 파일, Google 사진이나 Google 문서 등)은 이 청크 크기를 사용하여 
      멀티파트 업로드로 업로드됩니다.
      
      한 번에 transfer당 메모리에 upload_concurrency 크기의 청크가 버퍼링됩니다.
      
      고속 링크로 대용량 파일을 전송하고 메모리가 충분한 경우 이 값을 늘리면 전송 속도가 높아집니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드 할 때 10,000개의 청크 제한을 유지하기 위해 청크 크기를 자동으로 증가시킵니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크가 가능하므로 기본적으로 stream 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 
      더 큰 파일을 stream 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행률 통계의 정확도가 낮아집니다.
      

   --upload-concurrency
      멀티파트 업로드에 대한 동시성.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대용량 파일을 고속 링크로 작은 수의 파일을 업로드하는 경우에는 이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계치.
      
      서버 측 복사가 필요한 이보다 큰 파일은 이 크기로 청크 단위로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사 시간제한.
      
      복사는 비동기 작업이므로 복사 성공을 기다리기 위한 제한 시간을 지정합니다.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력 파일의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에는 좋지만 대용량 파일의 시작 부분에 대한 오랜 지연을 초래할 수 있습니다.

   --encoding
      백엔드의 인코딩.
      
      더 많은 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패시 모든 성공적으로 업로드된 파트를 S3에 남기고 업로드를 중단하지 않도록 설정합니다.
      
      이 기능은 다른 세션 간에 업로드를 다시 시작해야 할 때 true로 설정해야 합니다.
      
      경고: 미완료된 멀티파트 업로드의 일부를 저장하면 객체 저장소의 공간 사용량에 영향을 줍니다. 청소하지 않으면 추가 비용이 발생할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 줄이려는 경우에 유용합니다.
      
      사용자가 버킷 생성 권한을 갖고 있지 않은 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 객체와 관련된 AES-256 암호화 키의 base64 인코딩된 문자열이 포함된 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예:
         | <unset> | None

   --sse-customer-key
      SSE-C를 사용하려면 데이터의 암호화 또는 복호화에 사용할 선택적 헤더로서의 base64 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 내용은
      서버 측 암호화에 직접 키 사용(https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우 암호화 키의 base64 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값을 사용하여 암호화 키의 무결성을 확인합니다. 서버 측 암호화에 직접 키 사용(https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | None

   --sse-kms-key-id
      보관소에서 고유한 마스터 키를 사용하는 경우이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화하거나 복호화하기 위해 Key Management 서비스를 호출하기 위해 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예:
         | <unset> | None

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 "AES256"을 암호화 알고리즘으로 지정하는 선택적 헤더입니다.
      객체 저장소는 "AES256"을 암호화 알고리즘으로 지원합니다. 자세한 내용은
      서버 측 암호화에 직접 키 사용(https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | None
         | AES256  | AES256


옵션:
   --compartment value  객체 저장소 구획 OCID [$COMPARTMENT]
   --endpoint value     객체 저장소 API의 엔드포인트 [$ENDPOINT]
   --help, -h           도움말 표시
   --namespace value    객체 저장소 네임스페이스 [$NAMESPACE]
   --region value       객체 저장소 리전 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기 (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계치 (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사 시간제한 (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다 (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩 (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패시 모든 성공적으로 업로드된 파트를 S3에 남기고 업로드를 중단하지 않도록 설정합니다 (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않습니다 (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로서 "AES256"을 지정하는 선택적 헤더입니다 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터의 암호화 또는 복호화에 사용할 선택적 헤더로서의 base64 인코딩된 256비트 암호화 키를 지정합니다 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 객체와 관련된 AES-256 암호화 키의 base64 인코딩된 문자열이 포함된 파일입니다 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 base64 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보관소에서 고유한 마스터 키를 사용하는 경우이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화하거나 복호화하기 위해 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다 [$SSE_KMS_KEY_ID]
   --storage-tier value             새 객체를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성 (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계치 (기본값: "200Mi") [$UPLOAD_CUTOFF]

   일반

   --name value  저장소 이름 (기본값: 자동 생성)
   --path value  저장소 경로

```
{% endcode %}