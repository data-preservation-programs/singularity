# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
명령어:
   singularity storage create swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

사용법:
   singularity storage create swift [command options] [arguments...]

DESCRIPTION:
   --env-auth
      표준 OpenStack 형태의 환경 변수에서 swift 자격 증명을 가져옵니다.

      예시:
         | false | 다음 단계에서 swift 자격 증명을 입력합니다.
         | true  | 환경 변수에서 swift 자격 증명을 가져옵니다.
         |       | 이 옵션을 사용하는 경우 다른 필드는 비워 두세요.

   --user
      로그인할 사용자 이름 (OS_USERNAME).

   --key
      API 키 또는 비밀번호 (OS_PASSWORD).

   --auth
      서버의 인증 URL (OS_AUTH_URL).

      예시:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace 미국
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace 영국
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore 영국
         | https://auth.storage.memset.com/v2.0         | Memset Memstore 영국 v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      로그인할 사용자 ID - 선택 사항 - 대부분의 swift 시스템은 사용자를 사용하고 이 필드를 비워 둡니다 (v3 인증) (OS_USER_ID).

   --domain
      사용자 도메인 - 선택 사항 (v3 인증) (OS_USER_DOMAIN_NAME)

   --tenant
      테넌트 이름 - v1 인증의 경우 선택 사항이며, 그렇지 않으면 이 필드 또는 tenant_id가 필요합니다 (OS_TENANT_NAME 또는 OS_PROJECT_NAME).

   --tenant-id
      테넌트 ID - v1 인증의 경우 선택 사항이며, 그렇지 않으면 이 필드 또는 테넌트가 필요합니다 (OS_TENANT_ID).

   --tenant-domain
      테넌트 도메인 - 선택 사항 (v3 인증) (OS_PROJECT_DOMAIN_NAME).

   --region
      지역 이름 - 선택 사항 (OS_REGION_NAME).

   --storage-url
      스토리지 URL - 선택 사항 (OS_STORAGE_URL).

   --auth-token
      대체 인증에서의 인증 토큰 - 선택 사항 (OS_AUTH_TOKEN).

   --application-credential-id
      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID).

   --application-credential-name
      애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME).

   --application-credential-secret
      애플리케이션 자격 증명 비밀번호 (OS_APPLICATION_CREDENTIAL_SECRET).

   --auth-version
      인증 버전 - 선택 사항 - 인증 URL에 버전이 없는 경우 (1, 2, 3)로 설정합니다 (ST_AUTH_VERSION).

   --endpoint-type
      서비스 카탈로그에서 선택한 엔드포인트 유형 (OS_ENDPOINT_TYPE).

      예시:
         | public   | 공개 (기본값, 확실하지 않을 경우 선택하세요)
         | internal | 내부 (내부 서비스 네트워크 사용)
         | admin    | 관리자

   --leave-parts-on-error
      실패 시 업로드 중단 호출을 피하려면 true로 설정하세요.
      
      이 옵션은 세션 간에 업로드를 다시 시작할 때 true로 설정되어야 합니다.

   --storage-policy
      새 컨테이너 생성 시 사용할 스토리지 정책입니다.
      
      이 옵션을 사용하면 새 컨테이너 생성 시 해당 스토리지 정책이 적용됩니다.
      정책은 이후에 변경할 수 없습니다. 허용되는 구성 값과 의미는 사용 중인 Swift 스토리지 공급자에 따라 다릅니다.

      예시:
         | <unset> | 기본값
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --chunk-size
      이 크기 이상의 파일은 _segments 컨테이너로 분할됩니다.
      
      이 크기 이상의 파일은 _segments 컨테이너로 분할됩니다.
      이 옵션의 기본값은 최대 5GiB입니다.

   --no-chunk
      스트리밍 업로드 중 파일을 분할하지 않습니다.
      
      스트리밍 업로드(예: rcat 또는 mount 사용)를 수행할 때 이 플래그를 설정하면 스위프트 백엔드에서 파일을 분할하지 않습니다.
      
      이렇게 하면 최대 업로드 크기가 5GiB로 제한됩니다. 그러나 분할되지 않은 파일은 다루기 쉽고 MD5SUM이 있습니다.
      
      일반 복사 작업을 수행할 때 rclone은 여전히 chunk_size보다 큰 파일을 분할합니다.

   --no-large-objects
      정적 및 동적 큰 객체 지원을 비활성화합니다.
      
      Swift는 5GiB보다 큰 파일을 투명하게 저장할 수 없습니다. 이에는 정적 또는 동적 큰 객체 두 가지 방법이 있으며, API에서도 객체가 정적 또는 동적인 큰 객체인지를 HEAD 요청하지 않고 알 수 없습니다. 이를 처리하기 위해 파일이 객체인지 여부를 확인하기 위해 rclone이 HEAD 요청을 수행해야 합니다. 예를 들어 체크섬을 읽을 때와 같은 경우입니다.
      
      `no_large_objects`가 설정되면 rclone은 정적 또는 동적 큰 객체가 저장되지 않았다고 가정합니다. 따라서 rclone은 추가적인 HEAD 요청을 수행하지 않으므로 성능이 크게 향상됩니다. 특히 `--checksum`을 설정한 상태에서 swift에서 swift로 전송을 수행하는 경우에 더욱 그렇습니다.
      
      이 옵션을 설정하면 `no_chunk`도 설정됩니다. 또한 5GiB보다 큰 파일은 업로드하지 않으므로 업로드가 실패합니다.
      
      이 옵션을 설정하고 정적 또는 동적 큰 객체가 있는 경우 잘못된 해시가 반환됩니다. 다운로드는 성공하지만, 제거 및 복사와 같은 다른 작업은 실패합니다.
      

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 encoding 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --application-credential-id value      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID). [$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME). [$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  애플리케이션 자격 증명 비밀번호 (OS_APPLICATION_CREDENTIAL_SECRET). [$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           서버의 인증 URL (OS_AUTH_URL). [$AUTH]
   --auth-token value                     대체 인증에서의 인증 토큰 - 선택 사항 (OS_AUTH_TOKEN). [$AUTH_TOKEN]
   --auth-version value                   인증 버전 - 사용하지 않으면 (1, 2, 3)으로 설정합니다 (ST_AUTH_VERSION). (default: 0) [$AUTH_VERSION]
   --domain value                         사용자 도메인 - 선택 사항 (v3 인증) (OS_USER_DOMAIN_NAME) [$DOMAIN]
   --endpoint-type value                  서비스 카탈로그에서 선택한 엔드포인트 유형 (OS_ENDPOINT_TYPE). (default: "public") [$ENDPOINT_TYPE]
   --env-auth                             표준 OpenStack 형태의 환경 변수에서 swift 자격 증명을 가져옵니다. (default: false) [$ENV_AUTH]
   --help, -h                             도움말 표시
   --key value                            API 키 또는 비밀번호 (OS_PASSWORD). [$KEY]
   --region value                         지역 이름 - 선택 사항 (OS_REGION_NAME). [$REGION]
   --storage-policy value                 새 컨테이너 생성 시 사용할 스토리지 정책. [$STORAGE_POLICY]
   --storage-url value                    스토리지 URL - 선택 사항 (OS_STORAGE_URL). [$STORAGE_URL]
   --tenant value                         테넌트 이름 - v1 인증의 경우 선택 사항이며, 그렇지 않으면 이 필드 또는 tenant_id가 필요합니다 (OS_TENANT_NAME 또는 OS_PROJECT_NAME). [$TENANT]
   --tenant-domain value                  테넌트 도메인 - 선택 사항 (v3 인증) (OS_PROJECT_DOMAIN_NAME). [$TENANT_DOMAIN]
   --tenant-id value                      테넌트 ID - v1 인증의 경우 선택 사항이며, 그렇지 않으면 이 필드 또는 테넌트가 필요합니다 (OS_TENANT_ID). [$TENANT_ID]
   --user value                           로그인할 사용자 이름 (OS_USERNAME). [$USER]
   --user-id value                        로그인할 사용자 ID - 선택 사항 - 대부분의 swift 시스템은 사용자를 사용하고 이 필드를 비워 둡니다 (v3 인증) (OS_USER_ID). [$USER_ID]

   Advanced

   --chunk-size value      이 크기 이상의 파일은 _segments 컨테이너로 분할됩니다. (default: "5Gi") [$CHUNK_SIZE]
   --encoding value        백엔드의 인코딩입니다. (default: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  실패 시 업로드 중단 호출을 피하려면 true로 설정하세요. (default: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              스트리밍 업로드 중 파일을 분할하지 않습니다. (default: false) [$NO_CHUNK]
   --no-large-objects      정적 및 동적 큰 객체 지원을 비활성화합니다 (default: false) [$NO_LARGE_OBJECTS]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}