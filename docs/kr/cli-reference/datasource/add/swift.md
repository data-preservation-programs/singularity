# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

사용법:
   singularity datasource add swift [command options] <dataset_name> <source_path>

DESCRIPTION:
   --swift-application-credential-id
      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID).

   --swift-application-credential-name
      애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME).

   --swift-application-credential-secret
      애플리케이션 자격 증명 비밀 (OS_APPLICATION_CREDENTIAL_SECRET).

   --swift-auth
      서버의 인증 URL (OS_AUTH_URL).

      예:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --swift-auth-token
      대체 인증용 인증 토큰 - 선택 사항 (OS_AUTH_TOKEN).

   --swift-auth-version
      인증 버전 - 선택 사항 - 인증 URL에 버전이 없는 경우 (ST_AUTH_VERSION)에 (1,2,3)으로 설정. (기본값: "0")

   --swift-chunk-size
      이 크기보다 큰 파일은 _segments 컨테이너에 청크로 분할됩니다.
      
      이 크기보다 큰 파일은 _segments 컨테이너에 청크로 분할됩니다. 기본값은 5 GiB로, 최대 크기입니다.

   --swift-domain
      사용자 도메인 - 선택 사항 (v3 인증) (OS_USER_DOMAIN_NAME)

   --swift-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --swift-endpoint-type
      서비스 카탈로그에서 선택할 엔드포인트 유형 (OS_ENDPOINT_TYPE).

      예:
         | public   | Public (기본값, 확실하지 않을 경우 이것을 선택하세요)
         | internal | Internal (내부 서비스 네트워크 사용)
         | admin    | Admin

   --swift-env-auth
      표준 OpenStack 형식의 환경 변수에서 swift 자격 증명을 가져옵니다.

      예:
         | false | 다음 단계에서 swift 자격 증명을 입력하세요.
         | true  | 환경 변수에서 swift 자격 증명 가져오기.
                 | 이용하려면 다른 필드는 비워 두세요.

   --swift-key
      API 키 또는 비밀번호 (OS_PASSWORD).

   --swift-leave-parts-on-error
      실패 시 업로드 중단 호출을 회피하도록 true로 설정합니다.
      
      이 기능을 true로 설정하면 서로 다른 세션 간에 업로드를 재개할 수 있습니다.

   --swift-no-chunk
      스트리밍 업로드 중 파일을 청크로 나누지 않습니다.
      
      스트리밍 업로드(예: rcat 또는 mount 사용)를 진행할 때 이 플래그를 설정하면 swift 백엔드에서 청크로 나누지 않게 됩니다.
      
      이렇게 하면 최대 업로드 크기가 5 GiB로 제한됩니다. 그러나 청크로 나누지 않은 파일은 다루기 쉽고 MD5SUM이 있습니다.
      
      일반 복사 작업을 수행할 때 rclone은 여전히 chunk_size보다 큰 파일을 청크로 분할합니다.

   --swift-no-large-objects
      정적 및 동적 큰 객체를 지원하지 않도록 비활성화합니다.
      
      Swift는 5 GiB보다 큰 파일을 투명하게 저장할 수 없습니다. 이를 위해 두 가지 방법, 즉 정적 또는 동적 큰 객체가 있으며,
      API는 rclone이 HEAD로 객체가 정적인지 동적인지 알 수 없으므로 개별 객체에 대해 HEAD 요청을 수행해야 합니다.
      
      `no_large_objects`가 설정된 경우 rclone은 정적 또는 동적 클러스터가 저장되지 않은 것으로 가정합니다.
      이렇게 하면 추가 HEAD 호출을 수행하지 않아 성능이 크게 향상되며 특히 `--checksum`이 설정된 경우 swift에서 swift로 전송하는 경우입니다.
      
      이 옵션을 설정하면 `no_chunk`도 설정되며 파일이 5 GiB보다 크면 업로드에 실패합니다.
      
      이 옵션을 설정하고 정적 또는 동적 큰 객체가 있는 경우 이로 인해 잘못된 해시가 생성됩니다. 다운로드는 성공하지만 삭제나 복사와 같은
      기타 작업은 실패합니다.
      

   --swift-region
      리전 이름 - 선택 사항 (OS_REGION_NAME).

   --swift-storage-policy
      새 컨테이너를 생성할 때 사용할 저장 정책입니다.
      
      이 설정은 새 컨테이너를 생성할 때 지정된 저장 정책을 적용합니다. 정책은 후에 변경할 수 없습니다. 허용되는
      구성 값 및 그 의미는 사용하는 Swift 저장 공급자에 따라 달라집니다.

      예:
         | <unset> | 기본값
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive
         
   --swift-storage-url
      저장소 URL - 선택 사항 (OS_STORAGE_URL).

   --swift-tenant
      테넌트 이름 - v1 인증을 위한 선택 사항, 그렇지 않으면 이것 또는 tenant_id가 필요 (OS_TENANT_NAME 또는 OS_PROJECT_NAME).

   --swift-tenant-domain
      테넌트 도메인 - 선택 사항 (v3 인증) (OS_PROJECT_DOMAIN_NAME).

   --swift-tenant-id
      테넌트 ID - v1 인증을 위한 선택 사항, 그렇지 않으면 이것 또는 tenant가 필요 (OS_TENANT_ID).

   --swift-user
      로그인할 사용자 이름 (OS_USERNAME).

   --swift-user-id
      로그인할 사용자 ID - 선택 사항 - 대부분의 swift 시스템은 사용자를 사용하고 이 필드를 비워 둡니다 (v3 인증) (OS_USER_ID).


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터 세트를 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 지정된 간격이 경과할 때마다 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캐닝 상태를 설정합니다 (기본값: 준비)

   swift 옵션

   --swift-application-credential-id value      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID). [$SWIFT_APPLICATION_CREDENTIAL_ID]
   --swift-application-credential-name value    애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME). [$SWIFT_APPLICATION_CREDENTIAL_NAME]
   --swift-application-credential-secret value  애플리케이션 자격 증명 비밀 (OS_APPLICATION_CREDENTIAL_SECRET). [$SWIFT_APPLICATION_CREDENTIAL_SECRET]
   --swift-auth value                           서버의 인증 URL (OS_AUTH_URL). [$SWIFT_AUTH]
   --swift-auth-token value                     대체 인증용 인증 토큰 - 선택 사항 (OS_AUTH_TOKEN). [$SWIFT_AUTH_TOKEN]
   --swift-auth-version value                   인증 버전 - 선택 사항 - 인증 URL에 버전이 없는 경우 (ST_AUTH_VERSION). (기본값: "0") [$SWIFT_AUTH_VERSION]
   --swift-chunk-size value                     이 크기보다 큰 파일은 _segments 컨테이너에 청크로 분할됩니다. (기본값: "5Gi") [$SWIFT_CHUNK_SIZE]
   --swift-domain value                         사용자 도메인 - 선택 사항 (v3 인증) (OS_USER_DOMAIN_NAME) [$SWIFT_DOMAIN]
   --swift-encoding value                       백엔드의 인코딩입니다. (기본값: "Slash,InvalidUtf8") [$SWIFT_ENCODING]
   --swift-endpoint-type value                  서비스 카탈로그에서 선택할 엔드포인트 유형 (OS_ENDPOINT_TYPE). (기본값: "public") [$SWIFT_ENDPOINT_TYPE]
   --swift-env-auth value                       표준 OpenStack 형식의 환경 변수에서 swift 자격 증명을 가져옵니다. (기본값: "false") [$SWIFT_ENV_AUTH]
   --swift-key value                            API 키 또는 비밀번호 (OS_PASSWORD). [$SWIFT_KEY]
   --swift-leave-parts-on-error value           실패 시 업로드 중단 호출을 회피하도록 true로 설정합니다. (기본값: "false") [$SWIFT_LEAVE_PARTS_ON_ERROR]
   --swift-no-chunk value                       스트리밍 업로드 중 파일을 청크로 나누지 않습니다. (기본값: "false") [$SWIFT_NO_CHUNK]
   --swift-no-large-objects value               정적 및 동적 큰 객체 지원 비활성화 (기본값: "false") [$SWIFT_NO_LARGE_OBJECTS]
   --swift-region value                         리전 이름 - 선택 사항 (OS_REGION_NAME). [$SWIFT_REGION]
   --swift-storage-policy value                 새 컨테이너를 생성할 때 사용할 저장 정책입니다. [$SWIFT_STORAGE_POLICY]
   --swift-storage-url value                    저장소 URL - 선택 사항 (OS_STORAGE_URL). [$SWIFT_STORAGE_URL]
   --swift-tenant value                         테넌트 이름 - v1 인증을 위한 선택 사항, 그렇지 않으면 이것 또는 tenant_id가 필요 (OS_TENANT_NAME 또는 OS_PROJECT_NAME). [$SWIFT_TENANT]
   --swift-tenant-domain value                  테넌트 도메인 - 선택 사항 (v3 인증) (OS_PROJECT_DOMAIN_NAME). [$SWIFT_TENANT_DOMAIN]
   --swift-tenant-id value                      테넌트 ID - v1 인증을 위한 선택 사항, 그렇지 않으면 이것 또는 tenant가 필요 (OS_TENANT_ID). [$SWIFT_TENANT_ID]
   --swift-user value                           로그인할 사용자 이름 (OS_USERNAME). [$SWIFT_USER]
   --swift-user-id value                        로그인할 사용자 ID - 선택 사항 - 대부분의 swift 시스템은 사용자를 사용하고 이 필드를 비워둡니다 (v3 인증) (OS_USER_ID). [$SWIFT_USER_ID]

```
{% endcode %}