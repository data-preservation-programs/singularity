# SeaweedFS S3

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 seaweedfs - SeaweedFS S3

사용법:
   singularity storage update s3 seaweedfs [command options] <이름|ID>

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env 변수 또는 EC2/ECS 메타 데이터를 사용해서).
      
      access_key_id 및 secret_access_key 값이 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 (환경 변수 또는 IAM) AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두세요.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두세요.

   --region
      연결할 지역.
      
      S3 클론을 사용하고 지역을 가지지 않은 경우 비워 두세요.

      예시:
         | <지정하지 않음> | 확인할 수 없는 경우 사용하세요.
         |                 | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 동작하지 않을 때만 사용하세요.
         |                   | 예: Jewel/v10 이전 버전 CEPH.

   --endpoint
      S3 API의 엔드포인트.
      
      S3 클론을 사용하는 경우 필요합니다.

      예시:
         | localhost:8333 | SeaweedFS S3 localhost

   --location-constraint
      지역 제한 - 지역과 일치하도록 설정해야 합니다.
      
      확실하지 않은 경우 비워 두세요. 버킷을 생성할 때만 사용됩니다.

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL.
      
      이 ACL은 객체를 생성할 때 사용되며 bucket_acl 값이 설정되지 않은 경우에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl 을 참조하세요.
      
      S3에서 서버 측 복사하는 경우에만 이 ACL이 적용됩니다.
      원본에서 ACL을 복사하는 것이 아니라 새로 쓰기 때문입니다.
      
      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl 을 참조하세요.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예시:
         | private            | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹이 읽기 액세스를 얻습니다.
         | public-read-write  | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 액세스를 얻습니다.
         |                    | 이 기능을 버킷에 적용하는 것을 일반적으로 권장하지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AuthenticatedUsers 그룹이 읽기 액세스를 얻습니다.

   --upload-cutoff
      청크드 업로드로 전환하는 업로드 최대 크기.
      
      이 값보다 큰 파일은 chunk_size로 청크로 업로드됩니다.
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --chunk-size
      업로드시 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기가 알려지지 않은 파일(예: "rclone rcat"으로 업로드한 파일이나 "rclone mount" 또는 Google
      사진 또는 Google 문서로 업로드한 파일)은 이 청크 크기를 사용하여 multipart 업로드됩니다.
      
      참고로 "--s3-upload-concurrency"는 전송마다 이 크기의 청크를 메모리에 버퍼링합니다.
      
      고속 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 향상됩니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 최대 10,000개의 청크 제한을 유지합니다.
      
      알려지지 않는 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로,
      기본적으로 스트리밍 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트리밍 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확도가 낮아집니다. rclone은 청크가 AWS SDK에 의해 버퍼링되었을 때 청크를
      전송된 것으로 처리하지만 실제로는 아직 업로드 중일 수 있습니다.
      큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행 상태 보고서가 진실에서 벗어난다는 의미입니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 청크 수.
      
      이 옵션은 멀티파트 업로드 중에 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      10,000개의 청크 규격을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대용량 파일을 업로드할 때 자동으로 청크 크기를 증가시켜 이 청크 수 제한을 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 복사 최대 크기.
      
      서버 측에서 복사해야 하는 이 값보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      보통 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에는 좋지만
      파일 크기가 큰 경우 업로드 시작까지 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true 인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있는 경우 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 확인합니다. 환경 변수 값이 비어 있는 경우 기본값은
      현재 사용자의 홈 디렉토리입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일입니다.
      
      env_auth = true 인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로파일을 제어합니다.
      
      비워 두면 환경 변수 "AWS_PROFILE" 또는 "default"가 기본값으로 사용됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      고속 링크를 통해 고속으로 대량 파일을 업로드하는 경우이 업로드가 대역폭을 전부 사용하지 못하는 경우에는이 값을 증가시키면 전송 속도가 향상될
      수 있습니다.

   --force-path-style
      true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅된 스타일을 사용합니다.
      
      true (기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false인 경우 rclone은 가상 경로 스타일을 사용합니다.
      자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공자 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이 값을 false로 설정해야 합니다. rclone은 제공자 설정을 기반으로 이 작업을
      자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.
      
      false (기본값)인 경우 rclone은 v4 인증을 사용합니다.
      설정된 경우 rclone은 v4 인증을 사용하려고 시도할 것입니다.
      
      v4 서명이 동작하지 않는 경우에만 사용하세요. 예: Jewel/v10 이전 버전 CEPH.

   --list-chunk
      리스트 청크의 크기 (각 ListObject S3 요청의 응답 리스트).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청 크기를 1000개로 잘라냅니다. 이 값을 요청보다 크게 설정해도 응답 리스트는 1000개로 잘라내집니다.
      AWS S3에서 이 값은 전역 최대값이므로 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값 크기를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동(0).
      
      S3가 처음 출시될 때 버킷의 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 높은 성능을 제공하며 가능하면 사용해야 합니다.
      기본값 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 제공자는 목록 URL 인코딩을 지원하며, 사용 가능한 경우 파일 이름에 제어 문자를 사용할 때이 방법이 더 안정적입니다.
      unset으로 설정하면 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      또한 사용자가 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 전에는 버그로 인해 이 작업이 무시될 수 있었습니다.
      

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone이 PUT로 객체를 업로드한 후 200 OK 메시지를 수신하면 제대로 업로드되었다고 가정합니다.
      
      특히 다음 사항을 가정합니다:
      
      - 업로드되는 내용, 수정 시간, 저장 클래스 및 콘텐츠 유형은 업로드된 것과 동일합니다.
      - 크기는 업로드된 것과 동일합니다.
      
      단일 파트 PUT 응답에서 다음 항목을 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드인 경우 이러한 항목은 읽지 않습니다.
      
      크기가 알려지지 않은 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 확률이 높아져 일반 작업에는 권장되지 않습니다. 실제로 업로드 실패의 확률은 매우 낮습니다.
      

   --no-head-object
      GET 시 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 플러시할지를 결정합니다.
      
      메모리 풀이 필요한 버퍼(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용 비활성화.
      
      현재 S3 (특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. HTTP/2는 s3 백엔드의 기본값이지만 이곳에서 비활성화할 수 있습니다.
      이 문제가 해결될 때까지 이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 외부 데이터 전송을 제공합니다.

   --use-multipart-etag
      확인을 위해 multipart 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 기본값을 사용할지 여부를 설정해야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 값을 false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone < 1.59 버전에서는 단일 부분 객체를 업로드하기 위해 서명된 요청을 사용하는데, 해당 기능은 이러한 상황이나 테스트 외에는 불필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간 해당 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 시간("2006-01-02 15:04:05") 또는 그 이전을 의미하는 지속되는 기간("100d" 또는 "1h")여야 합니다.
      
      이를 사용하면 파일 쓰기 작업을 수행할 수 없으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      사용 가능한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"가 설정된 상태로 객체를 S3에 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 받은 상태로 "Content-Encoding: gzip"가 있는 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip될 수 있는 경우 설정하세요.
      
      일반적으로 제공자는 다운로드시에 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드시에도 해당 값이 설정되지 않습니다.
      
      그러나 일부 제공자는 객체를 "Content-Encoding: gzip"로 업로드하지 않았더라도 객체를 gzip 압축합니다(Cloudflare와 같은 경우).
      
      이러한 경우 다음과 같은 오류를 받게 될 수 있습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip으로 설정된 청크드 전송 코드를 사용하여 객체를 다운로드하면 rclone은 객체를 압축 풀어 향시키도록 합니다.
      
      unset으로 설정하면 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value             S3 API의 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env 변수 또는 EC2/ECS 메타 데이터를 사용해서). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  지역 제한 - 지역과 일치하도록 설정해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS 비밀 액세스 키 (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드시 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              복사로 전환하는 복사 최대 크기. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅된 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기 (각 ListObject S3 요청의 응답 리스트). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1, 2 또는 자동(0). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 플러시할지를 결정합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip될 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET 시 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로파일입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크드 업로드로 전환하는 업로드 최대 크기. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 multipart 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간 해당 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}