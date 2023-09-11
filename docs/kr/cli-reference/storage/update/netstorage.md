# Akamai NetStorage

{% code fullWidth="true" %}
```
이름:
   singularity storage update netstorage - Akamai NetStorage

사용법:
   singularity storage update netstorage [command options] <name|id>

설명:
   --protocol
      HTTP 또는 HTTPS 프로토콜 중 선택하세요.
      
      대부분의 사용자는 기본값인 HTTPS를 선택해야 합니다.
      HTTP는 주로 디버깅을 위해 제공됩니다.

      예:
         | http  | HTTP 프로토콜
         | https | HTTPS 프로토콜

   --host
      연결할 NetStorage 호스트의 도메인+경로입니다.
      
      형식은 `<도메인>/<내부 폴더>`여야 합니다.

   --account
      NetStorage 계정 이름을 설정합니다.

   --secret
      인증을 위한 NetStorage 계정 비밀/ G2O 키를 설정합니다.
      
      비밀번호를 설정하려면 'y' 옵션을 선택하고 비밀을 입력하세요.


옵션:
   --account value  NetStorage 계정 이름 설정 [$ACCOUNT]
   --help, -h       도움말 보기
   --host value     연결할 NetStorage 호스트의 도메인+경로 설정 [$HOST]
   --secret value   인증을 위한 NetStorage 계정 비밀/G2O 키 설정 [$SECRET]

   고급

   --protocol value  HTTP 또는 HTTPS 프로토콜 중 선택합니다. (기본값: "https") [$PROTOCOL]

```
{% endcode %}