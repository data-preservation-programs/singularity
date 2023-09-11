# 本番環境へのSingularityの展開

Singularityでは、セットアップの容易さから、デフォルトのデータベースバックエンドとして`sqlite3`を使用しています。ただし、特に複数のワーカーを使用する予定や、負荷分散のために取得を提供する場合など、本番環境への移行時には、より堅牢なデータベースバックエンドに切り替えることをお勧めします。バックエンドは、`$DATABASE_CONNECTION_STRING`環境変数を設定することで構成できます。

## サポートされているデータベースバックエンド

- **PostgreSQL**:  
  接続文字列の例:  
  `postgres://user:pass@example.com:5432/dbname`

- **MySQL**:  
  接続文字列の例:  
  `mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true`

## デプロイにDocker Composeを使用する

SingularityをPostgreSQLバックエンドとともに素早く展開したい場合は、提供されているDocker Composeテンプレートを使用することを検討してください。

```bash
wget https://raw.githubusercontent.com/data-preservation-programs/singularity/main/docker-compose.yml
docker-compose up
```

上記のコマンドを実行すると、PostgreSQLデータベースがセットアップされ、必要なSingularityサービス（APIおよびデータセットワーカー）が起動されます。