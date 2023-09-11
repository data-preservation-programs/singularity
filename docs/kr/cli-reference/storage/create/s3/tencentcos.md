# Tencent Cloud Object Storage (COS)

## 소개
Tencent Cloud Object Storage (COS)는 Scimgularity와 통합되어 쉽게 사용할 수 있는 객체 저장소입니다. 이 문서에서는 S3 프로토콜을 사용하여 COS와의 연결을 설정하는 방법에 대해 자세히 설명합니다.

## 사용법

```
singularity storage create s3 tencentcos [command options] [arguments...]
```

## 설명

### --env-auth

- 런타임에서 AWS 인증 정보(환경 변수 또는 엔드포인트의 EC2/ECS 메타데이터)를 가져옵니다.
- `access_key_id`와 `secret_access_key`가 비어 있을 경우에만 적용됩니다.
- 예시:
  - `false`: 다음 단계에서 AWS 인증 정보를 입력하세요.
  - `true`: 환경(환경 변수 또는 IAM)에서 AWS 인증 정보를 가져옵니다.

### --access-key-id

- AWS 액세스 키 ID를 입력하세요.
- 익명 액세스 또는 런타임 인증 정보를 사용하려면 비워 두세요.

### --secret-access-key

- AWS 시크릿 액세스 키(비밀번호)를 입력하세요.
- 익명 액세스 또는 런타임 인증 정보를 사용하려면 비워 두세요.

### --endpoint

- Tencent COS API의 엔드포인트를 입력하세요.
- 예시:
  - `cos.ap-beijing.myqcloud.com`: 베이징 리전
  - `cos.ap-nanjing.myqcloud.com`: 난징 리전
  - `cos.ap-shanghai.myqcloud.com`: 상하이 리전
  - `cos.ap-guangzhou.myqcloud.com`: 광저우 리전
  - `cos.ap-nanjing.myqcloud.com`: 난징 리전
  - `cos.ap-chengdu.myqcloud.com`: 청두 리전
  - `cos.ap-chongqing.myqcloud.com`: 초칭 리전
  - `cos.ap-hongkong.myqcloud.com`: 홍콩(중국) 리전
  - `cos.ap-singapore.myqcloud.com`: 싱가포르 리전
  - `cos.ap-mumbai.myqcloud.com`: 뭄바이 리전
  - `cos.ap-seoul.myqcloud.com`: 서울 리전
  - `cos.ap-bangkok.myqcloud.com`: 방콕 리전
  - `cos.ap-tokyo.myqcloud.com`: 도쿄 리전
  - `cos.na-siliconvalley.myqcloud.com`: 실리콘밸리 리전
  - `cos.na-ashburn.myqcloud.com`: 버지니아 리전
  - `cos.na-toronto.myqcloud.com`: 토론토 리전
  - `cos.eu-frankfurt.myqcloud.com`: 프랑크푸르트 리전
  - `cos.eu-moscow.myqcloud.com`: 모스크바 리전
  - `cos.accelerate.myqcloud.com`: 텐센트 COS 가속 엔드포인트 사용

### --acl

