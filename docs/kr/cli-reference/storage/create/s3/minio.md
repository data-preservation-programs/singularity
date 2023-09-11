# Minio Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 minio - Minio 오브젝트 스토리지

사용법:
   singularity storage create s3 minio [command options] [arguments...]

설명:
   --env-auth
      실행 시간(AWS 환경 변수 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옴.
      
      access_key_id와 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 실행 시간(AWS 환경 변수 또는 IAM)에서 AWS 자격 증명 가져오기.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 실행 시간 자격 증명을 사용하려면 빈 값으로 남깁니다.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호).
      
      익명 액세스 또는 실행 시간 자격 증명을 사용하려면 빈 값으로 남깁니다.

   --region
      연결할 리전.
      
      S3 클론을 사용하고 리전이 없는 경우 비워 둡니다.

      예제:
         | <unset>            | 확실하지 않은 경우에 사용합니다.
         |                    | v4 서명과 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않은 경우에만 사용합니다.
         |                    | 예: 이전 버전의 Jewel/v10 CEPH.

   --endpoint
      S3 API의 엔드포인트.
      
      S3 클론을 사용하는 경우 필요합니다.

   --location-constraint
      리전과 일치하는 위치 제한.
      
      확실하지 않은 경우 비워둡니다. 버킷을 만들 때만 사용됩니다.

   --acl
      버킷 및 객체 생성 또는 복사 시 사용되는 Canned ACL.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되어 있지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl 를 참조하십시오.
      
      S3에서 서버 쪽에서 객체를 복사할 때 이 ACL이 적용됩니다.
      S3가 소스에서 ACL을 복사하지 않고 새로 쓰므로 이 ACL이 적용됩니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본(프라이빗)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl 를 참조하십시오.
      
      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl"을 대신 사용합니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본(프라이빗)이 사용됩니다.
      

      예제:
         | private            | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 액세스를 얻습니다.
         | public-read-write  | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻습니다.
         |                    | 버킷에 대해 이 권한을 부여하면 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스를 얻습니다.

   --server-side-encryption
      S3에 객체를 저장할 때 사용되는 서버 쪽 암호화 알고리즘.

      예제:
         | <unset> | None
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 S3에 객체를 저장할 때 사용되는 서버 쪽 암호화 알고리즘.

      예제:
         | <unset> | None
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다.

      예제:
         | <unset>                 | None
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다.
      
      대안으로 --sse-customer-key-base64를 사용할 수도 있습니다.

      예제:
         | <unset> | None

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터를 암호화/복호화하기 위해 base64로 인코딩된 비밀 암호화 키를 제공해야 합니다.
      
      대안으로 --sse-customer-key를 사용할 수도 있습니다.

      예제:
         | <unset> | None

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 비밀 암호화 키 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 두면 sse_customer_key에서 자동으로 계산됩니다.
      

      예제:
         | <unset> | None

   --upload-cutoff
      청크 업로드로 전환하는 임계값.
      
      이보다 큰 파일은 chunk_size의 크기로 청크 업로드됩니다.
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 크거나 알 수 없는 크기의 파일(예: "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)을 업로드할 때 이 청크 크기를 사용하여 여러 부분으로 업로드됩니다.
      
      참고로 "--s3-upload-concurrency"은 이 크기만큼의 청크가 전송당 메모리에 버퍼링됩니다.
      
      높은 속도 링크에서 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 높이면 전송 속도가 향상됩니다.
      
      큰 파일(크기가 알려진)을 업로드할 때 rclone은 최대 10,000개의 청크 한도를 지키기 위해 청크 크기를 자동으로 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 낮아집니다. rclone은 청크를 AWS SDK(소프트웨어 개발 키트)에 버퍼링 할 때 청크가 보내진 것으로 처리하지만, 실제로 업로드 중일 수도 있습니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 통보의 결과로 이어집니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수.
      
      이 옵션은 멀티파트 업로드 시 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      S3 사양의 10,000개의 청크를 지원하지 않는 서비스라면 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드할 때 이 옵션에 따라 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값.
      
      이보다 큰 서버 사이 쪽 복사가 필요한 파일은 이 크기로 청크로 복사됩니다.
      
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만 대용량 파일의 업로드 시작에 대한 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true 경우에만 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본적으로 현재 사용자의 홈 디렉토리를 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일입니다.
      
      env_auth = true 경우에만 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용되는 프로파일을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default" 환경 변수가 설정되지 않은 경우 기본값으로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크를 동시에 업로드하는 수입니다.
      
      고속 링크에서 소수의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true 인 경우 path style 액세스를 사용하고 false 인 경우 virtual hosted style을 사용합니다.
      
      true(기본값)이면 rclone은 path style 액세스를 사용하고, false이면 rclone은 virtual path style을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이 작업을 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용하며, 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용합니다. 예: 이전 버전의 Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(S3 요청의 각 ListObject의 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청된 개체보다 많은 응답 목록을 획득하지 않습니다. AWS S3에서는 이것이 전역 최대값이며 변경할 수 없으며 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 통해 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 0은 자동.
      
      S3가 처음 시작되었을 때 버킷의 개체를 열거하기 위해 ListObjects 호출만을 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 소개되었습니다. 이는 훨씬 더 높은 성능을 제공하므로 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 자체 설정에 따라 어떤 목록 객체 방법을 호출할지 추측합니다. 틀렸다면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며, 가능한 경우 파일 이름에 제어 문자를 사용할 때 이 방법이 더 안정적입니다. unset으로 설정하면 rclone은 공급자 설정에 따라 무엇을 적용할지 선택합니다.

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 실행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      또한 사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. 버전 v1.52.0 이전에는 오류로 인해 무시되었을 것입니다.
      

   --no-head
      업로드한 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다.
      
      rclone은 PUT를 사용하여 객체를 업로드한 후 200 OK 메시지를 수신하면 올바르게 업로드된 것으로 간주합니다.
      
      특히 다음 항목을 가정합니다.
      
      - metadata, 수정 시간, 저장 클래스 및 콘텐츠 유형은 업로드한 것과 같음
      - 크기가 업로드한 것과 같음
      
      싱글 파트 PUT의 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드한 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      알 수 없는 길이의 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 확률이 높아지므로 일반 작업에는 권장되지 않습니다. 실제로 업로드 실패가 감지되는 확률은 매우 낮습니다.

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 비워질 때까지 얼마나 오래 소요되는지 지정합니다.
      
      추가 버퍼를 필요로 하는 업로드(예: 멀티파트)는 메모리 풀을 사용하여 할당을 수행합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관련된 미해결된 문제가 있습니다. HTTP/2은 기본적으로 s3 백엔드에 대해 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그는 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 이른바 egress를 제공합니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값은 true, false 또는 기본값을 사용합니다.

   --use-presigned-request
      싱글 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용해야 하는지 여부
      
      false로 설정하면 rclone은 단일 파트 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone의 버전 < 1.59은 테스트나 특별한 상황에서만 원활한 동작을 위해 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하며, 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜, "2006-01-02", datetime "2006-01-02 15:04:05" 또는 그 이래로 오랜 시간 사용하는 지속시간을 지정할 수 있습니다.
      
      이를 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      설정된 경우 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"이 설정된 상태에서 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 "Content-Encoding: gzip"으로 수신되는 즉시 이러한 파일을 압축 해제합니다. 이로 인해 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 해제됩니다.

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하십시오.
      
      일반적으로 공급자는 개체를 다운로드할 때 개체를 변경하지 않습니다. 만약 개체가 `Content-Encoding: gzip`로 업로드되지 않은 경우 다운로드 시 설정되지 않습니다.
      
      그러나 어떤 공급자는 `Content-Encoding: gzip`로 업로드되지 않은 경우에도 개체를 gzip으로 압축할 수 있습니다(예: Cloudflare).
      
      이런 경우 데이터가 `Content-Encoding: gzip`으로 설정되고 청크 전송 인코딩이 있는 객체를 rclone이 다운로드하면 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본값)으로 설정하면 rclone은 공급자 설정에 따라 무엇을 적용할지 선택하지만 사용자가 rclone의 선택을 여기서 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value           AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                     버킷 및 객체 생성 또는 복사 시 사용되는 Canned ACL. [$ACL]
   --endpoint value                S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                      실행 시간(AWS 환경 변수 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옴. (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     리전과 일치하는 위치 제한. [$LOCATION_CONSTRAINT]
   --region value                  연결할 리전. [$REGION]
   --secret-access-key value       AWS 비밀 액세스 키 (비밀번호). [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에 객체를 저장할 때 사용되는 서버 쪽 암호화 알고리즘. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다. [$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크 복사로 전환하는 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정된 경우 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 path style 액세스를 사용하고 false 인 경우 virtual hosted style을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(S3 요청의 각 ListObject의 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0은 자동. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 비워질 때까지 얼마나 오래 소요되는지 지정합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로파일입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 S3에 객체를 저장할 때 사용되는 서버 쪽 암호화 알고리즘. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우 데이터를 암호화/복호화하기 위해 base64로 인코딩된 비밀 암호화 키를 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우 비밀 암호화 키 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          싱글 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}