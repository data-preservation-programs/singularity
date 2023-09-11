# Storj (S3 호환 게이트웨이)

```bash
사용법:
   singularity storage update s3 storj [command options] <name|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명 가져오기 (환경 변수나 env 변수 또는 EC2/ECS 메타데이터에서).
      
      access_key_id와 secret_access_key가 비어 있을 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명 가져오기.

   --access-key-id
      AWS 엑세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --secret-access-key
      AWS 시크릿 액세스 키(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

   --endpoint
      Storj Gateway 엔드포인트.

      예제:
         | gateway.storjshare.io | 글로벌 호스팅 게이트웨이

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 내용은 [Amazon S3](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하십시오.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다. 설정되지 않으면 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사용자는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 읽기 권한이 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 권한을 부여 받습니다.
         |                    | 버킷에 대해 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹이 읽기 권한을 부여받습니다.

   --upload-cutoff
      청크 업로드로 전환할 파일의 크기 기준.
      
      이 크기보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 큰 파일이나 크기를 모르는 파일(예: "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일)은 이 청크 크기를 사용하여
      멀티파트 업로드로 업로드됩니다.
      
      참고로 "--s3-upload-concurrency" 이 크기의 청크는 전송마다 메모리에 버퍼링됩니다.
      
      높은 속도의 링크로 대용량 파일을 전송하는 경우에 충분한 메모리가 있고 전송을 가속화하려면
      이 값을 늘리십시오.
      
      Rclone은 알려진 크기의 큰 파일을 업로드할 때 10,000개 청크 제한을
      유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      
      알려진 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본적인 청크 크기는 5 MiB이며 최대
      10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는
      48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행정보의 정확도가 낮아집니다.
      Rclone은 청크가 AWS SDK에 의해 버퍼링될 때 청크 전송이 완료된 것으로 처리하므로
      아직 업로드 중일 수 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼와
      진행률 보고를 지연시킵니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 파트 수.
      
      이 옵션은 멀티파트 업로드 시 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이것은 10,000개 청크의 AWS S3 사양을 지원하지 않는 서비스에 유용합니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드할 때 10,000개 청크 제한을
      유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환할 파일의 크기 기준.
      
      이 크기보다 큰 파일은 이 크기로 청크를 복사하여 복사됩니다.
      
      최소 크기는 0이고 최대 크기는 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬 저장하지 않기.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여
      객체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 좋지만
      큰 파일이 업로드될 때 시작하기 전에 긴 지연 시간을 초래할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를
      찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 설정되지 않은 환경 변수 "AWS_PROFILE" 또는 "default"
      사용됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 병렬 처리 수.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      고속 링크로 대용량 파일을 소량 업로드하고 이러한 업로드가
      대역폭을 완전히 활용하지 않는 경우에는 이 값을 증가시킴으로써
      전송 속도를 향상시킬 수 있습니다.

   --force-path-style
      참인 경우 패스 스타일 액세스를 사용하고 거짓인 경우 가상 호스팅 스타일을 사용합니다.
      
      참(기본값)인 경우 rclone은 패스 스타일 액세스를, 거짓인 경우 rclone은 가상 패스 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 프로바이더(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는
      이 값을 거짓으로 설정해야 합니다(알아서 설정됨).

   --v2-auth
      전송 시 v2 인증을 사용할지 여부.
      
      거짓(기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정된 경우
      rclone은 v2 인증을 사용합니다.
      
      v4 시그니처가 작동하지 않는 경우에만 사용합니다. 예를 들어 Jewel/v10 이전 CEPH의 경우입니다.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 리스트 크기).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로
      알려져 있습니다.
      대부분의 서비스는 요청된 것보다 많은 수의 응답 목록을 잘라냅니다. AWS
      S3에서는 이 값이 전역 최대값이고 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을
      증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전: 1,2 또는 auto에 대한 0.
      
      S3가 처음 출시되었을 때 bucket의 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 호출은
      성능이 훨씬 더 우수하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 provider 설정에 따라 객체의 목록을 호출할
      어떤 ListObjects 방법을 추측합니다. 만약 추측하고 그것이 잘못 될 경우 수동으로
      이곳에서 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 프로바이더에서는 파일 이름에 제어 문자를 사용할 때 이 기능이 더
      신뢰성이 있을 수 있습니다. unset(기본값)로 설정된 경우 rclone은
      프로바이더 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의
      선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷을 확인하거나 생성하지 않으려면 true로 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone의 거래 횟수를 최소화하려는 경우 이 옵션을 사용할 수 있습니다.
      
      또한 사용자가 버킷 작성 권한이 없는 경우에 필요할 수 있습니다. v1.52.0 이전에는 이것은
      버그로 인해 암묵적으로 통과되었을 것입니다.
      

   --no-head
      업로드된 개체의 무결성을 확인하기 위해 HEAD를 수행하지 않으려면 true로 설정합니다.
      
      rclone은 이 옵션을 설정하면 개체를 PUT한 후 200 OK 응답 메시지를 받게 되면
      제대로 업로드된 것으로 간주합니다.
      
      특히 다음을 가정합니다.
      
      - 업로드된 것과 동일한 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)입니다.
      - 업로드된 것과 동일한 크기입니다.
      
      단일 파트 PUT의 응답에서 다음 아이템을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목을 읽지 않습니다.
      
      길이를 알 수 없는 소스 객체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 감지 확률이 증가하므로
      일반적인 운용에는 권장되지 않습니다. 실제로 업로드 실패의 감지 확률은
      이 플래그를 사용하지 않아도 매우 낮습니다.
      

   --no-head-object
      GET을 수행하기 전에 HEAD를 수행하지 않도록 설정하려면 true로 설정합니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.
      
      부가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용 비활성화.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. 
      s3 백엔드에서 HTTP/2는 기본적으로 활성화되어 있지만 여기서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우
      더 저렴한 대출 대상으로 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      확인용으로 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값을 true, false 또는 unset으로 설정해야 합니다.

   --use-presigned-request
      단일 파트 업로드에 서명된 요청 또는 PutObject를 사용할지 여부
      
      거짓으로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여
      객체를 업로드합니다.
      
      rclone 버전 < 1.59은 서명된 요청(UnsignedPayload), 단일
      파트 객체를 업로드하려면이 플래그를 true로 설정합니다. 이것은 예외적인
      경우나 테스트 외에는 필요하지 않습니다.
      

   --versions
      디렉터리 목록에 이전 버전 포함 여부.

   --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개 변수는 날짜(yyyy-mm-dd), 날짜 및 시간(yyyy-mm-dd HH:MM:SS) 또는
      그 때의 지속 시간, 예를들면 "100d" 또는 "1h"입니다.
      
      이를 사용할 때 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

   --decompress
      gzip으로 인코딩된 객체를 해제할지 여부.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다.
      일반적으로 rclone은 이러한 파일을 압축된 개체로 다운로드합니다.
      
      이 플래그가 설정된 경우 rclone은 개체를 수신하면
      "Content-Encoding: gzip"로 압축 해제합니다. 이는 rclone이
      크기와 해시를 확인할 수 없게 되지만 파일 내용은 압축이 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우 이를 설정합니다.
      
      일반적으로 제공자는 객체를 다운로드할 때 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지
      않은 경우 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 제공자(Cloudflare 등)는 객체를 업로드할 때 "Content-Encoding: gzip"로
      업로드하지 않았더라도 개체를 gzip으로 압축할 수 있습니다.
      
      이러한 경우 다음과 같은 오류를 받게 됩니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip으로 설정되고 청크된 전송 부호화로 객체를
      다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset(기본값)으로 설정된 경우 rclone은
      프로바이더 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기 억제


옵션:
   --access-key-id value     AWS Access Key ID. [$ACCESS_KEY_ID]
   --endpoint value          Storj Gateway 엔드포인트. [$ENDPOINT]
   --env-auth                런타임에서 AWS 자격 증명 가져오기 (환경 변수나 env 변수 또는 EC2/ECS 메타데이터에서). (기본값: false) [$ENV_AUTH]
   --help, -h                도움말 표시
   --secret-access-key value AWS Secret Access Key (password). [$SECRET_ACCESS_KEY]

   고급 옵션

   --bucket-acl value          버킷 생성 시 사용되는 Canned ACL. [$BUCKET_ACL]
   --chunk-size value          업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value         멀티파트 복사로 전환할 파일의 크기 기준. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                gzip으로 인코딩된 객체를 해제할지 여부. (기본값: false) [$DECOMPRESS]
   --disable-checksum          객체 메타데이터에 MD5 체크섬 저장하지 않기. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2             S3 백엔드에서 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value        다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value            백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style          참인 경우 패스 스타일 액세스를 사용하고 거짓인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value          목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 리스트 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value     목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value        사용할 ListObjects의 버전: 1,2 또는 auto에 대한 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value    멀티파트 업로드에서 사용할 최대 파트 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value How often internal memory buffer pools will be flushed. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value          백엔드가 개체를 gzip으로 압축할 수 있는 경우 이를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket           버킷을 확인하거나 생성하지 않으려면 true로 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                   업로드된 개체의 무결성을 확인하기 위해 HEAD를 수행하지 않으려면 true로 설정합니다. (기본값: false) [$NO_HEAD]
   --no-head-object            GET을 수행하기 전에 HEAD를 수행하지 않도록 설정하려면 true로 설정합니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata        시스템 메타데이터의 설정 및 읽기 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value             공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value       AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value  멀티파트 업로드의 병렬 처리 수. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value       청크 업로드로 전환할 파일의 크기 기준. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value  확인용으로 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request     단일 파트 업로드에 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                   전송 시 v2 인증을 사용할지 여부. (기본값: false) [$V2_AUTH]
   --version-at value          지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                  디렉터리 목록에 이전 버전 포함 여부. (기본값: false) [$VERSIONS]

```