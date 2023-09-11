# IONOS Cloud

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 ionos - IONOS Cloud

사용법:
   singularity storage create s3 ionos [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 받아옵니다(환경 변수나 env vars 또는 EC2/ECS 메타 데이터).
      
      access_key_id와 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명을 받아옵니다(환경 변수나 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두세요.

   --secret-access-key
      AWS 비밀 액세스 키(암호)입니다.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 두세요.

   --region
      버킷이 생성되고 데이터가 저장될 지역입니다.

      예제:
         | de           | 독일 프랑크푸르트
         | eu-central-2 | 독일 베를린
         | eu-south-2   | 스페인 로그로뇨

   --endpoint
      IONOS S3 객체 저장소의 엔드포인트입니다.

      동일한 지역의 엔드포인트를 지정하세요.

      예제:
         | s3-eu-central-1.ionoscloud.com | 독일 프랑크푸르트
         | s3-eu-central-2.ionoscloud.com | 독일 베를린
         | s3-eu-south-2.ionoscloud.com   | 스페인 로그로뇨

   --acl
      버킷 생성 및 객체 저장 또는 복사 시에 사용되는 기본 ACL입니다.

      이 ACL은 객체를 생성할 때 사용되며, bucket_acl이 설정되지 않은 경우에도 버킷 생성에 사용됩니다.

      자세한 내용은 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하세요.

      참고로, S3는 서버 측 복사 중에 원본의 ACL을 복사하지 않고, 새로운 ACL을 작성합니다.

      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본(개인)이 사용됩니다.

   --bucket-acl
      버킷을 생성할 때 사용되는 기본 ACL입니다.

      자세한 내용은 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)을 참조하세요.

      참고로, 이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않은 경우에는 "acl"이 대신 사용됩니다.

      "acl"과 "bucket_acl"이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본(개인)이 사용됩니다.

      예제:
         | private            | 소유자가 FULL_CONTROL을 받습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자가 FULL_CONTROL을 받습니다.
         |                    | AllUsers 그룹이 읽기 액세스 권한을 받습니다.
         | public-read-write  | 소유자가 FULL_CONTROL을 받습니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 액세스 권한을 받습니다.
         |                    | 일반적으로 버킷에 대해 이 권한을 허용하는 것은 권장되지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL을 받습니다.
         |                    | AuthenticatedUsers 그룹이 읽기 액세스 권한을 받습니다.

   --upload-cutoff
      청크 업로드로 전환하는 크기 임계값입니다.

      이보다 큰 파일은 청크 크기 단위로 업로드됩니다.
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 크거나 알려지지 않은 크기의 파일(예: "rclone rcat"으로 "rclone mount" 또는 Google Photos 또는 Google Docs로 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.

      --s3-upload-concurrency 과 chunk의 크기는 하나의 파일 전송당 메모리 버퍼에 버퍼링됩니다.

      고속 링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 증가시키면 전송 속도가 향상될 수 있습니다.

      rclone은 알려진 크기의 대형 파일을 전송할 때 10,000 청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.

      알려지지 않은 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000 청크까지 존재할 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림으로 업로드하려면 청크 크기를 증가시켜야 합니다.

      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행률 통계의 정확도가 감소합니다. rclone은 실제로 업로드 중인 청크를 버퍼링하는 경우 chunk를 보낸 것으로 처리하지만, 실제로 업로드 중인 경우일 수도 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 통계의 차이가 더 작은(Bigger) 청크 크기를 의미합니다.

   --max-upload-parts
      멀티파트 업로드에서 사용되는 최대 파트 수입니다.

      이 옵션은 멀티파트 업로드를 수행할 때 사용할 파트의 최대 수를 정의합니다.

      AWS S3의 10,000 청크 사양을 지원하지 않는 경우에 유용할 수 있습니다.

      rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜이 파트 수 제한을 유지할 수 있습니다.

   --copy-cutoff
      멀티파트 복사로 전환하는 크기 임계값입니다.

      이보다 큰 파일은 이 크기로 청크 복사본을 복사합니다.
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      보통 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이렇게 함으로써 데이터의 무결성을 확인할 수 있지만, 큰 파일의 경우 업로드 시작에 오랜 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      환경 변수 "AWS_SHARED_CREDENTIALS_FILE"에 값이 없는 경우 rclone은 다음과 같이 찾습니다.
      - Linux/OSX: "$HOME/.aws/credentials"
      - Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.

      비워 두면 환경 변수 "AWS_PROFILE" 또는 "default" (환경 변수가 설정되어 있지 않은 경우)로 지정됩니다.

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드를 위한 동시성입니다.

      동일한 파일의 멀티파트 청크 수를 동시에 업로드합니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.

      기본값인 true이면 rclone은 경로 스타일 액세스를 사용하고, false이면 가상 경로 스타일을 사용합니다.
      자세한 내용은 [AWS S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.

      일부 제공자(AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 이 값을 false로 설정해야 합니다. rclone은 공급자 설정에 따라 자동으로 설정합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.

      기본값인 false이면 rclone은 v4 인증을 사용합니다. 설정된 경우 v2 인증을 사용합니다.

      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: Jewel/v10 CEPH에서 사용합니다.

   --list-chunk
      목록 청크 크기(ListObject S3 요청마다 응답 목록 크기)입니다.

      이 옵션은 AWS S3 사양의 MaxKeys, max-items 또는 page-size로 알려져 있습니다.
      대부분의 서비스는 1000개 이상이 요청되었더라도 응답 목록을 1000개로 종료합니다.
      AWS S3에서는 이 값이 전역 최대 값이므로 변경할 수 없습니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.

   --list-version
      사용할 ListObjects 버전입니다: 1,2 또는 0(자동 설정).

      S3가 처음 출시될 때 버킷의 객체를 열거하는 ListObjects 호출만 제공했습니다.

      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하므로 가능한 경우 사용해야 합니다.

      기본값인 0으로 설정된 경우 rclone은 공급자 설정에 따라 호출할 목록 객체 메서드를 추측합니다. 예상과 다르게 추측한 경우 여기에서 수동으로 설정할 수 있습니다.

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset

      일부 제공자에서는 이러한 목록에 대한 URL 인코딩을 지원하며, 파일 이름에 제어 문자를 사용할 때 이렇게 인코딩하는 것이 더 신뢰할 수 있습니다. 설정되지 않은 경우(기본값) rclone은 공급자 설정에 따라 적용할 항목을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.

   --no-check-bucket
      버킷을 확인하거나 생성하지 않습니다.

      버킷이 이미 존재하는 경우 rclone의 트랜잭션 수를 최소화하기 위해 이 옵션을 사용할 수 있습니다.

      사용자당 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 정상적으로 전달되었습니다.

   --no-head
      업로드된 객체의 정합성을 확인하기 위해 HEAD를 하지 않습니다.

      rclone은 이후 PUT으로 객체를 업로드한 후에 200 OK 메시지를 수신하면 제대로 업로드된 것으로 가정합니다.

      특히 다음을 가정합니다.
      
      - 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형포함)이 업로드한 것과 동일하게 됩니다.
      - 크기가 업로드한 것과 동일합니다.

      단일 청크 PUT에 대한 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드 날짜

      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.

      크기를 알 수 없는 소스 개체가 업로드된 경우 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면, 올바르지 않은 크기로 인해 업로드 실패의 가능성이 커지므로, 정상적인 운영에는 권장되지 않습니다. 실제로 이 플래그를 사용하여 업로드 실패 가능성은 매우 낮습니다.

   --no-head-object
      객체를 가져오기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.

      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.

      현재 s3(특히 minio) 백엔드와 HTTP/2에 관련된 해결되지 않은 문제가 있습니다. HTTP/2는 기본값으로 설정되어 있지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      
      AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 이그레스를 제공하기 때문에 일반적으로 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부입니다.

      true, false 또는 unset 이어야 합니다.

      기본값을 사용하려면 unset으로 설정하세요.

   --use-presigned-request
      단일 파트 업로드에 대해 프리서인 단청 요청 또는 PutObject를 사용할지 여부입니다.

      false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.

      rclone 버전 < 1.59는 단일 파트 개체를 업로드하기 위해 프리시니드 요청을 사용하고, 이 플래그를 true로 설정하면 해당 기능을 다시 활성화합니다. 특수한 경우나 테스트 외에는 이 기능이 필요하지 않습니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에서 파일 버전을 표시합니다.

      매개변수는 "2006-01-02", "2006-01-02 15:04:05"와 같은 날짜, "100d" 또는 "1h"와 같은 이틀 이전의 지속 시간일 수 있습니다.

      이렇게 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      유효한 형식에 대한 자세한 내용은 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.

   --decompress
      gzip으로 인코딩된 개체를 압축 해제합니다.

      "Content-Encoding: gzip"으로 S3에 객체를 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.

      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"으로 받는 대로 압축을 해제합니다. 즉, rclone은 크기 및 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

   --might-gzip
      백엔드에서 개체를 gzip으로 압축할 수 있는 경우 이를 설정하세요.

      공급자는 일반적으로 다운로드될 때 개체를 수정하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 경우 다운로드되지 않습니다.

      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 개체도 Gzip으로 압축할 수 있습니다(Cloudflare 등).

      이를 위해 설정된 경우 rclone은 "Content-Encoding: gzip"으로 설정된 개체와 청크 전송 인코딩으로 개체를 다운로드할 때 개체를 실시간으로 압축 해제합니다.

      unset로 설정되면 rclone은 공급자 설정에 따라 무엇을 적용할지 선택합니다. rclone의 선택을 여기에서 재정의할 수 있습니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


옵션:
   --access-key-id value      AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                버킷 생성 및 객체 저장 또는 복사 시에 사용되는 기본 ACL입니다. [$ACL]
   --endpoint value           IONOS S3 객체 저장소의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명을 받아옵니다(환경 변수나 EC2/ECS 메타 데이터에서 env vars가 비어 있는 경우). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             버킷이 생성되고 데이터가 저장될 지역입니다. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키(암호)입니다. [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 생성할 때 사용되는 기본 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 크기 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기(ListObject S3 요청마다 응답 목록 크기)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전입니다: 1,2 또는 0(자동 설정). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용되는 최대 파트 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 개체를 gzip으로 압축할 수 있는 경우 이를 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷을 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 정합성을 확인하기 위해 HEAD를 하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져오기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드를 위한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 크기 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부입니다 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 프리서인 단청 요청 또는 PutObject를 사용할지 여부입니다. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        v2 인증을 사용할지 여부입니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에서 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   일반

   --name value  스토리지의 이름(자동 생성됨) (기본값: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}