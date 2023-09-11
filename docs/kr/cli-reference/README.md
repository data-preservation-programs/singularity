# CLI 참조

{% code fullWidth="true" %}
```
이름:
   singularity - Filecoin 네트워크에 대규모 클라이언트 및 PB 규모 데이터 온보딩을 위한 도구

사용법:
   singularity [전역 옵션] 명령어 [명령어 옵션] [인수...]

설명:
   데이터베이스 백엔드 지원:
     Singularity는 다중 데이터베이스 백엔드를 지원합니다: sqlite3, postgres, mysql5.7+
     데이터베이스 연결 문자열을 지정하려면 '--database-connection-string' 또는 $DATABASE_CONNECTION_STRING을 사용하십시오.
       postgres에 대한 예 - postgres://user:pass@example.com:5432/dbname
       mysql에 대한 예 - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       sqlite3에 대한 예 - sqlite:/absolute/path/to/database.db
                   또는 - sqlite:relative/path/to/database.db

   네트워크 지원:
     Singularity의 기본 설정은 메인넷을 위한 것입니다. 다른 네트워크를 위해 아래 환경 값들을 설정할 수 있습니다:
       Calibration 네트워크를 위해:
         * LOTUS_API를 https://api.calibration.node.glif.io/rpc/v1로 설정
         * MARKET_DEAL_URL을 https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst로 설정
         * LOTUS_TEST를 1로 설정
       다른 모든 네트워크를 위해:
         * LOTUS_API를 네트워크의 Lotus API 엔드포인트로 설정
         * MARKET_DEAL_URL을 빈 문자열로 설정
         * LOTUS_TEST를 네트워크 주소가 'f' 또는 't'로 시작하는지에 따라 0 또는 1로 설정
       같은 데이터베이스 인스턴스에서 서로 다른 네트워크 사이를 전환하는 것은 권장되지 않습니다.

명령어:
   version, v  버전 정보 출력
   help, h     명령어 목록 또는 특정 명령어에 대한 도움말 표시
   데몬:
     run  다른 singularity 구성 요소 실행
   작업:
     admin    관리자 명령어
     deal     복제 / 거래 생성 관리
     wallet   지갑 관리
     storage  스토리지 시스템 연결 생성 및 관리
     prep     데이터 준비 생성 및 관리
   유틸리티:
     ez-prep      로컬 경로에서 데이터셋 준비
     download     메타데이터 API에서 CAR 파일 다운로드
     extract-car  CAR 파일 폴더 또는 파일을 로컬 디렉토리로 추출

전역 옵션:
   --database-connection-string value  데이터베이스에 대한 연결 문자열 (기본값: sqlite:./singularity.db) [$DATABASE_CONNECTION_STRING]
   --help, -h                          도움말 표시
   --json                              JSON 출력 사용 (기본값: false)
   --verbose                           자세한 출력 사용. 결과에 더 많은 열과 전체 오류 추적을 포함하여 출력됩니다 (기본값: false)

   Lotus

   --lotus-api value    Lotus RPC API 엔드포인트 (기본값: "https://api.node.glif.io/rpc/v1") [$LOTUS_API]
   --lotus-test         런타임 환경이 Testnet을 사용하는지 여부 (기본값: false) [$LOTUS_TEST]
   --lotus-token value  Lotus RPC API 토큰 [$LOTUS_TOKEN]

```
{% endcode %}