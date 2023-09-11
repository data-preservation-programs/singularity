# Enterprise File Fabric

{% code fullWidth="true" %}
```
이름:
   singularity storage update filefabric - Enterprise File Fabric

사용법:
   singularity storage update filefabric [command options] <name|id>

설명:
   --url
      연결할 Enterprise File Fabric의 URL입니다.

      예시:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 자체 Enterprise File Fabric에 연결합니다

   --root-folder-id
      루트 폴더의 ID입니다.
      
      보통 비워놔두세요.
      
      특정 ID의 디렉토리로 rclone을 시작하려면 이 값을 채우세요.
      

   --permanent-token
      영구 인증 토큰입니다.
      
      어플리케이션 내 "마이 인증 토큰"이라는 이름의 항목에서 Enterprise File Fabric에서
      영구 인증 토큰을 만들 수 있습니다.
      
      이 토큰은 일반적으로 여러 년 동안 유효합니다.
      
      자세한 내용은 다음 링크를 참조하세요: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      세션 토큰입니다.
      
      이 값은 rclone이 구성 파일에 캐시하는 세션 토큰입니다.
      일반적으로 1시간 동안 유효합니다.
      
      이 값을 설정하지 마세요 - rclone이 자동으로 설정합니다.
      

   --token-expiry
      토큰 만료 시간입니다.
      
      이 값을 설정하지 마세요 - rclone이 자동으로 설정합니다.
      

   --version
      파일 패브릭에서 읽은 버전입니다.
      
      이 값을 설정하지 마세요 - rclone이 자동으로 설정합니다.
      

   --encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --help, -h               도움말 표시
   --permanent-token value  영구 인증 토큰입니다. [$PERMANENT_TOKEN]
   --root-folder-id value   루트 폴더의 ID입니다. [$ROOT_FOLDER_ID]
   --url value              연결할 Enterprise File Fabric의 URL입니다. [$URL]

   Advanced

   --encoding value      백엔드에 대한 인코딩입니다. (default: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         세션 토큰입니다. [$TOKEN]
   --token-expiry value  토큰 만료 시간입니다. [$TOKEN_EXPIRY]
   --version value       파일 패브릭에서 읽은 버전입니다. [$VERSION]

```
{% endcode %}