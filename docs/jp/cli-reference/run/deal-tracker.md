# 関連するウォレットの取引をトラッキングするディールトラッカーを開始します

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-tracker - 関連するウォレットの取引をトラッキングするディールトラッカーを開始します

使用法:
   singularity run deal-tracker [コマンドオプション] [引数...]

オプション:
   --market-deal-url value, -m value  ZST 圧縮状態のマーケットディール JSON の URL。空にすると Lotus API を使用します。 (デフォルト: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         新しい取引をチェックする頻度 (デフォルト: 1h0m0s)
   --help, -h                         ヘルプを表示
```
{% endcode %}