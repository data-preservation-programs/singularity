# RackCorp Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage create s3 rackcorp - RackCorp Object Storage

사용법:
   singularity storage create s3 rackcorp [command options] [arguments...]

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).
      
      access_key_id와 secret_access_key가 비어 있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 IAM).

   --access-key-id
      AWS Access Key ID.
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워둡니다.

   --secret-access-key
      AWS Secret Access Key (비밀번호).
      
      익명 액세스 또는 런타임 자격 증명을 위해 비워둡니다.

   --region
      지역 - 버킷이 생성되고 데이터가 저장될 위치입니다.
      

      예시:
         | global    | 글로벌 CDN (모든 위치) 지역
         | au        | 호주 (모든 주)
         | au-nsw    | 뉴 사우스 웨일즈 (호주) 지역
         | au-qld    | 퀸즐랜드 (호주) 지역
         | au-vic    | 빅토리아 (호주) 지역
         | au-wa     | 퍼스 (호주) 지역
         | ph        | 마닐라 (필리핀) 지역
         | th        | 방콕 (태국) 지역
         | hk        | 홍콩 (홍콩) 지역
         | mn        | 울란바토르 (몽골) 지역
         | kg        | 비슈케크 (키르기스스탄) 지역
         | id        | 자카르타 (인도네시아) 지역
         | jp        | 도쿄 (일본) 지역
         | sg        | 싱가포르 (싱가포르) 지역
         | de        | 프랑크푸르트 (독일) 지역
         | us        | 미국 (Anycast) 지역
         | us-east-1 | 뉴욕 (미국) 지역
         | us-west-1 | 프리몬트 (미국) 지역
         | nz        | 오클랜드 (뉴질랜드) 지역

   --endpoint
      RackCorp Object Storage의 엔드포인트입니다.

      예시:
         | s3.rackcorp.com           | 글로벌 (Anycast) 엔드포인트
         | au.s3.rackcorp.com        | 호주 (Anycast) 엔드포인트
         | au-nsw.s3.rackcorp.com    | 시드니 (호주) 엔드포인트
         | au-qld.s3.rackcorp.com    | 브리즈번 (호주) 엔드포인트
         | au-vic.s3.rackcorp.com    | 멜버른 (호주) 엔드포인트
         | au-wa.s3.rackcorp.com     | 퍼스 (호주) 엔드포인트
         | ph.s3.rackcorp.com        | 마닐라 (필리핀) 엔드포인트
         | th.s3.rackcorp.com        | 방콕 (태국) 엔드포인트
         | hk.s3.rackcorp.com        | 홍콩 (홍콩) 엔드포인트
         | mn.s3.rackcorp.com        | 울란바토르 (몽골) 엔드포인트
         | kg.s3.rackcorp.com        | 비슈케크 (키르기스스탄) 엔드포인트
         | id.s3.rackcorp.com        | 자카르타 (인도네시아) 엔드포인트
         | jp.s3.rackcorp.com        | 도쿄 (일본) 엔드포인트
         | sg.s3.rackcorp.com        | 싱가포르 (싱가포르) 엔드포인트
         | de.s3.rackcorp.com        | 프랑크푸르트 (독일) 엔드포인트
         | us.s3.rackcorp.com        | 미국 (Anycast) 엔드포인트
         | us-east-1.s3.rackcorp.com | 뉴욕 (미국) 엔드포인트
         | us-west-1.s3.rackcorp.com | 프리몬트 (미국) 엔드포인트
         | nz.s3.rackcorp.com        | 오클랜드 (뉴질랜드) 엔드포인트

   --location-constraint
      버킷이 위치하고 데이터가 저장될 위치입니다.
      

      예시:
         | global    | 글로벌 CDN 지역
         | au        | 호주 (모든 위치)
         | au-nsw    | 뉴 사우스 웨일즈 (호주) 지역
         | au-qld    | 퀸즐랜드 (호주) 지역
         | au-vic    | 빅토리아 (호주) 지역
         | au-wa     | 퍼스 (호주) 지역
         | ph        | 마닐라 (필리핀) 지역
         | th        | 방콕 (태국) 지역
         | hk        | 홍콩 (홍콩) 지역
         | mn        | 울란바토르 (몽골) 지역
         | kg        | 비슈케크 (키르기스스탄) 지역
         | id        | 자카르타 (인도네시아) 지역
         | jp        | 도쿄 (일본) 지역
         | sg        | 싱가포르 (싱가포르) 지역
         | de        | 프랑크푸르트 (독일) 지역
         | us        | 미국 (Anycast) 지역
         | us-east-1 | 뉴욕 (미국) 지역
         | us-west-1 | 프리몬트 (미국) 지역
         | nz        | 오클랜드 (뉴질랜드) 지역

   --acl
      버킷 및 객체 생성 또는 복사 시 사용되는 Canned ACL입니다.
      
      이 ACL은 객체 생성에 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 정보는 다음을 참조하십시오. [AWS S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)

      비어 있는 문자열로 설정할 경우 X-Amz-Acl: 헤더가 추가되지 않고
      기본값 (private)이 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 Canned ACL입니다.
      
      자세한 정보는 다음을 참조하십시오. [AWS S3 ACL 개요](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)
      
      이 ACL은 버킷 생성 시에만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
      
      "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고
      기본값 (private)이 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | 기타 사용자는 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 액세스 권한이 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AllUsers 그룹에게 READ 및 WRITE 액세스 권한이 부여됩니다.
         |                    | 버킷에 이러한 액세스 권한을 부여하는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL이 부여됩니다.
         |                    | AuthenticatedUsers 그룹에게 READ 액세스 권한이 부여됩니다.

   --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.
      
      이 임계값보다 큰 크기의 파일은 chunk_size로 잘라서 업로드됩니다.
      최소값은 0이며 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기입니다.
      
      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일 (예: "rclone rcat" 또는 "rclone mount" 또는 Google
      사진 또는 Google 문서에서 업로드한 파일)은 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      주의할 점은 "--s3-upload-concurrency" 단위 개수의 이러한 크기의 청크가 전송당 메모리에 버퍼링됩니다.
      
      고속 링크에서 대용량 파일을 전송하고 메모리가 충분한 경우 크기를 늘리면 전송 속도가 향상됩니다.
      
      rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 준수합니다.
      
      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk_size는 5 MiB이며 최대 10,000개의
      청크가 있을 수 있으므로 기본적으로 스트림 업로드 할 수 있는 파일의 최대 크기는 48 GiB입니다.
      용량이 더 큰 파일을 스트림 업로드하려면 청크 크기를 더 크게 설정해야 합니다.
      
      청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태 통계의 정확도가 감소합니다. rclone은 청크가 AWS SDK에
      의해 버퍼링 될 때 청크를 보낸 것으로 처리하지만 실제로는 업로드 중일 수 있습니다.
      더 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행 상태 보고에서 진실에서 벗어납니다.
      

   --max-upload-parts
      멀티파트 업로드의 최대 청크 수입니다.
      
      이 옵션은 멀티파트 업로드를 수행할 때 사용할 multipart 청크의 최대 수를 정의합니다.
      
      이는 서비스가 10,000개의 청크를 지원하지 않는 경우 유용할 수 있습니다.
      
      rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 초과하지 않도록 합니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 임계값입니다.
      
      서버 사이드 복사가 필요한 크기보다 큰 파일은 이 크기로 청크를 복사합니다.
      
      최소값은 0이며 최대값은 5 GiB입니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 이는 
      데이터 무결성 확인에 유용하지만 대용량 파일의 업로드가 시작될 때 소요 시간이 길어질 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수가 비어 있으면
      현재 사용자의 홈 디렉터리로 기본값으로 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이
      변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비어 있으면 환경 변수 "AWS_PROFILE" 또는 환경 변수가 설정되지 않은 경우 "default"로 기본값으로 설정됩니다.
      

   --session-token
      AWS 세션 토큰입니다.

   --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.
      
      파일의 동일한 청크가 동시에 업로드됩니다.
      
      대용량 파일을 고속 링크로 전송하고 이러한 업로드가 대역폭을 완전히 활용하지 않는 경우 이를 늘리면 전송 속도가 향상될 수 있습니다.

   --force-path-style
      true이면 경로 스타일 액세스, false이면 가상 호스팅 스타일 액세스를 사용합니다.
      
      이 값이 true (기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고
      false이면 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3
      문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)
     를 참조하십시오.
      
      일부 공급업체 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS)에서는 이 값을
      false로 설정해야 합니다 - rclone은 이를 프로바이더 설정을 기반으로 자동으로 수행합니다.

   --v2-auth
      true이면 v2 인증을 사용합니다.
      
      이 값이 false (기본값)인 경우 rclone은 v4 인증을 사용합니다. 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하십시오. 예: Jewel/v10 CEPH 이전.

   --list-chunk
      목록 청크의 크기 (각 ListObject S3 요청에 대한 응답 목록)입니다.
      
      이 옵션은 AWS S3 명세의 "MaxKeys", "max-items" 또는 "page-size"와 같이 알려져 있습니다.
      대부분의 서비스는 요청된 것보다 더 많은 목록을 요청해도 1000개의 객체로 자른다.
      AWS S3에서는 전역 최대값이고 변경할 수 없으므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이를 증가시킬 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 자동으로 0.
      
      S3가 처음 출시될 때에는 버킷의 객체를 열거하기 위해 ListObjects 호출만 제공했습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은
      훨씬 높은 성능을 제공하므로 가능하면 사용해야 합니다.
      
      기본값 0으로 설정하면 rclone은 프로바이더로 설정된대로 List Objects 메서드를 호출할 것입니다.
      잘못 추측하면 여기서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할지 여부: true/false/unset
      
      일부 공급업체는 목록을 URL 인코딩하는 것을 지원하고 가능한 경우 파일 이름에 제어 문자를 사용할 때 이는 더 신뢰할 수 있습니다. "unset"으로 설정된 경우 (기본값) rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기서 rclone의 선택을 재정의할 수 있습니다.
      

   --no-check-bucket
      버킷이 존재하는지 확인하거나 생성하지 않으려면 설정하세요.
      
      버킷이 이미 존재하는 경우 rclone의 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      사용자에게 버킷 생성 권한이 없는 경우 필요할 수도 있습니다. v1.52.0 이전에는 버그 때문에 무시되었을 것입니다.
      

   --no-head
      업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다.
      
      rclone은 각 PUT을 통해 객체를 업로드한 후 200 OK 메시지를 받으면 제대로 업로드되었다고 가정합니다.
      
      특히 다음을 가정합니다:
      
      - 메타데이터, 수정 시간, 저장 클래스 및 콘텐츠 유형은 업로드와 동일합니다.
      - 크기는 업로드와 동일합니다.
      
      단일 파트 PUT의 응답에서 다음 항목을 읽습니다.
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목들은 읽지 않습니다.
      
      알려지지 않은 길이의 소스 객체를 업로드하면 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 무결성 검사 횟수가 증가하므로 정상적인 작동에는 권장되지 않습니다. 실제로 무결성 검사 실패 가능성은 매우 작습니다.
      

   --no-head-object
      객체를 가져오기 전에 HEAD를 수행하지 않습니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 시간입니다.
      
      추가 버퍼가 필요한 업로드 (예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 시간을 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2 사용을 비활성화합니다.
      
      현재 s3 (특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. HTTP/2는 s3 백엔드의 기본값이지만 여기에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      
      

   --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드 할 경우 더 싼 egress를 제공합니다.

   --use-multipart-etag
      검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부
      
      이 값은 true, false 또는 기본 프로바이더에 따라 설정되지 않은 상태여야 합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부
      
      false인 경우 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용할 것입니다.
      
      rclone < 1.59 버전에서는 단일 파트 객체를 업로드하기 위해 사전 서명된 요청을 사용하고이 플래그를 true로 설정하면
      그 기능이 다시 사용됩니다. 이는 예외적인 경우나 테스트에만 필요합니다.
      

   --versions
      디렉토리 목록에 이전 버전을 포함합니다.

   --version-at
      지정된 시간에 파일 버전을 표시합니다.
      
      매개변수는 날짜, "2006-01-02", 시간 "2006-01-02 15:04:05" 또는 그때까지의 기간, 예: "100d" 또는 "1h"입니다.
      
      이를 사용할 때 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
      
      유효한 형식에 대해서는 [time 옵션 문서](/docs/#time-option)를 참조하십시오.
      

   --decompress
      gzip 인코딩된 객체를 압축 해제합니다.
      
      S3에 "Content-Encoding: gzip"로 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을
      압축된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 "Content-Encoding: gzip" 파일을 받는 대로 압축을 해제합니다.
      이는 rclone이 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하세요.
      
      일반적으로 제공자는 객체를 다운로드할 때 객체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지
      않은 경우 다운로드하지 않습니다.
      
      그러나 제공자 중 일부는 객체를 업로드시에 "Content-Encoding: gzip"로 업로드하지 않았더라도 gzip으로
      압축할 수 있습니다 (예: Cloudflare).
      
      이것의 증상은 다음과 같은 오류가 발생합니다.
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
      이 플래그를 설정하고 rclone이 "Content-Encoding: gzip"가 설정된 상태에서 청크 전송 인코딩으로 객체를
      다운로드할 때 rclone은 객체를 압축 해제합니다.
      
      unset로 설정된 경우 (기본값) rclone은 제공자 설정에 따라 적용할 내용을 선택하지만 여기서 rclone의 선택을
      재정의할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


OPTIONS:
   --access-key-id value        AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                  버킷과 객체를 생성하거나 복사할 때 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value             RackCorp Object Storage의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  버킷이 위치하고 데이터가 저장될 위치입니다. [$LOCATION_CONSTRAINT]
   --region value               버킷을 생성하고 데이터를 저장할 위치입니다. [$REGION]
   --secret-access-key value    AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

   고급

   --bucket-acl value               버킷을 생성할 때 사용되는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              청크로 복사 전환하는 기준값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip 인코딩된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true이면 경로 스타일 액세스를 사용합니다. false이면 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크의 크기입니다. (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할지 여부입니다. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1, 2 또는 자동으로 0. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드의 최대 청크 개수입니다. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간입니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축할 수 있는 경우 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷이 존재하는지 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 객체를 HEAD하여 무결성을 확인하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 객체를 가져오기 전에 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰입니다. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성입니다. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값입니다. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       검증을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 파트 업로드에 사전 서명된 요청 또는 PutObject을 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간에 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

   General

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}