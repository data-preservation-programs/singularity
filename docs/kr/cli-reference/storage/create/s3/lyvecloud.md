# Seagate Lyve Cloud

{% code fullWidth="true" %}
```
명령어:
   singularity storage create s3 lyvecloud - Seagate Lyve Cloud

사용법:
   singularity storage create s3 lyvecloud [command options] [arguments...]

설명:
   --env-auth
      런타임으로부터 AWS 자격증명을 가져옵니다 (환경 변수 또는 env 변수가 없는 경우 EC2/ECS 메타 데이터에서 가져옵니다).
      
      access_key_id 및 secret_access_key가 비어 있을 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 AWS 자격증명을 입력하세요.
         | true  | 런타임에서 AWS 자격증명을 가져옵니다 (환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID입니다.
      
      익명 액세스 또는 런타임 자격증명을 사용할 경우 비워둡니다.

   --secret-access-key
      AWS Secret 액세스 키(비밀번호)입니다.
      
      익명 액세스 또는 런타임 자격증명을 사용할 경우 비워둡니다.

   --region
      연결할 리전입니다.
      
      S3 클론을 사용하고 리전이 없는 경우 비워둡니다.

      예제:
         | <unset>            | 확실하지 않은 경우 사용하세요.
         |                    | v4 서명 및 빈 리전을 사용합니다.
         | other-v2-signature | v4 서명이 작동하지 않는 경우에만 사용하세요.
         |                    | 예: pre Jewel/v10 CEPH.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필수입니다.

      예제:
         | s3.us-east-1.lyvecloud.seagate.com      | Seagate Lyve Cloud 미국 동부 1 (버지니아)
         | s3.us-west-1.lyvecloud.seagate.com      | Seagate Lyve Cloud 미국 서부 1 (캘리포니아)
         | s3.ap-southeast-1.lyvecloud.seagate.com | Seagate Lyve Cloud AP 동남아시아 1 (싱가포르)

   --location-constraint
      리전과 일치하는 위치 제한입니다.
      
      확실하지 않을 경우 비워둡니다. 버킷을 생성할 때 사용됩니다.

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 선언적 ACL입니다.
      
      이 ACL은 객체를 생성할 때 사용되며, bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 [Amazon S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      S3는 서버 측 복사 시 소스의 ACL을 복사하지 않고 새로운 ACL을 작성합니다.
      
      acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

   --bucket-acl
      버킷을 만들 때 사용되는 선언적 ACL입니다.
      
      자세한 내용은 [Amazon S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참조하세요.
      
      acl을 설정하지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값 (private)이 사용됩니다.
      

      예제:
         | private            | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | 다른 사람에게 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | AllUsers 그룹에게 READ 권한을 부여합니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 권한을 부여합니다.
         |                    | 버킷에 이 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한을 부여합니다.
         |                    | 인증된 사용자 그룹에게 READ 권한을 부여합니다.

   --upload-cutoff
      청크 업로드로 전환하는 데 사용되는 분할을 위한 임계점입니다.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 알려지지 않은 크기의 파일("rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일)을 업로드할 경우, 이 청크 크기를 사용하여
      다중 부분 업로드가 진행됩니다.
      
      참고로, "--s3-upload-concurrency" 크기의 청크가 전송마다 메모리에 버퍼링됩니다.
      
      고속 링크로 대용량 파일을 전송하고 메모리 여유가 있는 경우, 이를 증가시켜 전송 속도를 높일 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 10,000 청크 제한을 지키기 위해
      자동으로 청크 크기를 증가시킵니다.
      
      알려지지 않은 크기의 파일은 구성된
      청크 크기로 업로드됩니다. 기본적으로 청크 크기는 5 MiB이며 최대 10,000 청크까지 존재할 수 있으므로,
      기본적으로 스트림으로 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림으로 업로드하려면
      청크 크기를 증가시켜야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태의 정확성이 감소합니다. Rclone은 청크를 AWS SDK에 버퍼로 전송한 경우 청크 전송으로 처리하지만,
      실제로는 여전히 업로드 중일 수 있습니다.
      청크 크기가 클수록 AWS SDK 버퍼가 커지므로 진행 상태 보고가 진실성에서 더욱 벗어나게 됩니다.
      

   --max-upload-parts
      다중 부분 업로드 중 사용할 최대 부분 수입니다.
      
      이 옵션은 다중 부분 업로드 시 사용할 최대 부분 수를 정의합니다.
      
      서비스가 10,000 개의 청크에 대한 AWS S3 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      Rclone은 알려진 크기의 큰 파일을 업로드 할 때 10,000 청크 제한을 지키기 위해
      자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      분할 복사로 전환하는 데 사용되는 임계치입니다.
      
      서버 측 복사가 필요한 이 임계치보다 큰 파일은
      이 사이즈의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이렇게 함으로써 데이터 무결성 확인에 유리하지만, 대용량 파일의 업로드 시작을 위해 오랜 지연을 발생시킬 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. env 값이 비어 있으면 사용자의 현재 홈 디렉토리를 기본값으로 사용합니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
       

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수를 사용하여 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워두면 "AWS_PROFILE" 또는
      그 환경 변수도 설정되어 있지 않은 경우에는 "default"로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      다중 부분 업로드에 대한 동시성입니다.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      고속 링크로 대용량 파일을 하이스피드 링크로 업로드하고 업로드가 제대로 이루어 지지 않을 경우, 이 값을 증가시켜 전송 속도를 높일 수 있습니다.

   --force-path-style
      true인 경우 경로 스타일 접근을 사용하고 false인 경우 가상 호스팅 스타일 접근을 사용합니다.
      
      기본값이 true이므로 rclone은 경로 스타일 접근을 사용합니다.
      false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 제공업체 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 false로 설정해야합니다.
      rclone은 제공자 설정에 따라 자동으로 이를 수행합니다.

   --v2-auth
      true인 경우 v2 인증을 사용합니다.
      
      기본값이 false이므로 rclone은 v4 인증을 사용합니다.
      이 값을 설정하면 rclone은 v2 인증을 사용합니다.
     
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예: pre Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록).
      
      AWS S3 사양에서는 이 옵션은 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 요청된 객체의 응답 목록을 1000 개로 잘라냅니다.
      AWS S3에서는 이것이 전역적인 최대값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1,2 또는 0은 자동으로 설정합니다.
      
      S3가 처음에 출시될 때 버킷의 개체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 더 높은 성능을 제공하고 가능하면 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 목록 개체 방법을 추측합니다. 잘못 추측하면 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL로 인코딩할지 여부: true/false/unset
      
      일부 제공자는 목록을 URL로 인코딩하는 것을 지원하며, 파일 이름에 제어 문자를 사용할 때 이것이 더 신뢰할 수 있습니다. 이 값이 unset으로 설정된 경우(rclone의 기본값) rclone은 제공자 설정에 따라 적용할 내용을 선택하지만, 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      버킷 생성 권한이 없는 사용자의 경우 필요할 수도 있습니다. v1.52.0 이전에는 이전에는 이것은 버그로 인해 조용히 통과되었습니다.
      

   --no-head
      청크 업로드의 무결성을 확인하기 위해 업로드된 개체를 HEAD하지 않습니다.
      
      rclone은 GET 이전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      더 자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 간격입니다.
      
      추가 버퍼(예: 멀티파트로 업로드되는 파일)가 필요한 업로드에는 메모리 풀이 사용됩니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용해야하는지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2 관련된 해결되지 않은 문제가 있습니다. S3 백엔드에서는 기본적으로 HTTP/2가 사용됩니다.
      그러나 여기에서 사용을 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      

   --download-url
      다운로드용 사용자 정의 엔드포인트입니다.
      일반적으로 AWS S3는 CloudFront 네트워크를 통해 데이터 다운로드에 대한 더 저렴한 출구를 제공하므로 이를 CloudFront CDN URL로 설정합니다.

   --use-multipart-etag
      확인을 위해 multipart 업로드에서 ETag를 사용하는지 여부
      
      true, false 또는 기본값 (공급자에 따름)으로 설정해야 합니다.
      

   --use-presigned-request
      단일 부분 업로드에 대해 사전 서명된 요청을 사용할지 여부
      
      이 값이 false이면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone의 버전 < 1.59에서는 단일 부분 객체를 업로드하기 위해 사전 서명된 요청을 사용하고, 이 플래그를 true로 설정하면 이 기능을 다시 활성화합니다. 이는 예외적인 경우 또는 테스트 목적 이외에 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간의 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 시간("2006-01-02 15:04:05") 또는 그 이전 시간 동안의 기간("100d" 또는 "1h")이어야합니다.
      
      이를 사용하면 파일을 업로드하거나 삭제할 수 없으므로,
      파일 쓰기 작업을 수행할 수 없습니다.
      
      유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참조하세요.
      

   --decompress
      gzip로 압축된 개체를 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 "Content-Encoding: gzip"로 수신받은 파일을 복원합니다. 이로 인해 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우이를 설정하세요.
      
      일반적으로 제공자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 제공자는 객체를 gzip으로 압축할 수 있습니다. (예: Cloudflare).
      
      이렇게 될 경우 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip이 설정된 chunked 전송 인코딩과 함께 개체를 다운로드하면 rclone은 개체를 실시간으로 해제합니다.
      
      unset으로 설정하면 (기본값인) rclone은 제공자 설정에 따라 적용할 내용을 선택하지만, 이 값을 재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value        AWS 액세스 키 ID입니다. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 선언적 ACL입니다. [$ACL]
   --endpoint value             S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   런타임으로부터 AWS 자격증명을 가져옵니다 (환경 변수 또는 env 변수가 없는 경우 EC2/ECS 메타 데이터에서 가져옵니다). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말을 표시합니다
   --location-constraint value  리전과 일치하는 위치 제한입니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 리전입니다. [$REGION]
   --secret-access-key value    AWS Secret 액세스 키(비밀번호)입니다. [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 만들 때 사용되는 선언적 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              분할 복사로 전환하는 데 사용되는 임계치입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 압축된 개체를 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드용 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 경로 스타일 접근을 사용하고 false인 경우 가상 호스팅 스타일 접근을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL로 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0은 자동으로 설정합니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         다중 부분 업로드 중 사용할 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 간격입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용해야하는지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는 경우이를 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        청크 업로드의 무결성을 확인하기 위해 업로드된 개체를 HEAD하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 수행하기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       다중 부분 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 데 사용되는 임계점입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       확인을 위해 multipart 업로드에서 ETag를 사용하는지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 사전 서명된 요청을 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  저장소 이름 (기본값: 자동 생성)
   --path value  저장소 경로

```
{% endcode %}