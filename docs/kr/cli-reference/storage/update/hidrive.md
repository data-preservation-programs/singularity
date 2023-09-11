# HiDrive

{% code fullWidth="true" %}
```
이름:
   singularity storage update hidrive - HiDrive

사용법:
   singularity storage update hidrive [command options] <이름|ID>

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워두십시오.

   --client-secret
      OAuth 클라이언트 비밀.
      
      일반적으로 비워두십시오.

   --token
      JSON blob으로 된 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      Provider 기본값을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL.
      
      Provider 기본값을 사용하려면 비워두십시오.

   --scope-access
      HiDrive에 대한 액세스 권한으로 rclone이 요청할 때 사용해야하는 권한입니다.

      예시:
         | rw | 리소스에 대한 읽기 및 쓰기 액세스.
         | ro | 리소스에 대한 읽기 전용 액세스.

   --scope-role
      HiDrive로부터 액세스를 요청할 때 rclone이 사용해야하는 사용자 수준입니다.

      예시:
         | user  | 관리 권한에 대한 사용자 수준 액세스.
         |       | 대부분의 경우에 충분합니다.
         | admin | 관리 권한에 대한 광범위한 액세스.
         | owner | 관리 권한에 대한 완전한 액세스.

   --root-prefix
      모든 경로의 루트/부모 폴더입니다.
      
      지정된 폴더를 모든 경로의 부모로 사용하려면 입력하십시오.
      이렇게하면 rclone은 시작점으로 모든 폴더를 사용할 수 있습니다.

      예시:
         | /       | rclone이 접근할 수 있는 가장 상위 디렉토리.
         |         | 이는 rclone이 일반 HiDrive 사용자 계정을 사용하는 경우 "root"와 동등합니다.
         | root    | HiDrive 사용자 계정의 최상위 디렉토리
         | <unset> | 경로에 대한 루트 접두사가 없음을 지정합니다.
         |         | 이를 사용하는 경우 항상 유효한 부모 (예: "remote:/path/to/dir" 또는 "remote:root/path/to/dir")로 경로를 지정해야합니다.

   --endpoint
      서비스의 엔드포인트입니다.
      
      API 호출이 이루어질 URL입니다.

   --disable-fetching-member-count
      디렉토리의 객체 수를 가져 오지 않습니다. 필요하지 않은 경우에만 가져옵니다.
      
      하위 디렉토리의 객체 수를 가져 오지 않으면 요청이 더 빠를 수 있습니다.

   --chunk-size
      청크 업로드의 청크 크기입니다.
      
      구성된 임계 값보다 큰 파일 (또는 알 수없는 크기의 파일)은 이 크기로 청크 단위로 업로드됩니다.
      
      이 값의 상한은 2147483647 바이트 (약 2.000Gi)입니다.
      이것은 단일 업로드 작업이 지원하는 최대 바이트 수입니다.
      이 값을 상한을 초과하거나 음수 값으로 설정하면 업로드가 실패합니다.
      
      이 값을 더 큰 값으로 설정하면 업로드 속도가 빨라질 수 있지만 더 많은 메모리를 사용합니다.
      메모리 사용을 줄이려면 더 작은 값을 설정할 수 있습니다.

   --upload-cutoff
      청크 업로드 임계 값입니다.
      
      이 값보다 큰 파일은 구성된 청크 크기로 청크 단위로 업로드됩니다.
      
      이 값의 상한은 2147483647 바이트 (약 2.000Gi)입니다.
      이것은 단일 업로드 작업이 지원하는 최대 바이트 수입니다.
      이 값을 상한 이상으로 설정하면 업로드가 실패합니다.

   --upload-concurrency
      청크 업로드의 동시성입니다.
      
      동일한 파일에 대해 동시에 실행되는 전송 수의 상한입니다.
      1보다 작은 값으로 설정하면 업로드가 데드락에 걸릴 수 있습니다.
      
      대역폭을 충분히 이용하지 못하고 높은 속도로 많은 수의 대용량 파일을 업로드하는 경우,
      이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --scope-access value   HiDrive에서 액세스 요청할 때 rclone이 사용해야하는 액세스 권한입니다. (기본값: "rw") [$SCOPE_ACCESS]

   Advanced

   --auth-url value                 인증 서버 URL. [$AUTH_URL]
   --chunk-size value               청크 업로드의 청크 크기. (기본값: "48Mi") [$CHUNK_SIZE]
   --disable-fetching-member-count  디렉토리의 객체 수를 가져 오지 않습니다. 필요하지 않은 경우에만 가져옵니다. (기본값: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 백엔드의 인코딩. (기본값: "Slash,Dot") [$ENCODING]
   --endpoint value                 서비스의 엔드포인트. (기본값: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              모든 경로의 루트/부모 폴더. (기본값: "/") [$ROOT_PREFIX]
   --scope-role value               HiDrive에서 액세스를 요청할 때 rclone이 사용해야하는 사용자 수준입니다. (기본값: "user") [$SCOPE_ROLE]
   --token value                    JSON blob으로 된 OAuth 액세스 토큰. [$TOKEN]
   --token-url value                토큰 서버 URL. [$TOKEN_URL]
   --upload-concurrency value       청크 업로드의 동시성. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드 임계 값. (기본값: "96Mi") [$UPLOAD_CUTOFF]

```
{% endcode %}