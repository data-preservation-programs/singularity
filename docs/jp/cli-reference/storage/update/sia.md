# Sia 分散クラウド

{% code fullWidth="true" %}
```
NAME:
   singularity storage update sia - Sia 分散クラウド

USAGE:
   singularity storage update sia [コマンドオプション] <名前|ID>

DESCRIPTION:
   --api-url
      SiaデーモンのAPI URL。例えば http://sia.daemon.host:9980のような形式です。
      
      注意：他のホストへAPIポートを開くためには、siadを--disable-api-securityオプションで実行する必要があります (非推奨)。
      Siaデーモンがlocalhost上で実行されている場合はデフォルトのままにしてください。

   --api-password
      SiaデーモンのAPIパスワード。
      
      ホームディレクトリ/.sia/ またはデーモンディレクトリにあるapipasswordファイルに記載されています。

   --user-agent
      Siad ユーザーエージェント
      
      セキュリティのために、Siaデーモンはデフォルトで'Sia-Agent'ユーザーエージェントを要求します。

   --encoding
      バックエンドのエンコーディング。

      詳細は[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


OPTIONS:
   --api-password value  SiaデーモンのAPIパスワード。 [$API_PASSWORD]
   --api-url value       SiaデーモンのAPI URL。例えば http://sia.daemon.host:9980 のような形式です。 (default: "http://127.0.0.1:9980") [$API_URL]
   --help, -h            ヘルプを表示する

   Advanced

   --encoding value    バックエンドのエンコーディング。 (default: "Slash,Question,Hash,Percent,Del,Ctl,InvalidUtf8,Dot") [$ENCODING]
   --user-agent value  Siad ユーザーエージェント (default: "Sia-Agent") [$USER_AGENT]

```
{% endcode %}