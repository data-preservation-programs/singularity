# FTP

{% code fullWidth="true" %}
```
NAME:
   singularity storage update ftp - FTP

USAGE:
   singularity storage update ftp [command options] <name|id>

DESCRIPTION:
   --host
      接続するFTPホスト。
      
      例: "ftp.example.com"。

   --user
      FTPユーザー名。

   --port
      FTPポート番号。

   --pass
      FTPパスワード。

   --tls
      暗黙的FTP経由のTLS (FTPS) を使用します。
      
      暗黙的FTP経由のTLSを使用すると、クライアントは開始時からTLSを使用して接続するため、
      非TLS対応のサーバーとの互換性がなくなります。通常、ポート21ではなくポート990経由で
      提供されます。明示的FTP経由のTLSとの組み合わせは使用できません。

   --explicit-tls
      明示的FTP経由のTLS (FTPS) を使用します。
      
      明示的FTP経由のTLSを使用すると、クライアントは平文接続を暗号化接続にアップグレードするために、
      サーバーにセキュリティを要求します。明示的FTP経由のTLSとの組み合わせは使用できません。

   --concurrency
      最大同時FTP接続数、0は制限なし。
      
      これを設定すると、デッドロックが発生する可能性が非常に高いため、注意して使用する必要があります。
      
      同期またはコピーを行っている場合は、`--transfers`と`--checkers`の合計より1つ多くなるように設定してください。
      
      `--check-first`を使用している場合は、`--checkers`と`--transfers`の最大値より1つ多くなる必要があります。
      
      したがって、`concurrency 3`の場合、`--checkers 2 --transfers 2 --check-first`または`--checkers 1 --transfers 1`を使用します。

   --no-check-certificate
      サーバーのTLS証明書を検証しません。

   --disable-epsv
      サーバーがサポートを広告していても、EPSVを使用しないようにします。

   --disable-mlsd
      サーバーがサポートを広告していても、MLSDを使用しないようにします。

   --disable-utf8
      サーバーがサポートを広告していても、UTF-8を使用しないようにします。

   --writing-mdtm
      MDTMを使用して修正時刻を設定します（VsFtpdの特異点）

   --force-list-hidden
      隠しファイルとフォルダを強制的にリストするためにLIST -aを使用します。これにより、MLSDの使用が無効になります。

   --idle-timeout
      アイドル接続を閉じるまでの最大時間。
      
      指定された時間内にコネクションプールに戻った接続がない場合、rcloneはコネクションプールを空にします。
      
      0に設定すると、接続を無期限に保持します。

   --close-timeout
      クローズ要求への応答待ちの最大時間。

   --tls-cache-size
      すべてのコントロールおよびデータ接続のTLSセッションキャッシュのサイズ。
      
      TLSキャッシュは、TLSセッションを再開し、接続間でPSKを再利用することができます。
      デフォルトのサイズが十分でない場合は、増やしてTLSの再開エラーを回避します。
      デフォルトでは有効になっています。無効にするには0を使用します。

   --disable-tls13
      TLS 1.3を無効にします（TLSのバグがあるFTPサーバー用の回避策）

   --shut-timeout
      データ接続のクローズステータス待ちの最大時間。

   --ask-password
      必要に応じてFTPパスワードを尋ねることを許可します。
      
      これが設定されていてパスワードが指定されていない場合、rcloneがパスワードを要求します。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

      例:
         | Asterisk,Ctl,Dot,Slash                               | ProFTPdはファイル名に'*'を処理できません
         | BackSlash,Ctl,Del,Dot,RightSpace,Slash,SquareBracket | PureFTPdはファイル名に'[]'または'*'を処理できません
         | Ctl,LeftPeriod,Slash                                 | VsFTPdはドットで始まるファイル名を処理できません


OPTIONS:
   --explicit-tls  明示的FTP経由のTLS (FTPS)を使用します。 (default: false) [$EXPLICIT_TLS]
   --help, -h      ヘルプの表示
   --host value    接続するFTPホスト。 [$HOST]
   --pass value    FTPパスワード。 [$PASS]
   --port value    FTPポート番号。 (default: 21) [$PORT]
   --tls           暗黙的FTP経由のTLS (FTPS)を使用します。 (default: false) [$TLS]
   --user value    FTPユーザー名。 (default: "$USER") [$USER]

   Advanced

   --ask-password          必要に応じてFTPパスワードを尋ねることを許可します。 (default: false) [$ASK_PASSWORD]
   --close-timeout value   クローズ要求への応答待ちの最大時間。 (default: "1m0s") [$CLOSE_TIMEOUT]
   --concurrency value     最大同時FTP接続数、0は制限なし。 (default: 0) [$CONCURRENCY]
   --disable-epsv          サーバーがサポートを広告していても、EPSVを使用しないようにします。 (default: false) [$DISABLE_EPSV]
   --disable-mlsd          サーバーがサポートを広告していても、MLSDを使用しないようにします。 (default: false) [$DISABLE_MLSD]
   --disable-tls13         TLS 1.3を無効にします（TLSのバグがあるFTPサーバー用の回避策） (default: false) [$DISABLE_TLS13]
   --disable-utf8          サーバーがサポートを広告していても、UTF-8を使用しないようにします。 (default: false) [$DISABLE_UTF8]
   --encoding value        バックエンドのエンコーディング。 (default: "Slash,Del,Ctl,RightSpace,Dot") [$ENCODING]
   --force-list-hidden     隠しファイルとフォルダを強制的にリストするためにLIST -aを使用します。これにより、MLSDの使用が無効になります。 (default: false) [$FORCE_LIST_HIDDEN]
   --idle-timeout value    アイドル接続を閉じるまでの最大時間。 (default: "1m0s") [$IDLE_TIMEOUT]
   --no-check-certificate  サーバーのTLS証明書を検証しません。 (default: false) [$NO_CHECK_CERTIFICATE]
   --shut-timeout value    データ接続のクローズステータス待ちの最大時間。 (default: "1m0s") [$SHUT_TIMEOUT]
   --tls-cache-size value  すべてのコントロールおよびデータ接続のTLSセッションキャッシュのサイズ。 (default: 32) [$TLS_CACHE_SIZE]
   --writing-mdtm          MDTMを使用して修正時刻を設定します（VsFtpdの特異点） (default: false) [$WRITING_MDTM]

```
{% endcode %}