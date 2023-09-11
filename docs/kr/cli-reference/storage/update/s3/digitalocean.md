# DigitalOcean Spaces

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 digitalocean - DigitalOcean Spaces

사용법:
   singularity storage update s3 digitalocean [옵션] <이름|ID>

설명:
   --env-auth
      실행중에 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없으면 EC2/ECS 메타데이터에서 가져옵니다).
      
      access_key_id 및 secret_access_key가 비어있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경 (env vars 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 실행중인 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (패스워드).
      
      익명 액세스 또는 실행중인 자격 증명을 위해 비워 둡니다.

   --region
      연결할 리전입니다.
      
      S3 클론을 사용하고 지역이 없는 경우 비워 둡니다.

      예시:
         | <unset>            | 확실하지 않은 경우 사용합니다.
         |                    | v4 서명 및 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 경우에만 사용합니다.
         |                    | 예 : Jewel/v10 CEPH 이전.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필요합니다.

      예시:
         | syd1.digitaloceanspaces.com | DigitalOcean Spaces Sydney 1
         | sfo3.digitaloceanspaces.com | DigitalOcean Spaces San Francisco 3
         | fra1.digitaloceanspaces.com | DigitalOcean Spaces Frankfurt 1
         | nyc3.digitaloceanspaces.com | DigitalOcean Spaces New York 3
         | ams3.digitaloceanspaces.com | DigitalOcean Spaces Amsterdam 3
         | sgp1.digitaloceanspaces.com | DigitalOcean Spaces Singapore 1

   --location-constraint
      위치 제약 조건 - 리전과 일치해야 합니다.
      
      확실하지 않은 경우 비워 둡니다. 버킷을 생성 할 때 사용됩니다.

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다.
      
      이 ACL은 객체를 생성하기 위해 사용되며, "bucket_acl"이 설정되지 않은 경우 버킷 생성을 위해서도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      서버 측 복사로 객체를 복사 할 때 이 ACL이 적용됩니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 생성 할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 모두 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본(비공개)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사람에게 액세스 권한이 없습니다 (기본).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 액세스가 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스가 부여됩니다.
         |                    | 버킷에서 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스가 부여됩니다.

   --upload-cutoff
      청크 업로드로 전환하는 임계치입니다.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수없는 파일 (예 : "rclone rcat" 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드 된 파일)을 업로드하는 경우, 이 청크 크기를 사용하여 멀티파트 업로드됩니다.
      
      참고로 "--s3-upload-concurrency" 이 크기의 청크는 전송당 메모리에 버퍼링됩니다.
      
      고속 링크로 대량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시켜 전송 속도를 높일 수 있습니다.
      
      Rclone은 알려진 큰 파일을 업로드 할 때 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본적인 청크 크기는 5MiB이며 최대 10,000 청크까지 있을 수 있으므로 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확도가 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링 될 때 청크를 보낸 것으로 처리하지만 실제 업로드 중일 수 있습니다. 청크 크기가 크면 AWS SDK 버퍼 및 진행 상태 보고가 실제로 다를 수 있습니다.
      

   --max-upload-parts
      멀티파트 업로드에 사용되는 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드 시 사용되는 최대 멀티파트 청크 수를 정의합니다.
      
      서비스가 10,000 청크의 AWS S3 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 큰 파일을 업로드 할 때 10,000 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계치입니다.
      
      서버 간에 복사해야 하는 이 임계치보다 큰 파일은 이 크기로 청크화하여 복사됩니다.
      
      최소값은 0이고 최대값은 5GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력 파일의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만 큰 파일을 업로드하기 시작하기 전에 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본 값은 현재 사용자의 홈 디렉토리입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      없으면 환경 변수 "AWS_PROFILE" 또는 "default"를 기본 값으로 사용합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대역폭을 완전히 활용하지 못하는 고속 링크에서 큰 파일의 작은 수를 업로드하고 있는 경우 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고, false이면 가상 호스팅 스타일을 사용합니다.
      
      true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고, false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급 업체 (예 : AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 false로 설정해야 합니다. rclone은 공급 업체 설정을 기반으로 자동으로 이 작업을 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명 작동 시에만 사용하세요. 예 : Jewel/v10 CEPH 이전.

   --list-chunk
      '- 최대 버킷 초과' 오류로 인해 S3 요청마다 응답 목록의 크기를 나타냅니다.
      
      이 옵션은 AWS S3 사양의 'MaxKeys', 'max-items' 또는 'page-size'로 알려진 것입니다.
      대부분의 서비스는 요청보다 더 많이 요청해도 응답 목록을 1000개로 잘라냅니다. AWS S3에서 이것은 전역 최댓값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동 설정을 위해 0.
      
      S3가 처음 출시될 때 버킷 내 개체를 나열하기 위해 ListObjects 호출만 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 방식은 훨씬 높은 성능을 제공하며 가능하면 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 list objects 방법을 공급 업체가 설정한대로 추측합니다. 추측이 잘못되면 여기에서 수동으로 설정 할 수 있습니다.
      

   --list-url-encode
      목록을 URL로 인코딩할지 여부: true/false/unset
      
      일부 공급 업체는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며 사용 가능한 경우 이 방법이 더 신뢰할 수 있습니다. unset으로 설정된 경우 (기본값) rclone은 공급 업체 설정에 따라 적용하는 값을 선택합니다. 하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      알다시피 버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      또한 사용자가 버킷 생성 권한을 가지고 있지 않으면 필요할 수 있습니다. 버전 1.52.0 이전에는 오류로 인해 무시되었을 것입니다.
      

   --no-head
      업로드된 객체의 정합성을 확인하기 위해 HEAD 요청을 보내지 않습니다.
      
      rclone은 일반적으로 PUT으로 객체를 업로드 한 후 200 OK 메시지를 받으면 올바르게 업로드된 것으로 가정합니다.
      
      특히 rclone은 다음 항목을 단일 부분 PUT의 응답에서 읽습니다:
      
      - MD5 체크섬
      - 업로드된 날짜
      
      멀티파트 업로드는 이러한 항목을 읽지 않습니다.
      
      길이를 모르는 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기를 포함한 업로드 실패의 가능성이 높아지므로 정상적인 작동에는 권장되지 않습니다. 실제로 이 플래그로 인한 업로드 실패의 가능성은 매우 낮습니다.
      

   --no-head-object
      객체를 가져 오기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [관련 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 시간입니다.
      
      부가 버퍼가 필요한 업로드는 메모리 풀을 사용하여 할당을 수행합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 조정합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. s3 백엔드의 HTTP/2 기능은 기본적으로 활성화되지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      이는 보통 AWS S3는 CloudFront 네트워크를 통해 내려받을 때 더 저렴한 대출이기 때문에 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 지원되는 경우 공백으로 설정하세요.
      

   --use-presigned-request
      단일 부분 업로드를위한 서명 된 요청 또는 PutObject를 사용할지 여부
      
      false로 설정하면 rclone은 AWS SDK에서 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone의 버전 < 1.59은 단일 부분 객체를 업로드하기 위해 서명 된 요청을 사용하며 이 플래그를 true로 설정하면이 기능이 다시 활성화됩니다. 특수한 경우나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함시킵니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", datetime "2006-01-02 15:04:05" 또는 그 전 시간에 대한 기간, 예를 들어 "100d" 또는 "1h"가 있습니다.
      
      이를 사용할 때는 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 수신 시 "Content-Encoding: gzip"로 이러한 파일을 압축 해제합니다. 즉, rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 gzip으로 객체를 압축 할 수 있는 경우 설정하세요.
      
      일반적으로 공급 업체는 객체를 다운로드 할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"으로 업로드되지 않은 객체의 경우 다운로드시 설정되지 않습니다.
      
      그러나 일부 공급 업체는 gzip으로 압축한 객체를 제공 할 수 있습니다 (예 : Cloudflare).
      
      이것의 증상은 다음과 같은 오류를받는 경우입니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 이러한 요구 사항이 충족되는 상태에서 Content-Encoding: gzip가 설정되고 청크 전송 인코딩을 수신하면 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset으로 설정하면(기본값) rclone은 공급 업체 설정에 따라 적용하는 값을 선택합니다. 하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 작성을 억제합니다


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value             S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                   실행중에 AWS 자격 증명 가져오기 (환경 변수 또는 env vars가 없으면 EC2/ECS 메타데이터에서 가져옵니다). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말을 표시합니다
   --location-constraint value  위치 제약 조건 - 리전과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전입니다. [$REGION]
   --secret-access-key value    AWS Secret Access Key (패스워드). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계치. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고, false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               '- 최대 버킷 초과' 오류로 인해 S3 요청마다 응답 목록의 크기입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL로 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동 설정을 위해 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에 사용되는 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 gzip으로 객체를 압축 할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 정합성을 확인하기 위해 HEAD 요청을 보내지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져 오기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 작성 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계치. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드를위한 서명 된 요청 또는 PutObject 사용 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전 표시. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함시킵니다. (기본값: false) [$VERSIONS]

``` 
{% endcode %}