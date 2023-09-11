# FTP

{% code fullWidth="true" %}
```
이름:
   singularity storage create ftp - FTP

사용법:
   singularity storage create ftp [command options] [arguments...]

설명:
   --host
      연결할 FTP 호스트입니다.
      
      예시: "ftp.example.com".

   --user
      FTP 사용자 이름입니다.

   --port
      FTP 포트 번호입니다.

   --pass
      FTP 비밀번호입니다.

   --tls
      암묵적인 FTPS(FTP over TLS) 사용.
      
      암묵적인 FTPS를 사용하면 클라이언트가 시작부터 TLS를 사용하여 연결하므로 
      TLS를 지원하지 않는 서버와 호환되지 않습니다. 
      이는 일반적으로 포트 21이 아닌 포트 990에서 서비스됩니다. 
      명시적인 FTPS와 함께 사용할 수 없습니다.

   --explicit-tls
      명시적인 FTPS(FTP over TLS) 사용.
      
      명시적인 FTPS를 사용하면 클라이언트가 서버에 보안을 요청하여 
      일반 텍스트 연결을 암호화된 연결로 업그레이드합니다. 
      암묵적인 FTPS와 함께 사용할 수 없습니다.

   --concurrency
      최대 FTP 동시 연결 수, 무제한은 0입니다.
      
      이 값을 설정하면 데드락이 발생할 확률이 매우 높으므로 주의해서 사용해야 합니다.
      
      동기화 또는 복사 작업을 수행하는 경우 concurrency를 `--transfers` 및 `--checkers`의 합보다 1 더 크게 설정해야 합니다.
      
      `--check-first`를 사용하는 경우 `--checkers`와 `--transfers` 중 최댓값보다 1 더 크게 설정해야 합니다.
      
      예를 들어 `concurrency 3`의 경우 `--checkers 2 --transfers 2 --check-first` 또는 `--checkers 1 --transfers 1`을 사용합니다.
      


   --no-check-certificate
      서버의 TLS 인증서를 검증하지 않습니다.

   --disable-epsv
      서버가 지원을 알리더라도 EPSV 사용하지 않음을 비활성화합니다.

   --disable-mlsd
      서버가 지원을 알리더라도 MLSD 사용하지 않음을 비활성화합니다.

   --disable-utf8
      서버가 지원을 알리더라도 UTF-8 사용하지 않음을 비활성화합니다.

   --writing-mdtm
      수정 시간을 설정하기 위해 MDTM 사용(VsFtpd Quirk)

   --force-list-hidden
      숨겨진 파일과 폴더의 리스트를 보기 위해 LIST -a 사용. 이는 MLSD 사용을 비활성화합니다.

   --idle-timeout
      유휴 연결을 닫기 전에 대기할 최대 시간입니다.
      
      주어진 시간 동안 연결 풀에 반환된 연결이 없는 경우 rclone은 연결 풀을 비울 것입니다.
      
      0으로 설정하면 연결을 영구적으로 유지합니다.
      

   --close-timeout
      닫기 응답을 기다리는 최대 시간입니다.

   --tls-cache-size
      모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다.
      
      TLS 캐시는 TLS 세션을 재개하고 연결 사이에서 PSK를 재사용할 수 있게 합니다.
      기본 크기가 충분하지 않을 경우 TLS 재개 오류가 발생하므로 크기를 늘리세요.
      기본값은 활성화됩니다. 비활성화하려면 0을 사용하세요.

   --disable-tls13
      TLS 1.3 비활성화(오류가 있는 FTP 서버를 위한 해결책)

   --shut-timeout
      데이터 연결 종료 상태를 기다리는 최대 시간입니다.

   --ask-password
      필요할 때 FTP 비밀번호를 요청할 수 있게 합니다.
      
      이를 설정하고 비밀번호를 제공하지 않으면 rclone은 비밀번호를 요청할 것입니다.
      

   --encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요 섹션](/overview/#encoding)의 인코딩 섹션을 참조하세요.

      예시:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd는 파일 이름에 '*'를 처리할 수 없습니다.
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd는 파일 이름에 '[]'나 '*'를 처리할 수 없습니다.
         | Ctl,LeftPeriod,Slash                                 | VsFTPd는 '.'으로 시작하는 파일 이름을 처리할 수 없습니다.


옵션:
   --explicit-tls  명시적인 FTPS(FTP over TLS) 사용. (기본값: false) [$EXPLICIT_TLS]
   --help, -h      도움말 표시
   --host value    연결할 FTP 호스트입니다. [$HOST]
   --pass value    FTP 비밀번호입니다. [$PASS]
   --port value    FTP 포트 번호입니다. (기본값: 21) [$PORT]
   --tls           암묵적인 FTPS(FTP over TLS) 사용. (기본값: false) [$TLS]
   --user value    FTP 사용자 이름입니다. (기본값: "$USER") [$USER]

   Advanced

   --ask-password          필요할 때 FTP 비밀번호를 요청할 수 있게 합니다. (기본값: false) [$ASK_PASSWORD]
   --close-timeout value   닫기 응답을 기다리는 최대 시간입니다. (기본값: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     최대 FTP 동시 연결 수, 무제한은 0입니다. (기본값: 0) [$CONCURRENCY]
   --disable-epsv          서버가 지원을 알리더라도 EPSV 사용하지 않음을 비활성화합니다. (기본값: false) [$DISABLE_EPSV]
   --disable-mlsd          서버가 지원을 알리더라도 MLSD 사용하지 않음을 비활성화합니다. (기본값: false) [$DISABLE_MLSD]
   --disable-tls13         TLS 1.3 비활성화(오류가 있는 FTP 서버를 위한 해결책) (기본값: false) [$DISABLE_TLS13]
   --disable-utf8          서버가 지원을 알리더라도 UTF-8 사용하지 않음을 비활성화합니다. (기본값: false) [$DISABLE_UTF8]
   --encoding value        백엔드의 인코딩입니다. (기본값: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     숨겨진 파일과 폴더의 리스트를 보기 위해 LIST -a 사용. 이는 MLSD 사용을 비활성화합니다. (기본값: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    유휴 연결을 닫기 전에 대기할 최대 시간입니다. (기본값: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  서버의 TLS 인증서를 검증하지 않습니다. (기본값: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    데이터 연결 종료 상태를 기다리는 최대 시간입니다. (기본값: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다. (기본값: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          수정 시간을 설정하기 위해 MDTM 사용(VsFtpd Quirk) (기본값: false) [$WRITING_MDTM]

   General

   --name value  저장소의 이름(기본값: 자동 생성)
   --path value  저장소의 경로

```
{% endcode %}