# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 lyvecloud - Seagate Lyve Cloud

사용법:
   singularity storage update s3 lyvecloud [command options] <name|id>

DESCRIPTION:
   --env-auth
      AWS 인증 정보를 런타임에서 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터).

      access_key_id 및 secret_access_key가 비어 있을 때만 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 인증 정보 입력
         | true  | 환경으로부터 AWS 인증 정보 가져오기 (env vars 또는 IAM).

   --access-key-id
      AWS Access Key ID.

      익명 액세스 또는 실행시 액세스 자격 증명의 경우 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호).

      익명 액세스 또는 실행시 액세스 자격 증명의 경우 비워 둡니다.

   --region
      연결할 지역.

      S3 클론을 사용하고 지역이없는 경우 비워 둡니다.

      예:
         | <unset>            | 불확실한 경우 사용하세요.
         |                    | v4 서명 및 빈 지역 사용
         | other-v2-signature | v4 서명이 작동하지 않을 때만 사용.
         |                    | 예 : 이전 CEPH의 Jewel/v10.

   --endpoint
      S3 API의 엔드포인트.

      S3 클론을 사용하는 경우 필수 입력 사항입니다.

      예:
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US East 1 (Virginia)
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud US West 1 (California)
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP Southeast 1 (Singapore)

   --location-constraint
      위치 제약 조건 - 지역과 일치해야 함.

      확실하지 않으면 비워 둡니다. 버킷 생성시에만 사용됩니다.

   --acl
      버킷을 생성하거나 객체를 저장하거나 복사 할 때 사용되는 canned ACL.

      이 ACL은 객체 생성에 사용되며 버킷_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.

      서버 측에서 복사 중에 이 ACL이 적용됩니다. S3는 소스의 ACL을 복사하지 않고 새로 작성합니다.

      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.


   --bucket-acl
      버킷을 생성할 때 사용되는 canned ACL.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하세요.

      이 ACL은 버킷을 생성 할 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.

      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.


      예:
         | private            | 소유자는 FULL_CONTROL이됩니다.
         |                    | 다른 사람들은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자는 FULL_CONTROL이됩니다.
         |                    | AllUsers 그룹에는 읽기 액세스가 있습니다.
         | public-read-write  | 소유자는 FULL_CONTROL이됩니다.
         |                    | AllUsers 그룹에는 읽기 및 쓰기 액세스가 있습니다.
         |                    | 일반적으로 버킷에 대해 이렇게 권한을 부여하는 것은 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL이됩니다.
         |                    | 인증 된 사용자 그룹에는 읽기 액세스가 있습니다.

   --upload-cutoff
      청크 업로드로 전환하는 데 사용되는 임계 값.

      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.

      upload_cutoff보다 크거나 알려지지 않은 크기의 파일 (예 : "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 도큐먼트에서 전송 한 파일)은이 청크 크기를 사용하여 멀티 파트 업로드로 업로드됩니다.

      유의하세요. "--s3-upload-concurrency"이 크기의 청크가 메모리당 전송으로
      버퍼링됩니다.

      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있는 경우이 값을 증가시키면 전송 속도가 향상됩니다.

      rclone은 알려진 크기의 대용량 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 10,000 청크 제한을 유지합니다.

      파일 크기를 모르는 경우 설정된
      chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이고 최대 10,000 청크가 있을 수 있으므로,
      기본적으로 스트리밍 업로드 할 수있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트리밍 업로드하려면
      청크 크기를 증가시켜야합니다.

      청크 크기를 증가시키면 프로그레스의 정확도가 감소합니다.
      "-P" 플래그와 함께 표시되는 진행 상황 통계의 ACCURACY. rclone은 청크가 S3에
      버퍼링 될 때 완료된 것으로 처리하며, 실제로 업로드 중인 경우.
      청크 크기가 클수록 AWS SDK 버퍼와 진행률
      보고는 진실에서 벗어났다.

   --max-upload-parts
      멀티 파트 업로드의 최대 부분 수.

      이 옵션은 멀티 파트 업로드시 사용할 최대 멀티파트 청크 수를 정의합니다.

      서비스가 AWS S3 사양의 10,000 청크를 지원하지 않는 경우 유용 할 수 있습니다.

      rclone은 알려진 크기의 대용량 파일을 업로드 할 때 청크 크기를 자동으로 증가시켜 이러한 청크 수 제한 아래로 유지합니다.


   --copy-cutoff
      멀티파트 복사로 전환하는 임계 값.

      서버 측에서 복사해야하는 이보다 큰 파일은
      이 크기의 청크로 복사됩니다.

      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체 메타 데이터에 추가합니다. 이는
      데이터 무결성 확인에 좋습니다만 큰 파일에 대해 많은 지연이 발생할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. env 값이 비어 있으면 기본값은 현재 사용자의 홈 디렉토리입니다.

          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용 할 프로필.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용되는 프로필을 제어합니다.

      비어 있으면 환경 변수 "AWS_PROFILE" 또는
      "default"라는 환경 변수가 설정되지 않은 경우 기본값으로 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.

      파일의 동일한 청크 수가 동시에 업로드됩니다.

      큰 파일을 고속 링크를 통해 작은 수의 큰 파일을 업로드하고 이 업로드가 대역폭을 완벽하게 활용하지 못하는 경우 업로드 속도를 높이기 위해 이 값을 증가 할 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 경로 스타일을 사용합니다.

      true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고
      false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3
      docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)을 참조하세요.

      제공자(예: AWS, Aliyun OSS, Netease COS, 또는 Tencent COS) 중 일부에서는이 값을
      false로 설정해야 합니다. rclone은 제공자에 따라 자동으로처리합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.

      false(기본값)인 경우 rclone은 v4 인증을 사용합니다.
      설정된 경우 rclone은 v2 인증을 사용합니다.

      v4 서명이 작동하지 않을 때만 사용하십시오, 예 : Jewel/v10 이전의 CEPH.

   --list-chunk
      목록 청크의 크기.

      이 옵션은 또한 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로
      알려진 것입니다. 대부분의 서비스는 1000 개 이상 요청하더라도 응답 리스트를 1000 개로 자른다.
      AWS S3에서는이 전역 최대값을 변경할 수 없으며 [AWS S3] 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전 : 1, 2 또는 자동으로 0.

      S3가 처음 시작되었을 때 bucket 내의 객체를 나열하기 위해 ListObjects 호출 만 제공되었습니다.

      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 높은 성능을 제공하며 가능한 경우 사용해야합니다.

      기본값인 0으로 설정되면 rclone은 공급자에 따라 추측하여 어떤 목록 객체 방법을 호출할지 추측합니다. 추측이 잘못될 경우 수동으로 여기에서 설정할 수 있습니다.
      

   --list-url-encode
      목록을 url 인코딩 여부 : true/false/unset
      
      일부 제공자는 URL 인코딩 목록을 지원하며 사용 가능한 경우 파일에서
      제어 문자를 사용할 때이 방법이 더 안정적입니다. unset으로 설정되어 있으면
      rclone은 공급자 설정에 따라 적용 할 것을 선택하나 rclone의 선택을 여기에서 무시 할 수 있습니다.
      

   --no-check-bucket
      설정된 경우, 버킷이 존재하는지 확인하거나 생성하지 않습니다.

      버킷이 이미 존재하는 경우 rclone의 수행 트랜잭션 수를 최소화하려는 경우 유용 할 수 있습니다.

      사용자가 버킷 생성 권한이 없는 경우 필요할 수도 있습니다. v1.52.0 이전에는 버그로 인해이가 정상적으로 전달되었습니다. 

   --no-head
      설정된 경우 요청의 무결성을 확인하기 위해 업로드 된 객체의 HEAD를 실행하지 않습니다.

      rclone은 요청자가 PUT 후에 200 OK 메시지를 수신하면 제대로 업로드 된 것으로 간주합니다.

      특히 다음과 같습니다:
      
      - 업로드 시 메타데이터, 수정 시간, 저장 클래스 및 콘텐츠 유형이 업로드 된 내용과 동일합니다.
      - 크기가 업로드 된 것과 동일합니다.
      
      다음 내용을 읽습니다.

      - MD5SUM
      - 업로드 날짜
      
      멀티 파트 업로드의 경우 이러한 항목을 읽을 수 없습니다.
      
      길이를 알 수없는 소스 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 가능성이 증가합니다.
      특히 잘못된 크기로 인한 것이므로 일반 작업에는 권장하지 않습니다. 
      실제로 업로드 실패가 발생할 가능성은 매우 적습니다.

   --no-head-object
      GET을 하기 전에 HEAD를 실행하지 않습니다.

   --encoding
      백엔드의 인코딩.

      자세한 내용은 [도시 개요](/overview/#encoding)의 인코딩 섹션을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시 될 간격.

      추가 버퍼를 필요로하는 업로드 (예 : 멀티 파트)는 할당을 위해 메모리 풀을 사용합니다. 이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 속도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부

   --disable-http2
      S3 백엔드의 http2 사용 중지

      현재 s3 (특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. S3 기본적으로 s3 백엔드에 대해 HTTP/2가 활성화되어 있지만 여기에서 비활성화 할 수 있습니다. 문제가 해결되면이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드 할 때 더 저렴하게 데이터를 내보냅니다.

   --use-multipart-etag
      확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      true, false 또는 기본값을 알고자하는 경우 설정하지 않은 상태여야합니다.

   --use-presigned-request
      단일 파트 업로드를위한 사전 서명 된 요청 또는 PutObject를 사용할지 여부

      이 플래그가 거짓이면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone < 1.59 버전에서는 단일 부분 객체를 업로드하기 위해 사전 서명 된 요청을 사용하고 이 플래그를 true로 설정하면 다시 활성화됩니다. 이 기능은 특정한 상황이나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉토리 목록에 이전 버전 포함.

   --version-at
      지정된 시간의 파일 버전으로 표시합니다.
      
      매개 변수는 날짜, "2006-01-02", 날짜시각 "2006-01-02
      15:04:05" 또는 그렇게 오래된 기간, 예 : "100d" 또는 "1h"가 될 수 있습니다.
      
      이를 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      올바른 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip로 인코딩 된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 보통 rclone은 이러한 파일을 압축 된 객체로 다운로드합니다.
      
      이 플래그가 설정되어 있으면 rclone은 "Content-Encoding: gzip"로 전송된 이러한 파일을 수신하는 동안 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인 할 수 없다는 것을 의미하지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 gzip을 압축 할 수 있다면 지정하십시오.
      
      보통 공급자는 객체를 다운로드 할 때 객체를 수정하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 경우 다운로드시 설정되지 않습니다.

      그러나 일부 공급 업체는 "Content-Encoding: gzip"로 업로드되지 않은 경우에도 gzip으로 객체를 압축 할 수 있습니다 (예 : Cloudflare).
      
      이러한 문제의 증상은 다음과 같습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM

      이 플래그를 설정하고 rclone이 "Content-Encoding: gzip"을 설정하고 청크 전송 인코딩으로 객체를 다운로드하는 경우 rclone은 객체를 실시간으로 압축 해제합니다.

      이 값을 unset로 설정(default)하면 rclone은 공급 업체 설정에 따라 적용할 대상을 선택하지만 여기에서 rclone의 선택을 무시 할 수 있습니다.

   --no-system-metadata
      시스템 메타 데이터의 설정 및 읽기를 억제합니다.

옵션:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하거나 객체를 저장하거나 복사 할 때 사용되는 canned ACL입니다. [$ACL]
   --endpoint value             S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   AWS 인증 정보를 런타임에서 가져옵니다(환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터). (default: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  위치 제약 조건 - 지역과 일치해야 함. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역. [$REGION]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (default: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티 파트 복사로 전환하는 임계 값입니다. (default: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip로 인코딩 된 객체를 압축 해제합니다. (default: false) [$DECOMPRESS]
   --disable-checksum               객체 메타 데이터에 MD5 체크섬을 저장하지 않습니다. (default: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드의 http2 사용 중지 (default: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (default: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 경로 스타일을 사용합니다. (default: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다. (default: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 url 인코딩 할지 여부 : true/false/unset (default: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전 : 1,2 또는 자동으로 0 (default: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 부분 수입니다. (default: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시 될 간격입니다. (default: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에 mmap 버퍼를 사용할지 여부입니다. (default: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 gzip을 압축 할 수 있다면 지정하십시오. (default: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정된 경우, 버킷이 존재하는지 확인하거나 생성하지 않습니다. (default: false) [$NO_CHECK_BUCKET]
   --no-head                        설정된 경우 업로드 된 객체의 HEAD를 실행하지 않습니다. (default: false) [$NO_HEAD]
   --no-head-object                 GET을 하기 전에 HEAD를 실행하지 않습니다. (default: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타 데이터의 설정 및 읽기를 억제합니다 (default: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용 할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (default: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계 값입니다. (default: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (default: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드를위한 사전 서명 된 요청 또는 PutObject를 사용할지 여부 (default: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (default: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전으로 표시됩니다. (default: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전 포함. (default: false) [$VERSIONS]

```
{% endcode %}