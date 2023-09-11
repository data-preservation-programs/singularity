# 엔터프라이즈 파일 패브릭

{% code fullWidth="true" %}
```
NAME:
   singularity storage create filefabric - 엔터프라이즈 파일 패브릭

USAGE:
   singularity storage create filefabric [command options] [arguments...]

DESCRIPTION:
   --url
      연결할 엔터프라이즈 파일 패브릭의 URL입니다.

      예시:
         | https://storagemadeeasy.com       | Storage Made Easy US
         | https://eu.storagemadeeasy.com    | Storage Made Easy EU
         | https://yourfabric.smestorage.com | 자체의 엔터프라이즈 파일 패브릭에 연결합니다.

   --root-folder-id
      루트 폴더의 ID입니다.
      
      일반적으로 비워 두세요.
      
      특정 ID의 디렉토리로 시작하도록 rclone을 설정하려면 값을 지정하세요.
      

   --permanent-token
      영구 인증 토큰입니다.
      
      영구 인증 토큰은 엔터프라이즈 파일 패브릭에서 생성할 수 있으며, 사용자의 대시보드에서
      보안 아래의 "내 인증 토큰" 항목에서 만들 수 있습니다. 관리 버튼을 클릭하여 생성하세요.
      
      이러한 토큰은 보통 몇 년 동안 유효합니다.
      
      자세한 정보는 다음 링크를 참조하세요: https://docs.storagemadeeasy.com/organisationcloud/api-tokens
      

   --token
      세션 토큰입니다.
      
      rclone이 설정 파일에 캐시하는 세션 토큰입니다. 대개 1시간 동안 유효합니다.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.
      

   --token-expiry
      토큰 만료 시간입니다.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.
      

   --version
      파일 패브릭에서 읽은 버전입니다.
      
      이 값을 설정하지 마세요. rclone이 자동으로 설정합니다.
      

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --help, -h               도움말 표시
   --permanent-token value  영구 인증 토큰입니다. [$PERMANENT_TOKEN]
   --root-folder-id value   루트 폴더의 ID입니다. [$ROOT_FOLDER_ID]
   --url value              연결할 엔터프라이즈 파일 패브릭의 URL입니다. [$URL]

   Advanced

   --encoding value      백엔드의 인코딩입니다. (기본값: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --token value         세션 토큰입니다. [$TOKEN]
   --token-expiry value  토큰 만료 시간입니다. [$TOKEN_EXPIRY]
   --version value       파일 패브릭에서 읽은 버전입니다. [$VERSION]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}
