# 自己取引提案のSPADEポリシーを作成する

{% code fullWidth="true" %}
```
NAME:
   singularity deal spade-policy create - 自己取引提案のSPADEポリシーを作成する

使い方:
   singularity deal spade-policy create [コマンドオプション] DATASET_NAME [...PROVIDER_ID]

オプション:
   --min-delay value     取引開始時の最小遅延日数（デフォルト: 3）
   --max-delay value     取引開始時の最大遅延日数（デフォルト: 3）
   --min-duration value  取引開始時の最小期間日数（デフォルト: 535）
   --max-duration value  取引開始時の最大期間日数（デフォルト: 535）
   --verified            検証済みの取引を提案するかどうか（デフォルト: true）
   --price value         取引価格（32GiB単位で計測し、全期間にわたる価格）（デフォルト: 0）
   --help, -h            ヘルプを表示する
```
{% endcode %}