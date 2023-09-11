# Citrix Sharefile

{% code fullWidth="true" %}
```
名前:
   singularity storage update sharefile - Citrix Sharefile

使用法:
   singularity storage update sharefile [コマンドオプション] <名前|ID>

説明:
   --upload-cutoff
      マルチパートアップロードに切り替えるためのカットオフポイント。

   --root-folder-id
      ルートフォルダのID。

      空白のままにすると「個人フォルダ」にアクセスします。ここでは、
      標準の値を使用するか、任意のフォルダID（長い16進数ID）を使用できます。

      例:
         | <未設定>    | パーソナルフォルダ（デフォルト）にアクセスします。
         | favorites  | お気に入りのフォルダにアクセスします。
         | allshared  | 共有されているすべてのフォルダにアクセスします。
         | connectors | 個別のコネクタにアクセスします。
         | top        | ホーム、お気に入り、共有フォルダとコネクタにアクセスします。

   --chunk-size
      アップロードのチャンクサイズ。

      2の累乗で、256k以上である必要があります。

      これを大きくするとパフォーマンスが向上しますが、注意してください。
      各チャンクは1つの転送ごとにメモリ上にバッファリングされます。

      これを減らすとメモリ使用量が減少しますが、パフォーマンスは低下します。

   --endpoint
      API呼び出しのエンドポイント。

      通常、oauthプロセスの一環として自動的に検出されますが、
      手動でhttps://XXX.sharefile.com などの値に設定することもできます。

   --encoding
      バックエンドのエンコーディング。

      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。


オプション:
   --help, -h              ヘルプを表示します
   --root-folder-id value  ルートフォルダのID [$ROOT_FOLDER_ID]

   Advanced

   --chunk-size value     アップロードのチャンクサイズ（デフォルト: "64Mi"）[$CHUNK_SIZE]
   --encoding value       バックエンドのエンコーディング（デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,Ctl,LeftSpace,LeftPeriod,RightSpace,RightPeriod,InvalidUtf8,Dot"）[$ENCODING]
   --endpoint value       API呼び出しのエンドポイント [$ENDPOINT]
   --upload-cutoff value  マルチパートアップロードに切り替えるためのカットオフポイント（デフォルト: "128Mi"）[$UPLOAD_CUTOFF]

```
{% endcode %}