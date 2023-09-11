# Cloudflare R2 스토리지

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 cloudflare - Cloudflare R2 스토리지

사용법:
   singularity storage update s3 cloudflare [command options] <이름|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key가 비어있는 경우에만 적용됩니다.

      예제:
         | false | AWS 자격 증명을 다음 단계에서 입력하세요.
         | true  | 환경 (환경 변수 또는 IAM)에서 AWS 자격 증명 가져오기.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --region
      연결할 지역.

      예제:
         | auto | 저렴한 지연 시간을 위해 R2 버킷은 Cloudflare 데이터 센터에 자동으로 분산됩니다.

   --endpoint
      S3 API를 사용할 때 엔드포인트.
      
      S3 클론을 사용할 때 필요합니다.

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 비어 있으면 X-Amz-Acl 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사용자에게 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 액세스가 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스가 부여됩니다.
         |                    | 일반적으로 버킷에서이 권한을 부여하는 것은 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스가 부여됩니다.

   --upload-cutoff
      청크 업로드로 전환하는 파일의 크기 기준.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기가 알 수없는 파일 (예 : "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드 된 파일)은이 청크 크기를 사용하여
      이 청크 크기를 사용하여 멀티파트 업로드를 사용하여 업로드됩니다.
      
      참고 : "--s3-upload-concurrency"는 메모리 당 이 크기의 청크를 버퍼링합니다.
      
      고속 링크를 통해 큰 파일을 전송하고 충분한 메모리가 있다면이 값을 높이면 전송 속도가 빨라집니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드 할 때 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      알려진 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000 청크가 있을 수 있으므로
      스트림 업로드 할 수있는 파일의 기본적 최대 크기는 48 GiB입니다.  더 큰 파일을 스트림으로 업로드하려면 청크 크기를 늘려야합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 통계의 정확도가 낮아집니다. Rclone은 청크를
      AWS SDK가 버퍼링 할 때 전송 된 것으로 간주합니다. 그러나 실제로 업로드 중 일 수도 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행 상황으로
      실제와 다른 보고를 만듭니다.

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수.
      
      이 옵션은 멀티파트 업로드시 사용할 부분(chunk)의 최대 수를 정의합니다.
      
      이는 10,000 청크의 AWS S3 사양을 지원하지 않는 서비스에 유용 할 수 있습니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드 할 때 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      서버 측 복사에 전환하는 파일의 기준 크기.
      
      이보다 큰 파일이 서버 측으로 복사되어야하는 경우이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      개체 메타 데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체에 추가합니다. 이는
      데이터 무결성을 확인하는데 유용하지만 대용량 파일의 경우 업로드 시작까지 오랜 지연을 야기할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE"
      환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는
      그 환경 변수가 설정되지 않은 경우 "default"로 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.
      
      이는 동시에 업로드되는 동일한 파일 청크의 수입니다.
      
      고속 링크에서 작은 수의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지
      못한다면 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true (기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고,
      false 인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) 문서를 참조하세요.
      
      일부 공급 업체 (예 : AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이 값을 설정하여
      false로 설정해야합니다. Rclone은 공급자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.
      
      이 값이 false (기본값)이면 rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만이 옵션을 사용하세요. 예를 들어 v10 CEPH로부터 보내기전에.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items"또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청보다 많은 수를 요청해도 응답 목록을 1000 개로 자르지만
      AWS S3에서는 전역 최대 값으로 1000 개로 이를 변경 할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여이 값을 높일 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전 : 1,2 또는 자동으로 0.
      
      S3가 처음 시작되었을 때는 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공되었습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이는
      훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 사용해야 할 목록 개체 방법을 추측합니다.
      잘못 추측하면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩 할지 여부 : true/false/unset
      
      일부 공급 업체는 URL 인코딩 목록을 지원하며 가능한 경우 파일의 제어 문자를
      사용할 때 이것이 더 신뢰할 수 있습니다. 이 값이 설정되지 않은 경우 (기본값) rclone은
      제공자 설정에 따라 적용할 내용을 선택하지만 rclone의 선택을 여기에서 무시 할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않으려면 설정하세요.
      
      버킷이 이미 존재한다는 것을 알고있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 이 옵션을 사용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. 이전 버전의 v1.52.0까지 발생한 버그로 인해 이로 인해 전체 검사가 수행되지 않았습니다.
      

   --no-head
      업로드 된 개체의 무결성을 확인하기 위해 HEAD를 수행하지 않도록 설정하세요.
      
      rclone은 가능한 적은 수의 트랜잭션을 수행하려고 할 때 유용합니다.
      
      설정하면 rclone은 PUT으로 개체를 업로드한 후에 200 OK 메시지를 받으면
      제대로 업로드되었다고 가정합니다.
      
      특히 다음과 같은 경우에 다음을 가정합니다.
      
      - 업로드 될 때의 메타 데이터, 변경 시간, 저장 클래스 및 콘텐츠 유형이 업로드와 동일합니다.
      - 크기가 업로드와 동일합니다.
      
      단일 부분의 PUT에 대한 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드 날짜
      
      이러한 항목은 멀티파트 업로드에 대해 읽히지 않습니다.
      
      길이를 알 수없는 소스 개체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패가 검출되지 않을 수 있는 가능성이 증가하며,
      특히 잘못된 크기 미스매치의 경우 일반 작업에는 권장되지 않습니다. 실제로 이 플래그로 인한
      업로드 실패 확률은 매우 낮습니다.

   --no-head-object
      GET 시 HEAD를 수행하기 전에 HEAD를 수행하지 않도록 설정합니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 encoding 섹션](/overview/#encoding)에서 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 플러시 할지 제어합니다.
      
      추가 버퍼가 필요한 업로드 (예 : multipart)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에서 http2의 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. s3 백엔드에 대해 HTTP/2는 기본적으로 활성화되어 있지만
      여기에서 사용을 비활성화 할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트.
      이는 일반적으로 AWS S3에서 퍼져 나가는 데이터에 대해 더 저렴한 egress를 제공하는 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에 ETag를 사용할지 여부
      
      이 값은 true, false 또는 공급 업체의 기본값을 사용하도록 설정되지 않은 경우여야합니다.
      

   --use-presigned-request
      단일 부분 업로드를위한 미리 서명 된 요청 또는 PutObject를 사용할지 여부
      
      false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone < 1.59 버전은 단일 부분 객체를 업로드하기 위해 미리 서명 된 요청을 사용하고이 플래그를 true로 설정하면
      해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않을 것입니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜, "2006-01-02" 날짜 시간 "2006-01-02
      15:04:05" 또는 그만큼 오래된 기간, 예 : "100d" 또는 "1h" 여야합니다.
      
      이를 사용하는 경우 파일 쓰기 작업을 수행할 수 없으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식은[time 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩 된 개체를 압축 해제합니다.
      
      "Content-Encoding: gzip"가 설정된 상태에서 S3에 개체를 업로드 할 수 있습니다. 보통 rclone은
      이러한 파일을 압축 된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 수신되는대로 "Content-Encoding: gzip"가 있는 이러한 파일을
      압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 콘텐츠가 압축 풀린 상태입니다.
      

   --might-gzip
      백엔드에서 개체를 gzip으로 압축 할 수 있는 경우 이를 설정합니다.
      
      일반적으로 제공 업체는 개체가 다운로드 될 때 개체를 변경하지 않습니다. `Content-Encoding: gzip`로
      업로드되지 않은 개체에는 설정되지 않습니다.
      
      그러나 일부 제공 업체는 `Content-Encoding: gzip`로 업로드하지 않았더라도 개체를 gzip으로 압축 할 수 있습니다.
      (예 : Cloudflare).
      
      이로 인해 다음과 같은 오류가 표시됩니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 `Content-Encoding: gzip` 및 청크 전송 인코딩으로 개체를 다운로드하면
      rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset으로 설정되어있는 경우 (기본값) rclone은
      제공자 설정에 따라 적용 할 내용을 선택하지만 여기서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --endpoint value           S3 API를 위한 엔드포인트. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역. [$REGION]
   --secret-access-key value  AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 기준 크기. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩 된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               개체 메타 데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2의 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩 할지 여부 : true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전 : 1,2 or 0 for auto. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 플러시 할지 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 개체를 gzip으로 압축 할 수 있는 경우 이를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 만들지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        HEAD를 수행하지 않도록 설정합니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않도록 설정합니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 파일의 크기 기준. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드를위한 미리 서명 된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}