# Netease Object Storage (NOS)

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 netease - Netease Object Storage (NOS)

사용법:
   singularity storage update s3 netease [command options] <name|id>

DESCRIPTION:
   --env-auth
      런타임에서 AWS 인증 자격 증명(환경 변수 또는 환경(EC2/ECS 메타 데이터)에 따라)을 얻습니다.
      
      access_key_id와 secret_access_key가 비어 있을 경우에만 적용됩니다.

      예시:
         | false | AWS 자격 증명을 다음 단계에서 입력합니다.
         | true  | 환경 (환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS Access Key ID입니다.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(암호)입니다.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --region
      연결할 지역입니다.
      
      S3 클론을 사용하고 지역이 없는 경우 비워 둡니다.

      예시:
         | <unset>            | 확실하지 않은 경우 사용하십시오.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용하십시오.
         |                    | 예: 이전 Jewel/v10 CEPH.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필수입니다.

   --location-constraint
      위치 제약 조건 - 지역과 일치하도록 설정해야 합니다.
      
      확실하지 않으면 비워두세요. 버킷을 만들 때만 사용됩니다.

   --acl
      버킷 및 객체 저장 또는 복사 시 사용하는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않았을 때도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      이 ACL은 S3이 객체를 서버 측 복사할 때 적용됩니다.
      S3은 소스의 ACL을 복사하지 않고 새로 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고
      기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL입니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      생성 버킷에만 적용되는 ACL입니다. 설정되지 않은 경우, "acl"이 대신 사용됩니다.
      
      "acl"와 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl:
      헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예시:
         | private            | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람들에게는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 READ 액세스를 얻습니다.
         | public-read-write  | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 READ 및 WRITE 액세스를 얻습니다.
         |                    | 버킷에 이를 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹은 READ 액세스를 얻습니다.

   --upload-cutoff
      청크 업로드로 전환하기 위한 각각의 파일의 임계값입니다.
      
      이 값을 초과하는 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 모르는 파일
      ("rclone rcat"이나 "rclone mount"나 google
      사진이나 google 문서에서 업로드된 파일 등)을 업로드할 때,
      이 청크 크기를 사용하여 multipart 업로드가 수행됩니다.

      사이즈의 파일에 대해서는 per transfer의 메모리를 “--s3-upload-concurrency” 청크 단위로 버퍼링하므로,
      대용량 파일을 고속 링크로 전송하고 충분한 메모리가 있는 경우
      청크 크기를 증가시켜 전송 속도를 높이는 것이 좋습니다.

      rclone은 10,000개의 청크 제한을 초과하지 않도록
      큰 크기의 알려진 파일을 업로드할 때 자동으로 청크 크기를 증가시킵니다.

      알려지지 않은 크기의 파일은 구성된
      청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대
      10,000개의 청크가 가능하므로, 기본적으로 스트림 업로드할 수 있는
      파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림으로 업로드하려면
      청크 크기를 증가시켜야 합니다.

      청크 크기를 늘릴수록 "-P" 플래그로 표시되는 진행 상태의
      정확성이 낮아집니다. rclone은 청크가 AWS SDK에 의해 버퍼링되는 경우에
      청크가 전송된 것으로 처리하지만, 실제로는 아직 업로드 중인 경우가
      있을 수 있습니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼와
      진행률 표시의 정확성 간의 더 큰 차이를 의미합니다.
      

   --max-upload-parts
      multipart 업로드의 최대 부분 수입니다.
      
      이 옵션은 multipart 업로드를 할 때 사용할 최대 multipart 청크 수를 정의합니다.
      
      10,000개의 청크에 대한 AWS S3 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      알려진 크기의 대용량 파일을 업로드할 때 Rclone은
      10,000개의 청크 제한을 초과하지 않도록 청크 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.
      
      서버 측에서 복사해야 할 이 임계값을 초과하는 파일은
      이 사이즈 단위의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만
      큰 파일을 업로드하는 데 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 매개변수 값이 비어 있으면
      현재 사용자의 홈 디렉터리를 기본값으로 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true 인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워 두면 환경 변수 "AWS_PROFILE"이나
      설정되지 않은 경우 "default"를 기본값으로 사용합니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동시에 업로드되는 동일한 파일 청크 수입니다.
      
      고속 링크를 통해 대량의 대용량 파일을 업로드하고
      이 업로드가 대역폭을 완전히 이용하지 못하는 경우 이 값을 증가시키면
      전송 속도가 향상될 수 있습니다.

   --force-path-style
      true인 경우 path 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true(기본값)인 경우 rclone은 path 스타일 액세스를 사용하고
      false인 경우 rclone은 가상 path 스타일을 사용합니다. [the AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      자세한 내용은 를 참조하십시오.
      
      일부 프로바이더(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 다음과 같은 설정에 따라
      false로 설정해야 합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false(기본값)인 경우 rclone은 v4 인증을 사용합니다.
      설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 사용하십시오, 예: 이전 Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(S3 요청별 응답 목록입니다).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys" 또는 "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 1000개 이상을 요청하더라도 응답 목록을 1000개로 줄입니다.
      AWS S3에서는 이것이 전체 최대값이므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동(0).
      
      S3가 처음 출시되었을 때 버킷의 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하며 가능한 경우에는 사용해야 합니다.
      
      기본값인 0으로 설정할 경우 rclone은 공급 업체에 맞게
      호출할 목록 개체 메서드를 추측합니다. 잘못된 추측일 경우,
      여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급 업체는 목록의 URL 인코딩을 지원하며, 이렇게 사용하는 것이
      파일 이름에 제어 문자를 사용할 때 더 안정적입니다. 이를 설정하면
      rclone은 크기와 해시를 확인하기 위해 공급 업체 설정에 따라
      선택한 대로 적용합니다. (기본값 "unset").
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는
      트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 갖고 있지 않을 경우에도 필요할 수 있습니다.
      
      이전 v1.52.0까지는 버그로 인해 이것이 방문하는 버그로 인해 이것은
      음성으로 전달되었습니다.
      

   --no-head
      업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다.
      
      rclone은 가장자리에 대해 PUT로 객체를 업로드한 후 200 OK 메시지를 받으면
      업로드가 제대로 되었다고 가정합니다.
      
      특히 다음을 가정합니다:
      
      - 업로드된 모든파일의 metadata, modtime, storage class 및 content type가 업로드된 것과 같은 속성임을 가정합니다.
      - 크기가 업로드된 것과 같은 속성임을 가정합니다.
      
      단일 파트 PUT의 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드에서는 이러한 항목을 읽지 않습니다.
      
      길이를 알 수 없는 소스 객체가 업로드되는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기를 포함한
      감지되지 않은 업로드 실패의 가능성이 증가하므로
      일반 작업에는 권장되지 않습니다. 실제로는 업로드 실패의
      가능성은 이 플래그를 사용하지 않아도 극히 적습니다.
      

   --no-head-object
      GET하는 동안 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 비울 것인지를 제어합니다.
      
      추가 버퍼(예: multipart가 필요한 업로드)를 위해 영역이 필요한
      업로드는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 S3 (특히 minio) 백엔드와 HTTP/2에 관한
      문제가 해결되지 않은 상태입니다. S3 백엔드에서는 기본적으로
      HTTP/2가 사용되지만 이 옵션을 사용하여 비활성화할 수 있습니다.
      문제가 해결되면이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는
      CloudFront 네트워크를 통해 다운로드 된 데이터에 대해
      더 저렴한 이그레스를 제공합니다.

   --use-multipart-etag
      확인을 위해 multipart 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 제공자에 대한 기본값을 사용해야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 대해 presigned request 또는 PutObject를 사용할지 여부
      
      false인 경우 rclone은 AWS SDK의 PutObject를 사용하여
      객체를 업로드합니다.
      
      rclone의 버전 1.59 미만에서는 기본값으로 단일
      파트 객체를 업로드하려면 presigned requests를 사용하고
      이 플래그를 true로 설정하면 이 기능을 다시 활성화합니다.
      이는 특수한 경우나 테스트를 제외하고 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그 이전의 기간, 예를 들어 "100d" 또는 "1h"일 수 있습니다.
      
      이를 사용하는 동안 파일 쓰기 작업은 허용되지 않으므로
      파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option) 를 참조하십시오.
      

   --decompress
      Gzip으로 인코딩된 객체의 압축을 해제합니다.
      
      S3로 "Content-Encoding: gzip"으로 객체를 업로드할 수 있습니다. 보통 rclone은
      이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을
      "Content-Encoding: gzip"으로 수신되는 대로 해제합니다. 이는 rclone이
      크기와 해시를 확인할 수 없지만 파일 콘텐츠가 해제됩니다.
      

   --might-gzip
      백엔드에서 개체에 gzip을 적용할 수 있는 경우 설정합니다.
      
      일반적으로 공급자가 다운로드될 때 객체를 수정하지 않습니다. 제어 문자가 존재하지 않는 경우도
      "Content-Encoding: gzip"로 업로드되지 않았으면 다운로드시
      그렇게 설정되지 않습니다.
      
      그러나 일부 공급 업체는 다음과 같은 설정에 관계없이 객체를 gzip으로 압축할 수 있습니다
      예: 클라우드플레어).
      
      이 플래그를 설정하고 rclone이
      Content-Encoding: gzip과 청크 전송 인코딩이 설정된 개체를 다운로드하면 rclone
      객체를 실시간으로 압축합니다.
      
      unset인 경우(기본값) rclone은 공급 업체 설정에
      따라 선택하나, 이곳에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value        AWS Access Key ID입니다. [$ACCESS_KEY_ID]
   --acl value                  버킷 및 객체의 생성 또는 복사 시 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value             S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 인증 자격 증명(환경 변수 또는 환경(EC2/ECS 메타 데이터)에 따라)을 얻습니다. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 조건 - 지역과 일치하도록 설정해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역입니다. [$REGION]
   --secret-access-key value    AWS Secret Access Key(암호)입니다. [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     Gzip으로 인코딩된 객체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 path 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(S3 요청별 응답 목록입니다). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동(0). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         multipart 업로드의 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 비울 것인지를 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 개체에 gzip을 적용할 수 있는 경우 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET하는 동안 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 각각의 파일의 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 multipart 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 presigned request 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}