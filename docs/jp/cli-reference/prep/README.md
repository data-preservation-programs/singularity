# データセットの準備の作成と管理

{% code fullWidth="true" %}
```
NAME:
   singularity prep - データセットの準備の作成と管理

使用法:
   singularity prep コマンド [コマンドオプション] [引数...]

コマンド:
   create                新しい準備を作成する
   list                  全ての準備を一覧表示する
   status                準備のジョブステータスを取得する
   attach-source         ソースストレージを準備にアタッチする
   attach-output         出力ストレージを準備にアタッチする
   detach-output         出力ストレージを準備からデタッチする
   start-scan            ソースストレージのスキャンを開始する
   pause-scan            スキャンジョブを一時停止する
   start-pack            全てのパックジョブまたは特定のジョブを開始／再開する
   pause-pack            全てのパックジョブまたは特定のジョブを一時停止する
   start-daggen          全てのフォルダ構造のスナップショットを作成する DAG 生成を開始する
   pause-daggen          DAG 生成ジョブを一時停止する
   list-pieces           準備に生成された全てのピースを一覧表示する
   add-piece             準備に手動でピース情報を追加する。これは外部ツールで準備されたピースに便利です。
   explore               パスによる準備されたソースのエクスプローラ
   attach-wallet         ウォレットを準備にアタッチする
   list-wallets          準備にアタッチされたウォレットを一覧表示する
   detach-wallet         ウォレットを準備からデタッチする
   help, h               コマンドのリストを表示するか、コマンドのヘルプを表示する

オプション:
   --help, -h   ヘルプを表示する
```
{% endcode %}