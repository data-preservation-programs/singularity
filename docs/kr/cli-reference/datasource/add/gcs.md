# Google Cloud Storage (Google Drive가 아닙니다)

{% code fullWidth="true" %}
```
명령:
   singularity datasource add gcs - Google Cloud Storage (Google Drive가 아닙니다)

사용법:
   singularity datasource add gcs [command options] <dataset_name> <source_path>

설명:
   --gcs-anonymous
      자격증명 없이 공개 버킷 및 객체에 액세스합니다.
      
      파일을 다운로드만 하고 자격증명을 구성하지 않으려면 'true'로 설정하십시오.

   --gcs-auth-url
      인증 서버 URL입니다.
      
      기본값을 사용하려면 비워 두십시오.

   --gcs-bucket-acl
      새로운 버킷에 대한 액세스 제어 목록입니다.

      예:
         | authenticatedRead | 프로젝트 팀 소유자는 소유자 액세스를 받습니다.
                             | 모든 인증된 사용자는 READER 액세스를 받습니다.
         | private           | 프로젝트 팀 소유자는 소유자 액세스를 받습니다.
                             | 비워 둘 경우 기본값으로 설정됩니다.
         | projectPrivate    | 프로젝트 팀 멤버는 역할에 따라 액세스를 받습니다.
         | publicRead        | 프로젝트 팀 소유자는 소유자 액세스를 받습니다.
                             | 모든 사용자는 READER 액세스를 받습니다.
         | publicReadWrite   | 프로젝트 팀 소유자는 소유자 액세스를 받습니다.
                             | 모든 사용자는 WRITER 액세스를 받습니다.

   --gcs-bucket-policy-only
      액세스 검사는 버킷 수준 IAM 정책을 사용해야 합니다.
      
      Bucket Policy Only가 설정된 버킷에 객체를 업로드하려면 이 설정을 해야 합니다.
      
      설정된 경우 rclone은 다음과 같이 동작합니다:
      
      - 버킷에 설정된 ACL을 무시합니다.
      - 객체에 설정된 ACL을 무시합니다.
      - Bucket Policy Only가 설정된 버킷을 만듭니다.
      
      문서: https://cloud.google.com/storage/docs/bucket-policy-only
      

   --gcs-client-id
      OAuth 클라이언트 ID입니다.
      
      일반적으로 비워둡니다.

   --gcs-client-secret
      OAuth 클라이언트 비밀입니다.
      
      일반적으로 비워둡니다.

   --gcs-decompress
      이 설정이 있으면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      "Content-Encoding: gzip"로 GCS에 객체를 업로드할 수 있습니다.
      일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정된 경우 rclone은 수신한 파일을 "Content-Encoding: gzip"로
      압축 해제합니다. 이는 rclone이 크기와 해시를 확인할 수 없게 되지만 파일 내용은 압축 해제됩니다.
      

   --gcs-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --gcs-endpoint
      서비스의 엔드포인트입니다.
      
      일반적으로 비워둡니다.

   --gcs-env-auth
      런타임에서 GCP IAM 자격증명을 가져옵니다 (환경 변수 또는 인스턴스 메타 데이터). 중요한 점은 service_account_file과 service_account_credentials가 비어 있을 때만 해당됩니다.

      예:
         | false | 다음 절차에서 자격증명을 입력합니다.
         | true  | 환경(환경 변수 또는 IAM)에서 GCP IAM 자격증명을 가져옵니다.

   --gcs-location
      새로 생성된 버킷의 위치입니다.

      예:
         | <unset>                 | 기본 위치(미국)
         | asia                    | 아시아를 위한 다중 지역 위치
         | eu                      | 유럽을 위한 다중 지역 위치
         | us                      | 미국을 위한 다중 지역 위치
         | asia-east1              | 대만
         | asia-east2              | 홍콩
         | asia-northeast1         | 도쿄
         | asia-northeast2         | 오사카
         | asia-northeast3         | 서울
         | asia-south1             | 뭄바이
         | asia-south2             | 델리
         | asia-southeast1         | 싱가포르
         | asia-southeast2         | 자카르타
         | australia-southeast1    | 시드니
         | australia-southeast2    | 멜버른
         | europe-north1           | 핀란드
         | europe-west1            | 벨기에
         | europe-west2            | 런던
         | europe-west3            | 프랑크푸르트
         | europe-west4            | 네덜란드
         | europe-west6            | 취리히
         | europe-central2         | 와르샤와
         | us-central1             | 아이오와
         | us-east1                | 사우스 캐롤라이나
         | us-east4                | 미국 버지니아 북부
         | us-west1                | 오레곤
         | us-west2                | 캘리포니아
         | us-west3                | 솔트레이크시티
         | us-west4                | 라스베이거스
         | northamerica-northeast1 | 몬트리올
         | northamerica-northeast2 | 토론토
         | southamerica-east1      | 상파울로
         | southamerica-west1      | 산티아고
         | asia1                   | 듀얼 리전: 아시아-북동1 및 아시아-북동2.
         | eur4                    | 듀얼 리전: 유럽-북부1 및 유럽-서부4.
         | nam4                    | 듀얼 리전: 미국-중부1 및 미국-동부1.

   --gcs-no-check-bucket
      설정되면 버킷의 존재를 확인하거나 만들지 않습니다.
      
      버킷이 이미 존재한다는 것을 알고 있는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      

   --gcs-object-acl
      새로운 객체에 대한 액세스 제어 목록입니다.

      예:
         | authenticatedRead      | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 모든 인증된 사용자는 READER 액세스를 받습니다.
         | bucketOwnerFullControl | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 프로젝트 팀 소유자는 소유자 액세스를 받습니다.
         | bucketOwnerRead        | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 프로젝트 팀 소유자는 READER 액세스를 받습니다.
         | private                | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 비워 둘 경우 기본값으로 설정됩니다.
         | projectPrivate         | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 프로젝트 팀 멤버는 역할에 따라 액세스를 받습니다.
         | publicRead             | 객체 소유자는 소유자 액세스를 받습니다.
                                  | 모든 사용자는 READER 액세스를 받습니다.

   --gcs-project-number
      프로젝트 번호입니다.
      
      개발자 콘솔을 참조하십시오.

   --gcs-service-account-credentials
      서비스 계정 자격증명 JSON 블롭입니다.
      
      일반적으로 비워둡니다.
      대화식 로그인 대신 SA를 사용하려는 경우에만 필요합니다.

   --gcs-service-account-file
      서비스 계정 자격증명 JSON 파일 경로입니다.
      
      일반적으로 비워둡니다.
      대화식 로그인 대신 SA를 사용하려는 경우에만 필요합니다.
      
      리딩 `~`은 파일 이름에서 확장되며 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수도 확장됩니다.

   --gcs-storage-class
      Google Cloud Storage에 객체를 저장할 때 사용할 스토리지 클래스입니다.

      예:
         | <unset>                      | 기본값
         | MULTI_REGIONAL               | 다중 지역 스토리지 클래스
         | REGIONAL                     | 지역별 스토리지 클래스
         | NEARLINE                     | Nearline 스토리지 클래스
         | COLDLINE                     | Coldline 스토리지 클래스
         | ARCHIVE                      | Archive 스토리지 클래스
         | DURABLE_REDUCED_AVAILABILITY | 내구성이 현저하게 낮은 가용성 스토리지 클래스

   --gcs-token
      JSON 블롭 형식의 OAuth 액세스 토큰입니다.

   --gcs-token-url
      토큰 서버 URL입니다.
      
      기본값을 사용하려면 비워 두십시오.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험한 옵션] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 이 인터벌이 지나면 소스 디렉터리를 자동으로 재스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: ready)

   gcs 옵션

   --gcs-anonymous value             자격증명 없이 공개 버킷 및 객체에 액세스합니다. (기본값: "false") [$GCS_ANONYMOUS]
   --gcs-auth-url value              인증 서버 URL입니다. [$GCS_AUTH_URL]
   --gcs-bucket-acl value            새로운 버킷에 대한 액세스 제어 목록입니다. [$GCS_BUCKET_ACL]
   --gcs-bucket-policy-only value    액세스 검사는 버킷 수준 IAM 정책을 사용해야 합니다. (기본값: "false") [$GCS_BUCKET_POLICY_ONLY]
   --gcs-client-id value             OAuth 클라이언트 ID입니다. [$GCS_CLIENT_ID]
   --gcs-client-secret value         OAuth 클라이언트 비밀입니다. [$GCS_CLIENT_SECRET]
   --gcs-decompress value            이 설정이 있으면 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: "false") [$GCS_DECOMPRESS]
   --gcs-encoding value              백엔드에 대한 인코딩입니다. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$GCS_ENCODING]
   --gcs-endpoint value              서비스의 엔드포인트입니다. [$GCS_ENDPOINT]
   --gcs-env-auth value              런타임에서 GCP IAM 자격증명을 가져옵니다 (환경 변수 또는 인스턴스 메타 데이터). (기본값: "false") [$GCS_ENV_AUTH]
   --gcs-location value              새로 생성된 버킷의 위치입니다. [$GCS_LOCATION]
   --gcs-no-check-bucket value       설정되면 버킷의 존재를 확인하거나 만들지 않습니다. (기본값: "false") [$GCS_NO_CHECK_BUCKET]
   --gcs-object-acl value            새로운 객체에 대한 액세스 제어 목록입니다. [$GCS_OBJECT_ACL]
   --gcs-project-number value        프로젝트 번호입니다. [$GCS_PROJECT_NUMBER]
   --gcs-service-account-file value  서비스 계정 자격증명 JSON 파일 경로입니다. [$GCS_SERVICE_ACCOUNT_FILE]
   --gcs-storage-class value         Google Cloud Storage에 객체를 저장할 때 사용할 스토리지 클래스입니다. [$GCS_STORAGE_CLASS]
   --gcs-token value                 JSON 블롭 형식의 OAuth 액세스 토큰입니다. [$GCS_TOKEN]
   --gcs-token-url value             토큰 서버 URL입니다. [$GCS_TOKEN_URL]

```
{% endcode %}