# 다른 싱귤래리티 컴포넌트 실행하기

{% code fullWidth="true" %}
```
명령어:
   singularity run command [command options] [arguments...]

사용법:
   여러 가지 싱귤래리티 컴포넌트 실행

명령어:
   api               싱귤래리티 API 실행
   dataset-worker    데이터셋 준비 워커 시작하여 데이터셋 스캐닝 및 준비 작업 처리
   content-provider  검색 요청을 제공하는 콘텐츠 제공자 시작
   deal-tracker      모든 관련 지갑들에 대한 거래를 추적하는 딜 트래커 시작
   dealmaker         거래 만들기/추적 워커 시작하여 거래 처리
   spade-api         저장소 공급자 거래 제안 자체 서비스에 대한 Spade 호환 API 시작
   help, h           명령어 리스트 표시 또는 특정 명령어에 대한 도움말 표시

옵션:
   --help, -h  도움말 표시
```
{% endcode %}