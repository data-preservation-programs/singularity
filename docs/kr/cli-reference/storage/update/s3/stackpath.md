# StackPath Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 stackpath - StackPath Object Storage

사용법:
   singularity storage update s3 stackpath [command options] <name|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명 (환경 변수 또는 env 변수나 EC2/ECS 메타데이터) 가져오기.
      
      access_key_id와 secret_access_key이 비어 있을 때만 적용됨.

      예시:
         | false | AWS 자격 증명을 다음 단계에서 입력합니다.
         | true  | 환경에서 (환경 변수 또는 IAM) AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(암호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --region
      연결할 지역.
      
      S3 클론을 사용하고 지역을 가지고 있지 않은 경우 비워 둡니다.

      예시:
         | <unset>            | 향후 결정하려면 사용하세요.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때 사용하세요.
         |                    | 예: 이전 Jewel/v10 CEPH.

   --endpoint
      StackPath Object Storage의 엔드포인트.

      예시:
         | s3.us-east-2.stackpathstorage.com    | US East 엔드포인트
         | s3.us-west-1.stackpathstorage.com    | US West 엔드포인트
         | s3.eu-central-1.stackpathstorage.com | EU 엔드포인트

   --acl
      버킷을 생성하고 개체를 저장하거나 복사할 때 사용되는 Canned ACL.
      
      이 ACL은 개체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3는 ACL을 소스에서 복사하는 대신 새로 작성하므로 서버 쪽 객체 복사시에도 이 ACL이 적용됩니다.
      
      ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자들은 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 읽기 액세스 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스 권한 부여.
         |                    | 버킷에 이를 설정하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스 권한 부여.

   --upload-cutoff
      청크 업로드로 전환하기 위한 파일 크기 기준.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat"에서 가져온 파일, "rclone mount" 또는 Google 사진 또는 Google 문서와 업로드한 파일)을 업로드할 때 이 청크 크기를 사용하여 멀티파트 업로드를 수행합니다.
      
      참고로 "--s3-upload-concurrency"의 청크는 전송 당 메모리에 이 크기로 버퍼링됩니다.
      
      고속링크를 통해 대용량 파일을 전송하고 메모리가 충분한 경우 이 값을 높이면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 최대 10,000개의 청크 제한을 유지하도록 청크 크기를 자동으로 증가시킵니다.
      
      크기를 알 수 없는 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크를 가질 수 있으므로 기본적으로 스트림 업로드 가능한 파일 크기는 최대 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행률 통계의 정확성이 감소합니다. Rclone은 AWS SDK에 의해 버퍼링된 청크를 보낼 때 청크가 전송된 것으로 간주하지만 실제로는 여전히 업로드 중일 수 있습니다.
      큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률이 실제와 더 차이 나는 진행률 보고를 의미합니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 파트 수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      S3 사양에서 10,000개의 청크를 지원하지 않는 경우에 유용합니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 최대 청크 수 제한을 유지하도록 청크 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 파일 크기 기준.
      
      서버 쪽으로 복사해야 하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력 데이터의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이것은 데이터 무결성 확인에는 훌륭하지만 대용량 파일의 경우 업로드 시작에 오랜 지연을 초래할 수 있습니다.

   --shared-credentials-file
      공용 자격 증명 파일의 경로.
      
      env_auth = true인 경우 rclone은 공용 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공용 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true인 경우 rclone은 공용 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워 둘 경우 "AWS_PROFILE"이나 "default"라는 환경 변수를 기본값으로 사용합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.
      
      동일한 파일의 청크를 동시에 업로드합니다.
      
      고속링크와 함께 대용량 파일을 속도 최적화하기 위해 이를 풀로 활용하지 못할 경우 이 값을 높이는 것이 전송 속도를 높일 수 있습니다.

   --force-path-style
      true인 경우 path 스타일 액세스 사용, false인 경우 virtual hosted 스타일 사용.
      
      true로 설정하면 rclone은 path 스타일 액세스를 사용하며, false로 설정하면 virtual path 스타일을 사용합니다. 자세한 내용은 [AWS S3 설명서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이를 수행합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false로 설정하면 rclone은 v4 인증을 사용합니다. 설정되어 있으면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 사용하세요. 예: 이전 Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(ListObject S3 요청마다 응답 목록이 설정되는 크기).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청 수를 무조건적으로 1000개로 자릅니다.
      AWS S3에서는 이 값이 전역 최대값으로 고정되어 있으며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 더 크게 설정할 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전: 1,2 또는 0은 자동.
      
      S3를 처음 시작할 때는 버킷의 개체를 열거하기 위해 ListObjects 호출만 사용할 수 있었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하므로 가능한 경우에는 사용해야 합니다.
      
      기본 설정인 0은 rclone이 공급자 설정에 따라 호출할 List Objects 방법을 추측합니다. 잘못 추측할 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 가능한 경우 이 방법은 보다 신뢰성 있습니다. 설정되지 않은 경우(기본 설정) rclone은 공급자 설정에 따라 적용할 내용을 선택합니다. 그러나 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않습니다.
      
      번거로운 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다. 버킷이 이미 존재하는 경우에만 필요합니다.
      
      사용자가 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. 버전 1.52.0 이전에는 버그로 인해 이러한 경우에도 오류가 발생하지 않았습니다.
      

   --no-head
      업로드한 개체에 대해 HEAD를 사용하여 무결성을 확인하지 않습니다.
      
      트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      이 옵션을 설정하면 PUT으로 개체를 업로드한 후에 200 OK 메시지를 수신하면 올바르게 업로드된 것으로 간주합니다.
      
      특히 다음의 것으로 간주합니다:
      
      - 메타데이터(수정 시간, 스토리지 클래스 및 컨텐츠 유형)가 업로드한 대로인 것
      - 크기가 업로드한 것
      
      다음의 정보를 단일 파트 PUT의 응답 헤더에서 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 정보를 읽을 수 없습니다.
      
      길이를 알 수 없는 원본 개체가 업로드되는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 오류 감지 확률이 증가하며 특히 크기가 잘못된 경우에는 추천되지 않습니다. 실제로 업로드 오류가 감지되지 않을 확률은 매우 적습니다.
      

   --no-head-object
      GET을 수행하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 정보는 [개요에서의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 간격.
      
      추가 버퍼가 필요한 업로드(예: multipart)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않는 버퍼가 풀에서 제거되는 간격을 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에 대한 http2 사용 비활성화.
      
      현재 S3(특히 minio) 백엔드와 HTTP/2에 대한 문제가 미해결 상태입니다. HTTP/2는 S3 백엔드의 기본 설정이지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면 이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는 CloudFront 네트워크를 통해 다운로드되는 데이터에 대해 더 저렴한 나가부 트래픽 가격을 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 사용자 정의를 설정하지 않고 공급자의 기본값을 사용합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부
      
      false로 설정하면 rclone은 AWS SDK에서 PutObject를 사용하여 개체를 업로드합니다.
      
      rclone의 버전 1.59 이하에서는 단일 파트 개체를 업로드하려면 사전 서명된 요청을 사용하고 설정을 true로 지정하면 해당 기능이 다시 활성화됩니다. 이는 특수한 상황이나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 날짜와 시간("2006-01-02 15:04:05") 또는 해당 시점에서의 이전 기간("100d" 또는 "1h")일 수 있습니다.
      
      이를 사용하는 경우 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [time 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩된 개체를 압축 해제합니다.
      
      AWS S3에서 "Content-Encoding: gzip"로 파일을 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 이러한 파일을 받으면 "Content-Encoding: gzip"로 개체를 압축 해제합니다. 이로 인해 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 gzip으로 개체를 압축할 수 있는 경우 설정하세요.
      
      일반적으로 공급자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드할 때 설정되지 않습니다.
      
      그러나 일부 제공자는 gzip으로 압축하지 않은 객체들을 gzip으로 압축할 수 있습니다(예: Cloudflare).
      
      이러한 경우 다음과 같은 오류 메시지를 수신합니다:
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip이 설정되고 청크 전송 인코딩이 있는 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset으로 설정할 경우 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 생성하고 개체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value           StackPath Object Storage의 엔드포인트. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명 (환경 변수 또는 env 변수나 EC2/ECS 메타데이터) 가져오기. (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키(암호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 파일 크기 기준. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 path 스타일 액세스 사용, false인 경우 virtual hosted 스타일 사용. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(ListObject S3 요청마다 응답 목록이 설정되는 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 url 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1,2 또는 0은 자동. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 파트 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 간격. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 gzip으로 개체를 압축할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 개체에 대해 HEAD를 사용하여 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공용 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공용 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 파일 크기 기준. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}