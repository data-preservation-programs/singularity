# 자격증명 필요 없음, 일반적으로 공개 버킷을 읽는 데 사용됩니다.

{% code fullWidth="true" %}
```
명령어:
   singularity storage create oos no_auth - 자격증명 필요 없음, 일반적으로 공개 버킷을 읽는 데 사용됩니다.

사용법:
   singularity storage create oos no_auth [옵션] [인수...]

설명:
   --namespace
      객체 저장소 네임스페이스

   --region
      객체 저장소 지역

   --endpoint
      객체 저장소 API의 엔드포인트.
      
      지역의 기본 엔드포인트를 사용하려면 비워둡니다.

   --storage-tier
      새로운 객체를 저장할 때 사용할 저장 클래스입니다. [스토리지 계층 이해](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm)

      예:
         | Standard         | 표준 저장 계층, 기본 저장 계층입니다
         | InfrequentAccess | 자주 사용되지 않는 저장 계층
         | Archive          | 아카이브 저장 계층

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값입니다.
      
      이보다 큰 파일은 청크 크기로 분할하여 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일
      ("rclone rcat"에서 나오는 파일 또는 "rclone mount"나 Google
      포토, Google 문서로 업로드된 파일)은 이 청크 크기로
      멀티파트 업로드로 업로드됩니다.
      
      이 전송 당 "upload_concurrency"개의 청크가 이 크기로
      메모리에 버퍼링됩니다.
      
      고속 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우,
      청크 크기를 증가시켜 전송 속도를 높일 수 있습니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드할 때
      청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 벗어나지
      않도록 합니다.
      
      알려지지 않은 크기의 파일은 구성된
      chunk_size로 업로드됩니다. 기본 청크 크기는
      5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로,
      기본적으로 스트림 업로드 가능한 파일의 최대 크기는
      48 GiB입니다. 더 큰 파일을 스트림 업로드하려면
      chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는
      진행 상태 통계의 정확도가 낮아집니다.
      

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 동일한 파일 청크의 수입니다.
      
      고속 링크로 대용량 파일을 작은 수로 업로드하고
      이러한 업로드가 대역폭을 모두 사용하지 않으면,
      이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --copy-cutoff
      분할 복사로 전환하기 위한 임계값입니다.
      
      서버 측 복사가 필요한 이 임계값보다 큰 파일은
      이 크기로 분할하여 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --copy-timeout
      복사에 대한 시간 제한입니다.
      
      복사는 비동기 작업이므로 성공에 대한 대기 시간을 지정합니다.
      

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      개체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에는 훌륭하지만,
      대용량 파일의 업로드 시작에 대한 긴 지연을 초래할 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --leave-parts-on-error
      실패 시 업로드를 중단하지 않고, 성공적으로 업로드된 파트를 S3에 남겨 수동으로 복구하도록합니다.
      
      다양한 세션 간 업로드를 재개하려면 true로 설정해야합니다.
      
      경고: 완성되지 않은 멀티파트 업로드의 일부를 저장하는 경우,
      객체 저장소에서 공간 사용량으로 취급되며,
      정리되지 않으면 추가 비용이 발생합니다.
      

   --no-check-bucket
      버킷의 존재를 확인하지 않거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone의 트랜잭션 수를
      최소화하려는 경우 유용할 수 있습니다.
      
      또한 사용자에게 버킷 생성 권한이없는 경우에도 필요할 수 있습니다.
      

   --sse-customer-key-file
      SSE-C를 사용하려면 객체와 연결된
      AES-256 암호화 키의 base64로 인코딩 된 문자열을 포함한
      파일을 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.'

      예:
         | <unset> | 없음

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데
      사용할 base64로 인코딩된 256비트 암호화 키를 지정하는
      선택적 헤더입니다. sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.
      자세한 내용은 [사용자 고유 키를 사용한 서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음

   --sse-customer-key-sha256
      SSE-C를 사용하는 경우 암호화 키의
      base64로 인코딩 된 SHA256 해시를 지정하는
      선택적 헤더입니다. 자세한 내용은
      [사용자 고유 키를 사용한 서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음

   --sse-kms-key-id
      보관함의 고유 마스터 키를 사용하는 경우
      이 헤더는 데이터 암호화 키를 생성하거나
      데이터 암호화 키를 암호화 또는 복호화하기
      위해 키 관리 서비스를 호출하기 위해 사용되는
      마스터 암호화 키의 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)를 지정합니다.
      sse_customer_key_file|sse_customer_key|sse_kms_key_id 중 하나만 필요합니다.

      예:
         | <unset> | 없음

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 선택적 헤더로서,
      암호화 알고리즘으로 "AES256"을 지정합니다.
      객체 저장소는 암호화 알고리즘으로 "AES256"을 지원합니다.
      자세한 내용은 [사용자 고유 키를 사용한 서버 측 암호화](https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm)를 참조하십시오.

      예:
         | <unset> | 없음
         | AES256  | AES256


옵션:
   --endpoint value   객체 저장소 API의 엔드포인트 [$ENDPOINT]
   --help, -h         도움말 표시
   --namespace value  객체 저장소 네임스페이스 [$NAMESPACE]
   --region value     객체 저장소 지역 [$REGION]

   고급

   --chunk-size value               업로드에 사용할 청크 크기 (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              분할 복사로 전환하기 위한 임계값 (기본값: "4.656Gi") [$COPY_CUTOFF]
   --copy-timeout value             복사에 대한 시간 제한 (기본값: "1m0s") [$COPY_TIMEOUT]
   --disable-checksum               개체 메타데이터에 MD5 체크섬을 저장하지 않음 (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩 (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --leave-parts-on-error           실패 시 업로드를 중단하지 않고, 성공적으로 업로드된 파트를 S3에 남겨 수동으로 복구하기 (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-check-bucket                버킷의 존재를 확인하지 않거나 생성하지 않도록 설정함 (기본값: false) [$NO_CHECK_BUCKET]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 암호화 알고리즘으로 "AES256"을 지정합니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화 또는 복호화하는 데 사용할 base64로 인코딩된 256비트 암호화 키를 지정합니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-file value    SSE-C를 사용하려면 객체와 연결된 AES-256 암호화 키의 base64로 인코딩 된 문자열을 포함한 파일을 지정합니다. [$SSE_CUSTOMER_KEY_FILE]
   --sse-customer-key-sha256 value  SSE-C를 사용하는 경우 암호화 키의 base64로 인코딩 된 SHA256 해시를 지정합니다. [$SSE_CUSTOMER_KEY_SHA256]
   --sse-kms-key-id value           보관함의 고유 마스터 키를 사용하는 경우 OCID(https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm)로 암호화 키를 지정합니다. [$SSE_KMS_KEY_ID]
   --storage-tier value             새로운 객체를 저장할 때 사용할 저장 클래스. [스토리지 계층 이해](https://docs.oracle.com/en-us/iaas/Content/Object/Concepts/understandingstoragetiers.htm) (기본값: "Standard") [$STORAGE_TIER]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성 (기본값: 10) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계값 (기본값: "200Mi") [$UPLOAD_CUTOFF]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}