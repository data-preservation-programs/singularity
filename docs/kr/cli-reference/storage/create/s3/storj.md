# Storj (S3 호환 게이트웨이)

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 storj - Storj (S3 호환 게이트웨이)

사용법:
   singularity storage create s3 storj [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경 변수가 없으면 EC2/ECS 메타 데이터에서 가져옴).
      
      access_key_id와 secret_access_key가 비어 있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명 입력.
         | true  | 환경에서 AWS 자격 증명 가져오기 (환경 변수 또는 IAM).

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격 증명인 경우 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명인 경우 비워 둡니다.

   --endpoint
      Storj Gateway의 엔드포인트.

      예시:
         | gateway.storjshare.io | 글로벌 호스팅 게이트웨이

   --bucket-acl
      버킷을 생성할 때 사용되는 Canned ACL.
      
      자세한 정보는 [아마존 S3 개발자 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      이 ACL은 버킷을 생성할 때에만 적용됩니다. 설정되지 않으면 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl" 모두 빈 문자열인 경우에는 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사용자는 액세스 권한을 부여받지 못합니다(기본 설정).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 READ 권한을 부여합니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 권한을 부여합니다.
         |                    | 버킷에 대한 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AuthenticatedUsers 그룹에게 READ 권한을 부여합니다.

   --upload-cutoff
      청크 업로드로 전환되는 파일의 임계값.
      
      이보다 큰 크기의 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat"으로부터, "rclone mount"로 업로드된 파일이나 구글 포토, 구글 문서와 같은 파일 등)은 이 청크 크기를 사용하여 분할 업로드됩니다.
      
      참고로 "--s3-upload-concurrency" 크기의 청크는 각 전송별로 메모리에 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 높이면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 최대 10,000개의 청크 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본적인 청크 크기는 5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로, 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼링되는 경우 청크가 전송된 것으로 처리하지만, 실제로 업로드 중인 경우도 있을 수 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 보고서를 혼란스럽게 만듭니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용하는 최대 청크 수.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      AWS S3의 10,000개 청크 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 최대 청크 수 제한을 초과하지 않도록 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      서버 사이드 복사에 전환되는 파일의 임계값.
      
      이보다 큰 파일을 서버 사이드에서 복사해야 할 경우 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에는 유용하지만, 대용량 파일을 업로드할 때 오랜 시간 지연을 발생시킬 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 기본값이 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시 처리 수.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대용량 파일을 고속 연결로 작은 양의 업로드하며 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 늘리면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true이면 path 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.
      
      이 값이 true(기본값)이면 rclone은 path 스타일 액세스를 사용하며, false이면 가상 path 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 제공업체(AWS, Aliyun OSS, Netease COS, Tencent COS 등)는 이 값을 false로 설정해야 합니다. rclone은 제공업체 설정에 따라 자동적으로 이 작업을 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정되었으면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용합니다(Jewel/v10 CEPH 이전).

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청된 것보다 많은 항목을 요청해도 응답 목록을 1000개로 자르게 됩니다.
      AWS S3에서 이것은 전역적인 최대값이므로, [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동(0).
      
      S3가 처음 출시될 때는 버킷의 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하므로 가능한 경우 사용해야 합니다.
      
      기본 설정인 0으로 설정되어 있으면 rclone은 제공자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못 추측하면 여기서 직접 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 제공 업체에서는 목록을 URL 인코딩하는 것을 지원하며 파일 이름에 제어 문자를 사용할 때에는 이 방법이 더 안정적입니다. 쿼리 매개변수인지 또는 HTTP 메소드 항목의 JSON 필드인지에 따라 URL 인코딩이 사용되는 경우가 있습니다. unset(기본값)로 설정되어 있으면 rclone은 프로바이더 설정에 따라 적용할 내용을 선택합니다.

   --no-check-bucket
      설정하면 버킷이 존재하는지 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용합니다.
      
      사용 중인 사용자에게 버킷 생성 권한이 없는 경우 필요할 수도 있습니다. v1.52.0 이전 버전에서는 버그 때문에 이 작업이 무시되었습니다.
      

   --no-head
      설정하면 업로드된 객체의 정합성을 확인하기 위해 HEAD 요청을 수행하지 않습니다.
      
      rclone은 수행하는 트랜잭션 수를 최소화하려는 경우에 유용합니다.
      
      이 플래그를 설정하면 PUT로 객체를 업로드한 후 200 OK 메시지를 수신하면 제대로 업로드되었다고 가정합니다.
      
      특히 다음과 같은 가정을 합니다:
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)이 업로드한 대로임.
      - 크기가 업로드한 대로임.
      
      다음 항목을 단일 부분 PUT의 응답에서 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      알려지지 않은 길이의 원본 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패가 감지되지 않을 확률이 높아지므로 일반적인 운영에는 권장되지 않습니다. 실제로 이 플래그로 인해 업로드 실패가 감지되지 않는 확률은 매우 낮습니다.
      

   --no-head-object
      설정하면 GET하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 언제 플러시할지 지정합니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당에 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에 대한 http2 사용 비활성화.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. s3 백엔드는 기본적으로 HTTP/2를 사용하지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드용 사용자 설정 엔드포인트.
      AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 보다 저렴한 이그레스를 제공하는데, 일반적으로 이에 설정됩니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 검증에 사용할지 여부
      
      이 값은 true, false 또는 프로바이더의 기본값으로 설정되어야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 대해 선행 서명 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false이면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone의 버전 1.59 이전 버전은 단일 파트 객체를 업로드하기 위해 선행 서명 요청을 사용하고 이 플래그를 true로 설정하면 그 기능이 다시 활성화됩니다. 이 기능은 예외적인 상황이나 테스트를 위해서만 필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함할지 여부.

   --version-at
      파일 버전을 지정된 시간에 표시합니다.
      
      매개변수는 날짜, "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그만큼 오래된 기간인 "100d" 또는 "1h"와 같을 수 있습니다.
      
      이를 사용하는 동안 파일 쓰기 작업이 허용되지 않기 때문에 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식은 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      설정하면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"이 설정된 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 수신되는 데이터에 "Content-Encoding: gzip"와 함께 이 파일을 압축 해제합니다. 이렇게되면 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip할 수 있으면 설정하십시오.
      
      일반적으로 공급자들은 다운로드될 때 객체를 수정하지 않습니다. "Content-Encoding: gzip"과 업로드되지 않은 객체에 설정하지 않았다면 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급 업체(예: Cloudflare)는 "Content-Encoding: gzip"과 업로드되지 않은 객체에 대해서도 gzip으로 압축할 수 있습니다.
      
      이 경우 rclone이 청크 전송 인코딩과 "Content-Encoding: gzip"이 설정된 객체를 다운로드하면 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본값)로 설정되어 있으면 rclone은 프로바이더 설정에 따라 적용할 내용을 선택합니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정과 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --endpoint value           Storj Gateway의 엔드포인트. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경 변수가 없으면 EC2/ECS 메타 데이터에서 가져옴). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS Secret Access Key(비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 복사되기 시작할 크기. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정하면 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드용 사용자 설정 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 path 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1, 2 또는 자동(0). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용하는 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 언제 플러시할지 지정합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip할 수 있으면 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정하면 버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        설정하면 업로드된 객체의 정합성을 확인하기 위해 HEAD 요청을 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 설정하면 GET하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정과 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시 처리 수. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환되는 파일의 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 검증에 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 선행 서명 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간에 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함할지 여부. (기본값: false) [$VERSIONS]

   General

   --name value  저장소의 이름(자동 생성됨)
   --path value  저장소의 경로

```
{% endcode %}