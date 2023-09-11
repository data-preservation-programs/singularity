# IDrive e2

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 idrive - IDrive e2

사용법:
   singularity storage update s3 idrive [command options] <name|id>

설명:
   --env-auth
      실행 시간에 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env 변수 또는 IAM).

      익명 액세스를 위해 빈 문자열로 남겨 둡니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경에서 AWS 자격 증명을 가져옵니다 (환경변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.

      익명 액세스 또는 실행 시간 자격 증명을 위해 빈 문자열로 남겨 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호)입니다.

      익명 액세스 또는 실행 시간 자격 증명을 위해 빈 문자열로 남겨 둡니다.

   --acl
      버킷을 만들거나 객체를 저장 또는 복사할 때 사용되는 간편 ACL입니다.

      이 ACL은 객체를 생성할 때 사용되며, 버킷_ACL이 설정되어 있지 않은 경우 버킷을 생성할 때도 사용됩니다.

      자세한 정보는 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하십시오.

      반드시 기억하세요. 이 ACL은 S3가 객체를 서버간 복사 할 때 적용됩니다. 여기서 ACL은 소스에서 ACL을 복사하는 대신 새로 작성합니다.

      ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더를 추가하지 않고 기본값 (비공개)을 사용합니다.

   --bucket-acl
      버킷을 만들 때 사용되는 간편 ACL입니다.

      자세한 정보는 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하십시오.

      이 ACL은 오직 버킷을 만들 때만 적용됩니다. 설정하지 않으면 "acl"이 대신 사용됩니다.

      "acl"과 "bucket_acl"이 모두 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (비공개)을 사용합니다.

      예제:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사람에게 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 읽기 액세스 권한이 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스 권한이 부여됩니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 인증된 사용자 그룹에게 읽기 액세스 권한이 부여됩니다.

   --upload-cutoff
      청크 업로드로 전환하기 위한 크기 임계값입니다.

      이보다 큰 파일들은 chunk_size 단위로 청크별로 업로드됩니다.
      최소값은 0이며 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 (예: "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.

      "--s3-upload-concurrency" 청크 크기의 청크가 전송마다 메모리에 버퍼링됩니다.

      고속 링크로 큰 파일을 전송하고 메모리가 충분하면 이 값을 증가시키면 전송 속도가 향상됩니다.

      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 10,000 청크 제한을 유지합니다.

      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 청크까지 가능하기 때문에 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드 하려면 chunk_size를 증가시켜야 합니다.

      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 감소합니다. Rclone은 AWS SDK에 의해 버퍼링되었을 때 청크를 전송된 것으로 처리하지만 업로드중일 수도 있습니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률이 진실에서 더 벗어나는 진행 보고를 의미합니다.

   --max-upload-parts
      멀티파트 업로드에서 최대로 사용될 청크 수입니다.

      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.

      제공 서비스가 10,000 청크의 AWS S3 사양을 지원하지 않는 경우 유용합니다.

      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 초과하지 않습니다.

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 크기 임계값입니다.

      서버간 복사해야 하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이며 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 객체를 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 첨부합니다. 이는 데이터 무결성 확인에 탁월하지만 대용량 파일의 업로드 시작에 대기 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본 설정됩니다.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.

      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 설정됩니다.

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드를 위한 동시 실행 횟수입니다.

      고속 링크로 소량의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우 이 숫자를 증가시킴으로써 전송 속도를 높일 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.

      기본값으로 true이므로 rclone은 경로 스타일 액세스를 사용하고 false인 경우 가상 경로 스타일을 사용합니다. [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.

      일부 프로바이더 (예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이렇게 설정합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.

      false 인 경우 (기본값) rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.

      v4 서명이 작동하지 않는 경우에만 사용합니다. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록)입니다.

      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 응답 목록을 1000개 이상 요청해도 1000개로 잘라냅니다.
      AWS S3에서는 이것이 전역 최대값이며 변경할 수 없으며 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 증가할 수 있습니다.

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동 (0).

      S3가 처음 출시될 때는 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.

      그러나 ListObjectsV2 호출은 2016년 5월에 도입되었습니다. 이는 매우 높은 성능을 제공하며 가능하다면 사용해야 합니다.

      기본값 0으로 설정하면 rclone은 원래 설정된 공급자에 따라 호출할 목록 객체 방법을 추측합니다. 잘못 추측한다면 여기서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 URL 인코딩 할지 여부: true/false/unset

      일부 프로바이더는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 사용 가능한 경우 이 옵션은 일관된 작동을 위해 제안하는 rclone의 선택을 무시하고 적용할 대상을 선택할 수 있습니다. rclone의 선택을 무시할 수 있습니다.

   --no-check-bucket
      버킷을 확인하거나 만들기 위해 시도하지 마십시오.

      버킷이 이미 존재하는 경우 rclone 작업 횟수를 최소화하려는 경우 유용할 수 있습니다.

      버킷 생성 권한이 없는 경우도 필요할 수 있습니다. v1.52.0 이전까지는 버그 때문에 무시되었습니다.

   --no-head
      업로드한 객체의 상태를 확인하기 위해 HEAD를 사용하지 마십시오.

      작업 횟수를 최소화하려는 경우 유용할 수 있습니다.

      200 OK 메시지를 받은 경우 rclone은 PUT로 객체를 업로드한 후 제대로 업로드되었다고 가정합니다.

      특히 다음 항목을 가정합니다.

      - Metadata (modtime, 저장 클래스 및 콘텐츠 유형)이 업로드 된 것처럼입니다.

      - 크기가 업로드된 것처럼입니다.

      단일 부분 PUT의 응답에서 다음 항목을 읽습니다.

      - MD5SUM

      - 업로드된 날짜

      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.

      길이를 알 수 없는 소스 객체가 업로드된 경우 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 업로드 확인 오류의 가능성이 높아지므로 정상적인 동작에는 권장되지 않습니다. 실제로, 이 플래그로 인한 업로드 오류의 가능성은 매우 낮습니다.

   --no-head-object
      GET하는 동안 HEAD를 수행하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.

      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.

      추가 버퍼가 필요한 업로드 (예 : 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 수집될 때까지의 시간 관계를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.

      S3 (특히 minio) 백엔드와 HTTP/2의 미해결된 문제가 현재 존재합니다. S3 백엔드의 기본값에는 HTTP/2가 활성화되어 있지만 여기서 비활성화할 수 있습니다. 문제가 해결되면이 플래그는 제거됩니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드를위한 사용자 정의 엔드포인트입니다.
      보다 저렴한 데이터 나가기를 제공하는 AWS S3의 CloudFront CDN URL로 일반적으로 설정됩니다.

   --use-multipart-etag
      확인을 위해 multipart 업로드에서 ETag를 사용할지 여부입니다.

      이것은 true, false 또는 provider의 기본값을 사용할 지 여부로 설정하십시오.

   --use-presigned-request
      단일 파트 업로드에 대해 사전 서명된 요청을 사용할지 여부입니다.

      이 값이 false이면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.

      rclone < 1.59의 버전은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하고이 플래그를 true로 설정하면 이 기능을 다시 활성화합니다. 이는 예외적인 경우나 테스트할 때에만 필요합니다.

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간대의 파일 버전을 표시합니다.

      매개 변수는 날짜, "2006-01-02", 날짜 시간 "2006-01-02 15:04:05" 또는 그 시간 전의 기간 (예: "100d" 또는 "1h") 일 수 있습니다.

      이 옵션을 사용할 때 파일 작성 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      유효한 형식은 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

   --decompress
      이 전달되면 gzip으로 인코딩된 객체를 압축 해제합니다.

      S3로 "Content-Encoding: gzip"을 설정하여 객체를 업로드 할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축되어 전송된 파일로 다운로드합니다.

      이 플래그가 설정되면 rclone은 받은대로 "Content-Encoding: gzip"로 이 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

   --might-gzip
      백엔드가 개체를 gzip으로 인코딩 할 수도 있다면이를 설정하십시오.

      일반적으로 공급자는 개체가 다운로드 될 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체에는 설정되지 않습니다.

      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 객체도 gzip으로 압축 할 수 있습니다. (예: Cloudflare).

      이런 경우 다음과 같은 오류가 발생합니다.

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      이 플래그를 설정하고 rclone이 컨텐츠가 chunked 전송 인코딩으로 gzip을 설정한 객체를 다운로드하면 rclone은 객체를 실시간으로 압축 해제합니다.

      이 변수가 설정되지 않으면 제공자 설정에 따라 rclone이 적용할 내용을 선택합니다. rclone의 선택을 무시할 수 있습니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                버킷을 만들거나 객체를 저장 또는 복사할 때 사용되는 간편 ACL입니다. [$ACL]
   --env-auth                 실행 시간에 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env 변수 또는 IAM). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말을 표시합니다.
   --secret-access-key value  AWS 비밀 액세스 키 (비밀번호)입니다. [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 간편 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. Rclone은 알 수없는 크기의 파일을 구성된 chunk_size로 업로드합니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 크기 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할 지 여부입니다: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0 (자동) (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 최대로 사용될 청크 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 인코딩 할 수도 있다면이를 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷을 확인하거나 만들기 위해 시도하지 마십시오. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 상태를 확인하기 위해 HEAD를 사용하지 마십시오. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET하는 동안 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드를 위한 동시 실행 횟수입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 크기 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 multipart 업로드에서 ETag를 사용할지 여부입니다 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명 된 요청을 사용할지 여부입니다 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간대의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}