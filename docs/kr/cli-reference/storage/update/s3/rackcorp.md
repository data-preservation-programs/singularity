# RackCorp Object Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 rackcorp - RackCorp Object Storage

사용법:
   singularity storage update s3 rackcorp [명령 옵션] <이름|아이디>

설명:
   --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터).

      액세스 키 ID와 비밀 액세스 키가 비어있는 경우에만 적용됩니다.

      예시:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경에서 AWS 자격 증명을 가져옵니다(환경 변수 또는 IAM).

   --access-key-id
      AWS 액세스 키 ID.

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워두세요.

   --secret-access-key
      AWS 비밀 액세스 키(비밀번호).

      익명 액세스 또는 런타임 자격 증명을 사용하려면 비워두세요.

   --region
      지역 - 버킷 및 데이터가 저장될 위치입니다.

      예시:
         | global    | 글로벌 CDN(모든 위치) 지역
         | au        | 호주 (모든 주)
         | au-nsw    | 뉴사우스웨일즈(호주) 지역
         | au-qld    | 퀸즈랜드(호주) 지역
         | au-vic    | 빅토리아(호주) 지역
         | au-wa     | 퍼스(호주) 지역
         | ph        | 마닐라(필리핀) 지역
         | th        | 방콕(태국) 지역
         | hk        | 홍콩 지역
         | mn        | 울란바토르(몽골) 지역
         | kg        | 비쉬케크(키르기스스탄) 지역
         | id        | 자카르타(인도네시아) 지역
         | jp        | 도쿄(일본) 지역
         | sg        | 싱가포르 지역
         | de        | 프랑크푸르트(독일) 지역
         | us        | 미국 (AnyCast) 지역
         | us-east-1 | 뉴욕(미국) 지역
         | us-west-1 | 프리몬트(미국) 지역
         | nz        | 오클랜드(뉴질랜드) 지역

   --endpoint
      RackCorp Object Storage의 엔드포인트입니다.

      예시:
         | s3.rackcorp.com           | 글로벌 (AnyCast) 엔드포인트
         | au.s3.rackcorp.com        | 호주 (Anycast) 엔드포인트
         | au-nsw.s3.rackcorp.com    | 시드니(호주) 엔드포인트
         | au-qld.s3.rackcorp.com    | 브리즈번(호주) 엔드포인트
         | au-vic.s3.rackcorp.com    | 멜버른(호주) 엔드포인트
         | au-wa.s3.rackcorp.com     | 퍼스(호주) 엔드포인트
         | ph.s3.rackcorp.com        | 마닐라(필리핀) 엔드포인트
         | th.s3.rackcorp.com        | 방콕(태국) 엔드포인트
         | hk.s3.rackcorp.com        | 홍콩 엔드포인트
         | mn.s3.rackcorp.com        | 울란바토르(몽골) 엔드포인트
         | kg.s3.rackcorp.com        | 비쉬케크(키르기스스탄) 엔드포인트
         | id.s3.rackcorp.com        | 자카르타(인도네시아) 엔드포인트
         | jp.s3.rackcorp.com        | 도쿄(일본) 엔드포인트
         | sg.s3.rackcorp.com        | 싱가포르 엔드포인트
         | de.s3.rackcorp.com        | 프랑크푸르트(독일) 엔드포인트
         | us.s3.rackcorp.com        | 미국 (AnyCast) 엔드포인트
         | us-east-1.s3.rackcorp.com | 뉴욕(미국) 엔드포인트
         | us-west-1.s3.rackcorp.com | 프리몬트(미국) 엔드포인트
         | nz.s3.rackcorp.com        | 오클랜드(뉴질랜드) 엔드포인트

   --location-constraint
      버킷이 위치하고 데이터가 저장될 위치입니다.

      예시:
         | global    | 글로벌 CDN 지역
         | au        | 호주 (모든 위치)
         | au-nsw    | 뉴사우스웨일즈(호주) 지역
         | au-qld    | 퀸즈랜드(호주) 지역
         | au-vic    | 빅토리아(호주) 지역
         | au-wa     | 퍼스(호주) 지역
         | ph        | 마닐라(필리핀) 지역
         | th        | 방콕(태국) 지역
         | hk        | 홍콩 지역
         | mn        | 울란바토르(몽골) 지역
         | kg        | 비쉬케크(키르기스스탄) 지역
         | id        | 자카르타(인도네시아) 지역
         | jp        | 도쿄(일본) 지역
         | sg        | 싱가포르 지역
         | de        | 프랑크푸르트(독일) 지역
         | us        | 미국 (AnyCast) 지역
         | us-east-1 | 뉴욕(미국) 지역
         | us-west-1 | 프리몬트(미국) 지역
         | nz        | 오클랜드(뉴질랜드) 지역

   --acl
      개체를 만들고 저장하거나 복사할 때 사용되는 고정된 ACL.

      이 ACL은 개체 생성 및 버킷_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
      
      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
      
      S3 서버 간 개체 복사시에만 적용되므로 대상으로부터 ACL을 복사하지 않습니다.
      
      ACL이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않습니다. 기본값인 private가 사용됩니다.
      

   --bucket-acl
      버킷 생성 시 사용되는 고정된 ACL.

      자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl를 참조하십시오.
      
      이 ACL은 버킷 생성에만 적용됩니다. 설정되지 않은 경우 "acl"이 사용됩니다.
      
      "acl"와 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않으며 기본값인 private가 사용됩니다.
      

      예시:
         | private            | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | 다른 사용자는 액세스 권한이 없습니다(기본값).
         | public-read        | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AllUsers 그룹은 READ 권한이 부여됩니다.
         | public-read-write  | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AllUsers 그룹은 READ 및 WRITE 권한이 부여됩니다.
         |                    | 버킷에 이 권한이 부여되는 것은 일반적으로 권장되지 않습니다.
         | authenticated-read | 소유자에게 FULL_CONTROL 권한이 부여됩니다.
         |                    | AuthenticatedUsers 그룹은 READ 권한이 부여됩니다.

   --upload-cutoff
      청크 업로드로 전환하는 경계 값.

      이보다 큰 파일은 chunk_size 단위로 청크 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

   --chunk-size
      업로드에 사용할 청크 크기.

      upload_cutoff보다 큰 파일이나 크기를 알 수없는 파일("rclone rcat"에서 또는 "rclone mount" 또는 google 사진 또는 google 문서에서 업로드 된 파일)을 업로드 할 때, 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
      
      주의할 점은 "--s3-upload-concurrency" 이 크기의 청크가 전송마다 메모리에 버퍼링된다는 것입니다.
      
      높은 속도의 링크로 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 값을 높이면 전송 속도가 향상될 수 있습니다.
      
      rclone은 10,000개의 청크 제한을 초과하지 않도록 큰 파일의 경우 자동으로 청크 크기를 증가시킵니다.
      
      크기를 알 수없는 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 개의 청크가 있을 수 있으므로, 기본적으로 스트림 업로드 할 수있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 청크 크기를 증가해야합니다.
      
      청크 크기를 증가시키면 "-P" 플래그와 함께 표시되는 진행 상태 통계의 정확성이 감소합니다. rclone은 AWS SDK에 의해 버퍼링 된 청크가 전송된 것으로 처리하지만 실제로는 여전히 업로드 중인 경우에도 청크를 전송으로 처리합니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼 및 진행률 보고의 신뢰도와 편차가 생깁니다.
      

   --max-upload-parts
      멀티파트 업로드에서 사용할 최대 청크 수.

      이 옵션은 멀티파트 업로드를 수행 할 때 사용할 최대 멀티파트 청크 수를 정의합니다.
      
      서비스가 10,000개의 청크 지정 사양을 지원하지 않는 경우 유용할 수 있습니다.
      
      rclone은 크기를 알 수 있는 큰 파일을 업로드 할 때 이 번호가 청크 수 제한보다 작은 상태로 유지하기 위해 자동으로 청크 크기를 증가시킵니다.
      

   --copy-cutoff
      멀티파트 복사로 전환하는 파일의 클릭 밸리 예금.

      이보다 큰 파일은 이 크기의 청크로 복사됩니다.
      
      최소값은 0이고 최대값은 5 GiB입니다.

   --disable-checksum
      개체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      보통 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 좋지만 큰 파일의 경우 업로드를 시작하기까지 시간이 오래 걸릴 수 있습니다.

   --shared-credentials-file
      공유 자격 증명 파일의 경로.

      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
      
      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어있는 경우 현재 사용자의 홈 디렉토리로 기본 설정됩니다.
      
          Linux/OSX: "$HOME/.aws/credentials"
          Windows:   "%USERPROFILE%\.aws\credentials"
      

   --profile
      공유 자격 증명 파일에서 사용할 프로필.
      
      env_auth = true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
      
      비워두면 환경 변수 "AWS_PROFILE" 또는 환경 변수가 설정되지 않은 경우 "default"로 기본 설정됩니다.
      

   --session-token
      AWS 세션 토큰.

   --upload-concurrency
      멀티파트 업로드의 동시성.

      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      높은 속도의 링크를 통해 대량의 큰 파일을 전송하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우 이를 증가시키는 것이 전송 속도를 높이는 데 도움이 될 수 있습니다.

   --force-path-style
      True이면 경로 스타일 액세스, False이면 가상 호스팅 스타일 액세스를 사용합니다.

      true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고,
      false인 경우 rclone은 가상 호스팅 스타일을 사용합니다. 자세한 내용은[the AWS S3 docs](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      일부 공급자(AWS, Aliyun OSS, Netease COS 또는 Tencent COS 등)는 false로 설정해야합니다. rclone은 공급자 설정에 따라 자동으로 수행합니다.

   --v2-auth
      True이면 v2 인증을 사용합니다.

      false(기본값)인 경우 rclone은 v4 인증을 사용하고, 설정된 경우 rclone은 v2 인증을 사용합니다.
      
      v4 서명이 작동하지 않는 경우에만 사용하십시오. 예를 들어 Jewel/v10 이전의 CEPH의 경우입니다.

   --list-chunk
      목록 청크 크기(ListObject S3 요청마다 응답 목록)입니다.

      이 옵션은 AWS S3 사양의 "MaxKeys" 또는 "max-items" 또는 "page-size"로 알려진 것과 동일합니다.
      대부분의 서비스는 요청 수를 초과하더라도 응답 목록을 1000 개로 잘라냅니다. AWS S3에서 이 최대값은 전역적이며 변경할 수 없으므로 [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 높일 수 있습니다.
      

   --list-version
      사용할 ListObjects 버전: 1, 2 또는 0 (자동).

      S3가 처음 출시되었을 때 버킷의 개체를 나열하기 위해 ListObjects 호출만 제공하였습니다.
      
      그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하여 가능한 경우 사용해야합니다.
      
      기본값 0으로 설정하면 rclone은 설정된 공급자에 따라 어떤 목록 개체 메서드를 호출할지 추측합니다. 추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.
      

   --list-url-encode
      목록을 URL 인코딩할 지 여부: true/false/unset
      
      일부 제공자는 URL 인코딩 목록을 지원하며 사용 가능한 경우 파일 이름에 컨트롤 문자를 사용할 때이 방법이 더 안정적입니다. unset으로 설정된 경우 rclone은 공급자 설정에 따라 어떤 방법을 적용할지 선택하지만 여기에서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-check-bucket
      버킷의 존재를 확인하거나 버킷을 생성하지 않으려면 설정하세요.

      알려진 버킷이 이미 존재하는 경우 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      
      버킷 생성 권한이 없는 경우도 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 조용히 전달되었습니다.
      

   --no-head
      업로드된 객체를 HEAD하여 무결성을 확인하지 마십시오.

      rclone은 보통 "Content-Encoding: gzip"로 업로드 된 객체를 압축 해제하여 다운로드합니다. 나중에 압축이 풀린 파일 내용을 사용하지만 사이즈와 해시를 확인할 수 없습니다.
      
      이 플래그가 설정되면 rclone은 PUT 이후 200 OK 메시지를 수신하면 제대로 업로드 된 것으로 간주합니다.
      
      특히 다음과 같은 것을 가정합니다:
      
      - 메타데이터(모디파이, 저장소 클래스 및 콘텐츠 유형)가 업로드한 대로임
      - 크기가 업로드한 대로임
      
      또한 rclone은 다음과 같은 반환 응답에서 단일 파트 PUT에 대해 읽습니다:
      
      - MD5SUM
      - 업로드된 날짜
      
      멀티파트 업로드의 경우 이러한 항목은 읽지 않습니다.
      
      길이를 알 수없는 소스 개체를 업로드하는 경우 rclone은 HEAD 요청을 수행합니다.
      
      이 플래그를 설정하면 올바르지 않은 크기와 같은 감지되지 않은 업로드 오류의 가능성이 증가하므로 일반 작업에 권장되지 않습니다. 실제로 감지되지 않은 업로드 오류의 가능성은 매우 적습니다.
      

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 마십시오.

   --encoding
      백엔드의 인코딩입니다.

      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀을 어느 정도의 간격으로 플러시 할 것인지 제어합니다.

      추가 버퍼를 필요로하는 업로드(예: 멀티파트)는 메모리 풀을 사용하여 할당을 수행합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 자주를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --disable-http2
      S3 백엔드에서 http2의 사용을 비활성화합니다.

      현재 s3 (특히 minio) 백엔드와 HTTP/2에 해결되지 않은 문제가 있습니다. 기본적으로 s3 백엔드에서는 HTTP/2가 활성화되지만 이곳에서 비활성화 할 수 있습니다. 문제가 해결되면이 플래그가 제거될 것입니다.
      
      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631
      

   --download-url
      다운로드에 대한 사용자 정의 엔드포인트입니다.
      이는 일반적으로 데이터를 CloudFront 네트워크를 통해 다운로드하는 경우 AWS S3가
      CloudFront 네트워크를 통해 데이터를 다운로드하는 데보다 저렴한 데이터 egress를 제공하기 때문에
      CloudFront CDN URL로 설정됩니다.

   --use-multipart-etag
      멀티파트 업로드에서 ETag를 사용하여 검증할지 여부

      true, false 또는 unset이어야합니다.
      

   --use-presigned-request
      단일 파트 업로드에 사전 서명 된 요청 또는 PutObject을 사용할지 여부

      이 플래그가 false로 설정되면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
      
      rclone 버전 < 1.59는 단일 파트 객체를 업로드하기 위해 사전 서명 된 요청을 사용하고이 기능을 다시 활성화하기위한 flag를 true로 설정할 것입니다. 이 flag는 예외적인 상황이나 테스트를 제외한 경우에는 필요하지 않습니다.
      

   --versions
      디렉토리 목록에 이전 버전 포함

   --version-at
      지정된 시간에있었던 파일 버전을 표시합니다.

      매개 변수는 날짜 "2006-01-02", datetime "2006-01-02 15:04:05" 또는 그 정도로 설정된 지 오래된 기간인 "100d" 또는 "1h" 일 수 있습니다.
      
      이를 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제 할 수 없습니다.
      
      유효한 형식에 대한 자세한 내용은 [타임 옵션 도움말](/docs/#time-option)을 참조하십시오.
      

   --decompress
      설정하면 gzip으로 압축 된 개체를 압축 해제합니다.

      S3로 "Content-Encoding: gzip"가 설정되어있는 상태에서 개체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축 된 개체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 "Content-Encoding: gzip"로 수신시 압축 해제합니다. 따라서 rclone은 크기와 해시를 확인할 수 없지만 파일의 내용은 압축 해제됩니다.
      

   --might-gzip
      백엔드에서 압축을 수행할 수 있으므로 설정하세요.

      일반적으로 제공자는 개체를 다운로드 할 때 개체를 수정하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 개체는 다운로드될 때 설정되지 않습니다.
      
      그러나 일부 제공자는 `Content-Encoding: gzip`로 업로드되지 않은 개체를 gzip으로 압축 할 수 있습니다(Cloudflare 예).
      
      이를 확인하려면
      
          ERROR corrupted on transfer: sizes differ NNN vs MMM
      
     와 같은 오류가 발생합니다.
      
      이 플래그를 설정하고 rclone이 Content-Encoding: gzip이 설정되고 청크 전송 인코딩으로 개체를 다운로드하는 경우 rclone은 개체를 실시간으로 압축 풉니다.
      
      unset으로 설정하면 rclone은 공급자 설정에 따라 어느 것을 적용할지 선택하지만 여기에서 rclone의 선택을 재정의 할 수 있습니다.
      

   --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제합니다.


옵션:
   --access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
   --acl value                  개체 저장 또는 복사 시 사용되는 고정된 ACL. [$ACL]
   --endpoint value             RackCorp Object Storage의 엔드포인트. [$ENDPOINT]
   --env-auth                   런타임에서 AWS 자격 증명을 가져옵니다(환경 변수 또는 환경 변수가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                   도움말 표시
   --location-constraint value  버킷이 위치하고 데이터가 저장될 위치입니다. [$LOCATION_CONSTRAINT]
   --region value               지역 - 버킷 및 데이터가 저장될 위치입니다. [$REGION]
   --secret-access-key value    AWS 비밀 액세스 키(비밀번호). [$SECRET_ACCESS_KEY]

   Advanced

   --bucket-acl value               버킷 생성 시 사용되는 고정된 ACL. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환하는 파일의 크릭 만료. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     설정하면 gzip으로 압축 된 개체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               개체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에서 http2의 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드에 대한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               True이면 경로 스타일 액세스를 사용하고 false이면 가상 호스팅 스타일 액세스를 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기(ListObject S3 요청마다 응답 목록). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          목록을 URL 인코딩할 지 여부: true/false/unset (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0 (자동). (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에서 사용할 최대 청크 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀을 어느 정도의 간격으로 플러시 할 것인지 조정합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드에서 압축을 수행할 수 있으므로 설정하세요. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 버킷을 생성하지 않으려면 설정하세요. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드된 개체를 HEAD하여 무결성을 확인하지 마십시오. (기본값: false) [$NO_HEAD]
   --no-head-object                 개체를 가져올 때 GET 전에 HEAD를 수행하지 마십시오. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 크기 제한. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       멀티파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          사전 서명 된 요청 또는 PutObject를 사용하여 단일 파트 업로드 할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        True이면 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시간의 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉토리 목록에 이전 버전 포함 (기본값: false) [$VERSIONS]

```
{% endcode %}