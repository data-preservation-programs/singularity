# Wasabi 객체 저장소

{% code fullWidth="true" %}
```
명령:
   singularity storage update s3 wasabi - Wasabi 객체 저장소

사용법:
   singularity storage update s3 wasabi [command options] <name|id>

설명:
   --env-auth
      런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경변수가 없는 경우 EC2/ECS 메타데이터).
      
      access_key_id 및 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경에서 AWS 자격 증명 가져오기 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키 (비밀번호).
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --region
      연결할 지역.
      
      S3 클론을 사용하고 지역이 없는 경우 비워 둡니다.

      예시:
         | <unset>            | 확실하지 않은 경우 이 값을 사용하세요.
         |                    | v4 서명 및 빈 지역을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용하세요.
         |                    | 예: 이전 Jewel/v10 CEPH.

   --endpoint
      S3 API 엔드포인트.
      
      S3 클론을 사용할 때 필요합니다.

      예시:
         | s3.wasabisys.com                | Wasabi 미국 동부 1 (N. 버지니아)
         | s3.us-east-2.wasabisys.com      | Wasabi 미국 동부 2 (N. 버지니아)
         | s3.us-central-1.wasabisys.com   | Wasabi 미국 중부 1 (텍사스)
         | s3.us-west-1.wasabisys.com      | Wasabi 미국 서부 1 (오레곤)
         | s3.ca-central-1.wasabisys.com   | Wasabi CA 중앙 1 (토론토)
         | s3.eu-central-1.wasabisys.com   | Wasabi EU 중앙 1 (암스테르담)
         | s3.eu-central-2.wasabisys.com   | Wasabi EU 중앙 2 (프랑크푸르트)
         | s3.eu-west-1.wasabisys.com      | Wasabi EU 서부 1 (런던)
         | s3.eu-west-2.wasabisys.com      | Wasabi EU 서부 2 (파리)
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP 북동 1 (도쿄) 엔드포인트
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP 북동 2 (오사카) 엔드포인트
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP 남동 1 (싱가포르)
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP 남동 2 (시드니)

   --location-constraint
      위치 제약 - 지역과 일치해야 합니다.
      
      확실하지 않은 경우 비워 둡니다. 버킷 생성 시에만 사용됩니다.

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 기본 ACL.
      
      이 ACL은 객체 생성 시 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      S3는 서버 측에서 객체를 복사할 때 소스의 ACL을 복사하지 않고 새로 작성하기 때문에
      이 ACL은 서버 측 객체 복사에 적용됩니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl 헤더가 추가되지 않고
      기본값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 기본 ACL.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.
      
      이 ACL은 버킷 생성 시에만 사용됩니다. 설정되어 있지 않으면 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl:
      헤더가 추가되지 않고 기본값(개인)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | 다른 사람에게 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 읽기 액세스 권한 부여.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AllUsers 그룹에게 읽기 및 쓰기 액세스 권한 부여.
         |                    | 버킷에 대해서는 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한 부여.
         |                    | AuthenticatedUsers 그룹에게 읽기 액세스 권한 부여.

   --upload-cutoff
      청크 업로드로 전환하는 크기 임계값.
      
      이 값을 초과하는 파일은 chunk_size 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      업로드_cutoff을 초과하는 파일이나 크기가 알려지지 않은 파일(예: "rclone rcat"에서 가져온 파일이나 "rclone mount" 또는 Google 사진 또는 Google 문서에 업로드한 파일)은 이 청크 크기를 사용하여 다량 업로드 형식으로 업로드됩니다.
      
      "--s3-upload-concurrency" 크기 단위로 이 크기의 청크가 전송당 메모리에 버퍼링됩니다.
      
      초고속 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 늘릴수록 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대용량 파일을 전송할 때 10,000개의 청크 제한을 초과하지 않도록 자동으로 청크 크기를 늘립니다.
      
      알려진 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size가 5 MiB이고 최대 10,000개의 청크로 제한되기 때문에 기본값으로 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
      
      청크 크기를 늘리면 진행 상태 통계의 정확성이 감소합니다. Rclone은 청크가 AWS SDK에서 버퍼링될 때 청크를 보낸 것으로 처리하지만, 실제로는 아직 업로드되고 있을 수 있습니다.
      

   --max-upload-parts
      다중 파트 업로드에 사용할 최대 파트 수.
      
      이 옵션은 다중 파트 업로드를 수행할 때 사용할 다중 파트 청크의 최대 수를 정의합니다.
      
      AWS S3 사양의 10,000개 청크를 지원하지 않는 서비스의 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 대용량 파일을 전송할 때 이 청크 크기를 초과하지 않도록 자동으로 청크 크기를 늘립니다.
      

   --copy-cutoff
      다중 부분 복사로 전환하는 크기 임계값.
      
      서버 측에서 복사해야 하는 이 크기를 초과하는 파일은 이 크기로 청크별로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      보통 rclone은 업로드 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가하여 데이터 무결성 확인에 사용합니다. 이는 큰 파일의 업로드 시작에 오랜 지연을 초래할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일 경로.
      
      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수가 비어 있다면 현재 사용자의 홈 디렉터리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"가 설정되지 않은 경우 기본값을 사용합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      다중 파트 업로드에 대한 동시성.
      
      동일한 파일의 여러 청크를 동시에 업로드하는 것입니다.
      
      고속 링크를 통해 대량의 대용량 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 늘릴수록 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.
      
      기본값인 true인 경우 rclone은 경로 스타일 액세스를 사용하고, false인 경우 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 설명서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS)는 이 값을 false로 설정해야 합니다. rclone은 이 값이 공급자 설정에 따라 자동으로 처리합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      false로 설정된 경우(기본값) rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: 이전 Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려진 목록 청크의 크기입니다. 대부분의 서비스는 응답 목록을 1000개의 객체로 잘라냅니다. AWS S3에서는 이것이 전역 최대값이며 수정할 수 없습니다. 자세한 내용은 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요. Ceph에서는 "rgw list buckets 최대 청크" 옵션으로 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 자동(0).
      
      S3가 처음 출시될 때 버킷의 객체를 열거하는 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 호출은 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야 합니다.
      
      기본값인 0으로 설정되면 rclone은 공급자 설정에 따라 어떤 목록 객체 방법을 호출할지 추측합니다. 추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 파일 이름에 제어 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 사용 가능한 경우이 방법을 사용하면 파일에 액세스하는 데 더 신뢰할 수 있습니다. unset으로 설정된 경우(기본값) rclone은 제공자 설정에 따라 적용할 항목을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없을 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 이를 오류 없이 전달했을 것입니다.
      

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다.
      
      rclone은 일반적으로 작은 수의 큰 파일을 전송하고 이러한 전송이 전체 대역폭을 활용하지 않는 경우 HEAD 요청 수를 최소화하려는 경우에 유용할 수 있습니다.
      
      이 플래그가 설정되면 rclone은 PUT로 객체를 전송한 후 200 OK 메시지를 받으면 올바르게 전송되었다고 가정합니다.
      
      특히, rclone은 다음 항목을 단일 파트 PUT의 응답에서 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      대형 파일에 대해서는 여러 파트 업로드에서 이러한 항목을 읽지 않습니다.
      
      크기가 알려지지 않은 원본 객체를 업로드하면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 감지 확률이 높아지므로 정상 작업에는 권장되지 않습니다. 실제로 업로드 실패의 가능성은 매우 적습니다.
      

   --no-head-object
      GET 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 시간 간격.
      
      추가 버퍼(예: 다중 파트)를 필요로 하는 업로드는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

   --disable-http2
      S3 백엔드에 대한 http2 사용 비활성화.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관련된 문제가 있습니다. s3 백엔드의 기본값으로 HTTP/2가 활성화되지만 이곳에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 탈출 트래픽을 제공합니다.

   --use-multipart-etag
      다중 파트 업로드에 ETag를 사용하여 검증할지 여부
      
      true, false 또는 기본값(설정되지 않은 경우)을 사용합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부
      
      이 값이 false이면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone의 버전이 1.59 미만인 경우 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하고이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이것은 예외적인 상황이나 테스트를 위해서만 필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전 포함.

   --version-at
      지정된 시간에 파일 버전을 그대로 표시합니다.
      
      매개변수는 날짜, "2006-01-02", 날짜시간 "2006-01-02 15:04:05" 또는 그보다 오래된 기간인 "100d" 또는 "1h"일 수 있습니다.
      
      이를 사용할 때 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 설명서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip으로 인코딩된 개체를 압축 해제합니다.
      
      "Content-Encoding: gzip"으로 S3에 파일을 업로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되어 있다면 rclone은 수신된 "Content-Encoding: gzip"로 이러한 파일을 압축 해제합니다. 즉, rclone은 파일의 크기 및 해시를 확인할 수는 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우 이를 설정하십시오.
      
      일반적으로 제공자는 객체를 다운로드할 때 객체를 수정하지 않습니다. "Content-Encoding: gzip"으로 업로드되지 않은 객체에는 이 값이 설정되지 않습니다.
      
      그러나 일부 공급자(예: Cloudflare)는 객체를 "Content-Encoding: gzip"으로 업로드되지 않았음에도 gzip으로 압축할 수 있습니다.
      
      이 값을 설정하고 rclone이 "Content-Encoding: gzip"이 설정된 청크 전송 인코딩으로 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      unset으로 설정되면 rclone은 제공자 설정에 따라 적용할 항목을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정과 읽기를 억제합니다.


옵션:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  버킷 생성 및 객체 저장 또는 복사 시 사용되는 기본 ACL. [$ACL]
   --endpoint value             S3 API 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명 가져오기 (환경 변수 또는 환경변수가 없는 경우 EC2/ECS 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 - 지역과 일치해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 기본 ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              다중 복사로 전환하는 크기 임계값. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용 비활성화. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 자동(0). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         다중 파트 업로드에 사용할 최대 파트 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간 간격. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는 경우 이를 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재 여부를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정과 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       다중 파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 크기 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       다중 파트 업로드에 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 그대로 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전 포함. (기본값: false) [$VERSIONS]

```
{% endcode %}