# HTTP

{% code fullWidth="true" %}
```
이름:
   singularity datasource add http - HTTP

사용법:
   singularity datasource add http [command options] <dataset_name> <source_path>

설명:
   --http-headers
      모든 트랜잭션에 대해 HTTP 헤더를 설정합니다.
      
      모든 트랜잭션에 대해 추가적인 HTTP 헤더를 설정하기 위해 사용합니다.
      
      입력 형식은 키, 값 쌍을 쉼표로 구분한 목록입니다. 표준 CSV 인코딩을 사용할 수 있습니다.
      
      예를 들어, Cookie를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'를 사용합니다.
      
      여러 개의 헤더를 설정할 수 있습니다. 예를 들어 '"Cookie","name=value","Authorization","xxx"'와 같습니다.

   --http-no-head
      HEAD 요청을 사용하지 않습니다.
      
      HEAD 요청은 주로 디렉터리 목록에서 파일 크기를 찾기 위해 사용됩니다.
      사이트가 매우 느린 경우 이 옵션을 사용해 볼 수 있습니다.
      일반적으로 rclone은 디렉터리 목록에서 각 잠재적인 파일에 대해 HEAD 요청을 수행하여 다음을 확인합니다.
      
      - 크기 확인
      - 존재 여부 확인
      - 디렉터리인지 확인
      
      이 옵션을 설정하면 rclone은 HEAD 요청을 수행하지 않습니다. 이렇게 하면 디렉터리 목록을 훨씬 더 빠르게 얻을 수 있지만, rclone은 파일의 시간 또는 크기를 갖지 않고 목록에 존재하지 않는 일부 파일이 포함될 수 있습니다.

   --http-no-slash
      사이트가 디렉터리를 /로 끝내지 않을 경우에 설정합니다.
      
      원하는 웹 사이트가 디렉터리 끝에 /를 사용하지 않는 경우에 사용합니다.
      
      경로 끝에 /를 사용하면 rclone은 일반적으로 파일과 디렉터리를 구분합니다. 이 플래그가 설정되면 rclone은 Content-Type: text/html인 모든 파일을 디렉터리로 처리하고 URL을 다운로드 대신 해당 파일에서 읽습니다.
      
      실제 HTML 파일과 디렉터리를 rclone이 혼동할 수 있음에 주의하십시오.

   --http-url
      연결할 HTTP 호스트의 URL입니다.
      
      예: "https://example.com" 또는 "https://user:pass@example.com"을 사용하여 사용자 이름과 암호를 사용합니다.


옵션:
   --help, -h  도움말 표시
   
   데이터 준비 옵션
   
   --delete-after-export    [주의] 데이터셋 파일을 CAR 파일로 내보낸 후에는 데이터셋 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 시간 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: ready)

   http에 대한 옵션
   
   --http-headers value   모든 트랜잭션에 대해 HTTP 헤더를 설정합니다. [$HTTP_HEADERS]
   --http-no-head value   HEAD 요청을 사용하지 않습니다. (기본값: "false") [$HTTP_NO_HEAD]
   --http-no-slash value  사이트가 디렉터리를 /로 끝내지 않을 경우에 설정합니다. (기본값: "false") [$HTTP_NO_SLASH]
   --http-url value       연결할 HTTP 호스트의 URL입니다. [$HTTP_URL]
```
{% endcode %}