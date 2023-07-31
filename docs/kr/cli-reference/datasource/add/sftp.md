# SSH/SFTP

{% code fullWidth="true" %}
```
이름:
   singularity datasource add sftp - SSH/SFTP

사용법:
   singularity datasource add sftp [command options] <데이터셋_이름> <소스_경로>

설명:
   --sftp-ask-password
      SFTP 암호를 필요할 때 묻도록 허용합니다.
      
      이 플래그가 설정되고 비밀번호가 제공되지 않은 경우 rclone이 다음과 같은 동작을 수행합니다:
      - 암호를 묻습니다.
      - ssh 에이전트에 연락하지 않습니다.
      

   --sftp-chunk-size
      업로드 및 다운로드 청크 크기입니다.
      
      이는 SFTP 프로토콜 패킷의 최대 페이로드 크기를 제어합니다.
      RFC는 이를 32768 바이트 (32k)로 제한합니다. 그러나
      많은 서버가 큰 크기를 지원하며, 일반적으로 최대
      256k의 총 패킷 크기로 제한됩니다. 큰 크기로 설정하면
      높은 지연 시간 링크에서 전송 속도가 크게 향상됩니다. OpenSSH를 포함하여
      예를 들어 255k의 값을 사용하면 256k의 총 패킷 크기 내에서
      오버헤드 공간을 충분히 확보 할 수 있습니다.
      
      32k보다 큰 값을 사용하기 전에 충분히 테스트하고
      항상 동일한 서버에 연결하거나 충분히 넓은 테스트 후에만 사용하세요.
      더 큰 파일을 복사 할 때 "failed to send packet payload: EOF", "connection lost"가
      많이 발생하거나 "corrupted on transfer"가 많이 발생하면
      값을 줄여보세요. [rclone serve sftp]에서 실행되는 서버는
      표준 32k 최대 페이로드로 패킷을 전송하므로 다른 chunk_size를 설정하지 말아야합니다.
      그러나 업로드의 경우 256k를 위한 OpenSSH 예제와 동일한
      chunk_size를 설정할 수 있습니다.
      

   --sftp-ciphers
      우선 순위에 따라 세션 암호화에 사용되는 암호의 공백으로 구분 된 목록입니다.
      
      최소한 하나는 서버 구성과 일치해야합니다. 예를 들어 ssh -Q cipher를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우 이 플래그를 설정하지 않아야합니다.
      
      예:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --sftp-concurrency
      하나의 파일에 대한 최대 대기 중인 요청 수입니다.
      
      이는 하나의 파일에 대한 최대 대기 중인 요청 수를 제어합니다.
      이를 증가시키면 높은 지연 시간 링크에서 처리량이 증가하지만
      더 많은 메모리가 필요합니다.
      

   --sftp-disable-concurrent-reads
      동시 읽기를 사용하지 않도록 설정합니다.
      
      일반적으로 동시 읽기는 안전하게 사용할 수 있으며 사용하지 않으면
      성능이 저하됩니다. 따라서 기본적으로이 옵션은 비활성화되어 있습니다.
      
      일부 서버는 파일을
      다운로드 할 수있는 횟수를 제한합니다. 동시 읽기를 사용하면이 한계를 트리거 할 수 있으므로
      다음과 같이 서버가 반환하는 경우
      
          Failed to copy: file does not exist
      
      이 플래그를 활성화해야 할 수 있습니다.
      
      동시 읽기가 비활성화되면 use_fstat 옵션이 무시됩니다.
      

   --sftp-disable-concurrent-writes
      동시 쓰기를 사용하지 않도록 설정합니다.
      
      일반적으로 rclone은 동시 쓰기를 사용하여 파일을 업로드합니다. 이것은
      성능을 크게 향상시킵니다. 특히 먼 서버에서는
      이 옵션은 필요한 경우 동시 쓰기를 비활성화합니다.
      

   --sftp-disable-hashcheck
      원격 파일 해시 확인을 위해 SSH 명령 실행을 비활성화합니다.
      
      해싱을 사용하려면 비워 두거나 false로 설정하십시오 (권장).
      해싱을 비활성화하려면 true로 설정하십시오.

   --sftp-host
      연결할 SSH 호스트입니다.
      
      예 : "example.com".

   --sftp-idle-timeout
      비활성 연결을 닫기 전의 최대 시간입니다.
      
      지정된 시간 동안 연결 풀에 연결이 반환되지 않은 경우, rclone은 연결 풀을 비웁니다.
      
      0으로 설정하여 연결을 계속 유지하십시오.
      

   --sftp-key-exchange
      우선 순위에 따라 공백으로 구분 된 키 교환 알고리즘의 목록입니다.
      
      최소한 하나는 서버 구성과 일치해야합니다. 예를 들어 ssh -Q kex를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우이를 설정하면 안됩니다.
      
      예:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --sftp-key-file
      PEM로 인코딩된 개인 키 파일의 경로입니다.
      
      키를 사용할 경우 비워 두거나 key-use-agent를 설정하세요.
      
      파일 이름에서 '~'는 확장되며 '${RCLONE_CONFIG_DIR}'와 같은 환경 변수도 확장됩니다.

   --sftp-key-file-pass
      PEM로 암호화 된 개인 키 파일을 복호화하는 비밀번호입니다.
      
      새로운 OpenSSH 형식의 암호화 된 키는 사용할 수 없습니다.

   --sftp-key-pem
      원시 PEM으로 인코딩 된 개인 키입니다.
      
      지정된 경우 key_file 매개 변수를 무시합니다.

   --sftp-key-use-agent
      ssh-agent의 사용을 강제로 요청합니다.
      
      key-file도 설정된 경우 지정된 key-file의 ".pub" 파일이 읽혀지고 `associated key` 과 관련된 것만
      ssh-agent에서 요청됩니다. ssh-agent에 많은 키가 포함된 경우
      'Too many authentication failures for *username*' 오류를 피할 수 있습니다.

   --sftp-known-hosts-file
      알려진 호스트 파일의 선택적 경로입니다.
      
      서버 호스트 키 유효성을 검사하려면이 값을 설정하세요.
      
      파일 이름에서 '~'는 확장되며 '${RCLONE_CONFIG_DIR}'와 같은 환경 변수도 확장됩니다.

      예제:
         | ~/.ssh/known_hosts | OpenSSH의 알려진 호스트 파일 사용

   --sftp-macs
      우선 순위에 따라 공백으로 구분 된 MAC (메시지 인증 코드) 알고리즘 목록입니다.
      
      최소한 하나는 서버 구성과 일치해야합니다. 예를 들어 ssh -Q mac을 사용하여 확인할 수 있습니다.
      
      예:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      

   --sftp-md5sum-command
      md5 해시를 읽기 위해 사용되는 명령입니다.
      
      자동 감지하려면 비워 두세요.

   --sftp-pass
      SSH 암호입니다. ssh-agent를 사용하려면 비워 두세요.

   --sftp-path-override
      SSH 쉘 명령에 사용되는 경로를 재정의합니다.
      
      SFTP 및 SSH 경로가 다른 경우 체크섬 계산에 영향을줍니다. 이 문제는 다음을 포함하여 기타 문제들에 영향을줍니다.
      Synology NAS 상자.
      
      예를 들어 공유 폴더가 볼륨을 나타내는 디렉토리 안에 있다면:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      예를 들어 홈 디렉토리가 "home"이라는 공유 폴더에 있다면:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --sftp-port
      SSH 포트 번호입니다.

   --sftp-pubkey-file
      공개 키 파일의 선택적 경로입니다.
      
      인증에 사용할 서명 된 인증서가 있는 경우에 설정합니다.
      
      파일 이름에서 '~'는 확장되며 '${RCLONE_CONFIG_DIR}'와 같은 환경 변수도 확장됩니다.

   --sftp-server-command
      원격 호스트에서 sftp 서버를 실행할 경로 또는 명령을 지정합니다.
      
      server_command이 정의 된 경우 subsystem 옵션은 무시됩니다.

   --sftp-set-env
      sftp 및 명령에 전달할 환경 변수
      
      다음 양식으로 환경 변수를 설정하세요:
      
          VAR=value
      
      sftp 클라이언트 및 실행되는 모든 명령에 전달됩니다 (예 : md5sum).
      
      여러 변수를 공백으로 구분하여 전달하려면 다음과 같이 입력하십시오.
      
          VAR1=value VAR2=value
      
      공백이있는 변수는 따옴표로 묶으세요.
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --sftp-set-modtime
      설정된 경우 원격에서 수정된 시간을 설정합니다.

   --sftp-sha1sum-command
      sha1 해시를 읽기 위해 사용되는 명령입니다.
      
      자동 감지하려면 비워 두세요.

   --sftp-shell-type
      원격 서버의 SSH 쉘 유형 (있는 경우)입니다.
      
      자동으로 감지하려면 비워 두세요.

      예:
         | none       | 쉘 액세스가 없음
         | unix       | UNIX 쉘
         | powershell | PowerShell
         | cmd        | Windows 명령 프롬프트

   --sftp-skip-links
      시링크 및 기타 정규 파일 이외의 파일을 건너 뛰도록 설정합니다.

   --sftp-subsystem
      원격 호스트의 SSH2 하위 시스템을 지정합니다.

   --sftp-use-fstat
      fstat 대신 stat을 사용합니다.
      
      일부 서버는 오픈 된 파일의 양을 제한하고 파일을 열은 후에
      Stat을 호출하면 서버에서 오류가 발생합니다. 이 플래그를 설정하면
      이미 열려있는 파일 핸들에서 호출되는 Fstat을 호출합니다.
      
      IBM Sterling SFTP 서버에 도움이되는 것으로 밝혀진 바 있으며,
      '추출 가능성' 수준이 1로 설정되어 있으므로 어떤 시간에 어떤 파일도
      열 수 있습니다.
      

   --sftp-use-insecure-cipher
      보안이 취약한 암호 및 키 교환 방법의 사용을 활성화합니다.
      
      다음의 보안이 취약한 암호 및 키 교환 방법의 사용을 활성화합니다.
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      이러한 알고리즘은 보안 상 취약하며 공격자에게 평문 데이터를 복구 할 수 있을 수 있습니다.
      
      이는 ciphers 또는 key_exchange 고급 옵션을 사용하는 경우 false 여야합니다.
      

      예:
         | false | 기본 암호화 목록 사용
         | true  | aes128-cbc 암호 및 diffie-hellman-group-exchange-sha256, diffie-hellman-group-exchange-sha1 키 교환 사용

   --sftp-user
      SSH 사용자 이름.


옵션:
   --help, -h  도움말 표시

   데이터 준비 옵션

   --delete-after-export    [주의] CAR 파일로 데이터셋을 내보낸 후 파일을 삭제합니다.  (기본값: false)
   --rescan-interval value  마지막 성공적인 스캔으로부터 경과된 이후 시간이면 자동으로 소스 디렉토리를 재스캔합니다 (기본값: 사용안함)
   --scanning-state value   초기 스캔 상태를 설정합니다 (기본값: 준비)

   sftp 옵션

   --sftp-ask-password value               필요할 때 SFTP 비밀번호를 묻습니다. (기본값: "false") [$SFTP_ASK_PASSWORD]
   --sftp-chunk-size value                 업로드 및 다운로드 청크 크기를 지정합니다. (기본값: "32Ki") [$SFTP_CHUNK_SIZE]
   --sftp-ciphers value                    우선 순위에 따라 세션 암호화에 사용되는 암호의 공백으로 구분 된 목록입니다. [$SFTP_CIPHERS]
   --sftp-concurrency value                하나의 파일에 대한 최대 대기 중인 요청 수입니다. (기본값: "64") [$SFTP_CONCURRENCY]
   --sftp-disable-concurrent-reads value   동시 읽기를 사용하지 않도록 합니다. (기본값: "false") [$SFTP_DISABLE_CONCURRENT_READS]
   --sftp-disable-concurrent-writes value  동시 쓰기를 사용하지 않도록 합니다. (기본값: "false") [$SFTP_DISABLE_CONCURRENT_WRITES]
   --sftp-disable-hashcheck value          원격 파일 해싱 가능 여부를 결정하기 위한 SSH 명령 실행을 비활성화합니다. (기본값: "false") [$SFTP_DISABLE_HASHCHECK]
   --sftp-host value                       연결할 SSH 호스트입니다. [$SFTP_HOST]
   --sftp-idle-timeout value               비활성 연결을 닫기 전의 최대 시간입니다. (기본값: "1m0s") [$SFTP_IDLE_TIMEOUT]
   --sftp-key-exchange value               우선 순위에 따라 공백으로 구분 된 키 교환 알고리즘의 목록입니다. [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value                   PEM로 인코딩된 개인 키 파일의 경로입니다. [$SFTP_KEY_FILE]
   --sftp-key-file-pass value              PEM로 암호화 된 개인 키 파일을 복호화하는 비밀번호입니다. [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value                    원시 PEM으로 인코딩 된 개인 키입니다. [$SFTP_KEY_PEM]
   --sftp-key-use-agent value              ssh-agent의 사용을 강제로 요청합니다. (기본값: "false") [$SFTP_KEY_USE_AGENT]
   --sftp-known-hosts-file value           알려진 호스트 파일의 선택적 경로입니다. [$SFTP_KNOWN_HOSTS_FILE]
   --sftp-macs value                       우선 순위에 따라 공백으로 구분 된 MAC (메시지 인증 코드) 알고리즘 목록입니다. [$SFTP_MACS]
   --sftp-md5sum-command value             md5 해시를 읽기 위해 사용되는 명령입니다. [$SFTP_MD5SUM_COMMAND]
   --sftp-pass value                       SSH 암호입니다. ssh-agent를 사용하려면 비워 두세요. [$SFTP_PASS]
   --sftp-path-override value              SSH 쉘 명령에 사용되는 경로를 재정의합니다. [$SFTP_PATH_OVERRIDE]
   --sftp-port value                       SSH 포트 번호입니다. (기본값: "22") [$SFTP_PORT]
   --sftp-pubkey-file value                공개 키 파일의 선택적 경로입니다. [$SFTP_PUBKEY_FILE]
   --sftp-server-command value             원격 호스트에서 sftp 서버를 실행할 경로 또는 명령을 지정합니다. [$SFTP_SERVER_COMMAND]
   --sftp-set-env value                    sftp 및 명령에 전달할 환경 변수 [$SFTP_SET_ENV]
   --sftp-set-modtime value                설정된 경우 원격에서 수정된 시간을 설정합니다. (기본값: "true") [$SFTP_SET_MODTIME]
   --sftp-sha1sum-command value            sha1 해시를 읽기 위해 사용되는 명령입니다. [$SFTP_SHA1SUM_COMMAND]
   --sftp-shell-type value                 원격 서버의 SSH 쉘 유형 (있는 경우)입니다. [$SFTP_SHELL_TYPE]
   --sftp-skip-links value                 시링크 및 기타 정규 파일 이외의 파일을 건너 뜁니다. (기본값: "false") [$SFTP_SKIP_LINKS]
   --sftp-subsystem value                  원격 호스트의 SSH2 하위 시스템을 지정합니다. (기본값: "sftp") [$SFTP_SUBSYSTEM]
   --sftp-use-fstat value                  fstat 대신 stat을 사용합니다. (기본값: "false") [$SFTP_USE_FSTAT]
   --sftp-use-insecure-cipher value        보안이 취약한 암호 및 키 교환 방법의 사용을 활성화합니다. (기본값: "false") [$SFTP_USE_INSECURE_CIPHER]
   --sftp-user value                       SSH 사용자 이름입니다. (기본값: "$USER") [$SFTP_USER]

```
{% endcode %}