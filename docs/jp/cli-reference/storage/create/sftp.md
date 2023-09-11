# SSH/SFTP

```
NAME:
   singularityストレージの作成 sftp - SSH/SFTP

USAGE:
   singularityストレージの作成 sftp [command options] [arguments...]

DESCRIPTION:
   --host
      接続先のSSHホスト。

      例: "example.com"

   --user
      SSHユーザー名。

   --port
      SSHポート番号。

   --pass
      SSHパスワード。ssh-agentを使用する場合は空白。

   --key-pem
      PEM形式の秘密鍵データ。

      指定した場合、key_fileパラメータを上書きします。

   --key-file
      PEM形式の秘密鍵ファイルのパス。

      空白またはkey-use-agentを設定すると、ssh-agentを使用します。

      先頭の`~`や`${RCLONE_CONFIG_DIR}`のような環境変数はファイル名で展開されます。

   --key-file-pass
      PEM形式の暗号化された秘密鍵ファイルのパスワード。

      このオプションは、古いOpenSSH形式の暗号化された鍵ファイルのみサポートされています。
      新しいOpenSSH形式の暗号化された鍵は使用できません。

   --pubkey-file
      公開鍵ファイルのオプションパス。

      認証に使用する署名済み証明書がある場合に設定します。

      先頭の`~`や`${RCLONE_CONFIG_DIR}`のような環境変数はファイル名で展開されます。

   --known-hosts-file
      接続先のknown_hostsファイルのオプションパス。

      この値を設定すると、サーバーホストキーの検証を有効にします。

      先頭の`~`や`${RCLONE_CONFIG_DIR}`のような環境変数はファイル名で展開されます。

      Examples:
         | ~/.ssh/known_hosts | OpenSSHのknown_hostsファイルを使用します。

   --key-use-agent
      ssh-agentの使用を強制します。

      key-fileも設定されている場合、指定したkey-fileの".pub"ファイルが読み込まれ、関連する鍵のみが
      ssh-agentから要求されます。これにより、ssh-agentに多くの鍵が含まれている場合に
      `Too many authentication failures for *username*`エラーを回避できます。

   --use-insecure-cipher
      セキュリティの脆弱性のある暗号化方式と鍵交換方式を有効にします。

      これにより、以下のセキュリティの脆弱性のある暗号化方式と鍵交換方式が使用できるようになります:

      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1

      これらのアルゴリズムは安全ではなく、攻撃者によって平文データが復号される可能性があります。

      ciphersまたはkey_exchangeの高度なオプションを使用する場合、この値はfalseにする必要があります。

      Examples:
         | false | デフォルトのCipherリストを使用します。
         | true  | aes128-cbc暗号化方式およびdiffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1鍵交換を有効にします。

   --disable-hashcheck
      リモートファイルのハッシュを確認するためのSSHコマンドの実行を無効にします。

      ハッシュの確認を有効にする場合は空白またはfalseを設定し、ハッシュの無効化にはtrueを設定します。

   --ask-password
      必要に応じて、SFTPパスワードの入力を許可します。

      このオプションが設定されていてパスワードが指定されていない場合、rcloneは次のようにします:
      - パスワードの入力をリクエストします
      - sshエージェントにアクセスしません

   --path-override
      SSHシェルコマンドで使用されるパスを上書きします。

      これにより、SFTPとSSHのパスが異なる場合にチェックサム計算が可能になります。
      この問題は、他のものの中でもSynology NASボックスに影響します。

      例えば、共有フォルダがボリュームを示すディレクトリに格納されている場合:
      
          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory

      例えば、ホームディレクトリが「home」という共有フォルダに格納されている場合:
      
          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      リモートファイルの変更日時を設定します。

   --shell-type
      リモートサーバーのSSHシェルの種類。

      自動検出する場合は空白にします。

      Examples:
         | none       | シェルアクセス無し
         | unix       | Unixシェル
         | powershell | PowerShell
         | cmd        | Windowsコマンドプロンプト

   --md5sum-command
      MD5ハッシュの読み取りに使用するコマンド。

      自動検出する場合は空白にします。

   --sha1sum-command
      SHA1ハッシュの読み取りに使用するコマンド。

      自動検出する場合は空白にします。

   --skip-links
      シンボリックリンクやその他の定規ファイルをスキップするように設定します。

   --subsystem
      リモートホストで使用するSSH2サブシステムを指定します。

   --server-command
      リモートホストでSFTPサーバーを実行するためのパスまたはコマンドを指定します。

      server_commandが定義されている場合、subsystemオプションは無視されます。

   --use-fstat
      statではなくfstatを使用します。

      一部のサーバーでは、オープンできるファイルの数に制限があり、ファイルを
      開いた後にStatを呼び出すとサーバーからエラーが返されます。
      このフラグを設定すると、すでにオープンされたファイルハンドルで呼び出される
      Fstatを呼び出します。

      これにより、"extractability"レベルが1に設定されたIBM Sterling SFTPサーバーや、
      1回につき1つのファイルしかオープンできないという制約があるサーバーで効果があります。

   --disable-concurrent-reads
      同時読み取りを使用しない場合は設定します。

      通常、同時読み取りは安全に使用できるため、デフォルトでは無効になっています。
      一部のサーバーでは、ファイルのダウンロード回数に制限があります。
      同時読み取りを使用すると、この制限がトリガーされることがあります。
      したがって、次のエラーが表示される場合は、このフラグを有効にする必要があります。

          Failed to copy: file does not exist

      同時読み取りが無効になっている場合、use_fstatオプションは無効になります。

   --disable-concurrent-writes
      同時書き込みを使用しない場合は設定します。

      通常、rcloneはファイルのアップロードに同時書き込みを使用します。
      これにより、特に遠隔サーバーの場合、パフォーマンスが大幅に向上します。

      必要がある場合にのみ、このオプションを無効にします。

   --idle-timeout
      アイドル接続を閉じる前の最大時間。

      指定した時間内に接続がコネクションプールに返されない場合、rcloneはコネクションプールをクリアします。

      0に設定すると、接続を無期限に保持します。

   --chunk-size
      アップロードおよびダウンロードのチャンクサイズ。

      これはSFTPプロトコルパケットのペイロードの最大サイズを制御します。
      RFCではこれを32768バイト（32k）までに制限していますが、多くのサーバーはより大きなサイズをサポートしています。
      通常、最大パケットサイズが256kを超えないように制限されており、その場合は転送速度が劇的に向上します。
      これにはOpenSSHも含まれます。たとえば、256kの値を使用すると、256kの総パケットサイズ内で
      余裕をもってオーバーヘッドが残ります。

      32kより大きな値を使用する場合は、十分なテストを行ってから使用するようにしてください。
      同じサーバーに常に接続するか、十分に広範なテストの後にのみ使用してください。
      大きなファイルのコピー時に「failed to send packet payload: EOF」や多数の「connection lost」や「corrupted on transfer」といったエラーが発生する場合は、値を下げて試してください。
      [rclone serve sftp](/commands/rclone_serve_sftp)によって実行されるサーバーは32kの標準最大ペイロードサイズでパケットを送信するため、ダウンロードファイルの場合は異なるchunk_sizeを設定しないでくださいが、256kのパケットサイズまでパケットを受け入れますので、アップロードのchunk_sizeはOpenSSHの例のように設定できます。

   --concurrency
      1つのファイルに対する最大リクエスト数。

      これは、1つのファイルに対する最大リクエスト数を制御します。
      この値を増やすと、高遅延リンクでのスループットが向上しますが、より多くのメモリを使用します。

   --set-env
      sftpおよびコマンドに渡すための環境変数を設定します。

      以下の形式で環境変数を設定します:

          VAR=value

      sftpクライアントおよび実行されるコマンド（例: md5sum）に渡すために設定します。

      複数の変数をスペース区切りで指定することもできます:

          VAR1=value VAR2=value

      値にスペースが含まれる場合は、引用符で囲んで指定します:

          "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere

   --ciphers
      優先順位で指定されたセッション暗号化に使用する暗号方式のスペース区切りのリスト。

      少なくとも1つがサーバーの設定と一致する必要があります。例えば、ssh -Q cipherを使用して確認できます。

      use_insecure_cipherがtrueの場合、この値を設定してはいけません。

      例:
      
          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com

   --key-exchange
      優先順位で指定された鍵交換アルゴリズムのスペース区切りのリスト。

      少なくとも1つがサーバーの設定と一致する必要があります。例えば、ssh -Q kexを使用して確認できます。

      use_insecure_cipherがtrueの場合、この値を設定してはいけません。

      例:
      
          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256

   --macs
      優先順位で指定されたMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。

      少なくとも1つがサーバーの設定と一致する必要があります。例えば、ssh -Q macを使用して確認できます。

      例:
      
          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com
      
OPTIONS:
   --disable-hashcheck    SSHコマンドを実行して、リモートファイルのハッシュの可用性を確認しない。(デフォルト: false) [$DISABLE_HASHCHECK]
   --help, -h             ヘルプを表示
   --host value           接続先のSSHホスト。[$HOST]
   --key-file value       PEM形式の秘密鍵ファイルのパス。[$KEY_FILE]
   --key-file-pass value  PEM形式の秘密鍵ファイルのパスワード。[$KEY_FILE_PASS]
   --key-pem value        PEM形式の秘密鍵データ。[$KEY_PEM]
   --key-use-agent        ssh-agentの使用を強制する場合はtrue。(デフォルト: false) [$KEY_USE_AGENT]
   --pass value           SSHパスワード。ssh-agentを使用する場合は空白。[$PASS]
   --port value           SSHポート番号。(デフォルト: 22) [$PORT]
   --pubkey-file value    公開鍵ファイルのオプションパス。[$PUBKEY_FILE]
   --use-insecure-cipher  セキュリティの脆弱性のある暗号化方式と鍵交換方式を使用する場合はtrue。(デフォルト: false) [$USE_INSECURE_CIPHER]
   --user value           SSHユーザー名。(デフォルト: "$USER") [$USER]

   Advanced

   --ask-password               必要に応じてSFTPパスワードの入力を許可する場合はtrue。(デフォルト: false) [$ASK_PASSWORD]
   --chunk-size value           アップロードおよびダウンロードのチャンクサイズ。(デフォルト: "32Ki") [$CHUNK_SIZE]
   --ciphers value              優先順位で指定されたセッション暗号化に使用する暗号方式のスペース区切りのリスト。[$CIPHERS]
   --concurrency value          1つのファイルに対する最大リクエスト数。(デフォルト: 64) [$CONCURRENCY]
   --disable-concurrent-reads   同時読み取りを使用しない場合はtrue。(デフォルト: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  同時書き込みを使用しない場合はtrue。(デフォルト: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         アイドル接続を閉じる前の最大時間。(デフォルト: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         優先順位で指定された鍵交換アルゴリズムのスペース区切りのリスト。[$KEY_EXCHANGE]
   --known-hosts-file value     接続先のknown_hostsファイルのオプションパス。[$KNOWN_HOSTS_FILE]
   --macs value                 優先順位で指定されたMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。[$MACS]
   --md5sum-command value       MD5ハッシュの読み取りに使用するコマンド。[$MD5SUM_COMMAND]
   --path-override value        SSHシェルコマンドで使用されるパスを上書きします。[$PATH_OVERRIDE]
   --server-command value       リモートホストでSFTPサーバーを実行するためのパスまたはコマンド。[$SERVER_COMMAND]
   --set-env value              sftpおよびコマンドに渡すための環境変数。[$SET_ENV]
   --set-modtime                リモートファイルの変更日時を設定する場合はtrue。(デフォルト: true) [$SET_MODTIME]
   --sha1sum-command value      SHA1ハッシュの読み取りに使用するコマンド。[$SHA1SUM_COMMAND]
   --shell-type value           リモートサーバーのSSHシェルの種類。[$SHELL_TYPE]
   --skip-links                 シンボリックリンクやその他の定規ファイルをスキップする場合はtrue。(デフォルト: false) [$SKIP_LINKS]
   --subsystem value            リモートホストで使用するSSH2サブシステム。(デフォルト: "sftp") [$SUBSYSTEM]
   --use-fstat                  statではなくfstatを使用する場合はtrue。(デフォルト: false) [$USE_FSTAT]

   General

   --name value  ストレージの名前。(デフォルト: 自動生成)
   --path value  ストレージのパス
```