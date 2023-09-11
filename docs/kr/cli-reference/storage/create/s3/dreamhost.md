# Dreamhost DreamObjects

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 dreamhost - Dreamhost DreamObjects

사용법:
   singularity storage create s3 dreamhost [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격증명 (환경 변수나 env vars이 없으면 EC2/ECS 메타데이터) 가져오기.
      
      access_key_id 및 secret_access_key가 공백인 경우에만 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격증명 가져오기.

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격증명의 경우 공백으로 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호).
      
      익명 액세스 또는 런타임 자격증명의 경우 공백으로 둡니다.

   --region
      연결할 지역.
      
      S3 클론을 사용하는 경우 지역 없음.

      예:
         | <unset>            | 확실하지 않은 경우 이 값을 사용하십시오.
         |                    | v4 시그니처와 빈 지역을 사용합니다.
         | other-v2-signature | v4 시그니처가 작동하지 않는 경우에만 사용하십시오.
         |                    | 예: Jewel/v10 CEPH 이전.

   --endpoint
      S3 API에 대한 엔드포인트.
      
      S3 클론을 사용하는 경우 필수입니다.

      예:
         | objects-us-east-1.dream.io | Dream Objects 엔드포인트

   --location-constraint
      위치 제약 - 지역과 일치해야 함.
      
      확실하지 않은 경우 공백으로 둡니다. 버킷 생성 시에만 사용됩니다.

   --acl
      버킷 생성 및 객체 저장 또는 복사시 사용되는 Canned ACL.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3 서버 측 복사 객체할 때 이 ACL이 적용됨에 유념하세요. S3는 소스의 ACL을 복사하는 것이 아니라 새로 작성할 뿐입니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본(개인)이 사용됩니다.
      

      예:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자에게 액세스 권한 없음(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 액세스 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스 권한 부여.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스 권한 부여.

   --upload-cutoff
      청크 업로드로 전환하는 임계값.
      
      이 값보다 큰 파일은 chunk_size 크기의 청크로 업로드됩니다.
      최소값은 0이며 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이거나 크기가 알려지지 않은 파일("rclone rcat" 또는 "rclone mount" 또는 google 사진 또는 google 문서에서 업로드된 파일 등)은이 청크 크기를 사용하여 다중파트 업로드로 업로드됩니다.
      
      또한 "--s3-upload-concurrency" 크기의 청크는 전송당 메모리에 버퍼링됩니다.
      
      대역폭을 충분히 사용하는 고속 링크로 대량의 파일을 전송하고 충분한 메모리가 있는 경우 이 크기를 늘리면 전송 속도가 높아집니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드 할 때 청크 크기를 자동으로 늘려 10,000 청크 라번 이하로 유지합니다.
      
      알려지지 않은 크기의 파일은 설정된 chunk_size로 업로드됩니다. 기본 청크 크기가 5 MiB이고 최대 10,000 청크가 있을 수 있기 때문에 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상태의 정확도가 떨어집니다. Rclone은 실제로 업로드 중인 청크를 AWS SDK에서 버퍼링 할 때 청크를 보낸 것으로 처리하지만 실제로 업로드 중인 것일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 통계에 대해 진행 보고 과도하게 벗어나게합니다.

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      10,000 청크의 AWS S3 사양을 지원하지 않는 서비스의 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드 할 때 청크 크기를 자동으로 늘려 10,000 청크 라번 이하로 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값.
      
      이 임계값보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이며 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬 저장하지 않기.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이렇게하면 큰 파일을 업로드하기 시작하는 데 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 빈 문자열이면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본적으로 현재 사용자의 홈 디렉토리를 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필.
      
      env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용되는 프로필을 제어합니다.
      
      비어 있으면 기본적으로 "AWS_PROFILE" 또는
      환경 변수가 설정되지 않은 경우 "default"로 사용됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드에 대한 병렬화.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      대역폭을 최대한 활용하지 못하고 High-speed링크를 통해 대량의 대용량 파일을 업로드하고 있는 경우 이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true 인 경우 경로 스타일 액세스를 사용하고, false 인 경우 가상 호스팅 스타일을 사용합니다.
      
      true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고,
      false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공자(예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이 값을 false로 설정해야 합니다. rclone은 제공자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.
      
      이 값이 false(기본값)로 설정되면 rclone은 v4 인증을 사용합니다. 설정하면 rclone은 v2 인증을 사용합니다.
      
      v4 시그니처가 작동하지 않는 경우에만 사용하십시오. 예 : Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items"또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 1,000개 이상을 요청하면 응답 목록을 1,000개로 줄입니다.
      AWS S3에서 이것은 전역 최대이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값이 증가할 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전 : 1,2 또는 자동 시용인 경우 0.
      
      S3가 처음 시작되었을 때 버킷 내 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 높은 성능이 있으며 가능한 경우 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출할 목록 개체 메소드를 추측합니다. 추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록 인코딩의 사용 여부: true/false/unset
      
      일부 제공자는 파일 이름에 제어 문자를 사용할 때 이러한 사항을 적용할 수 있으므로 목록을 URL로 인코딩하는 자체를 지원합니다. unset(기본값)로 설정된 경우 rclone은 할당 할 아이템이 있는지를 고려하여 공급자 설정에 따라 선택합니 출력을 선택할 수 있습니다.
      

   --no-check-bucket
      버킷을 확인하거나 생성하지 않도록 설정합니다.
      
      알려진 버킷이 이미 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 갖지 못한 경우에도 필요할 수 있습니다. v1.52.0 이전 버전에서는 버그 때문에 정상적으로 전달되었습니다.
      

   --no-head
      청크 완전성을 확인하기 위해 업로드된 개체 HEAD 하지 않습니다.
      
      rclone은 200 OK 메시지를 수신할 경우 PUT으로 객체를 업로드 한 후에 업로드가 올바르게 된 것으로 가정합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터(모드 타임, 스토리지 클래스 및 콘텐츠 유형)이 업로드와 동일했던 것.
      - 크기가 업로드됨과 동일했던 것
      
      아래에서 몇 가지 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      여러 부분 PUT의 경우 이러한 항목은 읽지 않습니다.
      
      길이를 알 수없는 소스 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패가 감지되지 않을 가능성이 높아집니다.
      특히 올바르지 않은 크기의 경우이므로 정상적인 운영에는 권장되지 않습니다. 실제로 업로드 실패가 감지되지 않을 가능성은 매우 작습니다.
      

   --no-head-object
      개체를 가져오기 전에 HEAD를 GET 앞에 두지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기.
      
      청크가 필요한 업로드(예: 멀티파트)는 할당을위한 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 시기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2 도 사용할 수있는 이슈가 있습니다. s3 백엔드는 HTTP/2가 기본적으로 활성화되어 있지만 여기에서 비활성화 할 수 있습니다. 이 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참고: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트.
      주로 AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터의 egress를 더 저렴하게 제공합니다.

   --use-multipart-etag
      멀티파트 업로드에서 확인용으로 ETag를 사용할지 여부
      
      true, false 또는 기본값(해당 공급자에 대한)을 사용하도록 설정해야합니다.
      

   --use-presigned-request
      단일 부분 업로드에 대해 사전 서명 된 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false로 설정되어 있으면 rclone은 객체를 업로드하기 위해 PutObject를 사용합니다.
      
      rclone < 1.59 버전은 단일 부분 객체를 업로드하기위한 사전 서명 된 요청을 사용하고 이 플래그 값을 true로 설정하면 이 기능이 다시 활성화됩니다. 예외적인 경우 또는 테스트를 위해서만 필요합니다.
      

   --versions
      디렉터리 목록에 이전 버전 포함.

   --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개 변수는 날짜, "2006-01-02", 날짜시간 "2006-01-02
      15:04:05" 또는 그시간전의 기간, 예를 들어 "100d" 또는 "1h"일 수 있습니다.
      
     이를 사용하면 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 압축 된 개체를 해제하려면 이 값을 설정하세요.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드 할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축 된 개체로 다운로드하지만이 플래그로 설정하면 수신된 대로 이러한 파일을 "Content-Encoding: gzip"로 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip할 수 있으므로이 값을 설정합니다.
      
      일반적으로 제공자는 개체를 다운로드할 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 경우 다운로드하지 않습니다.
      
      그러나 몇몇 제공자는 gzip을 지원하지 않는 개체도 gzip으로 압축 할 수 있습니다(Cloudflare 등).
      
      이렇게하면
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      와 같은 오류가 발생할 수 있습니다.
      
      이 플래그를 설정하면 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩으로 개체를 다운로드하는 경우 rclone은 개체를 실시간으로 압축 해제합니다.
      
      비활성화 수 있습니다(기본값으로), rclone은 공급자 설정에 따라 적용할 사항을 선택합니다. 그러나 여기에서 rclone의 선택사항을 재정의 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사시 사용되는 Canned ACL. [$ACL]
   --endpoint value             S3 API에 대한 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격증명 (환경 변수나 env vars이 없으면 EC2/ECS 메타데이터) 가져오기. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 - 지역과 일치해야 함. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS Secret Access Key(비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 압축 된 개체를 해제하려면 이 값을 설정하세요. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬 저장하지 않기. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 경로 스타일 액세스를 사용하고, false 인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록 인코딩의 사용 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전 : 1,2 또는 자동 시용인 경우 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip 할 수 있으므로이 값을 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷을 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        청크 완전성을 확인하기 위해 업로드된 개체 HEAD 하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져오기 전에 HEAD를 GET 앞에 두지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 병렬화. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 확인용으로 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 대해 사전 서명 된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전 포함. (기본값: false) [$VERSIONS]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}