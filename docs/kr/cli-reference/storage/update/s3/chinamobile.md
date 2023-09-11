# China Mobile Ecloud Elastic Object Storage (EOS)

{% code fullWidth="true" %}
```
명령어:
   singularity storage update s3 chinamobile - 중국 모바일 Ecloud 탄력형 객체 스토리지 (EOS)

사용법:
   singularity storage update s3 chinamobile [command options] <name|id>

설명:
   --env-auth
      런타임으로부터 AWS 자격 증명 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터)을 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경 (환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --endpoint
      중국 모바일 Ecloud 탄력형 객체 스토리지 (EOS) API의 엔드포인트입니다.

      예시:
         | eos-wuxi-1.cmecloud.cn      | 기본 엔드포인트 - 확신이 없는 경우 좋은 선택입니다.
         |                             | 중국 동쪽 (수주)
         | eos-jinan-1.cmecloud.cn     | 중국 동쪽 (진안)
         | eos-ningbo-1.cmecloud.cn    | 중국 동쪽 (항저우)
         | eos-shanghai-1.cmecloud.cn  | 중국 동쪽 (상하이-1)
         | eos-zhengzhou-1.cmecloud.cn | 중국 중간 (정저우)
         | eos-hunan-1.cmecloud.cn     | 중국 중간 (창사-1)
         | eos-zhuzhou-1.cmecloud.cn   | 중국 중간 (창사-2)
         | eos-guangzhou-1.cmecloud.cn | 중국 남쪽 (광저우-2)
         | eos-dongguan-1.cmecloud.cn  | 중국 남쪽 (광저우-3)
         | eos-beijing-1.cmecloud.cn   | 중국 북쪽 (베이징-1)
         | eos-beijing-2.cmecloud.cn   | 중국 북쪽 (베이징-2)
         | eos-beijing-4.cmecloud.cn   | 중국 북쪽 (베이징-3)
         | eos-huhehaote-1.cmecloud.cn | 중국 북쪽 (후허하오테)
         | eos-chengdu-1.cmecloud.cn   | 중국 남서쪽 (청두)
         | eos-chongqing-1.cmecloud.cn | 중국 남서쪽 (충칭)
         | eos-guiyang-1.cmecloud.cn   | 중국 남서쪽 (구이양)
         | eos-xian-1.cmecloud.cn      | 중국 남서쪽 (샤안)
         | eos-yunnan.cmecloud.cn      | 윈난 중국 (쿤밍)
         | eos-yunnan-2.cmecloud.cn    | 윈난 중국 (쿤밍-2)
         | eos-tianjin-1.cmecloud.cn   | 청진 중국 (청진)
         | eos-jilin-1.cmecloud.cn     | 지린 중국 (창춘)
         | eos-hubei-1.cmecloud.cn     | 후베이 중국 (샹양)
         | eos-jiangxi-1.cmecloud.cn   | 장시 중국 (난창)
         | eos-gansu-1.cmecloud.cn     | 간수 중국 (란저우)
         | eos-shanxi-1.cmecloud.cn    | 산시 중국 (타이위안)
         | eos-liaoning-1.cmecloud.cn  | 리아오닝 중국 (셴양)
         | eos-hebei-1.cmecloud.cn     | 허베이 중국 (시징저앙)
         | eos-fujian-1.cmecloud.cn    | 후난 중국 (샤먼)
         | eos-guangxi-1.cmecloud.cn   | 광시 중국 (난닝)
         | eos-anhui-1.cmecloud.cn     | 안후이 중국 (후아난)

   --location-constraint
      엔드포인트와 일치해야 하는 위치 제약 조건입니다.
      
      버킷을 생성할 때만 사용됩니다.

      예시:
         | wuxi1      | 중국 동쪽 (수주)
         | jinan1     | 중국 동쪽 (진안)
         | ningbo1    | 중국 동쪽 (항저우)
         | shanghai1  | 중국 동쪽 (상하이-1)
         | zhengzhou1 | 중국 중간 (정저우)
         | hunan1     | 중국 중간 (창사-1)
         | zhuzhou1   | 중국 중간 (창사-2)
         | guangzhou1 | 중국 남쪽 (광저우-2)
         | dongguan1  | 중국 남쪽 (광저우-3)
         | beijing1   | 중국 북쪽 (베이징-1)
         | beijing2   | 중국 북쪽 (베이징-2)
         | beijing4   | 중국 북쪽 (베이징-3)
         | huhehaote1 | 중국 북쪽 (후허하오테)
         | chengdu1   | 중국 남서쪽 (청두)
         | chongqing1 | 중국 남서쪽 (충칭)
         | guiyang1   | 중국 남서쪽 (구이양)
         | xian1      | 중국 남서쪽 (샤안)
         | yunnan     | 윈난 중국 (쿤밍)
         | yunnan2    | 윈난 중국 (쿤밍-2)
         | tianjin1   | 청진 중국 (청진)
         | jilin1     | 지린 중국 (창춘)
         | hubei1     | 후베이 중국 (샹양)
         | jiangxi1   | 장시 중국 (난창)
         | gansu1     | 간수 중국 (란저우)
         | shanxi1    | 산시 중국 (타이위안)
         | liaoning1  | 리아오닝 중국 (셴양)
         | hebei1     | 허베이 중국 (시징저앙)
         | fujian1    | 후난 중국 (샤먼)
         | guangxi1   | 광시 중국 (난닝)
         | anhui1     | 안후이 중국 (후아난)

   --acl
      객체를 생성하거나 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다.
      
      이 ACL은 객체를 만들거나 bucket_acl이 설정되지 않은 경우에도 생성하는 데 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3는 서버 측 객체를 복사할 때 ACL을 소스에서 복사하지 않고 새로 작성하기 때문에 이 ACL이 적용됩니다.
      
      acl이 비어 있는 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 canned ACL입니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      bucket_acl이 설정되지 않은 경우에만 적용됩니다.
      
      acl과 bucket_acl이 모두 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람들은 액세스 권한이 없습니다 (기본 설정).
         | public-read        | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹에 READ 액세스 권한이 있습니다.
         | public-read-write  | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹에 READ 및 WRITE 액세스 권한이 있습니다.
         |                    | 버킷에서 이 ACL을 부여하는 것은 일반적으로 권장하지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹에 READ 액세스 권한이 있습니다.

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

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화/복호화하는 데 사용하는 비밀 암호화 키를 제공할 수 있습니다.
      
      --sse-customer-key-base64를 대신 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터를 암호화/복호화하기 위해 base64 형식으로 인코딩된 비밀 암호화 키를 제공해야 합니다.
      
      --sse-customer-key를 대신 제공할 수도 있습니다.

      예시:
         | <unset> | 없음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 비밀 암호화 키 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 둘 경우 sse_customer_key에서 자동으로 계산됩니다.
      

      예시:
         | <unset> | 없음

   --storage-class
      ChinaMobile에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다.

      예시:
         | <unset>     | 기본값
         | STANDARD    | 표준 스토리지 클래스
         | GLACIER     | 아카이브 스토리지 모드
         | STANDARD_IA | 적은 액세스 스토리지 모드

   --upload-cutoff
      청크로 업로드 전환하는 데 사용되는 컷오프입니다.
      
      이보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --chunk-size
      업로드하는 데 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나 알 수 없는 크기의 파일(예:"rclone rcat"에서 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에 업로드된 파일)을 업로드할 때 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      알림: "--s3-upload-concurrency" 청크 크기는 전송당 메모리에 저장되는 것입니다.
      
      고속 링크를 통해 대용량 파일을 전송하며 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 빨라집니다.
      
      큰 파일을 업로드할 때, rclone은 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      크기를 알 수 없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기가 5 MiB이고 최대 10,000 청크까지 존재할 수 있기 때문에 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 통계의 정확성이 감소합니다. rclone은 AWS SDK에서 버퍼에 보관된 채크를 보낸 것으로 처리하지만, 실제로 업로드되는 경우가 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률을 의미합니다.

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이는 서비스가 AWS S3 10,000 청크 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      rclone은 크기가 알려진 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 초과하지 않도록 합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 데 사용되는 컷오프입니다.
      
      해당 크기보다 큰 파일을 서버 측에서 복사해야 하는 경우 이 크기의 청크로 복사됩니다.
      
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이것은 데이터 무결성 확인에는 좋지만 대용량 파일의 업로드 시작에는 긴 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일에 대한 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      비워 둘 경우 환경 변수 "AWS_PROFILE" 또는 설정되지 않은 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      대역폭을 완전히 활용하지 못하는 상황에서 대용량 파일을 적은 수로 업로드하는 경우 이 값을 증가시키면 전송을 가속화하는 데 도움이 될 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급 업체(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이를 false로 설정해야 합니다. rclone은 공급 업체 설정에 따라 자동으로 수행됩니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false(기본값)인 경우 rclone은 v4 인증을 사용하고 설정하면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요(예: Jewel/v10 CEPH 이전).

   --list-chunk
      리스트 청크의 크기(ListObject S3 요청마다 응답 리스트)입니다.
      
      이 옵션은 AWS S3 사양의 max-items 또는 page-size로 알려져 있는 최대 요청 청크 크기 해당 암호화 옵션이 지원되는 경우 파일 이름에 제어 문자를 사용할 때 더 신뢰할 수 있습니다.
      이 값이 unset(기본값)로 설정된 경우 rclone은 제공자 설정에 따라 적용할 것을 선택하지만 이곳에서 rclone의 선택을 재정의할 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동으로 0입니다.
      
      S3가 처음에 출시되었을 때, 버킷의 객체를 열거하기 위해 ListObjects 호출만으로 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본 설정인 0으로 설정하면 rclone은 제공자 설정에 따라 list 객체 방법을 추측합니다. 잘못된 추측인 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록 url 인코딩 여부: true/false/unset
      
      일부 공급 업체는 URL 인코딩 목록을 지원하며 가능한 경우 제어 문자를 포함하는 파일 이름을 사용할 때 이 방법이 더 신뢰할 수 있습니다. unset으로 설정하는 경우 rclone은 공급자 설정을 따라 어떤 것을 적용할지 선택합니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않으려면 설정하세요.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      또한 사용자에게 버킷 생성 권한이 없으면 필요할 수도 있습니다. 버전 1.52.0 이전의 경우 버그로 인해 조용히 전달되었습니다.
      

   --no-head
      업로드한 객체의 정합성을 확인하기 위해 HEAD를 사용하지 않습니다.
      
      rclone이 PUT로 객체를 업로드한 후 200 OK 메시지를 수신하면 제대로 업로드된 것으로 간주합니다.
      
      특히 다음 사항에 대해 가정합니다:
      
      - 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드한 것과 동일합니다.
      - 크기가 업로드한 것과 동일합니다.
      
      단일 부분 PUT의 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      길이를 알 수 없는 소스 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 전송 실패의 발생 가능성이 증가하며, 특히 잘못된 크기인 경우이므로 일반적인 작업에 권장되지 않습니다. 실제로, 이 플래그를 설정해도 전송 실패의 가능성은 매우 작습니다.

   --no-head-object
      객체를 가져올 때 HEAD를 GET보다 먼저 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당에 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관련된 이슈가 해결되지 않은 상태입니다. HTTP/2는 s3 백엔드의 기본 설정이며, 여기서는 비활성화할 수 있습니다. 이 이슈가 해결되면 이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 이이그레스를 제공합니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false, 또는 기본 제공자의 기본값으로 설정할 수 있습니다.
      

   --use-presigned-request
      단일 부분 업로드에 미리 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false이면 rclone은 개체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 < 1.59는 단일 부분 객체를 업로드하기 위해 미리 서명된 요청을 사용하며, 이 플래그를 true로 설정하면 해당 기능을 다시 활성화합니다. 이는 특수한 상황이나 테스트를 위해서만 필요하지만 예외가 아니라면 이 기능을 사용할 필요는 없습니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜인 "2006-01-02", 날짜와 시간인 "2006-01-02 15:04:05"이거나 그 이래서 만료된 기간인 "100d" 또는 "1h"입니다.
      
      이렇게 사용하는 경우 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"로 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"로 받아들이기 때문에 크기와 해시를 확인하지 못하지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 gzip으로 객체를 압축할 수 있는 경우 설정하세요.
      
      일반적으로 제공 업체는 파일을 다운로드할 때 객체를 변경하지 않습니다. 파일이 `Content-Encoding: gzip`로 업로드되지 않으면 다운로드할 때 설정되지 않습니다.
      
      그러나 일부 제공업체는 gzip으로 압축된 객체는 아니지만 압축합니다(예: Cloudflare).
      
      이러한 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip과 청크 전송 인코딩이 설정된 객체를 다운로드할 때 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본 설정)이 설정되면 rclone은 제공자 설정에 따라 적용할 것을 선택하지만 이곳에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기 지원을 억제합니다


옵션:
   --access-key-id value           AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                     객체를 만들거나 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다. [$ACL]
   --endpoint value                중국 모바일 Ecloud 탄력형 객체 스토리지 (EOS) API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                      런타임으로부터 AWS 자격 증명 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터)을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     엔드포인트와 일치해야 하는 위치 제약 조건입니다. [$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS 비밀 액세스 키(비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --storage-class value           ChinaMobile에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드하는 데 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 데 사용되는 컷오프입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기(ListObject S3 요청마다 응답 리스트)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          리스트 url 인코딩 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동으로 0입니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip으로 객체를 압축할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 정합성을 확인하기 위해 HEAD를 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 HEAD를 GET보다 먼저 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기 지원을 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일에 대한 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화/복호화하는 데 사용하는 비밀 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우 데이터를 암호화/복호화하기 위해 base64 형식으로 인코딩된 비밀 키를 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우 비밀 키 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크로 업로드 전환하는 데 사용되는 컷오프입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 미리 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}