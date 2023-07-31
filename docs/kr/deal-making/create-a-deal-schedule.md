# 거래 일정 만들기

저장 공급 업체와 거래를 하기 위한 시간입니다. 거래를 진행하기 위해 다음 명령어를 실행하십시오.

```
singularity run dealmaker
```

## 한 번에 모든 거래 보내기

데이터셋이 작다면, 모든 거래를 한 번에 저장 공급 업체에 보낼 수 있습니다. 이를 위해서는 다음 명령어를 사용할 수 있습니다.

```sh
singularity deal schedule create dataset_name provider_id
```

하지만 데이터셋이 큰 경우, 거래 제안이 만료되기 전에 저장 공급 업체가 처리해야 하는 거래의 수가 너무 많을 수 있습니다. 이 경우, 일정을 만들 수 있습니다.

## 일정을 사용하여 거래 보내기

같은 명령어를 사용하여 거래의 속도와 빈도를 제어하는 자체 일정을 만들 수 있습니다.

```
--schedule-deal-number value, --number value     트리거된 일정 당 최대 거래 수, 예: 30 (기본값: 무제한)
--schedule-deal-size value, --size value         트리거된 일정 당 최대 거래 크기, 예: 500GB (기본값: 무제한)
--schedule-interval value, --every value         배치 거래를 보내기 위한 Cron 일정 (기본값: 비활성화됨)
--total-deal-number value, --total-number value  이 요청에 대한 최대 총 거래 수, 예: 1000 (기본값: 무제한)
```