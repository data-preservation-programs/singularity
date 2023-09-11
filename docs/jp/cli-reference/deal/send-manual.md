# ブーストまたは旧来のマーケットに手動のディール提案を送信する

```
NAME:
   singularity deal send-manual - ブーストまたは旧来のマーケットに手動のディール提案を送信する

USAGE:
   singularity deal send-manual [コマンドオプション] <クライアント> <プロバイダ> <ピース_CID> <ピース_サイズ>

DESCRIPTION:
   ブーストまたは旧来のマーケットに手動のディール提案を送信する
     例: singularity deal send-manual f01234 f05678 bagaxxxx 32GiB
   注意:
     * クライアントアドレスは、'singularity wallet import' を使用してウォレットにインポートする必要があります
     * ディール提案はデータベースに保存されませんが、ディールトラッカーが実行されている場合は最終的に追跡されます
     * LOTUS_API および LOTUS_TOKEN を独自のlotusノードに設定することで、GLIF APIを使用したクイックアドレス検証が可能です

OPTIONS:
   --help, -h       ヘルプの表示
   --timeout value  ディール提案のタイムアウト（デフォルト: 1m）

   Boostのみ

   --file-size value                            BoostがCARファイルを取得するためのファイルサイズ（デフォルト: 0）
   --http-header value [ --http-header value ]  リクエストと一緒に渡すhttpヘッダ（キー=値形式）
   --ipni                                       ディールをIPNIに公開するかどうか（デフォルト: true）
   --url-template value                         CARファイルを取得するためのPIECE_CIDプレースホルダを持つURLテンプレート、例：http://127.0.0.1/piece/{PIECE_CID}.car

   ディール提案

   --client value                 ディールの送信元となるクライアントアドレス
   --duration value, -d value     エポックまたは期間形式での期間、例：1500000、2400h（デフォルト: 12840h[535日]）
   --keep-unsealed                未シール状態のコピーを保持するかどうか（デフォルト: true）
   --piece-cid value              ディールのピースCID
   --piece-size value             ディールのピースサイズ（デフォルト: "32GiB"）
   --price-per-deal value         ディールごとのFIL単価（デフォルト: 0）
   --price-per-gb value           GiBあたりのFIL単価（デフォルト: 0）
   --price-per-gb-epoch value     エポックごとのGiBあたりのFIL単価（デフォルト: 0）
   --provider value               ディールの送信先となるストレージプロバイダID
   --root-cid value               ディール提案の一部として必要なルートCID。空の場合、空のCIDに設定されます（デフォルト: 空のCID）
   --start-delay value, -s value  ディールの開始遅延をエポックまたは期間形式で指定します。例：1000、72h（デフォルト: 72h[3日]）
   --verified                     ディールを検証済みとして提案するかどうか（デフォルト: true）
```