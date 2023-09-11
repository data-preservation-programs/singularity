# 텐센트 클라우드 객체 스토리지 (COS)

{% code fullWidth="true" %}
```
명령어:
   singularity storage update s3 tencentcos - 텐센트 클라우드 객체 스토리지 (COS)

사용법:
   singularity storage update s3 tencentcos [command options] <name|id>

설명:
   --env-auth
      런타임으로부터 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터).
      
      access_key_id와 secret_access_key이 비어있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워둡니다.

   --secret-access-key
      AWS 비밀 액세스 키 (암호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워두거나 입력하세요.

   --endpoint
      텐센트 COS API endpoint입니다.

      예시:
         | cos.ap-beijing.myqcloud.com       | 베이징 지역
         | cos.ap-nanjing.myqcloud.com       | 난징 지역
         | cos.ap-shanghai.myqcloud.com      | 상하이 지역
         | cos.ap-guangzhou.myqcloud.com     | 광저우 지역
         | cos.ap-nanjing.myqcloud.com       | 난징 지역
         | cos.ap-chengdu.myqcloud.com       | 청두 지역
         | cos.ap-chongqing.myqcloud.com     | 충칭 지역
         | cos.ap-hongkong.myqcloud.com      | 홍콩 (중국) 지역
         | cos.ap-singapore.myqcloud.com     | 싱가포르 지역
         | cos.ap-mumbai.myqcloud.com        | 뭄바이 지역
         | cos.ap-seoul.myqcloud.com         | 서울 지역
         | cos.ap-bangkok.myqcloud.com       | 방콕 지역
         | cos.ap-tokyo.myqcloud.com         | 도쿄 지역
         | cos.na-siliconvalley.myqcloud.com | 실리콘밸리 지역
         | cos.na-ashburn.myqcloud.com       | 버지니아 지역
         | cos.na-toronto.myqcloud.com       | 토론토 지역
         | cos.eu-frankfurt.myqcloud.com     | 프랑크푸르트 지역
         | cos.eu-moscow.myqcloud.com        | 모스크바 지역
         | cos.accelerate.myqcloud.com       | 텐센트 COS 가속 엔드포인트 사용

   --acl
      버킷 및 객체 생성 및 복사 시 사용할 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 S3에서 server-side로 객체를 복사할 때 적용됩니다.
      S3는 원본에서 ACL을 복사하는 대신 새로운 ACL을 씁니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고, 기본 값(private)이 사용됩니다.
      

      예시:
         | default | 소유자가 Full_CONTROL 권한을 갖습니다.
         |         | 다른 사람은 액세스 권한이 없습니다 (기본값).

   --bucket-acl
      버킷 생성 시 사용할 Canned ACL입니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 모두 빈 문자열인 경우 X-Amz-Acl:
      헤더가 추가되지 않고, 기본 값(private)이 사용됩니다.
      

      예시:
         | private            | 소유자가 FULL_CONTROL 권한을 갖습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자가 FULL_CONTROL 권한을 갖습니다.
         |                    | AllUsers 그룹에 READ 액세스 권한을 부여합니다.
         | public-read-write  | 소유자가 FULL_CONTROL 권한을 갖습니다.
         |                    | AllUsers 그룹에 READ 및 WRITE 액세스 권한을 부여합니다.
         |                    | 버킷에 대해서 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자가 FULL_CONTROL 권한을 갖습니다.
         |                    | AuthenticatedUsers 그룹에 READ 액세스 권한을 부여합니다.

   --storage-class
      새로운 객체를 저장할 때 사용할 스토리지 클래스입니다.

      예시:
         | <unset>     | 기본값
         | STANDARD    | 표준 스토리지 클래스
         | ARCHIVE     | 아카이브 스토리지 모드
         | STANDARD_IA | 관람 빈도 접근 스토리지 모드

   --upload-cutoff
      청크 업로드로 전환되는 크기 제한입니다.
      
      이 크기보다 큰 파일은 chunk_size 단위로 청크 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드할 때 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 ("rclone rcat"이나 "rclone mount" 또는 구글
      사진 또는 구글 문서로부터 업로드된 파일)은 이 청크 크기를 사용하여 멀티파트 업로드됩니다.
      
      참고로, "--s3-upload-concurrency" 청크의 크기 단위는 이전과 동일한 크기입니다.
      
      높은 속도의 링크에서 큰 파일을 전송하고 메모리가 충분한 경우 이 값을 높이면 전송 속도가 향상됩니다.
      
      알려진 크기의 대용량 파일을 업로드할 때 rclone은 10,000개의 청크 제한을 지키도록 청크 크기를 자동으로 증가시킵니다.
      
      크기를 알 수 없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의 청크가
      있을 수 있으므로 기본적으로 스트림 업로드가 가능한 파일의 최대 크기는 48 GiB입니다.
      
      chunk 크기를 증가시키면 "-P" 플래그로 표시되는 진행 통계의 정확성이 감소합니다. rclone은 전송된 청크가
      실제로 업로드되는 것이 아니라 AWS SDK에 의해 버퍼링되어 전송된 것으로 처리하는데, 이 과정에서 진행률이 일정 부분 지체됩니다.
      chunk 크기가 크면 AWS SDK 버퍼가 크고 진행률 보고가 사실과 더 멀어집니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 부분 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 부분 수를 정의합니다.
      
      10,000개의 청크를 지원하지 않는 서비스에 유용할 수 있습니다.
      
      크기가 알려진 파일의 경우 rclone은 10,000개의 청크 제한을 지키기 위해 chunk 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      청크 복사로 전환되는 크기 제한입니다.
      
      이 크기보다 큰 서버 사이드로 복사해야 하는 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가하기 때문에 대용량 파일을
      업로드하기 전에 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth가 true인 경우, rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다.
      환경 변수 값이 비어 있다면 현재 사용자의 홈 디렉토리를 기본으로 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로파일입니다.
      
      env_auth가 true인 경우, rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 파일에서 사용할 프로파일을 제어합니다.
      
      비어있으면 환경 변수 "AWS_PROFILE" 또는 그 값이 설정되어 있지 않은 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      대역폭을 충분히 활용하지 못하는 고속링크에서 대량의 대용량 파일을 업로드하는 경우 이 값을 높이면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true이면 path style 엑세스를 사용하고, false이면 virtual hosted style을 사용합니다.
      
      true일 경우 rclone은 path style 엑세스를 사용하고, false일 경우 virtual path style을 사용합니다.
      자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 이 값을 false로 설정해야 합니다.
      rclone은 공급자 설정을 기반으로 자동으로 이 작업을 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)인 경우 rclone은 v4 인증을 사용합니다.
      설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예를 들어, Jewel/v10 CEPH 이전입니다.

   --list-chunk
      리스트 청크의 크기입니다 (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청보다 큰 것을 요청하더라도 응답 목록을 1000개로 잘라냅니다.
      AWS S3에서는 이것이 전역 최대값이므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 크기를 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 0(자동)입니다.
      
      S3가 처음 시작되었을 때 버킷에서 객체를 나열하기 위해 ListObjects 호출만 제공되었습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 provider에 설정된 대로 어떤 list objects 방법을 호출할지 추측합니다.
      그렇게 추측한 경우 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      리스트를 URL 인코딩해야 하는지 여부: true/false/unset
      
      일부 공급자는 URL 인코딩 목록을 지원하며, 파일 이름에 제어 문자를 사용할 때 더 안정적입니다.
      설정되지 않은 경우 (기본값) rclone은 제공자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않으려면 설정하세요.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용 중인 사용자가 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      v1.52.0 이전 버전에서는 버킷 존재 확인 테스트를 통과했습니다.

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다.
      
      rclone은 기본적으로 객체를 PUT한 후에 200 OK 메시지를 받으면 제대로 업로드된 것으로 간주합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터, 수정 시간, 스토리지 클래스 및 콘텐츠 유형이 업로드한 것과 동일함
      - 크기가 업로드한 것과 동일함
      
      다음 항목을 한 부분 PUT의 응답에서 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      대용량 파일의 올바른 크기를 확인하기 위해 HEAD 요청을 수행할 수 있습니다.
      
      이 플래그를 설정하면 업로드 실패 가능성이 높아지므로 정상적인 작업에는 권장되지 않습니다. 이 플래그가 설정된 경우에도
      업로드 실패 가능성은 매우 작습니다.

   --no-head-object
      GET을 하기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 개요의 [인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 주기입니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 메모리 풀을 사용하여 할당합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관한 미해결된 문제가 있습니다.
      s3 백엔드는 기본적으로 http2를 사용하지만 여기에서 사용을 비활성화할 수 있습니다.
      문제가 해결될 때까지이 플래그는 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 사용할 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는
      CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 이그레스 비용을 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용해야 하는지 여부
      
      true, false 또는 기본값을 사용하려면 true, false 또는 unset으로 설정하세요.
      

   --use-presigned-request
      단일파트 업로드에 사전 서명된 요청 또는 PutObject을 사용해야 하는지 여부
      
      이 값이 false이면 rclone은 객체 업로드에 AWS SDK의 PutObject를 사용합니다.
      
      rclone < 1.59 버전은 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하며,이 플래그를 true로 설정하면
      해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 위해 필요합니다.
      

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에서 파일 버전을 표시합니다.
      
      매개변수는 날짜 "2006-01-02", 날짜 및 시간 "2006-01-02
      15:04:05" 또는 그로부터 먼 시간을 나타내는 기간(예: "100d" 또는 "1h")일 수 있습니다.
      
      이를 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [time 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      Gzip으로 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"으로 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을
      압축 해제된 객체로 다운로드합니다.
      
      이 플래그를 설정하면 rclone은 "Content-Encoding: gzip"로 수신되는 파일을 분해 과정에서 이 객체를 압축 해제합니다.
      따라서 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는 경우 설정하세요.
      
      일반적으로 공급자는 객체가 다운로드될 때 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는
      다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급자는 "Content-Encoding: gzip"로 업로드되지 않은 객체를 gzip으로 압축할 수 있습니다(예: Cloudflare).
      
      다음과 같은 오류가 발생하는 것이 이러한 경우입니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 chunked 전송 인코딩으로 "Content-Encoding: gzip"가 설정된 객체를 다운로드하는 경우,
      rclone은 객체를 실시간으로 압축 해제합니다.
      
      unset(기본값)으로 설정한 경우 rclone은 제공자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                버킷 및 객체 생성 및 복사 시 사용할 Canned ACL. [$ACL]
   --endpoint value           텐센트 COS API endpoint. [$ENDPOINT]
   --env-auth                 런타임으로부터 AWS 자격 증명을 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS 비밀 액세스 키 (암호). [$SECRET_ACCESS_KEY]
   --storage-class value      새로운 객체를 저장할 때 사용할 스토리지 클래스입니다. [$STORAGE_CLASS]

   Advanced

   --bucket-acl value               버킷 생성 시 사용할 Canned ACL. [$BUCKET_ACL]
   --chunk-size value               업로드할 때 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크 복사로 전환되는 크기 제한입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     Gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 HTTP/2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 사용할 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 path style 엑세스를 사용하고, false이면 virtual hosted style을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기입니다 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          리스트를 URL 인코딩해야 하는지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0(자동)입니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로파일입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환되는 크기 제한입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용해야 하는지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          사전 서명된 요청 또는 PutObject을 사용해야 하는지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에서 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}