# SSH/SFTP

```
NAME:
   singularity datasource add sftp - SSH/SFTP

USAGE:
   singularity datasource add sftp [コマンドオプション] <dataset_name> <source_path>

DESCRIPTION:
   --sftp-ask-password
      SFTPパスワードの入力を許可します。
      
      これが設定されていてパスワードが指定されていない場合、rclone は次のように動作します:
      - パスワードを要求します
      - SSHエージェントに接続しません
      

   --sftp-chunk-size
      アップロードとダウンロードのチャンクサイズ。
      
      これはSFTPプロトコルパケットのペイロードの最大サイズを制御します。
      RFCではこれは32768バイト（32k）に制限されています（デフォルト値）。しかし、
      多くのサーバーはより大きなサイズをサポートしています。通常、最大の
      パケットサイズは256kを制限するようになっており、32kより大きな値を設定すると
      高レイテンシリンクでの転送速度が大幅に向上します。これには、OpenSSHを含むものがあります。
      たとえば、256kの値を使用して問題ありません。256kのパケットサイズの範囲内にありながら、
      余分なオーバーヘッド用の十分なスペースが残ります。
      
      32k以上の値を使用するには、事前に十分なテストを行い、常に同じサーバーに接続するか、
      十分に広範なテストの後で使用してください。32kよりも大きな値を設定すると、
      「failed to send packet payload: EOF」や多数の「connection lost」、
      「corrupted on transfer」といったエラーが発生する場合は、値を下げてみてください。
      [rclone serve sftp](/commands/rclone_serve_sftp)によって動作するサーバーは
      標準の32k最大ペイロードでパケットを送信するため、ダウンロード時には
      異なるchunk_sizeを設定してはいけませんが、アップロード時には
      OpenSSHの例に書かれているようにchunk_sizeを設定することができます。
      

   --sftp-ciphers
      優先順に並べたセッション暗号化に使用する暗号のスペース区切りのリスト。
      
      少なくとも1つはサーバーの設定と一致する必要があります。これは、たとえばssh -Q cipherを使用して
      チェックできます。
      
      use_insecure_cipher が true の場合は設定しないでください。
      
      例:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com
      

   --sftp-concurrency
      1つのファイルに対する最大のアウトスタンディングリクエスト数。
      
      これは1つのファイルに対する最大のアウトスタンディングリクエスト数を制御します。
      これを増やすと、ハイレイテンシリンクでのスループットが向上しますが、
      より多くのメモリを使用します。
      

   --sftp-disable-concurrent-reads
      シーケンシャルな読み取りを使用しない場合は設定します。
      
      通常、並行した読み取りは使用して安全であり、使用しないとパフォーマンスが低下します。
      したがって、このオプションはデフォルトでは無効になっています。
      
      いくつかのサーバーでは、ファイルのダウンロード回数を制限しています。
      同時読み取りを使用すると、この制限が発生する場合があります。したがって、
      次のようなサーバーを持っている場合は:
      
          Failed to copy: file does not exist
      
      このフラグを有効にする必要があるかもしれません。
      
      並行した読み取りが無効にされている場合、use_fstat オプションは無視されます。
      

   --sftp-disable-concurrent-writes
      シーケンシャルな書き込みを使用しない場合は設定します。
      
      通常、rcloneは同時書き込みを使用してファイルをアップロードします。これにより
      パフォーマンスが大幅に向上し、特に遠隔のサーバーに対して有効です。
      
      必要に応じて、このオプションを使用して同時書き込みを無効にすることができます。
      

   --sftp-disable-hashcheck
      リモートファイルのハッシュ値の有効性を確認するためにSSHコマンドの実行を無効にします。
      
      ハッシュの使用を無効にするには空白またはtrueを設定し、ハッシュを有効にするにはtrueを設定します。

   --sftp-host
      接続するSSHホスト。
      
      例: "example.com".

   --sftp-idle-timeout
      アイドル接続を閉じる前の最大時間。
      
      指定した時間以内に接続プールに戻った接続がない場合、rcloneは接続プールを空にします。
      
      接続を無期限に保持するには0に設定します。
      

   --sftp-key-exchange
      優先順に並べたキー交換アルゴリズムのスペース区切りのリスト。
      
      少なくとも1つはサーバーの設定と一致する必要があります。これは、たとえばssh -Q kexを使用して
      チェックできます。
      
      use_insecure_cipher が true の場合は設定しないでください。
      
      例:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256
      

   --sftp-key-file
      PEMエンコードされた秘密鍵ファイルのパス。
      
      空白の場合、または key-use-agent を設定すると、ssh-agent を使用します。
      
      先頭の `~` はファイル名で展開され、環境変数（`${RCLONE_CONFIG_DIR}` など）も展開されます。

   --sftp-key-file-pass
      PEMエンコードされた秘密鍵ファイルを復号するためのパスフレーズ。
      
      暗号化されたキーは、古いOpenSSH形式のみサポートされています。新しいOpenSSH形式の暗号化キーは使用できません。

   --sftp-key-pem
      生のPEMエンコードされた秘密鍵。
      
      指定した場合、key_fileパラメーターが上書きされます。

   --sftp-key-use-agent
      ssh-agent の使用を強制します。
      
      key-file も設定されている場合、指定されたキーファイルの ".pub" ファイルが読み取られ、関連付けられたキーのみが
      ssh-agent から要求されます。これにより、「Too many authentication failures for *username*」というエラーを
      回避することができます。このエラーは、ssh-agent に多数のキーが含まれている場合に発生します。

   --sftp-known-hosts-file
      既知のホストファイルへのオプションのパス。
      
      サーバーホストキーの検証を有効にするには、この値を設定します。
      
      先頭の `~` はファイル名で展開され、環境変数（`${RCLONE_CONFIG_DIR}` など）も展開されます。

      例:
         | ~/.ssh/known_hosts | OpenSSH の known_hosts ファイルを使用します。

   --sftp-macs
      優先順に並べたMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。
      
      少なくとも1つはサーバーの設定と一致する必要があります。これは、たとえばssh -Q macを使用して
      チェックできます。
      
      例:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      

   --sftp-md5sum-command
      md5ハッシュを読み取るために使用するコマンド。
      
      自動検出するために空白のままにします。

   --sftp-pass
      SSHパスワードを指定します。ssh-agentを使用するには空白のままにします。

   --sftp-path-override
      SSHシェルコマンドで使用するパスをオーバーライドします。
      
      これにより、SFTPとSSHのパスが異なる場合のチェックサムの計算が可能になります。
      これは、Synology NASボックスなどに影響します。
      
      たとえば、共有フォルダがディレクトリボリュームを表すディレクトリにある場合:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory
      
      たとえば、ホームディレクトリが "home" という名前の共有フォルダにある場合:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --sftp-port
      SSHポート番号。

   --sftp-pubkey-file
      公開鍵ファイルへのオプションのパス。
      
      認証に使用する署名済み証明書がある場合に設定します。
      
      先頭の `~` はファイル名で展開され、環境変数（`${RCLONE_CONFIG_DIR}` など）も展開されます。

   --sftp-server-command
      リモートホストでsftpサーバーを実行するためのパスまたはコマンドを指定します。
      
      デフォルトのパスの場合、subsystemオプションは無視されます。

   --sftp-set-env
      sftpおよびコマンドに渡す環境変数
      
      次の形式で環境変数を設定します:
      
          VAR=value
      
      sftpクライアントおよび実行されるコマンド（md5sumなど）に渡されます。
      
      複数の変数をスペース区切りで指定する場合は、次のようにします:
      
          VAR1=value VAR2=value
      
      変数にスペースを含める場合は、次のようにクォートで括ります:
      
          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere
      
      

   --sftp-set-modtime
      リモートで変更された時刻を設定する場合は設定します。

   --sftp-sha1sum-command
      sha1ハッシュを読み取るために使用するコマンド。
      
      自動検出するために空白のままにします。

   --sftp-shell-type
      リモートサーバーでのSSHシェルの種類（ある場合）。
      
      自動検出するために空白のままにします。

      例:
         | none       | シェルアクセスなし
         | unix       | Unixシェル
         | powershell | PowerShell
         | cmd        | Windowsコマンドプロンプト

   --sftp-skip-links
      シンボリックリンクや他の一般的ではないファイルをスキップする場合は設定します。

   --sftp-subsystem
      リモートホスト上のSSH2サブシステムを指定します。

   --sftp-use-fstat
      statの代わりにfstatを使用する場合は設定します。
      
      一部のサーバーでは、オープンファイルの数を制限しており、ファイルを
      開いた後にStatを呼び出すとサーバーからエラーが発生します。このフラグを設定すると、
      サーバーには既にオープンされたファイルハンドルが存在する場合に呼び出される
      Fstatが呼び出されます。
      
      IBM Sterling SFTPサーバーに効果があることがわかっており、
      "extractability" レベルが 1 の場合、すなわち任意の時点で
      1つのファイルしか開くことができない場合に役立ちます。
      

   --sftp-use-insecure-cipher
      安全でない暗号およびキー交換方法を使用する場合は設定します。
      
      これにより、次のような安全でない暗号およびキー交換方法が使用されます:
      
      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1
      
      これらのアルゴリズムは安全ではなく、攻撃者によって平文データが回復される可能性があります。
      
      ciphersまたはkey_exchangeの高度なオプションを使用する場合、これはfalseにする必要があります。
      

      例:
         | false | デフォルトのCipherリストを使用します。
         | true  | aes128-cbc 暗号と diffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1 キー交換を有効にします。

   --sftp-user
      SSHユーザー名。


OPTIONS:
   --help, -h  ヘルプの表示

   データ準備オプション

   --delete-after-export    重要] CARファイルへのエクスポート後にデータセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過した場合に自動的にソースディレクトリを再スキャンします（デフォルト: 無効）
   --scanning-state value   初期のスキャン状態を設定します（デフォルト: ready）

   SFTP用オプション

   --sftp-ask-password value               SFTPパスワードの入力を許可します。 (デフォルト: "false") [$SFTP_ASK_PASSWORD]
   --sftp-chunk-size value                 アップロードおよびダウンロードのチャンクサイズ。 (デフォルト: "32Ki") [$SFTP_CHUNK_SIZE]
   --sftp-ciphers value                    優先順に並べたセッション暗号化に使用する暗号のスペース区切りのリスト。 [$SFTP_CIPHERS]
   --sftp-concurrency value                1つのファイルに対する最大のアウトスタンディングリクエスト数 (デフォルト: "64") [$SFTP_CONCURRENCY]
   --sftp-disable-concurrent-reads value   使用しない場合はシーケンシャルリードを無効にします。 (デフォルト: "false") [$SFTP_DISABLE_CONCURRENT_READS]
   --sftp-disable-concurrent-writes value  使用しない場合はシーケンシャルライトを無効にします。 (デフォルト: "false") [$SFTP_DISABLE_CONCURRENT_WRITES]
   --sftp-disable-hashcheck value          リモートファイルのハッシュ確認用のSSHコマンドの実行を無効にします。 (デフォルト: "false") [$SFTP_DISABLE_HASHCHECK]
   --sftp-host value                       接続するSSHホスト。 [$SFTP_HOST]
   --sftp-idle-timeout value               アイドル接続を閉じる前の最大時間。 (デフォルト: "1m0s") [$SFTP_IDLE_TIMEOUT]
   --sftp-key-exchange value               優先順に並べたキー交換アルゴリズムのスペース区切りのリスト。 [$SFTP_KEY_EXCHANGE]
   --sftp-key-file value                   PEMエンコードされた秘密鍵ファイルのパス。 [$SFTP_KEY_FILE]
   --sftp-key-file-pass value              PEMエンコードされた秘密鍵ファイルを復号するためのパスフレーズ。 [$SFTP_KEY_FILE_PASS]
   --sftp-key-pem value                    生のPEMエンコードされた秘密鍵。 [$SFTP_KEY_PEM]
   --sftp-key-use-agent value              SSHエージェントの使用を強制します。 (デフォルト: "false") [$SFTP_KEY_USE_AGENT]
   --sftp-known-hosts-file value           既知のホストファイルへのオプションのパス。 [$SFTP_KNOWN_HOSTS_FILE]
   --sftp-macs value                       優先順に並べたMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。 [$SFTP_MACS]
   --sftp-md5sum-command value             md5ハッシュを読み取るために使用するコマンド。 [$SFTP_MD5SUM_COMMAND]
   --sftp-pass value                       SSHパスワード。ssh-agentを使用する場合は空白のままにします。 [$SFTP_PASS]
   --sftp-path-override value              SSHシェルコマンドで使用するパスをオーバーライドします。 [$SFTP_PATH_OVERRIDE]
   --sftp-port value                       SSHポート番号。 (デフォルト: "22") [$SFTP_PORT]
   --sftp-pubkey-file value                公開鍵ファイルへのオプションのパス。 [$SFTP_PUBKEY_FILE]
   --sftp-server-command value             リモートホストでsftpサーバーを実行するためのパスまたはコマンドを指定します。 [$SFTP_SERVER_COMMAND]
   --sftp-set-env value                    sftpおよびコマンドに渡す環境変数 [$SFTP_SET_ENV]
   --sftp-set-modtime value                リモートで変更された時刻を設定する場合は設定します。 (デフォルト: "true") [$SFTP_SET_MODTIME]
   --sftp-sha1sum-command value            sha1ハッシュを読み取るために使用するコマンド。 [$SFTP_SHA1SUM_COMMAND]
   --sftp-shell-type value                 リモートサーバーでのSSHシェルの種類。 [$SFTP_SHELL_TYPE]
   --sftp-skip-links value                 シンボリックリンクや他の一般的ではないファイルをスキップする場合は設定します。 (デフォルト: "false") [$SFTP_SKIP_LINKS]
   --sftp-subsystem value                  リモートホスト上のSSH2サブシステムを指定します。 (デフォルト: "sftp") [$SFTP_SUBSYSTEM]
   --sftp-use-fstat value                  statの代わりにfstatを使用する場合は設定します。 (デフォルト: "false") [$SFTP_USE_FSTAT]
   --sftp-use-insecure-cipher value        安全でない暗号およびキー交換方法を使用する場合は設定します。 (デフォルト: "false") [$SFTP_USE_INSECURE_CIPHER]
   --sftp-user value                       SSHユーザー名。 (デフォルト: "$USER") [$SFTP_USER]

```