# Dropbox

{% code fullWidth="true" %}
```
이름:
   singularity datasource add dropbox - Dropbox

사용법:
   singularity datasource add dropbox [command options] <데이터셋_이름> <소스_경로>

설명:
   --dropbox-auth-url
      인증 서버 URL.
      
      기본값을 사용하려면 비워 두세요.

   --dropbox-batch-commit-timeout
      배치 완료까지 대기하는 최대 시간

   --dropbox-batch-mode
      파일 업로드 배치 동기화|비동기화|사용 안 함.
      
      이는 rclone에서 사용하는 배치 모드를 설정합니다.
      
      자세한 내용은 [본 문서](https://rclone.org/dropbox/#batch-mode)를 참조하세요.
      
      3가지 값이 있습니다.
      
      - off - 배치 사용 안 함
      - sync - 배치 업로드 및 완료 확인 (기본값)
      - async - 배치 업로드 및 완료 확인 안 함
      
      Rclone은 종료 시 보류 중인 배치를 닫을 것이며, 이로 인해 지연이 발생할 수 있습니다.
      

   --dropbox-batch-size
      업로드 배치에 포함되는 파일의 최대 개수.
      
      이는 업로드할 파일의 배치 크기를 설정합니다. 1000보다 작아야 합니다.
      
      기본값은 0으로, rclone이 batch_mode 설정에 따라 배치 크기를 자동으로 계산합니다.
      
      - batch_mode: async - 기본 배치 크기는 100입니다.
      - batch_mode: sync - 기본 배치 크기는 --transfers와 동일합니다.
      - batch_mode: off - 사용 안 함
      
      Rclone은 종료 시 보류 중인 배치를 닫을 것이며, 이로 인해 지연이 발생할 수 있습니다.
      
      이 옵션은 많은 작은 파일을 업로드하는 경우에 매우 유용합니다. --transfers 32를 사용하여 처리량을 극대화할 수 있습니다.
      

   --dropbox-batch-timeout
      업로드 배치가 비활성화된 상태로 대기한 최대 시간.
      
      업로드 배치가 이 시간 이상 비활성화된 경우 해당 배치가 업로드됩니다.
      
      기본값은 0으로, rclone은 사용 중인 batch_mode에 따라 합리적인 기본값을 선택합니다.
      
      - batch_mode: async - 기본 배치 시간은 500ms입니다.
      - batch_mode: sync - 기본 배치 시간은 10초입니다.
      - batch_mode: off - 사용 안 함
      

   --dropbox-chunk-size
      업로드 청크 크기 (< 150Mi).
      
      이보다 큰 파일은 이 크기의 청크로 업로드됩니다.
      
      청크는 메모리에 버퍼링되며 (한 번에 하나씩), rclone은 재시도를 처리할 수 있습니다. 크기를 크게 설정하면 속도가 약간 증가합니다 (테스트에서 최대 10%의 128 MiB)하지만 더 많은 메모리를 사용합니다. 메모리가 부족한 경우 더 작은 값으로 설정할 수 있습니다.

   --dropbox-client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워 두세요.

   --dropbox-client-secret
      OAuth 클라이언트 비밀.
      
      일반적으로 비워 두세요.

   --dropbox-encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --dropbox-impersonate
      비즈니스 계정을 사용할 때 이 사용자를 변장합니다.
      
      실제로 변장을 사용하려면 "rclone config"를 실행할 때이 플래그를 설정해야 합니다. 이렇게하면 rclone이 API에서 사용하는 내부 ID로 회원의 이메일 주소를 조회하는 데 필요한 "members.read" 범위를 요청합니다.
      
      "members.read" 범위를 사용하려면 OAuth 흐름 중 Dropbox 팀 관리자의 승인이 필요합니다.
      
      이 옵션을 사용하려면 자체 앱을 사용해야 합니다 (고유한 client_id와 client_secret을 설정함). 현재 rclone의 기본 권한 집합에 "members.read"가 포함되어 있지 않기 때문입니다. 이는 v1.55 이후의 버전이 universally 적용되면 추가할 수 있습니다.
      

   --dropbox-shared-files
      개별 공유 파일에서 작동하도록 rclone에 지시합니다.
      
      이 모드에서 rclone의 기능은 매우 제한됩니다. 이 모드에서는 목록 (ls, lsl 등) 작업과 읽기 작업 (예 : 다운로드) 만 지원됩니다. 이외의 모든 작업은 비활성화됩니다.

   --dropbox-shared-folders
      공유 폴더에서 작동하도록 rclone에 지시합니다.
            
      이 플래그를 경로 없이 사용하면 목록 작업만 지원되며 사용 가능한 모든 공유 폴더가 표시됩니다. 경로를 지정하면 첫 번째 부분이 공유 폴더의 이름으로 해석됩니다. 그런 다음 rclone은이 공유 폴더를 루트 네임 스페이스에 마운트하려고합니다. 공유 폴더가 성공적으로 마운트되면 rclone은 정상적으로 계속 진행됩니다. 이제 공유 폴더는 일반적인 폴더와 거의 동일하며 모든 일반 작업이 지원됩니다. 
      
      공유 폴더를 마운트 해제하지 않으므로 특정 공유 폴더의 첫 번째 사용 이후 --dropbox-shared-folders를 생략할 수 있습니다.

   --dropbox-token
      OAuth 액세스 토큰을 JSON blob 형식으로 입력하세요.

   --dropbox-token-url
      토큰 서버 URL.
      
      기본값을 사용하려면 비워 두세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] CAR 파일로 데이터셋을 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공한 스캔으로부터 이 시간이 지나면 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화됨)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비 완료)

   dropbox에 대한 옵션

   --dropbox-auth-url value              인증 서버 URL. [$DROPBOX_AUTH_URL]
   --dropbox-batch-commit-timeout value  배치 완료를 대기하는 최대 시간 (기본값: "10m0s") [$DROPBOX_BATCH_COMMIT_TIMEOUT]
   --dropbox-batch-mode value            파일 업로드 배치 동기화|비동기화. (기본값: "sync") [$DROPBOX_BATCH_MODE]
   --dropbox-batch-size value            업로드 배치에 포함되는 파일의 최대 개수. (기본값: "0") [$DROPBOX_BATCH_SIZE]
   --dropbox-batch-timeout value         업로드 전 비활성화된 대기 시간. (기본값: "0s") [$DROPBOX_BATCH_TIMEOUT]
   --dropbox-chunk-size value            업로드 청크 크기 (< 150Mi). (기본값: "48Mi") [$DROPBOX_CHUNK_SIZE]
   --dropbox-client-id value             OAuth 클라이언트 ID. [$DROPBOX_CLIENT_ID]
   --dropbox-client-secret value         OAuth 클라이언트 비밀. [$DROPBOX_CLIENT_SECRET]
   --dropbox-encoding value              백엔드의 인코딩. (기본값: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$DROPBOX_ENCODING]
   --dropbox-impersonate value           비즈니스 계정을 사용할 때 이 사용자를 변장합니다. [$DROPBOX_IMPERSONATE]
   --dropbox-shared-files value          개별 공유 파일에서 작업하도록 rclone에 지시합니다. (기본값: "false") [$DROPBOX_SHARED_FILES]
   --dropbox-shared-folders value        공유 폴더에서 작업하도록 rclone에 지시합니다. (기본값: "false") [$DROPBOX_SHARED_FOLDERS]
   --dropbox-token value                 OAuth 액세스 토큰을 JSON blob 형식으로 입력하세요. [$DROPBOX_TOKEN]
   --dropbox-token-url value             토큰 서버 URL. [$DROPBOX_TOKEN_URL]

```
{% endcode %}