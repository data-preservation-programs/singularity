# StackPath Object Storage

{% code fullWidth="true" %}
```
명령:
   singularity storage create s3 stackpath - StackPath Object Storage

사용법:
   singularity storage create s3 stackpath [옵션] [인자...]

설명:
   --env-auth
      실행 중에 AWS 자격 증명 가져오기 (환경 변수 또는 env vars 또는 IAM의 EC2/ECS 메타 데이터).

      access_key_id와 secret_access_key가 비어 있을 때만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID.

      익명 액세스 또는 실행 중 자격 증명인 경우 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호).

      익명 액세스 또는 실행 중 자격 증명인 경우 비워 둡니다.

   --region
      연결할 지역.

      S3 클론을 사용하는 경우와 지역이 없는 경우 비워 둡니다.

      예제:
         | <unset>            | 알 수 없는 경우 사용합니다.
         |                    | v4 서명 및 빈 지역 사용
         | other-v2-signature | 작동하지 않을 때에만 사용합니다.
         |                    | 예: Jewel/v10 이전 CEPH.

   --endpoint
      StackPath Object Storage에 대한 엔드포인트.

      예제:
         | s3.us-east-2.stackpathstorage.com    | 미국 동부 엔드포인트
         | s3.us-west-1.stackpathstorage.com    | 미국 서부 엔드포인트
         | s3.eu-central-1.stackpathstorage.com | 유럽 엔드포인트

   --acl
      버킷을 만들고 객체를 저장하거나 복사할 때 사용되는 Canned ACL.

      이 ACL은 객체를 만들 때 및 bucket_acl이 설정되지 않은 경우 버킷을 만들 때 사용됩니다.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl에서 확인하십시오.

      S3는 서버 측 복사에서 객체 ACL을 복사하지 않고 새로운 ACL을 작성하는 것에 유의하세요.

      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.

   --bucket-acl
      버킷을 만들 때 사용되는 Canned ACL.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl에서 확인하십시오.

      bucket을 만들 때만 이 ACL이 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.

      "acl" 및 "bucket_acl"이 모두 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(개인)이 사용됩니다.

      예제:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자는 액세스 권한이 없습니다.(기본값)
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 액세스 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스 권한 부여.
         |                    | 이 권한은 버킷에 대해 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스 권한 부여.

   --upload-cutoff
      chunked 업로드로 전환하는 cutoff 값.

      이 값보다 큰 파일은 chunk_size로 분할하여 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)의 경우 이 청크 크기를 사용하여 multipart 업로드로 업로드됩니다.

      참고로, "--s3-upload-concurrency" 크기의 청크는 전송당 메모리에 버퍼링되어 있습니다.

      빠른 속도의 네트워크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시키면 전송 속도가 향상됩니다.

      rclone은 크기가 알려진 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 초과하지 않도록 합니다.

      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이고 최대 10,000개의 chunk가 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.

      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 감소합니다. Rclone은 AWS SDK로 버퍼에 청크를 보냈을 때 청크가 전송된 것으로 처리하지만 사실은 여전히 업로드 중일 수 있습니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 보고 값의 차이를 가져옵니다.


   --max-upload-parts
      multipart 업로드의 최대 파트 수.

      multipart 업로드를 수행할 때 사용할 최대 multipart 청크 수를 정의하는 옵션입니다.

      10,000개의 청크 규격을 지원하지 않는 서비스에 유용합니다.

      rclone은 파일 크기가 알려진 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 이러한 청크 수 제한 아래에 유지합니다.

   --copy-cutoff
      multipart 복사로 전환하는 cutoff 값.

      복사해야 하는 이 값보다 큰 파일은 이 크기의 청크로 복사됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만 큰 파일이 업로드되기 전에 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본값으로 현재 사용자의 홈 디렉토리를 사용합니다.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필.

      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용될 프로필을 제어합니다.

      비어있으면 환경 변수 "AWS_PROFILE"이나 설정되지 않은 경우 "default"를 기본값으로 사용합니다.

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      multipart 업로드에 대한 동시성.

      동일한 파일의 청크 수를 동시에 업로드합니다.

      고속 링크를 통해 대량의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우, 이 값을 증가시켜 전송 속도를 높일 수 있습니다.

   --force-path-style
      true로 설정하면 path 스타일 액세스를 사용하고 false로 설정하면 가상 호스팅 스타일을 사용합니다.

      true(기본값)인 경우 rclone은 path 스타일 액세스를 사용하고 false인 경우 rclone은 가상 path 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.

      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값을 false로 설정해야 합니다. rclone은 제공자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      v2 인증을 사용하려면 true로 설정합니다.

      false(기본값)인 경우 rclone은 v4 인증을 사용하고 설정되면 v2 인증을 사용합니다.

      v4 서명이 작동하지 않을 때만 사용하십시오, 예: Jewel/v10 이전 CEPH.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).

      이 옵션은 "MaxKeys", "max-items" 또는 "page-size"로 AWS S3 사양에서도 알려져 있습니다.
      대부분의 서비스는 요청한 객체를 1000개로 잘라냅니다.
      AWS S3에서는 이것이 전역 최대값이며 변경할 수 없습니다(https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html 참조).
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.

   --list-version
      사용할 ListObjects의 버전: 1, 2 또는 자동을 위해 0.

      S3가 처음 출시될 때는 버킷 내의 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.

      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하며 가능하면 사용해야 합니다.

      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못 추측하면 여기에서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 url 인코딩할 지 여부: true/false/unset

      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 사용 가능한 경우 파일 이름에 제어 문자를 사용할 때 신뢰할 수 있는 방법입니다. unset(기본값)으로 설정하면 rclone은 제공자 설정에 따라 선택할 것입니다.

   --no-check-bucket
      버킷의 존재를 확인하거나 생성을 시도하지 않습니다.

      건이 이미 존재한다는 것을 알고 있는 경우, rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.

      또한 사용자에게 버킷 작성 권한이없는 경우 필요할 수 있습니다.
      v1.52.0 이전에는 이는 버그로 인해 무시되었습니다.

   --no-head
      업로드된 객체의 HEAD를 확인하지 않습니다.

      rclone은 AEAD 이후 PUT으로 객체를 업로드한 후 200 OK 메시지를 받으면 예상대로 업로드된 것으로 간주합니다.

      특히 다음을 가정합니다.

      - 업로드된 메타데이터(수정 시간, 스토리지 클래스 및 컨텐츠 유형 포함)가 업로드된 것과 같음
      - 크기가 업로드된 것과 같음

      PUT로 단일 부의 응답에서 다음 항목을 읽습니다.

      - MD5SUM
      - 업로드된 날짜

      multipart 업로드인 경우 이러한 항목은 읽지 않습니다.

      알 수없는 길이의 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 잘못된 크기와 같은 업로드 실패 위험이 증가하므로 정상적인 작동에는 권장하지 않습니다. 실제로 업로드 실패가 감지되지 않을 확률은 매우 작습니다.

   --no-head-object
      객체를 가져올 때 HEAD를 GET하기 전에 수행하지 않습니다.

   --encoding
      백 엔드의 인코딩.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀의 플러시 빈도.

      추가 버퍼가 필요한 업로드(예: 다른 프로그램을 통한 업로드)은 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.

      현재 s3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. s3 백엔드의 기본 설정은 HTTP/2로 활성화되어 있지만 여기서 사용을 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드를위한 사용자 정의 엔드포인트.
      주로 AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 대출을 제공합니다.

   --use-multipart-etag
      ETag를 사용하여 multipart 업로드를 검증할지 여부

      true, false 또는 기본 옵션을 사용하려면 true, false 또는 기본 옵션으로 설정합니다.

   --use-presigned-request
      단일 부 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부

      이 값이 false이면 rclone은 단일 __부 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.

      rclone 버전 1.59보다 낮은 버전은 단일 부 객체를 업로드하기 위해 사전 서명된 요청을 사용하고 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      파일 버전을 지정된 시간에 있었던 버전으로 표시합니다.

      매개 변수는 날짜 "2006-01-02", 시간 "2006-01-02 15:04:05" 또는 그 이전 duration "100d" 또는 "1h"와 같은 형식이어야 합니다.

      이렇게 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.

      "Content-Encoding: gzip"이 설정된 상태로 S3에 객체를 업로드 할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.

      이 플래그가 설정되면 rclone은 "Content-Encoding: gzip"로 수신되는 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

   --might-gzip
      백엔드가 객체를 gzip으로 압축 할 수 있는 경우 이 플래그를 설정합니다.

      공급자는 일반적으로 객체를 다운로드할 때 변경하지 않습니다. 버킷이 "Content-Encoding: gzip"로 업로드되지 않으면 다운로드될 때 설정되지 않습니다.

      그러나 일부 공급자(예: Cloudflare)는 "Content-Encoding: gzip"로 업로드되지 않은 경우에도 객체를 gzip으로 압축 할 수 있습니다.

      이러한 경우 다음과 같은 오류가 발생합니다.

         ERROR corrupted on transfer: sizes differ NNN vs MMM

      이 플래그를 설정하면 rclone이 Content-Encoding: gzip이 설정되고 청크 전송 인코딩이 사용되는 객체를 다운로드 할 때 rclone은 객체를 실시간으로 압축 해제합니다.

      unset로 설정되어 있는 경우(rclone의 기본값) rclone은 제공자 설정에 따라 선택할 것입니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


옵션:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 만들고 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value           StackPath Object Storage의 엔드포인트. [$ENDPOINT]
   --env-auth                 실행 중에 AWS 자격 증명 가져오기 (환경 변수 또는 env vars 또는 IAM의 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키 (비밀번호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 만들 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              multipart 복사로 전환하는 cutoff 값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를위한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true로 설정하면 path 스타일 액세스를 사용하고 false로 설정하면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 url 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1,2 또는 자동을 위해 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         multipart 업로드의 최대 파트 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀의 플러시 빈도입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축 할 수 있는 경우 이 플래그를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 HEAD를 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 HEAD를 GET하기 전에 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       multipart 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            chunked 업로드로 전환하는 cutoff 값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       multipart 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 인증을 사용하려면 true로 설정합니다. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간에 있었던 버전으로 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}