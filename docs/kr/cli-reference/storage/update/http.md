# HTTP

{% code fullWidth="true" %}
```
이름:
   singularity storage update http - HTTP

사용법:
   singularity storage update http [command options] <name|id>

설명:
   --url
      연결할 HTTP 호스트의 URL입니다.
      
      예: "https://example.com" 또는 "https://user:pass@example.com", 
      사용자 이름과 비밀번호를 사용하려는 경우입니다.

   --headers
      모든 트랜잭션에 대한 HTTP 헤더를 설정합니다.
      
      이를 사용하여 모든 트랜잭션에 대해 추가 HTTP 헤더를 설정합니다.
      
      입력 형식은 쉼표로 구분된 키,값 쌍의 목록입니다. 
      표준 [CSV 인코딩](https://godoc.org/encoding/csv)을 사용할 수 있습니다.
      
      예를 들어, 쿠키를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'를 사용합니다.
      
      여러 개의 헤더를 설정할 수 있습니다. 예: '"Cookie","name=value","Authorization","xxx"'.

   --no-slash
      사이트가 디렉토리를 /로 종료하지 않는 경우에 이를 설정합니다.
      
      대상 웹 사이트가 디렉토리 끝에 /를 사용하지 않는 경우에 사용합니다.
      
      경로 끝의 /는 rclone이 보통 파일과 디렉토리를 구별하는 방법입니다. 
      이 플래그가 설정되면 rclone은 Content-Type: text/html을 가진 모든 파일을 디렉토리로 취급하고 파일을 다운로드하는 대신 
      URL을 읽어옵니다.
      
      이로 인해 rclone이 진짜 HTML 파일을 디렉토리로 오해할 수 있음에 유의하세요.

   --no-head
      HEAD 요청을 사용하지 않습니다.
      
      HEAD 요청은 디렉토리 목록에서 파일 크기를 찾는 데 주로 사용됩니다.
      사이트의 로드 속도가 매우 느리면 이 옵션을 사용해 볼 수 있습니다.
      일반적으로 rclone은 디렉토리 목록에 있는 각 파일에 대해 HEAD 요청을 수행하여 다음을 확인합니다:
      
      - 파일의 크기
      - 파일이 실제로 존재하는지 확인
      - 디렉토리인지 확인
      
      이 옵션을 설정하면 rclone은 HEAD 요청을 수행하지 않습니다. 
      이는 디렉토리 목록을 훨씬 빠르게 가져올 것이지만, rclone은 파일의 시간 또는 크기를 알 수 없게되며, 
      디렉토리 목록에는 존재하지 않는 일부 파일이 있을 수 있습니다.


옵션:
   --help, -h   도움말 표시
   --url value  연결할 HTTP 호스트의 URL입니다. [$URL]

   고급

   --headers value  모든 트랜잭션에 대한 HTTP 헤더를 설정합니다. [$HEADERS]
   --no-head        HEAD 요청을 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-slash       사이트가 디렉토리를 /로 종료하지 않는 경우에 설정합니다. (기본값: false) [$NO_SLASH]
```
{% endcode %}