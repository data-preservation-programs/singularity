# Huawei Object Storage Service (Huawei OBS) - 저장소 업데이트

{% code fullWidth="true" %}
```
이름:
   singularity storage update s3 huaweiobs - Huawei Object Storage Service

사용법:
   singularity storage update s3 huaweiobs [command options] <name|id>

설명:
    --env-auth
      런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터).

      access_key_id와 secret_access_key가 비어 있을 경우에만 적용됩니다.

      예:
         | false | 다음 단계에서 AWS 자격 증명을 입력하세요.
         | true  | 환경(환경 변수 또는 IAM)에서 AWS 자격 증명을 가져옵니다.

    --access-key-id
      AWS Access Key ID.

      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

    --secret-access-key
      AWS Secret Access Key (비밀번호).

      익명 액세스 또는 런타임 자격 증명을 위해 비워 둡니다.

    --region
      연결할 지역입니다. - 버킷이 생성되고 데이터가 저장되는 위치입니다. 엔드포인트와 동일해야 합니다.

      예:
         | af-south-1     | AF-Johannesburg
         | ap-southeast-2 | AP-Bangkok
         | ap-southeast-3 | AP-Singapore
         | cn-east-3      | CN East-Shanghai1
         | cn-east-2      | CN East-Shanghai2
         | cn-north-1     | CN North-Beijing1
         | cn-north-4     | CN North-Beijing4
         | cn-south-1     | CN South-Guangzhou
         | ap-southeast-1 | CN-Hong Kong
         | sa-argentina-1 | LA-Buenos Aires1
         | sa-peru-1      | LA-Lima1
         | na-mexico-1    | LA-Mexico City1
         | sa-chile-1     | LA-Santiago2
         | sa-brazil-1    | LA-Sao Paulo1
         | ru-northwest-2 | RU-Moscow2

    --endpoint
      OBS API의 엔드포인트입니다.

      예:
         | obs.af-south-1.myhuaweicloud.com     | AF-Johannesburg
         | obs.ap-southeast-2.myhuaweicloud.com | AP-Bangkok
         | obs.ap-southeast-3.myhuaweicloud.com | AP-Singapore
         | obs.cn-east-3.myhuaweicloud.com      | CN East-Shanghai1
         | obs.cn-east-2.myhuaweicloud.com      | CN East-Shanghai2
         | obs.cn-north-1.myhuaweicloud.com     | CN North-Beijing1
         | obs.cn-north-4.myhuaweicloud.com     | CN North-Beijing4
         | obs.cn-south-1.myhuaweicloud.com     | CN South-Guangzhou
         | obs.ap-southeast-1.myhuaweicloud.com | CN-Hong Kong
         | obs.sa-argentina-1.myhuaweicloud.com | LA-Buenos Aires1
         | obs.sa-peru-1.myhuaweicloud.com      | LA-Lima1
         | obs.na-mexico-1.myhuaweicloud.com    | LA-Mexico City1
         | obs.sa-chile-1.myhuaweicloud.com     | LA-Santiago2
         | obs.sa-brazil-1.myhuaweicloud.com    | LA-Sao Paulo1
         | obs.ru-northwest-2.myhuaweicloud.com | RU-Moscow2

    --acl
      버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다.

      객체를 만들고 버킷_acl이 설정되지 않은 경우에도 이 ACL이 사용됩니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.

      사실 서버 측 복사를 수행하는 경우, S3는 소스의 ACL을 복사하는 대신 새로운 ACL을 작성합니다.

      acl이 빈 문자열이면 X-Amz-Acl: 헤더가 추가되지 않고 기본값(private)이 사용됩니다.

    --bucket-acl
      버킷을 만들 때 사용하는 Canned ACL입니다.

      자세한 정보는 https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#canned-acl을 참조하십시오.

      이 ACL은 버킷을 만들 때만 적용됩니다. 설정되지 않은 경우 "acl"이 대신 사용됩니다.

      "acl" 및 "bucket_acl"이 빈 문자열인 경우 X-Amz-Acl: 헤더가 추가되지 않고 기본값(private)이 사용됩니다.

      예:
         | private            | 소유자가 모든 권한( FULL_CONTROL)입니다.
         |                    | 다른 사람들은 액세스 권한이 없습니다 (기본값).
         | public-read        | 소유자가 모든 권한(FULL_CONTROL)입니다.
         |                    | AllUsers 그룹이 읽기 액세스 권한을 갖습니다.
         | public-read-write  | 소유자가 모든 권한(FULL_CONTROL)입니다.
         |                    | AllUsers 그룹이 읽기 및 쓰기 액세스 권한을 갖습니다.
         |                    | 이 권한은 일반적으로 버킷에 부여되지 않습니다.
         | authenticated-read | 소유자가 모든 권한(FULL_CONTROL)입니다.
         |                    | AuthenticatedUsers 그룹이 읽기 액세스 권한을 갖습니다.

    --upload-cutoff
      청크 업로드로 전환하는 임계값입니다.

      이 값보다 큰 크기의 파일은 chunk_size의 청크로 업로드됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

    --chunk-size
      업로드에 사용할 청크 크기입니다.

      upload_cutoff보다 큰 파일이나 크기를 알 수 없는 파일("rclone rcat" 이나 "rclone mount" 또는 google 사진이나 google 문서에서 업로드한 파일과 같은)을 업로드하는 경우 이 청크 크기를 사용하여 청크 업로드로 업로드됩니다.

      참고로 "--s3-upload-concurrency"는 이 크기의 청크가 전송마다 메모리에 버퍼로 저장됩니다.

      빠른 네트워크 링크에서 큰 파일을 전송하고 충분한 메모리가 있는 경우 이 값을 증가시키면 전송 속도가 향상될 것입니다.

      Rclone은 알려진 파일의 크기가 커지면 청크 크기를 자동으로 증가시켜 최대 10,000개의 청크 제한을 넘지 않도록 합니다.

      알려지지 않은 크기의 파일은 구성된 chunk_size로 업로드됩니다. 기본 청크 크기는 5 MiB이며 최대 10,000개의 청크가 있을 수 있습니다. 따라서 기본적으로 스트림 업로드할 수 있는 파일의 최대 크기는 48 GiB입니다. 큰 파일을 스트림 업로드하려면 chunk_size를 증가해야 합니다.

      청크 크기를 증가시키면 진행률 통계에 대한 정확도가 감소합니다. Rclone은 청크가 AWS SDK에 의해 버퍼로 저장된 경우 청크를 전송한 것으로 처리하지만 실제로는 여전히 업로드 중일 수 있습니다. 큰 청크 크기는 큰 AWS SDK 버퍼 및 진행률 통계를 더 신뢰할 수 없게 만듭니다.

    --max-upload-parts
      멀티파트 업로드에 사용되는 최대 부분 수입니다.

      이 옵션은 멀티파트 업로드를 수행할 때 사용할 최대 멀티파트 청크 수를 정의합니다.

      이 옵션은 10,000개의 청크를 지원하지 않는 서비스의 경우 유용할 수 있습니다.

      Rclone은 알려진 크기의 큰 파일을 업로드할 때 이 청크 크기를 자동으로 증가시켜 이 청크 수 제한을 유지합니다.

    --copy-cutoff
      멀티파트 복사로 전환되는 임계값입니다.

      이 값보다 크기가 큰 파일은 본 청크 크기로 복사됩니다.
      최소값은 0이고 최대값은 5 GiB입니다.

    --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체 메타데이터에 추가합니다. 이는 데이터 무결성 검사에 유용하지만, 큰 파일을 업로드할 때엔 시작하기까지 오랜 지연이 발생할 수 있습니다.

    --shared-credentials-file
      공유 자격 증명 파일의 경로입니다.

      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.

      이 변수가 비어 있으면 rclone은 "AWS_SHARED_CREDENTIALS_FILE" 환경 변수를 찾습니다. 환경 값이 비어 있으면 현재 사용자의 홈 디렉터리가 기본값이 됩니다.

        Linux/OSX : "$HOME/.aws/credentials"
        Windows   : "%USERPROFILE%\.aws\credentials" 

    --profile
      공유 자격 증명 파일에서 사용할 프로필입니다.

      env_auth = true이면 rclone은 공유 자격 증명 파일을 사용할 수 있습니다.이 변수는 해당 파일에서 사용할 프로필을 제어합니다.

      비어 있으면 환경 변수 "AWS_PROFILE" 또는 "default"로 기본값이 됩니다.

    --session-token
      AWS 세션 토큰.

    --upload-concurrency
      멀티파트 업로드에 대한 동시성입니다.

      동일한 파일의 청크 수입니다.

      높은 속도로 큰 파일을 업로드하고 현재 대역폭을 완전히 활용하지 못하는 경우에는이 값을 증가시키면 전송을 가속화하는 데 도움이 될 수 있습니다.

    --force-path-style
      true인 경우 path 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다.

      값이 true인 경우(기본값), rclone은 경로 스타일 액세스를 사용합니다.
      false인 경우 rclone은 가상 경로 스타일을 사용합니다. 자세한 내용은 [AWS S3 문서](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingBucket.html#access-bucket-intro)를 참조하십시오.
      
      (AWS, Aliyun OSS, Netease COS또는 Tencent COS 등) 일부 제공자는이 설정에 따라 false로 설정해야합니다. rclone은 제공자 설정을 기반으로 자동으로 수행합니다.

    --v2-auth
      true로 설정하면 v2 인증을 사용합니다.

      false인 경우(기본값) rclone은 v4 인증을 사용합니다.
      설정된 경우 rclone은 v2 인증을 사용합니다.

      v4 서명이 작동하지 않을 경우에만 사용하십시오. (예 : Jewel/v10 CEPH 전에 사용)

    --list-chunk
      목록 청크 크기(ListObject S3 요청마다 응답 리스트 크기)입니다.

      이 옵션은 AWS S3 사양에서 알려진 "MaxKeys", "max-items" 또는 "page-size"로 알려져 있습니다.
      대부분의 서비스는 1000 개보다 많이 요청해도 응답 목록을 잘라냅니다.
      AWS S3에서는 전역 최대 값이며 변경할 수 없습니다. [AWS S3](https://docs.aws.amazon.com/cli/latest/reference/s3/ls.html) 참조.
      Ceph에서는 "rgw list buckets max chunk" 옵션으로 이 값을 증가시킬 수 있습니다.

    --list-version
      사용할 ListObjects 버전: 1,2 또는 자동인식을 위해 0.

      S3가 처음 시작될 때 버킷 내 객체를 열거하기 위해서는 ListObjects 호출만 제공합니다.

      그러나 2016 년 5 월에 ListObjectsV2 호출이 도입되었습니다. 이는 월등히 높은 성능을 제공하며 가능하면 사용해야합니다.

      기본값인 0으로 설정하면 rclone은 제공자 설정에 따라 호출되어야 할 List 객체 방법을 추측합니다.
      추측이 잘못된 경우 여기에서 수동으로 설정할 수 있습니다.

    --list-url-encode
     리스트의 URL 인코딩 여부: true/false/unset
      
      일부 제공자에서는 파일 이름에 제어 문자를 사용할 때 이용 가능한 URL 인코딩 지원을 제공합니다. unset으로 설정하면 rclone은 공급자 설정에 따라 적용할 대상을 선택하게 됩니다.

    --no-check-bucket
      버킷의 존재를 확인하거나 생성하지 않으려면 설정하세요.

      버킷이 이미 존재한다면 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.

      또는 사용중인 사용자에게 버킷 생성 권한이 없는 경우 필요할 수 있습니다. v1.52.0 이전에서는 버그 때문에 무시되었을 것입니다.

    --no-head
      업로드 된 객체의 정합성을 확인하기 위한 HEAD를 실행하지 않습니다.

      rclone이 PUT로 객체를 업로드 한 후 200 OK 메시지를 받으면 올바르게 업로드 된 것으로 간주합니다.

      특히 다음을 가정합니다.

      - 업로드된 것과 동일한 메타데이터(모디파이드 시간, 스토리지 클래스 및 콘텐츠 유형)입니다.
      - 업로드된 것과 동일한 크기입니다.

      다음 항목을 읽습니다(단일 부분 PUT에 대해):

      - MD5SUM
      - 업로드된 날짜

      멀티파트 업로드의 경우 이 항목은 읽지 않습니다.

      알려지지 않은 길이의 소스 객체를 업로드하는 경우에는 rclone은 HEAD 요청을 수행합니다.

      이 플래그를 설정하면 잘못된 크기와 같은 업로드 실패의 가능성이 증가하므로 정상적인 작동에 권장하지 않습니다. 실제로 이 플래그를 설정해도 업로드 실패의 가능성은 매우 적습니다.

    --no-head-object
      GET을 실행하기 전에 HEAD를 수행하지 마세요.

    --encoding
      백엔드에 대한 인코딩입니다.

      자세한 내용은 [개요에서 인코딩 섹션](/overview/#encoding)을 참조하십시오.

    --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시된는 시간입니다.

      추가 버퍼가 필요한 업로드(예: 멀티파트)는 메모리 풀에 대한 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거될 때까지의 시간을 제어합니다.

    --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부.

    --disable-http2
      S3 백엔드에 대한 http2 사용을 비활성화합니다.

      현재 s3(특히 minio) 백엔드와 HTTP/2에 대한 미해결된 문제가 있습니다. S3 백엔드의 HTTP/2는 기본적으로 활성화되어 있지만 여기에서 비활성화할 수 있습니다. 이 문제가 해결되면이 플래그가 제거될 것입니다.

      참조: https://github.com/rclone/rclone/issues/4673, https://github.com/rclone/rclone/issues/3631


    --download-url
      다운로드를 위한 사용자 정의 엔드포인트입니다.
      보통 AWS S3는 CloudFront 네트워크를 통해 데이터를 다운로드 할 때 더 저렴한 이그레스를 제공합니다.

    --use-multipart-etag
      다중파트 업로드에서 ETag를 사용하여 검증할지 여부입니다.

      이 플래그는 true, false 또는 제공자의 기본값으로 설정할 수 있습니다.

    --use-presigned-request
      단일 부분 업로드를위한 임시 서명 된 요청 또는 PutObject를 사용할지 여부입니다.

      이 값이 false로 설정되면 rclone은 객체를 업로드하기 위해 AWS SDK의 PutObject를 사용합니다.

      Rclone <v1.59 버전은 단일 부분 객체를 업로드하기 위해 임시 서명 된 요청을 사용하고이 플래그를 true로 설정하면 해당 기능이 다시 활성화됩니다. 이는 예외적인 상황이나 테스트를 위해 사용되어야합니다.

    --versions
      디렉터리 목록에 이전 버전을 포함합니다.

    --version-at
      지정된 시점에서 파일 버전을 표시합니다.

      매개 변수는 date, "2006-01-02", datetime "2006-01-02 15:04:05" 또는 해당 시간까지의 지속 시간, 예 "100d" 또는 "1h" 일 수 있습니다.

      이 플래그를 사용하면 파일 쓰기 작업이 허용되지 않으므로 파일을 업로드하거나 삭제할 수 없습니다.

      사용 가능한 형식에 대한 자세한 정보는 [시간 옵션 문서](/docs/#time-option)를 참조하십시오.

    --decompress
     이 값이 설정되어 있으면 gzip으로 인코딩된 개체를 압축 해제합니다.

      S3에 "Content-Encoding: gzip"로 업로드 된 개체는 일반적으로 압축된 객체로 다운로드됩니다.

      이 플래그를 설정하면 rclone은 수신되는 개체와 함께 "Content-Encoding: gzip"로 이 플래그가 설정되었을 때 이 파일들을 압축 해제합니다. 이로 인해 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.

    --might-gzip
      백엔드가 개체를 gzip으로 압축 할 수 ​​있는지 여부를 설정하세요.

      일반적으로 제공자는 개체를 다운로드 할 때는 개체를 변경하지 않습니다. "Content-Encoding: gzip"로 업로드되지 않은 개체의 경우 다운로드되지 않습니다.

      그러나 일부 제공자는 "Content-Encoding: gzip"로 업로드되지 않은 경우에도 객체를 gzip으로 압축 할 수 ​​있습니다(예 : Cloudflare).

      이 플래그를 설정하고 rclone이 Content-Encoding: gzip 및 청크 전송 인코딩이 있는 개체를 다운로드하면 rclone은 개체를 실시간으로 압축 해제합니다.

      이 값을 unset(기본값)로 설정하면 rclone은 해당 공급자 설정에 따라 적용할 대상을 선택합니다. 그러나 여기에서 rclone의 선택을 재정의 할 수 있습니다.

    --no-system-metadata
      시스템 메타데이터의 설정 및 읽기를 억제

옵션:
   --access-key-id value      AWS Access Key ID. [$ACCESS_KEY_ID]
   --acl value                버킷을 만들거나 객체를 저장하거나 복사할 때 사용되는 Canned ACL입니다. [$ACL]
   --endpoint value           OBS API의 엔드포인트입니다. [$ENDPOINT]
   --env-auth                 런타임에서 AWS 자격 증명을 가져옵니다 (환경 변수 또는 env vars가 없는 경우 EC2/ECS 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                 도움말 표시
   --region value             연결할 지역입니다. - 버킷이 생성되고 데이터가 저장되는 위치입니다. 엔드포인트와 동일해야 합니다. [$REGION]
   --secret-access-key value  AWS Secret Access Key (비밀번호). [$SECRET_ACCESS_KEY]

고급

   --bucket-acl value               버킷을 만들 때 사용하는 Canned ACL입니다. [$BUCKET_ACL]
   --chunk-size value               업로드에 사용할 청크 크기입니다. (기본값: "5Mi") [$CHUNK_SIZE]
   --copy-cutoff value              멀티파트 복사로 전환되는 임계값입니다. (기본값: "4.656Gi") [$COPY_CUTOFF]
   --decompress                     gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --disable-http2                  S3 백엔드에 대한 http2 사용을 비활성화합니다. (기본값: false) [$DISABLE_HTTP2]
   --download-url value             다운로드를 위한 사용자 정의 엔드포인트입니다. [$DOWNLOAD_URL]
   --encoding value                 백엔드에 대한 인코딩입니다. (기본값: "Slash,InvalidUtf8,Dot") [$ENCODING]
   --force-path-style               true인 경우 path 스타일 액세스를 사용하고 false인 경우 가상 호스팅 스타일을 사용합니다. (기본값: true) [$FORCE_PATH_STYLE]
   --list-chunk value               목록 청크 크기(ListObject S3 요청마다 응답 리스트 크기). (기본값: 1000) [$LIST_CHUNK]
   --list-url-encode value          리스트의 URL 인코딩 여부: true/false/unset. (기본값: "unset") [$LIST_URL_ENCODE]
   --list-version value             사용할 ListObjects 버전: 1,2 또는 0으로 자동. (기본값: 0) [$LIST_VERSION]
   --max-upload-parts value         멀티파트 업로드에 사용되는 최대 부분 수. (기본값: 10000) [$MAX_UPLOAD_PARTS]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시된는 시간. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --might-gzip value               백엔드가 개체를 gzip으로 압축 할 수 ​​있는지 여부를 설정합니다. (기본값: "unset") [$MIGHT_GZIP]
   --no-check-bucket                버킷의 존재를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --no-head                        업로드 된 객체의 정합성을 확인하기 위한 HEAD를 실행하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-head-object                 GET을 실행하기 전에 HEAD를 수행하지 마세요. (기본값: false) [$NO_HEAD_OBJECT]
   --no-system-metadata             시스템 메타데이터의 설정 및 읽기를 억제합니다 (기본값: false) [$NO_SYSTEM_METADATA]
   --profile value                  공유 자격 증명 파일에서 사용할 프로필입니다. [$PROFILE]
   --session-token value            AWS 세션 토큰. [$SESSION_TOKEN]
   --shared-credentials-file value  공유 자격 증명 파일의 경로입니다. [$SHARED_CREDENTIALS_FILE]
   --upload-concurrency value       멀티파트 업로드에 대한 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 임계값. (기본값: "200Mi") [$UPLOAD_CUTOFF]
   --use-multipart-etag value       다중파트 업로드에서 ETag를 사용하여 검증할지 여부 (기본값: "unset") [$USE_MULTIPART_ETAG]
   --use-presigned-request          단일 부분 업로드를위한 임시 서명 요청 또는 PutObject 사용. (기본값: false) [$USE_PRESIGNED_REQUEST]
   --v2-auth                        true인 경우 v2 인증을 사용합니다. (기본값: false) [$V2_AUTH]
   --version-at value               지정된 시점에서 파일 버전을 표시합니다. (기본값: "off") [$VERSION_AT]
   --versions                       디렉터리 목록에 이전 버전을 포함합니다. (기본값: false) [$VERSIONS]

```
{% endcode %}