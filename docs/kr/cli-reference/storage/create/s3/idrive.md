# IDrive e2

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 idrive - IDrive e2

사용법:
   singularity storage create s3 idrive [command options] [arguments...]

설명:
   --env-auth
      런타임(환경 변수 또는 env vars 또는 IAM 인증)에서 AWS 자격 증명을 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어있는 경우에만 적용됩니다.
      
      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 런타임(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --acl
      버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다.
      
      이 ACL은 객체를 만들 때 사용되며, bucket_acl 설정이 없는 경우에도 버킷을 만들 때 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3는 소스로부터 ACL을 복사하지 않고 새로 작성하므로 이 ACL은 서버 측 객체 복사시 적용됩니다.
      
      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 만들 때 사용되는 Canned ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl" 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자는 FULL_CONTROL을 획득합니다.
         |                    | 다른 사용자에게는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자는 FULL_CONTROL을 획득합니다.
         |                    | AllUsers 그룹은 읽기 액세스를 획득합니다.
         | public-read-write  | 소유자는 FULL_CONTROL을 획득합니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 획득합니다.
         |                    | 버킷에이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL을 획득합니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스를 획득합니다.

   --upload-cutoff
      청크 업로드로 전환하는 컷오프입니다.
      
      이보다 큰 파일은 chunk_size 단위로 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat"에서 가져온 파일이거나 "rclone mount" 또는 Google 사진 또는 Google 문서로 업로드된 파일)을 업로드할 때 이 청크 크기를 사용하여 다중 파트 업로드로 업로드됩니다.
      
      "--s3-upload-concurrency" 개수의 이 청크 크기가 전송당 메모리 버퍼에서 버퍼링됩니다.
      
      고속링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 높이면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드할 때 10,000청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      
      크기를 알 수없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야합니다.
      
      청크 크기를 늘리면 진행 상태의 정확도가 감소합니다. Rclone은 AWS SDK에 의해 버퍼링 된 청크가 보내진 것으로 간주하지만 실제로는 여전히 업로드 중 일 수 있습니다.
      청크 크기가 커지면 AWS SDK 버퍼도 커지므로 진행률에 대한 보고가 투명해집니다.
      

   --max-upload-parts
      다중 파트 업로드에서의 최대 파트 수입니다.
      
      이 옵션은 다중 파트 업로드시 사용할 멀티파트 청크 수를 정의합니다.
      
      10,000 청크의 AWS S3 사양을 지원하지 않는 서비스의 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 대형 파일을 업로드할 때 10,000청크 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 컷오프입니다.
      
      복사해야하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타 데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만 큰 파일에서 업로드를 시작하는 데 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 기본값을 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 그 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default" 환경 변수에 기본값을 사용합니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대역폭을 완전히 이용하지 못하는 고속링크에서 소수의 대형 파일을 업로드하고 있고 이러한 업로드가 너무 느린 경우, 이 값을 높이면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      참이면 경로 스타일 액세스를 사용하고, 그렇지 않으면 가상 호스팅 스타일 액세스를 사용합니다.
      
      이 값이 참인 경우(default), rclone은 경로 스타일 액세스를 사용하고, 그렇지 않으면 rclone은 가상 경로 스타일 액세스를 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는이 값이 false로 설정되어야합니다. rclone은이 값을 제공자 설정에 따라 자동으로 설정합니다.

   --v2-auth
      참이면 v2 인증을 사용합니다.
      
      이 값이 false인 경우(default), rclone은 v4 인증을 사용하고, true로 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 이 값을 사용하십시오. 예: Jewel/v10 CEPH 기준 버전 이전.

   --list-chunk
      목록 청크의 크기(ListObject S3 요청당 응답 목록)입니다.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청한 것보다 많은 객체의 응답 목록을 1000개로 자름니다.
      AWS S3에서는 이것이 전역 최대값으로 설정되어 있고 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전: 1, 2 또는 자동으로 설정하려면 0.
      
      S3가 처음 출시 될 때 버킷의 객체를 열거하기 위해 ListObjects 호출 만 제공했습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하며 가능하면 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출 할 목록 개체 방법을 추측합니다. 추측이 잘못되면 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 제공자는 URL 인코딩 목록을 지원하며, 사용 가능한 경우 파일 이름에서 제어 문자를 사용할 때 신뢰할 수 있습니다. 이 값이 unset(default)로 설정되어 있으면 rclone은 제공자 설정에 따라 적용할 것을 선택하지만, 여기에서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-check-bucket
      버킷 존재 여부를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      이것은 버킷 만들기 권한이없는 사용자를 사용하는 경우에도 필요할 수 있습니다. v1.52.0 이전에는이 버그 때문에 정상적으로 전달되지 않았을 것입니다.
      

   --no-head
      업로드된 객체의 정합성을 확인하기 위해 HEAD를 실행하지 않습니다.
      
      rclone은 적합한 후 200 OK 메시지를 받으면 PUT로 객체를 업로드 한 후에 제대로 업로드 된 것으로 간주합니다.

      특히 다음을 가정합니다:
      
      - 업로드 된 것과 동일한 메타 데이터, 수정 시간 및 저장 클래스, 콘텐츠 유형이었습니다.
      - 업로드 된 크기가었습니다.
      
      다음 항목을 단일 파트 이미지의 응답에서 읽습니다:

      - MD5SUM,
      - 업로드 날짜
      
      멀티 파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      알 수없는 길이의 원본 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 잘못된 크기와 같은 감지되지 않은 업로드 오류가 더 커질 수 있으므로 정상 운영에는 권장되지 않습니다. 실제로 이 플래그로 인한 감지되지 않은 업로드 오류의 가능성은 매우 적습니다.
      

   --no-head-object
      GET을 실행하기 전에 HEAD를 실행하지 마십시오.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 메모리 풀에 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드의 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. S3 백엔드의 HTTP/2는 기본적으로 사용됩니다. 하지만 여기에서 비활성화 할 수 있습니다. 문제가 해결되면이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 사용되는 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 클라우드 프론트 CDN URL로 설정된 경우 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 이그레스를 제공합니다.

   --use-multipart-etag
      확인을 위해 multipart 업로드에서 ETag를 사용할지 여부
      
      이 값은 true, false 또는 제공자의 기본값을 사용하도록 설정되거나 unset이어야합니다.
      

   --use-presigned-request
      단일 파트 업로드에 서명 된 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false이면 rclone은 객체를 업로드하는 데 AWS SDK의 PutObject를 사용합니다.
      
      rclone < 1.59의 버전은 단일 파트 개체를 업로드하기 위해 서명 된 요청을 사용하고이 플래그를 true로 설정하면이 기능을 다시 활성화합니다. 이는 특정 상황이나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함시킵니다.

   --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개 변수는 날짜 "2006-01-02", 날짜 및 시간 "2006-01-02 15:04:05" 또는 그 이후로 오래 된 시간에 대한 지속 시간 "100d" 또는 "1h" 일 수 있습니다.
      
      이 값을 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      올바른 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      Gzip로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"이 설정된 상태로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 이러한 파일은 압축 된 객체로 다운로드됩니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"로 받아들입니다. 이는 rclone이 크기와 해시를 확인 할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 객체를 gzip으로 압축 할 수 있으므로이 값을 설정하십시오.
      
      일반적으로 공급자는 개체를 다운로드 할 때 개체를 수정하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않았다면 다운로드시에도 설정되지 않습니다.
      
      그러나 일부 공급자(예 : Cloudflare)는 `Content-Encoding : gzip`로 업로드되지 않은 개체에 대해서도 gzip으로 압축 할 수 있습니다.
      
      이로 인해 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 `Content-Encoding : gzip`가 설정되고 청크 전송 인코딩을 사용하여 개체를 다운로드하면 rclone은 그 즉시 개체를 해제합니다.
      
      이 값이 unset(default)로 설정되어 있으면 rclone은 제공자 설정에 따라 적용할 것을 선택하지만, 여기에서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL. [$ACL]
   --env-auth                 런타임(환경 변수 또는 EC2/ECS 메타데이터)에서 AWS 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS Secret Access Key (password). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 컷오프 크기. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     Gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 사용되는 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               참이면 경로 스타일 액세스를 사용하고, 그렇지 않으면 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기 (각 ListObject S3 요청당 응답 목록 크기입니다). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전 : 1, 2 또는 0이면 자동으로 설정됩니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         다중 파트 업로드에서의 최대 파트 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 객체를 gzip으로 압축 할 수 있으므로이 값을 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        HEAD를 실행하지 않고 업로드된 객체의 정합성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 실행하기 전에 HEAD를 실행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 컷오프 크기. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 multipart 업로드에서 ETag를 사용할지 여부. (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          서명 된 요청 또는 PutObject를 사용하여 단일 파트 업로드할지 여부. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        참이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간대로 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       이전 버전을 디렉토리 목록에 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}