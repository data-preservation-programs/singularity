# IONOS Cloud

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 ionos - IONOS Cloud

사용법:
   singularity storage update s3 ionos [command options] <name|id>

설명:
   --env-auth
      런타임으로부터 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터에서 가져옴).

      access_key_id와 secret_access_key가 비어있을 경우에만 적용됩니다.

      예시:
         | false | AWS 자격 증명을 다음 단계에서 입력합니다.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호)입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --region
      버킷이 생성되고 데이터가 저장될 지역입니다.

      예시:
         | de           | 독일 프랑크푸르트
         | eu-central-2 | 독일 베를린
         | eu-south-2   | 스페인 로그로뇨

   --endpoint
      IONOS S3 Object Storage의 엔드포인트입니다.

      같은 지역의 엔드포인트를 지정하세요.

      예시:
         | s3-eu-central-1.ionoscloud.com | 독일 프랑크푸르트
         | s3-eu-central-2.ionoscloud.com | 독일 베를린
         | s3-eu-south-2.ionoscloud.com   | 스페인 로그로뇨

   --acl
      버킷 및 객체 생성 또는 복사 시 사용되는 전용 ACL입니다.

      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.

      주의사항: 이 ACL은 S3 서버간의 객체 복사 시 적용됩니다.
      S3는 소스로부터 ACL을 복사하는 것이 아니라 새로 작성할 뿐입니다.

      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(Private)이 사용됩니다.

   --bucket-acl
      버킷 생성 시 사용되는 전용 ACL입니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl.

      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않았다면 "acl"이 대신 사용됩니다.

      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(Private)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자에게 엑세스 권한이 없음(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 그룹 AllUsers에게 READ 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 그룹 AllUsers에게 READ 및 WRITE 권한 부여.
         |                    | 버킷에 대해 이를 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 그룹 AuthenticatedUsers에게 READ 권한 부여.

   --upload-cutoff
      청크로 업로드로 전환되는 파일의 임계값입니다.

      이보다 큰 파일은 chunk_size별로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일(예: "rclone rcat"이나 "rclone mount" 또는 Google Photos나 Google Docs에서 업로드된 파일)을 업로드할 때, 이 청크 크기를 사용하여 청크별로 멀티파트 업로드로 업로드됩니다.

      "--s3-upload-concurrency" 크기별로 이 청크 크기만큼 버퍼에 유지됩니다.

      높은 속도 링크에서 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시키면 전송 속도가 빨라집니다.

      큰 파일을 전송할 때 Rclone은 확인된 크기의 대부분 파일에서 10,000개의 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.

      크기를 알 수 없는 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 증가시켜야 합니다.

      청크 크기가 증가하면 "-P" 플래그와 함께 표시되는 진행 통계의 정확성이 낮아집니다. Rclone은 AWS SDK가 버퍼에 저장되는 청크를 전송한 것으로 처리하지만 사실은 여전히 업로드 중일 수 있습니다. 큰 청크 크기는 큰 AWS SDK 버퍼와 실제로 업로드 중인 것과는 더 다른 진행 보고를 생성합니다.

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 청크 수입니다.

      이 옵션은 멀티파트 업로드를 할 때 사용할 멀티파트 청크의 최대 수를 결정합니다.

      이 옵션은 10,000 청크로 제한된 AWS S3 사양을 지원하지 않는 서비스에 유용할 수 있습니다.

      Rclone은 확인된 크기의 대부분 파일에서 10,000개의 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.

   --copy-cutoff
      청크 단위로 복사해야하는 이 파일보다 큰 파일의 임계값입니다.

      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만 큰 파일을 업로드하기 전에 오랜 지체가 있을 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본 설정됩니다.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 그 파일에서 사용할 프로필을 제어합니다.

      비워 둔 경우 "AWS_PROFILE" 또는 그 환경 변수가 설정되어 있지 않은 경우 "default"로 기본 설정됩니다.

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.

      동일한 파일의 청크 수를 동시에 업로드합니다.

      높은 속도 링크에서 작은 수의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 path 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.

      true(기본값)인 경우 rclone은 path 스타일 액세스를 사용하고,
      false인 경우 rclone은 가상 path 스타일을 사용합니다. 자세한 내용은 [AWS S3
      문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.

      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값을
      false로 설정해야 합니다. rclone은 제공자 설정에 따라 자동으로 설정합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.

      false(기본값)인 경우 rclone은 v4 인증을 사용하고, 설정하면 rclone은 v2 인증을 사용합니다.

      v4 서명이 작동하지 않는 경우에만 사용하세요. 예를 들어 Jewel/v10 CEPH 이전입니다.

   --list-chunk
      목록 청크의 크기(ListObject S3 요청마다 응답 목록 크기)입니다.

      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 불립니다.
      대부분의 서비스는 요청된 개수가 1000개 이상이더라도 응답 목록을 1000개로 잘라냅니다.
      AWS S3에서는 전역 최대값으로 설정되어 도움말을 참조해도 변경할 수 없습니다.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 늘릴 수 있습니다.

   --list-version
      사용할 ListObjects의 버전: 1, 2 또는 자동으로 0.

      S3가 처음 출시되었을 때는 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.

      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훌륭한 성능을 제공하며 가능한 한 사용해야 합니다.

      기본값인 0으로 설정되어 있는 경우, rclone은 제공자에 따라 호출할 ListObjects 방법을 추측합니다. 잘못 추측하면 여기서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset

      일부 공급자는 URL 인코딩 목록을 지원하며, 가능한 경우 파일 이름에 제어 문자를 사용할 때 신뢰할 수 있는 방법입니다. unset으로 설정된 경우 rclone은 공급자 설정에 따라 적용할 항목을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 버킷을 생성하지 않으려면 설정하세요.

      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.

      이는 사용자에게 버킷 생성 권한이 없는 경우도 필요할 수 있습니다. 버전 1.52.0 이전까지 이는 버그로 인해 무시되었습니다.

   --no-head
      체크섬을 확인하기 위해 업로드 된 객체의 HEAD를 수행하지 않습니다.

      rclone은 가능한 경우 최대한 트랜잭션 수를 최소화하려는 경우 유용합니다.

      현재의 200 OK 메시지를 받으면 PUT으로 객체를 업로드한 후 올바르게 업로드된 것으로 가정합니다.

      특히 다음 항목을 가정합니다:

      - metadata(예: modtime, storage class 및 content type)가 업로드한 것과 같다.
      - 크기가 업로드한 것과 같다.

      단일 파트 PUT의 응답에서 다음 항목을 읽습니다:

      - MD5SUM
      - 업로드된 날짜

      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.

      길이를 알 수 없는 원본 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 업로드 실패 가능성이 증가하며 특히 올바르지 않은 크기 때문에 일반적인 운영에 권장되지 않습니다. 실제로 이 플래그로 인해 업로드 실패 가능성은 매우 작습니다.

   --no-head-object
      GET을 수행하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩 방식입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 플러시할지 제어합니다.

      추가 버퍼(예: 멀티파트)가 필요한 업로드는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.

      현재 s3(특히 minio) 백엔드에서 HTTP/2에 문제가 있습니다. HTTP/2는 s3 백엔드의 기본값이지만 여기서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트입니다.
      AWS S3는 CloudFront 네트워크를 통해 다운로드된 데이터에 대해 더 저렴한 이그레스를 제공하기 때문에 일반적으로 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에 ETag를 사용할지 여부입니다.

      true, false 또는 기본값을 사용하세요.

   --use-presigned-request
      단일 파트 업로드에 프리사인된 요청 또는 PutObject를 사용할지 여부입니다.

      false인 경우 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.

      rclone 버전 1.59 미만은 단일 파트 객체를 업로드하기 위해 프리사인된 요청을 사용하며, 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트 외에는 필요하지 않습니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.

      매개변수는 날짜("2006-01-02"), 날짜 및 시간("2006-01-02 15:04:05") 또는 그보다 오래된 기간("100d" 또는 "1h")이 될 수 있습니다.

      이 옵션을 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      유효한 형식에 대한 자세한 내용은 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.

   --decompress
      이 옵션을 설정하면 gzip으로 인코딩된 객체를 압축 해제합니다.

      S3에 "Content-Encoding: gzip"으로 업로드되는 객체가 있을 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.

      이 플래그를 설정하면 rclone은 "Content-Encoding: gzip"으로 수신했을 때 이러한 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 해제됩니다.

   --might-gzip
      백엔드가 객체를 gzip할 수 있는 경우 true로 설정하세요.

      일반적으로 공급자는 객체를 다운로드할 때 객체를 수정하지 않습니다. "Content-Encoding: gzip"으로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.

      그러나 어떤 공급자는 "Content-Encoding: gzip"이 아닌 파일이 업로드되어도 객체를 gzip할 수 있습니다(Cloudflare 예).

      다음과 같은 오류가 발생한다면 그것이 의심의 여지가 있으며,

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      이 플래그를 설정하면 rclone은 chunked 전송 인코딩 및 "Content-Encoding: gzip"이 설정된 객체를 실시간으로 압축 해제합니다.

      unset으로 설정된 경우 rclone은 공급자 설정에 따라 적용할 항목을 선택하지만 rclone의 선택을 여기에서 무시할 수 있습니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                버킷 및 객체 생성 또는 복사 시 사용되는 전용 ACL입니다. [$ACL]
   --endpoint value           IONOS S3 Object Storage의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 런타임으로부터 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터에서 가져옴). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말을 표시합니다
   --region value             버킷이 생성되고 데이터가 저장될 지역입니다. [$REGION]
   --secret-access-key value  AWS Secret Access Key(비밀번호)입니다. [$SECRET_ACCESS_KEY]

   고급 옵션

   --bucket-acl value               버킷 생성 시 사용되는 전용 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 복사하는 파일의 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩 방식입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 path 스타일 액세스를 사용합니다. false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(ListObject S3 요청마다 응답 목록 크기)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1, 2 또는 자동으로 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 청크 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 플러시할지 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip할 수 있는 경우 true로 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 버킷을 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        체크섬을 확인하기 위해 업로드 된 객체의 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크로 업로드로 전환되는 파일의 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에 ETag를 사용할지 여부입니다. (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 프리사인된 요청을 사용할지 여부입니다. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 인증을 사용할지 여부입니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]
```
{% endcode %}