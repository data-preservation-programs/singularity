# SeaweedFS S3

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 seaweedfs - SeaweedFS S3

사용법:
   singularity storage create s3 seaweedfs [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경 변수가없는 경우 EC2 / ECS 메타 데이터).

      access_key_id와 secret_access_key이 비어있을 때만 적용됩니다.

      Examples:
         | false | 다음 단계에서 AWS 자격 증명 입력.
         | true  | 환경에서 AWS 자격 증명 가져오기 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.

      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (암호).

      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --region
      연결할 리전.

      S3 클론을 사용하고 리전이 없는 경우 비워 둡니다.

      Examples:
         | <unset>            | 확실하지 않은 경우 사용합니다.
         |                    | v4 시그니처 및 빈 리전이 사용됩니다.
         | other-v2-signature | v4 시그니처가 작동하지 않는 경우에만 사용합니다.
         |                    | 예 : Jewel / v10 이전 CEPH.

   --endpoint
      S3 API의 엔드 포인트.

      S3 클론을 사용하는 경우 필수입니다.

      Examples:
         | localhost:8333 | SeaweedFS S3 로컬 호스트

   --location-constraint
      리전과 일치해야하는 위치 제한.

      확실하지 않은 경우 비워 둡니다. 버킷을 생성하는 경우 사용됩니다.

   --acl
      버킷을 생성하고 개체를 저장하거나 복사할 때 사용되는 공유된 ACL.

      이 ACL은 개체를 생성할 때 사용되며, bucket_acl이 설정되지 않은 경우에도 버킷 생성에 사용됩니다.

      자세한 내용은 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하십시오.

      서버 측 복사로 개체를 복사 할 때 S3는 소스의 ACL을 복사하지 않고 새로 씁니다.

      ACL이 빈 문자열 인 경우 X-Amz-Acl : 헤더가 추가되지 않고 기본 (private)이 사용됩니다.

   --bucket-acl
      버킷을 만들 때 사용되는 공유된 ACL입니다.

      자세한 내용은 [https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하십시오.

      bucket_acl이 설정되지 않은 경우에만 적용됩니다.

      "acl" 및 "bucket_acl"이 빈 문자열 인 경우 X-Amz-Acl : 헤더가 추가되지 않고 기본 (private)이 사용됩니다.

      Examples:
         | private            | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 액세스를 얻습니다.
         | public-read-write  | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻습니다.
         |                    | 버킷에서이를 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL을 얻습니다.
         |                    | 인증 된 사용자 그룹이 읽기 액세스를 얻습니다.

   --upload-cutoff
      청크 업로드로 전환하는 임계 값.

      이보다 큰 파일은 chunk_size로 청크로 업로드됩니다.
      최소 0이고 최대 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.

      upload_cutoff보다 큰 파일이거나 크기가 알려지지 않은 파일 (예 : "rclone rcat"에서 가져온 파일이나 "rclone mount" 또는 google 사진 또는 google 문서로 업로드 된 파일)은 이 청크 크기를 사용하여 다단계 업로드로 업로드됩니다.

      참고로 "--s3-upload-concurrency"이 크기의 청크는 전송 당 메모리에 버퍼링됩니다.

      고속 링크로 대량의 대용량 파일을 전송하면서 충분한 메모리가있는 경우 이 값을 증가시키면 전송 속도가 향상됩니다.

      대용량 파일의 경우 rclone은 알려진 크기의 대형 파일을 10000개의 청크 한도를 초과하지 않도록 chunk 크기를 자동으로 증가시킵니다.

      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트리밍 업로드하려면 chunk_size를 증가시켜야합니다.

      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태 통계의 정확도가 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링될 때 청크가 전송된 것으로 처리하므로 실제로 업로드 중일 수 있음에도 불구하고 전송 상태에서처럼 진행됩니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률보고 차이가 나는 것을 의미합니다.

   --max-upload-parts
      다중 부분 업로드의 최대 부분 수.

      이 옵션은 다중 부분 업로드를 수행 할 때 사용할 최대 부분 청크 수를 정의합니다.

      이는 서비스가 AWS S3 사양에서 10,000 개의 청크를 지원하지 않는 경우 유용 할 수 있습니다.

      알려진 크기의 대형 파일을 업로드 할 때 rclone은 청크 크기를 자동으로 증가시켜 이러한 청크 수 제한을 초과하지 않도록합니다.

   --copy-cutoff
      청크 복사로 전환하는 임계 값.

      서버 측에서 복사해야하는이보다 큰 파일은 이 크기로 청크로 복사됩니다.

      최소 0이고 최대 5 GiB입니다.

   --disable-checksum
      개체 메타 데이터에 MD5 체크 섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크 섬을 계산하여 개체의 메타 데이터에 추가합니다. 이는 대형 파일을 업로드하는 데 오랜 지연을 유발할 수 있지만 데이터 무결성 확인에 좋습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.

      env_auth = true 인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. env 값이 비어 있으면 현재 사용자의 홈 디렉토리로 기본 설정됩니다.

          Linux / OSX : "$ HOME / .aws / credentials"
          Windows : "% USERPROFILE% \ .aws \ credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필.

      env_auth = true 인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용되는 프로필을 제어합니다.

      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 기본값으로 설정됩니다.

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      다중 부분 업로드에 대한 동시성.

      동시에 업로드되는 동일한 파일의 청크 수입니다.

      대량의 대형 파일을 고속 연결로 업로드하는 경우이 업로드가 대역폭을 완전히 활용하지 못하면이 값을 증가시키는 것이 전송 속도를 높일 수 있습니다.

   --force-path-style
      true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅 스타일 액세스를 사용합니다.

      true (기본값)이면 rclone은 경로 스타일 액세스를 사용하고 false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) 문서를 참조하십시오.

      일부 공급 업체 (예 : AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이 작업을 자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.

      false (기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.

      v4 시그니처가 작동하지 않는 경우에만 사용하십시오. 예 : Jewel / v10 이전 CEPH.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).

      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.

      대부분의 서비스는 요청 된 것보다 많은 수를 요청하더라도 응답 목록을 1000 개의 개체로 자르기 때문에 전역 최대 수입니다.

      AWS S3에서는이 수는 변경할 수 없으며 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.

      Ceph의 경우 "rgw list buckets max chunk" 옵션을 사용하여이 크기를 늘릴 수 있습니다.

   --list-version
      사용할 ListObjects 버전 : 1,2 또는 자동으로 지정하기 위해 0.

      S3가 처음 출시되었을 때 버킷의 개체를 나열하는 ListObjects 호출만 제공되었습니다.

      그러나 2016 년 5 월에는 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 높은 성능이며 가능한 경우 사용해야합니다.

      기본값인 0으로 설정하면 rclone은 공급 업체 설정에 따라 어떤 목록 객체 방법을 호출 할 것으로 추측합니다. 이 추측이 잘못되면 여기서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 URL 인코딩 할지 여부 : true / false / unset

      일부 공급 업체는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 이를 사용하면 파일의 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다. 이 경우 rclone은 공급자 설정에 따라 rclone의 선택을 무시합니다.

   --no-check-bucket
      버킷의 존재 확인이나 생성을 시도하지 않습니다.

      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용 할 수 있습니다.

      또는 사용자가 버킷 생성 권한이없는 경우 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이것이 의미없이 전달되었습니다.

   --no-head
      업로드 된 개체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다.

      rclone은 단일 부분을 사용하여 객체를 업로드 한 후에 200 OK 메시지를 수신하면 제대로 업로드 된 것으로 가정합니다.

      특히 다음과 같이 가정합니다.

      - 업로드시 메타 데이터 (modtime, 저장 클래스 및 콘텐츠 유형)가 업로드와 동일했습니다.
      - 크기가 업로드와 동일했습니다.

      단일 부분 PUT의 응답에서 다음 항목을 읽습니다.

      - MD5SUM
      - 업로드 날짜

      다중 부분 업로드의 경우이 항목은 읽지 않습니다.

      알 수없는 길이의 소스 개체가 업로드 된 경우 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 부적절한 크기를 포함한 감지되지 않은 업로드 실패 가능성이 증가하므로 정상 작업에는 권장되지 않습니다. 실제로 이 플래그를 사용하여 감지되지 않은 업로드 오류 가능성은 매우 낮습니다.

   --no-head-object
      GET을 수행하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기.

      추가 버퍼가 필요한 업로드 (예 : 다중 부분)은 메모리 풀을 사용하여 할당합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용 비활성화.

      현재 s3 (특히 minio) 백엔드와 HTTP / 2의 미해결 된 문제가 있습니다. s3 백엔드의 HTTP / 2는 기본적으로 사용되지만이 여기에서 비활성화 할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.

      참조 : [https://github.com/rclone/rclone/issues/4673](https://github.com/rclone/rclone/issues/4673), [https://github.com/rclone/rclone/issues/3631](https://github.com/rclone/rclone/issues/3631)

   --download-url
      다운로드에 대한 사용자 정의 엔드 포인트.
      AWS S3는 CloudFront 네트워크를 통해 다운로드되는 데이터에 대한 경제적인 Egress를 제공하므로 일반적으로 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      검증을 위해 multipart 업로드에서 ETag를 사용할지 여부

      true, false 또는 제공자의 기본값을 사용하도록 설정해야합니다.

   --use-presigned-request
      단일 부분 업로드에 presigned 요청 또는 PutObject를 사용할지 여부

      이 값이 false 인 경우 rclone은 AWS SDK에서 PutObject를 사용하여 개체를 업로드합니다.

      rclone의 버전 1.59 이전 버전은 단일 부분 개체를 업로드하기 위해 서명 된 요청을 사용하며이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 예외적인 경우 또는 테스트를 위해서만이 필요합니다.

   --versions
      디렉토리 목록에 이전 버전 포함.

   --version-at
      지정된 시간에 대한 파일 버전으로 표시합니다.

      매개 변수는 날짜 (예 : "2006-01-02"), datetime "2006-01-02 15:04:05" 또는 그로부터 오래된 기간 (예 : "100d" 또는 "1h") 일 수 있습니다.

      이 옵션을 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제 할 수 없습니다.

      유효한 형식에 대해서는 [시간 옵션 도움말](https://github.com/rclone/rclone/blob/master/docs/content/docs.md#time-option)을 참조하십시오.

   --decompress
      gzip으로 압축 된 객체를 압축 해제합니다.

      S3로 파일을 "Content-Encoding : gzip"으로 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축 해제 된 개체로 다운로드합니다.

      이 플래그가 설정되면 rclone은 수신되는 파일을 "Content-Encoding : gzip"으로 압축 해제합니다. 이는 rclone이 크기와 해시를 확인 할 수 없지만 파일 내용은 압축 해제 된다는 것을 의미합니다.

   --might-gzip
      백엔드에서 gzip 개체를 압축 할 수 있습니다.

      일반적으로 공급 업체는 다운로드 될 때 개체를 변경하지 않습니다. "Content-Encoding : gzip"로 업로드되지 않은 개체는 다운로드되지 않습니다.

      그러나 일부 제공 업체는 gzip으로 압축될 수 있습니다 (예 : Cloudflare).

      이로 인해 다음과 같은 오류가 발생합니다.

          ERROR corrupted on transfer: sizes differ NNN vs MMM

      이 플래그를 설정하고 rclone이 "Content-Encoding: gzip"로 설정되고 청크 전송 인코딩 메시지를 수신하면 rclone은 실시간으로 개체를 압축 해제합니다.

      unset으로 설정하면 rclone은 공급자 설정에 따라 적용 할 내용을 선택하지만이에서 rclone의 선택을 무시 할 수 있습니다.

   --no-system-metadata
      시스템 메타 데이터의 설정 및 읽기를 억제


OPTIONS:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 개체를 저장하거나 복사할 때 사용되는 공유 ACL. [$ACL]
   --endpoint value             S3 API의 엔드 포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경 변수가없는 경우 EC2 / ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  리전과 일치해야하는 위치 제한. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전. [$REGION]
   --secret-access-key value    AWS Secret Access Key (암호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 만들 때 사용되는 공유 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크 복사로 전환하는 임계 값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 압축 된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               개체 메타 데이터에 MD5 체크 섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드 포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 경로 스타일 액세스를 사용하고 false 인 경우 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩 할지 여부 : true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전 : 1,2 또는 자동으로 지정하기 위해 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         다중 부분 업로드의 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip 개체를 압축 할 수 있습니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 확인이나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드 된 개체의 무결성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타 데이터의 설정 및 읽기를 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       다중 부분 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계 값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 multipart 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 presigned 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 대한 파일 버전으로 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전 포함. (기본값: false) [$VERSIONS]

   일반적인

   --name value  스토리지의 이름 (기본값: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}