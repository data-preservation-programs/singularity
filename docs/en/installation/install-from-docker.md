# Installing Singularity via Docker

Utilizing Docker, you can effortlessly pull and run a pre-configured Singularity image.

## Pulling the Docker Image

To acquire the pre-built Docker image, execute the following command:

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

## Running Singularity from the Docker Image
### Using Default SQLite3 Backend

By default, Singularity uses `sqlite3` as its database backend. To run it, you should mount a local path to the home directory within the container:

```bash
docker run -v $HOME:/root ghcr.io/data-preservation-programs/singularity:main -h
```

### Using an Alternate Database Backend (e.g., Postgres)

If you opt for another database backend like Postgres, set the `DATABASE_CONNECTION_STRING` environment variable during container execution:
```bash
docker run -e DATABASE_CONNECTION_STRING=your_connection_string_here ghcr.io/data-preservation-programs/singularity:main -h
```
