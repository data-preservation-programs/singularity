# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add ftp - FTP

USAGE:
   singularity datasource add ftp [command options] <dataset_name> <source_path>

DESCRIPTION:
   --ftp-ask-password
      必要な場合にFTPパスワードを入力することを許可します。
      
      これが設定されており、パスワードが提供されていない場合、rcloneはパスワードの入力を求めます。
      

   --ftp-close-timeout
      接続を閉じるための応答を待つ最大時間

   --ftp-concurrency
      同時に行うFTPの最大接続数。無制限の場合は0を指定します。
      
      ただし、これを設定するとデッドロックが発生する可能性が非常に高いため、注意して使用する必要があります。
      
      同期やコピーを行っている場合は、`--transfers`オプションと`--checkers`オプションの合計に1を加えた値を設定してください。
      
      `--check-first`を使用する場合は、`--checkers`と`--transfers`の最大値に1を足した値で十分です。
      
      つまり、`concurrency 3`の場合は、`--checkers 2 --transfers 2 --check-first`または`--checkers 1 --transfers 1`を使用します。
      
      

   --ftp-disable-epsv
      サーバーがサポートを表示していても、EPSVの使用を無効にします。

   --ftp-disable-mlsd
      サーバーがサポートを表示していても、MLSDの使用を無効にします。

   --ftp-disable-tls13
      TLS 1.3を無効にします（TLSのバグのあるFTPサーバーの問題の回避策）

   --ftp-disable-utf8
      サーバーがサポートを表示していても、UTF-8の使用を無効にします。

   --ftp-encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

      例:
         | Asterisk,Ctl,Dot,Slash                               | ファイル名に「*」を使用できない /* ProFTPd */
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | ファイル名に「[]」または「*」を使用できない /* PureFTPd */
         | Ctl,LeftPeriod,Slash                                 | ファイル名がドットで始まる場合に使用できない /* VsFTPd */

   --ftp-explicit-tls
      明示的なFTPS（TLSを使用したFTP）を使用します。
      
      明示的なFTP over TLSを使用すると、クライアントは平文の接続を暗号化された接続にアップグレードするためにサーバーにセキュリティを明示的に要求します。
      暗黙的なFTPSとの組み合わせは使用できません。

   --ftp-force-list-hidden
      隠しファイルやフォルダーのリストを強制的に表示するためにLIST -aを使用します。これにより、MLSDの使用が無効になります。

   --ftp-host
      接続するFTPホスト。
      
      例: "ftp.example.com"。

   --ftp-idle-timeout
      アイドル接続を閉じる前の最大時間。
      
      指定した時間内に接続が接続プールに返されなかった場合、rcloneは接続プールを空にします。
      
      無制限に接続を保持するには、0を設定します。
      

   --ftp-no-check-certificate
      サーバーのTLS証明書を検証しません。

   --ftp-pass
      FTPパスワード。

   --ftp-port
      FTPポート番号。

   --ftp-shut-timeout
      データ接続のクローズ状態を待つ最大時間。

   --ftp-tls
      暗黙的なFTPS（TLSを介したFTP）を使用します。
      
      暗黙的なFTP over TLSを使用すると、クライアントは最初からTLSを使用して接続します。これにより、TLSに対応していないサーバーとの互換性が失われます。
      通常、ポート21ではなくポート990で提供されます。明示的なFTPSとの組み合わせは使用できません。

   --ftp-tls-cache-size
      すべての制御接続およびデータ接続のTLSセッションキャッシュのサイズ。
      
      TLSキャッシュには、TLSセッションを再開し、接続間でPSKを再利用する機能があります。
      デフォルトのサイズが不十分でTLSの再開エラーが発生する場合は、値を増やしてください。
      デフォルトでは有効になっています。無効にするには0を使用してください。

   --ftp-user
      FTPユーザー名。

   --ftp-writing-mdtm
      MDTMを使用して変更時刻を設定します（VsFtpdのクワーク）


OPTIONS:
   --help, -h  ヘルプを表示

   データ準備オプション

   --delete-after-export    [危険] データセットのファイルをCARファイルにエクスポートした後に削除します。  (default: false)
   --rescan-interval value  最後の成功したスキャンからこの間隔が経過したら、ソースディレクトリを自動的に再スキャンします (default: 無効)
   --scanning-state value   初期のスキャン状態を設定します (default: ready)

   ftpのオプション

   --ftp-ask-password value          必要な場合にFTPパスワードを入力することを許可します。 (default: "false") [$FTP_ASK_PASSWORD]
   --ftp-close-timeout value         接続を閉じるための応答を待つ最大時間 (default: "1m0s") [$FTP_CLOSE_TIMEOUT]
   --ftp-concurrency value           同時に行うFTPの最大接続数。無制限の場合は0を指定します。 (default: "0") [$FTP_CONCURRENCY]
   --ftp-disable-epsv value          サーバーがサポートを表示していても、EPSVの使用を無効にします。 (default: "false") [$FTP_DISABLE_EPSV]
   --ftp-disable-mlsd value          サーバーがサポートを表示していても、MLSDの使用を無効にします。 (default: "false") [$FTP_DISABLE_MLSD]
   --ftp-disable-tls13 value         TLS 1.3を無効にします（TLSのバグのあるFTPサーバーの問題の回避策） (default: "false") [$FTP_DISABLE_TLS13]
   --ftp-disable-utf8 value          サーバーがサポートを表示していても、UTF-8の使用を無効にします。 (default: "false") [$FTP_DISABLE_UTF8]
   --ftp-encoding value              バックエンドのエンコーディング。 (default: "Slash,Del,Ctl,RightSpace,Dot") [$FTP_ENCODING]
   --ftp-explicit-tls value          明示的なFTPS（TLSを使用したFTP）を使用します。 (default: "false") [$FTP_EXPLICIT_TLS]
   --ftp-force-list-hidden value     隠しファイルやフォルダーのリストを強制的に表示するためにLIST -aを使用します。これにより、MLSDの使用が無効になります。 (default: "false") [$FTP_FORCE_LIST_HIDDEN]
   --ftp-host value                  接続するFTPホスト。 [$FTP_HOST]
   --ftp-idle-timeout value          アイドル接続を閉じる前の最大時間。 (default: "1m0s") [$FTP_IDLE_TIMEOUT]
   --ftp-no-check-certificate value  サーバーのTLS証明書を検証しません。 (default: "false") [$FTP_NO_CHECK_CERTIFICATE]
   --ftp-pass value                  FTPパスワード。 [$FTP_PASS]
   --ftp-port value                  FTPポート番号。 (default: "21") [$FTP_PORT]
   --ftp-shut-timeout value          データ接続のクローズ状態を待つ最大時間。 (default: "1m0s") [$FTP_SHUT_TIMEOUT]
   --ftp-tls value                   暗黙的なFTPS（TLSを介したFTP）を使用します。 (default: "false") [$FTP_TLS]
   --ftp-tls-cache-size value        すべての制御接続およびデータ接続のTLSセッションキャッシュのサイズ。 (default: "32") [$FTP_TLS_CACHE_SIZE]
   --ftp-user value                  FTPユーザー名。 (default: "$USER") [$FTP_USER]
   --ftp-writing-mdtm value          MDTMを使用して変更時刻を設定します（VsFtpdのクワーク） (default: "false")     
```