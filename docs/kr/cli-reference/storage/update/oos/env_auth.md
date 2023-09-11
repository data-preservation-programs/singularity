# 자격 증명 credentials를 런타임(env)에서 자동으로 가져옵니다. 인증을 제공하는 첫 번째 자격 증명이 우선적으로 사용됩니다.

{% code fullWidth="true" %}
```
이름:
   singularity storage update oos env_auth - 자격 증명 credentials를 런타임(env)에서 자동으로 가져옵니다. 인증을 제공하는 첫 번째 자격 증명이 우선적으로 사용됩니다.

사용법:
   singularity storage update oos env_auth [command options] <이름|ID>

설명:
   --namespace
      Object storage 네임스페이스

   --compartment
      Object storage 컴파트먼트 OCID

   --region
      Object storage 지역

   --endpoint
      Object storage API의 엔드포인트
     
      해당 지역의 기본 엔드포인트를 사용하려면 비워두세요.

   --storage-tier
      새로운 객체를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예시:
         | 표준                     | 표준 스토리지 클래스, 기본 값입니다.
         | InfrequentAccess         | InfrequentAccess 스토리지 클래스
         | 아카이브                 | 아카이브 스토리지 클래스

   --upload-cutoff
      청크 업로드로 전환하는 임계값.
     
      이 기준값보다 큰 파일은 청크 크기단위로 업로드됩니다.
      최소값은 0이고, 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
     
      upload_cutoff보다 크거나 알 수 없는 크기의 파일(예: "rclone rcat"에서 생성된 파일, "rclone mount" 또는 Googlefoto 또는 Google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
     
      이전된 업로드 서비스가 메모리 단위로 buffer에 이 크기의 청크를 사용합니다.
     
      고속 링크에서 대용량 파일을 전송하고 메모리가 충분한 경우, 청크 크기를 늘리면 전송 속도가 향상됩니다.
     
      Rclone은 알려진 크기의 대용량 파일 업로드시 10,000 청크의 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
     
      알려진 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기가 5 MiB이고, 최대 10,000 청크를 사용할 수 있기 때문에 기본적으로 스트림 업로드할 수 있는 최대 파일 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
     
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상황 통계의 정확도가 감소합니다.
     

   --upload-concurrency
      멀티파트 업로드에 대한 동시성.
     
      동일한 파일의 청크 수와 동시에 업로드됩니다.
     
      고속 링크에서 속도를 최대한 활용하지 못하는 작은 수의 대용량 파일을 업로드하고 있는 경우, 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값.
     
      서버 측 복사를 해야 하는 이 임계값보다 큰 파일은 이 크기의 청크로 복사됩니다.
     
      최소값은 0이고, 최대값은 5 GiB입니다.

   --copy-timeout
      복사 대기 시간.
     
      복사는 비동기 작업이므로, 성공적인 복사를 위해 대기할 시간을 지정합니다.
      

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
     
      일반적으로 rclone은 업로드 전에 입력 데이터의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이렇게 하면 데이터 무결성을 확인할 수 있지만, 대용량 파일을 업로드하기 위해 오래 기다려야 하는 경우가 발생할 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
     
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --leave-parts-on-error
      true로 설정하면 업로드 실패시 업로드 중이던 모든 청크를 중단하지 않고 S3에 성공적으로 업로드된 모든 청크를 수동으로 복구할 수 있습니다.
     
      다른 세션간에 업로드 재개에 사용되도록 true로 설정해야 합니다.
     
      경고: 미완성 멀티파트 업로드의 부분들을 저장하면 객체 저장소에서 공유 공간 사용량에 포함되며, 정리하지 않은 경우 추가 비용이 발생할 수 있습니다.
      

   --no-check-bucket
      설정된 경우, 버킷의 존재 여부를 확인하거나 버킷을 생성하지 않습니다.
     
      버킷이 이미 존재하는 경우에 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
     
      사용자가 버킷 생성 권한을 가지지 않은 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하기 위해, 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예시:
         | 설정 안함 | 없음

   --sse-customer-key
      SSE-C를 사용하기 위해, 데이터를 암호화 또는 복호화하는 데 사용할 256비트 암호화 키의 base64로 인코딩된 선택적 헤더입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.
      자세한 정보는 서버 측 암호화를 위해 자체 키 사용
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예시:
         | 설정 안함 | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값은 암호화 키의 무결성을 확인하는 데 사용됩니다.
      서버 측 암호화를 위해 자체 키 사용
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예시:
         | 설정 안함 | 없음

   --sse-kms-key-id
      보관소에서 자체 마스터 키를 사용하는 경우, 이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화 또는 복호화하기 위해 키 관리 서비스를 호출하기 위해 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예시:
         | 설정 안함 | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다.
      Object Storage는 "AES256"을 암호화 알고리즘으로 지원합니다.
      자세한 정보는 서버 측 암호화를 위해 자체 키 사용
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예시:
         | 설정 안함 | 없음
         | AES256  | AES256


옵션:
   --compartment value  Object Storage 컴파트먼트 OCID [$COMPARTMENT]
   --endpoint value     Object Storage API의 엔드포인트 [$ENDPOINT]
   --help, -h           도움말 표시
   --namespace value    Object Storage 네임스페이스 [$NAMESPACE]
   --region value       Object Storage 지역 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사 대기 시간. (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           true로 설정하면 업로드 실패시 업로드 중이던 모든 청크를 중단하지 않고 S3에 성공적으로 업로드된 모든 청크를 수동으로 복구할 수 있습니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                설정된 경우, 버킷의 존재 여부를 확인하거나 버킷을 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하기 위해, 데이터를 암호화 또는 복호화하는 데 사용할 256비트 암호화 키의 base64로 인코딩된 선택적 헤더입니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하기 위해, 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보관소에서 자체 마스터 키를 사용하는 경우, 이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화 또는 복호화하기 위해 키 관리 서비스를 호출하기 위해 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             새로운 객체를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}