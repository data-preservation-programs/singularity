# Zoho

{% code fullWidth="true" %}
```
이름:
   singularity storage update zoho - Zoho

사용법:
   singularity storage update zoho [command options] <name|id>

설명:
   --client-id
      OAuth 클라이언트 ID.
      
      일반적으로 비워두십시오.

   --client-secret
      OAuth 클라이언트 시크릿.
      
      일반적으로 비워두십시오.

   --token
      JSON blob 형식의 OAuth 액세스 토큰.

   --auth-url
      인증 서버 URL.
      
      기본값을 사용하려면 비워두십시오.

   --token-url
      토큰 서버 URL.
      
      기본값을 사용하려면 비워두십시오.

   --region
      연결할 Zoho 지역.
      
      조직이 등록된 지역을 사용해야 합니다. 확실하지 않은 경우, 브라우저를 통해
      연결하는 것과 동일한 최상위 도메인을 사용하십시오.

      예:
         | com    | 미국 / 글로벌
         | eu     | 유럽
         | in     | 인도
         | jp     | 일본
         | com.cn | 중국
         | com.au | 호주

   --encoding
      백엔드에 대한 인코딩 방식.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --client-id value      OAuth 클라이언트 ID. [$CLIENT_ID]
   --client-secret value  OAuth 클라이언트 시크릿. [$CLIENT_SECRET]
   --help, -h             도움말 표시
   --region value         연결할 Zoho 지역. [$REGION]

   고급

   --auth-url value   인증 서버 URL. [$AUTH_URL]
   --encoding value   백엔드에 대한 인코딩 방식. (기본값: "Del,Ctl,InvalidUtf8") [$ENCODING]
   --token value      JSON blob 형식의 OAuth 액세스 토큰. [$TOKEN]
   --token-url value  토큰 서버 URL. [$TOKEN_URL]

```
{% endcode %}