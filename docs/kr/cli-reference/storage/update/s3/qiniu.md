# Qiniu Object Storage (Kodo)

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 qiniu - Qiniu Object Storage (Kodo)

사용법:
   singularity storage update s3 qiniu [명령 옵션] <이름|ID>

설명:
   --env-auth
      런타임(환경 변수 또는 환경으로부터 EC2/ECS 메타데이터)으로부터 AWS 자격 증명을 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어있을 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명 입력
         | true  | 환경(환경 변수 또는 IAM)으로부터 AWS 자격 증명 가져오기

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --region
      연결할 리전.

      예제:
         | cn-east-1      | 기본 엔드포인트 - 확실하지 않은 경우 좋은 선택입니다.
         |                | 동중국 리전 1.
         |                | 위치 제약 조건 cn-east-1이 필요합니다.
         | cn-east-2      | 동중국 리전 2.
         |                | 위치 제약 조건 cn-east-2이 필요합니다.
         | cn-north-1     | 북중국 리전 1.
         |                | 위치 제약 조건 cn-north-1이 필요합니다.
         | cn-south-1     | 남중국 리전 1.
         |                | 위치 제약 조건 cn-south-1이 필요합니다.
         | us-north-1     | 북미 리전.
         |                | 위치 제약 조건 us-north-1이 필요합니다.
         | ap-southeast-1 | 동남아시아 리전 1.
         |                | 위치 제약 조건 ap-southeast-1이 필요합니다.
         | ap-northeast-1 | 동북아시아 리전 1.
         |                | 위치 제약 조건 ap-northeast-1이 필요합니다.

   --endpoint
      Qiniu Object Storage를 위한 엔드포인트.

      예제:
         | s3-cn-east-1.qiniucs.com      | 동중국 엔드포인트 1
         | s3-cn-east-2.qiniucs.com      | 동중국 엔드포인트 2
         | s3-cn-north-1.qiniucs.com     | 북중국 엔드포인트 1
         | s3-cn-south-1.qiniucs.com     | 남중국 엔드포인트 1
         | s3-us-north-1.qiniucs.com     | 북미 엔드포인트 1
         | s3-ap-southeast-1.qiniucs.com | 동남아시아 엔드포인트 1
         | s3-ap-northeast-1.qiniucs.com | 동북아시아 엔드포인트 1

   --location-constraint
      리전과 일치하는 위치 제약 조건을 설정합니다.
      
      버킷 생성 시에만 사용됩니다.

      예제:
         | cn-east-1      | 동중국 리전 1
         | cn-east-2      | 동중국 리전 2
         | cn-north-1     | 북중국 리전 1
         | cn-south-1     | 남중국 리전 1
         | us-north-1     | 북미 리전 1
         | ap-southeast-1 | 동남아시아 리전 1
         | ap-northeast-1 | 동북아시아 리전 1

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL.
      
      이 ACL은 객체 생성 시에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      주의: S3에서 서버 측 복사 객체 시 이 ACL이 적용됩니다.
      소스로부터 ACL을 복사하는 것이 아니라 새로운 ACL을 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않은 경우 "acl"이 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자에게 액세스 권한이 없음(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 액세스 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스 권한 부여.
         |                    | 버킷에 대한 권한 부여는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스 권한 부여.

   --storage-class
      Qiniu에 새로운 객체를 저장할 때 사용할 저장 클래스.

      예제:
         | STANDARD     | 표준 저장 클래스
         | LINE         | 편근 액세스 저장 모드
         | GLACIER      | 아카이브 저장 모드
         | DEEP_ARCHIVE | 딥 아카이브 저장 모드

   --upload-cutoff
      청크 업로드로 전환할 파일의 임계값입니다.
      
      이 값보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      이 값은 0 이상 5 GiB 이하입니다.

   --chunk-size
      업로드할 때 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일(예: "rclone rcat"으로 생성되었거나 "rclone mount" 또는 Google
      사진 또는 Google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트로 업로드됩니다.
      
      이 값으로 인해 하나의 전송 당 메모리에 "--s3-upload-concurrency" 개의 청크 크기가 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값은 전송 속도를 높일 수 있습니다.
      
      rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 초과하지 않도록 합니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크를 가질 수 있으므로,
      기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      chunk 크기를 늘리면 "-P" 플래그로 표시되는 진행 상황 통계의 정확도가 낮아집니다. rclone은 chunk가 AWS SDK에 의해 버퍼링될 때
      chunk를 전송한 것으로 처리하며, 사실은 여전히 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 보고의 더 큰 차이를 의미합니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용하는 최대 청크 수입니다.
      
      이 옵션은 멀티파트 업로드 시 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이 옵션은 10,000개의 청크로 AWS S3 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 이 청크 수 제한에 미치지 않도록 합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환할 파일의 임계값입니다.
      
      서버 사이드로 복사해야 하는 이 임계값보다 큰 파일은 이 크기로 청크 단위로 복사됩니다.
      
      이 값은 0 이상 5 GiB 이하입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만
      대용량 파일의 업로드에는 시작하기 전에 오랜 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본값으로 현재 사용자의 홈 디렉터리를 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 그것도 설정되어 있지 않은 경우 "default" 환경 변수로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      대용량 파일을 고속링크를 통해 업로드하는 경우에 활용도가 낮아 대역폭을 완전히 활용하지 못하는 경우 이 값을 늘리면 전송 속도를 향상시킬 수 있습니다.

   --force-path-style
      true이면 패스 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.
      
      true(기본값)인 경우 rclone은 패스 스타일 액세스를 사용하고 false이면 rclone은 가상 패스 스타일을 사용합니다. 자세한 내용은
      [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS, Tencent COS)는 이 값을 false로 설정해야 합니다.
      rclone은 이를 기반으로 자동으로 설정합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정되어 있으면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 v2 인증을 사용하세요. 예를 들어 pre Jewel/v10 CEPH입니다.

   --list-chunk
      목록 청크의 크기(S3 요청의 각 ListObject에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청보다 더 많은 수의 응답 목록을 요구해도 응답 목록을 1000개로 잘라냅니다.
      AWS S3에서는 이것이 전역 최대이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전: 1,2 또는 자동으로 0.
      
      S3가 처음 시작되었을 때는 버킷에 있는 객체를 열거하기 위해 ListObjects 호출만을 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 호출은 매우 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정된 경우 rclone은 공급자가 설정한 목록 객체 방법에 따라 추측합니다. 잘못된 추측을한 경우 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 목록을 URL 인코딩하고 이를 사용할 수 있는 경우 파일 이름에 제어 문자를 사용할 때 신뢰할 수 있습니다. 이 값이 "unset"으로 설정된 경우 (기본값) rclone은
      프로바이더 설정에 따라 적용할 사항을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      설정된 경우 버킷의 존재 여부를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      또한 사용자가 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에는 이러한 오류가 무시되었습니다.

   --no-head
      업로드한 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone은 수신된 PUT 후에 200 OK 메시지를 받으면 제대로 업로드되었다고 가정합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)이 업로드한 것과 같음
      - 크기가 업로드한 것과 같음
      
      단일 파트 PUT의 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      크기를 알 수 없는 소스 객체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 증가합니다.
      특히 잘못된 크기와 같은 업로드 실패의 가능성이 매우 적습니다.
      

   --no-head-object
      GET을 수행하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시될 때까지의 시간간격입니다.
      
      메모리를 추가로 필요로 하는 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3(특히 미니오) 배경과 HTTP/2에 대한 미해결된 문제가 있습니다. 
      AWS S3는 s3 백엔드에 대해 기본적으로 HTTP/2를 사용하지만 여기서 비활성화할 수 있습니다. 
      문제가 해결되면이 플래그는 제거될 예정입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      AWS S3는 CloudFront 네트워크를 통해 다운로드되는 데이터에 대해 더 저렴한 전송 출력을 제공하는 경우가 많습니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값은 true, false 또는 기본값을 사용하려면 빈 상태로 두어야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false이면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone 1.59 이전 버전은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하고 이 플래그를 true로 설정하면 이 기능을 다시 활성화할 수 있습니다.
      예외적인 상황이나 테스트에서만 필요할 것입니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 날짜 및 시간("2006-01-02 15:04:05") 또는 그로부터 그렇게 오랜 시간 전("100d" 또는 "1h")일 수 있습니다.
      
      이 옵션을 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      이 값을 설정하면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"을 설정하여 객체를 S3로 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축되어 있는
      객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 이러한 파일을 수신시 "Content-Encoding: gzip"로 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만
      파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 값을 설정합니다.
      
      일반적으로 공급자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드할 때
      설정되지 않습니다.
      
      그러나 몇몇 공급자는 객체를 "Content-Encoding: gzip"로 압축할 수도 있습니다(예: Cloudflare).
      
      이런 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 "Content-Encoding: gzip"로 설정된 객체와 청크 전송 인코딩으로 객체를 다운로드하면 rclone은
      객체를 실시간으로 압축 해제합니다.
      
      이 값을 설정하지 않으면(기본값) rclone은 프로바이더 설정에 따라 적용할 사항을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL. [$ACL]
   --endpoint value             Qiniu Object Storage를 위한 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임(환경 변수 또는 환경으로부터 EC2/ECS 메타데이터)으로부터 AWS 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  리전과 일치하는 위치 제약 조건을 설정합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전. [$REGION]
   --secret-access-key value    AWS Secret Access Key(비밀번호). [$SECRET_ACCESS_KEY]
   --storage-class value        Qiniu에 새로운 객체를 저장할 때 사용할 저장 클래스. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드할 때 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환할 파일의 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     이 값을 설정하면 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 패스 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(S3 요청의 각 ListObject에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1,2 또는 자동으로 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용하는 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시될 때까지의 시간간격. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 값을 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정된 경우 버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환할 파일의 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]


```
{% endcode %}