# 런타임(env)에서 자동으로 인증 자격 증명을 가져옵니다. 먼저 인증 자격 증명을 제공한 사람이 이기는 방식

{% code fullWidth="true" %}
```
이름:
   singularity storage create oos env_auth - 런타임(env)에서 자동으로 인증 자격 증명을 가져옵니다. 먼저 인증 자격 증명을 제공한 사람이 이기는 방식

사용법:
   singularity storage create oos env_auth [command options] [arguments...]

설명:
   --namespace
      오브젝트 스토리지 네임스페이스

   --compartment
      오브젝트 스토리지 컴파트먼트 OCID

   --region
      오브젝트 스토리지 리전

   --endpoint
      오브젝트 스토리지 API의 엔드포인트.
      
      리전의 기본 엔드포인트를 사용하려면 비워 두세요.

   --storage-tier
      스토리지에 새로운 오브젝트를 저장할 때 사용할 스토리지 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예제:
         | Standard         | 스탠다드 스토리지 클래스, 디폴트 티어입니다
         | InfrequentAccess | 아이템바탕접근 스토리지 클래스
         | Archive          | 아카이브 스토리지 클래스

   --upload-cutoff
      청크 업로드로 전환하기 위한 크기 제한입니다.
      
      이보다 큰 파일은 chunk_size의 크기로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나, 알 수 없는 크기의 파일("rclone rcat"이나 "rclone mount" 또는 Google 포토나 Google 문서에서 업로드된 파일) 업로드 시,
      이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      "upload_concurrency"만큼의 이 크기의 청크가 모든 전송마다 메모리에 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이를 더 크게 설정하면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 청크 크기를 자동으로 늘려서 10,000개의 청크 제한을 유지합니다.
      
      알 수 없는 크기의 파일은 설정된 chunk_size로 업로드됩니다. 디폴트 chunk_size는 5 MiB이며, 최대 10,000개의 청크가 가능하므로,
      기본 설정에서는 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상황 통계의 정확도가 낮아집니다.
      

   --upload-concurrency
      멀티파트 업로드를 위한 병렬처리 수입니다.
      
      동일한 파일의 청크 수가 동시에 업로드됩니다.
      
      대용량 파일을 고속 링크를 통해 작은 수의 업로드로 업로드하고 이 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 증가시키면 전송 속도를 향상시킬 수 있습니다.

   --copy-cutoff
      청크로 복사를 위한 크기 제한입니다.
      
      청크 단위로 복사해야 하는 이 크기보다 큰 파일은 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사에 대한 타임아웃입니다.
      
      복사는 비동기 작업이므로, 복사 완료를 기다리기 위해 타임아웃을 지정하세요.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다.
      이는 데이터 무결성 검사에 유용하지만, 큰 파일의 업로드 시작에 오랜 지연을 야기할 수 있습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --leave-parts-on-error
      실패 시 청크 업로드를 중단하지 않고 모든 성공적으로 업로드된 청크가 수동 복구를 위해 S3에 남아 있는지 여부입니다.
      
      다른 세션에서 업로드를 다시 시작해야 하는 경우, 이 값을 true로 설정하세요.
      
      경고: 완료되지 않은 멀티파트 업로드의 일부를 저장하면 객체 스토리지의 공간 사용 데이터에 포함되며,
      정리하지 않으면 추가 비용이 발생합니다.
      

   --no-check-bucket
      Bucket의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우, rclone에서 수행하는 트랜잭션 수를 최소화하려는 경우에 유용합니다.
      
      또는 사용 중인 사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면, 객체와 연관된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예제:
         | <unset> | None

   --sse-customer-key
      SSE-C를 사용하려면, 데이터를 암호화하거나 복호화하는 데 사용할 선택적 헤더로
      base64로 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 내용은
      [서버 측 암호화를 위한 고유 키 사용](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예제:
         | <unset> | None

   --sse-customer-key-sha256
      SSE-C를 사용하려면, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값은 암호화 키의 무결성을 확인하기 위해 사용됩니다.
      [서버 측 암호화를 위한 고유 키 사용](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예제:
         | <unset> | None

   --sse-kms-key-id
      보관소에서 자체 마스터 키를 사용하려면, 이 헤더는 데이터 암호화 키를 생성하거나 암호화 또는 복호화하는 데 사용하는
      마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예제:
         | <unset> | None

   --sse-customer-algorithm
      SSE-C를 사용하려면, 선택적 헤더로 암호화 알고리즘을 "AES256"으로 지정합니다.
      Object Storage는 "AES256"을 암호화 알고리즘으로 지원합니다. 자세한 내용은
      [서버 측 암호화를 위한 고유 키 사용](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하세요.

      예제:
         | <unset> | None
         | AES256  | AES256


옵션:
   --compartment value  오브젝트 스토리지 컴파트먼트 OCID [$COMPARTMENT]
   --endpoint value     오브젝트 스토리지 API의 엔드포인트 [$ENDPOINT]
   --help, -h           도움말 표시
   --namespace value    오브젝트 스토리지 네임스페이스 [$NAMESPACE]
   --region value       오브젝트 스토리지 리전 [$REGION]

   고급 옵션

   --chunk-size value               업로드에 사용할 청크 크기. (디폴트: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 복사를 위한 크기 제한. (디폴트: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사에 대한 타임아웃. (디폴트: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않음. (디폴트: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드에 대한 인코딩. (디폴트: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패 시 청크 업로드를 중단하지 않고 모든 성공적으로 업로드된 청크가 S3에 남아 있는지 여부. (디폴트: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않음. (디폴트: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하려면, 암호화 알고리즘으로 "AES256"을 선택적 헤더로 지정합니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면, 데이터를 암호화하거나 복호화하는 데 사용할 선택적 헤더로 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면, 객체와 연관된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하려면, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보관소에서 자체 마스터 키를 사용하려면, 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             스토리지에 새로운 오브젝트를 저장할 때 사용할 스토리지 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (디폴트: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드를 위한 병렬처리 수. (디폴트: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 크기 제한. (디폴트: "200Mi") [$UPLOAD_CUTOFF]

   일반 옵션

   --name value  스토리지의 이름 (디폴트: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}