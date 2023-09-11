# seafile

{% code fullWidth="true" %}
```
이름:
   싱기큘래리티 저장소 생성 seafile - seafile

사용법:
   싱기큘래리티 저장소 생성 seafile [command options] [arguments...]

설명:
   --url
      연결할 seafile 호스트의 URL입니다.

      예시:
         | https://cloud.seafile.com/ | cloud.seafile.com에 연결합니다.

   --user
      사용자 이름 (보통 이메일 주소).

   --pass
      비밀번호.

   --2fa
      이중 인증 (계정에 2FA가 사용 중인 경우 'true').

   --library
      라이브러리의 이름.
      
      암호화되지 않은 모든 라이브러리에 접근하려면 비워 두십시오.

   --library-key
      라이브러리 비밀번호 (암호화된 라이브러리에만 해당).
      
      명령 줄을 통해 전달하려면 비워 두십시오.

   --create-library
      라이브러리가 존재하지 않으면 생성할지 여부를 지정합니다.

   --auth-token
      인증 토큰.

   --encoding
      백엔드의 인코딩.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.


옵션:
   --2fa                이중 인증 (계정에 2FA가 사용 중인 경우 'true'). (기본값: false) [$2FA]
   --auth-token value   인증 토큰. [$AUTH_TOKEN]
   --help, -h           도움말 표시
   --library value      라이브러리의 이름. [$LIBRARY]
   --library-key value  라이브러리 비밀번호 (암호화된 라이브러리에만 해당). [$LIBRARY_KEY]
   --pass value         비밀번호. [$PASS]
   --url value          연결할 seafile 호스트의 URL. [$URL]
   --user value         사용자 이름 (보통 이메일 주소). [$USER]

   고급

   --create-library  라이브러리가 존재하지 않으면 rclone이 생성해야 하는지 여부입니다. (기본값: false) [$CREATE_LIBRARY]
   --encoding value  백엔드의 인코딩. (기본값: "슬래시,이중 인용부호,백슬래시,Ctl,유효하지 않은 Utf8") [$ENCODING]

   일반

   --name value  저장소의 이름 (기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}