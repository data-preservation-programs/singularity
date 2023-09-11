# 자격증명이 필요하지 않으며, 일반적으로 공개 버킷을 읽기 위한 용도입니다.

{% code fullWidth="true" %}
```
이름:
   singularity storage update oos no_auth - 자격증명이 필요하지 않으며, 일반적으로 공개 버킷을 읽기 위한 용도입니다.

사용법:
   singularity storage update oos no_auth [command options] <name|id>

설명:
   --namespace
      객체 저장소 네임스페이스

   --region
      객체 저장소 지역

   --endpoint
      객체 저장소 API의 엔드포인트.
      
      지역의 기본 엔드포인트를 사용하려면 비워두세요.

   --storage-tier
      저장소에 새로운 객체를 저장할 때 사용할 저장 등급입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예시:
         | Standard         | 표준 저장 등급, 기본 등급입니다.
         | InfrequentAccess | InfrequentAccess 저장 등급
         | Archive          | Archive 저장 등급

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값입니다.
      
      이보다 큰 파일은 청크 크기로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat"으로부터 또는 "rclone mount" 또는 google
      photos 또는 google docs로 업로드된)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      "upload_concurrency" 크기의 청크가 전송마다 메모리에 버퍼링됩니다.
      
      고속링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 늘리면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드할 때 10,000개의 청크 제한을 유지하기 위해 청크 크기를 자동으로 늘립니다.
      
      알려진 크기의 파일은 설정된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로
      기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 통계의 정확도가 감소합니다.
      

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대역폭을 완전히 활용하지 않는 고속링크에서 작은 수의 대형 파일을 업로드하는 경우 이 값을 늘리면 전송 속도가 향상될 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.
      
      서버 측에서 복사할 필요가 있는 이 이상 큰 파일은 이 크기로 청크 단위로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사에 대한 타임아웃입니다.
      
      복사는 비동기 작업이므로 복사 성공을 기다릴 타임아웃을 지정합니다.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가하므로 데이터 무결성 확인에 적합하지만 대형 파일의
      업로드 시작까지 시간이 오래 걸릴 수 있습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패한 경우에서는 업로드를 중단하지 않고 S3에 성공적으로 업로드된 모든 청크를 수동으로 복구합니다.
      
      서로 다른 세션에서 업로드를 재개해야 하는 경우 이 값을 true로 설정하십시오.
      
      경고: 완료되지 않은 멀티파트 업로드의 구성 요소를 보관하면 객체 저장소의 공간 사용량에 포함되며
      청소하지 않으면 추가 비용이 발생할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않도록 설정합니다.
      
      알려진 버킷이 이미 존재하는 경우에 rclone이 수행하는 트랜잭션 수를 최소화하려는 데 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 객체와 연결된 AES-256 암호화 키의 base64로 인코딩된 문자열이 포함된 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예시:
         | <unset> | None

   --sse-customer-key
      SSE-C를 사용하려면, 데이터를 암호화 또는 복호화하는 데 사용할 선택적인 헤더로, 암호화 키로서 base64로 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 내용은 
      사용자의 키를 사용한 [서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예시:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시 값을 지정하는 선택적인 헤더입니다.
      이 값은 암호화 키의 무결성을 확인하는 데 사용됩니다. [서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예시:
         | <unset> | None

   --sse-kms-key-id
      보관함에서 고유한 마스터 키를 사용하는 경우, 이 헤더는 데이타 암호화 키를 생성하거나 데이타 암호화 키를 암호화 또는 복호화하기 위해 키 관리 서비스를 호출하는 데 사용된 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예시:
         | <unset> | None

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"를 지정하는 선택적인 헤더입니다.
      Object Storage는 "AES256"을 암호화 알고리즘으로 지원합니다. 자세한 내용은
      사용자의 키를 사용한 [서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예시:
         | <unset> | None
         | AES256  | AES256


옵션:
   --endpoint value   객체 저장소 API의 엔드포인트 [$ENDPOINT]
   --help, -h         도움말 표시
   --namespace value  객체 저장소 네임스페이스 [$NAMESPACE]
   --region value     객체 저장소 지역 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기 (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계값 (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사에 대한 타임아웃 (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않음 (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩 (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패한 경우에서도 업로드를 중단하지 않고 S3에 성공적으로 업로드된 모든 청크를 수동으로 복구합니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않도록 설정 (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"를 지정하는 선택적인 헤더입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면, 데이터를 암호화 또는 복호화하는 데 사용할 선택적인 헤더로, 암호화 키로서 base64로 인코딩된 256비트 암호화 키를 지정합니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 객체와 연결된 AES-256 암호화 키의 base64로 인코딩된 문자열이 포함된 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시 값을 지정하는 선택적인 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보관함에서 고유한 마스터 키를 사용하는 경우 이 헤더는 데이타 암호화 키를 생성하거나 데이타 암호화 키를 암호화 또는 복호화하기 위해 키 관리 서비스를 호출하는 데 사용된 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             저장소에 새로운 객체를 저장할 때 사용할 저장 등급. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성 (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계값 (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}