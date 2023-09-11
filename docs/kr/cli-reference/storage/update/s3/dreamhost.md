# Dreamhost DreamObjects

## NAME:
   singularity storage update s3 dreamhost - Dreamhost DreamObjects

## USAGE:
   singularity storage update s3 dreamhost [command options] <name|id>

## DESCRIPTION:

--env-auth
  런타임에서 AWS 자격증명 가져오기 (환경 변수 또는 env vars 또는 EC2/ECS 메타데이터).
  
  access_key_id 및 secret_access_key이 비어 있을 때만 적용됩니다.

  예제:
     | false | 다음 단계에서 AWS 자격증명을 입력하십시오.
     | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격증명 가져오기.

--access-key-id
  AWS 액세스 키 ID.
  
  익명 액세스 또는 런타임 자격증명을 위해 비워 두십시오.

--secret-access-key
  AWS 비밀 액세스 키 (비밀번호).
  
  익명 액세스 또는 런타임 자격증명을 위해 비워 두십시오.

--region
  연결할 리전.
  
  S3 클론을 사용하고 리전이 없는 경우 비워 두십시오.

  예제:
     | <unset>            | 확신이 없는 경우 이렇게 사용하십시오.
     |                    | v4 서명 및 빈 리전을 사용합니다.
     | other-v2-signature | v4 서명이 작동하지 않을 때만 사용하십시오.
     |                    | 예: 이전 버전의 Jewel/v10 CEPH.

--endpoint
  S3 API 엔드포인트.
  
  S3 클론을 사용하는 경우 필수입니다.

  예제:
     | objects-us-east-1.dream.io | Dream Objects 엔드포인트

--location-constraint
  리전과 일치하는 위치 제약 조건.
  
  확실하지 않은 경우 비워 두십시오. 버킷 생성 시에만 사용됩니다.

--acl
  버킷 생성 및 객체 저장 또는 복사 시 사용되는 Canned ACL.
  
  이 ACL은 객체를 생성할 때 사용되며 bucket_acl이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
  
  자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
  
  S3에서 서버 측 복사 객체로 복사될 때 이 ACL이 적용됩니다.
  소스에서 ACL을 복사하는 대신 S3는 새로운 ACL을 작성합니다.
  
  acl이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.

--bucket-acl
  버킷을 생성할 때 사용되는 Canned ACL.
  
  자세한 내용은 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.
  
  버킷을 생성할 때만 적용되는 ACL입니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.
  
  "acl"과 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(비공개)이 사용됩니다.

  예제:
     | private            | 소유자는 FULL_CONTROL을 얻습니다.
     |                    | 다른 사람은 액세스 권한이 없습니다(기본 설정).
     | public-read        | 소유자는 FULL_CONTROL을 얻습니다.
     |                    | AllUsers 그룹은 읽기 액세스를 얻습니다.
     | public-read-write  | 소유자는 FULL_CONTROL을 얻습니다.
     |                    | AllUsers 그룹은 읽기 및 쓰기 액세스를 얻습니다.
     |                    | 일반적으로 버킷에서 이를 허용하는 것은 권장되지 않습니다.
     | authenticated-read | 소유자는 FULL_CONTROL을 얻습니다.
     |                    | 인증된 사용자 그룹은 읽기 액세스를 얻습니다.

--upload-cutoff
  청크 업로드로 전환하는 임계값.
  
  이보다 큰 파일은 chunk_size의 청크로 업로드됩니다.
  최소값은 0이고 최대값은 5 GiB입니다.

--chunk-size
  업로드에 사용할 청크 크기.
  
  upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일의 경우("rclone rcat" 또는 "rclone mount" 또는 google 사진 또는 google 문서에서 업로드한 파일) 이 청크 크기를 사용하여 멀티파트 업로드로 업로드됩니다.
  
  참고: "--s3-upload-concurrency"이 크기의 청크는 전송당 메모리에 버퍼링됩니다.
  
  고속 링크를 통해 큰 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 빨라집니다.
  
  rclone은 알려진 크기의 대형 파일을 업로드할 때 청크 크기를 자동으로 증가시켜 10,000개의 청크 제한을 유지합니다.
  
  알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 chunk 크기는 5 MiB이며 최대 10,000개의 청크가 있습니다. 따라서 기본 설정에서 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 더 큰 파일을 스트림 업로드하려면 chunk_size를 늘려야 합니다.
  
  청크 크기를 증가시키면 "-P" 플래그로 표시되는 진행 상태 통계의 정확성이 감소합니다. rclone은 청크가 AWS SDK에 의해 버퍼로 보낼 때 청크를 전송한 것으로 처리하지만 실제로는 아직 업로드 중일 수 있습니다. 큰 청크 크기는 더 큰 AWS SDK 버퍼와 진행률 보고로부터 더 많은 차이를 발생시킵니다.

