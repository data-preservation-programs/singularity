# 검색 요청을 처리하는 콘텐트 프로바이더를 시작합니다.

{% code fullWidth="true" %}
```
이름:
   singularity run content-provider - 검색 요청을 처리하는 콘텐트 프로바이더를 시작합니다.

사용법:
   singularity run content-provider [command options] [arguments...]

옵션:
   --help, -h  도움말 표시

   HTTP IPFS Gateway

   --enable-http-ipfs  Enable trustless IPFS gateway on /ipfs/ (default: true)

   HTTP 검색

   --enable-http      HTTP 검색 활성화 (기본값: true)
   --http-bind value  HTTP 서버를 바인딩할 주소 (기본값: "127.0.0.1:7777")

```
{% endcode %}