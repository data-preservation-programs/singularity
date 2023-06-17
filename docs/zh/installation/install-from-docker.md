# Docker 安装

要拉取预先构建的Docker映像，请使用以下命令

```bash
docker pull datapreservationprogram/singularity:latest
```

默认情况下，它将使用 `sqlite3` 作为后端。您需要将本地路径挂载到容器内的主目录下，例如

```bash
docker run -v $HOME:/root datapreservationprogram/singularity -h
```

如果您使用其他数据库（例如Postgres）作为数据库后端，则需要设置` DATABASE_CONNECTION_STRING`环境变量，例如

```bash
docker run -e DATABASE_CONNECTION_STRING datapreservationprogram/singularity -h
```