# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage create ftp - FTP

USAGE:
   singularity storage create ftp [command options] [arguments...]

DESCRIPTION:
   --host
      FTPに接続するためのホスト。

      例: "ftp.example.com"。

   --user
      FTPのユーザー名。

   --port
      FTPのポート番号。

   --pass
      FTPのパスワード。

   --tls
      暗黙のFTPS（TLSを使用したFTP）を使用する。

      暗黙のFTPSを使用すると、クライアントはTLSを使用して開始時から接続し、
      非TLS対応のサーバーとの互換性がなくなります。
      通常、これはポート21ではなくポート990で提供されます。
      明示的なFTPSとの組み合わせは使用できません。

   --explicit-tls
      明示的なFTPS（TLSを使用したFTP）を使用する。

      明示的なFTP over TLSを使用すると、クライアントは接続を平文から暗号化された接続にアップグレードするために、
      サーバーにセキュリティを明示的に要求します。
      暗黙のFTPSとの組み合わせは使用できません。

   --concurrency
      FTPの同時接続の最大数。制限なしの場合は0を指定してください。

      ただし、この値を設定すると、デッドロックの発生が非常に高いため、注意して使用する必要があります。

      同期やコピーを行う場合は、`--transfers`と`--checkers`の合計に1を加える必要があります。

      `--check-first`を使用する場合、`--checkers`と`--transfers`の最大値に1を加えるだけで十分です。

      例えば、`concurrency 3`の場合、`--checkers 2 --transfers 2 --check-first`または`--checkers 1 --transfers 1`を使用します。

   --no-check-certificate
      サーバーのTLS証明書を検証しない。

   --disable-epsv
       サーバーがサポートを宣伝していても、EPSVの使用を無効にします。

   --disable-mlsd
      サーバーがサポートを宣伝していても、MLSDの使用を無効にします。

   --disable-utf8
      サーバーがサポートを宣伝していても、UTF-8の使用を無効にします。

   --writing-mdtm
      修正時刻を設定するためにMDTMを使用します（VsFtpdのクワーク）。

   --force-list-hidden
      隠しファイルとフォルダのリスト表示にLIST -aを使用します。これにより、MLSDの使用が無効になります。

   --idle-timeout
      アイドル状態の接続を閉じる前の最大時間。

      タイムアウト時間内に接続がコネクションプールに返されていない場合、
      rcloneはコネクションプールを空にします。

      接続を無制限に保持するには、0を設定してください。

   --close-timeout
      クローズのレスポンスを待つ最大時間。

   --tls-cache-size
      全ての制御とデータ接続のTLSセッションキャッシュのサイズ。

      TLSキャッシュには、TLSセッションを再開し、接続間でPSKを再利用することができます。
      デフォルトのサイズが不十分な場合、TLSの再開エラーが発生します。
      デフォルトでは有効です。無効にする場合は0を使用してください。

   --disable-tls13
      TLS 1.3を無効にします（TLSが不具合のあるFTPサーバーの回避策）。

   --shut-timeout
      データ接続のクローズステータスを待つ最大時間。

   --ask-password
      必要な場合にFTPのパスワードの入力を許可します。

      このオプションが設定されており、パスワードが提供されていない場合、rcloneはパスワードを要求します。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

      例:
         | Asterisk,Ctl,Dot,Slash                           | ファイル名で'*'を使用できないProFTPd
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | ファイル名で'[]'や'*'を使用できないPureFTPd
         | Ctl,LeftPeriod,Slash                                 | ファイル名がドットで始まるのでVsFTPdで処理できません


OPTIONS:
   --explicit-tls  明示的なFTPS（TLSを使用したFTP）を使用する。 (デフォルト: false) [$EXPLICIT_TLS]
   --help, -h      ヘルプを表示
   --host value    FTPに接続するためのホスト。 [$HOST]
   --pass value    FTPのパスワード。 [$PASS]
   --port value    FTPのポート番号。 (デフォルト: 21) [$PORT]
   --tls           暗黙のFTPS（TLSを使用したFTP）を使用する。 (デフォルト: false) [$TLS]
   --user value    FTPのユーザー名。 (デフォルト: "$USER") [$USER]

   Advanced

   --ask-password          必要な場合にFTPのパスワードの入力を許可します。 (デフォルト: false) [$ASK_PASSWORD]
   --close-timeout value   クローズのレスポンスを待つ最大時間。 (デフォルト: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     FTPの同時接続の最大数。制限なしの場合は0を指定してください。 (デフォルト: 0) [$CONCURRENCY]
   --disable-epsv          サーバーがサポートを宣伝していても、EPSVの使用を無効にします。 (デフォルト: false) [$DISABLE_EPSV]
   --disable-mlsd          サーバーがサポートを宣伝していても、MLSDの使用を無効にします。 (デフォルト: false) [$DISABLE_MLSD]
   --disable-tls13         TLS 1.3を無効にします（TLSが不具合のあるFTPサーバーの回避策） (デフォルト: false) [$DISABLE_TLS13]
   --disable-utf8          サーバーがサポートを宣伝していても、UTF-8の使用を無効にします。 (デフォルト: false) [$DISABLE_UTF8]
   --encoding value        バックエンドのエンコーディング。 (デフォルト: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     隠しファイルとフォルダのリスト表示にLIST -aを使用します。これにより、MLSDの使用が無効になります。 (デフォルト: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    アイドル状態の接続を閉じる前の最大時間。 (デフォルト: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  サーバーのTLS証明書を検証しない。 (デフォルト: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    データ接続のクローズステータスを待つ最大時間。 (デフォルト: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  全ての制御とデータ接続のTLSセッションキャッシュのサイズ。 (デフォルト: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          修正時刻を設定するためにMDTMを使用します（VsFtpdのクワーク） (デフォルト: false) [$WRITING_MDTM]

   General

   --name value  ストレージの名前（デフォルト: 自動生成）
   --path value  ストレージのパス

```
{% endcode %}