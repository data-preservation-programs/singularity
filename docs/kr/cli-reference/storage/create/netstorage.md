# Akamai NetStorage

{% code fullWidth="true" %}
```
이름:
   singularity storage create netstorage - Akamai NetStorage

사용법:
   singularity storage create netstorage [command options] [arguments...]

설명:
   --protocol
      HTTP 또는 HTTPS 프로토콜 중 하나를 선택하세요.

      대부분의 사용자는 기본값인 HTTPS를 선택해야 합니다.
      HTTP는 주로 디버깅을 위해 제공됩니다.

      예시:
         | http  | HTTP 프로토콜
         | https | HTTPS 프로토콜

   --host
      연결할 NetStorage 호스트의 도메인+경로를 지정하세요.
      
      형식은 `<도메인>/<내부 폴더>`여야 합니다.

   --account
      NetStorage 계정 이름을 설정하세요.

   --secret
      인증을 위한 NetStorage 계정 비밀번호 또는 G2O 키를 설정하세요.
      
      비밀번호를 직접 설정하려면 'y' 옵션을 선택한 후 비밀번호를 입력하세요.


옵션:
   --account value  NetStorage 계정 이름을 설정하세요 [$ACCOUNT]
   --help, -h       도움말 표시
   --host value     연결할 NetStorage 호스트의 도메인+경로를 지정하세요 [$HOST]
   --secret value   인증을 위한 NetStorage 계정 비밀번호 또는 G2O 키를 설정하세요 [$SECRET]

   고급

   --protocol value  HTTP 또는 HTTPS 프로토콜 중 하나를 선택하세요. (기본값: "https") [$PROTOCOL]

   일반

   --name value  스토리지 이름 (기본값: 자동 생성됨)
   --path value  스토리지 경로

```
{% endcode %}