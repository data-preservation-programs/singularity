# Wasabi Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 wasabi - Wasabi 객체 스토리지

사용법:
   singularity storage create s3 wasabi [명령 옵션] [인수...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 받아옵니다 (환경 변수 또는 env 변수에 저장된 EC2/ECS 메타 데이터).
     
      만약 access_key_id와 secret_access_key가 비어 있다면 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력합니다.
         | true  | 환경 (환경 변수 또는 IAM) 에서 AWS 자격 증명을 받아옵니다.

   --access-key-id
      AWS Access Key ID입니다.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호) 입니다.
      
      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워 둡니다.

   --region
      연결할 지역입니다.
      
      S3 클론을 사용하고 지역이 필요하지 않은 경우 비워 둡니다.

      예시:
         | <unset>            | 확실하지 않을 경우 이 옵션을 사용하세요.
         |                    | v4 서명 및 빈 지역이 사용됩니다.
         | other-v2-signature | v4 서명이 작동하지 않는 경우에만 사용하세요.
         |                    | 예 : 이전 Jewel/v10 CEPH 입니다.

   --endpoint
      S3 API의 엔드포인트입니다.
      
      S3 클론을 사용하는 경우 필수입니다.

      예시:
         | s3.wasabisys.com                | Wasabi US East 1 (N. Virginia)
         | s3.us-east-2.wasabisys.com      | Wasabi US East 2 (N. Virginia)
         | s3.us-central-1.wasabisys.com   | Wasabi US Central 1 (Texas)
         | s3.us-west-1.wasabisys.com      | Wasabi US West 1 (Oregon)
         | s3.ca-central-1.wasabisys.com   | Wasabi CA Central 1 (Toronto)
         | s3.eu-central-1.wasabisys.com   | Wasabi EU Central 1 (Amsterdam)
         | s3.eu-central-2.wasabisys.com   | Wasabi EU Central 2 (Frankfurt)
         | s3.eu-west-1.wasabisys.com      | Wasabi EU West 1 (London)
         | s3.eu-west-2.wasabisys.com      | Wasabi EU West 2 (Paris)
         | s3.ap-northeast-1.wasabisys.com | Wasabi AP Northeast 1 (Tokyo) 엔드포인트
         | s3.ap-northeast-2.wasabisys.com | Wasabi AP Northeast 2 (Osaka) 엔드포인트
         | s3.ap-southeast-1.wasabisys.com | Wasabi AP Southeast 1 (Singapore)
         | s3.ap-southeast-2.wasabisys.com | Wasabi AP Southeast 2 (Sydney)

   --location-constraint
      지역 제한 - 지역과 일치하도록 설정해야 합니다.
      
      확실하지 않으면 비워 둡니다. 버킷을 생성할 때만 사용됩니다.

   --acl
      버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 사전 정의 ACL입니다.
      
      이 ACL은 객체를 만들 때 사용되며, bucket_acl이 설정되지 않은 경우 bucket을 만들 때도 사용됩니다.
      
      자세한 내용은 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) 를 참조하세요.
      
      S3는 해당 ACL을 소스에서 복사하지 않고 새로 생성합니다.
      
      acl으로 빈 문자열이 전달되면 X-Amz-Acl: 헤더를 추가하지 않으며 기본 값(개인)이 사용됩니다.
      

   --bucket-acl
      버킷을 생성할 때 사용되는 사전 정의 ACL입니다.
      
      자세한 내용은 [Amazon S3 ACL](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl) 를 참조하세요.
      
      이 ACL은 오직 버킷을 생성할 때만 사용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl" 및 "bucket_acl"이 빈 문자열인 경우에도 X-Amz-Acl: 헤더가 추가되지 않고 기본 값(개인)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | 다른 사용자에게 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹에게 READ 액세스 권한을 부여합니다.
         | public-read-write  | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AllUsers 그룹이 READ 및 WRITE 액세스 권한을 부여합니다.
         |                    | 이러한 작업은 일반적으로 버킷에 대해 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL을 부여합니다.
         |                    | AuthenticatedUsers 그룹이 READ 액세스 권한을 부여합니다.

   --upload-cutoff
      청크가 업로드로 전환되는 용량 임계치입니다.
      
      이 보다 큰 파일은 chunk_size의 크기로 청크 형식으로 업로드됩니다.
      최소값은 0이고 최대값은 5GB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 (예: "rclone rcat"에서 가져온 파일이거나 "rclone mount" 또는 google 
      photos 또는 google docs로 업로드 된 파일)은 이 청크 크기를 사용하여 multipart 업로드로 업로드됩니다.
      
      주의: "--s3-upload-concurrency"의 지정된 크기의 청크가 전송마다 메모리에 버퍼되어 있습니다.
      
      고속 링크를 통해 대용량 파일을 전송하고 충분한 메모리가 있으면 이 값을 키우면 전송 속도가 향상됩니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드할 때 자동으로 청크 크기를 증가시켜 10,000 개의 청크 제한을 
      유지합니다.
      
      알려진 크기가없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size가 5MiB이고 최대 10,000 
      청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수있는 파일 크기의 최대 크기는 48 GiB입니다. 더 큰 
      파일을 스트림 업로드하려면 chunk_size를 크게 설정해야합니다.
      
      chunk 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 통계의 정확도가 감소합니다. Rclone은 
      AWS SDK가 버퍼로 사용할때 청크 전송이 완료되었다고 판단하며 업로드 중일 수 있으므로 실제로 전송이 
      완료될 때보다 더 많이 변형될 것입니다.

   --max-upload-parts
      Multipart 업로드에서 최대 파트 수입니다.
      
      이 옵션은 multipart 업로드를 할 때 사용할 최대 multipart 청크 수를 정의합니다.
      
      이것은 10,000 청크의 AWS S3 사양을 지원하지 않는 서비스에 유용 할 수 있습니다.
      
      Rclone은 알려진 크기의 대용량 파일을 업로드 할 때 자동으로 청크 크기를 증가시켜 이 청크 수 제한 내에 
      유지합니다.
      

   --copy-cutoff
      서버 측에서 복사해야하는 이 임계치보다 큰 파일은 이 크기로 청크 복사를 합니다.
      
      최소값은 0이고 최대값은 5GB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 
      이는 데이터 무결성 확인에 좋지만 대용량 파일을 업로드하기 시작하는 데 많은 시간이 소요될 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있다면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 
      Env 값이 비어 있으면 기본적으로 현재 사용자의 홈 디렉토리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 
      이 변수는 그 파일에서 사용할 프로필을 제어합니다.
      
      비어있으면 환경 변수 "AWS_PROFILE" 또는 설정되지 않은 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      multipart 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 파일 청크 수입니다.
      
      고속 링크에서 대용량 파일을 업로드하고 대역폭을 완전히 활용하지 않는 경우 이 값을 늘리면 전송 속도를 
      높일 수 있습니다.

   --force-path-style
      true 인 경우 path style 액세스를 사용하고 false 인 경우 가상 호스팅 스타일을 사용합니다.
      
      이 플래그가 true (기본값)이면 rclone은 path style 액세스를 사용하고 
      false이면 가상 path style을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro) 를 참조하세요.
      
      몇몇 공급자 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는 false로 설정해야합니다. 
      rclone은 제공자 설정에 따라이 작업을 자동으로 수행합니다.

   --v2-auth
      true 인 경우 v2 인증을 사용합니다.
      
      false 이면 (기본값) rclone은 v4 인증을 사용합니다.
      설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하세요. 예 : 이전 Jewel/v10 CEPH 입니다.

   --list-chunk
      객체 목록의 크기입니다. (각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 
      알려져 있습니다.
      대부분의 서비스는 요청보다 많은 인수가 요청되어도 응답 목록을 1000 개로 
      자르지만 메가 방식을 사용하여 분할 시 반환하는 수는 
      증가할 수 있으므로 이 값을 조정할 수도 있습니다. S3의 경우에는 
      전역 최대 값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 를 참조하세요.)
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전입니다 : 1,2 또는 자동으로 0입니다.
      
      S3가 처음 시작될 때는 ListObjects 호출로 버킷의 객체를 나열하는 기능만 제공했습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 
      이것은 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야합니다.
      
      기본 설정인 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 
      목록 개체 방법을 추측합니다. 그것이 잘못되었다고 추측하면 
      수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급자는 목록 URL 인코딩을 지원하고 있으며, 
      파일 이름에 제어 문자를 사용할 때 이것이 가능한 경우 파일 이름이 
      안정적 일 때 더 믿을 수 있습니다. 
      이 값을 unset으로 설정하면 rclone은 
      제공자 설정에 따라 적용할 것을 선택합니다. 
      그러나 여기서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      버킷을 확인하거나 생성하지 않도록 설정합니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 
      경우 유용할 수 있습니다.
      
      사용자가 버킷 생성 권한을 가지고 있지 않은 경우 필요할 
      수 있습니다. v1.52.0 이전에는 버그로 인해 체크하지 못하고 
      무시됐을 것입니다.
      

   --no-head
      업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다.
      
      rclone이 객체를 업로드 한 후 PUT 후에 200 OK 메시지를 
      수신하면 올바르게 업로드 된 것으로 가정할 것입니다.
      
      특히, 다음을 가정합니다.
      
      - 업로드 된 대로 메타 데이터, 수정 시간, 저장 클래스 및 콘텐츠 유형
      - 업로드 된 크기
      
      다음 항목을 "PUT"로부터 단일 파트 PUT에 대한 응답에서 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      대용량 파일의 소스 객체가 업로드 될 때 rclone은 HEAD 요청을 실행합니다.
      
      이 플래그를 설정하면 올바르게 업로드되지 않은 경우 올바르게 
      인식될 수 있는 업로드 실패 위험성이 높아지기 때문에 일반적인 
      작업에는 권장되지 않습니다. 실제로 업로드 실패의 위험률은 
      이 플래그를 사용하지 않아도 매우 작습니다.
      

   --no-head-object
      GET 작업 전에 HEAD 작업을 수행하지 않습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요 섹션 인코딩](/overview/#encoding) 를 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시 될 주기입니다.
      
      추가 버퍼가 필요한 업로드 (예 : multipart)은 할당을 위해 
      메모리 풀을 사용합니다. 이 옵션은 사용하지 않는 버퍼가 
      풀에서 제거되는 주기를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.
      
      현재 s3 (구체적으로 minio) 백엔드와 HTTP/2에 문제가 있습니다. 
      HTTP/2는 s3 백엔드의 기본 설정이지만 여기에서 비활성화할 
      수 있습니다. 이 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참고 : https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      이것은 보통 AWS S3가 거친 네트워크를 통해 다운로드 된 데이터에 
      대해 더 경제적인 전송 제공하기 위해 AWS S3로부터 제공되는 경로로 
      설정됩니다.

   --use-multipart-etag
      검증을 위해 multipart 업로드의 ETag을 사용할지 여부
      
      true, false 또는 설정되지 않음으로 제공자의 기본값을 사용할 것입니다.
      

   --use-presigned-request
      단일 파트 업로드에 대해 사전 서명 된 요청 또는 PutObject를 사용할지 여부
      
      false 일 때 rclone은 AWS SDK에서 
      PutObject를 사용하여 객체를 업로드합니다.
      
      1.59 보다 낮은 버전의 rclone은 단일 파트 객체를 업로드하기 위해 
      사전 서명 된 요청을 사용했으며 이 플래그를 true 로 설정하면 
      해당 기능을 다시 활성화합니다. 이는 특수한 경우 또는 테스트를 
      위해서여야합니다 .

   --versions
      디렉토리 목록에 오래된 버전을 포함합니다.

   --version-at
      지정된 시간에 버전별로 파일을 표시합니다.
      
      매개 변수는 날짜 "2006-01-02", datetime "2006-01-02 15:04:05" 
      또는 그만큼 이전을 나타내는 기간, 예 : "100d" 또는 "1h" 일 수 있습니다.
      
      이 옵션을 사용하는 경우 파일 작성 작업은 허용되지 않으므로 파일을 
      업로드하거나 삭제 할 수 없습니다.
      
      유효한 형식은 [시간 옵션 설명서](/docs/#time-option) 를 참조하세요.
      

   --decompress
      gzip으로 압축 된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 
      일반적으로 rclone은 이러한 파일을 압축 된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 받을 때 
      "Content-Encoding: gzip"로 압축을 해제합니다. 
      이로 인해 rclone은 크기와 해시를 확인할 수 없지만 
      파일 내용은 압축이 풀립니다.
      

   --might-gzip
      백엔드가 객체를 gzip으로 압축 할 수 있는 경우이를 설정합니다.
      
      일반적으로 공급자는 객체가 다운로드 될 때 객체를 변경하지 않습니다. 
      "Content-Encoding: gzip"로 업로드하지 않은 객체는 
      다운로드되면 설정되지 않습니다.
      
      그러나 일부 공급자 (예 : Cloudflare)는 "Content-Encoding: gzip"가 
      설정되지 않은 파일에 대해서도 객체를 gzip으로 압축 할 수 있습니다.
      
      이 경우 rclone이 chunked transfer encoding으로 
      Content-Encoding: gzip이 설정된 객체를 다운로드하고 
      압축을 풀게됩니다.
      
      unset으로 설정되면 (기본값) rclone은 공급자 설정에 따라 
      적용할 것을 선택합니다. 그러나 여기에서 rclone의 선택을 
      무시할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다


옵션:
   --access-key-id value        AWS Access Key ID입니다. [$ACCESS_KEY_ID]
   --acl value                  버킷을 생성하고 객체를 저장하거나 복사할 때 사용되는 사전 정의 ACL입니다. [$ACL]
   --endpoint value             S3 API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명을 받아옵니다 (환경 변수 또는 env 변수에 저장된 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  지역 제한 - 지역과 일치하도록 설정해야 합니다. [$LOCATION_CONSTRAINT]
   --region value               연결할 지역입니다. [$REGION]
   --secret-access-key value    AWS Secret Access Key입니다. [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷을 생성할 때 사용되는 사전 정의 ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크가 복사로 전환되는 용량 임계치입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 압축 된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true 인 경우 path style 액세스를 사용하고 false 인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               객체 목록의 크기입니다. (각 ListObject S3 요청에 대한 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전입니다 : 1,2 또는 자동으로 0입니다. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         Multipart 업로드에서 최대 파트 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시 될 주기입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축 할 수 있는 경우이를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷을 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET 작업 전에 HEAD 작업을 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       multipart 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크가 업로드로 전환되는 용량 임계치입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 multipart 업로드에 ETag을 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 대해 사전 서명 된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true 인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 버전별로 파일을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 오래된 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}