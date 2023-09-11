# ストレージプロバイダにディールを送信するスケジュールを作成する

{% code fullWidth="true" %}
```
NAME:
   singularity deal schedule create - ストレージプロバイダにディールを送信するスケジュールを作成する

使用法:
   singularity deal schedule create [command options] [arguments...]

説明:
   CRON パターン '--schedule-cron': CRON パターンは、ディスクリプタまたはオプションのセカンドフィールドを持つ標準の CRON パターンであることができます。
     標準の CRON:
       ┌───────────── 分 (0 - 59)
       │ ┌───────────── 時 (0 - 23)
       │ │ ┌───────────── 月の日 (1 - 31)
       │ │ │ ┌───────────── 月 (1 - 12)
       │ │ │ │ ┌───────────── 曜日 (0 - 6) (日曜日から土曜日)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     オプションのセカンドフィールド:
       ┌─────────────  秒 (0 - 59)
       │ ┌─────────────  分 (0 - 59)
       │ │ ┌─────────────  時 (0 - 23)
       │ │ │ ┌─────────────  月の日 (1 - 31)
       │ │ │ │ ┌─────────────  月 (1 - 12)
       │ │ │ │ │ ┌─────────────  曜日 (0 - 6) (日曜日から土曜日)
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     ディスクリプタ:
       @yearly, @annually - 0 0 1 1 * と等価
       @monthly           - 0 0 1 * * と等価
       @weekly            - 0 0 * * 0 と等価
       @daily,  @midnight - 0 0 * * * と等価
       @hourly            - 0 * * * * と等価

オプション:
   --help, -h               ヘルプを表示
   --preparation value      プリパレーションのIDまたは名前
   --provider value         ディールを送信するストレージプロバイダのID

   Boost のみ

   --http-header value, -H value [ --http-header value, -H value ]  リクエストと一緒に渡される HTTP ヘッダー（キー=値の形式）。
   --ipni                                                           IPNI にディールを公開するかどうか（デフォルト: true）
   --url-template value, -u value                                   PIECE_CID のプレースホルダーを使用して、CAR ファイルをフェッチするための URL テンプレート。例: http://127.0.0.1/piece/{PIECE_CID}.car

   ディール提案

   --duration value, -d value     エポック形式または期間形式のデュレーション。例: 1500000、2400h（デフォルト: 12840h [535 日間]）
   --keep-unsealed                シールされていないコピーを保持するかどうか（デフォルト: true）
   --price-per-deal value         ディールあたりの FIL 価格（デフォルト: 0）
   --price-per-gb value           GiB あたりの FIL 価格（デフォルト: 0）
   --price-per-gb-epoch value     エポックあたりの GiB あたりの FIL 価格（デフォルト: 0）
   --start-delay value, -s value  ディール開始の遅延を秒単位または期間形式で指定。例: 1000、72h（デフォルト: 72h [3 日間]）
   --verified                     ディールを検証済みとして提案するかどうか（デフォルト: true）

   制限

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      このスケジュールで許可されているピース CID のリスト（デフォルト: 任意）
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  許可されているピース CID のリストが含まれているファイルのリスト
   --max-pending-deal-number value, --pending-number value                                                            このリクエスト全体での最大未処理のディール数。例: 100TiB（デフォルト: 無制限）
   --max-pending-deal-size value, --pending-size value                                                                このリクエスト全体での最大未処理のディールサイズ。例: 1000（デフォルト: 無制限）
   --total-deal-number value, --total-number value                                                                    このリクエスト全体での最大合計ディール数。例: 1000（デフォルト: 無制限）
   --total-deal-size value, --total-size value                                                                        このリクエスト全体での最大合計ディールサイズ。例: 100TiB（デフォルト: 無制限）

   スケジューリング

   --schedule-cron value, --cron value           バッチディールを送信するためのクーロンスケジュール（デフォルト: 無効）
   --schedule-deal-number value, --number value  トリガされたスケジュールあたりの最大ディール数。例: 30（デフォルト: 無制限）
   --schedule-deal-size value, --size value      トリガされたスケジュールあたりの最大ディールサイズ。例: 500GiB（デフォルト: 無制限）

   トラッキング

   --notes value, -n value  リクエストと一緒に保存するメモやタグ。トラッキングの目的
```
{% endcode %}