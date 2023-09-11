# Citrix Sharefile

{% code fullWidth="true" %}
```
名前:
   singularity storage create sharefile - Citrix Sharefile

使い方:
   singularity storage create sharefile [コマンドオプション] [引数...]

説明:
   --upload-cutoff
      マルチパートアップロードに切り替えるためのカットオフです。

   --root-folder-id
      ルートフォルダのIDです。

      空白のままにすると、"パーソナルフォルダ"にアクセスします。標準の値や任意のフォルダID（長い16進数のID）を使用できます。

      例:
         | <unset>    | パーソナルフォルダにアクセス（デフォルト）。
         | favorites  | お気に入りのフォルダにアクセス。
         | allshared  | 全ての共有フォルダにアクセス。
         | connectors | 個別のコネクタにアクセス。
         | top        | ホーム、お気に入り、共有フォルダおよびコネクタにアクセス。

   --chunk-size
      アップロードのチャンクサイズです。

      256キロバイト以上の2の累乗である必要があります。

      これを大きくするとパフォーマンスが向上しますが、注意点として各チャンクは1つの転送ごとにメモリにバッファリングされます。

      これを減らすとメモリの使用量が減りますが、パフォーマンスが低下します。

   --endpoint
      API呼び出しのエンドポイントです。

      通常、oauthプロセスの一部として自動的に検出されますが、次のように手動で設定することもできます: https://XXX.sharefile.com


   --encoding
      バックエンドのエンコーディングです。

      詳細については、[エンコーディングのセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h              ヘルプを表示
   --root-folder-id value  ルートフォルダのID [$ROOT_FOLDER_ID]

   上級者向け

   --chunk-size value     アップロードのチャンクサイズ (default: "64Mi") [$CHUNK_SIZE]
   --encoding value       バックエンドのエンコーディング (default: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot") [$ENCODING]
   --endpoint value       API呼び出しのエンドポイント [$ENDPOINT]
   --upload-cutoff value  マルチパートアップロードに切り替えるためのカットオフ (default: "128Mi") [$UPLOAD_CUTOFF]

   一般的なオプション

   --name value  ストレージの名前 (default: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}