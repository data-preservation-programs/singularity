# Dockerを使用してSingularityをインストールする

Dockerを使用すると、事前に構成済みのSingularityイメージを簡単に取得して実行できます。

## Dockerイメージの取得

事前にビルドされたDockerイメージを取得するには、次のコマンドを実行します:

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

## DockerイメージからSingularityを実行する
### デフォルトのSQLite3バックエンドを使用する場合

デフォルトでは、Singularityはデータベースバックエンドとして `sqlite3` を使用します。実行するには、コンテナ内のホームディレクトリにローカルパスをマウントする必要があります:

```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

### 別のデータベースバックエンドを使用する場合（例：Postgres）

Postgresのような別のデータベースバックエンドを選択する場合は、コンテナの実行時に`DATABASE_CONNECTION_STRING`環境変数を設定します:

```bash
docker run -e DATABASE_CONNECTION_STRING=your_connection_string_here ghcr.io/datapreservationprogram/singularity -h
```