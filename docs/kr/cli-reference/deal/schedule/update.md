# 기존 일정 업데이트

{% code fullWidth="true" %}
```
이름:
   singularity deal schedule update - 기존 일정 업데이트

사용법:
   singularity deal schedule update [명령어 옵션] <일정_아이디>

설명:
   CRON 패턴 '--schedule-cron': CRON 패턴은 기술자(descriptor) 또는 선택적 두 번째 필드를 가진 표준 CRON 패턴일 수 있습니다.
     표준 CRON:
       ┌───────────── 분 (0 - 59)
       │ ┌───────────── 시간 (0 - 23)
       │ │ ┌───────────── 월 일 (1 - 31)
       │ │ │ ┌───────────── 월 (1 - 12)
       │ │ │ │ ┌───────────── 요일 (0 - 6) (일요일부터 토요일까지)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     선택적 두 번째 필드:
       ┌─────────────  초 (0 - 59)
       │ ┌─────────────  분 (0 - 59)
       │ │ ┌─────────────  시간 (0 - 23)
       │ │ │ ┌─────────────  월 일 (1 - 31)
       │ │ │ │ ┌─────────────  월 (1 - 12)
       │ │ │ │ │ ┌─────────────  요일 (0 - 6) (일요일부터 토요일까지)
       │ │ │ │ │ │
       │ │ │ │ │ │
       * * * * * *

     기술자:
       @yearly, @annually - 0 0 1 1 *
       @monthly           - 0 0 1 * *
       @weekly            - 0 0 * * 0
       @daily,  @midnight - 0 0 * * *
       @hourly            - 0 * * * *

옵션:
   --help, -h  도움말 표시

   Boost Only

   --http-header value, -H value [ --http-header value, -H value ]  요청과 함께 전달될 HTTP 헤더(key=value 형식)입니다. 기존 헤더 값을 대체합니다. 헤더를 제거하려면 --http-header "key=""". 모든 헤더를 제거하려면 --http-header ""를 사용하십시오.
   --ipni                                                           IPNI에 거래를 알릴지 여부 (기본값: true)
   --url-template value, -u value                                   boost에서 CAR 파일을 가져올 PIECE_CID에 대한 URL 템플릿입니다. 예: http://127.0.0.1/piece/{PIECE_CID}.car

   거래 제안

   --duration value, -d value     이후 거래 종료까지의 기간(epoch 또는 기간 형식으로)입니다. 예: 1500000, 2400h
   --keep-unsealed                미완료된 사본을 유지할지 여부 (기본값: true)
   --price-per-deal value         거래 당 FIL 가격(기본값: 0)
   --price-per-gb value           GiB 당 FIL 가격(기본값: 0)
   --price-per-gb-epoch value     에포크 당 GiB 당 FIL 가격(기본값: 0)
   --start-delay value, -s value  거래 시작 지연(기간 형식 또는 epoch로)입니다. 예: 1000, 72h
   --verified                     검증된 거래로 제안할지 여부 (기본값: true)

   제한 사항

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      이 일정에 허용되는 피스 CID의 목록입니다. append(추가)만 됩니다.
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  허용할 피스 CID 목록이 포함된 파일 목록입니다. append(추가)만 됩니다.
   --max-pending-deal-number value, --pending-number value                                                            이 요청에 대한 전체 대기 거래 최대 개수입니다. 예: 100TiB (기본값: 0)
   --max-pending-deal-size value, --pending-size value                                                                이 요청에 대한 전체 대기 거래 최대 크기입니다. 예: 1000
   --total-deal-number value, --total-number value                                                                    이 요청에 대한 전체 거래 최대 개수입니다. 예: 1000 (기본값: 0)
   --total-deal-size value, --total-size value                                                                        이 요청에 대한 전체 거래 최대 크기입니다. 예: 100TiB

   스케줄링

   --schedule-cron value, --cron value           배치 거래를 보내기 위한 Cron 스케줄
   --schedule-deal-number value, --number value  트리거된 스케줄당 최대 거래 개수. 예: 30 (기본값: 0)
   --schedule-deal-size value, --size value      트리거된 스케줄당 최대 거래 크기. 예: 500GiB

   추적

   --notes value, -n value  추적 목적으로 요청과 함께 저장될 메모 또는 태그

```
{% endcode %}