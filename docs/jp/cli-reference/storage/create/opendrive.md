# OpenDrive

{% code fullWidth="true" %}
```
名前:
   singularity storage create opendrive - OpenDrive

使い方:
   singularity storage create opendrive [オプション] [引数...]

説明:
   --username
      ユーザー名。

   --password
      パスワード。

   --encoding
      バックエンドのエンコーディングです。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --chunk-size
      ファイルはこのサイズのチャンクでアップロードされます。
      
      これらのチャンクはメモリ上でバッファリングされるため、増やすとメモリ使用量が増加します。


オプション:
   --help, -h        ヘルプを表示します
   --password value  パスワード。[$PASSWORD]
   --username value  ユーザー名。[$USERNAME]

   高度なオプション

   --chunk-size value  ファイルはこのサイズのチャンクでアップロードされます。 (デフォルト: "10Mi") [$CHUNK_SIZE]
   --encoding value    バックエンドのエンコーディング。 (デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

   一般的なオプション

   --name value  ストレージの名前 (デフォルト: 自動生成)
   --path value  ストレージのパス

```
{% endcode %}