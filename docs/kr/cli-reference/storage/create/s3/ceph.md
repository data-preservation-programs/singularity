# Ceph Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 ceph - Ceph Object Storage

사용법:
   singularity storage create s3 ceph [command options] [arguments...]

설명:
   --env-auth
      런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터)에서 AWS 자격 증명 가져오기.
      
      이 옵션은 access_key_id와 secret_access_key가 비어 있을 때만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력하십시오.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명 가져오기.

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두십시오.

   --secret-access-key
      AWS Secret Access Key(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두십시오.

   --region
      연결할 지역.
      
      S3 클론을 사용하고 지역을 가지지 않은 경우 비워 두십시오.

      예제:
         | <미설정>            | 확신이 없는 경우 이 옵션을 사용하십시오.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용하십시오.
         |                    | 예: 이전 Jewel/v10 CEPH.

   --endpoint
      S3 API의 엔드포인트.
      
      S3 클론을 사용할 때 필수입니다.

   --location-constraint
      지역 제한 - 지역과 일치해야합니다.
      
      확실하지 않은 경우 비워 두십시오. 버킷을 생성할 때만 사용됩니다.

   --acl
      버킷을 만들거나 개체를 저장하거나 복사할 때 사용되는 canned ACL.
      
      이 ACL은 개체를 만들거나 bucket_acl이 설정되지 않은 경우 해당 버킷을 만들 때 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      S3는 개체를 서버 측 복사할 때 ACL을 소스에서 복사하지 않고 새로 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 만들 때 사용되는 canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      bucket_acl이 설정되지 않은 경우만 버킷을 만들 때 사용됩니다.
      
      acl 및 bucket_acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.

      예제:
         | private            | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | 다른 사람은 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 액세스 권한을 가집니다.
         | public-read-write  | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스 권한을 가집니다.
         |                    | 버킷에서 이 것을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL 권한을 가집니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스 권한을 가집니다.

   --server-side-encryption
      S3에이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘.

      예제:
         | <미설정> | 없음
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘.

      예제:
         | <미설정> | 없음
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 키의 ARN을 제공해야합니다.

      예제:
         | <미설정>                 | 없음
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다.
      
      대신 --sse-customer-key-base64를 제공할 수 있습니다.

      예제:
         | <미설정> | 없음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터를 암호화/복호화하는 비밀 암호화 키를 Base64 형식으로 인코딩하여 제공해야합니다.
      
      대신 --sse-customer-key를 제공할 수 있습니다.

      예제:
         | <미설정> | 없음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 두면 sse_customer_key에서 자동으로 계산됩니다.

      예제:
         | <미설정> | 없음

   --upload-cutoff
      청크로 전환하는 크기.
      
      이보다 큰 파일은 chunk_size의 크기로 나누어져 업로드됩니다.
      최소값은 0이고 최대값은 5GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수없는 파일(예: "rclone rcat" 또는 "rclone mount" 또는 google
      사진 또는 google 문서에서 업로드 된 파일)은이 청크 크기를 사용하여
      멀티파트 업로드로 업로드됩니다.
      
      --s3-upload-concurrency"이 크기의 청크가
      메모리 당 전송에서인 이 크기의 청크를 버퍼로 사용합니다.
      
      고속 연결로 큰 파일을 전송하고 충분한 메모리가 있는 경우 이 값을
      늘리면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 파일의 크기가 큰 경우 청크 크기를 자동으로 늘립니다.
      
      크기를 알 수없는 파일은 구성된 대로
      청크 크기가 기본값 5 MiB이며 최대 10,000 청크를 가질 수 있으므로
      기본값에서 데이터 스트림 업로드 할 수있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을
      스트림업로드하려면 청크 크기를 늘려야합니다.
      
      청크 크기를 늘리면 "-P" 플래그로 표시되는 진행 상황,
      통계의 정확도가 낮아집니다. Rclone은 청크를 전송 한 것으로 처리합니다.
      버퍼로 AWS SDK에서 기록 할 때 실제로 업로드 될수도있는 송신 대상에 충분히 큰 청크 크기 및 진행 상황
      신뢰도.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수.
      
      이 옵션은 멀티파트 업로드시 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이것은 10,000 개 청크에 대한 AWS S3 사양을 지원하지 않는 서비스에 유용 할 수 있습니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드 할 때 청크 크기를 자동으로 늘릴
      기회를 제공합니다.
      

   --copy-cutoff
      청크 크기로 전환하기 위한 기준.
      
      서버 측 복사해야하는 이보다 큰 파일은
      이 크기로 청크별로 복사됩니다.
      
      최소값은 0이고 최대값은 5GiB입니다.

   --disable-checksum
      개체 메타 데이터에 MD5 체크섬을 저장하지 마십시오.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      개체의 메타 데이터에 추가하므로 대용량 파일을 업로드하기 전에 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면
      현재 사용자의 홈 디렉터리로 기본값을 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는
      해당 환경 변수도 설정되어 있지 않으면 "default"로 기본값을 사용합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      대용량 파일을 고속 링크로 업로드하는 경우에는 방대한 대역폭을 사용하지 못하는 경우에
      이 옵션을 늘리는 것이 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스트 스타일을 사용합니다.
      
      true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고
      false이면 rclone은 가상 경로 스타일을 사용합니다. [AWS S3
      문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      자세한 내용은 링크를 참조하십시오.
      
      일부 제공자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는이를 설정해야합니다.
      false - rclone은 제공자에 따라 자동으로 설정합니다.
      

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다.
      설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 사용하십시오. 예: 이전 Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(S3 요청별 응답 목록).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청한 것보다 더 많은 목록을 최대 1000 개의 개체로 자르지만 허용합니다.
      AWS S3에서는 전역 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)을 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 auto의 0입니다.
      
      S3가 처음 출시될 때 버킷의 개체를 나열하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하며 가능하다면 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 별로
      ListObjects 방법을 추측합니다. 잘못 추측하면
      여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩 할지 여부: true/false/unset
      
      일부 제공자는 목록을 URL 인코딩하고 파일의 제어 문자를 사용하는 경우가 있습니다. 사용 가능한 경우
      이렇게 사용하는 것이 파일 이름에 제어 문자를 사용할 때 더 안정적입니다. unset으로 설정하면
      rclone은 제공 업체 설정에 따라 적용할 내용을 선택하지만 여기서 rclone의 선택을 무시 할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 만들려고하지 마십시오.
      
      알려진 버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용 할 수 있습니다.
      
      또는 사용자가 버킷 생성 권한이없는 사용자 일 경우 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 정상적으로
     실행되지 않았습니다.
      

   --no-head
      업로드 된 개체의 무결성을 확인하기 위해 HEAD하지 않으십시오.
      
      rclone은 개체를 PUT 한 후 200 OK 메시지를 받으면 올바르게 업로드 된 것으로 간주합니다.
      
      특히 다음과 같다고 가정합니다.
      
      - 메타 데이터(모드 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드와 동일했음
      - 크기가 업로드와 동일함
      
      다음 요소를 단일 부 PUT 응답에서 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      길이가 알려지지 않은 소스 개체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그가 설정되면 올바르게 업로드될 때 rclone은 PUT을 사용하여 객체를 업로드 한 후에 200 OK
      메시지를 받으면 업로드 된 것으로 간주합니다.
      
      특히 다음과 같다고 가정합니다.
      
      - 메타 데이터(모드 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드와 동일했음
      - 크기가 업로드와 동일함
      
      단일 부 업로드의 경우 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      

   --no-head-object
      GET을 수행하기 전에 HEAD를 실행하지 마십시오.

   --encoding
      백엔드에 대한 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 간격.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)
      할당에는 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2간에 해결되지 않은 문제가 있습니다. 기본적으로 s3 백엔드는
      HTTP/2를 사용하지만 여기에서 비활성화 할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트입니다.
      이것은 보통 AWS S3를 통해 다운로드 된 데이터에 대해
      **안전하게** 되는 훨씬 낮은 다운로드 업 최대 속도로
      CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 공급자에 대한 기본값을 사용하려면 true, false 또는 설정하지 않습니다.
      

   --use-presigned-request
      단일 부 업로드에 대한 사전 서명된 요청 또는 PutObject 사용 여부
      
      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone < 1.59의 버전은 단일 부 객체를 업로드하기 위해 사전 서명된 요청을 사용하고이
      플래그를 true로 설정하면이 기능이 다시 활성화됩니다. 이 기능은 특수한 경우 또는 테스트
      외에는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전 포함 (default: false) [$VERSIONS]

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜 "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그 시간 전의 기간, 예를 들어 "100d" 또는 "1h" 일 수 있습니다.
      
      이를 사용하는 경우 파일 쓰기 작업은 허용되지 않으므로
      파일을 업로드하거나 삭제 할 수 없습니다.
      
      유효한 형식에 대해서는 [time 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩 된 개체를 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 개체를 업로드 할 수 있습니다. 일반적으로 rclone은 이러한
      파일을 압축 된 개체로 다운로드합니다.
      
      이 플래그가 설정 된 경우 rclone은 이러한 파일을 "Content-Encoding: gzip"으로 수신되는 대로 해제합니다. 이것은 rclone
      크기 및 해시를 확인할 수 없지만 파일 내용은 해제됩니다.
      

   --might-gzip
      백엔드에서 gzip 개체를 압축 할 수 있으면 이를 설정하십시오.
      
      일반적으로 제공자는 다운로드시 개체를 변경하지 않습니다. 만약
      `Content-Encoding: gzip`로 업로드되지 않은 객체는 다운로드시
      설정되지 않습니다
      
      그러나 일부 제공자는 `Content-Encoding: gzip`로 업로드되지 않은 객체 (예. Cloudflare)를 gzip으로 압축 할 수 있습니다.
      
      이의 표시로는 다음과 같은 오류를받을 때입니다
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그가 설정되고 rclone이 Content-Encoding: gzip를 설정하고 청크된 전송 인코딩으로 개체를 다운로드하면 rclone
      개체를 실시간으로 해제합니다.
      
      이를 설정하지 않는 경우(rclone의 기본값) rclone은 공급자 설명에 따라 적용할 내용을 선택하지만
      여기서 rclone의 선택을 무시 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 억제합니다. (기본값: false)

옵션:
   --access-key-id value           AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                     버킷을 만들거나 개체를 저장하거나 복사할 때 사용되는 canned ACL. [$ACL]
   --endpoint value                S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                      런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터)에서 AWS 자격 증명 가져오기. (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     지역 제한 - 지역과 일치해야합니다. [$LOCATION_CONSTRAINT]
   --region value                  연결할 지역. [$REGION]
   --secret-access-key value       AWS Secret Access Key(비밀번호). [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value          KMS ID를 사용하는 경우 키의 ARN을 제공해야합니다. [$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 전환하기 위한 기준. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩 된 개체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               개체 메타 데이터에 MD5 체크섬을 저장하지 마십시오. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               올바른 path style 액세스를 사용하려면 true를 입력하십시오. 가상 호스팅 스타일을 사용하려면 false를 입력하십시오.
                                    (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(S3 요청별 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩 할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 auto의 0입니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에 대한 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 간격. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip 개체를 압축 할 수 있으면 이를 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 만들려고하지 마십시오. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드 된 개체의 무결성을 확인하기 위해 HEAD하지 않으십시오. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 실행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 S3에이 개체를 저장할 때 사용되는 서버 측 암호화 알고리즘. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화/복호화하는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우 데이터를 암호화/복호화하는 비밀 암호화 키를 Base64 형식으로 인코딩하여 제공해야합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크로 전환하기 위한 크기. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부 업로드에 대한 사전 서명된 요청 또는 PutObject 사용 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다 (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 특정 시간으로 표시합니다. (기본값: "off") [$VERSION_AT]
   --decompress                     gzip으로 인코딩 된 개체를 해제합니다 (기본값: false) [$DECOMPRESS]
   --might-gzip                     백엔드에서 gzip 객체를 압축 할 수 있도록 설정합니다 (기본값: "unset") [$MIGHT_GZIP]
   --no-system-metadata 가         설정하지 않으면 시스템 메타데이터 설정 및 읽기를 억제하지 않습니다.

일반

   --name value  스토리지의 이름(기본 생성)
   --path value  스토리지의 경로

```
{% endcode %}