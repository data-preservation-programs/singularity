# 부스트 또는 레거시 마켓으로 수동 거래 제안 보내기

{% code fullWidth="true" %}
```
이름:
   싱귤래리티 딜 수동 전송 - 부스트 또는 레거시 마켓으로 수동 거래 제안 보내기

사용법:
   싱귤래리티 딜 수동 전송 [명령 옵션] <클라이언트> <제공자> <piece_cid> <piece_size>

설명:
   부스트 또는 레거시 마켓으로 수동 거래 제안을 보냅니다.
     예: singularity deal send-manual f01234 f05678 bagaxxxx 32GiB
   참고:
     * 거래 제안은 데이터베이스에 저장되지 않으며, 거래 추적기가 실행 중인 경우 추적됩니다.
     * 클라이언트 주소는 'singularity wallet import'를 사용하여 월렛에 가져와야 합니다.
     * 자체 lotus 노드에 대한 LOTUS_API 및 LOTUS_TOKEN을 설정하여 GLIF API를 사용하여 빠른 주소 검증을 수행할 수 있습니다.

옵션:
   --help, -h       도움말 표시
   --timeout value  거래 제안에 대한 타임아웃 (기본값: 1m)

   부스트 전용

   --file-size value                            부스트가 CAR 파일을 가져오기 위한 파일 크기(바이트) (기본값: 0)
   --http-header value [ --http-header value ]  요청과 함께 전달할 HTTP 헤더(예: key=value)
   --ipni                                       거래를 IPNI에 발표할지 여부 (기본값: true)
   --url-template value                         부스트가 CAR 파일을 가져오기 위한 PIECE_CID 자리 표시자가 포함된 URL 템플릿, 예: http://127.0.0.1/piece/{PIECE_CID}.car

   거래 제안

   --client value                 거래를 보낼 클라이언트 주소
   --duration value, -d value     거래 기간을 epoch 형식이나 기간 형식으로 지정합니다 (기본값: 12840h[535일])
   --keep-unsealed                미완료된 사본을 유지할지 여부 (기본값: true)
   --piece-cid value              거래의 Piece CID
   --piece-size value             거래의 Piece 크기 (기본값: "32GiB")
   --price-per-deal value         거래당 FIL 가격 (기본값: 0)
   --price-per-gb value           GiB당 FIL 가격 (기본값: 0)
   --price-per-gb-epoch value     epoch 당 GiB당 FIL 가격 (기본값: 0)
   --provider value               거래를 보낼 저장 공급자 ID
   --root-cid value               거래 제안의 일부로 필요한 Root CID, 비어있는 경우 빈 CID로 설정됩니다 (기본값: 빈 CID)
   --start-delay value, -s value  거래 시작 지연 시간을 epoch 형식이나 기간 형식으로 지정합니다 (기본값: 72h[3일])
   --verified                     검증된 거래로 거래 제안을 제안할지 여부 (기본값: true)

```
{% endcode %}