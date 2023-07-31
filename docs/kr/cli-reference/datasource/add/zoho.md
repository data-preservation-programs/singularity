# Zoho

{% code fullWidth="true" %}
```
이름:
   singularity datasource add zoho - Zoho

사용법:
   singularity datasource add zoho [command options] <데이터셋_이름> <소스_경로>

설명:
   --zoho-auth-url
      인증 서버 URL입니다.
      
      매개변수 기본값을 사용하려면 비워 둡니다.

   --zoho-client-id
      OAuth 클라이언트 ID입니다.
      
      일반적으로 비워 둡니다.

   --zoho-client-secret
      OAuth 클라이언트 비밀입니다.
      
      일반적으로 비워 둡니다.

   --zoho-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

   --zoho-region
      연결할 Zoho 지역입니다.
      
      조직이 등록된 지역과 동일한 최상위 도메인을 사용해야 합니다. 확실하지 않으면 브라우저에서 연결하는 동일한 대상 수준 도메인을 사용하세요.

      예:
         | com    | 미국 / 글로벌
         | eu     | 유럽
         | in     | 인도
         | jp     | 일본
         | com.cn | 중국
         | com.au | 오스트레일리아

   --zoho-token
      OAuth 액세스 토큰(JSON blob)입니다.

   --zoho-token-url
      토큰 서버 URL입니다.
      
      매개변수 기본값을 사용하려면 비워 둡니다.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋의 파일을 내보낸 후에 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔 이후 지정된 시간 지난 후에 자동으로 소스 디렉토리를 다시 스캔합니다 (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비됨)

   Zoho 옵션

   --zoho-auth-url value       인증 서버 URL입니다. [$ZOHO_AUTH_URL]
   --zoho-client-id value      OAuth 클라이언트 ID입니다. [$ZOHO_CLIENT_ID]
   --zoho-client-secret value  OAuth 클라이언트 비밀입니다. [$ZOHO_CLIENT_SECRET]
   --zoho-encoding value       백엔드에 대한 인코딩입니다. (기본값: "Del,Ctl,InvalidUtf8") [$ZOHO_ENCODING]
   --zoho-region value         연결할 Zoho 지역입니다. [$ZOHO_REGION]
   --zoho-token value          OAuth 액세스 토큰(JSON blob)입니다. [$ZOHO_TOKEN]
   --zoho-token-url value      토큰 서버 URL입니다. [$ZOHO_TOKEN_URL]

```
{% endcode %}