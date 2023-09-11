# Scaleway 객체 스토리지

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 scaleway - 스케일웨이 객체 스토리지

사용법:
   singularity storage create s3 scaleway [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명 가져 오기 (환경 변수 또는 env vars 또는 EC2/ECS 메타 데이터).
      
      access_key_id 및 secret_access_key가 비어 있을 때만 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 자격 증명을 입력하십시오.
         | true  | 런타임 환경에서 AWS 자격 증명 가져 오기 (env vars 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 두십시오.

   --secret-access-key
      AWS 비밀 액세스 키 (암호).
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 두십시오.

   --region
      연결할 지역입니다.

      예:
         | nl-ams | 네덜란드 암스테르담
         | fr-par | 프랑스 파리
         | pl-waw | 폴란드 와르샤와

   --endpoint
      Scaleway 객체 스토리지의 엔드포인트입니다.

      예:
         | s3.nl-ams.scw.cloud | 암스테르담 엔드포인트
         | s3.fr-par.scw.cloud | 파리 엔드포인트
         | s3.pl-waw.scw.cloud | 와르샤와 엔드포인트

   --acl
      버킷을 만들거나 오브젝트를 저장하거나 복사 할 때 사용되는 canned ACL.
      
      이 ACL은 오브젝트를 작성하는 데 사용되며, bucket_acl이 설정되지 않은 경우 버킷도 작성하는 데 사용됩니다.
      
      자세한 정보는 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      S3는 서버 측으로 객체를 복사 할 때 ACL을 복사하지 않고 새로 작성하기 때문에 이 ACL이 적용됩니다.
      
      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷을 만들 때 사용되는 canned ACL.
      
      자세한 정보는 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      
      예:
         | private            | 소유자는 FULL_CONTROL을 받습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자는 FULL_CONTROL을 받습니다.
         |                    | AllUsers 그룹은 읽기 액세스 권한을 받습니다.
         | public-read-write  | 소유자는 FULL_CONTROL을 받습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 권한을 받습니다.
         |                    | 버킷에 대해이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL을 받습니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스 권한을 받습니다.

   --storage-class
      S3에 새로운 객체를 저장할 때 사용할 저장 클래스입니다.

      예:
         | <unset>  | 기본값.
         | STANDARD | 요청시 스트리밍 또는 CDN과 같은 온디맨드 콘텐츠에 적합합니다.
         | GLACIER  | 아카이브 스토리지입니다.
         |          | 가격은 낮지만 사용하기 위해서는 복구되어야 합니다.

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이보다 큰 파일은 chunk_size로 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 알 수없는 크기의 파일(rclone rcat으로부터 또는 "rclone mount" 또는 Google 사진 또는 Google 문서에서 업로드 된 경우)을
      이 청크 크기를 사용하여 다중 파트 업로드로 업로드합니다.
      
      참고로, "--s3-upload-concurrency"는 해당 크기의 청크가 전송 당 메모리에 버퍼링되도록 합니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 늘리면 전송 속도가 향상됩니다.
      
      rclone은 알려진 크기의 큰 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 초과하지 않도록 합니다.
      
      알 수없는 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있으므로,
      기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 큰 파일을 스트림 업로드하려면 chunk_size를 크게 늘려야 합니다.
      
      chunk size가 증가하면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 낮아집니다. rclone은 청크가 AWS SDK에 의해 버퍼링될 때
      청크 전송된 것으로 처리하지만 실제로는 업로드 중인 경우도 있습니다. 청크 크기가 클수록 AWS SDK 버퍼가 크므로 진행률 표시가 진실과 
      더 멀어집니다.
      

   --max-upload-parts
      멀티 파트 업로드에서 사용할 패킷의 최대 수입니다.
      
      이 옵션은 멀티 파트 업로드를 수행 할 때 사용할 멀티 파트 청크의 최대 수를 정의합니다.
      
      서비스가 10,000 개의 청크에 대한 AWS S3 사양을 지원하지 않는 경우 유용합니다.
      
      rclone은 알려진 크기의 큰 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 해당치 이하의 청크 제한을 유지합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.
      
      서버 측에서 복사가 필요한 이 파일보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타 데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타 데이터에 추가합니다. 이는 데이터 무결성 
      확인에 좋지만 대용량 파일의 업로드 시작에 대한 긴 대기 시간을 일으킬 수도 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 기본값은 
      현재 사용자의 홈 디렉토리입니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용합니다. 이 변수는 해당 파일에서 사용될 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 기본값을 설정합니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티 파트 업로드의 동시성입니다.
      
      파일 청크의 동일한 수를 동시에 업로드합니다.
      
      대용량 파일을 빠른 속도로 고속 링크로 업로드하지만 대역폭을 완전히 활용하지 못하는 경우,
      이 값을 크게 증가시켜 전송 속도를 향상시킬 수 있습니다.

   --force-path-style
      true 인 경우 path 스타일 액세스를 사용하고 false 인 경우 가상 호스트 스타일 액세스를 사용합니다.
      
      이 값이 true(기본값) 인 경우 rclone은 path 스타일 액세스를 사용하고 false인 경우 가상 경로 스타일을 사용합니다. 
      자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 제공 업체(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 false로 설정해야 합니다.
      rclone은 이를 공급자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      이 값이 false(기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않을 때만 사용하십시오. 예를 들어, Jewel/v10 CEPH 이전에 사용합니다.

   --list-chunk
      목록 청크 크기 (각 ListObject S3 요청에 대한 응답 목록 크기)입니다.
      
      이옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 1000 개의 객체를 요청하더라도 응답 목록을 잘라냅니다.
      AWS S3에서는 전역 최대값이며 변경할 수 없으며, [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 더 크게 할 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동으로 설정하려면 0입니다.
      
      S3가 처음 시작될 때 버킷의 객체를 열거하기 위해 ListObjects 호출을 제공하기 시작했습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone이 공급자 설정에 따라 어떤 목록 객체 메서드를 호출할지 추측합니다. 추측이 잘못되면 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩 여부: true/false/unset
      
      일부 제공 업체는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다.
      이것이 사용 가능한 경우 파일 이름에 제어 문자를 사용하는 경우 보다 신뢰할 수 있습니다. 
      unset으로 설정되면 (기본값) rclone은 제공자 설정에 따라 적용할 항목을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 만들지 않으려면 설정하십시오.
      
      버킷이 이미 존재하는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.
      
      사용자가 버킷 생성 권한을 갖지 않은 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 무시되었습니다.
      

   --no-head
      업로드 된 객체의 무결성을 검사하기 위해 HEAD을 수행하지 않도록 설정하십시오.
      
      rclone은 업로드 후에 PUT로부터 200 OK 메시지를 수신하면 올바르게 업로드 된 것으로 가정합니다.
      
      특히 다음과 같이 가정합니다.
      
      - 업로드시 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형)가 업로드 된 것과 같음
      - 크기는 업로드 한 것과 같음
      
      다음을 파일의 단일 부분 PUT에 대한 응답에서 읽습니다.
      
      - MD5SUM
      - 업로드 된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      알 수없는 길이의 원본 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 높아지므로 일반적인 운영에는 권장되지 않습니다. 실제로 업로드 실패 가능성은 매우 
      적지만이 플래그가 있는 경우에도 실제로는 매우 작습니다.
      

   --no-head-object
      GET을 수행하기 전에 HEAD을 수행하지 않도록 설정하십시오.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시 될 때까지 걸리는 시간입니다.
      
      추가 버퍼가 필요한 업로드는 할당에 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 S3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. 
      S3 백엔드의 경우 HTTP/2는 기본적으로 사용되지만 여기에서 비활성화 할 수 있습니다.
      이 문제가 해결되면 이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를위한 사용자 정의 엔드 포인트입니다.
      AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 사용량을 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티 파트 업로드에서 ETag을 사용할지 여부
      
      true, false 또는 기본값을 사용하려면 true, false 또는 설정되지 않은 값을 사용합니다.
      

   --use-presigned-request
      단일 파트 업로드에 서명된 요청 또는 PutObject을 사용할지 여부
      
      false인 경우 rclone은 단일 부분 객체를 업로드하기 위해 AWS SDK에서 PutObject를 사용합니다.
      
      rclone 버전 < 1.59의 경우 서명된 요청을 사용하여 단일 부분 객체를 업로드하며이 플래그를 true로 설정하면
      해당 기능이 다시 활성화됩니다. 특별한 경우 또는 테스트하는 경우를 제외하고는 그리 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함시킵니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개 변수는 날짜, "2006-01-02", 날짜 및 시간 "2006-01-02
      15:04:05" 또는 해당 시간 전의 기간 (예 : "100d" 또는 "1h")이어야 합니다.
      
      이를 사용할 때는 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [시간 옵션 설명서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip으로 인코딩 된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip" 설정으로 객체를 업로드하는 것이 가능합니다. 보통 rclone은 이러한 파일을 
      압축된 객체로 다운로드합니다.
      
      이 플래그가 설정된 경우 rclone은 "Content-Encoding: gzip"로 이러한 파일을 수신하는 동안 이러한 파일을
      압축 해제합니다. 이는 rclone이 사이즈와 해시를 확인할 수 없게하지만 파일 내용은 압축 해제 됩니다.
      

   --might-gzip
      백엔드에서 gzip된 객체를 사용할 수 있습니다.
      
      일반적으로 업로드하지 않은 경우 공급업체는 객체를 수정하지 않습니다. "Content-Encoding: gzip"로 
      업로드되지 않은 경우 다운로드시 설정되지 않습니다.
      
      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 객체도 gzip으로 압축 할 수 있습니다
      (예 : Cloudflare 게이트웨이).
      
      이 경우 rclone이 chunked transfer encoding 및 "Content-Encoding: gzip"로 객체를 다운로드하는 것과 같은 
      오류를 수신하는 경우이 플래그를 설정하고 rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본값)로 설정된 경우 rclone은 제공자 설정에 따라 적용할 항목을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 만들거나 오브젝트를 저장하거나 복사 할 때 사용되는 canned ACL. [$ACL]
   --endpoint value           Scaleway 객체 스토리지의 엔드 포인트. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명 가져 오기 (환경 변수 또는 env vars 또는 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역입니다. [$REGION]
   --secret-access-key value  AWS 비밀 액세스 키 (암호). [$SECRET_ACCESS_KEY]
   --storage-class value      S3에 새로운 객체를 저장할 때 사용할 스토리지 클래스입니다. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              multipart 복사로 전환하는 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩 된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타 데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를위한 사용자 정의 엔드 포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 path 스타일 액세스를 사용하고 false 인 경우 가상 호스트 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기 (각 ListObject S3 요청에 대한 응답 목록 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동으로 설정하려면 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티 파트 업로드에서 사용할 패킷의 최대 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시 될 때까지 걸리는 시간. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip된 객체를 사용할 수 있습니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷이 존재하는지 확인하거나 만들지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드 된 객체의 무결성을 검사하기 위해 HEAD을 수행하지 않도록 설정하십시오. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD을 수행하지 않도록 설정하십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다. (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티 파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티 파트 업로드에서 ETag을 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 서명된 요청 또는 PutObject을 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함시킵니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}