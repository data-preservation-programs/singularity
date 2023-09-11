# Liara Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 liara - Liara 객체 스토리지

사용법:
   singularity storage create s3 liara [옵션] [인자]

설명:
   --env-auth
      런타임(환경 변수 또는 env vars 또는 IAM)에서 AWS 자격 증명 가져오기
      액세스 키 ID(access_key_id)와 시크릿 액세스 키(secret_access_key)가 비어있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명 가져오기(env vars 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워두세요.

   --secret-access-key
      AWS 시크릿 액세스 키(비밀번호)

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워두세요.

   --endpoint
      Liara 객체 스토리지 API 엔드포인트

      예시:
         | storage.iran.liara.space | 기본 엔드포인트
         |                          | 이란

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용할 Canned ACL

      이 ACL은 객체 생성 및 (bucket_acl이 설정되지 않은 경우) 버킷 생성에 사용됩니다.

      자세한 내용은 다음 링크를 참조하세요: [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)

      Source에서 복사하는 경우 S3는 ACL을 복사하지 않고 새로운 ACL을 작성합니다.

      acl이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고 기본 설정(개인)이 사용됩니다.

   --bucket-acl
      버킷 생성 시 사용할 Canned ACL

      자세한 내용은 다음 링크를 참조하세요: [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)

      bucket_acl이 설정되지 않은 경우 "acl"이 대신 사용됩니다.

      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고 기본 설정(개인)이 사용됩니다.

      예시:
         | private            | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 다른 사용자는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 권한을 가집니다.
         | public-read-write  | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 가집니다.
         |                    | 버킷에 대해서는 general권장사항이 아닙니다.
         | authenticated-read | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 인증된 사용자 그룹은 읽기 권한을 가집니다.

   --storage-class
      Liara에 새로운 객체를 저장할 때 사용할 스토리지 클래스

      예시:
         | STANDARD | 표준 스토리지 클래스

   --upload-cutoff
      청크 전송으로 전환할 크기 기준

      이보다 큰 파일은 chunk_size 크기로 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일 등)을 업로드하는 경우, 이 청크 크기를 사용하여 청크로 업로드됩니다.

      참고로, "--s3-upload-concurrency" 크기의 청크는 각 전송에 대해 메모리에 버퍼링됩니다.

      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 늘리면 전송 속도가 향상됩니다.

      Rclone은 10,000개의 청크 제한을 유지하기 위해 알려진 크기의 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시킵니다.

      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의
      청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을
      스트림 업로드하려면 chunk_size를 증가시켜야 합니다.

      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확도가 감소합니다.
      Rclone은 청크가 AWS SDK에 의해 버퍼링될 때, chunk가 전송된 것으로 처리하지만 실제로는 여전히 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행 상태를 더욱 지속적으로 보고하고 있습니다.

   --max-upload-parts
      멀티파트 업로드의 최대 파트 수

      멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의하는 옵션입니다.

      서비스가 10,000개의 청크에 대한 AWS S3 사양을 지원하지 않는 경우 유용합니다.

      Rclone은 알려진 크기의 대용량 파일을 업로드할 때 청크 크기를 자동으로 증가시켜이러한 청크 수 제한을 준수합니다.

   --copy-cutoff
      멀티파트 복사로 전환할 임계값

      이보다 큰 파일을 서버 측에서 복사해야 하는 경우 이 크기로 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬 저장 안 함

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다.
      이는 데이터 무결성 검사에 유용하지만 대용량 파일의 업로드가 시작되기까지 오랜 지연을 초래할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일 경로

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 검색합니다.
      환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본 설정됩니다.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      이 변수는 해당 파일에서 사용할 프로필을 제어합니다.

      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 기본값으로 설정됩니다.

   --session-token
      AWS 세션 토큰

   --upload-concurrency
      멀티파트 업로드에 대한 동시성

      동시에 업로드되는 동일한 파일의 청크 수입니다.

      대용량 파일을 고속 링크로 작은 수의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우 이 값을 늘릴 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스 사용, false인 경우 가상 호스팅 스타일 액세스 사용

      true인 경우 rclone은 경로 스타일 액세스를 사용하고
      false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공업체(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값을 false로 설정해야 할 수 있으며, rclone은 제공 업체 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true인 경우 v2 인증 사용

      이 값이 false인 경우(기본값) rclone은 v4 인증을 사용합니다. 설정되면 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: pre Jewel/v10 CEPH.

   --list-chunk
      목록 청크 크기(각 ListObject S3 요청에 대한 응답 목록 크기)

      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있는 크기입니다.
      대부분의 서비스는 1000개 이상의 요청을 받아도 응답 목록을 1000개로 자르지만
      AWS S3에서는 이 값이 전역 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 증가시킬 수 있습니다.

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동(0)

      S3가 처음 시작되었을 때는 버킷의 객체를 열거하기 위해 ListObjects 호출만 가능했습니다.

      하지만 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 가지며 가능한 경우 사용해야 합니다.

      기본값인 0으로 설정하면 rclone은 제공업체 설정에 따라 호출할 목록 개체 방법을 추측합니다. 추측이 잘못된 경우 이곳에서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset

      일부 제공자는 목록을 URL 인코딩하는 것을 지원하며, 이를 사용하면 파일 이름에 제어 문자를 사용할 때 더 신뢰할 수 있습니다.
      이 값이 unset(기본값)인 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않음

      존재하는 버킷을 이미 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.

      또는 사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에는 버그 때문에 오류가 전달되었을 것입니다.

   --no-head
      업로드한 객체의 무결성을 확인하기 위해 HEAD 요청을 보내지 않음

      rclone은 기본적으로 PUT로 객체를 업로드한 후 200 OK 메시지를 받으면 올바르게 업로드되었다고 가정합니다.

      HEAD 요청을 보내지 않으면 다음을 가정합니다.

      - 업로드할 때의 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형)가 같음
      - 크기가 같음

      단일 파트 PUT의 응답에서 다음 항목을 읽습니다.

      - MD5SUM
      - 업로드된 일자

      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.

      알려지지 않은 길이의 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 업로드 실패가 더 잘 감지될 수 있으므로 정상 작동을 위해 권장되지 않습니다. 이 플래그로도 업로드 실패가 감지되지 않을 가능성이 매우 낮습니다.

   --no-head-object
      GET하기 전에 HEAD를 수행하지 않음

   --encoding
      백엔드에 대한 인코딩

      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding) 섹션을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀의 플러시 빈도

      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부

   --disable-http2
      S3 백엔드에서의 http2 사용 비활성화

      현재 S3 (특히 minio) 백엔드와 HTTP/2에 관련된 해결되지 않은 문제가 있습니다. s3 백엔드의
      HTTP/2는 기본적으로 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면 이 플래그가 제거될 것입니다.

      자세한 내용은 다음을 참조하십시오: [https://github.com/rclone/rclone/issues/4673](https://github.com/rclone/rclone/issues/4673), [https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 다운로드된 데이터에 대해 더 저렴한 외부 전송을 제공합니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 검증할지 여부

      true, false 또는 기본값(셋 되지 않음)으로 설정합니다.

   --use-presigned-request
      개별 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부

      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.

      rclone 버전 < 1.59는 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하며, 이 플래그를 true로 설정하면
      해당 기능을 다시 활성화할 수 있습니다. 이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉토리 목록에 이전 버전 포함

   --version-at
      지정된 시간에 시간대별 파일 버전 표시

      매개변수는 날짜 "2006-01-02", 날짜시간 "2006-01-02 15:04:05" 또는 그 시간 전의 기간인 "100d" 또는 "1h"입니다.

      이를 사용하는 경우 파일 쓰기 작업이 허용되지 않기 때문에 파일을 업로드하거나 삭제할 수 없습니다.

      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하세요.

   --decompress
      이 플래그가 설정되면 gzip으로 인코딩된 객체를 압축 해제함

      "Content-Encoding: gzip"으로 S3에 객체를 업로드하는 것이 가능합니다.
      일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.

      이 플래그가 설정되면 rclone은 받은 "Content-Encoding: gzip"의 파일을 압축 해제합니다.
      이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 해제됩니다.

   --might-gzip
      백엔드에서 객체에 gzip을 적용할 수 있는 경우 이 값을 설정하세요.

      일반적으로 제공자는 객체를 다운로드할 때 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은
      경우 다운로드 시 설정하지 않습니다.

      그러나 일부 제공자는 "Content-Encoding: gzip"(예: Cloudflare)로 업로드되지 않은 경우에도 gzip으로 객체를
      압축할 수 있습니다.

      이 값을 설정하고 rclone이 "Content-Encoding: gzip" 및 청크로 이동하는 객체를 다운로드하면 rclone은 객체를 실시간으로
      압축 해제합니다.

      이 값이 unset(기본값)인 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기 억제


옵션:
   --access-key-id 값      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl 값                버킷 생성 및 객체 저장 또는 복사 시 사용할 Canned ACL. [$ACL]
   --endpoint 값           Liara 객체 스토리지 API 엔드포인트. [$ENDPOINT]
   --env-auth              런타임(환경 변수 또는 env vars 또는 IAM)에서 AWS 자격 증명 가져오기 (default: false). [$ENV_AUTH]
   --help, -h              도움말 표시
   --secret-access-key 값  AWS 시크릿 액세스 키(비밀번호). [$SECRET_ACCESS_KEY]
   --storage-class 값      Liara에 새로운 객체를 저장할 때 사용할 스토리지 클래스. [$STORAGE_CLASS]

   Advanced

   --bucket-acl 값               버킷 생성 시 사용할 Canned ACL. [$BUCKET_ACL]
   --chunk-size 값               업로드에 사용할 청크 크기. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff 값              멀티파트 복사로 전환할 임계값. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                  이 플래그가 설정되면 gzip으로 인코딩된 객체를 압축 해제함. (default: false) [$DECOMPRESS]
   --disable-checksum            객체 메타데이터에 MD5 체크섬 저장 안 함. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2               S3 백엔드에서의 http2 사용 비활성화. (default: false) [$DISABLE_HTTP2]
   --download-url 값             다운로드를 위한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding 값                 백엔드에 대한 인코딩. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style            true인 경우 경로 스타일 액세스 사용, false인 경우 가상 호스팅 스타일 액세스 사용. (default: true) [$FORCE_PATH_STYLE]
   --list-chunk 값               목록 청크 크기(각 ListObject S3 요청에 대한 응답 목록 크기). (default: 1000) [$LIST_CHUNK]
   --list-url-encode 값          목록을 URL 인코딩할지 여부: true/false/unset. (default: "unset") [$LIST_URL_ENCODE]
   --list-version 값             사용할 ListObjects 버전: 1,2 or 0 for auto. (default: 0) [$LIST_VERSION]
   --max-upload-parts 값         멀티파트 업로드의 최대 파트 수. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time 값   내부 메모리 버퍼 풀의 플러시 빈도. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap        내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip 값               백엔드에서 객체에 gzip을 적용할 수 있는 경우 이 값을 설정하세요. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket             버킷의 존재 여부를 확인하거나 생성을 시도하지 않음. (default: false) [$NO_CHECK_BUCKET]
   --no-head                     업로드한 객체의 무결성을 확인하기 위해 HEAD 요청을 보내지 않음. (default: false) [$NO_HEAD]
   --no-head-object              GET하기 전에 HEAD를 수행하지 않음. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata          시스템 메타데이터의 설정 및 읽기 억제. (default: false) [$NO_SYSTEM_METADATA]
   --profile 값                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token 값            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file 값  공유 자격 증명 파일 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency 값       멀티파트 업로드에 대한 동시성. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff 값            청크 전송으로 전환할 크기 기준. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag 값       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부. (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request       개별 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부. (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                     true인 경우 v2 인증 사용. (default: false) [$V2_AUTH]
   --version-at 값               지정된 시간에 시간대별 파일 버전 표시. (default: "off") [$VERSION_AT]
   --versions                    디렉토리 목록에 이전 버전 포함. (default: false) [$VERSIONS]

   General

   --name 값  스토리지의 이름(자동 생성됨).(기본값: 자동 생성)
   --path 값  스토리지의 경로

```
{% endcode %}