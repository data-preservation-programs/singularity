# 从 Docker 安装

要拉取预先构建的 Docker 镜像，请使用以下命令

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

默认情况下，它将使用 `sqlite3` 作为后端。您需要将本地路径挂载到容器内的主目录中，即

```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

如果您使用其他数据库，例如 PostgreSQL 作为数据库后端，您需要设置 `DATABASE_CONNECTION_STRING` 环境变量，即

```bash
docker run -e DATABASE_CONNECTION_STRING ghcr.io/datapreservationprogram/singularity -h
```