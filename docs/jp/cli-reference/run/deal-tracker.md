# 関連するすべてのウォレットの取引を追跡するディールトラッカーを起動する

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-tracker - 関連するすべてのウォレットの取引を追跡するディールトラッカーを起動する

使用法:
   singularity run deal-tracker [オプション] [引数...]

オプション:
   --market-deal-url value, -m value  ZST圧縮状態のマーケット取引JSONのURL。LotusのAPIを使用する場合は空に設定します。 (デフォルト: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         新しい取引を確認する間隔 (デフォルト: 1h0m0s)
   --once                             一度だけ実行して終了する (デフォルト: false)
   --help, -h                         ヘルプを表示する
```
{% endcode %}