# 本番環境にデプロイする

デフォルトでは、Singularityは`sqlite3`をデータベースバックエンドとして使用しています。これはセットアップが不要であるためです。本番環境では、または複数のワーカーで使用する場合や、ロードバランサーの背後でのデータ取得を提供する場合は、`$DATABASE_CONNECTION_STRING`を使用して実際のデータベースバックエンドを使用することをお勧めします。

* Postgresの例: `postgres://user:pass@example.com:5432/dbname`
* Mysqlの例: `mysql://user:pass@tcp(localhost:3306)/dbname?charset=ascii&parseTime=true`

以下のDocker Composeテンプレートも、スタート地点として使用することができます。これは、Postgresデータベースサービスを起動し、関連するすべてのSingularityサービスを実行します。

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker compose up
```

これにより、Postgresデータベースがセットアップされ、Singularity APIと単一のデータセットワーカーが起動します。