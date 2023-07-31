# Sia 分散クラウド

{% code fullWidth="true" %}
```
NAME:
   singularity datasource add sia - Sia 分散クラウド

USAGE:
   singularity datasource add sia [command options] <dataset_name> <source_path>

DESCRIPTION:
   --sia-api-password
      Sia Daemon API パスワード。
      
      "HOME/.sia/" やデーモンディレクトリにある "apipassword" ファイルに見つけることができます。
      
   --sia-api-url
      Sia デーモン API URL、例えば http://sia.daemon.host:9980。
      
      注意: 他のホストに対してAPIポートを開くためには、siad が --disable-api-security で実行される必要があります（推奨されません）。
      Sia デーモンが localhost 上で実行されている場合は、デフォルトを使用してください。

   --sia-encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --sia-user-agent
      Siad ユーザーエージェント。
      
      セキュリティのため、Sia デーモンはデフォルトで 'Sia-Agent' ユーザーエージェントを要求します。


OPTIONS:
   --help, -h  ヘルプを表示します

   データの準備オプション

   --delete-after-export    [危険] データセットを CAR ファイルにエクスポートした後、データセットのファイルを削除します。  (default: false)
   --rescan-interval value  最後の成功したスキャンからの経過時間がこの間隔を超えた場合、自動的にソースディレクトリを再スキャンします (default: disabled)
   --scanning-state value   初期のスキャン状態を設定します (default: ready)

   Sia のオプション

   --sia-api-password value  Sia Daemon API パスワード。 [$SIA_API_PASSWORD]
   --sia-api-url value       Sia デーモン API URL、例えば http://sia.daemon.host:9980。 (default: "http://127.0.0.1:9980") [$SIA_API_URL]
   --sia-encoding value      バックエンドのエンコーディング。 (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$SIA_ENCODING]
   --sia-user-agent value    Siad ユーザーエージェント (default: "Sia-Agent") [$SIA_USER_AGENT]

```
{% endcode %}