# WebDAV

{% code fullWidth="true" %}
```
이름:
   singularity 데이터 소스 추가 webdav - WebDAV

사용법:
   singularity 데이터 소스 추가 webdav [command options] <데이터세트_이름> <소스_경로>

설명:
   --webdav-bearer-token
      사용자/비밀번호 대신 Bearer 토큰 사용 (예: 맥캐런).

   --webdav-bearer-token-command
      Bearer 토큰을 가져 오기 위해 실행 할 명령.

   --webdav-encoding
      백엔드에 대한 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.
      
      기본 인코딩은 Slash, LtGt, DoubleQuote, Colon, Question, Asterisk, Pipe, Hash, Percent, BackSlash, Del, Ctl, LeftSpace, LeftTilde, RightSpace, RightPeriod, InvalidUtf8로 설정됩니다. SharePoint-ntlm이나 그렇지 않은 경우 identity로 설정됩니다.

   --webdav-headers
      모든 트랜잭션에 대한 HTTP 헤더 설정.
      
      이를 사용하여 모든 트랜잭션에 대한 추가 HTTP 헤더를 설정합니다.
      
      입력 형식은 쉼표로 구분 된 key,value 쌍의 목록입니다. 표준
      [CSV 인코딩](https://godoc.org/encoding/csv)을 사용할 수 있습니다.
      
      예를 들어, 쿠키를 설정하려면 'Cookie,name=value' 또는 '"Cookie","name=value"'를 사용하십시오.
      
      여러 헤더를 설정할 수 있습니다. 예를 들어, '"Cookie","name=value","Authorization","xxx"'와 같이 설정할 수 있습니다.
      

   --webdav-pass
      비밀번호.

   --webdav-url
      연결할 http 호스트의 URL입니다.
      
      예: https://example.com.

   --webdav-user
      사용자 이름.
      
      NTLM 인증을 사용하는 경우, 사용자 이름은 'Domain\User' 형식이어야합니다.

   --webdav-vendor
      사용하는 WebDAV 사이트/서비스/소프트웨어의 이름입니다.

      예시:
         | nextcloud       | Nextcloud
         | owncloud        | Owncloud
         | sharepoint      | Microsoft 계정으로 인증 된 Sharepoint Online
         | sharepoint-ntlm | 일반적으로 자체 호스팅되거나 온 프레미스인 NTLM 인증이 있는 Sharepoint
         | other           | 다른 사이트/서비스 또는 소프트웨어


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [위험] 데이터세트를 CAR 파일로 내보낸 후 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 이 간격이 경과하면 소스 디렉토리를 자동으로 다시 스캔합니다. (기본값: 비활성화)
   --scanning-state value   초기 스캔 상태를 설정합니다. (기본값: 준비)

   WebDAV 옵션

   --webdav-bearer-token value          사용자/비밀번호 대신 Bearer 토큰 사용 (예: 맥캐런). [$WEBDAV_BEARER_TOKEN]
   --webdav-bearer-token-command value  Bearer 토큰을 가져 오기 위해 실행 할 명령. [$WEBDAV_BEARER_TOKEN_COMMAND]
   --webdav-encoding value              백엔드에 대한 인코딩. [$WEBDAV_ENCODING]
   --webdav-headers value               모든 트랜잭션에 대한 HTTP 헤더 설정. [$WEBDAV_HEADERS]
   --webdav-pass value                  비밀번호. [$WEBDAV_PASS]
   --webdav-url value                   연결할 http 호스트의 URL. [$WEBDAV_URL]
   --webdav-user value                  사용자 이름. [$WEBDAV_USER]
   --webdav-vendor value                사용하는 WebDAV 사이트/서비스/소프트웨어의 이름. [$WEBDAV_VENDOR]

```
{% endcode %}