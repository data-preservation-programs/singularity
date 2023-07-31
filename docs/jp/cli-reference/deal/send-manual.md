# ブーストまたはレガシーマーケットに手動でディールプロポーザルを送信する

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual - ブーストまたはレガシーマーケットに手動でディールプロポーザルを送信する

使用法:
   singularity deal send-manual [コマンドオプション] クライアントアドレス プロバイダーID PIECE_CID PIECE_SIZE

オプション:
   --help, -h       ヘルプを表示します
   --timeout value  ディールプロポーザルのタイムアウト（デフォルト: 1m）

   ブーストのみ

   --file-size value                            CARファイルをフェッチするためのファイルサイズ（デフォルト: 0）
   --http-header value [ --http-header value ]  リクエストに渡されるHTTPヘッダー（キー=値）
   --ipni                                       ディールをIPNIにアナウンスするかどうか（デフォルト: true）
   --url-template value                         ブーストがCARファイルをフェッチするためのPIECE_CIDプレースホルダを含むURLテンプレート、例: http://127.0.0.1/piece/{PIECE_CID}.car

   ディールプロポーザル

   --duration value, -d value     エポックまたは期間形式の期間（デフォルト: 12840h[535 days]）
   --keep-unsealed                未検封コピーを保持するかどうか（デフォルト: true）
   --price-per-deal value         ディールごとのFIL単位の価格（デフォルト: 0）
   --price-per-gb value           GiB単位のFILの価格（デフォルト: 0）
   --price-per-gb-epoch value     エポックごとのGiB単位のFILの価格（デフォルト: 0）
   --root-cid value               ディールプロポーザルの一部として必要なルートCID（空の場合、空のCIDに設定されます）（デフォルト: Empty CID）
   --start-delay value, -s value  エポックまたは期間形式のディール開始遅延時間（デフォルト: 72h[3 days]）
   --verified                     検証済みとしてディールを提案するかどうか（デフォルト: true）
```
{% endcode %}