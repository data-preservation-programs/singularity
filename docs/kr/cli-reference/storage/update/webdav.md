# WebDAV

{% code fullWidth="true" %}
```
NAME:
   singularity storage update webdav - 웹DAV

사용법:
   singularity storage update webdav [command options] <name|id>

설명:
   --url
      연결할 http 호스트의 URL입니다.
      
      예제: https://example.com.

   --vendor
      사용 중인 웹DAV 사이트/서비스/소프트웨어의 이름입니다.

      예시:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, Microsoft 계정으로 인증
         | sharepoint-ntlm | Sharepoint, NTLM 인증, 일반적으로 자체 호스팅 또는 온프레미스
         | other           | 기타 사이트/서비스 또는 소프트웨어

   --user
      사용자 이름입니다.
      
      NTLM 인증을 사용하는 경우, 사용자 이름은 'Domain\User' 형식이어야 합니다.

   --pass
      비밀번호입니다.

   --bearer-token
      사용자명/비밀번호 대신 베어러 토큰을 사용합니다. (예: Macaroon).

   --bearer-token-command
      베어러 토큰을 얻기 위해 실행할 명령입니다.

   --encoding
      백엔드에 대한 인코딩입니다.
      
      [개요](/overview/#encoding)의 [인코딩 섹션](/overview/#encoding)을 참조하십시오.
      
      sharepoint-ntlm의 기본 인코딩은 Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,Hash,Percent,BackSlash,Del,Ctl,LeftSpace,LeftTilde,RightSpace,RightPeriod,InvalidUtf8이며, 그렇지 않으면 identity입니다.

   --headers
      모든 트랜잭션에 대해 HTTP 헤더를 설정합니다.
      
      이를 사용하여 모든 트랜잭션에 대해 추가적인 HTTP 헤더를 설정하십시오.
      
      입력 형식은 쉼표로 구분된 키,값 쌍의 목록입니다. 표준 [CSV 인코딩](https://godoc.org/encoding/csv)을 사용할 수 있습니다.
      
      예를 들어, 쿠키를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'를 사용하십시오.
      
      여러 헤더를 설정할 수 있습니다. 예: '"Cookie","name=value","Authorization","xxx"'.
      


옵션:
   --bearer-token value  사용자명/비밀번호 대신 베어러 토큰을 사용합니다. (예: Macaroon). [$BEARER_TOKEN]
   --help, -h            도움말 표시
   --pass value          비밀번호입니다. [$PASS]
   --url value           연결할 http 호스트의 URL입니다. [$URL]
   --user value          사용자 이름입니다. [$USER]
   --vendor value        사용 중인 웹DAV 사이트/서비스/소프트웨어의 이름입니다. [$VENDOR]

   Advanced

   --bearer-token-command value  베어러 토큰을 얻기 위해 실행할 명령입니다. [$BEARER_TOKEN_COMMAND]
   --encoding value              백엔드에 대한 인코딩입니다. [$ENCODING]
   --headers value               모든 트랜잭션에 대해 HTTP 헤더를 설정합니다. [$HEADERS]

```
{% endcode %}