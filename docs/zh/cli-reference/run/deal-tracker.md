# 启动一个交易跟踪器，跟踪所有相关钱包的交易

{% code fullWidth="true" %}
```
NAME:
   singularity run deal-tracker - 启动一个交易跟踪器，跟踪所有相关钱包的交易

使用方法:
   singularity run deal-tracker [command options] [arguments...]

选项:
   --market-deal-url value, -m value  ZST压缩状态市场交易json的URL。设为空以使用Lotus API。 (默认值："https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         检查新交易的频率 (默认值：1h0m0s)
   --once                             运行一次并退出 (默认值：false)
   --help, -h                         显示帮助信息
```
{% endcode %}