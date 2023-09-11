# 인스턴스 제한자를 사용하여 API 호출을 수행하는 인스턴스를 인증하십시오.
각 인스턴스는 고유한 식별자를 가지며 인스턴스 메타데이터에서 읽은 인증서를 사용하여 인증합니다.
https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

{% code fullWidth="true" %}
```
NANE:
   singularity storage create oos instance_principal_auth - use instance principals to authorize an instance to make API calls. 
                                                            each instance has its own identity, and authenticates using the certificates that are read from instance metadata. 
                                                            https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/callingservicesfrominstances.htm

사용법:
   singularity storage create oos instance_principal_auth [command options] [arguments...]

설명:
   --네임스페이스
      객체 저장소 네임스페이스

   --구편
      객체 저장소 구편 OCID

   --지역
      객체 저장소 지역

   - 에피스텐트
      객체 저장소 API의 엔드포인트.
      
      지역의 기본 엔드포인트를 사용하려면 비워 둡니다.

   --스토리지 티어
      새로운 객체를 저장할 때 사용할 저장 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm

      예제:
         | Standard         | 표준 저장 티어, 기본 티어입니다.
         | InfrequentAccess | 적은 액세스 저장 티어
         | Archive          | 아카이브 저장 티어

   --업로드 끊김
      청크 업로드로 전환하는 끊김 값.
      
      이 값보다 큰 파일은 청크 크기로 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --청크 크기
      업로드에 사용할 청크 크기.
      
      업로드 끊김보다 크거나 알 수없는 크기의 파일(예: "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)를 업로드할 때
      이 청크 크기를 사용하여 멀티 파트 업로드로 업로드됩니다.
      
      이전 "upload_concurrency" 크기의 청크가 전송당 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 큰 파일을 전송하는 경우 충분한 메모리가 있는 경우 이 값을 높이면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드할 때 10,000 청크 제한 아래로 유지하기 위해 청크 크기를 자동으로 증가시킵니다.
      
      알 수없는 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본적으로 청크 크기는 5 MiB이며 최대 10,000 청크입니다.
      스트리밍 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트리밍 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그로 표시되는 진행 상태 통계의 정확도가 감소합니다.
      

   --업로드 동시성
      멀티 파트 업로드에 대한 동시성.
      
      한 파일의 동일한 청크를 동시에 업로드하는 수입니다.
      
      고속 링크를 통해 적은 수의 대형 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않을 경우 이 값을 늘리면 전송 속도가 향상될 수 있습니다.

   --복사 끊김
      멀티파트 복사로 전환하는 끊김 값.
      
      서버 측에서 복사해야 하는 이 크기보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --복사 제한 시간
      복사 시간 제한.
      
      복사는 비동기 작업이므로 성공한 복사를 기다리기 위해 제한 시간을 지정합니다.
      

   --체크섬 사용 안함
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체에 메타데이터로 추가합니다. 이는
      데이터 무결성 확인에 크게 도움이 되지만 대형 파일을 업로드하기 시작하는 데 오랜 지연을 초래할 수 있습니다.

   --인코딩
      백엔드에 대한 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --오류 시 파트 유지
      실패로 인해 업로드 중단 호출을 피할 경우 모든 성공적으로 업로드된 부분을 수동으로 복구하기 위해 S3에 유지합니다.
      
      서로 다른 세션에서 업로드를 재개해야하는 경우 true로 설정해야 합니다.
      
      경고: 불완전한 멀티파트 업로드의 부분을 저장하면 개체 저장소에서 공간 사용 용량에 영향을 주고
      정리하지 않으면 추가 비용이 발생할 수 있습니다.
      

   --버킷 확인 안함
      버킷 존재 여부를 확인하거나 만들지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 갖지 않은 경우에도 필요할 수 있습니다.
      

   --서비스 지원 고객 키 파일
      SSE-C를 사용하려면 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예제:
         | <설정 안됨> | 없음

   --서비스 지원 고객 키
      SSE-C를 사용하려면 데이터의 암호화 또는 복호화에 사용할 base64로 인코딩된 256-bit 암호화 키를 지정하는 선택적 헤더입니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다. 자세한 정보는
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)에서 "서버 측 암호화에 자체 키 사용" 를 참조하십시오.

      예제:
         | <설정 안됨> | 없음

   --서비스 지원 고객 키 SHA256
      SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정하는 선택적 헤더입니다.
      이 값은 암호화키의 무결성을 확인하기 위해 사용됩니다. 자체 키 사용에 대한 자세한 내용은
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예제:
         | <설정 안됨> | 없음

   --서비스 지원 KMS 키 ID
      소유한 마스터 키를 사용 중인 경우 이 헤더는 데이터 암호화 키를 생성하거나 데이터 암호화 키를 암호화하거나 복호화하기 위해
      키 관리 서비스를 호출하는 데 사용되는 마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예제:
         | <설정 안됨> | 없음

   --서비스 지원 고객 알고리즘
      SSE-C를 사용하는 경우 선택적 헤더로 "AES256"을 암호화 알고리즘으로 지정합니다.
      객체 저장소는 암호화 알고리즘으로 "AES256"을 지원합니다. 자세한 정보는
      (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm) 를 참조하십시오.

      예제:
         | <설정 안됨> | 없음
         | AES256  | AES256


옵션:
   --구편 값  객체 저장소 구편 OCID [$COMPARTMENT]
   --엔드포인트 값     객체 저장소 API의 엔드포인트. [$ENDPOINT]
   --help, -h           도움말 표시
   --네임스페이스 값    객체 저장소 네임스페이스 [$NAMESPACE]
   --지역 값       객체 저장소 지역 [$REGION]

   고급

   --청크 크기 값               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --복사 끊김 값              멀티파트 복사로 전환하는 끊김 값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --복사 제한 시간 값             복사 시간 제한. (기본값: "1m0s") [$COPY_TIMEOUT]
   --체크섬 사용 안함               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --인코딩 값                 백엔드를 위한 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --오류 시 파트 유지           실패로 인해 업로드 중단 호출을 피하고 모든 성공적으로 업로드된 부분을 S3에 수동으로 복구합니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --버킷 확인 안함                버킷 존재 여부를 확인하거나 만들지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --서비스 지원 고객 알고리즘 값   SSE-C를 사용하는 경우 선택적 헤더로 "AES256"을 지정합니다. [$SSE_CUSTOMER_ALGORITHM]
   --서비스 지원 고객 키 값         SSE-C를 사용하려면 데이터의 암호화 또는 복호화에 사용할 base64로 인코딩된 256-bit 암호화 키를 지정합니다. [$SSE_CUSTOMER_KEY]
   --서비스 지원 고객 키 파일 값    SSE-C를 사용하려면 객체와 관련된 AES-256 암호화 키의 base64로 인코딩된 문자열을 포함하는 파일입니다. [$SSE_CUSTOMER_KEY_FILE]
   --서비스 지원 고객 키 SHA256 값  SSE-C를 사용하는 경우, 암호화 키의 base64로 인코딩된 SHA256 해시를 지정합니다. [$SSE_CUSTOMER_KEY_SHA256]
   --서비스 지원 KMS 키 ID 값           소유한 키로 마스터 키를 사용 중인 경우 [$SSE_KMS_KEY_ID]
   --스토리지 티어 값             새로운 개체를 저장할 때 사용할 저장 클래스. https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm (기본값: "Standard") [$STORAGE_TIER]
   --업로드 동시성 값       멀티 파트 업로드에 대한 동시성. (기본값: 10) [$UPLOAD_CONCURRENCY]
   --업로드 끊김 값            청크 업로드로 전환하는 끊김 값. (기본값: "200Mi") [$UPLOAD_CUTOFF]

   General

   --이름 값  Storage의 이름 (기본값: Auto generated)
   --경로 값  Storage의 경로

```
{% endcode %}