# Deploy to production

By default, singularity is using `sqlite3` as the database backend since it needs zero setup. For production usage, or to use it with multiple workers or serving retrieval behind load balancer, you would like to use a real database backend using `$DATABASE_CONNECTION_STRING`

* Postgres example: `postgres://user:pass@example.com:5432/dbname`
* Mysql example: `mysql://user:pass@example.com:5432/dbname`
