# Dropbox

{% code lineWidth="true" %}
```
이름:
   singularity storage update dropbox - Dropbox

사용법:
   singularity storage update dropbox [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 아이디.
      
      보통 빈칸으로 남겨두십시오.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      보통 빈칸으로 남겨두십시오.

   --token
      OAuth 액세스 토큰(JSON blob).

   --auth-url
      인증 서버 URL.
      
      공급자 기본값을 사용하려면 빈칸으로 남겨두십시오.

   --token-url
      토큰 서버 URL.
      
      공급자 기본값을 사용하려면 빈칸으로 남겨두십시오.

   --chunk-size
      업로드 청크 크기 (< 150Mi).
      
      이보다 큰 파일은 이 크기의 청크로 업로드됩니다.
      
      청크는 메모리에 버퍼링됩니다(한 번에 하나씩), 그래서 rclone이 재시도를 다룰 수 있습니다. 이 크기를 늘리면 속도가 약간(테스트에서 128MiB에 대해서 최대 10%) 증가하지만 메모리를 더 사용합니다. 메모리가 부족한 경우 더 작게 설정할 수 있습니다.

   --impersonate
      비즈니스 계정을 사용할 때 이 사용자를 가장합니다.
      
      가장하기를 사용하려면 "rclone config"를 실행할 때 이 플래그가 설정되어 있는지 확인해야 합니다. 이렇게 하면 rclone이 일반적으로 수행하지 않는 "members.read" 범위를 요청하게 됩니다. 이는 dropbox이 API에서 멤버 이메일 주소를 내부 ID로 변환하는 데 사용됩니다.
      
      "members.read" 범위를 사용하려면 OAuth 플로우 중에 Dropbox의 팀 관리자에게 승인을 요청해야 합니다.
      
      이 옵션을 사용하려면 자체적인 앱(고유한 client_id와 client_secret을 설정)을 사용해야 합니다. 현재 rclone의 기본적인 권한 집합에 "members.read"가 포함되어 있지 않기 때문입니다. v1.55 이상이 모든 곳에서 사용되면 추가할 수 있습니다.

   --shared-files
      rclone이 개별 공유 파일에서 작동하도록 지시합니다.
      
      이 모드에서는 rclone의 기능이 극도로 제한됩니다. 이 모드에서는 목록(ls, lsl 등) 및 읽기 작업(예: 다운로드)만 지원됩니다. 모든 다른 작업은 비활성화됩니다.

   --shared-folders
      rclone이 공유 폴더에서 작업하도록 지시합니다.
            
      이 플래그를 사용하면 경로를 지정하지 않은 경우 목록 작업만 지원되며 모든 사용 가능한 공유 폴더가 나열됩니다. 경로를 지정하면 첫 번째 부분은 공유 폴더의 이름으로 해석됩니다. 그런 다음 rclone은 이 공유 폴더를 루트 네임스페이스에 마운트하려고 시도합니다. 공유 폴더가 성공하면 rclone은 정상적으로진행합니다. 
      
      공유 폴더는 이제 거의 일반 폴더와 같으며 모든 일반 작업이 지원됩니다. 
      
      참고로, 이후에도 공유 폴더를 언마운트하지 않으므로 특정 공유 폴더를 처음 사용한 후에는 --dropbox-shared-folders를 생략할 수 있습니다.

   --batch-mode
      파일 업로드 배치 동기화|비동기화|비활성화.
      
      이는 rclone이 사용하는 배치 모드를 설정합니다.
      
      자세한 내용은 [메인 문서](https://rclone.org/dropbox/#batch-mode)를 참조하십시오.
      
      가능한 값은 다음과 같습니다.
      
      - off - 배치 없음
      - sync - 배치 업로드 및 완료 확인(기본값)
      - async - 배치 업로드 및 완료 확인 미필요
      
      Rclone은 종료 시 앞서 정리되지 않은 배치를 닫을 수 있으며, 이로 인해 종료에 지연이 발생할 수 있습니다.

   --batch-size
      업로드 배치에 포함될 파일의 최대 수.
      
      이는 업로드할 파일의 배치 크기를 설정합니다. 1000 이하이어야 합니다.
      
      기본값은 0으로, rclone은 배치 모드 설정에 따라 배치 크기를 계산합니다.
      
      - batch_mode: async - 기본 배치 크기는 100입니다.
      - batch_mode: sync - 기본 배치 크기는 --transfers와 같습니다.
      - batch_mode: off - 사용 안 함
      
      Rclone은 종료 시 앞서 정리되지 않은 배치를 닫을 수 있으며, 이로 인해 종료에 지연이 발생할 수 있습니다.
      
      많은 작은 파일을 업로드하는 경우 이를 설정하는 것이 좋습니다. 그렇게 하면 빠르게 업로드할 수 있습니다. --transfers 32를 사용하여 처리량을 최대화할 수 있습니다.

   --batch-timeout
      업로드할 것이 없는 배치가 있는 경우 업로드 전에 배치를 허용할 수 있는 최대 시간.
      
      업로드 배치가 지정된 시간 동안 비활동한 경우, 배치가 업로드됩니다.
      
      이 값의 기본값은 0으로, rclone은 사용 중인 배치 모드에 따라 합리적인 기본값을 선택합니다.
      
      - batch_mode: async - 기본 배치 시간 초과는 500ms입니다.
      - batch_mode: sync - 기본 배치 시간 초과는 10초입니다.
      - batch_mode: off - 사용 안 함

   --batch-commit-timeout
      배치 처리의 완료를 기다리는 최대 시간

   --encoding
      백엔드에 대한 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 아이디. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 시크릿. [$CLIENT_SECRET]
   --help, -h             도움말 보기

   고급

   --auth-url value              인증 서버 URL. [$AUTH_URL]
   --batch-commit-timeout value  배치 처리가 완료될 때까지 기다리는 최대 시간(기본값: "10m0s") [$BATCH_COMMIT_TIMEOUT]
   --batch-mode value            파일 업로드 배치 동기화|비동기화|비활성화(기본값: "sync") [$BATCH_MODE]
   --batch-size value            업로드 배치에 포함될 파일의 최대 수(기본값: 0) [$BATCH_SIZE]
   --batch-timeout value         업로드할 것이 없는 배치가 있는 경우 업로드 전에 배치를 허용할 수 있는 최대 시간(기본값: "0s") [$BATCH_TIMEOUT]
   --chunk-size value            업로드 청크 크기 (< 150Mi)(기본값: "48Mi") [$CHUNK_SIZE]
   --encoding value              백엔드에 대한 인코딩(기본값: "Slash,BackSlash,Del,RightSpace,InvalidUtf8,Dot") [$ENCODING]
   --impersonate value           비즈니스 계정을 사용할 때 이 사용자를 가장합니다. [$IMPERSONATE]
   --shared-files                rclone이 개별 공유 파일에서 작동하도록 지시합니다(기본값: false) [$SHARED_FILES]
   --shared-folders              rclone이 공유 폴더에서 작동하도록 지시합니다(기본값: false) [$SHARED_FOLDERS]
   --token value                 OAuth 액세스 토큰(JSON blob) [$TOKEN]
   --token-url value             토큰 서버 URL [$TOKEN_URL]

```
{% endcode %}