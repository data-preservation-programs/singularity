# 프로덕션에 배포하기

기본적으로, Singularity는 설정이 필요하지 않기 때문에 `sqlite3`을 데이터베이스 백엔드로 사용합니다. 하지만 실제 프로덕션 환경에서는 또는 여러 작업자 또는 로드 밸런서 뒤에서 검색을 제공하기 위해 실제 데이터베이스 백엔드를 사용하고자 할 것입니다. 이를 위해서 `$DATABASE_CONNECTION_STRING`을 사용할 수 있습니다.

* Postgres 예시: `postgres://user:pass@example.com:5432/dbname`
* MySQL 예시: `mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true`

아래의 Docker Compose 템플릿을 사용하여 시작점으로 사용할 수도 있습니다. 이 템플릿은 포스트그레스 데이터베이스 서비스를 시작하고 모든 관련 Singularity 서비스를 실행합니다.

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker compose up
```

이를 통해 포스트그레스 데이터베이스를 설정하고, Singularity API와 단일 데이터셋 작업자를 시작할 수 있습니다.