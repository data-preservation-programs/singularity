# Deploying Singularity to Production

Singularity uses `sqlite3` as its default database backend due to its ease of setup. However, when transitioning to a production environment, especially if you're planning to use multiple workers or intend to serve retrievals behind a load balancer, it's recommended to switch to a more robust database backend. You can configure the backend by setting the `$DATABASE_CONNECTION_STRING` environment variable.

## Supported Database Backends

- **PostgreSQL**:  
  Connection String Example:  
  `postgres://user:pass@example.com:5432/dbname`

- **MySQL**:  
  Connection String Example:  
  `mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true`

## Using Docker Compose for Deployment

If you'd like to quickly deploy Singularity along with a PostgreSQL backend, consider using the provided Docker Compose template:

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker-compose up
```
Executing the above commands will set up a PostgreSQL database and launch the necessary Singularity services, including the API and a dataset worker.
