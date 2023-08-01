# HiDrive

{% code fullWidth="true" %}
```
이름:
   singularity datasource add hidrive - HiDrive

사용법:
   singularity datasource add hidrive [command options] <dataset_name> <source_path>

설명:
   --hidrive-auth-url
      인증 서버 URL.
      
      제공자 기본값을 사용하려면 비워 두세요.

   --hidrive-chunk-size
      청크 업로드용 청크 크기.
      
      설정된 기준점 이상의 크기의 파일(또는 알려지지 않은 크기의 파일)은 이 크기의 청크로 업로드됩니다.
      
      이 값의 상한은 2147483647바이트(약 2.000GB)입니다.
      이는 단일 업로드 작업에서 지원하는 최대 바이트 수입니다.
      이 값을 상한선보다 크게 설정하거나 음수로 설정하면 업로드가 실패합니다.
      
      이 값을 크게 설정하면 업로드 속도가 빨라질 수 있지만 메모리 사용량이 증가합니다.
      메모리 절약을 위해 이 값을 작게 설정할 수도 있습니다.

   --hidrive-client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워 두세요.

   --hidrive-client-secret
      OAuth 클라이언트 비밀.
      
      일반적으로 비워 두세요.

   --hidrive-disable-fetching-member-count
      디렉토리의 객체 수를 가져올 필요가 없는 경우 가져오지 않습니다.
      
      객체 수를 가져오지 않으면 요청이 더 빨라질 수 있습니다.

   --hidrive-encoding
      백엔드용 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --hidrive-endpoint
      서비스의 엔드포인트.
      
      API 호출이 이 URL로 이루어집니다.

   --hidrive-root-prefix
      모든 경로의 루트/상위 폴더입니다.
      
      지정된 폴더를 모든 원격 경로의 시작점으로 사용하려면 값을 입력하세요.
      이렇게 하면 rclone이 시작 위치로 어떤 폴더든 사용할 수 있습니다.
      
      예:
         | /       | rclone에서 액세스할 수 있는 가장 상위 디렉토리입니다.
                   | 일반 HiDrive 사용자 계정을 사용하는 경우 이것은 "root"와 동등합니다.
         | root    | HiDrive 사용자 계정의 가장 상위 디렉토리
         | <unset> | 이것은 경로에 대한 루트 접두사가 없음을 지정합니다.
                   | 이 경우 "remote:/path/to/dir" 또는 "remote:root/path/to/dir"와 같은 유효한 상위 폴더가 포함된 경로를 항상 지정해야 합니다.

   --hidrive-scope-access
      HiDrive에 액세스할 때 rclone이 사용할 액세스 권한.

      예:
         | rw | 리소스의 읽기 및 쓰기 액세스권한.
         | ro | 리소스의 읽기 전용 액세스권한.

   --hidrive-scope-role
      HiDrive에 액세스할 때 rclone이 사용할 사용자 수준.

      예:
         | user  | 관리 권한에 대한 사용자 수준 액세스권한.
                 | 대부분의 경우에 충분합니다.
         | admin | 관리 권한에 대한 광범위한 액세스권한.
         | owner | 관리 권한에 대한 전체 액세스권한.

   --hidrive-token
      OAuth 액세스 토큰(JSON 형식)입니다.

   --hidrive-token-url
      토큰 서버 URL.
      
      제공자 기본값을 사용하려면 비워 두세요.

   --hidrive-upload-concurrency
      청크 업로드용 동시성.
      
      동일한 파일에 대해 동시에 실행되는 전송 수의 상한입니다.
      1보다 작은 값을 설정하면 업로드가 데드락에 걸릴 수 있습니다.
      
      고속 링크를 통해 대량의 큰 파일을 업로드하면서 대역폭을 완전히 활용하지 못하는 경우
      이 값을 높이면 전송 속도를 높일 수도 있습니다.

   --hidrive-upload-cutoff
      청크 업로드를 위한 임계값.
      
      이보다 큰 파일은 설정된 청크 크기의 청크로 업로드됩니다.
      
      이 값의 상한은 2147483647바이트(약 2.000GB)입니다.
      단일 업로드 작업에서 지원하는 최대 바이트 수입니다.
      이 값을 상한 이상으로 설정하면 업로드가 실패합니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터셋을 CAR 파일로 내보낸 후 데이터 파일 삭제.  (기본값: false)
   --rescan-interval value  마지막 스캔 이후로 일정한 시간이 지나면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화됨)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: ready)

   HiDrive 옵션

   --hidrive-auth-url value                       인증 서버 URL. [$HIDRIVE_AUTH_URL]
   --hidrive-chunk-size value                     청크 업로드용 청크 크기. (기본값: "48Mi") [$HIDRIVE_CHUNK_SIZE]
   --hidrive-client-id value                      OAuth 클라이언트 ID. [$HIDRIVE_CLIENT_ID]
   --hidrive-client-secret value                  OAuth 클라이언트 비밀. [$HIDRIVE_CLIENT_SECRET]
   --hidrive-disable-fetching-member-count value  디렉토리의 객체 수를 필요한 경우만 가져오지 않습니다. (기본값: "false") [$HIDRIVE_DISABLE_FETCHING_MEMBER_COUNT]
   --hidrive-encoding value                       백엔드용 인코딩. (기본값: "Slash,Dot") [$HIDRIVE_ENCODING]
   --hidrive-endpoint value                       서비스의 엔드포인트. (기본값: "https://api.hidrive.strato.com/2.1") [$HIDRIVE_ENDPOINT]
   --hidrive-root-prefix value                    모든 경로의 루트/상위 폴더. (기본값: "/") [$HIDRIVE_ROOT_PREFIX]
   --hidrive-scope-access value                   rclone이 HiDrive에 액세스할 때 사용할 액세스 권한. (기본값: "rw") [$HIDRIVE_SCOPE_ACCESS]
   --hidrive-scope-role value                     rclone이 HiDrive에 액세스할 때 사용할 사용자 수준. (기본값: "user") [$HIDRIVE_SCOPE_ROLE]
   --hidrive-token value                          OAuth 액세스 토큰(JSON 형식). [$HIDRIVE_TOKEN]
   --hidrive-token-url value                      토큰 서버 URL. [$HIDRIVE_TOKEN_URL]
   --hidrive-upload-concurrency value             청크 업로드용 동시성. (기본값: "4") [$HIDRIVE_UPLOAD_CONCURRENCY]
   --hidrive-upload-cutoff value                  청크 업로드를 위한 임계값. (기본값: "96Mi") [$HIDRIVE_UPLOAD_CUTOFF]

```
{% endcode %}