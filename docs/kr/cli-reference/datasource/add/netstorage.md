# Akamai NetStorage

{% code fullWidth="true" %}
```
이름:
   singularity datasource add netstorage - Akamai NetStorage

사용법:
   singularity datasource add netstorage [command 옵션] <데이터셋_이름> <소스_경로>

설명:
   --netstorage-account
      NetStorage 계정 이름을 설정합니다.

   --netstorage-host
      연결할 NetStorage 호스트의 도메인+경로를 설정합니다.

      형식은 `<도메인>/<내부 폴더>`이어야 합니다.

   --netstorage-protocol
      HTTP 또는 HTTPS 프로토콜을 선택합니다.

      대부분의 사용자는 기본값인 HTTPS를 선택해야 합니다.
      HTTP는 주로 디버깅 용도로 제공됩니다.

      예시:
         | http  | HTTP 프로토콜
         | https | HTTPS 프로토콜

   --netstorage-secret
      인증을 위한 NetStorage 계정 비밀번호/키를 설정합니다.

      비밀번호를 직접 설정하려면 'y' 옵션을 선택한 후 비밀번호를 입력하세요.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공한 스캔 이후로 경과한 시간이 이 간격을 초과하면 자동으로 소스 디렉터리를 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   netstorage 옵션

   --netstorage-account value   NetStorage 계정 이름 설정 [$NETSTORAGE_ACCOUNT]
   --netstorage-host value      연결할 NetStorage 호스트의 도메인+경로 설정 [$NETSTORAGE_HOST]
   --netstorage-protocol value  HTTP 또는 HTTPS 프로토콜 선택 (기본값: "https") [$NETSTORAGE_PROTOCOL]
   --netstorage-secret value    인증을 위한 NetStorage 계정 비밀번호/키 설정 [$NETSTORAGE_SECRET]

```
{% endcode %}