# Zoho

{% code fullWidth="true" %}
```
NAME:
   singularity storage create zoho - Zoho

USAGE:
   singularity storage create zoho [command options] [arguments...]

DESCRIPTION:
   --client-id
      OAuth 클라이언트 식별자.
      
      일반적으로 비어둡니다.

   --client-secret
      OAuth 클라이언트 비밀번호.
      
      일반적으로 비어둡니다.

   --token
      JSON 형태의 OAuth 엑세스 토큰.

   --auth-url
      인증 서버 URL.
      
      제공자 기본값을 사용하려면 비워 둡니다.

   --token-url
      토큰 서버 URL.
      
      제공자 기본값을 사용하려면 비워 둡니다.

   --region
      연결할 Zoho 지역.
      
      기관이 등록된 지역을 사용해야 합니다. 확실하지 않은 경우, 브라우저에서 접속하는 것과 동일한 최상위 도메인을 사용하세요.

      예시:
         | com    | 미국 / 전역
         | eu     | 유럽
         | in     | 인도
         | jp     | 일본
         | com.cn | 중국
         | com.au | 호주

   --encoding
      백엔드의 인코딩 방식.
      
      자세한 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


OPTIONS:
   --client-id value      OAuth 클라이언트 식별자. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 비밀번호. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --region value         연결할 Zoho 지역. [$REGION]

   고급 옵션

   --auth-url value   인증 서버 URL. [$AUTH_URL]
   --encoding value   백엔드의 인코딩 방식. (기본값: "Del,Ctl,InvalidUtf8") [$ENCODING]
   --token value      JSON 형태의 OAuth 엑세스 토큰. [$TOKEN]
   --token-url value  토큰 서버 URL. [$TOKEN_URL]

   일반 옵션

   --name value  스토리지의 이름(기본값: 자동 생성됨)
   --path value  스토리지의 경로

```
{% endcode %}