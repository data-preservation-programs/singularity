# Arvan Cloud Object Storage (AOS)

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 arvancloud - Arvan Cloud Object Storage (AOS)

사용법:
   singularity storage update s3 arvancloud [옵션] <이름|ID>

설명:
   --env-auth
      런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 런타임 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스나 런타임 자격 증명을 위해 비워둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(암호)입니다.
      
      익명 액세스나 런타임 자격 증명을 위해 비워둡니다.

   --endpoint
      Arvan Cloud Object Storage (AOS) API의 엔드포인트입니다.

      예시:
         | s3.ir-thr-at1.arvanstorage.com | 기본 엔드포인트 - 확실하지 않을 경우 좋은 선택입니다.
         |                                | 이란, 테헤란 (아시아텍)
         | s3.ir-tbz-sh1.arvanstorage.com | 이란, 타브리즈 (샤리아르)

   --location-constraint
      위치 제약 조건 - 엔드포인트와 일치해야 합니다.
      
      버킷 생성 시에만 사용됩니다.

      예시:
         | ir-thr-at1 | 이란, 테헤란 (아시아텍)
         | ir-tbz-sh1 | 이란, 타브리즈 (샤리아르)

   --acl
      객체를 생성하거나 복사할 때 사용되는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되지 않은 경우에도 bucket 생성에 사용됩니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      참고로, 이 ACL은 S3의 서버 간 복사 시 적용됩니다.
      소스에서 ACL을 복사하지 않고 새로운 ACL을 기록합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL입니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자가 FULL_CONTROL을 갖습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다. (기본값)
         | public-read        | 소유자가 FULL_CONTROL을 갖습니다.
         |                    | AllUsers 그룹에게 읽기 액세스 권한이 있습니다.
         | public-read-write  | 소유자가 FULL_CONTROL을 갖습니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스 권한이 있습니다.
         |                    | 버킷에 대해서는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL을 갖습니다.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스 권한이 있습니다.

   --storage-class
      ArvanCloud에 새로운 객체를 저장할 때 사용할 저장 클래스입니다.

      예시:
         | STANDARD | 표준 저장 클래스

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값입니다.
      
      이보다 큰 크기의 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나 알 수 없는 크기의 파일(예: "rclone rcat"으로 업로드된 파일이나 "rclone mount" 또는 Google 포토 또는 Google 문서에서 업로드된 파일)을 업로드할 때, 이 청크 크기를 사용하여 다중 파트 업로드로 업로드됩니다.
      
      참고로, "--s3-upload-concurrency"의 크기 청크는 전송 당 메모리에 버퍼링됩니다.
      
      대역폭을 충분히 사용하지 않거나 고속 링크를 통해 큰 파일을 전송하는 경우, 이 값을 늘리면 전송 속도가 향상됩니다.
      
      rclone은 알려진 크기의 큰 파일을 업로드할 때 10,000 청크 제한을 준수하기 위해 청크 크기를 자동으로 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본적으로 청크 크기는 5 MiB이며 최대 10,000 청크를 사용할 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링될 때 전송된 것으로 처리하지만 실제로는 업로드 중일 수 있습니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 청크 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 multipart 청크의 최대 수를 정의합니다.
      
      10,000 청크와 같은 AWS S3 사양을 지원하지 않는 서비스의 경우 유용할 수 있습니다.
      
      알려진 크기의 큰 파일을 업로드할 때 chunk_size를 자동으로 증가시키는 rclone입니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.
      
      이보다 큰 파일을 서버 간 복사해야하는 경우 이 크기로 청크별로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      보통 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 데이터 무결성 검사에는 좋지만 큰 파일의 업로드를 시작하는 데 오랜 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있는 경우 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      이 값이 비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크를 동시에 업로드하는 수입니다.
      
      대역폭을 충분히 사용하지 않고 큰 파일을 소폭 업로드하는 경우 이 값을 늘리면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고, false인 경우 가상 호스트 스타일 액세스를 사용합니다.
      
      이 값이 true인 경우(기본값), rclone은 경로 스타일 액세스를 사용하고, false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 제공자(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 이 값을 false로 설정해야 할 수 있습니다. rclone은 제공자 설정을 기반으로 이를 자동으로 수행합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      이 값이 false인 경우(기본값) rclone은 v4 인증을 사용하고, 설정되어 있는 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 이 값 사용합니다. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기(ListObject S3 요청마다 응답 목록의 크기)입니다.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청한 개수보다 많은 응답 목록을 잘라냅니다. AWS S3에서 이 값은 전역 최대값이며 수정할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동(0).
      
      S3가 처음 출시되었을 때 버킷의 객체를 열거하는 ListObjects 호출만 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone이 설정된 제공자에 따라 어떤 목록 객체 메서드를 호출할지 추측합니다. 잘못 추측하면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩해야 하는지 여부: true/false/unset
      
      일부 제공자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 사용 가능한 경우 이렇게 사용하는 것이 더 안정적입니다. 이 값이 unset으로 설정되어 있는 경우 rclone은 제공자 설정에 따라 적용할 것을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 것을 알고 있다면 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이것이 무시되었을 것입니다.
      

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone이 PUT으로 객체를 업로드한 후 200 OK 메시지를 받으면 제대로 업로드된 것으로 간주합니다.
      
      특히 다음을 가정합니다:
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)은 업로드한 것과 동일합니다.
      - 크기가 업로드한 것과 동일합니다.
      
      단일 부분 PUT의 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      이 항목들은 멀티파트 업로드에서는 읽지 않습니다.
      
      알려지지 않은 길이의 원본 객체가 업로드된 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 검출 가능성이 높아져 일반 작업에는 권장되지 않습니다. 실제로 업로드 실패의 가능성은 이 플래그를 사용해도 매우 적습니다.
      

   --no-head-object
      GET을 할 때 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 플러시할지 정합니다.
      
      추가 버퍼가 필요한 업로드(s3 로 업로드)는 메모리 풀을 할당하기 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. s3 백엔드의 HTTP/2는 기본적으로 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면 이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 이전을 제공합니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 검증에 사용할지 여부
      
      이 값은 true, false 또는 기본값(provider에 따라)로 설정할 수 있습니다.
      

   --use-presigned-request
      단일 부분 업로드에 프리사인된 요청 또는 PutObject를 사용할지 여부
      
      false인 경우 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      rclone < 1.59 버전은 단일 부분 객체를 업로드하기 위해 프리사인된 요청을 사용하고, 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다.
      이는 특정한 경우나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 날짜와 시간("2006-01-02 15:04:05") 또는 그 만큼의 시간 동안으로 지정할 수 있습니다.
      
      이를 사용하면 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"로 올릴 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 "Content-Encoding: gzip"로 수신된 파일을 압축 해제합니다. 즉, rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는지 여부입니다.
      
      보통 제공자는 다운로드될 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체에는 설정되지 않습니다.
      
      그러나 일부 제공자는 gzip으로 압축하지 않은 개체를 gzip으로 압축할 수 있습니다(Cloudflare 등).
      
      이런 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip과 청크 전송 인코딩이 설정된 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      기본값인 unset으로 설정된 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value        AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                  객체를 생성하거나 복사할 때 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value             Arvan Cloud Object Storage (AOS) API의 엔드포인트입니다.[$ENDPOINT]
   --env-auth                   런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 조건 - 엔드포인트와 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --secret-access-key value    AWS 비밀 액세스 키(암호)입니다. [$SECRET_ACCESS_KEY]
   --storage-class value        ArvanCloud에 새로운 객체를 저장할 때 사용할 저장 클래스입니다. [$STORAGE_CLASS]

   Advanced
   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고, false인 경우 가상 호스트 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩해야 하는지 여부입니다. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1, 2 또는 0 (기본값: 자동) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 청크 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는지 여부입니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 할 때 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 검증에 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 프리사인된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}