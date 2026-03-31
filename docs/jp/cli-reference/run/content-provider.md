# 検索リクエストを提供するコンテンツプロバイダを開始する

{% code fullWidth="true" %}
```
NAME:
   singularity run content-provider - 検索リクエストを提供するコンテンツプロバイダを開始する

USAGE:
   singularity run content-provider [コマンドオプション] [引数...]

OPTIONS:
   --help, -h  ヘルプを表示する

   HTTP IPFS Gateway

   --enable-http-ipfs  Enable trustless IPFS gateway on /ipfs/ (default: true)

   HTTPリトリーバル

   --enable-http      HTTPリトリーバルを有効にする (デフォルト: true)
   --http-bind value  HTTPサーバをバインドするアドレス (デフォルト: "127.0.0.1:7777")

```
{% endcode %}