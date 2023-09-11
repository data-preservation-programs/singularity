# スケジュールを更新する

{% code fullWidth="true" %}
```
NAME:
   singularity deal schedule update - スケジュールを更新する

USAGE:
   singularity deal schedule update [command options] <schedule_id>

DESCRIPTION:
   CRONパターン '--schedule-cron': CRONパターンは、ディスクリプタまたはオプションの秒フィールドを持つ標準のCRONパターンである場合があります。
     標準のCRONパターン:
       ┌───────────── 分 (0 - 59)
       │ ┌───────────── 時 (0 - 23)
       │ │ ┌───────────── 月の日 (1 - 31)
       │ │ │ ┌───────────── 月 (1 - 12)
       │ │ │ │ ┌───────────── 曜日 (0 - 6) (日曜日から土曜日まで)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     オプションの秒フィールド:
       ┌─────────────  秒 (0 - 59)
       │ ┌─────────────  分 (0 - 59)
       │ │ ┌─────────────  時 (0 - 23)
       │ │ │ ┌─────────────  月の日 (1 - 31)
       │ │ │ │ ┌─────────────  月 (1 - 12)
       │ │ │ │ │ ┌─────────────  曜日 (0 - 6) (日曜日から土曜日まで)
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     ディスクリプタ:
       @yearly, @annually - 等価：0 0 1 1 *
       @monthly           - 等価：0 0 1 * *
       @weekly            - 等価：0 0 * * 0
       @daily,  @midnight - 等価：0 0 * * *
       @hourly            - 等価：0 * * * *

OPTIONS:
   --help, -h  ヘルプを表示

   Boostのみ

   --http-header value, -H value [ --http-header value, -H value ]  リクエストに渡すHTTPヘッダ（キー=値）。これによって既存のヘッダ値が置換されます。ヘッダを削除するには、--http-header "key="を使用します。すべてのヘッダを削除するには、--http-header ""を使用します。
   --ipni                                                           取引をIPNIに発表するかどうか（デフォルト：true）
   --url-template value, -u value                                   BoostがCARファイルを取得するためのPIECE_CIDプレースホルダを含むURLテンプレート、例：http://127.0.0.1/piece/{PIECE_CID}.car

   取引提案

   --duration value, -d value     エポックまたは期間形式での期間、例：1500000、2400h
   --keep-unsealed                非封印のコピーを保持するかどうか（デフォルト：true）
   --price-per-deal value         取引あたりのFIL単価（デフォルト：0）
   --price-per-gb value           GiB当たりのFIL単価（デフォルト：0）
   --price-per-gb-epoch value     エポックあたりのGiB当たりのFIL単価（デフォルト：0）
   --start-delay value, -s value  取引の開始遅延をエポックまたは期間形式で指定します。例：1000、72h
   --verified                     驗證された取引として提案するかどうか（デフォルト：true）

   制約

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      このスケジュールで許可されるピースCIDのリスト。追加のみ。
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  許可するピースCIDのリストが含まれるファイルのリスト。追加のみ。
   --max-pending-deal-number value, --pending-number value                                                            このリクエスト全体の最大の保留中の取引数、例：100TiB（デフォルト：0）
   --max-pending-deal-size value, --pending-size value                                                                このリクエスト全体の最大の保留中の取引サイズ、例：1000
   --total-deal-number value, --total-number value                                                                    このリクエストの最大の総取引数、例：1000（デフォルト：0）
   --total-deal-size value, --total-size value                                                                        このリクエストの最大の総取引サイズ、例：100TiB

   スケジューリング

   --schedule-cron value, --cron value           バッチ取引を送信するためのスケジュールのCRON
   --schedule-deal-number value, --number value  1度のトリガーごとの最大取引数、例：30（デフォルト：0）
   --schedule-deal-size value, --size value      1度のトリガーごとの最大取引サイズ、例：500GiB

   追跡

   --notes value, -n value  リクエストと一緒に保存するメモやタグ、追跡目的のため

```
{% endcode %}