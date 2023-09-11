# 운영 환경으로의 Singularity 배포하기

Singularity는 설정이 간편한 `sqlite3`를 기본 데이터베이스 백엔드로 사용합니다. 그러나 복수의 워커를 사용하거나 로드 밸런서 뒤에서 검색을 서비스하려는 경우와 같이 운영 환경으로 전환할 때는 더 견고한 데이터베이스 백엔드로 전환하는 것이 권장됩니다. 데이터베이스 백엔드를 구성하기 위해 `$DATABASE_CONNECTION_STRING` 환경 변수를 설정할 수 있습니다.

## 지원하는 데이터베이스 백엔드

- **PostgreSQL**:  
  연결 문자열 예시:  
  `postgres://user:pass@example.com:5432/dbname`

- **MySQL**:  
  연결 문자열 예시:  
  `mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true`

## 배포를 위해 Docker Compose 사용하기

만약 PostgreSQL 백엔드와 함께 Singularity를 빠르게 배포하고자 한다면, 제공되는 Docker Compose 템플릿을 사용해보세요:

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker-compose up
```
위 명령을 실행하면, PostgreSQL 데이터베이스를 설정하고 필요한 Singularity 서비스, API, 그리고 데이터셋 워커를 시작합니다.