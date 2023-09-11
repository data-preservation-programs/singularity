# HiDrive

{% code fullWidth="true" %}
```
명령어:
   singularity storage create hidrive - HiDrive

사용법:
   singularity storage create hidrive [command options] [arguments...]

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      보통 비워둡니다.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      보통 비워둡니다.

   --token
      JSON blob으로 표시된 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      제공자의 기본값을 사용하려면 비워둡니다.

   --token-url
      토큰 서버 URL.
      
      제공자의 기본값을 사용하려면 비워둡니다.

   --scope-access
      HiDrive를 요청할 때 rclone이 사용해야 할 액세스 권한입니다.

      예시:
         | rw | 리소스에 대한 읽기와 쓰기 액세스.
         | ro | 리소스에 대한 읽기 전용 액세스.

   --scope-role
      HiDrive를 요청할 때 rclone이 사용해야 할 사용자 수준입니다.

      예시:
         | user  | 관리 권한에 대한 사용자 수준 액세스.
         |       | 대부분의 경우 충분합니다.
         | admin | 확장된 관리 권한에 대한 완전한 액세스.
         | owner | 관리 권한에 대한 완전한 액세스.

   --root-prefix
      모든 경로의 루트/상위 폴더입니다.
      
      지정된 폴더를 모든 원격 경로의 시작 위치로 사용하려면 채웁니다.
      이렇게 하면 rclone이 시작점으로서 어떤 폴더를 사용할 수 있습니다.

      예시:
         | /       | rclone이 접근할 수 있는 가장 상위 디렉토리입니다.
         |         | rclone이 일반적인 HiDrive 사용자 계정을 사용하는 경우 이는 "root"와 동일합니다.
         | root    | HiDrive 사용자 계정의 가장 상위 디렉토리
         | <unset> | 경로에 대한 루트/접두사가 없음을 나타냅니다.
         |         | 이를 사용하려면 항상 유효한 상위 폴더를 지정하여 이 원격지로 경로를 지정해야 합니다. 예: "remote:/path/to/dir" 또는 "remote:root/path/to/dir".

   --endpoint
      서비스의 엔드포인트입니다.
      
      API 호출이 이 URL로 이루어집니다.

   --disable-fetching-member-count
      절대로 필요하지 않을 때는 디렉토리의 객채 수를 가져오지 않습니다.
      
      객체의 수를 가져오지 않으면 요청이 더 빠를 수 있습니다.

   --chunk-size
      청크 업로드용 청크 크기입니다.
      
      지정된 임계값 이상 또는 파일 크기를 알 수 없는 파일은 이 크기로 조각화되어 업로드됩니다.
      
      이 값의 상한은 2147483647바이트(약 2.000Gi)입니다.
      단일 업로드 작업에서 지원되는 최대 바이트 수입니다.
      이 상한을 초과하거나 음수 값을 설정하면 업로드가 실패합니다.
      
      이 값을 크게 설정하면 메모리를 더 사용하여 업로드 속도가 늘어날 수 있습니다.
      메모리 절약을 위해 이 값을 작게 설정할 수도 있습니다.

   --upload-cutoff
      청크 업로드용 임계값입니다.
      
      이 값 이상인 파일은 구성된 청크 크기로 조각화되어 업로드됩니다.
      
      이 값의 상한은 2147483647바이트(약 2.000Gi)입니다.
      단일 업로드 작업에서 지원되는 최대 바이트 수입니다.
      이 상한을 초과하면 업로드가 실패합니다.

   --upload-concurrency
      청크 업로드용 병렬 처리 수입니다.
      
      이는 동일한 파일에 대해 동시에 실행되는 전송 작업의 상한값입니다.
      1보다 작은 값으로 설정하면 업로드가 교착 상태에 빠질 수 있습니다.
      
      높은 속도의 링크를 통해 작은 수의 대용량 파일을 업로드하고
      이러한 업로드가 대역폭을 완전히 활용하지 못한다면
      이 값을 늘리면 전송 속도를 높일 수 있습니다.

   --encoding
      백엔드의 인코딩입니다.
      
      더 많은 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --scope-access value   HiDrive에서 요청할 때 rclone이 사용할 액세스 권한입니다. (기본값: "rw") [$SCOPE_ACCESS]

   고급

   --auth-url value                 인증 서버 URL. [$AUTH_URL]
   --chunk-size value               청크 업로드용 청크 크기. (기본값: "48Mi") [$CHUNK_SIZE]
   --disable-fetching-member-count  디렉토리의 객채 수를 가져오지 않습니다. (기본값: false) [$DISABLE_FETCHING_MEMBER_COUNT]
   --encoding value                 백엔드용 인코딩. (기본값: "Slash,Dot") [$ENCODING]
   --endpoint value                 서비스의 엔드포인트. (기본값: "https://api.hidrive.strato.com/2.1") [$ENDPOINT]
   --root-prefix value              모든 경로의 루트/상위 폴더입니다. (기본값: "/") [$ROOT_PREFIX]
   --scope-role value               HiDrive에서 요청할 때 rclone이 사용할 사용자 수준입니다. (기본값: "user") [$SCOPE_ROLE]
   --token value                    JSON blob으로 표시된 OAuth 액세스 토큰. [$TOKEN]
   --token-url value                토큰 서버 URL. [$TOKEN_URL]
   --upload-concurrency value       청크 업로드용 병렬 처리 수. (기본값: 4) [$UPLOAD_CONCURRENCY]
   --upload-cutoff value            청크 업로드용 임계값. (기본값: "96Mi") [$UPLOAD_CUTOFF]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}