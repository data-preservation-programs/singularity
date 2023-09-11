# Ceph Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 ceph - Ceph Object Storage

사용법:
   singularity storage update s3 ceph [command options] <name|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없을 경우 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key이 비어 있으면 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --region
      연결할 리전입니다.
      
      S3 복제품을 사용하고 영역이 없는 경우 비워 둡니다.

      예:
         | <unset>            | 확실하지 않으면 사용하세요.
         |                    | v4 서명 및 빈 영역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용하세요.
         |                    | 예: Jewel/v10 이전 CEPH.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 복제품을 사용하는 경우 필요합니다.

   --location-constraint
      위치 제약 조건 - 리전과 일치해야 합니다.
      
      확실하지 않으면 비워 두세요. 버킷 생성 시에만 사용됩니다.

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 캐넌 ACL입니다.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl에서 확인할 수 있습니다.
      
      S3는 객체를 서버 간에 복사할 때 소스에서 ACL을 복사하는 것이 아니라 새로운 ACL을 작성하므로 이 ACL이 적용됩니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 캐넌 ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl에서 확인할 수 있습니다.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않으면 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예:
         | private            | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | 다른 사람에게는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹에는 읽기 권한이 있습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹에는 읽기 및 쓰기 권한이 있습니다.
         |                    | 버킷에 대해 이 권한을 허용하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AuthenticatedUsers 그룹에는 읽기 권한이 있습니다.

   --server-side-encryption
      S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예:
         | <unset> | 암호화되지 않음
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우 S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다.

      예:
         | <unset> | 암호화되지 않음
         | AES256  | AES256

   --sse-kms-key-id
      KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다.

      예:
         | <unset>                 | 암호화되지 않음
         | arn:aws:kms:us-east-1:* | arn:aws:kms:*

   --sse-customer-key
      SSE-C를 사용하려면 데이터를 암호화/복호화하는 데 사용되는 비밀 암호화 키를 제공할 수 있습니다.
      
      대신 --sse-customer-key-base64를 제공할 수도 있습니다.

      예:
         | <unset> | 암호화되지 않음

   --sse-customer-key-base64
      SSE-C를 사용하는 경우 데이터를 암호화/복호화하는 데 사용되는 비밀 암호화 키를 Base64로 인코딩된 형식으로 제공해야 합니다.
      
      대신 --sse-customer-key를 제공할 수도 있습니다.

      예:
         | <unset> | 암호화되지 않음

   --sse-customer-key-md5
      SSE-C를 사용하는 경우 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비워 두면 이 값은 sse_customer_key에서 자동으로 계산됩니다.
      

      예:
         | <unset> | 암호화되지 않음

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat" 또는 "rclone mount" 또는 google
      photos 또는 google docs에서 업로드된 파일 등)을 업로드할 때 이 청크 크기를 사용하여 multipart 업로드로 업로드됩니다.
      
      참고로 "--s3-upload-concurrency"는 이 크기의 청크가 각 전송별로 메모리에 버퍼링됩니다.
      
      높은 속도의 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 빨라집니다.
      
      rclone은 알려진 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 초과하지 않도록 합니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로
      기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 감소합니다. rclone은 청크가 AWS SDK에
      버퍼링될 때 청크를 전송한 것으로 간주하지만 이후에도 업로드 중일 수 있습니다.
      큰 청크 크기는 큰 AWS SDK 버퍼와 진행 상태 보고가 진실과 더 다른 것을 야기합니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 응답 수를 정의합니다.
      
      AWS S3의 10,000개 청크 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      rclone은 알려진 큰 파일을 업로드할 때 chunk_size를 자동으로 증가시켜 이 청크 수 제한을 초과하지 않도록 합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값입니다.
      
      서버 간 복사가 필요한 이보다 큰 파일은 이 크기로 청크 단위로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는
      데이터 무결성 확인에 유용하지만 대용량 파일의 시작 전에 긴 지연을 초래할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      해당 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" env 변수를 찾습니다. env 값이 비어 있으면
      현재 사용자의 홈 디렉토리로 기본값으로 설정합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 해당 변수는 그 파일에서 사용할 프로필을 제어합니다.
      
      비워 둔 경우 환경 변수 "AWS_PROFILE" 또는 그 환경 변수가 설정되지 않은 경우 "default"로 기본값으로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      파일 변환이 완전히 대역폭을 활용하지 않고 높은 속도의 링크를 통해 작은 수의 대용량 파일을 업로드하는 경우 이 값을 증가시키면 전송을 빠르게 할 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스, false이면 가상 호스팅 스타일 액세스를 사용합니다.
      
      이 값이 true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고, false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 false로 설정해야 합니다. rclone은 해당 제공자에 설정에 따라 자동으로 이 작업을 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false이면 rclone은 v4 인증을 사용합니다. 설정되면 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 사용하세요. 예: Jewel/v10 CEPH와 같이.

   --list-chunk
      목록 청크의 크기(ListObject S3 요청마다 응답 목록의 크기)입니다.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청보다 많은 응답 목록을 1000개로 잘립니다.
      AWS S3에서는 이것이 전역적인 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 0(자동).
      
      처음에 S3가 출시되었을 때 버킷의 객체를 나열하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 높은 성능을 제공하므로 가능하면 사용해야 합니다.
      
      기본값 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 list objects 방법을 추측합니다. 잘못 추측하면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 제공자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며 가능한 경우 이 방법이 더 안정적입니다. unset으로 설정하면
      rclone은 공급자 설정에 따라 적용할 것을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷을 체크하거나 생성하는 시도를 하지 않습니다.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이것은 조용히 전달되었을 것입니다.
      

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone은 PUT로 객체 업로드 후에 200 OK 메시지를 받으면 올바르게 업로드되었다고 가정합니다.
      
      특히 rclone은 다음을 가정합니다:
      
      - 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드한 것과 동일하다.
      - 크기가 업로드한 것과 동일하다.
      
      단일 파트 PUT의 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드는 이러한 항목을 읽지 않습니다.
      
      크기를 알 수 없는 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 감지 확률이 증가하며 특히 잘못된 크기가 발생할 수 있으므로 일반적인 작업에는 권장되지 않습니다. 실제로 업로드 실패가
      감지되지 않을 확률은 매우 작습니다.
      

   --no-head-object
      GET할 때 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      더 자세한 내용은 [개요에서 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 시간 간격입니다.
      
      추가 버퍼(예: 멀티파트)가 필요한 업로드는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. HTTP/2는
      s3 백엔드에서 기본적으로 활성화되지만 여기서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 다운로드되는 데이터에 대해
      더 저렴한 출발이로 제공합니다.

   --use-multipart-etag
      확인을 위해 Multipart 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 기본값을 사용하려면 true, false 또는 unset을 사용하세요.
      

   --use-presigned-request
      단일 파트 업로드를위한 프리 서명된 요청 또는 PutObject 사용 여부
      
      이 값이 false이면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone의 버전 1.59 이하는 단일 파트 객체를 업로드하기 위해 프리 서명된 요청을 사용하고 true로 설정하면 해당 기능을 다시 활성화합니다.
      이것은 특수한 경우 또는 테스트를 제외하고 필요하지 않습니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 그대로 표시합니다.
      
      매개변수는 날짜, "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그만큼 오래된 기간(예: "100d" 또는 "1h")일 수 있습니다.
      
      이렇게 사용할 때 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      이 옵션을 설정하면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"가 설정된 상태로 S3에 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로
      다운로드합니다.
      
      이 플래그가 설정되면 rclone은 받는 동안 "Content-Encoding: gzip"와 함께 받은 파일을 압축 해제합니다. 이는 rclone
      크기와 해시를 확인할 수 없지만 파일 콘텐츠는 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축 할 수있는 경우 설정합니다.
      
      일반적으로 공급자는 개별 다운로드시 객체를 수정하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 객체가 넘어가면
      다운로드 할 때 `Content-Encoding: gzip`로 설정되지 않을 것입니다.
      
      그러나 일부 공급자는 `Content-Encoding: gzip`로 업로드되지 않은 객체를 gzip으로 압축 할 수 있습니다(Cloudflare 등).
      
      이 플래그를 설정하면 rclone은 Content-Encoding: gzip이 설정된 상태에서 청크 전송 인코딩이 있는 객체를 스트리밍하는 경우 객체를
      실시간으로 압축 풀어 줍니다.
      
      unset(기본값)으로 설정하면 rclone은 공급자 설정에 따라 적용할 것을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터에 대한 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value            AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                      버킷 생성 및 객체 저장 또는 복사 시 사용되는 캐넌 ACL입니다. [$ACL]
   --endpoint value                 S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                       런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없을 경우 EC2/ECS 메타 데이터). (default: false) [$ENV_AUTH]
   --help, -h                       도움말 표시
   --location-constraint value      위치 제약 조건 - 리전과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value                   연결할 리전입니다. [$REGION]
   --secret-access-key value        AWS 비밀 액세스 키(암호)입니다. [$SECRET_ACCESS_KEY]
   --server-side-encryption value   S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --sse-kms-key-id value           KMS ID를 사용하는 경우 키의 ARN을 제공해야 합니다. [$SSE_KMS_KEY_ID]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 캐넌 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값입니다. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (default: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용을 비활성화합니다. (default: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일 액세스를 사용합니다. (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(ListObject S3 요청마다 응답 목록의 크기)입니다. (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0(자동). (default: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수입니다. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간 간격입니다. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축 할 수있는 경우 설정합니다. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷을 체크하거나 생성하는 시도를 하지 않습니다. (default: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (default: false) [$NO_HEAD]
   --no-head-object                 GET할 때 HEAD를 수행하지 않습니다. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터에 대한 설정 및 읽기를 억제합니다 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우 S3에 이 객체를 저장할 때 사용되는 서버 측 암호화 알고리즘입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하려면 데이터를 암호화/복호화하는 데 사용되는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우 데이터를 암호화/복호화하는 데 사용되는 비밀 암호화 키를 Base64로 인코딩된 형식으로 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값입니다. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 Multipart 업로드에서 ETag를 사용할지 여부 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드를위한 프리 서명된 요청 또는 PutObject 사용 여부 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (default: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 그대로 표시합니다. (default: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (default: false) [$VERSIONS]

```
{% endcode %}