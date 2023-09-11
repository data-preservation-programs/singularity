# API 호출에 리소스 주체 사용

{% code fullWidth="true" %}
```
명령어:
   singularity storage update oos resource_principal_auth - API 호출에 리소스 주체 사용

사용법:
   singularity storage update oos resource_principal_auth [command options] <name|id>

설명:
   --namespace
      오브젝트 스토리지 네임스페이스

   --compartment
      오브젝트 스토리지 컴파트먼트 OCID

   --region
      오브젝트 스토리지 리전

   --endpoint
      오브젝트 스토리지 API의 엔드포인트입니다.
      
      리전의 기본 엔드포인트를 사용하려면 비워 두십시오.

   --storage-tier
      새로운 오브젝트를 저장할 때 사용할 스토리지 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예제:
         | Standard         | 기본 티어인 구성 요소
         | InfrequentAccess | 드문 액세스 티어
         | Archive          | 아카이브 티어

   --upload-cutoff
      청크 형식으로 전환할 파일의 크기 제한입니다.
      
      이보다 큰 파일은 chunk_size 크기로 분할하여 업로드됩니다.
      최소값은 0이고 최대값은 5GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나 크기를 알 수 없는 파일 (예 : "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서)을
      전송하는 경우 이 청크 크기를 사용하여 여러 부분으로 업로드됩니다.
      
      이전 전송당 메모리에 upload_concurrency 크기의 청크가 버퍼링됩니다.
      
      대량 전송을 위해 대역폭이 높은 링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시키면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 10,000개의 청크 제한을 유지하려고 청크 크기를 자동으로 증가시킵니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 개의 청크가 있을 수 있으므로,
      기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를
      증가시켜야합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행률 통계의 정확도가 감소합니다.
      

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대역폭을 완전히 활용하지 못하는 고속 링크를 통해 소량의 큰 파일을 업로드하는 경우,
      이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환할 파일의 크기 제한입니다.
      
      복사해야하는 이보다 큰 파일은 이 크기로 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5GiB입니다.

   --copy-timeout
      복사에 대한 제한 시간입니다.
      
      복사는 비동기 작업이므로 성공을 위해 제한 시간을 지정하십시오.
      

   --disable-checksum
      개체 메타데이터와 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력 파일의 MD5 체크섬을 계산하여
      개체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만
      대용량 파일의 업로드를 시작하는 데 오랜 지연을 발생시킬 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패 시 업로드 중단 호출을 피하고 S3에 모두 성공적으로 업로드된 부분을 수동으로 복구합니다.
      
      다른 세션 간에 업로드를 계속하려면 true로 설정해야합니다.
      
      경고: 미완료된 멀티파트 업로드의 일부를 저장하면 개체 스토리지의 공간 사용량에 포함되며,
      청소되지 않으면 추가 비용이 발생할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않으려면 설정하십시오.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 갖고 있지 않은 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 개체와 연관된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예제:
         | <미설정> | 없음

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할
      선택적 헤더로서, base64로 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.
      자세한 내용은 Server-Side Encryption을 위해 직접 키 사용 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)을 참조하십시오.

      예제:
         | <미설정> | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값은 암호화 키의 무결성을 확인하는 데 사용됩니다. Server-Side Encryption을 위해 직접 키 사용 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예제:
         | <미설정> | 없음

   --sse-kms-key-id
      보루트에 고유한 마스터 키를 사용하는 경우,
      이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화 또는 복호화하기 위해 Key Management 서비스를 호출하기 위해
      사용되는 마스터 암호화 키의 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예제:
         | <미설정> | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 선택적 헤더로서 암호화 알고리즘으로 "AES256"을 지정합니다.
      Object Storage는 "AES256"을 지원하는 암호화 알고리즘입니다. 자세한 내용은
      Server-Side Encryption을 위해 직접 키 사용 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예제:
         | <미설정> | 없음
         | AES256  | AES256


옵션:
   --compartment value  오브젝트 스토리지 컴파트먼트 OCID [$COMPARTMENT]
   --endpoint value     오브젝트 스토리지 API의 엔드포인트 [$ENDPOINT]
   --help, -h           도움말 표시
   --namespace value    오브젝트 스토리지 네임스페이스 [$NAMESPACE]
   --region value       오브젝트 스토리지 리전 [$REGION]

   Advanced

   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환할 크기 제한. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사에 대한 제한 시간. (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               개체 메타데이터와 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패 시 업로드 중단 호출을 피하고 S3에 모두 성공적으로 업로드된 부분을 수동으로 복구합니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정합니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 base64로 인코딩된 256비트 암호화 키를 지정합니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 개체와 연관된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보루트에 고유한 마스터 키를 사용하는 경우 마스터 암호화 키의 OCID를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             새로운 오브젝트를 저장할 때 사용할 스토리지 클래스 (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 형식으로 전환할 파일의 크기 제한. (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}