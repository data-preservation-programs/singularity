# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage update ftp - FTP

사용법:
   singularity storage update ftp [command options] <name|id>

DESCRIPTION:
   --host
      연결할 FTP 호스트입니다.
      
      예: "ftp.example.com".

   --user
      FTP 사용자 이름입니다.

   --port
      FTP 포트 번호입니다.

   --pass
      FTP 비밀번호입니다.

   --tls
      Implicit FTPS(FTP over TLS)를 사용합니다.
      
      Implicit FTPS를 사용하면 클라이언트가 TLS를 사용하여 연결을 시작하므로
      TLS를 지원하지 않는 서버와의 호환성이 깨집니다.
      일반적으로 포트 21보다 포트 990에서 사용됩니다.
      Explicit FTPS와 함께 사용할 수 없습니다.

   --explicit-tls
      Explicit FTPS(FTP over TLS)를 사용합니다.
      
      Explicit FTPS를 사용하면 클라이언트가 연결을 암호화된 연결로 업그레이드하기 위해
      서버에게 안전한 연결을 요청합니다.
      Implicit FTPS와 함께 사용할 수 없습니다.

   --concurrency
      최대 동시 FTP 연결 수입니다. 0은 제한 없음을 의미합니다.
      
      이 값을 설정하면 데드락이 발생할 가능성이 매우 크므로 주의해서 사용해야 합니다.
      
      동기화 또는 복사 작업을 수행하는 경우, concurrency는 `--transfers` 및 `--checkers`
      합계보다 1만큼 더 크도록 설정해야 합니다.
      
      `--check-first`를 사용하는 경우, `--checkers` 및 `--transfers`
      중 가장 큰 값보다 1만큼 더 크게 설정해야 합니다.
      
      예를 들어, `concurrency 3`인 경우 `--checkers 2 --transfers 2 --check-first`
      또는 `--checkers 1 --transfers 1`을 사용합니다.
      


   --no-check-certificate
      서버의 TLS 인증서를 확인하지 않습니다.

   --disable-epsv
      서버가 지원을 알리더라도 EPSV 사용을 비활성화합니다.

   --disable-mlsd
      서버가 지원을 알리더라도 MLSD 사용을 비활성화합니다.

   --disable-utf8
      서버가 지원을 알리더라도 UTF-8 사용을 비활성화합니다.

   --writing-mdtm
      수정 시간을 설정하기 위해 MDTM을 사용합니다 (VsFtpd 특이사항)

   --force-list-hidden
      LIST -a를 사용하여 숨겨진 파일과 폴더를 강제로 표시합니다.
      이렇게 하면 MLSD 사용이 비활성화됩니다.

   --idle-timeout
      유휴 연결을 닫기 전의 최대 시간입니다.
      
      주어진 시간 동안 연결 풀에 반환된 연결이 없는 경우,
      rclone은 연결 풀을 비웁니다.
      
      0으로 설정하여 연결을 계속 유지할 수 있습니다.
      

   --close-timeout
      닫기 응답을 기다리는 최대 시간입니다.

   --tls-cache-size
      모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다.
      
      TLS 캐시는 TLS 세션을 재개하고 연결 사이에서 PSK를 재사용할 수 있도록 합니다.
      기본 크기가 충분하지 않으면 TLS 재개 오류가 발생할 수 있으므로 크기를 늘립니다.
      기본적으로 활성화됩니다. 0을 사용하여 비활성화합니다.

   --disable-tls13
      TLS 1.3을 비활성화합니다 (버그 있는 FTP 서버를 위한 해결책)

   --shut-timeout
      데이터 연결 종료 상태를 기다리는 최대 시간입니다.

   --ask-password
      필요할 때 FTP 비밀번호를 요청할 수 있도록 허용합니다.
      
      이 값을 설정하고 비밀번호가 제공되지 않으면 rclone이 비밀번호를 요청합니다.
      

   --encoding
      백엔드에 대한 인코딩입니다.
      
      더 많은 정보는 [개요의 인코딩 섹션](/overview/#encoding)을 참조하십시오.

      예:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd는 파일 이름에 '*'을 처리할 수 없습니다.
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd는 파일 이름에 '[]' 또는 '*'을 처리할 수 없습니다.
         | Ctl,LeftPeriod,Slash                                 | VsFTPd는 마침표로 시작하는 파일 이름을 처리할 수 없습니다.


OPTIONS:
   --explicit-tls  Explicit FTPS(FTP over TLS)를 사용합니다. (기본값: false) [$EXPLICIT_TLS]
   --help, -h      도움말 표시
   --host value    연결할 FTP 호스트입니다. [$HOST]
   --pass value    FTP 비밀번호입니다. [$PASS]
   --port value    FTP 포트 번호입니다. (기본값: 21) [$PORT]
   --tls           Implicit FTPS(FTP over TLS)를 사용합니다. (기본값: false) [$TLS]
   --user value    FTP 사용자 이름입니다. (기본값: "$USER") [$USER]

   Advanced

   --ask-password          필요할 때 FTP 비밀번호를 요청할 수 있도록 허용합니다. (기본값: false) [$ASK_PASSWORD]
   --close-timeout value   닫기 응답을 기다리는 최대 시간입니다. (기본값: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     최대 동시 FTP 연결 수입니다. 0은 제한 없음을 의미합니다. (기본값: 0) [$CONCURRENCY]
   --disable-epsv          서버가 지원을 알리더라도 EPSV 사용을 비활성화합니다. (기본값: false) [$DISABLE_EPSV]
   --disable-mlsd          서버가 지원을 알리더라도 MLSD 사용을 비활성화합니다. (기본값: false) [$DISABLE_MLSD]
   --disable-tls13         TLS 1.3을 비활성화합니다 (버그 있는 FTP 서버를 위한 해결책) (기본값: false) [$DISABLE_TLS13]
   --disable-utf8          서버가 지원을 알리더라도 UTF-8 사용을 비활성화합니다. (기본값: false) [$DISABLE_UTF8]
   --encoding value        백엔드에 대한 인코딩입니다. (기본값: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     LIST -a를 사용하여 숨겨진 파일과 폴더를 강제로 표시합니다. 이렇게 하면 MLSD 사용이 비활성화됩니다. (기본값: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    유휴 연결을 닫기 전의 최대 시간입니다. (기본값: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  서버의 TLS 인증서를 확인하지 않습니다. (기본값: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    데이터 연결 종료 상태를 기다리는 최대 시간입니다. (기본값: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다. (기본값: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          수정 시간을 설정하기 위해 MDTM을 사용합니다 (VsFtpd 특이사항) (기본값: false) [$WRITING_MDTM]

```
{% endcode %}