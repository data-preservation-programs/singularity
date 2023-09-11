# 기타 S3 호환 공급자

{% code fullWidth="true" %}
```
이름:
  singularity 스토리지 업데이트 s3 other - 기타 S3 호환 공급자

사용법:
  singularity 스토리지 업데이트 s3 other [command options] <name|id>

설명:
  --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예:
         | false | 가급적 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경 변수 (env vars 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

  --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

  --secret-access-key
      AWS 비밀 액세스 키 (암호).
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

  --region
      연결할 리전.
      
      S3 복제본을 사용하고 리전이 없는 경우 비워둡니다.

      예:
         | <unset>            | 확실하지 않을 때 사용합니다.
         |                    | v4 서명 및 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용합니다.
         |                    | 예: Jewel/v10 CEPH 이전.

  --endpoint
      S3 API의 엔드포인트.
      
      S3 복제본을 사용하는 경우 필수입니다.

  --location-constraint
      위치 제약조건 - 리전과 일치해야 합니다.
      
      확실하지 않을 경우 비워둡니다. 버킷을 생성할 때만 사용됩니다.

  --acl
      버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL.
      
      이 ACL은 객체를 생성하고 bucket_acl이 설정되어 있지 않은 경우 버킷을 생성할 때도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3 서버 간 복사에서 이 ACL이 적용됩니다.
      S3가 소스의 ACL을 복사하지 않고 새로 쓰기 때문에 지정됩니다.
      
      ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며
      기본값 (private)이 사용됩니다.
      

  --bucket-acl
      버킷을 만들 때 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      ACL 및 버킷 ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며
      기본값 (private)이 사용됩니다.
      

      예:
         | private            | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람에게 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 액세스를 얻습니다.
         | public-read-write  | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻습니다.
         |                    | 버킷에서 이 작업을 수행하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스를 얻습니다.

  --upload-cutoff
      청크 업로드로 전환할 파일 크기의 임계값.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

  --chunk-size
      업로드에 사용될 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 (예: "rclone rcat" 또는 "rclone mount" 또는 google
      사진 또는 google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 다중 부분 업로드로 업로드됩니다.
      
      주의: "--s3-upload-concurrency" 크기의 청크가 전송당 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 높여 전송 속도를 높일 수 있습니다.
      
      큰 파일의 경우 rclone은 10,000개의 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      크기를 알 수 없는 파일은 설정된 chunk_size로 업로드됩니다.
      기본 청크 크기가 5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 크게 할수록 "-P" 플래그와 함께 표시되는 진행 상태의 정확성이 감소합니다. rclone은
      AWS SDK의 버퍼에 청크가 버퍼링되면 청크를 전송한 것으로 처리하지만 실제로는 아직 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 더 정확하지 않은 진행률 보고를 의미합니다.
      

  --max-upload-parts
      멀티파트 업로드에 사용할 최대 청크 수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 수 있는 최대 멀티파트 청크 수를 정의합니다.
      
      서비스가 AWS S3 10,000 청크 사양을 지원하지 않는 경우에 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 유지합니다.
      

  --copy-cutoff
      멀티파트 복사로 전환할 파일 크기의 임계값.
      
      서버 측에서 복사해야 하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

  --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는
      데이터 무결성 검사에 좋지만 큰 파일을 업로드하기 시작할 때 큰 지연을 일으킬 수 있습니다.

  --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면
      현재 사용자의 홈 디렉토리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

  --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는
      비어 있으면 "default"로 기본 설정됩니다.
      

  --session-token
      AWS 세션 토큰.

  --upload-concurrency
      멀티파트 업로드의 동시성.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      고속 링크를 통해 대량의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우에는
      이를 늘리는 것이 전송 속도를 높일 수 있습니다.

  --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일 액세스를 사용합니다.
      
      true (기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고
      false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3
      문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는
      false로 설정되어야 합니다 - rclone은 이를 제공자 설정을 기반으로 자동으로 수행합니다.

  --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false인 경우 (기본값) rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.
      
      이것은 v4 서명이 작동하지 않을 때만 사용합니다. 예: Jewel/v10 CEPH 이전.

  --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청된 것보다 많은 항목을 요청해도 응답 목록을 1000개로 잘라냅니다.
      AWS S3에서는 이것이 전역 최대이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

  --list-version
      사용할 ListObjects의 버전: 1,2 또는 자동이면 0.
      
      S3가 처음 출시되었을 때 버킷의 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정된 경우 rclone은 제공자 설정에 따라 어떤 목록 개체 방법을 호출할지 추측할 것입니다. 올바른 추측을 하지 못하면
      여기서 수동으로 설정할 수 있습니다.
      

  --list-url-encode
      목록을 URL 인코딩할 지 여부: true/false/unset
      
      일부 공급자는 목록을 URL 인코딩하고 이를 사용할 수 있는 경우 파일 이름에 제어 문자를 사용할 때 이 방법이 더 신뢰성이 있습니다. 
      이 값이 unset (기본값)로 설정되어 있으면 rclone은 제공자 설정에 따라 어떤 값을 적용할지 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

  --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. 버전 1.52.0 이전에서는
      버그로 인해 이 항목이 정상적으로 전달되었을 것입니다.
      

  --no-head
      업로드된 객체를 체크하기 위해 HEAD를 실행하지 않습니다.
      
      rclone이 PUT를 사용하여 객체를 업로드한 후 200 OK 메시지를 받으면 올바르게 업로드된 것으로 가정합니다.
      
      특히, 다음과 같이 가정합니다.
      
      - 메타데이터 (수정 시간, 저장 클래스 및 콘텐츠 유형 포함)는 업로드한 대로입니다.
      - 크기는 업로드한 대로입니다.
      
      다중 부분 업로드의 단일 부분 PUT에 대한 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      다중 부분 업로드의 경우 이러한 항목은 읽히지 않습니다.
      
      길이를 알 수 없는 출처 객체가 업로드되면 rclone은 HEAD 요청을 실행합니다.
      
      이 플래그를 설정하면 업로드 실패가 감지되지 않을 위험이 높아집니다.
      특히 올바르지 않은 크기의 경우이므로 정상적인 운영에는 권장되지 않습니다. 실제로 업로드 실패가 발생할 가능성은
      매우 적습니다.
      

  --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다.

  --encoding
      백엔드에 대한 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

  --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 메모리 풀을 사용하여 할당을 수행합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

  --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

  --disable-http2
      S3 백엔드에서 http2 사용 비활성화.
      
      S3 (구체적으로 minio) 백엔드와 HTTP/2의 문제가 현재 해결되지 않았습니다. s3 백엔드는 기본적으로
      HTTP/2가 활성화되어 있지만 여기서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

  --download-url
      다운로드에 대한 사용자 정의 엔드포인트.
      AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우
      더 저렴한 전송이 제공됩니다.

  --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값은 true, false 또는 제공자의 기본값을 사용하도록 설정해야 합니다.
      

  --use-presigned-request
      단일 부분 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      false인 경우 rclone은 PutObject를 사용하여 개체를 업로드합니다.
      
      rclone < 1.59 버전은 단일 부분 개체를 업로드하기 위해 사전 서명된 요청을 사용하며 이 플래그를 true로 설정하면
      해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트 이외에는 필요하지 않아야 합니다.
      

  --versions
      디렉터리 목록에 이전 버전 포함 (기본값: false).

  --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", 날짜 시간 "2006-01-02
      15:04:05" 또는 그 이후를 나타내는 기간 "100d" 또는 "1h"일 수 있습니다.
      
      이를 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      사용 가능한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

  --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"으로 설정된 상태로 S3에 객체를 업로드하는 것도 가능합니다. 일반적으로 rclone은
      이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 수신할 때 "Content-Encoding: gzip"로 이러한 파일을 압축 해제합니다. 이것은 rclone이
      크기와 해시를 확인할 수 없지만 파일 콘텐츠는 압축 해제됩니다.
      

  --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는지 여부입니다.
      
      일반적으로 제공자는 다운로드 시 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체의 경우
      다운로드 시 설정되지 않습니다.
      
      그러나 일부 공급자는 "Content-Encoding: gzip"으로 업로드되지 않은 객체를 gzip으로 압축할 수 있습니다(예: Cloudflare).
      
      set으로 설정하면 rclone이 Content-Encoding: gzip이 설정되고 청크 전송 인코딩인 객체를 다운로드하면 rclone은 객체를 압축 해제합니다.
      
      unset(기본값)로 설정되어 있으면 rclone은 제공자 설정에 따라 적용할 값을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

  --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value             S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   정보 표시
   --location-constraint value  위치 제약조건 - 리전과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전. [$REGION]
   --secret-access-key value    AWS 비밀 액세스 키 (암호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 만들 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용될 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환할 파일 크기의 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 HTTP/2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할 지 여부: true/false/unset. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전: 1,2 또는 자동이면 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에 사용할 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는지 여부입니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체를 체크하기 위해 HEAD를 실행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환할 파일 크기의 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전 포함 (기본값: false) [$VERSIONS]

```
{% endcode %}