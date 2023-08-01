# Mail.ru Cloud

{% code fullWidth="true" %}
```
名前:
   singularity datasource add mailru - Mail.ru Cloud

使用法:
   singularity datasource add mailru [コマンドオプション] <データセット名> <ソースパス>

概要:
   --mailru-check-hash
      ファイルのチェックサムが不一致または無効な場合に、copyコマンドがどのように動作するか選択します。

      例:
         | true  | エラーで失敗します。
         | false | 無視して続行します。

   --mailru-encoding
      バックエンドのエンコーディング方法です。

      詳細については[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --mailru-pass
      パスワードです。

      これはアプリパスワードである必要があります。通常のパスワードではrcloneを使用することはできません。アプリパスワードの作成方法については、ドキュメントの設定セクションを参照してください。

   --mailru-quirks
      コンマ区切りの内部メンテナンスフラグのリストです。

      このオプションは一般ユーザーには使用しないでください。バックエンドの問題をリモートでトラブルシューティングするためのものです。フラグの正確な意味は文書化されておらず、リリース間で保証されることもありません。バックエンドが安定すると、クセは削除されます。
      サポートされているクセ: atomicmkdir binlist unknowndirs

   --mailru-speedup-enable
      同じデータハッシュを持つ別のファイルがある場合、完全なアップロードをスキップします。

      この機能は「スピードアップ」または「ハッシュによるアップロード」と呼ばれます。一般的に利用可能なファイル（人気のある本、ビデオ、オーディオクリップなど）の場合、ファイルはすべてのmailruユーザーのアカウントでハッシュで検索されます。ソースファイルがユニークまたは暗号化されている場合、それは意味がなく効果がありません。rcloneは使用する前にコンテンツハッシュを事前に計算し、完全なアップロードが必要かどうかを判断するために、ローカルのメモリとディスクスペースを必要とする場合があります。また、rcloneがファイルサイズを事前に知らない場合（ストリーミングや部分的なアップロードの場合など）、この最適化は試されません。

      例:
         | true  | 有効にする
         | false | 無効にする

   --mailru-speedup-file-patterns
      スピードアップ（ハッシュによるアップロード）の対象となるファイル名パターンのコンマ区切りのリストです。

      パターンは大文字と小文字を区別せず、「*」または「?」のメタ文字を含むことができます。

      例:
         | <unset>                 | リストを空にすると、スピードアップ（ハッシュによるアップロード）が完全に無効になります。
         | *                       | すべてのファイルがスピードアップの対象になります。
         | *.mkv,*.avi,*.mp4,*.mp3 | 一般的なオーディオ/ビデオファイルのみがスピードアップの対象になります。
         | *.zip,*.gz,*.rar,*.pdf  | 一般的なアーカイブまたはPDFブックのみがスピードアップの対象になります。

   --mailru-speedup-max-disk
      大きなファイルの場合、このオプションによりスピードアップ（ハッシュによるアップロード）が無効にできます。

      理由は、事前のハッシングがRAMやディスクスペースを使い果たす可能性があるためです。

      例:
         | 0    | スピードアップ（ハッシュによるアップロード）を完全に無効にします。
         | 1G   | 1GBより大きいファイルは直接アップロードされます。
         | 3G   | ローカルディスクの空き容量が3GB未満の場合にこのオプションを選択します。

   --mailru-speedup-max-memory
      以下のサイズよりも大きなファイルは常にディスク上でハッシュが作成されます。

      例:
         | 0    | 事前ハッシュは常に一時的なディスク領域で行われます。
         | 32M  | 事前ハッシュのために32MB以上のRAMを割り当てないでください。
         | 256M | ハッシュ計算に最大で256MBのRAMが利用可能です。

   --mailru-user
      ユーザー名（通常はメールアドレス）です。

   --mailru-user-agent
      クライアントが内部で使用するHTTPユーザーエージェントです。

      デフォルトは "rclone/VERSION" またはコマンドラインで指定された "--user-agent" です。


オプション:
   --help, -h  ヘルプを表示する

   データの準備オプション

   --delete-after-export   [危険] データセットをCARファイルにエクスポートした後、データセットのファイルを削除します。  (デフォルト: false)
   --rescan-interval value 最後の正常スキャンから指定した間隔が経過すると、自動的にソースディレクトリを再スキャンします (デフォルト: 無効)
   --scanning-state value  初期のスキャン状態を設定します (デフォルト: ready)

   Mail.ruのオプション

   --mailru-check-hash value             ファイルのチェックサムが不一致または無効な場合に、copyコマンドがどのように動作するか選択します。 (デフォルト: "true") [$MAILRU_CHECK_HASH]
   --mailru-encoding value               バックエンドのエンコーディング方法です。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Del,Ctl,InvalidUtf8,Dot") [$MAILRU_ENCODING]
   --mailru-pass value                   パスワードです。 [$MAILRU_PASS]
   --mailru-speedup-enable value         同じデータハッシュを持つ別のファイルがある場合、完全なアップロードをスキップします。 (デフォルト: "true") [$MAILRU_SPEEDUP_ENABLE]
   --mailru-speedup-file-patterns value  スピードアップ（ハッシュによるアップロード）の対象となるファイル名パターンのコンマ区切りのリストです。 (デフォルト: "*.mkv,*.avi,*.mp4,*.mp3,*.zip,*.gz,*.rar,*.pdf") [$MAILRU_SPEEDUP_FILE_PATTERNS]
   --mailru-speedup-max-disk value       大きなファイルの場合、このオプションによりスピードアップ（ハッシュによるアップロード）が無効にできます。 (デフォルト: "3Gi") [$MAILRU_SPEEDUP_MAX_DISK]
   --mailru-speedup-max-memory value     以下のサイズよりも大きなファイルは常にディスク上でハッシュが作成されます。 (デフォルト: "32Mi") [$MAILRU_SPEEDUP_MAX_MEMORY]
   --mailru-user value                   ユーザー名（通常はメールアドレス）です。 [$MAILRU_USER]

```
{% endcode %}