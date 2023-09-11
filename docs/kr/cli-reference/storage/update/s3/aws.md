# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 aws - Amazon Web Services (AWS) S3

사용법:
   singularity storage update s3 aws [command options] <이름|ID>

설명:
   --env-auth
      AWS 자격증명을 런타임 환경(환경 변수 또는 env vars 또는 EC2/ECS 메타데이터)에서 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어 있을 때만 적용됩니다.
      
      예시:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(암호)입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --region
      연결할 리전입니다.

      예시:
         | us-east-1      | 기본 엔드포인트 - 확실하지 않은 경우 좋은 선택입니다.
         |                | US 리전, 북부 버지니아 또는 태평양 북서 지역입니다.
         |                | 위치 제약 조건을 비워 두세요.
         | us-east-2      | US 동부(오하이오) 리전입니다.
         |                | 위치 제약 조건은 us-east-2입니다.
         | us-west-1      | US 서부(북부 캘리포니아) 리전입니다.
         |                | 위치 제약 조건은 us-west-1입니다.
         | us-west-2      | US 서부(오레곤) 리전입니다.
         |                | 위치 제약 조건은 us-west-2입니다.
         | ca-central-1   | 캐나다(중부) 리전입니다.
         |                | 위치 제약 조건은 ca-central-1입니다.
         | eu-west-1      | EU(아일랜드) 리전입니다.
         |                | 위치 제약 조건은 EU 또는 eu-west-1입니다.
         | eu-west-2      | EU(런던) 리전입니다.
         |                | 위치 제약 조건은 eu-west-2입니다.
         | eu-west-3      | EU(파리) 리전입니다.
         |                | 위치 제약 조건은 eu-west-3입니다.
         | eu-north-1     | EU(스톡홀름) 리전입니다.
         |                | 위치 제약 조건은 eu-north-1입니다.
         | eu-south-1     | EU(밀라노) 리전입니다.
         |                | 위치 제약 조건은 eu-south-1입니다.
         | eu-central-1   | EU(프랑크푸르트) 리전입니다.
         |                | 위치 제약 조건은 eu-central-1입니다.
         | ap-southeast-1 | 아시아 태평양(싱가포르) 리전입니다.
         |                | 위치 제약 조건은 ap-southeast-1입니다.
         | ap-southeast-2 | 아시아 태평양(시드니) 리전입니다.
         |                | 위치 제약 조건은 ap-southeast-2입니다.
         | ap-northeast-1 | 아시아 태평양(도쿄) 리전입니다.
         |                | 위치 제약 조건은 ap-northeast-1입니다.
         | ap-northeast-2 | 아시아 태평양(서울) 리전입니다.
         |                | 위치 제약 조건은 ap-northeast-2입니다.
         | ap-northeast-3 | 아시아 태평양(오사카-로컬) 리전입니다.
         |                | 위치 제약 조건은 ap-northeast-3입니다.
         | ap-south-1     | 아시아 태평양(뭄바이) 리전입니다.
         |                | 위치 제약 조건은 ap-south-1입니다.
         | ap-east-1      | 아시아 태평양(홍콩) 리전입니다.
         |                | 위치 제약 조건은 ap-east-1입니다.
         | sa-east-1      | 남아메리카(상파울루) 리전입니다.
         |                | 위치 제약 조건은 sa-east-1입니다.
         | me-south-1     | 중동(바레인) 리전입니다.
         |                | 위치 제약 조건은 me-south-1입니다.
         | af-south-1     | 아프리카(케이프타운) 리전입니다.
         |                | 위치 제약 조건은 af-south-1입니다.
         | cn-north-1     | 중국(베이징) 리전입니다.
         |                | 위치 제약 조건은 cn-north-1입니다.
         | cn-northwest-1 | 중국(닝샤) 리전입니다.
         |                | 위치 제약 조건은 cn-northwest-1입니다.
         | us-gov-east-1  | AWS GovCloud(미국-동부) 리전입니다.
         |                | 위치 제약 조건은 us-gov-east-1입니다.
         | us-gov-west-1  | AWS GovCloud(미국) 리전입니다.
         |                | 위치 제약 조건은 us-gov-west-1입니다.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      AWS를 사용하는 경우 리전의 기본 엔드포인트를 사용하려면 비워 둡니다.

   --location-constraint
      리전과 일치해야 하는 위치 제한입니다.
      
      버킷을 생성할 때만 사용됩니다.

      예시:
         | <unset>        | US 리전, 북부 버지니아 또는 태평양 북서
         | us-east-2      | US 동부(오하이오) 리전
         | us-west-1      | US 서부(북부 캘리포니아) 리전
         | us-west-2      | US 서부(오레곤) 리전
         | ca-central-1   | 캐나다(중부) 리전
         | eu-west-1      | EU(아일랜드) 리전
         | eu-west-2      | EU(런던) 리전
         | eu-west-3      | EU(파리) 리전
         | eu-north-1     | EU(스톡홀름) 리전
         | eu-south-1     | EU(밀라노) 리전
         | EU             | EU 리전
         | ap-southeast-1 | 아시아 태평양(싱가포르) 리전
         | ap-southeast-2 | 아시아 태평양(시드니) 리전
         | ap-northeast-1 | 아시아 태평양(도쿄) 리전
         | ap-northeast-2 | 아시아 태평양(서울) 리전
         | ap-northeast-3 | 아시아 태평양(오사카-로컬) 리전
         | ap-south-1     | 아시아 태평양(뭄바이) 리전
         | ap-east-1      | 아시아 태평양(홍콩) 리전
         | sa-east-1      | 남아메리카(상파울루) 리전
         | me-south-1     | 중동(바레인) 리전
         | af-south-1     | 아프리카(케이프타운) 리전
         | cn-north-1     | 중국(베이징) 리전
         | cn-northwest-1 | 중국(닝샤) 리전
         | us-gov-east-1  | AWS GovCloud(미국-동부) 리전
         | us-gov-west-1  | AWS GovCloud(미국) 리전

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되어 있지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      참고: S3에서 객체를 서버 측 복사할 때 S3는 원본에서 ACL을 복사하지 않고 새로 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성시 사용되는 Canned ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      설정되지 않은 경우 "acl"을 대신 사용합니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사용자에게 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹이 읽기 액세스를 받습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 액세스를 받습니다.
         |                    | 이러한 액세스를 버킷에 부여하는 것을 일반적으로 권장하지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AuthenticatedUsers 그룹이 읽기 액세스를 받습니다.

   --requester-pays
      S3 버킷과 상호 작용 시 요청자 지불 옵션을 활성화합니다.

   --server-side-encryption
      S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예시:
         | <unset> | 없음
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예시:
         | <unset> | 없음
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 Key의 ARN을 제공해야 합니다.

      예시:
         | <unset>                 | 없음
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C에 사용할 경우 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다.
      
      대신 --sse-customer-key-base64를 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터를 암호화/복호화하기 위해 암호화 키를 Base64 형식으로 인코딩된 상태로 제공해야 합니다.
      
      대신 --sse-customer-key를 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 둔 경우 sse_customer_key에서 자동으로 계산됩니다.
      

      예시:
         | <unset> | 없음

   --storage-class
      S3에 새로운 객체를 저장할 때 사용할 저장 클래스입니다.

      예시:
         | <unset>             | 기본값
         | STANDARD            | 표준 저장 클래스
         | REDUCED_REDUNDANCY  | 감소된 중복 저장 클래스
         | STANDARD_IA         | 표준 비표준 액세스 저장 클래스
         | ONEZONE_IA          | 한 가지 지역 비표준 액세스 저장 클래스
         | GLACIER             | Glacier 저장 클래스
         | DEEP_ARCHIVE        | Glacier Deep Archive 저장 클래스
         | INTELLIGENT_TIERING | Intelligent-Tiering 저장 클래스
         | GLACIER_IR          | Glacier Instant Retrieval 저장 클래스

   --upload-cutoff
      청크 업로드로 전환하는 경계입니다.
      
      이 경계보다 큰 파일은 chunk_size 단위로 청크화하여 업로드됩니다.
      최솟값은 0이고 최댓값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이거나 크기를 알 수 없는 파일("rclone rcat"으로부터 제공되거나 "rclone mount" 또는 Google 사진 또는 Google 문서로 업로드된 파일)을 업로드할 때, 이 청크 크기를 사용하여 멀티파트로 업로드됩니다.
      
      주의: "--s3-upload-concurrency"의 청크가 전송마다 메모리에 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 늘리면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대량 파일을 업로드할 때 10,000개의 청크 제한을 유지하기 위해 청크 크기를 자동으로 늘립니다.
      
      크기를 알 수 없는 파일은 구성된 chunk_size로 업로드됩니다. 기본적인 청크 크기는 5 MiB이며, 최대 10,000개의 청크를 가질 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상태의 정확성이 감소합니다. Rclone은 청크를 AWS SDK에 의해 버퍼링될 때 전송되었다고 간주하고 있지만 실제로 업로드 중일 수 있습니다. 청크 크기가 클수록 AWS SDK 버퍼도 커지고 진행률 보고가 진실과 더 멀어집니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      AWS S3의 10,000개 청크 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 대량 파일을 업로드할 때 10,000개의 청크 제한을 유지하기 위해 청크 크기를 자동으로 늘립니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 경계입니다.
      
      서버 측 복사할 필요가 있는 이 경계보다 큰 파일은 이 크기로 청크별로 복사됩니다.
      
      최솟값은 0이고 최댓값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이렇게 하면 대용량 파일을 업로드하기 시작할 때 오랜 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우를 사용할 것입니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대용량 파일을 고속 링크에서 소수의 개의 대용량 파일로 업로드하고 이 업로드가 대역폭을 완전히 활용하지 않는다면 이 값을 늘릴 수 있습니다.
      
   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      이 값을 true(기본값)로 설정하면 rclone은 경로 스타일 액세스를 사용하고 false로 설정하면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 서비스(예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)에서는 이 값을 false로 설정해야 합니다. rclone은 이 값을 제공자 설정에 따라 자동으로 설정합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      일반적으로 rclone은 v4 인증을 사용하며, 이 값이 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 시그니처가 작동하지 않는 경우에만이 플래그를 사용하세요(예: Jewel/v10 CEPH 이전).

   --use-accelerate-endpoint
      true인 경우 AWS S3 가속화 엔드포인트를 사용합니다.
      
      참조: [AWS S3 전송 가속화](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)

   --leave-parts-on-error
      true인 경우 실패 시 중단 업로드를 호출하지 않고 모두 S3에 성공적으로 업로드된 부분을 수동으로 복구합니다.
      
      여러 세션 간 업로드를 계속하는 경우 이 값을 true로 설정해야 합니다.
      
      경고: 완료되지 않은 멀티파트 업로드의 일부를 저장하면 S3에 대한 공간 사용량에 포함되고, 정리하지 않으면 추가 비용이 발생할 수 있습니다.
      

   --list-chunk
      리스트 청크의 크기(각 ListObject S3 요청에 대한 응답 목록)입니다.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청보다 더 많이 요청되어도 응답 목록을 1,000개로 잘라냅니다.
      Amazon S3에서는 이것이 전체적인 최대값이며 변경할 수 없습니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동 설정 (0).
      
      S3가 처음 출시될 때 버킷에서 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 등장했습니다. 이것은 훨씬 높은 성능을 제공하며 가능하면 사용해야 합니다.
      
      기본값인 0으로 설정된 경우 rclone은 공급자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못 추측하는 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      리스트를 URL로 인코딩할지 여부: true/false/unset
      
      일부 공급자가 리스트를 URL로 인코딩하도록 지원하는 경우 파일 이름에 제어 문자를 사용하는 경우 이렇게 하는 것이 더 안정적입니다. 이 값을 unset(기본값)으로 설정하면 rclone은 제공자 설정에 따라 적용할 것을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      설정되면 버킷의 존재를 확인하거나 생성하지 않습니다.
      
      알려진 버킷이 이미 존재하는 경우 rclone의 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 정적으로 전달되었습니다.
      

   --no-head
      업로드한 객체를 HEAD하여 무결성을 확인하지 않습니다.
      
      rclone은 POST로 객체를 업로드한 후 200 OK 메시지를 수신하면 업로드가 올바르게 수행된 것으로 가정합니다.
      
      특히 다음을 가정합니다:
      
      - metadata(수정 시간, 저장 클래스 및 콘텐츠 유형 포함)가 업로드한 것과 동일하다.
      - 크기가 업로드한 것과 동일하다.
      
      단일 부분 PUT에 대한 응답에서 rclone은 다음을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이 항목들은 읽지 않습니다.
      
      크기를 알 수 없는 소스 객체는 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 높아져 크기가 잘못된 경우 등의 알 수 없는 업로드 실패 가능성이 커져 일반적인 운영에는 권장하지 않습니다. 실제로 알 수 없는 업로드 실패 가능성은 매우 낮습니다.
      


   --no-head-object
      객체를 가져오기 전에 HEAD를 수행하지 않습니다.

   --encoding
      Backend의 인코딩 방식입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 비울지를 결정합니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)은 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼를 풀에서 얼마나 자주 제거할지를 조절합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부를 결정합니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 문제가 발생합니다. HTTP/2는 기본적으로 s3 백엔드에서 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드용 사용자 지정 엔드포인트입니다.
      보통 AWS S3는 클라우드프론트 CDN URL로 설정합니다. 클라우드프론트 네트워크를 통해 다운로드되는 데이터에 대해 AWS S3는 더 저렴한 egress를 제공합니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에 ETag를 사용할지 여부
      
      true, false 또는 세트되어 있지 않은 값으로 설정할 수 있습니다.
      

   --use-presigned-request
      단일 부분 업로드에 대해 사전 사인된 요청 또는 PutObject을 사용할지 여부
      
      이 값을 false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      Rclone < 1.59 버전은 단일 부분 업로드에 대해 사전 사인된 요청을 사용하고 이 플래그를 true로 설정하면 해당 기능을 다시 활성화합니다. 이는 예외적인 경우나 테스트에만 필요합니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에있었던 파일 버전을 표시합니다.
      
      매개변수는 날짜, "2006-01-02", "2006-01-02 15:04:05" 또는 그보다 오래될 수 있도록 지속 시간, 예: "100d" 또는 "1h"입니다.
      
      이 값을 사용할 때는 파일 쓰기 작업을 허용하지 않으므로 파일 업로드 또는 삭제를 할 수 없습니다.
      
      다음을 참조하세요. [시간 옵션 설명서](/docs/#time-option)에서 유효한 형식을 확인하세요.
      

   --decompress
      이 설정을 사용하면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 수신되는 개체를 "Content-Encoding: gzip"로 압축 해제합니다. 이로 인해 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 개체를 gzip으로 압축할 수 있는 경우 이 설정을 설정합니다.
      
      일반적으로 제공자는 객체를 다운로드할 때 수정하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급자(예: Cloudflare)는 "Content-Encoding: gzip"로 업로드되지 않은 경우에도 객체를 gzip으로 압축 할 수 있습니다.
      
      이 설정을 설정하고 rclone이 "Content-Encoding: gzip"로 설정된 객체를 청크 전송 인코딩과 함께 다운로드한다면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset(기본값)로 설정되면 rclone은 제공자 설정에 따라 적용할 것을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 제한합니다

   --sts-endpoint
      STS의 엔드포인트입니다.
      
      AWS를 사용하는 경우 리전의 기본 엔드포인트를 사용하려면 비워 둡니다.


OPTIONS:
   --access-key-id value           AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                     버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value                S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                      AWS 자격증명을 런타임 환경(환경 변수 또는 env vars 또는 EC2/ECS 메타데이터)에서 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말을 표시합니다.
   --location-constraint value     리전과 일치해야 하는 위치 제한입니다. [$LOCATION_CONSTRAINT]
   --region value                  연결할 리전입니다. [$REGION]
   --secret-access-key value       AWS 비밀 액세스 키(암호)입니다. [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID를 사용하는 경우 Key의 ARN을 제공해야 합니다. [$SSE_KMS_KEY_ID]
   --storage-class value           S3에 새로운 객체를 저장할 때 사용할 저장 클래스입니다. [$STORAGE_CLASS]

고급

   --bucket-acl value               버킷 생성시 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 경계입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     이 설정을 사용하면 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드용 사용자 지정 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 Backend의 인코딩 방식입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
 