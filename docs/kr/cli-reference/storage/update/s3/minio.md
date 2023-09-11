# Minio Object Storage

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 minio - 미니오 객체 스토리지

사용법:
   singularity storage update s3 minio [command options] <name|id>

DESCRIPTION:
   --env-auth
      AWS 자격증명을 런타임에서 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key가 비어있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격증명을 입력합니다.
         | true  | 환경에서 AWS 자격증명을 가져옵니다 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격증명을 사용하려면 비워 둡니다.

   --secret-access-key
      AWS 시크릿 액세스 키 (비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격증명을 사용하려면 비워 둡니다.

   --region
      연결할 지역입니다.
      
      S3 클론을 사용하고 지역을 가지고 있지 않은 경우 비워 둡니다.

      예시:
         | <unset>            | 확실하지 않은 경우 사용하세요.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않는 경우에만 사용합니다.
         |                    | 예: Jewel/v10 CEPH의 경우.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필수입니다.

   --location-constraint
      지역 제약 사항 - 지역과 일치해야 합니다.
      
      확실하지 않은 경우 비워 둡니다. 버킷을 만들 때만 사용됩니다.

   --acl
      버킷 생성 및 개체 저장 또는 복사 시 사용되는 canned ACL입니다.
      
      이 ACL은 개체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      S3가 소스로부터 ACL을 복사하는 대신 새로운 ACL을 작성하기 때문에 이 ACL은
      서버 측 개체 복사시 적용됩니다.
      
      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고
      기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성시 사용되는 canned ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      이 기능이 설정되지 않은 경우 대신 "acl"이 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열이면 X-Amz-Acl:
      헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한이 있음.
         |                    | 다른 사람은 액세스 권한을 갖지 않음(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 있음.
         |                    | AllUsers 그룹에게 읽기 액세스 권한이 있음.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 있음.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스 권한이 있음.
         |                    | 버킷에 대한 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 있음.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스 권한이 있음.

   --server-side-encryption
      S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예시:
         | <unset> | 없음
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C 사용 시 S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예시:
         | <unset> | 없음
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다.

      예시:
         | <unset>                 | 없음
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C를 사용하려면 데이터의 암호화/복호화에 사용하는 비밀 암호화 키를 제공할 수 있습니다.
      
      --sse-customer-key-base64를 대신 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터의 암호화/복호화에 사용하는 비밀 암호화 키를 base64 형식으로 인코딩된 상태로 제공해야 합니다.
      
      --sse-customer-key를 대신 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 secret encryption key의 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 둔 경우 이는 sse_customer_key에서 자동으로 계산됩니다.
      

      예시:
         | <unset> | 없음

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나 크기가 알려지지 않은 파일(예: "rclone rcat"에서 가져온 파일 또는 "rclone mount" 또는 google photos 또는 google docs로 업로드된 파일)을 업로드할 때 이 청크 크기를 사용하여 multipart 업로드를 수행합니다.
      
      청크 크기당 "--s3-upload-concurrency"개의 청크가 전송마다 메모리에 버퍼링됩니다.
      
      대역폭이 높은 링크로 큰 파일을 전송하고 메모리가 충분하다면 이 값을 늘리면 전송 속도가 높아집니다.
      
      rclone은 알려진 크기의 큰 파일을 업로드할 때 10,000개의 청크 제한을 초과하지 않도록 청크 크기를 자동으로 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본적으로 chunk_size는 5 MiB이고 최대 10,000개의 청크가 있을 수 있습니다. 따라서 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 낮아집니다. rclone은 청크가 AWS SDK에 의해 버퍼링될 때 청크를 전송한 것으로 처리하지만 실제로는 아직 업로드 중일 수 있습니다. 청크 크기가 클수록 AWS SDK 버퍼와 진행률 보고가 실제와 더 다를 수 있습니다.
      

   --max-upload-parts
      multipart 업로드에 사용할 최대 부분 수입니다.
      
      이 옵션은 multipart 업로드시 사용할 전체 업로드 부분 수를 정의합니다.
      
      10,000개의 청크를 지원하지 않는 서비스인 경우 유용할 수 있습니다.
      
      rclone은 알려진 크기의 큰 파일을 업로드할 때 자동으로 청크 크기를 증가시켜 이러한 청크 수 제한을 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값입니다.
      
      서버 사이드 복사가 필요한 이 임계값보다 큰 파일은 이 크기로 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이것은 큰 파일의 업로드를 시작하기 전에 오랜 지연을 초래하지만 데이터의 무결성 검사에 매우 유용합니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true로 설정된 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면 현재 사용자의 홈 디렉터리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일입니다.
      
      env_auth = true로 설정된 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용될 프로필을 제어합니다.
      
      값이 비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"인 경우 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      multipart 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크가 동시에 업로드되는 수입니다.
      
      대역폭을 충분히 활용하지 못하고 높은 속도로 대량 파일을 업로드하는 경우 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.
      
      true인 경우(rclone의 기본값) rclone은 경로 스타일 액세스를 사용하고 false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 설명서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는이 값을 false로 설정해야 합니다 - rclone은 공급자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정되어 있으면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 CEPH의 경우.

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 응답 목록을 1000 개로 자르지만 요청보다 더 많이 요청하더라도 이에 기반합니다.
      AWS S3에서는 이것이 전역 최대이므로 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)을 참조하십시오.
      Ceph의 경우 "rgw list buckets max chunk" 옵션을 사용하여 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동(0).
      
      S3가 처음 출시될 때 버킷의 개체를 열거하는 데는 ListObjects 호출 만을 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자가 정의한 것에 따라 호출할 목록 개체 방법을 추측합니다. 만약 올바른 추측을 하지 못한다면 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며 사용 가능한 경우 이는 더 신뢰할 수 있습니다. 이 값이 unset으로 설정된 경우 rclone은 공급자 설정에 따라 적용할 항목을 선택합니다.

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 갖지 않은 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이는 무시되었기 때문에 매우 소음 없이 전달했을 것입니다.
      

   --no-head
      체크섬 유효성을 확인하기 위해 업로드된 개체의 HEAD를 사용하지 않습니다.
      
      rclone이 PUT으로 개체를 업로드한 후 200 OK 메시지를 수신하면 제대로 업로드된 것으로 간주합니다.

      특히 다음을 가정합니다:
      
      - 업로드시 메타데이터(수정 시간, 저장 클래스 및 컨텐츠 유형 포함)가 업로드한 것과 같았음
      - 크기가 업로드한 것과 같았음
      
      단일 부분 PUT 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드 날짜
      
      Multipart 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      알려지지 않은 길이의 소스 개체가 업로드된 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기와 같은 감지되지 않은 업로드 실패 가능성이 증가하므로 정상적인 작업에는 권장되지 않습니다. 실제로 감지되지 않은 업로드 실패 가능성은 매우 낮습니다.
      

   --no-head-object
      개체를 가져올 때 HEAD 전에 GET을 수행하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.
      
      추가 버퍼(예: 멀티파트가 필요한 업로드)는 메모리 풀을 사용하여 할당됩니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 미니오) 백엔드와 HTTP/2에 관련한 해결되지 않은 문제가 있습니다. HTTP/2는 기본적으로 s3 백엔드에서 활성화되지만 이곳에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드용 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 다운로드한 데이터에 대해 더 저렴한 대출을 제공합니다.

   --use-multipart-etag
      검증을 위해 multipart 업로드에 ETag를 사용할지 여부
      
      이것은 true, false 또는 공급자의 기본값을 사용하려면 설정하지 않습니다.
      

   --use-presigned-request
      단일 부분 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 플래그를 false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone의 버전 1.59 이전 버전은 단일 부분 객체를 업로드하기 위해 사전 서명된 요청을 사용하고이 플래그를 true로 설정하면 해당 기능을 다시 활성화합니다. 이는 예외적인 경우 또는 테스트를 위해서만 필요하지만 그 외의 경우는 권장되지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함시킵니다.

   --version-at
      지정된 시간대의 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그보다 오래된 기간(예: "100d" 또는 "1h")이어야 합니다.
      
      이를 사용할 때는 파일 쓰기 작업을 수행할 수 없으므로 파일 업로드 또는 삭제를 수행할 수 없습니다.
      
      유효한 형식에 대해서는 [time 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩된 개체를 해제합니다.
      
      S3에 "Content-Encoding: gzip"가 설정된 개체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 개체를 받을 때 "Content-Encoding: gzip"로 풀어 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 풀린 상태입니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있다면 이를 설정하십시오.
      
      일반적으로 공급자는 파일이 다운로드될 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"으로 업로드되지 않은 경우 다운로드되지 않을 것입니다.
      
      그러나 일부 공급자는 파일이 "Content-Encoding: gzip"으로 업로드되지 않았더라도(예: Cloudflare) 개체를 gzip으로 압축할 수 있습니다.
      
      이러한 경우 다음과 같은 오류 메시지를 수신하는 경우가 발생할 수 있습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩이 설정된 개체를 다운로드하면 rclone은 개체를 실시간으로 해제합니다.
      
      unset으로 설정하면(기본값) rclone은 공급자 설정에 따라 적용할 항목을 선택합니다. memcpy : (buffervoid로)인할rclone이 개체를 다운로드하면 기본적으로 푸는 것이고 해리rclone이 개체를 다운로드하면 memcpy-인하기 때문에 디코딩하면서 객체를 만듭니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value           AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                     버킷 생성 및 개체 저장 또는 복사 시 사용되는 canned ACL입니다. [$ACL]
   --endpoint value                S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                      실행 시점에서 AWS 자격증명을 가져옵니다(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     위치 제약 사항 - 지역과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value                  연결할 지역입니다. [$REGION]
   --secret-access-key value       AWS 시크릿 액세스 키 (비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다. [$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               버킷 생성시 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크 복사로 전환하는 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드용 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부 (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0 (자동). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         multipart 업로드에 사용할 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있다면 이를 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        체크섬 유효성을 확인하기 위해 업로드된 개체의 HEAD를 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져올 때 HEAD 전에 GET을 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로파일입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C 사용 시 S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터의 암호화/복호화에 사용하는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우 데이터의 암호화/복호화에 사용하는 비밀 암호화 키를 base64 형식으로 인코드 된 상태로 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하여 secret encryption key의 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       multipart 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 multipart 업로드에 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간대로 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함시킵니다. (기본값: false) [$VERSIONS]

```
{% endcode %}