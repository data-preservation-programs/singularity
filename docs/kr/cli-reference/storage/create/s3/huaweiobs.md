# Huawei Object Storage Service

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 huaweiobs - Huawei Object Storage Service

사용법:
   singularity storage create s3 huaweiobs [옵션] [인자...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars나 IAM이 없는 경우 EC2/ECS 메타데이터에서 가져옴).
      
      access_key_id와 secret_access_key가 비어있을 때만 적용됩니다.
      
      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하십시오.
         | true  | AWS 자격 증명을 환경으로부터 가져옵니다(env vars 또는 IAM).

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(암호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --region
      연결할 지역 - 버킷이 생성되고 데이터가 저장되는 위치여야 합니다. 엔드포인트와 동일해야 합니다.

      예시:
         | af-south-1     | AF-Johannesburg
         | ap-southeast-2 | AP-Bangkok
         | ap-southeast-3 | AP-Singapore
         | cn-east-3      | CN East-Shanghai1
         | cn-east-2      | CN East-Shanghai2
         | cn-north-1     | CN North-Beijing1
         | cn-north-4     | CN North-Beijing4
         | cn-south-1     | CN South-Guangzhou
         | ap-southeast-1 | CN-Hong Kong
         | sa-argentina-1 | LA-Buenos Aires1
         | sa-peru-1      | LA-Lima1
         | na-mexico-1    | LA-Mexico City1
         | sa-chile-1     | LA-Santiago2
         | sa-brazil-1    | LA-Sao Paulo1
         | ru-northwest-2 | RU-Moscow2

   --endpoint
      OBS API용 엔드포인트입니다.

      예시:
         | obs.af-south-1.myhuaweicloud.com     | AF-Johannesburg
         | obs.ap-southeast-2.myhuaweicloud.com | AP-Bangkok
         | obs.ap-southeast-3.myhuaweicloud.com | AP-Singapore
         | obs.cn-east-3.myhuaweicloud.com      | CN East-Shanghai1
         | obs.cn-east-2.myhuaweicloud.com      | CN East-Shanghai2
         | obs.cn-north-1.myhuaweicloud.com     | CN North-Beijing1
         | obs.cn-north-4.myhuaweicloud.com     | CN North-Beijing4
         | obs.cn-south-1.myhuaweicloud.com     | CN South-Guangzhou
         | obs.ap-southeast-1.myhuaweicloud.com | CN-Hong Kong
         | obs.sa-argentina-1.myhuaweicloud.com | LA-Buenos Aires1
         | obs.sa-peru-1.myhuaweicloud.com      | LA-Lima1
         | obs.na-mexico-1.myhuaweicloud.com    | LA-Mexico City1
         | obs.sa-chile-1.myhuaweicloud.com     | LA-Santiago2
         | obs.sa-brazil-1.myhuaweicloud.com    | LA-Sao Paulo1
         | obs.ru-northwest-2.myhuaweicloud.com | RU-Moscow2

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      서버 간 객체 복사 시 이 ACL은 S3가 소스로부터 ACL을 복사하는 대신 새로운 ACL을 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(private)이 사용됩니다.

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL입니다.
      
      자세한 내용은 [Amazon S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      이 ACL은 버킷을 생성할 때만 사용됩니다. 설정되지 않은 경우 "acl"을 대신 사용합니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(private)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사용자는 액세스 권한이 없음 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹은 읽기 권한을 부여받음.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 부여받음.
         |                    | 일반적으로 버킷에 이를 부여하는 것은 권장되지 않음.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹은 읽기 권한을 부여받음.

   --upload-cutoff
      청크 업로드로 전환하는 크기입니다.
      
      이 보다 큰 파일은 청크 크기로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 크거나 크기를 알 수 없는 파일(예: "rclone rcat"을 통해 업로드된 파일 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드된 파일)을 업로드하는 경우, 이 청크 크기를 사용하여 멀티파트 업로드되며 메모리당 "--s3-upload-concurrency" 청크를 버퍼링합니다.
      
      높은 속도의 링크를 통해 큰 파일을 전송하고 메모리가 충분한 경우 이 값을 증가시켜 전송 속도를 높일 수 있습니다.
      
      큰 파일(크기가 알려진 파일)을 업로드할 때 rclone은 10,000개의 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      
      크기를 모르는 파일은 설정된
      청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000개의 청크가 존재할 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상황 통계의 정확도가 감소합니다. rclone은 청크가 AWS SDK에서 버퍼링될 때 청크가 전송된 것으로 처리하지만 아직 업로드 중일 수 있습니다. 큰 청크 크기는 큰 AWS SDK 버퍼와 진행률 통계의 차이가 더 많이 발생하기 때문에 진행률 보고를 잘못된 것처럼 만듭니다.

   --max-upload-parts
      HTTP multipart 업로드에서 사용하는 파트의 최대 수입니다.
      
      이 옵션은 multipart 업로드 시 사용할 최대 multipart 청크 수를 정의합니다.
      
      이는 AWS S3의 10,000개 청크 사양을 지원하지 않는 서비스에서 유용합니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 이 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 준수합니다.

   --copy-cutoff
      멀티파트 복사로 전환하는 크기입니다.
      
      서버 측에서 복사해야 하는 이 크기보다 큰 파일은 이 크기로 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 좋지만 큰 파일을 업로드할 때 시작 지연이 길어질 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true일 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어있으면 현재 사용자의 홈 디렉토리가 기본값으로 사용됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워두면 환경 변수 "AWS_PROFILE" 또는 설정되지 않은 경우 "default"를 기본값으로 사용합니다.

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크를 동시에 업로드하는 수입니다.
      
      대역폭을 완전히 활용하지 못하는 상황에서 대량의 대용량 파일을 업로드하는 경우 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      기본값으로 true인 경우 rclone은 경로 스타일 액세스를 사용하고 false일 경우 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      어떤 공급자(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는이 값을
      false로 설정해야 합니다. rclone은 제공자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      기본값인 경우 false이며 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하십시오. 예: 예전 Jewel/v10 CEPH.

   --list-chunk
      리스트 청크의 크기입니다(ListObject S3 요청마다 응답 리스트).
      
      이 옵션은 AWS S3 명세의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 1000개 이상을 요청하더라도 응답 리스트를 1000개로 자름.
      AWS S3에서 이는 전역 최대값이며 변경할 수 없으며 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph의 경우 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동(0).
      
      S3가 처음에 출시되었을 때는 ListObjects 호출만을 사용하여 버킷의 객체를 열거하였습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 호출은 훨씬 더 높은 성능을 제공하며 가능한 경우에 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 list 객체 방법을 추측합니다. 추측이 잘못된 경우 여기서 수동으로 설정할 수 있습니다.

   --list-url-encode
      URL 인코딩을 수행할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원하며 사용 가능한 경우 이러한 목록을 사용할 수 있습니다. unset으로 설정된 경우 rclone은 제공자 설정에 따라 선택합니다.

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않으려면 설정하십시오.
      
      버킷이 이미 존재할 경우 rclone 동작을 최소화하는 데 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우 필요할 수도 있습니다. v1.52.0 이전에는 이는 오류로 넘어갔습니다.

   --no-head
      업로드한 객체의 HEAD를 검사하여 무결성을 확인하지 않습니다.
      
      rclone은 기본적으로 PUT로 객체를 업로드 한 후 200 OK 메시지를 수신하면 올바르게 업로드되었다고 가정합니다.
      
      특히, 다음을 가정합니다.
      
      - 업로드 시, 메타데이터(수정 시각, 스토리지 클래스, 콘텐츠 유형)가 업로드한 상태와 동일한 것
      - 크기가 업로드한 것과 동일한 것
      
      PUT의 응답에서 다음 항목을 읽습니다(단일 파트 PUT의 경우):
      
      - MD5SUM
      - 업로드 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      길이를 알 수 없는 소스 객체를 업로드하면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 가능성이 증가하며 특히 유효하지 않은 크기의 경우 그렇게 되므로 정상적인 동작에는 권장하지 않습니다. 실제로 업로드 실패 가능성은 매우 작습니다.

   --no-head-object
      개체를 가져오기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당에 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2 문제가 해결되지 않은 상태입니다. HTTP/2는 기본적으로 s3 백엔드에 대해 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드에 사용할 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우보다 저렴한 이직을 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에 ETag를 사용할지 여부
      
      이는 true, false 또는 입력되지 않은 상태로 지정할 수 있습니다.
      

   --use-presigned-request
      단일 청크 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부
      
      이 옵션이 false로 설정되면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 < 1.59는 사전 서명된 요청을 사용하여 단일 파트 객체를 업로드하고이 플래그를 true로 설정하면이 기능을 다시 활성화합니다. 이는 예외적인 경우나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정한 시간에 버전이 있는 대상 파일을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 날짜와 시간("2006-01-02 15:04:05") 또는 그 시간 전의 지속 시간, 예를 들어 "100d" 또는 "1h"입니다.
      
      이를 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일 업로드 또는 삭제를 수행할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

   --decompress
      gzip으로 인코딩된 개체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드할 수 있습니다. 보통 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 객체를 수신할 때 "Content-Encoding: gzip"로 이러한 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

   --might-gzip
      백엔드가 gzip 객체를 압축할 수 있는 경우에 설정하십시오.
      
      일반적으로 공급자는 객체를 다운로드할 때 객체를 수정하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않았다면 다운로드되지 않을 것입니다.
      
      그러나 일부 제공자는 객체를 "Content-Encoding: gzip"로 업로드하지 않았더라도 gzip으로 압축할 수 있습니다(예: Cloudflare).
      
      이 경우 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩이 설정된 상태로 개체를 다운로드하면 rclone은 해당 객체를 실시간으로 압축 해제합니다.
      
      unset으로 설정된 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택합니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


옵션:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --endpoint value           OBS API용 엔드포인트. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars나 IAM이 없는 경우 EC2/ECS 메타데이터에서 가져옴). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역. [$REGION]
   --secret-access-key value  AWS Secret Access Key(암호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 크기. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 사용할 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          URL 인코딩을 수행할지 여부. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0(자동) (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용하는 파트의 최대 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 gzip 객체를 압축할 수 있는 경우에 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정하면 버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드한 객체의 HEAD를 검사하여 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져오기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다. (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 크기. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에 ETag를 사용할지 여부. (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 청크 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정한 시간에 버전이 있는 대상 파일을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   일반

   --name value  스토리지의 이름(자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}