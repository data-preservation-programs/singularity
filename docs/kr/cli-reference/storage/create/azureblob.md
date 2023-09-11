# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
NAME:
   rclone 저장소 생성 azureblob - Microsoft Azure Blob Storage

사용법:
   rclone 저장소 생성 azureblob [command options] [arguments...]

설명:
   --account
      Azure Storage Account Name.
      
      사용 중인 Azure Storage Account Name을 설정합니다.
      
      빈칸으로 두고 SAS URL이나 에뮬레이터를 사용하는 경우에는 설정하지 않아도 됩니다.
      
      이 값이 비어있고 env_auth가 설정되어 있으면 환경변수 `AZURE_STORAGE_ACCOUNT_NAME`에서 읽힙니다.
      

   --env-auth
      런타임(환경변수, CLI 또는 MSI)에서 자격 증명을 읽습니다.
      
      자세한 정보는 [인증 문서](/azureblob#authentication)를 참조하십시오.

   --key
      스토리지 계정 공유 키.
      
      SAS URL이나 에뮬레이터를 사용하는 경우 비워두십시오.

   --sas-url
      컨테이너 수준 접근용 SAS URL.
      
      계정/키 또는 에뮬레이터를 사용하는 경우 비워두십시오.

   --tenant
      서비스 프린시팔의 테넌트 ID입니다. 또는 디렉터리 ID입니다.
      
      다음을 사용하는 경우 이 값을 설정하십시오.
      - 클라이언트 시크릿을 가진 서비스 프린시팔
      - 인증서를 가진 서비스 프린시팔
      - 사용자 이름과 비밀번호를 가진 사용자
      

   --client-id
      사용 중인 클라이언트의 ID입니다.
      
      다음을 사용하는 경우 이 값을 설정하십시오.
      - 클라이언트 시크릿을 가진 서비스 프린시팔
      - 인증서를 가진 서비스 프린시팔
      - 사용자 이름과 비밀번호를 가진 사용자
      

   --client-secret
      서비스 프린시팔의 클라이언트 시크릿 중 하나입니다.
      
      클라이언트 시크릿을 가진 서비스 프린시팔의 경우 이 값을 설정하십시오.
      

   --client-certificate-path
      개인 키를 포함하는 PEM 또는 PKCS12 인증서 파일의 경로입니다.
      
      인증서를 가진 서비스 프린시팔의 경우 이 값을 설정하십시오.
      

   --client-certificate-password
      인증서 파일의 암호입니다(선택 사항).
      
      인증서를 가진 서비스 프린시팔의 경우 선택적으로 이 값을 설정하십시오.
      

   --client-send-certificate-chain
      인증서 인증 시 인증 요청에 x5c 헤더를 포함할지 여부를 지정합니다.
      
      인증서를 가진 서비스 프린시팔의 경우 선택적으로 이 값을 설정하십시오.
      

   --username
      사용자 이름(일반적으로 이메일 주소)
      
      사용자 이름과 비밀번호를 가진 사용자의 경우 이 값을 설정하십시오.
      

   --password
      사용자의 비밀번호
      
      사용자 이름과 비밀번호를 가진 사용자의 경우 이 값을 설정하십시오.
      

   --service-principal-file
      서비스 프린시팔과 함께 사용할 자격 증명이 담긴 파일의 경로입니다.
      
      보통 비워둡니다. 대화형 로그인 대신 서비스 프린시팔을 사용하려면 필요합니다.
      
          $ az ad sp create-for-rbac --name "<name>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<subscription>/resourceGroups/<resource-group>/providers/Microsoft.Storage/storageAccounts/<storage-account>/blobServices/default/containers/<container>" \
            > azure-principal.json
      
      자세한 정보는 ["Azure 서비스 프린시팔 만들기"](https://docs.microsoft.com/en-us/cli/azure/create-an-azure-service-principal-azure-cli)와 ["블롭 데이터에 대한 액세스를 위한 Azure 역할 할당"](https://docs.microsoft.com/en-us/azure/storage/common/storage-auth-aad-rbac-cli) 페이지를 참조하십시오.
      
      `service_principal_file`를 설정하는 대신 자격 증명을 rclone 구성 파일의 `client_id`, `tenant` 및 `client_secret` 키에 직접 입력하는 것이 편리할 수 있습니다.
      

   --use-msi
      관리형 서비스 ID를 사용하여 인증(오직 Azure에서만 작동) (default: false).
      
      true로 설정하면, SAS 토큰이나 계정 키 대신 Azure Storage 인증에 [관리형 서비스 ID](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/)를 사용합니다.
      
      이 프로그램이 실행되는 VM(SS)에 시스템 할당 ID가 있는 경우에는 기본값으로 사용하게 됩니다. 리소스에 시스템 할당 ID가 없고 사용자 할당 ID가 정확히 하나 있는 경우 사용자 할당 ID를 기본값으로 사용하게 됩니다. 리소스에 여러 사용자 할당 ID가 있는 경우 msi_object_id, msi_client_id 또는 msi_mi_res_id 매개변수 중 정확히 하나를 명시적으로 지정해야 합니다.

   --msi-object-id
      사용할 사용자 할당 MSI의 개체 ID입니다.
      
      msi_client_id 또는 msi_mi_res_id가 지정된 경우 비워두십시오.

   --msi-client-id
      사용할 사용자 할당 MSI의 개체 ID입니다.
      
      msi_object_id 또는 msi_mi_res_id가 지정된 경우 비워두십시오.

   --msi-mi-res-id
      사용할 사용자 할당 MSI의 Azure 리소스 ID입니다.
      
      msi_client_id 또는 msi_object_id가 지정된 경우 비워두십시오.

   --use-emulator
      로컬 스토리지 에뮬레이터를 사용하려면 'true'로 제공합니다.
      
      실제 Azure 스토리지 엔드포인트를 사용하는 경우 비워두십시오.

   --endpoint
      서비스의 엔드포인트입니다.
      
      보통 비워두십시오.

   --upload-cutoff
      청크 업로드로 전환하기 위한 임계값 (<= 256 MiB) (사용되지 않음).

   --chunk-size
      업로드 청크 크기.
      
      이 값은 메모리에 저장되며 메모리에는
      "--transfers" * "--azureblob-upload-concurrency"개의 청크가 한 번에 저장될 수 있습니다.

   --upload-concurrency
      분할 업로드 작업의 동시성.
      
      동시에 업로드되는 작은 수의 큰 파일을 고속 링크로 업로드하는 경우 대역폭을 완전히 활용하지 못하는 경우 이 값을 높이면 전송 속도를 향상시킬 수 있습니다.
      
      테스트에서 업로드 속도는 업로드 동시성과 거의 선형으로 증가합니다. 예를 들어 기가비트 파이프를 채우려면 이 값을 64로 올려야 할 수도 있습니다. 메모리를 더 사용할 수 있습니다.
      
      청크가 메모리에 저장되며 메모리에는
      "--transfers" * "--azureblob-upload-concurrency"개의 청크가 한 번에 저장될 수 있습니다.

   --list-chunk
      블롭 목록의 크기.
      
      이 값은 각 목록 청크에서 요청되는 blob의 수입니다. 기본값은 최대 5000입니다. "블롭 목록" 요청에는 완료하는 데 1 메가바이트당 2분이 허용됩니다. 평균적으로 2분당 1MB 이상이 소요되는 작업은 시간 초과됩니다 (
      [출처](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval)
      ). 시간 초과를 피하려면 반환되는 blob 항목 수를 제한하는 데 사용할 수 있습니다.

   --access-tier
      blob의 액세스 티어: hot, cool 또는 archive.
      
      아카이브 티어 블롭은 hot 또는 cool로 액세스 티어를 설정하여 복원할 수 있습니다. 계정 수준에서 설정된 기본 액세스 티어를 사용할 계획인 경우 비워두십시오.
      
      "액세스 티어"가 지정되지 않은 경우 rclone은 어떤 티어도 적용하지 않습니다.
      rclone은 업로드 중에 "Set Tier" 작업을 수행합니다. 개체가 수정되지 않은 경우 "액세스 티어"를 새로운 티어로 지정하더라도 영향을 주지 않을 것입니다.
      원격에 "archive 티어"로 있는 경우 원격에서 데이터 전송 작업을 수행할 수 없습니다. 사용자는 먼저 "Hot" 또는 "Cool"로 blob을 teiring하여 복원해야 합니다.

   --archive-tier-delete
      덮어쓰기 전에 아카이브 티어 blob 삭제.
      
      아카이브 티어 블롭은 업데이트할 수 없습니다. 따라서 이 플래그 없이 아카이브 티어 블롭을 업데이트하려고 하면 rclone은 다음과 같은 오류를 반환합니다:
      
          can't update archive tier blob without --azureblob-archive-tier-delete
      
      이 플래그가 설정된 경우 아카이브 티어 블롭을 덮어쓰기 전에 기존 blob을 삭제한 후 새로운 blob을 업로드하기 전에 업로드합니다. 업로드가 실패하는 경우 데이터 손실 발생 가능성이 있으며 (일반 blob을 업데이트하는 것과 달리) 삭제하는 작업이 더 많은 비용이 발생할 수 있습니다.

   --disable-checksum
      객체 메타데이터에 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다. 데이터 무결성 검사에는 좋지만 큰 파일의 업로드를 시작하는 데 오랜 지연이 발생할 수 있습니다.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 플러시되는 빈도입니다.
      
      추가 버퍼를 필요로 하는 업로드(예: 분할)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 미사용 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --encoding
      배경 처리에 사용되는 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

   --public-access
      컨테이너의 공개 액세스 레벨: blob 또는 container.

      예시:
         | <unset>   | 컨테이너 및 해당 blob은 권한이 있는 요청으로만 액세스할 수 있습니다.
         |           | 이것이 기본 값입니다.
         | blob      | 이 컨테이너 내의 blob 데이터는 익명의 요청을 통해 읽을 수 있습니다.
         | container | 컨테이너 및 blob 데이터에 대해 완전한 공개 읽기 액세스를 허용합니다.

   --no-check-container
      컨테이너가 있는지 확인하거나 컨테이너를 생성하지 않으려면 설정합니다.
      
      컨테이너가 이미 존재하는 경우에는 rclone이 수행하는 트랜잭션 수를 최소화하려고 할 때 유용할 수 있습니다.
      

   --no-head-object
      객체를 가져올 때 GET 전에 HEAD를 수행하지 않습니다.


OPTIONS:
   --account value                      Azure Storage Account Name. [$ACCOUNT]
   --client-certificate-password value  인증서 파일의 암호입니다(선택 사항). [$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      PEM 또는 PKCS12 인증서 파일의 경로입니다. [$CLIENT_CERTIFICATE_PATH]
   --client-id value                    사용 중인 클라이언트의 ID입니다. [$CLIENT_ID]
   --client-secret value                서비스 프린시팔의 클라이언트 시크릿 중 하나 [$CLIENT_SECRET]
   --env-auth                           런타임(환경변수, CLI 또는 MSI)에서 자격 증명을 읽습니다. (기본값: false) [$ENV_AUTH]
   --help, -h                           도움말 표시
   --key value                          스토리지 계정 공유 키. [$KEY]
   --sas-url value                      컨테이너 수준 접근용 SAS URL. [$SAS_URL]
   --tenant value                       서비스 프린시팔의 테넌트 ID입니다. 또는 디렉터리 ID입니다. [$TENANT]

   고급

   --access-tier value              blob의 액세스 티어: hot, cool or archive. [$ACCESS_TIER]
   --archive-tier-delete            덮어쓰기 전에 아카이브 티어 blob 삭제. (기본값: false) [$ARCHIVE_TIER_DELETE]
   --chunk-size value               업로드 청크 크기 (기본값: "4Mi") [$CHUNK_SIZE]
   --client-send-certificate-chain  인증서 인증 시 인증 요청에 x5c 헤더를 포함할지 여부 (기본값: false) [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               객체 메타데이터에 MD5 체크섬을 저장하지 않음 (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 배경 처리에 사용되는 인코딩 (기본값: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$ENCODING]
   --endpoint value                 서비스의 엔드포인트 [$ENDPOINT]
   --list-chunk value               블롭 목록의 크기 (기본값: 5000) [$LIST_CHUNK]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 플러시되는 빈도 (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부 (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            사용할 사용자 할당 MSI의 개체 ID입니다. [$MSI_CLIENT_ID]
   --msi-mi-res-id value            사용할 사용자 할당 MSI의 Azure 리소스 ID입니다. [$MSI_MI_RES_ID]
   --msi-object-id value            사용할 사용자 할당 MSI의 개체 ID입니다. [$MSI_OBJECT_ID]
   --no-check-container             컨테이너가 있는지 확인하거나 컨테이너를 생성하지 않음 (기본값: false) [$NO_CHECK_CONTAINER]
   --no-head-object                 객체를 가져올 때 HEAD를 수행하지 않음 (기본값: false) [$NO_HEAD_OBJECT]
   --password value                 사용자의 비밀번호 [$PASSWORD]
   --public-access value            컨테이너의 공개 액세스 레벨: blob 또는 container [$PUBLIC_ACCESS]
   --service-principal-file value   서비스 프린시팔과 함께 사용할 자격 증명이 담긴 파일의 경로 [$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       분할 업로드 작업의 동시성 (기본값: 16) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하기 위한 임계값 (<= 256 MiB) (사용되지 않음) [$UPLOAD_CUTOFF]
   --use-emulator                   로컬 스토리지 에뮬레이터를 사용하려면 'true'로 제공 (기본값: false) [$USE_EMULATOR]
   --use-msi                        관리형 서비스 ID를 사용하여 인증(오직 Azure에서만 작동) (기본값: false) [$USE_MSI]
   --username value                 사용자 이름(일반적으로 이메일 주소) [$USERNAME]

   일반

   --name value  저장소의 이름(기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}