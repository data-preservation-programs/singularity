# 다른 S3 호환 공급자

{% code fullWidth="true" %}
```
NAME:
   singularity storage create s3 other - 다른 S3 호환 공급자

USAGE:
   singularity storage create s3 other [command options] [arguments...]

DESCRIPTION:
   --env-auth
      AWS 자격 증명을 런타임에서 가져옵니다 (환경 변수 또는 환경 변수가 없을 경우 EC2/ECS 메타 데이터).

      access_key_id와 secret_access_key이 비어 있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경 (환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호)입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워둡니다.

   --region
      연결할 리전입니다.

      S3 복제본을 사용하고 리전이 없다면 비워둡니다.

      예시:
         | <unset>            | 확실하지 않을 때 사용하세요.
         |                    | v4 서명 및 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때 사용합니다.
         |                    | 예: Jewel/v10 CEPH 이전 버전.

   --endpoint
      S3 API의 엔드포인트입니다.

      S3 복제본을 사용하는 경우 필수입니다.

   --location-constraint
      리전과 일치하는 위치 제약 조건입니다.

      확실하지 않으면 비워둡니다. 버킷 생성 시에만 사용됩니다.

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 저장소 클래스 액세스 제어(Canned ACL)입니다.

      이 ACL은 객체 생성에도 사용되며, bucket_acl이 설정되지 않았을 경우 버킷 생성에도 사용됩니다.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.

      참고 : S3는 서버 측 객체 복사 시 ACL을 복사하지 않고 새로 작성합니다.

      acl이 빈 문자열이면 X-Amz-Acl 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 저장소 클래스 액세스 제어(Canned ACL)입니다.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.

      bucket_acl이 설정되지 않으면 "acl" 대신 사용됩니다.

      "acl"과 "bucket_acl"이 빈 문자열이면 X-Amz-Acl 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사용자에게는 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹은 읽기 권한을 받습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 받습니다.
         |                    | 버킷에서 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AuthenticatedUsers 그룹은 읽기 권한을 받습니다.

   --upload-cutoff
      청크 업로드로 전환하는 기준입니다.

      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 큰 파일이나 크기가 알 수 없는 파일("rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일)은 이 청크 크기를 사용하여 multipart 업로드를 사용하여 업로드됩니다.
      
      주의 : "--s3-upload-concurrency" 크기의 청크는 전송당 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 크기를 늘리면 전송 속도가 빨라집니다.
      
      rclone은 10,000개의 청크 제한을 유지하기 위해 알려진 크기의 대형 파일을 업로드할 때 자동으로 청크 크기를 증가시킵니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로,
      기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그로 표시되는 진행 상태 통계의 정확성이 낮아집니다. rclone은 AWS SDK에서 버퍼링되는 청크를
      보낸 것으로 처리하며, 사실은 아직 업로드 중인 경우입니다. 큰 청크 크기는 큰 AWS SDK 버퍼와 진행률 표시의 참값과의
      편차가 증가합니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용하는 최대 파트 수입니다.

      멀티파트 업로드 시 이 옵션은 사용할 멀티파트 청크 수를 정의합니다.
      
      이 옵션은 AWS S3 사양인 10,000개 청크를 지원하지 않는 서비스에 유용합니다.
      
      rclone은 알려진 크기의 대형 파일을 업로드할 때 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로의 전환용 기준입니다.

      서버 측 복사가 필요한 이보다 큰 파일은 이 크기로 조각별로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 이렇게 하면
      대형 파일의 업로드가 시작되기 전까지 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" env 변수를 찾습니다. env 값이 비어 있으면
      현재 사용자의 홈 디렉토리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      이 값이 비어있으면 "AWS_PROFILE" 환경 변수 or "default" 환경 변수를 기본값으로 설정합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.

      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대용량 파일을 고속 링크를 통해 소수 개의 큰 파일을 업로드하고 이 업로드가 대역폭을 완전히 활용하지 않는 경우,
      이 값을 증가시켜 전송 속도를 높일 수 있습니다.

   --force-path-style
      true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅 스타일을 사용합니다.

      true (기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false인 경우 rclone은 가상 경로 스타일을 사용합니다.
      자세한 내용은 [the AWS S3 docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공자(AWS, Aliyun OSS, Netease COS, 또는 Tencent COS 등)는 false로 설정해야 합니다. rclone은 제공자 설정에
      기반하여 이 설정을 자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.

      이 값이 false(기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      이는 v4 시그니처가 작동하지 않을 때 사용하세요. 예: Jewel/v10 CEPH 이전 버전.

   --list-chunk
      리스트 청크의 크기(각 ListObject S3 요청의 응답 목록)입니다.

      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청된 개수와 관계없이 응답 목록을 최대 1000개로 잘라냅니다.
      AWS S3에서는 이것이 글로벌 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 더 큰 값으로 변경할 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 auto에 대한 0.

      S3가 처음에 출시되었을 때 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하고 가능하면 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못 추측하면 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      리스팅을 URL 인코딩할 것인지 여부: true/false/unset
      
      일부 제공자는 리스팅을 URL 인코딩하며, 파일 이름에서 제어 문자를 사용할 때 이것이 가능하면 신뢰성이 더 높습니다. unset으로 설정된 경우
      rclone은 공급자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않습니다.

      버킷이 이미 존재하는 것을 알고있을 때, rclone이 실행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.
      
      버킷 생성 권한이 없는 사용자를 사용하는 경우 필요할 수 있습니다.
      이전 버전(1.52.0 이전)에서는 버그 때문에 이전 버전은 정상적으로 지나갔을 것입니다.
      

   --no-head
      업로드한 객체의 정합성을 확인하기 위해 HEAD를 하지 않습니다.

      rclone이 PUT으로 객체를 업로드한 후 200 OK 메시지를 받으면 올바르게 업로드된 것으로 가정합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터, 수정 시간, 저장 클래스 및 콘텐츠 유형은 업로드한 것과 동일합니다.
      - 크기는 업로드한 것과 동일합니다.
      
      다음의 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      단일 부 PUT의 응답에 대해 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      multipart 업로드의 경우 이 항목을 읽지 않습니다.
      
      크기를 알 수 없는 원본 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 증가하며, 특히 크기가 잘못된 경우입니다. 따라서 정상 운영에는 권장되지 않습니다.
      실제로 이 플래그를 설정하면 업로드 실패 가능성은 거의 없습니다.
      

   --no-head-object
      GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.

      자세한 내용은 [개요의 인코딩](/overview/#encoding) 섹션을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.

      추가 버퍼가 필요한 업로드(예:Multipart)는 할당에 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에 대한 http2 사용 비활성화.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 관한 미해결된 문제가 있습니다. 기본적으로 s3 백엔드에는 HTTP/2가 enabled되어
      있지만 이곳에서 비활성화할 수 있습니다. 문제가 해결되면 이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      이는 보통 AWS S3가 클라우드 프론트 CDN URL에 설정됩니다.
      AWS S3는 클라우드 프론트 네트워크를 통해 데이터를 다운로드 할 때 더 저렴한 이중타입을 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 기본값을 사용할지 설정합니다.
      

   --use-presigned-request
      단일 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 < 1.59은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하며, 이 플래그를 true로 설정하면
      해당 기능을 다시 활성화합니다. 이는 특정한 경우나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함시킵니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 "2006-01-02", datetime "2006-01-02 15:04:05" 또는 그렇게 오래된 기간(예: "100d" 또는 "1h")이어야 합니다.
      
      이 옵션을 사용하면 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식은 [time 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      설정하면 gzip으로 인코딩된 개체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축 해제하여
      "Content-Encoding: gzip"으로 수신받습니다. 이렇게 하면 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 gzip으로 압축된 객체인 경우 설정하십시오.
      
      일반적으로 공급자는 개체를 다운로드할 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지
      않은 개체는 다운로드할 때 설정되지 않을 것입니다.
      
      그러나 공급자 중 일부는 개체를 "Content-Encoding: gzip"로 업로드하지 않았더라도 개체를 gzip할 수 있습니다(Cloudflare 등).
      
      이러한 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 chunked 전송 인코딩 및 Content-Encoding: gzip으로 개체를 다운로드하면 rclone은
      실시간으로 개체를 압축 해제합니다.
      
      unset으로 설정하면 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사 시 사용되는 저장소 클래스 액세스 제어(Canned ACL). [$ACL]
   --endpoint value             S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                   AWS 자격 증명을 런타임에서 가져옵니다 (환경 변수 또는 환경 변수가 없을 경우 EC2/ECS 메타 데이터). (default: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  리전과 일치하는 위치 제약 조건. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전. [$REGION]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷 생성 시 사용되는 저장소 클래스 액세스 제어(Canned ACL). [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로의 전환용 기준. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정하면 gzip으로 인코딩된 개체를 압축 해제합니다. (default: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (default: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기(각 ListObject S3 요청의 응답 목록). (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          리스팅을 URL 인코딩할 것인지 여부. (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 auto에 대한 0. (default: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용하는 최대 파트 수. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 gzip으로 압축된 객체인 경우 설정하십시오. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성하지 않습니다. (default: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 정합성을 확인하기 위해 HEAD를 하지 않습니다. (default: false) [$NO_HEAD]
   --no-head-object                 GET 전에 HEAD를 수행하지 않습니다. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 기준. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부. (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 사전 서명된 요청 또는 PutObject를 사용할지 여부 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (default: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (default: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함시킵니다. (default: false) [$VERSIONS]

   General

   --name value  스토리지의 이름(기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}