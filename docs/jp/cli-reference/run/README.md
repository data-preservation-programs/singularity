# 異なる Singularity コンポーネントを実行する

{% code fullWidth="true" %}
```
NAME:
   singularity run - 異なる Singularity コンポーネントを実行する

USAGE:
   singularity run コマンド [コマンドオプション] [引数...]

COMMANDS:
   api               Singularity API を実行する
   dataset-worker    データセットのスキャンと準備タスクを処理するためのデータセット準備ワーカーを起動します
   content-provider  検索リクエストを提供するコンテンツプロバイダーを起動します
   deal-tracker      関連するウォレットの取引を追跡するディールトラッカーを起動します
   dealmaker         取引の作成と追跡を処理するディール作成/追跡ワーカーを起動します
   spade-api         ストレージプロバイダーのディール提案のための Spade 互換の API を起動します
   help, h           コマンドのリストを表示するか、特定のコマンドのヘルプを表示します

OPTIONS:
   --help, -h  ヘルプを表示する
```
{% endcode %}