--max-upload-parts
  멀티파트 업로드에 사용되는 최대 파트 수.
  
  이 옵션은 멀티파트 업로드를 수행 할 때 사용될 최대 멀티파트 청크 수를 정의합니다.
  
  AWS S3 사양의 10,000 개 청크를 지원하지 않는 서비스에 유용 할 수 있습니다.
  
  Rclone은 알려진 크기의 대형 파일을 업로드 할 때 자동으로 청크 크기를 증가시켜 이러한 청크 수 제한을 유지합니다.

--copy-cutoff
  청크 복사로 전환하는 임계값.
  
  서버 측에서 복사해야 하는 이보다 큰 파일은 이 크기의 청크로 복사됩니다.
  
  최소값은 0이고 최대값은 5 GiB입니다.

--disable-checksum
  개체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
  
  일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만 대용량 파일의 업로드를 시작하는 데 오랜 시간이 걸릴 수 있습니다.

--shared-credentials-file
  공유 자격증명 파일의 경로.
  
  env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다.
  
  이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" env 변수를 찾습니다. env 값이 비어 있으면 기본값으로 현재 사용자의 홈 디렉터리를 사용합니다.
  
      Linux/OSX: "$HOME/.aws/credentials"
      Windows:   "%USERPROFILE%\.aws\credentials"

--profile
  공유 자격증명 파일에서 사용할 프로필.
  
  env_auth = true이면 rclone은 공유 자격증명 파일을 사용할 수 있습니다. 이 변수는 해당 파일에서 사용할 프로필을 제어합니다.
  
  비어 있으면 환경 변수 "AWS_PROFILE" 또는 설정되지 않은 경우 "default"로 기본 설정됩니다.

--session-token
  AWS 세션 토큰.

--upload-concurrency
  멀티파트 업로드의 동시성.
  
  파일 수가 적고 고속 링크를 통해 전송하는데 사용되지 않는 경우 이 값을 늘리면 전송 속도를 높일 수 있습니다.

