# SSH/SFTP

{% code fullWidth="true" %}
```
이름:
   singularity storage create sftp - SSH/SFTP

사용법:
   singularity storage create sftp [command options] [arguments...]

설명:
   --host
      연결할 SSH 호스트입니다.
      
      예시: "example.com".

   --user
      SSH 사용자 이름입니다.

   --port
      SSH 포트 번호입니다.

   --pass
      SSH 암호입니다. ssh-agent를 사용하려면 비워 두십시오.

   --key-pem
      PEM 인코딩된 개인 키입니다.
      
      지정된 경우 key_file 매개변수를 무시합니다.

   --key-file
      PEM 인코딩된 개인 키 파일의 경로입니다.
      
      비워 두거나 key-use-agent를 설정하여 ssh-agent를 사용하십시오.
      
      `~`로 시작하는 파일 이름은 파일 이름이 확장되고 `${RCLONE_CONFIG_DIR}`과 같은 환경 변수도 확장됩니다.

   --key-file-pass
      PEM 인코딩된 개인 키 파일을 복호화하기 위한 암호입니다.
      
      사용가능한 암호화된 PEM 키 파일은 오래된 OpenSSH 형식입니다. 새로운 OpenSSH 형식의 암호화된 키는 사용할 수 없습니다.

   --pubkey-file
      선택적인 공개 키 파일의 경로입니다.
      
      인증에 사용할 서명된 인증서가 있는 경우에만 설정하십시오.
      
      `~`로 시작하는 파일 이름은 파일 이름이 확장되고 `${RCLONE_CONFIG_DIR}`과 같은 환경 변수도 확장됩니다.

   --known-hosts-file
      선택적인 known_hosts 파일의 경로입니다.
      
      이 값을 설정하여 서버 호스트 키 유효성 검사를 활성화할 수 있습니다.
      
      `~`로 시작하는 파일 이름은 파일 이름이 확장되고 `${RCLONE_CONFIG_DIR}`과 같은 환경 변수도 확장됩니다.

      예시:
         | ~/.ssh/known_hosts | OpenSSH의 known_hosts 파일 사용.

   --key-use-agent
      ssh-agent의 사용을 강제하는 경우 설정합니다.
      
      key-file도 설정되어 있는 경우, 지정된 key-file의 ".pub" 파일을 읽고 해당 키만 ssh-agent에서 요청합니다.
      이를 통해 ssh-agent에 여러 개의 키가 있는 경우 'Too many authentication failures for *username*' 오류를 피할 수 있습니다.

   --use-insecure-cipher
      보안을 가지지 않은 암호와 키 교환 방법을 사용하도록 설정합니다.
      
      다음의 보안을 가지지 않은 암호와 키 교환 방법을 사용하도록 설정합니다:
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      이러한 알고리즘은 보안이 취약하며 공격자가 평문 데이터를 복구할 수 있을 수 있습니다.
      
      ciphers나 key_exchange의 고급 옵션을 사용하는 경우이 값을 false로 설정해야합니다.
      

      예시:
         | false | 기본 Cipher 목록 사용.
         | true  | aes128-cbc 암호 및 diffie-hellman-group-exchange-sha256, diffie-hellman-group-exchange-sha1 키 교환 사용.

   --disable-hashcheck
      원격 파일 해싱 가능 여부를 판별하기 위해 SSH 명령 실행 비활성화.
      
      해싱을 활성화하려면 비워두거나 false로 설정하고, 해싱을 비활성화하려면 true로 설정합니다.

   --ask-password
      필요할 때 SFTP 암호를 요청할 수 있도록 허용합니다.
      
      이 값이 설정되고 암호가 제공되지 않은 경우 rclone은 다음을 수행합니다:
      - 암호를 요청함
      - ssh 에이전트에 연결하지 않음
      

   --path-override
      SSH 셸 명령에서 사용할 경로를 덮어씁니다.
      
      이로써 SFTP 및 SSH 경로가 다를 때 체크섬 계산이 가능합니다. 이 문제는 Synology NAS 상자를 비롯한 다른 상자들에 영향을 미칩니다.
      
      예시로 공유 폴더가 볼륨을 표현하는 디렉토리에 존재하는 경우:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      예시로 홈 디렉터리가 "home"이라는 공용 폴더에 존재하는 경우:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      설정된 경우 원격지에서 수정된 시간을 설정합니다.

   --shell-type
      원격 서버의 SSH 셸 유형입니다. 있을 경우 설정하세요.
      
      자동 감지하려면 비워 두세요.

      예시:
         | none       | 셸 액세스 없음
         | unix       | Unix 셸
         | powershell | PowerShell
         | cmd        | Windows 명령 프롬프트

   --md5sum-command
      md5 해시를 읽는 데 사용되는 명령입니다.
      
      자동 감지하려면 비워 두세요.

   --sha1sum-command
      sha1 해시를 읽는 데 사용되는 명령입니다.
      
      자동 감지하려면 비워 두세요.

   --skip-links
      심볼릭 링크 및 일반 파일 외의 파일을 건너뛰도록 설정합니다.

   --subsystem
      원격 호스트의 SSH2 서브시스템을 지정합니다.

   --server-command
      원격 호스트에서 SFTP 서버를 실행할 경로 또는 명령을 지정합니다.
      
      server_command가 정의되어 있는 경우 subsystem 옵션은 무시됩니다.

   --use-fstat
      파일을 열었을 때 stat 대신 fstat을 사용합니다.
      
      일부 서버는 열 수있는 파일의 양을 제한하며 파일을 열고 나서 Stat을 호출하면 서버에서 오류가 발생합니다. 이 플래그를 설정하면 이미 열려있는 파일 핸들에서 호출되는 Fstat을 호출합니다.
      
      IBM Sterling SFTP 서버에서 "extractability" 수준이 1로 설정되어 있어 한 번에 파일을 1개만 열 수 있는 것으로 확인되었습니다.

   --disable-concurrent-reads
      동시 읽기를 사용하지 않도록 설정하려면 설정하십시오.
      
      일반적으로 동시 읽기는 안전하게 사용할 수 있으며 사용하지 않으면 성능이 저하됩니다. 이 옵션은기본적으로 사용되지 않도록 비활성화되어 있습니다.
      
      일부 서버는 한 파일을 다운로드 할 수있는 횟수를 제한합니다. 동시 읽기를 사용하면이 제한이 트리거될 수 있으므로 다음과 같은 오류 메시지가 표시될 수 있습니다:
      
          Failed to copy: file does not exist
      
      동시 읽기가 비활성화 된 경우 use_fstat 옵션이 무시됩니다.
      

   --disable-concurrent-writes
      동시 쓰기를 사용하지 않도록 설정하려면 설정하십시오.
      
      일반적으로 rclone은 파일을 업로드하기 위해 동시쓰기를 사용합니다. 이는 특히 먼 서버의 경우 성능을 크게 향상시킵니다.
      
      필요한 경우이 옵션을 사용하여 동시 쓰기를 비활성화할 수 있습니다.
      

   --idle-timeout
      유휴 연결을 닫기 전의 최대 시간입니다.
      
      지정한 시간 동안 연결 풀에 반환된 연결이 없으면 rclone은 연결 풀을 비웁니다.
      
      연결을 계속 유지하려면 0으로 설정하십시오.
      

   --chunk-size
      업로드 및 다운로드 청크 크기입니다.
      
      이것은 SFTP 프로토콜 패킷의 최대 페이로드 크기를 제어합니다.
      RFC는 이 값을 32768바이트(32k)로 제한합니다. 그러나 많은 서버가 큰 크기를 지원하며 일반적으로 총 패키지 크기가 256k로 제한됩니다. 크기를 더 크게 설정하면 고지연 링크에서 전송 속도가 급격히 향상됩니다. 이에는 OpenSSH도 포함되며, 예를 들어 값이 255k인 경우에는 여유 공간이 충분하고 256k의 총 패킷 크기 내에 있는 상태로 작동합니다.
      
      32k보다 높은 값을 사용하기 전에 꼼꼼하게 테스트하고 항상 동일한 서버에 연결하거나 충분한 범위의 테스트 후에만 사용하십시오. 큰 파일을 복사할 때 'failed to send packet payload: EOF' 및 'connection lost' 또는 'corrupted on transfer'와 같은 오류가 발생하는 경우 값이 낮아져야합니다. [rclone serve sftp](/commands/rclone_serve_sftp)로 실행되는 서버는 표준 32k 최대 페이로드의 패킷을 보내므로 다른 chunk_size를 설정해서는 안되지만, 256k까지의 패킷은 수용하므로 업로드의 경우 위 OpenSSH 예에 설정할 수 있습니다.
      

   --concurrency
      하나의 파일에 대한 최대 대기 요청 수를 제어합니다.
      
      이것은 하나의 파일에 대한 최대 대기 요청 수를 제어합니다.
      이 값을 늘리면 고지연 링크에서 처리량이 증가하지만 더 많은 메모리를 사용합니다.
      

   --set-env
      sftp 및 명령에 전달 할 환경 변수입니다.
      
      다음 형식으로 환경 변수를 설정합니다:
      
          VAR=value
      
      sftp 클라이언트 및 실행되는 명령 (예: md5sum) 에 전달되는 변수를 설정하기 위해 여러 변수를 공백으로 구분하여 전달합니다.
      
          VAR1=value VAR2=value
      
      변수에 공백이 포함된 경우 인용 부호로 변수를 전달하십시오.
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --ciphers
      선호하는 순서로 세션 암호화에 사용할 암호 목록입니다. 공백으로 구분됩니다.
      
      적어도 한 개의 서버 구성과 일치해야 합니다. 예를 들어 ssh -Q cipher를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우이 값을 설정하지 않아야 합니다.
      
      예시:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --key-exchange
      선호하는 순서대로 키 교환 알고리즘의 목록입니다. 공백으로 구분됩니다.
      
      적어도 한 개의 서버 구성과 일치해야 합니다. 예를 들어 ssh -Q kex를 사용하여 확인할 수 있습니다.
      
      use_insecure_cipher가 true로 설정된 경우이 값을 설정하지 않아야 합니다.
      
      예시:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --macs
      선호하는 순서로 MAC (메시지 인증 코드) 알고리즘의 목록입니다. 공백으로 구분됩니다.
      
      적어도 한 개의 서버 구성과 일치해야 합니다. 예를 들어 ssh -Q mac를 사용하여 확인할 수 있습니다.
      
      예시:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      


OPTIONS:
   --disable-hashcheck    원격 파일 해싱 가능 여부를 판별하기 위해 SSH 명령 실행 비활성화 (기본값: false) [$DISABLE_HASHCHECK]
   --help, -h             도움말 표시
   --host value           연결할 SSH 호스트입니다. [$HOST]
   --key-file value       PEM 인코딩된 개인 키 파일의 경로입니다. [$KEY_FILE]
   --key-file-pass value  PEM 인코딩된 개인 키 파일을 복호화하기 위한 암호입니다. [$KEY_FILE_PASS]
   --key-pem value        Raw PEM-encoded private key. [$KEY_PEM]
   --key-use-agent        ssh-agent의 사용을 강제합니다 (기본값: false) [$KEY_USE_AGENT]
   --pass value           SSH 암호, ssh-agent를 사용하려면 비워 두십시오. [$PASS]
   --port value           SSH 포트 번호 (기본값: 22) [$PORT]
   --pubkey-file value    선택적인 공개 키 파일의 경로입니다. [$PUBKEY_FILE]
   --use-insecure-cipher  보안을 가지지 않은 암호와 키 교환 방법을 사용하도록 설정합니다 (기본값: false) [$USE_INSECURE_CIPHER]
   --user value           SSH 사용자 이름 (기본값: "$USER") [$USER]

   고급

   --ask-password               필요할 때 SFTP 암호를 요청할 수 있도록 허용합니다 (기본값: false) [$ASK_PASSWORD]
   --chunk-size value           업로드 및 다운로드 청크 크기 (기본값: "32Ki") [$CHUNK_SIZE]
   --ciphers value              선호하는 순서로 세션 암호화에 사용할 암호 목록입니다. [$CIPHERS]
   --concurrency value          하나의 파일에 대한 최대 대기 요청 수 (기본값: 64) [$CONCURRENCY]
   --disable-concurrent-reads   동시 읽기를 사용하지 않도록 설정하려면 설정하십시오 (기본값: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  동시 쓰기를 사용하지 않도록 설정하려면 설정하십시오 (기본값: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         유휴 연결을 닫기 전의 최대 시간 (기본값: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         선호하는 순서대로 키 교환 알고리즘의 목록입니다. [$KEY_EXCHANGE]
   --known-hosts-file value     선택적인 known_hosts 파일의 경로입니다. [$KNOWN_HOSTS_FILE]
   --macs value                 선호하는 순서로 MAC (메시지 인증 코드) 알고리즘의 목록입니다. [$MACS]
   --md5sum-command value       md5 해시를 읽는 데 사용되는 명령입니다. [$MD5SUM_COMMAND]
   --path-override value        SSH 셸 명령에서 사용할 경로를 덮어씁니다. [$PATH_OVERRIDE]
   --server-command value       원격 호스트에서 SFTP 서버를 실행할 경로 또는 명령을 지정합니다. [$SERVER_COMMAND]
   --set-env value              sftp 및 명령에 전달 할 환경 변수입니다 [$SET_ENV]
   --set-modtime                설정된 경우 원격지에서 수정된 시간을 설정합니다 (기본값: true) [$SET_MODTIME]
   --sha1sum-command value      sha1 해시를 읽는 데 사용되는 명령입니다. [$SHA1SUM_COMMAND]
   --shell-type value           원격 서버의 SSH 셸 유형입니다. [$SHELL_TYPE]
   --skip-links                 심볼릭 링크 및 일반 파일 외의 파일을 건너뛰도록 설정합니다 (기본값: false) [$SKIP_LINKS]
   --subsystem value            원격 호스트의 SSH2 서브시스템을 지정합니다 (기본값: "sftp") [$SUBSYSTEM]
   --use-fstat                  파일을 열었을 때 stat 대신 fstat을 사용합니다 (기본값: false) [$USE_FSTAT]

   일반

   --name value  스토리지의 이름 (기본값: 자동 생성)
   --path value  스토리지의 경로

```
{% endcode %}