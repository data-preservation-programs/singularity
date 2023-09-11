# Dropbox

{% code fullWidth="true" %}
```
이름:
   singularity storage create dropbox - Dropbox

사용법:
   singularity storage create dropbox [명령 옵션] [인수...]

설명:
   --client-id
      OAuth 클라이언트 아이디.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      보통 비워둡니다.

   --token
      JSON blob으로 된 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      기본 공급자 값을 사용하려면 비워 둡니다.

   --token-url
      토큰 서버 URL.
      
      기본 공급자 값을 사용하려면 비워 둡니다.

   --chunk-size
      업로드 청크 크기 (< 150Mi).
      
      이보다 큰 파일은 이 크기의 청크로 업로드됩니다.
      
      청크는 (한 번에 하나씩) 메모리에 버퍼링되며 rclone이 다시 시도할 수 있도록 합니다. 이 값을 크게 설정하면 속도가 약간 향상될 수 있지만(테스트에서 최대 10% 이상인 128 MiB의 속도 향상), 메모리를 더 사용하게 됩니다. 메모리가 부족한 경우 이 값을 작게 설정할 수도 있습니다.

   --impersonate
      비즈니스 계정을 사용할 때 이 사용자를 가장하십시오.
      
      가장하려면 "rclone config"를 실행할 때 이 플래그가 설정되어 있다는 것을 확인해야 합니다. 이렇게 하면 rclone이 보통 요청하지 않는 "members.read" 범위를 요청합니다. 이것은 드롭박스 API에서 구성원의 이메일 주소를 내부 ID로 조회하기 위해 필요합니다.
      
      "members.read" 범위를 사용하려면 OAuth 플로우 중에 드롭박스 팀 관리자의 승인이 필요합니다.
      
      이 옵션을 사용하려면 자체 앱(고유한 클라이언트 아이디와 클라이언트 시크릿을 설정함)을 사용해야 합니다. 현재 rclone의 기본 권한 집합에 "members.read"가 포함되어 있지 않으므로 이 옵션은 v1.55 이상이 모두 사용될 때 추가될 수 있습니다.

   --shared-files
      개별 공유 파일에서 작동하도록 rclone에 지시합니다.
      
      이 모드에서는 rclone의 기능이 매우 제한됩니다. 목록(lsl, ls 등) 및 읽기 작업(다운로드 등)만이 이 모드에서 지원됩니다. 이 모드에서는 모든 다른 작업이 비활성화됩니다.

   --shared-folders
      공유 폴더에서 작동하도록 rclone에 지시합니다.
            
      이 플래그를 사용하고 경로를 지정하지 않으면 목록 작업만 지원되며 모든 사용 가능한 공유 폴더가 나열됩니다. 경로를 지정하면 첫 번째 부분이 공유 폴더의 이름으로 해석됩니다. 그런 다음 rclone은 이 공유 폴더를 루트 이름 공간에 마운트하려고 시도합니다. 공유 폴더가 성공적으로 마운트되면 rclone은 정상적으로 진행됩니다. 이제 공유 폴더는 일반 폴더와 거의 동일하며 모든 일반 작업이 지원됩니다. 

      공유 폴더를 끝낸 후에는 공유 폴더를 마운트 해제하지 않으므로 특정 공유 폴더를 처음 사용한 후에는 --shared-folders를 생략할 수 있습니다.

   --batch-mode
      파일 업로드 배치 동기|비동기|사용 안 함.
      
      rclone에 의해 사용되는 배치 모드를 설정합니다.
      
      자세한 내용은 [메인 문서](https://rclone.org/dropbox/#batch-mode)를 참조하세요.
      
      이에는 3가지 가능한 값이 있습니다.
      
      - off - 배치 사용 안 함
      - sync - 배치 업로드 및 완료 확인(기본값)
      - async - 배치 업로드 및 완료 확인 안 함
      
      Rclone은 종료될 때 미처 완료되지 않은 배치를 모두 종료합니다. 이로 인해 종료에 지연이 발생할 수 있습니다.

   --batch-size
      업로드 배치에 포함되는 파일의 최대 개수.
      
      이 값을 설정하여 업로드할 파일의 배치 크기를 지정합니다. 이 값은 1000보다 작아야 합니다.
      
      기본값은 0이며, 이는 배치 모드 설정에 따라 배치 크기를 rclone이 계산하도록 합니다.
      
      - 배치 모드: async - 기본 배치 크기는 100입니다.
      - 배치 모드: sync - 기본 배치 크기는 --transfers와 같습니다.
      - 배치 모드: off - 사용 안 함
      
      Rclone은 종료될 때 미처 완료되지 않은 배치를 모두 종료합니다. 이로 인해 종료에 지연이 발생할 수 있습니다.
      
      많은 수의 작은 파일을 업로드하는 경우 이 값을 설정하는 것이 좋습니다. 이렇게 하면 작업이 훨씬 빠르게 처리됩니다. --transfers 32를 사용하여 처리량을 극대화할 수 있습니다.

   --batch-timeout
      업로드 전 대기 중인 배치의 최대 시간.
      
      이 값보다 더 오랜 시간 동안 업로드 배치가 대기하면 업로드됩니다.
      
      기본값은 0이며, 이는 rclone이 사용 중인 배치 모드를 기반으로 합리적인 기본값을 선택합니다.
      
      - 배치 모드: async - 기본 배치 대기 시간은 500ms입니다.
      - 배치 모드: sync - 기본 배치 대기 시간은 10s입니다.
      - 배치 모드: off - 사용 안 함

   --batch-commit-timeout
      배치 완료를 기다릴 최대 시간

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --client-id value      OAuth 클라이언트 아이디. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 시크릿. [$CLIENT_SECRET]
   --help, -h             도움말 보기

   고급

   --auth-url value              인증 서버 URL. [$AUTH_URL]
   --batch-commit-timeout value  배치 완료를 기다릴 최대 시간 (기본값: "10m0s") [$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            파일 업로드 배치 동기|비동기|사용 안 함. (기본값: "sync") [$BATCH_MODE]
   --batch-size value            업로드 배치에 포함되는 파일의 최대 개수. (기본값: 0) [$BATCH_SIZE]
   --batch-timeout value         업로드 전 대기 중인 배치의 최대 시간. (기본값: "0s") [$BATCH_TIMEOUT]
   --chunk-size value            업로드 청크 크기 (< 150Mi). (기본값: "48Mi") [$CHUNK_SIZE]
   --encoding value              백엔드의 인코딩. (기본값: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --impersonate value           비즈니스 계정을 사용할 때 이 사용자를 가장하십시오. [$IMPERSONATE]
   --shared-files                개별 공유 파일에서 작업하도록 rclone에 지시합니다. (기본값: false) [$SHARED_FILES]
   --shared-folders              공유 폴더에서 작업하도록 rclone에 지시합니다. (기본값: false) [$SHARED_FOLDERS]
   --token value                 JSON blob으로 된 OAuth 액세스 토큰. [$TOKEN]
   --token-url value             토큰 서버 URL. [$TOKEN_URL]

   일반

   --name value  스토리지의 이름 (기본값: Auto generated)
   --path value  스토리지의 경로

```
{% endcode %}