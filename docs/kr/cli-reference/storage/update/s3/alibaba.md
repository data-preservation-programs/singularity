# 알리바바클라우드 오브젝트 스토리지 시스템 (OSS) 이전에 Aliyun이라고 불리었습니다.

{% code fullWidth="true" %}
```
NAME:
   singularity storage update s3 alibaba - 알리바바클라우드 오브젝트 스토리지 시스템 (OSS) 이전에 Aliyun

USAGE:
   singularity storage update s3 alibaba [command options] <name|id>

DESCRIPTION:
   --env-auth
      런타임에서 AWS 인증 정보를 가져옵니다 (환경 변수 또는 env vars이 없는 경우 EC2/ECS 메타데이터).

      access_key_id와 secret_access_key이 비어있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 인증 정보를 입력합니다.
         | true  | 런타임 환경 (환경 변수 또는 IAM)에서 AWS 인증 정보를 가져옵니다.

   --access-key-id
      AWS Access Key ID.

      익명 액세스 또는 런타임 자격 증명인 경우 비워 둡니다.

   --secret-access-key
      AWS Secret Access Key (암호).

      익명 액세스 또는 런타임 자격 증명인 경우 비워 둡니다.

   --endpoint
      OSS API의 엔드포인트입니다.

      예시:
         | oss-accelerate.aliyuncs.com          | 글로벌 가속
         | oss-accelerate-overseas.aliyuncs.com | 글로벌 가속 (중국 본토 이외)
         | oss-cn-hangzhou.aliyuncs.com         | 중국 동부 1 (항주)
         | oss-cn-shanghai.aliyuncs.com         | 중국 동부 2 (상하이)
         | oss-cn-qingdao.aliyuncs.com          | 중국 북부 1 (칭다오)
         | oss-cn-beijing.aliyuncs.com          | 중국 북부 2 (베이징)
         | oss-cn-zhangjiakou.aliyuncs.com      | 중국 북부 3 (장집)
         | oss-cn-huhehaote.aliyuncs.com        | 중국 북부 5 (호호트)
         | oss-cn-wulanchabu.aliyuncs.com       | 중국 북부 6 (울란차부)
         | oss-cn-shenzhen.aliyuncs.com         | 중국 남부 1 (셴젠)
         | oss-cn-heyuan.aliyuncs.com           | 중국 남부 2 (허얀)
         | oss-cn-guangzhou.aliyuncs.com        | 중국 남부 3 (관저우)
         | oss-cn-chengdu.aliyuncs.com          | 중국 서부 1 (칭두)
         | oss-cn-hongkong.aliyuncs.com         | 홍콩 (홍콩)
         | oss-us-west-1.aliyuncs.com           | 미서부 1 (실리콘밸리)
         | oss-us-east-1.aliyuncs.com           | 미동부 1 (버지니아)
         | oss-ap-southeast-1.aliyuncs.com      | 동남아시아 1 (싱가포르)
         | oss-ap-southeast-2.aliyuncs.com      | 태평양남동 2 (시드니)
         | oss-ap-southeast-3.aliyuncs.com      | 동남아시아 3 (쿠알라룸푸르)
         | oss-ap-southeast-5.aliyuncs.com      | 태평양남동 5 (자카르타)
         | oss-ap-northeast-1.aliyuncs.com      | 태평양북동 1 (대일)
         | oss-ap-south-1.aliyuncs.com          | 태평양남부 1 (뭄바이)
         | oss-eu-central-1.aliyuncs.com        | 중앙유럽 1 (프랑크푸르트)
         | oss-eu-west-1.aliyuncs.com           | 서유럽 (런던)
         | oss-me-east-1.aliyuncs.com           | 중동 1 (두바이)

   --acl
      버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL입니다.

      이 ACL은 객체 생성에 사용되고 bucket_acl이 설정되지 않은 경우
      버킷 생성에도 사용됩니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.

      서버 측 복사로 객체를 복사할 때 S3는 ACL을 복사하지 않고 새로 작성합니다.

      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고
      기본값(비공개)이 사용됩니다.

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL입니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.

      acl이 설정되지 않은 경우에만 버킷 생성 시 적용됩니다.
      
      acl과 bucket_acl이 빈 문자열이면 X-Amz-Acl:
      헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 다른 사용자에게 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 액세스가 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스가 부여됩니다.
         |                    | 버킷에서 이 권한을 부여하는 것은 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스가 부여됩니다.

   --storage-class
      새로운 객체를 OSS에 저장할 때 사용할 저장 클래스입니다.

      예시:
         | <unset>     | 기본값
         | STANDARD    | 표준 저장 클래스
         | GLACIER     | 아카이브 저장 모드
         | STANDARD_IA | 자주 액세스되지 않는 저장 모드

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값입니다.

      이보다 큰 파일은 chunk_size로 청크 단위로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 (예: "rclone rcat"에서 가져온 파일이나 "rclone mount" 또는 google 사진이나 google 문서로 업로드된 파일)은
      이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.

      "--s3-upload-concurrency"의 청크는 이 크기만큼 전송 당 메모리에 버퍼링됩니다.

      대역폭이 높은 링크를 통해 큰 파일을 전송하고 충분한 메모리를 가지고 있다면, 이 값을 증가시키면 전송 속도가 향상됩니다.

      rclone은 알려진 크기의 큰 파일을 업로드할 때 10,000 청크 제한을 준수하기 위해 자동으로 청크 크기를 증가시킵니다.

      크기를 알 수 없는 파일은 설정된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 청크가 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다.
      더 큰 파일을 스트림 업로드하려면 chunk_size를 크게 설정해야 합니다.

      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태 통계의 정확도가 감소합니다. Rclone은 AWS SDK에서 버퍼에 청크가 전송된 것으로 취급하지만,
      실제로는 아직 업로드 중일 수 있습니다. 청크 크기가 클수록 AWS SDK 버퍼와 실제로는 다를 수 있는 큰 청크 사용량과 진행 상태 보고의 정확성이 떨어집니다.

   --max-upload-parts
      멀티파트 업로드에 사용할 최대 청크 수입니다.

      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 multipart 청크 수를 정의합니다.

      이 옵션은 10,000 청크의 AWS S3 사양을 지원하지 않는 서비스에 유용할 수 있습니다.

      알려진 크기의 큰 파일을 업로드할 때 rclone은 10,000 청크 제한을 준수하기 위해 자동으로 청크 크기를 증가시킵니다.

   --copy-cutoff
      멀티파트 복사로 전환하기 위한 임계값입니다.

      서버 측 복사가 필요한 이보다 큰 파일은 이 크기로 청크 단위로 복사됩니다.

      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체에 추가하기 때문에 큰 파일의 업로드 시작에 오랜 시간이 소요될 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은
      "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있으면
      현재 사용자의 홈 디렉터리로 기본 설정됩니다.

          리눅스/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 그 파일에서 사용된 프로필을 제어합니다.

      비워 두면 환경 변수 "AWS_PROFILE" 또는
      설정되지 않은 경우 "default"로 기본 설정됩니다.

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드의 동시성입니다.

      동시에 업로드되는 동일 파일의 청크 수입니다.

      대역폭을 충분히 활용하지 못하고 높은 속도의 링크로 소량의 큰 파일을
      업로드하는 경우 이 값을 증가시키면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다.

      true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고,
      false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 정보는
      [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.

      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는이를 설정하지 않으면
      false로 설정해야 합니다. Rclone은 공급자 설정에 기반하여 자동으로 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.

      false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정된 경우 v2 인증을 사용합니다.

      v4 시그니처가 작동하지 않는 경우에만 사용하십시오. 예: 이전 Jewel/v10 CEPH의 경우.

   --list-chunk
      리스트 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록 크기)입니다.

      이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
      대부분의 서비스는 요청한 것보다 많이 요청해도 응답 목록을 1000개로 자르지만 더 많이 요청한 경우에도 해당 응답 목록이 최대 1000개로 줄어듭니다.
      AWS S3에서는 전역 최대값이 이것으로 그대로이 고, [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는`rgw list buckets max chunk` 옵션으로 증가할 수 있습니다.

   --list-version
      ListObjects의 버전: 1, 2 또는 0 (자동)을 사용하십시오.

      S3가 처음에 출시될 때 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공하였습니다.

      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      성능이 훨씬 더 높으므로 가능한 경우 사용해야 합니다.

      기본값인 0으로 설정하면 rclone은 공급자에 설정된 값에 따라 어떤 목록 객체 방법을 호출할 것이라고 추측합니다. 잘못 추측하는 경우에는 여기에서 수동으로 설정할 수 있습니다.

   --list-url-encode
      리스팅에 대한 URL 인코딩 여부: true/false/unset

      일부 제공자는 URL 인코딩 리스트를 지원하고 가능한 경우 파일 이름에 제어 문자를 사용할 때 이것이 더 안정적입니다. 이것이 unset으로 설정된 경우 (기본값) rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.

   --no-check-bucket
      설정된 경우 버킷이 존재하는지 확인하거나 생성하지 않으려면 이 옵션을 사용하지 마십시오.

      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용합니다.

      버킷 생성 권한이 없는 사용자를 사용해야 할 경우에도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 무시되었습니다.

   --no-head
      업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 마십시오.

      rclone이 PUT로 객체를 업로드한 후 200 OK 메시지를 받으면 제대로 업로드되었다고 가정할 수 있도록 하는 것이 유용합니다.

      특히 다음을 가정합니다:

      - 업로드된 metadata(수정 시간, 저장 클래스 및 콘텐츠 유형)가 업로드한 것과 같았다.
      - 크기가 업로드한 것과 같았다.

      싱글 파트 PUT에 대한 응답의 다음 항목을 읽습니다:

      - MD5SUM
      - 업로드된 날짜

      멀티파트 업로드의 경우 이 항목은 읽지 않습니다.

      길이를 알 수 없는 원본 객체가 업로드되는 경우 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 올바르지 않은 크기와 같은 업로드 실패의 가능성이 높아지므로 정상적인 작업에는 권장되지 않습니다. 실제로 이 플래그를 설정해도 업로드 실패가 감지되지 않는 가능성은 매우 적습니다.
      
   --no-head-object
      객체를 가져오기 전에 GET 직전에 HEAD를 실행하지 마십시오.

   --encoding
      백엔드의 인코딩입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 얼마나 자주 플러시 할 것인지 결정합니다.

      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용할 것입니다.
      이 옵션은 사용되지 않은 버퍼를 풀에서 제거하는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.

      현재 s3(특히 minio) 백엔드와 HTTP/2 문제가 아직 해결되지 않은 상태입니다. S3 백엔드의 HTTP/2는 기본적으로 활성화되지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우 더 저렴한 이그레스를 제공하므로 일반적으로 CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      ETag를 멀티파트 업로드에서 확인에 사용할지 여부

      이는 true, false 또는 빈 값으로 설정해야 합니다.

      기본값을 사용하려면 true, false 또는 설정하지 않으십시오.

   --use-presigned-request
      단일 파트 업로드에 사전 서명 요청 또는 PutObject을 사용할지 여부

      이 값이 false이면 rclone은 객체를 업로드하는 데 AWS SDK에서 PutObject를 사용할 것입니다.

      rclone < 1.59 버전은 단일 파트 객체를 업로드하기 위해 사전 서명 요청을 사용하고이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다.
      이는 예외적인 상황이나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간의 파일 버전과 동일한 방식으로 파일 버전을 표시합니다.

      매개변수는 날짜 "2006-01-02", datetime "2006-01-02
      15:04:05" 또는 그보다 오래된 지속 시간, 예: "100d" 또는 "1h" 일 수 있습니다.

      이를 사용하는 경우 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      사용 가능한 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

   --decompress
      Gzip으로 인코딩된 객체를 압축 해제합니다.

      S3에 "Content-Encoding: gzip"로 업로드된 객체를 다운로드할 수 있습니다. 일반적으로 rclone은 이러한 파일을 압축 해제하여 "Content-Encoding: gzip"로 받습니다. 이렇게 되면 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하십시오.

      일반적으로 제공업체는 객체를 다운로드할 때 객체를 변경하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.

      그러나 일부 제공자는 `Content-Encoding: gzip`로 업로드되지 않은 객체도 gzip으로 압축할 수 있습니다(Cloudflare와 같은 경우).
      
      이러한 경우에는
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      와 같은 오류가 발생합니다.
      
      이 플래그를 설정하고 rclone이 `Content-Encoding: gzip`가 설정되고 청크 전송 인코딩으로 객체를 다운로드하면 rclone은 해당 객체를 실시간으로 압축 해제합니다.
      
      unset으로 설정된 경우 (기본값) rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제

OPTIONS:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value           OSS API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 인증 정보를 가져옵니다 (환경 변수 또는 env vars이 없는 경우 EC2/ECS 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --secret-access-key value  AWS Secret Access Key (암호). [$SECRET_ACCESS_KEY]
   --storage-class value      새로운 객체를 OSS에 저장할 때 사용할 저장 클래스입니다. [$STORAGE_CLASS]

   고급

   --bucket-acl value               버킷 생성 시 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하기 위한 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     Gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               리스트 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록 크기)입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          리스팅에 대한 URL 인코딩 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             ListObjects의 버전: 1,2 or 0 (자동) (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에 사용할 최대 청크 수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 얼마나 자주 플러시 할 것인지 결정합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하십시오. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                설정된 경우 버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체의 무결성을 확인하기 위해 HEAD를 수행하지 마십시오. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져오기 전에 GET 직전에 HEAD를 실행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       ETag를 멀티파트 업로드에서 확인에 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          사전 서명 요청 또는 PutObject를 사용하여 단일 파트 업로드를 할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               파일 버전을 지정된 시간의 상태와 동일하게 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}