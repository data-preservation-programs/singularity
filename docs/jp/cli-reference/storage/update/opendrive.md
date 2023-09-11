# OpenDrive

{% code fullWidth="true" %}
```
名前:
   singularity storage update opendrive - OpenDrive

使用法:
   singularity storage update opendrive [コマンドオプション] <名前|ID>

説明:
   --username
      ユーザー名。

   --password
      パスワード。

   --encoding
      バックエンドのエンコーディング。
      
      詳細については、[概要のエンコーディングセクション](/overview/#encoding)を参照してください。

   --chunk-size
      ファイルはこのサイズでチャンクにアップロードされます。
      
      ただし、これらのチャンクはメモリにバッファされるため、増やすとメモリ使用量が増加します。


オプション:
   --help, -h        ヘルプを表示
   --password value  パスワード。 [$PASSWORD]
   --username value  ユーザー名。 [$USERNAME]

   高度なオプション

   --chunk-size value  ファイルはこのサイズでチャンクにアップロードされます。(デフォルト: "10Mi") [$CHUNK_SIZE]
   --encoding value    バックエンドのエンコーディング。(デフォルト: "Slash,LtGt,DoubleQuote,Colon,Question,Asterisk,Pipe,BackSlash,LeftSpace,LeftCrLfHtVt,RightSpace,RightCrLfHtVt,InvalidUtf8,Dot") [$ENCODING]

```