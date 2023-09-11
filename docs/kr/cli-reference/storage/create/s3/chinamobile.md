# 중국 모바일 이클라우드 엘라스틱 오브젝트 스토리지 (EOS)

{% code fullWidth="true" %}
```
명령어:
   singularity storage create s3 chinamobile - 중국 모바일 이클라우드 엘라스틱 오브젝트 스토리지 (EOS)

사용법:
   singularity storage create s3 chinamobile [명령어 옵션] [인수...]

설명:
   --env-auth
      런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타데이터)에서 AWS 자격 증명을 가져옵니다.
      
      access_key_id 및 secret_access_key가 비어 있을 때만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

   --access-key-id
      AWS 액세스 키 ID를 입력하세요.
      
      익명 액세스 또는 런타임 자격 증명을 사용하지 않으려면 비워 두세요.

   --secret-access-key
      AWS Secret Access Key(암호)를 입력하세요.
      
      익명 액세스 또는 런타임 자격 증명을 사용하지 않으려면 비워 두세요.

   --endpoint
      중국 모바일 이클라우드 엘라스틱 오브젝트 스토리지 (EOS) API의 엔드포인트를 입력하세요.

      예시:
         | eos-wuxi-1.cmecloud.cn      | 기본 엔드포인트(인증되지 않을 때 권장)
         |                             | 동부 중국(수주)
         | eos-jinan-1.cmecloud.cn     | 동부 중국(진남)
         | eos-ningbo-1.cmecloud.cn    | 동부 중국(항저우)
         | eos-shanghai-1.cmecloud.cn  | 동부 중국(상하이-1)
         | eos-zhengzhou-1.cmecloud.cn | 중부 중국(정저우)
         | eos-hunan-1.cmecloud.cn     | 중부 중국(청사-1)
         | eos-zhuzhou-1.cmecloud.cn   | 중부 중국(청사-2)
         | eos-guangzhou-1.cmecloud.cn | 남부 중국(광저우-2)
         | eos-dongguan-1.cmecloud.cn  | 남부 중국(광저우-3)
         | eos-beijing-1.cmecloud.cn   | 북부 중국(베이징-1)
         | eos-beijing-2.cmecloud.cn   | 북부 중국(베이징-2)
         | eos-beijing-4.cmecloud.cn   | 북부 중국(베이징-3)
         | eos-huhehaote-1.cmecloud.cn | 북부 중국(호흐하테)
         | eos-chengdu-1.cmecloud.cn   | 남서부 중국(청두)
         | eos-chongqing-1.cmecloud.cn | 남서부 중국(충칭)
         | eos-guiyang-1.cmecloud.cn   | 남서부 중국(구이양)
         | eos-xian-1.cmecloud.cn      | 남서부 중국(시안)
         | eos-yunnan.cmecloud.cn      | 윈난 중국(쿤밍)
         | eos-yunnan-2.cmecloud.cn    | 윈난 중국(쿤밍-2)
         | eos-tianjin-1.cmecloud.cn   | 톈진 중국(톈진)
         | eos-jilin-1.cmecloud.cn     | 지린 중국(창춘)
         | eos-hubei-1.cmecloud.cn     | 후베이 중국(샹양)
         | eos-jiangxi-1.cmecloud.cn   | 강시 중국(난창)
         | eos-gansu-1.cmecloud.cn     | 간수 중국(난징)
         | eos-shanxi-1.cmecloud.cn    | 산시 중국(타이위안)
         | eos-liaoning-1.cmecloud.cn  | 랴오닝 중국(선양)
         | eos-hebei-1.cmecloud.cn     | 허베이 중국(시징짜강)
         | eos-fujian-1.cmecloud.cn    | 후난 중국(샤먼)
         | eos-guangxi-1.cmecloud.cn   | 광시 중국(난닝)
         | eos-anhui-1.cmecloud.cn     | 안후이 중국(화난)

   --location-constraint
      엔드포인트와 일치해야 하는 위치 제약 조건을 입력하세요.
      
      버킷 생성 시에만 사용됩니다.

      예시:
         | wuxi1      | 동부 중국(수주)
         | jinan1     | 동부 중국(진남)
         | ningbo1    | 동부 중국(항저우)
         | shanghai1  | 동부 중국(상하이-1)
         | zhengzhou1 | 중부 중국(정저우)
         | hunan1     | 중부 중국(청사-1)
         | zhuzhou1   | 중부 중국(청사-2)
         | guangzhou1 | 남부 중국(광저우-2)
         | dongguan1  | 남부 중국(광저우-3)
         | beijing1   | 북부 중국(베이징-1)
         | beijing2   | 북부 중국(베이징-2)
         | beijing4   | 북부 중국(베이징-3)
         | huhehaote1 | 북부 중국(호흐하테)
         | chengdu1   | 남서부 중국(청두)
         | chongqing1 | 남서부 중국(충칭)
         | guiyang1   | 남서부 중국(구이양)
         | xian1      | 남서부 중국(시안)
         | yunnan     | 윈난 중국(쿤밍)
         | yunnan2    | 윈난 중국(쿤밍-2)
         | tianjin1   | 톈진 중국(톈진)
         | jilin1     | 지린 중국(창춘)
         | hubei1     | 후베이 중국(샹양)
         | jiangxi1   | 강시 중국(난창)
         | gansu1     | 간수 중국(난징)
         | shanxi1    | 산시 중국(타이위안)
         | liaoning1  | 랴오닝 중국(선양)
         | hebei1     | 허베이 중국(시징짜강)
         | fujian1    | 후난 중국(샤먼)
         | guangxi1   | 광시 중국(난닝)
         | anhui1     | 안후이 중국(화난)

   --acl
      버킷 및 오브젝트 생성 또는 복사 시에 사용되는 canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며, bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      주의: 이 ACL은 S3에서 server-side로 객체를 복사할 때 적용됩니다.
      소스에서 ACL을 복사하지 않고 새로운 ACL을 작성합니다.
      
      만약 acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고
      기본 값(Private)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시에 사용되는 canned ACL입니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하세요.
      
      이 ACL은 버킷을 생성할 때만 적용됩니다.
      설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      acl 및 bucket_acl이 빈 문자열인 경우 X-Amz-Acl:
      헤더가 추가되지 않고 기본 값(Private)이 사용됩니다.
      

      예시:
         | private            | 소유자는 FULL_CONTROL 권한을 얻음.
         |                    | 다른 사람은 액세스 권한을 가지지 않음(기본).
         | public-read        | 소유자는 FULL_CONTROL 권한을 얻음.
         |                    | AllUsers 그룹은 읽기 액세스를 얻음.
         | public-read-write  | 소유자는 FULL_CONTROL 권한을 얻음.
         |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻음.
         |                    | 버킷에 대해 적용하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자는 FULL_CONTROL 권한을 얻음.
         |                    | AuthenticatedUsers 그룹은 읽기 액세스를 얻음.

   --server-side-encryption
      S3에 객체를 저장할 때 사용되는 server-side 암호화 알고리즘입니다.

      예시:
         | <unset> | 암호화 안 함
         | AES256  | AES256

   --sse-customer-algorithm
      SSE-C를 사용하는 경우, S3에 객체를 저장할 때 사용되는 server-side 암호화 알고리즘입니다.

      예시:
         | <unset> | 암호화 안 함
         | AES256  | AES256

   --sse-customer-key
      SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는 데 사용하는 비밀 암호화 키를 제공할 수 있습니다.
      
      또는 --sse-customer-key-base64을 제공할 수 있습니다.

      예시:
         | <unset> | 암호화 안 함

   --sse-customer-key-base64
      SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는 데 사용할 비밀 암호화 키를 Base64로 인코딩된 형식으로 제공해야 합니다.
      
      또는 --sse-customer-key를 제공할 수 있습니다.

      예시:
         | <unset> | 암호화 안 함

   --sse-customer-key-md5
      SSE-C를 사용하는 경우, 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항).
      
      비어 있으면 sse_customer_key에서 자동으로 계산됩니다.
      

      예시:
         | <unset> | 암호화 안 함

   --storage-class
      ChinaMobile에 새로운 객체를 저장할 때 사용할 스토리지 클래스를 지정하세요.

      예시:
         | <unset>     | 기본(Default)
         | STANDARD    | 표준 스토리지 클래스
         | GLACIER     | 아카이브 스토리지 모드
         | STANDARD_IA | 드물게 액세스하는 스토리지 모드

   --upload-cutoff
      청크 업로드로 전환하는 임계값을 입력하세요.
      
      이보다 큰 파일은 chunk_size로 설정된 크기의 청크로 업로드됩니다.
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기를 지정하세요.
      
      upload_cutoff보다 큰 파일이나 크기를 모르는 파일("rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드된 파일 등)을 업로드할 때는 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      참고: "--s3-upload-concurrency" 개수의 이 크기의 청크가 전송당 메모리에 버퍼링됩니다.
      
      고속 링크로 대용량 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 높이면 전송 속도가 향상됩니다.
      
      rclone은 알려진 큰 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000 청크 제한을 맞추려고 합니다.
      
      알 수 없는 크기의 파일은 구성된 chunk_size로 업로드됩니다.
      기본 청크 크기가 5 MiB이며 최대 10,000 청크가 가능하기 때문에 기본적으로 Stream 업로드 가능한 최대 파일 크기는 48 GiB입니다.
      더 큰 파일을 Stream 업로드하려면 chunk_size를 증가시켜야 합니다.
      
      청크 크기가 증가하면 "-P" 플래그와 함께 표시되는 진행 정보의 정확성이 감소합니다. rclone은
      AWS SDK에 의해 버퍼링된 청크를 보냈다고 생각하며 실제로 업로드될 때가 될 수 있지만,
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 보고지연을 의미합니다.
      

   --max-upload-parts
      멀티파트 업로드에서 최대 청크 수를 지정하세요.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      이는 10,000 청크와 같은 AWS S3 사양을 지원하지 않는 서비스에 유용할 수 있습니다.
      
      rclone은 알려진 큰 파일을 업로드하면 최대 청크 수 제한을 유지하기 위해 청크 크기를 자동으로 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값을 입력하세요.
      
      최대 크기의 파일은 해당 크기로 청크로 복사됩니다.
      
      최소 값은 0이고 최대 값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 
      이는 데이터 무결성 확인에 유용하지만 큰 파일의 업로드를 시작하는 데 오랜 시간이 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로를 입력하세요.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 
      환경 값이 비어 있으면 현재 사용자의 홈 디렉터리로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필을 입력하세요.
      
      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 
      이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      빈 값으로 설정하면 환경 변수 "AWS_PROFILE" 또는 설정되지 않은 경우 "default"로 설정됩니다.
      

   --session-token
      AWS 세션 토큰을 입력하세요.

   --upload-concurrency
      멀티파트 업로드의 동시성을 입력하세요.
      
      동일한 파일의 청크 개수를 동시에 업로드합니다.
      
      고속 링크를 통해 고속 링크를 사용하여 대량의 큰 파일을 업로드하고 이러한 업로드가 대역폭을 충분히 활용하지 못하는 경우
      이 값을 높이면 전송 속도를 높일 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다.
      
      이 값이 true(기본값)이면 rclone은 경로 스타일 액세스를 사용하고
      false이면 rclone은 가상 경로 스타일 액세스를 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하세요.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 이 값을 false로 설정해야 합니다.
      rclone은 알아서 제공자 설정에 따라 자동으로 처리합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      이 값이 false(기본값)이면 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 시그니처가 작동하지 않는 경우에만 이 값만 사용하세요. 예: Jewel/v10 CEPH 이전 버전.

   --list-chunk
      목록 패치 크기(각 ListObject S3 요청에 대한 응답 목록 크기)를 입력하세요.
      
      이 옵션은 AWS S3 사양의 "MaxKeys", "max-items" 또는 "page-size"로 알려진 것과 같습니다.
      대부분의 서비스는 요청된 개수가 1000개 이상인 경우에도 응답 목록을 1000개로 축소합니다.
      AWS S3에서 전역 최대값이고 값을 변경할 수 없으므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하세요.
      Ceph에서는 "rgw list buckets max chunk" 옵션을 사용하여 이 값을 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전을 지정하세요: 1,2 또는 자동으로 0을 입력하세요.
      
      S3가 처음 시작된 시점에는 버킷에서 개체를 나열하는 데 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이 호출은
      성능이 훨씬 높으므로 가능하면 사용해야 합니다.
      
      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 ListObjects 메서드 호출 방법을 추측합니다.
      잘못 추측하면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL로 인코딩할지 여부: true/false/unset을 입력하세요.
      
      일부 공급자는 목록을 URL로 인코딩하고 가능한 경우 파일 이름에 제어 문자를 사용할 때 
      이 기능을 사용할 수 있습니다. 이 값을 unset(기본값)으로 설정하면 제공자 설정에 따라 rclone이 
      결정합니다.

   --no-check-bucket
      버킷의 존재를 확인하거나 생성을 시도하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우에 유용합니다.
      
      사용자에게 버킷 생성 권한이 없는 경우에도 필요할 수 있습니다.
      v1.52.0 이전에는 버그 때문에 조용히 통과되었습니다.

   --no-head
      청크의 무결성을 확인하기 위해 업로드된 객체의 HEAD를 하지 않습니다.
      
      rclone은 200 OK 메시지를 받으면 PUT로 객체를 업로드한 후 정상적으로 업로드된 것으로 가정합니다
      이 플래그를 설정하면 rclone은 PUT 후 200 OK 메시지를 수신하면 객체를 정상적으로 업로드된 것으로 간주합니다.
      
      특히 rclone은 다음 항목을 프롬프트 응답에서 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티 파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      크기를 모르는 소스 개체가 업로드되는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 업로드 실패의 가능성이 증가하며, 특히 잘못된 크기일 경우이며, 일반적인 운영에는 권장되지 않습니다.
      실제로 업로드 실패 가능성은 매우 적습니다.

   --no-head-object
      개체를 가져올 때 GET 전에 HEAD를 실행하지 않습니다.

   --encoding
      백엔드의 인코딩을 입력하세요.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀의 플러시 간격을 입력하세요.
      
      추가적인 버퍼를 필요로하는 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부를 입력하세요.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3(특히 minio) 백엔드와 HTTP/2에 관한 해결되지 않은 문제가 있습니다. 
      s3 백엔드의 HTTP/2는 기본적으로 활성화되지만 여기에서 비활성화할 수 있습니다.
      문제가 해결될 때이 플래그는 제거됩니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드하는 경우
      더 저렴한 이중 출고를 제공하므로 CloudFront CDN URL로 설정합니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 검증할지 여부
      
      true, false 또는 기본값을 사용하려면 true, false 또는 설정하지 않으면 공급자에 따라 기본값을 사용합니다.

   --use-presigned-request
      싱글 파트 업로드에 서명된 요청 또는 PutObject을 사용할지 여부
      
      이 값을 false로 설정하면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject을 사용합니다.
      
      rclone 버전 < 1.59에서는 싱글 파트 객체를 업로드하기 위해 서명된 요청을 사용하며, 
      이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 
      이는 특정 경우나 테스트를 제외하고는 필요하지 않습니다.

   --versions
      디렉터리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간대의 파일 버전을 표시합니다.
      
      매개변수는 날짜("2006-01-02"), 날짜시간("2006-01-02 15:04:05") 또는 해당 시간 전의 기간("100d" 또는 "1h")일 수 있습니다.
      
      이 값을 사용하는 경우 파일 쓰기 작업을 허용하지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [시간 옵션 문서](/docs/#time-option)를 참조하세요.

   --decompress
      객체에 gzip 인코딩이 적용된 경우 이를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 S3에 객체를 업로드할 수 있습니다. 
      일반적으로 rclone은 이러한 파일을 압축 해제하여 "Content-Encoding: gzip"으로 수신합니다. 
      이는 rclone이 크기와 해시를 확인할 수 없지만 파일 컨텐츠는 압축 해제됩니다.

   --might-gzip
      백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 플래그를 설정하세요.
      
      일반적으로 공급자는 개체를 다운로드할 때 개체를 수정하지 않습니다. 
      "Content-Encoding: gzip"로 업로드되지 않은 개체에는 다운로드되지 않습니다.
      
      그러나 일부 공급자(예: Cloudflare)는 `Content-Encoding: gzip`로 업로드되지 않은 개체도 
      압축할 수 있습니다.
      
      이로 인해 다음과 같은 오류가 발생할 수 있습니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하면 rclone은 chunked 전송 인코딩과 `Content-Encoding: gzip`로 개체를 다운로드하면 
      개체를 실시간으로 압축 해제합니다.
      
      이를 unset으로 설정하면 rclone은 설정된 대로 적용할 것입니다. 
      그러나 이 값을 재설정할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


옵션:
   --access-key-id value           AWS 액세스 키 ID를 입력하세요. [$ACCESS_KEY_ID]
   --acl value                     버킷 및 오브젝트를 생성하거나 저장 또는 복사할 때 사용되는 canned ACL입니다. [$ACL]
   --endpoint value                중국 모바일 이클라우드 엘라스틱 오브젝트 스토리지 (EOS) API의 엔드포인트를 입력하세요. [$ENDPOINT]
   --env-auth                      런타임(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타데이터)에서 AWS 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                      도움말 표시
   --location-constraint value     엔드포인트와 일치해야 하는 위치 제약 조건을 입력하세요. [$LOCATION_CONSTRAINT]
   --secret-access-key value       AWS Secret Access Key(암호)를 입력하세요. [$SECRET_ACCESS_KEY]
   --server-side-encryption value  S3에 객체를 저장할 때 사용되는 server-side 암호화 알고리즘입니다. [$SERVER_SIDE_ENCRYPTION]
   --storage-class value           ChinaMobile에 새로운 객체를 저장할 때 사용할 스토리지 클래스를 지정하세요. [$STORAGE_CLASS]

   고급

   --bucket-acl value               버킷 생성 시에 사용되는 canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기를 지정하세요. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 임계값을 입력하세요. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩을 입력하세요. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 패치 크기(각 ListObject S3 요청에 대한 응답 목록 크기)를 입력하세요. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL로 인코딩할지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전을 지정하세요: 1,2 또는 자동으로 0을 입력하세요. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 최대 청크 수를 지정하세요. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀의 플러시 간격을 입력하세요. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 플래그를 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성을 시도하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        청크의 무결성을 확인하기 위해 업로드된 객체의 HEAD를 하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져올 때 GET 전에 HEAD를 실행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필을 입력하세요. [$PROFILE]
   --session-token value            AWS 세션 토큰을 입력하세요. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로를 입력하세요. [$SHARED_CREDENTIALS_FILE]
   --sse-customer-algorithm value   SSE-C를 사용하는 경우, S3에 객체를 저장할 때 사용되는 server-side 암호화 알고리즘입니다. [$SSE_CUSTOMER_ALGORITHM]
   --sse-customer-key value         SSE-C를 사용하는 경우, 데이터를 암호화/복호화하는 데 사용하는 비밀 암호화 키를 제공할 수 있습니다. [$SSE_CUSTOMER_KEY]
   --sse-customer-key-base64 value  SSE-C를 사용하는 경우, 데이터를 암호화/복호화하기 위해 Base64로 인코딩된 비밀 암호화 키를 제공해야 합니다. [$SSE_CUSTOMER_KEY_BASE64]
   --sse-customer-key-md5 value     SSE-C를 사용하는 경우, 비밀 암호화 키의 MD5 체크섬을 제공할 수 있습니다(선택 사항). [$SSE_CUSTOMER_KEY_MD5]
   --upload-concurrency value       멀티파트 업로드의 동시성을 입력하세요. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값을 입력하세요. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          싱글 파트 업로드에 서명된 요청 또는 PutObject을 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간대의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   일반

   --name value  스토리지의 이름(기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}