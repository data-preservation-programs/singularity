# 異なるSingularityのコンポーネントを実行する

{% code fullWidth="true" %}
```
NAME:
   singularity run - 異なるSingularityのコンポーネントを実行する

USAGE:
   singularity run command [command options] [arguments...]

COMMANDS:
   api               SingularityのAPIを実行する
   dataset-worker    データセットの準備のためのワーカーを開始し、データセットのスキャンと準備のタスクを処理する
   content-provider  リクエストの提供を行うコンテンツプロバイダーを開始する
   deal-tracker      関連するウォレットの取引を追跡するディールトラッカーを開始する
   deal-pusher       取引スケジュールを監視し、ディールをストレージプロバイダーにプッシュするディールプッシャーを開始する
   help, h           コマンドのリストの表示または特定のコマンドのヘルプを表示する

OPTIONS:
   --help, -h  ヘルプの表示
```
{% endcode %}