# ストレージプロバイダにディールを送信するスケジュールを作成します

{% code fullWidth="true" %}
```
NAME:
   singularity deal schedule create - ストレージプロバイダにディールを送信するスケジュールを作成します

USAGE:
   singularity deal schedule create [command options] DATASET_NAME PROVIDER_ID

OPTIONS:
   --help, -h  ヘルプを表示する

   Boost のみ

   --http-header value, -H value [ --http-header value, -H value ]  リクエストと一緒に渡される HTTP ヘッダのリスト (例: key=value)
   --ipni                                                           ディールを IPNI に公開するかどうか (デフォルト: true)
   --url-template value, -u value                                   PIECE_CID のプレースホルダー PIECE_CID を使用して CAR ファイルを取得するための URL テンプレート (例: http://127.0.0.1/piece/{PIECE_CID}.car)

   ディール提案

   --duration value, -d value     エポックまたは期間形式で指定された期間 (例: 1500000, 2400h) (デフォルト: "12840h")
   --keep-unsealed                未密封コピーを保持するかどうか (デフォルト: true)
   --price-per-deal value         ディール単位の価格 (FIL 単位) (デフォルト: 0)
   --price-per-gb value           GiB 単位の価格 (FIL 単位) (デフォルト: 0)
   --price-per-gb-epoch value     エポックあたりの GiB 単位の価格 (FIL 単位) (デフォルト: 0)
   --start-delay value, -s value  ディールの開始遅延をエポックまたは期間形式で指定 (例: 1000, 72h) (デフォルト: "72h")
   --verified                     ディールを検証済みとして提案するかどうか (デフォルト: true)

   制限事項

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      このスケジュールで許可されたピース CID のリスト (デフォルト: Any)
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  許可するピース CID のリストが含まれるファイルのリスト
   --max-pending-deal-number value, --pending-number value                                                            このリクエスト全体での最大保留ディール数 (デフォルト: 無制限)
   --max-pending-deal-size value, --pending-size value                                                                このリクエスト全体での最大保留ディールサイズ (デフォルト: 無制限)
   --total-deal-size value, --total-size value                                                                        このリクエスト全体での最大総ディールサイズ (例: 100TB) (デフォルト: 無制限)

   スケジューリング

   --schedule-cron value, --cron value              バッチディールを送信するための Cron スケジュール (デフォルト: 無効)
   --schedule-deal-number value, --number value     1 つのトリガーされたスケジュールあたりの最大ディール数 (例: 30) (デフォルト: 無制限)
   --schedule-deal-size value, --size value         1 つのトリガーされたスケジュールあたりの最大ディールサイズ (例: 500GB) (デフォルト: 無制限)
   --total-deal-number value, --total-number value  このリクエスト全体での最大総ディール数 (例: 1000) (デフォルト: 無制限)

   追跡

   --notes value, -n value  リクエストと一緒に保存されるノートやタグ。追跡の目的で使用されます

```
{% endcode %}