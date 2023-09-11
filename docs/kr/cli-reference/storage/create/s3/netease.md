# Netease Object Storage (NOS)

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 netease - Netease Object Storage (NOS)

사용법:
   singularity storage create s3 netease [command options] [arguments...]

설명:
   --env-auth
      실행 시간(AWS 환경 변수 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옵니다.
      env_vars 또는 IAM에서 자격 증명을 가져오는 것입니다.
      
      access_key_id 및 secret_access_key가 비어있을 경우만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경 변수(env_vars 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 실행 시간 자격 증명인 경우 비워 두세요.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호)입니다.
      
      익명 액세스 또는 실행 시간 자격 증명인 경우 비워 두세요.

   --region
      연결할 지역입니다.
      
      S3 클론을 사용하고 지역이 없는 경우 비워 두세요.

      예제:
         | <설정되지 않음> | 확실하지 않은 경우 사용하세요.
         |                 | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않는 경우에만 사용하세요.
         |                 | 예: 이전 버전의 CEPH(Jewel/v10).

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필수입니다.

   --location-constraint
      위치 제약 조건 - 지역과 일치해야 합니다.
      
      확실하지 않으면 비워 두세요. 버킷 생성에만 사용됩니다.

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다.
      
      이 ACL은 객체를 만들 때 사용되며 bucket_acl이 설정되지 않은 경우에도 버킷을 만들 때 사용됩니다.
      
      자세한 내용은 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 S3가 서버 간에 객체를 복사할 때 ACL을 복사하지 않고 새롭게 쓰기 때문에 적용됩니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL입니다.
      
      자세한 내용은 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | 다른 사람에게는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 액세스가 있습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스가 있습니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스가 있습니다.

   --upload-cutoff
      청크로 전환하는 데 사용되는 파일의 크기 임계값입니다.
      
      이 값보다 큰 파일은 청크 크기로 업로드됩니다 (chunk_size).
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기가 알 수 없는 파일(예: "rclone rcat" 또는 "rclone mount"로 업로드되거나 Google 사진 또는 Google 문서에서 업로드된 파일)은 이 청크 크기를 사용하여 다른 부분의 업로드로 업로드됩니다.

      알아 두셔야 할 점은 "--s3-upload-concurrency" 청크 크기만큼의 청크가 전송별로 메모리에 버퍼링되어 있습니다.

      높은 속도로 큰 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 늘리면 전송 속도를 높일 수 있습니다.

      Rclone은 지정된 크기의 큰 파일을 업로드할 때 10,000개의 청크 제한에 대해 작은 청크 크기를 자동으로 증가시킵니다.

      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000의 저장소가 있을 수 있으므로 기본적으로 스트림으로 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림으로 업로드하려면 chunk_size를 늘려야 합니다.

      청크 크기를 늘리면 "-P" 플래그와 함께 표시된 진행률 통계의 정확성이 감소합니다. Rclone은 AWS SDK가 버퍼에 있는 청크를 보내면 전송 완료된 것으로 처리하지만 실제 전송 중인 경우가 있을 수 있습니다. 청크 크기가 크면 AWS SDK 버퍼가 크고 진행률 통계가 실제 값과 다를 수 있습니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 이음새 수의 최대값입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이 옵션은 서비스가 10,000개의 멀티파트 청크를 지원하지 않는 경우에 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 청크 크기를 증가시켜 이 채널의 최대 청크 수를 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.
      
      이 값보다 큰 서버 사이드 복사가 필요한 파일은 이 크기로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 큰 도움이 되지만 대용량 파일을 업로드할 때 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면 기본적으로 현재 사용자의 홈 디렉토리를 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워 두면 환경 변수 "AWS_PROFILE" 또는 미설정된 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      속도가 높은 링크를 통해 작은 수의 대용량 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 이용하지 못하는 경우 이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 패스 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.
      
      true(기본값)이면 rclone은 패스 스타일 액세스를 사용하고 false이면 rclone은 가상 패스 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS, 또는 Tencent COS 등)는 이 값을 false로 설정해야 합니다. rclone은 제공자 설정에 기반하여 자동으로 설정합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: 이전 버전의 CEPH(버전 Jewel/v10).

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기)입니다.
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청된 것보다 더 많은 객체를 요청해도 응답 목록을 1000개로 자릅니다.
      AWS S3에는 전역 최대치가 있으며 이를 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동(0).
      
      S3가 처음 시작될 때는 버킷의 객체를 나열하는 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 높은 성능을 제공하므로 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 어떤 ListObjects 방법을 호출할지 추측합니다. 추측이 잘못된 경우 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 URL 인코딩 목록을 지원하며 사용 가능할 때 파일 이름에 제어 문자를 사용하는 경우 이 방법이 더 안정적입니다. unset으로 설정된 경우 rclone은 제공자 설정에 따라 일부용법을 적용하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않으려면 설정하세요.
      
      버킷이 이미 존재하는 경우 rclone의 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. 버전 1.52.0 이전에는 버그 때문에 이것이 은밀하게 통과되었을 것입니다.
      

   --no-head
      업로드한 객체의 INTEGRITY를 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone이 PUT로 객체를 업로드한 후 200 OK 메시지를 받으면 제대로 업로드된 것으로 간주됩니다. 이 플래그를 설정하면 rclone은 PUT로 객체를 업로드한 후 200 OK 메시지를 수신한 경우 올바르게 업로드됐다고 가정합니다.
      
      특히 다음과 같다고 가정합니다:
      
      - 업로드된 것처럼 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)가 업로드된 대로였음
      - 업로드된 크기는 업로드된 대로였음
      
      단일 부분의 PUT에 대한 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽을 수 없습니다.
      
      길이를 알 수 없는 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패가 탐지되지 않을 확률이 높아집니다. 특히 잘못된 크기이므로 정상적인 운영에는 권장되지 않습니다. 실제로엔 업로드 실패가 탐지되지 않을 확률은 매우 작습니다.
      

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않으려면 설정하세요.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 비울지를 제어합니다.
      
      추가 버퍼(예: 멀티파트)가 필요한 업로드는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 관련된 해결되지 않은 문제가 있습니다. S3 백엔드는 기본적으로 HTTP/2를 사용하지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      대부분의 경우 저렴한 다운로드를 위해 AWS S3를 통해 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      ETag를 멀티파트 업로드에서 사용하여 검증할지 여부
      
      true, false 또는 unset이어야 합니다. 이는 공급자에 대한 기본값을 사용합니다.
      

   --use-presigned-request
      단일 부분 업로드에 서명된 요청 또는 PutObject를 사용할지 여부
      
      false이면 rclone은 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      rclone 버전 < 1.59은 단일 부분 객체를 업로드하기 위해 서명된 요청을 사용하고이 플래그를 true로 설정하면 해당 기능을 다시 활성화합니다. 이는 예외적인 상황이나 테스트를 위해서만 필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그보다 오래된 기간(예: "100d" 또는 "1h")일 수 있습니다.
      
      이 옵션을 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      [시간 옵션 문서](/docs/#time-option)를 참조하여 유효한 형식을 확인하세요.
      

   --decompress
      gzip으로 압축된 객체를 압축 해제해야 하는 경우 설정하세요.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 받은 대로 "Content-Encoding: gzip"로 이 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용을 압축 해제하는 것을 의미합니다.
      

   --might-gzip
      백엔드에서 객체를 gzip으로 압축할 수 있는 경우 설정하세요.
      
      일반적으로 제공자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 제공자는 gzip을 사용하여 객체를 압축할 수 있습니다(예: Cloudflare).
      
      이 경우에는
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      와 같은 오류가 발생할 수 있습니다.
      
      이 플래그를 설정하고 rclone이 "Content-Encoding: gzip"이 설정된 상태의 객체를 청크 전송 인코딩을 사용하여 다운로드하는 경우 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본값)로 설정된 경우 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value        AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value             S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   실행 시간(AWS 환경 변수 또는 EC2/ECS 메타 데이터)에서 AWS 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 조건 - 지역과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역입니다. [$REGION]
   --secret-access-key value    AWS 비밀 액세스 키(비밀번호)입니다. [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 압축된 객체를 압축 해제해야 하는 경우 설정합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 패스 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동(0). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 이음새 수의 최대값입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 비워지는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 객체를 gzip으로 압축할 수 있는 경우 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 INTEGRITY를 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 전에 HEAD를 수행하지 않으려면 설정하세요. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크로 전환하는 데 사용되는 파일의 크기 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          싱글 파트 업로드에 대해 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   일반

   --name value  저장소의 이름(자동 생성)
   --path value  저장소의 경로

```
{% endcode %}