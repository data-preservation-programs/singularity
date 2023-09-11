# OCI 사용자와 API 키를 사용하여 인증합니다.
인증에는 구성 파일에 가입 기관 OCID, 사용자 OCID, 리전, 경로 및 API 키의 지문을 입력해야 합니다.
https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage update oos user_principal_auth - OCI 사용자와 API 키를 사용하여 인증합니다.
                                                        인증에는 구성 파일에 가입 기관 OCID, 사용자 OCID, 리전, 경로 및 API 키의 지문을 입력해야 합니다.
                                                        https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

사용법:
   singularity storage update oos user_principal_auth [명령 옵션] <이름 또는 ID>

설명:
   --namespace
      객체 스토리지 이름 공간

   --compartment
      객체 스토리지 컴파트먼트 OCID

   --region
      객체 스토리지 리전

   --endpoint
      객체 스토리지 API의 엔드포인트
      
      리전의 기본 엔드포인트를 사용하려면 비워둡니다.

   --config-file
      OCI 구성 파일 경로

      예:
         | ~/.oci/config | oci 구성 파일 위치

   --config-profile
      OCI 구성 파일 내 프로파일 이름

      예:
         | Default | 기본 프로필 사용

   --storage-tier
      객체를 저장할 때 사용할 저장 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예:
         | Standard         | 표준 저장 클래스, 기본 클래스입니다.
         | InfrequentAccess | 기반보다 적은 액세스 저장 클래스
         | Archive          | 아카이브 저장 클래스

   --upload-cutoff
      청크 업로드로 전환하는 임계점입니다.
      
      이보다 큰 파일은 chunk_size 단위로 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일
      (예: "rclone rcat"에서 가져온 파일이나 "rclone mount"를 사용하여 업로드된 파일 또는 Google
      사진 또는 Google 문서)은이 청크 크기를 사용하여 대용량 업로드를 수행합니다.
      
      확인하려면 "upload_concurrency"개의 이 청크 크기의 청크가 이전에 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 대용량 파일을 전송 중이고 충분한 메모리가 있는 경우
      이를 증가하면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드 할 때
      10,000 번의 청크 제한을 초과하지 않도록 청크 크기를 자동으로 늘립니다.
      
      알 수없는 크기의 파일은 구성된
      청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대
      스트림 업로드 가능한 파일의 크기 48 GiB입니다. 큰 파일을 스트림으로 업로드하려면
      청크 크기를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상태의 정확성이 감소합니다.
      

   --upload-concurrency
      멀티파트 업로드를 위한 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대용량 파일을 고속 링크로 전송하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우
      이를 증가시켜 전송 속도를 향상시킬 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하는 임계점입니다.
      
      서버 측에서 복사해야 하는 이보다 큰 파일은
      이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사 시간 제한입니다.
      
      복사는 비동기 작업이므로 복사가 성공할 때까지 대기할 시간을 지정합니다.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가합니다. 데이터 무결성 확인에는 좋지만
      대용량 파일의 업로드 시작에는 긴 지연이 발생할 수 있습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패 시 업로드를 중단하지 않고 모든 성공적으로 업로드된 청크를 S3에 보존합니다.
      
      다른 세션 간에 업로드를 계속하기 위해 이를 true로 설정해야 합니다.
      
      경고: 불완전한 멀티파트 업로드의 일부를 객체 저장소의 공간 사용량으로 계산하며 추가
      삭제하지 않으면 추가 비용이 발생합니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않습니다.
      
      이미 버킷이 존재한다고 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 개체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예:
         | <unset> | 없음

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할 base64로 인코딩된 256비트 암호화 키를 지정하는 선택적 헤더입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 내용은
      내 키를 사용하여 서버 측 암호화 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 참조.

      예:
         | <unset> | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      암호화 키의 무결성을 확인하기 위해이 값을 사용합니다.
      내 키를 사용하여 서버 측 암호화 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 참조.

      예:
         | <unset> | 없음

   --sse-kms-key-id
      사용자 소유의 보루에 사용하는 사용자 소유 마스터 키를 사용하는 경우
      데이터 암호화 키를 생성하거나 암호화 또는 복호화하기 위해 키 관리 서비스를 호출하는 데 사용되는 마스터 암호화 키의 OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예:
         | <unset> | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다.
      객체 스토리지는 "AES256"을 암호화 알고리즘으로 지원합니다.
      자세한 내용은 서버 측 암호화에 사용자 키 사용 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)을 참조하십시오.

      예:
         | <unset> | 없음
         | AES256  | AES256


옵션:
   --compartment value     객체 스토리지 컴파트먼트 OCID [$COMPARTMENT]
   --config-file value     OCI 구성 파일 경로 (기본값: "~/.oci/config") [$CONFIG_FILE]
   --config-profile value  OCI 구성 파일 내 프로파일 이름 (기본값: "Default") [$CONFIG_PROFILE]
   --endpoint value        객체 스토리지 API의 엔드포인트 [$ENDPOINT]
   --help, -h              도움말 표시
   --namespace value       객체 스토리지 이름 공간 [$NAMESPACE]
   --region value          객체 스토리지 리전 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기 (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계점 (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사 시간 제한 (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다 (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드에 대한 인코딩 (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패 시 업로드를 중단하지 않고 모든 성공적으로 업로드된 청크를 S3에 보존합니다 (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않습니다 (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다 [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할 base64로 인코딩된 256비트 암호화 키를 지정합니다 [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다 [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다 [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           사용자 소유의 보루에 사용하는 사용자 소유 마스터 키를 사용하는 경우 [$SSE_KMS_KEY_ID]
   --storage-tier value             객체를 저장할 때 사용할 저장 클래스. (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드를 위한 동시성 (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계점 (기본값: "200Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}