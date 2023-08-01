# 부스트 또는 레거시 마켓에 수동 거래 제안 보내기

{% code fullWidth="true" %}
```
NAME:
   singularity deal send-manual - 부스트 또는 레거시 마켓에 수동 거래 제안 보내기

사용법:
   singularity deal send-manual [command options] CLIENT_ADDRESS PROVIDER_ID PIECE_CID PIECE_SIZE

OPTIONS:
   --help, -h       도움말 표시
   --timeout value  거래 제안을 위한 타임아웃 설정 (기본값: 1분)

   부스트 전용

   --file-size value                            CAR 파일을 가져오기 위한 파일 크기 (기본값: 0)
   --http-header value [ --http-header value ]  요청과 함께 전달될 HTTP 헤더 (예: key=value)
   --ipni                                       거래를 IPNI에 알리는지 여부 (기본값: true)
   --url-template value                         CAR 파일을 가져오기 위해 PIECE_CID 자리 표시자가 포함된 URL 템플릿 (예: http://127.0.0.1/piece/{PIECE_CID}.car)

   거래 제안

   --duration value, -d value     거래 기간(에포크 또는 기간 형식), 예: 1500000, 2400시간 (기본값: 12840시간[535일])
   --keep-unsealed                미인증된 사본을 유지할지 여부 (기본값: true)
   --price-per-deal value         거래당 FIL 단가 (기본값: 0)
   --price-per-gb value           1GB당 FIL 단가 (기본값: 0)
   --price-per-gb-epoch value     1GB당 FIL 단가/에포크 (기본값: 0)
   --root-cid value               거래 제안의 일부로 필요한 Root CID, 비어있는 경우 빈 CID로 설정됨 (기본값: 빈 CID)
   --start-delay value, -s value  거래 시작 지연 시간(에포크 또는 기간 형식), 예: 1000, 72시간 (기본값: 72시간[3일])
   --verified                     검증된 거래로 거래를 제안할지 여부 (기본값: true)

```
{% endcode %}