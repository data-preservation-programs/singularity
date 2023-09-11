# 在生产环境中部署 Singularity

Singularity默认使用`sqlite3`作为它的数据库后端，因为它易于设置。但是，当切换到生产环境时，特别是如果您计划使用多个工作节点或者打算在负载平衡器后面提供检索服务，建议切换到更强大的数据库后端。您可以通过设置`$DATABASE_CONNECTION_STRING`环境变量来配置后端。

## 支持的数据库后端

- **PostgreSQL**：  
  连接字符串示例：  
  `postgres://user:pass@example.com:5432/dbname`

- **MySQL**：  
  连接字符串示例：  
  `mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true`

## 使用 Docker Compose 进行部署

如果您想快速部署带有PostgreSQL后端的Singularity，请考虑使用提供的Docker Compose模板：

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker-compose up
```
执行上述命令将设置一个PostgreSQL数据库，并启动必要的Singularity服务，包括API和数据集工作节点。