- 버킷 및 객체 생성 시 사용되는 Canned ACL을 입력하세요.
- 객체 생성에 사용되며 `bucket_acl`이 설정되지 않은 경우 버킷 생성에도 사용됩니다.
- 자세한 내용은 [AWS S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참고하세요.
- 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.

### --bucket-acl

- 버킷 생성 시 사용되는 Canned ACL을 입력하세요.
- 자세한 내용은 [AWS S3 공식 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl)를 참고하세요.
- `acl` 및 `bucket_acl`이 빈 문자열이 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(개인)이 사용됩니다.

### --storage-class

- Tencent COS에 새로운 객체를 저장할 때 사용할 스토리지 클래스를 입력하세요.
- 예시:
  - `<unset>`: 기본값
  - `STANDARD`: 표준 스토리지 클래스
  - `ARCHIVE`: 아카이브 스토리지 모드
  - `STANDARD_IA`: 빈번한 액세스 스토리지 모드

### --upload-cutoff

- 청크 업로드로 전환하는 파일의 크기 기준을 입력하세요.
- 이 값보다 큰 파일은 청크 크기로 업로드됩니다.
- 최소 값은 0이고 최대 값은 5 GiB입니다.

### --chunk-size

- 업로드할 파일에 사용할 청크 크기를 입력하세요.
- `upload_cutoff`보다 큰 파일이나 크기를 알 수 없는 파일(`rclone rcat` 또는 "rclone mount" 또는 Google 사진 또는 Google 문서로 업로드된 파일)은 이 청크 크기를 사용하여 청크로 업로드됩니다.
- `--s3-upload-concurrency` 크기의 청크는 전송당 메모리에 버퍼링됩니다.
- 높은 속도의 링크로 큰 파일을 전송하고 충분한 메모리가 있는 경우, 이 옵션을 늘리면 전송 속도가 빨라집니다.
- 큰 파일을 전송하기 위해 rclone은 알려진 크기의 큰 파일을 청크 크기보다 작은 크기로 업로드할 때마다 자동으로 청크 크기를 증가시킵니다.
- 알려지지 않은 크기의 파일은 구성된 청크 크기로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000 청크까지 있을 수 있으므로 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다.
- 청크 크기를 증가시키면 `-P` 플래그로 표시되는 진행 상태 통계의 정확도가 감소합니다. rclone은 AWS SDK에서 버퍼로 해당 크기의 청크를 전송한 후 청크가 아직 업로드되고 있는 상태라고 인식하기 때문입니다.
- 청크 크기가 클수록 AWS SDK의 버퍼 크기와 진행 상태의 정확도가 비례합니다.

### --max-upload-parts

- 멀티파트 업로드에서 사용할 최대 청크 수를 입력하세요.
- 멀티파트 업로드를 수행하는 경우 이 옵션은 청크 수의 최대 값을 정의합니다.
- 일부 서비스는 10,000 청크(아마존 S3 사양)를 지원하지 않으므로 유용할 수 있습니다.
- rclone은 알려진 크기의 큰 파일을 전송할 때 청크 크기를 자동으로 증가시켜 이러한 청크 수 제한을 유지합니다.

### --copy-cutoff

- 멀티파트 복사로 전환하는 파일의 크기 기준을 입력하세요.
- 이 값을 초과하는 크기의 파일은 이 크기의 청크로 복사됩니다.
- 최소 값은 0이고 최대 값은 5 GiB입니다.

### --disable-checksum

- 객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
- 기본적으로 rclone은 업로드하기 전에 입력 파일의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다. 이는 데이터 무결성 확인에 유용하지만 대용량 파일의 업로드를 시작하는 데 시간이 오래 걸릴 수 있습니다.

### --shared-credentials-file

- 공유 자격 증명 파일의 경로를 입력하세요.
- `env_auth`가 true인 경우 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.
- 이 변수가 비어 있는 경우 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 변수 값이 비어 있다면 현재 사용자의 홈 디렉토리가 기본값으로 사용됩니다.
- Linux/OSX: "$HOME/.aws/credentials"
- Windows: "%USERPROFILE%\.aws\credentials"

### --profile

- 공유 자격 증명 파일에서 사용할 프로필 이름을 입력하세요.
- `env_auth`가 true인 경우 rclone은 공유 자격 증명 파일에서 이 변수로 지정된 프로필을 사용합니다.
- 비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 기본값을 사용합니다.

### --session-token

- AWS 세션 토큰을 입력하세요.

### --upload-concurrency

- 멀티파트 업로드의 동시성을 입력하세요.
- 동일한 파일의 청크를 동시에 업로드하는 개수입니다.
- 고속 링크로 대용량 파일을 업로드하고이 업로드가 대역폭을 완전히 활용하지 못하는 경우, 이 값을 늘리면 전송 속도가 향상될 수 있습니다.

### --force-path-style

- 경로 스타일 액세스를 사용하려면 true로 설정하세요. 가상 호스트 스타일 액세스를 사용하려면 false로 설정하세요.
- true로 설정하면 rclone은 경로 스타일 액세스를 사용하고, false로 설정하면 가상 경로 스타일 액세스를 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참고하세요.
- AWS, Aliyun OSS, Netease COS 또는 Tencent COS와 같은 일부 공급자는 이 값이 false로 설정되어야 하며, rclone은 이 값을 제공자 설정에 따라 자동으로 설정합니다.

### --v2-auth

- true로 설정하면 v2 인증을 사용합니다.
- false로 설정하면 rclone은 v4 인증을 사용합니다. 설정하지 않으면 rclone은 v4 인증을 사용합니다.
- v4 시그니처가 작동하지 않는 경우에만 v2 인증을 사용하세요. 예: Jewel/v10 이전의 CEPH.

### --list-chunk

- 리스트 청크의 크기를 입력하세요. (각 ListObject S3 요청에 대한 응답 목록의 크기)
- 이 옵션은 AWS S3 사양에서 "MaxKeys", "max-items" 또는 "page-size"로도 알려져 있습니다.
- 대부분의 서비스는 요청 수가 1000개 이상이어도 응답 목록을 1000개로 자르지만, 전역적으로 최대값으로 설정되어 있습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html)를 참고하세요.
- Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 높일 수 있습니다.

