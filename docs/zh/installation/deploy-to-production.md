# 部署到生产环境

默认情况下，Singularity 使用 `sqlite3` 作为数据库后端，因为它不需要任何设置。但如果要用于生产环境，或者与多个工作进程一起使用，或者在负载均衡器后面提供检索服务，则需要使用真正的数据库后端，使用 `$DATABASE_CONNECTION_STRING`

* Postgres 示例：`postgres://user:pass@example.com:5432/dbname`
* MySQL 示例：`mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true`

您还可以使用下面的 Docker Compose 模板作为起点。它会启动一个 Postgres 数据库服务并运行所有相关的 Singularity 服务。

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker compose up
```

这将设置一个 Postgres 数据库，并启动 Singularity API 和一个单独的数据集工作进程。