# Microsoft Azure Blob Storage

{% code fullWidth="true" %}
```
이름:
   singularity storage update azureblob - Microsoft Azure Blob Storage

사용법:
   singularity storage update azureblob [명령 옵션] <이름|아이디>

설명:
   --account
      Azure 저장소 계정 이름입니다.
      
      사용 중인 Azure 저장소 계정 이름을 설정합니다.
      
      사용하지 않는 경우 SAS URL 또는 에뮬레이터를 사용하고,
      그렇지 않은 경우 설정해야 합니다.
      
      이 값을 비워 두고 env_auth가 설정되어 있다면 환경 변수 `AZURE_STORAGE_ACCOUNT_NAME`에서 읽어옵니다.
      

   --env-auth
      실행 시간(Runtim)에서 인증 정보를 읽습니다.
      
      자세한 내용은 [인증 문서](/azureblob#authentication)를 참조하세요.

   --key
      스토리지 계정 공유 키입니다.
      
      SAS URL 또는 에뮬레이터를 사용하지 않으려면 비워 두세요.

   --sas-url
      컨테이너 수준 액세스에 대한 SAS URL입니다.
      
      계정/키 또는 에뮬레이터를 사용하는 경우 비워 둡니다.

   --tenant
      서비스 주체의 테넌트 ID입니다. (디렉터리 ID라고도 함)
      
      다음과 같은 경우 설정합니다.
      - 클라이언트 비밀키로 서비스 주체 사용
      - 인증서로 서비스 주체 사용
      - 사용자 아이디와 비밀번호를 사용하는 사용자

   --client-id
      사용 중인 클라이언트의 ID입니다.
      
      다음과 같은 경우 설정합니다.
      - 클라이언트 비밀키로 서비스 주체 사용
      - 인증서로 서비스 주체 사용
      - 사용자 아이디와 비밀번호를 사용하는 사용자

   --client-secret
      서비스 주체의 클라이언트 비밀키 중 하나입니다.
      
      다음과 같은 경우 설정합니다.
      - 클라이언트 비밀키로 서비스 주체 사용

   --client-certificate-path
      개인 키를 포함한 PEM 또는 PKCS12 인증서 파일의 경로입니다.
      
      다음과 같은 경우 설정합니다.
      - 인증서로 서비스 주체 사용

   --client-certificate-password
      인증서 파일의 비밀번호입니다. (선택 사항)
      
      다음과 같은 경우 설정합니다.
      - 인증서로 서비스 주체 사용
      
      인증서에 비밀번호가 있는 경우 선택적으로 설정합니다.

   --client-send-certificate-chain
      인증서 인증을 사용할 때 인증 요청에 x5c 헤더를 포함할지 여부를 지정합니다.
      
      true로 설정하면 인증 요청에 x5c 헤더가 포함됩니다.
      
      다음과 같은 경우 선택적으로 설정합니다.
      - 인증서로 서비스 주체 사용

   --username
      사용자 이름(일반적으로 이메일 주소)
      
      다음과 같은 경우 설정합니다.
      - 사용자 아이디와 비밀번호를 사용하는 사용자

   --password
      사용자의 비밀번호
      
      다음과 같은 경우 설정합니다.
      - 사용자 아이디와 비밀번호를 사용하는 사용자

   --service-principal-file
      서비스 주체를 사용하기 위한 자격 증명이 있는 파일의 경로입니다.
      
      일반적으로 비워 둡니다. 대화식 로그인 대신 서비스 주체를 사용하려는 경우에만 필요합니다.
      
          $ az ad sp create-for-rbac --name "<이름>" \
            --role "Storage Blob Data Owner" \
            --scopes "/subscriptions/<구독>/resourceGroups/<리소스 그룹>/providers/Microsoft.Storage/storageAccounts/<스토리지 계정>/blobServices/default/containers/<컨테이너>" \
            > azure-principal.json
      
      자세한 내용은 ["Azure 서비스 주체 생성"](https://docs.microsoft.com/ko-kr/cli/azure/create-an-azure-service-principal-azure-cli) 및 ["Blob 데이터에 대한 액세스를 위한 Azure 역할 할당"](https://docs.microsoft.com/ko-kr/azure/storage/common/storage-auth-aad-rbac-cli) 페이지를 참조하세요.
      
      `client_id`, `tenant`, `client_secret` 키를 설정하는 대신 자격 증명을 직접 rclone 설정 파일에 넣는 것이 더 편리할 수 있습니다.
      

   --use-msi
      관리되는 서비스 ID를 사용하여 인증(평의원) (Azure에서만 작동).
      
      true인 경우 [관리되는 서비스 ID](https://docs.microsoft.com/ko-kr/azure/active-directory/managed-identities-azure-resources/)를 사용하여
      SAS 토큰이나 계정 키 대신 Azure Storage에 인증합니다.
      
      이 프로그램이 실행되는 VM(SS)에 시스템 할당 ID가 있는 경우 기본적으로 사용됩니다. 
      시스템 할당이 없는 리소스에 정확히 하나의 사용자 할당 ID가 있는 경우 사용자 할당 ID가 기본적으로 사용됩니다. 
      여러 사용자 할당 ID가 있는 경우에는 msi_object_id, msi_client_id, msi_mi_res_id 중 하나만 사용하여 명시적으로 사용할 ID를 지정해야 합니다.

   --msi-object-id
      필요한 경우 사용할 할당된 사용자용 MSI의 개체 ID입니다.
      
      msi_client_id 또는 msi_mi_res_id가 지정된 경우 비워 두세요.

   --msi-client-id
      필요한 경우 사용할 할당된 사용자용 MSI의 개체 ID입니다.
      
      msi_object_id 또는 msi_mi_res_id가 지정된 경우 비워 두세요.

   --msi-mi-res-id
      필요한 경우 사용할 할당된 사용자용 MSI의 Azure 리소스 ID입니다.
      
      msi_client_id 또는 msi_object_id가 지정된 경우 비워 두세요.

   --use-emulator
      로컬 스토리지 에뮬레이터를 사용하려면 'true'로 설정합니다.
      
      실제 Azure 스토리지 엔드포인트를 사용하는 경우 비워 두세요.

   --endpoint
      서비스의 엔드포인트입니다.
      
      일반적으로 비워 둡니다.

   --upload-cutoff
      청크 업로드로 전환하는 데 사용되는 임계값(<= 256 MiB) (비권장).

   --chunk-size
      업로드 청크 크기입니다.
      
      이 값은 메모리에 저장되며, 메모리에는
      "--transfers" * "--azureblob-upload-concurrency" 만큼의 청크가 저장될 수 있습니다.

   --upload-concurrency
      다중 파트 업로드에 대한 동시성입니다.
      
      동시에 업로드되는 동일한 파일의 청크 수입니다.
      
      고속 링크로 대량의 큰 파일을 업로드하고
      이러한 업로드가 대역폭을 완전히 활용하지 못하는 경우,
      이 값을 높이면 전송 속도가 향상될 수 있습니다.
      
      테스트에서 업로드 속도는 거의 선형적으로 증가합니다.
      예를 들어 1기가비트 파이프를 채우려면 64로 이 값을 높여야 할 수 있습니다.
      이렇게 하면 더 많은 메모리가 사용됩니다.
      
      청크가 메모리에 저장되며, 메모리에는
      "--transfers" * "--azureblob-upload-concurrency" 만큼의 청크가 저장될 수 있습니다.

   --list-chunk
      블롭 목록의 크기입니다.
      
      이 값은 각 목록 청크에서 요청하는 블롭 수를 설정합니다. 기본값은 최대값인 5000입니다. 
      ["목록 블롭" 요청은 1MB당 2분의 시간이 허용됩니다](https://docs.microsoft.com/en-us/rest/api/storageservices/setting-timeouts-for-blob-service-operations#exceptions-to-default-timeout-interval).
      평균 1MB당 2분 이상 소요되는 작업은 시간이 초과됩니다.
      이 값을 사용하여 반환할 블롭 항목 수를 제한하여 시간 초과를 피할 수 있습니다.

   --access-tier
      블롭의 액세스 계층: hot, cool 또는 archive입니다.
      
      보관된 블롭은 액세스 계층을 hot 또는 cool로 설정하여 복원할 수 있습니다.
      액세스 계층을 지정하지 않으면 rclone은 모든 계층에 적용하지 않습니다.
      rclone은 업로드 중 "Set Tier" 작업을 수행하며, 객체가 수정되지 않으면 새 액세스 계층에 대한 영향을 주지 않습니다.
      원격지의 블롭이 "archive 계층"에 있는 경우 원격지에서 데이터 전송 작업을 수행할 수 없습니다.
      먼저 블롭을 "Hot" 또는 "Cool" 계층으로 전환하여 복원해야 합니다.

   --archive-tier-delete
      덮어쓰기 전에 보관된 계층의 블롭을 삭제합니다.
      
      보관된 계층 블롭은 업데이트할 수 없습니다. 이 플래그 없이 보관된 계층 블롭을 업데이트하려고 하면,
      rclone은 다음과 같은 오류를 발생시킵니다:
      
          --azureblob-archive-tier-delete없이 보관된 상위 블롭을 업데이트할 수 없음
      
      해당 플래그를 설정하면 rclone이 보관된 계층 블롭을 덮어쓰기 전에 기존 블롭을 삭제한 다음 대체 블롭을 업로드합니다.
      업로드에 실패하면(일반적인 블롭 업데이트와 달리) 데이터 손실이 발생할 수 있으며,
      미리 보관된 계층 블롭을 삭제하는 것은 추가 비용이 발생할 수 있습니다.
      

   --disable-checksum
      객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다.
      
      일반적으로 rclone은 객체를 업로드하기 전에 입력의 MD5 체크섬을 계산하여 객체의 메타데이터에 추가합니다.
      이는 데이터 무결성 확인에 좋지만 큰 파일의 업로드 시작에 오랜 시간이 소요될 수 있습니다.

   --memory-pool-flush-time
      내부 메모리 버퍼 풀이 얼마나 자주 플러시되는지 제어합니다.
      
      추가 버퍼가 필요한 업로드(예: 멀티파트)는 할당을 위해 메모리 풀을 사용합니다.
      이 옵션은 사용하지 않는 버퍼가 풀에서 제거되는 빈도를 제어합니다.

   --memory-pool-use-mmap
      내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --public-access
      컨테이너의 공개 액세스 수준: blob 또는 container.

      예:
         | <unset>   | 컨테이너와 해당 블롭은 인증된 요청으로만 액세스할 수 있습니다.
         |           | 기본값입니다.
         | blob      | 이 컨테이너 내의 블롭 데이터는 익명 요청을 통해 읽을 수 있습니다.
         | container | 컨테이너와 블롭 데이터에 대한 완전한 공개 읽기 액세스를 허용합니다.

   --no-check-container
      컨테이너가 존재하는지 확인하거나 생성하지 않도록 설정합니다.
      
      컨테이너가 이미 존재하는 경우 rclone이 수행하는 트랜잭션 수를 최소화하려는 경우 유용할 수 있습니다.
      

   --no-head-object
      GET 작업 수행 시 HEAD를 수행하지 않습니다.


옵션:
   --account value                      Azure 저장소 계정 이름입니다. [$ACCOUNT]
   --client-certificate-password value  인증서 파일의 비밀번호입니다. (선택 사항) [$CLIENT_CERTIFICATE_PASSWORD]
   --client-certificate-path value      개인 키를 포함한 PEM 또는 PKCS12 인증서 파일의 경로입니다. [$CLIENT_CERTIFICATE_PATH]
   --client-id value                    사용 중인 클라이언트의 ID입니다. [$CLIENT_ID]
   --client-secret value                서비스 주체의 클라이언트 비밀키 중 하나입니다 [$CLIENT_SECRET]
   --env-auth                           실행 시간(Runtime)에서 인증 정보를 읽습니다. (기본값: false) [$ENV_AUTH]
   --help, -h                           도움말 표시
   --key value                          스토리지 계정 공유 키입니다. [$KEY]
   --sas-url value                      컨테이너 수준 액세스에 대한 SAS URL입니다. [$SAS_URL]
   --tenant value                       서비스 주체의 테넌트 ID입니다. (디렉터리 ID라고도 함) [$TENANT]

   고급 옵션

   --access-tier value              블롭의 액세스 계층: hot, cool 또는 archive입니다. [$ACCESS_TIER]
   --archive-tier-delete            덮어쓰기 전에 보관된 계층의 블롭을 삭제합니다. (기본값: false) [$ARCHIVE_TIER_DELETE]
   --chunk-size value               업로드 청크 크기입니다. (기본값: "4Mi") [$CHUNK_SIZE]
   --client-send-certificate-chain  인증 요청에 인증서 체인을 포함할지 여부입니다. (기본값: false) [$CLIENT_SEND_CERTIFICATE_CHAIN]
   --disable-checksum               객체 메타데이터와 함께 MD5 체크섬을 저장하지 않습니다. (기본값: false) [$DISABLE_CHECKSUM]
   --encoding value                 백엔드의 인코딩입니다. (기본값: "Slash,BackSlash,Del,Ctl,RightPeriod,InvalidUtf8") [$ENCODING]
   --endpoint value                 서비스의 엔드포인트입니다. [$ENDPOINT]
   --list-chunk value               블롭 목록의 크기입니다. (기본값: 5000) [$LIST_CHUNK]
   --memory-pool-flush-time value   내부 메모리 버퍼 풀이 얼마나 자주 플러시되는지 제어합니다. (기본값: "1m0s") [$MEMORY_POOL_FLUSH_TIME]
   --memory-pool-use-mmap           내부 메모리 풀에서 mmap 버퍼를 사용할지 여부입니다. (기본값: false) [$MEMORY_POOL_USE_MMAP]
   --msi-client-id value            필요한 경우 사용할 할당된 사용자용 MSI의 개체 ID입니다. [$MSI_CLIENT_ID]
   --msi-mi-res-id value            필요한 경우 사용할 할당된 사용자용 MSI의 Azure 리소스 ID입니다. [$MSI_MI_RES_ID]
   --msi-object-id value            필요한 경우 사용할 할당된 사용자용 MSI의 개체 ID입니다. [$MSI_OBJECT_ID]
   --no-check-container             컨테이너가 존재하는지 확인하거나 생성하지 않도록 설정합니다. (기본값: false) [$NO_CHECK_CONTAINER]
   --no-head-object                 GET 작업 수행 시 HEAD를 수행하지 않습니다. (기본값: false) [$NO_HEAD_OBJECT]
   --password value                 사용자의 비밀번호 [$PASSWORD]
   --public-access value            컨테이너의 공개 액세스 수준: blob 또는 container. [$PUBLIC_ACCESS]
   --service-principal-file value   서비스 주체를 사용하기 위한 자격 증명이 있는 파일의 경로입니다. [$SERVICE_PRINCIPAL_FILE]
   --upload-concurrency value       다중 파트 업로드에 대한 동시성입니다. (기본값: 16) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드로 전환하는 데 사용되는 임계값(<= 256 MiB) (비권장). [$UPLOAD_CUTOFF]
   --use-emulator                   로컬 스토리지 에뮬레이터를 사용하려면 'true'로 설정합니다. (기본값: false) [$USE_EMULATOR]
   --use-msi                        관리되는 서비스 ID를 사용하여 인증(평의원) (Azure에서만 작동). (기본값: false) [$USE_MSI]
   --username value                 사용자 이름(일반적으로 이메일 주소) [$USERNAME]

```
{% endcode %}