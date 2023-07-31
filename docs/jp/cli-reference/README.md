# CLIリファレンス

{% code fullWidth="true" %}
```
NAME:
   singularity - ファイルコインネットワークへのPBスケールデータのオンボーディングを行う大規模クライアント向けのツール

USAGE:
   singularity [global options] command [command options] [arguments...]

DESCRIPTION:
   データベースバックエンドのサポート：
     Singularityは複数のデータベースバックエンドをサポートしています：sqlite3、postgres、mysql5.7+
     データベース接続文字列を指定するには、'--database-connection-string'または$DATABASE_CONNECTION_STRINGを使用します。
       Postgresの例 - postgres://user:pass@example.com:5432/dbname
       MySQLの例 - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       SQLite3の例 - sqlite:/absolute/path/to/database.db
                   または  - sqlite:relative/path/to/database.db

   ネットワークのサポート：
     Singularityのデフォルト設定はMainnet用です。他のネットワークを使用する場合は、以下の環境変数を設定できます：
       Calibrationネットワークの場合：
         * LOTUS_APIをhttps://api.calibration.node.glif.io/rpc/v1に設定する
         * MARKET_DEAL_URLをhttps://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zstに設定する
       その他のすべてのネットワークの場合：
         * LOTUS_APIを自分のネットワークのロータスAPIエンドポイントに設定する
         * MARKET_DEAL_URLを空の文字列に設定する
       同じデータベースインスタンスで異なるネットワークを切り替えることはお勧めしません。

COMMANDS:
   version, v  バージョン情報を表示する
   help, h     コマンドのリストを表示するか、特定のコマンドのヘルプを表示する
   守護プログラム：
     run  異なるSingularityコンポーネントを実行する
   簡単なコマンド：
     ez-prep  ローカルパスからデータセットを準備する
   オペレーション：
     admin       管理コマンド
     deal        レプリケーション/ディール管理
     dataset     データセット管理
     datasource  データソース管理
     wallet      ウォレット管理
   ツール：
     tool  開発とデバッグに使用されるツール
   ユーティリティ：
     download  メタデータAPIからCARファイルをダウンロードする

GLOBAL OPTIONS:
   --database-connection-string value  データベースへの接続文字列（デフォルト：sqlite:./singularity.db）[$DATABASE_CONNECTION_STRING]
   --help, -h                          ヘルプを表示する
   --json                              JSON出力を有効にする（デフォルト：無効）

   Lotus

   --lotus-api value    Lotus RPC APIエンドポイント（デフォルト："https://api.node.glif.io/rpc/v1"）[$LOTUS_API]
   --lotus-token value  Lotus RPC APIトークン[$LOTUS_TOKEN]

```
{% endcode %}