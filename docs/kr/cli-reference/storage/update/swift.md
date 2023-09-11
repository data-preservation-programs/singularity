# OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

{% code fullWidth="true" %}
```
이름:
   singularity storage update swift - OpenStack Swift (Rackspace Cloud Files, Memset Memstore, OVH)

사용법:
   singularity storage update swift [command options] <이름|아이디>

설명:
   --env-auth
      표준 OpenStack 형식의 환경 변수에서 swift 자격 증명을 가져옵니다.

      예제:
         | false | 다음 단계에서 swift 자격 증명을 입력합니다.
         | true  | 환경 변수에서 swift 자격 증명을 가져옵니다.
         |       | 이 옵션을 사용하는 경우 다른 필드는 비워 두십시오.

   --user
      로그인할 사용자 이름입니다 (OS_USERNAME).

   --key
      API 키 또는 암호 (OS_PASSWORD).

   --auth
      서버의 인증 URL (OS_AUTH_URL).

      예제:
         | https://auth.api.rackspacecloud.com/v1.0     | Rackspace US
         | https://lon.auth.api.rackspacecloud.com/v1.0 | Rackspace UK
         | https://identity.api.rackspacecloud.com/v2.0 | Rackspace v2
         | https://auth.storage.memset.com/v1.0         | Memset Memstore UK
         | https://auth.storage.memset.com/v2.0         | Memset Memstore UK v2
         | https://auth.cloud.ovh.net/v3                | OVH

   --user-id
      로그인할 사용자 ID (선택 사항) - 대부분의 swift 시스템은 사용자를 사용하고 이 필드는 비워 둡니다 (v3 인증) (OS_USER_ID).

   --domain
      사용자 도메인 (선택 사항) (v3 인증) (OS_USER_DOMAIN_NAME).

   --tenant
      입주자 이름 (v1 인증의 경우 선택 사항이며 이 필드 또는 tenant_id가 필요합니다. 그렇지 않은 경우) (OS_TENANT_NAME 또는 OS_PROJECT_NAME).

   --tenant-id
      입주자 ID (v1 인증의 경우 선택 사항이며 이 필드 또는 입주자가 필요합니다. 그렇지 않은 경우) (OS_TENANT_ID).

   --tenant-domain
      입주자 도메인 (선택 사항) (v3 인증) (OS_PROJECT_DOMAIN_NAME).

   --region
      리전 이름 (선택 사항) (OS_REGION_NAME).

   --storage-url
      저장소 URL (선택 사항) (OS_STORAGE_URL).

   --auth-token
      대체 인증의 인증 토큰 (선택 사항) (OS_AUTH_TOKEN).

   --application-credential-id
      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID).

   --application-credential-name
      애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME).

   --application-credential-secret
      애플리케이션 자격 증명 비밀 (OS_APPLICATION_CREDENTIAL_SECRET).

   --auth-version
      인증 버전 (옵션) - 인증 URL에 버전이 없는 경우 (ST_AUTH_VERSION)로 설정합니다.

   --endpoint-type
      서비스 목록에서 선택할 엔드포인트 유형 (OS_ENDPOINT_TYPE).

      예제:
         | public   | 공용 (기본값, 확실하지 않은 경우 이 옵션을 선택하십시오)
         | internal | 내부 (내부 서비스 넷 사용)
         | admin    | 관리자

   --leave-parts-on-error
      실패 시 업로드를 중단하지 않고 호출을 피합니다.
      
      이 옵션은 다른 세션에서 업로드를 재개하기 위해 true로 설정해야 합니다.

   --storage-policy
      새 컨테이너를 만들 때 사용할 저장소 정책입니다.
      
      이렇게 하면 새로운 컨테이너를 만들 때 지정된 저장소 정책이 적용됩니다.
      이후에 정책을 변경할 수 없습니다.
      허용되는 구성 값 및 그 의미는 Swift 저장소 공급자에 따라 다릅니다.

      예제:
         | <unset> | 기본값
         | pcs     | OVH Public Cloud Storage
         | pca     | OVH Public Cloud Archive

   --chunk-size
      이 크기를 초과하는 파일은 _segments 컨테이너로 분할됩니다.
      
      이 크기를 초과하는 파일은 _segments 컨테이너로 분할됩니다. 기본값은 최대값인 5 GiB입니다.

   --no-chunk
      스트리밍 업로드 중에 파일을 청크로 나누지 않습니다.
      
      스트리밍 업로드(예: rcat 또는 마운트 사용)를 할 때 이 플래그를 설정하면
      swift 백엔드에서 청크로 파일을 업로드하지 않습니다.
      
      이렇게 하면 최대 업로드 크기가 5 GiB로 제한됩니다.
      그러나 청크로 나누지 않은 파일은 다루기 쉽고 MD5SUM이 있습니다.
      
      일반 복사 작업을 수행할 때 rclone은 여전히 chunk_size보다 큰 파일을 청크로 나눕니다.

   --no-large-objects
      정적 및 동적 대형 개체 지원을 비활성화합니다.
      
      Swift는 5 GiB보다 큰 파일을 투명하게 저장할 수 없습니다. 이를 위해
      두 가지 방법이 있습니다: 정적 또는 동적 대형 개체이며
      API에서는 rclone에서 파일이 정적인지 동적인지 판단할 수 없습니다.
      이로 인해 checksum을 읽을 때 object에 대한 HEAD 요청을 수행해야 하는 등의 작업이 필요합니다.
      
      `no_large_objects`을 설정하면 rclone은 정적 또는 동적 대형 개체가 저장되어 있지 않다고 가정합니다.
      이는 더 이상 추가 HEAD 호출을 수행하지 않아 성능이 크게 향상되며
      `--checksum`을 설정한 경우에 특히 swift에서 swift로 전송하는 경우에 유리합니다.
      
      이 옵션을 설정하면 `no_chunk`를 의미하며
      청크되지 않은 파일도 업로드되지 않으므로 5 GiB보다 큰 파일은 업로드에 실패합니다.
      
      이 옵션을 설정하고 정적 또는 동적 대형 개체가 있는 경우 이것은 정확하지 않은 해시를 제공합니다.
      다운로드는 성공하지만 삭제 및 복사와 같은 기타 작업은 실패합니다.
      
      
   --encoding
      백엔드의 인코딩 방식입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --application-credential-id value      애플리케이션 자격 증명 ID (OS_APPLICATION_CREDENTIAL_ID). [$APPLICATION_CREDENTIAL_ID]
   --application-credential-name value    애플리케이션 자격 증명 이름 (OS_APPLICATION_CREDENTIAL_NAME). [$APPLICATION_CREDENTIAL_NAME]
   --application-credential-secret value  애플리케이션 자격 증명 비밀 (OS_APPLICATION_CREDENTIAL_SECRET). [$APPLICATION_CREDENTIAL_SECRET]
   --auth value                           서버의 인증 URL (OS_AUTH_URL). [$AUTH]
   --auth-token value                     대체 인증의 인증 토큰 (선택 사항) (OS_AUTH_TOKEN). [$AUTH_TOKEN]
   --auth-version value                   인증 버전 - 인증 URL에 버전이 없는 경우 (ST_AUTH_VERSION)로 설정합니다. (기본값: 0) [$AUTH_VERSION]
   --domain value                         사용자 도메인 - 선택 사항 (v3 인증) (OS_USER_DOMAIN_NAME) [$DOMAIN]
   --endpoint-type value                  서비스 목록에서 선택할 엔드포인트 유형 (기본값: "public") (OS_ENDPOINT_TYPE) [$ENDPOINT_TYPE]
   --env-auth                             표준 OpenStack 형식의 환경 변수에서 swift 자격 증명을 가져옵니다. (기본값: false) [$ENV_AUTH]
   --help, -h                             도움말 표시
   --key value                            API 키 또는 암호 (OS_PASSWORD). [$KEY]
   --region value                         리전 이름 - 선택 사항 (OS_REGION_NAME) [$REGION]
   --storage-policy value                 새 컨테이너를 만들 때 사용할 저장소 정책입니다. [$STORAGE_POLICY]
   --storage-url value                    저장소 URL - 선택 사항 (OS_STORAGE_URL) [$STORAGE_URL]
   --tenant value                         입주자 이름 - v1 인증의 경우 선택 사항이며 이 필드 또는 tenant_id가 필요합니다. 그렇지 않은 경우 (OS_TENANT_NAME 또는 OS_PROJECT_NAME) [$TENANT]
   --tenant-domain value                  입주자 도메인 - 선택 사항 (v3 인증) (OS_PROJECT_DOMAIN_NAME) [$TENANT_DOMAIN]
   --tenant-id value                      입주자 ID - v1 인증의 경우 선택 사항이며 이 필드 또는 입주자가 필요합니다. 그렇지 않은 경우 (OS_TENANT_ID) [$TENANT_ID]
   --user value                           로그인할 사용자 이름 (OS_USERNAME) [$USER]
   --user-id value                        로그인할 사용자 ID - 선택 사항 - 대부분의 swift 시스템은 사용자를 사용하고 이 필드는 비워 둡니다 (v3 인증) (OS_USER_ID) [$USER_ID]

   고급

   --chunk-size value      이 크기를 초과하는 파일은 _segments 컨테이너로 분할됩니다. (기본값: "5Gi") [$CHUNK_SIZE]
   --encoding value        백엔드의 인코딩 방식입니다. (기본값: "Slash,InvalidUtf8") [$ENCODING]
   --leave-parts-on-error  실패 시 업로드를 중단하지 않고 호출을 피합니다. (기본값: false) [$LEAVE_PARTS_ON_ERROR]
   --no-chunk              스트리밍 업로드 중에 파일을 청크로 나누지 않습니다. (기본값: false) [$NO_CHUNK]
   --no-large-objects      정적 및 동적 대형 개체 지원을 비활성화합니다. (기본값: false) [$NO_LARGE_OBJECTS]

```
{% endcode %}