--force-path-style
  참이면 경로 스타일 액세스를 사용하고 거짓이면 가상 호스팅 스타일 액세스를 사용합니다.
  
  이 값이 true(기본값)인 경우 rclone은 경로 스타일 액세스를 사용하고 false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 AWS S3 설명서 (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
  
  일부 제공업체는 (예: AWS, Aliyun OSS, Netease COS 또는 Tencent COS) 이 값을 false로 설정해야 합니다. rclone은 제공자 설정에 따라 이 작업을 자동으로 수행합니다.

--v2-auth
  V2 인증을 사용할 경우 true로 설정합니다.
  
  false인 경우 rclone은 V4 인증을 사용합니다. 설정된 경우 rclone은 V2 인증을 사용합니다.
  
  이 값을 사용할 경우 v4 서명 작동하지 않을 때만 사용하세요. 예: 이전 버전의 Jewel/v10 CEPH.

--list-chunk
  목록의 청크 크기 (각 ListObject S3 요청에 대한 응답 목록 크기).
  
  이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
  대부분의 서비스는 응답 목록을 1000 개의 개체로 잘라 내도록 제한하지만 요청한 개수보다 많은 응답 목록이 있어도 1000 개로 잘립니다.
  AWS S3에서는 이것이 전역 최대치이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참조하십시오.
  Ceph에서는 "rgw list buckets max chunk" 옵션으로 증가시킬 수 있습니다.

--list-version
  사용할 ListObjects 버전: 1,2 또는 자동(0).
  
  S3가 처음 출시되었을 때 버킷의 객체를 나열하기 위해 ListObjects 호출만 제공했습니다.
  
  그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이것은 훨씬 더 높은 성능을 제공하며 가능하면 사용해야 합니다.
  
  기본값인 0으로 설정되어 있으면 rclone은 제공자 설정에 따라 어떤 list 객체 메서드를 호출할지 추측합니다. 잘못 추측하면 여기에서 수동으로 설정할 수 있습니다.

--list-url-encode
  목록을 URL 인코딩할지 여부: true/false/unset
  
  일부 제공업체는 파일 이름에 특수 문자를 사용할 때 URL 인코딩 목록을 지원합니다. 이를 사용할 수 있으면 파일을 수신할 때 신뢰성이 높아질 수 있습니다. unset으로 설정할 경우 rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의할 수 있습니다.

--no-check-bucket
  버킷을 확인하거나 생성하지 않으려면 설정하세요.
  
  버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
  
  또는 사용자가 버킷 생성 권한을 갖지 않은 경우 필요할 수 있습니다. v1.52.0 이전에는 이는 버그로 인해 정상적으로 전달되지 않았습니다.

--no-head
  업로드 된 개체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다.
  
  rclone은 PUT으로 객체 업로드 후에 200 OK 메시지를 받으면 올바르게 업로드된 것으로 간주합니다.
  
  특히 다음과 같이 가정됩니다:
  
  - Metadata, modtime, 저장 클래스 및 콘텐츠 유형이 업로드한 것과 동일합니다.
  - 크기는 업로드한 것과 동일합니다.
  
  rclone은 단일 부분 PUT의 응답에 대해 다음 항목을 읽습니다:
  
  - MD5SUM
  - 업로드된 날짜
  
  멀티파트 업로드의 경우 이러한 항목들은 읽지 않습니다.
  
  알려지지 않은 길이의 소스 개체가 업로드되면 rclone은 HEAD 요청을 수행합니다.
  
  이 플래그를 설정하면 업로드 실패가 감지되지 않을 가능성이 높아지므로 일반 작업에는 권장되지 않습니다. 실제로 두 번째 업로드 실패 확률은 매우 작습니다.

--no-head-object
  GET을 이용해 오브젝트를 가져오기 전에 HEAD를 사용하지 않습니다.

--encoding
  백엔드에 대한 인코딩입니다.
  
  자세한 내용은 개요의 [encoding section](/overview/#encoding)을 참조하십시오.

--memory-pool-flush-time
  내부 메모리 버퍼 풀이 플러시되는 횟수입니다.
  
  추가 버퍼(예: 멀티파트)가 필요한 업로드에는 메모리 풀을 사용하여 할당됩니다.
  이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 주기를 제어합니다.

--memory-pool-use-mmap
  내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

--disable-http2
  S3 백엔드에 대한 HTTP/2 사용을 비활성화합니다.
  
  현재 s3(특히 minio) 백엔드와 HTTP/2에 문제가 있습니다. s3 백엔드의 HTTP/2는 기본적으로 활성화되어 있지만 이곳에서 비활성화할 수 있습니다. 문제가 해결되면이 플래그는 제거될 것입니다.
  
  참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

--download-url
  다운로드에 대한 사용자 정의 엔드포인트.
  일반적으로 데이터를 CloudFront 네트워크를 통해 다운로드하는 경우 더 저렴한 데이터 운임을 제공하는 AWS S3를 사용할 때 CloudFront CDN URL로 설정됩니다.

--use-multipart-etag
  확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부입니다.
  
  이 값은 true, false 또는 공급자의 기본값을 사용하려면 비워 둡니다.

--use-presigned-request
  단일 부분 업로드에 대해 미리 서명된 요청 또는 PutObject를 사용할지 여부입니다.
  
  이 값이 false인 경우 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.
  
  rclone의 버전 < 1.59은 단일 부분 객체를 업로드하기 위해 미리 서명된 요청을 사용하며이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이는 예외적인 경우나 테스트 용도를 제외하고는 필요하지 않습니다.

--versions
  디렉토리 목록에 이전 버전을 포함합니다.

--version-at
  지정된 시간에 파일 버전을 표시합니다.
  
  매개변수는 날짜 "2006-01-02", 날짜시간 "2006-01-02 15:04:05" 또는 그렇게 오래된 기간 "100d" 또는 "1h"일 수 있습니다.
  
  이 옵션을 사용하는 경우 파일 쓰기 작업은 허용되지 않으므로 파일 업로드나 삭제는 할 수 없습니다.
  
  사용 가능한 형식은 [time option docs](/docs/#time-option)를 참조하십시오.

--decompress
  gzip으로 압축된 개체를 압축 해제합니다.
  
  "Content-Encoding: gzip"로 S3에 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
  
  이 플래그가 설정되면 rclone은 수신 시 "Content-Encoding: gzip"로 이러한 파일을 압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없게하지만 파일 내용은 압축 해제됩니다.

--might-gzip
  백엔드가 개체를 gzip으로 압축 할 수 있다면이 값을 설정하십시오.
  
  일반적으로 공급자는 객체를 다운로드할 때 개체를 수정하지 않습니다. "Content-Encoding: gzip"으로 업로드되지 않은 경우 다운로드되지 않습니다.
  
  그러나 Cloudflare와 같은 일부 공급업체는 "Content-Encoding: gzip"로 업로드되지 않은 개체를 gzip으로 압축 할 수 있습니다.
  
  이 경우 "ERROR corrupted on transfer: sizes differ NNN vs MMM"과 같은 오류를 받게됩니다.
  
  이 플래그를 설정하고 rclone이 Content-Encoding: gzip이 설정된 청크된 전송 인코딩으로 객체를 다운로드하면 rclone은 객체를 차례로 압축 해제합니다.
  
  unset으로 설정된 경우 (기본값) rclone은 공급자 설정에 따라 적용할 내용을 선택하지만 여기에서 rclone의 선택을 재정의 할 수 있습니다.

--no-system-metadata
  시스템 메타데이터 설정 및 읽기를 제한합니다.


## OPTIONS:

--access-key-id value        AWS 액세스 키 ID. [$ACCESS_KEY_ID]
--acl value                  버킷 생성 및 저장 또는 복사 개체에 사용되는 Canned ACL. [$ACL]
--endpoint value             S3 API 엔드포인트. [$ENDPOINT]
--env-auth                   런타임에서 AWS 자격증명 가져오기 (환경 변수 또는 EC2/ECS 메타데이터) (기본값: false) [$ENV_AUTH]
--help, -h                   도움말 표시
--location-constraint value  리전과 일치하는 위치 제약 조건. [$LOCATION_CONSTRAINT]
--region value               연결할 리전. [$REGION]
--secret-access-key value    AWS 비밀 액세스 키 (비밀번호). [$SECRET_ACCESS_KEY]

## Advanced

--bucket-acl value               버킷 생성에 사용되는 Canned ACL. [$BUCKET_ACL]
--chunk-size value               업로드에 사용할 청크 크기 (기본값: "5Mi") [$CHUNK_SIZE]
--copy-cutoff value              청크 복사로 전환하는 임계값 (기본값: "4.656Gi") [$COPY_CUTOFF]
--decompress                     gzip으로 압축된 객체를 압축 해제합니다 (기본값: false) [$DECOMPRESS]
--disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다 (기본값: false) [$DISABLE_CHECKSUM]
--disable-http2                  S3 백엔드에 대한 HTTP/2 사용을 비활성화합니다 (기본값: false) [$DISABLE_HTTP2]
--download-url value             다운로드에 대한 사용자 지정 엔드포인트 [$DOWNLOAD_URL]
--encoding value                 백엔드에 대한 인코딩 (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
--force-path-style               경로 스타일 액세스를 사용할지 여부 (기본값: true) [$FORCE_PATH_STYLE]
--list-chunk value               목록의 청크 크기 (각 ListObject S3 요청에 대한 응답 목록 크기) (기본값: 1000) [$LIST_CHUNK]
--list-url-encode value          목록을 URL 인코딩할지 여부 (기본값: "unset") [$LIST_URL_ENCODE]
--list-version value             사용할 ListObjects 버전: 1,2 또는 자동 (0) (기본값: 0) [$LIST_VERSION]
--max-upload-parts value         멀티파트 업로드에 사용되는 최대 파트 수 (기본값: 10000) [$MAX_UPLOAD_PARTS]
--memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 시간 (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
--memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부 (기본값: false) [$MEMORY_POOL_USE_MMAP]
--might-gzip value               백엔드가 개체를 gzip으로 압축 할 수 있다면이 값을 설정하세요 (기본값: "unset") [$MIGHT_GZIP]
--no-check-bucket                버킷을 확인하거나 생성하지 않습니다 (기본값: false) [$NO_CHECK_BUCKET]
--no-head                        업로드 된 개체의 무결성을 확인하기 위해 HEAD를 사용하지 않습니다 (기본값: false) [$NO_HEAD]
--no-head-object                 GET을 이용해 오브젝트를 가져오기 전에 HEAD를 사용하지 않습니다 (기본값: false) [$NO_HEAD_OBJECT]
--no-system-metadata             시스템 메타데이터 설정 및 읽기를 제한합니다 (기본값: false) [$NO_SYSTEM_METADATA]
--profile value                  공유 자격증명 파일에서 사용할 프로필 [$PROFILE]
--session-token value            AWS 세션 토큰 [$SESSION_TOKEN]
--shared-credentials-file value  공유 자격증명 파일의 경로 [$SHARED_CREDENTIALS_FILE]
--upload-concurrency value       멀티파트 업로드의 동시성 (기본값: 4) [$UPLOAD_CONCURRENCY]
--upload-cutoff value            청크 업로드로 전환하는 임계값 (기본값: "200Mi") [$UPLOAD_CUTOFF]
--use-multipart-etag value       확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
--use-presigned-request          단일 부분 업로드에 대해 미리 서명된 요청 또는 PutObject를 사용할지 여부 (기본값: false) [$USE_PRESIGNED_REQUEST]
--v2-auth                        V2 인증을 사용할 경우 true로 설정합니다 (기본값: false) [$V2_AUTH]
--version-at value               지정된 시간의 파일 버전을 표시합니다 (기본값: "off") [$VERSION_AT]
--versions                       디렉토리 목록에 이전 버전을 포함합니다 (기본값: false) [$VERSIONS]