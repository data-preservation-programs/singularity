# IBM COS S3

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 ibmcos - IBM COS S3

사용법:
   singularity storage update s3 ibmcos [command options] <name|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터).
      
      access_key_id 및 secret_access_key이 비어 있을 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --region
      연결할 지역.
      
      S3 클론을 사용하는 경우 지역이 없으면 빈칸으로 둡니다.

      예시:
         | <unset>            | 확실하지 않을 때 사용합니다.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용합니다.
         |                    | 예: Jewel/v10 CEPH 이전.

   --endpoint
      IBM COS S3 API를 위한 엔드포인트.
      
      IBM COS On Premise를 사용하는 경우 지정합니다.

      예시:
         | s3.us.cloud-object-storage.appdomain.cloud               | 미국 크로스 리전 엔드포인트
         | s3.dal.us.cloud-object-storage.appdomain.cloud           | 미국 크로스 리전 Dallas 엔드포인트
         | s3.wdc.us.cloud-object-storage.appdomain.cloud           | 미국 크로스 리전 Washington DC 엔드포인트
         | s3.sjc.us.cloud-object-storage.appdomain.cloud           | 미국 크로스 리전 San Jose 엔드포인트
         | s3.private.us.cloud-object-storage.appdomain.cloud       | 미국 크로스 리전 개인 엔드포인트
         | s3.private.dal.us.cloud-object-storage.appdomain.cloud   | 미국 크로스 리전 Dallas 개인 엔드포인트
         | s3.private.wdc.us.cloud-object-storage.appdomain.cloud   | 미국 크로스 리전 Washington DC 개인 엔드포인트
         | s3.private.sjc.us.cloud-object-storage.appdomain.cloud   | 미국 크로스 리전 San Jose 개인 엔드포인트
         | s3.us-east.cloud-object-storage.appdomain.cloud          | 미국 동부 리전 엔드포인트
         | s3.private.us-east.cloud-object-storage.appdomain.cloud  | 미국 동부 리전 개인 엔드포인트
         | s3.us-south.cloud-object-storage.appdomain.cloud         | 미국 남부 리전 엔드포인트
         | s3.private.us-south.cloud-object-storage.appdomain.cloud | 미국 남부 리전 개인 엔드포인트
         | s3.eu.cloud-object-storage.appdomain.cloud               | 유럽 크로스 리전 엔드포인트
         | s3.fra.eu.cloud-object-storage.appdomain.cloud           | 유럽 크로스 리전 Frankfurt 엔드포인트
         | s3.mil.eu.cloud-object-storage.appdomain.cloud           | 유럽 크로스 리전 Milan 엔드포인트
         | s3.ams.eu.cloud-object-storage.appdomain.cloud           | 유럽 크로스 리전 Amsterdam 엔드포인트
         | s3.private.eu.cloud-object-storage.appdomain.cloud       | 유럽 크로스 리전 개인 엔드포인트
         | s3.private.fra.eu.cloud-object-storage.appdomain.cloud   | 유럽 크로스 리전 Frankfurt 개인 엔드포인트
         | s3.private.mil.eu.cloud-object-storage.appdomain.cloud   | 유럽 크로스 리전 Milan 개인 엔드포인트
         | s3.private.ams.eu.cloud-object-storage.appdomain.cloud   | 유럽 크로스 리전 Amsterdam 개인 엔드포인트
         | s3.eu-gb.cloud-object-storage.appdomain.cloud            | 영국 엔드포인트
         | s3.private.eu-gb.cloud-object-storage.appdomain.cloud    | 영국 개인 엔드포인트
         | s3.eu-de.cloud-object-storage.appdomain.cloud            | 유럽 DE 리전 엔드포인트
         | s3.private.eu-de.cloud-object-storage.appdomain.cloud    | 유럽 DE 리전 개인 엔드포인트
         | s3.ap.cloud-object-storage.appdomain.cloud               | APAC 크로스 리전 엔드포인트
         | s3.tok.ap.cloud-object-storage.appdomain.cloud           | APAC 크로스 리전 Tokyo 엔드포인트
         | s3.hkg.ap.cloud-object-storage.appdomain.cloud           | APAC 크로스 리전 HongKong 엔드포인트
         | s3.seo.ap.cloud-object-storage.appdomain.cloud           | APAC 크로스 리전 Seoul 엔드포인트
         | s3.private.ap.cloud-object-storage.appdomain.cloud       | APAC 크로스 리전 개인 엔드포인트
         | s3.private.tok.ap.cloud-object-storage.appdomain.cloud   | APAC 크로스 리전 Tokyo 개인 엔드포인트
         | s3.private.hkg.ap.cloud-object-storage.appdomain.cloud   | APAC 크로스 리전 HongKong 개인 엔드포인트
         | s3.private.seo.ap.cloud-object-storage.appdomain.cloud   | APAC 크로스 리전 Seoul 개인 엔드포인트
         | s3.jp-tok.cloud-object-storage.appdomain.cloud           | APAC 리전 일본 엔드포인트
         | s3.private.jp-tok.cloud-object-storage.appdomain.cloud   | APAC 리전 일본 개인 엔드포인트
         | s3.au-syd.cloud-object-storage.appdomain.cloud           | APAC 리전 호주 엔드포인트
         | s3.private.au-syd.cloud-object-storage.appdomain.cloud   | APAC 리전 호주 개인 엔드포인트
         | s3.ams03.cloud-object-storage.appdomain.cloud            | Amsterdam 단일 사이트 엔드포인트
         | s3.private.ams03.cloud-object-storage.appdomain.cloud    | Amsterdam 단일 사이트 개인 엔드포인트
         | s3.che01.cloud-object-storage.appdomain.cloud            | Chennai 단일 사이트 엔드포인트
         | s3.private.che01.cloud-object-storage.appdomain.cloud    | Chennai 단일 사이트 개인 엔드포인트
         | s3.mel01.cloud-object-storage.appdomain.cloud            | Melbourne 단일 사이트 엔드포인트
         | s3.private.mel01.cloud-object-storage.appdomain.cloud    | Melbourne 단일 사이트 개인 엔드포인트
         | s3.osl01.cloud-object-storage.appdomain.cloud            | Oslo 단일 사이트 엔드포인트
         | s3.private.osl01.cloud-object-storage.appdomain.cloud    | Oslo 단일 사이트 개인 엔드포인트
         | s3.tor01.cloud-object-storage.appdomain.cloud            | Toronto 단일 사이트 엔드포인트
         | s3.private.tor01.cloud-object-storage.appdomain.cloud    | Toronto 단일 사이트 개인 엔드포인트
         | s3.seo01.cloud-object-storage.appdomain.cloud            | Seoul 단일 사이트 엔드포인트
         | s3.private.seo01.cloud-object-storage.appdomain.cloud    | Seoul 단일 사이트 개인 엔드포인트
         | s3.mon01.cloud-object-storage.appdomain.cloud            | Montreal 단일 사이트 엔드포인트
         | s3.private.mon01.cloud-object-storage.appdomain.cloud    | Montreal 단일 사이트 개인 엔드포인트
         | s3.mex01.cloud-object-storage.appdomain.cloud            | Mexico 단일 사이트 엔드포인트
         | s3.private.mex01.cloud-object-storage.appdomain.cloud    | Mexico 단일 사이트 개인 엔드포인트
         | s3.sjc04.cloud-object-storage.appdomain.cloud            | San Jose 단일 사이트 엔드포인트
         | s3.private.sjc04.cloud-object-storage.appdomain.cloud    | San Jose 단일 사이트 개인 엔드포인트
         | s3.mil01.cloud-object-storage.appdomain.cloud            | Milan 단일 사이트 엔드포인트
         | s3.private.mil01.cloud-object-storage.appdomain.cloud    | Milan 단일 사이트 개인 엔드포인트
         | s3.hkg02.cloud-object-storage.appdomain.cloud            | Hong Kong 단일 사이트 엔드포인트
         | s3.private.hkg02.cloud-object-storage.appdomain.cloud    | Hong Kong 단일 사이트 개인 엔드포인트
         | s3.par01.cloud-object-storage.appdomain.cloud            | Paris 단일 사이트 엔드포인트
         | s3.private.par01.cloud-object-storage.appdomain.cloud    | Paris 단일 사이트 개인 엔드포인트
         | s3.sng01.cloud-object-storage.appdomain.cloud            | Singapore 단일 사이트 엔드포인트
         | s3.private.sng01.cloud-object-storage.appdomain.cloud    | Singapore 단일 사이트 개인 엔드포인트

   --location-constraint
      엔드포인트와 일치해야하는 위치 제약조건 (IBM Cloud Public 사용시).
      
      온프레미스 COS의 경우 이 목록에서 선택하지 마십시오.

      예시:
         | us-standard       | 미국 크로스 리전 표준
         | us-vault          | 미국 크로스 리전 보안금고
         | us-cold           | 미국 크로스 리전 차가운 저장
         | us-flex           | 미국 크로스 리전 플렉스
         | us-east-standard  | 미국 동부 리전 표준
         | us-east-vault     | 미국 동부 리전 보안금고
         | us-east-cold      | 미국 동부 리전 차가운 저장
         | us-east-flex      | 미국 동부 리전 플렉스
         | us-south-standard | 미국 남부 리전 표준
         | us-south-vault    | 미국 남부 리전 보안금고
         | us-south-cold     | 미국 남부 리전 차가운 저장
         | us-south-flex     | 미국 남부 리전 플렉스
         | eu-standard       | 유럽 크로스 리전 표준
         | eu-vault          | 유럽 크로스 리전 보안금고
         | eu-cold           | 유럽 크로스 리전 차가운 저장
         | eu-flex           | 유럽 크로스 리전 플렉스
         | eu-gb-standard    | 영국 표준
         | eu-gb-vault       | 영국 보안금고
         | eu-gb-cold        | 영국 차가운 저장
         | eu-gb-flex        | 영국 플렉스
         | ap-standard       | APAC 표준
         | ap-vault          | APAC 보안금고
         | ap-cold           | APAC 차가운 저장
         | ap-flex           | APAC 플렉스
         | mel01-standard    | Melbourne 표준
         | mel01-vault       | Melbourne 보안금고
         | mel01-cold        | Melbourne 차가운 저장
         | mel01-flex        | Melbourne 플렉스
         | tor01-standard    | Toronto 표준
         | tor01-vault       | Toronto 보안금고
         | tor01-cold        | Toronto 차가운 저장
         | tor01-flex        | Toronto 플렉스

   --acl
      버킷 생성 및 객체 저장 또는 복사시 사용할 Canned ACL.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않으면 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하십시오.
      
      S3는 소스의 ACL을 복사하지 않고 새로운 ACL을 작성하기 때문에 이 ACL은 S3에서 객체를 서버 간 복사할 때 적용됩니다.
      
      ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 모든 사용자에게는 액세스 권한이 없습니다 (기본값입니다).
         |                    | 이 ACL은 IBM Cloud (Infra), IBM Cloud (Storage), 온프레미스 COS에서 사용할 수 있습니다.
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹은 READ 액세스 권한을 얻습니다.
         |                    | 이 ACL은 IBM Cloud (Infra), IBM Cloud (Storage), 온프레미스 IBM COS에서 사용할 수 있습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹은 READ 및 WRITE 액세스 권한을 얻습니다.
         |                    | 이 ACL은 IBM Cloud (Infra), 온프레미스 IBM COS에서 사용할 수 있습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹은 READ 액세스 권한을 얻습니다.
         |                    | 버킷에서 지원되지 않습니다.
         |                    | 이 ACL은 IBM Cloud (Infra) 및 온프레미스 IBM COS에서 사용할 수 있습니다.

   --bucket-acl
      버킷 생성 시 사용할 Canned ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      이 ACL은 버킷을 생성할 때만 적용되며 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 모두 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 모든 사용자에게 액세스 권한이 없습니다 (기본값입니다).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹은 READ 액세스 권한을 얻습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹은 READ 및 WRITE 액세스 권한을 얻습니다.
         |                    | 버킷에서는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹은 READ 액세스 권한을 얻습니다.

   --upload-cutoff
      청크드 업로드로 전환할 임계값.
      
      이보다 큰 파일은 청크 크기로 업로드됩니다. 최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 크거나 파일 크기를 모르는 파일 ("rclone rcat"에서 가져온 파일이거나 "rclone mount" 또는 구글 포토 또는 구글 문서로 업로드된 파일)을 업로드하는 경우 청크 크기를 사용하여 멀티파트 업로드를 사용하여 업로드됩니다.
      
      참고로, "--s3-upload-concurrency" 대용량 파일의 청크 수만큼의 이 크기의 버퍼가 전송마다 메모리에 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하고 충분한 메모리가 있다면 이 값을 증가시켜 전송 속도를 높일 수 있습니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 최대 10,000개의 청크 제한을 유지합니다.
      
      알려진 크기를 가지지 않는 파일의 경우 설정된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 낮아집니다. rclone은 AWS SDK에서 버퍼링될 때 청크를 보낼 때-보낸 것으로 처리하지만 아직 업로드 중일 수 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 보고 사이의 진행률 통계의 차이를 의미합니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 청크 수.
      
      이 옵션은 멀티파트 업로드시 사용할 청크 수의 최댓값을 정의합니다.
      
      이는 서비스가 AWS S3의 10,000개 청크 사양을 지원하지 않을 경우 유용할 수 있습니다.
      
      큰 크기의 알려진 파일을 업로드할 때 rclone은 수의 청크 한도를 유지하기 위해 청크 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환할 임계값.
      
      이보다 큰 파일을 서버 측으로 복사해야 한다면 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만 큰 파일을 업로드하기 시작하는 데 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면 현재 사용자의 홈 디렉토리를 기본값으로 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로파일을 제어합니다.
      
      값이 비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default" 인 경우에는 환경 변수 값을 기본값으로 사용합니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      높은 속도의 링크에서 대용량 파일을 업로드하므로 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true로 설정하면 경로 스타일 액세스를 사용하고, false로 설정하면 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고, false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값이 false로 설정되어야 합니다. rclone은 제공자 설정에 따라 자동으로 설정합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false(기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청한 객체가 1000개 이상인 경우에도 응답 목록을 잘라냅니다.
      AWS S3에서는 글로벌 최대값이므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동 선택을 위한 0.
      
      S3가 처음 출시될 때는 ListObjects 호출만 제공해서 버킷의 객체를 열람할 수 있었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 높은 성능을 제공하며 가능하면 사용해야 합니다.
      
      기본값으로 설정된 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못된 추측을 하는 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 목록을 URL 인코딩하는 것을 지원하며 가능한 경우 파일 이름에 제어 문자를 사용할 때 이 기능이 더 신뢰할 수 있습니다. unset(기본값)로 설정된 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      버킷 생성 권한이 없는 경우 이것이 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이전에 정상적으로 전달되었습니다.
      

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD 요청을 수행하지 않습니다.
      
      rclone이 PUT로 객체를 업로드한 후 200 OK 메시지를 수신하면 제대로 업로드되었다고 가정합니다. 

      특히 다음에 해당하는 것으로 가정합니다.
      
      - 메타데이터, 변경 시간, 저장 유형 및 콘텐츠 유형이 업로드한 것과 같음
      - 크기가 업로드한 것과 같음
      
      다음을 읽어 단일 부분 PUT 응답에서 다음 항목을 읽어옵니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      알려지지 않은 길이의 소스 객체가 업로드된 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 가능성이 증가하며 특히 올바르지 않은 크기의 경우이므로 일반적인 작동에는 권장되지 않습니다. 실제로 이 플래그를 설정하더라도 업로드 실패 가능성은 매우 적습니다.
      

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀의 몇 번의 메모리 버퍼를 플러시할지 정의합니다.
      
      추가 버퍼(예: 멀티파트가 필요한 업로드)는 주어진 메모리 버퍼 풀을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않는 버퍼를 얼마나 자주 풀에서 제거할지를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. s3 백엔드의 경우 HTTP/2가 기본적으로 활성화되지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면 이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트입니다.
      일반적으로 AWS S3는 표준 egress: CloudFront 네트워크를 통해 데이터를 다운로드할 때 더 저렴한 egress를 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값을 true, false 또는 unset(기본값)으로 설정할 수 있습니다.
      

   --use-presigned-request
      단일 파트 업로드에서 사전 서명된 요청 또는 PutObject을 사용할지 여부
      
      false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone < 1.59 버전에서는 단일 부분 객체를 업로드하기 위해 사전 서명된 요청을 사용하고, 이 플래그를 true로 설정하면 이 기능이 다시 활성화됩니다. 이것은 특수한 경우 또는 테스트에만 필요합니다.


   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 버전이 어떻게 보였는지 파일을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", datetime "2006-01-02 15:04:05" 또는 그 동안의 기간을 나타내는 '100d' 또는 '1h'과 같이 길었던 경우에 사용합니다.
      
      이를 사용할 때 파일 쓰기 작업을 수행할 수 없으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [time 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"으로 설정된 객체를 S3에 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"으로 받을 때 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 압축할 수 있다면 이것을 설정하세요.
      
      일반적으로 제공자는 객체가 다운로드될 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 경우 다운로드되지 않습니다.
      
      그러나 일부 제공자는 객체를 "Content-Encoding: gzip"로 압축할 수도 있습니다(예: Cloudflare).
      
      이러한 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip를 설정하고 청크 전송 인코딩으로 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset(기본값)로 설정된 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하게 됩니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사시 사용할 Canned ACL. [$ACL]
   --endpoint value             IBM COS S3 API를 위한 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  IBM Cloud Public 사용시 엔드포인트와 일치해야하는 위치 제약조건. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용할 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환할 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동 선택을 위한 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀의 몇 번의 메모리 버퍼를 씻을지 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 압축할 수 있다면 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD 요청을 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로파일. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크드 업로드로 전환할 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부. (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 인증을 사용할지 여부. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정한 시간에 어떻게 보여줄지. (기본값: "off") [$VERSION_AT]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --might-gzip value               백엔드가 객체를 압축할 수 있다면 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]

```
{% endcode %}