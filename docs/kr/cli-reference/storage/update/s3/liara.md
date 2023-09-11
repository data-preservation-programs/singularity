# Liara Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 liara - 리아라 오브젝트 스토리지

사용법:
   singularity storage update s3 liara [옵션] <이름|ID>

설명:
   --env-auth
      실행 중인 환경에서 AWS 자격 증명을 가져옵니다. (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id와 secret_access_key가 비어있으면 적용됩니다.

      예제:
         | false | AWS 자격 증명을 다음 단계에서 입력하세요.
         | true  | 실행 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 엑세스 키 ID입니다.
      
      익명 액세스 또는 실행 중인 자격 증명인 경우 비워둡니다.

   --secret-access-key
      AWS 비밀 액세스 키입니다.
      
      익명 액세스 또는 실행 중인 자격 증명인 경우 비워둡니다.

   --endpoint
      리아라 오브젝트 스토리지 API의 엔드포인트입니다.

      예제:
         | storage.iran.liara.space | 기본 엔드포인트
         |                          | 이란

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다.
      
      이 ACL은 개체를 작성할 때 및 bucket_acl이 설정되지 않은 경우 버킷을 작성할 때 모두 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      S3가 ACL을 복사하기보다는 새로운 ACL을 작성할 때 이 ACL이 적용됩니다.
      
      acl이 비어있는 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본 (private)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 canned ACL입니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본 (private)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 READ 액세스권을 부여합니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스권을 부여합니다.
         |                    | 버킷에 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스 권한을 부여합니다.

   --storage-class
      Liara에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다.

      예제:
         | STANDARD | 표준 스토리지 클래스

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이보다 큰 크기의 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 (예: "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일)은 이 청크 크기를 사용하여
      멀티파트 업로드로 업로드됩니다.
      
      주의: "--s3-upload-concurrency" 크기의 청크는 전송 당 메모리에 버퍼링됩니다.
      
      고속 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우 청크 크기를 늘리면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드할 때 청크 크기를 자동으로 증가시켜
      10,000 청크 제한을 초과하지 않도록합니다.
      
      알 수 없는 크기의 파일은 구성된
      청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대
      10,000 청크를 사용할 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는
      48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 늘려야합니다.
      
      청크 크기를 늘리면 "-P" 플래그와 함께 표시되는 진행 상태의 정확도가 감소합니다. Rclone은 청크를 전송으로 처리합니다.
      AWS SDK에서 버퍼링 될 때, 실제로 업로드 중인 경우로 표시될 수 있습니다.
      청크 크기가 클수록 AWS SDK 버퍼와 진행 상태가 정확도가 떨어지게됩니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드을 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      10,000 개의 청크로 지정된 AWS S3 사양을 지원하지 않는 서비스의 경우 유용할 수 있습니다.
      
      알려진 크기의 대용량 파일을 업로드할 때 Rclone은 청크 크기를 자동으로 증가시켜
      이 청크 수 제한을 초과하지 않도록합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값입니다.
      
      이보다 큰 파일은 서버 측에서 복사해야하는 경우이 크기로 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타 데이터에 추가합니다.
      이는 데이터 무결성 검사에 좋지만 대용량 파일의 업로드 시작을 위해 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      변수가 비어있는 경우 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면
      현재 사용자의 홈 디렉터리로 기본값으로 설정합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 기본값으로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 사용되는 동시성 수입니다.
      
      대용량 파일을 고속 링크로 업로드하는 경우가 아니고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우
      이 값을 늘려 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스트 스타일을 사용합니다.
      
      true이면 rclone은 경로 스타일 액세스를 사용하고
      false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 공급자 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이를 요청에 따라 비활성화합니다.
      기본값으로 rclone이 이를 자동으로 수행합니다.
      Provider 설정에 따라.
      

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      즉 v4 서명이 작동하지 않는 경우에만 사용하세요. (예: Jewel/v10 CEPH 이전).

   --list-chunk
      목록 청크 크기 (각 ListObject S3 요청에 대한 응답 목록)입니다.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청한 것보다 더 큰 응답 목록을 1000 개의 개체로 자르지만
      AWS S3에서는 글로벌 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 0은 자동으로 됩니다.
      
      S3가 처음 출시되었을 때 버킷의 객체를 열거하기 위해 ListObjects 호출만 사용할 수 있었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하며 가능하면 사용해야합니다.
      
      자동으로 설정되는 기본값인 0으로 설정되어 있으면 rclone은 공급자에 따라
      호출할 목록 개체 방법을 추측합니다. 추측이 잘못되면
      여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩 할지 여부: true/false/미설정
      
      일부 공급자는 URL 인코딩 목록을 지원하며 사용 가능한 경우 파일에서 컨트롤 문자를 사용할 때이 방법이 
      더 안정적입니다. 이 값은 unset으로 설정될 경우 (기본값) rclone은
      공급자 설정에 따라 무엇을 적용할지 선택하지만 여기에서 rclone의 선택을 무시 할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 있는 경우 rclone이 수행하는 거래 수를 최소화하려는 경우에 유용합니다.
      
      또한 사용자가 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에는이 버그로 인해 음부 전달되었어야했습니다.
      문제없이 전달되었습니다.
      

   --no-head
      업로드 된 객체의 정합성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone은 청크로 업로드 된 후 PUT로부터 200 OK 메시지를 수신하면 제대로 업로드된 것으로 간주합니다.
      
      특히 다음을 가정합니다.
      
      - 메타 데이터 (모드 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드와 동일합니다.
      - 크기가 업로드된 것입니다.
      
      다음 항목을 한부분 PUT의 응답에서 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      알 수없는 길이의 소스 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 가능성이 높아지며,
      특히 잘못된 크기로 인한 것이므로 정상적인 운영에는 권장하지 않습니다. 실제로 미검출
      전송 오류의 가능성은이 플래그를 사용하지 않아도 매우 낮습니다.
      

   --no-head-object
      개체를 가져올 때 HEAD를 실행하지 마십시오.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 어느 정도 자주 플러시할지 제어합니다.
      
      추가 버퍼가 필요한 업로드 (예: 멀티파트)는 할당에 메모리 풀을 사용할 것입니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 관련된 해결되지 않은 문제가 있습니다. 
      S3 백엔드에서는 기본적으로 HTTP/2가 활성화되지만 여기에서 비활성화 할 수 있습니다.
      문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      주로 AWS S3는
      CloudFront 네트워크를 통해 다운로드 된 데이터의 더 저렴한 Egress를 제공합니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 확인할지 여부
      
      true, false 또는 기본값으로 잡아 둘 수 있습니다. 공급자에 대한 기본값을 사용하려면
      설정되지 않은 경우입니다.
      

   --use-presigned-request
      단일 파일 업로드에 미리 서명 된 요청 또는 PutObject를 사용할지 여부
      
      false이면 rclone은 AWS SDK의 PutObject를 사용하여 개체를 업로드합니다.
      
      rclone의 버전 < 1.59은 단일
      파트 개체를 업로드하기 위해 미리 서명된 요청을 사용하며이 플래그를 true로 설정하면
      그 기능을 다시 활성화합니다. 이는 예외적인 경우나 테스트 외에도 필요하지 않은 경우입니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜, "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그 이전의 기간, 예를 들어 "100d" 또는 "1h"일 수 있습니다.
      
      이 옵션을 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제 할 수 없습니다.
      
      사용 가능한 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩 된 개체를 해제합니다.
      
      "Content-Encoding: gzip"로 설정하여 S3에 객체를 업로드 할 수 있습니다. 일반적으로 rclone은 이러한 파일을
      압축된 객체로 다운로드합니다.
      
      이 플래그가 설정된 경우 rclone은 수신되는 "Content-Encoding: gzip"으로 이러한 파일을 해제합니다.
      이렇게하면 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 객체를 gzip으로 압축 할 수있는 경우이를 설정합니다.
      
      일반적으로 제공자는 객체를 다운로드 할 때 객체를 수정하지 않습니다. 
      객체가 `Content-Encoding: gzip`로 업로드되지 않았다면 다운로드될 때 설정되지 않을 것입니다.
      
      그러나 일부 제공자는 개체가 `Content-Encoding: gzip`로 업로드되지 않았더라도 객체를 gzip으로 압축 할 수 있습니다
      (예: Cloudflare).
      
      이 경우 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩으로 개체를 다운로드 한 경우
      rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset으로 설정되는 경우 (기본값) rclone은
      공급자 설정에 따라 무엇을 적용할지 선택하지만 여기에서 rclone의 선택을 무시 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타 데이터의 설정 및 읽기를 억제합니다.

옵션:
   --access-key-id value      AWS 엑세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 canned ACL입니다. [$ACL]
   --endpoint value           리아라 오브젝트 스토리지 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 실행 중인 환경에서 AWS 자격 증명을 가져옵니다. (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS 비밀 액세스 키입니다. [$SECRET_ACCESS_KEY]
   --storage-class value      Liara에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다. [$STORAGE_CLASS]

   고급

   --bucket-acl value               버킷을 생성할 때 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코드 된 개체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스트 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기 (각 ListObject S3 요청에 대한 응답 목록)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩 할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0은 자동으로 됩니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 플러시 할 것인지 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 객체를 gzip으로 압축 할 수있는 경우이를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드 된 객체의 정합성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져올 때 HEAD를 실행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타 데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 사용되는 동시성 수입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 확인할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파일 업로드에 미리 서명 된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}