# 도커로 설치하기

미리 빌드된 도커 이미지를 가져오려면 아래 명령을 사용하세요.

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

기본적으로 `sqlite3`가 백엔드로 사용됩니다. 컨테이너 내부의 홈 경로에 로컬 경로를 마운트해야 합니다. 즉,

```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

데이터베이스 백엔드로 다른 데이터베이스(예: PostgreSQL)를 사용하는 경우, `DATABASE_CONNECTION_STRING` 환경 변수를 설정해야 합니다. 즉,

```bash
docker run -e DATABASE_CONNECTION_STRING ghcr.io/datapreservationprogram/singularity -h
```