# 모든 관련 지갑의 거래를 추적하는 거래 추적기 시작하기

{% code fullWidth="true" %}
```
이름:
   singularity run deal-tracker - 모든 관련 지갑의 거래를 추적하는 거래 추적기 시작하기

사용법:
   singularity run deal-tracker [command options] [arguments...]

옵션:
   --market-deal-url value, -m value  ZST 압축 상태 시장 거래 json의 URL. Lotus API를 사용하려면 비워 두십시오. (기본값: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         새로운 거래를 확인하는 간격 (기본값: 1h0m0s)
   --help, -h                         도움말 표시
```
{% endcode %}