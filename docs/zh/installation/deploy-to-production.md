# 部署到生产环境

默认情况下，Singularity使用`sqlite3`作为数据库后端，因为它不需要任何设置。对于生产环境的使用或使用多个工作进程或在负载均衡器后方提供检索，则需要使用`$DATABASE_CONNECTION_STRING`使用真实数据库后端。

* Postgres示例：'postgres：// user：pass @ example.com：5432 / dbname'
* Mysql示例：'mysql：// user：pass @ tcp（localhost：3306）/ dbname？charset = ascii＆parseTime = true'

您还可以使用以下Docker Compose模板作为起点。它启动了一个postgres数据库服务并运行所有相关的singularity服务。

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker compose up
```

这将设置一个postgres数据库，启动singularity API和单个数据集工作进程。