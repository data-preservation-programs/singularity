# Google Cloud Storage (Google 드라이브가 아님)

{% code fullWidth="true" %}
```
이름:
   singularity storage update gcs - Google Cloud Storage (Google 드라이브가 아님)

사용법:
   singularity storage update gcs [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 Id.
      
      일반적으로 비워둡니다.

   --client-secret
      OAuth 클라이언트 Secret.
      
      일반적으로 비워둡니다.

   --token
      OAuth 접근 토큰(JSON blob).

   --auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 비워둡니다.

   --project-number
      프로젝트 번호.
      
      목록/생성/삭제 버킷에만 필요합니다. 개발자 콘솔 참조.

   --service-account-file
      서비스 계정 인증 정보 JSON 파일 경로.
      
      일반적으로 비워둡니다.
      대화형 로그인 대신 서비스 계정(SA)을 사용하려면 필요합니다.
      
      `~`은 파일 이름에서 확장되며, `${RCLONE_CONFIG_DIR}`과 같은 환경 변수도 확장됩니다.

   --service-account-credentials
      서비스 계정 인증 정보 JSON blob.
      
      일반적으로 비워둡니다.
      대화형 로그인 대신 서비스 계정(SA)을 사용하려면 필요합니다.

   --anonymous
      인증 정보 없이 공개 버킷과 객체에 액세스합니다.
      
      파일을 다운로드하기만 하려는 경우 'true'로 설정합니다.

   --object-acl
      새 객체에 대한 액세스 제어 목록.

      예:
         | authenticatedRead      | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 모든 인증된 사용자가 READER 액세스를 받습니다.
         | bucketOwnerFullControl | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 프로젝트 팀 소유자가 소유자 액세스를 받습니다.
         | bucketOwnerRead        | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 프로젝트 팀 소유자가 READER 액세스를 받습니다.
         | private                | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 비워둘 경우 기본값입니다.
         | projectPrivate         | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 프로젝트 팀 멤버는 역할에 따라 액세스를 받습니다.
         | publicRead             | 객체 소유자가 소유자 액세스를 받습니다.
         |                        | 모든 사용자가 READER 액세스를 받습니다.

   --bucket-acl
      새 버킷에 대한 액세스 제어 목록.

      예:
         | authenticatedRead | 프로젝트 팀 소유자가 소유자 액세스를 받습니다.
         |                   | 모든 인증된 사용자가 READER 액세스를 받습니다.
         | private           | 프로젝트 팀 소유자가 소유자 액세스를 받습니다.
         |                   | 비워둘 경우 기본값입니다.
         | projectPrivate    | 프로젝트 팀 멤버는 역할에 따라 액세스를 받습니다.
         | publicRead        | 프로젝트 팀 소유자가 소유자 액세스를 받습니다.
         |                   | 모든 사용자가 READER 액세스를 받습니다.
         | publicReadWrite   | 프로젝트 팀 소유자가 소유자 액세스를 받습니다.
         |                   | 모든 사용자가 WRITER 액세스를 받습니다.

   --bucket-policy-only
      액세스 확인은 버킷 수준의 IAM 정책을 사용해야 합니다.
      
      버킷 정책만 설정한 버킷에 파일을 업로드하려면 이 옵션을 설정해야 합니다.
      
      이 설정을 사용하면 rclone은 아래와 같이 작동합니다.
      
      - 버킷에 설정된 ACL을 무시합니다.
      - 객체에 설정된 ACL을 무시합니다.
      - 버킷에 "Bucket Policy Only" 설정으로 생성합니다.
      
      문서: https://cloud.google.com/storage/docs/bucket-policy-only


   --location
      새로 생성된 버킷의 위치.

      예:
         | <unset>                 | 기본 위치(미국)
         | asia                    | 아시아용 다중 지역 위치
         | eu                      | 유럽용 다중 지역 위치
         | us                      | 미국용 다중 지역 위치
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
         | europe-central2         | 바르샤바
         | us-central1             | 아이오와
         | us-east1                | 사우스캐롤라이나
         | us-east4                | 버지니아 북부
         | us-west1                | 오레곤
         | us-west2                | 캘리포니아
         | us-west3                | 솔트레이크시티
         | us-west4                | 라스베이거스
         | northamerica-northeast1 | 몬트리올
         | northamerica-northeast2 | 토론토
         | southamerica-east1      | 상파울루
         | southamerica-west1      | 산티아고
         | asia1                   | 아시아-북동1과 아시아-북동2의 듀얼 리전
         | eur4                    | 유럽-북부1과 유럽-서부4의 듀얼 리전
         | nam4                    | 미국-중앙1과 미국-동부지역1의 듀얼 리전

   --storage-class
      Google Cloud Storage에 객체를 저장할 때 사용할 스토리지 클래스.

      예:
         | <unset>                      | 기본값
         | MULTI_REGIONAL               | 멀티-리전 스토리지 클래스
         | REGIONAL                     | 리전 스토리지 클래스
         | NEARLINE                     | Nearline 스토리지 클래스
         | COLDLINE                     | Coldline 스토리지 클래스
         | ARCHIVE                      | 아카이브 스토리지 클래스
         | DURABLE_REDUCED_AVAILABILITY | 내구성 변경 전 가용성 스토리지 클래스

   --no-check-bucket
      버킷의 존재 여부를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      

   --decompress
      gzip으로 압축된 객체를 압축 해제합니다.
      
      GCS에 "Content-Encoding: gzip"으로 객체를 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 "Content-Encoding: gzip"으로 수신된 이러한 파일을 압축 해제합니다. 이렇게 하면 rclone은 크기와 해시를 확인할 수 없지만 파일 내용은 압축이 해제됩니다.
      

   --endpoint
      서비스의 엔드포인트.
      
      일반적으로 비워둡니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --env-auth
      실행 중인 환경(runtim)에서 GCP IAM 자격 증명을 가져옵니다(환경 변수 또는 인스턴스 메타데이터).
      
      service_account_file 및 service_account_credentials이 비어 있는 경우에만 적용됩니다.

      예:
         | false | 다음 단계에서 자격 증명 입력.
         | true  | 환경(환경 변수 또는 IAM)에서 GCP IAM 자격 증명을 가져옵니다.


OPTIONS:
   --anonymous                          인증 정보 없이 공개 버킷과 객체에 액세스합니다. (기본값: false) [$ANONYMOUS]
   --bucket-acl value                   새 버킷에 대한 액세스 제어 목록. [$BUCKET_ACL]
   --bucket-policy-only                 액세스 확인은 버킷 수준의 IAM 정책을 사용해야 합니다. (기본값: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuth 클라이언트 Id. [$CLIENT_ID]
   --client-secret value                OAuth 클라이언트 Secret. [$CLIENT_SECRET]
   --env-auth                           실행 중인 환경(runtim)에서 GCP IAM 자격 증명을 가져옵니다(환경 변수 또는 인스턴스 메타데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                           도움말 표시
   --location value                     새로 생성된 버킷의 위치. [$LOCATION]
   --object-acl value                   새 객체에 대한 액세스 제어 목록. [$OBJECT_ACL]
   --project-number value               프로젝트 번호. [$PROJECT_NUMBER]
   --service-account-credentials value  서비스 계정 인증 정보 JSON blob. [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         서비스 계정 인증 정보 JSON 파일 경로. [$SERVICE_ACCOUNT_FILE]
   --storage-class value                Google Cloud Storage에 객체를 저장할 때 사용할 스토리지 클래스. [$STORAGE_CLASS]

   Advanced

   --auth-url value   인증 서버 URL. [$AUTH_URL]
   --decompress       gzip으로 압축된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --encoding value   백엔드의 인코딩. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value   서비스의 엔드포인트. [$ENDPOINT]
   --no-check-bucket  버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --token value      OAuth 접근 토큰(JSON blob). [$TOKEN]
   --token-url value  토큰 서버 URL. [$TOKEN_URL]

```
{% endcode %}