# IBM COS S3

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 ibmcos - IBM COS S3

사용법:
   singularity storage create s3 ibmcos [command options] [arguments...]

설명:
   --env-auth
      AWS 자격증명을 런타임(환경 변수 또는 EC2/ECS 메타 데이터)에서 가져옵니다.
      
      access_key_id 및 secret_access_key가 빈 값인 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격증명을 가져옵니다.

   --access-key-id
      AWS Access Key ID입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워 둡니다.

   --region
      연결할 리전입니다.
      
      S3 클론을 사용하는 경우 리전이 없으므로 비워 둡니다.

      예제:
         | <unset>            | 확실하지 않은 경우 이 값을 사용하세요.
         |                    | v4 서명와 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 경우에만 사용합니다.
         |                    | 예: Jewel/v10 CEPH 이전.
  
   --endpoint
      IBM COS S3 API의 엔드포인트입니다.
      
      IBM COS 온프레미스를 사용하는 경우 지정하세요.

      예제:
         | s3.us.cloud-object-storage.appdomain.cloud               | Cross Region US 엔드포인트
         | s3.dal.us.cloud-object-storage.appdomain.cloud           | Cross Region Dallas 엔드포인트
         | s3.wdc.us.cloud-object-storage.appdomain.cloud           | Cross Region Washington DC 엔드포인트
         | s3.sjc.us.cloud-object-storage.appdomain.cloud           | Cross Region San Jose 엔드포인트
         | s3.private.us.cloud-object-storage.appdomain.cloud       | Cross Region Private 엔드포인트
         | s3.private.dal.us.cloud-object-storage.appdomain.cloud   | Cross Region Dallas Private 엔드포인트
         | s3.private.wdc.us.cloud-object-storage.appdomain.cloud   | Cross Region Washington DC Private 엔드포인트
         | s3.private.sjc.us.cloud-object-storage.appdomain.cloud   | Cross Region San Jose Private 엔드포인트
         | s3.us-east.cloud-object-storage.appdomain.cloud          | East Region US 엔드포인트
         | s3.private.us-east.cloud-object-storage.appdomain.cloud  | East Region US Private 엔드포인트
         | s3.us-south.cloud-object-storage.appdomain.cloud         | South Region US 엔드포인트
         | s3.private.us-south.cloud-object-storage.appdomain.cloud | South Region US Private 엔드포인트
         | s3.eu.cloud-object-storage.appdomain.cloud               | Cross Region EU 엔드포인트
         | s3.fra.eu.cloud-object-storage.appdomain.cloud           | Cross Region Frankfurt 엔드포인트
         | s3.mil.eu.cloud-object-storage.appdomain.cloud           | Cross Region Milan 엔드포인트
         | s3.ams.eu.cloud-object-storage.appdomain.cloud           | Cross Region Amsterdam 엔드포인트
         | s3.private.eu.cloud-object-storage.appdomain.cloud       | Cross Region Private 엔드포인트
         | s3.private.fra.eu.cloud-object-storage.appdomain.cloud   | Cross Region Frankfurt Private 엔드포인트
         | s3.private.mil.eu.cloud-object-storage.appdomain.cloud   | Cross Region Milan Private 엔드포인트
         | s3.private.ams.eu.cloud-object-storage.appdomain.cloud   | Cross Region Amsterdam Private 엔드포인트
         | s3.eu-gb.cloud-object-storage.appdomain.cloud            | Great Britain 엔드포인트
         | s3.private.eu-gb.cloud-object-storage.appdomain.cloud    | Great Britain Private 엔드포인트
         | s3.eu-de.cloud-object-storage.appdomain.cloud            | EU Region DE 엔드포인트
         | s3.private.eu-de.cloud-object-storage.appdomain.cloud    | EU Region DE Private 엔드포인트
         | s3.ap.cloud-object-storage.appdomain.cloud               | Cross Regional APAC 엔드포인트
         | s3.tok.ap.cloud-object-storage.appdomain.cloud           | Cross Regional Tokyo 엔드포인트
         | s3.hkg.ap.cloud-object-storage.appdomain.cloud           | Cross Regional HongKong 엔드포인트
         | s3.seo.ap.cloud-object-storage.appdomain.cloud           | Cross Regional Seoul 엔드포인트
         | s3.private.ap.cloud-object-storage.appdomain.cloud       | Cross Regional Private 엔드포인트
         | s3.private.tok.ap.cloud-object-storage.appdomain.cloud   | Cross Regional Tokyo Private 엔드포인트
         | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud   | Cross Regional HongKong Private 엔드포인트
         | s3.private.seo.ap.cloud-object-storage.appdomain.cloud   | Cross Regional Seoul Private 엔드포인트
         | s3.jp-tok.cloud-object-storage.appdomain.cloud           | Region Japan 엔드포인트
         | s3.private.jp-tok.cloud-object-storage.appdomain.cloud   | Region Japan Private 엔드포인트
         | s3.au-syd.cloud-object-storage.appdomain.cloud           | Region Australia 엔드포인트
         | s3.private.au-syd.cloud-object-storage.appdomain.cloud   | Region Australia Private 엔드포인트
         | s3.ams03.cloud-object-storage.appdomain.cloud            | Amsterdam Single Site 엔드포인트
         | s3.private.ams03.cloud-object-storage.appdomain.cloud    | Amsterdam Single Site Private 엔드포인트
         | s3.che01.cloud-object-storage.appdomain.cloud            | Chennai Single Site 엔드포인트
         | s3.private.che01.cloud-object-storage.appdomain.cloud    | Chennai Single Site Private 엔드포인트
         | s3.mel01.cloud-object-storage.appdomain.cloud            | Melbourne Single Site 엔드포인트
         | s3.private.mel01.cloud-object-storage.appdomain.cloud    | Melbourne Single Site Private 엔드포인트
         | s3.osl01.cloud-object-storage.appdomain.cloud            | Oslo Single Site 엔드포인트
         | s3.private.osl01.cloud-object-storage.appdomain.cloud    | Oslo Single Site Private 엔드포인트
         | s3.tor01.cloud-object-storage.appdomain.cloud            | Toronto Single Site 엔드포인트
         | s3.private.tor01.cloud-object-storage.appdomain.cloud    | Toronto Single Site Private 엔드포인트
         | s3.seo01.cloud-object-storage.appdomain.cloud            | Seoul Single Site 엔드포인트
         | s3.private.seo01.cloud-object-storage.appdomain.cloud    | Seoul Single Site Private 엔드포인트
         | s3.mon01.cloud-object-storage.appdomain.cloud            | Montreal Single Site 엔드포인트
         | s3.private.mon01.cloud-object-storage.appdomain.cloud    | Montreal Single Site Private 엔드포인트
         | s3.mex01.cloud-object-storage.appdomain.cloud            | Mexico Single Site 엔드포인트
         | s3.private.mex01.cloud-object-storage.appdomain.cloud    | Mexico Single Site Private 엔드포인트
         | s3.sjc04.cloud-object-storage.appdomain.cloud            | San Jose Single Site 엔드포인트
         | s3.private.sjc04.cloud-object-storage.appdomain.cloud    | San Jose Single Site Private 엔드포인트
         | s3.mil01.cloud-object-storage.appdomain.cloud            | Milan Single Site 엔드포인트
         | s3.private.mil01.cloud-object-storage.appdomain.cloud    | Milan Single Site Private 엔드포인트
         | s3.hkg02.cloud-object-storage.appdomain.cloud            | Hong Kong Single Site 엔드포인트
         | s3.private.hkg02.cloud-object-storage.appdomain.cloud    | Hong Kong Single Site Private 엔드포인트
         | s3.par01.cloud-object-storage.appdomain.cloud            | Paris Single Site 엔드포인트
         | s3.private.par01.cloud-object-storage.appdomain.cloud    | Paris Single Site Private 엔드포인트
         | s3.sng01.cloud-object-storage.appdomain.cloud            | Singapore Single Site 엔드포인트
         | s3.private.sng01.cloud-object-storage.appdomain.cloud    | Singapore Single Site Private 엔드포인트

   --location-constraint
      버킷을 생성할 때 지정한 엔드포인트와 일치해야 하는 위치 제약입니다.
      
      온프레미스 COS를 사용하는 경우 이 목록에서 선택하지 마세요. 그냥 Enter 키를 입력하세요.

      예제:
         | us-standard       | Cross Region US 표준
         | us-vault          | Cross Region US 보관소
         | us-cold           | Cross Region US 차가운 저장소
         | us-flex           | Cross Region US Flex
         | us-east-standard  | US East 표준
         | us-east-vault     | US East 보관소
         | us-east-cold      | US East 차가운 저장소
         | us-east-flex      | US East Flex
         | us-south-standard | US South 표준
         | us-south-vault    | US South 보관소
         | us-south-cold     | US South 차가운 저장소
         | us-south-flex     | US South Flex
         | eu-standard       | Cross Region EU 표준
         | eu-vault          | Cross Region EU 보관소
         | eu-cold           | Cross Region EU 차가운 저장소
         | eu-flex           | Cross Region EU Flex
         | eu-gb-standard    | Great Britain 표준
         | eu-gb-vault       | Great Britain 보관소
         | eu-gb-cold        | Great Britain 차가운 저장소
         | eu-gb-flex        | Great Britain Flex
         | ap-standard       | APAC 표준
         | ap-vault          | APAC 보관소
         | ap-cold           | APAC 차가운 저장소
         | ap-flex           | APAC Flex
         | mel01-standard    | Melbourne 표준
         | mel01-vault       | Melbourne 보관소
         | mel01-cold        | Melbourne 차가운 저장소
         | mel01-flex        | Melbourne Flex
         | tor01-standard    | Toronto 표준
         | tor01-vault       | Toronto 보관소
         | tor01-cold        | Toronto 차가운 저장소
         | tor01-flex        | Toronto Flex

   --acl
      버킷 및 객체 생성 및 복사시 사용되는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되지 않으면 버킷 생성에도 사용됩니다.
      
      자세한 정보는 [아마존 S3 개발자 가이드](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 S3에서 서버 측 복사 객체에 적용됩니다.
      S3는 소스에서 ACL을 복사하지 않고 새로운 ACL을 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 다른 사용자는 접근 권한이 없습니다 (기본값).
         |                    | 이 ACL은 IBM Cloud(Infra), IBM Cloud(Storage), On-Premise COS에서 사용할 수 있습니다.
         | public-read        | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 권한을 가집니다.
         |                    | 이 ACL은 IBM Cloud(Infra), IBM Cloud(Storage), On-Premise IBM COS에서 사용할 수 있습니다.
         | public-read-write  | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 가집니다.
         |                    | 이 ACL은 IBM Cloud(Infra), On-Premise IBM COS에서 사용할 수 있습니다.
         | authenticated-read | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 인증된 사용자 그룹은 읽기 권한을 가집니다.
         |                    | 버킷에는 지원되지 않습니다.
         |                    | 이 ACL은 IBM Cloud(Infra) 및 On-Premise IBM COS에서 사용할 수 있습니다.

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL입니다.
      
      자세한 정보는 [아마존 S3 개발자 가이드](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않으면 "acl"만 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 다른 사용자는 접근 권한이 없습니다 (기본값).
         | public-read        | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 권한을 가집니다.
         | public-read-write  | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 가집니다.
         |                    | 버킷에 대해서는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 인증된 사용자 그룹은 읽기 권한을 가집니다.

   --upload-cutoff
      청크 업로드로 전환하는 조각 크기입니다.
      
      이 크기보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일(예: "rclone rcat"에서 업로드되거나 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      참고: "--s3-upload- 동시 수"는 per transfer마다 이 크기의 청크가 메모리에 버퍼링되는 것을 의미합니다.
      
      고속 링크로 큰 파일을 전송하고 충분한 메모리가 있다면 이 값을 높이면 전송 속도가 향상됩니다.
      
      rclone은 10,000개의 청크 제한을 유지하기 위해 대용량 파일을 전송할 때 자동으로 청크 크기를 늘립니다.
      
      크기를 알 수 없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림으로 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 정보의 정확도가 낮아집니다. rclone은 청크가 AWS SDK에 의해 버퍼링되었을 때 청크를 전송된 것으로 처리하지만 실제로는 업로드 중일 수 있습니다.
      청크 크기가 클수록 AWS SDK 버퍼와 진행률 보고의 차이가 커집니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수입니다.
      
     이 옵션은 멀티파트 업로드 시 사용할 최대 복수 파트 수를 정의합니다.
      
     사용 가능한 서비스에서는 AWS S3 사양의 10,000개 청크를 지원하지 않을 수 있습니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드 할 때 청크 크기를 한도 내로 유지하기 위해 자동으로 청크 크기를 늘릴 것입니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 파일의 크기 제한입니다.
      
      서버 측에서 복사해야 하는 이 크기보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      개체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체 메타데이터에 추가합니다. 이는 데이터의 무결성 검사에 큰 도움이 됩니다. 하지만 큰 파일을 업로드하기 전에 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본값을 설정합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워 둘 경우 환경 변수 "AWS_PROFILE" 또는 환경 변수도 설정되어 있지 않으면 "default"로 기본값이 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 파일의 청크 수입니다.
      
      고속 링크로 대량의 대용량 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못한다면 크기를 늘리는 것이 전송 속도를 높이는 데 도움이 될 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false인 경우 가상 경로 스타일을 사용합니다. 자세한 내용은 [아마존 S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값을 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이 작업을 수행할 것입니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      이 값이 false(기본값)인 경우 rclone은 v4 인증을 사용하고 설정되어 있으면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 경우에만 사용하십시오. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기입니다(각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려진 것입니다.
      대부분의 서비스는 응답 목록을 1000개의 객체로 잘라내도록 설정되어 있습니다. AWS S3에서는 이 값이 전역 최대치이고 제한을 늘릴 수 없습니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요. Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 0(자동).
      
      S3가 처음 출시되었을 때는 버킷의 개체를 나열하기 위해 ListObjects 호출만 이용할 수 있었습니다.
      
      그러나 2016년 5월에는 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 좋은 성능을 가지며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출할 목록 객체 메서드를 추측합니다. 추측이 틀린 경우 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록에 URL 인코딩을 사용할지 여부: true/false/unset
      
      일부 공급자는 목록에 URL 인코딩을 지원하며 사용 가능한 경우 파일 이름에서 제어 문자를 사용할 때 이 방법이 더 안정적입니다. 이 값이 unset(기본값)인 경우 rclone은 공급자 설정에 따라 적용할 방법을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      설정된 경우 버킷을 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      또한 사용자가 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에는 이 버그로 인해 오류가 발생했으므로 무시되었습니다.
      

   --no-head
      설정된 경우 업로드한 개체의 무결성을 확인하기 위해 HEAD하지 않습니다.
      
      rclone은 최소한 PUT으로 객체를 업로드 한 후에 200 OK 메시지를 수신하면 정상적으로 업로드된 것으로 간주합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터, 수정 시간, 스토리지 클래스 및 콘텐츠 유형이 업로드된 것과 동일하다.
      - 크기는 업로드된 것과 동일하다.
      
      단일 부분 PUT의 응답으로 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      크기를 알 수 없는 소스 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 감지 확률이 올라가므로 정상적인 작업에는 권장되지 않습니다. 실제로 업로드 실패의 확률이 매우 낮습니다.
      

   --no-head-object
      설정된 경우 GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용할 것입니다.
      이 옵션은 미사용된 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3 백엔드(특히 minio)와 HTTP/2에 관한 해결되지 않은 문제가 있습니다. HTTP/2는 S3 백엔드에 대해 기본적으로 활성화되어 있지만이 플래그에서 비활성화할 수 있습니다. 문제가 해결될 때 이 플래그는 제거될 예정입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 다운로드된 데이터에 대해 더 저렴한 트래픽 비용을 제공합니다.

   --use-multipart-etag
      멀티파트 업로드의 확인을 위해 ETag를 사용할지 여부
      
      true, false 또는 unset으로 true, false를 사용할지 기본값을 사용할지 설정하세요.
      

   --use-presigned-request
      단일 부분 업로드에 대해 사전 서명 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false인 경우 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      rclone < 1.59 버전에서는 단일 부분 객체를 업로드하기 위해 서명된 요청을 사용하고이 플래그를 true로 설정하면이 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜, "2006-01-02", date의 "2006-01-02 15:04:05" 또는 해당 시간이 얼마나 오래되었는지를 나타내는 기간인 "100d" 또는 "1h"가 있어야 합니다.
      
      이 설정을 사용하는 경우 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 정보는 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      설정된 경우 gzip으로 압축된 객체를 복원합니다.
      
      S3에 "Content-Encoding: gzip"가 설정된 상태로 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 파일을 받으면서 "Content-Encoding: gzip"로 받은 압축 파일을 복원합니다. 이는 rclone이 크기와 해시를 확인하지 못하지만 파일 내용은 복원됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 설정값을 지정하십시오.
      
      일반적으로 공급자는 다운로드할 때 객체를 변경하지 않습니다. 객체가 `Content-Encoding: gzip`로 업로드되지 않았다면 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급자는 gzip으로 압축된 객체를 (Cloudflare 등) `Content-Encoding: gzip`로 업로드하지 않았더라도 gzip으로 압축할 수 있습니다.
      
      이러한 항목을 수신하는 경우
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      같은 오류가 발생할 수 있습니다.
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip이 설정된 개체를 chunked 전송 인코딩으로 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset로 설정하면 rclone은 공급자 설정에 따라 적용할 방법을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 및 객체 생성, 저장 또는 복사 시 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value             IBM COS S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   런타임(환경 변수 또는 EC2/ECS 메타 데이터)에서 AWS 자격증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  IBM Cloud Public을 사용하는 경우 엔드포인트와 일치해야 하는 위치 제약입니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전입니다. [$REGION]
   --secret-access-key value    AWS Secret Access Key(비밀번호)입니다. [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 크기 제한입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정된 경우 gzip으로 압축된 객체를 복원합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               개체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다(각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록에 URL 인코딩을 사용할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0(자동). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 설정값을 지정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정된 경우 버킷을 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        설정된 경우 HEAD한 후에 업로드한 개체의 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 설정된 경우 GET 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 조각 크기입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드의 확인을 위해 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 대해 사전 서명 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름(자동 생성)
   --path value  스토리지 경로

```
{% endcode %}