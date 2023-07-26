# Install from docker

To pull a pre-built docker image, use below command

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

By default, it will be using `sqlite3` as the backend. You will need to mount a local path to the home path inside the container, i.e.

```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

If you are using other database such as Postgres as the database backend, you will need to set the  `DATABASE_CONNECTION_STRING` environment variable, i.e.

```bash
docker run -e DATABASE_CONNECTION_STRING ghcr.io/datapreservationprogram/singularity -h
```
