# seafile

{% code fullWidth="true" %}
```
이름:
   singularity storage update seafile - seafile

사용법:
   singularity storage update seafile [command options] <name|id>

설명:
   --url
      연결할 seafile 호스트의 URL입니다.

      예제:
         | https://cloud.seafile.com/ | cloud.seafile.com에 연결합니다.

   --user
      사용자 이름 (일반적으로 이메일 주소).

   --pass
      비밀번호.

   --2fa
      두 단계 인증이 활성화되어 있는 경우 'true'입니다.

   --library
      라이브러리 이름.
      
      암호화되지 않은 모든 라이브러리에 액세스하려면 비워 두세요.

   --library-key
      라이브러리 비밀번호 (암호화된 라이브러리 전용).
      
      명령 줄을 통해 전달되면 비워 두세요.

   --create-library
      라이브러리가 존재하지 않는 경우 rclone이 라이브러리를 생성해야 할지 여부입니다.

   --auth-token
      인증 토큰.

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.


옵션:
   --2fa                두 단계 인증이 활성화되어 있는 경우 'true'입니다. (기본값: false) [$2FA]
   --auth-token value   인증 토큰. [$AUTH_TOKEN]
   --help, -h           도움말 표시
   --library value      라이브러리 이름. [$LIBRARY]
   --library-key value  라이브러리 비밀번호 (암호화된 라이브러리 전용). [$LIBRARY_KEY]
   --pass value         비밀번호. [$PASS]
   --url value          연결할 seafile 호스트의 URL. [$URL]
   --user value         사용자 이름 (일반적으로 이메일 주소). [$USER]

   고급

   --create-library  라이브러리가 존재하지 않는 경우 rclone이 라이브러리를 생성해야 할지 여부입니다. (기본값: false) [$CREATE_LIBRARY]
   --encoding value  백엔드의 인코딩. (기본값: "Slash,DoubleQuote,BackSlash,Ctl,InvalidUtf8") [$ENCODING]

```
{% endcode %}