# CLI 참조

{% code fullWidth="true" %}
```
이름:
   singularity - 파일코인 네트워크에 대규모 클라이언트 및 PB 스케일 데이터 온보딩을 위한 도구

사용법:
   singularity [일반 옵션] 명령어 [명령어 옵션] [인수...]

설명:
   데이터베이스 백엔드 지원:
     Singularity는 다양한 데이터베이스 백엔드를 지원합니다: sqlite3, postgres, mysql5.7+
     데이터베이스 연결 문자열을 지정하려면 '--database-connection-string' 또는 $DATABASE_CONNECTION_STRING을 사용하십시오.
       예시(포스트그레스)     - postgres://user:pass@example.com:5432/dbname
       예시(마이스톤)        - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       예시(스퀘이트3)       - sqlite:/absolute/path/to/database.db
                   또는        - sqlite:relative/path/to/database.db

   네트워크 지원:
     Singularity의 기본 설정은 Mainnet입니다. 다른 네트워크에 대해 아래 환경값을 설정할 수 있습니다:
       Calibration 네트워크의 경우:
         * LOTUS_API를 https://api.calibration.node.glif.io/rpc/v1으로 설정
         * MARKET_DEAL_URL을 https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst로 설정
       기타 모든 네트워크의 경우:
         * LOTUS_API를 해당 네트워크의 Lotus API 엔드포인트로 설정
         * MARKET_DEAL_URL을 빈 문자열로 설정
       동일한 데이터베이스 인스턴스에서 다른 네트워크로 전환하는 것은 권장되지 않습니다.

명령어:
   version, v  버전 정보 출력
   help, h     명령어 목록 또는 특정 명령어의 도움말 보기
   Daemons:
     run  다른 Singularity 구성 요소 실행
   쉬운 명령어:
     ez-prep  로컬 경로에서 데이터셋 준비
   작업:
     admin       관리자 명령어
     deal        복제/딜 진행 관리
     dataset     데이터셋 관리
     datasource  데이터 소스 관리
     wallet      지갑 관리
   도구:
     tool  개발 및 디버깅에 사용되는 도구
   유틸리티:
     download  메타데이터 API에서 CAR 파일 다운로드

일반 옵션:
   --database-connection-string value  데이터베이스 연결 문자열 (기본값: sqlite:./singularity.db) [$DATABASE_CONNECTION_STRING]
   --help, -h                          도움말 표시
   --json                              JSON 출력 사용 (기본값: false)

   Lotus

   --lotus-api value    Lotus RPC API 엔드포인트 (기본값: "https://api.node.glif.io/rpc/v1") [$LOTUS_API]
   --lotus-token value  Lotus RPC API 토큰 [$LOTUS_TOKEN]

```
{% endcode %}