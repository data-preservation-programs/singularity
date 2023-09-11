# SSH/SFTP

{% code fullWidth="true" %}
```
名前:
   singularity storage update sftp - SSH/SFTP

使用法:
   singularity storage update sftp [コマンドオプション] <名前|ID>

説明:
   --host
      接続するSSHホスト。

      例: "example.com"。

   --user
      SSHユーザー名。

   --port
      SSHポート番号。

   --pass
      SSHパスワード。ssh-agentを使用する場合は空白のままにします。

   --key-pem
      RawのPEMエンコードされた秘密鍵。

      指定すると、key_fileパラメーターは上書きされます。

   --key-file
      PEMエンコードされた秘密鍵ファイルへのパス。

      空白にするか、key-use-agentをssh-agentの使用に設定します。

      先頭の `~` はファイル名内で展開されますし、`${RCLONE_CONFIG_DIR}` のような環境変数も展開されます。

   --key-file-pass
      PEMエンコードされた秘密鍵ファイルを復号化するためのパスフレーズ。

      古いOpenSSH形式のPEM暗号化キーファイルのみサポートされています。新しいOpenSSH形式の暗号化キーは使用できません。

   --pubkey-file
      使用する署名付き証明書のオプションのパス。

      ファイル名内の先頭の `~` は展開されますし、`${RCLONE_CONFIG_DIR}` のような環境変数も展開されます。

   --known-hosts-file
      オプションのknown_hostsファイルへのパス。

      サーバーホストキーの検証を有効にする場合は、この値を設定します。

      ファイル名内の先頭の `~` は展開されますし、`${RCLONE_CONFIG_DIR}` のような環境変数も展開されます。

      例:
         | ~/.ssh/known_hosts | OpenSSHのknown_hostsファイルを使用する。

   --key-use-agent
      ssh-agentの使用を強制します。

      key-fileも設定されている場合、指定されたkey-fileの ".pub" ファイルが読み込まれ、関連するキーのみ
      ssh-agentから要求されます。これにより、ssh-agentに多くのキーが含まれている場合に発生する
      `Too many authentication failures for *username*` エラーを回避することができます。

   --use-insecure-cipher
      セキュリティの脆弱な暗号とキーエクスチェンジ手法の使用を有効にします。

      次のセキュリティの脆弱な暗号とキーエクスチェンジ手法を使用できるようにします：

      - aes128-cbc
      - aes192-cbc
      - aes256-cbc
      - 3des-cbc
      - diffie-hellman-group-exchange-sha256
      - diffie-hellman-group-exchange-sha1

      これらのアルゴリズムはセキュリティの問題があり、攻撃者によって平文データを回復される可能性があります。

      これは、ciphersまたはkey_exchangeの高度なオプションを使用する場合はfalseに設定する必要があります。

      例:
         | false | デフォルトの暗号リストを使用します。
         | true  | aes128-cbc 暗号と diffie-hellman-group-exchange-sha256、diffie-hellman-group-exchange-sha1 キーエクスチェンジを使用できるようにします。

   --disable-hashcheck
      リモートファイルハッシングが有効かどうかを判断するためのSSHコマンドの実行を無効にします。

      ハッシュの有効化には空白またはfalseを設定し、ハッシングの無効化にはtrueを設定します。

   --ask-password
      必要な場合にSFTPパスワードを要求できるようにします。

      このオプションが設定され、パスワードが提供されていない場合、rcloneは以下の操作を行います：
      - パスワードの入力を要求する
      - sshエージェントには連絡しない

   --path-override
      SSHシェルコマンドによって使用されるパスを上書きします。

      これにより、SFTPとSSHのパスが異なる場合にチェックサムの計算が可能になります。これは他のものにも影響を与えます、
      例えばSynology NASボックス。

      例えば、共有フォルダがボリュームを表すディレクトリ内に見つかる場合：

          rclone sync /home/local/directory remote:/directory --sftp-path-override /volume2/directory

      例えば、ホームディレクトリが "home" という共有フォルダにある場合：

          rclone sync /home/local/directory remote:/home/directory --sftp-path-override /volume1/homes/USER/directory

   --set-modtime
      リモートの更新時間が設定されている場合に、更新時間を設定します。

   --shell-type
      リモートサーバー上のSSHシェルのタイプ。

      自動検出する場合は空白にします。

      例:
         | none       | シェルアクセスなし
         | unix       | Unixシェル
         | powershell | PowerShell
         | cmd        | Windowsコマンドプロンプト

   --md5sum-command
      MD5ハッシュを読み取るために使用するコマンド。

      自動検出する場合は空白にします。

   --sha1sum-command
      SHA1ハッシュを読み取るために使用するコマンド。

      自動検出する場合は空白にします。

   --skip-links
      シンボリックリンクとその他の通常のファイルをスキップするように設定します。

   --subsystem
      リモートホスト上のSSH2サブシステムを指定します。

   --server-command
      リモートホストでSFTPサーバーを実行するためのパスまたはコマンドを指定します。

      server_commandが定義されている場合、subsystemオプションは無視されます。

   --use-fstat
      fstatをstatの代わりに使用します。

      一部のサーバーは開いているファイルの数を制限していて、ファイルを開いた後にStatを呼び出すと、サーバーからエラーが発生します。
      このフラグを設定すると、既に開いているファイルハンドルに対して呼び出されるFstatが呼び出されます。

      これによって、"extractability" レベルが1に設定されているIBM Sterling SFTPサーバーで問題が発生しなくなります。
      これはつまり、与えられた時間で1ファイルしか開けないことを意味します。

   --disable-concurrent-reads
      ファイルの同時読み取りを使用しない場合は設定します。

      通常、ファイルの同時読み取りは安全であり、使用しない場合はパフォーマンスが低下します。そのため、
      このオプションはデフォルトで無効になっています。

      一部のサーバーでは、同時にファイルをダウンロードできる回数に制限があります。同時読み取りはこの制限を引き起こす可能性があるため、
      次のエラーメッセージが表示される場合：

          Failed to copy: file does not exist

      このフラグを有効にする必要があります。

      同時読み取りが無効にされている場合、use_fstatオプションは無視されます。
      

   --disable-concurrent-writes
      ファイルの同時書き込みを使用しない場合は設定します。

      通常、rcloneはファイルの同時書き込みを使用してファイルをアップロードします。これにより、特に遠隔サーバーの場合にパフォーマンスが大幅に向上します。

      このオプションは必要な場合に同時書き込みを無効にします。

   --idle-timeout
      アイドル接続を閉じる前の最大時間。

      コネクションプールに接続が返されなかった場合、指定された時間内にrcloneはコネクションプールを空にします。

      ゼロに設定すると、接続が無期限に保持されます。

   --chunk-size
      アップロードとダウンロードのチャンクサイズ。

      これはSFTPプロトコルパケットのペイロードの最大サイズを制御します。
      RFCではこれを32768バイト（32k）に制限しています（デフォルト値）。ただし、多くのサーバーは
      より大きなサイズをサポートしており、通常は最大256kのパケットサイズに制限されています。
      この値を大きくすると、遅延のあるリンクでの転送速度が劇的に向上します。これにはOpenSSHも含まれます。
      たとえば、OpenSSHの場合、値を255kに設定すると、オーバーヘッドの十分な余地が残されますが、
      256kの総パケットサイズ内に収まります。

      32kよりも大きな値を使用する場合は、十分なテストを行ってから使用してください。
      ファイルが大きい場合に "failed to send packet payload: EOF" や多くの "connection lost"、
      または "corrupted on transfer" のエラーが発生した場合は、値を下げてみてください。
      [rclone serve sftp](/commands/rclone_serve_sftp) で実行されるサーバーは、
      標準の32kの最大ペイロードを持つパケットを送信しますので、ダウンロード時に異なるchunk_sizeを設定しないでくださいが、
      256kの総サイズまでのパケットを受け入れるため、アップロード時にはOpenSSHの例と同じchunk_sizeを設定できます。

   --concurrency
      1つのファイルに対する最大未処理リクエスト数

      これにより、1つのファイルに対する最大未処理リクエスト数を制御できます。数を増やすと、
      高レイテンシリンクでのスループットが増加しますが、メモリの使用量も増えます。

   --set-env
      sftpおよびコマンドに渡す環境変数

      以下の形式で環境変数を設定します：

          VAR=value

      sftpクライアントおよび実行するコマンド（例：md5sum）に渡すためです。

      複数の変数をスペース区切りで指定します。例えば：

          VAR1=value VAR2=value

      スペースを含む変数は引用符で囲んで指定します。例えば：

         "VAR3=value with space" "VAR4=value with space" VAR5=nospacehere

   --ciphers
      優先順位に基づいてセッション暗号化に使用するスペース区切りの暗号リスト。

      少なくとも1つはサーバーの設定と一致する必要があります。ssh -Q cipherなどを使用して確認できます。

      これはuse_insecure_cipherがtrueの場合は設定しないでください。

      例：

          aes128-ctr aes192-ctr aes256-ctr aes128-gcm@openssh.com aes256-gcm@openssh.com

   --key-exchange
      優先順位に基づいてキーエクスチェンジアルゴリズムのスペース区切りのリスト。

      少なくとも1つはサーバーの設定と一致する必要があります。ssh -Q kexなどを使用して確認できます。

      use_insecure_cipherがtrueの場合は設定しないでください。

      例：

          sntrup761x25519-sha512@openssh.com curve25519-sha256 curve25519-sha256@libssh.org ecdh-sha2-nistp256

   --macs
      優先順位に基づいてMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。

      少なくとも1つはサーバーの設定と一致する必要があります。ssh -Q macなどを使用して確認できます。

      例：

          umac-64-etm@openssh.com umac-128-etm@openssh.com hmac-sha2-256-etm@openssh.com

オプション:
   --disable-hashcheck    リモートファイルハッシングが有効かどうかを判断するためのSSHコマンドの実行を無効にします。 (デフォルト: false) [$DISABLE_HASHCHECK]
   --help, -h             ヘルプを表示
   --host value           接続するSSHホスト。 [$HOST]
   --key-file value       PEMエンコードされた秘密鍵ファイルへのパス。 [$KEY_FILE]
   --key-file-pass value  PEMエンコードされた秘密鍵ファイルを復号化するためのパスフレーズ。 [$KEY_FILE_PASS]
   --key-pem value        RawのPEMエンコードされた秘密鍵。 [$KEY_PEM]
   --key-use-agent        ssh-agentの使用を強制します。 (デフォルト: false) [$KEY_USE_AGENT]
   --pass value           SSHパスワード。ssh-agentを使用する場合は空白のままにします。 [$PASS]
   --port value           SSHポート番号。 (デフォルト: 22) [$PORT]
   --pubkey-file value    使用する署名付き証明書のオプションのパス。 [$PUBKEY_FILE]
   --use-insecure-cipher  セキュリティの脆弱な暗号とキーエクスチェンジ手法の使用を有効にします。 (デフォルト: false) [$USE_INSECURE_CIPHER]
   --user value           SSHユーザー名。 (デフォルト: "$USER") [$USER]

   高度なオプション

   --ask-password               必要な場合にSFTPパスワードを要求できるようにします。 (デフォルト: false) [$ASK_PASSWORD]
   --chunk-size value           アップロードとダウンロードのチャンクサイズ。 (デフォルト: "32Ki") [$CHUNK_SIZE]
   --ciphers value              優先順位に基づいてセッション暗号化に使用するスペース区切りの暗号リスト。 [$CIPHERS]
   --concurrency value          1つのファイルに対する最大未処理リクエスト数 (デフォルト: 64) [$CONCURRENCY]
   --disable-concurrent-reads   ファイルの同時読み取りを使用しない場合は設定します。 (デフォルト: false) [$DISABLE_CONCURRENT_READS]
   --disable-concurrent-writes  ファイルの同時書き込みを使用しない場合は設定します。 (デフォルト: false) [$DISABLE_CONCURRENT_WRITES]
   --idle-timeout value         アイドル接続を閉じる前の最大時間。 (デフォルト: "1m0s") [$IDLE_TIMEOUT]
   --key-exchange value         優先順位に基づいてキーエクスチェンジアルゴリズムのスペース区切りのリスト。 [$KEY_EXCHANGE]
   --known-hosts-file value     オプションのknown_hostsファイルへのパス。 [$KNOWN_HOSTS_FILE]
   --macs value                 優先順位に基づいてMAC（メッセージ認証コード）アルゴリズムのスペース区切りのリスト。 [$MACS]
   --md5sum-command value       MD5ハッシュを読み取るために使用するコマンド。 [$MD5SUM_COMMAND]
   --path-override value        SSHシェルコマンドによって使用されるパスを上書きします。 [$PATH_OVERRIDE]
   --server-command value       リモートホストでSFTPサーバーを実行するためのパスまたはコマンドを指定します。 [$SERVER_COMMAND]
   --set-env value              sftpおよびコマンドに渡す環境変数 [$SET_ENV]
   --set-modtime                リモートの更新時間が設定されている場合に、更新時間を設定します。 (デフォルト: true) [$SET_MODTIME]
   --sha1sum-command value      SHA1ハッシュを読み取るために使用するコマンド。 [$SHA1SUM_COMMAND]
   --shell-type value           リモートサーバー上のSSHシェルのタイプ。 [$SHELL_TYPE]
   --skip-links                 シンボリックリンクとその他の通常のファイルをスキップするように設定します。 (デフォルト: false) [$SKIP_LINKS]
   --subsystem value            リモートホスト上のSSH2サブシステムを指定します。 (デフォルト: "sftp") [$SUBSYSTEM]
   --use-fstat                  fstatをstatの代わりに使用します。 (デフォルト: false) [$USE_FSTAT]

```
{% endcode %}