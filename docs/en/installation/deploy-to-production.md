# Deploy to production

By default, singularity is using `sqlite3` as the database backend since it needs zero setup. For production usage, or to use it with multiple workers or serving retrieval behind load balancer, you would like to use a real database backend using `$DATABASE_CONNECTION_STRING`

* Postgres example: `postgres://user:pass@example.com:5432/dbname`
* Mysql example: `mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true`

You can also use below docker compose template as a start point. It starts up a postgres database service and runs all relevant singularity services.

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker compose up
```

This will setup a postgres database, starts singularity API and a single dataset worker.
