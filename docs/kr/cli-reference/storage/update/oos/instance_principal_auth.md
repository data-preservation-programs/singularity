# 인스턴스 고유한 신원을 사용하여 API 호출 권한 부여
각 인스턴스는 자체 신원을 가지고 있으며 인스턴스 메타데이터에서 읽은 인증서를 사용하여 인증합니다.
https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

{% code fullWidth="true" %}
```
적용 방법:
   singularity storage update oos instance_principal_auth [command options] <name|id>

설명:
   --namespace
      Object storage 네임스페이스

   --compartment
      Object storage 컴파트먼트 OCID

   --region
      Object storage 리전

   --endpoint
      Object storage API의 엔드포인트.
      
      리전의 기본 엔드포인트를 사용하려면 공란으로 두십시오.

   --storage-tier
      새로운 객체를 저장할 때 사용할 스토리지 클래스입니다.
      https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예시:
         | Standard         | 기본적인 스토리지 티어
         | InfrequentAccess | 드문 접근 스토리지 티어
         | Archive          | 아카이브 스토리지 티어

   --upload-cutoff
      청크 업로드로 전환하는 기준값.
      
      이 값을 초과하는 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 크거나 크기를 알 수 없는 파일(예: "rclone rcat" 또는 "rclone mount" 또는 Google 또는 Google 문서에서 업로드된 파일)을 업로드할 때 이 청크 크기를 사용하여 
      멀티파트 업로드로 업로드됩니다.
      
      메모리 당 10000개의 chunk 크기와 "-P" 플래그와 함께 표시되는 진행 통계의 정확성이 감소됩니다.
      
      높은 속도의 링크에서 대용량 파일을 전송하고 메모리가 충분하면 이 값을 늘리면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 10000개의 청크 제한을 유지하려고 청크 크기를 자동으로 늘립니다.
      
      알 수없는 크기의 파일은 구성된 chunk_size로 업로드됩니다. chunk_size는 5 MiB로 기본값이며 10000개의 청크까지 있을 수 있으므로 기본값으로 스트림 업로드 할 수 있는 파일의 
      최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 통계의 정확성이 감소됩니다.
      

   --upload-concurrency
      멀티파트 업로드에 대한 동시성.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      높은 속도의 링크로 대량의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우 동시성을 늘리면 전송 속도가 향상될 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하는 기준값.
      
      서버 측 복사가 필요한 이 기준값보다 큰 파일은 이 크기로 청크 단위로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사에 대한 제한 시간.
      
      복사는 비동기 작업이며 제한 시간을 지정하여 복사가 성공할 때까지 대기합니다.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만 대용량 파일의 업로드 시작까지 
      오랜 지연을 유발할 수 있습니다.

   --encoding
      백엔드의 인코딩 방식입니다.
      
      자세한 내용은 [개요에서 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패한 경우 업로드 중단 요청을 호출하지 않고 S3에 모든 업로드된 첨부 파일을 수동으로 복구합니다.
      
      다른 세션 간에 업로드를 재개해야 할 경우 true로 설정하면 됩니다.
      
      경고: 불완전한 멀티파트 업로드의 부분을 저장하면 객체 스토리지의 공간 사용량에 추가되고 정리하지 않으면 추가 비용이 발생합니다.
      

   --no-check-bucket
      버킷의 존재 확인 또는 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한이 없을 경우 필요할 수도 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 객체와 연결된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함한 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예시:
         | <unset> | 없음

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화하거나 해독하는 데 사용할 base64로 인코딩된 256-bit 암호화 키를 지정하는 선택적 헤더입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.
      자세한 내용은 서버 측 암호화를 위해 자체 키 사용하기
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 참조

      예시:
         | <unset> | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값은 암호화 키의 무결성을 확인하는 데 사용됩니다. 서버 측 암호화를 위해 자체 키 사용하기
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 참조.

      예시:
         | <unset> | 없음

   --sse-kms-key-id
      본 사용자의 볼트에서 자체 마스터 키를 사용하는 경우, 이 헤더는 데이터 암호화 키를 생성하거나 암호화 키를 암호화하거나 해독하기 위해 
      키 관리 서비스를 호출하는 데 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예시:
         | <unset> | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다.
      Object Storage는 암호화 알고리즘으로 "AES256"을 지원합니다. 자세한 내용은
      서버 측 암호화를 위해 자체 키 사용하기 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 참조.

      예시:
         | <unset> | 없음
         | AES256  | AES256


옵션:
   --compartment value  Object storage 컴파트먼트 OCID [$COMPARTMENT]
   --endpoint value     Object storage API의 엔드포인트 [$ENDPOINT]
   --help, -h           도움말 표시
   --namespace value    Object storage 네임스페이스 [$NAMESPACE]
   --region value       Object storage 리전 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 기준값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사에 대한 제한 시간. (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩 방식. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패한 경우 업로드 중단 요청을 호출하지 않고 S3에서
   이미 업로드한 모든 파트를 수동으로 복구합니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷의 존재 확인 또는 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화하거나 해독하는 데 사용할 base64로 인코딩된 256-bit 암호화 키를 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 객체와 연결된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함한 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           본 사용자의 볼트에서 자체 마스터 키를 사용하는 경우, 이 헤더는 데이터 암호화 키를 생성하거나 암호화 키를 암호화하거나 해독하기 위해 키 관리 서비스를 호출하는 데 사용되는 마스터 암호화 키의 OCID를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             새로운 객체를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 기준값. (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}