# SSH/SFTP

{% code fullWidth="true" %}
```
명령:
   singularity storage update sftp - SSH/SFTP

사용법:
   singularity storage update sftp [command options] <name|id>

설명:
   --host
      연결할 SSH 호스트입니다.
      
      예: "example.com".

   --user
      SSH 사용자 이름입니다.

   --port
      SSH 포트 번호입니다.

   --pass
      SSH 비밀번호입니다. ssh-agent를 사용하려면 비워 두십시오.

   --key-pem
      PEM 인코딩된 개인 키입니다.
      
      지정하면 key_file 매개변수를 무시합니다.

   --key-file
      PEM 인코딩된 개인 키 파일의 경로입니다.
      
      비워 두거나 key-use-agent를 설정하여 ssh-agent를 사용하십시오.
      
      "~"로 시작하는 파일 이름이나 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수는 확장됩니다.

   --key-file-pass
      PEM 인코딩된 개인 키 파일을 복호화하는 패스워드입니다.
      
      새로운 OpenSSH 형식의 암호화된 키를 사용할 수 없으며 예전 OpenSSH 형식의 키만 지원됩니다.

   --pubkey-file
      공개 키 파일의 경로입니다.
      
      인증에 사용할 서명된 인증서가 있는 경우 지정하세요.
      
      "~"로 시작하는 파일 이름이나 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수는 확장됩니다.

   --known-hosts-file
      known_hosts 파일의 경로입니다.
      
      이 값을 설정하여 서버 호스트 키 유효성 검사를 활성화합니다.
      
      "~"로 시작하는 파일 이름이나 `${RCLONE_CONFIG_DIR}`와 같은 환경 변수는 확장됩니다.

      예제:
         | ~/.ssh/known_hosts | OpenSSH의 known_hosts 파일 사용

   --key-use-agent
      ssh-agent 사용을 강제로 설정합니다.
      
      key-file도 설정되어 있는 경우, 지정된 key-file의 ".pub" 파일을 읽고 연결된 키만 ssh-agent에서 요청합니다. 이렇게 하면 ssh-agent에 많은 키가 포함되어 있을 때 `Too many authentication failures for *username*` 오류를 피할 수 있습니다.

   --use-insecure-cipher
      (보안에 안전하지 않음) 암호화 및 키 교환 방법에 보안에 안전하지 않은 암호와 키 교환 방법의 사용을 활성화합니다.
      
      다음 보안에 안전하지 않은 암호와 키 교환 방법을 사용할 수 있게 됩니다:
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      이러한 알고리즘은 보안에 취약할 수 있으며, 공격자가 평문 데이터를 복구할 수 있을 수도 있습니다.
      
      만약 암호나 key_exchange 고급 옵션을 사용하는 경우, 이 값을 false로 설정해야 합니다.
      

      예제:
         | false | 기본 암호화 목록 사용
         | true  | aes128-cbc 암호와 diffie-hellman-group-exchange-sha256, diffie-hellman-group-exchange-sha1 키 교환 사용

   --disable-hashcheck
      원격 파일 해싱의 사용가능 여부를 확인하기 위해 SSH 명령 실행 비활성화합니다.
      
      해싱 사용을 활성화하려면 비워 두거나 false로 설정하고, 해싱을 비활성화하려면 true로 설정하세요.

   --ask-password
      SFTP 암호를 요청할 수 있도록 허용합니다.
      
      이 값을 설정하고 암호를 제공하지 않으면 rclone은 다음과 같은 동작을 합니다:
      - 암호를 요청합니다.
      - ssh 에이전트에 문의하지 않습니다.
      

   --path-override
      SSH 쉘 명령에서 사용할 경로를 재정의합니다.
      
      이렇게 하면 SFTP와 SSH 경로가 다를 때 체크섬 계산이 가능합니다. 이 문제는 Synology NAS 상자 등에 영향을 줍니다.
      
      예를 들어, 공유 폴더가 볼륨을 나타내는 디렉토리에 있다고 가정합니다:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      다른 예를 들어, 홈 디렉토리가 "home"이라는 공유 폴더에 있다고 가정합니다:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      변경된 시간이 있는 경우 원격에 변경된 시간을 설정합니다.

   --shell-type
      원격 서버의 SSH 쉘 유형입니다. (자동 감지하지 않으려면 비워 둡니다.)

      예제:
         | none       | 쉘 접근 불가
         | unix       | Unix 쉘
         | powershell | PowerShell
         | cmd        | Windows 명령 프롬프트

   --md5sum-command
      md5 해시를 읽는 데 사용되는 명령입니다.
      
      자동 감지하려면 비워 두십시오.

   --sha1sum-command
      sha1 해시를 읽는 데 사용되는 명령입니다.
      
      자동 감지하려면 비워 두십시오.

   --skip-links
      심볼릭 링크 및 정규 파일이 아닌 파일을 건너뜁니다.

   --subsystem
      원격 호스트의 SSH2 서브시스템을 지정합니다.

   --server-command
      원격 호스트에서 SFTP 서버를 실행하기 위한 경로 또는 명령을 지정합니다.
      
      server_command가 정의되어 있는 경우 서브시스템 옵션이 무시됩니다.

   --use-fstat
      fstat 대신에 사용하여 stat을 사용합니다.
      
      일부 서버는 열린 파일의 수를 제한하며 파일을 열고 난 후에 Stat을 호출하면 서버에서 오류가 발생합니다. 이 플래그를 설정하면 이미 열린 파일 핸들에서 호출되는 Fstat을 사용합니다.
      
      "extractability" 레벨이 1로 설정된 IBM Sterling SFTP 서버의 경우라고 알려진 서버에서 도움이 된다고 알려져 있습니다. 이 레벨은 한 번에 1개의 파일만 열 수 있는 것을 의미합니다.
      

   --disable-concurrent-reads
      동시 읽기를 사용하지 않도록 설정하십시오.
      
      일반적으로 동시 읽기는 사용해도 안전하며 사용하지 않으면 성능이 감소합니다. 따라서 이 옵션은 기본적으로 비활성화되어 있습니다.
      
      일부 서버는 파일을 다운로드할 수 있는 횟수를 제한합니다. 동시 읽기를 사용하면 이 제한을 활성화할 수 있습니다. 그러므로 다음과 같은 오류 메시지를 받는 서버가 있는 경우:
      
          Failed to copy: file does not exist
      
      이 플래그를 활성화해야 할 수 있습니다.
      
      동시 읽기가 비활성화되면 use_fstat 옵션은 무시됩니다.
      

   --disable-concurrent-writes
      동시 쓰기를 사용하지 않도록 설정하십시오.
      
      일반적으로 rclone은 파일을 업로드하기 위해 동시 쓰기를 사용합니다. 이렇게 하면 성능이 크게 향상되며 특히 먼 서버의 경우에 유용합니다.
      
      필요한 경우 이 옵션을 사용하여 동시 쓰기를 비활성화시킵니다.
      

   --idle-timeout
      비활성 연결을 닫기 전의 최대 시간입니다.
      
      주어진 시간 동안 연결이 연결 풀로 반환되지 않으면 rclone은 연결 풀을 비웁니다.
      
      연결을 계속 유지하려면 0으로 설정하세요.
      

   --chunk-size
      업로드 및 다운로드 청크 크기입니다.
      
      이 값은 SFTP 프로토콜 패킷에서 페이로드의 최대 크기를 제어합니다.
      RFC는 이 값을 32768바이트(32k)로 제한하지만 많은 서버는 더 큰 크기를 지원합니다. 일반적으로 최대 패킷 크기가 256k로 제한되며, 이 값을 크게 설정하면 고지연 링크에서 전송 속도가 크게 향상됩니다. 이것에는 OpenSSH도 포함되며, 예를 들어 OpenSSH에서 255k 값을 사용하면 충분한 여유 공간을 남기면서 패킷 크기는 총 256k가 됩니다.
      
      32k보다 큰 값으로 설정하기 전에 충분한 테스트를 수행한 후에 사용하세요. "failed to send packet payload: EOF", "connection lost" 또는 "corrupted on transfer"와 같은 오류가 발생하는 경우 큰 파일을 복사할 때, 이 값을 낮추세요. [rclone serve sftp](/commands/rclone_serve_sftp)로 실행되는 서버는 표준 32k 최대 페이로드를 갖는 패킷을 보내므로 다른 chunk_size를 설정해서는 안되지만, 최대 256k의 패킷까지 수용합니다. 따라서 업로드에 대해서는 위의 OpenSSH 예제와 동일한 chunk_size를 설정할 수 있습니다.
      

   --concurrency
      한 파일에 대한 최대 대기 중인 요청 수입니다.
      
      이 값은 한 파일에 대한 최대 대기 중인 요청 수를 제어합니다. 이 값을 늘리면 고지연 링크에서 처리량이 더욱 향상되지만 메모리를 더 사용합니다.
      

   --set-env
      sftp와 명령에 전달할 환경 변수들입니다.
      
      다음 형식으로 환경 변수를 설정하세요:
      
          VAR=value
      
      sftp 클라이언트와 실행되는 모든 명령(md5sum과 같음)에 전달됩니다.
      
      여러 변수를 공백으로 구분하여 전달할 수 있습니다. 예를 들어:
      
          VAR1=value VAR2=value
      
      변수에는 공백이 있는 경우 따옴표(")로 묶어 전달하세요. 예를 들어:
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --ciphers
      암호화 세션에 사용할 암호의 우선순위도를 나타내는 공백으로 분리된 암호 목록입니다.
      
      적어도 하나는 서버의 구성과 일치해야 합니다. 예를 들어 ssh -Q cipher를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우 이 값을 설정하지 마십시오.
      
      예제:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --key-exchange
      공백으로 분리된 우선순위별 키 교환 알고리즘 목록입니다.
      
      적어도 하나는 서버의 구성과 일치해야 합니다. 예를 들어 ssh -Q kex를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우 이 값을 설정하지 마십시오.
      
      예제:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --macs
      공백으로 분리된 우선순위별 MACs(메시지 인증 코드) 알고리즘 목록입니다.
      
      적어도 하나는 서버의 구성과 일치해야 합니다. 예를 들어 ssh -Q mac을 사용하여 확인할 수 있습니다.
      
      예제:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      


OPTIONS:
   --disable-hashcheck    원격 파일 해싱의 사용가능 여부를 확인하기 위해 SSH 명령 실행 비활성화. (default: false) [$DISABLE_HASHCHECK]
   --help, -h             도움말 표시
   --host value           연결할 SSH 호스트. [$HOST]
   --key-file value       PEM 인코딩된 개인 키 파일의 경로. [$KEY_FILE]
   --key-file-pass value  PEM 인코딩된 개인 키 파일을 복호화하는 패스워드. [$KEY_FILE_PASS]
   --key-pem value        PEM 인코딩된 개인 키. [$KEY_PEM]
   --key-use-agent        ssh-agent의 사용 강제화. (default: false) [$KEY_USE_AGENT]
   --pass value           SSH 비밀번호, ssh-agent를 사용하려면 비워 두세요. [$PASS]
   --port value           SSH 포트 번호. (default: 22) [$PORT]
   --pubkey-file value    공개 키 파일의 경로. [$PUBKEY_FILE]
   --use-insecure-cipher  (보안에 안전하지 않음) 보안에 안전하지 않은 암호와 키 교환 방법의 사용 활성화. (default: false) [$USE_INSECURE_CIPHER]
   --user value           SSH 사용자 이름. (default: "$USER") [$USER]

   Advanced

   --ask-password               SFTP 암호를 요청할 수 있도록 허용. (default: false) [$ASK_PASSWORD]
   --chunk-size value           업로드 및 다운로드 청크 크기. (default: "32Ki") [$CHUNK_SIZE]
   --ciphers value              암호화 세션에 사용할 암호의 우선순위도를 나타내는 공백으로 분리된 암호 목록. [$CIPHERS]
   --concurrency value          한 파일에 대한 최대 대기 중인 요청 수. (default: 64) [$CONCURRENCY]
   --disable-concurrent-reads   동시 읽기를 사용하지 않도록 설정. (default: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  동시 쓰기를 사용하지 않도록 설정. (default: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         비활성 연결을 닫기 전의 최대 시간. (default: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         공백으로 분리된 우선순위별 키 교환 알고리즘 목록. [$KEY_EXCHANGE]
   --known-hosts-file value     known_hosts 파일의 경로. [$KNOWN_HOSTS_FILE]
   --macs value                 공백으로 분리된 우선순위별 MACs(메시지 인증 코드) 알고리즘 목록. [$MACS]
   --md5sum-command value       md5 해시를 읽는 데 사용되는 명령. [$MD5SUM_COMMAND]
   --path-override value        SSH 쉘 명령에서 사용할 경로를 재정의. [$PATH_OVERRIDE]
   --server-command value       원격 호스트에서 SFTP 서버를 실행하기 위한 경로 또는 명령. [$SERVER_COMMAND]
   --set-env value              sftp와 명령에 전달할 환경 변수들. [$SET_ENV]
   --set-modtime                변경된 시간이 있는 경우 원격에 변경된 시간을 설정. (default: true) [$SET_MODTIME]
   --sha1sum-command value      sha1 해시를 읽는 데 사용되는 명령. [$SHA1SUM_COMMAND]
   --shell-type value           원격 서버의 SSH 쉘 유형. [$SHELL_TYPE]
   --skip-links                 심볼릭 링크 및 정규 파일이 아닌 파일을 건너뜁니다. (default: false) [$SKIP_LINKS]
   --subsystem value            원격 호스트의 SSH2 서브시스템을 지정합니다. (default: "sftp") [$SUBSYSTEM]
   --use-fstat                  fstat 대신에 사용하여 stat을 사용. (default: false) [$USE_FSTAT]

```
{% endcode %}