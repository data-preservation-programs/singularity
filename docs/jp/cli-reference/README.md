# CLI リファレンス

{% code fullWidth="true" %}
```
NAME:
   singularity - FilecoinネットワークへのPBスケールのデータオンボードを行う大規模クライアントのためのツール

USAGE:
   singularity [グローバルオプション] コマンド [コマンドオプション] [引数...]

DESCRIPTION:
   データベースバックエンドのサポート:
     Singularityでは複数のデータベースバックエンドをサポートしています：sqlite3、postgres、mysql5.7以上
     データベース接続文字列を指定するには、'--database-connection-string' または $DATABASE_CONNECTION_STRING を使用します。
       postgresの例  - postgres://user:pass@example.com:5432/dbname
       mysqlの例     - mysql://user:pass@tcp(localhost:3306)/dbname?parseTime=true
       sqlite3の例   - sqlite:/絶対パス/データベース.db
                   または   - sqlite:相対パス/データベース.db

   ネットワークサポート:
     Singularityのデフォルト設定はMainnetです。他のネットワークに対しては、以下の環境変数を設定できます：
       Calibrationネットワークの場合：
         * LOTUS_API を https://api.calibration.node.glif.io/rpc/v1 に設定します
         * MARKET_DEAL_URL を https://marketdeals-calibration.s3.amazonaws.com/StateMarketDeals.json.zst に設定します
         * LOTUS_TEST を 1 に設定します
       その他のネットワークの場合：
         * LOTUS_API をネットワークのLotus APIエンドポイントに設定します
         * MARKET_DEAL_URL を空の文字列に設定します
         * network address が 'f' または 't' で始まるかに応じて、LOTUS_TEST を 0 または 1 に設定します
       同じデータベースインスタンスで異なるネットワーク間を切り替えることは推奨されません。

COMMANDS:
   version, v  バージョン情報を表示します
   help, h     コマンドの一覧または特定のコマンドのヘルプを表示します
   サービス:
     run  異なる Singularity コンポーネントを実行します
   オペレーション:
     admin    管理コマンド
     deal     レプリケーション / 取引管理
     wallet   ウォレット管理
     storage  ストレージシステムの接続の作成と管理
     prep     データセットの準備の作成と管理
   ユーティリティ:
     ez-prep      ローカルパスからデータセットを準備します
     download     メタデータAPIからCARファイルをダウンロードします
     extract-car  CARファイルのフォルダまたはファイルをローカルディレクトリに抽出します

GLOBAL OPTIONS:
   --database-connection-string value  データベースへの接続文字列 (デフォルト: sqlite:./singularity.db) [$DATABASE_CONNECTION_STRING]
   --help, -h                          ヘルプを表示します
   --json                              JSON出力を有効にします (デフォルト: false)
   --verbose                           詳細な出力を有効にします。結果にさらなる列や完全なエラートレースも表示されます (デフォルト: false)

   Lotus

   --lotus-api value    Lotus RPC APIエンドポイント (デフォルト: "https://api.node.glif.io/rpc/v1") [$LOTUS_API]
   --lotus-test         ランタイム環境がテストネットを使用しているかどうかを指定します (デフォルト: false) [$LOTUS_TEST]
   --lotus-token value  Lotus RPC APIトークン [$LOTUS_TOKEN]

```
{% endcode %}
