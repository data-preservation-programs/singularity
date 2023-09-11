# 야네 덱

{% code fullWidth="true" %}
```
이름:
   singularity storage update yandex - 야네 덱

사용법:
   singularity storage update yandex [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 아이디.
      
      일반적으로 비워두세요.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      일반적으로 비워두세요.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      공급자 기본 설정을 사용하려면 비워두세요.

   --token-url
      토큰 서버 URL.
      
      공급자 기본 설정을 사용하려면 비워두세요.

   --hard-delete
      파일을 휴지통에 넣지 않고 영구적으로 삭제합니다.

   --encoding
      백엔드의 인코딩 설정.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --client-id value      OAuth 클라이언트 아이디. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 시크릿. [$CLIENT_SECRET]
   --help, -h             도움말 표시하기

   고급 옵션

   --auth-url value   인증 서버 URL. [$AUTH_URL]
   --encoding value   백엔드의 인코딩 설정. (기본값: "Slash,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --hard-delete      파일을 휴지통에 넣지 않고 영구적으로 삭제합니다. (기본값: false) [$HARD_DELETE]
   --token value      JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value  토큰 서버 URL. [$TOKEN_URL]

```
{% endcode %}