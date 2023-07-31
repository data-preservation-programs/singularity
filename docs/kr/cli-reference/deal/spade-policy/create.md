# 자체 거래 제안에 대한 SPADE 정책 생성

{% code fullWidth="true" %}
```
이름:
   singularity deal spade-policy create - 자체 거래 제안에 대한 SPADE 정책 생성

사용법:
   singularity deal spade-policy create [command options] DATASET_NAME [...PROVIDER_ID]

옵션:
   --min-delay value     거래 시작 에포크를 위한 최소 지연 일수 (기본값: 3)
   --max-delay value     거래 시작 에포크를 위한 최대 지연 일수 (기본값: 3)
   --min-duration value  거래 시작 에포크를 위한 최소 지속 일수 (기본값: 535)
   --max-duration value  거래 시작 에포크를 위한 최대 지속 일수 (기본값: 535)
   --verified            검증된 거래로 거래 제안 여부 (기본값: true)
   --price value         거래의 가격 (32GiB당) (기본값: 0)
   --help, -h            도움말 표시
```
{% endcode %}