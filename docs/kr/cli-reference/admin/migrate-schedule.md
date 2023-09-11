# 오래된 싱귤래리 모든 사용자의 파일들을 이관한다.

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-schedule - 오래된 싱귤래리 모든 사용자의 파일들을 이관한다.

USAGE:
   singularity admin migrate-schedule [command options] [arguments...]

DESCRIPTION:
   싱귤래리 V1에서 V2로 스케줄을 이관한다. 주의사항:
     1. 데이터셋 이관을 먼저 완료해야 한다.
     2. 모든 새로운 스케줄은 '일시 정지' 상태로 생성된다.
     3. 거래 상태는 자동으로 거래 추적기로 채워질 것이므로 이관되지 않는다.
     4. --output-csv는 더 이상 지원되지 않는다. 앞으로 새로운 도구를 제공할 것이다.
     5. 레플리카 개수는 스케줄의 일부로 지원되지 않는다. 앞으로 이것은 설정 가능한 정책으로 만들 것이다.
     6. --force는 더 이상 지원되지 않는다. 앞으로 정책 제한을 모두 무시할 수 있는 유사한 지원을 추가할 수도 있다.
     7. --offline는 더 이상 지원되지 않는다. URL 템플릿이 설정된 경우 레거시 마켓에는 항상 오프라인 거래로 설정되고 부스트 마켓에는 온라인 거래로 설정될 것이다.

OPTIONS:
   --mongo-connection-string value  MongoDB 연결 문자열 (기본값: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --help, -h                       도움말 표시
```
{% endcode %}