### --list-version

- 사용할 ListObjects 버전을 입력하세요: 1,2 또는 자동으로 설정하려면 0을 입력하세요.
- AWS S3 처음 출시 시 버킷의 객체를 나열하기 위해 ListObjects 호출만 제공했습니다.
- 그러나 2016년 5월에 ListObjectsV2 호출이 도입되었습니다. 이는 훨씬 더 높은 성능을 제공하므로 가능한 경우 사용해야 합니다.
- 기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 호출할 목록 객체 방법을 추측합니다. 추측이 잘못되면 여기서 수동으로 설정할 수 있습니다.

### --list-url-encode

- 목록을 URL 인코딩할지 여부를 입력하세요: true/false/unset
- 일부 공급자는 URL 인코딩 목록을 지원하며 파일 이름에 제어 문자를 사용할 때 이 방법이 더 신뢰할 수 있습니다. unset으로 설정하면 rclone은 제공자 설정에 따라 적용할 값을 선택합니다. rclone의 선택을 재정의할 수 있습니다.

### --no-check-bucket

- 버킷의 존재를 확인하거나 생성하지 않으려면 설정하세요.
- 버킷이 이미 존재하는 경우에는 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
- 버킷 생성 권한이 없는 사용자인 경우 필요할 수 있습니다. v1.52.0 이전에는 버그로 인해 무시되었습니다.

### --no-head

- 업로드된 객체의 정합성을 확인하기 위해 HEAD 요청을 수행하지 않습니다.
- rclone은 PUT로 객체를 업로드 한 후 200 OK 메시지를 수신하면 제대로 업로드된 것으로 간주합니다.
- 특히 다음 항목을 가정합니다.
  - 메타데이터(수정 시간, 스토리지 클래스 및 콘텐츠 유형)은 업로드 시와 동일합니다.
  - 크기는 업로드된 것과 동일합니다.
- 단일 파트 PUT의 응답에서는 다음 항목을 읽습니다.
  - MD5SUM
  - 업로드된 날짜
- 멀티파트 업로드의 경우 이러한 항목은 읽히지 않습니다.
- 크기를 모르는 원본 개체가 업로드되는 경우 rclone은 HEAD 요청을 수행합니다.
- 이 플래그를 설정하면 잘못된 크기를 비롯한 업로드 실패 발생 가능성이 증가하므로 보통의 운영에는 권장되지 않습니다. 실제로 이 플래그를 설정하여 업로드 실패가 발생할 가능성은 매우 작습니다.

### --no-head-object

- 객체를 가져올 때 HEAD를 GET 전에 수행하지 않습니다.

### --encoding

- 백엔드의 인코딩을 입력하세요.
- 자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참고하세요.

### --memory-pool-flush-time

- 내부 메모리 버퍼 풀을 얼마나 자주 플러시할지 입력하세요.
- 추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
- 이 옵션은 사용되지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

### --memory-pool-use-mmap

- 내부 메모리 풀에서 mmap 버퍼를 사용할지 여부를 입력하세요.

### --disable-http2

- S3 백엔드에 대해 http2 사용을 비활성화합니다.
- 현재 s3 (특히 minio) 백엔드와 HTTP/2에 관한 문제가 해결되지 않은 상태입니다. 기본적으로 s3 백엔드에서는 HTTP/2가 활성화되지만 여기서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그가 제거될 것입니다.
- 참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631

