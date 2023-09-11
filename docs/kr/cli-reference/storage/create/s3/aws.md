# Amazon Web Services (AWS) S3

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 aws - Amazon Web Services (AWS) S3

사용법:
   singularity storage create s3 aws [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격증명 가져옵니다 (환경 변수 또는 env vars 또는 EC2/ECS meta 데이터).
      
      access_key_id 및 secret_access_key이 비워져 있는 경우에만 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 자격증명 입력.
         | true  | 환경 (env vars 또는 IAM)에서 AWS 자격증명 가져옵니다.

   --access-key-id
      AWS Access Key ID입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --region
      연결할 리전입니다.

      예:
         | us-east-1      | 기본 엔드포인트 - 확실하지 않은 경우 좋은 선택입니다.
         |                | US Region. Northern Virginia 또는 Pacific Northwest.
         |                | 위치 제약 조건을 비워 둡니다.
         | us-east-2      | US East (Ohio) 리전입니다.
         |                | 위치 제약 조건: us-east-2
         | us-west-1      | US West (Northern California) 리전입니다.
         |                | 위치 제약 조건: us-west-1
         | us-west-2      | US West (Oregon) 리전입니다.
         |                | 위치 제약 조건: us-west-2
         | ca-central-1   | Canada (Central) 리전입니다.
         |                | 위치 제약 조건: ca-central-1
         | eu-west-1      | EU (Ireland) 리전입니다.
         |                | 위치 제약 조건: EU 또는 eu-west-1
         | eu-west-2      | EU (London) 리전입니다.
         |                | 위치 제약 조건: eu-west-2
         | eu-west-3      | EU (Paris) 리전입니다.
         |                | 위치 제약 조건: eu-west-3
         | eu-north-1     | EU (Stockholm) 리전입니다.
         |                | 위치 제약 조건: eu-north-1
         | eu-south-1     | EU (Milan) 리전입니다.
         |                | 위치 제약 조건: eu-south-1
         | eu-central-1   | EU (Frankfurt) 리전입니다.
         |                | 위치 제약 조건: eu-central-1
         | ap-southeast-1 | 아시아 태평양 (싱가포르) 리전입니다.
         |                | 위치 제약 조건: ap-southeast-1
         | ap-southeast-2 | 아시아 태평양 (시드니) 리전입니다.
         |                | 위치 제약 조건: ap-southeast-2
         | ap-northeast-1 | 아시아 태평양 (도쿄) 리전입니다.
         |                | 위치 제약 조건: ap-northeast-1
         | ap-northeast-2 | 아시아 태평양 (서울) 리전입니다.
         |                | 위치 제약 조건: ap-northeast-2
         | ap-northeast-3 | 아시아 태평양 (오사카-로컬) 리전입니다.
         |                | 위치 제약 조건: ap-northeast-3
         | ap-south-1     | 아시아 태평양 (뭄바이) 리전입니다.
         |                | 위치 제약 조건: ap-south-1
         | ap-east-1      | 아시아 태평양 (홍콩) 리전입니다.
         |                | 위치 제약 조건: ap-east-1
         | sa-east-1      | 남아메리카 (상파울로) 리전입니다.
         |                | 위치 제약 조건: sa-east-1
         | me-south-1     | 중동 (바레인) 리전입니다.
         |                | 위치 제약 조건: me-south-1
         | af-south-1     | 아프리카 (케이프 타운) 리전입니다.
         |                | 위치 제약 조건: af-south-1
         | cn-north-1     | 중국 (북경) 리전입니다.
         |                | 위치 제약 조건: cn-north-1
         | cn-northwest-1 | 중국 (닝샤) 리전입니다.
         |                | 위치 제약 조건: cn-northwest-1
         | us-gov-east-1  | AWS GovCloud (US-East) 리전입니다.
         |                | 위치 제약 조건: us-gov-east-1
         | us-gov-west-1  | AWS GovCloud (US) 리전입니다.
         |                | 위치 제약 조건: us-gov-west-1

   --endpoint
      S3 API의 엔드포인트입니다.
      
      AWS의 기본 엔드포인트를 사용하려면 비워 두세요.

   --location-constraint
      위치 제약 조건 - 리전과 일치해야 합니다.
      
      버킷을 생성할 때만 사용됩니다.

      예:
         | <unset>        | US Region. Northern Virginia 또는 Pacific Northwest일 때 비웁니다.
         | us-east-2      | US East (Ohio) 리전일 때
         | us-west-1      | US West (Northern California) 리전일 때
         | us-west-2      | US West (Oregon) 리전일 때
         | ca-central-1   | Canada (Central) 리전일 때
         | eu-west-1      | EU (Ireland) 리전일 때
         | eu-west-2      | EU (London) 리전일 때
         | eu-west-3      | EU (Paris) 리전일 때
         | eu-north-1     | EU (Stockholm) 리전일 때
         | eu-south-1     | EU (Milan) 리전일 때
         | EU             | EU 리전일 때
         | ap-southeast-1 | 아시아 태평양 (싱가포르) 리전일 때
         | ap-southeast-2 | 아시아 태평양 (시드니) 리전일 때
         | ap-northeast-1 | 아시아 태평양 (도쿄) 리전일 때
         | ap-northeast-2 | 아시아 태평양 (서울) 리전일 때
         | ap-northeast-3 | 아시아 태평양 (오사카-로컬) 리전일 때
         | ap-south-1     | 아시아 태평양 (뭄바이) 리전일 때
         | ap-east-1      | 아시아 태평양 (홍콩) 리전일 때
         | sa-east-1      | 남아메리카 (상파울로) 리전일 때
         | me-south-1     | 중동 (바레인) 리전일 때
         | af-south-1     | 아프리카 (케이프 타운) 리전일 때
         | cn-north-1     | 중국 (북경) 리전일 때
         | cn-northwest-1 | 중국 (닝샤) 리전일 때
         | us-gov-east-1  | AWS GovCloud (US-East) 리전일 때
         | us-gov-west-1  | AWS GovCloud (US) 리전일 때

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 canned ACL입니다.
      
      이 ACL은 버킷을 생성할 때와 객체를 생성할 때 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3에서 객체 복사 시, S3는 A CL을 원본에서 복사하지 않고 새로 생성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고, 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 canned ACL입니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 생성할 때에만 적용됩니다. 설정되지 않은 경우 "acl"을 대신 사용합니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고, 기본값(비공개)이 사용됩니다.
      

      예:
         | private            | 소유자만 FULL_CONTROL을 얻습니다.
         |                    | 다른 사용자는 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자만 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 액세스를 얻습니다.
         | public-read-write  | 소유자만 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻습니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자만 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스를 얻습니다.

   --requester-pays
      S3 버킷과 상호 작용할 때 요청자 지불 옵션 활성화.

   --server-side-encryption
      S3에 이 객체를 저장할 때 사용하는 서버 측 암호화 알고리즘.

      예:
         | <unset> | 없음
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, S3에 이 객체를 저장할 때 사용하는 서버 측 암호화 알고리즘.

      예:
         | <unset> | 없음
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 Key의 ARN을 제공해야 합니다.

      예:
         | <unset>                 | 없음
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다.
      
      대신 --sse-customer-key-base64를 사용할 수도 있습니다.

      예:
         | <unset> | 없음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는 비밀 암호화 키를 base64 형식으로 인코딩하여 제공해야 합니다.
      
      대신 --sse-customer-key를 사용할 수도 있습니다.

      예:
         | <unset> | 없음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우, 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다. (선택 사항)
      
      비워두면 sse_customer_key에서 자동으로 계산됩니다.
      

      예:
         | <unset> | 없음

   --storage-class
      S3에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다.

      예:
         | <unset>             | 기본값
         | STANDARD            | 기본 저장 수준
         | REDUCED_REDUNDANCY  | 감소된 중복 저장 수준
         | STANDARD_IA         | 표준 희귀 액세스 저장 수준
         | ONEZONE_IA          | 단일 지역 희귀 액세스 저장 수준
         | GLACIER             | Glacier 저장 수준
         | DEEP_ARCHIVE        | Glacier Deep Archive 저장 수준
         | INTELLIGENT_TIERING | Intelligent-Tiering 저장 수준
         | GLACIER_IR          | Glacier Instant Retrieval 저장 수 

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일 또는 크기가 알려지지 않은 파일(예: "rclone rcat"에서 또는 "rclone mount" 또는 구글 사진 또는 구글 문서로 업로드된 파일)을 업로드할 때, 이 청크 크기로 multipart 업로드를 통해 업로드됩니다.
      
      참고로 "--s3-upload-concurrency" 크기의 청크는 버퍼링됩니다.
      
      빠른 링크에서 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 높이면 전송 속도가 빨라집니다.
      
      라클론은 알려진 크기의 큰 파일을 전송 시 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000 청크입니다.
      
      chunk 크기를 증가시키면 '-P' 플래그와 함께 표시되는 진행 상태 통계의 정확도가 감소합니다. rclone은 AWS SDK에 의해 버퍼링될 때 chunk가 전송되었으므로 chunk 크기가 지정된 곳이 실제로 업로드되었다고 간주합니다. chunk 크기가 클수록 AWS SDK 버퍼와 진행률 보고는 진실과 더 멀어집니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 파트 수입니다.
      
      이 옵션은 멀티파트 업로드 수행 시 사용되는 멀티파트 청크의 최대 개수를 정의합니다.
      
      10,000 청크의 AWS S3 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때이 청크 크기를 자동으로 증가시켜 10,000 청크 수 제한을 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값입니다.
      
      이보다 큰 파일을 서버 측 복사해야 하는 경우 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 데이터 무결성 확인에는 이점이 있지만 크기가 큰 파일을 업로드하는 데 오랜 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리가 기본값입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로파일입니다.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 파일로부터 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 기본값이 됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크는 이 동시에 업로드됩니다.
      
      고속 링크에서 큰 파일을 업로드하고 대역폭을 완전히 활용하지 않으면 이 값을 높이면 전송 속도가 증가할 수 있습니다.

   --force-path-style
      true이면 path 스타일 액세스를, false이면 가상 호스팅 스타일 액세스를 사용합니다.
      
      이 값을 true(기본)로 설정하면 path 스타일 액세스를 사용하며, false이면 가상 path 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 이 값을 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이를 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      이 값이 false(기본값)로 설정되어 있으면 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 이전 CEPH의 경우.

   --use-accelerate-endpoint
      true이면 AWS S3 가속 엔드포인트를 사용합니다.
      
      자세한 내용은 [AWS S3 Transfer acceleration](https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration-examples.html)을 참조하세요.

   --leave-parts-on-error
      true이면 실패 시 업로드 중단을 호출하지 않고 S3에 모든 성공적으로 업로드된 파트를 수동으로 복구합니다.
      
      다른 세션 간에 업로드를 다시 시작해야 하는 경우 true로 설정해야 합니다.
      
      경고: 완료되지 않은 멀티파트 업로드의 일부를 S3에 저장하면 S3의 공간 사용과 추가 비용이 발생합니다.
      

   --list-chunk
      리스트 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기)입니다.
      
      이 옵션은 AWS S3 사양에 따라 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있는 것입니다.
      대부분의 서비스는 요청을 통해 요청한 모든 객체 목록을 최대 1000개까지로 자르지만 요청한 것보다 많은 경우에도 최대 1000개로 잘립니다.
      AWS S3에서는 이것이 전역 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동(0).
      
      S3가 처음 출시되었을 때 버킷의 객체를 열람하기 위해 ListObjects 호출만 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 소개되었습니다. 이 호출은 훨씬 높은 성능을 제공하며 가능한한 사용해야 합니다.
      
      기본값인 0으로 설정된 경우 rclone은 공급자 설정에 따라 어떤 객체 목록 방법을 호출할지 추측합니다. 추측이 잘못된 경우 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할 것인지 여부: true/false/unset
      
      일부 공급자는 목록 URL이 지원되고 가능한 경우 파일 이름의 제어 문자를 사용하는 경우에 더 신뢰할 수 있습니다. 이 값이 unset으로 설정된 경우 rclone은 공급자 설정에 따라 적용할 항목을 결정하지만 이곳에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      설정되면 버킷을 확인하거나 생성하지 않습니다.
      
      알려진 버킷이 이미 있는 경우 rclone이 수행하는 거래 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우 필요할 수도 있습니다. v1.52.0 이전에는 버그로 인해 이 설정이 무시되었을 것입니다.
      

   --no-head
      설정하면 업로드한 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone은 PUT로 객체를 업로드한 후에 200 OK 메세지를 받으면 업로드가 제대로 된 것으로 가정합니다.
      
      특히 다음과 같이 가정합니다.:
      
      - 메타데이터, 수정 시간, 저장 수준 및 콘텐츠 유형이 업로드한 대로 있었음
      - 크기가 업로드한 대로 있었음
      
      PUT로 단일 부분 객체에 대한 응답에서 rclone은 다음 항목을 읽습니다.:
      
      - MD5SUM
      - 업로드한 날짜
      
      여러 파트 UPLOAD인 경우 이러한 항목을 읽지 않습니다.
      
      알 수 없는 길이의 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 증가하지만, 큰 파일의 업로드 실패 가능성은 실제로 매우 낮기 때문에 정상적인 운영에는 권장되지 않습니다.
      

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 시간 간격입니다.
      
      추가 버퍼를 필요로하는 업로드(예: multipart)은 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼를 풀에서 제거하는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용 비활성화.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 대해 아직 해결되지 않은 문제가 있습니다. 기본적으로 s3 백엔드에서는 HTTP/2가 활성화되지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 다운로드된 데이터를 통해 더 저렴한 egress를 제공합니다.

   --use-multipart-etag
      검증을 위해 multipart 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 unset으로 설정해야 합니다. 기본값은 공급자를 기반으로 선택됩니다.
      

   --use-presigned-request
      단일 부분 업로드에 대해 서명된 요청 또는 PutObject을 사용할지 여부
      
      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      rclone <1.59 버전은 단일 부분 객체 업로드에 대해 서명된 요청을 사용하고 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이 기능은 예외적인 상황이나 테스트 외에는 필요하지 않을 것입니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정한 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜인 "2006-01-02", datetime인 "2006-01-02 15:04:05" 또는 해당 시간까지의 기간, 예: "100d" 또는 "1h"여야 합니다.
      
      이를 사용할 때 파일 쓰기 작업은 허용되지 않으므로 파일 업로드 또는 삭제를 수행할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [타임 옵션 도움말](/docs/#time-option)을 참조하세요.
      

   --decompress
      gzip 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"로 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone이 이러한 파일을 "Content-Encoding: gzip"와 함께 수신 시 압축을 해제합니다. rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는 경우 설정하십시오.
      
      일반적으로 공급자는 객체를 다운로드 할 때 객체를 변경하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급자는 `Content-Encoding: gzip`로 업로드되지 않은 객체를 gzip으로 압축할 수 있습니다 (예: Cloudflare).
      
      이러한 상황이 발생하는 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하면 rclone이 chunked 전송 인코딩 및 `Content-Encoding: gzip`가 설정된 개체를 다운로드 시 개체를 압축 해제합니다.
      
      unset으로 설정되어 있는 경우 rclone은 공급자 설정에 따라 적용하는 방법을 선택하지만 이곳에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.

   --sts-endpoint
      STS의 엔드포인트입니다.
      
      기본 엔드포인트를 사용하려면 비워 두세요.


옵션:
   --access-key-id value           AWS Access Key ID입니다. [$ACCESS_KEY_ID]
   --acl value                     버킷 생성 및 객체 저장 또는 복사 시 사용되는 canned ACL입니다. [$ACL]
   --endpoint value                S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                      런타임에서 AWS 자격증명 가져옵니다 (환경 변수 또는 EC2 / ECS 메타 데이터 또는 env vars). (default: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     위치 제약 조건 - 리전과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value                  연결할 리전입니다. [$REGION]
   --secret-access-key value       AWS Secret Access Key (비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에 이 객체를 저장할 때 사용하는 서버 측 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID를 사용하는 경우 Key의 ARN을 제공해야 합니다. [$SSE_KMS_KEY_ID]
   --storage-class value           S3에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값입니다. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정하면 gzip으로 압축된 객체를 압축 해제합니다. (default: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용을 비활성화합니다. (default: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 path 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다. (default: true) [$FORCE_PATH_STYLE]
   --leave-parts-on-error           true이면 실패 시 업로드 중단을 호출하지 않고 S3에 모든 성공적으로 업로드된 파트를 수동으로 복구합니다. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --list-chunk value               리스트 청크의 크기입니다. (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할 것인지 여부입니다. (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전입니다: 1, 2 또는 0은 자동입니다. (default: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 파트 수입니다. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간 간격입니다. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는 경우 설정하십시오. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정되면 버킷을 확인하거나 생성하지 않습니다. (default: false) [$NO_CHECK_BUCKET]
   --no-head                        설정하면 업로드한 객체를 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (default: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다. (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로파일입니다. [$PROFILE]
   --requester-pays                 S3 버킷과 상호 작용할 때 요청자 지불 옵션 활성화. (default: false) [$REQUESTER_PAYS]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우, S3에 이 객체를 저장할 때 사용하는 서버 측 암호화 알고리즘입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는 비밀 암호화 키를 base64 형식으로 인코딩하여 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우, 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다. (선택 사항) [$SSE_CUSTOMER_KEY_MD5]
   --sts-endpoint value             STS의 엔드포인트입니다. [$STS_ENDPOINT]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값입니다. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-accelerate-endpoint        true이면 AWS S3 가속 엔드포인트를 사용합니다. (default: false) [$USE_ACCELERATE_ENDPOINT]
   --use-multipart-etag value       검증을 위해 multipart 업로드에서 ETag를 사용할지 여부 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 대해 서명된 요청 또는 PutObject를 사용할지 여부 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (default: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정한 시간에 표시합니다. (default: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (default: false) [$VERSIONS]

   General

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}