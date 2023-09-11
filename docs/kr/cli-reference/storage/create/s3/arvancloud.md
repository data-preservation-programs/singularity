# Arvan Cloud Object Storage (AOS)

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 arvancloud - Arvan Cloud Object Storage (AOS)

사용법:
   singularity storage create s3 arvancloud [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격증명 가져오기 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터).
      
      access_key_id 및 secret_access_key가 비어 있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 환경에서 AWS 자격증명 가져오기 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격증명을 위해 비워둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호).
      
      익명 액세스 또는 런타임 자격증명을 위해 비워둡니다.

   --endpoint
      Arvan Cloud Object Storage (AOS) API의 엔드포인트.

      예시:
         | s3.ir-thr-at1.arvanstorage.com | 기본 엔드포인트 - 확실하지 않은 경우 좋은 선택입니다.
         |                                | 이란 테헤란 (Asiatech)
         | s3.ir-tbz-sh1.arvanstorage.com | 이란 타브리즈 (Shahriar)

   --location-constraint
      위치 제약 조건 - 엔드포인트와 일치해야 합니다.
      
      버킷을 만들 때만 사용됩니다.

      예시:
         | ir-thr-at1 | 이란 테헤란 (Asiatech)
         | ir-tbz-sh1 | 이란 타브리즈 (Shahriar)

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 기본 ACL.
      
      이 ACL은 객체를 생성하는 데 사용되며, bucket_acl이 설정되지 않은 경우 버킷을 생성하는 데에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      기억하세요, S3는 소스의 ACL을 복사하지 않고 새로운 ACL을 작성하므로이 ACL은 서버 측 복사 객체에 적용됩니다.
      
      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

   --bucket-acl
      버킷을 만들 때 사용되는 기본 ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에는 READ 액세스 권한이 있습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에는 READ 및 WRITE 액세스 권한이 있습니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에는 READ 액세스 권한이 있습니다.

   --storage-class
      ArvanCloud에서 새로운 객체를 저장할 때 사용할 저장 클래스.

      예시:
         | STANDARD | 표준 저장 클래스

   --upload-cutoff
      청크 업로드로 전환하는 용량 제한.
      
      이보다 큰 파일은 chunk_size 단위로 청크 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 알 수없는 크기의 파일 (예 : "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)을 업로드 할 때, 이 청크 크기를 사용하여 멀티파트 업로드됩니다.
      
      "--s3-upload-concurrency" 스레드는 이 크기의 청크를 전송 당 메모리에 버퍼링합니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 전송할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한 이하로 유지합니다.
      
      알 수없는 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드 가능한 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가해야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확도가 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링되면 해당 청크를 전송 한 것으로 처리하지만 여전히 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 통보를 실제와 더 많이 다르게 할 수 있습니다.
      

   --max-upload-parts
      대용량 업로드에 사용되는 멀티파트 청크의 최대 개수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 멀티파트 청크의 최대 개수를 정의합니다.
      
      이는 서비스에서 AWS S3 사양의 10,000 청크를 지원하지 않는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 이 청크의 개수 제한을 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 용량 제한.
      
      이보다 큰 파일이 서버 측에서 복사되어 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 대용량 파일을 업로드하는 데 오랜 지연을 초래할 수 있지만 데이터 무결성 확인에 유용합니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로.
      
      env_auth가 true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필.
      
      env_auth가 true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워 두면 환경 변수 "AWS_PROFILE" 또는 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.
      
      동시에 업로드되는 동일한 파일의 청크 개수입니다.
      
      고속링크를 통해 대량의 큰 파일을 업로드하고 이 업로드에서 대역폭을 완전히 활용하지 못하는 경우 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고, false이면 가상 호스팅 스타일을 사용합니다.
      
      true (기본값)이면 rclone은 경로 스타일 액세스를 사용하고, false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자 (예 : AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false (기본값)이면 rclone은 v4 인증을 사용하며, 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 이전 CEPH.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청보다 많은 수의 응답 목록을 1000개로 잘라냅니다.
      AWS S3의 경우 이것은 전역 최대값이며, [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 0 (자동).
      
      S3가 처음 출시될 때 버킷의 객체를 나열하기 위해 ListObjects 호출 만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 제대로 추측하지 못하면 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 사용 가능한 경우 이것이 더 신뢰할 수 있습니다.
      이것이 unset으로 설정되면 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.
      
      또는 사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. 버전 1.52.0 이전에는 이 버그로 인해 에러 없이 통과되었습니다.
      

   --no-head
      청크 완정성을 확인하기 위해 업로드된 객체 HEAD하지 않습니다.
      
      rclone은 HEAD로 업로드 된 객체에 대한 200 OK 메시지를 받으면 올바르게 업로드된 것으로 가정합니다.
      
      특히 다음을 가정합니다:
      
      - 업로드하는 동안 메타데이터, 수정 시간, 저장 클래스 및 콘텐츠 유형이 업로드한 것과 같습니다.
      - 크기가 업로드된 것과 같습니다.
      
      단일 파트 PUT의 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      단, 멀티파트 업로드의 경우 이러한 항목은 읽히지 않습니다.
      
      알 수없는 길이의 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기와 같은 감지되지 않은 업로드 실패 가능성이 증가하므로 정상적인 작동에 권장되지 않습니다. 실제로 감지되지 않은 업로드 실패 가능성은 매우 작습니다.
      

   --no-head-object
      객체를 가져올 때 GET 앞에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 완전히 플러시되는 주기입니다.
      
      추가 버퍼가 필요한 업로드 (예: 멀티파트)는 메모리 풀을 사용하여 할당됩니다.
      이 옵션은 사용되지 않는 버퍼를 풀에서 정기적으로 제거하는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 S3 (특히 minio) 백엔드와 HTTP/2와 관련된 해결되지 않은 문제가 있습니다. S3 백엔드의 경우 HTTP/2가 기본적으로 사용되지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3가 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 전송을 제공하기 때문에 보통 CloudFront CDN URL로 설정합니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 provider의 기본값을 사용할지 여부입니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명 요청 또는 PutObject를 사용할지 여부
      
      false인 경우 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone의 버전 1.59 이전은 사전 서명 요청을 사용하여 단일 파트 객체를 업로드하며, 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이렇게 설정하는 경우 예외적인 상황이나 테스트만을 위해 필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 있었던 대로 파일 버전을 표시합니다.
      
      매개 변수는 날짜 "2006-01-02", 날짜시간 "2006-01-02 15:04:05" 또는 해당 시간 이전의 지속 시간 "100d" 또는 "1h" 여야 합니다.
      
      주의: 이를 사용하는 동안 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 압축된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"로 수신한 대로 압축을 해제합니다. 이렇게 되면 rclone은 크기와 해시를 확인하지 못하지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 gzip 객체를 사용할 수 있는 경우 이를 설정합니다.
      
      일반적으로 공급자는 객체를 다운로드할 때 객체를 변경하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 객체라면 다운로드할 때도 설정되지 않습니다.
      
      그러나 일부 공급자는 개체를 `Content-Encoding: gzip`로 압축하지 않았더라도 개체를 gzip으로 압축 할 수 있습니다 (예 : Cloudflare).
      
      이 경우 다음과 같은 오류를 받을 수 있습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 `Content-Encoding: gzip`와 청크전송 인코딩이 설정된 객체를 다운로드하면 rclone은 객체를 실시간으로 압축 해제합니다.
      
      이를 unset으로 설정하면 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사 시 사용되는 기본 ACL. [$ACL]
   --endpoint value             Arvan Cloud Object Storage (AOS) API의 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격증명 가져오기 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 조건 - 엔드포인트와 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]
   --storage-class value        ArvanCloud에서 새로운 객체를 저장할 때 사용할 저장 클래스. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 기본 ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 복사하는 용량 제한. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip 으로 압축된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고, false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1, 2 또는 0 (자동). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         대용량 업로드에 사용되는 멀티파트 청크의 최대 개수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 완전히 플러시되는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip 객체를 사용할 수 있는 경우 이를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        청크 완정성을 확인하기 위해 업로드된 객체 HEAD하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져올 때 GET 앞에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 용량 제한. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 있었던 대로 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}