# Cloudflare R2 Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 cloudflare - Cloudflare R2 Storage

사용법:
   singularity storage create s3 cloudflare [옵션] [인자...]

설명:
   --env-auth
      실행 시점에서 AWS 자격 증명 가져오기 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경 (env vars 또는 IAM)에서 AWS 자격 증명 가져오기.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 실행 시점 자격 증명인 경우 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호).
      
      익명 액세스 또는 실행 시점 자격 증명인 경우 비워 둡니다.

   --region
      연결할 지역.

      예시:
         | auto | R2 버킷은 지연 시간을 최소화하기 위해 Cloudflare의 데이터 센터에 자동으로 분산됩니다.

   --endpoint
      S3 API에 대한 엔드포인트.
      
      S3 클론을 사용하는 경우 필수입니다.

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되어 있지 않으면 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 "X-Amz-Acl:" 헤더는 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹이 읽기 액세스 권한을 갖습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 액세스 권한을 갖습니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AuthenticatedUsers 그룹이 읽기 액세스 권한을 갖습니다.

   --upload-cutoff
      청크 업로드로 전환하는 크기 제한.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 모르는 파일 (예: "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서로 업로드 됨)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      참고로 "--s3-upload-concurrency" 덩어리(크기는 이 크기로 설정됩니다)는 전송당 메모리에 버퍼링됩니다.
      
      대역폭이 높은 링크를 통해 대량의 파일을 전송하고 충분한 메모리가 있는 경우 이를 증가시켜 전송 속도를 높일 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 10,000개의 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      
      크기를 모르는 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크가 있을 수 있습니다.
      따라서 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 통계의 정확성이 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링 될 때 청크가 보낸 것으로
      처리하지만 실제로는 업로드 중인 경우도 있을 수 있습니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 보고 더 정확성이 떨어지기 때문에 진행률 통계를
      잘못 표시할 것이다.

   --max-upload-parts
      멀티파트 업로드의 최대 부분 개수.
      
      이 옵션은 멀티파트 업로드 시 사용할 멀티파트 청크의 최대 수를 정의합니다.
      
      AWS S3 명세의 10,000 청크를 지원하지 않는 서비스가 있는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 10,000개의 청크 제한을 유지하기 위해 청크 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 크기 제한.
      
      서버 측에서 복사해야 하는 이 크기보다 큰 파일은 이 크기로 청크 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬 저장하지 않음.
      
      일부 대형 파일의 업로드 시작에 긴 지연을 유발할 수 있는 MD5 체크섬을 rclone은 업로드 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가할 수
      있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로.
      
      env_auth=true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있는 경우 기본값은 현재 사용자의 홈 디렉토리입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필.
      
      env_auth=true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 그 파일에서 사용할 프로필을 제어합니다.
      
      비워 두면 환경 변수 "AWS_PROFILE" 또는 그 환경 변수가 설정되지 않은 경우 "default"로 기본값이 됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.
      
      파일의 동일한 청크를 동시에 업로드합니다.
      
      대역폭을 완전히 활용하지 못하는 경우에 대용량 파일 일부를 빠르게 전송하기 위해 이 값을 증가시킬 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스, false이면 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3
      문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는이 값을 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이 작업을
      수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정되어 있는 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예를 들어, 이전 버전의 Jewel/v10 CEPH 경우입니다.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 명세의 "MaxKeys", "max-items" 또는 "page-size"라고도 알려져 있습니다.
      대부분의 서비스는 요청한 이상의 응답 목록을 1000개로 자릅니다.
      AWS S3에서 전역 최대값이므로 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph는 "rgw list buckets max chunk" 옵션으로 이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전: 1,2 또는 자동 (0).
      
      S3가 처음 시작되었을 때는 버킷의 객체를 나열하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하며 가능하면 사용해야 합니다.
      
      기본값 0으로 설정하면 rclone은 공급자가 설정하는 것에 따라 어떤 객체 목록 방법을 호출할지 추측합니다.
      잘못 추측된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록 URL 인코딩 여부: true/false/unset
      
      일부 공급자는 목록을 URL로 인코딩하고 파일 이름에 컨트롤 문자를 사용할 때 이 기능을 지원합니다. 이 값이 unset(기본값)로 설정된 경우 rclone은 공급자
      설정에 따라 적용할 내용을 선택합니다. 그러나 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부 점검이나 생성을 시도하지 않음.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      버킷 생성 권한이 없는 사용자의 경우 필요할 수도 있습니다. 이전 버전의 v1.52.0 이전에는 버그로 인해 이 작업이 무음으로 전달되었습니다.
      

   --no-head
      업로드된 객체를 HEAD하여 무결성을 확인하지 않음.
      
      rclone은 기본적으로 PUT으로 객체를 업로드한 후 200 OK 메시지를 받으면 제대로 업로드된 것으로 가정합니다.
      
      특히 다음을 가정합니다:
      
      - 업로드된 메타데이터, 수정 시간, 저장소 클래스 및 콘텐츠 유형이 업로드된 대로이었음
      - 크기가 업로드된 대로
      

   --no-head-object
      GET하기 전에 HEAD를 수행하지 않을 경우 설정.

   --encoding
      백엔드용 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기.
      
      추가 버퍼를 필요로 하는 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거될 때를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에서 HTTP/2 사용 비활성화.
      
      현재 s3 (특히 minio) 백엔드에는 HTTP/2에 관한 문제가 있습니다. s3 백엔드의 경우 HTTP/2가 기본적으로 활성화되지만 여기에서 비활성화할 수
      있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참고: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트.
      AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 이그레스를 제공하므로
      일반적으로 이것은 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 검증할지 여부
      
      true, false 또는 기본값(공급자 설정에 따름) 중 하나여야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 값을 false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone < 1.59의 버전은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하고 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다.
      이 경우에는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전 포함.

   --version-at
      지정된 시간마다 파일 버전을 표시합니다.
      
      매개변수는 날짜인 "2006-01-02", 날짜 및 시간인 "2006-01-02 15:04:05" 또는 그 시간 전까지의 지속 시간인 "100d" 또는 "1h"일 수 있습니다.
      
      이 사용 시 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      있으면 gzip로 인코딩된 객체를 압축 해제합니다.
      
      S3로 "Content-Encoding: gzip"으로 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 객체를 "Content-Encoding: gzip"로 수신받을 때 압축 해제합니다. 따라서 rclone은 크기와 해시를 확인할 수
      없지만 파일 콘텐츠는 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 객체를 gzip으로 인코딩할 수 있는 경우 이 값을 설정합니다.
      
      일반적으로 공급자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체의 경우
      다운로드 시 설정되지 않습니다.
      
      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 상태에서도 객체를 gzip으로 압축할 수 있습니다 (예 : Cloudflare).
      
      이렇게 하면 다음과 같은 오류를 받습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip을 설정한 청크 전송 인코딩을 통해 객체를 다운로드하면 rclone은 객체를 실시간으로
      압축 해제합니다.
      
      unset으로 설정된 경우 (기본값) rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기 억제


옵션:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --endpoint value           S3 API에 대한 엔드포인트. [$ENDPOINT]
   --env-auth                 실행 시점에서 AWS 자격 증명 가져오기 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터 설정).(기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키 (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 크기 제한. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     있으면 gzip로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬 저장하지 않음. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 HTTP/2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩. (기본값: "슬래시, 잘못된 UTF8, 점") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스, false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL로 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1,2 또는 0으로 설정하면 자동 추측됨. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 개수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 객체를 gzip으로 인코딩할 수 있는 경우 이 값을 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부 점검이나 생성을 시도하지 않음. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체를 HEAD하여 무결성을 확인하지 않음. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET하기 전에 HEAD를 수행하지 않을 경우 설정. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 크기 제한. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간마다 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전 포함. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}