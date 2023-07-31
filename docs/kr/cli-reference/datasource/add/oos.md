# Oracle Cloud Infrastructure Object Storage

{% code fullWidth="true" %}
```
이름:
   Singularity 데이터 소스 추가 oos - Oracle Cloud Infrastructure Object Storage

사용법:
   singularity datasource add oos [명령 옵션] <데이터 세트 이름> <소스 경로>

설명:
   --oos-chunk-size
      업로드에 사용할 청크 크기.
      
      업로드 허용값(upload_cutoff)보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat" 또는 "rclone mount" 또는 google
      photos 또는 google docs에서 가져온 파일)의 경우, 이 청크 크기를 사용하여 멀티파트 업로드로 전송됩니다.
      
      참고로, 전송마다 "upload_concurrency" 개의 이 청크 크기가 메모리에 버퍼링됩니다.
      
      높은 속도의 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시키면 전송 속도가 높아집니다.
      
      Rclone은 알려진 큰 파일을 업로드할 때 10,000개의 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      
      크기를 알 수 없는 파일은 구성된 chunk_size로 업로드됩니다. 기본적인 청크 크기는 5MiB이며, 한 번에 업로드할 수 있는
      청크의 최대 개수는 10,000개이므로, 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48GiB입니다. 더 큰 파일을
      스트림으로 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태 통계의 정확도가 낮아집니다.
      

   --oos-compartment
      [공급자] - user_principal_auth
         객체 저장 컴파트먼트 OCID

   --oos-config-file
      [공급자] - user_principal_auth
         OCI 구성 파일 경로

         예시:
            | ~/.oci/config | oci 구성 파일 위치

   --oos-config-profile
      [공급자] - user_principal_auth
         OCI 구성 파일 내 프로필 이름

         예시:
            | Default | 기본 프로필 사용

   --oos-copy-cutoff
      멀티파트 복사를 위한 커트오프 값.
      
      이 값을 초과하는 파일을 서버 측에서 복사해야 할 경우, 이 크기로 청크별로 복사가 수행됩니다.
      
      최소값은 0이고, 최대값은 5GiB입니다.

   --oos-copy-timeout
      복사를 위한 대기 시간.
      
      복사는 비동기 작업이므로, 복사가 성공하기를 기다릴 시간을 지정하세요.
      

   --oos-disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      보통 rclone은 업로드하기 전에 입력 파일의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다.
      이는 데이터 무결성 검사에 유용하지만, 대용량 파일의 업로드 시작이 지연되는 경우가 있습니다.

   --oos-encoding
      백엔드에 대한 인코딩입니다.
      
      더 자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --oos-endpoint
      객체 스토리지 API의 엔드포인트입니다.
      
      기본 엔드포인트를 사용하려면 비워 두세요.

   --oos-leave-parts-on-error
      문제 발생 시 업로드 중단하기 전에 업로드에 성공한 모든 청크를 S3에 남깁니다.
      
      다른 세션 간에 업로드를 재개해야 할 경우 이 값을 true로 설정해야 합니다.
      
      경고: 완성되지 않은 멀티파트 업로드의 일부를 보관하면 객체 저장 공간 사용량에 포함되며, 정리하지 않으면 추가 비용이 발생할 수 있습니다.
      

   --oos-namespace
      객체 저장 네임스페이스

   --oos-no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone 동작에서 수행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.
      
      또한 사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      

   --oos-provider
      인증 공급자를 선택하세요.

      예시:
         | env_auth                | 런타임(env)에서 자격 증명 자동 가져오기, 인증을 제공하는 첫 번째 자격 증명이 우선순위
         | user_principal_auth     | OCI 사용자 및 API 키를 사용한 인증.
                                   | OCI 테넌시 OCID, 사용자 OCID, 리전, API 키에 대한 경로, 지문을 OCI 구성 파일에 넣어야 합니다.
                                   | https://docs.oracle.com/en-us/iaas/Content/API/Concepts/sdkconfig.htm
         | instance_principal_auth | 인스턴스 원칙을 사용하여 인스턴스가 API 호출을 승인하는 데 사용되는 것입니다.
                                   | 각 인스턴스는 자체 식별자를 가지며 인스턴스 메타데이터에서 읽은 인증서를 사용하여 인증합니다.
                                   | https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm
         | resource_principal_auth | 리소스 원칙을 사용하여 API 호출 수행
         | no_auth                 | 자격 증명이 필요하지 않은 경우, 일반적으로 공개 버킷을 읽는 데 사용됩니다.

   --oos-region
      객체 저장 리전

   --oos-sse-customer-algorithm
      SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다.
      객체 저장은 암호화 알고리즘으로 "AES256"을 지원합니다. 자세한 내용은 아래 문서를 참조하세요.
      용자체 암호화 (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).

      예시:
         | <unset> | 없음
         | AES256  | AES256

   --oos-sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할 선택적 헤더로서, 
      Base64로 인코딩된 256비트 암호화 키를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.
      더 자세한 내용은 아래 문서를 참조하세요. 
      Using Your Own Keys for Server-Side Encryption 
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)

      예시:
         | <unset> | 없음

   --oos-sse-customer-key-file
      SSE-C를 사용하려면, 객체와 연결된 AES-256 암호화 키의 Base64로 인코딩된 문자열이 들어있는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예시:
         | <unset> | 없음

   --oos-sse-customer-key-sha256
      SSE-C를 사용하는 경우 암호화 키의 Base64로 인코딩된 SHA256 해시입니다.
      이 값을 사용하여 암호화 키의 무결성을 확인합니다.
      Using Your Own Keys for 
      Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)을 참조하세요.

      예시:
         | <unset> | 없음

   --oos-sse-kms-key-id
      보관함에서 고유한 마스터 키를 사용하는 경우, 이 헤더는
      데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화하거나 복호화하기 위해 Key Management 서비스를 호출하는 데 
      사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예시:
         | <unset> | 없음

   --oos-storage-tier
      새로운 객체를 저장할 때 사용할 저장 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예시:
         | Standard         | 표준 저장 클래스로, 기본값입니다.
         | InfrequentAccess | 응급 액세스 저장 클래스
         | Archive          | 아카이브 저장 클래스

   --oos-upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크 여러 개를 동시에 업로드합니다.
      
      소수의 대용량 파일을 고속 링크로 업로드하고 이 업로드가 대역폭을 완전히 활용하지 않는 경우 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --oos-upload-cutoff
      청크 업로드를 위한 커트오프 값.
      
      이 값을 초과하는 파일은 chunk_size의 청크로 업로드됩니다. 최소값은 0이고, 최대값은 5GiB입니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터 세트를 CAR 파일로 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터이 시간이 경과할 때마다 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   oos용 옵션

   --oos-chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$OOS_CHUNK_SIZE]
   --oos-compartment value              객체 저장 컴파트먼트 OCID [$OOS_COMPARTMENT]
   --oos-config-file value              OCI 구성 파일 경로 (기본값: "~/.oci/config") [$OOS_CONFIG_FILE]
   --oos-config-profile value           OCI 구성 파일 내 프로필 이름 (기본값: "Default") [$OOS_CONFIG_PROFILE]
   --oos-copy-cutoff value              멀티파트 복사를 위한 커트오프 값입니다. (기본값: "4.656Gi") [$OOS_COPY_CUTOFF]
   --oos-copy-timeout value             복사를 위한 대기 시간입니다. (기본값: "1m0s") [$OOS_COPY_TIMEOUT]
   --oos-disable-checksum value         객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: "false") [$OOS_DISABLE_CHECKSUM]
   --oos-encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$OOS_ENCODING]
   --oos-endpoint value                 객체 저장 API의 엔드포인트입니다. [$OOS_ENDPOINT]
   --oos-leave-parts-on-error value     실패 시 업로드를 중지하지 말고 S3에 성공적으로 업로드된 모든 청크를 수동으로 유지합니다. (기본값: "false") [$OOS_LEAVE_PARTS_ON_ERROR]
   --oos-namespace value                객체 저장 네임스페이스 [$OOS_NAMESPACE]
   --oos-no-check-bucket value          버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: "false") [$OOS_NO_CHECK_BUCKET]
   --oos-provider value                 인증 공급자를 선택하세요 (기본값: "env_auth") [$OOS_PROVIDER]
   --oos-region value                   객체 저장 리전 [$OOS_REGION]
   --oos-sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정하는 선택적 헤더입니다. [$OOS_SSE_CUSTOMER_ALGORITHM]
   --oos-sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할 선택적 헤더로서, [$OOS_SSE_CUSTOMER_KEY]
   --oos-sse-customer-key-file value    SSE-C를 사용하려면, 객체와 연결된 AES-256 암호화 키의 Base64로 인코딩된 문자열이 들어있는 파일입니다. [$OOS_SSE_CUSTOMER_KEY_FILE]
   --oos-sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 Base64로 인코딩된 SHA256 해시입니다. [$OOS_SSE_CUSTOMER_KEY_SHA256]
   --oos-sse-kms-key-id value           보관함에서 고유한 마스터 키를 사용하는 경우, 이 헤더는 [$OOS_SSE_KMS_KEY_ID]
   --oos-storage-tier value             새로운 객체를 저장할 때 사용할 저장 클래스입니다. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$OOS_STORAGE_TIER]
   --oos-upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (기본값: "10") [$OOS_UPLOAD_CONCURRENCY]
   --oos-upload-cutoff value            청크 업로드를 위한 커트오프 값입니다. (기본값: "200Mi") [$OOS_UPLOAD_CUTOFF]

```
{% endcode %}