### --download-url

- 다운로드용 사용자 정의 엔드포인트를 입력하세요.
- AWS S3는 CloudFront 네트워크를 통해 다운로드 된 데이터에 대해 더 저렴한 이그레스를 제공하므로 일반적으로 CloudFront CDN URL로 설정됩니다.

### --use-multipart-etag

- 확인을 위해 멀티파트 업로드에서 ETag를 사용할지 여부를 입력하세요.
- true, false 또는 기본값(공급자에 따름)을 사용하세요.

### --use-presigned-request

- 단일 파트 업로드에 대해 선행 서명 된 요청을 사용할지 여부를 입력하세요.
- false로 설정하면 rclone은 AWS SDK의 PutObject를 사용하여 개체를 업로드합니다.
- rclone의 버전이 1.59보다 작은 경우, 단일 파트 객체를 업로드하기 위해 선행 서명된 요청을 사용하고, 이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 위해서만 필요합니다.

### --versions

- 디렉토리 목록에 이전 버전을 포함합니다.

### --version-at

- 지정한 시간에 파일 버전을 표시합니다.
- 매개변수는 날짜, "2006-01-02" 날짜 시간 "2006-01-02 15:04:05" 또는 그 이전으로 얼마나 먼 시간인지를 나타내는 기간(예: "100d" 또는 "1h")일 수 있습니다.
- 이 값을 사용할 때 파일 쓰기 작업은 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.
- 유효한 형식에 대해서는 [시간 옵션 문서](/docs/#time-option)를 참고하세요.

### --decompress

- gzip으로 인코딩된 객체를 압축 해제합니다.
- S3에 "Content-Encoding: gzip"로 객체를 업로드할 수도 있습니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
- 이 플래그가 설정되면 rclone은 "Content-Encoding: gzip"로 받은 파일을 해제합니다. 이는 파일의 크기 및 해시를 확인할 수 없지만 파일 내용은 해제됩니다.

### --might-gzip

- 백엔드가 객체를 gzip으로 압축할 수 있는 경우 이 값을 설정하세요.
- 일반적으로 공급자는 다운로드할 때 객체를 변경하지 않습니다. `Content-Encoding: gzip`로 업로드되지 않은 객체는 다운로드될 때 설정되지 않습니다.
- 그러나 일부 공급자는 gzip으로 압축하지 않은 객체에 대해서도 gzip으로 연산하거나 압축할 수 있습니다(예: Cloudflare).
- 다음과 같은 오류가 발생하는 경우, 이 플래그를 설정하고 rclone이 `Content-Encoding: gzip`와 청크전송 인코딩을 사용하여 객체를 다운로드하면 rclone은 객체를 동시에 압축 해제합니다.
  ```
  ERROR corrupted on transfer: sizes differ NNN vs MMM
  ```
- unset(기본값)으로 설정하면 rclone은 제공자 설정에 따라 적용할 값을 선택합니다. rclone의 선택을 재정의할 수 있습니다.

### --no-system-metadata

- 시스템 메타데이터의 설정 및 읽기를 억제합니다.

## 옵션

### --access-key-id value

- AWS 액세스 키 ID를 입력하세요.
- 기본값: $ACCESS_KEY_ID

### --acl value

- 버킷 및 객체 생성 시 사용되는 Canned ACL을 입력하세요.
- 기본값: $ACL

### --endpoint value

- Tencent COS API의 엔드포인트를 입력하세요.
- 기본값: $ENDPOINT

### --env-auth

- 런타임에서 AWS 인증 정보(환경 변수 또는 EC2/ECS 메타데이터)를 가져옵니다.
- 기본값: false
- $ENV_AUTH

### --help, -h

- 도움말을 표시합니다.

### --secret-access-key value

- AWS 시크릿 액세스 키(비밀번호)를 입력하세요.
- 기본값: $SECRET_ACCESS_KEY

### --storage-class value

- Tencent COS에 새로운 객체를 저장할 때 사용할 스토리지 클래스를 입력하세요.
- 기본값: $STORAGE_CLASS

### 일반

### --name value

- 스토리지의 이름을 입력하세요.
- 기본값: 자동 생성됨

### --path value

- 스토리지의 경로를 입력하세요.

