# FTP

{% code fullWidth="true" %}
```
이름:
   singularity 데이터소스 추가 ftp - FTP

사용법:
   singularity 데이터소스 추가 ftp [command options] <데이터셋_이름> <소스_경로>

설명:
   --ftp-ask-password
      FTP 암호를 요청할 수 있도록 합니다.
      
      이 옵션이 설정되어 있고 암호가 제공되지 않았다면 rclone은 암호를 요청할 것입니다.
      

   --ftp-close-timeout
      닫기 응답까지 대기할 수 있는 최대 시간입니다.

   --ftp-concurrency
      FTP 동시 연결의 최대 수입니다. 무제한인 경우 0을 설정합니다.
      
      이 값을 설정하면 데드락이 발생할 가능성이 매우 크므로 신중하게 사용해야 합니다.
      
      동기화 또는 복사를 수행하는 경우 concurrency는 `--transfers`와 `--checkers`의 합보다 1개 많아야 합니다.
      
      `--check-first`를 사용하는 경우 `--checkers`와 `--transfers` 중 가장 큰 항목보다 1개 많아야 합니다.
      
      예를 들어 concurrency 3의 경우 `--checkers 2 --transfers 2 --check-first` 또는 `--checkers 1 --transfers 1`을 사용합니다.
      
      

   --ftp-disable-epsv
      서버에서 EPSV 사용을 비활성화합니다.

   --ftp-disable-mlsd
      서버에서 MLSD 사용을 비활성화합니다.

   --ftp-disable-tls13
      TLS 1.3 사용 비활성화(오작동하는 FTP 서버에 대한 해결책)

   --ftp-disable-utf8
      서버가 지원을 선언해도 UTF-8 사용을 비활성화합니다.

   --ftp-encoding
      백엔드의 인코딩입니다.
      
      자세한 내용은 [개요의 인코딩 섹션](/overview/#encoding)을 참조하세요.

      예시:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPd에서 파일 이름에 '*' 사용이 불가능합니다.
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPd에서 '[]' 또는 '*'가 포함된 파일 이름 사용이 불가능합니다.
         | Ctl,LeftPeriod,Slash                                 | VsFTPd에서 점으로 시작하는 파일 이름 사용이 불가능합니다.

   --ftp-explicit-tls
      명시적 FTPS (TLS를 통한 FTP) 사용합니다.
      
      명시적 FTP over TLS를 사용하면 클라이언트는 평문 연결을 암호화된 연결로 업그레이드하기 위해 서버에 안전성을 요청합니다.
      명시적 FTPS와 암묵적 FTPS를 함께 사용할 수 없습니다.

   --ftp-force-list-hidden
      LIST -a를 사용하여 숨겨진 파일 및 폴더의 나열을 강제합니다. 이렇게 하면 MLSD 사용이 비활성화됩니다.

   --ftp-host
      연결할 FTP 호스트입니다.
      
      예: "ftp.example.com".

   --ftp-idle-timeout
      유휴 연결을 닫기 전 최대 시간입니다.
      
      주어진 시간 동안 연결 풀에 반환된 연결이 없는 경우 rclone은 연결 풀을 비웁니다.
      
      연결을 무기한으로 유지하려면 0으로 설정하세요.
      

   --ftp-no-check-certificate
      서버의 TLS 인증서를 확인하지 않습니다.

   --ftp-pass
      FTP 암호입니다.

   --ftp-port
      FTP 포트 번호입니다.

   --ftp-shut-timeout
      데이터 연결 종료 상태를 기다리는 최대 시간입니다.

   --ftp-tls
      암묵적 FTPS (TLS를 통한 FTP) 사용합니다.
      
      암묵적 FTP over TLS를 사용하면 클라이언트는 TLS를 시작부터 사용하여 연결하며 이는 TLS를 인식하지 못하는 서버와 호환성을 끊게 됩니다. 
      일반적으로 포트 990에 서비스됩니다(기본 포트인 21 대신). 명시적 FTPS와 함께 사용할 수 없습니다.

   --ftp-tls-cache-size
      모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다.
      
      TLS 캐시는 TLS 세션을 재개하고 연결 사이에 PSK를 재사용할 수 있도록 합니다.
      기본 크기가 충분하지 않을 경우 TLS 재개 오류가 발생하므로 이 값을 증가시키세요.
      기본값으로 사용합니다. 0을 사용하여 비활성화하세요.

   --ftp-user
      FTP 사용자 이름입니다.

   --ftp-writing-mdtm
      MDTM을 사용하여 수정 시간을 설정합니다 (VsFtpd의 특이점)


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] 데이터셋을 CAR 파일로 내보낸 후에 파일을 삭제합니다. (기본값: false)
   --rescan-interval value  마지막 성공적인 검색으로부터 지정된 간격 이후에 자동으로 소스 디렉토리를 검색합니다 (기본값: 비활성화)
   --scanning-state value   초기 검색 상태를 설정합니다 (기본값: 준비됨)

   ftp 옵션

   --ftp-ask-password value          FTP 암호를 요청할 수 있도록 합니다. (기본값: "false") [$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         닫기 응답까지 대기할 수 있는 최대 시간입니다. (기본값: "1m0s") [$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           FTP 동시 연결의 최대 수입니다. (기본값: "0") [$FTP_CONCURRENCY]
   --ftp-disable-epsv value          서버에서 EPSV 사용을 비활성화합니다. (기본값: "false") [$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          서버에서 MLSD 사용을 비활성화합니다. (기본값: "false") [$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         TLS 1.3 사용 비활성화(오작동하는 FTP 서버에 대한 해결책) (기본값: "false") [$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          서버가 지원을 선언해도 UTF-8 사용을 비활성화합니다. (기본값: "false") [$FTP_DISABLE_UTF8]
   --ftp-encoding value              백엔드의 인코딩입니다. (기본값: "Slash,Del,Ctl,RightSpace,Dot") [$FTP_ENCODING]
   --ftp-explicit-tls value          명시적 FTPS (TLS를 통한 FTP) 사용합니다. (기본값: "false") [$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     LIST -a를 사용하여 숨겨진 파일 및 폴더의 나열을 강제합니다. 이렇게 하면 MLSD 사용이 비활성화됩니다. (기본값: "false") [$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  연결할 FTP 호스트입니다. [$FTP_HOST]
   --ftp-idle-timeout value          유휴 연결을 닫기 전 최대 시간입니다. (기본값: "1m0s") [$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  서버의 TLS 인증서를 확인하지 않습니다. (기본값: "false") [$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTP 암호입니다. [$FTP_PASS]
   --ftp-port value                  FTP 포트 번호입니다. (기본값: "21") [$FTP_PORT]
   --ftp-shut-timeout value          데이터 연결 종료 상태를 기다리는 최대 시간입니다. (기본값: "1m0s") [$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   암묵적 FTPS (TLS를 통한 FTP) 사용합니다. (기본값: "false") [$FTP_TLS]
   --ftp-tls-cache-size value        모든 제어 및 데이터 연결에 대한 TLS 세션 캐시의 크기입니다. (기본값: "32") [$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTP 사용자 이름입니다. (기본값: "$USER") [$FTP_USER]
   --ftp-writing-mdtm value          MDTM을 사용하여 수정 시간을 설정합니다 (VsFtpd의 특이점) (기본값: "false") [$FTP_WRITING_MDTM]

```
{% endcode %}