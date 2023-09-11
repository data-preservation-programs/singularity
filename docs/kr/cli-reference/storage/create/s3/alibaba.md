# Alibaba Cloud 객체 스토리지 시스템 (OSS) - 이전에는 Aliyun이라고 불렸습니다.

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 alibaba - Alibaba Cloud 객체 스토리지 시스템 (OSS) - 이전에는 Aliyun이라고 불렸습니다.

사용법:
   singularity storage create s3 alibaba [command options] [arguments...]

설명:
   --env-auth
      런타임(환경 변수 또는 env vars나 IAM이 없는 경우 EC2/ECS 메타데이터)에서 AWS 자격 증명 가져오기.
      
      access_key_id와 secret_access_key가 비어 있을 때만 적용됩니다.

      예:
         | false | AWS 자격 증명을 다음 단계에 입력합니다.
         | true  | AWS 자격 증명을 환경에서 가져옵니다(환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호).
      
      익명 액세스 또는 런타임 자격 증명의 경우 비워 둡니다.

   --endpoint
      OSS API의 엔드포인트.

      예:
         | oss-accelerate.aliyuncs.com          | 글로벌 가속
         | oss-accelerate-overseas.aliyuncs.com | (중국 본토 외부) 글로벌 가속
         | oss-cn-hangzhou.aliyuncs.com         | 중국 동부 1(항저우)
         | oss-cn-shanghai.aliyuncs.com         | 중국 동부 2(상하이)
         | oss-cn-qingdao.aliyuncs.com          | 중국 북부 1(칭다오)
         | oss-cn-beijing.aliyuncs.com          | 중국 북부 2(베이징)
         | oss-cn-zhangjiakou.aliyuncs.com      | 중국 북부 3(장강 구)
         | oss-cn-huhehaote.aliyuncs.com        | 중국 북부 5(흔헤 허허트)
         | oss-cn-wulanchabu.aliyuncs.com       | 중국 북부 6(울란차부)
         | oss-cn-shenzhen.aliyuncs.com         | 중국 남부 1(선전)
         | oss-cn-heyuan.aliyuncs.com           | 중국 남부 2(헬렌)
         | oss-cn-guangzhou.aliyuncs.com        | 중국 남부 3(광저우)
         | oss-cn-chengdu.aliyuncs.com          | 중국 서부 1(청두)
         | oss-cn-hongkong.aliyuncs.com         | 홍콩(홍콩)
         | oss-us-west-1.aliyuncs.com           | US West 1(실리콘 밸리)
         | oss-us-east-1.aliyuncs.com           | US East 1(버지니아)
         | oss-ap-southeast-1.aliyuncs.com      | 동남 아시아 동남 1(싱가포르)
         | oss-ap-southeast-2.aliyuncs.com      | 아시아 태평양 동남 2(시드니)
         | oss-ap-southeast-3.aliyuncs.com      | 동남 아시아 동남 3(쿠알라룸푸르)
         | oss-ap-southeast-5.aliyuncs.com      | 아시아 태평양 동남 5(자카르타)
         | oss-ap-northeast-1.aliyuncs.com      | 아시아 태평양 북동 1(일본)
         | oss-ap-south-1.aliyuncs.com          | 아시아 태평양 남부 1(뭄바이)
         | oss-eu-central-1.aliyuncs.com        | 중앙 유럽 1(프랑크푸르트)
         | oss-eu-west-1.aliyuncs.com           | 서부 유럽(런던)
         | oss-me-east-1.aliyuncs.com           | 중동 1(두바이)

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL.
      
      이 ACL은 객체를 생성할 때와 bucket_acl이 설정되어 있지 않은 경우에도 사용됩니다.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      S3는 ACL을 확인하지 않지만 새로 쓴다는 점을 유의하십시오(소스에서 ACL을 복사하지 않음).
      
      ACL이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(비공개)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL.
      
      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값(비공개)이 사용됩니다.
      

      예:
         | private            | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | 다른 사람은 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹은 읽기 액세스 권한이 있습니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스 권한이 있습니다.
         |                    | 버킷에서이 기능을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 있습니다.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스 권한이 있습니다.

   --storage-class
      OSS에 새로운 객체를 저장할 때 사용할 저장 클래스.

      예:
         | <unset>     | 기본값
         | STANDARD    | 표준 저장 클래스
         | GLACIER     | 보관 저장 모드
         | STANDARD_IA | 편리한 액세스 저장 모드

   --upload-cutoff
      청크 끊어짐으로 전환되는 파일 크기.
      
      이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.
      
      upload_cutoff보다 크거나 알 수없는 크기의 파일(예 : "rclone rcat" 또는 "rclone mount" 또는 구글 사진 또는 구글 문서에서 업로드된 파일)을 업로드하는 경우, 이 청크 크기를 사용하여 여러 부분으로 업로드됩니다.
      
      주의할 점은 transfer별로 "--s3-upload-concurrency" 청크 크기 만큼의 청크가 transfer당 메모리에 버퍼링됩니다.
      
      고속 링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이를 더 크게 설정하면 전송 속도가 향상됩니다.
      
      rclone은 알려진 크기의 대규모 파일을 전송할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 유지합니다.
      
      크기를 알 수없는 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기가 5 MiB이고 최대 10,000개의 청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트리밍 업로드하려면 청크 크기를 증가시켜야합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 감소합니다. rclone은 실제로 아직 업로드 중일 수있는 경우에도 AWS SDK에서 버퍼로 처리되고 있을 때 청크가 전송된 것으로 처리합니다. 더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행을 표시하는 데 진실과 더 거리가 있는 진행 보고를 의미합니다.
      

   --max-upload-parts
      여러 부분 업로드에 대한 최대 부분 수.
      
      이 옵션은 여러 부분 업로드 시 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      AWS S3 10,000 청크 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대규모 파일을 업로드 할 때 해당 청크 수 제한을 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      Multipart 복사로 전환되는 파일 크기 컷 오프.
      
      이보다 큰 파일을 서버 간 복사해야하는 경우 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가하여 데이터 무결성 확인에 유용하지만 큰 파일은 업로드를 시작하기 전에 긴 지연을 야기할 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.
      
      env_auth = true의 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. env 값이 비어 있으면 현재 사용자의 홈 디렉토리가 기본값으로 사용됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true의 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에 사용되는 프로필을 제어합니다.
      
      비어 있으면 기본값으로 "AWS_PROFILE" 또는 "default" 환경 변수를 사용합니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      다중 부분 업로드에 대한 동시성.
      
      동일한 파일의 청크 수를 동시에 업로드합니다.
      
      넓대역망에서 대용량 파일을 빠른 속도로 전송하고 대역폭을 완전히 활용하지 못하는 경우 이를 늘리면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다.
      
      이것이 true인 경우(기본값) rclone은 경로 스타일 액세스를 사용하고 false이면 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서]를 참조하십시오. (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
      
      일부 공급 업체(AWS, Aliyun OSS, Netease COS 또는 Tencent COS)는이를 false로 설정해야합니다. rclone은 제공 업체 설정에 따라 자동으로 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정되면 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만이 옵션을 사용하십시오. 예 : pre Jewel/v10 CEPH.

   --list-chunk
      목록 청크의 크기(각 ListObject S3 요청에 대한 응답 목록).
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청이 1,000개 이상이라도 응답 목록을 1,000개로 자름.
      AWS S3에서는이 것은 전역 최대이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 통해 이 값을 늘릴 수 있습니다.
      

   --list-version
      사용할 ListObjects의 버전 : 1,2 또는 0(auto)입니다.
      
      S3가 처음 시작되었을 때 버킷의 객체를 열거하기 위해 ListObjects 호출 만 제공되었습니다.
      
      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하며 가능한 경우 사용해야합니다.
      
      기본값인 0으로 설정하면 rclone은 공급자 설정에 따라 호출할 목록 개체 메서드를 추측합니다. 추측이 잘못되면 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL로 인코딩할지 여부: true/false/unset
      
      몇몇 공급 업체는 파일 이름에 제어 문자를 사용할 때 URL로 인코딩 목록을 지원하며 가능한 경우 파일 내용 압축이 더 안정적입니다. 설정되지 않으면 (기본값) rclone은 provider 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 무시할 수 있습니다.
      

   --no-check-bucket
      설정되면 버킷의 존재를 확인하거나 생성하지 않으려고 시도하지 않습니다.
      
      버킷이 이미 존재하는 것을 알고있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용 할 수 있습니다.
      
      사용자가 버킷 생성 권한이없는 경우에도 필요할 수 있습니다. v1.52.0 이전에는이 버그로 인해 무음 통과되었습니다.
      

   --no-head
      업로드된 객체의 정합성을 확인하기 위해 HEAD를 수행하지 않습니다.
      
      rclone은 PUT 후 200 OK 메시지를 수신하면 제대로 업로드되었다고 가정하기 때문에 이 옵션을 설정하면 rclone은 PUT로 객체를 업로드 한 후 제대로 업로드되었다고 가정합니다.
      
      특히 다음을 가정합니다.
      
      - 메타데이터(수정 시간, 저장 클래스 및 콘텐츠 유형)이 업로드 한 것과 같음
      - 크기가 업로드 한 것과 같음
      
      PUT로 단일 부분을 업로드 할 때 다음 항목을 읽습니다.

      - MD5SUM
      - 업로드된 날짜
      
      여러 부분의 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      알 수없는 길이의 원본 객체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 중 무시 된 업로드 오류의 가능성이 높아지므로 정상 작동에는 권장하지 않습니다. 사실이 플래그로 인해 전송되지 않은 업로드 오류의 가능성은 매우 적습니다.
      

   --no-head-object
      개체를 가져오기 전에 HEAD를 수행하지 마십시오.

   --encoding
      백엔드에 대한 인코딩.
      
      개요의 [인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.
      
      추가 버퍼를 필요로하는 업로드(예 : 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 HTTP2 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. S3 백엔드의 HTTP/2는 기본적으로 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드에 대한 사용자 지정 엔드포인트.
      보통 AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 배출을 제공합니다.

   --use-multipart-etag
      여러 부분 업로드에서 ETag를 검증에 사용할지 여부
      
      true, false 또는 provider에 대한 기본값으로 설정하는 것이 좋습니다.
      

   --use-presigned-request
      단일 부분 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부
      
      이 flag를 false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 객체를 업로드합니다.
      
      rclone 버전 < 1.59는 단일 부분 객체 업로드에 사전 서명된 요청을 사용하고이 flag를 true로 설정하면이 기능을 다시 활성화합니다. 이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.
      

   --versions
      디렉터리 목록에 이전 버전 포함.

   --version-at
      지정된 시간에 파일 버전 표시.
      
      매개 변수는 날짜 "2006-01-02", 날짜 시간 "2006-01-02 15:04:05" 또는 그만큼 먼 시간에 대한 기간, 예를 들어 "100d" 또는 "1h" 일 수 있습니다.
      
      이를 사용하는 경우 파일 쓰기 작업이 허용되지 않으므로 파일 업로드 또는 삭제를 할 수 없습니다.
      
      유효한 형식에 대한 자세한 정보는 [time 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      설정하면 gzip으로 인코딩된 개체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축 해제하여 "Content-Encoding: gzip"로 수신될 때이 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 gzip 객체를 수신할 수 있으면 설정하십시오.
      
      일반적으로 공급 업체는 객체를 다운로드 할 때 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 개체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 공급 업체는 "Content-Encoding: gzip"로 업로드하지 않았더라도 개체를 gzip으로 압축 할 수 있습니다(예 : Cloudflare).
      
      이와 같은 경우 12121 같은 오류를 받을 수 있습니다.
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩이 설정된 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.
      
      이 값을 설정하지 않으면(기본값) rclone은 provider 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터 설정 및 읽기를 억제합니다


OPTIONS:
   --access-key-id value      AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                버킷 생성 및 저장 또는 복사 개체에 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value           OSS API의 엔드포인트. [$ENDPOINT]
   --env-auth                 런타임(환경 변수 또는 env vars나 IAM이 없는 경우 EC2/ECS 메타데이터)에서 AWS 자격 증명 가져오기입니다. (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS 비밀 액세스 키(비밀번호). [$SECRET_ACCESS_KEY]
   --storage-class value      새로운 객체를 저장할 때 사용할 저장 클래스입니다. [$STORAGE_CLASS]

   고급

   --bucket-acl value               버킷 생성에 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              Multipart 복사로 전환되는 파일 크기입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정하면 gzip으로 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 HTTP2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 지정 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL로 인코딩할지 여부입니다. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects의 버전입니다. 1,2 또는 0(auto) (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         여러 부분 업로드에 대한 최대 부분 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 gzip 객체를 수신할 수 있으면 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정되면 버킷의 존재를 확인하거나 생성하지 않으려고 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 개체의 정합성을 확인하기 위해 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져오기 전에 HEAD를 수행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       다중 부분 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 끊어짐으로 전환되는 파일 크기입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       여러 부분 업로드에서 ETag를 검증에 사용할지 여부입니다. (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부입니다. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에서 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전 포함. (기본값: false) [$VERSIONS]

   일반

   --name value  스토리지의 이름(기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}