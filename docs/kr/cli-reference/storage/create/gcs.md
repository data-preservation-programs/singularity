# Google Cloud Storage (이는 Google 드라이브가 아닙니다)

{% code fullWidth="true" %}
```
이름:
   singularity storage create gcs - Google Cloud Storage (이는 Google 드라이브가 아닙니다)

사용법:
   singularity storage create gcs [명령 옵션] [인수...]

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      보통은 비워 둡니다.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      보통은 비워 둡니다.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 비워 둡니다.

   --token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 비워 둡니다.

   --project-number
      프로젝트 번호.
      
      선택 사항 - 리스트/생성/삭제 버킷에만 필요합니다. 개발자 콘솔에서 확인하세요.

   --service-account-file
      서비스 계정 정보 JSON 파일 경로.
      
      보통은 비워 둡니다.
      대화식 로그인 대신에 서비스 계정을 사용하려면 필요합니다.
      
      파일 이름에 `~`가 포함되거나 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수가 포함되어 확장됩니다.

   --service-account-credentials
      서비스 계정 정보 JSON blob.
      
      보통은 비워 둡니다.
      대화식 로그인 대신에 서비스 계정을 사용하려면 필요합니다.

   --anonymous
      자격 증명 없이 공개 버킷과 객체에 액세스합니다.
      
      파일을 다운로드만 하고 자격 증명을 구성하지 않으려면 'true'로 설정하세요.

   --object-acl
      새로운 객체에 대한 액세스 제어 목록.

      예제:
         | authenticatedRead      | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 모든 인증된 사용자가 READER 액세스를 얻습니다.
         | bucketOwnerFullControl | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 프로젝트 팀 소유자가 OWNER 액세스를 얻습니다.
         | bucketOwnerRead        | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 프로젝트 팀 소유자가 READER 액세스를 얻습니다.
         | private                | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 비워 둔 경우 기본값입니다.
         | projectPrivate         | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 프로젝트 팀 구성원은 역할에 따라 액세스를 받습니다.
         | publicRead             | 객체 소유자가 OWNER 액세스를 얻습니다.
         |                        | 모든 사용자가 READER 액세스를 얻습니다.

   --bucket-acl
      새로운 버킷에 대한 액세스 제어 목록.

      예제:
         | authenticatedRead | 프로젝트 팀 소유자가 OWNER 액세스를 얻습니다.
         |                   | 모든 인증된 사용자가 READER 액세스를 얻습니다.
         | private           | 프로젝트 팀 소유자가 OWNER 액세스를 얻습니다.
         |                   | 비워 둔 경우 기본값입니다.
         | projectPrivate    | 프로젝트 팀 구성원은 역할에 따라 액세스를 받습니다.
         | publicRead        | 프로젝트 팀 소유자가 OWNER 액세스를 얻습니다.
         |                   | 모든 사용자가 READER 액세스를 얻습니다.
         | publicReadWrite   | 프로젝트 팀 소유자가 OWNER 액세스를 얻습니다.
         |                   | 모든 사용자가 WRITER 액세스를 얻습니다.

   --bucket-policy-only
      액세스 확인은 버킷 수준의 IAM 정책을 사용해야 합니다.
      
      Bucket Policy Only가 설정된 버킷에 객체를 업로드하려면 이 옵션을 설정해야 합니다.
      
      이 옵션이 설정되면 rclone은 다음과 같이 동작합니다:
      
      - 버킷에 설정된 ACL을 무시합니다.
      - 객체에 설정된 ACL을 무시합니다.
      - Bucket Policy Only가 설정된 상태로 버킷을 생성합니다.
      
      문서: https://cloud.google.com/storage/docs/bucket-policy-only
      

   --location
      새롭게 생성되는 버킷의 위치.

      예제:
         | <unset>                 | 기본 위치(미국)을 비워 둡니다.
         | asia                    | 아시아의 멀티 지역 위치
         | eu                      | 유럽의 멀티 지역 위치
         | us                      | 미국의 멀티 지역 위치
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
         | us-east4                | 북부 버지니아
         | us-west1                | 오레곤
         | us-west2                | 캘리포니아
         | us-west3                | 솔트레이크시티
         | us-west4                | 라스베이거스
         | northamerica-northeast1 | 몬트리올
         | northamerica-northeast2 | 토론토
         | southamerica-east1      | 상파울로
         | southamerica-west1      | 산티아고
         | asia1                   | 이중 지역: 아시아-북동1, 아시아-북동2
         | eur4                    | 이중 지역: 유럽-북부1, 유럽-서부4
         | nam4                    | 이중 지역: 미국-중부1, 미국-동부1

   --storage-class
      Google Cloud Storage에 객체를 저장할 때 사용할 저장 클래스.

      예제:
         | <unset>                      | 기본값
         | MULTI_REGIONAL               | 멀티 리전 저장 클래스
         | REGIONAL                     | 리전 저장 클래스
         | NEARLINE                     | Nearline 저장 클래스
         | COLDLINE                     | Coldline 저장 클래스
         | ARCHIVE                      | Archive 저장 클래스
         | DURABLE_REDUCED_AVAILABILITY | 내구성이 낮은 가용성 저장 클래스

   --no-check-bucket
      설정하면 버킷의 존재 여부를 확인하거나 생성하지 않습니다.
      
      버킷이 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      

   --decompress
      설정하면 gzip으로 인코딩된 객체를 압축 해제합니다.
      
      GCS에 "Content-Encoding: gzip"가 설정된 파일을 업로드하는 것이 가능합니다. 일반적으로 rclone은 이러한 파일을 압축된 객체로 다운로드합니다.
      
      이 플래그가 설정되면 rclone은 이러한 파일을 수신할 때 "Content-Encoding: gzip"로 압축 해제합니다. 이는 rclone이 파일의 크기와 해시를 확인할 수 없지만 파일 내용은 압축 해제됩니다.
      

   --endpoint
      서비스의 엔드포인트.
      
      보통은 비워 둡니다.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --env-auth
      런타임에서 GCP IAM 자격 증명을 가져옵니다(환경 변수 또는 인스턴스 메타 데이터).
      
      service_account_file과 service_account_credentials가 비어 있는 경우에만 적용됩니다.

      예제:
         | false | 다음 단계에서 자격 증명을 입력하세요.
         | true  | 환경에서 GCP IAM 자격 증명을 가져옵니다(환경 변수 또는 IAM).


옵션:
   --anonymous                          자격 증명 없이 공개 버킷과 객체에 액세스합니다. (기본값: false) [$ANONYMOUS]
   --bucket-acl value                   새로운 버킷에 대한 액세스 제어 목록. [$BUCKET_ACL]
   --bucket-policy-only                 액세스 확인은 버킷 수준의 IAM 정책을 사용해야 합니다. (기본값: false) [$BUCKET_POLICY_ONLY]
   --client-id value                    OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value                OAuth 클라이언트 시크릿. [$CLIENT_SECRET]
   --env-auth                           런타임에서 GCP IAM 자격 증명을 가져옵니다(환경 변수 또는 인스턴스 메타 데이터). (기본값: false) [$ENV_AUTH]
   --help, -h                           도움말 표시
   --location value                     새롭게 생성되는 버킷의 위치. [$LOCATION]
   --object-acl value                   새로운 객체에 대한 액세스 제어 목록. [$OBJECT_ACL]
   --project-number value               프로젝트 번호. [$PROJECT_NUMBER]
   --service-account-credentials value  서비스 계정 정보 JSON blob. [$SERVICE_ACCOUNT_CREDENTIALS]
   --service-account-file value         서비스 계정 정보 JSON 파일 경로. [$SERVICE_ACCOUNT_FILE]
   --storage-class value                Google Cloud Storage에 객체를 저장할 때 사용할 저장 클래스. [$STORAGE_CLASS]

   고급

   --auth-url value    인증 서버 URL. [$AUTH_URL]
   --decompress        설정하면 gzip으로 인코딩된 객체를 압축 해제합니다. (기본값: false) [$DECOMPRESS]
   --encoding value    백엔드의 인코딩. (기본값: "Slash,CrLf,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value    서비스의 엔드포인트. [$ENDPOINT]
   --no-check-bucket   설정하면 버킷의 존재 여부를 확인하거나 생성하지 않습니다. (기본값: false) [$NO_CHECK_BUCKET]
   --token value       JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value   토큰 서버 URL. [$TOKEN_URL]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}