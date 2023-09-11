# Yandex Disk

{% code fullWidth="true" %}
```
명령어:
   singularity storage create yandex - Yandex Disk

사용법:
   singularity storage create yandex [옵션] [인자...]

설명:
   --client-id
      OAuth 클라이언트 ID입니다.
      
      일반적인 경우 비워 두십시오.

   --client-secret
      OAuth 클라이언트 비밀입니다.
      
      일반적인 경우 비워 두십시오.

   --token
      JSON blob 형태의 OAuth 액세스 토큰입니다.

   --auth-url
      인증 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워 두십시오.

   --token-url
      토큰 서버 URL입니다.
      
      제공자 기본값을 사용하려면 비워 두십시오.

   --hard-delete
      파일을 휴지통으로 보내는 대신 영구적으로 삭제합니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀. [$CLIENT_SECRET]
   --help, -h             도움말 표시

   고급 옵션

   --auth-url value   인증 서버 URL. [$AUTH_URL]
   --encoding value   백엔드의 인코딩. (기본값: "슬래시,딜리트,컨트롤,유효하지 않은 UTF-8,점") [$ENCODING]
   --hard-delete      파일을 휴지통으로 보내는 대신 영구적으로 삭제합니다. (기본값: false) [$HARD_DELETE]
   --token value      JSON blob 형태의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value  토큰 서버 URL입니다. [$TOKEN_URL]

   일반 옵션

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}