# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
이름:
   singularity datasource add azureblob - Microsoft Azure Blob Storage

사용법:
   singularity datasource add azureblob [command options] <dataset_name> <source_path>

설명:
   --azureblob-access-tier
      블롭의 액세스 티어: hot, cool 또는 archive입니다.

      보관된 블롭은 그 액세스 티어를 hot 또는 cool로 설정하여 복원할 수 있습니다.
      계정 수준에서 설정된 기본 액세스 티어를 사용하려면 비워두세요.

      구체적으로 "액세스 티어"가 지정되지 않은 경우, rclone은 어떠한 티어도 적용하지 않습니다.
      객체가 수정되지 않으면 "액세스 티어"를 새로 지정해도 효과가 없습니다.
      원격지에 "보관 티어"의 블롭이 있는 경우 원격에서 데이터 전송 작업을 수행할 수 없습니다. 사용자는
      "Hot" 또는 "Cool" 티어로 블롭을 티어링하여 먼저 복원해야 합니다.

   --azureblob-account
      Azure 스토리지 계정 이름입니다.

      사용 중인 Azure 스토리지 계정 이름을 여기에 설정합니다.

      SAS URL 또는 에뮬레이터를 사용하려면 비워두세요.
      
      지금 비어있고 env_auth가 설정된 경우 가능한 경우에는 환경 변수 `AZURE_STORAGE_ACCOUNT_NAME`에서 읽습니다.

   --azureblob-archive-tier-delete
      보관 티어 블롭을 덮어쓰기 전에 삭제합니다.

      보관 티어 블롭은 업데이트할 수 없습니다. 따라서 이 플래그를 설정하지 않으면
      보관 티어 블롭을 업데이트하려고 하면 rclone은 오류를 발생시킵니다:

          --azureblob-archive-tier-delete 없이는 보관 티어 블롭을 업데이트할 수 없습니다.

      이 플래그를 설정하면 rclone이 보관 티어 블롭을 덮어쓰기 전에 기존 블롭을 삭제합니다.
      업로드가 실패하면 데이터 손실이 발생할 수 있으며 (일반적인 블롭을 업데이트하는 것과
      달리), 이렇게 하면 삭제하기 전에 보관된 티어 블롭을 삭제하면 비용이 추가로 발생할 수도 있습니다.

   --azureblob-chunk-size
      업로드 청크 크기입니다.

      이는 메모리에 저장되며 메모리에 한 번에 "--transfers" * "--azureblob-upload-concurrency"
      청크가 최대로 저장될 수 있습니다.

   --azureblob-client-certificate-password
      (선택사항) 인증서 파일의 암호입니다.

      사용하는 경우 이를 설정하십시오.
      - 인증서가 있는 서비스 원칙

      그리고 인증서에 비밀번호가 있는 경우입니다.

   --azureblob-client-certificate-path
      개인 키를 포함한 PEM 또는 PKCS12 인증서 파일의 경로입니다.

      사용하는 경우 이를 설정하십시오.
      - 인증서가 있는 서비스 원칙

   --azureblob-client-id
      사용 중인 클라이언트의 ID입니다.

      사용하는 경우 이를 설정하십시오.
      - 클라이언트 비밀로 서비스 원칙
      - 인증서가 있는 서비스 원칙
      - 사용자 이름과 암호를 가진 사용자

   --azureblob-client-secret
      서비스 원칙의 클라이언트 비밀 중 하나입니다.

      사용하는 경우 이를 설정하십시오.
      - 클라이언트 비밀로 서비스 원칙

   --azureblob-client-send-certificate-chain
      인증서 인증을 사용할 때 인증 요청이 x5c 헤더를 포함하는지 여부를 지정합니다.

      true로 설정하면 인증 요청에 x5c 헤더가 포함됩니다.

      인증서로 인증할 때 사용하는 경우 이를 설정하십시오.
      - 인증서가 있는 서비스 원칙

   --azureblob-disable-checksum
      개체 메타데이터에 MD5 체크섬을 저장하지 않습니다.

      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 개체의 메타데이터에 추가합니다.
      이는 데이터 무결성 검사에 좋지만 큰 파일의 경우 업로드를 시작하기까지 오래 걸릴 수 있습니다.

   --azureblob-encoding
      백엔드의 인코딩입니다.

      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --azureblob-endpoint
      서비스의 엔드포인트입니다.

      보통 비워두십시오.

   --azureblob-env-auth
      런타임에서 자격증명을 읽습니다. (환경 변수, CLI 또는 MSI)

      전체 정보는 [인증 문서](/azureblob#authentication)를 참조하십시오.

   --azureblob-key
      스토리지 계정의 공유 키입니다.

      SAS URL 또는 에뮬레이터를 사용하려면 비워두세요.

   --azureblob-list-chunk
      블롭 목록의 크기입니다.

      이렇게 하면 각 목록 청크에서 요청하는 블롭 수가 설정됩니다. 기본값은 최대 5000입니다.
      "List blobs" 요청에 사용 가능한 시간은 완료까지 1MB당 2분입니다. 작업의 평균 작업 시간이
      1MB당 2분보다 길게 걸리는 경우 시간 초과됩니다.
      ([원본](https://docs.microsoft.com/ko-kr/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval))
      이를 사용하여 반환하는 블롭 항목 수를 제한할 수 있습니다.

   --azureblob-memory-pool-flush-time
      내부 메모리 버퍼 풀이 얼마나 자주 플러시될지를 설정합니다.

      추가 버퍼를 필요로 하는 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용되지 않은 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --azureblob-memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --azureblob-msi-client-id
      사용할 사용자 할당된 MSI의 개체 ID, 있는 경우

      msi_client_id 또는 msi_mi_res_id가 지정된 경우 비워둡니다.

   --azureblob-msi-mi-res-id
      사용할 사용자 할당된 MSI의 Azure 리소스 ID, 있는 경우

      msi_client_id 또는 msi_object_id가 지정된 경우 비워둡니다.

   --azureblob-msi-object-id
      사용할 사용자 할당된 MSI의 개체 ID, 있는 경우

      msi_client_id 또는 msi_mi_res_id가 지정된 경우 비워둡니다.

   --azureblob-no-check-container
      컨테이너의 존재를 확인하거나 생성하지 않으려면 설정하세요.

      이미 컨테이너가 존재하는 경우 rclone에서 수행하는 트랜잭션 수를 최소화하려는 경우에 유용합니다.

   --azureblob-no-head-object
      GET을 수행할 때 HEAD를 수행하지 마세요.

   --azureblob-password
      사용자의 암호입니다.

      사용하는 경우 이를 설정하십시오.
      - 사용자 이름과 암호를 가진 사용자

   --azureblob-public-access
      컨테이너의 공개 액세스 수준입니다: blob 또는 container입니다.

      예:
         | <unset>   | 컨테이너 및 해당 블롭은 인증된 요청으로만 액세스할 수 있습니다.
                      | 기본값입니다.
         | blob      | 이 컨테이너 안의 Blob 데이터는 익명 요청을 통해 읽을 수 있습니다.
         | container | 컨테이너 및 블롭 데이터에 대해 전체 공개 읽기 액세스를 허용합니다.

   --azureblob-sas-url
      컨테이너 수준 접근에 대한 SAS URL입니다.

      계정/키 또는 Emulator를 사용하려면 비워두세요.

   --azureblob-service-principal-file
      서비스 원칙으로 사용할 자격증명이 포함된 파일의 경로입니다.

      보통 비워두십시오. 대화형 로그인 대신 서비스 원칙을 사용하려는 경우에만 필요합니다.

          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      자세한 내용은 ["Azure 서비스 프린시팔 생성"](https://docs.microsoft.com/ko-kr/cli/azure/create-an-azure-service-principal-azure-cli) 및 ["Blob 데이터에 액세스하기 위해 Azure 역할 할당"](https://docs.microsoft.com/ko-kr/azure/storage/common/storage-auth-aad-rbac-cli) 페이지를 참조하십시오.
      
      `service_principal_file`를 설정하는 대신 자격증명을 직접 rclone 구성 파일에
      `client_id`, `tenant` 및 `client_secret` 키 아래에 넣는 것이 더 편리할 수 있습니다.

   --azureblob-tenant
      서비스 원칙의 테넌트 ID입니다. (디렉터리 ID로도 알려짐)

      사용하는 경우 이를 설정하십시오.
      - 클라이언트 비밀로 서비스 원칙
      - 인증서가 있는 서비스 원칙
      - 사용자 이름과 암호를 가진 사용자

   --azureblob-upload-concurrency
      멀티파트 업로드의 동시성입니다.

      동시에 업로드되는 동일한 파일 청크의 수입니다.

      고속 연결을 통해 대량의 대형 파일을 업로드하고 이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우에는
      이 값을 높여 전송 속도를 높일 수 있습니다.

      테스트에서 업로드 속도는 업로드 동시성과 거의 직선적으로 증가합니다.
      예를 들어 기가비트 업로드를 채우려면 64로 이 값을 높여야 할 수 있습니다. 이렇게 하면 메모리를 더 사용합니다.

      청크는 메모리에 저장되며 한 번에 최대 "--transfers" * "--azureblob-upload-concurrency" 청크가 저장될 수 있습니다.

   --azureblob-upload-cutoff
      청크 업로드로 전환하는 기준(<= 256 MiB) (사용이 중단됨).

   --azureblob-use-emulat'order'ument
      제공된 경우 로컬 스토리지 에뮬레이터를 사용합니다.

      정식 Azure 저장소 엔드포인트를 사용하는 경우 비워두세요.

   --azureblob-use-msi
      관리형 서비스 ID를 사용하여 인증합니다(Azure에서만 작동).

      true로 설정하면 SAS 토큰이나 계정 키 대신 [관리형 서비스 ID](https://docs.microsoft.com/ko-kr/azure/active-directory/managed-identities-azure-resources/)
      를 사용하여 Azure Storage에 인증합니다.

      이 프로그램이 실행 중인 VM(SS)에 시스템 할당 ID가 있으면 기본적으로 사용됩니다.
      시스템 할당 ID가 없지만 사용자 할당 ID가 정확히 하나만 있는 경우, 기본적으로 사용됩니다.
      리소스에 여러 사용자 할당 ID가 있는 경우 msi_object_id, msi_client_id 또는 msi_mi_res_id
      매개변수 중 정확히 한 가지를 명시적으로 사용해야 합니다.

   --azureblob-username
      사용자 이름입니다(보통 이메일 주소).

      사용하는 경우 이를 설정하십시오.
      - 사용자 이름과 암호를 가진 사용자


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후에 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔에서 이 간격이 경과하면 소스 디렉터리를 자동으로 다시 스캔합니다. (기본값: 사용 안 함)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비 완료)

   azureblob 옵션

   --azureblob-access-tier value                    블롭의 액세스 티어: hot, cool 또는 archive입니다. [$AZUREBLOB_ACCESS_TIER]
   --azureblob-account value                        Azure 스토리지 계정 이름입니다. [$AZUREBLOB_ACCOUNT]
   --azureblob-archive-tier-delete value            덮어쓰기 전에 보관 티어 블롭을 삭제합니다. (기본값: "false") [$AZUREBLOB_ARCHIVE_TIER_DELETE]
   --azureblob-chunk-size value                     업로드 청크 크기입니다. (기본값: "4Mi") [$AZUREBLOB_CHUNK_SIZE]
   --azureblob-client-certificate-password value    인증서 파일의 암호입니다. (선택사항) [$AZUREBLOB_CLIENT_CERTIFICATE_PASSWORD]
   --azureblob-client-certificate-path value        PEM 또는 PKCS12 인증서 파일의 경로입니다. [$AZUREBLOB_CLIENT_CERTIFICATE_PATH]
   --azureblob-client-id value                      사용 중인 클라이언트의 ID입니다. [$AZUREBLOB_CLIENT_ID]
   --azureblob-client-secret value                  서비스 원칙의 클라이언트 비밀 중 하나입니다. [$AZUREBLOB_CLIENT_SECRET]
   --azureblob-client-send-certificate-chain value  인증 요청이 인증서를 기반으로 항목 이름/발행자를 지원하도록 인증 요청에 x5c 헤더를 포함할지 여부를 지정합니다. (기본값: "false") [$AZUREBLOB_CLIENT_SEND_CERTIFICATE_CHAIN]
   --azureblob-disable-checksum value               개체 메타데이터에 MD5 체크섬을 저장하지 않습니다. (기본값: "false") [$AZUREBLOB_DISABLE_CHECKSUM]
   --azureblob-encoding value                       백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$AZUREBLOB_ENCODING]
   --azureblob-endpoint value                       서비스의 엔드포인트입니다. [$AZUREBLOB_ENDPOINT]
   --azureblob-env-auth value                       런타임에서 자격증명을 읽습니다. (환경 변수, CLI 또는 MSI) (기본값: "false") [$AZUREBLOB_ENV_AUTH]
   --azureblob-key value                            스토리지 계정의 공유 키입니다. [$AZUREBLOB_KEY]
   --azureblob-list-chunk value                     블롭 목록의 크기입니다. (기본값: "5000") [$AZUREBLOB_LIST_CHUNK]
   --azureblob-memory-pool-flush-time value         내부 메모리 버퍼 풀이 얼마나 자주 플러시될지를 설정합니다. (기본값: "1m0s") [$AZUREBLOB_MEMORY_POOL_FLUSH_TIME]
   --azureblob-memory-pool-use-mmap value           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: "false") [$AZUREBLOB_MEMORY_POOL_USE_MMAP]
   --azureblob-msi-client-id value                  사용할 사용자 할당된 MSI의 개체 ID, 있는 경우 [$AZUREBLOB_MSI_CLIENT_ID]
   --azureblob-msi-mi-res-id value                  사용할 사용자 할당된 MSI의 Azure 리소스 ID, 있는 경우 [$AZUREBLOB_MSI_MI_RES_ID]
   --azureblob-msi-object-id value                  사용할 사용자 할당된 MSI의 개체 ID, 있는 경우 [$AZUREBLOB_MSI_OBJECT_ID]
   --azureblob-no-check-container value             컨테이너의 존재를 확인하거나 생성하지 않습니다. (기본값: "false") [$AZUREBLOB_NO_CHECK_CONTAINER]
   --azureblob-no-head-object value                 GET을 수행할 때 HEAD를 수행하지 마세요. (기본값: "false") [$AZUREBLOB_NO_HEAD_OBJECT]
   --azureblob-password value                       사용자의 암호 [$AZUREBLOB_PASSWORD]
   --azureblob-public-access value                  컨테이너의 공개 액세스 수준입니다: blob 또는 container입니다. [$AZUREBLOB_PUBLIC_ACCESS]
   --azureblob-sas-url value                        컨테이너 수준 접근에 대한 SAS URL입니다. [$AZUREBLOB_SAS_URL]
   --azureblob-service-principal-file value         서비스 원칙으로 사용할 자격증명이 포함된 파일의 경로입니다. [$AZUREBLOB_SERVICE_PRINCIPAL_FILE]
   --azureblob-tenant value                         서비스 원칙의 테넌트 ID입니다. (디렉터리 ID로도 알려짐) [$AZUREBLOB_TENANT]
   --azureblob-upload-concurrency value             멀티파트 업로드의 동시성입니다. (기본값: "16") [$AZUREBLOB_UPLOAD_CONCURRENCY]
   --azureblob-upload-cutoff value                  청크 업로드로 전환하는 기준(<= 256 MiB) (사용이 중단됨) [$AZUREBLOB_UPLOAD_CUTOFF]
   --azureblob-use-emulat'order'ument value                   제공된 경우 로컬 스토리지 에뮬레이터를 사용합니다. (기본값: "false") [$AZUREBLOB_USE_EMULATOR]
   --azureblob-use-msi value                        관리형 서비스 ID를 사용하여 인증합니다(작동은 Azure에서만). (기본값: "false") [$AZUREBLOB_USE_MSI]
   --azureblob-username value                       사용자 이름입니다(보통 이메일 주소). [$AZUREBLOB_USERNAME]

```
{% endcode %}