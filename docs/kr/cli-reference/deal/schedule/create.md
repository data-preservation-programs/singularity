# 스토리지 제공자에게 거래를 전송하기 위한 일정 생성

{% code fullWidth="true" %}
```
이름:
   singularity deal schedule create - 스토리지 제공자에게 거래를 전송하기 위한 일정 생성

사용법:
   singularity deal schedule create [command options] DATASET_NAME PROVIDER_ID

옵션:
   --help, -h  도움말 표시

   부스트 옵션

   --http-header value, -H value [ --http-header value, -H value ]  요청과 함께 전달할 HTTP 헤더 (예: key=value)
   --ipni                                                           거래를 IPNI에 알리는지 여부 (기본값: true)
   --url-template value, -u value                                   부스트가 CAR 파일을 가져오기 위한 PIECE_CID를 포함한 URL 템플릿, 예: http://127.0.0.1/piece/{PIECE_CID}.car

   거래 제안

   --duration value, -d value     기간을 에포크 시간 또는 기간 형식으로 지정 (예: 1500000, 2400h) (기본값: "12840h")
   --keep-unsealed                미립자를 유지할 지 여부 (기본값: true)
   --price-per-deal value         거래 당 FIL 가격 (기본값: 0)
   --price-per-gb value           GiB 당 FIL 가격 (기본값: 0)
   --price-per-gb-epoch value     에포크 당 GiB 당 FIL 가격 (기본값: 0)
   --start-delay value, -s value  거래 시작 지연 시간을 에포크 시간 또는 기간 형식으로 지정 (예: 1000, 72h) (기본값: "72h")
   --verified                     검증된 거래로 제안할지 여부 (기본값: true)

   제한 사항

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      이 일정에서 허용된 조각 CID 목록 (기본값: Any)
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  조각 CID 목록이 포함된 파일 목록을 허용
   --max-pending-deal-number value, --pending-number value                                                            이 요청에 대한 전체 대기 중인 거래 번호 제한 (기본값: 무제한)
   --max-pending-deal-size value, --pending-size value                                                                이 요청에 대한 전체 대기 중인 거래 크기 제한 (기본값: 무제한)
   --total-deal-size value, --total-size value                                                                        이 요청에 대한 최대 전체 거래 크기, 예: 100TB (기본값: 무제한)

   일정 설정

   --schedule-cron value, --cron value              기록된 거래를 전송하는데 사용되는 cron 일정 (기본값: 비활성화)
   --schedule-deal-number value, --number value     트리거된 일정당 최대 거래 수, 예: 30 (기본값: 무제한)
   --schedule-deal-size value, --size value         트리거된 일정당 최대 거래 크기, 예: 500GB (기본값: 무제한)
   --total-deal-number value, --total-number value  이 요청에 대한 최대 총 거래 수, 예: 1000 (기본값: 무제한)

   추적

   --notes value, -n value  추적을 위해 요청과 함께 저장되는 추가 정보 또는 태그

```
{% endcode %}