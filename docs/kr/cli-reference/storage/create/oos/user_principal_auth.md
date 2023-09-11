# 사용자 인증을 위해 OCI 사용자 및 API 키를 사용합니다.
테넌시 OCID, 사용자 OCID, 리전, 경로 및 API 키의 fingerprint를 구성 파일에 입력해야 합니다.
https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

{% code fullWidth="true" %}
```
NAME:
   singularity storage create oos user_principal_auth - 사용자 인증을 위해 OCI 사용자 및 API 키를 사용합니다.
                                                        테넌시 OCID, 사용자 OCID, 리전, 경로 및 API 키의 fingerprint를 구성 파일에 입력해야 합니다.
                                                        https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm

USAGE:
   singularity storage create oos user_principal_auth [command options] [arguments...]

DESCRIPTION:
   --namespace
      객체 스토리지 네임스페이스

   --compartment
      객체 스토리지 compartment OCID

   --region
      객체 스토리지 리전

   --endpoint
      객체 스토리지 API의 엔드포인트입니다.
      
      리전의 기본 엔드포인트를 사용하려면 비워 두십시오.

   --config-file
      OCI 구성 파일의 경로

      예:
         | ~/.oci/config | oci 구성 파일 위치

   --config-profile
      OCI 구성 파일 내 프로파일 이름

      예:
         | Default | 기본 프로파일 사용

   --storage-tier
      객체 저장소에 새로운 객체를 저장할 때 사용할 저장 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예:
         | Standard | 표준 저장 등급으로, 이것이 기본 등급입니다.
         | InfrequentAccess | 출입 횟수가 적은 저장 등급
         | Archive | 아카이브 저장 등급

   --upload-cutoff
      청크 업로드로 전환하는 기준.
      
      이 기준을 초과하는 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일은 이 청크 크기를 사용하여 여러 부분으로 업로드됩니다.
      알려지지 않은 크기의 파일(예: "rclone rcat"에서 사용되는 파일이나 "rclone mount" 또는 Google 사진이나 Google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 multipart 업로드로 업로드됩니다.
      
      주의: 각 전송당 "upload_concurrency" 청크가 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값(청크 크기)을 늘리면 전송 속도가 향상됩니다.
      
      Rclone은 라스트출 규 카운트(10,000 청크 제한)를 유지하기 위해 알려진 크기의 큰 파일을 업로드하는 경우 자동으로 청크 크기를 늘립니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 청크까지 있을 수 있으므로 기본적으로 파일 스트림 업로드 가능한 최대 파일 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 크게 늘려야 합니다.
      
      청크 크기를 증가시킬수록 "-P" 플래그로 표시되는 진행 통계의 정확성이 감소합니다.
      

   --upload-concurrency
      멀티파트 업로드의 동시성.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      고속 링크를 통해 작은 수의 대용량 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --copy-cutoff
      청크 복사로 전환하는 기준.
      
      서버 쪽 복사가 필요한 최소한 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사 제한 시간.
      
      복사는 비동기 작업이므로 복사가 성공할 때까지 대기할 제한 시간을 지정합니다.
      

   --disable-checksum
      객체 메타데이터와 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만 큰 파일을 업로드하는 데 오랜 시간이 걸릴 수 있습니다.

   --encoding
      백엔드의 인코딩 방식.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패 시 업로드를 중단하지 않고 모든 성공적으로 업로드된 일부를 수동으로 회수하기 위해 S3에 그대로 두는 경우 true로 설정하십시오.
      
      이 경우 다른 세션에서 업로드를 이어서 진행해야 합니다.
      
      경고: 완료하지 못한 멀티파트 업로드의 일부를 저장공간에 보관하면 객체 스토리지의 공간 사용량에 포함되어 추가 비용이 발생할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않으려면 설정하십시오.
      
      버킷이 이미 존재하는 경우 rclone의 수행 횟수를 최소화하려는 경우 유용할 수 있습니다.
      
      또한 사용자가 버킷 생성 권한을 갖고 있지 않은 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하기 위해 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 담은 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예:
         | <unset> | 없음

   --sse-customer-key
      SSE-C를 사용하기 위해, 데이터를 암호화 또는 복호화하는 데 사용할 선택적 헤더로 base64로 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 내용은
      Using Your Own Keys for Server-Side Encryption 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우, 선택사항 헤더로 암호화 키의 base64로 인코딩된 SHA256 해시를 지정합니다.
      이 값은 암호화 키의 무결성을 확인하는 데 사용됩니다. Using Your Own Keys for 
      Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음

   --sse-kms-key-id
      사용자 고유의 마스터 키를 자체에서 사용하는 경우에는 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)로 지정된 마스터 암호화 키의 헤더입니다. 
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예:
         | <unset> | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다.
      객체 스토리지는 "AES256"을 암호화 알고리즘으로 지원합니다. 자세한 내용은
      Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음
         | AES256  | AES256


OPTIONS:
   --compartment value     객체 스토리지 compartment OCID 값 [$COMPARTMENT]
   --config-file value     OCI 구성 파일의 경로 (기본값: "~/.oci/config") [$CONFIG_FILE]
   --config-profile value  OCI 구성 파일 내 프로파일 이름 (기본값: "Default") [$CONFIG_PROFILE]
   --endpoint value        객체 스토리지 API의 엔드포인트입니다. [$ENDPOINT]
   --help, -h              도움말 표시
   --namespace value       객체 스토리지 네임스페이스 값 [$NAMESPACE]
   --region value          객체 스토리지 리전 값 [$REGION]

   고급

   --chunk-size value               청크 업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크 복사로 전환하는 기준입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사 제한 시간입니다. (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               객체 메타데이터와 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩 방식입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패 시 업로드를 중단하지 않고 모든 성공적으로 업로드된 일부를 S3에서 수동 복구하기 위해 true로 설정하십시오. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않으려면 설정하십시오. (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하기 위해 데이터를 암호화 또는 복호화하는 데 사용할 선택적 헤더로 base64로 인코딩된 256비트 암호화 키를 지정합니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하기 위해 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 담은 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택사항 헤더입니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           사용자 고유의 마스터 키를 자체에서 사용하는 경우에는 OCID로 지정된 마스터 암호화 키의 헤더입니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             객체 저장소에 새로운 객체를 저장할 때 사용할 저장 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 기준입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]

   공통

   --name value  스토리지 이름 (기본값: 자동 생성)
   --path value  스토리지 경로

```
{% endcode %}