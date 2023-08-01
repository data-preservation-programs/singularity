# Enterprise File Fabric

{% code fullWidth="true" %}
```
이름:
   singularity 데이터소스 추가 filefabric - Enterprise File Fabric

사용법:
   singularity datasource add filefabric [command options] <dataset_name> <source_path>

설명:
   --filefabric-encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요 섹션의 인코딩](/overview/#encoding)을 참조하세요.

   --filefabric-permanent-token
      영구적인 인증 토큰.
      
      영구적인 인증 토큰은 Enterprise File Fabric에서 생성할 수 있습니다.
      사용자 대시보드의 보안 섹션에 있는 "My Authentication Tokens"이라는 항목을 확인할 수 있습니다.
      생성하기 위해 "Manage" 버튼을 클릭하세요.
      
      이러한 토큰은 일반적으로 여러 년간 유효합니다.
      
      자세한 내용은 다음을 참조하세요: [https://docs.storagemadeeasy.com/organisationcloud/api-tokens](https://docs.storagemadeeasy.com/organisationcloud/api-tokens)

   --filefabric-root-folder-id
      루트 폴더의 식별자(ID).
      
      보통 비워둡니다.
      
      지정하면 rclone이 지정된 ID의 디렉토리로 시작합니다.

   --filefabric-token
      세션 토큰.
      
      이는 rclone이 구성 파일에 캐시하는 세션 토큰입니다. 보통 1시간 동안 유효합니다.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.

   --filefabric-token-expiry
      토큰 만료 시간.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.

   --filefabric-url
      연결할 Enterprise File Fabric의 URL입니다.

      예시:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 연결할 Enterprise File Fabric

   --filefabric-version
      파일 패브릭에서 읽은 버전.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 시간 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태 설정 (기본값: 준비 완료)

   filefabric 옵션

   --filefabric-encoding value         백엔드의 인코딩입니다. (기본값: "Slash,Del,Ctl,InvalidUtf8,Dot") [$FILEFABRIC_ENCODING]
   --filefabric-permanent-token value  영구적인 인증 토큰입니다. [$FILEFABRIC_PERMANENT_TOKEN]
   --filefabric-root-folder-id value   루트 폴더의 식별자(ID)입니다. [$FILEFABRIC_ROOT_FOLDER_ID]
   --filefabric-token value            세션 토큰입니다. [$FILEFABRIC_TOKEN]
   --filefabric-token-expiry value     토큰 만료 시간입니다. [$FILEFABRIC_TOKEN_EXPIRY]
   --filefabric-url value              연결할 Enterprise File Fabric의 URL입니다. [$FILEFABRIC_URL]
   --filefabric-version value          파일 패브릭에서 읽은 버전입니다. [$FILEFABRIC_VERSION]

```
{% endcode %}