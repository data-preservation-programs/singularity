# Qiniu Object Storage (Kodo)

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 qiniu - Qiniu Object Storage (Kodo)

사용법:
   singularity storage create s3 qiniu [command options] [arguments...]

설명:
   --env-auth
      런타임(환경 변수 또는 env vars 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명 가져오기.
      
      access_key_id와 secret_access_key이 비어있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 런타임에서 AWS 자격 증명을 환경으로부터 가져옵니다. (env vars 또는 IAM)

   --access-key-id
      AWS 액세스 키 ID. (anonymous access 또는 런타임 자격 증명을 위해 비워둘 수 있음.)

   --secret-access-key
      AWS 비밀 액세스 키(암호).
      
      anonymous access 또는 런타임 자격 증명을 위해 비워둘 수 있음.

   --region
      연결할 지역.

      예시:
         | cn-east-1      | 기본 엔드포인트(알 수 없다면 이것을 선택.)
         |                | 동 중국 1 지역.
         |                | 위치 제약 cn-east-1이 필요함.
         | cn-east-2      | 동 중국 2 지역.
         |                | 위치 제약 cn-east-2가 필요함.
         | cn-north-1     | 북 중국 1 지역.
         |                | 위치 제약 cn-north-1이 필요함.
         | cn-south-1     | 남 중국 1 지역.
         |                | 위치 제약 cn-south-1이 필요함.
         | us-north-1     | 북 아메리카 지역.
         |                | 위치 제약 us-north-1이 필요함.
         | ap-southeast-1 | 남동 아시아 지역 1.
         |                | 위치 제약 ap-southeast-1이 필요함.
         | ap-northeast-1 | 북동 아시아 지역 1.
         |                | 위치 제약 ap-northeast-1이 필요함.

   --endpoint
      Qiniu Object Storage를 위한 엔드포인트.

      예시:
         | s3-cn-east-1.qiniucs.com      | 동 중국 엔드포인트 1
         | s3-cn-east-2.qiniucs.com      | 동 중국 엔드포인트 2
         | s3-cn-north-1.qiniucs.com     | 북 중국 엔드포인트 1
         | s3-cn-south-1.qiniucs.com     | 남 중국 엔드포인트 1
         | s3-us-north-1.qiniucs.com     | 북 아메리카 엔드포인트 1
         | s3-ap-southeast-1.qiniucs.com | 남동 아시아 엔드포인트 1
         | s3-ap-northeast-1.qiniucs.com | 북동 아시아 엔드포인트 1

   --location-constraint
      지리적 제약 - 지역과 일치해야 합니다.
      
      버킷 생성 시에만 사용됩니다.

      예시:
         | cn-east-1      | 동 중국 1 지역
         | cn-east-2      | 동 중국 2 지역
         | cn-north-1     | 북 중국 1 지역
         | cn-south-1     | 남 중국 1 지역
         | us-north-1     | 북 아메리카 지역 1
         | ap-southeast-1 | 남동 아시아 지역 1
         | ap-northeast-1 | 북동 아시아 지역 1

   --acl
      버킷 생성 및 객체 저장 또는 복사 시에 사용되는 Canned ACL.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl 설정이 없는 경우, 버킷 생성에도 사용됩니다.
      
      자세한 내용은 [Amazon S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      S3는 소스에서 ACL을 복사하는 대신 새로운 ACL을 작성합니다.
      
      acl이 비어 있는 문자열인 경우에는 X-Amz-Acl: 헤더가 추가되지 않고
      기본값(private)이 사용됩니다.

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 내용은 [Amazon S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되어 있지 않은 경우
      "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 모두 비어 있는 문자열인 경우 X-Amz-Acl:
      헤더가 추가되지 않고 기본값(private)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | 다른 사용자에게는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | AllUsers 그룹에게 읽기 권한이 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 권한이 부여됩니다.
         |                    | 버킷에 대해서는 특별히 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | AuthenticatedUsers 그룹에게 읽기 권한이 부여됩니다.

   --storage-class
      Qiniu에서 새로운 객체를 저장할 때 사용할 저장 클래스.

      예시:
         | STANDARD     | 표준 저장 클래스
         | LINE         | 거의 사용되지 않는 접근 저장 모드
         | GLACIER      | 보관 저장 모드
         | DEEP_ARCHIVE | 심층 보관 저장 모드

   --upload-cutoff
      청크 업로드로 전환하는 업로드 크기 기준.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 크거나 알 수 없는 크기의 파일(예: "rclone rcat"으로
      업로드된 파일 또는 "rclone mount" 또는 google photos 또는 google
      docs로 업로드된 파일)의 경우, 이 청크 크기를 사용하여 멀티파트
      업로드를 수행합니다.
      
      알려진 크기의 큰 파일을 업로드할 때는 "--s3-upload-concurrency"명령어
      줄 단위 복사본당(transfer) 이 크기의 청크로 버퍼링됩니다.
      
      초고속 링크를 통해 대용량 파일을 전송하는 경우에 충분한 메모리가 있다면
      이 값을 높이면 전송 속도가 향상됩니다.
      
      대용량 파일을 전송하는 경우 맥스 10,000 청크 제한을 낮추기 위해 Rclone은
      청크 크기를 자동으로 증가시킵니다.
      
      알려진 크기의 파일은 구성된 청크 크기로 업로드됩니다.
      기본 청크 크기는 5 MiB이며 최대 10,000 청크를 가장 많이 가질 수 있으므로,
      기본적으로 48 GiB까지 파일을 스트림으로 업로드 할 수 있습니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태의 정확도가
      감소합니다. Rclone은 AWS SDK가 버퍼에 있는 청크를 보낼 때 청크가 보내진 것으로
      처리하지만 실제로 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 보고를 가져옵니다.
      

   --max-upload-parts
      멀티파트 업로드에서의 청크 최대 개수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 멀티파트 청크의 최대 개수를
      정의합니다.
      
      서비스가 10,000 청크의 AWS S3 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때
      이 청크 크기를 자동으로 증가시킬 것입니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 파일의 크기 기준.
      
      서버 측에서 복사해야 하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고, 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 객체를 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 아주 유용하지만,
      큰 파일을 업로드하는 데 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth 설정이 true(기본값)인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있는 경우 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어
      있다면 현재 사용자의 홈 디렉토리로 기본값이 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 공유 자격 증명 파일에서 사용할 프로필을 정의합니다.
      
      비워 둘 경우 환경 변수 "AWS_PROFILE" 또는
      설정되지 않은 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드에 대한 연속성.
      
      동일한 파일의 청크 개수를 동시에 업로드합니다.
      
      대용량 파일을 고속 링크로 업로드하고 이러한 업로드가 전체 대역폭을
      활용하지 못하는 경우, 이 값을 증가시키면 전송 속도를 향상시킬 수 있습니다.

   --force-path-style
      경로 스타일 액세스 시 true, 가상 호스팅 스타일 액세스 시 false를 사용합니다.
      
      true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고,
      false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은
      [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      참조.
      
      이 설정에 따라 rclone이 자동으로 설정합니다.
      (AWS, Aliyun OSS, Netease COS 또는 Tencent COS)과 같이 일부 제공자는
      false로 설정해야합니다.

   --v2-auth
      v2 인증을 사용하려면 true로 설정합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용하고, 설정되어 있으면
      rclone은 v2 인증을 사용합니다.
      
      v4 시그니처 작동하지 않을 때만 사용하세요 (예: Jewel/v10 CEPH 전 버전).

   --list-chunk
      목록 단위 크기(각 ListObject S3 요청에 대한 목록 리스트 크기).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로
     도 알려져있습니다.
      대부분의 서비스는 목록 응답을 1000개의 객체로 자름이 선호되며 그 이상의
      요청에 대해서도 1000개로 응답합니다.
      AWS S3에서는 이 값이 전역 최대값이며 변경할 수 없습니다. 자세한 내용은
      [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를
      참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을 증가시킬
      수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동 설정을 위한 0.
      
      S3가 처음 시작될 때 버킷의 객체를 나열할 수 있는 ListObjects 호출만
      제공했습니다.
      
      그러나 2016년 5월부터 ListObjectsV2 호출이 도입되었습니다. 이것이
      훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정할 경우 rclone은 사전에 지정된 제공자에 따라 목록
      객체 방법을 선택합니다. 틀린 경우 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록의 URL 인코드 여부: true/false/unset
      
      일부 제공자는 파일 이름에 특수 문자를 사용할 때 URL 인코딩 목록을 지원하며 해당
      경우 파일 이름에 특수 문자를 포함하는 경우 신뢰할 수 있는 방법입니다. unset으로
      설정하면(rclone의 기본 설정인 경우에) 제공자 설정에 따라 rclone이 적용할
      내용을 선택합니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우에 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자가 네 가지 버킷 생성 권한이 없는 경우 필요할 수 있습니다.
      v1.52.0 이전에는 버그로 인해 정상적으로 전달되었습니다.
      

   --no-head
      체크섬을 확인하기 위해 업로드된 객체에 HEAD를 수행하지 않습니다.
      
      rclone은 업로드한 후에 PUT로 객체에 대한 200 OK 메시지를 받으면
      제대로 업로드되었다고 가정합니다.
      
      특히 다음을 가정합니다:
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)이 업로드한 대로임
      - 크기가 업로드한 것과 같음
      
      하나의 부분 PUT의 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽어지지 않습니다.
      
      알려지지 않은 길이의 소스 객체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 오류가 감지되지 않는 확률이 높아지므로,
      일반적인 운영에는 권장되지 않습니다. 실제로 이 플래그를 사용해도
      업로드 오류가 무시될 가능성은 매우 낮습니다.
      

   --no-head-object
      객체를 가져오기 전에 HEAD를 수행하지 마십시오.

   --encoding
      백엔드에 대한 인코딩.
      
      더 많은 정보를 보려면 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도 설정.
      
      추가 버퍼가 필요한 업로드(예: multipart)는 메모리 풀을 사용하여
      할당됩니다.
      이 옵션은 미사용 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에 mmap 버퍼를 사용할 것인지 여부.

   --disable-http2
      S3 백엔드에서 http2 사용 중지.
      
      현재 s3 (구체적으로 minio) 백엔드에서 HTTP/2와 관련하여 해결되지 않은 문제가 있습니다.
      s3 백엔드의 HTTP/2는 기본적으로 활성화되어 있지만 여기에서 비활성화할 수 있습니다.
      문제가 해결되면이 플래그는 제거될 것입니다.
      
      참고: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위해 사용자 정의 엔드포인트.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우
      더 저렴한 아웃바운드(Egress)를 제공합니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 검증할지 여부
      
      true, false 또는 기본값(provider에 따름)을 사용해야 합니다.
      

   --use-presigned-request
      단일 부분 업로드를 위해 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 < 1.59는 단일 부분 객체를 업로드하기 위해 사전 서명된 요청을
      사용하고이 플래그를 true로 설정하면이 기능을 다시 활성화합니다. 이는
      특수한 상황 또는 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정한 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜(예: "2006-01-02"), 날짜 및 시간(예: "2006-01-02
      15:04:05") 또는 그 전시간에 대한 기간(예: "100d" 또는 "1h") 일 수 있습니다.
      
      이 설정을 사용하는 경우 파일 쓰기 작업을 수행할 수 없으므로
      파일을 업로드하거나 삭제할 수 없습니다.
      
      올바른 포맷에 대한 자세한 내용은 [시간 옵션 설명서](/docs/#time-option)를
      참조하십시오.
      

   --decompress
      gzip으로 인코딩된 객체를 해제하려면 설정하세요.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드할 수 있습니다.
      일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 이러한 파일을 "Content-Encoding: gzip"
     로 받을 때 객체를 해제합니다. 이로 인해 rclone은 크기와 해시를
      확인할 수 없지만 파일 내용은 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 압축할 수 있는 경우 설정하세요.
      
      일반적으로는 공급자가 객체를 다운로드할 때 변경하지 않습니다. 만약
      객체가 `Content-Encoding: gzip`로 업로드되지 않았다면,
      다운로드될 때 `Content-Encoding: gzip`가 설정되지 않습니다.
      
      그러나 일부 제공자는 `Content-Encoding: gzip`로 업로드되지
      않았더라도 객체를 gzip으로 압축 할 수 있습니다(예: Cloudflare).
      
      이 경우 `Content-Encoding: gzip`가 설정되어 있는 chunked
      transfer encoding으로 rclone이 객체를 다운로드하면 rclone은 객체를 실시간으로
      해제합니다.
      
      이 값을 unset으로 설정하면(rclone의 기본값인 경우) rclone이 적용할
      내용을 제공자 설정에 따라 선택합니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value             Qiniu Object Storage를 위한 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임(환경 변수 또는 env vars 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명 가져오기. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  지리적 제약 - 지역과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS 비밀 액세스 키(암호). [$SECRET_ACCESS_KEY]
   --storage-class value        Qiniu에서 새로운 객체를 저장할 때 사용할 저장 클래스. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              복제를 위해 청크로 전환하는 크기 기준. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용 중지. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               경로 스타일 액세스 시 true, 가상 호스팅 스타일 액세스 시 false를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 단위 크기(각 ListObject S3 요청에 대한 목록 리스트 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록의 URL 인코드 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 or 0 for auto. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서의 청크 최대 개수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도 설정. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에 mmap 버퍼를 사용할 것인지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 압축할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        체크섬을 확인하기 위해 업로드된 객체에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져오기 전에 HEAD를 수행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 연속성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 업로드 크기 기준. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드를 위해 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 인증을 사용하려면 true로 설정합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정한 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}