# 启动一个交易跟踪器，用来跟踪所有相关钱包的交易

{% code fullWidth="true" %}
```
命令名称:
   singularity run deal-tracker - 启动一个交易跟踪器，用来跟踪所有相关钱包的交易

使用方法:
   singularity run deal-tracker [命令选项] [参数...]

选项:
   --market-deal-url value, -m value  ZST 压缩状态市场交易 JSON 的 URL。设置为空以使用 Lotus API。 (默认值: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         检查新交易的频率 (默认值: 1h0m0s)
   --help, -h                         显示帮助
```
{% endcode %}