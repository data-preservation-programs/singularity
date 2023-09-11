# 스토리지 제공업체에 거래를 보내기위한 일정 생성

{% code fullWidth="true" %}
```
명령어:
   singularity deal schedule create - 스토리지 제공업체에 거래를 보내기위한 일정 생성

사용법:
   singularity deal schedule create [command options] [arguments...]

설명:
   CRON 패턴 '--schedule-cron': CRON 패턴은 기술자(dexcriptor) 또는 선택적으로 두 번째 필드를 사용한 표준 CRON 패턴이 될 수 있습니다.
     표준 CRON:
       ┌───────────── 분 (0 - 59)
       │ ┌───────────── 시간 (0 - 23)
       │ │ ┌───────────── 월의 일 (1 - 31)
       │ │ │ ┌───────────── 월 (1 - 12)
       │ │ │ │ ┌───────────── 주의 일 (0 - 6) (일요일 - 토요일)
       │ │ │ │ │                                   
       │ │ │ │ │
       │ │ │ │ │
       * * * * *

     선택적으로 두 번째 필드:
       ┌─────────────  초 (0 - 59)
       │ ┌─────────────  분 (0 - 59)
       │ │ ┌─────────────  시간 (0 - 23)
       │ │ │ ┌─────────────  월의 일 (1 - 31)
       │ │ │ │ ┌─────────────  월 (1 - 12)
       │ │ │ │ │ ┌─────────────  주의 일 (0 - 6) (일요일 - 토요일)
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
   --help, -h             도움말 표시
   --preparation value    준비 ID 또는 이름
   --provider value       거래를 보낼 스토리지 제공업체 ID

   Boost Only

   --http-header value, -H value [ --http-header value, -H value ]  요청과 함께 전달할 HTTP 헤더 (예: key=value)
   --ipni                                                           거래를 IPNI에 알리는지 여부(default: true)
   --url-template value, -u value                                   boost가 CAR 파일을 가져올때 PIECE_CID 플레이스홀더와 함께 사용하는 URL 템플릿, 예: http://127.0.0.1/piece/{PIECE_CID}.car

   거래 제안

   --duration value, -d value     기간을 에포크 또는 기간 형식으로 지정, 예: 1500000, 2400h (default: 12840h[535일])
   --keep-unsealed                미완성된 복사본을 유지할지 여부(default: true)
   --price-per-deal value         거래 당 FIL 가격(default: 0)
   --price-per-gb value           GiB 당 FIL 가격(default: 0)
   --price-per-gb-epoch value     GiB당 FIL/에포크 가격(default: 0)
   --start-delay value, -s value  거래 시작 지연 시간을 에포크 또는 기간 형식으로 지정, 예: 1000, 72h (default: 72h[3일])
   --verified                     검증된 거래로 거래를 제안할지 여부(default: true)

   제한사항

   --allowed-piece-cid value, --piece-cid value [ --allowed-piece-cid value, --piece-cid value ]                      이 일정에서 허용된 piece CID 목록(Default: Any)
   --allowed-piece-cid-file value, --piece-cid-file value [ --allowed-piece-cid-file value, --piece-cid-file value ]  허용된 piece CID 목록이 포함 된 파일 목록
   --max-pending-deal-number value, --pending-number value                                                            요청 전체에 대한 최대 대기중인 거래 번호(Default: 무제한)
   --max-pending-deal-size value, --pending-size value                                                                요청 전체에 대한 최대 대기중인 거래 크기(Default: 무제한)
   --total-deal-number value, --total-number value                                                                    요청에 대한 최대 거래 번호(Default: 무제한)
   --total-deal-size value, --total-size value                                                                        요청에 대한 최대 거래 크기(Default: 무제한)

   스케줄링

   --schedule-cron value, --cron value           배치 거래 전송하는 cron 일정 설정(Default: disabled)
   --schedule-deal-number value, --number value  트리거된 일정 당 최대 거래 번호, 예: 30 (default: 무제한)
   --schedule-deal-size value, --size value      트리거된 일정 당 최대 거래 크기, 예: 500GiB (default: 무제한)

   추적

   --notes value, -n value  요청과 함께 저장되는 추적 용도의 메모나 태그

```
{% endcode %}