# 이전 싱귤래리티 MongoDB로부터 데이터셋 마이그레이션하기

{% code fullWidth="true" %}
```
NAME:
   singularity admin migrate-dataset - 이전 싱귤래리티 MongoDB로부터 데이터셋 마이그레이션하기

사용법:
   singularity admin migrate-dataset [command options] [arguments...]

설명:
   싱귤래리티 V1에서 V2로 데이터셋을 마이그레이션합니다. 다음과 같은 단계를 포함합니다.
     1. 소스 스토리지와 출력 스토리지를 생성하고 V2의 데이터준비(DataPrep)에 연결합니다.
     2. 새로운 데이터셋에 모든 폴더 구조와 파일을 생성합니다.
   주의사항:
     1. 생성된 준비는 새로운 데이터셋 워커와 호환되지 않습니다.
        따라서 데이터 준비를 재개하거나 마이그레이션된 데이터셋에 새로운 파일을 추가하지 마십시오.
        이러한 작업은 문제 없이 거래를 진행하거나 데이터셋을 찾아볼 수 있습니다.
     2. 폴더 CID는 복잡성으로 인해 생성되거나 마이그레이션되지 않습니다.

옵션:
   --mongo-connection-string value  MongoDB 연결 문자열 (기본값: "mongodb://localhost:27017") [$MONGO_CONNECTION_STRING]
   --skip-files                     파일과 폴더에 대한 세부 정보 마이그레이션을 건너뜁니다. 이렇게 하면 마이그레이션 속도가 훨씬 빨라집니다.
                                    거래만 수행하고자 할 경우 유용합니다. (기본값: false)
   --help, -h                       도움말 표시
```
{% endcode %}