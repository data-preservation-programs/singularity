# WebDAV

{% code fullWidth="true" %}
```
이름:
   singularity storage create webdav - WebDAV

사용법:
   singularity storage create webdav [command options] [arguments...]

설명:
   --url
      연결하려는 HTTP 호스트의 URL입니다.
      
      예시: https://example.com.

   --vendor
      사용 중인 WebDAV 사이트/서비스/소프트웨어의 이름입니다.

      예시:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Sharepoint Online, Microsoft 계정으로 인증
         | sharepoint-ntlm | NTLM 인증을 사용하는 Sharepoint, 일반적으로 자체 호스트나 온프레미스에서 운영
         | other           | 기타 사이트/서비스 또는 소프트웨어

   --user
      사용자 이름입니다.
      
      NTLM 인증을 사용하는 경우, 사용자 이름은 '도메인\사용자' 형식이어야 합니다.

   --pass
      비밀번호입니다.

   --bearer-token
      사용자/비밀번호 대신에 Bearer 토큰(마카룬)을 사용합니다.

   --bearer-token-command
      Bearer 토큰을 가져오기 위해 실행할 명령입니다.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.
      
      sharepoint-ntlm의 경우 기본 인코딩은 Slash, LtGt, DoubleQuote, Colon, Question, Asterisk, Pipe, Hash, Percent, BackSlash, Del, Ctl, LeftSpace, LeftTilde, RightSpace, RightPeriod, InvalidUtf8입니다. 그 외에는 identity를 사용합니다.

   --headers
      모든 트랜잭션에 대해 HTTP 헤더를 설정합니다.
      
      이를 사용하여 모든 트랜잭션에 대한 추가적인 HTTP 헤더를 설정합니다.
      
      입력 형식은 쉼표로 구분된 키,값 쌍의 목록입니다. 표준 [CSV 인코딩](https://godoc.org/encoding/csv)을 사용할 수 있습니다.

      예를 들어 쿠키를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'을 사용하세요.

      여러 개의 헤더를 설정할 수 있습니다. 예: '"Cookie","name=value","Authorization","xxx"'.


옵션:
   --bearer-token value  사용자/비밀번호 대신에 Bearer 토큰(마카룬)을 사용합니다. [$BEARER_TOKEN]
   --help, -h            도움말 표시
   --pass value          비밀번호입니다. [$PASS]
   --url value           연결하려는 HTTP 호스트의 URL입니다. [$URL]
   --user value          사용자 이름입니다. [$USER]
   --vendor value        사용 중인 WebDAV 사이트/서비스/소프트웨어의 이름입니다. [$VENDOR]

   Advanced

   --bearer-token-command value  Bearer 토큰을 가져오기 위해 실행할 명령입니다. [$BEARER_TOKEN_COMMAND]
   --encoding value              백엔드의 인코딩입니다. [$ENCODING]
   --headers value               모든 트랜잭션에 대해 HTTP 헤더를 설정합니다. [$HEADERS]

   General

   --name value  저장소의 이름 (기본: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}