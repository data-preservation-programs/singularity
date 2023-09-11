# Scaleway Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 scaleway - Scaleway Object Storage

사용법:
   singularity storage update s3 scaleway [command options] <name|id>

설명:
   --env-auth
      권한을 실행 시간(runtime)에서 가져옵니다 (환경 변수 또는 env vars나 IAM이 없다면 EC2/ECS 메타 데이터에서).
      
      access_key_id 및 secret_access_key이 비어 있을 경우만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 환경에서 AWS 자격증명을 가져옵니다 (env vars 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스나 실행 시간(runtime) 자격증명을 위해 비워둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호)입니다.
      
      익명 액세스나 실행 시간(runtime) 자격증명을 위해 비워둡니다.

   --region
      연결할 지역입니다.

      예제:
         | nl-ams | 네덜란드 암스테르담
         | fr-par | 프랑스 파리
         | pl-waw | 폴란드 바르샤바

   --endpoint
      Scaleway Object Storage에 대한 엔드포인트입니다.

      예제:
         | s3.nl-ams.scw.cloud | 암스테르담 엔드포인트
         | s3.fr-par.scw.cloud | 파리 엔드포인트
         | s3.pl-waw.scw.cloud | 바르샤바 엔드포인트

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다.
      
      이 ACL은 객체를 생성할 때와 bucket_acl이 설정되어 있지 않을 때에도 사용됩니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      S3는 복사 과정에서 ACL을 복사하지 않고 새로운 ACL을 작성하므로 이 ACL은 서버 측 객체 복사 시 적용됩니다.
      
      acl이 빈 문자열이라면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 canned ACL입니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      이 ACL은 버킷을 생성할 때에만 적용됩니다. bucket_acl이 설정되어 있지 않을 경우 "acl"이 대신 사용됩니다.
      
      만약 "acl"과 "bucket_acl"이 빈 문자열이라면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사용자에게는 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 액세스가 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스가 부여됩니다.
         |                    | 버킷에 대해서는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스가 부여됩니다.

   --storage-class
      S3에 새로운 객체를 저장할 때 사용할 저장 클래스입니다.

      예제:
         | <unset>  | 기본값입니다.
         | STANDARD | 스트리밍이나 CDN과 같은 주문형 콘텐츠에 적합한 표준 클래스입니다.
         |          | 기본값입니다.
         | GLACIER  | 아카이브 저장소입니다.
         |          | 가격은 낮지만, 먼저 복원해야 액세스할 수 있습니다.

   --upload-cutoff
      청크화 업로드로 전환하는데 필요한 임계값입니다.
      
      이보다 큰 파일은 chunk_size 단위로 청크화해서 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 모르는 파일("rclone rcat"으로 만든 파일이나 "rclone mount"나 구글
      사진이나 구글 문서로 업로드한 파일 등)을 업로드할 때 이 청크 크기를 사용하여
      multipart 업로드로 업로드합니다.
      
      "--s3-upload-concurrency" 크기의 청크는 전송 당 메모리에 버퍼로 유지됩니다.
      
      고속링크로 큰 파일을 전송하고 충분한 메모리가 있다면 이 값을 늘리면 전송 속도가 향상됩니다.
      
      큰 파일(크기를 알 수 있는)을 업로드할 때 Rclone은 10,000개의 청크 제한을 지키기 위해
      청크 크기를 자동으로 늘립니다.
      
      크기를 알 수 없는 파일은 구성된 청크 크기로 업로드됩니다.
      기본 청크 크기는 5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로,
      기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다.
      더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야 합니다.
      
      청크 크기를 늘리면 "-P" 플래그로 표시되는 진행 통계의 정확성이 감소합니다.
      Rclone은 청크가 AWS SDK에 의해 buffered될 때 청크를 보낸 것으로 처리하지만
      실제로는 여전히 업로드 중일 수 있습니다.
      청크 크기가 클수록 AWS SDK 버퍼가 크고 진행률 표시가 실제와 다르게 표시됩니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 청크 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      일부 서비스에서는 AWS S3 사양의 10,000개의 청크를 지원하지 않을 수 있으므로
      이 옵션이 유용할 수 있습니다.
      
      Rclone은 10,000개의 청크 제한을 지키기 위해 크기를 알 수 있는 큰 파일을 업로드할 때
      청크 크기를 자동으로 늘립니다.
      

   --copy-cutoff
      파티션화 복사로 전환하는데 필요한 임계값입니다.
      
      이보다 크기가 큰 파일을 서버 측에서 복사할 때 이 크기로 파일을 복사합니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만
      크기가 큰 파일을 업로드하는 데 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE"
      환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리가 기본값입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이
      변수는 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 "default"가 기본값입니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      대용량 파일을 고속 링크로 여러 개 업로드하고이 업로드가 제대로 이루어지는 동안 대역폭을 충분히 이용하지 못하는 경우에는
      이 값을 늘려 전송 속도를 높일 수 있습니다.

   --force-path-style
      참이면 경로 스타일 액세스를 사용하고, 그렇지 않으면 가상 호스트형 스타일 액세스를 사용합니다.
      
      이 값이 true인 경우(rclone의 기본 설정) rclone은 경로 스타일 액세스를 사용합니다.
      그렇지 않으면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 AWS S3
      문서 [링크](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는
      이 값을 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 이를 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false로 설정하면(rclone의 기본 설정) v4 인증을 사용합니다. 설정되어 있으면 v2 인증을 사용합니다.
      
      v4 시그니처가 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기입니다 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청 수보다 많이 요청한 경우에도 응답 목록을 1000개로 자릅니다.
      AWS S3에서 이 수는 전체 최대값이므로 변경할 수 없습니다. [AWS S3 문서](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2, 자동 선택을 위해 0.
      
      S3가 처음 출시될 때는 버킷의 개체를 열람하는 ListObjects 호출만 지원했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 소개되었습니다. 이것은
      매우 높은 성능을 제공하므로 가능하면 사용해야 합니다.
      
      기본 설정인 0으로 설정하면 rclone은 공급자 설정에 따라 어떤 목록 객체 방법을 호출할 지 추측합니다.
      추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 url 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며,
      이 경우 사용하면 파일 내용에 대해 더 안정적입니다. unset으로 설정된 경우(rclone의 기본 설정) rclone은
      공급자 설정에 따라 적용할 내용을 결정하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 만들려고 시도하지 않습니다.
      
      버킷이 이미 존재하는 것을 알고 있다면 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      또한 버킷 생성 권한이 없는 사용자를 사용해야 하는 경우 필요할 수 있습니다.
      v1.52.0 이전에는 이 버그로 인해 무음으로 통과되었습니다.
      

   --no-head
      업로드된 객체의 무결성 확인을 위해 HEAD을 사용하지 않습니다.
      
      rclone은 기본적으로 객체를 PUT한 후 200 OK 메시지를 받으면 제대로 업로드된 것으로 가정합니다.
      따라서 다음과 같이 가정합니다.
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형 등)는 업로드한 것과 동일
      - 크기는 업로드한 것과 동일
      
      단일 파트 PUT에 대한 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      크기를 알 수 없는 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기와 같은 감지되지 않은 업로드 오류가 발생할 수있는
      위험성이 높아지므로 일반적인 작업에는 권장하지 않습니다. 실제로
      이 플래그가 설정되어도 감지되지 않은 업로드 오류의 가능성은 매우 작습니다.
      

   --no-head-object
      GET을 실행하기 전에 HEAD을 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 언제 플러시할지를 결정합니다.
      
      (분할) 업로드에 추가 버퍼가 필요한 경우 할당을 위해 메모리 풀을 사용할 수 있습니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관한 문제가 해결되지 않은 상태입니다. HTTP/2는
      s3 백엔드의 기본 설정이지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면
      이 플래그는 제거될 것입니다.
      
      참고: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      AWS S3는 클라우드프론트 CDN URL로 설정하는 것이 일반적이며
      클라우드프론트 네트워크를 통해 데이터를 다운로드하는 경우 AWS S3보다 더 저렴한 데이터 이탈료를 제공합니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 기본값을 사용할지 여부로 설정해야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      false로 설정된 경우 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 1.59 미만은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하며
      이 플래그를 true로 설정하면 해당 기능이 다시 사용됩니다. 이는 특수한 경우나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", 날짜와 시간 "2006-01-02
      15:04:05" 또는 그만큼 예전에 대한 기간으로 설정할 수 있습니다. ex) "100d" 또는 "1h".
      
      이를 사용할 때는 파일 쓰기 동작을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      [time 옵션 문서](/docs/#time-option)에서 유효한 형식을 확인하세요.
      

   --decompress
      지정된 경우 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에는 "Content-Encoding: gzip"로 설정된 파일을 업로드하는 것이 가능합니다.
      보통 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 수신시 "Content-Encoding: gzip"로 이러한 파일을
      압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없게 만듭니다.
      

   --might-gzip
      백엔드에서 gzip을 적용할 수 있는지 여부입니다.
      
      일반적으로 공급자는 다운로드시 객체를 수정하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 경우
      다운로드시 설정되지 않습니다.
      
      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 파일(예: Cloudflare)도
      gzip으로 압축할 수 있습니다.
      
      이런 경우에는
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      와 같은 오류 메시지를 수신할 수 있습니다.
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩으로 객체를 다운로드하면
      rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset으로 설정된 경우(rclone의 기본 설정) rclone은 공급자 설정에 따라 적용할 내용을 결정하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value      AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다. [$ACL]
   --endpoint value           Scaleway Object Storage에 대한 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 권한을 실행 시간(runtime)에서 가져옵니다 (환경 변수 또는 env vars나 IAM이 없다면 EC2/ECS 메타 데이터에서). (default: false) [$ENV_AUTH]
   --help, -h                 도움말을 출력합니다
   --region value             연결할 지역입니다. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키(비밀번호)입니다. [$SECRET_ACCESS_KEY]
   --storage-class value      S3에 새로운 객체를 저장할 때 사용할 저장 클래스입니다. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              파티션화 복사로 전환하는데 필요한 임계값입니다. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     지정된 경우 gzip으로 인코딩된 객체를 압축 해제합니다. (default: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (default: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               참이면 경로 스타일 액세스를 사용하고, 그렇지 않으면 가상 호스트형 스타일 액세스를 사용합니다. (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다 (각 ListObject S3 요청에 대한 응답 목록). (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 url 인코딩할지 여부: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동 선택을 위해 0. (default: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 청크 수입니다. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 언제 플러시할지를 결정합니다. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip을 적용할 수 있는지 여부입니다. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 만들려고 시도하지 않습니다. (default: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성 확인을 위해 HEAD을 사용하지 않습니다. (default: false) [$NO_HEAD]
   --no-head-object                 GET을 실행하기 전에 HEAD을 수행하지 않습니다. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크화 업로드로 전환하는데 필요한 임계값입니다. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (default: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간에 표시합니다. (default: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (default: false) [$VERSIONS]

```
{% endcode %}