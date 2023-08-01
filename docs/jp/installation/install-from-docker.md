# Dockerからのインストール

プリビルドされたDockerイメージを取得するには、以下のコマンドを使用してください。

```bash
docker pull ghcr.io/data-preservation-programs/singularity:main
```

デフォルトでは、バックエンドとして`sqlite3`が使用されています。コンテナ内のホームパスにローカルパスをマウントする必要があります。

```bash
docker run -v $HOME:/root ghcr.io/datapreservationprogram/singularity -h
```

データベースのバックエンドとしてPostgresなどの他のデータベースを使用する場合は、`DATABASE_CONNECTION_STRING`環境変数を設定する必要があります。

```bash
docker run -e DATABASE_CONNECTION_STRING ghcr.io/datapreservationprogram/singularity -h
```