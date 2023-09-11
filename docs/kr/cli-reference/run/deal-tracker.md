# 모든 관련 지갑 거래를 추적하는 거래 추적기 시작하기

{% code fullWidth="true" %}
```
명령:
   singularity run deal-tracker - 모든 관련 지갑 거래를 추적하는 거래 추적기 시작

사용법:
   singularity run deal-tracker [command 옵션] [arguments...]

옵션:
   --market-deal-url value, -m value  ZST 압축 상태 거래 json에 대한 URL입니다. Lotus API를 사용하려면 빈 값으로 설정하십시오. (기본값: "https://marketdeals.s3.amazonaws.com/StateMarketDeals.json.zst") [$MARKET_DEAL_URL]
   --interval value, -i value         새로운 거래를 확인하는 간격입니다. (기본값: 1h0m0s)
   --once                             한 번 실행 후 종료합니다. (기본값: false)
   --help, -h                         도움말 보기
```
{% endcode %}