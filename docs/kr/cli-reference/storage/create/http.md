# HTTP

{% code fullWidth="true" %}
```
이름:
   singularity storage create http - HTTP

사용법:
   singularity storage create http [command options] [arguments...]

설명:
   --url
      연결할 HTTP 호스트의 URL입니다.
      
      예: "https://example.com" 또는 "https://user:pass@example.com"와 같이 사용자 이름과 암호를 사용할 수 있습니다.

   --headers
      모든 트랜잭션에 대한 HTTP 헤더를 설정합니다.
      
      이를 사용하여 모든 트랜잭션에 대해 추가적인 HTTP 헤더를 설정할 수 있습니다.
      
      입력 형식은 key,value 쌍의 쉼표로 구분된 목록입니다. 표준 [CSV 인코딩](https://godoc.org/encoding/csv)을 사용할 수 있습니다.
      
      예를 들어, 쿠키를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'를 사용합니다.
      
      여러 개의 헤더를 설정할 수도 있습니다. 예: '"Cookie","name=value","Authorization","xxx"'.

   --no-slash
      웹 사이트가 디렉토리 마지막에 /을 사용하지 않는 경우에만 설정합니다.
      
      대상 웹 사이트가 디렉토리의 끝에 /을 사용하지 않는 경우에만 사용합니다.
      
      경로의 끝에 /가 있는지 여부는 rclone이 일반적으로 파일과 디렉토리를 구분하는 방법입니다. 이 플래그가 설정된 경우 rclone은 Content-Type: text/html인 경우 모든 파일을 디렉토리로 처리하고 해당 파일에서 URL을 읽지 않고 다운로드합니다.
      
      이로 인해 rclone은 실제로 HTML 파일을 디렉토리로 오해할 수 있음을 유의하세요.

   --no-head
      HEAD 요청을 사용하지 않습니다.
      
      HEAD 요청은 주로 디렉토리 목록에서 파일 크기를 찾기 위해 사용됩니다. 웹 사이트가 로드되는데 매우 느린 경우 이 옵션을 시도할 수 있습니다.
      일반적으로 rclone은 디렉토리 목록의 각 잠재적인 파일에 대해 다음을 위해 HEAD 요청을 수행합니다:
      
      - 크기를 찾기 위해
      - 실제로 존재하는지 확인하기 위해
      - 디렉토리인지 확인하기 위해
      
      이 옵션을 설정하면 rclone은 HEAD 요청을 수행하지 않습니다. 따라서 디렉토리 목록이 훨씬 빨리 로드되지만, rclone은 파일의 시간이나 크기를 가지지 않으며, 목록에 존재하지 않는 일부 파일이 있을 수 있습니다.


옵션:
   --help, -h   도움말 표시
   --url value  연결할 HTTP 호스트의 URL입니다. [$URL]

   Advanced

   --headers value  모든 트랜잭션에 대한 HTTP 헤더를 설정합니다. [$HEADERS]
   --no-head        HEAD 요청을 사용하지 않습니다. (기본값: false) [$NO_HEAD]
   --no-slash       웹 사이트가 디렉토리 마지막에 /을 사용하지 않는 경우만 설정합니다. (기본값: false) [$NO_SLASH]

   General

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}