# 검색 요청을 처리하는 콘텐트 프로바이더를 시작합니다.

{% code fullWidth="true" %}
```
이름:
   singularity run content-provider - 검색 요청을 처리하는 콘텐트 프로바이더를 시작합니다.

사용법:
   singularity run content-provider [command options] [arguments...]

옵션:
   --help, -h  도움말 표시

   Bitswap 검색

   --enable-bitswap                                 Bitswap 검색 활성화 (기본값: false)
   --libp2p-identity-key value                      libp2p 피어의 base64로 인코딩된 개인 키 (기본값: 자동생성)
   --libp2p-listen value [ --libp2p-listen value ]  libp2p 연결을 위해 듣기 대기할 주소

   HTTP 검색

   --enable-http      HTTP 검색 활성화 (기본값: true)
   --http-bind value  HTTP 서버를 바인딩할 주소 (기본값: "127.0.0.1:7777")

```
{% endcode %}